package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/go-toast/toast"
)

type eve struct {
	Name    string `json:"name"`
	EndTime string `json:"endTime"`
}

func pushMes(mes string) {

	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "到期提醒",
		Message: mes,
		Icon:    "C:\\Windows\\WinSxS\\amd64_microsoft-windows-dxp-deviceexperience_31bf3856ad364e35_10.0.19041.746_none_251e769058968366\\settings.ico", // 文件必须存在
		Actions: []toast.Action{
			// {"protocol", "知道了", "https://www.google.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func timeLine(allTime []eve) {

	nowTime := time.Now()

	for _, v := range allTime {
		endtime, _ := time.Parse("2006-01-02", v.EndTime)
		tempDays := endtime.Sub(nowTime).Hours() / 24
		days := int(math.Ceil(tempDays))
		if days < 0 {
			mes := fmt.Sprintf("%s已过期!!!!!", v.Name)
			pushMes(mes)
		}
		switch days {
		case 30, 20, 10, 7, 5, 3, 2, 1:
			mes := fmt.Sprintf("距离%s到期剩余%d天", v.Name, days)
			pushMes(mes)
		case 0:
			mes := fmt.Sprintf("%s今日到期", v.Name)
			pushMes(mes)
		}

	}
}

func main() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	var allList []eve
	json.Unmarshal(b, &allList)

	timeLine(allList)

}
