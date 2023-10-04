package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var files [1024]string

	cmd := exec.Command("ls", "./")
	out, err := cmd.Output()

	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: \n", string(out))

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for index, e := range entries {
		files[index] = e.Name()
		fmt.Println(strings.Split(files[index], "."))
	}

}

//gs -dPDFA -dBATCH -dNOPAUSE -dUseCIEColor -sProcessColorModel=DeviceCMYK -sDEVICE=pdfwrite -sPDFACompatibilityPolicy=1 -sOutputFile=output_filename.pdf input_filename.pdf
/*"gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=" + '"' + sPDF + '" "' + sPDFA + '"'*/
