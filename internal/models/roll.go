package models

import "math/rand"

type Roll struct {
    Dice int
    DiceCount int
    Bonus int
}

func NewRoll(dice, diceCount, bonus int) Roll {
    return Roll {
        Dice: dice,
        DiceCount: diceCount,
        Bonus: bonus,
    }
}

func (r Roll) Roll() int {
    sum := 0
    for i := 0; i < r.DiceCount; i++ {
        sum += rand.Intn(r.Dice) + 1
    }
    return sum + r.Bonus
}
