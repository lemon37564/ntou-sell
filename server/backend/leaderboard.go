package backend

import (
	"encoding/json"
	"se/database"
	"strconv"
	"time"
)

func AddLeader(name string, point string, str string) (string, error) {
	point_int, err := strconv.Atoi(point)
	if err != nil {
		return "cannot convert " + point + " into integer", err
	}

	strength_int, err := strconv.Atoi(str)
	if err != nil {
		return "cannot conver " + str + " into integer", err
	}

	err = database.AddLeader(name, point_int, strength_int, time.Now())
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
