package main
import ("fmt")

func main(){
	//Podemos hacer switch sin condicion de switch y si en case 
	//Mas lindo que multiples if 
	//var dias []string={"Lunes","Martes","Miercoles","Jueves","Viernes"}
	var dias string="Lunes"

	switch {

	case dias=="Lunes":
		fmt.Println("Es Lunes")
	case dias=="Martes":
		fmt.Println("Es Martes")
	case dias=="Miercoles":
		fmt.Println("Es Miercoles")
	case dias=="Jueves":
		fmt.Println("Es Jueves")
	case dias=="Viernes":
		fmt.Println("Es Viernes")
	default:
		fmt.Println("ES Finde de semana")
	}









}