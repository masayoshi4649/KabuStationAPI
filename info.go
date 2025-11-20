package kabusapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ReqGetInfoBoardSymbol は **GET /board/{symbol}** のリクエスト。
//
// ### 概要
// 時価情報・板情報
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoBoardSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。<br> ※SOR市場は取扱っておりませんのでご注意ください。<b>市場コード</b><br> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoBoardSymbol は **GET /board/{symbol}** のレスポンス。
//
// ### 概要
// 時価情報・板情報
type ResGetInfoBoardSymbol struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名
	SymbolName string `json:"SymbolName,omitempty"`
	// 市場コード<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange,omitempty"`
	// 市場名称<br>※株式・先物・オプション銘柄の場合のみ
	ExchangeName string `json:"ExchangeName,omitempty"`
	// 現値
	CurrentPrice float64 `json:"CurrentPrice,omitempty"`
	// 現値時刻
	CurrentPriceTime string `json:"CurrentPriceTime,omitempty"`
	// 現値前値比較 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0056</td> <td>変わらず</td> </tr> <tr> <td>0057</td> <td>UP</td> </tr> <tr> <td>0058</td> <td>DOWN</td> </tr> <tr> <td>0059</td> <td>中断板寄り後の初値</td> </tr> <tr> <td>0060</td> <td>ザラバ引け</td> </tr> <tr> <td>0061</td> <td>板寄り引け</td> </tr> <tr> <td>0062</td> <td>中断引け</td> </tr> <tr> <td>0063</td> <td>ダウン引け</td> </tr> <tr> <td>0064</td> <td>逆転終値</td> </tr> <tr> <td>0066</td> <td>特別気配引け</td> </tr> <tr> <td>0067</td> <td>一時留保引け</td> </tr> <tr> <td>0068</td> <td>売買停止引け</td> </tr> <tr> <td>0069</td> <td>サーキットブレーカ引け</td> </tr> <tr> <td>0431</td> <td>ダイナミックサーキットブレーカ引け</td> </tr> </tbody> </table>
	CurrentPriceChangeStatus string `json:"CurrentPriceChangeStatus,omitempty"`
	// 現値ステータス <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>現値</td> </tr> <tr> <td>2</td> <td>不連続歩み</td> </tr> <tr> <td>3</td> <td>板寄せ</td> </tr> <tr> <td>4</td> <td>システム障害</td> </tr> <tr> <td>5</td> <td>中断</td> </tr> <tr> <td>6</td> <td>売買停止</td> </tr> <tr> <td>7</td> <td>売買停止・システム停止解除</td> </tr> <tr> <td>8</td> <td>終値</td> </tr> <tr> <td>9</td> <td>システム停止</td> </tr> <tr> <td>10</td> <td>概算値</td> </tr> <tr> <td>11</td> <td>参考値</td> </tr> <tr> <td>12</td> <td>サーキットブレイク実施中</td> </tr> <tr> <td>13</td> <td>システム障害解除</td> </tr> <tr> <td>14</td> <td>サーキットブレイク解除</td> </tr> <tr> <td>15</td> <td>中断解除</td> </tr> <tr> <td>16</td> <td>一時留保中</td> </tr> <tr> <td>17</td> <td>一時留保解除</td> </tr> <tr> <td>18</td> <td>ファイル障害</td> </tr> <tr> <td>19</td> <td>ファイル障害解除</td> </tr> <tr> <td>20</td> <td>Spread/Strategy</td> </tr> <tr> <td>21</td> <td>ダイナミックサーキットブレイク発動</td> </tr> <tr> <td>22</td> <td>ダイナミックサーキットブレイク解除</td> </tr> <tr> <td>23</td> <td>板寄せ約定</td> </tr> </tbody> </table>
	CurrentPriceStatus int `json:"CurrentPriceStatus,omitempty"`
	// 計算用現値
	CalcPrice float64 `json:"CalcPrice,omitempty"`
	// 前日終値
	PreviousClose float64 `json:"PreviousClose,omitempty"`
	// 前日終値日付
	PreviousCloseTime string `json:"PreviousCloseTime,omitempty"`
	// 前日比
	ChangePreviousClose float64 `json:"ChangePreviousClose,omitempty"`
	// 騰落率
	ChangePreviousClosePer float64 `json:"ChangePreviousClosePer,omitempty"`
	// 始値
	OpeningPrice float64 `json:"OpeningPrice,omitempty"`
	// 始値時刻
	OpeningPriceTime string `json:"OpeningPriceTime,omitempty"`
	// 高値
	HighPrice float64 `json:"HighPrice,omitempty"`
	// 高値時刻
	HighPriceTime string `json:"HighPriceTime,omitempty"`
	// 安値
	LowPrice float64 `json:"LowPrice,omitempty"`
	// 安値時刻
	LowPriceTime string `json:"LowPriceTime,omitempty"`
	// 売買高<br>※株式・先物・オプション銘柄の場合のみ
	TradingVolume float64 `json:"TradingVolume,omitempty"`
	// 売買高時刻<br>※株式・先物・オプション銘柄の場合のみ
	TradingVolumeTime string `json:"TradingVolumeTime,omitempty"`
	// 売買高加重平均価格（VWAP）<br>※株式・先物・オプション銘柄の場合のみ
	VWAP float64 `json:"VWAP,omitempty"`
	// 売買代金<br>※株式・先物・オプション銘柄の場合のみ
	TradingValue float64 `json:"TradingValue,omitempty"`
	// 最良売気配数量 ※①<br>※株式・先物・オプション銘柄の場合のみ
	BidQty float64 `json:"BidQty,omitempty"`
	// 最良売気配値段 ※①<br>※株式・先物・オプション銘柄の場合のみ
	BidPrice float64 `json:"BidPrice,omitempty"`
	// 最良売気配時刻 ※①<br>※株式銘柄の場合のみ
	BidTime string `json:"BidTime,omitempty"`
	// 最良売気配フラグ ※①<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0101</td> <td>一般気配</td> </tr> <tr> <td>0102</td> <td>特別気配</td> </tr> <tr> <td>0103</td> <td>注意気配</td> </tr> <tr> <td>0107</td> <td>寄前気配</td> </tr> <tr> <td>0108</td> <td>停止前特別気配</td> </tr> <tr> <td>0109</td> <td>引け後気配</td> </tr> <tr> <td>0116</td> <td>寄前気配約定成立ポイントなし</td> </tr> <tr> <td>0117</td> <td>寄前気配約定成立ポイントあり</td> </tr> <tr> <td>0118</td> <td>連続約定気配</td> </tr> <tr> <td>0119</td> <td>停止前の連続約定気配</td> </tr> <tr> <td>0120</td> <td>買い上がり売り下がり中</td> </tr> </tbody> </table>
	BidSign string `json:"BidSign,omitempty"`
	// 売成行数量<br>※株式銘柄の場合のみ
	MarketOrderSellQty float64 `json:"MarketOrderSellQty,omitempty"`
	// 売気配数量1本目
	Sell1 struct {
		Time  string  `json:"Time,omitempty"`  // 時刻<br>※株式銘柄の場合のみ
		Sign  string  `json:"Sign,omitempty"`  // 気配フラグ<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0101</td> <td>一般気配</td> </tr> <tr> <td>0102</td> <td>特別気配</td> </tr> <tr> <td>0103</td> <td>注意気配</td> </tr> <tr> <td>0107</td> <td>寄前気配</td> </tr> <tr> <td>0108</td> <td>停止前特別気配</td> </tr> <tr> <td>0109</td> <td>引け後気配</td> </tr> <tr> <td>0116</td> <td>寄前気配約定成立ポイントなし</td> </tr> <tr> <td>0117</td> <td>寄前気配約定成立ポイントあり</td> </tr> <tr> <td>0118</td> <td>連続約定気配</td> </tr> <tr> <td>0119</td> <td>停止前の連続約定気配</td> </tr> <tr> <td>0120</td> <td>買い上がり売り下がり中</td> </tr> </tbody> </table>
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell1,omitempty"`
	// 売気配数量2本目
	Sell2 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell2,omitempty"`
	// 売気配数量3本目
	Sell3 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell3,omitempty"`
	// 売気配数量4本目
	Sell4 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell4,omitempty"`
	// 売気配数量5本目
	Sell5 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell5,omitempty"`
	// 売気配数量6本目
	Sell6 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell6,omitempty"`
	// 売気配数量7本目
	Sell7 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell7,omitempty"`
	// 売気配数量8本目
	Sell8 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell8,omitempty"`
	// 売気配数量9本目
	Sell9 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell9,omitempty"`
	// 売気配数量10本目
	Sell10 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell10,omitempty"`
	// 最良買気配数量 ※①<br>※株式・先物・オプション銘柄の場合のみ
	AskQty float64 `json:"AskQty,omitempty"`
	// 最良買気配値段 ※①<br>※株式・先物・オプション銘柄の場合のみ
	AskPrice float64 `json:"AskPrice,omitempty"`
	// 最良買気配時刻 ※①<br>※株式銘柄の場合のみ
	AskTime string `json:"AskTime,omitempty"`
	// 最良買気配フラグ ※①<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0101</td> <td>一般気配</td> </tr> <tr> <td>0102</td> <td>特別気配</td> </tr> <tr> <td>0103</td> <td>注意気配</td> </tr> <tr> <td>0107</td> <td>寄前気配</td> </tr> <tr> <td>0108</td> <td>停止前特別気配</td> </tr> <tr> <td>0109</td> <td>引け後気配</td> </tr> <tr> <td>0116</td> <td>寄前気配約定成立ポイントなし</td> </tr> <tr> <td>0117</td> <td>寄前気配約定成立ポイントあり</td> </tr> <tr> <td>0118</td> <td>連続約定気配</td> </tr> <tr> <td>0119</td> <td>停止前の連続約定気配</td> </tr> <tr> <td>0120</td> <td>買い上がり売り下がり中</td> </tr> </tbody> </table>
	AskSign string `json:"AskSign,omitempty"`
	// 買成行数量<br>※株式銘柄の場合のみ
	MarketOrderBuyQty float64 `json:"MarketOrderBuyQty,omitempty"`
	// 買気配数量1本目
	Buy1 struct {
		Time  string  `json:"Time,omitempty"`  // 時刻<br>※株式銘柄の場合のみ
		Sign  string  `json:"Sign,omitempty"`  // 気配フラグ<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0101</td> <td>一般気配</td> </tr> <tr> <td>0102</td> <td>特別気配</td> </tr> <tr> <td>0103</td> <td>注意気配</td> </tr> <tr> <td>0107</td> <td>寄前気配</td> </tr> <tr> <td>0108</td> <td>停止前特別気配</td> </tr> <tr> <td>0109</td> <td>引け後気配</td> </tr> <tr> <td>0116</td> <td>寄前気配約定成立ポイントなし</td> </tr> <tr> <td>0117</td> <td>寄前気配約定成立ポイントあり</td> </tr> <tr> <td>0118</td> <td>連続約定気配</td> </tr> <tr> <td>0119</td> <td>停止前の連続約定気配</td> </tr> <tr> <td>0120</td> <td>買い上がり売り下がり中</td> </tr> </tbody> </table>
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy1,omitempty"`
	// 買気配数量2本目
	Buy2 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy2,omitempty"`
	// 買気配数量3本目
	Buy3 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy3,omitempty"`
	// 買気配数量4本目
	Buy4 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy4,omitempty"`
	// 買気配数量5本目
	Buy5 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy5,omitempty"`
	// 買気配数量6本目
	Buy6 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy6,omitempty"`
	// 買気配数量7本目
	Buy7 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy7,omitempty"`
	// 買気配数量8本目
	Buy8 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy8,omitempty"`
	// 買気配数量9本目
	Buy9 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy9,omitempty"`
	// 買気配数量10本目
	Buy10 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy10,omitempty"`
	// OVER気配数量<br>※株式銘柄の場合のみ
	OverSellQty float64 `json:"OverSellQty,omitempty"`
	// UNDER気配数量<br>※株式銘柄の場合のみ
	UnderBuyQty float64 `json:"UnderBuyQty,omitempty"`
	// 時価総額<br>※株式銘柄の場合のみ
	TotalMarketValue float64 `json:"TotalMarketValue,omitempty"`
	// 清算値<br>※先物銘柄の場合のみ
	ClearingPrice float64 `json:"ClearingPrice,omitempty"`
	// インプライド・ボラティリティ<br>※オプション銘柄かつ日通しの場合のみ
	IV float64 `json:"IV,omitempty"`
	// ガンマ<br>※オプション銘柄かつ日通しの場合のみ
	Gamma float64 `json:"Gamma,omitempty"`
	// セータ<br>※オプション銘柄かつ日通しの場合のみ
	Theta float64 `json:"Theta,omitempty"`
	// ベガ<br>※オプション銘柄かつ日通しの場合のみ
	Vega float64 `json:"Vega,omitempty"`
	// デルタ<br>※オプション銘柄かつ日通しの場合のみ
	Delta float64 `json:"Delta,omitempty"`
	// 銘柄種別 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>指数</td> </tr> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>101</td> <td>日経225先物</td> </tr> <tr> <td>103</td> <td>日経225OP</td> </tr> <tr> <td>107</td> <td>TOPIX先物</td> </tr> <tr> <td>121</td> <td>JPX400先物</td> </tr> <tr> <td>144</td> <td>NYダウ</td> </tr> <tr> <td>145</td> <td>日経平均VI</td> </tr> <tr> <td>154</td> <td>グロース250先物</td> </tr> <tr> <td>155</td> <td>TOPIX_REIT</td> </tr> <tr> <td>171</td> <td>TOPIX CORE30</td> </tr> <tr> <td>901</td> <td>日経平均225ミニ先物</td> </tr> <tr> <td>907</td> <td>TOPIXミニ先物</td> </tr> </tbody> </table>
	SecurityType int `json:"SecurityType,omitempty"`
}

// GetInfoBoardSymbol は **GET /board/{symbol}** を呼び出します。
//
// ### 概要
// 時価情報・板情報
//
// 指定した銘柄の時価情報・板情報を取得します<br> レスポンスの一部にnullが発生した場合、該当銘柄を銘柄登録をしてから、 <br>再度時価情報・板情報APIを実行してください。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoBoardSymbol(ReqGetInfoBoardSymbol{ /* TODO: フィールドを設定 */ })
func GetInfoBoardSymbol(req ReqGetInfoBoardSymbol) (code int, res ResGetInfoBoardSymbol, err error) {
	p := "/board/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoSymbolSymbol は **GET /symbol/{symbol}** のリクエスト。
//
// ### 概要
// 銘柄情報
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。<br> ※SOR市場は取扱っておりませんのでご注意ください。<b>市場コード</b><br> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
	// 追加情報出力フラグ（未指定時：true）<br> ※追加情報は、「時価総額」、「発行済み株式数」、「決算期日」、「清算値」を意味します。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>true</td> <td>追加情報を出力する</td> </tr> <tr> <td>false</td> <td>追加情報を出力しない</td> </tr> </tbody> </table>
	Addinfo string `json:"-" query:"addinfo"`
}

// ResGetInfoSymbolSymbol は **GET /symbol/{symbol}** のレスポンス。
//
// ### 概要
// 銘柄情報
type ResGetInfoSymbolSymbol struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名
	SymbolName string `json:"SymbolName,omitempty"`
	// 銘柄略称<br>※株式・先物・オプション銘柄の場合のみ
	DisplayName string `json:"DisplayName,omitempty"`
	// 市場コード<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange,omitempty"`
	// 市場名称<br>※株式・先物・オプション銘柄の場合のみ
	ExchangeName string `json:"ExchangeName,omitempty"`
	// 業種コード名<br>※株式銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0050</td> <td>水産・農林業</td> </tr> <tr> <td>1050</td> <td>鉱業</td> </tr> <tr> <td>2050</td> <td>建設業</td> </tr> <tr> <td>3050</td> <td>食料品</td> </tr> <tr> <td>3100</td> <td>繊維製品</td> </tr> <tr> <td>3150</td> <td>パルプ・紙</td> </tr> <tr> <td>3200</td> <td>化学</td> </tr> <tr> <td>3250</td> <td>医薬品</td> </tr> <tr> <td>3300</td> <td>石油・石炭製品</td> </tr> <tr> <td>3350</td> <td>ゴム製品</td> </tr> <tr> <td>3400</td> <td>ガラス・土石製品</td> </tr> <tr> <td>3450</td> <td>鉄鋼</td> </tr> <tr> <td>3500</td> <td>非鉄金属</td> </tr> <tr> <td>3550</td> <td>金属製品</td> </tr> <tr> <td>3600</td> <td>機械</td> </tr> <tr> <td>3650</td> <td>電気機器</td> </tr> <tr> <td>3700</td> <td>輸送用機器</td> </tr> <tr> <td>3750</td> <td>精密機器</td> </tr> <tr> <td>3800</td> <td>その他製品</td> </tr> <tr> <td>4050</td> <td>電気・ガス業</td> </tr> <tr> <td>5050</td> <td>陸運業</td> </tr> <tr> <td>5100</td> <td>海運業</td> </tr> <tr> <td>5150</td> <td>空運業</td> </tr> <tr> <td>5200</td> <td>倉庫・運輸関連業</td> </tr> <tr> <td>5250</td> <td>情報・通信業</td> </tr> <tr> <td>6050</td> <td>卸売業</td> </tr> <tr> <td>6100</td> <td>小売業</td> </tr> <tr> <td>7050</td> <td>銀行業</td> </tr> <tr> <td>7100</td> <td>証券、商品先物取引業</td> </tr> <tr> <td>7150</td> <td>保険業</td> </tr> <tr> <td>7200</td> <td>その他金融業</td> </tr> <tr> <td>8050</td> <td>不動産業</td> </tr> <tr> <td>9050</td> <td>サービス業</td> </tr> <tr> <td>9999</td> <td>その他</td> </tr> </tbody> </table>
	BisCategory string `json:"BisCategory,omitempty"`
	// 時価総額<br>※株式銘柄の場合のみ<br>追加情報出力フラグ：falseの場合、null
	TotalMarketValue float64 `json:"TotalMarketValue,omitempty"`
	// 発行済み株式数（千株）<br>※株式銘柄の場合のみ<br>追加情報出力フラグ：falseの場合、null
	TotalStocks float64 `json:"TotalStocks,omitempty"`
	// 売買単位<br>※株式・先物・オプション銘柄の場合のみ
	TradingUnit float64 `json:"TradingUnit,omitempty"`
	// 決算期日<br>※株式銘柄の場合のみ<br>追加情報出力フラグ：falseの場合、null
	FiscalYearEndBasic int `json:"FiscalYearEndBasic,omitempty"`
	// 呼値グループ<br> ※株式・先物・オプション銘柄の場合のみ<br> ※各呼値コードが対応する商品は以下となります。<BR> 株式の呼値の単位の詳細は [JPXページ](https://www.jpx.co.jp/equities/trading/domestic/07.html) をご覧ください。<BR> 10000：株式(通常の呼値単位の銘柄) <br> 10003：株式(TOPIX500構成銘柄※売買単位が10口以上のETF等含む)<br> 10004：株式(売買単位が1口のETF等)<br> 10118 : 日経平均先物<br> 10119 : 日経225mini<br> 10318 : 日経平均オプション<br> 10706 : ﾐﾆTOPIX先物<br> 10718 : TOPIX先物<br> 12122 : JPX日経400指数先物<br> 14473 : NYダウ先物<br> 14515 : 日経平均VI先物<br> 15411 : グロース250先物<br> 15569 : 東証REIT指数先物<br> 17163 : TOPIXCore30指数先物<br> <table> <thead> <tr> <th>呼値コード</th> <th>値段の水準</th> <th>呼値単位</th> </tr> </thead> <tbody> <tr> <td>10000</td> <td>3000円以下</td> <td>1</td> </tr> <tr> <td>10000</td> <td>5000円以下</td> <td>5</td> </tr> <tr> <td>10000</td> <td>30000円以下</td> <td>10</td> </tr> <tr> <td>10000</td> <td>50000円以下</td> <td>50</td> </tr> <tr> <td>10000</td> <td>300000円以下</td> <td>100</td> </tr> <tr> <td>10000</td> <td>500000円以下</td> <td>500</td> </tr> <tr> <td>10000</td> <td>3000000円以下</td> <td>1000</td> </tr> <tr> <td>10000</td> <td>5000000円以下</td> <td>5000</td> </tr> <tr> <td>10000</td> <td>30000000円以下</td> <td>10000</td> </tr> <tr> <td>10000</td> <td>50000000円以下</td> <td>50000</td> </tr> <tr> <td>10000</td> <td>50000000円超</td> <td>100000</td> </tr> <tr> <td>10003</td> <td>1000円以下</td> <td>0.1</td> </tr> <tr> <td>10003</td> <td>3000円以下</td> <td>0.5</td> </tr> <tr> <td>10003</td> <td>10000円以下</td> <td>1</td> </tr> <tr> <td>10003</td> <td>30000円以下</td> <td>5</td> </tr> <tr> <td>10003</td> <td>100000円以下</td> <td>10</td> </tr> <tr> <td>10003</td> <td>300000円以下</td> <td>50</td> </tr> <tr> <td>10003</td> <td>1000000円以下</td> <td>100</td> </tr> <tr> <td>10003</td> <td>3000000円以下</td> <td>500</td> </tr> <tr> <td>10003</td> <td>10000000円以下</td> <td>1000</td> </tr> <tr> <td>10003</td> <td>30000000円以下</td> <td>5000</td> </tr> <tr> <td>10003</td> <td>30000000円超</td> <td>10000</td> </tr> <tr> <td>10004</td> <td>10000円以下</td> <td>1</td> </tr> <tr> <td>10004</td> <td>30000円以下</td> <td>5</td> </tr> <tr> <td>10004</td> <td>100000円以下</td> <td>10</td> </tr> <tr> <td>10004</td> <td>300000円以下</td> <td>50</td> </tr> <tr> <td>10004</td> <td>1000000円以下</td> <td>100</td> </tr> <tr> <td>10004</td> <td>3000000円以下</td> <td>500</td> </tr> <tr> <td>10004</td> <td>10000000円以下</td> <td>1000</td> </tr> <tr> <td>10004</td> <td>30000000円以下</td> <td>5000</td> </tr> <tr> <td>10004</td> <td>30000000円超</td> <td>10000</td> </tr> <tr> <td>10118</td> <td>-</td> <td>10</td> </tr> <tr> <td>10119</td> <td>-</td> <td>5</td> </tr> <tr> <td>10318</td> <td>100円以下</td> <td>1</td> </tr> <tr> <td>10318</td> <td>1000円以下</td> <td>5</td> </tr> <tr> <td>10318</td> <td>1000円超</td> <td>10</td> </tr> <tr> <td>10706</td> <td>-</td> <td>0.25</td> </tr> <tr> <td>10718</td> <td>-</td> <td>0.5</td> </tr> <tr> <td>12122</td> <td>-</td> <td>5</td> </tr> <tr> <td>14473</td> <td>-</td> <td>1</td> </tr> <tr> <td>14515</td> <td>-</td> <td>0.05</td> </tr> <tr> <td>15411</td> <td>-</td> <td>1</td> </tr> <tr> <td>15569</td> <td>-</td> <td>0.5</td> </tr> <tr> <td>17163</td> <td>-</td> <td>0.5</td> </tr> </tbody> </table>
	PriceRangeGroup string `json:"PriceRangeGroup,omitempty"`
	// 一般信用買建フラグ<br>※trueのとき、一般信用(長期)または一般信用(デイトレ)が買建可能<br>※株式銘柄の場合のみ
	KCMarginBuy bool `json:"KCMarginBuy,omitempty"`
	// 一般信用売建フラグ<br>※trueのとき、一般信用(長期)または一般信用(デイトレ)が売建可能<br>※株式銘柄の場合のみ
	KCMarginSell bool `json:"KCMarginSell,omitempty"`
	// 制度信用買建フラグ<br>※trueのとき制度信用買建可能<br>※株式銘柄の場合のみ
	MarginBuy bool `json:"MarginBuy,omitempty"`
	// 制度信用売建フラグ<br>※trueのとき制度信用売建可能<br>※株式銘柄の場合のみ
	MarginSell bool `json:"MarginSell,omitempty"`
	// 値幅上限<br>※株式・先物・オプション銘柄の場合のみ
	UpperLimit float64 `json:"UpperLimit,omitempty"`
	// 値幅下限<br>※株式・先物・オプション銘柄の場合のみ
	LowerLimit float64 `json:"LowerLimit,omitempty"`
	// 原資産コード<br>※先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>NK225</td> <td>日経225</td> </tr> <tr> <td>NK300</td> <td>日経300</td> </tr> <tr> <td>GROWTH</td> <td>グロース250先物</td> </tr> <tr> <td>JPX400</td> <td>JPX日経400</td> </tr> <tr> <td>TOPIX</td> <td>TOPIX</td> </tr> <tr> <td>NKVI</td> <td>日経平均VI</td> </tr> <tr> <td>DJIA</td> <td>NYダウ</td> </tr> <tr> <td>TSEREITINDEX</td> <td>東証REIT指数</td> </tr> <tr> <td>TOPIXCORE30</td> <td>TOPIX Core30</td> </tr> </tbody> </table>
	Underlyer string `json:"Underlyer,omitempty"`
	// 限月-年月<br>※「限月-年月」は「年(yyyy)/月(MM)」で表示します。<br>※先物・オプション銘柄の場合のみ
	DerivMonth string `json:"DerivMonth,omitempty"`
	// 取引開始日<br>※先物・オプション銘柄の場合のみ
	TradeStart int `json:"TradeStart,omitempty"`
	// 取引終了日<br>※先物・オプション銘柄の場合のみ
	TradeEnd int `json:"TradeEnd,omitempty"`
	// 権利行使価格<br>※オプション銘柄の場合のみ
	StrikePrice float64 `json:"StrikePrice,omitempty"`
	// プット/コール区分<br>※オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>プット</td> </tr> <tr> <td>2</td> <td>コール</td> </tr> </tbody> </table>
	PutOrCall int `json:"PutOrCall,omitempty"`
	// 清算値<br>※先物銘柄の場合のみ<br>追加情報出力フラグ：falseの場合、null
	ClearingPrice float64 `json:"ClearingPrice,omitempty"`
}

