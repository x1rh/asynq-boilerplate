package asynqctx

import "asynq-boilerplate/config"

type Context struct {
	Config config.Config
}

func New(c config.Config) *Context {
	return &Context{
		Config: c,
	}
}
