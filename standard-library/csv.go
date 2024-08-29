package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func writeCsv() {
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	_ = writer.Write([]string{"Nama", "Alamat", "Okupansi"})
	_ = writer.Write([]string{"Rifqi", "Tangerang", "Backend"})
	_ = writer.Write([]string{"Fadhillah", "Ciledug", "Cloud Engineer"})

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing csv writer:", err)
	} else {
		fmt.Println("CSV written successfully.")
	}
}

func bufIo() {

	file, _ := os.Open("output.csv")
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Print(line) // Each line is a raw string
	}

}

func readCsv() {
	file, _ := os.Open("output.csv")
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	defer fmt.Println("CSV Read Completed")
}

func main() {
	// writeCsv()
	readCsv()
	// bufIo()
}
