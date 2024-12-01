package useful_things

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IndexStr(data []string, findChar string) int {
	for i, v := range data {
		if v == findChar {
			return i
		}
	}
	return -1
}

func Call() {
	var data string
	var shift string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text to encrypt:")
	data, err := reader.ReadString('\n')
	fmt.Println("Enter shift encrypt:")
	shift, err2 := reader.ReadString('\n')

	if err != nil || err2 != nil {
		return
	}

	shiftInt, _ := strconv.Atoi(strings.Trim(shift, "\n"))

	fmt.Printf("%v\n", ASCIICaesarsCipher(data, shiftInt))
}

func ASCIICaesarsCipher(data string, shift int) string {
	encrypted := ""
	letters := [33]string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z",
	}

	data = strings.Trim(strings.ToLower(data), "\n")

	for i := 0; i < len(data); i++ {
		index := IndexStr(letters[:], string(data[i]))

		if index == -1 {
			encrypted += string(data[i])
			continue
		}

		encrypted += letters[(index+shift)%len(letters)]
	}

	return encrypted
}
