package run

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/config"
)

var andarielStartingPosition = data.Position{
	X: 22561,
	Y: 9553,
}

type Andariel struct {
	baseRun
}

func (a Andariel) Name() string {
	return string(config.AndarielRun)
}

func (a Andariel) BuildActions() (actions []action.Action) {
	actions = append(actions,
		a.builder.WayPoint(area.CatacombsLevel2), // Moving to starting point (Catacombs Level 2)
		a.builder.MoveToArea(area.CatacombsLevel3),
		a.builder.MoveToArea(area.CatacombsLevel4),
	)
	// Let's move to a safe area and open the portal in companion mode
	if a.CharacterCfg.Companion.Enabled && a.CharacterCfg.Companion.Leader {
		actions = append(actions,
			a.builder.MoveToCoords(data.Position{
				//22571, 9578
				X: 22571,
				Y: 9578,
			}),
			a.builder.OpenTPIfLeader(),
		)
	}
	actions = append(actions,
		a.builder.MoveToCoords(andarielStartingPosition), // Travel to boss position
		a.char.KillAndariel(),                            // Kill Andariel
	)
	return actions
}
