package model

type OfficialDoubleColor struct {
	State     int                       `json:"state"`
	Message   string                    `json:"message"`
	PageCount int                       `json:"pageCount"`
	CountNum  int                       `json:"countNum"`
	Tflag     int                       `json:"Tflag"`
	Result    []OfficialDoubleColorData `json:"result"`
}

type OfficialDoubleColorData struct {
	Name        string       `json:"name,omitempty"`
	Code        string       `json:"code,omitempty"`
	DetailsLink string       `json:"detailsLink,omitempty"`
	VideoLink   string       `json:"videoLink,omitempty"`
	Date        string       `json:"date,omitempty"`
	Week        string       `json:"week,omitempty"`
	Red         string       `json:"red,omitempty"`
	Blue        string       `json:"blue,omitempty"`
	Blue2       string       `json:"blue2,omitempty"`
	Sales       string       `json:"sales,omitempty"`
	PoolMoney   string       `json:"poolmoney,omitempty"`
	Content     string       `json:"content,omitempty"`
	AddMoney    string       `json:"addmoney,omitempty"`
	AddMoney2   string       `json:"addmoney2,omitempty"`
	Msg         string       `json:"msg,omitempty"`
	Z2add       string       `json:"z2add,omitempty"`
	M2add       string       `json:"m2add,omitempty"`
	PrizeGrades []PrizeGrade `json:"prizeGrades"`
}

type PrizeGrade struct {
	Type      int    `json:"type"`
	TypeNum   string `json:"typenum"`
	TypeMoney string `json:"typemoney"`
}
