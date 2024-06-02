package partyrobot

import (
    "fmt"
)
const welcomeTemplate string = "Welcome to my party, %v!"
const happyBirthdayTemplate string = "Happy birthday %v! You are now %v years old!"
const assignTableTemplate string = "You have been assigned to table %03d. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v."

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf(welcomeTemplate, name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf(happyBirthdayTemplate, name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    greet := fmt.Sprintf(welcomeTemplate,name)
    greet += "\n"
	return (greet + fmt.Sprintf(assignTableTemplate, table, direction, distance, neighbor))
}
