package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"os"
	"strings"

)

type headerData struct {
	Algorythm string `json:"alg"`
	Type string `json:"typ"`
}

func EncodeJWT(data any) (string, error) {
	var token string
	header, err := encodedHeader()
	if err != nil {
		return token, errors.Join(err, errors.New("could not get jwt header"))
	}

	payload, err := encodePayload(data)
	if err != nil {
		return token, errors.Join(err, errors.New("could not encode jwt payload"))
	}

	signature, err := encodeSignature(header, payload)
	if err != nil {
		return token, errors.Join(err, errors.New("could not get jwt signature"))
	}

	token = strings.Join([]string{header, payload, signature}, ".")

	return token, nil
}

func encodedHeader() (string, error){
	var encodedHeader string
	alg, envWasSet := os.LookupEnv("JWT_ALG")
	if !envWasSet{
		// fallback if env wasnt set. ONLY FOR TESTING!!!
		alg = "HS256"
	}

	headerdata := headerData{
		Algorythm: alg,
		Type: "JWT",
	}
	headerJSON, err := json.Marshal(headerdata)
	if err != nil{
		return encodedHeader, errors.Join(err, errors.New("could not get a json object from given data"))
	}
	encodedHeader = base64.StdEncoding.EncodeToString(headerJSON)

	return encodedHeader, nil
}

func encodePayload(data any) (string, error) {
	var encodedPayload string

	jsonString, err := json.Marshal(data)
	if err != nil {
		return encodedPayload, errors.Join(err, errors.New("could not get a json object from given data"))
	}

	encodedPayload = base64.StdEncoding.EncodeToString(jsonString)
	return encodedPayload, nil
}

func encodeSignature(header, payload string) (string, error) {
	var encodedSignature string
	// ideally i should read env vars in init, and then store them in memory
	secret, envWasSet := os.LookupEnv("JWT_SECRET")
	if !envWasSet{
		// fallback if env wasnt set. ONLY FOR TESTING!!!
		secret = "4l4MYgc4hjAtUK0INR2xfO7kDXJieWub4JNkPcqQ4S5g7uoYgGAtE52cqVs5DwGm"
	}

	var alg string

	decodedHeader, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		alg, envWasSet = os.LookupEnv("JWT_ALG")
		if !envWasSet{
			// fallback if env wasnt set. ONLY FOR TESTING!!!
			alg = "HS256"
		}
	} else {
		var headerdata headerData
		if err = json.Unmarshal(decodedHeader, &headerdata); err != nil {
			return encodedSignature, errors.Join(err, errors.New("could not read header data"))
		}
		alg = headerdata.Algorythm
	}

	var hasher hash.Hash
	switch alg{
		case "HS256":
		hasher = hmac.New(sha256.New, []byte(secret))
		default:
		return encodedSignature, fmt.Errorf("algorythm %q is unsupported", alg)
	}

	hasher.Write([]byte(header))
	hasher.Write([]byte("."))
	hasher.Write([]byte(payload))

	signature := hasher.Sum(nil)
	encodedSignature = base64.StdEncoding.EncodeToString(signature)

	return encodedSignature, nil
}

func DecodeJWT[T any](token string) (T, error) {
	var result T
	splitted := strings.Split(token, ".")
	if len(splitted) != 3{
		return result, errors.New("incorrect token")
	}
	header := splitted[0]
	payload := splitted[1]
	signature := splitted[2]

	expectedSignature, err := encodeSignature(header, payload)
	if err != nil {
		return result, errors.Join(err, errors.New("could not get jwt signature to check"))
	}
	if signature != expectedSignature {
		return result, errors.New("incorrect JWT signature")
	}

	decodedPayload, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return result, errors.New("could not decode payload")
	}
	err = json.Unmarshal(decodedPayload, &result)
	if err != nil {
		return result, errors.Join(err, errors.New("could not parse json"))
	}

	return result, nil
}
