package cmd

/*
EncoDec is intended to work with only floats. It encodes and decodes floats to a binary representation.
*/
type EncoDec interface {
	Encoder
	Decoder
}

type Encoder interface {
	Encode(v float64, b []byte)
}

type Decoder interface {
	Decode(b []byte) (v float64)
}

func main() {

}
