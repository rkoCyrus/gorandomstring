package gorandomstring

import (
	"math/rand"
	"time"
)

//These charther is excluded "I" and "l" which easy making confuse one
var validChar = [55]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'p', 'r', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

//Generate human redable string in bunch of charther (You can decide that reusing the same charther generated befor or not)
func GenString(charLength int, repeatChar bool) string {
	//For loop only support int (with negative), so I decide add this
	if charLength <= 0 {
		panic("Generate nothing")
	}
	var result string
	//First seed
	rand.Seed(time.Now().UnixNano())
	for letter := 0; letter < charLength; letter++ {
		prePick := 0
		for {
			picker := rand.Intn(54)
			if picker == prePick && !repeatChar {
				continue
			} else {
				result += string(validChar[picker])
				prePick = picker
				break
			}
		}
		//Adding refreshing seed to look like random
		if prePick % letter == 0 && letter != 0 && prePick != 0 {
			rand.Seed(time.Now().UnixNano())
		}
	}
	return result
}
