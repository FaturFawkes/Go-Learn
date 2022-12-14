package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// READ WRITE MUTEX
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


// SIMULASI DEADLOCK
type UserBalance struct {
	Mutex sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user-A ", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user-B ", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{Name: "Fatur", Balance: 1000000}
	user2 := UserBalance{Name: "Ahmad", Balance: 1000000}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000	)
	
	time.Sleep(3 * time.Second)

	fmt.Println("User:", user1.Name, "; Balance:", user1.Balance)
	fmt.Println("User:", user2.Name, "; Balance:", user2.Balance)

}