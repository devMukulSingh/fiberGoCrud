package model

import(
)


type Blog struct{
	Id    int			`json:"id" gorm:"primaryKey"`	
	Title string 		`json:"title" gorm:"not null"`
	Post string			`json:"post" gorm:"not null"`
}