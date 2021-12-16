# loan-calculator
 
http verb: POST 
url: https://learned-shell-332900.ts.r.appspot.com/calculate-loan

Input to POST Body:

{
  "loan_amount": 350000,
  "loan_type": "PrincipalAndInterest", 
  "payment_frequency": "Fortnightly",
  "interest_rate": 7.23,
  "loan_term": 1  
}


Output:

{
    "monthly_repayments": 13973,
    "total_interest_payable": 13291,
    "amount_owing": [
        {
            "year": 0,
            "principal": 350000,
            "interest": 13292,
            "total": 363292
        },
        {
            "year": 1,
            "principal": 337001,
            "interest": 12318,
            "total": 349318
        },
        {
            "year": 2,
            "principal": 323965,
            "interest": 11381,
            "total": 335346
        },
        {
            "year": 3,
            "principal": 310893,
            "interest": 10480,
            "total": 321373
        },
        {
            "year": 4,
            "principal": 297785,
            "interest": 9615,
            "total": 307400
        },
        {
            "year": 5,
            "principal": 284640,
            "interest": 8787,
            "total": 293427
        },
        {
            "year": 6,
            "principal": 271459,
            "interest": 7996,
            "total": 279455
        },
        {
            "year": 7,
            "principal": 258241,
            "interest": 7241,
            "total": 265482
        },
        {
            "year": 8,
            "principal": 244986,
            "interest": 6523,
            "total": 251509
        },
        {
            "year": 9,
            "principal": 231695,
            "interest": 5842,
            "total": 237537
        },
        {
            "year": 10,
            "principal": 218367,
            "interest": 5197,
            "total": 223564
        },
        {
            "year": 11,
            "principal": 205001,
            "interest": 4590,
            "total": 209591
        },
        {
            "year": 12,
            "principal": 191598,
            "interest": 4020,
            "total": 195618
        },
        {
            "year": 13,
            "principal": 178158,
            "interest": 3487,
            "total": 181646
        },
        {
            "year": 14,
            "principal": 164681,
            "interest": 2992,
            "total": 167673
        },
        {
            "year": 15,
            "principal": 151166,
            "interest": 2534,
            "total": 153700
        },
        {
            "year": 16,
            "principal": 137614,
            "interest": 2113,
            "total": 139727
        },
        {
            "year": 17,
            "principal": 124024,
            "interest": 1731,
            "total": 125755
        },
        {
            "year": 18,
            "principal": 110396,
            "interest": 1386,
            "total": 111782
        },
        {
            "year": 19,
            "principal": 96730,
            "interest": 1079,
            "total": 97809
        },
        {
            "year": 20,
            "principal": 83026,
            "interest": 810,
            "total": 83836
        },
        {
            "year": 21,
            "principal": 69285,
            "interest": 579,
            "total": 69864
        },
        {
            "year": 22,
            "principal": 55505,
            "interest": 386,
            "total": 55891
        },
        {
            "year": 23,
            "principal": 41686,
            "interest": 232,
            "total": 41918
        },
        {
            "year": 24,
            "principal": 27829,
            "interest": 116,
            "total": 27945
        },
        {
            "year": 25,
            "principal": 13934,
            "interest": 39,
            "total": 13973
        },
        {
            "year": 26,
            "interest": -0
        }
    ]
}
