package libsodium

/*
#cgo pkg-config: libsodium
#include <stdlib.h>
#include <sodium.h>
*/
import "C"

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/cristalhq/acmd"
)

func main() {
	cmds := []acmd.Command{
		{
			Name:        "hash",
			Description: "hashes the provided password using argon2id",
			ExecFunc:    HashCommand,
		},
		{
			Name:        "status",
			Description: "prints status of the system",
			ExecFunc: func(ctx context.Context, args []string) error {
				fmt.Println("status")
				return nil
			},
		},
	}

	// all the acmd.Config fields are optional
	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName:        "acmd-example",
		AppDescription: "Example of acmd package",
		Version:        "the best v0.x.y",
		// Context - if nil `signal.Notify` will be used
		// Args - if nil `os.Args[1:]` will be used
		// Usage - if nil default print will be used
	})

	if err := r.Run(); err != nil {
		r.Exit(err)
	}

	// salt, hash := CryptoPwhash(os.Args[1])

	// fmt.Println(os.Args[1])
	// fmt.Println(base64.RawURLEncoding.EncodeToString(salt))
	// fmt.Println(hash)
	// fmt.Println(base64.RawURLEncoding.EncodeToString(hash))
}

func HashCommand(ctx context.Context, args []string) error {
	salt, hash := CryptoPwhash1(args[0])
	fmt.Println("Salt:", base64.RawURLEncoding.EncodeToString(salt))
	fmt.Println("Hash:", base64.RawURLEncoding.EncodeToString(hash))
	return nil
}

func GetRandomBytes1(length int) (random_bytes []byte) {
	random_bytes = make([]byte, length)
	rand.Read(random_bytes)
	return random_bytes
}

func CryptoPwhash1(password string) (salt, hash []byte) {
	HASH_LENGTH := 32
	SALT_LENGTH := 16
	CryptoPWHashOpsLimitModerate := int(C.crypto_pwhash_opslimit_moderate())
	CryptoPWHashMemLimitModerate := int(C.crypto_pwhash_memlimit_moderate())
	CryptoPWHashAlgDefault := int(C.crypto_pwhash_alg_default())

	pwc := C.CString(password)
	hash = make([]byte, HASH_LENGTH)
	salt = GetRandomBytes(SALT_LENGTH)

	C.crypto_pwhash(
		(*C.uchar)(&hash[0]),
		C.ulonglong(HASH_LENGTH),
		pwc,
		C.ulonglong(len(password)),
		(*C.uchar)(&salt[0]),
		C.ulonglong(CryptoPWHashOpsLimitModerate),
		C.ulong(CryptoPWHashMemLimitModerate),
		C.int(CryptoPWHashAlgDefault),
	)

	return salt, hash
}
