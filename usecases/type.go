package usecases

type Command[T any] interface {
	Validate() error
	Execute() (T, error)
}

type Query[T any] interface {
	Validate() error
	Execute() (T, error)
}
