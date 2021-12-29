package twelve

import (
	"fmt"
	"strings"
)

func Song() string {
	s := ""
	for i := 1; i <= 11; i++ {
		s += Verse(i) + "\n"
	}
	return s + Verse(12)
}

func Verse(i int) string {
	return fmt.Sprintf("%s: %s", getVersePrefix(i), getVerseGifts(i))
}

func getVersePrefix(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me", getDayNumber(i))
}

func getDayNumber(i int) string {
	switch i {
	case 1:
		return "first"
	case 2:
		return "second"
	case 3:
		return "third"
	case 4:
		return "fourth"
	case 5:
		return "fifth"
	case 6:
		return "sixth"
	case 7:
		return "seventh"
	case 8:
		return "eighth"
	case 9:
		return "ninth"
	case 10:
		return "tenth"
	case 11:
		return "eleventh"
	case 12:
		return "twelfth"
	default:
		return "unknown"
	}
}

func getVerseGifts(i int) string {
	if i == 1 {
		return "a Partridge in a Pear Tree."
	}

	giftsAll := []string{
		"twelve Drummers Drumming",
		"eleven Pipers Piping",
		"ten Lords-a-Leaping",
		"nine Ladies Dancing",
		"eight Maids-a-Milking",
		"seven Swans-a-Swimming",
		"six Geese-a-Laying",
		"five Gold Rings",
		"four Calling Birds",
		"three French Hens",
		"two Turtle Doves",
		"and a Partridge in a Pear Tree.",
	}
	giftsPartial := giftsAll[len(giftsAll)-i:]
	return strings.Join(giftsPartial, ", ")
}
