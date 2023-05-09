package entity

type Content struct {
	ID                 int64  `db:"id"`
	Author             string `db:"author"`
	Header             string `db:"header"`
	Content            string `db:"content"`
	LastCommitedClicks int64  `db:"last_commited_clicks"`
	ClickExists        bool   `db:"click_exists"`
}

type Click struct {
	ContentID int64  `db:"content_id"`
	Date      string `db:"date"`
	Address   string `db:"address"`
}
