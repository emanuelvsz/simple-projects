package main

import (
	"encoding/csv"
	"fmt"
	"list/src/core/domain"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PersonDTO struct {
	ID        uuid.UUID `json:"id"`
	LastName  string `json:"surname"`
	ShirtSize string `json:"shirt_size"`
	Number    int8   `json:"number"`
	Shirt     bool   `json:"want_shirt"`
	Short     bool   `json:"want_short"`
}

func main() {
	e := echo.New()

	apiGroup := e.Group("/api")

	apiGroup.GET("/people", handlePeople)
	e.Logger.Fatal(e.Start(":8080"))
}

func handlePeople(context echo.Context) error {
	peopleInstance, err := GetPeople()
	if err != nil {
		return err
	}

	people := make([]PersonDTO, 0)
	for _, p := range peopleInstance {
		people = append(people, PersonDTO{
			ID: *p.ID(),
			LastName: strings.ToUpper(p.Surname()),
			ShirtSize: p.ShirtSize(),
			Number: p.Number(),
			Shirt: p.Shirt(),
			Short: p.Short(),
		})
	}

	return context.JSON(http.StatusOK, people)
}

func GetPeople() ([]domain.Person, error) {
	file, err := os.Open("./src/utils/data/data.csv")
	if err != nil {
		fmt.Println("Um erro na leitura do arquivo csv ocorreu.(", err, ")")
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Um erro na leitura dos dados do arquivo csv ocorreu.(", err, ")")
		return nil, err
	}

	p, err := FindPerson(records)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func FindPerson(records [][]string) ([]domain.Person, error) {
	var people = make([]domain.Person, len(records))
	for k, record := range records {
		surname := strings.Split(record[0], ";")[0]
		shirtSize := strings.Split(record[0], ";")[4]
		numberStr := strings.Split(record[0], ";")[3]
		shortStr := strings.Split(record[0], ";")[1]
		shirtStr := strings.Split(record[0], ";")[2]

		number, err := strconv.ParseInt(numberStr, 10, 8)
		if err != nil {
			fmt.Println("Erro ao converter number para int8:", err)
		}

		short, err := strconv.ParseBool(shortStr)
		if err != nil {
			fmt.Println("Erro ao converter short para bool:", err)
		}

		shirt, err := strconv.ParseBool(shirtStr)
		if err != nil {
			fmt.Println("Erro ao converter shirt para bool:", err)
		}

		person := domain.NewPerson(
			uuid.New(),
			surname,
			shirtSize,
			int8(number),
			shirt,
			short,
		)
		people[k] = *person
	}

	return people, nil
}
