package database

import (
	"database/sql"
	"log"
	"time"
)

const leaderBoardTable = `
CREATE TABLE IF NOT EXISTS leaderboard(
	player_name varchar(32) NOT NULL,
	point int NOT NULL,
	strength int NOT NULL,
	played_date timestamp NOT NULL,
	PRIMARY KEY(player_name, played_date)
);`

type LeaderBoard struct {
	Name     string    `json:"player_name"`
	Point    int       `json:"point"`
	Strength int       `json:"strength"`
	GameDate time.Time `json:"game_date"`
}

var (
	leadAdd *sql.Stmt
	leadGet *sql.Stmt
)

func leaderBoardPrepare(db *sql.DB) {
	var err error

	const (
		add = "INSERT INTO leaderboard VALUES(?,?,?,?);"
		get = "SELECT * FROM leaderboard;"
	)

	if leadAdd, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if leadGet, err = db.Prepare(get); err != nil {
		log.Println(err)
	}
}

func AddLeader(name string, point int, str int, date time.Time) error {
	_, err := leadAdd.Exec(name, point, str, date)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func GetLeader() (all []LeaderBoard) {
	var (
		rows     *sql.Rows
		err      error
		name     string
		point    int
		strength int
		date     time.Time
	)
	all = make([]LeaderBoard, 0)

	rows, err = leadGet.Query()

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&name, &point, &strength, &date)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, LeaderBoard{Name: name, Point: point, Strength: strength, GameDate: date})
	}
	rows.Close()

	return
}
