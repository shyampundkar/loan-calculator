package swagger

import "errors"

func CreateCalculator(calculatorType string) (calculator ICalculator, err error) {

	switch calculatorType {
	case "PrincipalAndInterest":
		calculator = InterestPrincipalCalculator{}
	case "InterestOnly":
		calculator = InterestCalculator{}
	default:
		err = errors.New(calculatorType + " Calculator is not supported yet")
	}

	return calculator, err
}
