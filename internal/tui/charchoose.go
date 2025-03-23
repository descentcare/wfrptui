package tui

import (
	"fmt"
    "github.com/descentcare/wfrptui/internal/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/descentcare/wfrptui/internal/fsclient"
	"github.com/descentcare/wfrptui/internal/models"
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
    //m.list.DisableQuitKeybindings()
    m.list.KeyMap.Quit.SetEnabled(false)
    m.list.Title = "Choose or create a character"
    return m
}

func (m Charchoose) Init() tea.Cmd {
    return nil
}

func (m Charchoose) Update(msg tea.Msg) (Charchoose, tea.Cmd) {
    var cmds []tea.Cmd
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if key.Matches(msg, constants.Keymap.Enter) {
            selectedItem := m.list.SelectedItem().(item)
            cmds = append(cmds, func () tea.Msg {
                c := models.NewCharacter(selectedItem.title)
                if err := fsclient.Load(&c, selectedItem.title); err != nil {
                    panic(err)
                }
                return c
            })
        } else if key.Matches(msg, constants.Keymap.NewCharacter) {
            fmt.Print("skhdajhfdsallfads")
            cmds = append(cmds, func () tea.Msg {
                fmt.Print("ERFTGYUHIJ")
                c := models.NewCharacter(fmt.Sprintf("New_%d", len(m.list.Items())))
                if err := fsclient.Save(c); err != nil {
                    panic(err)
                }
                return c
            })
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width - h, msg.Height - v)
    }
    m.list, cmd = m.list.Update(msg)
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}

func (m Charchoose) View() string {
    return docStyle.Render(m.list.View())
}
