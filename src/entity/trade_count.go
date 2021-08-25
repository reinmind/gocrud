package entity

type TradeCount struct {
	MainId    string `gorm:"column:main_id"`
	AreaCode  string `gorm:"column:area_code"`
	Kind      string `gorm:"column:type"`
	TradeCode string `gorm:"column:trade_code"`
	Count     string `gorm:"column:cnt"`
	TradeName string `gorm:"column:trade_name"`
}
