package models

import (
	// "fmt"
	"fmt"
	"testing"
)

func init(){
	ConnectToDB("127.0.0.1", 27017, "test")
}

func TestInsertTest(t *testing.T) {
	test := Test{Name: "TEEEEEEEEEEEST"}
	err := test.Insert()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTest(t *testing.T) {
	test := Test{Name: "TEEEEEEEEEEEST"}
	err := test.Insert()
	if err != nil {
		t.Error(err)
	}
	err = test.Delete()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTest(t *testing.T) {
	test := Test{Name: "TEEEEEEEEEEEST"}
	err := test.Insert()
	if err != nil {
		t.Error(err)
	}
	test.Name = "Updated Test"
	err = test.Update()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(test)
}
func TestFindTest(t *testing.T) {
	test := Test{Name: "TEEEEEEEEEEEST"}
	err := test.Insert()
	if err != nil {
		t.Error(err)
	}
	found,err := FindTest(test.Id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(found)
}

func TestAllTests(t *testing.T) {
	tests, err := AllTests()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tests)
}
