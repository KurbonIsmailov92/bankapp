package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_positive() {
	myCard := &types.Card{
		Balance: 10_000_00,
		Active:  true,
	}

	Deposit(myCard, 5_000_00)
	fmt.Println(myCard.Balance)
	// Output: 1500000
}

func ExampleDeposit_inactive() {
	myCard := &types.Card{
		Balance: 10_000_00,
		Active:  false,
	}
	Deposit(myCard, 5_000_00)
	fmt.Println(myCard.Balance)
	// Output: 1000000
}

//больше лимита
func ExampleDeposit_outOfLimit() {

	myCard := &types.Card{
		Balance: 10_000_00,
		Active:  true,
	}
	Deposit(myCard, 50_000_01)
	fmt.Println(myCard.Balance)
	// Output: 1000000
}

// Тесты для бонуса

func ExampleAddBonus_positive() {
	myCard := &types.Card{
		Balance:    100_00,
		Active:     true,
		MinBalance: 50_00,
	}
	AddBonus(myCard, 3, 30, 365)
	fmt.Println(myCard.Balance)

	// Output: 10012
}

func ExampleAddBonus_inactive() {
	myCard := &types.Card{
		Balance:    100_00,
		Active:     false,
		MinBalance: 50_00,
	}
	AddBonus(myCard, 3, 30, 365)
	fmt.Println(myCard.Balance)

	// Output: 10000
}

func ExampleAddBonus_negativeBalance() {
	myCard := &types.Card{
		Balance:    0,
		Active:     true,
		MinBalance: 50_00,
	}
	AddBonus(myCard, 3, 30, 365)
	fmt.Println(myCard.Balance)

	// Output: 0
}

func ExampleAddBonus_limitAbove() {
	myCard := &types.Card{
		Balance:    3_000_000_00,
		Active:     true,
		MinBalance: 3_000_000_00,
	}
	AddBonus(myCard, 3, 30, 365)
	fmt.Println(myCard.Balance)

	// Output: 300000000
}

func ExampleTotal() {

	fmt.Println(Total([]types.Card{
		{Balance: 1_000_00, Active: true},
	}))
	
	fmt.Println(Total([]types.Card{
		{Balance: 1_000_00, Active: true},
		{Balance: 2_000_00, Active: true},
	}))
	
	fmt.Println(Total([]types.Card{
		{Balance: 1_000_00, Active: false},
	}))
	
	fmt.Println(Total([]types.Card{
		{Balance: -1_000_00, Active: true},
	}))			

	// Output: 
	// 100000
	// 300000
	// 0
	// 0
}
