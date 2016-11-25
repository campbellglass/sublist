package main

import (
	"regexp"

	"github.com/valyala/fasthttp"
)

type Endpoint struct {
	Regexp  *regexp.Regexp
	Handler func(*fasthttp.RequestCtx)
}

func NewEndpoint(rgx string, handler func(*fasthttp.RequestCtx)) *Endpoint {
	endpoint := Endpoint{
		Regexp:  regexp.MustCompile(rgx),
		Handler: handler,
	}
	return &endpoint
}
