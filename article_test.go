package main

import "testing"
import "ModRestApi/app/model"

func TestFindArticle(t *testing.T) {
	noData := model.FindArticle("10")
	if noData.Id == "0" {
		t.Errorf("Failed, no data find with this Id %v", noData.Id)
	} else {
		t.Logf("Success, data find with this Id %v", noData.Id)
	}

	data := model.FindArticle("1")
	if data.Id != "1" {
		t.Errorf("Failed, no data find with this Id %v", data.Id)
	} else {
		t.Logf("Success, data find with this Id %v", data.Id)
	}
}
