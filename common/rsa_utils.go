package service

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "errors"
    "fmt"
)

func GenRSAKey() (prvkey, pubkey []byte, keyLength int) {
    privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
    if err != nil {
        panic(err)
    }
    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    }
    prvkey = pem.EncodeToMemory(block)
    publicKey := &privateKey.PublicKey
    derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        panic(err)
    }
    block = &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    }
    pubkey = pem.EncodeToMemory(block)
    return
}

func RSASignWithSha256(data []byte, keyBytes []byte) []byte {
    h := sha256.New()
    h.Write(data)
    hashed := h.Sum(nil)
    block, _ := pem.Decode(keyBytes)
    if block == nil {
        panic(errors.New("private key error"))
    }
    privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        fmt.Println("ParsePKCS8PrivateKey err", err)
        panic(err)
    }

    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
    if err != nil {
        fmt.Printf("Error from signing: %s\n", err)
        panic(err)
    }

    return signature
}

func RSAVerifySignWithSha256(data, signData, keyBytes []byte) bool {
    block, _ := pem.Decode(keyBytes)
    if block == nil {
        panic(errors.New("public key error"))
    }
    pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        panic(err)
    }

    hashed := sha256.Sum256(data)
    err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
    if err != nil {
        panic(err)
    }
    return true
}

func RSAEncrypt(data, keyBytes []byte) []byte {
    block, _ := pem.Decode(keyBytes)
    if block == nil {
        panic(errors.New("public key error"))
    }
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        panic(err)
    }
    pub := pubInterface.(*rsa.PublicKey)
    ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
    if err != nil {
        panic(err)
    }
    return ciphertext
}

func RSADecrypt(ciphertext, keyBytes []byte) []byte {
    block, _ := pem.Decode(keyBytes)
    if block == nil {
        panic(errors.New("private key error!"))
    }
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        panic(err)
    }
    data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
    if err != nil {
        panic(err)
    }
    return data
}