package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)

	newcsv, err := os.Create("export.csv")
	checkError(err)

	defer newcsv.Close()

	w := csv.NewWriter(newcsv)

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		}

		//Adding three new columns to the record (you can chain the appends, or create new ones)
		record = append(record, "New column", "Another new column")
		record = append(record, "And again a new append")

		// Ranging over the record array
		// for index, col := range record {
		// 	fmt.Println(index, ":", col)
		// }

		writeErr := w.Write(record)
		checkError(writeErr)

		w.Flush()
		if err := w.Error(); err != nil {
			fmt.Println("Error: Nothing Written", err)
		}

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
