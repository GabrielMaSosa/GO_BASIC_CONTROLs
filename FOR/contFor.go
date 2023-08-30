package main
import "fmt"

func main(){
//la sentencia continue salta a la otra iteracion y deja sin ejecutar 
//el codigo de abjo y vuelve a correr el for en la prox iteracion
	for i:=0;i<10;i++{
		if i%2==0{
			fmt.Println(i)
			continue
		}
		fmt.Println(i,"es impar")
	}
}