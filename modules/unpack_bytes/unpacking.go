package unpackbytes

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
)

func Unpack() {

	resp, err := http.Get("https://hackattic.com/challenges/help_me_unpack/problem?access_token=8e80fec0cbe25049")
	if err != nil {

	}

	readBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	var anyJson map[string]interface{}
	json.Unmarshal(readBytes, &anyJson)

	stringReceived := anyJson["bytes"].(string)

	bytesToExtract, err := base64.StdEncoding.DecodeString(stringReceived)

	if err != nil {
	}

	integer := int32(binary.LittleEndian.Uint32(bytesToExtract[0:4]))
	println(integer)

	unsignedInteger := uint32(binary.LittleEndian.Uint32(bytesToExtract[4:8]))
	println(unsignedInteger)

	signedShort := int16(binary.LittleEndian.Uint16(bytesToExtract[8:12]))
	println(signedShort)

	signedFloatTemp := binary.LittleEndian.Uint32(bytesToExtract[12:16])
	signedFloat := float64(math.Float32frombits(signedFloatTemp))
	println(signedFloat)

	signedDoubleTemp := binary.LittleEndian.Uint64(bytesToExtract[16:24])
	signedDouble := math.Float64frombits(signedDoubleTemp)
	println(signedDouble)

	signedDoubleBigTemp := binary.BigEndian.Uint64(bytesToExtract[24:32])
	signedBigDouble := math.Float64frombits(signedDoubleBigTemp)
	println(signedBigDouble)

	type Solution struct {
		Int               int32   `json:"int"`
		Uint              uint32  `json:"uint"`
		Short             int16   `json:"short"`
		Float             float64 `json:"float"`
		Double            float64 `json:"double"`
		Big_endian_double float64 `json:"big_endian_double"`
	}

	solution := Solution{integer, unsignedInteger, signedShort, signedFloat, signedDouble, signedBigDouble}

	solutionBytes, err := json.Marshal(solution)
	if err != nil {
	}

	bytesReader := bytes.NewReader(solutionBytes)

	resp, err = http.Post("https://hackattic.com/challenges/help_me_unpack/solve?access_token=8e80fec0cbe25049&playground=1", "application/json", bytesReader)
	readBytes, err = ioutil.ReadAll(resp.Body)
	println(string(readBytes))
}
