package cast

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInterfaceCast(t *testing.T) {
	m := make(map[string]interface{})
	var a_int64 int64 = 10
	var a_int32 int32 = 20
	var a_int int = 300
	m["a_int64"] = a_int64
	m["a_int32"] = a_int32
	m["a_int"] = a_int
	for _, v := range m {
		vv := fmt.Sprintf("%v", v)
		intValue, err := strconv.ParseInt(vv, 10, 64)
		fmt.Println(intValue, err)
	}
	d, d1 := m["a_null"]
	ddd := fmt.Sprintf("%v", d)
	fmt.Println(d, d1, ddd)
}
