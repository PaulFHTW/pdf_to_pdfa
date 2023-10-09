package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	currentTime := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Error: no folder provided")
		return
	}

	folderName := os.Args[1]

	var files [1024]string

	os.Chdir(folderName)
	entries, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for i, e := range entries {
		files[i] = e.Name()

		fileInfo, err := os.Stat(files[i])
		if err != nil {
			log.Fatal(err)
		}
		// Gives the modification time
		modificationTime := fileInfo.ModTime()

		//now.Add(-24 * time.Hour) = yesterday
		if currentTime.Add(-24 * time.Hour).Before(modificationTime) {
			input_file := files[i]
			output_file := "PDFA_" + files[i]
			ghostscript := fmt.Sprintf("gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=%v %v", output_file, input_file)
			cmd := exec.Command("cmd", "/C", ghostscript)
			out, err := cmd.Output()
			if err != nil {
				fmt.Println("could not run command: ", err)
			}
			fmt.Println("Output: \n", string(out))
		}
	}
}

//"gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=output_filename.pdf input_filename.pdf"
