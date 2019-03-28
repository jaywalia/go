package main

import "fmt"

func main(){
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C \n", freezingF, fToC(freezingF) )
	fmt.Printf("%g F = %g C \n", boilingF, fToC(boilingF) )
	fmt.Printf("%g", 1.0* 5/9)
}

func fToC(f float64) float64 {
	return ( f - 32 ) * 5/9
}