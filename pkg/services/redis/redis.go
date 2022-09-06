package redis

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client       *redis.Client
	queryTimeout time.Duration
}

func CreateRedisService() *RedisService {
	return &RedisService{
		client:       createClient(),
		queryTimeout: queryTimeout(),
	}
}

func (s *RedisService) Shutdown() error {
	return s.client.Close()
}

type QueryFunc func(cli *redis.Client, ctx context.Context, cancel context.CancelFunc) (any, error)
type QueryFuncVoid func(cli *redis.Client, ctx context.Context, cancel context.CancelFunc) error

func (s *RedisService) WithTimeout(f QueryFunc) func() (any, error) {
	cli := s.client
	timeout := s.queryTimeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return func() (any, error) {
		defer cancel()
		result, err := f(cli, ctx, cancel)
		if err != nil {
			return result, err
		}
		return result, err
	}
}
func (s *RedisService) WithTimeoutVoid(f QueryFuncVoid) func() error {
	cli := s.client
	timeout := s.queryTimeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return func() error {
		defer cancel()
		return f(cli, ctx, cancel)
	}
}

func createClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     utils.EnvVar("REDIS_HOST") + ":" + utils.EnvVar("REDIS_POST"),
		Username: utils.EnvVar("REDIS_USER"),
		Password: utils.EnvVar("REDIS_PASSWORD"),
		DB:       utils.EnvVarIntDefault("REDIS_DATABASE_NUMBER", 0),
	})
}

func queryTimeout() time.Duration {
	valueStr := utils.EnvVarDefault("REDIS_QUERY_TIMEOUT_IN_SECONDS", "30")

	valueInt, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Printf("Unable to read 'REDIS_QUERY_TIMEOUT_IN_SECONDS' from config, using default value for 30 seconds")
		return 30 * time.Second
	}

	return time.Duration(valueInt) * time.Second
}
