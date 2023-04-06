package models

import (
	"log"
	// "time"
)

// type User struct {
// 	ID        int
// 	UUID      string
// 	Name      string
// 	Email     string
// 	PassWord  string
// 	CreatedAt time.Time
// 	// Todos     []Todo
// }
// ↓追加
type User struct {
	ID        int
	UUID      string
	Name      string
	Birth_day string
	Age       int
	Height    float64
	Weight    float64
}



func (u *User) CreateUser() (err error) {
	// cmd := `insert into users (
	// 	uuid,
	// 	name,
	// 	email,
	// 	password,
	// 	created_at) values (?, ?, ?, ?, ?)`
// ↓追加
	cmd := `insert into users (
		uuid,
		name,
		birth_day,
		age,
		height,
		weight)`

	// _, err = Db.Exec(cmd,
	// 	createUUID(),
	// 	u.Name,
	// 	u.Email,
	// 	Encrypt(u.PassWord),
	// 	time.Now())
// ↓追加
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Birth_day,
		u.Age,
		u.Height,
		u.Weight)
		
		if err != nil {
			log.Fatalln(err)
		}
		return err
	}

// func GetUser(id int) (user User, err error) {
// 	user = User{}
// 	cmd := `select id, uuid, name, birth_day, age, height, weight
// 	from users where id = ?`
// 	err = Db.QueryRow(cmd, id).Scan(
// 		&user.ID,
// 		&user.UUID,
// 		&user.Name,
// 		&user.Birth_day,
// 		&user.Age,
// 		&user.Height,
// 		&user.Weight,
// 	)
// 	return user, err
// }

// func (u *User) UpdateUser() (err error) {
// 	cmd := `update users set name = ?, age = ? where id = ?`
// 	_, err = Db.Exec(cmd, u.Name, u.Age, u.ID)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return err
// }

// func (u *User) DeleteUser() (err error) {
// 	cmd := `delete from users where id = ?`
// 	_, err = Db.Exec(cmd, u.ID)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return err
// }
