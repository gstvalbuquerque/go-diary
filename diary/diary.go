package diary

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type DiaryEntry struct {
	Date      string `json:"date"`
	Morning   string `json:"morning"`
	Afternoon string `json:"afternoon"`
	Evening   string `json:"evening"`
}

type Diary struct {
	Entries  map[string]DiaryEntry `json:"entries"`
	Filename string                `json:"-"`
}

func NewDiary(filename string) (*Diary, error) {
	diary := &Diary{
		Entries:  make(map[string]DiaryEntry),
		Filename: filename,
	}

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open diary file: %w", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&diary.Entries); err != nil {
			return nil, fmt.Errorf("failed to decode diary data: %w", err)
		}
	}

	return diary, nil
}

func (d *Diary) Save() error {
	file, err := os.Create(d.Filename)
	if err != nil {
		return fmt.Errorf("failed to create diary file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(d.Entries); err != nil {
		return fmt.Errorf("failed to encode diary data: %w", err)
	}

	return nil
}

func (d *Diary) AddEntry(date, section, content string) error {
	entry, exists := d.Entries[date]
	if !exists {
		entry = DiaryEntry{Date: date}
	}

	switch section {
	case "morning":
		entry.Morning = content
	case "afternoon":
		entry.Afternoon = content
	case "evening":
		entry.Evening = content
	default:
		return fmt.Errorf("invalid section: %s (must be morning, afternoon, or evening)", section)
	}

	d.Entries[date] = entry
	return d.Save()
}

func (d *Diary) GetEntry(date string) (DiaryEntry, bool) {
	entry, exists := d.Entries[date]
	return entry, exists
}

func (d *Diary) ListDates() []string {
	dates := make([]string, 0, len(d.Entries))
	for date := range d.Entries {
		dates = append(dates, date)
	}
	return dates
}

// FormatDate formats the current date as DD-MM-YYYY
func FormatDate(t time.Time) string {
	return t.Format("02-01-2006")
}
