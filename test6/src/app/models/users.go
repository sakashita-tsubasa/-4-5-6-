package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int
	Name     string
	BirthDay string
	Age      int
	Height   float64
	Weight   float64
}

func (u *User) createUser() error {

	cmd := `insert into users (
		name,
		birthday,
		age,
		height,
		weight) values (?, ?, ?, ?, ?)`
	_, err := Db.Exec(cmd,
		u.Name,
		u.BirthDay,
		u.Age,
		u.Height,
		u.Weight)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil

}

func Create(c *gin.Context) {

	u := User{}
	type requestParams struct {
		Name     string  `json:"Name"`
		BirthDay string  `json:"birthDay"`
		Age      int     `json:"Age"`
		Height   float64 `json:"Height"`
		Weight   float64 `json:"Weight"`
	}
	var params requestParams
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "POSTMANからリクエスト取得に失敗しました"})
		return
	}

	u.Name = params.Name
	u.BirthDay = params.BirthDay
	u.Age = params.Age
	u.Height = params.Height
	u.Weight = params.Weight
	// ↓dbフォルダーに移動させるとここでこける
	if err := u.createUser(); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "ユーザー作成に失敗しました"})
		return
	}
	json, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(200, gin.H{"message": string(json)})

}

func getUser(id string, user *User) error {

	var err error
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.Name,
		&user.BirthDay,
		&user.Age,
		&user.Height,
		&user.Weight,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return nil

}

func Get(c *gin.Context) {

	user := User{}
	id := c.Param("id")
	if err := getUser(id, &user); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "ユーザー取得に失敗しました"})
		return
	}
	json, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(200, gin.H{"message": string(json)})

}

func (u *User) updateUser() error {

	var err error
	cmd := `update users set name = ?, birthday = ?, age = ?, height = ?, weight = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.BirthDay, u.Age, u.Height, u.Weight, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return nil

}

func Update(c *gin.Context) {

	u := User{}
	type requestParams struct {
		ID       int     `json:"ID"`
		Name     string  `json:"Name"`
		BirthDay string  `json:"birthDay"`
		Age      int     `json:"Age"`
		Height   float64 `json:"Height"`
		Weight   float64 `json:"Weight"`
	}
	var params requestParams
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "POSTMANからリクエスト取得に失敗しました"})
		return
	}
	u.ID = params.ID
	u.Name = params.Name
	u.BirthDay = params.BirthDay
	u.Age = params.Age
	u.Height = params.Height
	u.Weight = params.Weight
	if err := u.updateUser(); err != nil {
		fmt.Println(err)
		return
	}
	json, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(200, gin.H{"message": string(json)})

}

func (u *User) deleteUser(id string) error {

	var err error
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return nil

}

func Delete(c *gin.Context) {

	u := User{}
	id := c.Param("id")
	u.ID, _ = strconv.Atoi(id)
	// u.DeleteUser(id)
	if err := u.deleteUser(id); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "ユーザー削除に失敗しました"})
		return
	}

	c.JSON(200, gin.H{"message": "選択したユーザーを削除しました"})

}
