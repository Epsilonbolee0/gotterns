package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetPercentage(name string) float64
}

const epsilon = 0.00001

type singletonDatabase struct {
	chart map[string]float64
}

func (db *singletonDatabase) GetPercentage(name string) float64 {
	return db.chart[name]
}

func GetTotalPercentageEx(db Database, languages []string) float64 {
	result := 0.0
	for _, language := range languages {
		result += db.GetPercentage(language)
	}

	return result
}

func readData(path string) (map[string]float64, error) {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(ex + path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]float64{}

	for scanner.Scan() {
		key := scanner.Text()

		scanner.Scan()
		cleansedText := scanner.Text()
		value, _ := strconv.ParseFloat(cleansedText, 64)

		result[key] = value
	}

	return result, nil

}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		chart, e := readData("\\Singleton\\language_chart.txt")
		db := singletonDatabase{chart}
		if e == nil {
			db.chart = chart
		}
		instance = &db
	})
	return instance
}

type DummyDatabase struct {
	dummyData map[string]float64
}

func (d *DummyDatabase) GetPercentage(name string) float64 {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]float64{
			"alpha": 1,
			"betta": 2,
			"gamma": 3,
		}
	}

	return d.dummyData[name]
}

func main() {
	languages := []string{"C#", "JavaScript", "Java"}
	tp := GetTotalPercentageEx(GetSingletonDatabase(), languages)

	ok := math.Abs(tp-47.1) < epsilon
	fmt.Println(ok)

	dummyNames := []string{"alpha", "gamma"}
	tp = GetTotalPercentageEx(&DummyDatabase{}, dummyNames)
	ok = math.Abs(tp-4) < epsilon
	fmt.Println(ok)
}
