package kabusapi

import (
	"encoding/json"
	"net/url"
)

// ReqPostOrderSendorder は **POST /sendorder** のリクエスト。
//
// ### 概要
// 注文発注（現物・信用）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPostOrderSendorder struct {
	// 銘柄コード
	Symbol string `json:"Symbol" validate:"required"`
	// 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>9</td> <td>SOR</td> </tr> <tr> <td>27</td> <td>東証+</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange" validate:"required"`
	// 商品種別 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>株式</td> </tr> </tbody> </table>
	SecurityType int `json:"SecurityType" validate:"required"`
	// 売買区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"Side" validate:"required"`
	// 信用区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>2</td> <td>新規</td> </tr> <tr> <td>3</td> <td>返済</td> </tr> </tbody> </table>
	CashMargin int `json:"CashMargin" validate:"required"`
	// 信用取引区分<br>※現物取引の場合は指定不要。<br>※信用取引の場合、必須。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>制度信用</td> </tr> <tr> <td>2</td> <td>一般信用（長期）</td> </tr> <tr> <td>3</td> <td>一般信用（デイトレ）</td> </tr> </tbody> </table>
	MarginTradeType int `json:"MarginTradeType"`
	// １株あたりのプレミアム料(円)<br> ※プレミアム料の刻値は、プレミアム料取得APIのレスポンスにある"TickMarginPremium"にてご確認ください。<br> ※入札受付中(19:30～20:30)プレミアム料入札可能銘柄の場合、「MarginPremiumUnit」は必須となります。<br> ※入札受付中(19:30～20:30)のプレミアム料入札可能銘柄以外の場合は、「MarginPremiumUnit」の記載は無視されます。<br> ※入札受付中以外の時間帯では、「MarginPremiumUnit」の記載は無視されます。
	MarginPremiumUnit float64 `json:"MarginPremiumUnit"`
	// 受渡区分<br>※現物買は指定必須。<br>※現物売は「0(指定なし)」を設定<br>※信用新規は「0(指定なし)」を設定<br>※信用返済は指定必須 <br>※auマネーコネクトが有効の場合にのみ、「3」を設定可能 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>指定なし</td> </tr> <tr> <td>2</td> <td>お預り金</td> </tr> <tr> <td>3</td> <td>auマネーコネクト</td> </tr> </tbody> </table>
	DelivType int `json:"DelivType" validate:"required"`
	// 資産区分（預り区分）<br>※現物買は、指定必須。<br>※現物売は、「' '」 半角スペース2つを指定必須。<br>※信用新規と信用返済は、「11」を指定するか、または指定なしでも可。指定しない場合は「11」が自動的にセットされます。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>(半角スペース2つ)</td> <td>現物売の場合</td> </tr> <tr> <td>02</td> <td>保護</td> </tr> <tr> <td>AA</td> <td>信用代用</td> </tr> <tr> <td>11</td> <td>信用取引</td> </tr> </tbody> </table>
	FundType string `json:"FundType"`
	// 口座種別 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>一般</td> </tr> <tr> <td>4</td> <td>特定</td> </tr> <tr> <td>12</td> <td>法人</td> </tr> </tbody> </table>
	AccountType int `json:"AccountType" validate:"required"`
	// 注文数量<br>※信用一括返済の場合、返済したい合計数量を入力してください。
	Qty int `json:"Qty" validate:"required"`
	// 決済順序<br>※信用返済の場合、必須。<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>日付（古い順）、損益（高い順）</td> </tr> <tr> <td>1</td> <td>日付（古い順）、損益（低い順）</td> </tr> <tr> <td>2</td> <td>日付（新しい順）、損益（高い順）</td> </tr> <tr> <td>3</td> <td>日付（新しい順）、損益（低い順）</td> </tr> <tr> <td>4</td> <td>損益（高い順）、日付（古い順）</td> </tr> <tr> <td>5</td> <td>損益（高い順）、日付（新しい順）</td> </tr> <tr> <td>6</td> <td>損益（低い順）、日付（古い順）</td> </tr> <tr> <td>7</td> <td>損益（低い順）、日付（新しい順）</td> </tr> </tbody> </table>
	ClosePositionOrder int `json:"ClosePositionOrder"`
	// 返済建玉指定<br>※信用返済の場合、必須。<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。<br>※信用一括返済の場合、各建玉IDと返済したい数量を入力してください。<br>※建玉IDは「E」から始まる番号です。
	ClosePositions []struct {
		HoldID string `json:"HoldID,omitempty"` // 返済建玉ID
		Qty    int    `json:"Qty,omitempty"`    // 返済建玉数量
	} `json:"ClosePositions"`
	// 執行条件 ※SOR以外は以下、全て指定可能です。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> <th>”Price"の指定</th> <th>SORで発注可</th> </tr> </thead> <tbody> <tr> <td>10</td> <td>成行</td> <td>0</td> <td>〇</td> </tr> <tr> <td>13</td> <td>寄成（前場）</td> <td>0</td> <td> </td> </tr> <tr> <td>14</td> <td>寄成（後場）</td> <td>0</td> <td> </td> </tr> <tr> <td>15</td> <td>引成（前場）</td> <td>0</td> <td> </td> </tr> <tr> <td>16</td> <td>引成（後場）</td> <td>0</td> <td> </td> </tr> <tr> <td>17</td> <td>IOC成行</td> <td>0</td> <td> </td> </tr> <tr> <td>20</td> <td>指値</td> <td>発注したい金額</td> <td>〇</td> </tr> <tr> <td>21</td> <td>寄指（前場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>22</td> <td>寄指（後場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>23</td> <td>引指（前場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>24</td> <td>引指（後場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>25</td> <td>不成（前場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>26</td> <td>不成（後場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>27</td> <td>IOC指値</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>30</td> <td>逆指値</td> <td>指定なし<br>※AfterHitPriceで指定ください</td> <td>〇</td> </tr> </tbody> </table>
	FrontOrderType int `json:"FrontOrderType" validate:"required"`
	// 注文価格<br>※FrontOrderTypeで成行を指定した場合、0を指定する。<br>※詳細について、”FrontOrderType”をご確認ください。
	Price float64 `json:"Price" validate:"required"`
	// 注文有効期限<br> yyyyMMdd形式。<br> 「0」を指定すると、kabuステーション上の発注画面の「本日」に対応する日付として扱います。<br> 「本日」は直近の注文可能日となり、以下のように設定されます。<br> 引けまでの間 : 当日<br> 引け後 : 翌取引所営業日<br> 休前日 : 休日明けの取引所営業日<br> ※ 日替わりはkabuステーションが日付変更通知を受信したタイミングです。
	ExpireDay int `json:"ExpireDay" validate:"required"`
	// 逆指値条件<br> ※FrontOrderTypeで逆指値を指定した場合のみ必須。
	ReverseLimitOrder struct {
		TriggerSec        int     `json:"TriggerSec,omitempty"`        // トリガ銘柄<br> ※未設定の場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>発注銘柄</td> </tr> <tr> <td>2</td> <td>NK225指数</td> </tr> <tr> <td>3</td> <td>TOPIX指数</td> </tr> </tbody> </table>
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // トリガ価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。
		UnderOver         int     `json:"UnderOver,omitempty"`         // 以上／以下<br> ※未設定の場合はエラーになります。<br> ※1、2以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>以下</td> </tr> <tr> <td>2</td> <td>以上</td> </tr> </tbody> </table>
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // ヒット後執行条件<br> ※未設定の場合はエラーになります。<br> ※1、2、3以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>成行</td> </tr> <tr> <td>2</td> <td>指値</td> </tr> <tr> <td>3</td> <td>不成</td> </tr> </tbody> </table>
		AfterHitPrice     float64 `json:"AfterHitPrice,omitempty"`     // ヒット後注文価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。<br><br> ヒット後執行条件に従い、下記のようにヒット後注文価格を設定してください。 <table> <thead> <tr> <th>ヒット後執行条件</th> <th>設定価格</th> </tr> </thead> <tbody> <tr> <td>成行</td> <td>0</td> </tr> <tr> <td>指値</td> <td>指値の単価</td> </tr> <tr> <td>不成</td> <td>不成の単価</td> </tr> </tbody> </table>
	} `json:"ReverseLimitOrder"`
}

