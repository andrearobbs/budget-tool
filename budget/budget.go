package budget

import "database/sql"

var (
	budget []Expense
)

type Expense struct {
	Id       int
	Name     string
	Cost     float64
	BudgetId int
}

type Budget struct {
	Id   int
	Name string
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
	insertNewBudgetQuery = "INSERT INTO budget (budget_name) VALUES (?); SELECT LAST_INSERT_ID();"

	selectBudgetQuery = "SELECT id, budget_name FROM budget WHERE budget_name = ?"

	insertExpenseQuery = "INSERT INTO expense (expense_name, expense_cost) VALUES (?,?)"

	selectExpensesQuery = "SELECT id, expense_name, expense_cost, budget_id FROM expense"
)

func (a *BudgetService) FindOrCreateBudget(budgetName string) (Budget, error) {

	// a.db.QueryRow budget to see if it exists, if it does return it
	row := a.db.QueryRow(selectBudgetQuery, budgetName)

	var budget Budget

	err := row.Scan(
		&budget.Id,
		&budget.Name,
	)
	if err == nil {
		return budget, nil
	}

	row = a.db.QueryRow(insertNewBudgetQuery, budgetName)

	err = row.Scan(
		&budget.Id,
	)
	if err != nil {
		return Budget{}, err
	}

	budget.Name = budgetName

	return budget, nil
}

func (a *BudgetService) AddExpense(expense Expense) {
	budget = append(budget, expense)
}

func (a *BudgetService) ListExpenses(budgetId int) ([]Expense, error) {
	rows, err := a.db.Query(selectExpensesQuery)
	if err != nil {
		return nil, err
	}

	var expenses []Expense
	for rows.Next() {
		var expense Expense

		err := rows.Scan(
			&expense.Id,
			&expense.Name,
			&expense.Cost,
			&expense.BudgetId,
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}
	return budget, nil
}

func (a *BudgetService) CalculateGrandTotal() float64 {
	var sum float64

	for _, x := range budget {
		sum += x.Cost
	}

	return sum
}
