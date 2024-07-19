package main

import (
	"fmt"

	c "github.com/douglasdoro/cachelru/cachelru"
)

func main() {
	cacheLRU := c.NewCacheLRU(10)
	cacheLRU.Set("item1", "value1")
	cacheLRU.Set("item1", "value1")
	cacheLRU.Set("item3", "value3")
	cacheLRU.Set("item4", "value4")

	fmt.Println("Cache LRU: ", cacheLRU)
}
