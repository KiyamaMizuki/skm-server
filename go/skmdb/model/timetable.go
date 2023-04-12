package model

//timetable テーブルの定義
type Timetables struct {
	ID         uint   `gorm:"primaryKey"` //ID
	Classname  string //授業名
	Classroom  string //教室名
	Teachers   string //先生
	Credit     string //単位
	Department string //学部
	Semester   string //後期or前期
	Day        string //曜日
	Required   string //必修or選択・・・追加
	ClassTime  string //授業時間 2コマある時は・で区切る
	Grade      string //受講可能年次
}

//return用の構造体
type ReturnTimetables struct {
	Classname  string   //授業名
	Classroom  string   //教室名
	Teachers   string   //先生
	Credit     string   //単位
	Department string   //学部
	Semester   string   //後期or前期
	Day        string   //曜日
	Required   string   //必修or選択・・・追加
	ClassTime  []string //授業時間 2コマある時は・で区切る
	Grade      []string //受講可能年次
}
