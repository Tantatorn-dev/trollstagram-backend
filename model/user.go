package model

import (
	"fmt"
	"log"
	"trollstagram-backend/db"

	"github.com/lib/pq"
)

//User a model of user
type User struct {
	ID       int
	Username string
	Password string
	Name     string
	ImgPath  []string
}

//GetByID get user by id
func GetByID(id int) (*User, error) {
	user := new(User)

	row := db.DB.QueryRow(`SELECT id,username,password,name FROM usr WHERE id=$1`, id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

//AddFilePath add filepath to db
func AddFilePath(filePath string) error {
	statement := fmt.Sprintf("UPDATE usr SET imgpath=array_append(imgpath,'%s') WHERE id=1", filePath)
	_, err := db.DB.Exec(statement)
	return err
}

//CountPosts count a number of posts
func CountPosts() int {
	var count int
	statement := "SELECT array_length(imgpath,1) FROM usr WHERE id=1"

	row := db.DB.QueryRow(statement)

	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//GetFilePaths get every file paths
func GetFilePaths() *[]string {
	var str []string
	statement := "SELECT imgpath FROM usr WHERE id=1"

	row := db.DB.QueryRow(statement)
	err := row.Scan(pq.Array(&str))
	if err != nil {
		log.Fatal(err)
	}
	return &str
}
