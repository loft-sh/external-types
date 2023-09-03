package server

// +genclient
// +genclient:nonNamespaced

// ChatAuthCreateInput is the required input data for chat authentication
type ChatAuthCreateInput struct {
	*InstanceTokenAuth

	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// ChatAuthCreateOutput is returned by the server
type ChatAuthCreateOutput struct{}
