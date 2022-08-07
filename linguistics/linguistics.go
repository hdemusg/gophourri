package linguistics

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/csv"
	"strconv"
	"strings"
)

/*
func Translate(f string, t string, p string) {
    
}
*/
/*
func LoadModels(modelloc string) (models map[string]map[string]int) {
	files, err := ioutil.ReadDir(modelloc)
	if err != nil {
		fmt.Println(err)
        return nil
	}
	for _, file := range files {
		f, err := os.Open(modelloc + "/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		var transitions = make(map[string]int)
		reader := csv.NewReader(f)
		model, e := reader.ReadAll()
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("Now reading " + file.Name())
		for _, line := range model {
            if err != nil {
				fmt.Println(err)
				return nil
			}
			k := line[0]
			if k == "sum" {
                sum, e := strconv.Atoi(line[1])
				if e != nil {
					return nil
				}
			} else {
				if len(k) == 2 {
					if k[0] == '+' {
						k = " " + k
					} else {
						k += " "
					}
				}
				transitions[k], err = strconv.Atoi(line[1])
				if err != nil {
					return nil
				}
			}
		}
		fn := file.Name()
        models[fn[:len(fn) - 4]] = transitions
	}
	return
}
*/

func TrainModels(dataloc string, targetloc string) {
	files, err := ioutil.ReadDir(dataloc)
	if err != nil {
		fmt.Println(err)
		return
	}
	var languages []string
	for _, file := range files {
		fn := file.Name()
		f, err := os.Open(dataloc + "/" + fn)
		language := fn[:len(fn) - 4]
		languages = append(languages, language)
		//fmt.Println(language)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		var transitions = make(map[string]int)
		transitions["sum"] = 0
		reader := csv.NewReader(f)
	    corpus, e := reader.ReadAll()
		if e != nil {
			fmt.Println(e)
			return
		}
		for _, line := range corpus {
			if err != nil {
				fmt.Println(err)
				return
			}
			l := line[0]
			for i := 0; i < len(l) - 1; i += 1 {
				key := string(l[i]) + "+" + string(l[i+1])
				// fmt.Println(transitions[key])
				if transitions[key] == 0 {
					transitions[key] = 1
				} else {
					transitions[key] = transitions[key] + 1
				}
				// fmt.Println(transitions[key])
				transitions["sum"] += 1
			}
		}
        modelfile, err := os.Create(targetloc + "/" + language + ".csv")
        //modelfile.Close()

		w := csv.NewWriter(modelfile)

		for k, v := range transitions {
			pair := []string{strings.TrimSpace(k),strconv.Itoa(v)}
			fmt.Println(pair)
			w.Write(pair)
		}
		w.Flush()
		modelfile.Close()
	}
}

func PredictLanguageInput(dataloc string) {
	fmt.Println("Enter a phrase. Don't include any commas or punctuation.")
    var phrase string
	fmt.Scanln(&phrase)
	files, err := ioutil.ReadDir(dataloc)
	if err != nil {
		fmt.Println(err)
		return
	}
	var languages []string
	p := strings.TrimSpace(phrase)
	var scores = make(map[string]float64)
	scoresum := 0.0
	for _, file := range files {
		fn := file.Name()
		f, err := os.Open(dataloc + fn)
		language := fn[:len(fn) - 4]
		languages = append(languages, language)
		scores[language] = 1.0
		//fmt.Println(language)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		var transitions = make(map[string]int)
		transitions["sum"] = 0
		reader := csv.NewReader(f)
	    corpus, e := reader.ReadAll()
		if e != nil {
			fmt.Println(e)
			return
		}
		for _, line := range corpus {
			if err != nil {
				fmt.Println(err)
				return
			}
			l := line[0]
			for i := 0; i < len(l) - 1; i += 1 {
				key := string(l[i]) + "+" + string(l[i+1])
				// fmt.Println(transitions[key])
				if transitions[key] == 0 {
					transitions[key] = 1
				} else {
					transitions[key] = transitions[key] + 1
				}
				// fmt.Println(transitions[key])
				transitions["sum"] += 1
			}
		}
		den := float64(transitions["sum"] + 1)
		for j := 0; j < len(p) - 1; j += 1 {
			key := string(p[j]) + "+" + string(p[j+1])
			num := float64((transitions[key] + 1)) / den
			//println(transitions[key] + 1, den, num)
            scores[language] = scores[language] * num
		    //println(scores[language])
		}
		scoresum += scores[language]
	}
	var most_likely string
	var ml_score float64
	for _, lang := range languages {
		score := scores[lang] / scoresum
		fmt.Println(lang + ": " + fmt.Sprintf("%f", score))
	    if score > ml_score {
			ml_score = score
			most_likely = lang
		}
	}
	fmt.Println("Prediction: " + most_likely)
	fmt.Println("Confidence: " + fmt.Sprintf("%f", ml_score))
}

func PredictLanguage(phrase string, dataloc string) {
    files, err := ioutil.ReadDir(dataloc)
	if err != nil {
		fmt.Println(err)
		return
	}
	var languages []string
	p := strings.TrimSpace(phrase)
	var scores = make(map[string]float64)
	scoresum := 0.0
	for _, file := range files {
		fn := file.Name()
		f, err := os.Open(dataloc + fn)
		language := fn[:len(fn) - 4]
		languages = append(languages, language)
		scores[language] = 1.0
		//fmt.Println(language)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		var transitions = make(map[string]int)
		transitions["sum"] = 0
		reader := csv.NewReader(f)
	    corpus, e := reader.ReadAll()
		if e != nil {
			fmt.Println(e)
			return
		}
		for _, line := range corpus {
			if err != nil {
				fmt.Println(err)
				return
			}
			l := line[0]
			for i := 0; i < len(l) - 1; i += 1 {
				key := string(l[i]) + "+" + string(l[i+1])
				// fmt.Println(transitions[key])
				if transitions[key] == 0 {
					transitions[key] = 1
				} else {
					transitions[key] = transitions[key] + 1
				}
				// fmt.Println(transitions[key])
				transitions["sum"] += 1
			}
		}
		den := float64(transitions["sum"] + 1)
		for j := 0; j < len(p) - 1; j += 1 {
			key := string(p[j]) + "+" + string(p[j+1])
			num := float64((transitions[key] + 1)) / den
			//println(transitions[key] + 1, den, num)
            scores[language] = scores[language] * num
		    //println(scores[language])
		}
		scoresum += scores[language]
	}
	var most_likely string
	var ml_score float64
	for _, lang := range languages {
		score := scores[lang] / scoresum
		fmt.Println(lang + ": " + fmt.Sprintf("%f", score))
	    if score > ml_score {
			ml_score = score
			most_likely = lang
		}
	}
	fmt.Println("Prediction: " + most_likely)
	fmt.Println("Confidence: " + fmt.Sprintf("%f", ml_score))
}