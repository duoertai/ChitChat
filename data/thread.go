package data

import "time"

const layoutFormat = "Jan 2, 2006 at 3:04pm"

type Thread struct {
	ID int
	UUID string
	Topic string
	UserID int
	CreatedAt time.Time
}

type Post struct {
	ID int
	UUID string
	Body string
	UserID int
	ThreadID int
	CreatedAt time.Time
}

func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format(layoutFormat)
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format(layoutFormat)
}

