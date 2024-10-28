package main

import (
	"fmt"

	"github.com/humanbeeng/lld-go/splitwise/internal/expense"
	"github.com/humanbeeng/lld-go/splitwise/internal/user"
)

func main() {
	fmt.Println("Splitwise")

	us := user.NewUserService()
	u1 := user.User{
		Id:   1,
		Name: "Nithin",
	}
	u2 := user.User{
		Id:   2,
		Name: "Ramu",
	}

	u3 := user.User{
		Id:   3,
		Name: "Somu",
	}

	us.Add(u1)
	us.Add(u2)
	us.Add(u3)

	em := expense.NewExpenseManager(&us)

	e1 := expense.Expense{
		Id:           1,
		PaidByUserId: 1,
		Shares: map[int]float64{
			2: 0,
		},
		SplitType: expense.EQUAL,
		Total:     100,
		Note:      "Chai sutta",
	}
	em.AddExpense(e1)
	em.View(1)

	e2 := expense.Expense{
		Id:           2,
		PaidByUserId: 2,
		Shares: map[int]float64{
			1: 100,
		},
		SplitType: expense.EQUAL,
		Total:     100,
		Note:      "Chai sutta",
	}
	em.AddExpense(e2)

	em.View(1)
	em.View(2)

}
