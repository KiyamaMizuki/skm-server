package skmdb

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"local.com/db-module/model"
	"local.com/db-module/val"
)

func Maketimetable() {
	//読み込むPDFへのパスを書き込む
	file, err := os.Open("./test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string
	var lenght = 0
	var capacity = 0
	var array = make([][]string, lenght, capacity)

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}

		array = append(array, line)

	}
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		//キャストする必要がある

		SaveData(array[i][0], array[i][1], array[i][2], array[i][3], array[i][4], array[i][5], array[i][6], array[i][7], array[i][8], array[i][9])
	}

}

func SaveData(name string, room string, Teachers string, Credit string, Department string, Semester string, Day string, Required string, classTime string, grade string) (timetables model.Timetables) {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()
	timetable := model.Timetables{}

	//マイグレーション
	db.AutoMigrate(model.Timetables{})

	//値を設定
	timetable.Classname = name
	timetable.Classroom = room
	timetable.Teachers = Teachers
	timetable.Credit = Credit
	timetable.Department = Department
	timetable.Semester = Semester
	timetable.Day = Day
	timetable.Required = Required
	timetable.ClassTime = classTime
	timetable.Grade = grade

	//データベースに値を格納
	db.Create(&timetable)
	return timetable
}

//授業名がDBに格納されていたらその情報を返す。
func ReturnClassInformation(class string) []model.ReturnTimetables {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()

	//マイグレーション
	db.AutoMigrate(model.Timetables{})

	timetable := model.Timetables{}
	returntimetable := []model.ReturnTimetables{}
	//capacity := 0
	//var info = make([]string, capacity)

	db.Where("Classname = ?", class).Find(&timetable)
	classitme := strings.Split(timetable.ClassTime, "・")
	grade := strings.Split(timetable.Grade, "・")

	//さっき取得したDBの情報を変数に代入
	var timetable_classname = timetable.Classname
	var timetable_classroom = timetable.Classroom
	var timetable_teachers = timetable.Teachers
	var timetable_credit = timetable.Credit
	var timetable_department = timetable.Department
	var timetable_semester = timetable.Semester
	var timetable_day = timetable.Day
	var timetable_required = timetable.Required

	m := model.ReturnTimetables{
		Classname:  timetable_classname,
		Classroom:  timetable_classroom,
		Teachers:   timetable_teachers,
		Credit:     timetable_credit,
		Department: timetable_department,
		Semester:   timetable_semester,
		Day:        timetable_day,
		Required:   timetable_required,
		ClassTime:  classitme,
		Grade:      grade,
	}

	returntimetable = append(returntimetable, m)

	//info = append(info, timetable.Classname, timetable.Classroom, timetable.Teachers, timetable.Credit, timetable.Department, timetable.Semester, timetable.Day, timetable.Required, timetable.ClassTime, timetable.Grade)

	return returntimetable

}

//授業名が登録されているか判定
func JudgeClassname(classname string) bool {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()
	//マイグレーション
	db.AutoMigrate(&model.Timetables{})

	timetable := model.Timetables{}

	var judge bool

	judge = db.Where("Classname = ?", classname).Find(&timetable).RecordNotFound()

	//ないならfalseあるならture
	return !judge
}

//最初にアクセスした時に授業名を全部渡す処理
func JudgToken(token string) bool {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()
	user := model.Users{}
	var existtoken bool = false
	var count int = 0
	db.Where("token = ?", token).Find(&user).Count(&count)
	//.Error
	if count > 0 {
		existtoken = true
	}
	return existtoken
}

//授業名をすべて返す関数
func ReturnClassName() []string {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()
	timetable := model.Timetables{}
	var classnames []string
	db.Model(&timetable).Pluck("Classname", &classnames)
	return classnames
}