// GetInfoSymbolSymbol は **GET /symbol/{symbol}** を呼び出します。
//
// ### 概要
// 銘柄情報
//
// 指定した銘柄情報を取得します
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoSymbolSymbol(ReqGetInfoSymbolSymbol{ /* TODO: フィールドを設定 */ })
func GetInfoSymbolSymbol(req ReqGetInfoSymbolSymbol) (code int, res ResGetInfoSymbolSymbol, err error) {
	p := "/symbol/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Addinfo != "" {
		v.Set("addinfo", fmt.Sprint(req.Addinfo))
	}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoOrders は **GET /orders** のリクエスト。
//
// ### 概要
// 注文約定照会
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoOrders struct {
	// 取得する商品 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>すべて </td> </tr> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>2</td> <td>信用</td> </tr> <tr> <td>3</td> <td>先物</td> </tr> <tr> <td>4</td> <td>OP</td> </tr> </tbody> </table>
	Product string `json:"-" query:"product"`
	// 注文番号<br> ※指定された注文番号と一致する注文のみレスポンスします。<br> ※指定された注文番号との比較では大文字小文字を区別しません。<br> ※複数の注文番号を指定することはできません。
	ID string `json:"-" query:"id"`
	// 更新日時<br> ※形式：yyyyMMddHHmmss （例：20201201123456）<br> ※指定された更新日時以降（指定日時含む）に更新された注文のみレスポンスします。<br> ※複数の更新日時を指定することはできません。
	Updtime string `json:"-" query:"updtime"`
	// 注文詳細抑止 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>true</td> <td>注文詳細を出力する（デフォルト）</td> </tr> <tr> <td>false</td> <td>注文詳細の出力しない</td> </tr> </tbody> </table>
	Details string `json:"-" query:"details"`
	// 銘柄コード<br>※指定された銘柄コードと一致する注文のみレスポンスします。<br>※複数の銘柄コードを指定することができません。
	Symbol string `json:"-" query:"symbol"`
	// 状態<br> ※指定された状態と一致する注文のみレスポンスします。<br> ※フィルタには数字の入力のみ受け付けます。<br> ※複数の状態を指定することはできません。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>待機（発注待機）</td> </tr> <tr> <td>2</td> <td>処理中（発注送信中）</td> </tr> <tr> <td>3</td> <td>処理済（発注済・訂正済）</td> </tr> <tr> <td>4</td> <td>訂正取消送信中</td> </tr> <tr> <td>5</td> <td>終了（発注エラー・取消済・全約定・失効・期限切れ）</td> </tr> </tbody> </table>
	State string `json:"-" query:"state"`
	// 売買区分<br> ※指定された売買区分と一致する注文のみレスポンスします。<br> ※フィルタには数字の入力のみ受け付けます。<br> ※複数の売買区分を指定することができません。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"-" query:"side"`
	// 取引区分<br> ※指定された取引区分と一致する注文のみレスポンスします。<br> ※フィルタには数字の入力のみ受け付けます。<br> ※複数の取引区分を指定することができません。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>新規</td> </tr> <tr> <td>3</td> <td>返済</td> </tr> </tbody> </table>
	Cashmargin string `json:"-" query:"cashmargin"`
}

