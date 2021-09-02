package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

/*
nama : Wahyu Dwi Ramadhan
kode : GLNG-KS01-008
*/

type Friend struct {
	Id int
	Name string
	Address string
	Job string
	Reason string
}

func (friend Friend) PrintData() {
	fmt.Println("nama :",friend.Name)
	fmt.Println("alamat :",friend.Address)
	fmt.Println("pekerjaan :",friend.Job)
	fmt.Println("alasan memilih golang :",friend.Reason)
}

func searchData(key int, friends []Friend) (Friend, error) {
	for _,friend := range friends {
		if friend.Id == key {
			return friend, nil
		}
	}
	return Friend{}, errors.New("Data not found")
}

func main(){
	friends := []Friend{ 
		{1, "MUHAMMAD ZUNAN ALFIKRI", "Jakarta", "Student", "Go has high demand"},
		{2, "ARFAN JADULHAQ", "Bandung", "Student", "Go is easy to learn"},
		{3, "TRIYONO", "Surakarta", "Student", "There are many companies use go"},
		{4, "ADITYA RIZKI PRATAMA", "Semarang", "Student", "Go has high demand"},
		{5, "YULYANO THOMAS DJAYA", "Jakarta", "Student", "Go is easy to learn"},
		{6, "ARFINAL", "ACEH", "Student", "Go is simple"},
		{7, "FELIX YANGSEN", "Makasar", "Student", "There are many companies use go"},
		{8, "WAHYU DWI RAMADHAN", "Jepara", "Student", "Go is powerfull"},
		{9, "MUHAMMAD HANIF NAUFAL EKA W", "Bali", "Student", "I love Golang"},
		{10, "THOBIB KHOIRUL ANNAS", "Serang", "Student", "Go is east to learn"},
	}

	key := os.Args
	if len(key) < 2 {
		fmt.Println("Insert key first")
		return
	}

	intKey,err := strconv.Atoi(key[1])
	if err != nil {
		fmt.Println("Key must be number")
		return
	}

	result, err := searchData(intKey, friends)

	if err  != nil {
		fmt.Println(err)
		return
	}
	
	result.PrintData()
}
