package sshkey

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ssh"
	"strings"
)

// GenerateSSHKeyPair returns a new ED25519 key pair (public, private keys) prepared for SSH
func GenerateSSHKeyPair() (string, string, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	publicKeyString, err := publicKeyToMarshal(publicKey)
	return publicKeyString, string(privateKeyToPEM(privateKey)), err
}

func publicKeyToMarshal(key ed25519.PublicKey) (string, error) {
	publicRsaKey, err := ssh.NewPublicKey(key)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(ssh.MarshalAuthorizedKey(publicRsaKey)), " \n\r"), nil
}

func privateKeyToPEM(privateKey ed25519.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privateKey),
	})
}
