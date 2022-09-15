package cast

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFloatToString(t *testing.T) {
	fList := []float64{0., 323.12344, 0.10, 0.1, 0.12, -11.11}
	for _, f := range fList {
		var fs string
		// origin:3.123440, Sprintf:3.123440
		fs = fmt.Sprintf("%f", f)
		fmt.Printf("origin:%f, Sprintf:%s\n", f, fs)

		// 推荐，几乎和想要的一致
		// origin:3.123440, FormatFloatPrec-1:3.12344
		fs = strconv.FormatFloat(f, 'f', -1, 64)
		fmt.Printf("origin:%f, FormatFloatPrec-1:%s\n", f, fs)

		// origin:3.127440, FormatFloatPrec2:3.13
		fs = strconv.FormatFloat(f, 'f', 2, 64)
		fmt.Printf("origin:%f, FormatFloatPrec2:%s\n", f, fs)

	}

}

func TestFloatToString2(t *testing.T) {
	fList := []float64{0., 323.12344, 0.110, 0.1, 0.12, -11.11}
	for _, f := range fList {
		s := fmt.Sprintf("%v\n", f)
		fmt.Println(s)
	}

}
