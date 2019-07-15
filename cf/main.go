package main

import (
	"cf/conv"
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func mainPart(str string) {
	t, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	var1 := conv.Celsius(t)
	var2 := conv.Fahrenheit(t)
	var3 := conv.Kelvin(t)
	var4 := conv.Foot(t)
	var5 := conv.Meter(t)
	cf := conv.CToF(var1)	
	ck := conv.CToK(var1)	
	fc := conv.FToC(var2)
	fk := conv.FToK(var2)
	kf := conv.KToF(var3)
	kc := conv.KToC(var3)
	fm := conv.FToM(var4)
	mf := conv.MToF(var5)
	fmt.Println("Rounded values")
	fmt.Printf(" 1: %v -> %v -> %v\n 2: %v -> %v -> %v\n 3: %v -> %v -> %.7v\n 4: %v -> %v\n 5: %v -> %v\n", var1, cf, ck, var2, fc, fk, var3, kf, kc, var4, fm, var5, mf)	
}

func main() {
numbers := os.Args[1:]
if len(numbers) == 0 {
		input := bufio.NewScanner(os.Stdin) 
		for input.Scan() {
			mainPart(input.Text())
	}
} else {
		for _, arg := range os.Args[1:] {
			mainPart(arg)
		}
	}
}
