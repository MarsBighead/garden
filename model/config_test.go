package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssign(t *testing.T) {
	var p Pointor
	assign(&p)
	assert.Equal(t, 1, p.num, "Num in Pointor error")
}
