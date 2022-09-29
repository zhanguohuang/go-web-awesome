package array

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestNotIn(t *testing.T) {
	//has := []int64{
	//	59757148,
	//	60069407,
	//	60069408,
	//	60069410,
	//	60069412,
	//	60069413,
	//	60069414,
	//	60069416,
	//	60069417,
	//	60069871,
	//	60069872,
	//	60069874,
	//	60069875,
	//	60069877,
	//	60069878,
	//	60649789,
	//	60834795,
	//	60939938,
	//	60940012,
	//	60940038,
	//	60940051,
	//	60940096,
	//	60940142,
	//	60940149,
	//	60940156,
	//	60940190,
	//	60940211,
	//	60940244,
	//	60940247,
	//	60940272,
	//	60940294,
	//	60940313,
	//	60940322,
	//	60940325,
	//	60940336,
	//	60940349,
	//	60940370,
	//	60940389,
	//	60940414,
	//	60940422,
	//	60943414,
	//	60943426,
	//	60943443,
	//	60944186,
	//}
	all := []int64{
		60649789,
		59757148,
		60069878,
		60069877,
		60069875,
		60069874,
		60069872,
		60069871,
		60069417,
		60069416,
		60069414,
		60069413,
		60069412,
		60069410,
		60069408,
		60069407,
		60834795,
		60939938,
		60940012,
		60940038,
		60940051,
		60940096,
		60940142,
		60940149,
		60940156,
		60940190,
		60940211,
		60940238,
		60940244,
		60940247,
		60940272,
		60940294,
		60940313,
		60940322,
		60940325,
		60940336,
		60940342,
		60940349,
		60940370,
		60940389,
		60940414,
		60940422,
		60943414,
		60943426,
		60943443,
		60944186,
	}
	m := make(map[int64]bool)

	for _, a := range all {
		if _, ok := m[a]; ok {
			fmt.Println(a)
		}
		m[a] = true
	}

}

func TestLenMap(t *testing.T) {
	m := make(map[int64]int64)
	fmt.Println(len(m))
	m[1] = 2
	fmt.Println(len(m))
	m = nil
	fmt.Println(len(m))
}

func TestIntArr(t *testing.T) {
	var arr []int64
	fmt.Println(arr == nil)
	arrJson, err := json.Marshal(arr)
	fmt.Println(arrJson, err)
	err = json.Unmarshal(arrJson, &arr)
	fmt.Println(err)

	arr = nil
	fmt.Println(arr == nil)
	arrJson, err = json.Marshal(arr)
	fmt.Println(arrJson, err)
}

type ErrorCode int64

type ResponseJson struct {
	St        ErrorCode   `json:"st"`
	Msg       string      `json:"msg"`
	Code      ErrorCode   `json:"code"`
	Data      interface{} `json:"data"`
	Page      int64       `json:"page"`
	Size      int64       `json:"size"`
	Total     int64       `json:"total"`
	HasMaster *bool       `json:"has_master,omitempty"`
}

func TestMap(t *testing.T) {
	data := make(map[string]interface{})
	respJson := &ResponseJson{
		St:        ErrorCode(0),
		Msg:       "",
		Code:      ErrorCode(1000811234231),
		Data:      nil,
		Page:      0,
		Size:      0,
		Total:     0,
		HasMaster: nil,
	}
	respJsonByte, _ := json.Marshal(respJson)
	decoder := json.NewDecoder(bytes.NewReader(respJsonByte))
	decoder.UseNumber()
	_ = decoder.Decode(&data)

	cd, existCd := data["code"]
	st, existSt := data["st"]
	// 两个都有值的继续比较
	codeStr := fmt.Sprintf("%v", cd)
	stStr := fmt.Sprintf("%v", st)
	codeInt, pCdErr := strconv.ParseInt(codeStr, 10, 64)
	stInt, pStErr := strconv.ParseInt(stStr, 10, 64)
	fmt.Println(cd, existCd, st, existSt, codeStr, stStr, codeInt, pCdErr, stInt, pStErr)
}
