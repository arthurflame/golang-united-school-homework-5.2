package cache

import (
	"time"
)

type Data struct {
	value          string
	expirationTime time.Time
}

type Cache struct {
	data map[string]Data
}

func NewCache() Cache {
	return Cache{
		make(map[string]Data),
	}
}

func (c *Cache) Put(key, value string) {
	data := Data{
		value:          value,
		expirationTime: time.Time{},
	}

	if _, ok := c.data[key]; !ok {
		c.data[key] = data
	}
	c.data[key] = data
	
}

func (c *Cache) Get(key string) (string, bool) {
	// TODO: figure out why this doesn't work
	var value string
	var exists bool
	var z time.Time
	
	if i, ok := c.data[key]; ok && time.Now().Before(i.expirationTime) || i.expirationTime == z {
		exists = ok
		value = i.value
	} else {
		return "", false
	}
	return value, exists
}


func (c *Cache) Keys() []string {
	// TODO: figure out why this doesn't fcking !!!!work
	var keys []string
	var z time.Time

	for i := range c.data {
		if _, ok := c.data[i]; ok && time.Now().Before(v.expirationTime) || v.expirationTime == z {
			keys = append(keys, i)
		}
	}
	return keys
}


func (c *Cache) PutTill(key, value string, deadline time.Time) {
	data := Data{
		value:          value,
		expirationTime: deadline,
	}

	if _, ok := c.data[key]; !ok {
		c.data[key] = data
	}
	c.data[key] = data
	
}

