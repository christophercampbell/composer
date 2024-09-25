package composer

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapE(t *testing.T) {
	s := "5"

	result := MapE(MapE(ParseInt(s), Reciprocal).Right(), formatString) // Something is wrong here...
	// result := MapE(MapE(ParseInt(s), Reciprocal)) formatString)
	require.True(t, result.IsRight())
	require.NotNil(t, result.Right())
}

func TestFlatMapE(t *testing.T) {
	s := "5"
	result := FlatMapE(FlatMapE(ParseInt(s), Reciprocal), formatString)
	require.True(t, result.IsRight())
	require.NotNil(t, result.Right())
	require.Equal(t, "Result: 0.20", result.Right())

	s = "A"
	result = FlatMapE(FlatMapE(ParseInt(s), Reciprocal), formatString)
	require.True(t, result.IsLeft())
	require.NotNil(t, result.Left())
	require.Equal(t, `strconv.Atoi: parsing "A": invalid syntax`, result.Left().Error())
}

// ParseInt attempts to parse a string into an integer.
func ParseInt(s string) Either[error, int] {
	n, err := strconv.Atoi(s)
	if err != nil {
		return Left[error, int](err)
	}
	return Right[error, int](n)
}

// Reciprocal computes the reciprocal of a non-zero integer.
func Reciprocal(n int) Either[error, float64] {
	if n == 0 {
		return Left[error, float64](errors.New("division by zero"))
	}
	return Right[error, float64](1 / float64(n))
}

// FormatResult formats the float64 result into a string.
func formatString(f float64) Either[error, string] {
	return Right[error, string](fmt.Sprintf("Result: %.2f", f))
}
