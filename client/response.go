package client

// Response is a container for contents of a template http response.
type Response struct {
	HTTPStatus int
	Message    string `json:"message"`
	Payload    interface{}
}
