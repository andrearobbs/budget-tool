package main

import (
	"errors"
	"fmt"
	"strconv"

	"gitlab.com/parallellearning/lessons/lesson-08/andreas-palace/budget"

	fixerio "github.com/fadion/gofixerio"
	"github.com/manifoldco/promptui"
)

const (
	viewBudgetCmd    = "View Budget"
	addExpenseCmd    = "Add an Expense to Budget"
	currConverterCmd = "Convert Total to Pesos (MXN)"
)

func main() {
	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				viewBudgetCmd,
				addExpenseCmd,
				currConverterCmd,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case viewBudgetCmd:
			ViewBudget()

		case addExpenseCmd:
			err := AddExpense()
			if err != nil {
				fmt.Printf("Prompt failed %n\n", err)
				return
			}

		case currConverterCmd:
			currencyConverter(budget.CalculateGrandTotal)
		}
	}

}

func ViewBudget() {
	budgetItems := budget.ViewBudget()
	for _, x := range budgetItems {
		fmt.Println(x)
	}

	total := budget.CalculateGrandTotal()

	fmt.Println("Your grand total is $", total)
}

func AddExpense() error {
	namePrompt := promptui.Prompt{
		Label: "Name this Expense",
	}
	name, err := namePrompt.Run()
	if err != nil {
		return err
	}

	cost, err := numberPromptHelper("What is the total cost for this expense?")
	if err != nil {
		return err
	}

	newExpense := budget.Item{
		Name: name,
		Cost: cost,
	}

	budget.AddExpense(newExpense)

	fmt.Println("Added new expense to budget!")

	return nil
}

func numberPromptHelper(label string) (float64, error) {

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("You need to type a number, dummy!")
		}
		return nil
	}

	costPrompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	numberStr, err := costPrompt.Run()
	if err != nil {
		return 0, err
	}
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func currencyConverter(func() float64) error {

	total := budget.CalculateGrandTotal()

	exchange := fixerio.New()
	exchange.Symbols(fixerio.USD, fixerio.MXN)

	rates, err := exchange.GetRates()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rates)

	fmt.Println("Your grand total is $", total, "(USD)")

	return nil

}
