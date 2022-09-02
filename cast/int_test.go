package cast

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInt64ToString(t *testing.T) {
	iList := []int64{0, 34, 1000, 20}
	for _, i := range iList {
		var is string
		is = fmt.Sprintf("%d", i)
		fmt.Printf("origin:%d, Sprintf:%s\n", i, is)

		// 推荐
		is = strconv.FormatInt(i, 10)
		fmt.Printf("origin:%d, FormatInt:%s\n", i, is)

		is = strconv.Itoa(int(i))
		fmt.Printf("origin:%d, Itoa:%s\n", i, is)
	}
}
