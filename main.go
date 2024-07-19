package main

import (
	"fmt"

	c "github.com/douglasdoro/lrucache/cache"
)

func main() {
	cacheLRU := c.New(10)
	cacheLRU.Set("item1", "value1")
	cacheLRU.Set("item1", "value1")
	cacheLRU.Set("item3", "value3")
	cacheLRU.Set("item4", "value4")

	fmt.Println("Cache LRU: ", cacheLRU)
}
