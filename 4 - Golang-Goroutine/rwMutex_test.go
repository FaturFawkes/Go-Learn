package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(ammount int) {
	account.RWMutex.Lock()
	account.Balance += ammount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() (balance int) {
	account.RWMutex.RLock()
	balance = account.Balance
	account.RWMutex.RUnlock()
	return
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}
	
	for i := 0; i < 1000; i++ {
		go func ()  {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
		
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total balance =", account.GetBalance())
}