package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name string
	CPF  string
	Age  int
}

func main() {
	readDoc()  // this function read the csv document
	writeDoc() // this function create a csv document
}

func readDoc() {
	file, err := os.Open("./csv/people.csv")
	if err != nil {
		fmt.Println("Um erro na leitura do arquivo csv ocorreu.(", err, ")")
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Um erro na leitura dos dados do arquivo csv ocorreu.(", err, ")")
	}

	var people = make([]Person, len(records))
	for k, record := range records {
		idade, _ := strconv.Atoi(record[2])
		person := Person{
			Name: record[0],
			CPF:  record[1],
			Age:  idade,
		}
		people[k] = person
	}

	fmt.Println(people)
}

func writeDoc() {
	file, err := os.Create("./csv/new-people.csv")
	if err != nil {
		fmt.Println("Um erro na leitura do arquivo csv ocorreu.(", err, ")")
	}
	defer file.Close()

	records := [][]string{
		{"Tiago Temporin", "tiago.temporin@provedor.com"},
		{"Amanda Moreira", "amanda.moreira@hotmail.com"},
		{"João Santos", "joao.santos@gmail.me"},
		{"Valentina Silva", "valentina.silva@gmail.com"},
		{"Emanuel Vilela", "thundGames@outlook.cum"},
	}

	err = csv.NewWriter(file).WriteAll(records)
	if err != nil {
		panic(err)
	}
}
