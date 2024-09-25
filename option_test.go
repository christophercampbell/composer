package composer

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	optNumber := None[int]()
	twiceOpt := Map(optNumber, doubleValue)
	require.False(t, twiceOpt.isDefined)

	optNumber = Some(5)
	twiceOpt = Map(optNumber, doubleValue)
	require.True(t, twiceOpt.isDefined)
	require.Equal(t, 10, twiceOpt.value)
}

func doubleValue(n int) int {
	return n * 2
}

func TestFlatMap(t *testing.T) {
	s := "5"
	result := FlatMap(FlatMap(parseInt(s), reciprocal), someOtherOperation)
	require.True(t, result.isDefined)
	require.NotNil(t, result.Get())
	require.Equal(t, "Result: 0.20", result.Get())
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
