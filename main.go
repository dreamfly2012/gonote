package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var db *sql.DB

//RootCommand default command
var RootCommand = &cobra.Command{
	Use:   "go-note",
	Short: "go-note is note cli",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// WriteContentSubCommand write content command
var WriteContentSubCommand = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
	},
}

// CreateSubCommand create sub command
var CreateSubCommand = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println(len(args))
		if len(args) >= 1 {
			log.Println("insert note title")
			//insertNote(args[0])

			insertContent()
		}
	},
}

func insertContent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter note content: ")
	content, _ := reader.ReadString('\x10')
	fmt.Println(content)
}

func init() {
	RootCommand.AddCommand(CreateSubCommand)
	// RootCommand.AddCommand(TestSubCommand)

	// RootCommand.Flags().BoolP("update", "u", false, "")
	db, err := sql.Open("sqlite3", "./note.db")
	if err != nil {
		panic(err)
	}

	var count int
	err = db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='note'").Scan(&count)
	if err != nil {
		log.Fatalln("query select failed", err)
	}
	if count == 0 {
		_, err = db.Exec("create table note (id integer, title varchar(30), content TEXT)")
		if err != nil {
			log.Fatal("Failed to create table:", err)
		}
	}
}

func insertNote(title string) {
	db, err := sql.Open("sqlite3", "./note.db")
	if err != nil {
		panic(err)
	}
	stmt, _ := db.Prepare("INSERT INTO note (title) VALUES (?)")
	stmt.Exec(title)
	defer stmt.Close()
	//    db.Exec("insert into note (title) values ()")
	//	rows, err := db.Query("select id, title, content from note order by id desc")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer rows.Close()
	//	for rows.Next() {
	//		var id int64
	//		var title string
	//		var content string
	//		if err := rows.Scan(&id, &title, &content); err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Printf("id=%d title=%s content=%s\n", id, title, content)
	//	}
}

func main() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
