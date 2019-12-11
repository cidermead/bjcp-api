package model

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Range struct{
	Low   string  `json:"Low"`
	High  string  `json:"High"`
}

type Stats struct{
	IBU   Range  `json:"IBU"`
	OG    Range  `json:"OG"`
	FG    Range  `json:"FG"`
	SRM   Range  `json:"SRM"`
	ABV   Range  `json:"ABV"`
}

type Category struct {
	gorm.Model
	CategoryId         string `json:"CategoryId"           gorm:"type:varchar(2)"`
	Type               string `json:"Type"                 gorm:"type:varchar(5)"`
	Name               string `json:"Name"                 gorm:"type:varchar(256)"`
	Notes              string `json:"Notes"                gorm:"type:varchar(256)"`
  Styles            []Style `gorm:"foreignkey:StyleId;association_foreignkey:CategoryId"` // One-To-Many relationship (has many - use Style's StyleID as foreign key)
}

type Style struct {
	gorm.Model
	StyleId            string `json:"styleId,omitempty"              gorm:"type:varchar(3)"`
	Name               string `json:"name,omitempty"                 gorm:"type:varchar(256)"`
	Aroma              string `json:"aroma,omitempty"                gorm:"type:text"`
	Appearance         string `json:"appearance,omitempty"           gorm:"type:text"`
	Flavor             string `json:"flavor,omitempty"               gorm:"type:text"`
	Mouthfeel          string `json:"mouthfeel,omitempty"            gorm:"type:text"`
	Impression         string `json:"impression,omitempty"           gorm:"type:text"`
	Comments           string `json:"comments,omitempty"             gorm:"type:text"`
	History            string `json:"history,omitempty"              gorm:"type:text"`
	Ingredients        string `json:"ingredients,omitempty"          gorm:"type:text"`
	Comparison         string `json:"comparison,omitempty"           gorm:"type:text"`
	EntryInstructions  string `json:"entryInstructions,omitempty"    gorm:"type:text"`

	SimilarStyles     []Style `json:",omitempty"                     gorm:"many2many:similar_styles;association_jointable_foreignkey:similar_id"`

	Examples   pq.StringArray `json:"examples,omitempty"             gorm:"type:varchar(64)[]"`
	Varieties  pq.StringArray `json:"varieties,omitempty"            gorm:"type:varchar(24)[]"`
	Tags       pq.StringArray `json:"tags,omitempty"                 gorm:"type:varchar(24)[]"`
	Stats      postgres.Jsonb `json:"stats,omitempty"                gorm:"type:json"`
	BeerExam             bool `json:"beerExam,omitempty"             gorm:"type:boolean"`

	CategoryId           uint `gorm:"index,omitempty"`   // Foreign key (belongs to)
	Category         Category `json:"category,omitempty"             gorm:"foreignkey:StyleId;association_foreignkey:ID"`

	//Author       *User      `json:"author"`
	//AuthorID     int        `json:"-"`
}
