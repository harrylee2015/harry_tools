package main

import (
	"time"
	"sync/atomic"
	"fmt"
	"strconv"
	"math"
)
const (
	//剪刀
	Scissor = iota + 1
	//石头
	Rock
	//布
	Paper
)
const  (
	//游戏结果
	//平局
	IsDraw = iota + 1
	IsCreatorWin
	IsMatcherWin
	//开奖超时
	IsTimeOut

	//从有matcher参与游戏开始计算本局游戏开奖的有效时间，单位为天
	Active_Time = 1
)
//1534478054
func main() {

	/*
  math包：
   */
	i := -100
	fmt.Println(math.Abs(float64(i))) //绝对值
	fmt.Println(math.Ceil(5.0))       //向上取整
	fmt.Println(math.Floor(5.8))      //向下取整
	fmt.Println(math.Mod(11.0, 3.0))      //取余数，同11%3
	fmt.Println(math.Modf(5.26))      //取整数，取小数
	fmt.Println(math.Pow(3, 2))       //x的y次方
	fmt.Println(math.Pow10(4))        // 10的n次方
	fmt.Println(math.Sqrt(8))         //开平方
	fmt.Println(math.Cbrt(8))         //开立方
	fmt.Println(math.Pi)

	var total int64
	for i:=1;i<10;i++{
		atomic.AddInt64(&total,int64(i))
		fmt.Println(total)
	}
	//fmt.Println(math.Remainder(4,2))
	//fmt.Println(math.Remainder(5,2))
	//fmt.Println(math.Mod(4,2))
	//fmt.Println(math.Mod(7,2))
	fmt.Println(IsMatcherWin)
	fmt.Println(IsTimeOut)
	//SetTimeDelta(1)
	//time1 := Now().Unix()
	//fmt.Println(time1)
	//time.Sleep(time.Second)
	//fmt.Println(Now().Unix())

	//格式化为字符串,tm为Time类型
   //timestamp:=1534478054
	tm := time.Unix(1534478054, 0)
	fmt.Println(strconv.FormatInt(10, 10))

	count, _ := strconv.ParseInt(strconv.FormatInt(10, 10), 10, 64)
	fmt.Println(count)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
}
var deltaTime int64
var timeCalibration bool
var NtpHosts = []string{
	"time.windows.com:123",
	"ntp.ubuntu.com:123",
	"pool.ntp.org:123",
	"cn.pool.ntp.org:123",
	"time.asia.apple.com:123",
}

func SetFixTime(openTimeCalibration bool) {
	timeCalibration = openTimeCalibration
}

func IsFixTime() bool {
	return timeCalibration
}

//realtime - localtime
//超过60s 不做修正
//为了系统的安全，我们只做小范围时间错误的修复
func SetTimeDelta(dt int64) {
	if dt > 300*int64(time.Second) || dt < -300*int64(time.Second) {
		dt = 0
	}
	atomic.StoreInt64(&deltaTime, dt)
}

func Now() time.Time {
	dt := time.Duration(atomic.LoadInt64(&deltaTime))
	return time.Now().Add(dt)
}

func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}