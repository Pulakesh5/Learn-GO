package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0
    for _, birds := range birdsPerDay {
        birdCount += birds
    }
    return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInWeek := 0
    for i:=(week-1)*7; i<(week*7); i++ {
        birdsInWeek+=birdsPerDay[i]
    }
    return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, _ := range birdsPerDay {
        if(index%2==0) {
            birdsPerDay[index]+=1
        }
    }
    return birdsPerDay
}
