package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow { /* Değer dizisinin indislerini ya da o indise karşılık gelen değerleri _ ataması ile atlayabilirsiniz.*/

		/*Eğer yalnızca indisi kullanmak istiyorsanız ", value" çıkar */
		fmt.Printf("%d\n", value)
	}
}

