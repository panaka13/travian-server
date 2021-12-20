package model

import (
	"strconv"
	"strings"
	"time"
)

type Village struct {
	Name          string      `json:"name"`
	Id            int         `json:"id",gorm:"primaryKey"`
	User          User        `gorm:"foreignKey:Id"`
	Structures    []Structure `gorm:"-"`
	StructureByte string
	Production    ResourceSet `gorm:"-"`
	Resource      ResourceSet `gorm:"-"`
	ResourceByte  string
	UpdatedTime   time.Time
}

func (v *Village) UpdateResource(wood, clay, iron, wheat int) {
	v.UpdatedTime = time.Now()
	v.Resource.Wood = wood
	v.Resource.Clay = clay
	v.Resource.Iron = iron
	v.Resource.Wheat = wheat
	v.SerializeResource()
}

func (v *Village) Serialize() {
	v.SerializeResource()
	v.SerializeStructure()
}

func (v *Village) Deserialize() {
	v.DeserializeStructure()
	v.DeserializeProduction()
	v.DeserializeResource()
}

func (v *Village) SerializeResource() {
	v.ResourceByte = v.Resource.serialize()
}

func (v *Village) DeserializeResource() {
	v.Resource.deserialize(v.ResourceByte)
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
	v.DeserializeProduction()
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
	v.DeserializeProduction()
	v.SerializeStructure()
}

func (v *Village) DeserializeProduction() {
	var r ResourceSet
	for i := 0; i < 18; i++ {
		switch v.Structures[i].Building {
		case WOODCUTTER:
			r.Wood += v.Structures[i].GetProduction()
		case CLAY_PIT:
			r.Clay += v.Structures[i].GetProduction()
		case IRON_MINE:
			r.Iron += v.Structures[i].GetProduction()
		case CROPLAND:
			r.Wheat += v.Structures[i].GetProduction()
		}
	}
	r.Wood *= v.User.Speed
	r.Clay *= v.User.Speed
	r.Iron *= v.User.Speed
	r.Wheat *= v.User.Speed
	v.Production = r
}
