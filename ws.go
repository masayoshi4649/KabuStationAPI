package kabusapi

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

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

// OpenQuote は endpoint に接続して Quote を受け取り、handlers に流します。
// 戻り値は (closeFunc, error)。closeFunc を呼ぶと接続を閉じます。
func OpenQuote(handlers ...func(Quote)) (func(), error) {
	if endpoint == "" {
		return nil, errors.New("kabusapi: endpoint is empty")
	}

	dialer := websocket.DefaultDialer
	conn, _, dErr := dialer.Dial(endpoint, nil)
	if dErr != nil {
		return nil, dErr
	}

	done := make(chan struct{})
	var once sync.Once

	// safe wrapper: handler 内の panic を吸収
	safeCall := func(h func(Quote), q Quote) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("kabusapi: handler panic recovered: %v", r)
			}
		}()
		h(q)
	}

	// 読み取りループ
	go func() {
		// ここで組み込み close(done) を呼ぶ -> done チャネルを閉じる
		defer close(done)

		for {
			_, msg, rErr := conn.ReadMessage()
			if rErr != nil {
				log.Printf("kabusapi: websocket read error: %v", rErr)
				return
			}

			var q Quote
			if jErr := json.Unmarshal(msg, &q); jErr != nil {
				log.Printf("kabusapi: json unmarshal error: %v; raw: %s", jErr, string(msg))
				continue
			}

			for _, h := range handlers {
				if h == nil {
					continue
				}
				go safeCall(h, q)
			}
		}
	}()

	// closeFn: 何度呼んでも安全
	closeFn := func() {
		once.Do(func() {
			_ = conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(1*time.Second))
			_ = conn.Close()
			// 読み取り goroutine が done を閉じるのを待つ
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
		})
	}

	return closeFn, nil
}
