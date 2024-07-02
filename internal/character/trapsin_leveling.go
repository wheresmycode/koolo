package character

import (
	"sort"
	"time"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/d2go/pkg/data/skill"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/action/step"
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/helper"
	"github.com/hectorgimenez/koolo/internal/pather"
)

const (
	assnLvlMaxAttacksLoop = 5
	assnLvlMinDistance    = 2
	assnLvlMaxDistance    = 35
)

type TrapsinLeveling struct {
	BaseCharacter
}

func (s TrapsinLeveling) ShouldResetSkills(d game.Data) bool {
	lvl, _ := d.PlayerUnit.FindStat(stat.Level, 0)
	if lvl.Value >= 35 && d.PlayerUnit.Skills[skill.WakeOfFire].Level > 15 {
		return true
	}

	return false
}

func (s TrapsinLeveling) SkillsToBind(d game.Data) (skill.ID, []skill.ID) {
	primarySkill := skill.AttackSkill
	skillBindings := []skill.ID{
		skill.BladeShield,
		skill.ShadowMaster,
		skill.TomeOfTownPortal,
	}
	if _, found := d.KeyBindings.KeyBindingForSkill(skill.BurstOfSpeed); !found {
		if _, found := d.KeyBindings.KeyBindingForSkill(skill.Fade); !found {
			skillBindings = append(skillBindings, skill.BurstOfSpeed) //default to Burst of Speed if no keybinding for it or fade was found
		}
	}

	if d.PlayerUnit.Skills[skill.LightningSentry].Level > 0 {
		skillBindings = append(skillBindings, skill.LightningSentry)
	}
	if d.PlayerUnit.Skills[skill.DeathSentry].Level > 0 {
		skillBindings = append(skillBindings, skill.DeathSentry)
	} else if d.PlayerUnit.Skills[skill.WakeOfFire].Level > 0 {
		skillBindings = append(skillBindings, skill.WakeOfFire)
	}
	if d.PlayerUnit.Skills[skill.FireBlast].Level > 0 {
		primarySkill = skill.FireBlast
	}
	lvl, _ := d.PlayerUnit.FindStat(stat.Level, 0)
	if lvl.Value < 15 {
		skillBindings = append(skillBindings, skill.AttackSkill)
	}
	return primarySkill, skillBindings
}

func (s TrapsinLeveling) StatPoints(d game.Data) map[stat.ID]int {
	lvl, _ := d.PlayerUnit.FindStat(stat.Level, 0)
	if lvl.Value < 9 {
		return map[stat.ID]int{
			stat.Strength: 20,
			stat.Vitality: 9999,
		}
	}

	if lvl.Value < 15 {
		return map[stat.ID]int{
			stat.Energy:   30,
			stat.Strength: 30,
			stat.Vitality: 9999,
		}
	}

	return map[stat.ID]int{
		stat.Energy:   40,
		stat.Strength: 60,
		stat.Vitality: 9999,
	}
}

func (s TrapsinLeveling) SkillPoints(d game.Data) []skill.ID {
	lvl, _ := d.PlayerUnit.FindStat(stat.Level, 0)
	if lvl.Value < 35 {
		return []skill.ID{
			skill.FireBlast,
			skill.FireBlast,
			skill.FireBlast,
			skill.FireBlast,
			skill.FireBlast,
			skill.ClawMastery,
			skill.BurstOfSpeed,
			skill.BurstOfSpeed,
			skill.BurstOfSpeed,
			skill.BurstOfSpeed,
			skill.BurstOfSpeed,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WeaponBlock,
			skill.ShadowWarrior,
			skill.ShadowMaster,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.WakeOfFire,
			skill.FireBlast,
			skill.FireBlast,
			skill.FireBlast,
			skill.FireBlast,
		}
	}

	return []skill.ID{
		skill.ClawMastery,
		skill.BurstOfSpeed,
		skill.WeaponBlock,
		skill.PsychicHammer,
		skill.CloakOfShadows,
		skill.ShadowWarrior,
		skill.ShadowMaster,
		skill.Fade,
		skill.FireBlast,
		skill.ShockWeb,
		skill.ChargedBoltSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.DeathSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.LightningSentry,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ShockWeb,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.ChargedBoltSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.DeathSentry,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
		skill.FireBlast,
	}
}

