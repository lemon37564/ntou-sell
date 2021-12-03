package backend

import (
	"encoding/json"
	"se/database"
	"strconv"
	"time"
)

func AddLeader(name string, selfPoint string, enemyPoint string, str string) (string, error) {
	selfPointInt, err := strconv.Atoi(selfPoint)
	if err != nil {
		return "cannot convert " + selfPoint + " into integer", err
	}

	enemyPointInt, err := strconv.Atoi(enemyPoint)
	if err != nil {
		return "cannot convert " + enemyPoint + " into integer", err
	}

	strengthInt, err := strconv.Atoi(str)
	if err != nil {
		return "cannot conver " + str + " into integer", err
	}

	err = database.AddLeader(name, selfPointInt, enemyPointInt, strengthInt, time.Now())
	if err != nil {
		return "failed", err
	}

	return "", nil
}

func GetLeaders() (string, error) {
	pd := database.GetLeader()
	str, err := json.Marshal(pd)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}
