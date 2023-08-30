package main
import "fmt"

func main(){

	
	//fallthrough ejecuta el otro case por mas que no sea cierto !!!
	//Cuidado!!
	switch dias:="Lunes";dias{

	case "Lunes","Martes","Miercoles":
		fmt.Println("Son los primeros 3 dias")
		fallthrough
	case "Jueves","Viernes":
		fmt.Println("Son lo 2 ultimos dias")

	
default:
	fmt.Println("Es Finde de semana")
	}
}
