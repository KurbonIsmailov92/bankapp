package card

import "bank/pkg/bank/types"

const withdrawLimit types.Money = 20_000_00 //лимит снятия суммы с карты
const topupLimit types.Money = 50_000_00    //лимит пополнения карты
const bonusLimit int64 = 5_000_00           // лимит на проценты
//выпус карты
func IssueCard(currency types.Currency, color string, name string) types.Card {

	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card

}

//снятия с баланса
func Withdraw(card types.Card, amount types.Money) types.Card {

	if !card.Active {
		return card
	}

	if amount > card.Balance {
		return card
	}

	if amount < 0 {
		return card
	}

	if amount > withdrawLimit {
		return card
	}

	card.Balance -= amount

	return card
}

//пополнение баланса карты
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount <= 0 {
		return
	}
	if amount > topupLimit {
		return
	}
	card.Balance += amount
}

//начисление бонуса
func AddBonus(card *types.Card, persent int, daysInMonth int, daysInYear int) {

	bonus := (int64(card.MinBalance) * int64(persent) * int64(daysInMonth) / int64(daysInYear)) / 100

	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	if bonus > bonusLimit {
		bonus = 0
	}

	card.Balance += types.Money(bonus)
}

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}

		sum += card.Balance
	}
	return sum
}
