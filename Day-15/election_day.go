package electionday
import (
    "fmt"
)
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	votes := initialVotes
    return &votes
}
// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if(counter==nil){
        return 0
    }
	return *counter
}
// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	result := ElectionResult{candidateName, votes}
    return &result
}
// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	displayResult := fmt.Sprintf("%v (%v)",result.Name, result.Votes)
    return displayResult
}
// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for key, _ := range results {
        if(key==candidate) {
            results[candidate]-=1
        }
    }
}
