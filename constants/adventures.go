package constants

const (
	ADVENTURE_POND_CONTRACT       = "0x85e66216fb0e80f87b54eb39a415c3bbd40e37f9"
	ADVENTURE_STREAM_CONTRACT     = "0x780feb71117157a039e682668d79584d18579e90"
	ADVENTURE_SWAMP_CONTRACT      = "0xec7e923e7e0bd2dc7bb2ac0fabccf4e650c5418c"
	ADVENTURE_RIVER_CONTRACT      = "0x4eef52b71bd64d54d736cf2f3073e6dbbfcc7e31"
	ADVENTURE_FOREST_CONTRACT     = "0xcd32ed513a86484688cd3dbada05a9ed3c0c0eb6"
	ADVENTURE_GREAT_LAKE_CONTRACT = "0x1009cba3c0a50a2a0e8a92bc070ac5ffb8a3efe2"
)

type Adventure int

const (
	AdventurePond Adventure = iota
	AdventureStream
	AdventureSwamp
	AdventureRiver
	AdventureForest
	AdventureGreatLake
)
