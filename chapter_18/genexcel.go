package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func genExcel() {
	// Öffnen Sie die Excel-Datei
	xlFile, err := xlsx.OpenFile("example.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// Suchen Sie das Blatt mit dem Namen "Klamotten"
	sheet, found := xlFile.Sheets["Klamotten"]
	if !found {
		log.Fatalf("Blatt mit dem Namen 'Klamotten' nicht gefunden.")
	}

	// Suchen Sie die Spalte mit dem Namen "Mein" (Annahme: "Mein" ist die Überschrift)
	var columnIndex int
	for index, cell := range sheet.Rows[0].Cells {
		if cell.String() == "Mein" {
			columnIndex = index
			break
		}
	}

	// Durchlaufen Sie die Zeilen und nehmen Sie nur die auf, in denen die Spalte "Mein" den Wert "x" enthält
	var selectedRows [][]string
	for rowIndex, row := range sheet.Rows {
		// Überspringen Sie die Überschriftenzeile
		if rowIndex == 0 {
			continue
		}
		cell := row.Cells[columnIndex]
		if cell.String() == "x" {
			// Fügen Sie die gesamte Zeile (als Zeichenfolgenarray) zur ausgewählten Zeilenliste hinzu
			var rowValues []string
			for _, cell := range row.Cells {
				rowValues = append(rowValues, cell.String())
			}
			selectedRows = append(selectedRows, rowValues)
		}
	}

	// Öffnen Sie die Textdatei für den Ausgabestrom
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Erstellen Sie einen Schreiber für die Textdatei
	writer := bufio.NewWriter(outputFile)

	// Schreiben Sie die ausgewählten Zeilen in die Textdatei
	for _, row := range selectedRows {
		fmt.Fprintln(writer, strings.Join(row, "\t"))
	}

	// Flush und schließen Sie den Schreiber
	writer.Flush()
}
