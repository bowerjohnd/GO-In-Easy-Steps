package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	//tmpFile, err := ioutil.TempFile("", "Data-*")	// book uses deprecated ioutil
	tmpFile, err := os.CreateTemp("", "Data-*")
	check(err)
	fmt.Printf("\nCreated File:\n%v \n", tmpFile.Name())

	nb, err := tmpFile.WriteString("Go Programming Fun!\n")
	check(err)

	txt, err := os.ReadFile(tmpFile.Name()) // book uses deprecated ioutil
	check(err)
	fmt.Printf("\nRead: %v bytes - %v \n", nb, string(txt))

	tmpFile.Close()

	fmt.Println("Removing", tmpFile.Name())
	if os.Remove(tmpFile.Name()) != nil {
		fmt.Println("error removing")
	} else {
		fmt.Println("Removal succeeded")
	}

	_, err = os.Stat(tmpFile.Name())
	check(err)
}
