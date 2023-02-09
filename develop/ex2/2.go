/*Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func StringUnpacking(str string) (string, error) {
	var (
		res   strings.Builder
		s     string
		flag  bool
		count int
	)
	for i, v := range str {
		temp := string(v)
		if temp == `\` {
			count++
		}

		if num, err := strconv.Atoi(temp); err == nil && count != 1 {
			if flag == true || i == 0 {
				return "", errors.New("incorrect string input")
			}
			flag = true
			for ; num > 1; num-- {
				res.WriteString(s)
			}
			continue
		}
		s = temp
		if s != `\` || count == 2 {
			count = 0
			res.WriteString(s)
		}
		flag = false
	}

	return res.String(), nil
}

func main() {
	s, err := StringUnpacking("qwe\\45")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
