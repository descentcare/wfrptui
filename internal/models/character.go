package models

type Character struct {
    Name string
    Characteristics map[string]*Characteristic
    BasicSkills map[string]*Skill
    AdvancedSkills map[string]*Skill
    Talents []Talent
    Species Species
}

func NewCharacter(name string) Character {
    characteristics := GetNewCharacteristics()
    basicSkills := make(map[string]*Skill)
    advancedSkills := make(map[string]*Skill)
    for k, v := range GetNewSkills(characteristics) {
        if v.Common {
            basicSkills[k] = &v
        } else {
            advancedSkills[k] = &v
        }
    }
    char := Character {
        Name: name,
        Characteristics: characteristics,
        BasicSkills: basicSkills,
        AdvancedSkills: advancedSkills,
    }
    return char
}
