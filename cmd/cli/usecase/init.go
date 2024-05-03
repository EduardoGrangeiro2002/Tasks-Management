package usecase

import (
	"fmt"
	"log"
	"tasks/cmd/cli/services"
	"time"

	"github.com/spf13/cobra"
)

func InitTask() *cobra.Command {
	cmdInit := &cobra.Command{
		Use: "init", 
		Short: "Initialize a task", 
		Run: func (cmd *cobra.Command, args []string)  {
			query := "INSERT INTO tasks (name, project, start_date, end_date, closed) VALUES (?, ?, ?, ?, ?)"
			db := services.Db()
			defer db.Close()
			name, _ := cmd.Flags().GetString("name")
			project, _ := cmd.Flags().GetString("project")
			stmt, err := db.Prepare(query)

			if err != nil {
				log.Fatal(err)
			}
			now := time.Now()
			dateTimeString := now.Format("2006-01-02 15:04:05")
			_, err = stmt.Exec(name, project, dateTimeString, nil, false)

			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			fmt.Printf("Task %s created with success in project %s\n", name, project)
		},
	}
	cmdInit.Flags().String("name", "", "Name of the task")
	cmdInit.Flags().String("project", "", "Project of the task")
	cmdInit.MarkFlagRequired("name")

	return cmdInit
}
