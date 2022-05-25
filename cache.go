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
	var value string
	var exists bool
	
	if i, ok := c.data[key]; ok && !time.Now().Before(i.expirationTime) {
		exists = ok
		value = i.value
	} else {
		return "", false
	}
	return value, exists
}


func (c *Cache) Keys() []string {
	var keys []string

	for i := range c.data {
		if k, ok := c.data[i]; ok && !time.Now().Before(k.expirationTime) {
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

