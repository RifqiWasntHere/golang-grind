package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, content string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(content)

	return nil
}

// Buatkan fungsi untuk read file, WElllllll
func readFile(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var content string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		content += string(line)
	}
	return content, err
}

// Bangg kalo mau nambahin konten ke existing file gimana bang ? siyaaaaaap
func appendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(content)

	return nil
}

func main() {
	// createNewFile("Biodata", "Namaku Rifqi AHIHIHIHIHIH")

	appendToFile("biodata.txt", "Tapi boong cihuuuuuy")

	result, err := readFile("biodata.txt")
	fmt.Println(result, err)
}
