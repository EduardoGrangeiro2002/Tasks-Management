package usecase

import (
	"fmt"
	"log"
	"tasks/cmd/cli/services"
	"time"

	"github.com/spf13/cobra"
)

func FinishTask() *cobra.Command {
	cmdFinish := &cobra.Command{
		Use:   "finish",
		Short: "Finish a task",
		Run: func(cmd *cobra.Command, args []string) {
			query := "UPDATE tasks SET closed = true, end_date = ? WHERE id = ?"
			db := services.Db()
			defer db.Close()
			id, err := cmd.Flags().GetInt("id")
			if err != nil {
				log.Fatal("Error to finish task")
			}

			stmt, err := db.Prepare(query)
			if err != nil {
				log.Fatal(err)
			}

			now := time.Now()
			dateTimeString := now.Format("2006-01-02 15:04:05")
			_, err = stmt.Exec(dateTimeString, id)
			
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			fmt.Printf("Task %d finished with success\n", id)
		},
	}
	cmdFinish.Flags().Int("id", 0, "ID of the task")
	cmdFinish.MarkFlagRequired("id")

	return cmdFinish
}