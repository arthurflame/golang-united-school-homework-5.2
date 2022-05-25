package cache

import (
	"fmt"
	"time"
)

type Data struct {
	value          string
	canExpire      bool
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
		canExpire:      false,
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
	if i, ok := c.data[key]; ok && !i.canExpire {
		exists = ok
		value = i.value
	} else if ok && i.canExpire {
		if calcTime(time.Now(), i.expirationTime) {
			exists = ok
			value = i.value
		}
	} else {
		return fmt.Sprintf("the requested key: [%v] has expired or doesn't exist.\n", key), false
	}
	return value, exists
}


func (c *Cache) Keys() []string {
	var keys []string

	for i := range c.data {
		if k, ok := c.data[i]; ok && !k.canExpire {
			keys = append(keys, i)
		} else if ok && k.canExpire {
			if calcTime(time.Now(), k.expirationTime) {
				keys = append(keys, i)
			}
		}
	}
	return keys
}


func (c *Cache) PutTill(key, value string, deadline time.Time) {
	data := Data{
		value:          value,
		canExpire:      true,
		expirationTime: deadline,
	}

	if _, ok := c.data[key]; !ok {
		c.data[key] = data
	}
	c.data[key] = data

}


func calcTime(timeNow, deadline time.Time) bool {
	now := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), time.UTC)
	expiration := time.Date(deadline.Year(), deadline.Month(), deadline.Day(), deadline.Hour(), deadline.Minute(), deadline.Second(), deadline.Nanosecond(), time.UTC)

	if now.Before(expiration) {
		return true
	}
	return false
}
