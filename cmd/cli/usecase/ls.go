package usecase

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"tasks/cmd/cli/services"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type Task struct {
	ID        int
	Name      string
	Project   string
	Classification string
	StartDate string
	EndDate   string
	Closed    bool
}

func LsTask() *cobra.Command {
	cmdLs := &cobra.Command{
		Use:   "ls",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			query := "SELECT * FROM tasks"
			db := services.Db()
			defer db.Close()
			rows, err := db.Query(query)
			if err != nil {
				log.Fatal(err)
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Project", "Classification", "Start Date", "End Date", "Closed"})

			for rows.Next() {
				var id int
				var name string
				var project string
				var classification string
				var startDate string
				var endDate sql.NullString
				var closed bool
				err = rows.Scan(&id, &name, &project, &classification, &startDate, &endDate, &closed)
				if err != nil {
					log.Fatal(err)
				}
				task := Task{ID: id, Name: name, Project: project, Classification: classification, StartDate: startDate, EndDate: endDate.String, Closed: closed}
				table.Append([]string{strconv.Itoa(task.ID), task.Name, task.Project, task.Classification, task.StartDate, task.EndDate, strconv.FormatBool(task.Closed)})
			}
			table.Render()
		},
	}

	return cmdLs
}