package budget

var (
	budget []Item
)

type Item struct {
	Name string
	Cost float64
}

func AddExpense(expense Item) {
	budget = append(budget, expense)
}

func ViewBudget() []Item {
	return budget
}

func SetBudget(b []Item) {
	budget = b
}

func ListItems() []Item {
	return budget
}

func CalculateGrandTotal() float64 {
	var sum float64

	for _, x := range budget {
		sum += x.Cost
	}

	return sum
}
