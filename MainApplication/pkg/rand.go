package pkg

import (
	"crypto/rand"
	"math/big"
)

func GenId() (uint64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		return 0, err
	}
	return nBig.Uint64(), nil
}