func (s TrapsinLeveling) BuffSkills(d game.Data) []skill.ID {
	skillsList := make([]skill.ID, 0)
	if _, found := d.KeyBindings.KeyBindingForSkill(skill.BladeShield); found {
		skillsList = append(skillsList, skill.BladeShield)
	}

	armors := []skill.ID{skill.BurstOfSpeed, skill.Fade}
	for _, armor := range armors {
		if _, found := d.KeyBindings.KeyBindingForSkill(armor); found {
			skillsList = append(skillsList, armor)
			return skillsList
		}
	}

	return skillsList
}

func (s TrapsinLeveling) PreCTABuffSkills(d game.Data) []skill.ID {
	armor := skill.ShadowWarrior
	armors := []skill.ID{skill.ShadowWarrior, skill.ShadowMaster}
	hasShadow := false
	for _, arm := range armors {
		if _, found := d.KeyBindings.KeyBindingForSkill(arm); found {
			armor = arm
			hasShadow = true
		}
	}

	if hasShadow {
		return []skill.ID{armor}
	}

	return []skill.ID{}
}

func (s TrapsinLeveling) KillMonsterSequence(
	monsterSelector func(d game.Data) (data.UnitID, bool),
	skipOnImmunities []stat.Resist,
	opts ...step.AttackOption,
) action.Action {
	return action.NewStepChain(func(d game.Data) (steps []step.Step) {
		id, found := monsterSelector(d)
		if !found {
			return []step.Step{}
		}
		if !s.preBattleChecks(d, id, skipOnImmunities) {
			return []step.Step{}
		}

		opts := []step.AttackOption{step.Distance(assnLvlMinDistance, assnLvlMaxDistance)}

		helper.Sleep(100)

		// if missing fire blast or below lvl 15 and less than 25% mana just do melee attack
		lvl, _ := d.PlayerUnit.FindStat(stat.Level, 0)
		if (lvl.Value < 15 && d.PlayerUnit.MPPercent() < 25) || d.PlayerUnit.Skills[skill.FireBlast].Level == 0 {
			steps = append(steps, step.SecondaryAttack(skill.AttackSkill, id, 5, step.Distance(1, 3)))
		} else if _, found := d.KeyBindings.KeyBindingForSkill(skill.DeathSentry); found {
			steps = append(steps,
				step.SecondaryAttack(skill.LightningSentry, id, 3, opts...),
				step.SecondaryAttack(skill.DeathSentry, id, 2, opts...),
				step.PrimaryAttack(id, 5, true, step.Distance(assnLvlMinDistance, assnLvlMaxDistance)),
			)
		} else if _, found := d.KeyBindings.KeyBindingForSkill(skill.LightningSentry); found {
			steps = append(steps,
				step.SecondaryAttack(skill.LightningSentry, id, 5, opts...),
				step.PrimaryAttack(id, 5, true, step.Distance(assnLvlMinDistance, assnLvlMaxDistance)),
			)
		} else if _, found := d.KeyBindings.KeyBindingForSkill(skill.WakeOfFire); found {
			steps = append(steps,
				step.SecondaryAttack(skill.WakeOfFire, id, 3, opts...),
				step.PrimaryAttack(id, 5, true, step.Distance(assnLvlMinDistance, assnLvlMaxDistance)),
			)
		} else {
			steps = append(steps,
				step.PrimaryAttack(id, 3, true, step.Distance(assnLvlMinDistance, assnLvlMaxDistance)),
			)
		}

		return
	}, action.RepeatUntilNoSteps())
}

