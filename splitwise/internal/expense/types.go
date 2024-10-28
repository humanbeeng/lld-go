package expense

type SplitType string

const (
	EXACT   SplitType = "exact"
	EQUAL   SplitType = "equal"
	PERCENT SplitType = "percent"
)

type BalanceSheetItem struct {
	CreditorId int
	DebitorId  int
	Amount     float64
}

type Expense struct {
	Id           int
	PaidByUserId int
	Shares       map[int]float64
	Total        float64
	SplitType    SplitType
	Note         string
}

type OwedDetails struct {
	CreditorsId int
	DebitorsId  int
	Amount      float32
}