// ResPostOrderSendorder は **POST /sendorder** のレスポンス。
//
// ### 概要
// 注文発注（現物・信用）
type ResPostOrderSendorder struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PostOrderSendorder は **POST /sendorder** を呼び出します。
//
// ### 概要
// 注文発注（現物・信用）
//
// 注文を発注します。<br> 同一の銘柄に対しての注文は同時に5件ほどを上限としてご利用ください。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PostOrderSendorder(ReqPostOrderSendorder{ /* TODO: フィールドを設定 */ })
func PostOrderSendorder(req ReqPostOrderSendorder) (code int, res ResPostOrderSendorder, err error) {
	p := "/sendorder"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
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

// ReqPostOrderSendorderFuture は **POST /sendorder/future** のリクエスト。
//
// ### 概要
// 注文発注（先物）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPostOrderSendorderFuture struct {
	// 銘柄コード<br>※取引最終日に「先物銘柄コード取得」でDerivMonthに0（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	Symbol string `json:"Symbol" validate:"required"`
	// 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> <tr> <td>32</td> <td>SOR日通し</td> </tr> <tr> <td>33</td> <td>SOR日中</td> </tr> <tr> <td>34</td> <td>SOR夜間</td> </tr> </tbody> </table> SOR日中、SOR夜間、SOR日通しは一部銘柄のみ対象となります。<br> SOR対象銘柄は以下をご参照ください。 <table> <thead> <tr> <th>先物SOR取扱銘柄</th> <th>有効限月</th> </tr> </thead> <tbody> <tr> <td>日経225先物ラージ</td> <td>直近2限月</td> </tr> <tr> <td>日経225先物ミニ</td> <td>直近4限月</td> </tr> <tr> <td>TOPIX先物ラージ</td> <td>直近2限月</td> </tr> <tr> <td>TOPIX先物ミニ </td> <td>直近3限月</td> </tr> <tr> <td>東証マザーズ指数先物</td> <td>直近2限月</td> </tr> <tr> <td>JPX日経400先物</td> <td>直近2限月</td> </tr> <tr> <td>NYダウ先物</td> <td>直近2限月</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange" validate:"required"`
	// 取引区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>新規</td> </tr> <tr> <td>2</td> <td>返済</td> </tr> </tbody> </table>
	TradeType int `json:"TradeType" validate:"required"`
	// 有効期間条件 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>FAS</td> </tr> <tr> <td>2</td> <td>FAK</td> </tr> <tr> <td>3</td> <td>FOK</td> </tr> </tbody> </table> ※執行条件(FrontOrderType)、有効期限条件(TimeInForce)、市場コード(Exchange)で選択できる組み合わせは下表のようになります。<br> <br>■日中、夜間、日通し対応表 <table> <thead> <tr> <th rowspan="2">執行条件</th> <th rowspan="2">有効<br>期間条件</th> <th colspan="3">市場コード</th> </tr> <tr> <th>日中</th> <th>夜間</th> <th>日通し</th> </tr> </thead> <tbody> <tr> <td>指値</td> <td>FAS</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>指値</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>指値</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>逆指値（指値）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>逆指値（成行）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引成</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引指</td> <td>FAS</td> <td>●</td> <td>●</td> <td>-</td> </tr> </tbody> </table> <br>■SOR日中、SOR夜間、SOR日通し対応表 <table> <thead> <tr> <th rowspan="2">執行条件</th> <th rowspan="2">有効<br>期間条件</th> <th colspan="3">市場コード</th> </tr> <tr> <th>SOR日中</th> <th>SOR夜間</th> <th>SOR日通し</th> </tr> </thead> <tbody> <tr> <td>指値</td> <td>FAS</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>指値</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>指値</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>逆指値（指値）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>逆指値（成行）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引成</td> <td>FAK</td> <td>-</td> <td>-</td> <td>-</td> </tr> <tr> <td>引指</td> <td>FAS</td> <td>-</td> <td>-</td> <td>-</td> </tr> </tbody> </table>
	TimeInForce int `json:"TimeInForce" validate:"required"`
	// 売買区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"Side" validate:"required"`
	// 注文数量
	Qty int `json:"Qty" validate:"required"`
	// 決済順序<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>日付（古い順）、損益（高い順）</td> </tr> <tr> <td>1</td> <td>日付（古い順）、損益（低い順）</td> </tr> <tr> <td>2</td> <td>日付（新しい順）、損益（高い順）</td> </tr> <tr> <td>3</td> <td>日付（新しい順）、損益（低い順）</td> </tr> <tr> <td>4</td> <td>損益（高い順）、日付（古い順）</td> </tr> <tr> <td>5</td> <td>損益（高い順）、日付（新しい順）</td> </tr> <tr> <td>6</td> <td>損益（低い順）、日付（古い順）</td> </tr> <tr> <td>7</td> <td>損益（低い順）、日付（新しい順）</td> </tr> </tbody> </table>
	ClosePositionOrder int `json:"ClosePositionOrder"`
	// 返済建玉指定<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。
	ClosePositions []struct {
		HoldID string `json:"HoldID,omitempty"` // 返済建玉ID
		Qty    int    `json:"Qty,omitempty"`    // 返済建玉数量
	} `json:"ClosePositions"`
	// 執行条件 <table> <thead> <tr> <th>定義値</th> <th>説明</th> <th>”Price”の指定</th> </tr> </thead> <tbody> <tr> <td>18</td> <td>引成（派生）<br>※TimeInForceは、「FAK」のみ有効</td> <td>0</td> </tr> <tr> <td>20</td> <td>指値</td> <td>発注したい金額</td> </tr> <tr> <td>28</td> <td>引指（派生）<br>※TimeInForceは、「FAS」のみ有効</td> <td>発注したい金額</td> </tr> <tr> <td>30</td> <td>逆指値</td> <td>指定なし<br>※AfterHitPriceで指定ください</td> </tr> <tr> <td>120</td> <td>成行（マーケットオーダー）</td> <td>0</td> </tr> </tbody> </table>
	FrontOrderType int `json:"FrontOrderType" validate:"required"`
	// 注文価格<br>※FrontOrderTypeで成行を指定した場合、0を指定する。<br>※詳細について、”FrontOrderType”をご確認ください。
	Price float64 `json:"Price" validate:"required"`
	// 注文有効期限<br> yyyyMMdd形式。<br> 「0」を指定すると、kabuステーション上の発注画面の「本日」に対応する日付として扱います。<br> 「本日」は直近の注文可能日となり、以下のように設定されます。<br> その市場の引けまでの間 : 当日<br> その市場の引け後 : 翌取引所営業日<br> その市場の休前日 : 休日明けの取引所営業日<br> ※ 日替わりはkabuステーションが日付変更通知を受信したタイミングです。<br> ※ 日通しの場合、夜間取引の引け後に日付が更新されます。
	ExpireDay int `json:"ExpireDay" validate:"required"`
	// 逆指値条件<br> ※FrontOrderTypeで逆指値を指定した場合のみ必須。
	ReverseLimitOrder struct {
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // トリガ価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。
		UnderOver         int     `json:"UnderOver,omitempty"`         // 以上／以下<br> ※未設定の場合はエラーになります。<br> ※1、2以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>以下</td> </tr> <tr> <td>2</td> <td>以上</td> </tr> </tbody> </table>
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // ヒット後執行条件<br> ※未設定の場合はエラーになります。<br> ※日通の注文で2以外が指定された場合はエラーになります。<br> ※日中、夜間の注文で1、2以外が指定された場合はエラーになります。<br> ※逆指値（成行）で有効期間条件(TimeInForce)にFAK以外を指定された場合はエラーになります。<br> ※逆指値（指値）で有効期間条件(TimeInForce)にFAS以外を指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>成行</td> </tr> <tr> <td>2</td> <td>指値</td> </tr> </tbody> </table>
		AfterHitPrice     float64 `json:"AfterHitPrice,omitempty"`     // ヒット後注文価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。<br><br> ヒット後執行条件に従い、下記のようにヒット後注文価格を設定してください。 <table> <thead> <tr> <th>ヒット後執行条件</th> <th>設定価格</th> </tr> </thead> <tbody> <tr> <td>成行</td> <td>0</td> </tr> <tr> <td>指値</td> <td>指値の単価</td> </tr> </tbody> </table>
	} `json:"ReverseLimitOrder"`
}

