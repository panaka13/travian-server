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
