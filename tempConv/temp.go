package temp

func ToFarenheit(t float64, Indicator int) float64 {
	var FarenheitTemp float64
	if Indicator == 1 {
		FarenheitTemp = 1.8*(t-273.15) + 32
	} else if Indicator == 2 {
		FarenheitTemp = t
	} else {
		FarenheitTemp = 1.8*(t) + 32
	}
	return FarenheitTemp
}
func ToKelvin(t float64, Indicator int) float64 {
	var KelvinTemp float64
	if Indicator == 1 {
		KelvinTemp = t
	} else if Indicator == 2 {
		KelvinTemp = ((t - 32) / 1.8) + 273.15
	} else {
		KelvinTemp = t + 273.15
	}
	return KelvinTemp
}
func ToCelsius(t float64, Indicator int) float64 {
	var CelsiusTemp float64
	if Indicator == 1 {
		CelsiusTemp = t - 273.15
	} else if Indicator == 2 {
		CelsiusTemp = ((t - 32) * 5) / 9
	} else {
		CelsiusTemp = t
	}
	return CelsiusTemp
}