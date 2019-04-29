package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"

	"strconv"
	"strings"
	"time"
)

var logger * log.Logger

func main()  {
	app := newApp();
	app.Run(iris.Addr(":8000"))
}

func newApp() *iris.Application {
	app := iris.New();
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	return app
}


func initLog()  {
	f,_ := os.Create("/var/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

type lotteryController struct {
	Ctx iris.Context
}

type gift struct {
	id      int    // 奖品ID
	name    string // 奖品名称
	pic     string // 照片链接
	link    string // 链接
	inuse   bool   // 是否使用中
	rate    int    // 中奖概率，十分之N,0-9
	rateMin int    // 大于等于，中奖的最小号码,0-10
	rateMax int    // 小于，中奖的最大号码,0-10
}

const rateMax = 10


func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate","4,3,2,1,0")
	giftlist := giftRate(rate)
	return fmt.Sprintf("%v\n", giftlist)
}


func (c *lotteryController) GetLucky() map[string]interface{}{
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate","4,3,2,1")
	result := make (map[string]interface{})

	giftList := giftRate(rate)
	code := luckyCode()
	ok :=  false
	for _, data := range giftList{

		if data.rateMin <= int(code) && data.rateMax > int(code){

			ok = true
			if ok{
				result["code"]  = code
				result["uid"]  = uid
				result["id"]  = data.id
				result["name"]  = data.name
				result["link"]  = data.link
				result["pic"]  = data.pic
				saveLuckyData(uid, int(code), data.id , data.name, data.link, data.pic)
				break;
			}
		}
	}
	return result
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

func saveLuckyData(uid, code, id int, name, link, sendData string)  {
	logger.Printf("lucky, uid=%d, code=%d, gift=%d, name=%s, link=%s, data=%s ", uid, code, id, name, link, sendData)
	fmt.Printf("lucky, uid=%d, code=%d, gift=%d, name=%s, link=%s, data=%s ", uid, code, id, name, link, sendData)
}
func newGift() *[5]gift {
	giftlist := new([5]gift)

	g1 := gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		inuse:   true,
		rate:    4,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[0] = g1
	// 2 实物小奖
	g2 := gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		inuse:   true,
		rate:    3,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[1] = g2
	// 3 虚拟券，相同的编码
	g3 := gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		inuse:   true,
		rate:    2,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[2] = g3
	// 4 虚拟券，不相同的编码
	g4 := gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		inuse:   true,
		rate:    1,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[3] = g4
	// 5 虚拟币
	g5 := gift{
		id:      5,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[4] = g5
	return giftlist
}

func giftRate(rate string) *[5]gift {

	giftlist := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rate)

	rateStart := 0

	for i, data := range giftlist{

		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen{
			grate, _ = strconv.Atoi(rates[i])
		}
		giftlist[i].rate = grate
		giftlist[i].rateMin = rateStart
		giftlist[i].rateMax = rateStart + grate
		if giftlist[i].rateMax  >= rateMax {
			giftlist[i].rateMax = rateMax
			rateStart = 0
		}else{
			rateStart += grate
		}
	}
	fmt.Printf("giftlist=%v\n", giftlist)
	return giftlist;
}
