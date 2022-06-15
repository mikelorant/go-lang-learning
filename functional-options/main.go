package main

import (
	"fmt"
	"github.com/mikelorant/go-lang-learning-functional-options/example"
	"github.com/mikelorant/go-lang-learning-functional-options/option"
)

type Option interface{
	Apply(*option.Options)
}

type WithAPIKey string
func (k WithAPIKey) Apply(o *option.Options) {
	o.APIKey = string(k)
}

type WithAPISecret string
func (s WithAPISecret) Apply(o *option.Options) {
	o.APISecret = string(s)
}

type WithBaseURL string
func (u WithBaseURL) Apply(o *option.Options) {
	o.BaseURL = string(u)
}

func main() {
	New(
		WithAPIKey("key"),
		WithAPISecret("secret"),
	)
}

func New(opts ...Option) {
	o := &option.Options{}

	for _, opt := range opts {
		opt.Apply(o)
	}

	fmt.Printf("%+v\n", o)

	example.New(o, WithBaseURL("http://www.example.org"))
}
