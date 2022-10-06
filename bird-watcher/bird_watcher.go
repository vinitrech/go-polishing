package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0

	for _, x := range birdsPerDay {
		sum += x
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	currentWeek := 0

	for index, value := range birdsPerDay {
		if index%7 == 0 {
			currentWeek += 1
		}

		if currentWeek == week {
			sum += value
		}
	}

	return sum

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, value := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] = value + 1
		}
	}

	return birdsPerDay
}
