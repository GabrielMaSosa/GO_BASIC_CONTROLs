package main

func suma(valores []float64) {

}
func multiplica(valores []float64) {

}
func divide(valores []float64) {

}
func resta(valores []float64) {

}

func operacionAritmetica(tipo string, valores ...float64) {

	switch tipo {
	case "+":
		suma(valores)
	case "-":
		resta(valores)
	case "*":
		multiplica(valores)
	case "/":
		divide(valores)
	}

}

func main() {
	const (
		suma  = "+"
		resta = "-"
		multi = "*"
		divi  = "/"
	)
	a, b, c, d := 1.0, 2.0, 6.0, 8.0

	operacionAritmetica(suma, a, b, c, d)

}
