package repository

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Login string
	Salt  string
	Role  Role
}

type Role struct {
	gorm.Model
	ID          int
	Name        string
	Description string
}

type Journal struct {
	gorm.Model
	ID           int
	AttackType   Attack
	ActivityType Activity
	Description  string
	Timestamp    time.Time
}

type UserActivity struct {
	gorm.Model
	ID           int
	ActivityType Activity
	Login        string
	Timestamp    time.Time
}

type Attack struct {
	gorm.Model
	ID          int
	Name        string
	Description string
}

type Activity struct {
	gorm.Model
	ID          int
	Name        string
	Description string
}
