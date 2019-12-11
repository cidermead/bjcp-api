package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)
type Report struct{
	Correct int  `json:"correct,omitempty"`
	Wrong   int  `json:"wrong,omitempty"`
}

type Question struct {
	gorm.Model
	Question           string `json:"question"   gorm:"type:varchar(355)"`
	Options    pq.StringArray `json:"options"    gorm:"type:text[]"`
	Answer                int `json:"answer"     gorm:"type:integer"`
	Topic              string `json:"topic"      gorm:"type:varchar(20)"`
	Exam               string `json:"exam"       gorm:"type:varchar(5)"`

	Stats      	      []uint8 `json:"stats"      gorm:"type:json"`
	Active               bool `json:"active"     gorm:"type:boolean"`
	Deleted              bool `json:"deleted"    gorm:"type:boolean"`

	// CreatedAt       time.Time `json:"created_at"`
	// UpdatedAt       time.Time `json:"updated_at"`
	// Description     string `json:"Description"  gorm:"type:text"`
	// Tags             []Tag  // One-To-Many relationship (has many - use Tag's UserID as foreign key)
}

