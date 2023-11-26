package main

func PrintEvenValues(list []int) {
	if len(list) == 0 {
		return
	}

	val := list[0]
	if val%2 == 0 {
		println(val)
	}

	PrintEvenValues(list[1:])
}

func main() {
	l := []int{1, 2, 5, 53, 44, 46}
	PrintEvenValues(l)
}
