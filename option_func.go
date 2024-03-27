package goya

type OptionFunc func(opt *Option)

// WithData will add data to the request body
// data can be a structure or a map
func WithData(data any) OptionFunc {
	return func(opt *Option) {
		opt.Data = data
	}
}

// WithParams will add params to the request URL
// params can be a structure or a map
func WithParams(params any) OptionFunc {
	return func(opt *Option) {
		opt.Params = params
	}
}
