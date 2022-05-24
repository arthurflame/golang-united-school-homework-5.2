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
	data []map[string]Data


}

func NewCache() Cache {
	return Cache{}
}

func main() {
	cache := NewCache()

	cache.Put("1", "1111111111111")
	//cache.Put("2", "2222222222222")
	//cache.Put("2", "3333333333333")
	fmt.Println("cache:", cache)
}

func (c *Cache) Put(key, value string) {

	data := Data{
		value:          value,
		canExpire:      false,
		expirationTime: 0,
	}


	//c.data = append(c.data, map[string]Data{key:data})

	//fmt.Println(c.folder[0], data)


	fmt.Println("c.data", c.data)

	for i := 0; i < len(c.data); i++ {
		if v, ok := c.data[i][key]; !ok {
			c.data = append(c.data, map[string]Data{key:data})
			//fmt.Println(k)
			//c.data[0]["1"] = Data{value: "wewewew"}
		} else {
			c.data[i][key] = Data{value: value}
		}
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
