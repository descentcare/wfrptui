package main

import (
    "bufio"
    "os"
    "fmt"
    "github.com/descentcare/wfrptui/internal/tui"
    "github.com/descentcare/wfrptui/internal/models"
    tea "github.com/charmbracelet/bubbletea"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Character Name: ")
    name, _ := reader.ReadString('\n')
    character := models.NewCharacter(name)
    m := tui.NewCharacteristicsTabModel(&character)
    //if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
