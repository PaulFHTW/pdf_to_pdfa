package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no folder provided")
		return
	}

	folderName := os.Args[1]
	var files [1024]string

	entries, err := os.ReadDir(folderName)

	if err != nil {
		log.Fatal(err)
	}

	for i, e := range entries {
		files[i] = e.Name()
		fmt.Println(files[i])
		ghostscript := fmt.Sprintf("gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=%v %v", files[i], files[i])

		cmd := exec.Command("cmd", "/C", ghostscript)
		//fmt.Println(cmd)
		out, err := cmd.Output()

		if err != nil {
			fmt.Println("could not run command: ", err)
		}
		fmt.Println("Output: \n", string(out))
	}
}

//gs -dPDFA -dBATCH -dNOPAUSE -dUseCIEColor -sProcessColorModel=DeviceCMYK -sDEVICE=pdfwrite -sPDFACompatibilityPolicy=1 -sOutputFile=output_filename.pdf input_filename.pdf
/*"gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=" + '"' + sPDF + '" "' + sPDFA + '"'*/
//"gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=output_filename.pdf input_filename.pdf"
