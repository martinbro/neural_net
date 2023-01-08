package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func ReadDataFile(filename string, a *mat.Dense) {

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf("line: %d text %s\n", i, t)
		v := strings.Split(t, "\t")
		j := 0
		var src = make([]float64, 5)
		for _, val := range v {
			varF, _ := strconv.ParseFloat(val, 64)
			fmt.Println(varF, ", j:", j)
			src[j] = varF
			j++

		}
		a.SetRow(i, src)
		i++
	}
}
func readWholeFile(filename string) {
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(text))
}
func readStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("filename: %s\n", stat.Name())
	fmt.Printf("file last modified: %v\n", stat.ModTime().Format("15:04:05"))
}
