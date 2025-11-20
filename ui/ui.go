package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"diary/diary"
)

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return strings.TrimSpace(text)
}

func displayMenu() {
	fmt.Println("\n===== Daily Diary Application =====")
	fmt.Println("1. Add entry for today")
	fmt.Println("2. Add entry for specific date")
	fmt.Println("3. View entry by date")
	fmt.Println("4. List entries by period")
	fmt.Println("5. List latest entries (7 days)")
	fmt.Println("6. List all entries")
	fmt.Println("7. Delete entry by date")
	fmt.Println("8. Update entry by date")
	fmt.Println("9. Exit")
	fmt.Print("Select an option: ")
}

func StartInteractiveMenu(diaryApp *diary.Diary) {
	for {
		displayMenu()

		choice := readInput("")

		switch choice {
		case "1":
			handleTodayEntry(diaryApp)
		case "2":
			handleSpecificDateEntry(diaryApp)
		case "3":
			viewEntryByDate(diaryApp)
		case "4":
			listEntriesByPeriod(diaryApp)
		case "5":
			listLatestEntries(diaryApp)
		case "6":
			listAllEntries(diaryApp)
		case "7":
			deleteEntryByDate(diaryApp)
		case "8":
			updateEntryByDate(diaryApp)
		case "9":
			fmt.Println("Thank you for using Daily Diary. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func handleTodayEntry(diaryApp *diary.Diary) {
	date := diary.FormatDate(time.Now())
	fmt.Printf("Adding entry for today (%s):\n", date)

	morning := readInput("Morning: ")
	if morning != "" {
		if err := diaryApp.AddEntry(date, "morning", morning); err != nil {
			fmt.Println("Error saving morning entry:", err)
		}
	}

	afternoon := readInput("Afternoon: ")
	if afternoon != "" {
		if err := diaryApp.AddEntry(date, "afternoon", afternoon); err != nil {
			fmt.Println("Error saving afternoon entry:", err)
		}
	}

	evening := readInput("Evening: ")
	if evening != "" {
		if err := diaryApp.AddEntry(date, "evening", evening); err != nil {
			fmt.Println("Error saving evening entry:", err)
		}
	}

	fmt.Println("Entry saved successfully!")
}

func handleSpecificDateEntry(diaryApp *diary.Diary) {
	dateInput := readInput("Enter date (DD-MM-YYYY): ")

	section := strings.ToLower(readInput("Enter section (morning/afternoon/evening): "))
	if section != "morning" && section != "afternoon" && section != "evening" {
		fmt.Println("Invalid section. Must be morning, afternoon, or evening.")
		return
	}

	content := readInput("Enter content: ")
	if err := diaryApp.AddEntry(dateInput, section, content); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Entry saved successfully!")
	}
}

func viewEntryByDate(diaryApp *diary.Diary) {
	dateInput := readInput("Enter date to view (DD-MM-YYYY): ")
	entry, exists := diaryApp.GetEntry(dateInput)

	if !exists {
		fmt.Println("No entry found for that date.")
		return
	}

	fmt.Printf("\n==== Entry for %s ====\n", dateInput)
	fmt.Printf("Morning: %s\n", entry.Morning)
	fmt.Printf("Afternoon: %s\n", entry.Afternoon)
	fmt.Printf("Evening: %s\n", entry.Evening)
}

func listEntriesByPeriod(diaryApp *diary.Diary) {
	startDate := readInput("Enter start date (DD-MM-YYYY): ")
	endDate := readInput("Enter end date (DD-MM-YYYY): ")

	if startDate == "" || endDate == "" {
		fmt.Println("Please enter both start and end dates.")
		return
	}

	if startDate > endDate {
		fmt.Println("Start date must be before end date.")
		return
	}

	dates := diaryApp.ListDates()
	//refact this in the future
	for _, date := range dates {
		if date >= startDate && date <= endDate {
			entry, _ := diaryApp.GetEntry(date)
			fmt.Printf("\n==== Entry for %s ====\n", date)
			fmt.Printf("Morning: %s\n", entry.Morning)
			fmt.Printf("Afternoon: %s\n", entry.Afternoon)
			fmt.Printf("Evening: %s\n", entry.Evening)
		}
	}
}

func listLatestEntries(diaryApp *diary.Diary) {
	dates := diaryApp.ListDates()
	if len(dates) == 0 {
		fmt.Println("No entries found.")
		return
	}

	// get entries for last 7 days
	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	dates = diaryApp.ListDates()
	//refact this in the future
	for _, date := range dates {
		if date >= startDate && date <= endDate {
			entry, _ := diaryApp.GetEntry(date)
			fmt.Printf("\n==== Entry for %s ====\n", date)
			fmt.Printf("Morning: %s\n", entry.Morning)
			fmt.Printf("Afternoon: %s\n", entry.Afternoon)
			fmt.Printf("Evening: %s\n", entry.Evening)
		}
	}
}

func listAllEntries(diaryApp *diary.Diary) {
	dates := diaryApp.ListDates()
	if len(dates) == 0 {
		fmt.Println("No entries found.")
		return
	}

	for _, date := range dates {
		entry, _ := diaryApp.GetEntry(date)
		fmt.Printf("\n==== Entry for %s ====\n", date)
		fmt.Printf("Morning: %s\n", entry.Morning)
		fmt.Printf("Afternoon: %s\n", entry.Afternoon)
		fmt.Printf("Evening: %s\n", entry.Evening)
	}
}

func deleteEntryByDate(diaryApp *diary.Diary) {
	dateInput := readInput("Enter date to delete (DD-MM-YYYY): ")
	if err := diaryApp.DeleteEntry(dateInput); err != nil {
		fmt.Println("Error:", err)
	} else {
	}
}

func updateEntryByDate(diaryApp *diary.Diary) {
	dateInput := readInput("Enter date to update (DD-MM-YYYY): ")
	entry, exists := diaryApp.GetEntry(dateInput)

	if !exists {
		fmt.Println("No entry found for that date.")
		return
	}

	fmt.Printf("\n==== Current Entry for %s ====\n", dateInput)
	fmt.Printf("Morning: %s\n", entry.Morning)
	fmt.Printf("Afternoon: %s\n", entry.Afternoon)
	fmt.Printf("Evening: %s\n", entry.Evening)

	section := strings.ToLower(readInput("\nEnter section to update (morning/afternoon/evening): "))
	if section != "morning" && section != "afternoon" && section != "evening" {
		fmt.Println("Invalid section. Must be morning, afternoon, or evening.")
		return
	}

	var currentContent string
	switch section {
	case "morning":
		currentContent = entry.Morning
	case "afternoon":
		currentContent = entry.Afternoon
	case "evening":
		currentContent = entry.Evening
	}

	fmt.Printf("Current content: %s\n", currentContent)
	newContent := readInput("Enter new content: ")

	if err := diaryApp.AddEntry(dateInput, section, newContent); err != nil {
		fmt.Println("Error updating entry:", err)
	} else {
		fmt.Println("Entry updated successfully!")
	}
}
