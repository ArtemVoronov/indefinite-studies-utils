package mongo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/api"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	ConnectTimeout time.Duration
	QueryTimeout   time.Duration
	client         *mongo.Client
}

func (s *MongoService) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.ConnectTimeout)
	defer cancel()
	return s.client.Disconnect(ctx)
}

func (s *MongoService) GetCollection(dbName string, collectionName string) *mongo.Collection {
	return s.client.Database(dbName).Collection(collectionName)
}

func (s *MongoService) Insert(dbName string, collectionName string, document interface{}) (*primitive.ObjectID, error) {
	collection := s.GetCollection(dbName, collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), s.QueryTimeout)
	defer cancel()

	insertResult, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("unable to insert document '%v'. Error: %w", document, err)
	}

	result, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("unable to insert document: %v", api.ERROR_ASSERT_RESULT_TYPE)
	}

	return &result, nil
}

func (s *MongoService) Upsert(dbName string, collectionName string, filter any, update interface{}) (any, error) {
	return s.update(dbName, collectionName, filter, update, true)
}

func (s *MongoService) Update(dbName string, collectionName string, filter any, update interface{}) error {
	_, err := s.update(dbName, collectionName, filter, update, false)
	return err
}

func (s *MongoService) GetCollectionNames(databaseName string, filter any) ([]string, error) {
	db := s.client.Database(databaseName)

	ctx, cancel := context.WithTimeout(context.Background(), s.QueryTimeout)
	defer cancel()

	return db.ListCollectionNames(ctx, filter)
}

func (s *MongoService) Delete(dbName string, collectionName string, filter bson.D) error {
	collection := s.GetCollection(dbName, collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), s.QueryTimeout)
	defer cancel()

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("unable to delete document. filter: '%v'. Error: %w", filter, err)
	}
	return err
}

func (s *MongoService) update(dbName string, collectionName string, filter any, update any, isUpsert bool) (any, error) {
	collection := s.GetCollection(dbName, collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), s.QueryTimeout)
	defer cancel()

	opts := options.Update().SetUpsert(isUpsert)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, fmt.Errorf("unable to update document. Filter: '%v'. Update: '%v'. Error: %w", filter, update, err)
	}

	if result.MatchedCount != 0 {
		return nil, nil
	}

	if isUpsert && result.UpsertedCount != 0 {
		return result.UpsertedID, nil
	}

	return nil, nil
}

func CreateMongoService() (*MongoService, error) {
	connectTimeout := connectTimeout()
	queryTimeout := queryTimeout()
	client, err := createClient(connectTimeout)
	if err != nil {
		return nil, err
	}
	return &MongoService{
		ConnectTimeout: connectTimeout,
		QueryTimeout:   queryTimeout,
		client:         client,
	}, nil
}

func createClient(connectTimeout time.Duration) (*mongo.Client, error) {
	var result *mongo.Client

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	result, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionURL()))
	if err != nil {
		return result, fmt.Errorf("unable to create mongo client: %w", err)
	}

	return result, nil
}

func mongoConnectionURL() string {
	username := utils.EnvVarDefault("MONGO_USERNAME", "mongo_admin")
	password := utils.EnvVarDefault("MONGO_PASSWORD", "mongo_admin_password")
	host := utils.EnvVarDefault("MONGO_HOST", "mongo")
	port := utils.EnvVarDefault("MONGO_PORT", "27017")
	return "mongodb://" + username + ":" + password + "@" + host + ":" + port
}

func connectTimeout() time.Duration {
	valueStr := utils.EnvVarDefault("MONGO_CONNECT_TIMEOUT_IN_SECONDS", "30")

	valueInt, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Info("Unable to read 'MONGO_CONNECT_TIMEOUT_IN_SECONDS' from config, using default value for 30 seconds")
		return 30 * time.Second
	}

	return time.Duration(valueInt) * time.Second
}

func queryTimeout() time.Duration {
	valueStr := utils.EnvVarDefault("MONGO_QUERY_TIMEOUT_IN_SECONDS", "30")

	valueInt, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Info("Unable to read 'MONGO_QUERY_TIMEOUT_IN_SECONDS' from config, using default value for 30 seconds")
		return 30 * time.Second
	}

	return time.Duration(valueInt) * time.Second
}

// add processing of case when we have replica set of mongos and need to use sessions + tx
type QueryFuncVoid func(sc mongo.SessionContext) error

func (s *MongoService) Tx(f QueryFuncVoid) func() error {

	ctx, cancel := context.WithTimeout(context.Background(), s.QueryTimeout)
	defer cancel()

	return func() error {
		session, err := s.client.StartSession()
		if err != nil {
			return fmt.Errorf("unable to start session: %w", err)
		}
		defer session.EndSession(ctx)

		err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
			err = session.StartTransaction()
			if err != nil {
				return fmt.Errorf("unable to start tx: %w", err)
			}
			defer session.AbortTransaction(sc)

			err := f(sc)
			if err != nil {
				return err
			}

			err = session.CommitTransaction(sc)
			if err != nil {
				return fmt.Errorf("unable to commit tx: %w", err)
			}
			return nil
		})

		return err
	}
}
