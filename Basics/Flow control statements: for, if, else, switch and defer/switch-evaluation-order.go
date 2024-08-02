package main

import (
	"fmt"
	"time"
)



func main(){
	fmt.Println(time.Now())
	fmt.Println("When's Saturday?")
	//今日の曜日を取得して変数に代入する
	today := time.Now().Weekday()

	fmt.Println(time.Saturday)


	//time.Weekday型の定数で、0から6までの整数値に対応しています（0がSunday、6がSaturday）。これを直接加算することはできません。

	//土曜日は6
	switch time.Saturday{
		//今日の曜日(0である日曜日)+0 == 6　の場合　土曜日は今日！
	case today + 0:
		fmt.Println("Today.")
		//今日の曜日+1(5である金曜日) == 6の場合　明日
	case today + 1:
		fmt.Println("Tomorrow.")	
		//今日の曜日+2(4である木曜日) == 6の場合　明後日
	case today +2:
		fmt.Println("IN two days.")
		//それ以外の曜日
	default:
		fmt.Println("Too far away.")
	}
}

//一致する処理があったらその時点で処理終了

//time.Saturdayは、Go言語の標準ライブラリであるtimeパッケージに定義されている曜日の定数の一つです。
// const (
//     Sunday time.Weekday = iota
//     Monday
//     Tuesday
//     Wednesday
//     Thursday
//     Friday
//     Saturday
// )