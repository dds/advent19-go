package aoc19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	t.Run("1+1=2", func(t *testing.T) {
		program := []int{1, 0, 0, 0, 99}
		data := Run(program)
		require.Equal(t, data[0], 2)
	})

	t.Run("3 * 2 = 6", func(t *testing.T) {
		program := []int{2, 3, 0, 3, 99}
		data := Run(program)
		require.Equal(t, data[3], 6)
	})

	t.Run("99 * 99 = 9801", func(t *testing.T) {
		program := []int{2, 4, 4, 5, 99, 9801}
		data := Run(program)
		require.Equal(t, data[5], 9801)
	})

	t.Run("test4", func(t *testing.T) {
		program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		data := Run(program)
		require.Equal(t, data, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
	})
}
