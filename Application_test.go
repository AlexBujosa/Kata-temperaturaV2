package main

import (
	BO "mypkg/basicOperation"
	CT "mypkg/tempConv"
	"strconv"
	"testing"
)

func TestCase1(t *testing.T){
	want := "55"
	expected :=BO.Add(5,50)
	if want != expected {
		t.Error("No devuelve el resultado esperado en cuanto a la suma")
	}
}

func TestCase2(t *testing.T){
	want := "3"
	expected := BO.Substract(5, 2)
	if want != expected {
		t.Error("No devuelve el resultado esperado en cuanto a la resta")
	}
}

func TestCase3(t *testing.T){
	want := "8"
	expected :=BO.MultiplyBy(4,2)
	if want != expected {
		t.Error("No devuelve el resultado esperado en cuanto a la multiplicacion")
	}
}

func TestCase4(t *testing.T){
	want := "10"
	expected := BO.DivideBy(20, 2)
	if want != expected {
		t.Error("No devuelve el resultado esperado en cuanto a la division")
	}
}
func TestCase5(t *testing.T){
	want := "Error, no puedes dividir entre 0"
	expected := BO.DivideBy(30, 0)
	if want != expected {
		t.Error("No devuelve el resultado esperado en cuanto a las divisiones entre 0")
	}
}

func TestCase6(t *testing.T){
	want := "50"
	expected := CT.ToFarenheit(10, 3)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a farenheit")
	}
}

func TestCase7(t *testing.T){
	want := "283.15"
	expected := CT.ToKelvin(10, 3)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a kelvin")
	}
}

func TestCase8(t *testing.T){
	want := "10"
	expected := CT.ToCelsius(10, 3)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a kelvin")
	}
}

func TestCase9(t *testing.T){
	want := "50"
	expected := CT.ToFarenheit(50, 2)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a farenheit")
	}
}

func TestCase10(t *testing.T){
	want := "283.15"
	expected := CT.ToKelvin(50, 2)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a kelvin")
	}
}

func TestCase11(t *testing.T){
	want := "10"
	expected := CT.ToCelsius(50, 2)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a kelvin")
	}
}

func TestCase12(t *testing.T){
	want := "-369.67"
	expected := CT.ToFarenheit(50, 1)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a farenheit")
	}
}

func TestCase13(t *testing.T){
	want := "50"
	expected := CT.ToKelvin(50, 1)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a kelvin")
	}
}

func TestCase14(t *testing.T){
	want := "-223.15"
	expected := CT.ToCelsius(50, 1)
	if want != strconv.FormatFloat(expected, 'f', -1, 32) {
		t.Error("No devuelve el resultado esperado en cuanto a la conversion a kelvin")
	}
}