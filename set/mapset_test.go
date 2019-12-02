package set

import (
	"github.com/deckarep/golang-set"
	"gotest.tools/assert"
	"testing"
)

func TestMapSet(t *testing.T) {
	ms := mapset.NewSet(1, 2, 3)
	ms.Add(3)
	ms.Add(4)
	assert.Assert(t, ms.Equal(mapset.NewSetFromSlice([]interface{}{1, 2, 3, 4})))
	assert.Assert(t, mapset.NewSet(1, 2, 3).Union(mapset.NewSet(3, 4, 5)).Equal(mapset.NewSet(1, 2, 3, 4, 5)))
}
