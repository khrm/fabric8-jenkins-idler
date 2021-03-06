package openshift

import (
	"strconv"
	"sync"
	"testing"

	"github.com/fabric8-services/fabric8-jenkins-idler/internal/idler"
	"github.com/stretchr/testify/assert"
)

func Test_retrieve_idler_for_unknown_namespace_returns_nil(t *testing.T) {
	m := NewUserIdlerMap()
	idler, ok := m.Load("foo")
	assert.False(t, ok, "There should be no entry mapped")
	assert.Nil(t, idler, "No reference should be returned")
}

func Test_retrieve_idler_for_empty_namespace_returns_nil(t *testing.T) {
	m := NewUserIdlerMap()
	idler, ok := m.Load("")
	assert.False(t, ok, "There should be no entry mapped")
	assert.Nil(t, idler, "No reference should be returned")
}

func Test_len(t *testing.T) {
	m := NewUserIdlerMap()
	assert.Equal(t, 0, m.Len(), "Empty map should return 0 len")
	m.Store("foo", &idler.UserIdler{})
	assert.Equal(t, 1, m.Len(), "Len should return number of items")
}

func Test_len_mutli(t *testing.T) {
	m := NewUserIdlerMap()
	assert.Equal(t, 0, m.Len(), "Empty map should return 0 len")

	const n = 50
	wg := &sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(x int) {
			defer wg.Done()

			key := strconv.Itoa(x)
			m.Store(key, &idler.UserIdler{})
		}(i)
	}

	wg.Wait()
	assert.Equal(t, n, m.Len(), "Len should return number of items")
}
