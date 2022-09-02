package pattern

import "testing"

// 0.1~999.9 先搞定999.9，然后0.1
func TestLessThanOneThousandWithOneDecimal(t *testing.T) {
	// 正确的
	if !LessThanOneThousandWithOneDecimal.MatchString("0.1") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("9") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("999") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("999.9") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("1.1") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("20.1") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("1") {
		t.Fail()
	}
	if !LessThanOneThousandWithOneDecimal.MatchString("99") {
		t.Fail()
	}
	// 错误的
	if LessThanOneThousandWithOneDecimal.MatchString("-1") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("1.") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("1.-") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("1000") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("9.0") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("11.99") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("11.09") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("01.9") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("0") {
		t.Fail()
	}
	if LessThanOneThousandWithOneDecimal.MatchString("0.") {
		t.Fail()
	}
}

// 0~30 先搞定0.01，然后30，然后29.99或9.99
func TestBetweenZeroAndThirtyWithTwoDecimal(t *testing.T) {
	// 正确的
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("0") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("30") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("0.1") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("29") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("29.99") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("1.1") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("20.1") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("1") {
		t.Fail()
	}
	if !BetweenZeroAndThirtyWithTwoDecimal.MatchString("9") {
		t.Fail()
	}
	// 错误的
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("-1") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("1.") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("1.-") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("31") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("9.0") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("20.10") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("11.099") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("01.9") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("09") {
		t.Fail()
	}
	if BetweenZeroAndThirtyWithTwoDecimal.MatchString("0.") {
		t.Fail()
	}
}

// 0.01~999.99 先搞定999.99，然后0.01
func TestLessThanOneThousandWithTwoDecimal(t *testing.T) {
	// 正确的
	if !LessThanOneThousandWithTwoDecimal.MatchString("30") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("0.1") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("0.11") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("29") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("29.99") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("329.99") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("999.99") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("1.1") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("1.19") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("1.09") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("20.1") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("20.09") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("1") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("9") {
		t.Fail()
	}
	if !LessThanOneThousandWithTwoDecimal.MatchString("999") {
		t.Fail()
	}
	// 错误的
	if LessThanOneThousandWithTwoDecimal.MatchString("-1") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("1.") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("1.-") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("1000") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("9.0") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("9.001") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("0.001") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("09") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("20.109") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString(".099") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("01.9") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("09") {
		t.Fail()
	}
	if LessThanOneThousandWithTwoDecimal.MatchString("0.") {
		t.Fail()
	}
}

func TestMoreThanZeroAndLessThanOneThousand(t *testing.T) {
	// 正确的
	if !MoreThanZeroAndLessThanOneThousand.MatchString("1") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneThousand.MatchString("9") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneThousand.MatchString("99") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneThousand.MatchString("999") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneThousand.MatchString("10") {
		t.Fail()
	}
	// 错误的
	if MoreThanZeroAndLessThanOneThousand.MatchString("0") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneThousand.MatchString("1000") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneThousand.MatchString("10.1") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneThousand.MatchString("9909") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneThousand.MatchString("00") {
		t.Fail()
	}
}

// 0~99 先搞定99.99，然后0.01
func TestLessThanOneHundredWithTwoDecimal(t *testing.T) {
	// 正确的
	if !LessThanOneHundredWithTwoDecimal.MatchString("30") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("0.1") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("0.11") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("29") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("29.99") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("99.99") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("1.1") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("1.19") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("1.09") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("20.1") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("20.09") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("1") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("9") {
		t.Fail()
	}
	if !LessThanOneHundredWithTwoDecimal.MatchString("99") {
		t.Fail()
	}
	// 错误的
	if LessThanOneHundredWithTwoDecimal.MatchString("-1") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("1.") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("1.-") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("100") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("9.0") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("9.001") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("0.001") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("09") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("20.109") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString(".099") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("01.9") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("09") {
		t.Fail()
	}
	if LessThanOneHundredWithTwoDecimal.MatchString("0.") {
		t.Fail()
	}
}

// 0~9999
func TestMoreThanZeroAndLessThanTenThousand(t *testing.T) {
	// 正确的
	if !MoreThanZeroAndLessThanTenThousand.MatchString("1") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanTenThousand.MatchString("9") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanTenThousand.MatchString("99") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanTenThousand.MatchString("999") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanTenThousand.MatchString("9999") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanTenThousand.MatchString("10") {
		t.Fail()
	}
	// 错误的
	if MoreThanZeroAndLessThanTenThousand.MatchString("0") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanTenThousand.MatchString("10000") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanTenThousand.MatchString("10.1") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanTenThousand.MatchString("99099") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanTenThousand.MatchString("00") {
		t.Fail()
	}
}

// 0~99999
func TestMoreThanZeroAndLessThanOneHundredThousand(t *testing.T) {
	// 正确的
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("1") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("9") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("99") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("999") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("9999") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("99999") {
		t.Fail()
	}
	if !MoreThanZeroAndLessThanOneHundredThousand.MatchString("10") {
		t.Fail()
	}
	// 错误的
	if MoreThanZeroAndLessThanOneHundredThousand.MatchString("0") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneHundredThousand.MatchString("100000") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneHundredThousand.MatchString("10.1") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneHundredThousand.MatchString("990999") {
		t.Fail()
	}
	if MoreThanZeroAndLessThanOneHundredThousand.MatchString("00") {
		t.Fail()
	}
}
