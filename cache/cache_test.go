package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/douglasdoro/lrucache/cache"
)

func TestCGet(t *testing.T) {
	t.Run("Returns -1 when the key doesn't exist", func(t *testing.T) {
		lrucache := c.New(10)

		assert.Equal(t, -1, lrucache.Get("sample"))
	})

	t.Run("Returns Item when the key exist", func(t *testing.T) {
		lrucache := c.New(10)
		lrucache.Set("foo", "bar")

		cacheItem := lrucache.Get("foo")

		assert.NotEqual(t, -1, cacheItem)
		assert.Equal(t, "bar", cacheItem)
	})

	t.Run("The consulted Item should be the first of items when has only one Item", func(t *testing.T) {
		lrucache := c.New(10)

		lrucache.Set("item1", "value1")

		expectedItem := lrucache.Get("item1")

		assert.Equal(t, lrucache.Items()[0].Value, expectedItem)
	})

	t.Run("The consulted item should be the first of items", func(t *testing.T) {
		lrucache := c.New(10)

		lrucache.Set("item1", "value1")
		lrucache.Set("item2", "value2")
		lrucache.Set("item3", "value3")

		expectedItem := lrucache.Get("item2")

		assert.Equal(t, lrucache.Items()[0].Value, expectedItem)
	})

}

func TestCSet(t *testing.T) {
	t.Run("Save a new item when the capacity is not filled", func(t *testing.T) {
		lrucache := c.New(10)

		lrucache.Set("item1", "value1")
		lrucache.Set("item2", "value2")
		lrucache.Set("item3", "value3")

		assert.Equal(t, 3, len(lrucache.Items()))
	})

	t.Run("Save a new item when the capacity is filled", func(t *testing.T) {
		lrucache := c.New(2)

		lrucache.Set("item1", "value1")
		lrucache.Set("item2", "value2")
		lrucache.Set("item3", "value3")
		lrucache.Set("item4", "value4")

		assert.Equal(t, 2, len(lrucache.Items()))
	})

	t.Run("Save an existing item", func(t *testing.T) {
		lrucache := c.New(5)

		lrucache.Set("item1", "value1")
		lrucache.Set("item1", "value2")
		lrucache.Set("item2", "value3")
		lrucache.Set("item2", "value4")

		assert.Equal(t, 2, len(lrucache.Items()))
	})

	t.Run("Save an existing item (randon add)", func(t *testing.T) {
		lrucache := c.New(10)

		lrucache.Set("item1", "value2")
		lrucache.Set("item2", "value3")
		lrucache.Set("item2", "value3")
		lrucache.Set("item1", "value1")
		lrucache.Set("item1", "value1")
		lrucache.Set("item2", "value4")
		lrucache.Set("item2", "value3")
		lrucache.Set("item1", "value1")

		assert.Equal(t, 2, len(lrucache.Items()))
	})

	t.Run("Save an existing item when the capacity is not filled", func(t *testing.T) {
		lrucache := c.New(3)

		lrucache.Set("item1", "value2")
		lrucache.Set("item2", "value3")
		lrucache.Set("item3", "value3")
		lrucache.Set("item4", "value1")
		lrucache.Set("item1", "value1")
		lrucache.Set("item2", "value4")
		lrucache.Set("item2", "value3")
		lrucache.Set("item1", "value1")

		assert.Equal(t, 3, len(lrucache.Items()))
	})

}
