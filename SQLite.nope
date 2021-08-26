package wiz

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"strconv"
)

//	Special: This simple database package is limited: It creates only SQLITE
//		file databases, with any number of tables, which each have an integer
//		primary key column called "key", and a second string column called "value".

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed struct
//			type SimpleDatabase struct

//		Exposed functions:
//			SQLiteOpen(err *error, dbName string) Database

//		Exposed functions on SimpleDatabase struct
//			Close()
//				Closes the db
//			MakeTable(table string) error
//				Creates a table, with columns key(int) and value(string)
//			ClearTable(table string) error
//				Deletes every entry in a given table
//			AddItemAt(table string, key uint64, data string) error
//				Adds an item at a given key in a given table (will not overwrite)
//			GetItemAt(table string, key uint64) (string, error)
//				Gets the item with a given key in a given table
//			DeleteItemAt(table string, key uint64) error
//				Deletes the item with a given key in a given table
//			GetKeys(table string) ([]uint64, error)
//				Gets a list of all the keys in the key column of a given table
//			CheckOrder(table string) (uint64, error)
//				Checks that keys in a table follow pattern: 1,2,3,...
//				Returns nil error only if pattern is unbroken for all keys.
//				On error, returns the first key found to be missing from pattern
//				When error is nil, uint64 is the next available key
//			Latest(table string) (uint64, error)
//				Returns the highest key of all entries in the key column of a table

//		Errors created are usually database shenanigans

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

type SimpleDatabase struct {
	name     string
	database *sql.DB
}

func (db *SimpleDatabase) Close() {
	if db.database != nil {
		db.database.Close()
		db.database = nil
	}
}

func SQLiteOpen(dbName string) (SimpleDatabase, error) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		return SimpleDatabase{}, errors.Wrap(err, "SQLiteOpen")
	}
	return SimpleDatabase{name: dbName, database: db}, nil
}

func (db *SimpleDatabase) MakeTable(table string) error {
	if db.database == nil {
		return errors.New("SimpleDatabase.MakeTable: database is nil")
	}
	query := "CREATE TABLE IF NOT EXISTS " + table + "(key INTEGER PRIMARY KEY, value TEXT);"
	statement, err := db.database.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "SimpleDatabase.MakeTable")
	}
	_, err = statement.Exec()
	return errors.Wrap(err, "SimpleDatabase.MakeTable")
}

func (db *SimpleDatabase) ClearTable(table string) error {
	if db.database == nil {
		return errors.New("SimpleDatabase.ClearTable: database is nil")
	}
	query := "DELETE FROM " + table + ";"
	statement, err := db.database.Prepare(query)
	if err == nil {
		_, err = statement.Exec()
	}
	return errors.Wrap(err, "SimpleDatabase.ClearTable")
}

func (db *SimpleDatabase) AddItemAt(table string, key uint64, data string) error {
	if db.database == nil {
		return errors.New("SimpleDatabase.AddItemAt: database is nil")
	}
	if data == "" {
		return errors.New("SimpleDatabase.AddItemAt: does not accept empty string as data")
	}
	query := "INSERT INTO " + table + " VALUES (" + uToString(key) + ", '" + data + "');"
	statement, err := db.database.Prepare(query)
	if err == nil {
		_, err = statement.Exec()
	}
	return errors.Wrap(err, "SimpleDatabase.AddItemAt")
}

func (db *SimpleDatabase) GetItemAt(table string, key uint64) (string, error) {
	if db.database == nil {
		return "", errors.New("SimpleDatabase.GetItemAt: database is nil")
	}
	query := "SELECT value FROM " + table + " WHERE key=" + uToString(key) + ";"
	row := db.database.QueryRow(query)
	data := ""
	err := row.Scan(&data)
	return data, errors.Wrap(err, "SimpleDatabase.GetItemAt")
}

func (db *SimpleDatabase) DeleteItemAt(table string, key uint64) error {
	if db.database == nil {
		return errors.New("SimpleDatabase.DeleteItemAt: database is nil")
	}
	query := "DELETE FROM " + table + " WHERE key=" + uToString(key) + ";"
	statement, err := db.database.Prepare(query)
	if err == nil {
		_, err = statement.Exec()
	}
	return errors.Wrap(err, "SimpleDatabase.DeleteItemAt")
}

func (db *SimpleDatabase) GetKeys(table string) ([]uint64, error) {
	if db.database == nil {
		return []uint64{}, errors.New("SimpleDatabase.GetKeys: database is nil")
	}
	key := uint64(0)
	keys := []uint64{}
	query := "SELECT key FROM " + table + " ORDER BY key ASC;"
	rows, err := db.database.Query(query)
	defer rows.Close()
	for err == nil && rows.Next() {
		err = rows.Scan(&key)
		keys = append(keys, key)
	}
	if err == nil {
		err = rows.Err()
	}
	return keys, errors.Wrap(err, "SimpleDatabase:GetKeys")
}

func (db *SimpleDatabase) Latest(table string) (uint64, error) {
	if db.database == nil {
		return 0, errors.New("SimpleDatabase.Latest: database is nil")
	}
	latest := uint64(0)
	query := "SELECT MAX(key) FROM " + table + ";"
	row := db.database.QueryRow(query)
	err := row.Scan(&latest)
	return latest, errors.Wrap(err, "SimpleDatabase.Latest")
}

//If your database is expected to have items at primary keys [1,2,3,4...]
// then this function will return the next open key following that order.
// Error is nil only if all (non-zero, non-negative) keys were in order with
// no gaps between them. IGNORES KEY=0. If error is non-nil, the function still
// returns the next open key.
func (db *SimpleDatabase) CheckOrder(table string) (uint64, error) {
	if db.database == nil {
		return 0, errors.New("SimpleDatabase.CheckOrder: database is nil")
	}
	keys, err := db.GetKeys(table)
	if err != nil {
		return 0, errors.Wrap(err, "SimpleDatabase.CheckOrder")
	}
	for len(keys) > 0 && keys[0] < 1 {
		keys = keys[1:] //Ignore key that is less than one
	}
	for i, v := range keys {
		expected := uint64(i + 1)
		if v != expected {
			return expected, errors.New("SimpleDatabase.CheckOrder: gap in expected patern")
		}
	}
	//If keys[n] = n+1, and keys[len(keys)] is open, then next value is len(keys)+1
	return uint64(len(keys) + 1), nil
}

//
//
//
//
//

func uToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}
