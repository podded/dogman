package main

import (
	"fmt"
	"github.com/gobuffalo/envy"
	"github.com/podded/dogman/engine"
	"log"
	"strconv"
	"time"
)

func main() {
	mysqlAddress := envy.Get("DB_ADDR", "127.0.0.1")
	mysqlPortEnv := envy.Get("DB_PORT", "3306")
	mysqlUser := envy.Get("DB_USER", "root")
	mysqlPass := envy.Get("DB_PASS", "password")
	mysqlDB := envy.Get("DB_DATABASE", "sde")

	log.SetFlags(log.Ltime | log.Lshortfile)

	mysqlPort := 3306
	i, err := strconv.Atoi(mysqlPortEnv)
	if err == nil {
		mysqlPort = i
	}

	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPass, mysqlAddress, mysqlPort, mysqlDB)

	init := time.Now()

	dgm, err := engine.NewSolo(uri)
	if err != nil {
		log.Fatalln(err)
	}
	defer dgm.Close()

	// This is kind of a benchmark..... But not really, should implement a proper one later
	start := time.Now()

	err = dgm.InjectAllSkills()
	if err != nil {
		log.Fatalln(err)
	}
	err = dgm.SetAllSkillsLevel(0)
	if err != nil {
		log.Fatalln(err)
	}

	err = dgm.SetSkillLevel(3392, 5)
	if err != nil {
		log.Fatalln(err)
	}

	// Set ship type to Caracal
	err = dgm.SetShipID(621)
	if err != nil {
		log.Fatalln(err)
	}

	err = dgm.BuildAffectorTree()
	if err != nil {
		log.Fatalln(err)
	}

	err = dgm.PrintAffectorTree()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Dogma engine spent %v initialising\n", start.Sub(init))
	log.Printf("Dogma engine spent %v processing\n", time.Now().Sub(start))

}
