package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcome_msg := "Welcome to my party, " + name + "!"
	return welcome_msg
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	happy_birthday_msg := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return happy_birthday_msg
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	table_instructions := ""
	table_instructions = fmt.Sprintf("Welcome to my party, %s!\n", name)
	table_instructions = table_instructions + fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %0.1f meters from here.\nYou will be sitting next to %s.", table, direction, distance, neighbor)
	return table_instructions
}
