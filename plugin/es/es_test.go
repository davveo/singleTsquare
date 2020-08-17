package es

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestEsClient(t *testing.T) {
	esIndex := "test_index"
	user := &User{Name: "ago", Age: 18}
	resp, err := client.Insert(esIndex, user)
	if err != nil {
		t.Error(err)
	}
	bytes, err := client.GetById(esIndex, resp.Id)
	if err != nil {
		t.Error(err)
	}
	result := new(User)
	_ := json.Unmarshal(bytes, result)
	fmt.Println(result)
}
