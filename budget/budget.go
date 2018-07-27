package budget

import "database/sql"

var (
	budget []Item
)

type Item struct {
	Name string
	Cost float64
}

type BudgetService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *BudgetService {
	return &BudgetService{
		db: db,
	}
}

const (
	insertNewBudgetQuery = "INSERT INTO budget (budget_name) VALUES (?)"

	selectBudgetQuery = "SELECT id, budget_name FROM budget"

	insertExpenseQuery = "INSERT INTO expense (expense_name, expense_cost) VALUES (?,?)"
)

func (a *BudgetService) AddExpense(expense Item) {
	budget = append(budget, expense)
}

func (a *BudgetService) ViewBudget() []Item {
	return budget
}

func (a *BudgetService) SetBudget(b []Item) {
	budget = b
}

func (a *BudgetService) ListItems() []Item {
	return budget
}

func (a *BudgetService) CalculateGrandTotal() float64 {
	var sum float64

	for _, x := range budget {
		sum += x.Cost
	}

	return sum
}
