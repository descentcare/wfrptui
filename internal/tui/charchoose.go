package tui

import (
    "fmt"
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/lipgloss"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/descentcare/wfrptui/internal/fsclient"
)

var docStyle = lipgloss.NewStyle().Margin(1,2)

type item struct {
    index int
    title string
}

func (i item) Title() string { return fmt.Sprintf("%d. %s", i.index, i.title) }
func (i item) Description() string { return i.title }
func (i item) FilterValue() string { return i.title }

type Charchoose struct {
    list list.Model
}

func InitCharchoose() *Charchoose {
    names, err := fsclient.ScanSaves()
    if err != nil {
        panic(err)
    }
    items := make([]list.Item, 0)
    for i, name := range names {
        items = append(items, item{ i + 1, name })
    }
    d := list.NewDefaultDelegate()
    d.ShowDescription = false
    m := &Charchoose{ list.New(items, d, 0, 0) }
    m.list.DisableQuitKeybindings()
    m.list.Title = "Choose or create a character"
    return m
}

func (m Charchoose) Init() tea.Cmd {
    return nil
}

func (m Charchoose) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "enter":
            selectedItem := m.list.SelectedItem().(item)
            selectedItem.title
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width - h, msg.Height - v)
    }
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m Charchoose) View() string {
    return docStyle.Render(m.list.View())
}
