package probabilities

import (
	"math/rand"
    "time"
    "errors"
    "os"
    "encoding/csv"
    "strconv"
    "strings"
)

// Takes in a list of integer weights and randomly selects a number in the range of the length of said list
func WeightedDrawing(weights ...int) (int, error) {
    sum := 0
    for _, num := range weights {
        sum = sum + num
    }
    x1 := rand.NewSource(time.Now().UnixNano())
    y1 := rand.New(x1)
	r := y1.Intn(sum)
    counter := 0
    for i, num := range weights {
        counter = counter + num
        if counter > r {
            return i, nil
        }
    }
	return len(weights), errors.New("Weighted drawing either got invalid inputs or broke during execution.")
}

// Takes in a list of float weights and randomly selects a number in the range of the length of said list
func WeightedDrawingFloat(weights ...float64) (int, error) {
    sum := 0.0
    for _, num := range weights {
        sum = sum + num
    }
    x1 := rand.NewSource(time.Now().UnixNano())
    y1 := rand.New(x1)
	r := (y1.Float64() * sum)
    counter := 0.0
    for i, num := range weights {
        counter = counter + num
        if counter > r {
            return i, nil
        }
    }
	return len(weights), errors.New("Weighted drawing either got invalid inputs or broke during execution.")
}

// Reads in a CSV containing unique labels and weights and randomly returns a label
func WeightedDrawingCSV(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "Could not open file " + filename, err
    } else {
        defer f.Close()
        csvReader := csv.NewReader(f)
        data, err := csvReader.ReadAll()
        if err != nil {
            return "Could not read file " + filename, err
        }
        m := make(map[string]float64)
        var sum float64
        for _, line := range data {
            var label string
            for j, unit := range line {
                if j == 0 {
                    label = string(strings.TrimSpace(unit))
                } 
                if j == 1 {
                    val, err := strconv.ParseFloat(strings.TrimSpace(unit), 64)
                    if err != nil {
                        return "Make sure your first column only consists of labels and your second column only consists of numeric weights.", err
                    }
                    sum = sum + val
                    if label != "" {
                        if _, ok := m[label]; ok {
                            return "Labels must be unique.", errors.New("Labels must be unique.")
                        }
                        m[label] = sum
                    } else {
                        return "Missing label", errors.New("Missing label")
                    }
                }
            }
        }
        x1 := rand.NewSource(time.Now().UnixNano())
        y1 := rand.New(x1)
	    r := (y1.Float64() * sum)
        for k, v := range m {
            if r < v {
                return k, nil
            }
        }
        return "None of the labels", errors.New("No label selected.")
    }
}

// Flips a positive number of coins and returns the results of each flip.
func FlipCoins(coins int) (flips []string, e error) {
    if coins <= 0 {
        e = errors.New("coins must be a positive integer")
    }
    for i := 0; i < coins; i++ {
        x1 := rand.NewSource(time.Now().UnixNano())
        y1 := rand.New(x1)
	    r := y1.Intn(2)
        if r == 0 {
            flips = append(flips, "tails")
        } else {
            flips = append(flips, "heads")
        }
    }
    return
}

// Rolls a positive number of dice and returns the sum of the rolls.
func RollDice(dice int) (sum int, e error) {
    if dice <= 0 {
        e = errors.New("dice must be a positive integer")
    }
    for i := 0; i < dice; i++ {
        x1 := rand.NewSource(time.Now().UnixNano())
        y1 := rand.New(x1)
	    r := y1.Intn(6) + 1
        sum += r
    }
    return
}

// Rolls a positive number of dice and returns the sum of the rolls.
func DnDRoll(value int) (r int) {
    x1 := rand.NewSource(time.Now().UnixNano())
    y1 := rand.New(x1)
	r = y1.Intn(value) + 1
    return
}