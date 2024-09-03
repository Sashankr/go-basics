package main

import (
	"fmt"
	"math"
)

func profit_calculator(){
	// inputs: revenue, expenses, tax_rate
	// outputs : EBT, Earning after tax(profit)
	// outputs : ratio(EBT/profit)
	// Final outputs : EBT, profit and ratio

	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Enter your revenue: ")
	// fmt.Scan(&revenue)

	revenue = get_user_input("Enter your revenue: ")

	// fmt.Print("Enter your expenses: ")
	// fmt.Scan(&expenses)
	expenses = get_user_input("Enter your expenses: ")

	// fmt.Print("Enter Tax Rate: ")
	// fmt.Scan(&taxRate)

	taxRate = get_user_input("Enter Tax Rate: ")


	// earningsBeforeTax := revenue - expenses;
	// fmt.Println("Earnings Before Tax: ",earningsBeforeTax);

	// earningsAfterTax := earningsBeforeTax - ((taxRate/100)*earningsBeforeTax)
	// fmt.Println("Earnings After Tax: ", earningsAfterTax)

	// ratio := earningsBeforeTax / earningsAfterTax
	// fmt.Println("Ratio: ", ratio);

	ebt, profit, ratio := calculateEarnings(revenue,expenses,taxRate)
	fmt.Printf("Eearning before tax: %.1f\nProfit: %1.f\nRatio: %.3f\n",ebt,profit,ratio)

}

func get_user_input(promptText string)(float64)  {
	var tempVariable float64
	fmt.Print(promptText)
	fmt.Scan(&tempVariable)
	return tempVariable
}

func calculateEarnings(revenue, expenses float64,taxRate float64)(float64, float64, float64){
	ebt := revenue - expenses
	profit :=ebt - ((taxRate/100)*ebt)
	ratio := ebt / profit	
	return ebt, profit, ratio
}

const inflationRate = 2.5

func investment_calculator(){
	
	var investmentAmount float64 
	var expectedReturnRate float64
	var years float64

	// fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount);

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate);

	fmt.Print("Enter investment duration in years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow((1 + (expectedReturnRate /  100)),years); 
	// futureRealValue := futureValue / math.Pow(1 + (inflationRate / 100),years)
	futureValue, futureRealValue:= calculateFutureValues(investmentAmount,expectedReturnRate,years)

	// fmt.Println(futureValue)
	fmt.Printf("Future Value: %.1f\nFuture Value Adjusted for inflation:%.1f",futureValue,futureRealValue);

	// formattedFutreValue := fmt.Sprintf("Future Value: %f.1\n",futureValue)
	// formattedFutureAdjustedInflationValue := fmt.Sprintf("Future Adjusted for inflation: %f.1\n",futureRealValue)

	// fmt.Print(formattedFutreValue)
	// fmt.Print((formattedFutureAdjustedInflationValue))

	// multi line
	// fmt.Printf(`Future Value: %.1f
	// Future Value Adjusted for inflation:%.1f`,futureValue,futureRealValue);



	// fmt.Println(futureRealValue)
}

func main(){
	// investment_calculator()
	profit_calculator()
} 

func outputText(value string){
	fmt.Println(value);
}

func calculateFutureValues(investmentAmount ,expectedReturnRate ,years float64) (float64,float64) {
	fv := investmentAmount * math.Pow((1 + (expectedReturnRate /  100)),years)
	rfv := fv / math.Pow(1 + (inflationRate / 100),years)
	return  fv, rfv 	
}

// alternative return syntax
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64 )(fv float64,rfv float64){
// 	fv = investmentAmount * math.Pow((1 + (expectedReturnRate /  100)),years)
// 	rfv = fv / math.Pow(1 + (inflationRate / 100),years)
// 	return
// }
