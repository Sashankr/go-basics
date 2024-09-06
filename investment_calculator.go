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
	// profit_calculator()
	bank()
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


func bank(){

	var account_balance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice);

	wantsCheckBalance := choice == 1
	wantsToDepositMoney := choice == 2
	wantsToWithdrawMoney := choice == 3 

	if wantsCheckBalance {
		fmt.Println("Your balance is: ",account_balance)
	} else if wantsToDepositMoney{
		var deposit_amount float64
		fmt.Print("Enter Deposit Amount: ")
		fmt.Scan(&deposit_amount)
		account_balance +=  deposit_amount // account_balance = account_balance + deposit_amount
		fmt.Printf("New Balance: %.1f\n",account_balance);
	} else if wantsToWithdrawMoney{
		var withdraw_amount float64
		fmt.Print("Enter withdrawal amount: ")
		fmt.Scan(&withdraw_amount)
		
		if withdraw_amount > account_balance{
			fmt.Println("Insufficient Funds, Available Balance: ",account_balance)
			return
		}
		account_balance -= withdraw_amount
		fmt.Printf("Withdraw successfull, Current Balance: %1.f\n",account_balance)

	}

}