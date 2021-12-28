package backend

import (
	"encoding/json"
	"errors"
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

	// if cheat
	sum := selfPointInt + enemyPointInt
	// it is impossible that someone < 0 or sum > 64 or sum < 10
	if selfPointInt < 0 || enemyPointInt < 0 || selfPointInt > 64 || enemyPointInt > 64 || sum > 64 || sum < 10 {
		return "failed", errors.New("not accepted")
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

	if limitInt > 100 {
		limitInt = 100
	}

	pd := database.GetLeaderOrdered(strengthInt, limitInt)
	str, err := json.Marshal(pd)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}
