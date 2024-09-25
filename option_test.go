package composer

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipelining(t *testing.T) {
	s := "5"

	// Chain operations using FlatMap functions.
	// Since we cannot have methods with type parameters, we use functions.
	result := FlatMap(parseInt(s), func(n int) Option[string] {
		return FlatMap(reciprocal(n), func(r float64) Option[string] {
			return someOtherOperation(r)
		})
	})

	require.True(t, result.isDefined)
	require.NotNil(t, result.Get())
}

func parseInt(s string) Option[int] {
	n, err := strconv.Atoi(s)
	if err != nil {
		return None[int]()
	}
	return Some(n)
}

func reciprocal(n int) Option[float64] {
	if n == 0 {
		return None[float64]()
	}
	return Some(1 / float64(n))
}

func someOtherOperation(f float64) Option[string] {
	return Some(fmt.Sprintf("Result: %.2f", f))
}