// ResPostOrderSendorderFuture は **POST /sendorder/future** のレスポンス。
//
// ### 概要
// 注文発注（先物）
type ResPostOrderSendorderFuture struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PostOrderSendorderFuture は **POST /sendorder/future** を呼び出します。
//
// ### 概要
// 注文発注（先物）
//
// 先物銘柄の注文を発注します。<br> 同一の銘柄に対しての注文は同時に5件ほどを上限としてご利用ください。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PostOrderSendorderFuture(ReqPostOrderSendorderFuture{ /* TODO: フィールドを設定 */ })
func PostOrderSendorderFuture(req ReqPostOrderSendorderFuture) (code int, res ResPostOrderSendorderFuture, err error) {
	p := "/sendorder/future"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
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

// ReqPostOrderSendorderOption は **POST /sendorder/option** のリクエスト。
//
// ### 概要
// 注文発注（オプション）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPostOrderSendorderOption struct {
	// 銘柄コード<br>※取引最終日に「オプション銘柄コード取得」でDerivMonthに0（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	Symbol string `json:"Symbol" validate:"required"`
	// 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange" validate:"required"`
	// 取引区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>新規</td> </tr> <tr> <td>2</td> <td>返済</td> </tr> </tbody> </table>
	TradeType int `json:"TradeType" validate:"required"`
	// 有効期間条件 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>FAS</td> </tr> <tr> <td>2</td> <td>FAK</td> </tr> <tr> <td>3</td> <td>FOK</td> </tr> </tbody> </table> ※執行条件(FrontOrderType)、有効期限条件(TimeInForce)、市場コード(Exchange)で選択できる組み合わせは下表のようになります。 <table> <thead> <tr> <th rowspan="2">執行条件</th> <th rowspan="2">有効期間条件</th> <th colspan="3">市場コード</th> </tr> <tr> <th>日中</th> <th>夜間</th> <th>日通し</th> </tr> </thead> <tbody> <tr> <td>指値</td> <td>FAS</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>指値</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>指値</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>逆指値（指値）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>逆指値（成行）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引成</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引指</td> <td>FAS</td> <td>●</td> <td>●</td> <td>-</td> </tr> </tbody> </table>
	TimeInForce int `json:"TimeInForce" validate:"required"`
	// 売買区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"Side" validate:"required"`
	// 注文数量
	Qty int `json:"Qty" validate:"required"`
	// 決済順序<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>日付（古い順）、損益（高い順）</td> </tr> <tr> <td>1</td> <td>日付（古い順）、損益（低い順）</td> </tr> <tr> <td>2</td> <td>日付（新しい順）、損益（高い順）</td> </tr> <tr> <td>3</td> <td>日付（新しい順）、損益（低い順）</td> </tr> <tr> <td>4</td> <td>損益（高い順）、日付（古い順）</td> </tr> <tr> <td>5</td> <td>損益（高い順）、日付（新しい順）</td> </tr> <tr> <td>6</td> <td>損益（低い順）、日付（古い順）</td> </tr> <tr> <td>7</td> <td>損益（低い順）、日付（新しい順）</td> </tr> </tbody> </table>
	ClosePositionOrder int `json:"ClosePositionOrder"`
	// 返済建玉指定<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。
	ClosePositions []struct {
		HoldID string `json:"HoldID,omitempty"` // 返済建玉ID
		Qty    int    `json:"Qty,omitempty"`    // 返済建玉数量
	} `json:"ClosePositions"`
	// 執行条件 <table> <thead> <tr> <th>定義値</th> <th>説明</th> <th>”Price”の指定</th> </tr> </thead> <tbody> <tr> <td>18</td> <td>引成（派生）<br>※TimeInForceは、「FAK」のみ有効</td> <td>0</td> </tr> <tr> <td>20</td> <td>指値</td> <td>発注したい金額</td> </tr> <tr> <td>28</td> <td>引指（派生）<br>※TimeInForceは、「FAS」のみ有効</td> <td>発注したい金額</td> </tr> <tr> <td>30</td> <td>逆指値</td> <td>指定なし<br>※AfterHitPriceで指定ください</td> </tr> <tr> <td>120</td> <td>成行（マーケットオーダー）</td> <td>0</td> </tr> </tbody> </table>
	FrontOrderType int `json:"FrontOrderType" validate:"required"`
	// 注文価格<br>※FrontOrderTypeで成行を指定した場合、0を指定する。<br>※詳細について、”FrontOrderType”をご確認ください。
	Price float64 `json:"Price" validate:"required"`
	// 注文有効期限<br> yyyyMMdd形式。<br> 「0」を指定すると、kabuステーション上の発注画面の「本日」に対応する日付として扱います。<br> 「本日」は直近の注文可能日となり、以下のように設定されます。<br> その市場の引けまでの間 : 当日<br> その市場の引け後 : 翌取引所営業日<br> その市場の休前日 : 休日明けの取引所営業日<br> ※ 日替わりはkabuステーションが日付変更通知を受信したタイミングです。<br> ※ 日通しの場合、夜間取引の引け後に日付が更新されます。
	ExpireDay int `json:"ExpireDay" validate:"required"`
	// 逆指値条件<br> ※FrontOrderTypeで逆指値を指定した場合のみ必須。
	ReverseLimitOrder struct {
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // トリガ価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。
		UnderOver         int     `json:"UnderOver,omitempty"`         // 以上／以下<br> ※未設定の場合はエラーになります。<br> ※1、2以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>以下</td> </tr> <tr> <td>2</td> <td>以上</td> </tr> </tbody> </table>
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // ヒット後執行条件<br> ※未設定の場合はエラーになります。<br> ※日通の注文で2以外が指定された場合はエラーになります。<br> ※日中、夜間の注文で1、2以外が指定された場合はエラーになります。<br> ※逆指値（成行）で有効期間条件(TimeInForce)にFAK以外を指定された場合はエラーになります。<br> ※逆指値（指値）で有効期間条件(TimeInForce)にFAS以外を指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>成行</td> </tr> <tr> <td>2</td> <td>指値</td> </tr> </tbody> </table>
		AfterHitPrice     float64 `json:"AfterHitPrice,omitempty"`     // ヒット後注文価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。<br><br> ヒット後執行条件に従い、下記のようにヒット後注文価格を設定してください。 <table> <thead> <tr> <th>ヒット後執行条件</th> <th>設定価格</th> </tr> </thead> <tbody> <tr> <td>成行</td> <td>0</td> </tr> <tr> <td>指値</td> <td>指値の単価</td> </tr> </tbody> </table>
	} `json:"ReverseLimitOrder"`
}

