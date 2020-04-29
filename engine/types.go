package engine

type (
	InvType struct {
		BasePrice     float64  `json:"basePrice"`
		Capacity      float64  `json:"capacity"`
		Description   string   `json:"description"`
		GraphicID     int32    `json:"graphicID"`
		Group         InvGroup `json:"group"`
		IconID        int32    `json:"iconID"`
		MarketGroupID int32    `json:"marketGroupID"`
		Mass          float64  `json:"mass"`
		PortionSize   int32    `json:"portionSize"`
		Published     int32    `json:"published"`
		RaceID        int32    `json:"raceID"`
		SoundID       int32    `json:"soundID"`
		TypeID        int32    `json:"typeID"`
		TypeName      string   `json:"typeName"`
		Volume        float64  `json:"volume"`
	}

	InvGroup struct {
		Anchorable           int32       `json:"anchorable"`
		Anchored             int32       `json:"anchored"`
		Category             InvCategory `json:"category"`
		FittableNonSingleton int32       `json:"fittableNonSingleton"`
		GroupID              int32       `json:"groupID"`
		GroupName            string      `json:"groupName"`
		IconID               int32       `json:"iconID"`
		Published            int32       `json:"published"`
		UseBasePrice         int32       `json:"useBasePrice"`
	}

	InvCategory struct {
		CategoryID   int32  `json:"categoryID"`
		CategoryName string `json:"categoryName"`
		IconID       int32  `json:"iconID"`
		Published    int32  `json:"published"`
	}

	DogmaEffect struct {
		Description                    string         `json:"description"`
		DisallowAutoRepeat             int32          `json:"disallowAutoRepeat"`
		DischargeAttributeID           int32          `json:"dischargeAttributeID"`
		DisplayName                    string         `json:"displayName"`
		Distribution                   int32          `json:"distribution"`
		DurationAttributeID            int32          `json:"durationAttributeID"`
		EffectCategory                 int32          `json:"effectCategory"`
		EffectID                       int32          `json:"effectID"`
		EffectName                     string         `json:"effectName"`
		ElectronicChance               int32          `json:"electronicChance"`
		FalloffAttributeID             int32          `json:"falloffAttributeID"`
		FittingUsageChanceAttributeID  int32          `json:"fittingUsageChanceAttributeID"`
		GUID                           string         `json:"guid"`
		IconID                         int32          `json:"iconID"`
		IsAssistance                   int32          `json:"isAssistance"`
		IsOffensive                    int32          `json:"isOffensive"`
		IsWarpSafe                     int32          `json:"isWarpSafe"`
		ModifierInfoString             string         `json:"-"`
		ModifierInfo                   []ModifierInfo `json:"modifierInfo"`
		NpcActivationChanceAttributeID int32          `json:"npcActivationChanceAttributeID"`
		NpcUsageChanceAttributeID      int32          `json:"npcUsageChanceAttributeID"`
		PostExpression                 int32          `json:"postExpression"`
		PreExpression                  int32          `json:"preExpression"`
		PropulsionChance               int32          `json:"propulsionChance"`
		Published                      int32          `json:"published"`
		RangeAttributeID               int32          `json:"rangeAttributeID"`
		RangeChance                    int32          `json:"rangeChance"`
		SfxName                        string         `json:"sfxName"`
		TrackingSpeedAttributeID       int32          `json:"trackingSpeedAttributeID"`
	}

	ModifierInfo struct {
		Domain               string `json:"domain" yaml:"domain"`
		Function             string `json:"func" yaml:"func"`
		ModifiedAttributeID  int32  `json:"modifiedAttributeID" yaml:"modifiedAttributeID"`
		ModifyingAttributeID int32  `json:"modifyingAttributeID" yaml:"modifyingAttributeID"`
		Operation            string `json:"operation" yaml:"operation"`
		GroupID              int32  `json:"groupID,omitempty" yaml:"groupID"`
		SkillTypeID          int32  `json:"skillTypeID,omitempty" yaml:"skillTypeID"`
	}

	DogmaAttribute struct {
		AttributeID   int32                  `json:"attributeID"`
		AttributeName string                 `json:"attributeName"`
		Category      DogmaAttributeCategory `json:"categoryID"`
		DefaultValue  float64                `json:"defaultValue"`
		Description   string                 `json:"description"`
		DisplayName   string                 `json:"displayName"`
		HighIsGood    int32                  `json:"highIsGood"`
		IconID        int32                  `json:"iconID"`
		Published     int32                  `json:"published"`
		Stackable     int32                  `json:"stackable"`
		UnitID        int32                  `json:"unitID"`
		Unit          DogmaUnit              `json:"unit,omitempty"` // TODO Populate the unit
	}

	DogmaAttributeCategory struct {
		CategoryDescription string `json:"categoryDescription"`
		CategoryID          int32  `json:"categoryID"`
		CategoryName        string `json:"categoryName"`
	}

	DogmaTypeAttribute struct {
		AttributeID int32          `json:"attributeID"`
		TypeID      int32          `json:"typeID"`
		Value       float64        `json:"value"`
		Info        DogmaAttribute `json:"info"`

		Affectors []*DogmaTypeModifier `json:"-"`
		Affected  bool                 `json:"-"`
	}

	DogmaTypeModifier struct {
		ModifyingAttribute *DogmaTypeAttribute
		Operation          string
		SUID               string
	}

	DogmaTypeEffect struct {
		EffectID  int32       `json:"effectID"`
		Effect    DogmaEffect `json:"effect"`
		IsDefault int32       `json:"isDefault"`
		TypeID    int32       `json:"typeID"`
	}

	DogmaUnit struct {
		Description string `json:"description"`
		DisplayName string `json:"displayName"`
		UnitID      int    `json:"unitID;primary_key"`
		UnitName    string `json:"unitName"`
	}
)
