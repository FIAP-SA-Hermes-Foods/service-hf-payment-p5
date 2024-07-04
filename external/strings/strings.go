package strings

import (
	"encoding/json"

	l "service-hf-payment-p5/external/logger"
)

func MarshalString(s interface{}) string {
	if s == nil {
		return ""
	}

	o, err := json.Marshal(s)
	if err != nil {
		l.Errorf("error in MarshalString payment ", " | ", err)
		return ""
	}

	return string(o)
}
