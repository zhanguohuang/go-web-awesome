package cast

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStringToFloat(t *testing.T) {
	fList := []string{"0.", "323.12344", "0.100", "0.1", "0.12", "-11.11", ".1"}
	for _, fs := range fList {
		var ff float64
		ff, _ = strconv.ParseFloat(fs, 64)
		fmt.Println(ff)
	}
}
