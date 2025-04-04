package tui

import (
    "strconv"
    "fmt"
    "math/rand"
    "github.com/descentcare/wfrptui/internal/models"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/table"
)

var baseStyle = lipgloss.NewStyle().
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("240"))

type CharacteristicsTabModel struct {
    characteristicsTable table.Model
    char *models.Character
    text string
}

func (m CharacteristicsTabModel) Init() tea.Cmd {
    return nil
}

func (m CharacteristicsTabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "esc":
            if m.characteristicsTable.Focused() {
                m.characteristicsTable.Blur()
            } else {
                m.characteristicsTable.Focus()
            }
        case "+":
            m.char.Characteristics[m.characteristicsTable.SelectedRow()[1]].Advance(1)
            m.characteristicsTable.SetRows(getCharacteristicsRows(m.char))
        case "-":
            m.char.Characteristics[m.characteristicsTable.SelectedRow()[1]].Advance(-1)
            m.characteristicsTable.SetRows(getCharacteristicsRows(m.char))
        case "enter":
            m.text = fmt.Sprintf("%s: %v             ",
                m.characteristicsTable.SelectedRow()[0],
                rand.Int() % 100)
            return m, nil
        }
    }
    m.characteristicsTable, cmd  = m.characteristicsTable.Update(msg)
    return m, cmd
}

func (m CharacteristicsTabModel) View() string {
    return baseStyle.Render(m.characteristicsTable.View()) + "\n" + m.text
}

func NewCharacteristicsTabModel(character *models.Character) CharacteristicsTabModel {
    columns := []table.Column{
        {Title: "Characteristic", Width: 15},
        {Title: "Short", Width: 5},
        {Title: "Initial", Width: 8},
        {Title: "Advances", Width: 9},
        {Title: "Current", Width: 8},
        {Title: "Rating", Width: 6},
    }
    rows := getCharacteristicsRows(character)
    t := table.New(
        table.WithColumns(columns),
        table.WithRows(rows),
        table.WithFocused(true),
        table.WithHeight(11),
    )
    s := table.DefaultStyles()
    s.Header = s.Header.
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("240")).
        BorderBottom(true).
        Bold(false)
    s.Selected = s.Selected.
        Foreground(lipgloss.Color("229")).
        Background(lipgloss.Color("57")).
        Bold(false)
    t.SetStyles(s)
    m := CharacteristicsTabModel{t, character, ""}
    return m
}

func getCharacteristicsRows(character *models.Character) []table.Row {
    rows := make([]table.Row, 0)
    
    for _, abb := range models.OrderedCharacteristicsKeys {
        c := character.Characteristics[abb]
        rows = append(rows, table.Row{
            c.GetName(),
            abb,
            strconv.Itoa(c.Initial),
            strconv.Itoa(c.Advances),
            strconv.Itoa(c.GetCurrent()),
            strconv.Itoa(c.GetRating()),
        })
    }

    return rows
}
