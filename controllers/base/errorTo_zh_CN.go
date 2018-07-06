package base

const (
	always_haven_it = "200 OK"
	forbiden_403    = "403 forbiden"
)

func ErrorTo_zh_CN(err error) string {
	switch err.Error() {
	case always_haven_it:
		return always_haven_it
	case forbiden_403:
		return forbiden_403
	default:
		return err.Error()
	}
}
