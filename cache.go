package cache

import "time"

type keyPair struct {
	value    string
	deadline time.Time
}

type Cache struct {
	keyPairs map[string]keyPair
}

func NewCache() Cache {
	return Cache{
		keyPairs: make(map[string]keyPair),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	cache, ok := c.keyPairs[key]
	currentTime := time.Now()

	if !cache.deadline.IsZero() && currentTime.After(cache.deadline) {
		delete(c.keyPairs, key)

		return "", false
	}

	if !ok {
		return "", false
	}

	return cache.value, true
}

func (c *Cache) Put(key, value string) {
	c.keyPairs[key] = keyPair{
		value:    value,
		deadline: time.Time{},
	}
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.keyPairs))
	currentTime := time.Now()

	for k, val := range c.keyPairs {
		if !val.deadline.IsZero() && currentTime.After(val.deadline) {
			delete(c.keyPairs, k)

			continue
		}
		keys = append(keys, k)
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.keyPairs[key] = keyPair{
		value:    value,
		deadline: deadline,
	}
}
