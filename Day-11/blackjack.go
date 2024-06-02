package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int = 0
	switch card {
        case "ace" :
        	value = 11
        case "eight" :	
        	value = 8
        case "two" :
        	value = 2
        case "nine" :
        	value = 9
        case "three" :
        	value = 3 
        case "ten" :
        	value = 10
        case "four" :
        	value = 4
        case "jack" :
        	value = 10
        case "five" :
        	value = 5
        case "queen" :
        	value = 10 
        case "six" :
        	value = 6
        case "king" :
        	value = 10 
        case "seven" :
        	value = 7
    	case "other" :
        	value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var option string = ""
    
    value1, value2, dealerCardValue := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
    
    sum := (value1+value2)
    
    switch {
        case (value1==11 && value2==11) :
			option = "P"
        case (value1+value2 ==21 && (dealerCardValue>=10)) :
        	option = "S"
    	case (value1+value2 ==21 && (dealerCardValue<10)) :
        	option = "W"
        case (sum>=17 && sum<=20) :
        	option = "S"
        case (sum>=12 && sum<=16 && dealerCardValue<7) :
        	option = "S"
        case (sum>=12 && sum<=16 && dealerCardValue>=7) :
        	option = "H"
        case (sum<=11) :
        	option = "H"
    }

    return option
}
