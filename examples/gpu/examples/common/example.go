package common

type ExampleInterface interface {
	String() string
	Init(context *Context) error
	Update(context *Context) error
	Draw(context *Context) error
	Quit(context *Context)
}
