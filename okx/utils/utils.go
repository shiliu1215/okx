package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"okex/models"
	"strconv"
	"strings"
	"time"
)

func GetOkxSign(timestamp, path string) string {
	r := timestamp + "GET" + path
	hash := hmac.New(sha256.New, []byte("9F69653B7E2906F30464DE4C8FD94864"))
	hash.Write([]byte(r))
	sign := hash.Sum(nil)
	base64Encoded := base64.StdEncoding.EncodeToString(sign)
	return base64Encoded
}
func SendGetRequest(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return body, nil
}
func GetIsoTime() string {
	now := time.Now()
	utcTime := now.UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

// GetOkxTime 获取带有时间偏移量的 OKX 时间戳
func GetOkxTime(offsetSeconds int64) string {
	now := time.Now().UTC()
	expiryTime := now.Add(time.Duration(offsetSeconds) * time.Second)
	return expiryTime.Format("2006-01-02T15:04:05.999Z")
}
func GetOkxTimeNew() string {
	// 获取 ISO 8601 格式时间戳
	now := time.Now().UTC()
	timestamp := now.Format("2006-01-02T15:04:05.999Z")
	timestamp = strings.TrimSuffix(timestamp, "Z") + "Z" // 去除毫秒部分后重新加上
	return timestamp
}
func EpochTime() string {
	millisecond := time.Now().UnixNano() / 1000000
	epoch := strconv.Itoa(int(millisecond))
	epochBytes := []byte(epoch)
	epoch = string(epochBytes[:10])
	return epoch
}
func PreHashString(timestamp string, method string, requestPath string, body string) string {
	return timestamp + strings.ToUpper(method) + requestPath + body
}
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func OkxPost(path string, req interface{}) (code int, body []byte, err error) {

	timestamp := GetIsoTime()

	marshal, _ := json.Marshal(req)
	preHash := PreHashString(timestamp, "POST", path, string(marshal))
	sign, err := HmacSha256Base64Signer(preHash, models.SecretKey)
	httpReq, err := http.NewRequest("POST", models.ApiUrl+path, bytes.NewBuffer(marshal))
	if err != nil {
		log.Printf("发送OKX POST请求 失败 url=%v err=%v ", path, err)
		return 0, nil, err
	}
	// 添加请求头
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-simulated-trading", "1") //测试环境
	httpReq.Header.Set("OK-ACCESS-KEY", models.AccessKey)
	httpReq.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	httpReq.Header.Set("OK-ACCESS-PASSPHRASE", models.PASSPHRASE)
	httpReq.Header.Set("OK-ACCESS-SIGN", sign)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("发送OKX POST请求 失败 url=%v err=%v ", path, err)
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("获取body失败 url=%v err=%v ", path, err)
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}
func OkxGet(path string) (code int, body []byte, err error) {

	timestamp := GetIsoTime()
	preHash := PreHashString(timestamp, "GET", path, "")
	sign, err := HmacSha256Base64Signer(preHash, models.SecretKey)
	httpReq, err := http.NewRequest("GET", models.ApiUrl+path, nil)
	if err != nil {
		log.Printf("发送OKX GET请求 失败 url=%v err=%v ", path, err)
		return 0, nil, err
	}
	// 添加请求头
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-simulated-trading", "1") //测试环境
	httpReq.Header.Set("OK-ACCESS-KEY", models.AccessKey)
	httpReq.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	httpReq.Header.Set("OK-ACCESS-PASSPHRASE", models.PASSPHRASE)
	httpReq.Header.Set("OK-ACCESS-SIGN", sign)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("发送OKX GET请求 失败 url=%v err=%v ", path, err)
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("发送OKX GET请求 失败 url=%v err=%v ", path, err)
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}
