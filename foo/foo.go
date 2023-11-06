package foo

const (
	Max = 100
	//小文字だと外部から参照できない
	min = 1
)

func ReturnMin() int {
	return min
}
