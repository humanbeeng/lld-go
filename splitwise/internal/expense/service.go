package expense

import (
	"fmt"

	"github.com/humanbeeng/lld-go/splitwise/internal/user"
)

type ExpenseManager struct {
	ExpenseLog   []Expense
	BalanceSheet []BalanceSheetItem
	userService  *user.UserService
}

func NewExpenseManager(userService *user.UserService) ExpenseManager {
	return ExpenseManager{
		ExpenseLog:   make([]Expense, 0),
		BalanceSheet: make([]BalanceSheetItem, 0),
		userService:  userService,
	}
}

func (em *ExpenseManager) SubmitExpense(expense Expense) error {
	// validate
	err := em.validateExpense(expense)
	if err != nil {
		return err
	}

	em.ExpenseLog = append(em.ExpenseLog, expense)
	numOfContributors := len(expense.Contributions)

	switch expense.ExpenseType {
	case EQUAL:
		{
			split := expense.Total / float32(numOfContributors)
			for _, cont := range expense.Contributions {
				if cont.UserId != expense.PaidByUserId {
					item := BalanceSheetItem{
						DebtorId:   cont.UserId,
						CreditorId: expense.PaidByUserId,
						Amount:     expense.Total - split,
					}
					em.BalanceSheet = append(em.BalanceSheet, item)
				}
			}
		}

	case PERCENT:
		{
			for _, cont := range expense.Contributions {
				if cont.UserId != expense.PaidByUserId {
					item := BalanceSheetItem{
						DebtorId:   cont.UserId,
						CreditorId: expense.PaidByUserId,
						Amount:     expense.Total - (expense.Total*cont.Value)/100,
					}

					em.BalanceSheet = append(em.BalanceSheet, item)
				}
			}
		}

	case EXACT:
		{
			for _, cont := range expense.Contributions {
				if cont.UserId != expense.PaidByUserId {
					item := BalanceSheetItem{
						DebtorId:   cont.UserId,
						CreditorId: expense.PaidByUserId,
						Amount:     expense.Total - cont.Value,
					}
					em.BalanceSheet = append(em.BalanceSheet, item)
				}

			}
		}
	}

	return nil
}

func (em *ExpenseManager) Display(userId int) error {
	if _, err := em.userService.Get(userId); err != nil {
		return err
	}

	if len(em.BalanceSheet) == 0 {
		return fmt.Errorf("no balances")
	}

	for _, item := range em.BalanceSheet {
		if item.CreditorId == userId {
			debtor, _ := em.userService.Get(item.DebtorId)
			creditor, _ := em.userService.Get(item.CreditorId)
			fmt.Printf("%v owes %v to %v\n", debtor.Username, item.Amount, creditor.Username)
		}
	}

	return nil
}

func (em *ExpenseManager) validateExpense(expense Expense) error {

	if len(expense.Contributions) == 0 {
		return fmt.Errorf("no contributors found")
	}

	var total float32
	for _, cont := range expense.Contributions {

		// check if user exists
		if _, err := em.userService.Get(cont.UserId); err != nil {
			return err
		}

		if cont.Value < 0 {
			return fmt.Errorf("invalid contribution value: %v", cont.Value)
		}

		total += cont.Value
	}

	switch expense.ExpenseType {
	case EXACT:
		{
			if total != expense.Total {
				return fmt.Errorf("contributions amount mismatch. calculated total: %v. submitted total: %v", total, expense.Total)
			}
		}
	case PERCENT:
		{
			if total != float32(100) {
				return fmt.Errorf("contributions amount mismatch. expected 100%%. got %v", total)
			}
		}
	}

	return nil
}
