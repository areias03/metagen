package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

var (
	p *tea.Program
)

type sessionSate int

const (
	searchView sessionSate = iota
	listView
)

type SearchModel struct {
	state         sessionSate
	textinput     tea.Model
	list          tea.Model
	query         string
	searchResults map[string]int
}

func (m SearchModel) Init() tea.Cmd {
	return nil
}

func (m SearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case listui.FindMsg:
		m.state = searchView
	case searchui.SelectMsg:
		m.state = listView
	}
	switch m.state {
	case searchView:
	}
	return m, tea.Batch(cmds...)
}

func (m SearchModel) View() string {
	switch m.state {
	case listView:
		return m.list.View()
	default:
		return m.textinput.View()
	}
}
