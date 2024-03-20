package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
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
