package usecase

import (
	"database/sql"
	"log"
	"strconv"
	"tasks/cmd/cli/services"
	"tasks/utils"
	"time"

	"github.com/spf13/cobra"
)

func ReportTask() *cobra.Command {
	cmdLs := &cobra.Command{
		Use:   "report",
		Short: "Generate excel's tasks report",
		Run: func(cmd *cobra.Command, args []string) {
			var table [][]string
			query := "SELECT id, name, project, classification, start_date, end_date FROM tasks"
			db := services.Db()
			defer db.Close()
			rows, err := db.Query(query)
			if err != nil {
				log.Fatal(err)
			}
			for rows.Next() {
				var id int
				var name string
				var project string
				var classification string
				var startDate string
				var endDate sql.NullString
				var startHour string
				var endHour string
				err = rows.Scan(&id, &name, &project, &classification, &startDate, &endDate)
				if err != nil {
					log.Fatal(err)
				}

				startHour = getTimeParse(startDate)

				if endDate.Valid {
					endHour = getTimeParse(endDate.String)
				} else {
					endHour = ""
				}

				startDate = getDateParse(startDate)
				task := []string{strconv.Itoa(id), name, project, classification, startDate, startHour, endHour}
				table = append(table, task)
			}
			utils.GenerateTaskExcelFile([]string{"ID", "Nome", "Projeto", "Classificação", "Dia", "Hora Inicial", "Hora Final"}, table)
		},
	}

	return cmdLs
}

func getDateParse(date string) string {
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	layoutBrazilian := "02/01/2006"
	formattedDate := dateTime.Format(layoutBrazilian)
	return formattedDate
}

func getTimeParse(date string) string {
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	time := dateTime.Format("15:04")

	return time
}
