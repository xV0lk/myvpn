package tui

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type netModel struct {
	status int
	err    error
}

type statusMsg int
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}

	res, err := c.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)
}

func (m netModel) Init() tea.Cmd {
	return checkServer
}

func (m netModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m netModel) View() string {
	if m.err != nil {
		return titleStyle.Render(fmt.Sprintf("\nHubo un error: %v\n\n", m.err))
	}

	s := focusedStyle.Render(fmt.Sprintf("Conectando con %s...\n\n", url))

	if m.status > 0 {
		s += focusedStyle.Render(fmt.Sprintf("%d %s\n", m.status, http.StatusText(m.status)))
	}
	return "\n" + s + "/n/n"
}
