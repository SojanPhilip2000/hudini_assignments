package main

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard1(card string) int {
    switch card{
    case "ace":
    	return 11
    case "ten", "jack" ,"queen", "king":
    	return 10
    case "nine":
    	return 9
    case "eight":
    	return 8
    case "seven":
    	return 7
    case "six":
    	return 6
    case "five":
    	return 5
    case "four":
    	return 4
    case "three":
    	return 3
    case "two":
    	return 2
    default:
    	return 0
    }
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var sum = ParseCard(card1)+ParseCard(card2)
    if (card1 == "ace" && card2 =="ace"){
        return "P"
    } else if(sum==21){
        if (ParseCard(dealerCard)==11 || ParseCard(dealerCard)==10){
            return "S"
        } else{
            return "W"
        }
    } else if (sum>=17 && sum<=20){
        return "S"
    } else if (sum>=12 && sum<=16){
        if  ParseCard(dealerCard)>=7{
            return "H"
        } else{
            return "S"
        } 
    } else if sum<= 11{
        return "H"
    } else {
        return ""
    } 
	panic("Please implement the FirstTurn function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn1(card1, card2, dealerCard string) string {
	//panic("Please implement the FirstTurn function")
    playerScore := ParseCard(card1)+ParseCard(card2);
    switch {
       	case card1 == "ace" && card2 =="ace": return "P";
        case playerScore == 21 && dealerCard != "ace" && ParseCard(dealerCard) != 10: return "W";
        case playerScore == 21 && (dealerCard == "ace" || ParseCard(dealerCard) == 10): return "S";
        case playerScore >= 17 && playerScore <= 20: return "S";
        case playerScore >= 12 && playerScore <= 16 && ParseCard(dealerCard) < 7: return "S";
		case playerScore >= 12 && playerScore <= 16 && ParseCard(dealerCard) >= 7: return "H";	
        case playerScore <= 11: return "H";
    }
    return "";
}
