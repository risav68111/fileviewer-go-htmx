package services

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

// ExcelSheet represents a sheet in Excel file
type ExcelSheet struct {
	SheetName string                   `json:"sheetName"`
	Headers   []string                 `json:"headers"`
	Data      []map[string]interface{} `json:"data"`
	RowCount  int                      `json:"rowCount"`
}

// Convert Excel to JSON structure
func ExcelToJSON(filePath string) ([]ExcelSheet, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %v", err)
	}
	defer f.Close()

	// Get all sheet names
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in Excel file")
	}

	var excelSheets []ExcelSheet

	// Process each sheet
	for _, sheetName := range sheets {
		sheet, err := processSheet(f, sheetName)
		if err != nil {
			continue // Skip sheets with errors
		}
		excelSheets = append(excelSheets, *sheet)
	}

	return excelSheets, nil
}

func processSheet(f *excelize.File, sheetName string) (*ExcelSheet, error) {
	// Get all rows
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet %s: %v", sheetName, err)
	}

	if len(rows) == 0 {
		return &ExcelSheet{
			SheetName: sheetName,
			Headers:   []string{},
			Data:      []map[string]interface{}{},
			RowCount:  0,
		}, nil
	}

	// First row is headers
	headers := make([]string, len(rows[0]))
	for i, header := range rows[0] {
		if strings.TrimSpace(header) == "" {
			headers[i] = fmt.Sprintf("Column %d", i+1)
		} else {
			headers[i] = strings.TrimSpace(header)
		}
	}
	
	// Process data rows
	var data []map[string]interface{}
	for i := 1; i < len(rows); i++ {
		row := make(map[string]interface{})
		for j, cell := range rows[i] {
			header := headers[j]
			row[header] = strings.TrimSpace(cell)
		}
		data = append(data, row)
	}

	return &ExcelSheet{
		SheetName: sheetName,
		Headers:   headers,
		Data:      data,
		RowCount:  len(data),
	}, nil
}

// Keep original function for backward compatibility
func XLSXToTable(filePath string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
