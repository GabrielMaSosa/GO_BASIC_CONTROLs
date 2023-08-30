La palabra reservada range itera por los elementos de estructuras de datos, retornando siempre 2 variables. Tenemos:
Array or slice: primero el índice, segundo el elemento de la lista.
String: primero el índice, segunda es la representación del carácter en rune int. 
Map: primero la key, segundo el value del par clave-valor.
Channel: primero el elemento del canal, segundo está vacío.
Veamos un ejemplo con un array:


frutas := []string{"manzana", "banana", "pera"}
for i, fruta := range frutas {
   fmt.Println(i, fruta)
}
