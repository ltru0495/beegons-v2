
package models

import (
	"testing"
)

func init() {
	ConnectToDB("127.0.0.1", 27017, "test")
}

func TestInsertModule(t *testing.T) {
	module := Module{}
	err := module.Insert()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteModule(t *testing.T) {
	module := Module{}
	err := module.Insert()
	if err != nil {
		t.Error(err)
	}
	err = module.Delete()
}


func TestUpdateModule(t *testing.T) {
	module := Module{}
	err := module.Insert()
	if err != nil {
		t.Error(err)
	}
	err = module.Update()
	if err != nil {
		t.Error(err)
	}
}

func TestFindModule(t *testing.T) {
	module:= Module{}
	err :=  module.Insert()
	if err != nil {
		t.Error(err)
	}
	_, err = FindModule(module.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestAllModules(t *testing.T) {
	_, err := AllModules()
	if err != nil {
		t.Error(err)
	}
}

