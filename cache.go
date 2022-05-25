package main
import (
	"fmt"
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

func main() {
	cache := NewCache()
	cache.Put("1", "222")
	cache.Put("2", "23")

	cache.PutTill("2", "should exist", time.Now().Add(time.Second*20))
	cache.PutTill("3", "shouldn't exist", time.Now())

	//i, ok := cache.Get("3")


	fmt.Println(cache.Keys())

	fmt.Println(cache)
	fmt.Println("\n\n")
	//fmt.Println(i, ok)
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
	if i, ok := c.data[key]; ok {
		exists = ok
		value = i.value
	} else {
		return fmt.Sprintf("the requested key: [%v] has expired or doesn't exist.\n", key), false
	}

	deleteExpired(c.data)
	return value, exists
}


func (c *Cache) Keys() []string {
	var keys []string

	for i := range c.data {
		if _, ok := c.data[i]; ok {
			deleteExpired(c.data)
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


func deleteExpired(data map[string]Data) {
	for i, v := range data {
		if time.Now() == time.Now().Add(time.Until(v.expirationTime)) {
			delete(data, i)
		}
	}
}