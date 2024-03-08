package main

import (
	"flag"
	"fmt"
)
import "math"

const diff = "diff"
const annuity = "annuity"

func main() {
	var _type string
	var payment float64
	var principal float64
	var periods float64
	var interest float64

	var typePointer *string = &_type
	var paymentPointer *float64 = &payment
	var principalPointer *float64 = &principal
	var periodsPointer *float64 = &periods
	var interestPointer *float64 = &interest

	ParseArgs(typePointer, paymentPointer, principalPointer, periodsPointer, interestPointer)

	if !(_type == annuity || _type == diff) || interest == float64(0) {
		fmt.Println("Incorrect parameters")
		return
	}

	switch _type {
	case annuity:
		if periods < float64(0) || principal < float64(0) || payment < float64(0) {
			fmt.Println("Incorrect parameters")
			return
		}
		if periods == float64(0) && principal == float64(0) && payment == float64(0) {
			fmt.Println("Incorrect parameters")
			return
		}
		if periods == float64(0) {
			result := CalculateNumberOfPayments(*paymentPointer, CalculateNominalInterest(*interestPointer), *principalPointer)
			years := result / 12
			months := result % 12
			fmt.Print("It will take ")
			if years != 0 {
				fmt.Printf("%d years\n", years)
			}
			if months != 0 {
				fmt.Printf("%d months\n", months)
			}
			fmt.Println("to repay this loan!")
			fmt.Printf("Overpayment = %d\n", int(math.Abs(math.Ceil((float64(result)*payment)-principal))))
		} else if principal == float64(0) {
			result := CalculateLoanPrinciple(*paymentPointer, CalculateNominalInterest(*interestPointer), *periodsPointer)
			fmt.Printf("Your loan principal = %d!\n", int(math.Floor(result)))
			fmt.Printf("Overpayment = %d\n", int(math.Abs(math.Ceil((payment*periods)-result))))
			//fmt.Printf("Your loan principal = %d!", int(math.Round(result)))
		} else if payment == float64(0) {
			result := CalculateMonthlyPayment(*principalPointer, *periodsPointer, CalculateNominalInterest(*interestPointer))
			fmt.Printf("Your annuity payment = %d!\n", int(math.Ceil(result)))
			fmt.Printf("Overpayment = %d\n", int(math.Abs(math.Ceil(principal-float64(int(math.Ceil(result)))*periods))))
			//fmt.Printf("Your monthly payment = %d!", int(math.Ceil(result)))
		}
	case diff:
		if periods < float64(0) || principal < float64(0) || payment < float64(0) || interest < float64(0) {
			fmt.Println("Incorrect parameters")
			return
		}
		if periods == float64(0) || principal == float64(0) || payment != float64(0) {
			fmt.Println("Incorrect parameters")
			return
		}
		calculateTotalDiff(principal, periods, CalculateNominalInterest(*interestPointer), periods)
	}

}
func calculateTotalDiff(P float64, n float64, i float64, m float64) {
	var sum float64 = 0

	for index := 1; index <= int(m); index++ {
		result := calculateDiff(P, n, i, float64(index))
		sum += float64(int(math.Ceil(result)))
		fmt.Printf("Month %d: payment is %d\n", index, int(math.Ceil(result)))
	}
	fmt.Println()
	fmt.Printf("Overpayment = %d\n", int(math.Abs(math.Ceil(P-sum))))
}

func ParseArgs(
	typePointer *string,
	paymentPointer *float64,
	principalPointer *float64,
	periodsPointer *float64,
	interestPointer *float64) {

	flag.StringVar(typePointer, "type", "", "Enter type")
	flag.Float64Var(paymentPointer, "payment", 0, "Enter payment")
	flag.Float64Var(principalPointer, "principal", 0, "Enter principal")
	flag.Float64Var(periodsPointer, "periods", 0, "Enter periods")
	flag.Float64Var(interestPointer, "interest", 0, "Enter interest")

	flag.Parse()
}

func CalculateLoanPrinciple(A float64, i float64, n float64) float64 {
	P := A / ((i * math.Pow(1+i, n)) / (math.Pow(1+i, n) - 1))
	return P
}

func CalculateNumberOfPayments(P float64, i float64, A float64) int {
	n := math.Log(P/(P-i*A)) / math.Log(1+i)
	return int(math.Ceil(n))
}

func CalculateMonthlyPayment(P float64, n float64, i float64) float64 {
	A := P * ((i * math.Pow(1+i, n)) / (math.Pow(1+i, n) - 1))
	return A
}

func CalculateNominalInterest(interest float64) float64 {
	return interest / (float64(12) * float64(100))
}

func calculateDiff(P float64, n float64, i float64, m float64) float64 {
	result := (P / n) + (i * (P - (P*(m-1))/n))
	return result
}
