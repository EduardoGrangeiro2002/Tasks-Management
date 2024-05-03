package utils

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func GenerateTaskExcelFile(headers []string, tasks [][]string) {
    f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

    index, err := f.NewSheet("Tasks")

    if err != nil {
        log.Fatal(err)
    }

    for i, header := range headers {
        cell, _ := excelize.CoordinatesToCellName(i+1, 1)
        f.SetCellValue("Tasks", cell, header)
    }

    for i, task := range tasks {
        for j, attr := range task {
            cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
            f.SetCellValue("Tasks", cell, attr)
        }
    }

    f.SetActiveSheet(index)

    if err := f.SaveAs("tasks.xlsx"); err != nil {
        fmt.Println(err)
    }
}

