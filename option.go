package composer

// Option represents a value that may or may not be present.
// It is the Go equivalent of Scala's Option[T] monad.
type Option[T any] struct {
	isDefined bool // Indicates whether the Option contains a value.
	value     T    // The value contained within the Option, if any.
}

// Some creates an Option that contains a value.
// Equivalent to Scala's Some(value).
func Some[T any](value T) Option[T] {
	return Option[T]{isDefined: true, value: value}
}

// None creates an Option that contains no value.
// Equivalent to Scala's None.
func None[T any]() Option[T] {
	var zero T
	return Option[T]{isDefined: false, value: zero}
}

// IsDefined checks whether the Option contains a value.
// Returns true if the Option contains a value, false otherwise.
func (o Option[T]) IsDefined() bool {
	return o.isDefined
}

// Get retrieves the value contained in the Option.
// Should only be called if IsDefined() returns true.
func (o Option[T]) Get() T {
	return o.value
}

// FlatMap applies a function to the Option's value if it is defined,
// returning a new Option. If the Option is not defined, it returns None.
// Since methods cannot have type parameters, FlatMap is defined as a function.
func FlatMap[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.isDefined {
		return f(o.value)
	}
	return None[U]()
}

// Map applies a function to the Option's value if it is defined,
// wrapping the result in a new Option. If the Option is not defined, it returns None.
// Map is also defined as a function for the same reason.
func Map[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.isDefined {
		return Some(f(o.value))
	}
	return None[U]()
}
