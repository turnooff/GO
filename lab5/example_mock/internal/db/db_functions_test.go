package dbfunctions_test

import (
	"database/sql"
	"errors"
	"reflect"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	db_functions "example_mock/internal/db"
)

func TestNew(t *testing.T) {
	mockDB := new(db_functions.Database)
	result := db_functions.New(*mockDB)
	dbService := db_functions.Service{DB: *mockDB}
	if result != dbService {
		t.Errorf("expected to be %v, got %v", dbService, result)
	}
}

var testTableGetName = []struct {
	nameTest    string
	names       []string
	errExpected error
}{
	{
		nameTest:    "Success Test",
		names:       []string{"Tim", "Alex"},
		errExpected: nil,
	},
	{
		nameTest:    "Error Test",
		names:       []string{},
		errExpected: errors.New("the slice is empty"),
	},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	dbService := db_functions.Service{DB: mockDB}

	for _, row := range testTableGetName {

		rows := sqlmock.NewRows([]string{"name"})

		for _, name := range row.names {
			rows.AddRow(name)
		}

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows).WillReturnError(row.errExpected)

		names, err := dbService.GetNames()
		if err != nil {
			t.Fatal(err, "in", row.nameTest)
		}

		if row.errExpected != nil {
			t.Fatal(err, "in", row.nameTest)
		}

		if !reflect.DeepEqual(names, row.names) {
			t.Fatalf("expected names to be %v, got %v in %v", row.names, names, row.nameTest)
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

var testTableSelectUniqueValues = []struct {
	nameTest   string
	columnName string
	tableName  string
	want       []string
	wantErr    bool
	err        error
	rows       *sqlmock.Rows
}{
	{
		nameTest:   "Success Test",
		columnName: "column",
		tableName:  "table",
		want:       []string{"Val1", "Val2", "Val3"},
		wantErr:    false,
		err:        nil,
		rows: sqlmock.NewRows([]string{"value"}).
			AddRow("Val1").
			AddRow("Val2").
			AddRow("Val3"),
	},
	{
		nameTest:   "Erorr Test",
		columnName: "column",
		tableName:  "table",
		want:       nil,
		wantErr:    true,
		err:        sql.ErrNoRows,
		rows:       sqlmock.NewRows([]string{"value"}),
	},
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	dbService := db_functions.New(mockDB)

	for _, row := range testTableSelectUniqueValues {
		rows := sqlmock.NewRows([]string{"name"})

		for _, name := range row.nameTest {
			rows.AddRow(name)
		}
		query := "SELECT DISTINCT " + row.columnName + " FROM " + row.tableName
		if row.wantErr {
			mock.ExpectQuery(query).WillReturnError(row.err)
		} else {
			mock.ExpectQuery(query).WillReturnRows(row.rows)
		}
		got, err := dbService.SelectUniqueValues(row.columnName, row.tableName)
		if err != nil {
			t.Fatal(err, "in", row.nameTest)
		}
		if !reflect.DeepEqual(got, row.want) {
			t.Fatalf("expected names to be %v, got %v in %v", row.want, got, row.nameTest)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
