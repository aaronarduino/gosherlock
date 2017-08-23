package helpers

var (
	MonthToInt = map[string]int{"jan": 0, "feb": 1, "mar": 2, "apr": 3, "may": 4, "jun": 5, "jul": 6, "aug": 7, "sep": 8, "oct": 9, "nov": 10, "dec": 11}

	// mapping of words to numbers
	WordsToInt = map[string]int{
		"one":     1,
		"first":   1,
		"two":     2,
		"second":  2,
		"three":   3,
		"third":   3,
		"four":    4,
		"fourth":  4,
		"five":    5,
		"fifth":   5,
		"six":     6,
		"sixth":   6,
		"seven":   7,
		"seventh": 7,
		"eight":   8,
		"eighth":  8,
		"nine":    9,
		"ninth":   9,
		"ten":     10,
		"tenth":   10,
	}

	// mapping of number to words
	IntToWords = []string{
		"one|first",
		"two|second",
		"three|third",
		"four|fourth",
		"five|fifth",
		"six|sixth",
		"seven|seventh",
		"eight|eighth",
		"nine|ninth",
		"ten|tenth",
	}
)
