package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	answerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("10"))

	responseStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1)

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

func PrintError(err error) {
	fmt.Println(errorStyle.Render("Error: " + err.Error()))
}