package main

import (
	"fmt"
	"os"
	"path/filepath"

	"diary/auth"
	"diary/diary"
	"diary/ui"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	baseDir := filepath.Join(homeDir, ".diary")
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		fmt.Println("Error creating application directory:", err)
		return
	}

	usersFile := filepath.Join(baseDir, "users.json")
	userStore, err := auth.NewUserStore(usersFile)

	if err != nil {
		fmt.Println("Error initializing user store:", err)
		return
	}

	username, err := auth.HandleUserAuth(userStore)
	if err != nil {
		fmt.Println("Exiting:", err)
		return
	}

	diaryFile := filepath.Join(baseDir, fmt.Sprintf("%s-diary.json", username))

	diaryApp, err := diary.NewDiary(diaryFile)
	if err != nil {
		fmt.Println("Error initializing diary:", err)
		return
	}

	fmt.Printf("Welcome to your Daily Diary, %s!\n", username)
	fmt.Printf("Diary data stored at: %s\n", diaryFile)

	ui.StartInteractiveMenu(diaryApp)
}
