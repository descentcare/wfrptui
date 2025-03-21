package fsclient

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
    f, err := os.Create(fmt.Sprintf("%s/%s.json", sharedir, character.Name))
    defer f.Close()
    if err != nil {
        return err
    }
    
    if _, err = f.Write(bytes); err != nil {
        return err
    }

    return f.Sync()
}

func Load(character *models.Character, name string) error {
    dir, err := getShareDir()
    if err != nil {
        return err
    }
    data, err := os.ReadFile(dir + name + ".json")
    if err != nil {
        return err
    }
    return json.Unmarshal(data, character)
}

func ScanSaves() ([]string, error) {
    dir, err := getShareDir()
    if err != nil {
        return nil, err
    }
    entries, err := os.ReadDir(dir)
    if err != nil {
        return nil, err
    }
    saves := make([]string, 0)
    for _, e := range entries {
        if !e.IsDir() {
            name := e.Name()
            saves = append(saves, name[:len(name)-5])
        }
    }
    return saves, nil
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
