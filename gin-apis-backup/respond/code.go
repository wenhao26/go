package respond

type Codes struct {
	SUCCESS uint
	FAILED  uint
	Message map[uint]string
}

var APICode = &Codes{
	SUCCESS: 1,
	FAILED:  0,
}

func init() {
	APICode.Message = map[uint]string{
		APICode.SUCCESS: "请求成功",
		APICode.FAILED:  "请求失败",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
