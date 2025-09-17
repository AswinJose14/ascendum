package userEngine

import (
	"ascendum/models"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type UserEngine struct {
	Count int
}

func (e *UserEngine) GenerateUsers() {
	fmt.Println("Generating Users")
	for i := range e.Count {
		go func(i int) {
			user := &models.User{
				UserId: i,
				Answer: e.getAnswer(),
			}
			e.submitAnserWithRandomDelay(user)
		}(i)
	}
}

func (e *UserEngine) getAnswer() string {
	number := rand.Int31n(2)
	switch number {
	case 0:
		return "No"
	case 1:
		return "Yes"
	default:
		return "No"
	}
}

func (e *UserEngine) submitAnserWithRandomDelay(u *models.User) {
	delay := 10 + rand.Intn(991)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	e.submitAnser(u)
}

func (e *UserEngine) submitAnser(u *models.User) {
	jsonPayload, _ := json.Marshal(u)
	url := "http://127.0.0.1:8080/submit"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req) 
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
}
