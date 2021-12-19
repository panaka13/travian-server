package model

type BuildingType int

const (
	EMPTY                  BuildingType = 0
	WOODCUTTER             BuildingType = 1
	CLAY_PIT               BuildingType = 2
	IRON_MINE              BuildingType = 3
	CROPLAND               BuildingType = 4
	SAWMILL                BuildingType = 5
	BRICKYARD              BuildingType = 6
	IRON_FOUNDRY           BuildingType = 7
	GRAIN_MILL             BuildingType = 8
	BAKERY                 BuildingType = 9
	WAREHOUSE              BuildingType = 10
	GRANARY                BuildingType = 11
	SMITHY                 BuildingType = 13
	TOURNAMENT_SQUARE      BuildingType = 14
	MAIN_BUILDING          BuildingType = 15
	RALLY_POINT            BuildingType = 16
	EMBASSY                BuildingType = 18
	BARRACKS               BuildingType = 19
	STABLE                 BuildingType = 20
	WORKSHOP               BuildingType = 21
	ACADEMY                BuildingType = 22
	CRANNY                 BuildingType = 23
	TOWN_HALL              BuildingType = 24
	RESIDENCE              BuildingType = 25
	PALACE                 BuildingType = 26
	TREASURY               BuildingType = 27
	TRADE_OFFICE           BuildingType = 28
	GREAT_BARRACKS         BuildingType = 29
	GREAT_STABLE           BuildingType = 30
	CITY_WALL              BuildingType = 31
	EARTH_WALL             BuildingType = 32
	PALLISADE              BuildingType = 33
	STONEMASONS_LODGE      BuildingType = 34
	TRAPPER                BuildingType = 36
	HEROS_MANSION          BuildingType = 37
	GREAT_WAREHOUSE        BuildingType = 38
	GREAT_GRANARY          BuildingType = 39
	HORSE_DRINKING_TROUGHT BuildingType = 41
	MAKESHIFT_WALL         BuildingType = 43
	COMMAND_CENTER         BuildingType = 44
	HOSPITAL               BuildingType = 46
)

type Structure struct {
	Building BuildingType
	Level    int
}

var WarehouseCapacity [21]int = [21]int{800, 1200, 1700, 2300, 3100, 4000, 5000, 6300, 7800, 9600, 11800, 14400, 17600, 21400, 25900, 31300, 37900, 45700, 55100, 66400, 80000}
var ResourceProduction [21]int = [21]int{3, 7, 13, 21, 31, 46, 70, 98, 140, 203, 280, 392, 525, 693, 889, 1120, 1400, 1820, 2240, 2800, 3430}

func (s *Structure) GetCapacity() int {
	switch s.Building {
	case WAREHOUSE, GRANARY:
		return WarehouseCapacity[s.Level-1]
	case GREAT_WAREHOUSE, GREAT_GRANARY:
		return WarehouseCapacity[s.Level-1] * 3
	}
	return 0
}

func (s *Structure) GetProduction() int {
	switch s.Building {
	case WOODCUTTER, CLAY_PIT, IRON_MINE, CROPLAND:
		return ResourceProduction[s.Level]
	}
	return 0
}
