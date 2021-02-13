package main

import (
	"fmt"
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

func main() {
	e := errors.New("ERROR original")
	wrapedE := errs.Wrap(e, "additional info")
	fmt.Println(e)
	fmt.Println(wrapedE)
	fmt.Println(shortid.MustGenerate())
}