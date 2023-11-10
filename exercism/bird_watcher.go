package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (totalBirds int) {

    for _, birds := range birdsPerDay {
		totalBirds += birds
	}

	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInAWeek := 0

	end := week * 7
	start := end - 7

	sevenDaysOfBirds := birdsPerDay[start:end]

	for i := 0; i < 7; i++ {
		birdsInAWeek += sevenDaysOfBirds[i]
	}

	return birdsInAWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}

