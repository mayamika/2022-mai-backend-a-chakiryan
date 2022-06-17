package token

const hmacSecretBytes = 4096

var hmacSecret []byte

func init() {
	// TODO: Yes.
	hmacSecret = make([]byte, hmacSecretBytes)
}
