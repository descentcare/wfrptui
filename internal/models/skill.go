package models

type Skill struct {
    Name string
    Base *Characteristic
    Advances int
    Common bool
}

func (s Skill) GetName() string {
    return s.Name
}

func (s Skill) GetCurrent() int {
    return s.Base.GetCurrent() + s.Advances
}

func (s *Skill) Advance(points int) {
    s.Advances += points
}

func (s Skill) GetRating() int {
    return s.GetCurrent() / 10
}

func NewSkill(name string, base *Characteristic, advances int, common bool) Skill {
    return Skill {
        Name: name,
        Base: base,
        Advances: advances,
        Common: common,
    }
}

func GetNewSkills(characteristics map[string]*Characteristic) map[string]Skill {
    return map[string]Skill {
        "Азартные игры":          NewSkill("Азартные игры",         characteristics[Intelligence],      0, true),
        "Артистизм":              NewSkill("Артистизм",             characteristics[Fellowship],        0, true),   //Группа навыков 
        "Атлетика":               NewSkill("Атлетика",              characteristics[Agility],           0, true),
        "Верховая езда":          NewSkill("Верховая езда",         characteristics[Agility],           0, true),   //Группа навыков 
        "Взлом":                  NewSkill("Взлом",                 characteristics[Dexterity],         0, false),
        "Вождение":               NewSkill("Вождение",              characteristics[Agility],           0, true),
        "Выживание":              NewSkill("Выживание",             characteristics[Intelligence],      0, true),
        "Выслеживание":           NewSkill("Выслеживание",          characteristics[Initiative],        0, false),
        "Гребля":                 NewSkill("Гребля",                characteristics[Strength],          0, true),
        "Дрессировка":            NewSkill("Дрессировка",           characteristics[Intelligence],      0, false),  //Группа навыков 
        "Запугивание":            NewSkill("Запугивание",           characteristics[Strength],          0, true),
        "Знание":                 NewSkill("Знание",                characteristics[Intelligence],      0, false),  //Группа навыков 
        "Интуиция":               NewSkill("Интуиция",              characteristics[Initiative],        0, true),
        "Искусство":              NewSkill("Искусство",             characteristics[Dexterity],         0, true),   //Группа навыков 
        "Книжные изыскания":      NewSkill("Книжные изыскания",     characteristics[Intelligence],      0, false),
        "Концентрация":           NewSkill("Концентрация",          characteristics[Willpower],         0, false),  //Группа навыков
        "Кутёж":                  NewSkill("Кутёж",                 characteristics[Toughness],         0, true),
        "Лазание":                NewSkill("Лазание",               characteristics[Strength],          0, true),
        "Лечение":                NewSkill("Лечение",               characteristics[Intelligence],      0, false),
        "Лидерство":              NewSkill("Лидерство",             characteristics[Fellowship],        0, true),
        "Ловкость рук":           NewSkill("Ловкость рук",          characteristics[Dexterity],         0, false),
        "Молитвословие":          NewSkill("Молитвословие",         characteristics[Fellowship],        0, false),
        "Музицирование":          NewSkill("Музицирование",         characteristics[Dexterity],         0, false),  //Группа навыков 
        "Наблюдательность":       NewSkill("Наблюдательность",      characteristics[Initiative],        0, true),
        "Обаяние":                NewSkill("Обаяние",               characteristics[Fellowship],        0, true),
        "Обращение с животными":  NewSkill("Обращение с животными", characteristics[Intelligence],      0, false),
        "Обращение с ловушками":  NewSkill("Обращение с ловушками", characteristics[Dexterity],         0, false),
        "Ориентирование":         NewSkill("Ориентирование",        characteristics[Initiative],        0, true),
        "Оценка":                 NewSkill("Оценка",                characteristics[Intelligence],      0, false),
        "Плавание":               NewSkill("Плавание",              characteristics[Strength],          0, false),
        "Подкуп":                 NewSkill("Подкуп",                characteristics[Fellowship],        0, true),
        "Ремесло":                NewSkill("Ремесло",               characteristics[Dexterity],         0, false),  //Группа навыков 
        "Рукопашный бой":         NewSkill("Рукопашный бой",        characteristics[WeaponSkill],       0, true),   //Группа навыков
        "Скрытность":             NewSkill("Скрытность",            characteristics[Agility],           0, true),   //Группа навыков 
        "Сплетничество":          NewSkill("Сплетничество",         characteristics[Fellowship],        0, true),
        "Стойкость":              NewSkill("Стойкость",             characteristics[Toughness],         0, true),
        "Стрельба":               NewSkill("Стрельба",              characteristics[BallisticSkill],    0, false),  //Группа навыков
        "Сценическое искусство":  NewSkill("Сценическое искусство", characteristics[Agility],           0, false),  //Группа навыков 
        "Тайные знаки":           NewSkill("Тайные знаки",          characteristics[Intelligence],      0, false),  //Группа навыков 
        "Торговля":               NewSkill("Торговля",              characteristics[Fellowship],        0, true),
        "Уклонение":              NewSkill("Уклонение",             characteristics[Agility],           0, true),
        "Усмирение животных":     NewSkill("Усмирение животных",    characteristics[Willpower],         0, true),
        "Хладнокровие":           NewSkill("Хладнокровие",          characteristics[Willpower],         0, true),
        "Хождение под парусом":   NewSkill("Хождение под парусом",  characteristics[Agility],           0, false),  //Группа навыков 
        "Язык":                   NewSkill("Язык",                  characteristics[Intelligence],      0, false),  //Группа навыков 
    }
}
