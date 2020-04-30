package sde

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Data struct {

	// Types maps a type id to the Data representation of the type
	Types map[int32]InvType
	// Groups maps a group id to the Data representation of the group
	Groups map[int32]InvGroup
	// Categories maps a category id to the Data representation of the category
	Categories map[int32]InvCategory

	// TypeEffects maps an effect id to the Data data of the effect
	Effects map[int32]DogmaEffect
	// TypeAttributes maps an attribute id to the Data data of the attribute
	Attributes map[int32]DogmaAttribute
	// AttributeCategories maps an attribute id to the Data data of the attribute category
	AttributeCategories map[int32]DogmaAttributeCategory

	// TypeEffects maps a type id to the list of effects that apply to it
	TypeEffects map[int32][]DogmaTypeEffect
	// TypeAttributes maps a type id to the list of attributes that apply to it
	TypeAttributes map[int32][]DogmaTypeAttribute
}

func New(db *sqlx.DB) (*Data, error) {
	sde := Data{}

	tps, err := getSDETypes(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde types: %v", err)
	}
	sde.Types = tps

	gps, err := getSDEGroups(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde groups: %v", err)
	}
	sde.Groups = gps

	cts, err := getSDECategories(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde categories: %v", err)
	}
	sde.Categories = cts

	efs, err := getSDEEffects(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde effects: %v", err)
	}
	sde.Effects = efs

	ats, err := getSDEAttributes(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde attributes: %v", err)
	}
	sde.Attributes = ats

	atc, err := getSDEAttributeCategory(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde attribute categories: %v", err)
	}
	sde.AttributeCategories = atc

	tas, err := getSDETypeAttributes(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde type attributes: %v", err)
	}
	sde.TypeAttributes = tas

	tes, err := getSDETypeEffects(db)
	if err != nil {
		return nil, fmt.Errorf("error getting sde type effects: %v", err)
	}
	sde.TypeEffects = tes

	return &sde, nil
}
