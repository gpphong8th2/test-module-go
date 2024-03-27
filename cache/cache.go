package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var wrapMemcache *gocache.Cache

func InitCache() {
	wrapMemcache = gocache.New(5*time.Minute, 10*time.Minute)
}

func GetCache(key string) (interface{}, bool) {
	return wrapMemcache.Get(key)
}

func SetCache(key string, value interface{}, ttl time.Duration) {
	wrapMemcache.Set(key, value, ttl)
}
