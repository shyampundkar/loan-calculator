/*
 * Loan Calculator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CalculateloanBody struct {
	LoanAmount float64 `json:"loan_amount,omitempty"`

	LoanType string `json:"loan_type,omitempty"`

	PaymentFrequency string `json:"payment_frequency,omitempty"`

	InterestRate float64 `json:"interest_rate,omitempty"`

	LoanTerm int32 `json:"loan_term,omitempty"`
}
