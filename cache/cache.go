package cache

type CacheItem struct {
	Key   string
	Value string
}

type Cache struct {
	Capacity int
	items    []CacheItem
}

func New(capacity int) *Cache {
	return &Cache{Capacity: capacity, items: []CacheItem{}}
}

func (c *Cache) Get(key string) interface{} {
	for i, v := range c.items {
		if v.Key == key {
			c.addToBeginningCacheItems(v, i)

			return v
		}
	}

	return -1
}

func (c *Cache) Items() []CacheItem {
	items := make([]CacheItem, len(c.items))
	copy(items, c.items)

	return items
}

func (c *Cache) Set(key, value string) {
	if len(c.items) < c.Capacity {
		c.addToBeginningCacheItems(CacheItem{Key: key, Value: value}, 0)
	} else {
		c.addToBeginningCacheItems(CacheItem{Key: key, Value: value}, -1)
	}
}

func (c *Cache) addToBeginningCacheItems(cacheItem CacheItem, currentPosition int) {
	slc := []CacheItem{}
	slc = append(slc, cacheItem)

	if currentPosition == 0 {
		c.items = append(slc, c.items...)
	} else if currentPosition == -1 {
		c.items = append(slc, c.items[1:len(c.items)-1]...)
	} else {
		c.items = append(slc, append(c.items[:currentPosition], c.items[currentPosition+1:]...)...)
	}

}
