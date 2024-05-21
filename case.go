package main

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
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