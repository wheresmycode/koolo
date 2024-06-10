package run

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/game"
)

type Bishibosh struct {
	baseRun
}

func (a Bishibosh) Name() string {
	return string(config.BishiboshRun)
}

func (a Bishibosh) BuildActions() (actions []action.Action) {
	a.logger.Debug("Building actions for Bishibosh")
	actions = append(actions,
		a.builder.WayPoint(area.ColdPlains), // Moving to starting point (Cold Plains)
		a.builder.MoveTo(func(d game.Data) (data.Position, bool) {
			m, found := d.NPCs.FindOne(734) // 734 is the superunique preset id for bishibosh
			return m.Positions[0], found
		}), // Travel to boss position

		a.char.KillMonsterSequence(func(d game.Data) (data.UnitID, bool) {
			if m, found := d.Monsters.FindOne(npc.FallenShaman, data.MonsterTypeSuperUnique); found {
				return m.UnitID, true
			}

			return 0, false
		}, nil),
		a.builder.ItemPickup(false, 35),
	)

	return
}
