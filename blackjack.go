package blackjack

import "strconv"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "10":
		return 10
	case "9", "8", "7", "6", "5", "4", "3", "2":
		val, _ := strconv.Atoi(card)
		return val
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Parse the values of the cards
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	// Determine the total value of the player's cards
	totalValue := value1 + value2

	// Determine the recommended action based on the rules
	if totalValue == 21 {
		return "Blackjack!"
	} else if value1 == 11 && value2 == 11 {
		return "Split"
	} else if totalValue >= 17 {
		return "Stay"
	} else if (totalValue == 16 || totalValue == 15) && dealerValue >= 7 {
		return "Hit"
	} else if (totalValue == 14 || totalValue == 13) && dealerValue >= 8 {
		return "Hit"
	} else if totalValue == 12 && (dealerValue == 2 || dealerValue == 3 || dealerValue >= 7) {
		return "Hit"
	} else {
		return "Stay"
	}
}
