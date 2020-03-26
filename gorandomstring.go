package gorandomstring

import "math/rand"

var validChar = [55]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'p', 'r', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func genString(charLength int) string {
	var result string
	for letter := 0; letter < charLength; letter++ {
		prePick := 0
		for {
			picker := rand.Int() % 55
			if picker == prePick {
				continue
			} else {
				result += string(validChar[picker])
				break
			}
		}
	}
	return string(result)
}
