package swagger

type ICalculator interface {

	// Calculate how much needs to be paid back for each month/fortnightly/weekly
	CalculateRepayment(interestRate float64, loanTerm int32, loanAmount float64, totalNumberOfPayments int32) (repayment float64)

	// Calculate total interest payable for the whole loan term (e.g 1,2,3year...etc)
	CalculateTotalInterestPayable(loanAmount float64, repayment float64, totalNumberOfPayments int32) float64

	// Calculate the reducing interest and principal for the loan term
	CalculateAmountOwning(interestRate float64, loanTerm int32, loanAmount float64, monthlyRepayment float64, totalNumberOfPayments int32) (loanRepaymentsAmountOwing []LoanRepaymentsAmountOwing)
}
