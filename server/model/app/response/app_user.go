package response

type IdImageFrontResult struct {
	Reason    string `json:"reason"`
	Result    Front  `json:"result"`
	ErrorCode int    `json:"error_code"`
}

type IdImageBackResult struct {
	Reason    string `json:"reason"`
	Result    Back   `json:"result"`
	ErrorCode int    `json:"error_code"`
}

type Front struct {
	RealName string `json:"realname"`
	Sex      string `json:"sex"`
	Nation   string `json:"nation"`
	Born     string `json:"born"`
	Address  string `json:"address"`
	Idcard   string `json:"idcard"`
	Side     string `json:"side"`
	OrderId  string `json:"orderid"`
}

type Back struct {
	Begin      string `json:"begin"`
	Department string `json:"department"`
	End        string `json:"end"`
	Side       string `json:"side"`
	OrderId    string `json:"orderid"`
}

type BankCardResult struct {
	Reason    string   `json:"reason"`
	Result    BankCard `json:"result"`
	ErrorCode int      `json:"error_code"`
}

type BankCard struct {
	BankCardNumber string `json:"bank_card_number"`
	ValidDate      string `json:"valid_date"`
	BankCardType   string `json:"bank_card_type"`
	BankName       string `json:"bank_name"`
}

//营业执照查询结果
type BusResult struct {
	Reason    string `json:"reason"`
	Result    Bus    `json:"result"`
	ErrorCode int    `json:"error_code"`
}

type Bus struct {
	ResultNum int   `json:"result_num"`
	Res       []res `json:"res"`
}

type res struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
