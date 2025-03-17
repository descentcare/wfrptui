package models

type Characteristic struct {
    name string
    Initial, Advances int
}

func (c Characteristic) GetName() string {
    return c.name
}

func (c Characteristic) GetCurrent() int {
    return c.Initial + c.Advances
}

func (c *Characteristic) Advance(points int) {
    c.Advances += points
}

func (c Characteristic) GetRating() int {
    return c.GetCurrent() / 10
}

func NewCharacteristic(name string, initial, advances int) *Characteristic {
    return &Characteristic {
        name: name,
        Initial: initial,
        Advances: advances,
    }
}

const (
    WeaponSkill      = "ББ"     // Ближний бой
    BallisticSkill   = "ДБ"     // Дальний бой
    Strength         = "С"      // Сила
    Toughness        = "В"      // Выносливость
    Initiative       = "И"      // Инициатива
    Agility          = "Пр"     // Проворство
    Dexterity        = "Л"      // Ловкость
    Intelligence     = "Инт"    // Интеллект
    Willpower        = "СВ"     // Сила воли
    Fellowship       = "Х"      // Харизма
)
var OrderedCharacteristicsKeys = [...]string {
    WeaponSkill,
    BallisticSkill,
    Strength,
    Toughness,
    Initiative,
    Agility,
    Dexterity,
    Intelligence,
    Willpower,
    Fellowship,
}


func GetNewCharacteristics() map[string]*Characteristic {
    return map[string]*Characteristic {
        WeaponSkill:    NewCharacteristic("Ближний бой", 0, 0),
        BallisticSkill: NewCharacteristic("Дальний бой", 0, 0),
        Strength:       NewCharacteristic("Сила", 0, 0),
        Toughness:      NewCharacteristic("Выносливость", 0, 0),
        Initiative:     NewCharacteristic("Инициатива", 0, 0),
        Agility:        NewCharacteristic("Проворство", 0, 0),
        Dexterity:      NewCharacteristic("Ловкость", 0, 0),
        Intelligence:   NewCharacteristic("Интеллект", 0, 0),
        Willpower:      NewCharacteristic("Сила воли", 0, 0),
        Fellowship:     NewCharacteristic("Харизма", 0, 0),
    }
}
