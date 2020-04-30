package engine

import (
	"errors"
	"github.com/podded/dogman/sde"
	"gopkg.in/yaml.v2"
)

type (
	Ship struct {
		Type           InvType
		TypeAttributes []*DogmaTypeAttribute
		TypeEffects    []*DogmaTypeEffect
	}
)

func NewShip(typeID int32, sd *sde.Data) (*Ship, error) {

	ship := &Ship{}

	tid, found := sd.Types[typeID]
	if !found {
		return nil, errors.New("ship typeID not found")
	}

	grp, found := sd.Groups[tid.GroupID.Int32]
	if !found {
		return nil, errors.New("ship groupID not found")
	}

	cat, found := sd.Categories[grp.CategoryID.Int32]
	if !found {
		return nil, errors.New("ship typeID not found")
	}

	tp := InvType{
		BasePrice:   tid.BasePrice.Float64,
		Capacity:    tid.Capacity.Float64,
		Description: tid.Description.String,
		GraphicID:   tid.GraphicID.Int32,
		Group: InvGroup{
			Anchorable: grp.Anchorable.Int32,
			Anchored:   grp.Anchored.Int32,
			Category: InvCategory{
				CategoryID:   cat.CategoryID,
				CategoryName: cat.CategoryName.String,
				IconID:       cat.IconID.Int32,
				Published:    cat.Published.Int32,
			},
			FittableNonSingleton: grp.FittableNonSingleton.Int32,
			GroupID:              grp.GroupID,
			GroupName:            grp.GroupName.String,
			IconID:               grp.IconID.Int32,
			Published:            grp.Published.Int32,
			UseBasePrice:         grp.UseBasePrice.Int32,
		},
		IconID:        tid.IconID.Int32,
		MarketGroupID: tid.MarketGroupID.Int32,
		Mass:          tid.Mass.Float64,
		PortionSize:   tid.PortionSize.Int32,
		Published:     tid.Published.Int32,
		RaceID:        tid.RaceID.Int32,
		SoundID:       tid.SoundID.Int32,
		TypeID:        tid.TypeID,
		TypeName:      tid.TypeName.String,
		Volume:        tid.Volume.Float64,
	}

	ship.Type = tp

	// Now get all of the attributes associated with the type

	ta := sd.TypeAttributes[ship.Type.TypeID]
	tas := make([]*DogmaTypeAttribute, len(ta))
	for i, t := range ta {
		sta := sd.Attributes[t.AttributeID]

		val := float64(0)
		if t.ValueFloat.Valid && !t.ValueInt.Valid {
			val = t.ValueFloat.Float64
		} else if !t.ValueFloat.Valid && t.ValueInt.Valid {
			val = float64(t.ValueInt.Int32)
		} else {
			// TBH not sure what to do here, but at time of writing all fields are valueFloat, no valueInts so go with that
			val = t.ValueFloat.Float64
		}

		ac := sd.AttributeCategories[sta.CategoryID.Int32]

		dta := DogmaTypeAttribute{
			AttributeID: t.AttributeID,
			TypeID:      ship.Type.TypeID,
			Value:       val,
			Info: DogmaAttribute{
				AttributeID:   t.AttributeID,
				AttributeName: sta.AttributeName.String,
				Category: DogmaAttributeCategory{
					CategoryDescription: ac.CategoryName.String,
					CategoryID:          ac.CategoryID,
					CategoryName:        ac.CategoryName.String,
				},
				DefaultValue: sta.DefaultValue.Float64,
				Description:  sta.Description.String,
				DisplayName:  sta.DisplayName.String,
				HighIsGood:   sta.HighIsGood.Int32,
				IconID:       sta.IconID.Int32,
				Published:    sta.Published.Int32,
				Stackable:    sta.Stackable.Int32,
				UnitID:       sta.UnitID.Int32,
			},
			Affectors: nil,
			Affected:  false,
		}

		tas[i] = &dta
	}

	ship.TypeAttributes = tas

	// Now the type effects

	te := sd.TypeEffects[ship.Type.TypeID]
	tefs := make([]*DogmaTypeEffect, len(te))

	for i, t := range te {

		des := sd.Effects[t.EffectID]

		tef := DogmaTypeEffect{
			EffectID: t.EffectID,
			Effect: DogmaEffect{
				Description:                    des.Description.String,
				DisallowAutoRepeat:             des.DisallowAutoRepeat.Int32,
				DischargeAttributeID:           des.DischargeAttributeID.Int32,
				DisplayName:                    des.DisplayName.String,
				Distribution:                   des.Distribution.Int32,
				DurationAttributeID:            des.DurationAttributeID.Int32,
				EffectCategory:                 des.EffectCategory.Int32,
				EffectID:                       des.EffectID,
				EffectName:                     des.EffectName.String,
				ElectronicChance:               des.ElectronicChance.Int32,
				FalloffAttributeID:             des.FalloffAttributeID.Int32,
				FittingUsageChanceAttributeID:  des.FittingUsageChanceAttributeID.Int32,
				GUID:                           des.GUID.String,
				IconID:                         des.IconID.Int32,
				IsAssistance:                   des.IsAssistance.Int32,
				IsOffensive:                    des.IsOffensive.Int32,
				IsWarpSafe:                     des.IsWarpSafe.Int32,
				ModifierInfoString:             des.ModifierInfo.String,
				NpcActivationChanceAttributeID: des.NpcActivationChanceAttributeID.Int32,
				NpcUsageChanceAttributeID:      des.NpcUsageChanceAttributeID.Int32,
				PostExpression:                 des.PostExpression.Int32,
				PreExpression:                  des.PreExpression.Int32,
				PropulsionChance:               des.PropulsionChance.Int32,
				Published:                      des.Published.Int32,
				RangeAttributeID:               des.RangeAttributeID.Int32,
				RangeChance:                    des.RangeChance.Int32,
				SfxName:                        des.SfxName.String,
				TrackingSpeedAttributeID:       des.TrackingSpeedAttributeID.Int32,
			},
			IsDefault: t.IsDefault.Int32,
			TypeID:    t.TypeID,
		}

		// Parse the ModifierInfoString into the ModifierInfo field
		if len(tef.Effect.ModifierInfoString) > 0 {
			var mi []ModifierInfo
			err := yaml.Unmarshal([]byte(tef.Effect.ModifierInfoString), &mi)
			if err == nil {
				tef.Effect.ModifierInfo = mi
			}
		}

		tefs[i] = &tef
	}
	ship.TypeEffects = tefs

	return ship, nil
}
