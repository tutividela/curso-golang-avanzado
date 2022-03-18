package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database
var lock sync.Mutex

func  getDatabaseInstance() *Database  {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB connection")
		db = &Database{}
		db.createSingleConnection()
		return db
	}else{
		fmt.Println("DB connection already created")
	}
	return db
}
func (Database) createSingleConnection(){
	fmt.Printf("Creating singleton for database")
	time.Sleep(2*time.Second)
}
func main(){
	var wg sync.WaitGroup
	for i:= 0 ; i<10;i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}