package save

import (
    models "github.com/descentcare/wfrptui/internal/models"
    "encoding/json"
    "os"
    "fmt"
)

func Save(character models.Character) error {
    bytes, err := json.Marshal(character)
    if err != nil {
        return err
    }
    sharedir, err := getShareDir();
    if err != nil {
        return err
    }
    f, err := os.Create(fmt.Sprintf("%s/%s.txt", sharedir, character.Name))
    defer f.Close()
    if err != nil {
        return err
    }
    
    if _, err = f.Write(bytes); err != nil {
        return err
    }

    return f.Sync()
}

func Unload(character *models.Character, name string) error {
    dir, err := getShareDir()
    if err != nil {
        return err
    }
    data, err := os.ReadFile(dir + name + ".txt")
    if err != nil {
        return err
    }
    return json.Unmarshal(data, character)
}

func makeShareDir() (string, error) {
    sharedir, err := getShareDir()
    if err != nil {
        return sharedir, err
    }
    return sharedir, os.MkdirAll(sharedir, 0770)
}

func getShareDir() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    sharedir := home + "/.local/share/wfrptui/"
    return sharedir, nil
}
