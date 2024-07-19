package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCGet(t *testing.T) {
	t.Run("Returns -1 when the key doesn't exist", func(t *testing.T) {
		c := NewCacheLRU(10)

		assert.Equal(t, -1, c.Get("sample"))
	})

	t.Run("Returns 1 when the key exist", func(t *testing.T) {
		c := NewCacheLRU(10)
		c.items = append(c.items, Cache{key: "foo", value: "bar"})

		assert.NotEqual(t, -1, c.Get("foo"))
	})

	t.Run("The consulted item should be the first of items", func(t *testing.T) {
		c := NewCacheLRU(10)

		item2 := Cache{key: "item2", value: "value2"}

		c.items = append(c.items, Cache{key: "item1", value: "value1"})
		c.items = append(c.items, item2)
		c.items = append(c.items, Cache{key: "item3", value: "value3"})

		c.Get(item2.key)

		assert.Equal(t, item2, c.items[0])
	})

}

func TestCSet(t *testing.T) {
	t.Run("Save a new item when the capacity is not filled", func(t *testing.T) {
		c := NewCacheLRU(10)

		c.Set("item1", "value1")
		c.Set("item1", "value1")
		c.Set("item3", "value3")

		assert.Equal(t, 3, len(c.items))
	})

	t.Run("Save a new item when the capacity is filled", func(t *testing.T) {
		c := NewCacheLRU(2)

		c.Set("item1", "value1")
		c.Set("item1", "value1")
		c.Set("item3", "value3")
		c.Set("item4", "value4")

		assert.Equal(t, 2, len(c.items))
		assert.Equal(t, "item4", c.items[0].key)
	})

}
