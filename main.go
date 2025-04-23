package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"strconv"
)

const (
	MaxGuessesAmount = 10
	MaxGuess         = 300
	MinGuess         = 0
)

func main() {
	// Создаем приложение и окно
	myApp := app.New()
	myWindow := myApp.NewWindow("Угадайка!")

	var tryGuessButton *widget.Button
	var guessDisplay *widget.Label
	var userGuessInput *widget.Entry
	var err error

	currentAttemptsCount := 0
	playerValue := -1
	gameValue := -1

	// 1/5 компонент – Лэйбл который показывает кол-во попыток
	computerGuessLabel := widget.NewLabel("Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!")

	// 2/5 компонент – Кнопка старта игры, которая назначает новое значение computerGuess, guessesLeft и начинает игру
	startGameButton := widget.NewButton("Загадать новое число!", func() {
		gameValue = rand.Int() % MaxGuess
		currentAttemptsCount = 0
		playerValue = -1
		computerGuessLabel.SetText("Компьютер загадал число, угадай его!")
		tryGuessButton.Enable()
		userGuessInput.Enable()
		guessDisplay.SetText("Введите ваше первое число!")

	})

	// 3/5 компонент – Лэйбл который подсказывает выше ли наше число загаданного или ниже
	guessDisplay = widget.NewLabel("Введите ваше первое число!")

	// 4/5 компонент – Инпут для пользовательского ввода
	userGuessInput = widget.NewEntry()
	userGuessInput.SetPlaceHolder("Введите ваше число: ")

	// 5/5 компонент – Кнопка "Попробовать"
	tryGuessButton = widget.NewButton("Попробовать", func() {
		if gameValue == -1 {
			guessDisplay.SetText("Вы забыли дать загадать число компьютеру!")
		}

		userInput := userGuessInput.Text
		playerValue, err = strconv.Atoi(userInput)
		if err != nil {
			guessDisplay.SetText("Упс, похоже вы ввели не число! Попробуйте еще раз")
			return
		}
		if playerValue > MinGuess && playerValue < MaxGuess {
			currentAttemptsCount++
			if currentAttemptsCount == MaxGuessesAmount {
				guessDisplay.SetText("Увы, вы исчерпали все попытки и оказались в числе проигравших!")
				tryGuessButton.Disable()
				userGuessInput.Disable()
				return
			}
			if gameValue == playerValue {
				currentAttemptsCount--
				guessDisplay.SetText(fmt.Sprintf("Поздравляю, вы угадали число! У вас оставалось еще %d попыток", MaxGuessesAmount-currentAttemptsCount))
				return
			}
			if playerValue > gameValue {
				guessDisplay.SetText(fmt.Sprintf("Ваше число больше загаданного! У вас осталось еще %d попыток", MaxGuessesAmount-currentAttemptsCount))
				return
			}
			if playerValue < gameValue {
				guessDisplay.SetText(fmt.Sprintf("Ваше число меньше загаданного! У вас осталось еще %d попыток", MaxGuessesAmount-currentAttemptsCount))
				return
			}
		} else {
			guessDisplay.SetText("Ваше число должно быть от 0 до 300")
		}
	})
	tryGuessButton.Disable()
	userGuessInput.Disable()
	// пихаем все виджеты которые мы создали в наше окно
	myWindow.SetContent(
		container.NewVBox(
			computerGuessLabel,
			startGameButton,
			guessDisplay,
			userGuessInput,
			tryGuessButton,
		),
	)

	// назначаем размер окна и начинаем программу!
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
