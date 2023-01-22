package budget

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"
)

func InputMonthlyBudget() error {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	var monthlyString string
	fmt.Println("Enter monthly budget (numerical values only):")

	numInput, err := fmt.Scanln(&monthlyString)
	if err != nil {
		fmt.Println("error 1")
		sugar.Errorf("error inputting budget: %s", err)
	}

	if numInput > 1 {
		fmt.Println("error 2")
		sugar.Error("error: too many different values inputted")
	}

	monthlyBudget, err := strconv.ParseFloat(monthlyString, 32)
	if err != nil {
		fmt.Println("error 3")
		return err
	}

	needs := monthlyBudget * 0.5
	savings := monthlyBudget * 0.3
	wants := monthlyBudget * 0.2

	sliceVals := formatCalculatedValues(needs, savings, wants)

	fmt.Printf("%s\n%s\n%s\n", sliceVals[0], sliceVals[1], sliceVals[2])
	return nil
}

func formatCalculatedValues(needs float64, savings float64, wants float64) []string {
	// first step is to convert all floats back into strings
	needString := fmt.Sprintf("budget for needs: $%f", needs)
	savingString := fmt.Sprintf("budget for savings: $%f", savings)
	wantString := fmt.Sprintf("budget for wants: $%f", wants)

	// now add them all to a slice
	budgetSlice := make([]string, 3)
	budgetSlice[0] = needString
	budgetSlice[1] = savingString
	budgetSlice[2] = wantString

	return budgetSlice
}
