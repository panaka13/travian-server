package model

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Village struct {
	gorm.Model
	Name          string
	Id            int         `gorm:"primaryKey"`
	User          User        `gorm:"foreignKey:ID"`
	Structures    []Structure `gorm:"-"`
	StructureByte string
}

func (v *Village) SerializeStructure() {
	n := len(v.Structures)
	v.StructureByte = ""
	for i := 0; i < n; i++ {
		v.StructureByte += strconv.Itoa(int(v.Structures[i].Building)) + ";"
		v.StructureByte += strconv.Itoa(v.Structures[i].Level) + ";"
	}
	for i := n + 1; i < 40; i++ {
		v.StructureByte += "0;0;"
	}
}

func (v *Village) DeserializeStructure() {
	tokens := strings.Split(v.StructureByte, ";")
	v.Structures = make([]Structure, 40)
	for i := 0; i < 40; i++ {
		level, err := strconv.ParseInt(tokens[i*2+1], 10, 32)
		if err != nil {
			level = 0
		}
		building, err := strconv.ParseInt(tokens[i*2], 10, 32)
		if err != nil {
			building = 0
		}
		v.Structures = append(v.Structures, Structure{Building: BuildingType(building), Level: int(level)})
	}
}
