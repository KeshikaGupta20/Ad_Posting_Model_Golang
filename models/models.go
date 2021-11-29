package models

import (
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model

	// @All people details in the system
	UId      string    `gorm:"primary_key;auto_increment" json:"UId"`
	Name     string
	Email    string `gorm:"typevarchar(100);unique_index"`
	Phone    int64
	Address  string
	Password string `gorm:"size:100;not null;"`
	Posts  []Post
}

type Ads struct {
	Title       string
	Type        string
	Description string
	Price       float64
}

type Response struct {
	Message string `json:"message"`
}

type Post struct {

	gorm.Model

	PID       string  
	Title     string    
	Content   string    
	//Owner   []Register      
	OwnerID  string    
	Image string 

	
}

