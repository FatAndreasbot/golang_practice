package t2

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var connection *sql.DB
var initialyzer sync.Once

func getDBConnection() (*sql.DB){
	var err error
	initialyzer.Do(func(){
		connection, err = sql.Open("sqlite3", "./data.db")
	})
	if err != nil {
		// panics in production is a bad habit, but i feel it makes sense in server init
		panic(err)
	}
	return connection
}

func GetAllComments() (chan *Comment, error){
	conn := getDBConnection()
	stmt := "SELECT c.id, c.User_ID, c.text FROM Comments AS c"

	rows, err := conn.Query(stmt)
	if err != nil{
		return nil, errors.Join(err, errors.New("could not fetch with query"))
	}
	defer rows.Close()

	var allComments chan *Comment

	go func(){
		for rows.Next(){
			var comment Comment
			err := rows.Scan(&comment.id, &comment.authorID, &comment.text)
			if err != nil{
				log.Println(errors.Join(err, errors.New("error when reading db response")))
			}
			allComments <- &comment
		}
		close(allComments)
	}()
	return allComments, nil
}

func GetAuthor(c *Comment) (User, error){
	conn := getDBConnection()
	stmt := fmt.Sprintf(
		"SELECT u.ID, u.Username FROM \"User\" u where u.ID = %d", c.authorID,
	)

	rows, err := conn.Query(stmt)
	if err != nil{
		return User{}, errors.Join(err, errors.New("could not fetch with query"))
	}
	defer rows.Close()

	var user User

	if !rows.Next(){
		return User{}, errors.New("could not find author")
	}
	err = rows.Scan(&user.id, &user.username)
	if err != nil{
		return User{}, err
	}

	return user, nil
}
