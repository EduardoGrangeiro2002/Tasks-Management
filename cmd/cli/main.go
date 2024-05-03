package main

import (
	"log"
	"tasks/cmd/cli/services"
	"tasks/cmd/cli/usecase"

	"github.com/spf13/cobra"
)

type Task struct {
	ID int
	Name string
	Project string
	Classification string
	StartDate string
	EndDate string
	Closed bool
}

func main() {
	services.CreateDirIfNotExist()
	initializeDatabase()
	var rootCmd = &cobra.Command{Use: "task"}
	var cmdInit = usecase.InitTask()
	var cmdFinish = usecase.FinishTask()
	var cmdLs = usecase.LsTask()
	var cmdRm = usecase.RmTask()
	rootCmd.AddCommand(cmdInit, cmdFinish, cmdLs, cmdRm)
	rootCmd.Execute()
}


func initializeDatabase() {
	db := services.Db()
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		project VARCHAR(255),
		classification VARCHAR(255),
		start_date DATETIME NOT NULL,
		end_date DATETIME,
		closed BOOLEAN NOT NULL
	)`
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}


