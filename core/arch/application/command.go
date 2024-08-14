package application

type ICommand interface {
	Validate() error
}
