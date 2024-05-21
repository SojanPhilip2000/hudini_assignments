var a int
a = 2

var p *int
p = &a // the variable 'p' contains the memory address of 'a'


var b int
b = *p // b == 2

package electionday

import "fmt"


// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    var a *int
    a = &initialVotes
    return a
	panic("Please implement the NewVoteCounter() function")
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter == nil{
        return 0
    } 
    return *counter
	panic("Please implement the VoteCount() function")
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
    *counter = *counter+increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}
// => John (32)



// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    results[candidate]=results[candidate]-1
}