// ResGetInfoOrders は **GET /orders** のレスポンス。
//
// ### 概要
// 注文約定照会
type ResGetInfoOrders struct {
}

// GetInfoOrders は **GET /orders** を呼び出します。
//
// ### 概要
// 注文約定照会
//
// 注文一覧を取得します。<br> ※下記Queryパラメータは任意設定となります。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoOrders(ReqGetInfoOrders{ /* TODO: フィールドを設定 */ })
func GetInfoOrders(req ReqGetInfoOrders) (code int, res ResGetInfoOrders, err error) {
	p := "/orders"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Product != "" {
		v.Set("product", fmt.Sprint(req.Product))
	}
	if req.ID != "" {
		v.Set("id", fmt.Sprint(req.ID))
	}
	if req.Updtime != "" {
		v.Set("updtime", fmt.Sprint(req.Updtime))
	}
	if req.Details != "" {
		v.Set("details", fmt.Sprint(req.Details))
	}
	if req.Symbol != "" {
		v.Set("symbol", fmt.Sprint(req.Symbol))
	}
	if req.State != "" {
		v.Set("state", fmt.Sprint(req.State))
	}
	if req.Side != "" {
		v.Set("side", fmt.Sprint(req.Side))
	}
	if req.Cashmargin != "" {
		v.Set("cashmargin", fmt.Sprint(req.Cashmargin))
	}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoPositions は **GET /positions** のリクエスト。
//
// ### 概要
// 残高照会
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoPositions struct {
	// 取得する商品 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>すべて</td> </tr> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>2</td> <td>信用</td> </tr> <tr> <td>3</td> <td>先物</td> </tr> <tr> <td>4</td> <td>OP</td> </tr> </tbody> </table>
	Product string `json:"-" query:"product"`
	// 銘柄コード<br>※指定された銘柄コードと一致するポジションのみレスポンスします。<br>※複数の銘柄コードを指定することはできません。
	Symbol string `json:"-" query:"symbol"`
	// 売買区分フィルタ<br> 指定された売買区分と一致する注文を返す <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"-" query:"side"`
	// 追加情報出力フラグ（未指定時：true）<br> ※追加情報は、「現在値」、「評価金額」、「評価損益額」、「評価損益率」を意味します。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>true</td> <td>追加情報を出力する</td> </tr> <tr> <td>false</td> <td>追加情報を出力しない</td> </tr> </tbody> </table>
	Addinfo string `json:"-" query:"addinfo"`
}

// ResGetInfoPositions は **GET /positions** のレスポンス。
//
// ### 概要
// 残高照会
type ResGetInfoPositions struct {
}

// GetInfoPositions は **GET /positions** を呼び出します。
//
// ### 概要
// 残高照会
//
// 残高一覧を取得します。<br>※下記Queryパラメータは任意設定となります。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoPositions(ReqGetInfoPositions{ /* TODO: フィールドを設定 */ })
func GetInfoPositions(req ReqGetInfoPositions) (code int, res ResGetInfoPositions, err error) {
	p := "/positions"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Product != "" {
		v.Set("product", fmt.Sprint(req.Product))
	}
	if req.Symbol != "" {
		v.Set("symbol", fmt.Sprint(req.Symbol))
	}
	if req.Side != "" {
		v.Set("side", fmt.Sprint(req.Side))
	}
	if req.Addinfo != "" {
		v.Set("addinfo", fmt.Sprint(req.Addinfo))
	}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoSymbolnameFuture は **GET /symbolname/future** のリクエスト。
//
// ### 概要
// 先物銘柄コード取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolnameFuture struct {
	// 先物コード<br> ※大文字小文字は区別しません。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>NK225</td> <td>日経平均先物</td> </tr> <tr> <td>NK225mini</td> <td>日経225mini先物</td> </tr> <tr> <td>TOPIX</td> <td>TOPIX先物</td> </tr> <tr> <td>TOPIXmini</td> <td>ミニTOPIX先物</td> </tr> <tr> <td>GROWTH</td> <td>グロース250先物</td> </tr> <tr> <td>JPX400</td> <td>JPX日経400先物</td> </tr> <tr> <td>DOW</td> <td>NYダウ先物</td> </tr> <tr> <td>VI</td> <td>日経平均VI先物</td> </tr> <tr> <td>Core30</td> <td>TOPIX Core30先物</td> </tr> <tr> <td>REIT</td> <td>東証REIT指数先物</td> </tr> <tr> <td>NK225micro</td> <td>日経225マイクロ先物</td> </tr> </tbody> </table>
	FutureCode string `json:"-" query:"FutureCode"`
	// 限月<br> ※限月はyyyyMM形式で指定します。0を指定した場合、直近限月となります。<br> ※取引最終日に「0」（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、 取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。<br>
	DerivMonth string `json:"-" query:"DerivMonth" required:"true"`
}

// ResGetInfoSymbolnameFuture は **GET /symbolname/future** のレスポンス。
//
// ### 概要
// 先物銘柄コード取得
type ResGetInfoSymbolnameFuture struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名称
	SymbolName string `json:"SymbolName,omitempty"`
}

