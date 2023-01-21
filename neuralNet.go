package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/martinbro/neural_net/data"
	"gonum.org/v1/gonum/mat"
)

func prtMatrix(m *mat.Dense) {
	f := mat.Formatted(m, mat.Prefix("\t"), mat.Squeeze())
	fmt.Printf("\n\t%v\n", f)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	input := mat.NewDense(5, 1, nil)

	outp := data.Ouput25()
	inp := data.Input25()

	// b := fmt.Sprint("./data/dat", time.Now().Format("13_50_01"), ".csv")
	start := time.Now()
	f, err := os.Create("./data/dat.csv")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "Voltage (V);Current (A);Temp (C);SOC (estimated);SOC (messured)\n")
	w.Flush()
	for i := 0; i < 475; i++ {
		for j := 0; j < 5; j++ {
			input.Set(j, 0, inp.At(j, i*100))

		}
		inputVals := fmt.Sprint(
			(4.23-2.79)*input.At(0, 0)+2.79, ";",
			(5.99+18.09)*input.At(1, 0)-18.09, ";",
			(26.81+10.30)*input.At(2, 0)-10.30)

		//Normalization,zerocenter
		input.Sub(input, data.Layers_1_mean())
		//første 'hidden layer' beregnes
		var lag1 mat.Dense

		lag1.Mul(data.Layers_2_weights(), input)
		lag1.Add(data.Layers_2_bias(), &lag1)
		lag1.Apply(func(i, j int, v float64) float64 { return math.Tanh(v) }, &lag1) //aktivation

		lag1.Mul(data.Layers_4_weights(), &lag1)
		lag1.Add(&lag1, data.Layers_4_bias())
		lag1.Apply(func(i, j int, v float64) float64 { return math.Max(v, 0.3*v) }, &lag1) //aktivation LekyRELU med trenshold på 0.3

		var lag2 mat.Dense
		lag2.Mul(data.Layers_6_weights(), &lag1)
		lag2.Add(data.Layers_6_bias(), &lag2)
		lag2.Apply(func(i, j int, v float64) float64 { return math.Min(1, math.Max(0, v)) }, &lag2)

		// fmt.Printf("%v;%v;%v\n", inputVals, lag2.At(0, 0), outp.At(0, i*100))
		p := fmt.Sprint(inputVals, ";", lag2.At(0, 0), ";", outp.At(0, i*100))
		s := strings.ReplaceAll(p, ".", ",")
		fmt.Fprintf(w, "%v\n", s)
		w.Flush()
	}
	fmt.Printf("%v", time.Since(start))

	//
	// lag2w := data.Layers_2_weights()
	// lag2b := data.Layers_2_bias()
	// lag4w := data.Layers_4_weights()
	// lag4b := data.Layers_4_bias()

	// prtMatrix(lag4w)
	// prtMatrix(lag4b)

}
