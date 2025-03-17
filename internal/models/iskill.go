package models

type Iskill interface {
    GetName() string
    GetCurrent() int
    Advance(int)
    GetRating() int
}
