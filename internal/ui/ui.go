package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	answerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("10"))

	commandLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("11"))

	commandStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1).
			Width(76).
			Foreground(lipgloss.Color("11"))

	responseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1).
			Width(76)

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("9"))
)

func PrintResponse(response string) {
	title := answerStyle.Render("Answer")
	body := responseStyle.Render(response)

	fmt.Println(title)
	fmt.Println(body)
}

func PrintCommand(command string, explanation string) {
	commandTitle := commandLabelStyle.Render("Command")
	commandBody := commandStyle.Render(command)

	explanationTitle := answerStyle.Render("Explanation")
	explanationBody := responseStyle.Render(explanation)

	fmt.Println(commandTitle)
	fmt.Println(commandBody)
	fmt.Println(explanationTitle)
	fmt.Println(explanationBody)
}

func PrintError(err error) {
	fmt.Println(errorStyle.Render("Error: " + err.Error()))
}