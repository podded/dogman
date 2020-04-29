package sde

import (
	"github.com/jmoiron/sqlx"
)

func getSDETypes(db *sqlx.DB) (map[int32]InvType, error) {

	tps := make(map[int32]InvType)

	var t []InvType
	err := db.Select(&t, "SELECT * FROM invTypes")

	for _, tp := range t {
		tps[tp.TypeID] = tp
	}

	return tps, err
}

func getSDEGroups(db *sqlx.DB) (map[int32]InvGroup, error) {

	gps := make(map[int32]InvGroup)

	var g []InvGroup
	err := db.Select(&g, "SELECT * FROM invGroups")

	for _, gp := range g {
		gps[gp.GroupID] = gp
	}

	return gps, err
}

func getSDECategories(db *sqlx.DB) (map[int32]InvCategory, error) {

	cts := make(map[int32]InvCategory)

	var c []InvCategory
	err := db.Select(&c, "SELECT * FROM invCategories")

	for _, ct := range c {
		cts[ct.CategoryID] = ct
	}

	return cts, err
}

func getSDEEffects(db *sqlx.DB) (map[int32]DogmaEffect, error) {

	efs := make(map[int32]DogmaEffect)

	var e []DogmaEffect
	err := db.Select(&e, "SELECT * FROM dgmEffects")

	for _, ef := range e {
		efs[ef.EffectID] = ef
	}

	return efs, err
}

func getSDEAttributes(db *sqlx.DB) (map[int32]DogmaAttribute, error) {

	ats := make(map[int32]DogmaAttribute)

	var a []DogmaAttribute
	err := db.Select(&a, "SELECT * FROM dgmAttributeTypes")

	for _, at := range a {
		ats[at.AttributeID] = at
	}

	return ats, err
}

func getSDEAttributeCategory(db *sqlx.DB) (map[int32]DogmaAttributeCategory, error) {

	acs := make(map[int32]DogmaAttributeCategory)

	var a []DogmaAttributeCategory
	err := db.Select(&a, "SELECT * FROM dgmAttributeCategories")

	for _, ac := range a {
		acs[ac.CategoryID] = ac
	}

	return acs, err
}

func getSDETypeAttributes(db *sqlx.DB) (map[int32][]DogmaTypeAttribute, error) {

	tas := make(map[int32][]DogmaTypeAttribute)

	var t []DogmaTypeAttribute
	err := db.Select(&t, "SELECT * FROM dgmTypeAttributes")

	for _, ta := range t {
		if tae, ok := tas[ta.TypeID]; ok {
			tas[ta.TypeID] = append(tae, ta)
		} else {
			tas[ta.TypeID] = []DogmaTypeAttribute{ta}
		}
	}
	
	return tas, err
}

func getSDETypeEffects(db *sqlx.DB) (map[int32][]DogmaTypeEffect, error) {
	
	tes := make(map[int32][]DogmaTypeEffect)

	var t []DogmaTypeEffect
	err := db.Select(&t, "SELECT * FROM dgmTypeEffects")

	for _, te := range t {
		if tae, ok := tes[te.TypeID]; ok {
			tes[te.TypeID] = append(tae, te)
		} else {
			tes[te.TypeID] = []DogmaTypeEffect{te}
		}
	}

	return tes, err
}