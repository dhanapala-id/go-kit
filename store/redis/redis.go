package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"

	"github.com/dhanapala-id/go-idempotency/store"
)

type redisStore struct {
	db *redis.Client
}

func New(addr, password string, db int) store.Store {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisStore{db: rdb}
}

func (s *redisStore) Lock(ctx context.Context, key string, exp time.Duration) error {
	cacheKey := fmt.Sprintf("idemlock_%s", key)

	txfn := func(tx *redis.Tx) error {
		exists, err := tx.Get(ctx, cacheKey).Bool()
		if err != nil && err != redis.Nil {
			return err
		}

		if exists {
			return store.ErrUnableToLock
		}

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, cacheKey, true, exp)
			return nil
		})
		return err
	}

	if err := s.db.Watch(ctx, txfn, cacheKey); err == redis.TxFailedErr {
		return store.ErrUnableToLock
	} else if err != nil {
		return err
	}

	return nil
}

func (s *redisStore) Unlock(ctx context.Context, key string) error {
	cacheKey := fmt.Sprintf("idemlock_%s", key)
	if err := s.db.Del(ctx, cacheKey).Err(); err != nil {
		return err
	}
	return nil
}

func (s *redisStore) Get(ctx context.Context, key string) (*store.Data, error) {
	cacheKey := fmt.Sprintf("idemdata_%s", key)
	b, err := s.db.Get(ctx, cacheKey).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var data store.Data
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *redisStore) Set(ctx context.Context, key string, data *store.Data, exp time.Duration) error {
	if data != nil && data.StatusCode == 0 {
		data.StatusCode = 200
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("idemdata_%s", key)
	if err := s.db.SetEX(ctx, cacheKey, b, exp).Err(); err != nil {
		return err
	}
	return err
}