// ResPostOrderSendorderOption は **POST /sendorder/option** のレスポンス。
//
// ### 概要
// 注文発注（オプション）
type ResPostOrderSendorderOption struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PostOrderSendorderOption は **POST /sendorder/option** を呼び出します。
//
// ### 概要
// 注文発注（オプション）
//
// オプション銘柄の注文を発注します。<br> 同一の銘柄に対しての注文は同時に5件ほどを上限としてご利用ください。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PostOrderSendorderOption(ReqPostOrderSendorderOption{ /* TODO: フィールドを設定 */ })
func PostOrderSendorderOption(req ReqPostOrderSendorderOption) (code int, res ResPostOrderSendorderOption, err error) {
	p := "/sendorder/option"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
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

// ReqPutOrderCancelorder は **PUT /cancelorder** のリクエスト。
//
// ### 概要
// 注文取消
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutOrderCancelorder struct {
	// 注文番号<br>sendorderのレスポンスで受け取るOrderID。
	OrderId string `json:"OrderId" validate:"required"`
}

// ResPutOrderCancelorder は **PUT /cancelorder** のレスポンス。
//
// ### 概要
// 注文取消
type ResPutOrderCancelorder struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PutOrderCancelorder は **PUT /cancelorder** を呼び出します。
//
// ### 概要
// 注文取消
//
// 注文を取消します
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := PutOrderCancelorder(ReqPutOrderCancelorder{ /* TODO: フィールドを設定 */ })
func PutOrderCancelorder(req ReqPutOrderCancelorder) (code int, res ResPutOrderCancelorder, err error) {
	p := "/cancelorder"
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
