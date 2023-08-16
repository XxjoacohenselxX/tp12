package main

import (
	"os"
	"testing"
)

func TestSaveAndLoadToCSV(t *testing.T) {
	filename := "test_data.csv"
	defer os.Remove(filename)

	deleteAllRecords()
	createRecord("Source1", 10.5, "Event1")
	createRecord("Source2", 20.0, "Event2")

	err := saveToCSV(filename)
	if err != nil {
		t.Errorf("Error saving to CSV: %v", err)
	}

	err = loadFromCSV(filename)
	if err != nil {
		t.Errorf("Error loading from CSV: %v", err)
	}

	if len(records) != 2 {
		t.Errorf("Expected 2 records after loading from CSV, got %d", len(records))
	}

	// Check record details
	record := records[0]
	if record.PKID != 1 || record.Source != "Source1" || record.Measurement != 10.5 || record.Event != "Event1" {
		t.Errorf("Read record doesn't match expectations")
	}

	record = records[1]
	if record.PKID != 2 || record.Source != "Source2" || record.Measurement != 20.0 || record.Event != "Event2" {
		t.Errorf("Read record doesn't match expectations")
	}
}

func TestCRUDOperations(t *testing.T) {
	// Prueba de creación
	deleteAllRecords()
	createRecord("Source1", 10.5, "Event1")
	createRecord("Source2", 20.0, "Event2")

	if len(records) != 2 {
		t.Errorf("Expected 2 records after creation, got %d", len(records))
	}

	// Prueba de lectura
	record := records[0]
	if record.PKID != 1 || record.Source != "Source1" || record.Measurement != 10.5 || record.Event != "Event1" {
		t.Errorf("Read record doesn't match expectations")
	}

	// Prueba de lectura de todos los registros
	readAllRecords()

	// Prueba de actualización
	updateRecord(2, "UpdatedSource", 15.0, "UpdatedEvent")
	record = records[1]
	if record.Source != "UpdatedSource" || record.Measurement != 15.0 || record.Event != "UpdatedEvent" {
		t.Errorf("Update failed")
	}

	// Prueba de eliminación
	deleteRecord(1)
	if len(records) != 1 {
		t.Errorf("Expected 1 record after deletion, got %d", len(records))
	}

	// Prueba de borrar todos los registros
	deleteAllRecords()
	if len(records) != 0 {
		t.Errorf("Delete all records failed")
	}
}

func TestMain(m *testing.M) {
	filename := "test_data.csv"
	// Crear un archivo de prueba
	file, _ := os.Create(filename)
	file.Close()

	// Ejecutar pruebas
	result := m.Run()

	// Limpiar después de las pruebas
	os.Remove(filename)

	os.Exit(result)
}
