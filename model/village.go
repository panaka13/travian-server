package model

import (
	"strconv"
	"strings"
)

type Village struct {
	Name          string      `json:"name"`
	Id            int         `json:"id",gorm:"primaryKey"`
	User          User        `gorm:"foreignKey:Id"`
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
	for len(tokens) < 80 {
		tokens = append(tokens, "0")
	}
	v.Structures = make([]Structure, 0, 40)
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

func (v *Village) PartialDeserialize(structureByte string, part string) {
	v.DeserializeStructure()
	tokens := strings.Split(structureByte, ";")
	var start, end int
	if part == "1" {
		start = 0
		end = 18
	} else {
		start = 18
		end = 40
	}
	for i, j := start, 0; i < end; i, j = i+1, j+1 {
		level, err := strconv.ParseInt(tokens[j*2+1], 10, 32)
		if err != nil {
			level = 0
		}
		building, err := strconv.ParseInt(tokens[j*2], 10, 32)
		if err != nil {
			building = 0
		}
		v.Structures[i].Level = int(level)
		v.Structures[i].Building = BuildingType(building)
	}
	v.SerializeStructure()
}
