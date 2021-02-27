package mini


type result struct {
	rest		[]string
	flags		map[string]interface{}
}

func (r *result) GetRest() []string {
	return r.rest
}

func (r *result) GetFlag(flag string) interface{} {
	return r.flags[flag]
}

func (r *result) GetFlags() map[string]interface{} {
	return r.flags
}
