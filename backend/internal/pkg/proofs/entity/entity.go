package entity

import "github.com/zero-pkg/null"

const (
	SubscriptionPrice = 100
	DateFormat        = "2006-01-02"
)

type Proof struct {
	V     int64  `json:"v"`
	R     string `json:"r"`
	S     string `json:"s"`
	Hash  string `json:"hash"`
	Date  string `json:"date"`
	Value int64  `json:"value"`
}

type DBProof struct {
	V                 int64       `db:"v" json:"v"`
	R                 string      `db:"r" json:"r"`
	S                 string      `db:"s" json:"s"`
	Hash              string      `db:"hash" json:"hash"`
	Date              string      `db:"date" json:"date"`
	Value             int64       `db:"value" json:"value"`
	Address           string      `db:"address" json:"address"`
	LastCommitedDate  null.String `db:"last_commited_date" json:"last_commited_date"`
	LastCommitedValue null.Int    `db:"last_commited_value" json:"last_commited_value"`
}
