package thefarm
import (
    "errors"
    "fmt"
)
// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, numOfCows int) (foodPerCow float64, err error) {
	totalFood, err1 := fodderCalculator.FodderAmount(numOfCows)
    if( err1 != nil) {
		return 0, err1
    }
    fatteningFactor, err2 := fodderCalculator.FatteningFactor()
    if( err2 != nil) {
		return 0, err2
    }
    return (totalFood*fatteningFactor)/float64(numOfCows), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numOfCows int) (foodPerCow float64, err error) {
    if (numOfCows <=0 ) {
        return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fodderCalculator, numOfCows)
}
// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    numOfCows int
    errorMessage string
}
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %s", e.numOfCows, e.errorMessage)
}
func ValidateNumberOfCows(numOfCows int) *InvalidCowsError {
    if(numOfCows>0) {
        return nil
    }
    
    cowsError := &InvalidCowsError{
        numOfCows,
        "",
    }
    if(numOfCows<0) {
        cowsError.errorMessage = "there are no negative cows"
    } else if(numOfCows == 0) {
        cowsError.errorMessage = "no cows don't need food"
    }
	return cowsError
}
