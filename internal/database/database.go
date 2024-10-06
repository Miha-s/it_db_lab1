package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Miha-s/it_db_lab1/internal/database/attributes"
)

type Database struct {
	name         string
	storage_path string
	tables       []*Table
}

func NewDatabase(storage_path string, name string) (*Database, error) {
	db := &Database{
		name:         name,
		storage_path: storage_path + "/" + name,
	}

	err := os.MkdirAll(db.storage_path, 0755)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func LoadFromStorage(storage_path string, name string) (*Database, error) {
	db, err := NewDatabase(storage_path, name)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(db.storage_path)
	if err != nil {
		return nil, fmt.Errorf("could not read storage directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && file.Type() == 0 {
			fileName := file.Name()
			if len(fileName) < 5 || fileName[len(fileName)-4:] != ".csv" {
				continue
			}

			tableName := fileName[:len(fileName)-4]
			table := LoadFromFile(db.storage_path, tableName)
			if table == nil {
				return nil, fmt.Errorf("failed to load table from file: %v", fileName)
			}

			db.tables = append(db.tables, table)
		}
	}

	return db, nil
}

func LoadDatabases(storage_path string) ([]*Database, error) {
	files, err := os.ReadDir(storage_path)
	var databases []*Database

	if err != nil {
		return nil, fmt.Errorf("could not read storage directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			new_db, err := LoadFromStorage(storage_path, file.Name())
			if err != nil {
				log.Printf("failed to load database - %v", file.Name())
			} else {
				databases = append(databases, new_db)
			}
		}
	}

	return databases, nil
}

func (db *Database) CreateTable(name string, attributes []attributes.Attribute) error {
	if db.GetTable(name) != nil {
		return fmt.Errorf("%v table already exists", name)
	}

	db.tables = append(db.tables, NewTable(db.storage_path, name, attributes))
	return nil
}

func (db *Database) RemoveTable(name string) error {
	for index, table := range db.tables {
		if table.Name() == name {
			table.Delete()
			db.tables = append(db.tables[:index], db.tables[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("%v table not found", name)
}

func (db *Database) GetTable(name string) *Table {
	for _, table := range db.tables {
		if table.Name() == name {
			return table
		}
	}

	return nil
}
