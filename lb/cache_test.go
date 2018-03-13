package lb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRRCache(t *testing.T) {
	c := newRRCache()

	res := c.Next(testServiceName, testAddrs)

	assert.Contains(t, testAddrs[0], res)
}

func TestRRCacheOrder(t *testing.T) {
	c := newRRCache()

	res1 := c.Next(testServiceName, testAddrs)

	l := len(testAddrs)
	rev := make([]string, l)
	for i, x := range testAddrs {
		rev[l-i-1] = x
	}

	res2 := c.Next(testServiceName, rev)

	assert.Equal(t, testAddrs[0], res1)
	assert.Equal(t, testAddrs[1], res2)
}

func TestRRCacheAdditon(t *testing.T) {
	c := newRRCache()

	res1 := c.Next(testServiceName, testAddrs)

	l := len(testAddrs)
	rev := make([]string, l+1)
	for i, x := range testAddrs {
		rev[l-i-1] = x
	}
	rev[0] = "10.0.0.5"

	res2 := c.Next(testServiceName, rev)

	assert.Equal(t, testAddrs[0], res1)
	assert.Equal(t, testAddrs[1], res2)
}
