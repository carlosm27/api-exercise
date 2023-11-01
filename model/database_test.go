package model

import (
	"testing"

	"gorm.io/gorm"
)

var testDb *gorm.DB
var err error

func TestInit(t *testing.T) {
	testDb, err = SetupDatabase() // connection of your test database
	if err != nil {
		t.Errorf("Error in initializing test DB: %v", err)
	}
}
