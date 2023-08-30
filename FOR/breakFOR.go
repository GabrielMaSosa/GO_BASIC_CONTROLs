package main
import "fmt"
func main(){
	i:=0
	//el break sale del for especial para buscar algo y mejorar el rendimiento
	
	for {
		fmt.Println(i)
		
		if i==10{
			break
		}
		i++;

	}



}