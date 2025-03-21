package constants

import (
    "github.com/charmbracelet/bubbles/key"
)

type keymap struct {
    QuitProgram key.Binding
    Enter key.Binding
}

var Keymap = keymap {
    QuitProgram: key.NewBinding(
        key.WithKeys("ctrl+c"),
        key.WithHelp("ctrl+c", "Quit program"),
    ),
    Enter: key.NewBinding(
        key.WithKeys("enter"),
        key.WithHelp("Enter", "Enter"),
    ),
}

type CurrentTab int
const (
    Choose CurrentTab = iota
    Create
    Char
)
