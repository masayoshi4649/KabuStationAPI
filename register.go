package kabusapi

import (
	"encoding/json"
	"net/url"
)

// ReqPutRegisterRegister は **PUT /register** のリクエスト。
//
// ### 概要
// 銘柄登録
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutRegisterRegister struct {
	Symbols []struct {
		Symbol   string `json:"Symbol,omitempty"`   // 銘柄コード
		Exchange int    `json:"Exchange,omitempty"` // 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	} `json:"Symbols"`
}

// ResPutRegisterRegister は **PUT /register** のレスポンス。
//
// ### 概要
// 銘柄登録
type ResPutRegisterRegister struct {
	// 現在登録されている銘柄のリスト
	RegistList []struct {
		Symbol   string `json:"Symbol,omitempty"`   // 銘柄コード
		Exchange int    `json:"Exchange,omitempty"` // 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	} `json:"RegistList,omitempty"`
}

// PutRegisterRegister は **PUT /register** を呼び出します。
//
// ### 概要
// 銘柄登録
//
// PUSH配信する銘柄を登録します。<br> API登録銘柄リストを開くには、kabuステーションAPIインジケーターを右クリックし「API登録銘柄リスト」を選択してください。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PutRegisterRegister(ReqPutRegisterRegister{ /* TODO: フィールドを設定 */ })
func PutRegisterRegister(req ReqPutRegisterRegister) (code int, res ResPutRegisterRegister, err error) {
	p := "/register"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
	code, data, err := doRequest("PUT", p, v, b, needAuth)
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

// ReqPutRegisterUnregister は **PUT /unregister** のリクエスト。
//
// ### 概要
// 銘柄登録解除
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutRegisterUnregister struct {
	// ※為替銘柄を登録する場合、銘柄名は"通貨A" + "/" + "通貨B"、市場コードは"300"で指定してください。 例：'Symbol': 'EUR/USD', "Exchange": 300
	Symbols []struct {
		Symbol   string `json:"Symbol,omitempty"`   // 銘柄コード
		Exchange int    `json:"Exchange,omitempty"` // 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	} `json:"Symbols"`
}

// ResPutRegisterUnregister は **PUT /unregister** のレスポンス。
//
// ### 概要
// 銘柄登録解除
type ResPutRegisterUnregister struct {
	// 現在登録されている銘柄のリスト
	RegistList []struct {
		Symbol   string `json:"Symbol,omitempty"`   // 銘柄コード
		Exchange int    `json:"Exchange,omitempty"` // 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	} `json:"RegistList,omitempty"`
}

// PutRegisterUnregister は **PUT /unregister** を呼び出します。
//
// ### 概要
// 銘柄登録解除
//
// # API登録銘柄リストに登録されている銘柄を解除します
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PutRegisterUnregister(ReqPutRegisterUnregister{ /* TODO: フィールドを設定 */ })
func PutRegisterUnregister(req ReqPutRegisterUnregister) (code int, res ResPutRegisterUnregister, err error) {
	p := "/unregister"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
	code, data, err := doRequest("PUT", p, v, b, needAuth)
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

// ReqPutRegisterUnregisterAll は **PUT /unregister/all** のリクエスト。
//
// ### 概要
// 銘柄登録全解除
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutRegisterUnregisterAll struct {
}

// ResPutRegisterUnregisterAll は **PUT /unregister/all** のレスポンス。
//
// ### 概要
// 銘柄登録全解除
type ResPutRegisterUnregisterAll struct {
	// 現在登録されている銘柄のリスト<br>※銘柄登録解除が正常に行われれば、空リストを返します。<br> 登録解除でエラー等が発生した場合、現在登録されている銘柄のリストを返します
	RegistList map[string]interface{} `json:"RegistList,omitempty"`
}

// PutRegisterUnregisterAll は **PUT /unregister/all** を呼び出します。
//
// ### 概要
// 銘柄登録全解除
//
// # API登録銘柄リストに登録されている銘柄をすべて解除します
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PutRegisterUnregisterAll(ReqPutRegisterUnregisterAll{ /* TODO: フィールドを設定 */ })
func PutRegisterUnregisterAll(req ReqPutRegisterUnregisterAll) (code int, res ResPutRegisterUnregisterAll, err error) {
	p := "/unregister/all"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("PUT", p, v, b, needAuth)
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
