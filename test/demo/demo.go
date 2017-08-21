package demo

func Myfun(number int) int  {
	if number < 2 {
		return number
	}
	return Myfun(number - 1) + Myfun(number - 2)
}