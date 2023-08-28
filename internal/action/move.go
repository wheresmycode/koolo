package action

import (
	"fmt"
	"time"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/object"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/koolo/internal/action/step"
	"github.com/hectorgimenez/koolo/internal/helper"
	"github.com/hectorgimenez/koolo/internal/pather"
	"github.com/hectorgimenez/koolo/internal/reader"
	"go.uber.org/zap"
)

func (b Builder) MoveToArea(dst area.Area) Action {
	toFun := func(d data.Data) (data.Position, bool) {
		if d.PlayerUnit.Area == dst {
			b.logger.Debug("Already in area", zap.Any("area", dst))
			return data.Position{}, false
		}

		for _, a := range d.AdjacentLevels {
			if a.Area == dst {
				// To correctly detect the two possible exits from Lut Gholein
				if dst == area.RockyWaste && d.PlayerUnit.Area == area.LutGholein {
					if _, _, found := pather.GetPath(d, data.Position{X: 5004, Y: 5065}); found {
						return data.Position{X: 4989, Y: 5063}, true
					} else {
						return data.Position{X: 5096, Y: 4997}, true
					}
				}

				// This means it's a cave, we don't want to load the map, just find the entrance and interact
				if a.IsEntrance {
					return a.Position, true
				}

				lvl, _ := reader.CachedMapData.GetLevelData(a.Area)
				_, _, objects, _ := reader.CachedMapData.NPCsExitsAndObjects(lvl.Offset, a.Area)

				// Let's try to find any random object to use as a destination point, once we enter the level we will exit this flow
				for _, obj := range objects {
					_, _, found := pather.GetPath(d, obj.Position)
					if found {
						return obj.Position, true
					}
				}

				return a.Position, true
			}
		}

		b.logger.Debug("Destination area not found", zap.Any("area", dst))

		return data.Position{}, false
	}

	return NewChain(func(d data.Data) []Action {
		return []Action{
			b.MoveTo(toFun),
			BuildStatic(func(d data.Data) []step.Step {
				return []step.Step{
					step.InteractEntrance(dst),
				}
			}),
		}
	})
}

func (b Builder) MoveToCoords(to data.Position) *Factory {
	return b.MoveTo(func(d data.Data) (data.Position, bool) {
		return to, true
	})
}

func (b Builder) MoveTo(toFunc func(d data.Data) (data.Position, bool)) *Factory {
	pickupBeforeMoving := false
	openedDoors := make(map[object.Name]data.Position)
	previousIterationPosition := data.Position{}
	var currentStep step.Step

	return NewFactory(func(d data.Data) Action {
		to, found := toFunc(d)
		if !found {
			return nil
		}

		// To stop the movement, not very accurate
		if pather.DistanceFromMe(d, to) < 5 {
			return nil
		}

		// Let's go pickup more pots if we have less than 2 (only during leveling)
		_, isLevelingChar := b.ch.(LevelingCharacter)
		if isLevelingChar {
			_, healingPotsFound := d.Items.Belt.GetFirstPotion(data.HealingPotion)
			_, manaPotsFound := d.Items.Belt.GetFirstPotion(data.ManaPotion)
			if (!healingPotsFound || !manaPotsFound) && d.PlayerUnit.TotalGold() > 1000 {
				return NewChain(func(d data.Data) []Action {
					return b.InRunReturnTownRoutine()
				})
			}
		}

		if helper.CanTeleport(d) {
			// If we can teleport, and we're not on leveling sequence, just return the normal MoveTo step and stop here
			if !isLevelingChar {
				return BuildStatic(func(d data.Data) []step.Step {
					return []step.Step{step.MoveTo(to)}
				})
			}
			// But if we are leveling and have enough money (to buy pots), let's teleport. We add the timeout
			// to re-trigger this action, so we can get back to town to buy pots in case of empty belt
			if d.PlayerUnit.TotalGold() > 10000 {
				return BuildStatic(func(d data.Data) []step.Step {
					return []step.Step{step.MoveTo(to, step.WithTimeout(5*time.Second))}
				})
			}
		}

		// Check if there is a door blocking our path
		for _, o := range d.Objects {
			if o.IsDoor() && pather.DistanceFromMe(d, o.Position) < 10 && openedDoors[o.Name] != o.Position {
				if o.Selectable {
					return BuildStatic(func(d data.Data) []step.Step {
						b.logger.Info("Door detected and teleport is not available, trying to open it...")
						openedDoors[o.Name] = o.Position
						return []step.Step{step.InteractObject(o.Name, func(d data.Data) bool {
							for _, obj := range d.Objects {
								if obj.Name == o.Name && obj.Position == o.Position && !obj.Selectable {
									return true
								}
							}
							return false
						})}
					}, CanBeSkipped())
				}
			}
		}

		// Detect if there are monsters close to the player
		closestMonster := data.Monster{}
		closestMonsterDistance := 9999999
		targetedNormalEnemies := make([]data.Monster, 0)
		targetedElites := make([]data.Monster, 0)
		minDistance := 6
		minDistanceForElites := 20                                       // This will make the character to kill elites even if they are far away, ONLY during leveling
		stuck := pather.DistanceFromMe(d, previousIterationPosition) < 5 // Detect if character was not able to move from last iteration
		for _, m := range d.Monsters.Enemies() {
			// Skip if monster is already dead
			if m.Stats[stat.Life] <= 0 {
				continue
			}

			dist := pather.DistanceFromMe(d, m.Position)
			appended := false
			if m.IsElite() && dist <= minDistanceForElites {
				targetedElites = append(targetedElites, m)
				appended = true
			}

			if dist <= minDistance {
				targetedNormalEnemies = append(targetedNormalEnemies, m)
				appended = true
			}

			if appended {
				if dist < closestMonsterDistance {
					closestMonsterDistance = dist
					closestMonster = m
				}
			}
		}

		if len(targetedNormalEnemies) > 3 || len(targetedElites) > 0 {
			if stuck {
				b.logger.Info("Character stuck and monsters detected, trying to kill monsters around")
			} else {
				b.logger.Info(fmt.Sprintf("At least %d monsters detected close to the character, targeting closest one: %d", len(targetedNormalEnemies)+len(targetedElites), closestMonster.Name))
			}

			pickupBeforeMoving = true

			return b.ch.KillMonsterSequence(func(d data.Data) (data.UnitID, bool) {
				return closestMonster.UnitID, true
			}, nil)
		}

		if pickupBeforeMoving {
			pickupBeforeMoving = false
			return b.ItemPickup(false, 30)
		}

		// Continue moving
		return BuildStatic(func(d data.Data) []step.Step {
			previousIterationPosition = d.PlayerUnit.Position
			if currentStep == nil {
				currentStep = step.MoveTo(
					to,
					step.ClosestWalkable(),
					step.WithTimeout(time.Millisecond*1000),
				)
			} else {
				currentStep.Reset()
			}

			return []step.Step{currentStep}
		})
	})
}