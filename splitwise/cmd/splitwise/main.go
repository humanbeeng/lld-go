package main

import (
	"fmt"

	"github.com/humanbeeng/lld-go/splitwise/internal/expense"
	"github.com/humanbeeng/lld-go/splitwise/internal/user"
)

func main() {

	fmt.Println("Splitwise")
	userService := user.NewUserService()
	u1 := user.User{
		Id:       1,
		Username: "Nithin",
	}

	u2 := user.User{
		Id:       2,
		Username: "Ramu",
	}

	err := userService.Add(u1)
	if err != nil {
		fmt.Println(err)
	}

	userService.Add(u2)

	expenseManager := expense.NewExpenseManager(&userService)

	expense1 := expense.Expense{
		Id:           1,
		PaidByUserId: 1,
		Total:        100,
		Contributions: map[int]expense.Contribution{
			2: {
				UserId: 2,
				Value:  0,
			},
		},
		ExpenseType: expense.EXACT,
		Note:        "Chai sutta",
	}

	expense2 := expense.Expense{
		Id:           2,
		PaidByUserId: 2,
		Total:        100,
		Contributions: map[int]expense.Contribution{
			1: {
				UserId: 1,
				Value:  50,
			},
		},
		ExpenseType: expense.EXACT,
		Note:        "Idli vada",
	}

	err = expenseManager.SubmitExpense(expense1)
	if err != nil {
		fmt.Println(err)
	}

	err = expenseManager.SubmitExpense(expense2)
	if err != nil {
		fmt.Println(err)
	}

	err = expenseManager.Display(1)
	if err != nil {
		fmt.Println(err)
	}

	err = expenseManager.Display(2)
	if err != nil {
		fmt.Println(err)
	}

}
