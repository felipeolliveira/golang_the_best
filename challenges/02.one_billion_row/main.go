package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurementsFile, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurementsFile.Close()

	data := make(map[string]Measurement)
	scanner := bufio.NewScanner(measurementsFile)

	start := time.Now()
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolorPosition := strings.Index(rawData, ";")
		location := rawData[:semicolorPosition]
		rawTemp := rawData[semicolorPosition+1:]
		temperature, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count++
		}
		data[location] = measurement
	}

	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	for _, name := range locations {
		measurement := data[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")
	fmt.Println("Time:", time.Since(start))
}
