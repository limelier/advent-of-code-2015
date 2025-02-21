package solvers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
func Solve04(input string) (int, int) {
	iFiveZeroes := -1
	iSixZeroes := -1
	for i := 0; iFiveZeroes == -1 || iSixZeroes == -1; i++ {
		in := fmt.Sprintf("%s%d", input, i)
		out := md5.Sum([]byte(in))
		outString := hex.EncodeToString(out[:])
		if strings.HasPrefix(outString, "00000") && iFiveZeroes == -1 {
			iFiveZeroes = i
		}
		if strings.HasPrefix(outString, "000000") && iSixZeroes == -1 {
			iSixZeroes = i
		}
	}
	return iFiveZeroes, iSixZeroes
}
