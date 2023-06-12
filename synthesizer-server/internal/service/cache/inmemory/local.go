package inmemory

import (
	"synthesizer-server/internal/entity"
	"synthesizer-server/internal/service"

	"github.com/patrickmn/go-cache"
)

type inMemory struct {
	c *cache.Cache
}

func (i *inMemory) Get(hash string) (entity.Data[string], bool, error) {
	if res, ok := i.c.Get(hash); ok {
		return res.(entity.Data[string]), ok, nil
	}
	return entity.Data[string]{}, false, nil
}

func (i *inMemory) Set(k string, v entity.Data[string]) error {
	i.c.SetDefault(k, v)
	return nil
}

func New() service.DataCache {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	return &inMemory{
		c: c,
	}
}
