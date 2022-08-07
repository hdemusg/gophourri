// Mathematical and Scientific conversions
package conversions

import (
	"errors"
	"fmt"
	"strconv"
	// "os"
	// "encoding/csv"
	// "strings"
)

// case 0: Fahrenheit to Celsius
func ftoc(f float64) (float64, error) {
    baseline := 32.0
	offset := f - baseline
	convFactor := 5.0 / 9
	return (convFactor * offset), nil
}

// case 1: Fahrenheit to Kelvin
func ftok(f float64) (float64, error) {
    baseline := 32.0
	offset := f - baseline 
	convFactor := 5.0 / 9
	return (convFactor * offset) + 273.0, nil
}

// case 2: Celsius to Fahrenheit
func ctof(f float64) (float64, error) {
    baseline := 32.0
	convFactor := 9.0 / 5
	return (convFactor * f) + baseline, nil
}

// case 3: Celsius to Kelvin
func ctok(f float64) (float64, error) {
    return f + 273.0, nil
}

// case 4: Kelvin to Celsius
func ktoc(f float64) (float64, error) {
	return f - 273.0, nil
}

//case 5: Kelvin to Fahrenheit
func ktof(f float64) (float64, error) {
    celsius := f - 273.0
	baseline := 32.0
	convFactor := 9.0 / 5
	return (convFactor * celsius) + baseline, nil
}

// Converts between temperature scales
//  f: input Fahrenheit degree value
//  c: conversion code
//   0: Fahrenheit to Celsius
//   1: Fahrenheit to Kelvin
//   2: Celsius to Fahrenheit
//   3: Celsius to Kelvin
//   4: Kelvin to Celsius
//   5: Kelvin to Fahrenheit
//  output: corresponding Celsius degree value
func ConvertTemp(f float64, c int) (float64, error) {
	switch (c) {
	  case 0:
        return ftoc(f)
	  case 1:
		return ftok(f)
	  case 2:
		return ctof(f)
	  case 3:
		return ctok(f)
	  case 4:
		return ktoc(f)
	  case 5:
		return ktof(f)
	  default:
		return 0.0, errors.New("Invalid code")
	}
}

// Asks user for input before running temperature conversion. 
// TODO: Can improve the UI to include the scales we are converting between
func ConvertTempInput() {
    fmt.Println("Enter a degree value. This should be a non-integer value.")
    var degrees string
	fmt.Scanln(&degrees)
	fmt.Println("Select an option from the following: \n 0: Fahrenheit to Celsius \n 1: Fahrenheit to Kelvin \n 2: Celsius to Fahrenheit \n 3: Celsius to Kelvin \n 4: Kelvin to Celsius \n 5: Kelvin to Fahrenheit")
	var code string
	fmt.Scanln(&code)
	d, err1 := strconv.ParseFloat(degrees, 64)
	c, err2 := strconv.Atoi(code)
	if err1 != nil {
		fmt.Printf("%s\n", errors.New("The degree value should be a numeric value."))
	} else if err2 != nil {
		fmt.Printf("%s\n", errors.New("The conversion code should be a numeric value."))
	} else {
		conversion, error := ConvertTemp(d, c)
	    if error != nil {
		  fmt.Println(error)
        } else {
		  fmt.Println(conversion)
	    }
	}
}

// contains the conversion between meters (the base unit) and all other supported units of measuring distance
var dist_model = map[string]float64 {
    "m": 1,
	"mi": 1609.344, 
	"km": 1000,
	"cm": 0.01,
	"mm": 0.001,
	"f": 1609.344 / 8.0,
}

// Converts between distance units
//  d: distance (float)
//  f: starting unit (string)
//  t: target unit (string)
//  Units:
//   m, meters (base)
//   km, kilometers
//   mi, miles
//   cm, centimeters
//   mm, millimeters
//   f, furlongs
//  NOTE: if f is not supported, you will get 0, while if t is not supported, you will get either Inf or -Inf
func ConvertDist(d float64, f string, t string) (nd float64) {
    //m, err := os.Open("distance_model.csv")
	nd = (d * dist_model[f]) / dist_model[t]
	return
}

// contains the conversion between grams (base unit) and all other supported units of measuring mass
var mass_model = map[string]float64 {
	"g": 1, 
	"kg": 1000,
	"mg": 0.001,
	"t": 1000000,
	"lb": 453.59,
	"Mt": 1000000000000,
	"Gt": 1000000000000000,
}

// Converts between distance units
//  m: mass (float)
//  f: starting unit (string)
//  t: target unit (string)
//  Units:
//   g, grams (base unit)
//   kg, kilograms
//   mg, milligrams
//   t, tonnes
//   lb, pounds
//   Mt, megatonne
//   Gt, gigatonne
//  NOTE: if f is not supported, you will get 0, while if t is not supported, you will get either Inf or -Inf
func ConvertMass(m float64, f string, t string) (nm float64) {
	nm = (m * mass_model[f]) / mass_model[t]
	return
}

