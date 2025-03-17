package models

type Species struct {
    Name string
    CharacteristicRolls map[string]Roll
    AgeRoll Roll
    Resilience int
    Fate int
    Speed int
    Skills []Skill
}

var r2d10p20 Roll = NewRoll(10, 2, 20)

var PlayableSpecies []Species = []Species {
    {
        Name: "Человек",
        CharacteristicRolls: map[string]Roll {
            WeaponSkill: r2d10p20,
            BallisticSkill: r2d10p20,
            Strength: r2d10p20,
            Toughness: r2d10p20,
            Initiative: r2d10p20,
            Agility: r2d10p20,
            Dexterity: r2d10p20,
            Intelligence: r2d10p20,
            Willpower: r2d10p20,
            Fellowship: r2d10p20,
        },
        AgeRoll: NewRoll(10, 1, 15),
        Resilience: 1,
        Fate: 2,
        Speed: 4,
        Skills: []Skill {

        },
    },
}
