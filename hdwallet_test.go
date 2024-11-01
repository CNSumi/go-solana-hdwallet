package hdwallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// DO NOT USE THIS MNEMONIC IN PRODUCT ENVIRONMENT!
const testMnemonic = "beach liar addict wrap pause runway evolve front grab print jungle mimic"

func TestNewNode(t *testing.T) {
	type Case struct {
		opts      []NewNodeOption
		secretKey string
		address   string
	}

	cases := []*Case{
		{
			[]NewNodeOption{},
			"2rh6oFzfX5efaXPRmVSz1U4DAMaaum8PKBD8x72ssUTCpx1AoEGCjW5CYGQZ83PNLVBQFjvY6K5i6QFMMbi2dMLZ",
			"HDEXA1PAaHBnhjvrqc7vr1GrfV2qtpAjaYmRANj97zCV",
		},
		{
			[]NewNodeOption{WithIndex(1)},
			"4S51QvMwJoD8U6eejyYn9SAtfm3a7jWBR9S4ftZfdW7XswPJ614fYyqYEjd3RnmHGHTVpFPEYcV62Zo7HC28opQp",
			"EZJsGZAdqZB8oHtKjVkPWaduvRcNUov2oJgSkNqkZmiQ",
		},
	}

	for _, c := range cases {
		node := NewNode(testMnemonic, c.opts...)
		assert.Equal(t, node.SecretKey(), c.secretKey)
		assert.Equal(t, node.Address(), c.address)
	}
}
