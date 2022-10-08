package http

type HttpInvoker struct {
	Request Request
}

// GetBodyParam /*
func (httpInvoker *HttpInvoker) GetBodyParam(r *Response, jpath string) string {
	return GetBodyParam(r, jpath)
}

func (httpInvoker *HttpInvoker) Call(r Request) (*Response, error) {
	//TODO: robert
	return nil, nil
}
