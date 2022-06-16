package shared

type Args map[string]interface{}

func (a Args) GetValue(v string) interface{} {
	return a[v]
}
