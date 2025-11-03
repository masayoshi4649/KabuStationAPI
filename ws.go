package kabusapi

import "time"

const endpoint = "ws://localhost:18080/kabusapi/websocket"

// Websocket受信データ
type Quote struct {
	AskPrice float64 `json:"AskPrice"`
	AskQty   float64 `json:"AskQty"`
	AskSign  string  `json:"AskSign"`
	BidPrice float64 `json:"BidPrice"`
	BidQty   float64 `json:"BidQty"`
	BidSign  string  `json:"BidSign"`
	Buy1     struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
		Sign  string  `json:"Sign"`
	} `json:"Buy1"`
	Buy10 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy10"`
	Buy2 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy2"`
	Buy3 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy3"`
	Buy4 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy4"`
	Buy5 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy5"`
	Buy6 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy6"`
	Buy7 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy7"`
	Buy8 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy8"`
	Buy9 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy9"`
	CalcPrice                float64   `json:"CalcPrice"`
	ChangePreviousClose      float64   `json:"ChangePreviousClose"`
	ChangePreviousClosePer   float64   `json:"ChangePreviousClosePer"`
	ClearingPrice            float64   `json:"ClearingPrice"`
	CurrentPrice             float64   `json:"CurrentPrice"`
	CurrentPriceChangeStatus string    `json:"CurrentPriceChangeStatus"`
	CurrentPriceStatus       float64   `json:"CurrentPriceStatus"`
	CurrentPriceTime         time.Time `json:"CurrentPriceTime"`
	Exchange                 float64   `json:"Exchange"`
	ExchangeName             string    `json:"ExchangeName"`
	HighPrice                float64   `json:"HighPrice"`
	HighPriceTime            time.Time `json:"HighPriceTime"`
	LowPrice                 float64   `json:"LowPrice"`
	LowPriceTime             time.Time `json:"LowPriceTime"`
	OpeningPrice             float64   `json:"OpeningPrice"`
	OpeningPriceTime         time.Time `json:"OpeningPriceTime"`
	PreviousClose            float64   `json:"PreviousClose"`
	PreviousCloseTime        time.Time `json:"PreviousCloseTime"`
	SecurityType             float64   `json:"SecurityType"`
	Sell1                    struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
		Sign  string  `json:"Sign"`
	} `json:"Sell1"`
	Sell10 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell10"`
	Sell2 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell2"`
	Sell3 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell3"`
	Sell4 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell4"`
	Sell5 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell5"`
	Sell6 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell6"`
	Sell7 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell7"`
	Sell8 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell8"`
	Sell9 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell9"`
	Symbol            string    `json:"Symbol"`
	SymbolName        string    `json:"SymbolName"`
	TradingValue      float64   `json:"TradingValue"`
	TradingVolume     float64   `json:"TradingVolume"`
	TradingVolumeTime time.Time `json:"TradingVolumeTime"`
	VWAP              float64   `json:"VWAP"`
}
