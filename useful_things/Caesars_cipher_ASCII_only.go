package useful_things

import (
	"fmt"
)

func caesarEncrypt(plaintext string, shift int) string {
	ciphertext := ""
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			ciphertext += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			ciphertext += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			ciphertext += string(char)
		}
	}
	return ciphertext
}

func caesarDecrypt(ciphertext string, shift int) string {
	return caesarEncrypt(ciphertext, 26-shift)
}

func Main() {
	plaintext := "Привет, мир!"
	shift := 26 // It breaks if shift > 26 or < 0

	encryptedText := caesarEncrypt(plaintext, shift)
	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Encrypted:", encryptedText)

	decryptedText := caesarDecrypt(encryptedText, shift)
	fmt.Println("Decrypted:", decryptedText)
}
