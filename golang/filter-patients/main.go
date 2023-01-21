package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	var name *string
	var n string
	n = "felipe2"
	name = &n

	fmt.Println(*name)

	getChoose()

}

func getChoose() {
	var n int

	fmt.Println("\n1 - List all patients")
	fmt.Println("2 - List active patients")
	fmt.Println("3 - List patients by CPF")
	fmt.Println("4 - List patients by RG")
	fmt.Println("5 - List patients names")
	fmt.Println()

	fmt.Scan(&n)

	returnChoose(n)
}

func returnChoose(n int) {
	if n == 1 {
		listPatients()
	} else if n == 2 {
		listActivePatients()
	} else if n == 3 {
		listPatientsByCPF()
	} else if n == 4 {
		listPatientsByRG()
	} else if n == 5 {
		listWords()
	}
}

type Patient struct {
	Name     string
	CPF      string
	RG       string
	IsActive bool
}

var patients = []Patient{
	{
		Name:     "Felipe da Silva Santos",
		CPF:      "10820795020",
		RG:       "467123603",
		IsActive: true,
	},
	{
		Name:     "Emanuel Vilela de Souza",
		CPF:      "10553971409",
		RG:       "467123603",
		IsActive: true,
	},
	{
		Name:     "Douglas Pereira Santos Silva",
		CPF:      "29807641080",
		RG:       "104996493",
		IsActive: false,
	},
	{
		Name:     "Pedro da Silva Domingos",
		CPF:      "40642123047",
		RG:       "334987684",
		IsActive: true,
	},
}

var patientActive []Patient
var patientCPF []Patient
var patientRG []Patient
var patientName []Patient

func isActive(pt Patient) bool {
	return pt.IsActive == true
}

func isCPFEqual(pt Patient) string {

	return pt.CPF
}

func isRGEqual(pt Patient) string {
	return pt.RG
}

func isNameEqual(pt Patient) string {
	return pt.Name
}

// list methods

func listPatients() {
	fmt.Println(" ")
	fmt.Println("All Patients: ")

	fmt.Println(patients)
}

func listActivePatients() {
	for _, patient := range patients {
		if isActive(patient) { // aqui ficaria o queryParam, no local do booleano do if
			patientActive = append(patientActive, patient)
		}
	}
	fmt.Println("\nIsActives: ")

	fmt.Println(patientActive)
}

func listPatientsByCPF() {
	for _, patient := range patients {
		if isCPFEqual(patient) == "29807641080" { // aqui ficaria o queryParam, no local da string escrita
			patientCPF = append(patientCPF, patient)
		}
	}

	fmt.Println("\nByCPF: ")

	for _, u := range patientCPF {
		fmt.Println(u)
	}
}

func listPatientsByRG() {
	for _, patient := range patients {
		if isRGEqual(patient) == "334987684" { // aqui ficaria o queryParam, no local da string escrita
			patientRG = append(patientRG, patient)
		}
	}
	fmt.Println("\nByRG: ")

	for _, u := range patientRG {
		fmt.Println(u)
	}

	checkDuplicatedValues := checkDuplicatedValues(rgs)
	fmt.Println(checkDuplicatedValues)

}

func listWords() (int, error) {

	var patientEqualName []Patient

	for _, patient := range patients {
		fmt.Println(filterWord(isNameEqual(patient)))

		if filterWord(isNameEqual(patient)) == true {
			patientEqualName = append(patientEqualName, patient)
		}
	}

	fmt.Println("-----------------------------")

	return fmt.Println(patientEqualName)
}

func filterWord(w string) bool {
	w = strings.ToLower(w)
	matched, err := regexp.MatchString(strings.ToLower(`Silva`), w)
	if err != nil {
		fmt.Println(err)
	}

	return matched
}

// Filters

func checkDuplicatedValues(intSlice []string) []string { // tambem pensei em deixar essa função somente, melhor do que criar duas pro cpf e rg, dai eu mudaria a mensagem abaixo de lugar
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			fmt.Println("\nAlgo está errado com seu RG(334987684), pois há outro igual no nosso sistema. Por favor, entre em contato com a unidade de saúde") // aqui nesse cpf eu estava tentando fazer uma validação, caso tenha um RG duplo, também pensava em implementar ao CPF
		}
	}
	return list
}

var rgs = []string{ // aqui, no caso, seria o patient.RG
	"334987684",
	"334987684",
	"334987684",
	"334987684",
	"334987684",
}
