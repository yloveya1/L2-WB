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
)

type fflags struct {
	sortCol int
	sortNum bool
	sortRev bool
	notDupl bool
}

func main() {
	flags := fflags{}
	flags.sortCol = *flag.Int("k", 1, "указание колонки для сортировки")
	flags.sortNum = *flag.Bool("n", false, "сортировать по числовому значению")
	flags.sortRev = *flag.Bool("r", false, "сортировать в обратном порядке")
	flags.notDupl = *flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	filename := flag.Arg(0)
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}

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

// writeLines writes the lines to the given file.txt.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
