package mymath

func MyPower(base int, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func MyAbs(number int) int { //value function !!!
	if number < 0 {
		return number * -1
	}
	return number
}
