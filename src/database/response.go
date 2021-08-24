package database

import "crud-demo/src/database/entity"

type Response struct {
	Msg string
	Code int
	Value entity.TradeCount
}