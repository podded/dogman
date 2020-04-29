package sde

import "database/sql"

type (
	InvType struct {
		BasePrice     sql.NullFloat64 `db:"basePrice"`
		Capacity      sql.NullFloat64 `db:"capacity"`
		Description   sql.NullString  `db:"description"`
		GraphicID     sql.NullInt32   `db:"graphicID"`
		GroupID       sql.NullInt32   `db:"groupID"`
		IconID        sql.NullInt32   `db:"iconID"`
		MarketGroupID sql.NullInt32   `db:"marketGroupID"`
		Mass          sql.NullFloat64 `db:"mass"`
		PortionSize   sql.NullInt32   `db:"portionSize"`
		Published     sql.NullInt32   `db:"published"`
		RaceID        sql.NullInt32   `db:"raceID"`
		SoundID       sql.NullInt32   `db:"soundID"`
		TypeID        int32           `db:"typeID"`
		TypeName      sql.NullString  `db:"typeName"`
		Volume        sql.NullFloat64 `db:"volume"`
	}

	InvGroup struct {
		Anchorable           sql.NullInt32  `db:"anchorable"`
		Anchored             sql.NullInt32  `db:"anchored"`
		CategoryID           sql.NullInt32  `db:"categoryID"`
		FittableNonSingleton sql.NullInt32  `db:"fittableNonSingleton"`
		GroupID              int32          `db:"groupID"`
		GroupName            sql.NullString `db:"groupName"`
		IconID               sql.NullInt32  `db:"iconID"`
		Published            sql.NullInt32  `db:"published"`
		UseBasePrice         sql.NullInt32  `db:"useBasePrice"`
	}

	InvCategory struct {
		CategoryID   int32          `db:"categoryID"`
		CategoryName sql.NullString `db:"categoryName"`
		IconID       sql.NullInt32  `db:"iconID"`
		Published    sql.NullInt32  `db:"published"`
	}

	DogmaEffect struct {
		Description                    sql.NullString `db:"description"`
		DisallowAutoRepeat             sql.NullInt32  `db:"disallowAutoRepeat"`
		DischargeAttributeID           sql.NullInt32  `db:"dischargeAttributeID"`
		DisplayName                    sql.NullString `db:"displayName"`
		Distribution                   sql.NullInt32  `db:"distribution"`
		DurationAttributeID            sql.NullInt32  `db:"durationAttributeID"`
		EffectCategory                 sql.NullInt32  `db:"effectCategory"`
		EffectID                       int32          `db:"effectID"`
		EffectName                     sql.NullString `db:"effectName"`
		ElectronicChance               sql.NullInt32  `db:"electronicChance"`
		FalloffAttributeID             sql.NullInt32  `db:"falloffAttributeID"`
		FittingUsageChanceAttributeID  sql.NullInt32  `db:"fittingUsageChanceAttributeID"`
		GUID                           sql.NullString `db:"guid"`
		IconID                         sql.NullInt32  `db:"iconID"`
		IsAssistance                   sql.NullInt32  `db:"isAssistance"`
		IsOffensive                    sql.NullInt32  `db:"isOffensive"`
		IsWarpSafe                     sql.NullInt32  `db:"isWarpSafe"`
		ModifierInfo                   sql.NullString `db:"modifierInfo"`
		NpcActivationChanceAttributeID sql.NullInt32  `db:"npcActivationChanceAttributeID"`
		NpcUsageChanceAttributeID      sql.NullInt32  `db:"npcUsageChanceAttributeID"`
		PostExpression                 sql.NullInt32  `db:"postExpression"`
		PreExpression                  sql.NullInt32  `db:"preExpression"`
		PropulsionChance               sql.NullInt32  `db:"propulsionChance"`
		Published                      sql.NullInt32  `db:"published"`
		RangeAttributeID               sql.NullInt32  `db:"rangeAttributeID"`
		RangeChance                    sql.NullInt32  `db:"rangeChance"`
		SfxName                        sql.NullString `db:"sfxName"`
		TrackingSpeedAttributeID       sql.NullInt32  `db:"trackingSpeedAttributeID"`
	}

	DogmaAttribute struct {
		AttributeID   int32           `db:"attributeID"`
		AttributeName sql.NullString  `db:"attributeName"`
		CategoryID    sql.NullInt32   `db:"categoryID"`
		DefaultValue  sql.NullFloat64 `db:"defaultValue"`
		Description   sql.NullString  `db:"description"`
		DisplayName   sql.NullString  `db:"displayName"`
		HighIsGood    sql.NullInt32   `db:"highIsGood"`
		IconID        sql.NullInt32   `db:"iconID"`
		Published     sql.NullInt32   `db:"published"`
		Stackable     sql.NullInt32   `db:"stackable"`
		UnitID        sql.NullInt32   `db:"unitID"`
	}

	DogmaAttributeCategory struct {
		CategoryDescription sql.NullString `db:"categoryDescription"`
		CategoryID          int32          `db:"categoryID"`
		CategoryName        sql.NullString `db:"categoryName"`
	}

	DogmaTypeAttribute struct {
		AttributeID int32           `db:"attributeID"`
		TypeID      int32           `db:"typeID"`
		ValueFloat  sql.NullFloat64 `db:"valueFloat"`
		ValueInt    sql.NullInt32   `db:"valueInt"`
	}

	DogmaTypeEffect struct {
		EffectID  int32         `db:"effectID"`
		IsDefault sql.NullInt32 `db:"isDefault"`
		TypeID    int32         `db:"typeID"`
	}

	DogmaUnit struct {
		Description sql.NullString `db:"description"`
		DisplayName sql.NullString `db:"displayName"`
		UnitID      int            `db:"unitID;primary_key"`
		UnitName    sql.NullString `db:"unitName"`
	}
)
