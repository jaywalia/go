package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func createEmptyFile() {
	ef, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer ef.Close()

	log.Println(ef)
}

func createDirectory(dir string) {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		panic(err)
	}
}

func fileMetadata(file string) {
	fileStat, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()

}

func readFile(fileName string) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(buf)
	data := bufio.NewScanner(strings.NewReader(contents))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}
}

func appendContent(file, content string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", content)
}

func fileExists(file string) {
	// Test File existence.
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	fmt.Printf("%s File exists!\n", file)
}

func main() {
	//createEmptyFile()
	//createDirectory("test_dir")
	//fileMetadata("files.go")
	//readFile("files.go")
	// appendContent("empty.txt", "Not so empty anymore!")
	fileExists("empty.txt")
}
