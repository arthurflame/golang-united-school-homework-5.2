package main

import (
	"fmt"
)
type Data struct {
	value string
	canExpire bool
	expirationTime int
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
	cache.Put("2", "3333333333333")
	//cache.Put("3", "3333333333333")
	fmt.Println("cache:", cache)
}


func (c *Cache) Put(key, value string) {

	data := Data{
		value:          value,
		canExpire:      false,
		expirationTime: 0,
	}

	if _, ok := c.data[key]; !ok {
		c.data[key] = data
	} else {
		c.data[key] = data
	}

}



//func (c *Cache) Get(key string) (string, bool) {
//
//}
//
//
//func (receiver) Keys() []string {
//}
//
//func (receiver) PutTill(key, value string, deadline time.Time) {
//}
