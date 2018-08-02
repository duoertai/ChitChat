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

	return threads, err
}
