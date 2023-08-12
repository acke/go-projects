package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var stars string
    
	for x := 0; x < numStarsPerLine; x++ {
    	stars = stars + "*"
    }

    return stars + "\n" + welcomeMsg + "\n" + stars 
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    newString := ""
    
     newString = strings.ReplaceAll(oldMsg, "*", "")
    
    newString = strings.TrimSpace(newString)
	
    return newString
}

