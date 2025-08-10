package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"TestGO/internal/database_models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&database_models.User{},
		&database_models.Company{},
		&database_models.TestRun{},
		&database_models.TestSuite{},
		&database_models.TestResult{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}