package data

import "time"

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

type Session struct {
	ID int
	UUID string
	Email string
	UserID int
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	preparedStatement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return session, err
	}
	defer func() {
		err = stmt.Close()
	}()

	err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.CreatedAt)
	return session, err
}

func (user *User) GetSession() (session Session, err error) {
	session = Session{}
	preparedStatement := "select id, uuid, email, user_id, created_at from sessions where user_id = $1"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return session, err
	}
	defer func() {
		err = stmt.Close()
	}()

	err = stmt.QueryRow(user.ID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return session, err
}

func (session *Session) CheckSession() (valid bool, err error) {
	preparedStatement := "select id, uuid, email, user_id, created_at from sessions where uuid = $1"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		valid = false
		return valid, err
	}
	defer func() {
		err = stmt.Close()
	}()


	return valid, err
}

func (session *Session) DeleteByUUID() (err error) {
	preparedStatement := "delete from sessions where uuid = $1"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return err
	}
	defer func() {
		err = stmt.Close()
	}()

	stmt.Exec(session.UUID)
	return err
}
