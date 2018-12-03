package utils

import (
	"math"
	"strconv"
	"strings"
)

func CalcValCode(code string) (string, error) {
	_, err := strconv.ParseInt(code, 10, 64)
	if err != nil || len(code) < 9 {
		return "Not a number or invalid length!", err
	}

	valCode := make([]string, 8)
	codeInt := make([]int, 9)

	for i := 0; i < 9; i++ {
		codeInt[i], _ = strconv.Atoi(code[i : i+1])
	}

	r0 := (-codeInt[8]) % 10
	c68 := codeInt[6]*10 + codeInt[8]
	vt := c68 - int(math.Ceil(float64(c68)/10))
	if codeInt[3] != 0 {
		r0 = (codeInt[3] - codeInt[8] - 1) % 10
		vt = c68 - int(math.Ceil(float64(c68-codeInt[3]+1)/10))
	}
	valCode[0] = strconv.Itoa(r0)
	valCode[4] = strconv.Itoa(vt / 10)
	valCode[5] = strconv.Itoa(vt % 10)

	valCode[1] = strconv.Itoa(10 - codeInt[3]%10)
	valCode[2] = strconv.Itoa(codeInt[0])
	valCode[3] = strconv.Itoa(codeInt[8])
	valCode[6] = strconv.Itoa(codeInt[6])
	valCode[7] = strconv.Itoa(codeInt[3])

	return strings.Join(valCode, ""), nil
}
