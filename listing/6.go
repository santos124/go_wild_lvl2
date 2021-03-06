// Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее 
// устройство слайсов и что происходит при передачи их в качестве 
// аргументов функции.

package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	// fmt.Printf("1 %p\n", s)
	modifySlice(s)
	fmt.Println(s)
	// fmt.Printf("2 %p\n", s)
}

func modifySlice(i []string) {
	// fmt.Printf("3 %p\n", i)
	i[0] = "3"
	//append возвращает новый указатель, но срез был получен не по указателю, поэтому изменение адреса  черевато потерей будущих изменений в этой функции
	i = append(i, "4")
	// fmt.Printf("4 %p\n", i)
	i[1] = "5"
	i = append(i, "6")
}
