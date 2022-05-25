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
	cache.Put("1", "222")
	cache.Put("2", "222")
	cache.Put("1", "333")
	cache.PutTill("33", "111", time.Now().Add(time.Second*10))
	cache.PutTill("15", "111", time.Now())

	fmt.Println(cache.Keys())

	fmt.Println(cache)
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

	if !calcTime(time.Now(), data.expirationTime) {
		//c.data[key]
	}

}

func calcTime(timeNow, deadline time.Time) bool {
	if timeNow.Before(deadline) {
		return true
	}
	return false
}
