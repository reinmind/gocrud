package db

import "gocrud/entity"

type Response struct {
	Msg   string
	Code  int
	Value entity.TradeCount
}
