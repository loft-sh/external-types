package server

// +genclient
// +genclient:nonNamespaced

// IntercomHashCreateInput is the required input data for generating a hash for a user for intercom.
type IntercomHashCreateInput struct {
	*InstanceTokenAuth

	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// IntercomHashCreateOutput is the struct holding all information returned from "intercom
// generate user hash" requests.
type IntercomHashCreateOutput struct {
	CurrentTime int64  `json:"currentTime"`
	Hash        string `json:"hash"`
}
