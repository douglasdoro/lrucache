package cache

type CacheItem struct {
	Key   string
	Value any
	prev  *CacheItem
	next  *CacheItem
}

type Cache struct {
	Capacity      int
	items         map[string]*CacheItem
	itemsPosition *CacheItem
	head          *CacheItem
}

func New(capacity int) *Cache {
	return &Cache{
		Capacity: capacity,
		items:    map[string]*CacheItem{}}
}

func (c *Cache) Get(key string) any {
	item, ok := c.items[key]
	if !ok {
		return -1
	}

	if item == c.head {
		return item.Value
	}

	if item.prev != nil {
		item.prev.next = item.next
	}

	if item.next != nil {
		item.next.prev = item.prev
	}

	if item == c.itemsPosition {
		c.itemsPosition = item.next
	}

	c.head.next = item
	item.prev = c.head
	c.head = item
	c.head.next = nil

	return item.Value
}

func (c *Cache) Set(key string, value any) {
	// if key already exist, remove from linked list
	if item, ok := c.items[key]; ok {
		if item.prev != nil {
			item.prev.next = item.next
		}

		if item.next != nil {
			item.next.prev = item.prev
		}

		if item == c.itemsPosition {
			c.itemsPosition = c.itemsPosition.next
		}

		if item == c.head {
			c.head = item.prev
		}

		delete(c.items, item.Key)
	}

	if len(c.items) == c.Capacity {
		leastItem := c.itemsPosition

		delete(c.items, leastItem.Key)

		c.itemsPosition.Key = key
		c.itemsPosition.Value = value
		c.items[key] = c.itemsPosition

		return
	}

	newItem := CacheItem{Key: key, Value: value}

	c.items[key] = &newItem
	newItem.next = c.itemsPosition

	if c.itemsPosition != nil {
		c.itemsPosition.prev = &newItem
	}

	c.itemsPosition = &newItem

	if c.itemsPosition.next == nil {
		c.head = c.itemsPosition
	}
}

func (c *Cache) Items() []CacheItem {
	items := []CacheItem{}

	current := c.head

	for current != nil {
		items = append(items, *current)
		current = current.prev
	}

	return items
}
