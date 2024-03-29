package database

import (
	"database/sql"
	"log"
	"time"
)

const leaderBoardTable = `
CREATE TABLE IF NOT EXISTS leaderboard(
	player_name varchar(32) NOT NULL,
	self_point int NOT NULL,
	enemy_point int NOT NULL,
	strength int NOT NULL,
	played_date timestamp NOT NULL,
	PRIMARY KEY(player_name, played_date)
);`

type LeaderBoard struct {
	Name       string    `json:"player_name"`
	SelfPoint  int       `json:"self_point"`
	EnemyPoint int       `json:"enemy_point"`
	Strength   int       `json:"strength"`
	GameDate   time.Time `json:"game_date"`
}

const TIMEOUT = time.Second * 5

var (
	leadAdd        *sql.Stmt
	leadGetRaw     *sql.Stmt
	leadGetOrdered *sql.Stmt

	cache      [3][]LeaderBoard // [easy, medium, hard]
	lastUpdate time.Time
	needUpdate [3]bool = [3]bool{true, true, true}
)

func leaderBoardPrepare(db *sql.DB) {
	var err error

	const (
		add        = "INSERT INTO leaderboard VALUES($1,$2,$3,$4,$5);"
		getRaw     = "SELECT * FROM leaderboard;"
		getOrdered = `
		SELECT player_name, self_point, enemy_point, played_date
		FROM leaderboard WHERE strength=$1
		ORDER BY (CAST(self_point AS FLOAT)/CAST((self_point+enemy_point) AS FLOAT)) DESC
		LIMIT $2;`
	)

	if leadAdd, err = db.Prepare(add); err != nil {
		panic(err)
	}

	if leadGetRaw, err = db.Prepare(getRaw); err != nil {
		panic(err)
	}

	if leadGetOrdered, err = db.Prepare(getOrdered); err != nil {
		panic(err)
	}
}

func AddLeader(name string, selfPoint int, enemyPoint int, str int, date time.Time) error {
	_, err := leadAdd.Exec(name, selfPoint, enemyPoint, str, date)
	if err != nil {
		log.Println(err)
		return err
	}
	needUpdate[str] = true

	return err
}

func GetLeaderRaw() (all []LeaderBoard) {
	var (
		rows       *sql.Rows
		err        error
		name       string
		selfPoint  int
		enemyPoint int
		strength   int
		date       time.Time
	)
	all = make([]LeaderBoard, 0)

	rows, err = leadGetRaw.Query()

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&name, &selfPoint, &enemyPoint, &strength, &date)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, LeaderBoard{Name: name, SelfPoint: selfPoint, EnemyPoint: enemyPoint, Strength: strength, GameDate: date})
	}
	rows.Close()

	return
}

func GetLeaderOrdered(strength int, limit int) (all []LeaderBoard) {
	var (
		rows       *sql.Rows
		err        error
		name       string
		selfPoint  int
		enemyPoint int
		date       time.Time
	)

	if !needUpdate[strength] && time.Since(lastUpdate) < TIMEOUT {
		return cache[strength]
	}

	all = make([]LeaderBoard, 0)

	rows, err = leadGetOrdered.Query(strength, limit)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&name, &selfPoint, &enemyPoint, &date)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, LeaderBoard{Name: name, SelfPoint: selfPoint, EnemyPoint: enemyPoint, Strength: strength, GameDate: date})
	}
	rows.Close()

	lastUpdate = time.Now()
	cache[strength] = all
	needUpdate[strength] = false

	return
}
