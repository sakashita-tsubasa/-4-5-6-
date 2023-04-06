package main

import (
	"fmt"
	// "log"
	// "test6/config"
	"test6/app/models"
)

func main () {

	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

// ↓追加
	u := &models.User{}
	u.Name = "坂下"
	u.Birth_day = "19970523"
	u.Age = 25
	u.Height = 168.5
	u.Weight = 61.3
	fmt.Println(u)


	u.CreateUser()

	// u, _ := models.GetUser(2)

	// fmt.Println(u)
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
}





// 【テスト６】CRUD実装
// DBへの「登録、更新、表示、削除」を実装すること。
// いずれもJSONで値を返却すること（POSTMANでチェックします）
// 画面および仕様書は（あるに越したことはないが）必須ではない
// 想定している環境は Django on docker だが、動作すれば環境は問わない。 ←docker
// 言語＆FWも自分の好きなもので良い ←go
// DBの項目は以下の通り
// | name      | type   | key     | length | nullable | auto increment | description 
// | :-------- | :----- | :------ | -----: | :------: | :------------- | :-----------
// | id        | uint   | primary |        | x        | ○              | ID
// | name      | string |         | 255    | x        |                | 名前
// | birth_day | date   |         |        | x        |                | 誕生日
// | age       | int    |         |        | x        |                | 年齢
// | height    | float  |         |        | x        |                | 身長
// | weight    | float  |         |        | x        |                | 体重
