package skmdb

import (
	"strings"

	"local.com/db-module/model"

	"local.com/db-module/val"
)

//ユーザー情報と授業情報を登録
func Register(classname string, username string) string {
	db := val.MysqlConnect()
	timetable := model.Timetables{}
	user := model.Users{}
	registration := model.Registrations{}
	var count int
	var message string

	// 実行完了後DB接続を閉じる
	defer db.Close()

	//マイグレーション
	db.AutoMigrate(&registration)

	db.Where("Classname = ?", classname).Find(&timetable)
	db.Where("Username = ?", username).Find(&user)
	if timetable.Classname == classname && user.Username == username {
		//データベースに値を格納
		db.Where("Classname = ?", classname).Find(&registration).Count(&count)
		if count == 0 {
			db.Create(&model.Registrations{Userid: int(user.ID), Username: username, Classid: int(timetable.ID), Classname: classname})
			message = "ok"
		} else {
			message = "already registered"
		}
	} else if timetable.Classname != classname {
		message = "Non-existent class"
	} else if user.Username != username {
		message = "Non-existent user"
	}
	return message
}

//ユーザーが授業を一個でも登録しているのかを確認
func JudgeRegisterClass(username string) bool {
	db := val.MysqlConnect()
	registration := model.Registrations{}
	// 実行完了後DB接続を閉じる
	defer db.Close()
	//マイグレーション
	db.AutoMigrate(&registration)
	var judge bool
	var count int
	db.Table("registrations").Where("Username = ?", username).Find(&registration).Count(&count)
	if count == 0 {
		judge = false
	} else {
		judge = true
	}
	return judge
}

//GetRegisterClass ユーザーが登録した授業情報を全取得する。
func GetRegisterClass(username string) []model.ReturnTimetables {
	db := val.MysqlConnect()
	//timetable := model.Timetables{}

	registration := model.Registrations{}

	// 実行完了後DB接続を閉じる
	defer db.Close()

	//マイグレーション
	db.AutoMigrate(&registration)

	var registerclassid []int
	returntimetable := []model.ReturnTimetables{}

	db.Table("registrations").Where("Username = ?", username).Pluck("classid", &registerclassid)
	for i := 0; i < len(registerclassid); i++ {
		timetable := model.Timetables{}
		classid := registerclassid[i]
		db.Where("id = ?", classid).Find(&timetable)
		classitme := strings.Split(timetable.ClassTime, "・")
		grade := strings.Split(timetable.Grade, "・")
		var timetableClassname = timetable.Classname
		var timetableClassroom = timetable.Classroom
		var timetableTeachers = timetable.Teachers
		var timetableCredit = timetable.Credit
		var timetableDepartment = timetable.Department
		var timetableSemester = timetable.Semester
		var timetableDay = timetable.Day
		var timetableRequired = timetable.Required

		m := model.ReturnTimetables{
			Classname:  timetableClassname,
			Classroom:  timetableClassroom,
			Teachers:   timetableTeachers,
			Credit:     timetableCredit,
			Department: timetableDepartment,
			Semester:   timetableSemester,
			Day:        timetableDay,
			Required:   timetableRequired,
			ClassTime:  classitme,
			Grade:      grade,
		}
		returntimetable = append(returntimetable, m)

	}

	return returntimetable
}
