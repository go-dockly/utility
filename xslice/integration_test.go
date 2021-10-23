package xslice_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-dockly/utility/xslice"
	"github.com/stretchr/testify/assert"
)

var syncslice = xslice.NewSyncSlice([]interface{}{"one", "two"}...)

func TestIntegrationShift(t *testing.T) {
	var res = <-syncslice.Shift()
	spew.Dump(res)
	assert.Equal(t, res.Val, "one")
}

func TestIntegrationInsert(t *testing.T) {
	syncslice.Append("three", "four")
	assert.GreaterOrEqual(t, syncslice.Len(), 3)
}
