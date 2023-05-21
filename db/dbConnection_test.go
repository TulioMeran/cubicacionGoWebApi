package db

import (
	"testing"
)

func TestDbConnection(t *testing.T) {

	DbConnection()

	if err != nil {
		t.Error("The DB connection is down!")
	}

}
