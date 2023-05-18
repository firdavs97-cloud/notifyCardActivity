package card

type MockActivityEvent struct {
	ActivityEventInput
	Results map[string][]interface{}
}

func (a MockActivityEvent) Save() (int64, error) {
	var res1 int64
	var res2 error
	if len(a.Results["Save"]) > 0 {
		res1 = a.Results["Save"][0].(int64)
	}
	if len(a.Results["Save"]) > 1 {
		res2 = a.Results["Save"][1].(error)
	}
	return res1, res2
}

func (a MockActivityEvent) Notify(i int64) error {
	var res error
	if len(a.Results["Notify"]) > 0 {
		res = a.Results["Notify"][0].(error)
	}
	return res
}

func (a MockActivityEvent) Load(id int64) error {
	var res error
	if len(a.Results["Load"]) > 0 {
		res = a.Results["Load"][0].(error)
	}
	return res
}