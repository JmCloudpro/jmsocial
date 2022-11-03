package repository

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type User struct {
	db *sql.DB
}

// NewUserRepo creates a new repository of users receiving db
func NewUsersRepo(db *sql.DB) *User {

	return &User{db}
}

//fmt.Println(user)

//Create is a Method  and it inserts a user in  the database

func (repository User) Create(user models.User) (uint64, error) {

	statement, err := repository.db.Prepare("insert into users (name, nick, email, passwd) values(?, ?, ?, ?)")

	if err != nil {

		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(user.Name, user.Nick, user.Email, user.Passwd)
	if err != nil {

		return 0, err
	}

	lastIDInserted, err := res.LastInsertId()
	if err != nil {

		return 0, err
	}
	return uint64(lastIDInserted), nil

}

// Search  searches all users
func (repository User) Search(nameornick string) ([]models.User, error) {
	// fmt.Sprintf("%%%s%")	  primeira % = escape,  2 entra na string,  3 é o %s que aguarda valor, e depois %%
	nameornick = fmt.Sprintf("%%%s%%", nameornick)

	lines, err := repository.db.Query("select id, name, nick, email, createdin from users where name LIKE ? or nick LIKE ?", nameornick, nameornick)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

// SearchID  searches one user by id
func (repository User) SearchID(userID uint64) (models.User, error) {
	lines, err := repository.db.Query("select id, name, nick, email, createdin from users where id = ?", userID)
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// UpdateUser updates the user data
func (repository User) UpdateUser(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?")

	if err != nil {

		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}
	return nil
}

// DeleteUser  deletes the user
func (repository User) DeleteUser(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {

		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository User) SearchEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, passwd from users where email = ?", email)

	if err != nil {

		return models.User{}, err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Passwd); err != nil {
			return models.User{}, err
		}

	}
	return user, nil

}

func (repository User) Follow(userID, followid uint64) error {

	statement, err := repository.db.Prepare("insert into followers (userid, followerid) values (?, ?)")
	if err != nil {

		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followid); err != nil {
		errocode := strings.Split(fmt.Sprint(err), " ")
		if errocode[1] == "1062:" {
			return errors.New("You already follow this user")

		}

		return err
	}

	return nil

}

func (repository User) Unfollow(userID, followid uint64) error {

	statement, err := repository.db.Prepare("delete from followers  where userid = ? and followerid = ?")
	if err != nil {

		return err
	}
	defer statement.Close()

	data, err := statement.Exec(userID, followid)
	if err != nil {
		return err
	}
	none, err := data.RowsAffected()
	if none == 0 {
		return errors.New("You don't  follow this user")

	}
	return nil

}

func (repository User) SearchFollowers(userID uint64) ([]models.User, error) {

	lines, err := repository.db.Query(
		"select u.id, u.name, u.nick, u.email, u.createdin from users u inner join followers f on u.id = f.followerid where f.userid = ?", userID)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User
		if err = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedIn,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)

	}
	if followers == nil {
		return nil, errors.New("This user doesnt have followers")
	}
	return followers, nil

}

// SearchPasswd searches the currnent password usind the ID get on token.
func (repository User) SearchPasswd(userID uint64) (string, error) {

	line, err := repository.db.Query("select passwd from users where id = ?", userID)

	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		fmt.Println(err)
		if err = line.Scan(&user.Passwd); err != nil {
			return "", err

		}
	}
	return user.Passwd, nil
}

// UpdatePasswd updates password from uer
func (repository User) UpdatePasswd(userID uint64, passhash string) error {

	statement, err := repository.db.Prepare("update users set passwd = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(passhash, userID); err != nil {
		return err
	}
	return nil

}
