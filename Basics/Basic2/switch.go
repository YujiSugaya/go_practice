package main

import (
	"fmt"
	//runtimeパッケージは、Goのランタイムシステムに関する操作を提供します。ここでは、OSの情報を取得するために使用されます。
	"runtime"
)

func main(){
	fmt.Print("Go runs on")
	//runtime.GOOSは現在のOSを返す　macOSでは"darwin"、Linuxでは"linux"、Windowsでは"windows"などです。
	//os変数の中に現在のOSの値を入れる
	switch os := runtime.GOOS; os{
		//osがdarwinの場合実行、その後処理終了
	case "darwin":
		fmt.Println("OS X.")
		//osがLinuxの場合
	case "linux":
		fmt.Println("Linuasgsfdx.")
	//どのcaseにも一致しない場合はdefaultが実行される
	default:
		fmt.Printf("%s.\n", os)
	}
}

//他の言語との違い
//各caseブロックの末尾にbreak文を明示的に書く必要がありません。デフォルトで、caseが実行された後、switch文は終了します。

//複数の条件: 一つのcaseで複数の値を評価することができます。例えば、case "a", "b":のように記述できます。

// switch number {
// case 1, 3, 5, 7, 9:
// 	fmt.Println("It's an odd number.")
// case 2, 4, 6, 8, 10:
// 	fmt.Println("It's an even number.")
// default:
// 	fmt.Println("Number is out of range.")
// }

//もしcaseに一致した場合でも処理を続けたい場合は明示的に fallthroughを入れてあげる必要あり


//他の言語は整数型のみ: 伝統的に、C/C++のswitch文は整数型の値のみを受け付けます。
//switch文では、文字列、浮動小数点数、ブール値、さらにはカスタム型を含む任意の型を評価することができます。



// switch fruit {
// case "apple":
// 	fmt.Println("It's an apple.")
// case "banana":
// 	fmt.Println("It's a banana.")
// case "orange":
// 	fmt.Println("It's an orange.")
// default:
// 	fmt.Println("Unknown fruit.")
// }

// package main

// import "fmt"

// func main() {
//     isRaining := true

//     switch isRaining {
//     case true:
//         fmt.Println("Take an umbrella.")
//     case false:
//         fmt.Println("No need for an umbrella.")
//     }
// }