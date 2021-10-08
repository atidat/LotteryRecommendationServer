package model

type Conn struct {
	User string `json:"user"`
	Pawd string `json:"pawd"`
	Host string `json:"host"`
	Port int `json:"port"`
	DbName string `json:"dbname"`
}

type TblDoubleColorHist struct {
	Date string `gorm:"primaryKey"`
	Red string
	Blue string
}

type DoubleColorRed struct {
	One int `json:"one"`
	Two int `json:"two"`
	Three int `json:"three"`
	Four int `json:"four"`
	Five int `json:"five"`
	Six int `json:"six"`
	Seven int `json:"seven"`
	Eight int `json:"eight"`
	Nine int `json:"nine"`
	Ten int `json:"ten"`
	Eleven int `json:"eleven"`
	Twelve int `json:"twelve"`
	Thirteen int `json:"thirteen"`
	Fourteen int `json:"fourteen"`
	Fifteen int `json:"fifteen"`
	Sixteen int `json:"sixteen"`
	Seventeen int `json:"seventeen"`
	Eighteen int `json:"eighteen"`
	Nineteen int `json:"nineteen"`
	Twenty int `json:"twenty"`
	Twentyone  int `json:"twentyone"`
	Twentytwo int `json:"twentytwo"`
	Twentythree int `json:"twentythree"`
	Twentyfour int `json:"twentyfour"`
	Twentyfive int `json:"twentyfive"`
	Twentysix int `json:"twentysix"`
	Twentyseven int `json:"twentyseven"`
	Twentyeight int `json:"twentyeight"`
	Twentynine int `json:"twentynine"`
	Thirty int `json:"thirty"`
	Thirtyone int `json:"thirtyone"`
	Thirtytwo int `json:"thirtytwo"`
	Thirtythree int `json:"thirtythree"`
}

type DoubleColorBlue struct {
	One int `json:"one"`
	Two int `json:"two"`
	Three int `json:"three"`
	Four int `json:"four"`
	Five int `json:"five"`
	Six int `json:"six"`
	Seven int `json:"seven"`
	Eight int `json:"eight"`
	Nine int `json:"nine"`
	Ten int `json:"ten"`
	Eleven int `json:"eleven"`
	Twelve int `json:"twelve"`
	Thirteen int `json:"thirteen"`
	Fourteen int `json:"fourteen"`
	Fifteen int `json:"fifteen"`
	Sixteen int `json:"sixteen"`
}

func(dcb *DoubleColorRed)Add1() { dcb.One++ }
func(dcb *DoubleColorRed)Add2() { dcb.Two++ }
func(dcb *DoubleColorRed)Add3() { dcb.Three++ }
func(dcb *DoubleColorRed)Add4() { dcb.Four++ }
func(dcb *DoubleColorRed)Add5() { dcb.Five++ }
func(dcb *DoubleColorRed)Add6() { dcb.Six++ }
func(dcb *DoubleColorRed)Add7() { dcb.Seven++ }
func(dcb *DoubleColorRed)Add8() { dcb.Eight++ }
func(dcb *DoubleColorRed)Add9() { dcb.Nine++ }
func(dcb *DoubleColorRed)Add10() { dcb.Ten++ }
func(dcb *DoubleColorRed)Add11() { dcb.Eleven++ }
func(dcb *DoubleColorRed)Add12() { dcb.Twelve++ }
func(dcb *DoubleColorRed)Add13() { dcb.Thirteen++ }
func(dcb *DoubleColorRed)Add14() { dcb.Fourteen++ }
func(dcb *DoubleColorRed)Add15() { dcb.Fifteen++ }
func(dcb *DoubleColorRed)Add16() { dcb.Sixteen++ }
func(dcb *DoubleColorRed)Add17() { dcb.Seventeen++ }
func(dcb *DoubleColorRed)Add18() { dcb.Eighteen++ }
func(dcb *DoubleColorRed)Add19() { dcb.Nineteen++ }
func(dcb *DoubleColorRed)Add20() { dcb.Twenty++ }
func(dcb *DoubleColorRed)Add21() { dcb.Twentyone++ }
func(dcb *DoubleColorRed)Add22() { dcb.Twentytwo++ }
func(dcb *DoubleColorRed)Add23() { dcb.Twentythree++ }
func(dcb *DoubleColorRed)Add24() { dcb.Twentyfour++ }
func(dcb *DoubleColorRed)Add25() { dcb.Twentyfive++ }
func(dcb *DoubleColorRed)Add26() { dcb.Twentysix++ }
func(dcb *DoubleColorRed)Add27() { dcb.Twentyseven++ }
func(dcb *DoubleColorRed)Add28() { dcb.Twentyeight++ }
func(dcb *DoubleColorRed)Add29() { dcb.Twentynine++ }
func(dcb *DoubleColorRed)Add30() { dcb.Thirty++ }
func(dcb *DoubleColorRed)Add31() { dcb.Thirtyone++ }
func(dcb *DoubleColorRed)Add32() { dcb.Thirtytwo++ }
func(dcb *DoubleColorRed)Add33() { dcb.Thirtythree++ }

func(dcb *DoubleColorBlue)Add1() { dcb.One++ }
func(dcb *DoubleColorBlue)Add2() { dcb.Two++ }
func(dcb *DoubleColorBlue)Add3() { dcb.Three++ }
func(dcb *DoubleColorBlue)Add4() { dcb.Four++ }
func(dcb *DoubleColorBlue)Add5() { dcb.Five++ }
func(dcb *DoubleColorBlue)Add6() { dcb.Six++ }
func(dcb *DoubleColorBlue)Add7() { dcb.Seven++ }
func(dcb *DoubleColorBlue)Add8() { dcb.Eight++ }
func(dcb *DoubleColorBlue)Add9() { dcb.Nine++ }
func(dcb *DoubleColorBlue)Add10() { dcb.Ten++ }
func(dcb *DoubleColorBlue)Add11() { dcb.Eleven++ }
func(dcb *DoubleColorBlue)Add12() { dcb.Twelve++ }
func(dcb *DoubleColorBlue)Add13() { dcb.Thirteen++ }
func(dcb *DoubleColorBlue)Add14() { dcb.Fourteen++ }
func(dcb *DoubleColorBlue)Add15() { dcb.Fifteen++ }
func(dcb *DoubleColorBlue)Add16() { dcb.Sixteen++ }

type TblDoubleColorRed1 struct {
	DoubleColorRed
}

type TblDoubleColorRed2 struct {
	DoubleColorRed
}

type TblDoubleColorRed3 struct {
	DoubleColorRed
}

type TblDoubleColorRed4 struct {
	DoubleColorRed
}

type TblDoubleColorRed5 struct {
	DoubleColorRed
}

type TblDoubleColorRed6 struct {
	DoubleColorRed
}

type TblDoubleColorBlue struct {
	DoubleColorBlue
}

type TblDoubleColorBalls struct {
	Red1 int
	Red2 int
	Red3 int
	Red4 int
	Red5 int
	Red6 int
	Blue int
}