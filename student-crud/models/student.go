package models

type Student struct {
    ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Grade string `json:"grade"`
}
