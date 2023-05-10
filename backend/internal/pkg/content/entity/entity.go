package entity

type Content struct {
	ID      int64  `db:"id"`
	Author  string `db:"author"`
	Header  string `db:"header"`
	Content string `db:"content"`

	ClickExists bool `db:"click_exists"`
}

type ContentClickAggregate struct {
	ID     int64  `db:"id"`
	Author string `db:"author"`
	Header string `db:"header"`

	ClicksCommited   int64 `db:"clicks_commited"`
	ClicksUncommited int64 `db:"clicks_uncommited"`
}

type AuthorClickAggregate struct {
	Author   string `db:"author"`
	Commited bool   `db:"commited"`
	Count    int64  `db:"count"`
}

type Click struct {
	ContentID int64  `db:"content_id"`
	Date      string `db:"date"`
	Address   string `db:"address"`
	Commited  bool   `db:"commited"`
}
