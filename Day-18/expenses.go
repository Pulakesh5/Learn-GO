package expenses

import (
    "errors"
    "fmt"
)
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(expense []Record, predicate func(Record) bool) []Record {
	filteredRecord := []Record{}
    for _, record := range expense{
        if(predicate(record)) {
            filteredRecord = append(filteredRecord, record)
        }
    }
    return filteredRecord
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
        return record.Day>=p.From && record.Day<=p.To
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(category string) func(Record) bool {
	return func(record Record) bool {
        return record.Category == category
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	totalExpense := 0.0
    filteredExpense := Filter(in, ByDaysPeriod(p))
    for _, record := range filteredExpense {
        totalExpense += record.Amount
    }
    return totalExpense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, category string) (float64, error) {
	totalExpense := 0.0
    
    filteredExpense := Filter(in, ByCategory(category))
    fmt.Println(filteredExpense)
    if(len(filteredExpense)==0) {
        return 0, errors.New(fmt.Sprintf("error(unknown category %v)",category))
    }
    
    filteredExpense = Filter(filteredExpense, ByDaysPeriod(p))
    for _, record := range filteredExpense {
        totalExpense += record.Amount
    }
    return totalExpense, nil
}
