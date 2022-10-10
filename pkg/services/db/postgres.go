package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	_ "github.com/lib/pq"
)

type DBParams struct {
	host         string
	port         string
	username     string
	password     string
	databaseName string
	sslMode      string
}

type PostgreSQLService struct {
	client       *sql.DB
	queryTimeout time.Duration
}

type SqlQueryFunc func(transaction *sql.Tx, ctx context.Context, cancel context.CancelFunc) (any, error)

type SqlQueryFuncVoid func(transaction *sql.Tx, ctx context.Context, cancel context.CancelFunc) error

func CreatePostgreSQLServiceDefault() *PostgreSQLService {
	return &PostgreSQLService{
		client:       createClientDefault(),
		queryTimeout: queryTimeout(),
	}
}

func CreatePostgreSQLService(params *DBParams) *PostgreSQLService {
	return &PostgreSQLService{
		client:       createClient(params),
		queryTimeout: queryTimeout(),
	}
}

func (s *PostgreSQLService) Shutdown() error {
	return s.client.Close()
}

func (s *PostgreSQLService) Tx(f SqlQueryFunc) func() (any, error) {
	database := s.client
	timeout := s.queryTimeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return func() (any, error) {
		defer cancel()
		tx, err := database.BeginTx(ctx, nil)
		if err != nil {
			return -1, fmt.Errorf("error at creating tx: %s", err)
		}
		defer tx.Rollback()
		result, err := f(tx, ctx, cancel)
		if err != nil {
			return result, err
		}
		err = tx.Commit()
		if err != nil {
			return -1, fmt.Errorf("error at commiting tx: %s", err)
		}
		return result, err
	}
}

func (s *PostgreSQLService) TxVoid(f SqlQueryFuncVoid) func() error {
	database := s.client
	timeout := s.queryTimeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return func() error {
		defer cancel()
		tx, err := database.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("error at creating tx: %s", err)
		}
		defer tx.Rollback()
		err = f(tx, ctx, cancel)
		if err != nil {
			return err
		}
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error at commiting tx: %s", err)
		}
		return err
	}
}

func createClientDefault() *sql.DB {
	defaultParams := &DBParams{
		host:         utils.EnvVar("DATABASE_HOST"),
		port:         utils.EnvVar("DATABASE_PORT"),
		username:     utils.EnvVar("DATABASE_USER"),
		password:     utils.EnvVar("DATABASE_PASSWORD"),
		databaseName: utils.EnvVar("DATABASE_NAME"),
		sslMode:      utils.EnvVar("DATABASE_SSL_MODE"),
	}

	return createClient(defaultParams)
}

func createClient(params *DBParams) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		params.host,
		params.port,
		params.username,
		params.password,
		params.databaseName,
		params.sslMode,
	)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to open connection to database: %s. Error: %s", params.databaseName, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = result.PingContext(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s. Error: %s", params.databaseName, err)
	}

	return result
}

func queryTimeout() time.Duration {
	valueStr := utils.EnvVarDefault("DATABASE_QUERY_TIMEOUT_IN_SECONDS", "30")

	valueInt, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Info("Unable to read 'DATABASE_QUERY_TIMEOUT_IN_SECONDS' from config, using default value for 30 seconds")
		return 30 * time.Second
	}

	return time.Duration(valueInt) * time.Second
}
