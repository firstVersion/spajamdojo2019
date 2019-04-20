package model

import (
  "fmt"

  "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Score struct {
	//UserID    int    `db:"user_id" json:"user_id"`
  //UserName  string `db:"user_name" json:"user_name"`
	//Term      int    `db:"term"    json:"term"`
	//TeamID    int    `db:"team_id" json:"team_id"`
  Distance  int    `db:"distance" json:"distance"`
	//Amount    int    `db:"amount"  json:"amount"`
}

type Event struct {
	FROM    string  `db:"t_from" json:"from"`
  TO      string  `db:"t_to" json:"to"`
}


func GetScores() Score {
	db, err := sqlx.Open("postgres", "host=postgres user=root password=root dbname=spajamdojo2019 sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	scores := []Score{}
	db.Select(&scores, `
    SELECT
    --u.user_id
    --,u.user_name
    --,s.term
    --,s.team_id
    s.distance
    --,b.amount
    FROM public.bets AS b
    , public.users AS u
    , public.scores AS s
    WHERE s.user_id = u.user_id
    AND b.user_id = s.user_id
    AND u.user_id = 1
    AND s.term = 1
    AND s.team_id = 1;`)
	return scores[0]
}

func GetEvents() []Event {
	db, err := sqlx.Open("postgres", "host=postgres user=root password=root dbname=spajamdojo2019 sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	events := []Event{}
	db.Select(&events, `
    SELECT
     MAX(t.t_from) AS t_from
    ,MAX(t.t_to) AS t_to
    FROM public.bets AS b
    , public.users  AS u
    , public.scores AS s
    , public.terms  AS t
    WHERE s.user_id = u.user_id
    AND b.user_id = u.user_id
    AND t.term_id = s.term
    AND u.user_id = 1
    GROUP BY s.term;
    `)
	fmt.Println(events)
	return events
}

func AddBet(term int, bet int) {
	db, err := sqlx.Open("postgres", "host=postgres user=root password=root dbname=spajamdojo2019 sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
  _, err = db.Exec(`INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 1, $1, 1, 0);`,term)
  if err != nil {
		fmt.Println(err)
	}
  _, err = db.Exec(`INSERT INTO bets ( user_id, amount, term, team_id ) VALUES( 1, $2, $1, 1);`,term,bet)
  if err != nil {
		fmt.Println(err)
	}
}

func AddDistance() {
	db, err := sqlx.Open("postgres", "host=postgres user=root password=root dbname=spajamdojo2019 sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
  _, err = db.Exec(`UPDATE public.scores SET distance=distance+1 WHERE user_id=1 AND term=1 AND team_id=1;`)
  if err != nil {
		fmt.Println(err)
	}
}
