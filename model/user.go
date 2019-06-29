package model

import (
	"fmt"
	"trollstagram-backend/db"
)

//User a model of user
type User struct {
	ID       int
	Username string
	Password string
	Name     string
	ImgPath  []string
	Posts    int
}

//GetByID get user by id
func GetByID(id int) (*User, error) {
	user := new(User)

	row := db.DB.QueryRow(`SELECT id,username,password,name,posts FROM usr WHERE id=$1`, id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Posts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

//AddFilePath add filepath to db
func AddFilePath(filePath string) error {
	statement := fmt.Sprintf("UPDATE usr SET imgpath=array_append(image,%s) WHERE id=1", filePath)
	_, err := db.DB.Exec(statement)
	return err
}
