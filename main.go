package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var files [1024]string

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for index, e := range entries {
		files[index] = e.Name()
		fmt.Println(strings.Split(files[index], "."))
	}
}

/*"gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=" + '"' + sPDF + '" "' + sPDFA + '"'*/
