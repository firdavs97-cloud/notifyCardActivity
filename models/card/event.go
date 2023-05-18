package card

import (
	"context"
	"fmt"
	"stash.transparent.com/gol/notifyCardActivity/services/mysql"
)

type IFace interface {
	Save() (int64, error)
	Notify(int64) error
	Load(id int64) error
}

func NewActivityEvent(ctx context.Context, input ActivityEventInput) IFace {
	if ctx.Value("MockActivityEvent") != nil {
		return MockActivityEvent{input, ctx.Value("MockActivityEvent").(map[string][]interface{})}
	}
	return &ActivityEvent{input}
}

type ActivityEvent struct {
	ActivityEventInput
}

type ActivityEventInput struct {
	OrderType  string `json:"orderType"`g
	SessionId  string `json:"sessionId"`
	Card       string `json:"card"`
	EventDate  string `json:"eventDate"`
	WebsiteUrl string `json:"websiteUrl"`
}

func (a ActivityEvent) Notify(id int64) error {
	db, err := mysql.NewDB()
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Card activity event: %s on %s", a.OrderType, a.EventDate))
	err = db.Update("UPDATE cardActivityEvent set status = 'done' WHERE id = ?", id)
	return err
}

func (a ActivityEvent) Save() (int64, error) {
	db, err := mysql.NewDB()
	if err != nil {
		return 0, err
	}
	id, err := db.Insert("INSERT INTO cardActivityEvent(orderType, sessionId, card, eventDate, websiteUrl, status) VALUES (?, ?, ?, ?, ?, ?)",
		a.OrderType, a.SessionId, a.Card, a.EventDate, a.WebsiteUrl, "pending")
	return id, err
}

func (a *ActivityEvent) Load(id int64) error {
	db, err := mysql.NewDB()
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT orderType, sessionId, card, eventDate, websiteUrl FROM cardActivityEvent WHERE id = ?",
		id, &a.OrderType, &a.SessionId, &a.Card, &a.EventDate, &a.WebsiteUrl)
	return err
}
