package database

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/Miha-s/it_db_lab1/internal/database/attributes"
)

type AcceptRow func([]string) bool

type Table struct {
	storage_path string
	name         string
	attributes   []attributes.Attribute
	rows         [][]string
}

func NewTable(storage_path string, name string, attributes []attributes.Attribute) *Table {
	table := &Table{
		storage_path: storage_path,
		name:         name,
		attributes:   attributes,
	}

	table.saveTable()
	return table
}

func LoadFromFile(storage_path string, name string) *Table {
	t := &Table{
		storage_path: storage_path,
		name:         name,
	}
	filePath := t.storage_path + "/" + t.name + ".csv"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if len(records) < 2 {
		panic("CSV file does not contain enough rows to load table")
	}

	dataTypes := records[0]
	columnNames := records[1]

	t.attributes = make([]attributes.Attribute, len(columnNames))
	for i, name := range columnNames {
		t.attributes[i] = attributes.CreateAttribute(dataTypes[i], name)
	}

	t.rows = records[2:]

	return t
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) Attributes() []attributes.Attribute {
	return t.attributes
}

func (t *Table) AddRow(row []string) error {
	if len(row) != len(t.attributes) {
		return errors.New("invalid number of arguments")
	}

	if err := t.validateRow(row); err != nil {
		return err
	}

	t.rows = append(t.rows, row)

	return nil
}

func (t *Table) UpdateRow(row []string, row_to_update AcceptRow) error {
	if len(row) != len(t.attributes) {
		return errors.New("invalid number of arguments")
	}

	if err := t.validateRow(row); err != nil {
		return err
	}

	for index, value := range t.rows {
		if row_to_update(value) {
			t.rows[index] = row
		}
	}

	return nil
}

func (t *Table) Sync() {
	t.saveTable()
}

func (t *Table) Delete() error {
	t.rows = [][]string{}

	filePath := t.storage_path + "/" + t.name + ".csv"
	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return nil
}

func (t *Table) RemoveDuplicates() {
	rowMap := make(map[string]bool)
	var uniqueRows [][]string

	for _, row := range t.rows {
		rowStr := fmt.Sprintf("%v", row)

		if !rowMap[rowStr] {
			uniqueRows = append(uniqueRows, row)
			rowMap[rowStr] = true
		}
	}

	t.rows = uniqueRows
}

func (t *Table) validateRow(row []string) error {
	for index, value := range row {
		err := t.attributes[index].Validate(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Table) saveTable() {
	filePath := t.storage_path + "/" + t.name + ".csv"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	dataTypes := make([]string, len(t.attributes))
	for i, attr := range t.attributes {
		dataTypes[i] = attr.Type()
	}
	if err := writer.Write(dataTypes); err != nil {
		panic(err)
	}

	columnNames := make([]string, len(t.attributes))
	for i, attr := range t.attributes {
		columnNames[i] = attr.Name()
	}
	if err := writer.Write(columnNames); err != nil {
		panic(err)
	}

	for _, row := range t.rows {
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}
