package expense

import (
	"fmt"
	"math"

	"github.com/humanbeeng/lld-go/splitwise/internal/user"
)

type ExpenseManager struct {
	userService *user.UserService
	expenseMap  map[int]map[int]float64
}

func NewExpenseManager(userService *user.UserService) ExpenseManager {
	return ExpenseManager{
		userService: userService,
		expenseMap:  make(map[int]map[int]float64, 0),
	}
}

func (em *ExpenseManager) AddExpense(expense Expense) error {

	// TODO: add validation

	_, err := em.userService.Get(expense.PaidByUserId)
	if err != nil {
		return err
	}

	splitDetails, err := em.calculateSplit(expense.Shares, expense.SplitType, expense.Total)

	if err != nil {
		return err
	}

	balanceMap, ok := em.expenseMap[expense.PaidByUserId]
	if !ok {

		for id, val := range splitDetails {
			balanceMap = make(map[int]float64)
			balanceMap[id] = val
			em.expenseMap[id] = map[int]float64{expense.PaidByUserId: -val}
		}

		em.expenseMap[expense.PaidByUserId] = balanceMap

	} else {

		for id, val := range splitDetails {
			balanceMap[id] += val

			if balanceMap[id] == 0 {
				delete(balanceMap, id)
			}

			m := em.expenseMap[id]
			m[expense.PaidByUserId] -= val

			if m[expense.PaidByUserId] == 0 {
				delete(m, expense.PaidByUserId)
			}

			em.expenseMap[id] = m
		}

		em.expenseMap[expense.PaidByUserId] = balanceMap
	}

	return nil
}

func (em *ExpenseManager) calculateSplit(shares map[int]float64, splitType SplitType, total float64) (map[int]float64, error) {

	splitDetails := make(map[int]float64, 0)

	for id, val := range shares {
		switch splitType {
		case EQUAL:
			{
				splitDetails[id] = math.Round((total * 100) / float64(len(shares)+1) / 100)
			}

		case EXACT:
			{
				splitDetails[id] = val
			}

		case PERCENT:
			{
				splitDetails[id] = (total * val) / 100
			}
		default:
			{
				return nil, fmt.Errorf("invalid split type %v", splitType)
			}
		}
	}

	return splitDetails, nil
}

func (em *ExpenseManager) View(userId int) error {
	expenseMap, ok := em.expenseMap[userId]

	if !ok {
		return fmt.Errorf("no balance")
	}
	creditor, err := em.userService.Get(userId)
	if err != nil {
		return err
	}

	fmt.Printf("--------------------------\n")
	if len(expenseMap) == 0 {
		fmt.Println("No balances")
		return nil
	}

	for id, val := range expenseMap {
		debitor, err := em.userService.Get(id)
		if err != nil {
			return err
		}

		fmt.Printf("%v owes %v %v\n", debitor.Name, creditor.Name, val)
	}
	fmt.Printf("--------------------------\n\n")

	return nil
}