// GetInfoSymbolnameFuture は **GET /symbolname/future** を呼び出します。
//
// ### 概要
// 先物銘柄コード取得
//
// 先物銘柄コード取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoSymbolnameFuture(ReqGetInfoSymbolnameFuture{ /* TODO: フィールドを設定 */ })
func GetInfoSymbolnameFuture(req ReqGetInfoSymbolnameFuture) (code int, res ResGetInfoSymbolnameFuture, err error) {
	p := "/symbolname/future"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	v.Set("FutureCode", fmt.Sprint(req.FutureCode))
	v.Set("DerivMonth", fmt.Sprint(req.DerivMonth))

	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoSymbolnameOption は **GET /symbolname/option** のリクエスト。
//
// ### 概要
// オプション銘柄コード取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolnameOption struct {
	// オプションコード<br> ※指定なしの場合は、日経225オプションを対象とする。<br> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>NK225op</td> <td>日経225オプション</td> </tr> <tr> <td>NK225miniop</td> <td>日経225ミニオプション</td> </tr> </tbody> </table>
	OptionCode string `json:"-" query:"OptionCode"`
	// 限月<br>※限月はyyyyMM形式で指定します。0を指定した場合、直近限月となります。<br>※取引最終日に「0」（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivMonth string `json:"-" query:"DerivMonth" required:"true"`
	// コール or プット<br> ※大文字小文字は区別しません。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>P</td> <td>PUT</td> </tr> <tr> <td>C</td> <td>CALL</td> </tr> </tbody> </table>
	PutOrCall string `json:"-" query:"PutOrCall" required:"true"`
	// 権利行使価格<br>※0を指定した場合、APIを実行した時点でのATMとなります。
	StrikePrice int `json:"-" query:"StrikePrice" required:"true"`
}

