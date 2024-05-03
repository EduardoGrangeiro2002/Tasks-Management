package main

import (
	"log"
	"tasks/cmd/cli/services"
	"tasks/cmd/cli/usecase"

	"github.com/spf13/cobra"
)

/*
	Objetivo desse projeto é criar um CLI de gerenciamento de Tarefas para desenvolvedores.
	Através desse CLI, o desenvolvedor poderá criar, listar, atualizar e deletar tarefas.
	Por exemplo:
		task init -n "Criar um novo projto" -p "Projeto" -> Cria uma nova tarefa
		task ls -> Lista todas as tarefas pode se passar a data e ele filtrara pelo dia. Por default ele lista todas as tarefas
		task rm -i #id -> remove a tarefa da lista
		task finish -i #id -> Finaliza a tarefa marcando a data final e excluindo do task ls
*/

type Task struct {
	ID int
	Name string
	Project string
	StartDate string
	EndDate string
	Closed bool
}

func main() {
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
		start_date DATETIME NOT NULL,
		end_date DATETIME,
		closed BOOLEAN NOT NULL
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}


