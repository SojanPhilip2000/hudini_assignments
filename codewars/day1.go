import (
	"strings"
)

//1
func SequenceSum(start, end, step int) int {
	sum := 0
	i := start
	for i <= end {
		sum = sum + i
		i = i + step
	}
	return sum
}

//2
func FindShort(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")
	shortest := words[0]
	for i := 0; i < len(words)-1; i++ {
		if len(shortest) > len(words[i+1]) {
			shortest = words[i+1]
		}
	}
	return len(shortest)

}

//3
func Rps(p1, p2 string) string {
	if p1 == "scissors" {
		if p2 == "paper" {
			return "Player 1 won!"
		} else if p2 == "rock" {
			return "Player 2 won!"
		} else {
			return "Draw!"
		}
	} else if p1 == "rock" {
		if p2 == "paper" {
			return "Player 2 won!"
		} else if p2 == "scissors" {
			return "Player 1 won!"
		} else {
			return "Draw!"
		}
	} else if p1 == "paper" {
		if p2 == "rock" {
			return "Player 1 won!"
		} else if p2 == "scissors" {
			return "Player 2 won!"
		} else {
			return "Draw!"
		}
	}
	return "Draw!"
}

//4
func FindOdd(seq []int) int {
	map1 := make(map[int]int)
	for _, num := range seq {
		map1[num] += 1
	}
	for k, v := range map1 {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}

//5

func IsLeapYear(year int) bool {

	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false

	//   if year%4==0{
	//     if year%400==0{
	//         return true
	//         } else if year%100==0{
	//           return false
	//         } else {
	//       return true
	//     }
	//     } else {
	//       return false
	//     }

}