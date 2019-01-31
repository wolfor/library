package Http

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// description:
// http post file
// return:
// http request variable
func HttpPostFile(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)

	if err != nil {
		return nil, err
	}
	//     这里的io.Copy实现,会把file文件都读取到内存里面，然后当做一个buffer传给NewRequest. 对于大文件来说会占用很多内存
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()

	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)

	request.Header.Set("Content-Type", writer.FormDataContentType())

	return request, err
}

func HttpRequestDo(request *http.Request) ([]byte, error) {
	var resp *http.Response

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Println("http.Do failed,[err=%s]", err)
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("http.Do failed,[err=%s]", err)
		return nil, err
	}

	return b, err
}

func Get(url string) string {

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return ""
	}

	//	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	//	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	//	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	//	reqest.Header.Set("Cache-Control", "max-age=0")
	//	reqest.Header.Set("Connection", "keep-alive")
	//	reqest.Header.Set("User-Agent", "chrome 100")

	response, err := client.Do(reqest)

	if err != nil {
		return ""
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {

		body, _ := ioutil.ReadAll(response.Body)
		bodyStr := string(body)
		//		log.Println(bodyStr)
		return bodyStr
	} else {
		return ""
	}
}
