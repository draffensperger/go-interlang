package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetHarmonic(n int) float64 {
	resp, err := http.Get("http://localhost:8000/" + strconv.Itoa(n))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var harmonicBytes []byte
	harmonicBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var harmonic float64
	harmonic, err = strconv.ParseFloat(string(harmonicBytes[:]), 64)
	if err != nil {
		log.Fatal(err)
	}
	return harmonic
}

func HarmonicSum(n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += GetHarmonic(i)
	}
	return sum
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println("\n\nGo with n times HTTP calls:")
	fmt.Printf("Sum for n=1..%v of 1/n =\n", n)
	fmt.Printf("%v", HarmonicSum(n))
}
