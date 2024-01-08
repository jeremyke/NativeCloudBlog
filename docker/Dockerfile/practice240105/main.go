package main

import (
	"fmt"
	"github.com/shakinm/xlsReader/xls"
	"log"
	"strconv"
)

func main() {

	workbook, err := xls.OpenFile("./file.xls")

	if err != nil {
		log.Panic(err.Error())
	}

	sheet, err := workbook.GetSheet(0)

	if err != nil {
		log.Panic(err.Error())
	}

	for i := 0; i <= sheet.GetNumberRows(); i++ {
		row, rowErr := sheet.GetRow(i)
		if rowErr != nil {
			continue
		}
		fmt.Println("-----第" + strconv.FormatInt(int64(i), 10) + "行----------")
		for j := 0; j < len(row.GetCols()); j++ {
			cell, cellErr := row.GetCol(j)
			if cellErr != nil {
				continue
			}
			//fmt.Println("类型：", cell.GetType())
			xfIndex := cell.GetXFIndex()
			formatIndex := workbook.GetXFbyIndex(xfIndex)
			format := workbook.GetFormatByIndex(formatIndex.GetFormatIndex())
			fmt.Println(format.GetFormatString(cell))
		}

	}
}
