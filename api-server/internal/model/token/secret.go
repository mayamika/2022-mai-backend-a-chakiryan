package token

import "crypto/rand"

const hmacSecretBytes = 4096

var hmacSecret []byte

func init() {
	hmacSecret = make([]byte, hmacSecretBytes)
	_, err := rand.Read(hmacSecret)
	if err != nil {
		panic(err)
	}
}
