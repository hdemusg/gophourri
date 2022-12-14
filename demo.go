package main

import (
	"fmt"
	"time"
	//"reflect"
	"os"

	"github.com/hdemusg/gophourri/conversions"
	"github.com/hdemusg/gophourri/probabilities"
	"github.com/hdemusg/gophourri/linguistics"
)

// A sandbox for testing new functions
func main() {
    fun := os.Args[1]
    switch(fun) {
    case "temp":
	// Testing temperature conversion 
        conversions.ConvertTempInput()
    case "draw": 
	r, e := probabilities.WeightedDrawingCSV("data.csv")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(r)
	}
    case "roll-d20":
	fmt.Println(probabilities.DnDRoll(20))
    case "roll-d12":
	fmt.Println(probabilities.DnDRoll(12))
    case "roll-d8":
	fmt.Println(probabilities.DnDRoll(8))
    case "roll-d4":
	fmt.Println(probabilities.DnDRoll(4))
    case "roll":
	r, _ := probabilities.RollDice(1)
	fmt.Println(r)
    case "time":
	time := time.Now()
        fmt.Println(time.String())
    case "lang":
	linguistics.PredictLanguageInput("./language_data/")
    default:
	fmt.Println("YOU SHALL NOT PASS.")
	fmt.Println("Supported commands: \n temp: temperature conversion \n draw: weighted drawing \n roll-<die>: rolls a die \n roll: rolls a 6-sided die \n time: tells the time \n lang: Predicts language of a user-provided quote.")
}
