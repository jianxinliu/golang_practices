package main

import (
	"fmt"
	"time"
)

func main() {

	// 高峰期瞬间涌入 10 个请求
	requests := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		requests <- i
	}

	close(requests)

	// Tick 返回一个 channel， 该 channel 会被按指定的时间间隔塞入当前时间点
	// 即，该 channel 的接收者 goroutine, 只能被限制在以指定的时间间隔执行（阻塞）
	limter := time.Tick(time.Millisecond * 2000)

	for range requests {
		// 使用 limter 控制 request 的处理
		<- limter
		handlerRequest()
	}
}

func handlerRequest() {
	fmt.Println("deal request at: ", time.Now().Format("2006-01-02 15:04:05"))
}

// golang time format layout 

// const (
//     stdLongMonth      = "January"
//     stdMonth          = "Jan"
//     stdNumMonth       = "1"
//     stdZeroMonth      = "01"          
//     stdLongWeekDay    = "Monday"
//     stdWeekDay        = "Mon"
//     stdDay            = "2"
//     stdUnderDay       = "_2"
//     stdZeroDay        = "02"
//     stdHour           = "15"
//     stdHour12         = "3"
//     stdZeroHour12     = "03"
//     stdMinute         = "4"
//     stdZeroMinute     = "04"
//     stdSecond         = "5"
//     stdZeroSecond     = "05"
//     stdLongYear       = "2006"
//     stdYear           = "06"
//     stdPM             = "PM"
//     stdpm             = "pm"
//     stdTZ             = "MST"
//     stdISO8601TZ      = "Z0700"  // prints Z for UTC
//     stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
//     stdNumTZ          = "-0700"  // always numeric
//     stdNumShortTZ     = "-07"    // always numeric
//     stdNumColonTZ     = "-07:00" // always numeric
// )