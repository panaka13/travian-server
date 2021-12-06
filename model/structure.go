package model

import (
	"gorm.io/gorm"
)

type BuildingType int

const (
	EMPTY         BuildingType = 0
	WOODCUTTER    BuildingType = 1
	CLAY_PIT      BuildingType = 2
	IRON_MINE     BuildingType = 3
	CROPLAND      BuildingType = 4
	WAREHOUSE     BuildingType = 10
	GRANARY       BuildingType = 11
	MAIN_BUILDING BuildingType = 15
)

type Structure struct {
	gorm.Model
	Building BuildingType
	Level    int
}
