package swagger

import (
	"errors"
	"math"
)

func CalculateLoanRepayments(calculateloanBody CalculateloanBody) (loanRepayment LoanRepayments, err error) {

	calculator, err := CreateCalculator(calculateloanBody.LoanType)

	if err != nil {
		return LoanRepayments{}, err
	}

	totalNumberOfPayments, err := CalculateTotalNumberOfPayments(calculateloanBody)

	if err != nil {
		return LoanRepayments{}, err
	}

	repayment := calculator.CalculateRepayment(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, totalNumberOfPayments)
	loanRepayment.MonthlyRepayments = math.Round(repayment)
	loanRepayment.TotalInterestPayable = math.Round(calculator.CalculateTotalInterestPayable(calculateloanBody.LoanAmount, repayment, totalNumberOfPayments))
	loanRepayment.AmountOwning = calculator.CalculateAmountOwning(calculateloanBody.InterestRate, calculateloanBody.LoanTerm, calculateloanBody.LoanAmount, repayment, totalNumberOfPayments)

	return
}

func CalculateTotalNumberOfPayments(calculateloanBody CalculateloanBody) (totalNumberOfPayments int32, err error) {
	switch calculateloanBody.PaymentFrequency {
	case "Monthly":
		totalNumberOfPayments = 12 * calculateloanBody.LoanTerm
	case "Fortnightly":
		totalNumberOfPayments = 26 * calculateloanBody.LoanTerm
	case "Weekly":
		totalNumberOfPayments = 52 * calculateloanBody.LoanTerm
	default:
		totalNumberOfPayments = 0
		err = errors.New(calculateloanBody.PaymentFrequency + "payment frequency is not supported yet")
	}
	return totalNumberOfPayments, err
}
