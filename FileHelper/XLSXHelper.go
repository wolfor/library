package FileHelper

import (
	"github.com/tealeg/xlsx"
)

type XLSXHelper struct {
	XLSXFile *xlsx.File
}

func NewXLSXHelper(filePath string) *XLSXHelper {
	helper := new(XLSXHelper)
	helper.XLSXFile = helper.open(filePath)

	return helper
}

func (this *XLSXHelper) open(filePath string) *xlsx.File {
	xlsx, err := xlsx.OpenFile(filePath)

	if err != nil {
		return nil
	}

	return xlsx
}

//根据sheet索引获取sheet名，sheet索引从1开始
func (this *XLSXHelper) GetSheetName(sheetIndex int) string {
	if sheetIndex < 0 && sheetIndex > len(this.XLSXFile.Sheets) {
		return ""
	}

	return this.XLSXFile.Sheets[sheetIndex].Name
}

func (this *XLSXHelper) MaxCol(sheetIndex int) int {
	if sheetIndex < 0 && sheetIndex > len(this.XLSXFile.Sheets) {
		return -1
	}

	return this.XLSXFile.Sheets[sheetIndex].MaxCol
}

//获取单元格值
func (this *XLSXHelper) GetCellValue(sheetIndex int, rowIndex int, colIndex int) string {
	return this.XLSXFile.Sheets[sheetIndex].Cell(rowIndex, colIndex).String()
}

//获取指定表页所有行
func (this *XLSXHelper) GetRows(sheetIndex int) [][]string {
	if sheetIndex < 0 && sheetIndex > len(this.XLSXFile.Sheets) {
		return nil
	}

	collectData := make([][]string, 0)

	sheet := this.XLSXFile.Sheets[sheetIndex]

	for _, row := range sheet.Rows {

		rowData := make([]string, 0)

		for _, cell := range row.Cells {
			text := cell.String()

			rowData = append(rowData, text)
		}

		collectData = append(collectData, rowData)
	}

	return collectData
}
