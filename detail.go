package problem

type (
	FieldDetail struct {
		Field  string `json:"field"`
		Value  string `json:"value"`
		Reason string `json:"reason"`
	}
)

func NewFieldDetail(field, value, reason string) FieldDetail {
	return FieldDetail{
		Field:  field,
		Value:  value,
		Reason: reason,
	}
}

func (e FieldDetail) Error() string {
	return e.Field + ": " + e.Reason
}
