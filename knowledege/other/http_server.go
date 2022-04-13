package knowledege

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func handle1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	str := `<h1 style="color:red">hello world</h1>`
	w.Write([]byte(str))
}

func Http_server() {
	http.HandleFunc("/posts/go/", handle1)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func Http_client() {
	//对于get请求 参数都放在url中 请求体body中没有数据的
	// response, err := http.Get("http://127.0.0.1:8080/xxx/")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//如何定制一个请求 和上面一样的请求
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "奥斯")
	data.Set("age", "10")
	quertStr := data.Encode()
	urlObj.RawQuery = quertStr
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	response, _ := http.DefaultClient.Do(req)
	fmt.Println(quertStr)
	fmt.Println(urlObj)
	fmt.Println(response)
	//从response中把服务器放回的数据读出来
	b, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))
}
