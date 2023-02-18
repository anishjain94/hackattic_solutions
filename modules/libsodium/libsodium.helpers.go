package libsodium

import "C"
import "encoding/base64"

func ToUnsignedChar(str []byte) *C.uchar {
	return (*C.uchar)(&str[0])
}

func MakeByte(len int) []byte {
	return make([]byte, HASH_LENGTH)
}

func ToUnsignedLongLong(n int) C.ulonglong {
	return C.ulonglong(n)
}

func ToString(str []byte) string {
	return base64.RawURLEncoding.EncodeToString(str)
}
