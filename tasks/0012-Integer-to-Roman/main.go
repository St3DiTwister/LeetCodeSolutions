package _012_Integer_to_Roman

func romanCount(symbol string, count int) string {
	res := ""
	for x := 0; x < count; x++ {
		res += symbol
	}
	return res
}

func digitToRoman(num int, one, five, ten string) string {
	if 1 <= num && num <= 3 {
		return romanCount(one, num)
	} else if num == 4 {
		return one + five
	} else if 5 <= num && num <= 8 {
		return five + romanCount(one, num-5)
	} else if num == 9 {
		return one + ten
	} else {
		return ""
	}
}

func intToRoman(num int) string {
	thousands := num / 1000
	hundreds := (num / 100) % 10
	tens := (num / 10) % 10
	ones := num % 10

	res := ""
	res += romanCount("M", thousands)
	res += digitToRoman(hundreds, "C", "D", "M")
	res += digitToRoman(tens, "X", "L", "C")
	res += digitToRoman(ones, "I", "V", "X")
	return res
}
