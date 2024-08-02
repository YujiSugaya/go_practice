package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	switch {
		//trueの場合実行
	case t.Hour() < 12:
		fmt.Println("Good morning!")
		//trueの場合実行
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
		//trueの場合実行
	default:
		fmt.Println("Good evening")
	}
}

//戻り値: Hour() メソッドは int 型の値を返します。これは0から23までの範囲で、現在の時刻の時を表します。たとえば、午前0時ならば0、午前1時ならば1、午後1時ならば13、午後11時ならば23などです。

