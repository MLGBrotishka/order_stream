package entity

import "time"

type OrderUID string

// Order
type Order struct {
	OrderUID          OrderUID  `json:"order_uid" example:"b563feb7b2b84b6test" validate:"required"`
	TrackNumber       string    `json:"track_number" example:"WBILMTESTTRACK"`
	Entry             string    `json:"entry" example:"WBIL"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required,dive"`
	Locale            string    `json:"locale" example:"en"`
	InternalSignature string    `json:"internal_signature" example:""`
	CustomerID        string    `json:"customer_id" example:"test" validate:"required"`
	DeliveryService   string    `json:"delivery_service" example:"meest"`
	ShardKey          string    `json:"shardkey" example:"9"`
	SMID              int       `json:"sm_id" example:"99"`
	DateCreated       time.Time `json:"date_created" example:"2021-11-26T06:22:19Z" validate:"required"`
	OOFShard          string    `json:"oof_shard" example:"1"`
}

// Delivery
type Delivery struct {
	Name    string `json:"name" example:"Test Testov" validate:"required"`
	Phone   string `json:"phone" example:"+9720000000"`
	Zip     string `json:"zip" example:"2639809"`
	City    string `json:"city" example:"Kiryat Mozkin"`
	Address string `json:"address" example:"Ploshad Mira 15" validate:"required"`
	Region  string `json:"region" example:"Kraiot"`
	Email   string `json:"email" example:"test@gmail.com"`
}

// Payment
type Payment struct {
	Transaction  string `json:"transaction" example:"b563feb7b2b84b6test" validate:"required"`
	RequestID    string `json:"request_id" example:""`
	Currency     string `json:"currency" example:"USD" validate:"required"`
	Provider     string `json:"provider" example:"wbpay"`
	Amount       int    `json:"amount" example:"1817"`
	PaymentDT    int64  `json:"payment_dt" example:"1637907727" validate:"required"`
	Bank         string `json:"bank" example:"alpha" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" example:"1500"`
	GoodsTotal   int    `json:"goods_total" example:"317"`
	CustomFee    int    `json:"custom_fee" example:"0"`
}

// Item
type Item struct {
	CHRTID      int    `json:"chrt_id" example:"9934930" validate:"required"`
	TrackNumber string `json:"track_number" example:"WBILMTESTTRACK"`
	Price       int    `json:"price" example:"453"`
	RID         string `json:"rid" example:"ab4219087a764ae0btest"`
	Name        string `json:"name" example:"Mascaras"`
	Sale        int    `json:"sale" example:"30"`
	Size        string `json:"size" example:"0"`
	TotalPrice  int    `json:"total_price" example:"317"`
	NMID        int    `json:"nm_id" example:"2389212"`
	Brand       string `json:"brand" example:"Vivienne Sabo"`
	Status      int    `json:"status" example:"202"`
}
