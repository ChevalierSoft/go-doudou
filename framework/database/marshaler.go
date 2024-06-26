package database

import (
	"context"
	"encoding/json"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/gocache/lib/cache"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/gocache/lib/store"
)

// Marshaler is the struct that marshal and unmarshal cache values
type Marshaler struct {
	cache cache.CacheInterface[any]
}

// NewMarshaler creates a new marshaler that marshals/unmarshals cache values
func NewMarshaler(cache cache.CacheInterface[any]) *Marshaler {
	return &Marshaler{
		cache: cache,
	}
}

// Get obtains a value from cache and unmarshal value with given object
func (c *Marshaler) Get(ctx context.Context, key any, returnObj any) (any, error) {
	result, err := c.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {
	case []byte:
		err = json.Unmarshal(v, returnObj)
	case string:
		err = json.Unmarshal([]byte(v), returnObj)
	}

	if err != nil {
		return nil, err
	}

	return returnObj, nil
}

// Set sets a value in cache by marshaling value
func (c *Marshaler) Set(ctx context.Context, key, object any, options ...store.Option) error {
	bytes, err := json.Marshal(object)
	if err != nil {
		return err
	}

	return c.cache.Set(ctx, key, bytes, options...)
}

// Delete removes a value from the cache
func (c *Marshaler) Delete(ctx context.Context, key any) error {
	return c.cache.Delete(ctx, key)
}

// Invalidate invalidate cache values using given options
func (c *Marshaler) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	return c.cache.Invalidate(ctx, options...)
}

// Clear reset all cache data
func (c *Marshaler) Clear(ctx context.Context) error {
	return c.cache.Clear(ctx)
}
