package example

import (
	"fmt"
	"github.com/mikelorant/go-lang-learning-functional-options/option"
)

type Option interface{ apply(*option.Options) }

func New(o *option.Options, opts ...Option) {
	for _, opt := range opts {
		opt.apply(o)
	}

	fmt.Printf("%+v\n", o)
}
