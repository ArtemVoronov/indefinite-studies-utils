package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	_ "github.com/lib/pq"
)

type PostgreSQLService struct {
	client       *sql.DB
	queryTimeout time.Duration
}

type SqlQueryFunc func(transaction *sql.Tx, ctx context.Context, cancel context.CancelFunc) (any, error)

type SqlQueryFuncVoid func(transaction *sql.Tx, ctx context.Context, cancel context.CancelFunc) error

func CreatePostgreSQLService() *PostgreSQLService {
	return &PostgreSQLService{
		client:       createClient(),
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

func createClient() *sql.DB {
	dbEnvVars := [6]string{"DATABASE_HOST", "DATABASE_PORT", "DATABASE_USER", "DATABASE_PASSWORD", "DATABASE_NAME", "DATABASE_SSL_MODE"}
	var variables []string
	for _, element := range dbEnvVars {
		value := utils.EnvVar(element)
		variables = append(variables, value)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", variables[0], variables[1], variables[2], variables[3], variables[4], variables[5])
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to database : %s", err)
	}

	// log.Printf("----- Database service setup succeed. Database name: %s -----", variables[4])

	return result
}

func queryTimeout() time.Duration {
	valueStr := utils.EnvVarDefault("DATABASE_QUERY_TIMEOUT_IN_SECONDS", "30")

	valueInt, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Printf("Unable to read 'DATABASE_QUERY_TIMEOUT_IN_SECONDS' from config, using default value for 30 seconds")
		return 30 * time.Second
	}

	return time.Duration(valueInt) * time.Second
}
