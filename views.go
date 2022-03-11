package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	// selected is the global that stores toggle state
	selected = map[string]bool{}
)

func (i choice) Title() string {

	if selected[i.name] {
		return "[x] " + i.name
	}
	return "[ ] " + i.name
}
func (i choice) String() string      { return i.name }
func (i choice) Description() string { return "    " + i.desc }
func (i choice) FilterValue() string { return i.name }

type listKeyMap struct {
	togglechoice  key.Binding
	finishchoices key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		togglechoice: key.NewBinding(
			key.WithKeys(""),
			key.WithHelp("<space>", "toggle"),
		),
		finishchoices: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("<enter>", "execute selected"),
		),
	}
}

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			selected = map[string]bool{}
			return m, tea.Quit
		}

		if msg.String() == " " {
			i, ok := m.list.SelectedItem().(choice)
			if ok {
				if selected[i.name] {
					selected[i.name] = false
				} else {
					selected[i.name] = true
				}
			}
			return m, nil
		}

		if msg.String() == "enter" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func selectChoices(_ context.Context, choices []choice) ([]choice, error) {
	items := []list.Item{}
	byName := map[string]choice{}

	for _, c := range choices {
		byName[c.name] = c
		items = append(items, c)
		selected[c.name] = true
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "ioc-bench"

	listKeys := newListKeyMap()
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{listKeys.finishchoices, listKeys.togglechoice}
	}

	m := model{list: l}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	sChoices := []choice{}
	for name, enabled := range selected {
		if enabled {
			sChoices = append(sChoices, byName[name])
		}
	}

	sort.SliceStable(sChoices, func(i, j int) bool {
		return sChoices[i].name < sChoices[j].name
	})

	return sChoices, nil
}

func termWidth() int {
	width, _, _ := terminal.GetSize(0)
	if width < 1 {
		return 78
	}
	return width
}

func status(title string) {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4"))
	fmt.Print(style.Render(title))
	style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3D16A4"))
	fmt.Println(style.Render(" ..."))
}

func msg(title string) {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#CDF456"))
	fmt.Println(style.Render(title))
}

func announce(title string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		MarginTop(1).
		PaddingLeft(4).
		PaddingRight(4).
		Width(termWidth())

	fmt.Println(style.Render(title))
}

func subtitle(title string) {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999")).
		Background(lipgloss.Color("#3D16A4")).
		PaddingLeft(4).
		PaddingRight(4).
		Width(termWidth())

	fmt.Println(style.Render(title))
}

func notice(title string) {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFF00")).
		Background(lipgloss.Color("#551111")).
		PaddingLeft(4).
		PaddingRight(4).
		Width(termWidth())

	fmt.Println(style.Render(title))
}