// ResGetInfoSymbolnameOption は **GET /symbolname/option** のレスポンス。
//
// ### 概要
// オプション銘柄コード取得
type ResGetInfoSymbolnameOption struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名称
	SymbolName string `json:"SymbolName,omitempty"`
}

// GetInfoSymbolnameOption は **GET /symbolname/option** を呼び出します。
//
// ### 概要
// オプション銘柄コード取得
//
// オプション銘柄コード取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoSymbolnameOption(ReqGetInfoSymbolnameOption{ /* TODO: フィールドを設定 */ })
func GetInfoSymbolnameOption(req ReqGetInfoSymbolnameOption) (code int, res ResGetInfoSymbolnameOption, err error) {
	p := "/symbolname/option"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.OptionCode != "" {
		v.Set("OptionCode", fmt.Sprint(req.OptionCode))
	}
	v.Set("DerivMonth", fmt.Sprint(req.DerivMonth))
	if req.PutOrCall != "" {
		v.Set("PutOrCall", fmt.Sprint(req.PutOrCall))
	}
	if req.StrikePrice != 0 {
		v.Set("StrikePrice", fmt.Sprint(req.StrikePrice))
	}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoSymbolnameMinioptionweekly は **GET /symbolname/minioptionweekly** のリクエスト。
//
// ### 概要
// ミニオプション（限週）銘柄コード取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolnameMinioptionweekly struct {
	// 限月<br>※限月はyyyyMM形式で指定します。0を指定した場合、直近限月となります。<br>※取引最終日に「0」（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivMonth string `json:"-" query:"DerivMonth" required:"true"`
	// 限週<br>※限週は0,1,3,4,5のいずれかを指定します。0を指定した場合、指定した限月の直近限週となります。<br>※取引最終日に「0」（直近限週）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限週の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivWeekly int `json:"-" query:"DerivWeekly" required:"true"`
	// コール or プット<br> ※大文字小文字は区別しません。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>P</td> <td>PUT</td> </tr> <tr> <td>C</td> <td>CALL</td> </tr> </tbody> </table>
	PutOrCall string `json:"-" query:"PutOrCall" required:"true"`
	// 権利行使価格<br>※0を指定した場合、APIを実行した時点でのATMとなります。
	StrikePrice int `json:"-" query:"StrikePrice" required:"true"`
}

