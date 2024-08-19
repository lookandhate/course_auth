package user

import (
	"context"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/lookandhate/course_auth/internal/cache/model"
	"github.com/lookandhate/course_auth/internal/config"
	"github.com/lookandhate/course_platform_lib/pkg/cache/redis"
)

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisPool *redigo.Pool, redisCfg config.RedisConfig) *RedisCache {
	client := redis.NewClient(redisPool, time.Duration(redisCfg.IdleTimeout))

	return &RedisCache{redisClient: client}
}

// Create user record in cache as model.UserModel.
func (r RedisCache) Create(ctx context.Context, user *model.UserModel) error {
	err := r.redisClient.HashSet(ctx, r.userKey(user.ID), user)
	if err != nil {
		return err
	}

	return nil
}

// Get user record from cache.
func (r RedisCache) Get(ctx context.Context, id int) (*model.UserModel, error) {
	response, err := r.redisClient.HGetAll(ctx, r.userKey(int64(id)))
	if err != nil {
		return nil, err
	}

	var user model.UserModel
	err = redigo.ScanStruct(response, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete record from cache.
func (r RedisCache) Delete(ctx context.Context, id int) error {
	return r.redisClient.Del(ctx, r.userKey(int64(id)))
}

// userKey returns redis key for user.
func (r RedisCache) userKey(id int64) string {
	return fmt.Sprintf("user_id-%d", id)
}
