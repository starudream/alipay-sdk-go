package alipaysdk

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

func RsaSign(content, privateKeyString string, hash crypto.Hash) (string, error) {
	h := crypto.Hash.New(hash)
	h.Write([]byte(content))
	contentBytes := h.Sum(nil)

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyString)
	if err != nil {
		return "", errors.New("decode private key fail")
	}

	var privateKey *rsa.PrivateKey
	if hash == crypto.SHA1 {
		privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBytes)
		if err != nil {
			return "", errors.New("private key is incorrect")
		}
	} else {
		privateKeyInterface, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
		if err != nil {
			return "", errors.New("private key is incorrect")
		}
		privateKey = privateKeyInterface.(*rsa.PrivateKey)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hash, contentBytes)
	if err != nil {
		return "", errors.New("sign fail")
	}

	sign := base64.StdEncoding.EncodeToString(signature)

	return sign, nil
}

func RsaVerify(content, signString, alipayPublicKey string, hash crypto.Hash) error {
	h := crypto.Hash.New(hash)
	h.Write([]byte(content))
	contentBytes := h.Sum(nil)

	signBytes, err := base64.StdEncoding.DecodeString(signString)
	if err != nil {
		return errors.New("decode sign fail")
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(alipayPublicKey)
	if err != nil {
		return errors.New("decode alipay public key fail")
	}

	var publicKey *rsa.PublicKey
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return errors.New("alipay public key is incorrect")
	}
	publicKey = publicKeyInterface.(*rsa.PublicKey)

	err = rsa.VerifyPKCS1v15(publicKey, hash, contentBytes, signBytes)
	if err != nil {
		return errors.New("verify fail")
	}

	return nil
}
