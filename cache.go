package main

import (
	"fmt"
	"time"
)
type Data struct {
	value string
	canExpire bool
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
	cache.Put("2", "2222222222222")
	//cache.Put("2", "3333333333333")
	//cache.Put("3", "3333333333333")
	k, ok := cache.Get("2")
	fmt.Println(k, ok)
	//fmt.Println("cache:", cache)
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

	if i, ok := c.data[key]; ok  && !i.canExpire {
		exists = ok
		value = i.value
	}
	return value, exists

}


//
//
//func (receiver) Keys() []string {
//}
//
//func (receiver) PutTill(key, value string, deadline time.Time) {
//}
