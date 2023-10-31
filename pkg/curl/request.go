package curl

import (
	"bytes"
	"io/ioutil"
	"net/http"

	syslog "snail/pkg/log"

	"github.com/gin-gonic/gin"

	"github.com/idoubi/goz"
	"github.com/tidwall/gjson"
)

// Get 请求
func Get(c *gin.Context, url string, data interface{}, headers map[string]interface{}) (gjson.Result, error) {
	var res gjson.Result
	cli := goz.NewClient()
	resp, err := cli.Get(url, goz.Options{
		Headers: headers,
		Query:   data,
	})
	if err != nil {
		syslog.Error(c, "post-error", err.Error())
		return res, err
	}
	body, err := resp.GetBody()
	if err != nil {
		syslog.Error(c, "post-error", err.Error())
		return res, err
	}
	res = gjson.Parse(body.GetContents())
	return res, nil
}

// PostForm 请求
func PostForm(c *gin.Context, url string, data map[string]interface{}, headers map[string]interface{}) (gjson.Result, error) {
	var res gjson.Result
	cli := goz.NewClient()
	resp, err := cli.Post(url, goz.Options{
		Headers:    headers,
		FormParams: data,
	})
	if err != nil {
		syslog.Error(c, "post-error", err.Error())
		return res, err
	}
	body, err := resp.GetBody()
	if err != nil {
		syslog.Error(c, "post-error", err.Error())
		return res, err
	}
	res = gjson.Parse(body.GetContents())
	return res, nil
}

// PostJson 请求
func PostJson(c *gin.Context, url string, data string, headers map[string]string) (gjson.Result, error) {
	var res gjson.Result
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	for index := range headers {
		req.Header.Set(index, headers[index])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		syslog.Error(c, "post-error", err.Error())
		return res, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		syslog.Error(c, "post-error", err.Error())
		return res, err
	}
	res = gjson.Parse(string(content))
	return res, nil
}
