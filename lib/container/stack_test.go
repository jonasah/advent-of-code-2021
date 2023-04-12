package container_test

import (
	"testing"

	"github.com/jonasah/advent-of-code-2021/lib/container"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := container.NewStack[int]()
	assert.Equal(t, 0, s.Len())

	s.Push(1337)
	s.Push(42)
	assert.Equal(t, 2, s.Len())

	v, err := s.Pop()
	assert.Equal(t, 42, *v)
	assert.Nil(t, err)
	assert.Equal(t, 1, s.Len())

	s.Pop()
	assert.Equal(t, 0, s.Len())

	v, err = s.Pop()
	assert.Error(t, err)
	assert.Nil(t, v)
}
