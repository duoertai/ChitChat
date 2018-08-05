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

	err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
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

	err = stmt.QueryRow(session.UUID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	return session.ID != 0, err
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

func (session *Session) GetUserFromSession() (user User, err error) {
	user = User{}

	preparedStatement := "select id, uuid, name, email, created_at from users where id = $1"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
	}
	defer func() {
		err = stmt.Close()
	}()

	err = stmt.QueryRow(session.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func (user *User) CreateUser() (err error) {
	preparedStatement := "insert into users (uuid, name, email, password, created_at) values($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return err
	}
	defer func() {
		err = stmt.Close()
	}()

	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.ID, &user.UUID, &user.CreatedAt)
	return err
}

func (user *User) CreateThread(topic string) (thread Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), topic, user.ID, time.Now()).Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt)
	return
}

func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), body, user.ID, thread.ID, time.Now()).Scan(&post.ID, &post.UUID, &post.Body, &post.UserID, &post.ThreadID, &post.CreatedAt)
	return
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	preparedStatement := "select id, uuid, name, email, password, created_at from users where email = $1"
	stmt, err := DB.Prepare(preparedStatement)
	if err != nil {
		return user, err
	}
	defer func() {
		err = stmt.Close()
	}()

	err = stmt.QueryRow(email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}
