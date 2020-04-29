package engine

import (
	"errors"
	"fmt"
	"github.com/emicklei/dot"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/podded/dogman/sde"
	"log"
	"strconv"
	"time"
)

type (
	Solo struct {
		db  *sqlx.DB
		sde *sde.Data

		ship   *Ship
		skills map[int32]*Skill
	}
)

const (
	minSkillLevel = 0
	maxSkillLevel = 5
)

func NewSolo(dbURI string) (*Solo, error) {

	solo := &Solo{}

	conn, err := sqlx.Connect("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	solo.db = conn

	sde, err := sde.New(conn)
	if err != nil {
		return nil, err
	}

	solo.sde = sde

	solo.ship = &Ship{}

	return solo, nil
}

func (solo *Solo) Close() error {
	return solo.db.Close()
}

func (solo *Solo) Reset() {

	solo.ship = &Ship{}
}

func (solo *Solo) SetShipID(shipTypeID int32) error {

	shp, err := NewShip(shipTypeID, solo.sde)
	if err != nil {
		return err
	}
	solo.ship = shp
	return nil
}

func (solo *Solo) InjectAllSkills() error {
	// Skill category is 16

	start := time.Now()

	solo.skills = make(map[int32]*Skill)

	grps := map[int32]struct{}{}

	// Get all groups with category 16
	for _, grp := range solo.sde.Groups {
		if grp.CategoryID.Int32 != 16 {
			continue
		}
		grps[grp.GroupID] = struct{}{}
	}

	for _, t := range solo.sde.Types {
		if _, match := grps[t.GroupID.Int32]; match {
			skill, err := NewSkill(t.TypeID, solo.sde)
			if err != nil {
				log.Println("Error creating skill")
				log.Printf("\t%v\n", err)
				continue
			}
			solo.skills[skill.Type.TypeID] = skill
		}
	}

	log.Printf("Loaded %d skills! This took %v\n", len(solo.skills), time.Now().Sub(start))

	return nil

}

func (solo *Solo) SetAllSkillsLevel(level int) error {
	if level < minSkillLevel || level > maxSkillLevel {
		return errors.New(fmt.Sprintf("invalid skill level. must satisfy %d < %d <= %d", minSkillLevel, level, maxSkillLevel))
	}

	for _, skill := range solo.skills {
		skill.SetLevel(5)
	}

	return nil
}

func (solo *Solo) SetSkillLevel(skillID int32, level int) error {
	if level < minSkillLevel || level > maxSkillLevel {
		return errors.New(fmt.Sprintf("invalid skill level. must satisfy %d < %d <= %d", minSkillLevel, level, maxSkillLevel))
	}

	skill, ok := solo.skills[skillID]
	if !ok {
		return errors.New(fmt.Sprintf("skillid %d has not been injected yet", skillID))
	}

	skill.SetLevel(level)

	return nil
}

func (solo *Solo) BuildAffectorTree() error {

	// First add the effects of the ship hull
	if solo.ship.TypeEffects != nil {
		for _, eff := range solo.ship.TypeEffects {
			if eff == nil {
				continue
			}
			for _, mi := range eff.Effect.ModifierInfo {
				switch mi.Domain {
				case "shipID":
					switch mi.Function {
					default:
						log.Printf("function '%s' is not implemented for the '%s' domain skipping", mi.Function, mi.Domain)
					}
				default:
					log.Printf("modifier info domain '%s' not implemented, skipping", mi.Domain)
					continue
				}
			}
		}

	}

	// Now add the effects of the skills implemented
	for _, sk := range solo.skills {
		if sk.GetLevel() == 0 {
			continue
		}

		if sk.TypeEffects == nil || len(sk.TypeEffects) == 0 {
			continue
		}

		for _, te := range sk.TypeEffects {
			if te == nil {
				continue
			}

			for _, mi := range te.Effect.ModifierInfo {
				switch mi.Domain {

				case "ItemID":
					// ItemID is a way of referring to itself as far as I can tell
					switch mi.Function {
					case "ItemModifier":
						var modifying *DogmaTypeAttribute
						var modified *DogmaTypeAttribute
						for _, ta := range sk.TypeAttributes {
							if ta.AttributeID == mi.ModifyingAttributeID {
								modifying = ta
								break
							}
						}
						for _, ta := range sk.TypeAttributes {
							if ta.AttributeID == mi.ModifiedAttributeID {
								modified = ta
								break
							}
						}
						// If either of these is null, we have a problem
						if modified == nil || modifying == nil {
							return errors.New(fmt.Sprintf("nil type attributes for effect '%d' in skill '%d'", te.Effect.EffectID, sk.Type.TypeID))
						}
						// Both attrs are ok, so proceed

						// Add the affector to the modified attribute
						mod := DogmaTypeModifier{
							ModifyingAttribute: modifying,
							Operation:          mi.Operation,
							SUID:               fmt.Sprintf("%d-%d", sk.Type.TypeID, te.EffectID),
						}

						modified.Affectors = append(modified.Affectors, &mod)
					}

				case "shipID":
					switch mi.Function {
					case "ItemModifier":
						var modifying *DogmaTypeAttribute
						var modified *DogmaTypeAttribute

						for _, ta := range solo.ship.TypeAttributes {
							if ta.AttributeID == mi.ModifiedAttributeID {
								modified = ta
								break
							}
						}
						if modified == nil {
							continue
						}

						for _, ta := range sk.TypeAttributes {
							if ta.AttributeID == mi.ModifyingAttributeID {
								modifying = ta
								break
							}
						}
						// If mpdifying is nil at this point we have a problem
						if modifying == nil {
							return errors.New(fmt.Sprintf("nil modifying attribute for effect '%d' in skill '%d'", te.Effect.EffectID, sk.Type.TypeID))
						}

						log.Printf("Adding a modifier to attr %d of type %d. modifiying attr is %d of type %d", modified.AttributeID, modified.TypeID, modifying.AttributeID, modifying.TypeID)

						// Add the affector to the modified attribute
						mod := DogmaTypeModifier{
							ModifyingAttribute: modifying,
							Operation:          mi.Operation,
							SUID:               fmt.Sprintf("%d-%d", sk.Type.TypeID, te.EffectID),
						}

						modified.Affectors = append(modified.Affectors, &mod)
					default:
						log.Printf("function '%s' is not implemented for the '%s' domain skipping", mi.Function, mi.Domain)
					}
				default:
					log.Printf("modifier info domain '%s' not implemented, skipping", mi.Domain)
					continue
				}
			}
		}
	}

	if solo.ship.TypeEffects != nil {
		for _, eff := range solo.ship.TypeEffects {
			if eff == nil {
				continue
			}
			for _, mi := range eff.Effect.ModifierInfo {
				switch mi.Domain {
				case "shipID":
					switch mi.Function {
					default:
						log.Printf("function '%s' is not implemented for the '%s' domain skipping", mi.Function, mi.Domain)
					}
				default:
					log.Printf("modifier info domain '%s' not implemented, skipping", mi.Domain)
					continue
				}
			}
		}

	}
	return nil
}

func (solo *Solo) PrintAffectorTree() error {

	// For the calculations we are only care about the attributes of the ship

	g := dot.NewGraph(dot.Directed)

	// Create subgraphs for each of the typeIDs that are involved
	// Key is 'typeID'
	subgraphs := make(map[int32]*dot.Graph)
	// Nodes will hold a list of all nodes created so that the graph can converge if required
	// Key is 'typeID-attributeID'
	nodes := make(map[string]*dot.Node)

	// Start off by adding the one cluster that we know will exist
	shp := g.Subgraph(fmt.Sprintf("Ship - %s", solo.ship.Type.TypeName), dot.ClusterOption{})
	subgraphs[solo.ship.Type.TypeID] = shp

	for _, ta := range solo.ship.TypeAttributes {
		// The root level typeIDs are all for the ship
		id := fmt.Sprintf("%d-%d", ta.TypeID, ta.AttributeID)
		nodeLabelHeader := ""
		if ta.Info.AttributeName != "" {
			nodeLabelHeader = ta.Info.AttributeName
		} else {
			nodeLabelHeader = strconv.FormatInt(int64(ta.AttributeID), 10)
		}
		atNode := shp.Node(id).Attr("label", fmt.Sprintf("%s\n%f", nodeLabelHeader, ta.Value))
		nodes[id] = &atNode

		// Now handle any modifiers that need to be applied to the node
		if len(ta.Affectors) > 0 {
			printTraverseAffectorList(ta, ta, g, subgraphs, nodes, solo.sde)
		}
	}

	fmt.Println(g.String())

	return nil
}

func printTraverseAffectorList(dta *DogmaTypeAttribute, root *DogmaTypeAttribute, rootGraph *dot.Graph, subgraphs map[int32]*dot.Graph, nodes map[string]*dot.Node, sd *sde.Data) error {
	// For each affector, make sure both modifying and modified node exist,
	// make sure that the modifying node is marked as affected, if not build that before making the edge
	for _, tm := range dta.Affectors {
		typeID := tm.ModifyingAttribute.TypeID
		log.Printf("Checking for existing subgraph of type %d", int32(typeID))
		log.Printf("Current subgraphs - %#v", subgraphs)
		// Check to see if we already have a subgraph for this type
		if sg, ok := subgraphs[int32(typeID)]; ok {
			// We already have a subgraph, check if the node already exists
			if nd, ok := nodes[fmt.Sprintf("%d-%d", tm.ModifyingAttribute.TypeID, tm.ModifyingAttribute.AttributeID)]; ok {
				// So we already have the node, make sure it has been affected before using it as an affector
				if !tm.ModifyingAttribute.Affected {
					// This attribute has not been affected, do that first
					err := printTraverseAffectorList(tm.ModifyingAttribute, root, rootGraph, subgraphs, nodes, sd)
					if err != nil {
						return err
					}
				}
				// We now know that the modifying attribute has been affected. Create the edge between the affector and affected
				// Get the current attribute node
				cn, ok := nodes[fmt.Sprintf("%d-%d", dta.TypeID, dta.AttributeID)]
				if !ok {
					return errors.New(fmt.Sprintf("trying to make edge to a modified node that does not exist for %d-%d", dta.TypeID, dta.AttributeID))
				}
				nd.Edge(*cn)
				// Edge done, move onto the next one
				continue
			} else {
				// By reaching here, we have a subgraph for the type id of the modifier, but not the attribute id.
				// Therefore we need to create the node
				id := fmt.Sprintf("%d-%d", tm.ModifyingAttribute.TypeID, tm.ModifyingAttribute.AttributeID)
				nd := sg.Node(id)
				nodes[id] = &nd

				// This should always run but I am going to check it anyway
				if !tm.ModifyingAttribute.Affected {
					// This attribute has not been affected, do that first
					err := printTraverseAffectorList(tm.ModifyingAttribute, root, rootGraph, subgraphs, nodes, sd)
					if err != nil {
						return err
					}
				}

				// We now know that the modifying attribute has been affected. Create the edge between the affector and affected
				// Get the current attribute node
				cn, ok := nodes[fmt.Sprintf("%d-%d", dta.TypeID, dta.AttributeID)]
				if !ok {
					return errors.New(fmt.Sprintf("trying to make edge to a modified node that does not exist for %d-%d", dta.TypeID, dta.AttributeID))
				}
				nd.Edge(*cn)
				// Edge done, move onto the next one
				continue
			}
		} else {
			// By reaching this point, we know that there is no subgraph created for that type id yet
			// Therefore we will need to create both the subgraph and the node for the object

			// Need to look up the type from our sde records
			tname := sd.Types[typeID].TypeName.String
			ttype := "Type"
			if sd.Groups[sd.Types[typeID].GroupID.Int32].CategoryID.Int32 == 16 {
				ttype = "Skill"
			}

			tg := rootGraph.Subgraph(fmt.Sprintf("%s - %s", ttype, tname), dot.ClusterOption{})
			subgraphs[dta.TypeID] = tg

			log.Printf("Creating new subgraph for Type - %s, ID:%d", tname, int32(dta.TypeID))

			id := fmt.Sprintf("%d-%d", tm.ModifyingAttribute.TypeID, tm.ModifyingAttribute.AttributeID)
			nd := tg.Node(id)
			nodes[id] = &nd

			// This should always run but I am going to check it anyway
			if !tm.ModifyingAttribute.Affected {
				// This attribute has not been affected, do that first
				err := printTraverseAffectorList(tm.ModifyingAttribute, root, rootGraph, subgraphs, nodes, sd)
				if err != nil {
					return err
				}
			}

			// We now know that the modifying attribute has been affected. Create the edge between the affector and affected
			// Get the current attribute node
			cn, ok := nodes[fmt.Sprintf("%d-%d", dta.TypeID, dta.AttributeID)]
			if !ok {
				return errors.New(fmt.Sprintf("trying to make edge to a modified node that does not exist for %d-%d", dta.TypeID, dta.AttributeID))
			}
			nd.Edge(*cn)
			// Edge done, move onto the next one
			continue

		}
	}

	// All of the modifiers have been handled, this node is done
	return nil
}

func (solo *Solo) ResetAffectorTree() error {
	return nil
}
