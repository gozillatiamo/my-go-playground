package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	privateKey, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(privateKey)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	fmt.Println(key)

	bBody, err := ioutil.ReadFile("body.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bBody))

	nPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err.Error)
	}

	nPublicKey := nPrivateKey.PublicKey
	fmt.Println(nPublicKey)

	signature := SignPKCS1v15(string(bBody), *nPrivateKey)
	fmt.Println(string(signature))
	if err != nil {
		panic(err)
	}

	publicKey, err := ioutil.ReadFile("public_key.pem")

	block, _ = pem.Decode(publicKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)

	rsaPublickey, _ := pub.(*rsa.PublicKey)

	fmt.Println(rsaPublickey)

	Verify(&nPublicKey, bBody, signature)
}

func SignPKCS1v15(plaintext string, privKey rsa.PrivateKey) string {

	rng := rand.Reader
	hashed := HashSha256(plaintext)
	signature, err := rsa.SignPKCS1v15(rng, &privKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return "Error from signing"
	}
	return base64.RawURLEncoding.EncodeToString(signature)
}

func Verify(public *rsa.PublicKey, val []byte, signature string) {
	b, err := base64.RawURLEncoding.DecodeString(signature)
	fmt.Println("Decoded: ", string(b))

	hashed := HashSha256(string(val))
	err = rsa.VerifyPKCS1v15(public, crypto.SHA256, hashed[:], b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Verified")
}

func HashSha256(content string) []byte {
	data := []byte(content)

	h := sha256.New()

	h.Write(data)

	digest := h.Sum(nil)

	return digest
	// return sha256.Sum256(data)
}
