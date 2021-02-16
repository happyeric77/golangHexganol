package msgpack

import (
	"project5/shortener"
	errs "github.com/pkg/errors"
	msgpack "github.com/vmihailenco/msgpack/v5"
)

type Redirect struct {}

func (r Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	encoded, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.msgpack.Encode")
	}
	return encoded, nil
}

func (r Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	err := msgpack.Unmarshal(input, redirect)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.msgpack.Decode")
	}
	return redirect, nil
}