func (s TrapsinLeveling) killMonster(npc npc.ID, t data.MonsterType) action.Action {
	return s.KillMonsterSequence(func(d game.Data) (data.UnitID, bool) {
		m, found := d.Monsters.FindOne(npc, t)
		if !found {
			return 0, false
		}

		return m.UnitID, true
	}, nil)
}

func (s TrapsinLeveling) KillCountess() action.Action {
	return s.killMonster(npc.DarkStalker, data.MonsterTypeSuperUnique)
}

func (s TrapsinLeveling) KillAndariel() action.Action {
	return s.killMonster(npc.Andariel, data.MonsterTypeNone)
}

func (s TrapsinLeveling) KillSummoner() action.Action {
	return s.killMonster(npc.Summoner, data.MonsterTypeNone)
}

func (s TrapsinLeveling) KillDuriel() action.Action {
	return s.killMonster(npc.Duriel, data.MonsterTypeNone)
}

func (s TrapsinLeveling) KillPindle(_ []stat.Resist) action.Action {
	return s.killMonster(npc.DefiledWarrior, data.MonsterTypeSuperUnique)
}

func (s TrapsinLeveling) KillMephisto() action.Action {
	return s.killMonster(npc.Mephisto, data.MonsterTypeNone)
}

func (s TrapsinLeveling) KillNihlathak() action.Action {
	return s.killMonster(npc.Nihlathak, data.MonsterTypeSuperUnique)
}

func (s TrapsinLeveling) KillDiablo() action.Action {
	timeout := time.Second * 20
	startTime := time.Time{}
	diabloFound := false
	return action.NewChain(func(d game.Data) []action.Action {
		if startTime.IsZero() {
			startTime = time.Now()
		}

		if time.Since(startTime) > timeout && !diabloFound {
			s.logger.Error("Diablo was not found, timeout reached")
			return nil
		}

		diablo, found := d.Monsters.FindOne(npc.Diablo, data.MonsterTypeNone)
		if !found || diablo.Stats[stat.Life] <= 0 {
			// Already dead
			if diabloFound {
				return nil
			}

			// Keep waiting...
			return []action.Action{action.NewStepChain(func(d game.Data) []step.Step {
				return []step.Step{step.Wait(time.Millisecond * 100)}
			})}
		}

		diabloFound = true
		s.logger.Info("Diablo detected, attacking")

		return []action.Action{
			s.killMonster(npc.Diablo, data.MonsterTypeNone),
			s.killMonster(npc.Diablo, data.MonsterTypeNone),
			s.killMonster(npc.Diablo, data.MonsterTypeNone),
		}
	}, action.RepeatUntilNoSteps())
}

func (s TrapsinLeveling) KillIzual() action.Action {
	return s.killMonster(npc.Izual, data.MonsterTypeNone)
}

func (s TrapsinLeveling) KillCouncil() action.Action {
	return s.KillMonsterSequence(func(d game.Data) (data.UnitID, bool) {
		var councilMembers []data.Monster
		for _, m := range d.Monsters {
			if m.Name == npc.CouncilMember || m.Name == npc.CouncilMember2 || m.Name == npc.CouncilMember3 {
				councilMembers = append(councilMembers, m)
			}
		}

		// Order council members by distance
		sort.Slice(councilMembers, func(i, j int) bool {
			distanceI := pather.DistanceFromMe(d, councilMembers[i].Position)
			distanceJ := pather.DistanceFromMe(d, councilMembers[j].Position)

			return distanceI < distanceJ
		})

		if len(councilMembers) > 0 {
			return councilMembers[0].UnitID, true
		}

		return 0, false
	}, nil)
}

func (s TrapsinLeveling) KillBaal() action.Action {
	return s.killMonster(npc.BaalCrab, data.MonsterTypeNone)
}

func (p TrapsinLeveling) KillAncients() action.Action {
	return action.NewChain(func(d game.Data) (actions []action.Action) {
		for _, m := range d.Monsters.Enemies(data.MonsterEliteFilter()) {
			actions = append(actions,
				p.killMonster(m.Name, data.MonsterTypeSuperUnique),
			)
		}
		return actions
	})
}
