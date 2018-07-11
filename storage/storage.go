package storage

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/parallellearning/lessons/lesson-08/andreas-palace/budget"
)

const filename = "budget.json"

func Load() error {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var savedItems []budget.Item
	err = json.Unmarshal(fileContents, &savedItems)
	if err != nil {
		return err
	}

	budget.SetBudget(savedItems)

	return nil
}

func Save() error {
	itemList := budget.ListItems()

	itemListBytes, err := json.MarshalIndent(itemList, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, itemListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}
