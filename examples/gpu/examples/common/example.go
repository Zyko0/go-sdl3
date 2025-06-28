package common

type Example struct {
	Name   string
	Init   func(context *Context) error
	Update func(context *Context) error
	Draw   func(context *Context) error
	Quit   func(context *Context)
}
