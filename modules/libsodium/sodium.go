package libsodium

/*
#cgo pkg-config: libsodium
#include <stdlib.h>
#include <sodium.h>
*/
import "C"

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GetRandomBytes(length int) (random_bytes []byte) {
	random_bytes = make([]byte, length)
	rand.Read(random_bytes)
	return random_bytes
}

func CryptoPwhash() ([]byte, []byte, []byte, []byte, []byte) {

	decryptMasterKeyStr()
	password := "somerandompassword"

	pwc := C.CString(password)
	hash := make([]byte, HASH_LENGTH)
	salt := GetRandomBytes(SALT_LENGTH)

	CryptoPWHashOpsLimitModerate := int(C.crypto_pwhash_opslimit_moderate())
	CryptoPWHashMemLimitModerate := int(C.crypto_pwhash_memlimit_moderate())
	CryptoPWHashAlgDefault := int(C.crypto_pwhash_alg_default())

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

	// authKey := GetRandomBytes(HASH_LENGTH)

	// ctx := C.CString(AUTH_CONTEXT)

	// C.crypto_kdf_derive_from_key((*C.uchar)(&authKey[0]), (C.ulong)(HASH_LENGTH), (C.ulonglong)(1),
	// 	ctx, (*C.uchar)(&hash[0]))

	masterKey := make([]byte, MASTERKEY_LENGTH)
	C.crypto_kdf_keygen((*C.uchar)(&masterKey[0]))
	fmt.Println(masterKey)
	// masterKey := GetRandomBytes(MASTERKEY_LENGTH)
	C.crypto_secretbox_keygen((*C.uchar)(&masterKey[0]))

	nonce := GetRandomBytes(NONCE_LENGTH)

	encMasterKey := make([]byte, MASTERKEY_LENGTH+CIPHER_LENGTH)

	C.crypto_secretbox_easy(
		(*C.uchar)(&encMasterKey[0]),
		(*C.uchar)(&masterKey[0]),
		C.ulonglong(MASTERKEY_LENGTH),
		(*C.uchar)(&nonce[0]), (*C.uchar)(&hash[0]))

	fmt.Println(encMasterKey)
	fmt.Println("Ciper: ", base64.RawURLEncoding.EncodeToString(encMasterKey))
	fmt.Println("key: ", base64.RawURLEncoding.EncodeToString(hash))
	fmt.Println("nonce: ", base64.RawURLEncoding.EncodeToString(nonce))

	decryptMasterKey(encMasterKey, nonce, hash)

	return salt, hash, masterKey, nonce, masterKey
}

func decryptMasterKey(encMasterKey []byte, nonce []byte, hash []byte) {

	//

	decryptedMsg := make([]byte, MASTERKEY_LENGTH)
	succes := C.crypto_secretbox_open_easy((*C.uchar)(&decryptedMsg[0]), (*C.uchar)(&encMasterKey[0]), (C.ulonglong)(len(encMasterKey)), (*C.uchar)(&nonce[0]), (*C.uchar)(&hash[0]))

	fmt.Println(succes)
	// fmt.Println(decryptedMsg)
	// fmt.Println("decrypted: ", base64.RawURLEncoding.EncodeToString(decryptedMsg))
}

func decryptMasterKeyStr() {

	encMasterKey := "4Q3SbwYjH0QGHXo6oChTKK4Pw2yGd-x5J3gKKX1VgHV3-hH2GEiiu6doq7U6gQ9h"
	nonce := "-e5vyl0-Mql1v2idkq8v8_A5xLfKT8kN"
	hash := "4e77dUDRf_SQgmR46sP2srrtwSqOGA8tm12R-r1WXyE"

	encBytes, _ := base64.RawURLEncoding.DecodeString(encMasterKey)
	nondeBytes, _ := base64.RawURLEncoding.DecodeString(nonce)
	authKeyBytes, _ := base64.RawURLEncoding.DecodeString(hash)

	decryptedMsg := make([]byte, MASTERKEY_LENGTH)
	_ = C.crypto_secretbox_open_easy((*C.uchar)(&decryptedMsg[0]), (*C.uchar)(&encBytes[0]), (C.ulonglong)(len(encBytes)), (*C.uchar)(&nondeBytes[0]), (*C.uchar)(&authKeyBytes[0]))

	// fmt.Println(succes)
	// fmt.Println(decryptedMsg)
	// fmt.Println("decrypted: ", base64.RawURLEncoding.EncodeToString(decryptedMsg))
}

// {salt: "TnSnrUueSQcL8vRAtkq2WA", nonce: "-e5vyl0-Mql1v2idkq8v8_A5xLfKT8kN", encMasterKey: "4Q3SbwYjH0QGHXo6oChTKK4Pw2yGd-x5J3gKKX1VgHV3-hH2GEiiu6doq7U6gQ9h", authKey: "ahiNAuKJeu71MXse7krWrMQfwAxq8qi-mgIiRPgYvjU"}
// salt: "TnSnrUueSQcL8vRAtkq2WA"
// nonce: "-e5vyl0-Mql1v2idkq8v8_A5xLfKT8kN"
// encMasterKey: "4Q3SbwYjH0QGHXo6oChTKK4Pw2yGd-x5J3gKKX1VgHV3-hH2GEiiu6doq7U6gQ9h"
// authKey: "ahiNAuKJeu71MXse7krWrMQfwAxq8qi-mgIiRPgYvjU"
