package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Initialize the database
	err := initDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		os.Exit(1)
	}

	// Set up the main command
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Set up subcommand flags
	showID := showCmd.Bool("id", false, "Show entry IDs")
	showCmd.BoolVar(showID, "i", false, "Show entry IDs (short)")
	showTimestamps := showCmd.Bool("timestamps", false, "Show timestamps")
	showCmd.BoolVar(showTimestamps, "t", false, "Show timestamps (short)")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'show', 'edit', or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		entry := strings.Join(addCmd.Args(), " ")
		addEntry(entry)
	case "show":
		showCmd.Parse(os.Args[2:])
		showEntries(*showID, *showTimestamps)
	case "edit":
		editCmd.Parse(os.Args[2:])
		if len(editCmd.Args()) < 2 {
			fmt.Println("Expected ID and new entry text")
			os.Exit(1)
		}
		id := editCmd.Arg(0)
		newEntry := strings.Join(editCmd.Args()[1:], " ")
		editEntry(id, newEntry)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if len(deleteCmd.Args()) != 1 {
			fmt.Println("Expected entry ID")
			os.Exit(1)
		}
		id := deleteCmd.Arg(0)
		deleteEntry(id)
	default:
		fmt.Println("Expected 'add', 'show', 'edit', or 'delete' subcommands")
		os.Exit(1)
	}
}

func addEntry(entry string) {
	id, err := addEntryToDB(entry)
	if err != nil {
		fmt.Println("Error adding entry:", err)
		return
	}
	color.Green("Entry added successfully.")
	color.Yellow("ID: %s", id)
}

func showEntries(showID, showTimestamps bool) {
	entries, err := getEntriesFromDB()
	if err != nil {
		fmt.Println("Error retrieving entries:", err)
		return
	}

	var currentDate string
	for _, entry := range entries {
		date := entry.AddedAt.Format("2006-01-02")
		if date != currentDate {
			if currentDate != "" {
				fmt.Println()
			}
			color.Cyan(date + ":")
			currentDate = date
		}

		if showID && showTimestamps {
			color.Yellow("[%s] - Added: %s, Updated: %s", entry.ID, entry.AddedAt.Format("2006-01-02 15:04:05"), entry.UpdatedAt.Format("2006-01-02 15:04:05"))
			color.White(entry.Text)
		} else if showID {
			color.Yellow("[%s] - %s", entry.ID, entry.Text)
		} else if showTimestamps {
			color.Magenta("Added: %s, Updated: %s", entry.AddedAt.Format("2006-01-02 15:04:05"), entry.UpdatedAt.Format("2006-01-02 15:04:05"))
			color.White(entry.Text)
		} else {
			color.White("- %s", entry.Text)
		}
	}
}

func editEntry(id string, newEntry string) {
	rowsAffected, err := editEntryInDB(id, newEntry)
	if err != nil {
		fmt.Println("Error editing entry:", err)
		return
	}
	if rowsAffected == 0 {
		color.Red("No entry found with ID %s", id)
	} else {
		color.Green("Entry updated successfully.")
	}
}

func deleteEntry(id string) {
	rowsAffected, err := deleteEntryFromDB(id)
	if err != nil {
		fmt.Println("Error deleting entry:", err)
		return
	}
	if rowsAffected == 0 {
		color.Red("No entry found with ID %s", id)
	} else {
		color.Green("Entry deleted successfully.")
	}
}
