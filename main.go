package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: no folders provided")
		fmt.Println("Usage: ./main.exe C:\\folder to read from C:\\folder to write in")
		return
	}

	currentTime := time.Now()
	var files [1024]string

	folderToRead := os.Args[1]
	folderToWrite := os.Args[2]

	if _, err := os.Stat(folderToRead); os.IsNotExist(err) {
		fmt.Printf("%v does not exsist", folderToRead)
		return
	}

	if _, err := os.Stat(folderToWrite); os.IsNotExist(err) {
		fmt.Println("Folder did not exsist, it will be created now")
		os.Mkdir(folderToWrite, 777)
	}

	os.Chdir(folderToRead)

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

		if strings.Contains(files[i], ".pdf") {
			//currentTime.Add(-24 * time.Hour) = currentTime - 24 hours
			if currentTime.Add(-24 * time.Hour).Before(modificationTime) {
				input_file := folderToRead + "\\" + files[i]
				output_file := "PDFA_" + files[i]
				ghostscript := fmt.Sprintf("gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=%v %v", output_file, input_file)
				os.Chdir(folderToWrite)
				cmd := exec.Command("cmd", "/C", ghostscript)
				out, err := cmd.Output()
				if err != nil {
					fmt.Println("could not run command: ", err)
				}
				fmt.Println("Output: \n", string(out))
				//os.Remove(input_file)
			}
			os.Chdir(folderToRead)
		}
	}
}

//"gswin64c -dPDFA -dBATCH -dNOPAUSE -sColorConversionStrategy=UseDeviceIndependentColor -sDEVICE=pdfwrite -dPDFACompatibilityPolicy=2 -sOutputFile=output_filename.pdf input_filename.pdf"
