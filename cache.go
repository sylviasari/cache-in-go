package main

import(
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
)

var c *cache.Cache

func initCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)

}

func setCache(key, value string) {
	c.Set(key, value, cache.DefaultExpiration)
}

func getCache(key string) (err error,result string){
	if x, found := c.Get(key); found {
		result = x.(string)
		return nil, result
	}
	return fmt.Errorf("not found"), result
}