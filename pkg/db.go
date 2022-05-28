package hermes

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sort"
	"strings"
	"time"
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
	TotalMessages      int       `json:"total_messages" yaml:"total-messages"`
	ReceivedMessages   int       `json:"received_messages" yaml:"received-messages"`
	SentMessages       int       `json:"sent_messages" yaml:"sent-messages"`
	AvgDailyMessages   float64   `json:"avg_daily_messages" yaml:"avg-daily-messages"`
	AvgMonthlyMessages float64   `json:"avg_monthly_messages" yaml:"avg-monthly-messages"`
	AvgYearlyMessages  float64   `json:"avg_yearly_messages" yaml:"avg-yearly-messages"`
	Chats              int       `json:"chats" yaml:"chats"`
	FirstMessage       time.Time `json:"first_message" yaml:"first-message"`
	LastMessage        time.Time `json:"last_message" yaml:"last-message"`
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

	row = db.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM chat")
	err = row.Scan(&stats.Chats)
	if err != nil {
		return nil, err
	}

	row = db.db.QueryRowContext(ctx, "SELECT MIN(date), MAX(date) FROM message")
	var firstMessage int64
	var lastMessage int64
	err = row.Scan(&firstMessage, &lastMessage)
	if err != nil {
		return nil, err
	}
	// Epoch in iMessage starts at 2001-01-01 as opposed to 1970-01-01 in Go.
	stats.FirstMessage = time.UnixMicro(firstMessage/1e3).
		AddDate(31, 0, 0).
		Round(time.Second)
	stats.LastMessage = time.UnixMicro(lastMessage/1e3).
		AddDate(31, 0, 0).
		Round(time.Second)

	// Total days are defined as full 24-hour blocks
	totalDays := float64(int(stats.LastMessage.Sub(stats.FirstMessage).Hours() / 24))
	stats.AvgDailyMessages = float64(stats.TotalMessages) / totalDays

	totalMonths := monthsSince(stats.FirstMessage, stats.LastMessage)
	if totalMonths > 0 {
		stats.AvgMonthlyMessages = float64(stats.TotalMessages) / float64(totalMonths)
	}

	totalYears := yearsSince(stats.FirstMessage, stats.LastMessage)
	if totalYears > 0 {
		stats.AvgYearlyMessages = float64(stats.TotalMessages) / float64(totalYears)
	}

	return stats, nil
}

// monthsSince counts the months between two points in time, rounding up to the next full month.
func monthsSince(t1 time.Time, t2 time.Time) int {
	if t2.Before(t1) {
		return 0
	}
	count := 1
	month := t1.UTC().Month()
	for t1.Before(t2) {
		if t1.Month() != month {
			count += 1
			month = t1.UTC().Month()
		}
		t1 = t1.AddDate(0, 0, 1)
	}
	return count
}

// yearsSince returns the number of full years between two timestamps.
func yearsSince(t1 time.Time, t2 time.Time) int {
	return monthsSince(t1, t2) / 12
}

// Participant is represented by a handle in the database
type Participant struct {
	ID      int    `json:"id" yaml:"id"`
	Number  string `json:"number" yaml:"number"`
	Country string `json:"country" yaml:"country"`
	Service string `json:"service" yaml:"service"`
}

// Chat has a unique ID and at least one participant.
type Chat struct {
	ID           int            `json:"id" yaml:"id"`
	Participants []*Participant `json:"participants,omitempty" yaml:"participants,omitempty"`
	Messages     []*Message     `json:"messages,omitempty" yaml:"messages,omitempty"`
}

type Message struct {
	SenderID int       `json:"sender_id" yaml:"sender-id"`
	Text     string    `json:"text" yaml:"text"`
	Date     time.Time `json:"date" yaml:"date"`
}

// Conversation returns all messages from a conversation with the given identifier if available.
func (db *Database) Conversation(ctx context.Context, chatID int) (chat *Chat, err error) {
	rows, err := db.db.QueryContext(ctx, `SELECT
    datetime (message.date / 1000000000 + strftime ("%s", "2001-01-01"), "unixepoch", "localtime") AS message_date,
    message.text,
    message.is_from_me,
    handle."ROWID", 
    handle.id, 
    handle.country, 
    handle.service
FROM
    chat
    JOIN chat_message_join ON chat. "ROWID" = chat_message_join.chat_id
    JOIN message ON chat_message_join.message_id = message. "ROWID"
    LEFT JOIN handle on message.handle_id = handle."ROWID"
WHERE
    chat."ROWID" = ?`, chatID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	participantMap := make(map[int]*Participant)
	chat = &Chat{
		Participants: make([]*Participant, 0),
		Messages:     make([]*Message, 0),
	}
	for rows.Next() {
		var dateString string
		p := &Participant{}
		m := &Message{}
		var isFromMe bool
		var text sql.NullString
		var participantID sql.NullInt64
		var participantNumber sql.NullString
		var participantCountry sql.NullString
		var participantService sql.NullString
		err = rows.Scan(&dateString, &text, &isFromMe, &participantID, &participantNumber, &participantCountry,
			&participantService)
		if err != nil {
			return nil, err
		}

		if participantID.Valid {
			p.ID = int(participantID.Int64)
		}
		if participantNumber.Valid {
			p.Number = participantNumber.String
		}
		if participantCountry.Valid {
			p.Country = participantCountry.String
		}
		if participantService.Valid {
			p.Service = participantService.String
		}

		// Copy sender ID to message
		m.SenderID = p.ID

		if text.Valid {
			m.Text = text.String
		}
		m.Date, err = time.Parse("2006-01-02 15:04:05", dateString)
		if err != nil {
			panic(err)
		}

		chat.Messages = append(chat.Messages, m)

		if p.ID != 0 {
			if _, ok := participantMap[p.ID]; !ok {
				participantMap[p.ID] = p
			}
		}
	}

	for _, p := range participantMap {
		chat.Participants = append(chat.Participants, p)
	}

	return chat, nil
}

func (db *Database) ListConversations(ctx context.Context) (chats []*Chat, err error) {
	rows, err := db.db.QueryContext(ctx, `SELECT
    chat."ROWID", handle."ROWID", handle.id, handle.country, handle.service
FROM
    chat
    JOIN chat_handle_join ON chat."ROWID" = chat_handle_join.chat_id
    JOIN handle ON chat_handle_join.handle_id = handle."ROWID"`)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	chatMap := make(map[int]*Chat)
	for rows.Next() {
		var chatID int
		var p Participant
		err = rows.Scan(&chatID, &p.ID, &p.Number, &p.Country, &p.Service)
		if err != nil {
			return nil, err
		}

		// Normalize values
		p.Service = strings.ToUpper(p.Service)
		p.Country = strings.ToUpper(p.Country)

		if _, ok := chatMap[chatID]; !ok {
			chatMap[chatID] = &Chat{
				ID:           chatID,
				Participants: make([]*Participant, 0),
			}
		}
		chatMap[chatID].Participants = append(chatMap[chatID].Participants, &p)
	}

	chats = make([]*Chat, 0)
	for _, chat := range chatMap {
		chats = append(chats, chat)
	}

	// Sort chats by ID so that they are always shown in the same order.
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].ID < chats[j].ID
	})

	return chats, nil
}