// ResGetInfoSymbolnameMinioptionweekly は **GET /symbolname/minioptionweekly** のレスポンス。
//
// ### 概要
// ミニオプション（限週）銘柄コード取得
type ResGetInfoSymbolnameMinioptionweekly struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名称
	SymbolName string `json:"SymbolName,omitempty"`
}

// GetInfoSymbolnameMinioptionweekly は **GET /symbolname/minioptionweekly** を呼び出します。
//
// ### 概要
// ミニオプション（限週）銘柄コード取得
//
// ミニオプション（限週）銘柄コード取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoSymbolnameMinioptionweekly(ReqGetInfoSymbolnameMinioptionweekly{ /* TODO: フィールドを設定 */ })
func GetInfoSymbolnameMinioptionweekly(req ReqGetInfoSymbolnameMinioptionweekly) (code int, res ResGetInfoSymbolnameMinioptionweekly, err error) {
	p := "/symbolname/minioptionweekly"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	v.Set("DerivMonth", fmt.Sprint(req.DerivMonth))
	if req.DerivWeekly != 0 {
		v.Set("DerivWeekly", fmt.Sprint(req.DerivWeekly))
	}
	if req.PutOrCall != "" {
		v.Set("PutOrCall", fmt.Sprint(req.PutOrCall))
	}
	if req.StrikePrice != 0 {
		v.Set("StrikePrice", fmt.Sprint(req.StrikePrice))
	}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoRanking は **GET /ranking** のリクエスト。
//
// ### 概要
// 詳細ランキング
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoRanking struct {
	// 種別<br> ※信用情報ランキングに「福証」「札証」を指定した場合は、空レスポンスになります <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>値上がり率（デフォルト）</td> </tr> <tr> <td>2</td> <td>値下がり率</td> </tr> <tr> <td>3</td> <td>売買高上位</td> </tr> <tr> <td>4</td> <td>売買代金</td> </tr> <tr> <td>5</td> <td>TICK回数</td> </tr> <tr> <td>6</td> <td>売買高急増</td> </tr> <tr> <td>7</td> <td>売買代金急増</td> </tr> <tr> <td>8</td> <td>信用売残増</td> </tr> <tr> <td>9</td> <td>信用売残減</td> </tr> <tr> <td>10</td> <td>信用買残増</td> </tr> <tr> <td>11</td> <td>信用買残減</td> </tr> <tr> <td>12</td> <td>信用高倍率</td> </tr> <tr> <td>13</td> <td>信用低倍率</td> </tr> <tr> <td>14</td> <td>業種別値上がり率</td> </tr> <tr> <td>15</td> <td>業種別値下がり率</td> </tr> </tbody> </table>
	Type string `json:"-" query:"Type" required:"true"`
	// 市場<br> ※業種別値上がり率・値下がり率に市場を指定しても無視されます <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>ALL</td> <td>全市場（デフォルト）</td> </tr> <tr> <td>T</td> <td>東証全体</td> </tr> <tr> <td>TP</td> <td>東証プライム</td> </tr> <tr> <td>TS</td> <td>東証スタンダード</td> </tr> <tr> <td>TG</td> <td>グロース250</td> </tr> <tr> <td>M</td> <td>名証</td> </tr> <tr> <td>FK</td> <td>福証</td> </tr> <tr> <td>S</td> <td>札証</td> </tr> </tbody> </table>
	ExchangeDivision string `json:"-" query:"ExchangeDivision" required:"true"`
}

// ResGetInfoRanking は **GET /ranking** のレスポンス。
//
// ### 概要
// 詳細ランキング
type ResGetInfoRanking struct {
}

// GetInfoRanking は **GET /ranking** を呼び出します。
//
// ### 概要
// 詳細ランキング
//
// 詳細ランキング画面と同様の各種ランキングを返します。 <br>ランキングの対象日はkabuステーションが保持している当日のデータとなります。 <br>※株価情報ランキング、業種別指数ランキングは、下記の時間帯でデータがクリアされるため、 <br>その間の詳細ランキングAPIは空レスポンスとなります。 <br>データクリア：平日7:53頃-9:00過ぎ頃 <br>※信用情報ランキングは毎週第３営業日の7:55頃にデータが更新されます。
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoRanking(ReqGetInfoRanking{ /* TODO: フィールドを設定 */ })
func GetInfoRanking(req ReqGetInfoRanking) (code int, res ResGetInfoRanking, err error) {
	p := "/ranking"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Type != "" {
		v.Set("Type", fmt.Sprint(req.Type))
	}
	if req.ExchangeDivision != "" {
		v.Set("ExchangeDivision", fmt.Sprint(req.ExchangeDivision))
	}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoExchangeSymbol は **GET /exchange/{symbol}** のリクエスト。
//
// ### 概要
// 為替情報
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoExchangeSymbol struct {
	// 通貨 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>usdjpy</td> <td>USD/JPY</td> </tr> <tr> <td>eurjpy</td> <td>EUR/JPY</td> </tr> <tr> <td>gbpjpy</td> <td>GBP/JPY</td> </tr> <tr> <td>audjpy</td> <td>AUD/JPY</td> </tr> <tr> <td>chfjpy</td> <td>CHF/JPY</td> </tr> <tr> <td>cadjpy</td> <td>CAD/JPY</td> </tr> <tr> <td>nzdjpy</td> <td>NZD/JPY</td> </tr> <tr> <td>zarjpy</td> <td>ZAR/JPY</td> </tr> <tr> <td>eurusd</td> <td>EUR/USD</td> </tr> <tr> <td>gbpusd</td> <td>GBP/USD</td> </tr> <tr> <td>audusd</td> <td>AUD/USD</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoExchangeSymbol は **GET /exchange/{symbol}** のレスポンス。
//
// ### 概要
// 為替情報
type ResGetInfoExchangeSymbol struct {
	// 通貨
	Symbol string `json:"Symbol,omitempty"`
	// BID
	BidPrice float64 `json:"BidPrice,omitempty"`
	// SP
	Spread float64 `json:"Spread,omitempty"`
	// ASK
	AskPrice float64 `json:"AskPrice,omitempty"`
	// 前日比
	Change float64 `json:"Change,omitempty"`
	// 時刻 <br>※HH:mm:ss形式
	Time string `json:"Time,omitempty"`
}

// GetInfoExchangeSymbol は **GET /exchange/{symbol}** を呼び出します。
//
// ### 概要
// 為替情報
//
// マネービューの情報を取得する
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoExchangeSymbol(ReqGetInfoExchangeSymbol{ /* TODO: フィールドを設定 */ })
func GetInfoExchangeSymbol(req ReqGetInfoExchangeSymbol) (code int, res ResGetInfoExchangeSymbol, err error) {
	p := "/exchange/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoRegulationsSymbol は **GET /regulations/{symbol}** のリクエスト。
//
// ### 概要
// 規制情報
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoRegulationsSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。 <b>市場コード</b> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoRegulationsSymbol は **GET /regulations/{symbol}** のレスポンス。
//
// ### 概要
// 規制情報
type ResGetInfoRegulationsSymbol struct {
	// 銘柄コード<br> ※対象商品は、株式のみ
	Symbol string `json:"Symbol,omitempty"`
	// 規制情報
	RegulationsInfo []struct {
		Exchange      int    `json:"Exchange,omitempty"`      // 規制市場 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>全対象</td> </tr> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>9</td> <td>SOR</td> </tr> <tr> <td>10</td> <td>CXJ</td> </tr> <tr> <td>21</td> <td>JNX</td> </tr> </tbody> </table>
		Product       int    `json:"Product,omitempty"`       // 規制取引区分<br> ※空売り規制の場合、「4：新規」 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>全対象</td> </tr> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>2</td> <td>信用新規（制度）</td> </tr> <tr> <td>3</td> <td>信用新規（一般）</td> </tr> <tr> <td>4</td> <td>新規</td> </tr> <tr> <td>5</td> <td>信用返済（制度）</td> </tr> <tr> <td>6</td> <td>信用返済（一般）</td> </tr> <tr> <td>7</td> <td>返済</td> </tr> <tr> <td>8</td> <td>品受</td> </tr> <tr> <td>9</td> <td>品渡</td> </tr> </tbody> </table>
		Side          string `json:"Side,omitempty"`          // 規制売買<br> ※空売り規制の場合、「1：売」 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>全対象</td> </tr> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
		Reason        string `json:"Reason,omitempty"`        // 理由<br>※空売り規制の場合、「空売り規制」
		LimitStartDay string `json:"LimitStartDay,omitempty"` // 制限開始日<br>yyyy/MM/dd HH:mm形式 <br>※空売り規制の場合、null
		LimitEndDay   string `json:"LimitEndDay,omitempty"`   // 制限終了日<br>yyyy/MM/dd HH:mm形式 <br>※空売り規制の場合、null
		Level         int    `json:"Level,omitempty"`         // コンプライアンスレベル<br> ※空売り規制の場合、null <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>０</td> <td>規制無し</td> </tr> <tr> <td>１</td> <td>ワーニング</td> </tr> <tr> <td>２</td> <td>エラー</td> </tr> </tbody> </table>
	} `json:"RegulationsInfo,omitempty"`
}

// GetInfoRegulationsSymbol は **GET /regulations/{symbol}** を呼び出します。
//
// ### 概要
// 規制情報
//
// 規制情報＋空売り規制情報を取得する
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoRegulationsSymbol(ReqGetInfoRegulationsSymbol{ /* TODO: フィールドを設定 */ })
func GetInfoRegulationsSymbol(req ReqGetInfoRegulationsSymbol) (code int, res ResGetInfoRegulationsSymbol, err error) {
	p := "/regulations/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoPrimaryexchangeSymbol は **GET /primaryexchange/{symbol}** のリクエスト。
//
// ### 概要
// 優先市場
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoPrimaryexchangeSymbol struct {
	// 銘柄コード
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoPrimaryexchangeSymbol は **GET /primaryexchange/{symbol}** のレスポンス。
//
// ### 概要
// 優先市場
type ResGetInfoPrimaryexchangeSymbol struct {
	// 銘柄コード<br>※対象商品は、株式のみ
	Symbol string `json:"Symbol,omitempty"`
	// 優先市場 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> </tbody> </table>
	PrimaryExchange int `json:"PrimaryExchange,omitempty"`
}

// GetInfoPrimaryexchangeSymbol は **GET /primaryexchange/{symbol}** を呼び出します。
//
// ### 概要
// 優先市場
//
// 株式の優先市場を取得する
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoPrimaryexchangeSymbol(ReqGetInfoPrimaryexchangeSymbol{ /* TODO: フィールドを設定 */ })
func GetInfoPrimaryexchangeSymbol(req ReqGetInfoPrimaryexchangeSymbol) (code int, res ResGetInfoPrimaryexchangeSymbol, err error) {
	p := "/primaryexchange/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoApisoftlimit は **GET /apisoftlimit** のリクエスト。
//
// ### 概要
// ソフトリミット
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoApisoftlimit struct {
}

// ResGetInfoApisoftlimit は **GET /apisoftlimit** のレスポンス。
//
// ### 概要
// ソフトリミット
type ResGetInfoApisoftlimit struct {
	// 現物のワンショット上限<br>※単位は万円
	Stock float64 `json:"Stock,omitempty"`
	// 信用のワンショット上限<br>※単位は万円
	Margin float64 `json:"Margin,omitempty"`
	// 先物のワンショット上限<br>※単位は枚
	Future float64 `json:"Future,omitempty"`
	// ミニ先物のワンショット上限<br>※単位は枚
	FutureMini float64 `json:"FutureMini,omitempty"`
	// マイクロ先物のワンショット上限<br>※単位は枚
	FutureMicro float64 `json:"FutureMicro,omitempty"`
	// オプションのワンショット上限<br>※単位は枚
	Option float64 `json:"Option,omitempty"`
	// ミニオプションのワンショット上限<br>※単位は枚
	MiniOption float64 `json:"MiniOption,omitempty"`
	// kabuステーションのバージョン
	KabuSVersion string `json:"KabuSVersion,omitempty"`
}

// GetInfoApisoftlimit は **GET /apisoftlimit** を呼び出します。
//
// ### 概要
// ソフトリミット
//
// kabuステーションAPIのソフトリミット値を取得する
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoApisoftlimit(ReqGetInfoApisoftlimit{ /* TODO: フィールドを設定 */ })
func GetInfoApisoftlimit(req ReqGetInfoApisoftlimit) (code int, res ResGetInfoApisoftlimit, err error) {
	p := "/apisoftlimit"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetInfoMarginMarginpremiumSymbol は **GET /margin/marginpremium/{symbol}** のリクエスト。
//
// ### 概要
// プレミアム料取得
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoMarginMarginpremiumSymbol struct {
	// 銘柄コード
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoMarginMarginpremiumSymbol は **GET /margin/marginpremium/{symbol}** のレスポンス。
//
// ### 概要
// プレミアム料取得
type ResGetInfoMarginMarginpremiumSymbol struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 一般信用（長期）
	GeneralMargin struct {
		MarginPremiumType  int     `json:"MarginPremiumType,omitempty"`  // プレミアム料入力区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>null</td> <td>一般信用（長期）非対応銘柄</td> </tr> <tr> <td>0</td> <td>プレミアム料がない銘柄</td> </tr> <tr> <td>1</td> <td>プレミアム料が固定の銘柄</td> </tr> <tr> <td>2</td> <td>プレミアム料が入札で決定する銘柄</td> </tr> </tbody> </table>
		MarginPremium      float64 `json:"MarginPremium,omitempty"`      // 確定プレミアム料<br> ※入札銘柄の場合、入札受付中は随時更新します。受付時間外は、確定したプレミアム料を返します。<br> ※非入札銘柄の場合、常に固定値を返します。<br> ※信用取引不可の場合、nullを返します。<br> ※19:30~翌営業日のプレミアム料になります。
		UpperMarginPremium float64 `json:"UpperMarginPremium,omitempty"` // 上限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		LowerMarginPremium float64 `json:"LowerMarginPremium,omitempty"` // 下限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		TickMarginPremium  float64 `json:"TickMarginPremium,omitempty"`  // プレミアム料刻値<br> ※入札可能銘柄以外は、nullを返します。
	} `json:"GeneralMargin,omitempty"`
	// 一般信用（デイトレ）
	DayTrade struct {
		MarginPremiumType  int     `json:"MarginPremiumType,omitempty"`  // プレミアム料入力区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>null</td> <td>一般信用（デイトレ）非対応銘柄</td> </tr> <tr> <td>0</td> <td>プレミアム料がない銘柄</td> </tr> <tr> <td>1</td> <td>プレミアム料が固定の銘柄</td> </tr> <tr> <td>2</td> <td>プレミアム料が入札で決定する銘柄</td> </tr> </tbody> </table>
		MarginPremium      float64 `json:"MarginPremium,omitempty"`      // 確定プレミアム料<br> ※入札銘柄の場合、入札受付中は随時更新します。受付時間外は、確定したプレミアム料を返します。<br> ※非入札銘柄の場合、常に固定値を返します。<br> ※信用取引不可の場合、nullを返します。<br> ※19:30~翌営業日のプレミアム料になります。
		UpperMarginPremium float64 `json:"UpperMarginPremium,omitempty"` // 上限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		LowerMarginPremium float64 `json:"LowerMarginPremium,omitempty"` // 下限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		TickMarginPremium  float64 `json:"TickMarginPremium,omitempty"`  // プレミアム料刻値<br> ※入札可能銘柄以外は、nullを返します。
	} `json:"DayTrade,omitempty"`
}

// GetInfoMarginMarginpremiumSymbol は **GET /margin/marginpremium/{symbol}** を呼び出します。
//
// ### 概要
// プレミアム料取得
//
// 指定した銘柄のプレミアム料を取得するAPI
//
// - 認証: 必須（X-API-KEY ヘッダー）
//
// ### 使い方
//
//	code, res, err := GetInfoMarginMarginpremiumSymbol(ReqGetInfoMarginMarginpremiumSymbol{ /* TODO: フィールドを設定 */ })
func GetInfoMarginMarginpremiumSymbol(req ReqGetInfoMarginMarginpremiumSymbol) (code int, res ResGetInfoMarginMarginpremiumSymbol, err error) {
	p := "/margin/marginpremium/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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
