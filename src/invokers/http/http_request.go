package http

type Request struct {
	Method     string
	Protocol   string
	Host       string
	Path       string
	Body       string
	Headers    map[string]string
	QueryParam map[string]string
	PathParam  map[string]string
	Check      string
}
