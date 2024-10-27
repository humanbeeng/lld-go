package expense

type ExpenseType string

const (
	EXACT   ExpenseType = "exact"
	EQUAL   ExpenseType = "equal"
	PERCENT ExpenseType = "percent"
)

type Contribution struct {
	UserId int
	Value  float32
}

type BalanceSheetItem struct {
	DebtorId   int
	CreditorId int
	Amount     float32
}

type Expense struct {
	Id            int
	PaidByUserId  int
	Total         float32
	Contributions map[int]Contribution
	ExpenseType   ExpenseType
	Note          string
}
