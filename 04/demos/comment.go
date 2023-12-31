package main

import (
	"log"
	"time"
)

type comment struct {
	ID        uint
	UserID    string `sql:"type:varchar(36);primary key"` // Cognito UUID
	PhotoID   uint
	Text      string
	CreatedAt time.Time
}

// InsertComment inserts a comment record
func InsertComment(photoid uint, userid string, text string) (uint, error) {
	comment := &comment{
		Text:      text,
		PhotoID:   photoid,
		UserID:    userid,
		CreatedAt: time.Now(),
	}

	if err := db.Create(comment); err.Error != nil {
		return 0, err.Error
	}

	log.Println("Inserted comment record:", comment.ID)

	return comment.ID, nil
}

func (c *comment) Username() string {
	user, _ := findUserByID(c.UserID)
	return user.Username
}
