// Package bank provides a concurrency-safe bank with one account.
package main

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawData)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

type withdrawData struct {
	amount   int
	resultCh chan bool
}

func Withdraw(amount int) bool {
	resultCh := make(chan bool)
	withdraws <- withdrawData{amount, resultCh}
	return <-resultCh
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawData := <-withdraws:
			if balance > withdrawData.amount {
				balance -= withdrawData.amount
				withdrawData.resultCh <- true
			} else {
				withdrawData.resultCh <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
