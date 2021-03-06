// Создать Go функцию, осуществляющую примитивную распаковку
// строки, содержащую повторяющиеся символы / руны, например:
// ● "a4bc2d5e" => "aaaabccddddde"
// ● "abcd" => "abcd"
// ● "45" => "" (некорректная строка)
// ● "" => ""
// Дополнительное задание: поддержка escape -
// последовательностей
// ● qwe\4\5 => qwe45 (*)
// ● qwe\45 => qwe44444 (*)
// ● qwe\\5 => qwe\\\\\ (*)
// В случае если была передана некорректная строка функция
// должна возвращать ошибку. Написать unit-тесты.
package main

import (
	"fmt"
)


func main() {
	// fmt.Println(Unpack(`a3bc2d2e1`))
	// Unpack("22")
	fmt.Println(Unpack(`e\\3a`))
}

