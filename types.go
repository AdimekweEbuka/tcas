package main

//Airport struct contains templates fo our airports
type Airport struct {
	name string
}

//Airplane struct contains info for airplanes
type airplane struct {
	Name      string
	CallSign  string
	Bearing   int
	Altitude  int
	HasTaxied bool
}
