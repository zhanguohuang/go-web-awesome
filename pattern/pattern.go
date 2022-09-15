package pattern

import "regexp"

var (

	// 时间 02:00-23:59
	OpeningDayTimeRangePattern, _ = regexp.Compile(`^[0-2][0-9]:[0-5][0-9]-[0-2][0-9]:[0-5][0-9]$`)
	// 联系方式
	ContactMobilePattern, _ = regexp.Compile(`^[+-0123456789]{1,20}$`)
	// 0.1~0.9
	ZeroWithOneDecimal, _ = regexp.Compile(`^(0\.[1-9])$`)
	// 0.01~0.99，注意0.0是不行的
	ZeroWithTwoDecimal, _ = regexp.Compile(`^(0\.[0-9]?[0-9])$`)
	// 如果是9的还是有优势的，开头可以直接[1-9]
	// 0.1~999.9 先搞定999.9，然后0.1
	LessThanOneThousandWithOneDecimal, _ = regexp.Compile(`^(([1-9][0-9]{0,2}(\.[1-9])?)|(0\.[1-9]))$`)
	// 不是9开头，得搞[1-2]开头，然后处理少一位的9开头
	// 0~30 先搞定0.01，然后30，然后29.99或9.99
	BetweenZeroAndThirtyWithTwoDecimal, _ = regexp.Compile(`^((0(\.[0-9]{1,2})?)|(30)|((([1-2][0-9])|([1-9]))?(\.[0-9]{1,2})?))$`)
	// 0.01~999.99 先搞定999.99，然后0.01
	LessThanOneThousandWithTwoDecimal, _ = regexp.Compile(`^(([1-9][0-9]{0,2}(\.[0-9]{1,2})?)|(0\.[0-9]?[0-9]))$`)
	// 1~999的整数
	MoreThanZeroAndLessThanOneThousand, _ = regexp.Compile(`^[1-9][0-9]{0,2}$`)
	// 0~99 先搞定99.99，然后0.01
	LessThanOneHundredWithTwoDecimal, _ = regexp.Compile(`^(([1-9][0-9]?(\.[0-9]{1,2})?)|(0\.[0-9]?[0-9]))$`)
	// 0~9999
	MoreThanZeroAndLessThanTenThousand, _ = regexp.Compile(`^[1-9][0-9]{0,3}$`)
	// 0~99999
	MoreThanZeroAndLessThanOneHundredThousand, _ = regexp.Compile(`^[1-9][0-9]{0,4}$`)

	// 最多3个整数和1个小数 999.9,0.1,34.,.123
	threeIntegerAndOneDecimal, _ = regexp.Compile(`^(([0-9]{1,3}(\.[0-9]?)?)|([0-9]?\.[0-9]))$`)
	// 最多2个整数和2个小数 22.22,0.22
	twoIntegerAndTwoDecimal, _ = regexp.Compile(`^(([0-9]{1,2}(\.[0-9]{0,2})?)|([0-9]?\.[0-9]{1,2}))$`)
	// 最多3个整数和2个小数 999.99,0.11,34.,.123
	threeIntegerAndTwoDecimal, _ = regexp.Compile(`^(([0-9]{1,3}(\.[0-9]{0,2})?)|([0-9]?\.[0-9]{1,2}))$`)
	// 3个整数
	threeInteger, _ = regexp.Compile(`^[0-9]{1,3}$`)
	// 最多1个整数和2个小数 2.22,0.22
	oneIntegerAndTwoDecimal, _ = regexp.Compile(`^(([0-9](\.[0-9]{0,2})?)|([0-9]?\.[0-9]{1,2}))$`)
)
