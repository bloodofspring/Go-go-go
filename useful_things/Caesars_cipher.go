package useful_things

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Encrypt(plaintext string, shift int) string {
	ciphertext := ""

	for _, char := range plaintext {
		ciphertext += string(char + int32(shift))
	}

	return ciphertext
}

func Decrypt(ciphertext string, shift int) string {
	if shift < 0 {
		return Encrypt(ciphertext, -shift)
	} else {
		return Encrypt(ciphertext, -shift)
	}
}

func CipherMain() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text:")
	plaintext, err := reader.ReadString('\n')
	fmt.Println("Enter shift:")
	shift, err2 := reader.ReadString('\n')
	shiftInt, err3 := strconv.Atoi(strings.Trim(shift, "\n "))

	if err != nil || err2 != nil || err3 != nil {
		panic(err)
	}

	if shiftInt < -32 || shiftInt > 256000 {
		panic("Shift must be between -32 and 32000.")
	}

	fmt.Println("Choose operation [d/e] (decrypt/encrypt)")
	operation, _ := reader.ReadString('\n')

	fmt.Print("Your text: ", plaintext)

	switch strings.Trim(operation, "\n ") {
	case "e":
		encryptedText := Encrypt(plaintext, shiftInt)
		fmt.Println("Encrypted: ", encryptedText)
	case "d":
		decryptedText := Decrypt(plaintext, shiftInt)
		fmt.Print("Decrypted: ", decryptedText)
	default:
		println("Invalid operation.")
	}
}
