package data

import (
	"time"
	"fmt"
)

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

func (thread *Thread) NumReplies() (count int) {
	rows, err := DB.Query("SELECT count(*) FROM posts where thread_id = $1", thread.ID)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

func (thread *Thread) User() (user User) {
	user = User{}

	preparedStatement := "SELECT id, uuid, name, email, created_at FROM users WHERE id = $1"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return user
	}
	defer func() {
		err = stmt.Close()
	}()

	err = stmt.QueryRow(thread.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}
	}

	return user
}



func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format(layoutFormat)
}

func GetAllThreads() (threads []Thread, err error) {
	preparedStatement := "select id, uuid, topic, user_id, created_at from threads order by created_at desc"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return threads, err
	}
	defer func() {
		err = stmt.Close()
	}()

	rows, err := stmt.Query()
	if err != nil {
		return threads, err
	}

	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt); err != nil {
			return threads, err
		}

		threads = append(threads, thread)
	}
	err = rows.Close()
	if err != nil {
		return threads, err
	}

	fmt.Println(threads)
	return threads, err
}
