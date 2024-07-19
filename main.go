package main

import "fmt"

func main() {
	cacheLRU := NewCacheLRU(10)
	cacheLRU.Set("item1", "value1")
	cacheLRU.Set("item1", "value1")
	cacheLRU.Set("item3", "value3")
	cacheLRU.Set("item4", "value4")

	fmt.Println("Cache LRU: ", cacheLRU)
}

type Cache struct {
	key   string
	value string
}

type CacheLRU struct {
	Capacity int
	items    []Cache
}

func NewCacheLRU(capacity int) *CacheLRU {
	return &CacheLRU{Capacity: capacity, items: []Cache{}}
}

func (c *CacheLRU) Get(key string) interface{} {
	for i, v := range c.items {
		if v.key == key {
			c.addToBeginningCacheItems(v, i)

			return v
		}
	}

	return -1
}

func (c *CacheLRU) Set(key, value string) {
	if len(c.items) < c.Capacity {
		c.addToBeginningCacheItems(Cache{key: key, value: value}, 0)
	} else {
		c.addToBeginningCacheItems(Cache{key: key, value: value}, -1)
	}
}

func (c *CacheLRU) addToBeginningCacheItems(cacheItem Cache, currentPosition int) {
	slc := []Cache{}
	slc = append(slc, cacheItem)

	if currentPosition == 0 {
		c.items = append(slc, c.items...)
	} else if currentPosition == -1 {
		c.items = append(slc, c.items[1:len(c.items)-1]...)
	} else {
		c.items = append(slc, append(c.items[:currentPosition], c.items[currentPosition+1:]...)...)
	}

}
