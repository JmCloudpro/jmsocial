package repository

import (
	"api/src/models"
	"database/sql"
)

// Publications represents a repository of publications
type Publications struct {
	db *sql.DB
}

// NewPublicationsRepo creates a new repository of Publicationss receiving db
func NewPubliRepo(db *sql.DB) *Publications {

	return &Publications{db}
}

//Create is a Method  and it inserts a Publications in  the database

func (repository Publications) Create(Publication models.Publication) (uint64, error) {

	statement, err := repository.db.Prepare(
		"insert into publications (title, content, creatorid) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(Publication.Title, Publication.Content, Publication.CreatorID)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastIDInserted), nil

}

/*
// Search  searches all Publicationss
func (repository Publications) SearchAll() ([]models.Publication, error) {
	// fmt.Sprintf("%%%s%")	  primeira % = escape,  2 entra na string,  3 é o %s que aguarda valor, e depois %%

	lines, err := repository.db.Query("select p.id, p.title, p.content, u.nick p.likes, p.createdin from publicarions p inner join users u on u.id = p.creatorid")
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var Publications []models.Publication

	for lines.Next() {
		var Publication models.Publication
		if err = lines.Scan(
			&Publication.ID,
			&Publication.Title,
			&Publication.Content,
			&Publication.CreatorNick,
			&Publication.Createdin,
		); err != nil {
			return nil, err
		}
		Publications = append(Publications, Publication)

	}
	return Publications, nil
}



// SearchID  searches one Publications by id
func (repository Publications) SearchID(PublicationsID uint64) (models.Publications, error) {
	lines, err := repository.db.Query("select id, name, nick, email, createdin from Publicationss where id = ?", PublicationsID)
	if err != nil {
		return models.Publications{}, err
	}

	defer lines.Close()

	var Publications models.Publications
	if lines.Next() {
		if err = lines.Scan(
			&Publications.ID,
			&Publications.Name,
			&Publications.Nick,
			&Publications.Email,
			&Publications.CreatedIn,
		); err != nil {
			return models.Publications{}, err
		}
	}
	return Publications, nil
}

// UpdatePublications updates the Publications data
func (repository Publications) UpdatePublications(ID uint64, Publications models.Publications) error {
	statement, err := repository.db.Prepare(
		"update Publicationss set name = ?, nick = ?, email = ? where id = ?")

	if err != nil {

		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(Publications.Name, Publications.Nick, Publications.Email, ID); err != nil {
		return err
	}
	return nil
}

// DeletePublications  deletes the Publications
func (repository Publications) DeletePublications(ID uint64) error {
	statement, err := repository.db.Prepare("delete from Publicationss where id = ?")
	if err != nil {

		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Publications) SearchEmail(email string) (models.Publications, error) {
	line, err := repository.db.Query("select id, passwd from Publicationss where email = ?", email)

	if err != nil {

		return models.Publications{}, err
	}
	defer line.Close()

	var Publications models.Publications
	if line.Next() {
		if err = line.Scan(&Publications.ID, &Publications.Passwd); err != nil {
			return models.Publications{}, err
		}

	}
	return Publications, nil

}

func (repository Publications) Follow(PublicationsID, followid uint64) error {

	statement, err := repository.db.Prepare("insert into followers (Publicationsid, followerid) values (?, ?)")
	if err != nil {

		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(PublicationsID, followid); err != nil {
		errocode := strings.Split(fmt.Sprint(err), " ")
		if errocode[1] == "1062:" {
			return errors.New("You already follow this Publications")

		}

		return err
	}

	return nil

}

func (repository Publications) Unfollow(PublicationsID, followid uint64) error {

	statement, err := repository.db.Prepare("delete from followers  where Publicationsid = ? and followerid = ?")
	if err != nil {

		return err
	}
	defer statement.Close()

	data, err := statement.Exec(PublicationsID, followid)
	if err != nil {
		return err
	}
	none, err := data.RowsAffected()
	if none == 0 {
		return errors.New("You don't  follow this Publications")

	}
	return nil

}

func (repository Publications) SearchFollowers(PublicationsID uint64) ([]models.Publications, error) {

	lines, err := repository.db.Query(
		"select u.id, u.name, u.nick, u.email, u.createdin from Publicationss u inner join followers f on u.id = f.followerid where f.Publicationsid = ?", PublicationsID)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var followers []models.Publications

	for lines.Next() {
		var follower models.Publications
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
		return nil, errors.New("This Publications doesnt have followers")
	}
	return followers, nil

}

// SearchPasswd searches the currnent password usind the ID get on token.
func (repository Publications) SearchPasswd(PublicationsID uint64) (string, error) {

	line, err := repository.db.Query("select passwd from Publicationss where id = ?", PublicationsID)

	if err != nil {
		return "", err
	}
	defer line.Close()

	var Publications models.Publications

	if line.Next() {
		fmt.Println(err)
		if err = line.Scan(&Publications.Passwd); err != nil {
			return "", err

		}
	}
	return Publications.Passwd, nil
}

// UpdatePasswd updates password from uer
func (repository Publications) UpdatePasswd(PublicationsID uint64, passhash string) error {

	statement, err := repository.db.Prepare("update Publicationss set passwd = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(passhash, PublicationsID); err != nil {
		return err
	}
	return nil

}
*/
