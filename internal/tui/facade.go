package tui

import (
    "github.com/descentcare/wfrptui/internal/models"
    "github.com/descentcare/wfrptui/internal/fsclient"
    "github.com/descentcare/wfrptui/internal/constants"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/key"

)

type Facade struct {
    Char *models.Character
    CharTab CharacteristicsTabModel
    ChooseList *Charchoose
    CurrentTab constants.CurrentTab
}

func (m Facade) Init() tea.Cmd {
    return nil
}

func (m Facade) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if key.Matches(msg, constants.Keymap.QuitProgram) {
            if err := fsclient.Save(*m.Char); err != nil {
                panic(err)
            }
            return m, tea.Quit
        }
    }
    switch m.CurrentTab {
        case constants.Choose:
            return m.ChooseList.Update(msg)
        case constants.Char:
            return m.CharTab.Update(msg)
    }
    return m, cmd
}

func (m Facade) View() string {
    return ""
}
