package domain

import (
	"time"

	"github.com/degalaxy/gp-common/gormtype"
)

// --------------------
// User
// --------------------

type UserStatus int

const (
	USER_STATUS_REGISTERING UserStatus = 0
	USER_STATUS_ACTIVE      UserStatus = 1
	USER_STATUS_INACTIVE    UserStatus = 2
)

type User struct {
	Id        string `gorm:"type:varchar(255);primary_key;"`
	Status    UserStatus
	Did       string
	UserId    string
	CreatedAt time.Time `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt time.Time `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*User) TableName() string {
	return "user"
}

// --------------------
// User Tag
// --------------------

type UserTagStatus int

const (
	USER_TAG_STATUS_UNKNOWN UserTagStatus = 0
	USER_TAG_STATUS_NORMAL  UserTagStatus = 1
)

type UserTag struct {
	Id        gormtype.UUIDBinary `gorm:"type:uuid;primary_key;"`
	Name      string              `gorm:"uniqueIndex"`
	Status    UserTagStatus
	CreatedAt time.Time `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt time.Time `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*UserTag) TableName() string {
	return "user_tag"
}

// --------------------
// User Tag Mappiing
// --------------------

type UserTagMapping struct {
	UserId    string              `gorm:"primaryKey"`
	UserTagId gormtype.UUIDBinary `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time           `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt time.Time           `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*UserTagMapping) TableName() string {
	return "user_tag_mapping"
}

// --------------------
// User Roles
// --------------------

type UserRoleStatus int

const (
	USER_ROLE_STATUS_UNKNOWN    UserRoleStatus = 0
	USER_ROLE_STATUS_IN_USE     UserRoleStatus = 1
	USER_ROLE_STATUS_NOT_IN_USE UserRoleStatus = 2
)

type UserRole struct {
	Id        gormtype.UUIDBinary `gorm:"type:uuid;primary_key;"`
	Name      string              `gorm:"uniqueIndex"`
	Image     string
	Status    UserRoleStatus
	CreatedAt time.Time `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt time.Time `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*UserRole) TableName() string {
	return "user_role"
}

// ---------------------------
// User Role Mapping
// ---------------------------

type UserRoleMapping struct {
	UserId    string              `gorm:"primary_key;"`
	RoleId    gormtype.UUIDBinary `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time           `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt time.Time           `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*UserRoleMapping) TableName() string {
	return "user_role_mapping"
}
