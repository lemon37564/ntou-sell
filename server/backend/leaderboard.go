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

func GetLeadersRaw() (string, error) {
	pd := database.GetLeaderRaw()
	str, err := json.Marshal(pd)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}

func GetLeadersOrdered(strength string, limit string) (string, error) {
	strengthInt, err := strconv.Atoi(strength)
	if err != nil {
		return "cannot convert " + strength + " into integer", err
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "cannot convert " + limit + " into integer", err
	}

	pd := database.GetLeaderOrdered(strengthInt, limitInt)
	str, err := json.Marshal(pd)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}
