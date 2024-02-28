package reverse

import  "strconv"

// returns decimal reversal of integer i

func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))

	return i
}