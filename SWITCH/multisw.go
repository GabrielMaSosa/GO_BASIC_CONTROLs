package main
import "fmt"

func main(){

	

	switch dias:="Domingo";dias{

	case "Lunes","Martes","Miercoles":
		fmt.Println("Son los primeros 3 dias")
	case "Jueves","Viernes":
		fmt.Println("Son lo 2 ultimos dias")

	
default:
	fmt.Println("Es Finde de semana")
	}
}
