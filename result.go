package mini


type result struct {
	rest		[]string
	flags		map[string]string
}

func (r *result) GetRest() []string {
	return r.rest
}

func (r *result) GetFlag(flag string) string {
	return r.flags[flag]
}

func (r *result) GetFlags() map[string]string {
	return r.flags
}
