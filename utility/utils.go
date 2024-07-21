package utility

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"strings"
	"time"
)

func MD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateJWT(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId // 将用户ID加入到token
	claims["exp"] = time.Now().Add(120 * time.Minute).Unix()

	JwtSecretKey := os.Getenv("JWT_SECRET")
	signedString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", fmt.Errorf("Signature generation error %v", err)
	}
	return signedString, nil
}

func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	JwtSecretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	} else {
		return nil, nil, fmt.Errorf("invalid token")
	}
}

var (
	hostUrl   = "wss://aichat.xf-yun.com/v1/chat"
	appid     = "0c20dd9e"
	apiSecret = "NDgyYTk0N2M0MzNhY2Q4NDcyNzdmMzhl"
	apiKey    = "1295738dfd63e9411200132c8449c855"
)

func QueryAI(query string) (string, error) {
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	conn, _, err := d.Dial(assembleAuthUrl1(hostUrl, apiKey, apiSecret), nil)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	data := genParams1(appid, query)
	conn.WriteJSON(data)

	var answer string
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return "", err
		}

		var data map[string]interface{}
		err = json.Unmarshal(msg, &data)
		if err != nil {
			return "", err
		}

		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			return "", fmt.Errorf("error code: %v", code)
		}
		status := choices["status"].(float64)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		answer += content

		if status == 2 {
			break
		}
	}

	return answer, nil
}

func genParams1(appid, question string) map[string]interface{} {
	messages := []Message{
		{Role: "user", Content: question},
	}

	return map[string]interface{}{
		"header": map[string]interface{}{
			"app_id": appid,
		},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				"domain":      "general",
				"temperature": float64(0.8),
				"top_k":       int64(6),
				"max_tokens":  int64(2048),
				"auditing":    "default",
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{
				"text": messages,
			},
		},
	}
}

func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		panic(err)
	}
	date := time.Now().UTC().Format(time.RFC1123)
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	sgin := strings.Join(signString, "\n")
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
