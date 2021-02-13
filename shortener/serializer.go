package shortener

type Serializeer interface {
	Encode(input *Redirect) ([]byte, error)
	Decode(input []byte) (*Redirect, error)
}