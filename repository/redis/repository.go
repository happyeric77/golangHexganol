package redis

import(
	"github.com/go-redis/redis"
	errs "github.com/pkg/errors"
	"project5/shortener"
	"fmt"
	"strconv"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient (redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err !=nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string) (shortener.RedirectRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errs.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r *redisRepository) Find(code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errs.Wrap(err, "repository.Redirect.Find - get data")
	}
	if len(data) == 0 {
		return nil, errs.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
	}
	createAt, err := strconv.ParseInt(data["create_at"], 10, 64)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Redirect.Find - get create_at")
	}
	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createAt
	return redirect, nil
}

func (r *redisRepository) Store(redirect *shortener.Redirect) error {
	key := r.generateKey(redirect.Code)
	data := map[string]interface{}{
		"code": redirect.Code,
		"url": redirect.URL,
		"create_at": redirect.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errs.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}

