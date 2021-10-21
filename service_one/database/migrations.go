package database

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

func runMigrations() error {
	var completedMigrations = []string{}
	_, err := DbConnection.Select("name").From("migrations").Load(&completedMigrations)
	if err != nil {
		return err
	}

	var migrationsInDirectory = []string{}
	files, err := ioutil.ReadDir("./database/migrations")
	if err != nil {
		return err
	}

	for _, f := range files {
		migrationsInDirectory = append(migrationsInDirectory, f.Name())
	}
	if err != nil {
		return err
	}

	err = execMigrations(completedMigrations, migrationsInDirectory)
	if err != nil {
		return err
	}

	return nil
}

func execMigrations(completedMigrations []string, migrationsInDirectory []string) error {
	for _, migration := range migrationsInDirectory{
		if !stringInSlice(migration, completedMigrations){
			err := execute(migration)
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func execute(migration string) error {
	tx , err := DbConnection.Begin()
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	path := filepath.Join("database", "migrations", fmt.Sprintf("%s", migration))

	c, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	sql := strings.Split(string(c), ";")
	for i, s := range sql {
		if sql[i] == "" {
			break
		}
		_, err = tx.Exec(s)
		if err != nil {
			return err
		}
	}

	_ , err = tx.InsertInto(MigrationTable).Columns("name", "created_at").Values(migration, time.Now()).Exec()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}