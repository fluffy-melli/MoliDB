package handlers

import (
	"os"

	"github.com/fluffy-melli/MoliDB/pkg/crypto"
	"github.com/fluffy-melli/MoliDB/pkg/gzip"
)

func CryptoEncrypt(store []byte) (string, error) {
	body, err := gzip.Compress(store)
	if err != nil {
		return "", err
	}
	b, err := crypto.AESEncrypt(os.Getenv("SECRET_KEY"), string(body))
	if err != nil {
		return "", err
	}
	return b, nil
}

func CryptoDecrypt(data string) ([]byte, error) {
	a, err := crypto.AESDecrypt(os.Getenv("SECRET_KEY"), data)
	if err != nil {
		return nil, err
	}
	ab, err := gzip.Decompress([]byte(a))
	if err != nil {
		return nil, err
	}
	return ab, nil
}
