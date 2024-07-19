package cachelru

type Cache struct {
	Key   string
	Value string
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
		if v.Key == key {
			c.addToBeginningCacheItems(v, i)

			return v
		}
	}

	return -1
}

func (c *CacheLRU) Items() []Cache {
	items := make([]Cache, len(c.items))
	copy(items, c.items)

	return items
}

func (c *CacheLRU) Set(key, value string) {
	if len(c.items) < c.Capacity {
		c.addToBeginningCacheItems(Cache{Key: key, Value: value}, 0)
	} else {
		c.addToBeginningCacheItems(Cache{Key: key, Value: value}, -1)
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
