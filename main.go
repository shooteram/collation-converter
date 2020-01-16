package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

var collation string

func main() {
	file := flag.String("file", "", "file location to convert columns's collations from")
	flag.StringVar(&collation, "collation", "utf8mb4_unicode_ci", "new collation")
	guid := flag.Bool("without-guid", false, "avoid changing collations of columns with a Doctrine2 \"guid\" comment")
	flag.Parse()

	if "" == *file {
		fmt.Println("a file location is required to proceed\n\n--help")
		flag.PrintDefaults()
		return
	}

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	regex := `\x60([a-z_]+)\x60 [a-z\(\)0-9]+ COLLATE ([a-z0-9_]+)`
	if true == *guid {
		regex = regex + `((.*)\'\(DC2Type:(guid)\)\')?`
	}

	fmt.Printf(
		"%s\n", regexp.MustCompile(regex).ReplaceAllStringFunc(string(data[:]), replace))
}

func replace(text string) string {
	return regexp.MustCompile(`[a-z0-9_]+$`).ReplaceAllString(text, collation)
}
