package undian

import (
	"context"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *service) ExportExcel(ctx context.Context) (string, error) {

	data, err := s.LihatSemuaPemenang(ctx)
	if err != nil {
		return "", err
	}
	mainSheet := "List Pemenang"
	i := 2
	f := excelize.NewFile()
	f.NewSheet(mainSheet)
	f.SetActiveSheet(2)
	f.DeleteSheet("Sheet1")
	f.SetCellValue(mainSheet, "A1", "No")
	f.SetCellValue(mainSheet, "B1", "Zona")
	f.SetCellValue(mainSheet, "C1", "Kategori")
	f.SetCellValue(mainSheet, "D1", "Nomor Tiket")
	f.SetCellValue(mainSheet, "E1", "Nama Toko")

	for _, zona := range data {
		for _, kategori := range zona.ZonaPemenang {
			for _, pemenang := range kategori.Pemenang {
				f.SetCellValue(mainSheet, fmt.Sprint("A", i), i-1)
				f.SetCellValue(mainSheet, fmt.Sprint("B", i), zona.Zona)
				f.SetCellValue(mainSheet, fmt.Sprint("C", i), kategori.Kategori)
				f.SetCellValue(mainSheet, fmt.Sprint("D", i), pemenang.Tiket)
				f.SetCellValue(mainSheet, fmt.Sprint("E", i), pemenang.NamaToko)
				i++
			}
		}
	}

	//adjust width
	f.SetColWidth(mainSheet, "A", "A", 4)
	f.SetColWidth(mainSheet, "B", "B", 22.3)
	f.SetColWidth(mainSheet, "C", "C", 10)
	f.SetColWidth(mainSheet, "D", "D", 25)
	f.SetColWidth(mainSheet, "E", "E", 15)

	//format table
	f.AddTable(mainSheet, "A1", fmt.Sprintf("E%d", i-1), `{
		"table_name": "table",
		"table_style": "TableStyleMedium17",
		"show_first_column": true,
		"show_last_column": false,
		"show_row_stripes": true,
		"show_column_stripes": false
	}`)

	buf, _ := f.WriteToBuffer()
	contents := buf.String()
	fmt.Println(contents)

	return contents, err
}
