package config

type Run string

const (
	BishiboshRun        Run = "bishibosh"
	CountessRun         Run = "countess"
	AndarielRun         Run = "andariel"
	AncientTunnelsRun   Run = "ancient_tunnels"
	SummonerRun         Run = "summoner"
	MephistoRun         Run = "mephisto"
	CouncilRun          Run = "council"
	EldritchRun         Run = "eldritch"
	PindleskinRun       Run = "pindleskin"
	NihlathakRun        Run = "nihlathak"
	TristramRun         Run = "tristram"
	LowerKurastRun      Run = "lower_kurast"
	LowerKurastChestRun Run = "lower_kurast_chest"
	StonyTombRun        Run = "stony_tomb"
	PitRun              Run = "pit"
	ArachnidLairRun     Run = "arachnid_lair"
	TalRashaTombsRun    Run = "tal_rasha_tombs"
	BaalRun             Run = "baal"
	DiabloRun           Run = "diablo"
	CowsRun             Run = "cows"
	LevelingRun         Run = "leveling"
	TerrorZoneRun       Run = "terror_zone"
)

var AvailableRuns = map[Run]interface{}{
	BishiboshRun:        nil,
	CountessRun:         nil,
	AndarielRun:         nil,
	AncientTunnelsRun:   nil,
	SummonerRun:         nil,
	MephistoRun:         nil,
	CouncilRun:          nil,
	EldritchRun:         nil,
	PindleskinRun:       nil,
	NihlathakRun:        nil,
	TristramRun:         nil,
	LowerKurastRun:      nil,
	LowerKurastChestRun: nil,
	StonyTombRun:        nil,
	PitRun:              nil,
	ArachnidLairRun:     nil,
	TalRashaTombsRun:    nil,
	BaalRun:             nil,
	DiabloRun:           nil,
	CowsRun:             nil,
	LevelingRun:         nil,
	TerrorZoneRun:       nil,
}
