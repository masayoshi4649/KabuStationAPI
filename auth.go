package kabusapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// BaseURL は API のベースURL。
// 変更したい場合は SetBaseURL("http://localhost:18081/kabusapi") のように設定してください。
var BaseURL = "http://localhost:18080/kabusapi"

// HTTPClient は REST 呼び出しに使用されるクライアントです。必要に応じて差し替え可能です。
var HTTPClient = &http.Client{Timeout: 30 * time.Second}

// SetHTTPClient は使用する http.Client を差し替えます。
func SetHTTPClient(c *http.Client) {
	if c != nil {
		HTTPClient = c
	}
}

// SetBaseURL は BaseURL を差し替えます。
func SetBaseURL(u string) {
	if u != "" {
		BaseURL = strings.TrimRight(u, "/")
	}
}

// apiKey は X-API-KEY ヘッダーに設定するトークンを保持します。
var apiKey string

// SetAPIKey は X-API-KEY ヘッダー用のトークンを設定します。
func SetAPIKey(k string) { apiKey = k }

// APIKey は現在設定されているトークンを返します。
func APIKey() string { return apiKey }

// ErrorResponse はエラー時のレスポンスボディです。
type ErrorResponse struct {
	Code    int    `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// APIError は非2xxのHTTPレスポンスを表すエラーです。
type APIError struct {
	StatusCode int    // HTTPステータスコード
	Code       int    // API固有のエラーコード
	Message    string // APIメッセージ
	Body       string // 生のレスポンスボディ
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return fmt.Sprintf("api error: status=%d code=%d msg=%s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("api error: status=%d body=%s", e.StatusCode, e.Body)
}

// doRequest は内部用のHTTP呼び出しです。必要に応じて X-API-KEY を付与します。
func doRequest(method, path string, query url.Values, body []byte, needAuth bool) (code int, data []byte, err error) {
	u := BaseURL + path
	if query != nil && len(query) > 0 {
		qs := query.Encode()
		if qs != "" {
			if strings.Contains(u, "?") {
				u += "&" + qs
			} else {
				u += "?" + qs
			}
		}
	}
	var r io.Reader
	if len(body) > 0 {
		r = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, u, r)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Accept", "application/json")
	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	if needAuth {
		if apiKey == "" {
			return 0, nil, fmt.Errorf("missing API key: SetAPIKey() が必要です")
		}
		req.Header.Set("X-API-KEY", apiKey)
	}
	resp, err := HTTPClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, b, nil
}

// ReqPostAuthToken は **POST /token** のリクエスト。
//
// ### 概要
// トークン発行
type ReqPostAuthToken struct {
	// APIパスワード
	APIPassword string `json:"APIPassword" validate:"required"`
}

// ResPostAuthToken は **POST /token** のレスポンス。
//
// ### 概要
// トークン発行
type ResPostAuthToken struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	ResultCode int `json:"ResultCode,omitempty"`
	// APIトークン
	Token string `json:"Token,omitempty"`
}

// PostAuthToken は **POST /token** を呼び出します。
//
// ### 概要
// トークン発行
//
// APIトークンを発行します。<br> 発行したトークンは有効である限り使用することができ、リクエストごとに発行する必要はありません。<br> 発行されたAPIトークンは以下のタイミングで無効となります。<br> ・kabuステーションを終了した時<br> ・kabuステーションからログアウトした時<br> ・別のトークンが新たに発行された時<br> ※kabuステーションは早朝、強制的にログアウトいたしますのでご留意ください。<br>
//
// ### 使い方
//
//	code, res, err := PostAuthToken(ReqPostAuthToken{ /* TODO: フィールドを設定 */ })
func PostAuthToken(req ReqPostAuthToken) (code int, res ResPostAuthToken, err error) {
	p := "/token"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := false
	code, data, err := doRequest("POST", p, v, b, needAuth)
	if err != nil {
		return code, res, err
	}
	if code >= 200 && code < 300 {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &res); err != nil {
				return code, res, err
			}
		}
		return code, res, nil
	}
	var apiErr ErrorResponse
	if len(data) > 0 {
		_ = json.Unmarshal(data, &apiErr)
	}
	return code, res, &APIError{StatusCode: code, Code: apiErr.Code, Message: apiErr.Message, Body: string(data)}
}
