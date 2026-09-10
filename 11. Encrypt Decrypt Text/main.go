package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"
)

// AES-256 needs a 32-byte key.
const key = "go-mini-projects-secret-key-32b!"

// go run main.go -text="Milwad Khosravi"
func main() {
	text := flag.String("text", "", "Text to encrypt")
	flag.Parse()

	if *text == "" {
		log.Fatal("-text is required")
	}

	encrypted, err := encrypt(*text)
	if err != nil {
		log.Fatal(err)
	}

	decrypted, err := decrypt(encrypted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted:", encrypted)
	fmt.Println("Decrypted:", decrypted)
}

// encrypt seals the text with AES-256-GCM and returns it base64 encoded.
func encrypt(text string) (string, error) {
	gcm, err := newGCM()
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("reading random data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(text), nil)), nil
}

// decrypt reverses encrypt.
func decrypt(encoded string) (string, error) {
	message, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("decoding text: %w", err)
	}

	gcm, err := newGCM()
	if err != nil {
		return "", err
	}
	if len(message) < gcm.NonceSize() {
		return "", errors.New("text is too short to be an encrypted message")
	}

	nonce, ciphertext := message[:gcm.NonceSize()], message[gcm.NonceSize():]

	text, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.New("the text was modified")
	}

	return string(text), nil
}

// newGCM wraps the key in AES-GCM.
func newGCM() (cipher.AEAD, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	return cipher.NewGCM(block)
}
