package json

import (
	"encoding/json"
	"project5/shortener"
	errs "github.com/pkg/errors"
)

type Redirect struct {}

func (r Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	encoded, err := json.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.json.Encode")
	}
	return encoded, nil
}

func (r Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	err := json.Unmarshal(input, redirect)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.json.Decode")
	}
	return redirect, nil
}