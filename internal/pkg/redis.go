package pkg

import "github.com/go-redis/redis"

type redisClient struct {
	client *redis.Client
}

func New(address string) redisClient {
	return redisClient{
		client: redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}

func (r redisClient) SetPair(key, value string) error {
	return r.client.Set(key, value, 0).Err()
}

func (r redisClient) GetPair(key string) (string, error) {
	return r.client.Get(key).Result()
}
