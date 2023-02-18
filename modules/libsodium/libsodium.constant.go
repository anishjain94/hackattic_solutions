package libsodium

/*
#cgo pkg-config: libsodium
#include <stdlib.h>
#include <sodium.h>
*/
import "C"

const (
	HASH_LENGTH       int = 32
	SALT_LENGTH       int = 16
	NONCE_LENGTH      int = 24
	MASTERKEY_LENGTH  int = 32
	KDF_CONTEXT_BYTES int = 8
	CIPHER_LENGTH     int = 16
)

var AUTH_CONTEXT string = "AUTH_KEY"

const (
	crypto_secretstream_xchacha20poly1305_KEYBYTES    int = 100
	crypto_secretstream_xchacha20poly1305_HEADERBYTES int = 100
	crypto_secretstream_xchacha20poly1305_ABYTES      int = 100
)
