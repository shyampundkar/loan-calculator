Loan calculator
-
Calculator Information:  
The Loan Repayments Calculator calculates the repayment amount depending upon the repayment frequency requested, which is dependent upon amount, term and interest rate.  

Based on:
https://www.serviceone.com.au/BankingAndFinance/Loans/Calculators/LoanRepaymentCalculator  

ReadMe in html format is also available at https://learned-shell-332900.ts.r.appspot.com/

Technologies:
-
#Language: Go V1.17  
#CI/CD: Github actions  
#Lint: golangci-lint run  


How to use:
-



HTTP verb: POST  
Google App Engine url: https://learned-shell-332900.ts.r.appspot.com/calculate-loan


Input to POST Body:

```
{
  "loan_amount": 350000,
  "loan_type": "PrincipalAndInterest",
  "payment_frequency": "Monthly",
  "interest_rate": 7.23,
  "loan_term": 1
}
```

-loan types: PrincipalAndInterest / InterestOnly  
-payment frequency: Weekly / Fortnightly / Monthly  


##Output:

```
{
    "monthly_repayments": 30321,
    "total_interest_payable": 13858,
    "amount_owing": [
        {
            "year": 0,
            "principal": 350000,
            "interest": 13858,
            "total": 363858
        },
        {
            "year": 1,
            "principal": 321787,
            "interest": 11749,
            "total": 333536
        },
        {
            "year": 2,
            "principal": 293405,
            "interest": 9810,
            "total": 303215
        },
        {
            "year": 3,
            "principal": 264851,
            "interest": 8043,
            "total": 272893
        },
        {
            "year": 4,
            "principal": 236125,
            "interest": 6447,
            "total": 242572
        },
        {
            "year": 5,
            "principal": 207226,
            "interest": 5024,
            "total": 212250
        },
        {
            "year": 6,
            "principal": 178153,
            "interest": 3776,
            "total": 181929
        },
        {
            "year": 7,
            "principal": 148905,
            "interest": 2702,
            "total": 151607
        },
        {
            "year": 8,
            "principal": 119481,
            "interest": 1805,
            "total": 121286
        },
        {
            "year": 9,
            "principal": 89879,
            "interest": 1085,
            "total": 90964
        },
        {
            "year": 10,
            "principal": 60099,
            "interest": 544,
            "total": 60643
        },
        {
            "year": 11,
            "principal": 30140,
            "interest": 182,
            "total": 30321
        },
        {
            "year": 12,
            "interest": 0
        }
    ]
}
```

