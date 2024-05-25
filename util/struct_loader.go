package util

import (
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v5"
)

func ScanStruct(rows pgx.Rows, dest interface{}) error {
	// Type check to ensure destination is a pointer to struct
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dest must be a pointer to a struct")
	}

	// Get the types of the struct fields
	values := make([]interface{}, len(rows.FieldDescriptions()))
	for i, fd := range rows.FieldDescriptions() {
		// Create a pointer to the struct field to scan into
		field := v.Elem().FieldByName(fd.Name) // Assumes field names match column names
		if field.IsValid() && field.CanSet() {
			values[i] = field.Addr().Interface()
		} else {
			// If field not found or not settable, create a dummy variable
			var dummy interface{}
			values[i] = &dummy
		}
	}

	// Scan the values into the struct fields
	if err := rows.Scan(values...); err != nil {
		return err
	}

	return nil
}

func ScanStructRow(row pgx.Row, dest interface{}) error {
	// Type check to ensure destination is a pointer to a struct.
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dest must be a pointer to a struct")
	}

	// Get the types of the struct fields
	fields := v.Elem()
	values := make([]interface{}, fields.NumField())

	// Create a pointer to each struct field to scan into
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		if field.CanSet() {
			values[i] = field.Addr().Interface()
		} else {
			// If the field is not settable, create a dummy variable
			var dummy interface{}
			values[i] = &dummy
		}
	}

	// Scan the values into the struct fields
	if err := row.Scan(values...); err != nil {
		return err
	}

	return nil
}
