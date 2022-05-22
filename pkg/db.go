package hermes

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db   *sql.DB
	conn string
}

func NewDatabase(conn string) *Database {
	return &Database{conn: conn}
}

func (db *Database) Connect() (err error) {
	db.db, err = sql.Open("sqlite3", db.conn)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() (err error) {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

// Statistics contains all metrics that are shown to the user.
type Statistics struct {
	TotalMessages    int     `json:"total_messages" yaml:"total-messages"`
	ReceivedMessages int     `json:"received_messages" yaml:"received-messages"`
	SentMessages     int     `json:"sent_messages" yaml:"sent-messages"`
	AvgDailyMessages float64 `json:"avg_daily_messages" yaml:"avg-daily-messages"`
	Chats            int     `json:"chats" yaml:"chats"`
}

func (db *Database) Statistics(ctx context.Context) (stats *Statistics, err error) {
	stats = &Statistics{}

	// All messages
	row := db.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM message")
	err = row.Scan(&stats.TotalMessages)
	if err != nil {
		return nil, err
	}

	// Sent messages
	row = db.db.QueryRowContext(ctx, "select count(*) from message where is_from_me = 1;")
	err = row.Scan(&stats.SentMessages)
	if err != nil {
		return nil, err
	}

	// Received messages
	row = db.db.QueryRowContext(ctx, "select count(*) from message where is_from_me = 0;")
	err = row.Scan(&stats.ReceivedMessages)
	if err != nil {
		return nil, err
	}

	// Average daily messages
	row = db.db.QueryRowContext(ctx, "SELECT AVG(m.c) FROM (select count(*) AS c from message GROUP BY strftime('%Y-%m-%d', datetime(date/1000000000 + strftime('%s','2001-01-01'), 'unixepoch'))) AS m")
	err = row.Scan(&stats.AvgDailyMessages)
	if err != nil {
		return nil, err
	}

	row = db.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM chat")
	err = row.Scan(&stats.Chats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

type Conversation struct {
	Messages []*Message `json:"messages"`
}

type Message struct {
	Text string `json:"text"`
}

func (db *Database) Conversation(ctx context.Context) (conversation *Conversation, err error) {

	return nil, nil
}

func (db *Database) ListConversations(ctx context.Context) (conversations []string, err error) {
	rows, err := db.db.QueryContext(ctx, "SELECT guid FROM chat")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	conversations = make([]string, 0)
	for rows.Next() {
		var guid string
		err = rows.Scan(&guid)
		if err != nil {
			return nil, err
		}
		conversations = append(conversations, guid)
	}
	return conversations, nil
}
