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
