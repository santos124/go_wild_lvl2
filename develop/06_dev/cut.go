package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

// Arguments содержит
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем
type Arguments struct {
	f int
	d string
	s bool
}

//проверка символа на цифру
func isDigit(c rune) bool {
	match := false
	if c >= '0' && c <= '9' {
		match = true
	}
	return match
}

//парсинг аргументов из аргументов запуска
func parseArgs() Arguments {
	var args Arguments
	args.d = " "
	for i := 1; i < len(os.Args); i++ {
		count := 0
		if len(os.Args[i]) > 0 && strings.HasPrefix(os.Args[i], "-") {
			if strings.Contains(os.Args[i], "s") {
				args.s = true
				count++
			}
			if strings.Contains(os.Args[i], "d") {
				temp := []rune{}
				temp = append(temp, rune(os.Args[i][strings.Index(os.Args[i], "d") + 1]))
				args.d = string(temp)
				count++
			}
			if strings.Contains(os.Args[i], "f") {
				index := strings.Index(os.Args[i], "f")
				if isDigit(rune(os.Args[i][index + 1])) {
					f, err := strconv.Atoi(os.Args[i][index + 1:])
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					} else {
						args.f = f
					}
				}
			}
			if count > 1 {
				fmt.Println("Ошибка ввода. Попробуйте еще раз, после '-' идет только 1 флаг, для следующего флага нужна новая черточка '-'")
				os.Exit(1)
			}
		}
	}
	return args
}

//перенос текста из ввода консоли в [][]string
func parseArrOfStrings(args Arguments) [][]string {
	strs := make([][]string, 0, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		s := sc.Text()
		if args.s && !strings.Contains(s, args.d){
			continue
		}
		ss := strings.Split(s, args.d)
		ss = append([]string{"basic input"}, ss...)
		strs = append(strs, ss)
	}
	return strs
}

func outputStr(strs [][]string, args Arguments) {
	for i := range strs {
		for j := 0; j < len(strs[i]); j++ { 
			if j == args.f {
				fmt.Println(strs[i][j])
			}
		}
		
	}
}
func main() {
	
	args := parseArgs()
	strs := parseArrOfStrings(args)
	outputStr(strs, args)
	

}