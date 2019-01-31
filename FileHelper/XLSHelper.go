package FileHelper

//2018-09-14 15:50 github.com/extrame/xls 对xls文件解析存在问题
import (
	"github.com/extrame/xls"
)

type XLSHelper struct {
	XLSFile *xls.WorkBook
}

func NewXLSHelper(filePath string) *XLSHelper {
	helper := new(XLSHelper)
	helper.XLSFile = helper.open(filePath)

	return helper
}

func (this *XLSHelper) open(filePath string) *xls.WorkBook {
	xlsFile, err := xls.Open(filePath, "utf-8")

	if err != nil {
		return nil
	}

	return xlsFile
}

//根据sheet索引获取sheet名，sheet索引从0开始
func (this *XLSHelper) GetSheetName(sheetIndex int) string {
	return this.XLSFile.GetSheet(sheetIndex).Name
}

//获取单元格值
func (this *XLSHelper) GetCellValue(sheetIndex int, rowIndex int, colIndex int) string {
	return this.XLSFile.GetSheet(sheetIndex).Row(rowIndex).Col(colIndex)
}

func (this *XLSHelper) GetSheet(sheetIndex int) *xls.WorkSheet {
	return this.XLSFile.GetSheet(sheetIndex)
}

func (this *XLSHelper) GetSheetMaxRow(sheetIndex int) uint16 {
	return this.GetSheet(sheetIndex).MaxRow
}

//获取指定表页所有行
func (this *XLSHelper) GetRows(sheetIndex int) [][]string {
	collection := make([][]string, 0)

	curSheet := this.XLSFile.GetSheet(sheetIndex)

	for rowIndex := 0; rowIndex <= (int(curSheet.MaxRow)); rowIndex++ {
		row := curSheet.Row(rowIndex)

		if row == nil {
			collection = nil
			break
		}

		cells := make([]string, 0)
		for colIndex := 0; colIndex <= (row.LastCol() + 1); colIndex++ {
			cells = append(cells, row.Col(colIndex))
		}

		collection = append(collection, cells)
	}

	return collection
}
