package auth

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserStore struct {
	Users    map[string]User `json:"users"`
	Filename string          `json:"-"`
}

func NewUserStore(filename string) (*UserStore, error) {
	store := &UserStore{
		Users:    make(map[string]User),
		Filename: filename,
	}

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open users file: %w", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&store.Users); err != nil {
			return nil, fmt.Errorf("failed to decode users data: %w", err)
		}
	}

	return store, nil
}

func (s *UserStore) Save() error {
	file, err := os.Create(s.Filename)
	if err != nil {
		return fmt.Errorf("failed to create users file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.Users); err != nil {
		return fmt.Errorf("failed to encode users data: %w", err)
	}

	return nil
}

func (s *UserStore) Register(username, password string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	if _, exists := s.Users[username]; exists {
		return fmt.Errorf("username '%s' already exists", username)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	s.Users[username] = User{
		Username:     username,
		PasswordHash: string(hash),
	}

	return s.Save()
}

func (s *UserStore) Authenticate(username, password string) bool {
	user, exists := s.Users[username]
	if !exists {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

func PromptCredentials(reader *bufio.Reader, prompt string) (string, string, error) {
	fmt.Println(prompt)

	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	password = strings.TrimSpace(password)

	return username, password, nil
}

func HandleUserAuth(userStore *UserStore) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== Daily Diary Authentication =====")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		option = strings.TrimSuffix(option, "\n")
		option = strings.TrimSuffix(option, "\r")

		switch option {
		case "1": // Login
			username, password, err := PromptCredentials(reader, "Please enter your login credentials:")
			if err != nil {
				return "", err
			}

			if !userStore.Authenticate(username, password) {
				fmt.Println("Invalid username or password. Please try again.")
				continue
			}
			fmt.Printf("User '%s' logged in successfully!\n", username)
			return username, nil

		case "2": // Register
			username, password, err := PromptCredentials(reader, "Please create your account:")
			if err != nil {
				return "", err
			}

			if err := userStore.Register(username, password); err != nil {
				fmt.Printf("Registration failed: %s\n", err)
				continue
			}

			fmt.Printf("User '%s' registered successfully!\n", username)
			return username, nil

		case "3": // Exit
			return "", errors.New("user exited")

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
