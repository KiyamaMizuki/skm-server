package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"local.com/db-module/model"
	"local.com/go-api/server"
	"local.com/go-api/skmdb"
)

//Retutn 返り値用の構造体を定義トークン発行用
type RetutnAuth struct {
	Token string `json:"token"`
}

type ErrAuth struct {
	Token string `json:"error"`
}

type Errtimetables struct {
	Time string `json:"error"`
}

func main() {
	g := gin.Default()
	node := g.Group("/node")
	{
		node.POST("", nodeSave)
		node.GET("", nodeGet)
	}
	road := g.Group("/road")
	{
		road.POST("", roadSave)
		road.GET("", roadGet)

	}
	location := g.Group("/location")
	{
		location.POST("/altitude", GetAltitude)
	}
	route := g.Group("/route")
	{
		route.POST("", GetRoute)
	}
	mail := g.Group("/mail")
	{
		mail.POST("", sendMail)
		mail.POST("/token", checkAuth)

	}
	timetable := g.Group("/timetable")
	{
		//ユーザーが検索をかけた時にアクセスするurl
		timetable.POST("/search", checkClassname)
		//初回ログイン時全授業名を投げるurl
		timetable.GET("/FirstAccess", Sendclass)
		//登録するかしないかの確認
		timetable.POST("/register", saveTimetable)
		//ユーザーが登録した授業情報を取得
		timetable.GET("/getclass", getclassinfo)
	}
	g.Run(":1323")
}

//ユーザーが登録した授業情報を取得
func getclassinfo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	token := c.Request.Header.Get("token")
	username := skmdb.GetUsername(token)
	if skmdb.JudgeRegisterClass(username) {
		c.JSON(http.StatusOK, skmdb.GetRegisterClass(username))
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Not registered",
		})
	}

}

//ユーザー情報と授業情報を登録
func saveTimetable(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	classname := c.PostForm("classname")
	//username := c.PostForm("username")
	token := c.Request.Header.Get("token")
	username := skmdb.GetUsername(token)

	//skmdb.Register(classname, username)
	c.JSON(http.StatusOK, gin.H{
		"message": skmdb.Register(classname, username),
	})
}

//授業名が一致していたら情報を返す
func checkClassname(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	//var retutn RetutnCalss
	var err Errtimetables
	classname := c.PostForm("classname")
	if !skmdb.JudgeClassname(classname) {
		err = Errtimetables{
			Time: "not found",
		}
		c.JSON(http.StatusOK, err)
	} else {
		c.JSON(http.StatusOK, skmdb.ReturnClassInformation(classname))
	}
}

//端末に授業名を全部送る
func Sendclass(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	token := c.Request.Header.Get("authorization")
	if skmdb.JudgToken(token) {
		c.JSON(http.StatusOK, skmdb.ReturnClassName())
	} else {
		c.JSON(401, "Unauthorized")
	}
}

//端末にトークン発行
func checkAuth(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	var retutn RetutnAuth
	var err ErrAuth
	mail := c.PostForm("mail")
	authcode := c.PostForm("authcode")
	if skmdb.Judge(mail, authcode) {
		retutn = RetutnAuth{
			Token: skmdb.Entry(mail, authcode),
		}
	} else {
		err = ErrAuth{
			Token: skmdb.Entry(mail, authcode),
		}
		c.JSON(http.StatusOK, err)
	}
	c.JSON(http.StatusOK, retutn)
}

// メールに対して認証コードを送る
func sendMail(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	mail := c.PostForm("mail")
	server.Send(mail)
	c.JSON(http.StatusOK, gin.H{
		"mail": mail,
	})
}

func nodeSave(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	latitude, _ := strconv.ParseFloat(c.PostForm("latitude"), 64)
	longitude, _ := strconv.ParseFloat(c.PostForm("longitude"), 64)
	nodeType, _ := strconv.Atoi(c.PostForm("node_type"))
	floor, _ := strconv.Atoi(c.PostForm("floor"))
	name := c.PostForm("name")
	data := server.SaveNode(latitude, longitude, nodeType, floor, name)
	c.JSON(http.StatusOK, gin.H{
		"node": data,
	})
}
func nodeGet(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(http.StatusOK, server.NodeGet())
}
func roadGet(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	print(server.GetRoad())
	c.JSON(http.StatusOK, gin.H{
		"road": server.GetRoad(),
	})
}
func roadSave(c *gin.Context) {
	var nodeStart model.Node
	var nodeEnd model.Node
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	distance, _ := strconv.ParseFloat(c.PostForm("distance"), 64)
	nodeone, _ := strconv.Atoi(c.PostForm("nodeid1"))
	nodetwo, _ := strconv.Atoi(c.PostForm("nodeid2"))
	floor, _ := strconv.Atoi(c.PostForm("floor"))
	c.JSON(http.StatusOK, gin.H{
		"road": server.SetRoad(nodeone, nodeStart, nodeEnd, nodetwo, distance, floor),
	})
}
func GetRoute(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	lat, _ := strconv.ParseFloat(c.PostForm("startLat"), 64)
	lng, _ := strconv.ParseFloat(c.PostForm("startLng"), 64)
	floor := c.PostForm("floor")
	nodeEnd := c.PostForm("nodeEnd")
	data := server.CalculateRoute(lat, lng, floor, nodeEnd)
	c.JSON(200, data)
}
func GetAltitude(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	pressure := c.PostForm("pressure")
	data := skmdb.CalculateAltitude(pressure)
	fmt.Print(data)
	c.JSON(200, data)
}
