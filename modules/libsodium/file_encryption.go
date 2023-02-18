package libsodium

/*
#cgo pkg-config: libsodium
#include <stdlib.h>
#include <sodium.h>
*/
import "C"

import "fmt"

func FileEncrypt() {
	msg1 := "Arbitrary data to encrypt"
	msg2 := "splitInto"
	msg3 := "3 messages"

	message1 := ToUnsignedChar([]byte(msg1))
	message2 := ToUnsignedChar([]byte(msg2))
	message3 := ToUnsignedChar([]byte(msg3))

	CIPHERTEXT_PART1_LEN := len(msg1) + crypto_secretstream_xchacha20poly1305_ABYTES
	CIPHERTEXT_PART2_LEN := len(msg2) + crypto_secretstream_xchacha20poly1305_ABYTES
	CIPHERTEXT_PART3_LEN := len(msg3) + crypto_secretstream_xchacha20poly1305_ABYTES

	key := GetRandomBytes(crypto_secretstream_xchacha20poly1305_KEYBYTES)
	header := MakeByte(crypto_secretstream_xchacha20poly1305_HEADERBYTES)

	state := C.crypto_secretstream_xchacha20poly1305_state{}

	cypherText1 := MakeByte(CIPHERTEXT_PART1_LEN)
	cypherText2 := MakeByte(CIPHERTEXT_PART2_LEN)
	cypherText3 := MakeByte(CIPHERTEXT_PART3_LEN)

	C.crypto_secretstream_xchacha20poly1305_keygen(ToUnsignedChar(key))

	C.crypto_secretstream_xchacha20poly1305_init_push(&state, ToUnsignedChar(header), ToUnsignedChar(key))

	C.crypto_secretstream_xchacha20poly1305_push(&state, ToUnsignedChar(cypherText1), nil, message1, ToUnsignedLongLong(len(msg1)), nil, 0, 0)
	C.crypto_secretstream_xchacha20poly1305_push(&state, ToUnsignedChar(cypherText2), nil, message2, ToUnsignedLongLong(len(msg2)), nil, 0, 0)
	C.crypto_secretstream_xchacha20poly1305_push(&state, ToUnsignedChar(cypherText3), nil, message3, ToUnsignedLongLong(len(msg3)), nil, 0, 0)

	fmt.Println(ToString(cypherText1))
	fmt.Println(ToString(cypherText2))
	fmt.Println(ToString(cypherText3))

	FileDecrypt(header, key, cypherText1, cypherText2, cypherText3)
}

func FileDecrypt(header []byte, key []byte, cypherText1 []byte, cypherText2 []byte, cypherText3 []byte) {

	tag := MakeByte(100)
	state := C.crypto_secretstream_xchacha20poly1305_state{}

	msg1 := MakeByte(100)
	msg2 := MakeByte(100)
	msg3 := MakeByte(100)

	if C.crypto_secretstream_xchacha20poly1305_init_pull(&state, ToUnsignedChar(header), ToUnsignedChar(key)) != 0 {
		panic("invalid header")
	}

	if C.crypto_secretstream_xchacha20poly1305_pull(&state, ToUnsignedChar(msg1), nil, ToUnsignedChar(tag), ToUnsignedChar(cypherText1), ToUnsignedLongLong(crypto_secretstream_xchacha20poly1305_KEYBYTES), nil, 0) != 0 {
		panic("cannot decrypt")
	}

	if C.crypto_secretstream_xchacha20poly1305_pull(&state, ToUnsignedChar(msg2), nil, ToUnsignedChar(tag), ToUnsignedChar(cypherText1), ToUnsignedLongLong(crypto_secretstream_xchacha20poly1305_KEYBYTES), nil, 0) != 0 {
		panic("cannot decrypt")
	}

	if C.crypto_secretstream_xchacha20poly1305_pull(&state, ToUnsignedChar(msg3), nil, ToUnsignedChar(tag), ToUnsignedChar(cypherText1), ToUnsignedLongLong(crypto_secretstream_xchacha20poly1305_KEYBYTES), nil, 0) != 0 {
		panic("cannot decrypt")
	}
}
