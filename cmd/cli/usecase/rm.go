package usecase

import (
	"fmt"
	"log"
	"tasks/cmd/cli/services"

	"github.com/spf13/cobra"
)

func RmTask() *cobra.Command {
	cmdRm := &cobra.Command{
		Use:   "rm",
		Short: "Remove a task",
		Run: func(cmd *cobra.Command, args []string) {
			query := `DELETE FROM tasks WHERE id = ?`
			db := services.Db()
			defer db.Close()
			id, err := cmd.Flags().GetInt("id")
			if err != nil {
				log.Fatal("Error to remove task")
			}

			stmt, err := db.Prepare(query)

			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			stmt.Exec(id) 
			
			fmt.Printf("Task %d removed with success\n", id)
		},
	}

	cmdRm.Flags().Int("id", 0, "ID of the task")
	cmdRm.MarkFlagRequired("id")

	return cmdRm
}