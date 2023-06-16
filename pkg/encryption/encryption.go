package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"

	"github.com/the-go-dragons/final-project/pkg/config"
)

var key = []byte(fmt.Sprintf("%d", config.GetEnv("ENCRYPTION_SECRET_KEY", "encryptionSeretKey")))

func Encrypt(plaintext interface{}) (string, error) {
    // Convert the plaintext to a byte array
    var plaintextBytes []byte
    switch v := plaintext.(type) {
    case int:
        plaintextBytes = []byte(fmt.Sprintf("%d", v))
    case uint:
        plaintextBytes = []byte(fmt.Sprintf("%d", v))
    case string:
        plaintextBytes = []byte(v)
    default:
        return "", fmt.Errorf("unsupported plaintext type")
    }

    // Generate a new AES cipher block using the key
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    // Generate a new initialization vector (IV)
    iv := make([]byte, aes.BlockSize)
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    // Pad the plaintext to a multiple of the block size
    paddedPlaintext := pad(plaintextBytes, block.BlockSize())

    // Create a new AES CBC encrypter using the key and IV
    encrypter := cipher.NewCBCEncrypter(block, iv)

    // Encrypt the padded plaintext
    ciphertext := make([]byte, len(paddedPlaintext))
    encrypter.CryptBlocks(ciphertext, paddedPlaintext)

    // Concatenate the IV and ciphertext and return a base64-encoded string
    return base64.StdEncoding.EncodeToString(append(iv, ciphertext...)), nil
}

// Decrypt a base64-encoded ciphertext using the given key and return the decrypted plaintext
func Decrypt(ciphertext string) (interface{}, error) {
    // Decode the base64-encoded ciphertext
    ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return nil, err
    }

    // Generate a new AES cipher block using the key
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // Split the ciphertext into the IV and encrypted data
    iv := ciphertextBytes[:aes.BlockSize]
    ciphertextBytes = ciphertextBytes[aes.BlockSize:]

    // Create a new AES CBC decrypter using the key and IV
    decrypter := cipher.NewCBCDecrypter(block, iv)

    // Decrypt the ciphertext
    plaintext := make([]byte, len(ciphertextBytes))
    decrypter.CryptBlocks(plaintext, ciphertextBytes)

    // Unpad the plaintext and return the result as an integer or string
    plaintext, err = unpad(plaintext)
    if err != nil {
        return nil, err
    }

    if isAllDigits(plaintext) {
        // The plaintext is a number
        return strconv.Atoi(string(plaintext))
    } else {
        // The plaintext is a string
        return string(plaintext), nil
    }
}

// Pad the input to a multiple of the block size using PKCS#7 padding
func pad(input []byte, blockSize int) []byte {
    padding := blockSize - len(input)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(input, padtext...)
}

// Unpad the input using PKCS#7 padding
func unpad(input []byte) ([]byte, error) {
    length := len(input)
    unpadding := int(input[length-1])
    if unpadding > length {
        return nil, fmt.Errorf("invalid padding")
    }
    return input[:(length - unpadding)], nil
}

// Check if a byte array contains only digits
func isAllDigits(input []byte) bool {
    for _, b := range input {
        if b < '0' || b > '9' {
            return false
        }
    }
    return true
}