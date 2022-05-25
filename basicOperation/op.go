package op

import "strconv"

func Add(fNumber float64, sNumber float64) string {
	return strconv.FormatFloat(fNumber + sNumber, 'f', -1, 32)
}
func Substract(fNumber float64, sNumber float64) string {
	return strconv.FormatFloat(fNumber - sNumber, 'f', -1, 32)
}
func MultiplyBy(fNumber float64, sNumber float64) string {
	return strconv.FormatFloat(fNumber * sNumber, 'f', -1, 32)
}
func DivideBy(fNumber float64, sNumber float64) string {
	if sNumber == 0 {
		return "Error, no puedes dividir entre 0"
	}
	return strconv.FormatFloat(fNumber / sNumber, 'f', -1, 32)
}