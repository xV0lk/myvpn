package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
)

type ListModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	width    int
	height   int
}

const (
	primary   = lip.Color("#15F5BA")
	secondary = lip.Color("#836FFF")
	accent    = lip.Color("#211951")
	text      = lip.Color("#F0F3FF")
)

var (
	appStyle     = lip.NewStyle().Padding(1, 0)
	cursorStyle  = lip.NewStyle().Foreground(lip.Color(primary)).Bold(true)
	titleStyle   = lip.NewStyle().Background(lip.Color(primary)).Foreground(lip.Color(accent)).Padding(0, 1).Bold(true)
	focusedStyle = lip.NewStyle().Foreground(lip.Color(primary)).Bold(true)
	actionStyle  = lip.NewStyle().Foreground(lip.Color(secondary)).Bold(true)
	textStyle    = lip.NewStyle().Foreground(lip.Color(text)).Faint(true)
)

func InitialListModel() ListModel {
	return ListModel{
		choices:  []string{"Conectar VPN", "Agregar VPN", "Eliminar VPN"},
		selected: make(map[int]struct{}),
	}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), "q":
			return m, tea.Quit
		case tea.KeyUp.String(), "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown.String(), "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case tea.KeyEnter.String(), tea.KeySpace.String():
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m ListModel) View() string {
	s := titleStyle.Render("Que deseas hacer?")
	s += "\n\n"
	for i, choice := range m.choices {

		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("->")
			choice = focusedStyle.Render(choice)
		}

		// if _, ok := m.selected[i]; ok {
		// 	choice = selectedStyle.Render(choice)
		// }

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += textStyle.Render(fmt.Sprintf("\nPress %s to quit\n", actionStyle.Render("q")))
	// view := appStyle.Render(lip.JoinVertical(lip.Left, h, s, f))
	// return lip.Place(m.width, m.height, lip.Center, lip.Center, view)
	return appStyle.Render(s)
}
