package main

import (
	a "./fctAssurance"
	t "./fctmoteur"
)

func main() {
	ta := t.Ta(200000, 300, 1, "m")
	//fmt.Printf(string(ta))
	// ass := a.Assurance(ta)
	a.Assurance(ta)
	// fmt.Printf(string(ass))
}
