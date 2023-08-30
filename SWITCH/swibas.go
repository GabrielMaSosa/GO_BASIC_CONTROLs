package main
import ("fmt")

func main(){

	//var dias []string={"Lunes","Martes","Miercoles","Jueves","Viernes"}
	var dias string="Domin"

	switch dias{

	case "Lunes":
		fmt.Println("Es Lunes")
	case "Martes":
		fmt.Println("Es Martes")
	case "Miercoles":
		fmt.Println("Es Miercoles")
	case "Jueves":
		fmt.Println("Es Jueves")
	case "Viernes":
		fmt.Println("Es Viernes")
	default:
		fmt.Println("ES Finde de semana")


	}









}