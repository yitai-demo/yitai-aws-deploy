package domain

import (
	"time"

	"github.com/degalaxy/gp-common/gormtype"
)

// --------------------
// Space
// --------------------

type SpaceStatus int

const (
	SPACE_STATUS_CREATING SpaceStatus = 0
	SPACE_STATUS_NORMAL   SpaceStatus = 1

	SPACE_ENTRY_POINT_TYPE_URL = "url"
)

type Space struct {
	Id             gormtype.UUIDBinary `gorm:"type:uuid;primary_key;"`
	Name           string
	Description    string
	EntryPointType string
	EntryPointUrl  string
	ImageUrl       string
	CosFolder      string
	DisplayOrder   int
	Status         SpaceStatus
	CreatedAt      time.Time `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt      time.Time `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*Space) TableName() string {
	return "space"
}

// --------------------
// Space Tag Mapping
// --------------------

type SpaceTagMapping struct {
	UserTagId gormtype.UUIDBinary `gorm:"type:uuid;primary_key;"`
	SpaceId   gormtype.UUIDBinary `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time           `gorm:"<-:create"` // Set to current time if it is zero on creating
	UpdatedAt time.Time           `gorm:"<-:update"` // Set to current time on updating or if it is zero on creating
}

func (*SpaceTagMapping) TableName() string {
	return "space_tag_mapping"
}
