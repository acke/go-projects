package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total_birds := 0

	for i := 0; i < len(birdsPerDay); i++ {
		total_birds = total_birds + birdsPerDay[i]
	}

	return total_birds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birds_in_a_week := 0

	end := week * 7
	start := end - 7

	seven_days_of_birds := birdsPerDay[start:end]

	for i := 0; i < 7; i++ {
		birds_in_a_week = birds_in_a_week + seven_days_of_birds[i]
	}

	return birds_in_a_week
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
