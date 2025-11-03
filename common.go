package kabusapi

import (
	"net/http"
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
