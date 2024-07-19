package cachelru_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/douglasdoro/cachelru/cachelru"
)

func TestCGet(t *testing.T) {
	t.Run("Returns -1 when the key doesn't exist", func(t *testing.T) {
		cacheLRU := c.NewCacheLRU(10)

		assert.Equal(t, -1, cacheLRU.Get("sample"))
	})

	t.Run("Returns 1 when the key exist", func(t *testing.T) {
		cacheLRU := c.NewCacheLRU(10)
		cacheLRU.Set("foo", "bar")

		assert.NotEqual(t, -1, cacheLRU.Get("foo"))
	})

	t.Run("The consulted item should be the first of items", func(t *testing.T) {
		cacheLRU := c.NewCacheLRU(10)

		cacheLRU.Set("item1", "value1")
		cacheLRU.Set("item2", "value2")
		cacheLRU.Set("item3", "value3")

		cacheLRU.Get("item2")

		expectedItem := cacheLRU.Items()[0]

		assert.Equal(t, "item2", expectedItem.Key)
	})

}

func TestCSet(t *testing.T) {
	t.Run("Save a new item when the capacity is not filled", func(t *testing.T) {
		cacheLRU := c.NewCacheLRU(10)

		cacheLRU.Set("item1", "value1")
		cacheLRU.Set("item1", "value1")
		cacheLRU.Set("item3", "value3")

		assert.Equal(t, 3, len(cacheLRU.Items()))
	})

	t.Run("Save a new item when the capacity is filled", func(t *testing.T) {
		cacheLRU := c.NewCacheLRU(2)

		cacheLRU.Set("item1", "value1")
		cacheLRU.Set("item1", "value1")
		cacheLRU.Set("item3", "value3")
		cacheLRU.Set("item4", "value4")

		assert.Equal(t, 2, len(cacheLRU.Items()))
		assert.Equal(t, "item4", cacheLRU.Items()[0].Key)
	})

}
