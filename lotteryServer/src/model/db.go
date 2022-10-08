package model

type RedArray [34]int
type BlueArray [17]int

type CacheConn struct {
	Addr		string `json:"addr"`
	Password	string `json:"password"`
}

type DBConn struct {
	User   string `json:"user"`
	Pawd   string `json:"pawd"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DbName string `json:"dbname"`
}

type TblDoubleColorHist struct {
	Date string `gorm:"primaryKey"`
	Red  string
	Blue string
}

type TblDoubleColorTotal struct {
	Reds  string
	Blues string
}

type DoubleColorRed struct {
	RSlots RedArray
}

type DoubleColorBlue struct {
	BSlots BlueArray
}

func (ball *DoubleColorRed) Add(indexes []int) {
	for _, index := range indexes {
		ball.RSlots[index]++
	}
}

func (ball *DoubleColorBlue) Add(indexes []int) {
	for _, index := range indexes {
		ball.BSlots[index]++
	}
}
