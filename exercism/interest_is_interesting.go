package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    var interest float32

    switch {
        case balance < 0:
          interest = 3.213
        case balance < 1000:
          interest = 0.5
        case balance < 5000:
          interest = 1.621
        default:
          interest = 2.475
    }

    return float32(interest)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0

    for targetBalance > balance {
        balance = AnnualBalanceUpdate(balance)
        years++
    }

    return years
}

