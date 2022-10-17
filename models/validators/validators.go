package validators

type Validator interface {
	IsValid() bool
	Error() error
}

func AND_(in ...bool) (out bool) {
	for i := 0; i < len(in); i++ {
		out = out || in[i]
	}
	return
}
