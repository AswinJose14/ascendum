package gameEngine

import (
	"ascendum/models"
	"fmt"
	"sync"
)

type GameEngine struct {
	CorrectAnswer      string
	Winner             *models.User
	CorrectAnswerCount int
	WrongAnswerCount   int
	mutex              *sync.Mutex
}

var gameEngine *GameEngine
var once sync.Once

func GetGameEngine() *GameEngine {
	if gameEngine == nil {
		once.Do(
			func() {
				gameEngine = &GameEngine{
					CorrectAnswer: "Yes",
					mutex:         &sync.Mutex{},
				}
			})
	}
	return gameEngine
}

func (ge *GameEngine) EvaluvateAnswer(user *models.User) {
	ge.mutex.Lock()
	if user.Answer == ge.CorrectAnswer {
		if ge.Winner == nil {
			ge.Winner = user

		}
		ge.CorrectAnswerCount++
	} else {
		ge.WrongAnswerCount++
	}
	ge.mutex.Unlock()
}

func AnnonceWinners() {
	ge := GetGameEngine()
	fmt.Println("Winner is userId: ", ge.Winner.UserId)
	fmt.Println("No of corrrect Answers: ", ge.CorrectAnswerCount)
	fmt.Println("No of wrong Answers: ", ge.WrongAnswerCount)
}
