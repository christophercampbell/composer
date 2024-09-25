package composer

// Either represents a value of one of two possible types (a disjoint union).
type Either[L any, R any] struct {
	isRight bool // Indicates if the Either contains a Right value.
	left    L    // Holds the Left value.
	right   R    // Holds the Right value.
}

// Left creates an Either that contains a Left value.
func Left[L any, R any](value L) Either[L, R] {
	var zeroR R
	return Either[L, R]{isRight: false, left: value, right: zeroR}
}

// Right creates an Either that contains a Right value.
func Right[L any, R any](value R) Either[L, R] {
	var zeroL L
	return Either[L, R]{isRight: true, left: zeroL, right: value}
}

// IsLeft checks if the Either is a Left value.
func (e Either[L, R]) IsLeft() bool {
	return !e.isRight
}

// IsRight checks if the Either is a Right value.
func (e Either[L, R]) IsRight() bool {
	return e.isRight
}

// Left returns the Left value.
func (e Either[L, R]) Left() L {
	return e.left
}

// Right returns the Right value.
func (e Either[L, R]) Right() R {
	return e.right
}

// Map applies a function to the Right value, if it exists.
// Since methods cannot have type parameters, Map is defined as a function.
func MapE[L any, R any, R2 any](e Either[L, R], f func(R) R2) Either[L, R2] {
	if e.IsRight() {
		return Right[L, R2](f(e.Right()))
	}
	return Left[L, R2](e.Left())
}

// FlatMap chains computations, applying the function to the Right value.
// Defined as a function due to Go's language constraints.
func FlatMapE[L any, R any, R2 any](e Either[L, R], f func(R) Either[L, R2]) Either[L, R2] {
	if e.IsRight() {
		return f(e.Right())
	}
	return Left[L, R2](e.Left())
}
