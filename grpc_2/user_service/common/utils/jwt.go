package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type headerData struct {
	Algorythm string `json:"alg"`
	Type string `json:"typ"`
}

var loadKeysOnce sync.Once
var privateKey *rsa.PrivateKey = nil
var publicKey *rsa.PublicKey = nil
func loadKeys(){
	// -------------------- private key --------------------
	data, err := os.ReadFile("private_key.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	block, _ := pem.Decode(data)
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	castedPrivateKeyKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok{
		fmt.Println("could not cast private key")
		return
	}
	privateKey = castedPrivateKeyKey

	// -------------------- public key --------------------
	data, err = os.ReadFile("public_key.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	block, _ = pem.Decode(data)
	parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	castedPublicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok{
		fmt.Println("could not cast public key")
		privateKey = nil
		return
	}
	publicKey = castedPublicKey
}

func EncodeJWT(data any) (string, error) {
	var token string
	loadKeysOnce.Do(loadKeys)
	if privateKey == nil {
		return token, errors.New("could not load keys")
	}

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
	alg := "RS256"

	headerdata := headerData{
		Algorythm: alg,
		Type: "JWT",
	}
	headerJSON, err := json.Marshal(headerdata)
	if err != nil{
		return encodedHeader, errors.Join(err, errors.New("could not get a json object from given data"))
	}
	encodedHeader = base64.URLEncoding.EncodeToString(headerJSON)

	return encodedHeader, nil
}

func encodePayload(data any) (string, error) {
	var encodedPayload string

	jsonString, err := json.Marshal(data)
	if err != nil {
		return encodedPayload, errors.Join(err, errors.New("could not get a json object from given data"))
	}

	encodedPayload = base64.URLEncoding.EncodeToString(jsonString)
	return encodedPayload, nil
}

func encodeSignature(header, payload string) (string, error) {
	var signature string
	hash := sha256.Sum256([]byte(header + "." + payload))
	sign, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		hash[:],
	)
	if err != nil {
		return signature, err
	}
	signature = base64.URLEncoding.EncodeToString(sign)

	return signature, nil
}

func VerifySignature(token string) error {
	loadKeysOnce.Do(loadKeys)
	if publicKey == nil {
		return errors.New("could not load keys")
	}

	splitted := strings.Split(token, ".")
	if len(splitted) != 3{
		return errors.New("incorrect token")
	}
	signature, err := base64.URLEncoding.DecodeString(splitted[2])
	if err != nil{
		return err
	}
	hash := sha256.Sum256([]byte(splitted[0] + "." + splitted[1]))

	return rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hash[:],
		signature,
	)
}

func DecodeJWT[T any](token string) (T, error) {
	var result T
	err := VerifySignature(token)
	if err != nil{
		return result, err
	}

	splitted := strings.Split(token, ".")
	if len(splitted) != 3{
		return result, errors.New("incorrect token")
	}

	payload := splitted[1]

	decodedPayload, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		return result, errors.New("could not decode payload")
	}
	err = json.Unmarshal(decodedPayload, &result)
	if err != nil {
		return result, errors.Join(err, errors.New("could not parse json"))
	}

	return result, nil
}
