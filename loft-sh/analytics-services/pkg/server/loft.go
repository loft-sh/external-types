package server

import "github.com/loft-sh/external-types/loft-sh/admin-services/pkg/server"

// +genclient
// +genclient:nonNamespaced

// AnalyticsInput is a struct holding analytics data sent from the Loft instance ot the analytics
// backend. Note that no sensitive data is sent and usernames and similar identifiable fields are
// sent simply as hashed values.
type AnalyticsInput struct {
	*server.InstanceTokenAuth

	LoftVersion string `form:"version"     json:"version"     validate:"required"`
	KubeVersion string `form:"kubeVersion" json:"kubeVersion" validate:"required"`

	KubeSystemNamespaceUID string `form:"kubeSystemNamespaceUID" json:"kubeSystemNamespace" validate:"required"`

	// Requests holds all requests that have occurred in the api gateway
	// and are requested by the analytics server
	Requests []Request `json:"requests"`

	// UserActivity holds information about user activity since
	// the last request
	UserActivity []UserActivity `json:"userActivity"`
}

// Request holds information about an analytics relevant request that has occurred in the loft
// api gateway.
type Request struct {
	// Timestamp the request occurred
	Timestamp int64 `json:"timestamp"`

	// Hostname of the request
	Hostname string `json:"hostname"`

	// ProxyHostname of the request
	ProxyHostname string `json:"proxyHostname"`

	// The cluster (if any) the request was sent to
	Cluster string `json:"cluster"`

	// The kubernetes api group of the request
	Group string `json:"group" binding:"required"`

	// The kubernetes api version of the request
	Version string `json:"version" binding:"required"`

	// The kubernetes verb of the request
	Verb string `json:"verb" binding:"required"`

	// The kubernetes resource of the request
	Resource string `json:"resource" binding:"required"`

	// The namespace of the request
	Namespace string `json:"namespace"`

	// The kubernetes object name of the request
	Name string `json:"name"`

	// The json encoded object that was returned by the server
	Object string `json:"object"`

	// User uid of the user that sent the request
	UserUID string `json:"user"`

	// Username (hash) of the user that sent the request
	Username string `json:"username"`
}

// UserActivity holds information about a certain user and his
// activity in loft
type UserActivity struct {
	// User uid of the user that sent the request
	UserUID string `json:"user"`

	// Username (hash) of the user that sent the request
	Username string `json:"username"`

	// Activity describes the activity of an user within loft.
	// The map holds key value pairs of the 'cluster/namespace' that were targeted
	// by the user as key and as value the amount of requests that have targeted
	// this cluster/namespace.
	Activity map[string]int `json:"activity"`
}
