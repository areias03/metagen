package listui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	Name, Desc string
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Name }

type ListModel struct {
	List list.Model
}

func InitialListModel(queryResults map[string]int) ListModel {
	items := []list.Item{}
	for key, value := range queryResults {
		switch value {
		case 0:
			items = append(items, Item{Name: key, Desc: "Not Found"})
		case 1:
			items = append(items, Item{Name: key, Desc: "Found"})
		}
	}

	return ListModel{
		List: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
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
	m.List.Title = "Found Items"
	return docStyle.Render(m.List.View())
}
