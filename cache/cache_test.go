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

	t.Run("Returns 1 when the key exist", func(t *testing.T) {
		lrucache := c.New(10)
		lrucache.Set("foo", "bar")

		assert.NotEqual(t, -1, lrucache.Get("foo"))
	})

	t.Run("The consulted item should be the first of items", func(t *testing.T) {
		lrucache := c.New(10)

		lrucache.Set("item1", "value1")
		lrucache.Set("item2", "value2")
		lrucache.Set("item3", "value3")

		lrucache.Get("item2")

		expectedItem := lrucache.Items()[0]

		assert.Equal(t, "item2", expectedItem.Key)
	})

}

func TestCSet(t *testing.T) {
	t.Run("Save a new item when the capacity is not filled", func(t *testing.T) {
		lrucache := c.New(10)

		lrucache.Set("item1", "value1")
		lrucache.Set("item1", "value1")
		lrucache.Set("item3", "value3")

		assert.Equal(t, 3, len(lrucache.Items()))
	})

	t.Run("Save a new item when the capacity is filled", func(t *testing.T) {
		lrucache := c.New(2)

		lrucache.Set("item1", "value1")
		lrucache.Set("item1", "value1")
		lrucache.Set("item3", "value3")
		lrucache.Set("item4", "value4")

		assert.Equal(t, 2, len(lrucache.Items()))
		assert.Equal(t, "item4", lrucache.Items()[0].Key)
	})

}
