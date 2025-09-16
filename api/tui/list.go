package tui

import (
	"fmt"
	"os"

	"github.com/areias03/metagen/api/db"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var DBs db.Databases = **db.Conf

type Item struct {
	Name, Desc string
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Name }

type ListModel struct {
	List list.Model
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {
	return docStyle.Render(m.List.View())
}

func main() {
	items := []list.Item{}
	for _, v := range DBs.Databases {
		items = append(items, Item{Name: v.Name, Desc: v.Match})

	}

	m := ListModel{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Found Items"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
