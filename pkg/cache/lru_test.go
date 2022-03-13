package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/pkg/cache"
)

type LRU = cache.LRU[string, string]

type Op func(t *testing.T, c *LRU)

func set(k, v string) Op {
	return func(t *testing.T, c *LRU) {
		c.Set(k, v)
	}
}

func get(k, want string) Op {
	return func(t *testing.T, c *LRU) {
		v, ok := c.Get(k)
		assert.Equal(t, want != "", ok)
		assert.Equal(t, want, v)
	}
}

func remove(key string) Op {
	return func(t *testing.T, c *LRU) {
		c.Remove(key)
	}
}

type TestCase struct {
	Name string
	Cap  int
	Ops  []Op
}

func TestLRU(t *testing.T) {
	t.Run("Homework", testHomework)

	testCases := []TestCase{
		{
			Name: "Capacity",
			Cap:  3,
			Ops: []Op{
				set("a", "A"),
				set("b", "B"),
				set("c", "C"),
				set("d", "D"),
				set("e", "E"),
				get("a", ""),
				get("b", ""),
				get("c", "C"),
				get("d", "D"),
				get("e", "E"),
			},
		},
		{
			Name: "RemoveEmpty",
			Cap:  3,
			Ops: []Op{
				set("a", "A"),
				set("b", "B"),
				remove("c"),
				get("a", "A"),
				get("b", "B"),
			},
		},
		{
			Name: "RemoveSingle",
			Cap:  3,
			Ops: []Op{
				set("a", "A"),
				remove("a"),
				get("a", ""),
			},
		},
		{
			Name: "RemoveHead",
			Cap:  3,
			Ops: []Op{
				set("a", "A"),
				set("b", "B"),
				set("c", "C"),
				remove("c"),
				get("a", "A"),
				get("b", "B"),
				get("c", ""),
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			c := cache.NewLRU[string, string](tc.Cap)
			for _, op := range tc.Ops {
				op(t, c)
			}
		})
	}
}

func testHomework(t *testing.T) {
	r := require.New(t)

	c := cache.NewLRU[string, string](100)
	c.Set("Jesse", "Pinkman")
	c.Set("Walter", "White")
	c.Set("Jesse", "James")

	v, ok := c.Get("Jesse")
	r.True(ok)
	r.Equal("James", v)

	c.Remove("Walter")
	v, ok = c.Get("Walter")
	r.False(ok)
	r.Equal(v, "")
}
