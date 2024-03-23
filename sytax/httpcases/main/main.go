package main

// "bytes"
// "encoding/json"
// "fmt"
// "net/http"

const (
	Ldate         = 1 << iota     // 日期
	Ltime                         // 时间
	Lmicroseconds                 // 微秒
	Llongfile                     // 完成文件名称
	Lshortfile                    // 短文件名称
	LUTC                          // 时区
	Lmsgprefix                    // 前缀
	LstdFlags     = Ldate | Ltime // 初始值
)

func main() {
	// var str string = "hello,world!"
	// json, err := json.Marshal(str)
	// if err != nil {
	// 	fmt.Println("json.Marshal(str)错误= ", err.Error())
	// }
	// rs := bytes.NewReader(json)
	// resp, err := http.Post("https://baidu.com", "application/json;charset=utf-8", rs)
	// if err != nil {
	// 	fmt.Println("http.Post错误= ", err.Error())
	// }
	// defer resp.Body.Close()

}
