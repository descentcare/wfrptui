package main

import (
    "os"
    "fmt"
    "github.com/descentcare/wfrptui/internal/tui"
    tea "github.com/charmbracelet/bubbletea"
)


func main() {
    m := tui.Facade{ ChooseList: tui.InitCharchoose() }
    //if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
