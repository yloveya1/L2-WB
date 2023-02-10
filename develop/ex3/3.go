/*Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим
описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fflags struct {
	sortCol int
	sortNum bool
	sortRev bool
	notDupl bool
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//// writeLines writes the lines to the given file.txt.
//func writeLines(lines []string, path string) error {
//	file, err := os.Create(path)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	w := bufio.NewWriter(file)
//	for _, line := range lines {
//		fmt.Fprintln(w, line)
//	}
//	return w.Flush()
//}

func sortCol(flags fflags, max int, splitStr [][]string) {
	k := flags.sortCol
	if k < 1 {
		log.Fatalf("sort: -k %d: Invalid argument", k)
	} else if k > 0 {

		if k > max {
			k = 1
		}

		sort.Slice(splitStr, func(i, j int) bool {

			if flags.sortRev {
				return strings.Join(splitStr[i][k-1:], " ") > strings.Join(splitStr[j][k-1:], " ")
			}
			return strings.Join(splitStr[i][k-1:], " ") < strings.Join(splitStr[j][k-1:], " ")

		})

	}
}

func main() {
	flags := fflags{}
	flag.IntVar(&flags.sortCol, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&flags.sortRev, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&flags.notDupl, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&flags.sortNum, "n", false, "сортировать по числовому значению")
	flag.Parse()
	filename := flag.Arg(0)
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var strFile []string

	for _, line := range lines {
		strFile = append(strFile, line)
	}

	if flags.notDupl {
		var uniqueFile []string
		var temp string
		for _, str := range strFile {
			if str != temp {
				uniqueFile = append(uniqueFile, str)
			}
			temp = str
		}
		strFile = uniqueFile
	}

	max := 0
	var splitStr [][]string

	for _, v := range strFile {
		str := strings.Fields(v)
		splitStr = append(splitStr, str)
		if len(str) > max {
			max = len(str)
		}
	}

	if flags.sortNum {
		sort.Slice(splitStr, func(i, j int) bool {
			f, _ := strconv.ParseFloat(splitStr[i][0], 64)
			s, _ := strconv.ParseFloat(splitStr[j][0], 64)
			return f < s
		})
	} else {
		sortCol(flags, max, splitStr)

	}
	//splitStr := sortCol(flags, strFile)

	for _, v := range splitStr {
		fmt.Println(strings.Join(v, " "))
	}

}
