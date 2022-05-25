package main

import (
	"bufio"
	"fmt"
	BO "mypkg/basicOperation"
	TC "mypkg/tempConv"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type kfc int

const (
	Kelvin    kfc = 1
	Farenheit kfc = 2
	Celsius   kfc = 3
)

type temp struct {
	TempNum   float64
	Indicator kfc
}
func WeAreInWindows() bool {
	return strings.Contains(runtime.GOOS, "windows")
}
func clearConsole(indq bool){
	if indq == true {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    	cmd.Stdout = os.Stdout
    	cmd.Run()
	}else {
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
    
}
var weAreIn bool
func convTemp(t1 temp, t2 temp) string{
	var newNumber float64
	var message string
	var grad string 
	if t1.Indicator == 1 {
		newNumber = TC.ToKelvin(t2.TempNum, int(t2.Indicator))
		grad = "K"
	}else if t1.Indicator == 2 {
		newNumber = TC.ToFarenheit(t2.TempNum, int(t2.Indicator))
		grad = "F"
	}else {
		newNumber = TC.ToCelsius(t2.TempNum, int(t2.Indicator))
		grad = "C"
	}
	op := basicOperation()
	if op == 1 {
		message = BO.Add(t1.TempNum, newNumber)
	}else if op == 2 {
		message = BO.Substract(t1.TempNum, newNumber)
	} else if op == 3 {
		message = BO.MultiplyBy(t1.TempNum, newNumber)
	}else if op == 4 {
		message = BO.DivideBy(t1.TempNum, newNumber)
		if message == "Error, no puedes dividir entre 0" {
			return message
		}
	}
	return message + grad


}
func basicOperation() int {
	var option int
	var err error
	for {
		clearConsole(weAreIn)
		fmt.Print("Que quieres? 1. Suma, 2. Restar, 3. Multiplicar y 4. Dividir: ")
		readerChoose := bufio.NewScanner(os.Stdin)
		readerChoose.Scan()
		optionText := readerChoose.Text()
		option, err = strconv.Atoi(optionText) 
		if err == nil  && (option == 1 || option == 2 || option == 3 || option == 4 ){
			break
		}
	}
	return option
	

}
func main() {
	weAreIn = WeAreInWindows()
	var fTemp, sTemp int
	var fKFC kfc
	var err error
	for {
		clearConsole(weAreIn)
		fmt.Print("Ingrese su primera temperatura: ")
		readerTemp1 := bufio.NewScanner(os.Stdin)
		readerTemp1.Scan()
		fTempText := readerTemp1.Text()
		fTemp, err = strconv.Atoi(fTempText) 
		fmt.Print("Ingrese si es 1. Kelvin, 2. Farenheit o 3. Celsius: ")
		readerKFC := bufio.NewScanner(os.Stdin)
		readerKFC.Scan()
		kfcTEXT := readerKFC.Text()
		KFC, er := strconv.Atoi(kfcTEXT)
		fKFC = kfc(KFC)
		if err == nil && er == nil && (fKFC == 1 || fKFC== 2 || fKFC == 3 ){
			break
		}
		
	}
	firstTemp := temp {float64(fTemp),fKFC}
	for {
		clearConsole(weAreIn)
		fmt.Print("Ingrese su segunda temperatura: ")
		readerTemp2 := bufio.NewScanner(os.Stdin)
		readerTemp2.Scan()
		sTempText := readerTemp2.Text()
		sTemp, err = strconv.Atoi(sTempText) 
		fmt.Print("Ingrese si es 1. Kelvin, 2. Farenheit o 3. Celsius: ")
		readerKFC := bufio.NewScanner(os.Stdin)
		readerKFC.Scan()
		kfcTEXT := readerKFC.Text()
		KFC, er := strconv.Atoi(kfcTEXT)
		fKFC = kfc(KFC)
		if err == nil && er == nil && (fKFC == 1 || fKFC== 2 || fKFC == 3 ){
			break
		}
	}
	secondTemp := temp{float64(sTemp), fKFC}
	message := convTemp(firstTemp, secondTemp)
	fmt.Println("El resultad es "+ message)
	
	
}