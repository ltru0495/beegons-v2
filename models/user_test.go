
package models

import (
	"testing"
)

func init() {
	ConnectToDB()
}

func TestInsertUser(t *testing.T) {
	user := User{}
	err := user.Insert()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	user := User{}
	err := user.Insert()
	if err != nil {
		t.Error(err)
	}
	err = user.Delete()
}


func TestUpdateUser(t *testing.T) {
	user := User{}
	err := user.Insert()
	if err != nil {
		t.Error(err)
	}
	err = user.Update()
	if err != nil {
		t.Error(err)
	}
}

func TestFindUser(t *testing.T) {
	user:= User{}
	err :=  user.Insert()
	if err != nil {
		t.Error(err)
	}
	_, err = FindUser(user.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestAllUsers(t *testing.T) {
	_, err := AllUsers()
	if err != nil {
		t.Error(err)
	}
}

