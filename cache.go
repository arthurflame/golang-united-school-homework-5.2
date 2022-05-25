package main

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

func main() {
	cache := NewCache()

	cache.Put("1", "1111111111111")
	//cache.Put("2", "2222222222222")
	cache.PutTill("2", "2222222222222", time.Now().Add(time.Second * 20))
	k, ok := cache.Get("2")
	fmt.Println(k, ok)


	fmt.Println("cache:", cache)
	keys := cache.Keys()
	fmt.Println("keys", keys)
}

func (c *Cache) Put(key, value string) {
	data := Data{
		value:          value,
		canExpire:      false,
		expirationTime: time.Time{},
	}

	if _, ok := c.data[key]; !ok {
		c.data[key] = data
	} else {
		c.data[key] = data
	}
}

func (c *Cache) Get(key string) (string, bool) {
	var value string
	var exists bool
	if i, ok := c.data[key]; ok && !i.canExpire {
		exists = ok
		value = i.value
	} else if ok && i.canExpire {
		now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
		expiration := time.Date(i.expirationTime.Year(), i.expirationTime.Month(), i.expirationTime.Day(), i.expirationTime.Hour(), i.expirationTime.Minute(), i.expirationTime.Second(), i.expirationTime.Nanosecond(), time.UTC)
		if now.Before(expiration) {
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
			now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
			expiration := time.Date(k.expirationTime.Year(), k.expirationTime.Month(), k.expirationTime.Day(), k.expirationTime.Hour(), k.expirationTime.Minute(), k.expirationTime.Second(), k.expirationTime.Nanosecond(), time.UTC)
			if now.Before(expiration) {
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
	} else {
		c.data[key] = data
	}
}
