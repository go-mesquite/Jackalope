package jackalope

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sync"
)

// jackalopeDB is a json relational database
type jackalopeDB struct {
	path       string     // Where the outermost database file is located
	json_path  string     // Where the main json database is stored
	files_path string     // Where all the files are stored
	mu         sync.Mutex // Lock and unlock the database for now
}

// Make sure the files exist and are intact
func validateDB(db *jackalopeDB) error {
	// Check if the database folder itself is there first
	_, err := os.Stat(db.path)
	if err != nil {
		return err
	}
	fmt.Println("Database found! Validating...")
	_, err = os.Stat(db.json_path)
	if err != nil {
		return fmt.Errorf(db.json_path, " is missing, cannot initialize the database", err)
	}
	_, err = os.Stat(db.files_path)
	if err != nil {
		return fmt.Errorf(db.files_path, " is missing, cannot initialize the database", err)
	}
	// TODO return an integrity error if the database schema does not match the files
	// Use the vars trying to open the files above for this
	return nil
}

// Create the files for a database
func createDB(db *jackalopeDB) error {
	// Make the directories first
	err := os.MkdirAll(db.files_path, os.ModePerm)
	if err != nil {
		return err
	}

	// Make the json file
	file, err := os.Create(db.json_path)
	if err != nil {
		// Return an error if file creation fails
		return fmt.Errorf("failed to create database file: %v", err)
	}
	defer file.Close() // Close the file when done
	return nil
}

// Creates a new database or validates an existing one
func NewDB(filepath string) (*jackalopeDB, error) {
	db := &jackalopeDB{
		path:       filepath,              // Filepath of the outermost database folder
		json_path:  filepath + "/db.json", // Structured data stored as json
		files_path: filepath + "/files",   // Files referenced by the json are stored in this folder
	}

	// Make sure the files exist and are intact
	err := validateDB(db)
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("No database files found. Creating:", db.path)

		// Create the database
		err = createDB(db)
		if err != nil {
			return nil, err
		}

	} else if err != nil {
		// Catch that the files don't exist and make them
		panic(err)
	}

	// Return the initialized jackalopeDB instance
	return db, nil
}

// TODO start here next time

// Adds a table, represented as a struct, to the database schema
func (db *jackalopeDB) AddTable(table interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Convert the table to its JSON representation
	tableJSON, err := json.Marshal(table)
	if err != nil {
		return fmt.Errorf("failed to marshal table to JSON: %v", err)
	}

	// Read existing JSON data from the database file
	file, err := os.Open(db.json_path)
	if err != nil {
		return fmt.Errorf("failed to open database file: %v", err)
	}
	defer file.Close()

	var existingData []byte
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	if stat.Size() > 0 {
		existingData = make([]byte, stat.Size())
		_, err := file.Read(existingData)
		if err != nil {
			return fmt.Errorf("failed to read existing data: %v", err)
		}
	}

	// Append the new table JSON to the existing data
	newData := append(existingData, tableJSON...)

	// Write the updated data back to the database file
	err = os.WriteFile(db.json_path, newData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write updated data to file: %v", err)
	}

	return nil
}

/*
// Adds the record to the database
func (db *jackalopeDB) Create(model interface{}) error {
	jsonData, err := json.MarshalIndent(db.data, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(db.filename, jsonData, 0644)
}
*/
