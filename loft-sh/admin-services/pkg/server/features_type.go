package server

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced

// Button is an object that represents a button in the Loft UI that links to some external service
// for handling operations for licensing for example.
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type Feature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FeatureSpec   `json:"spec,omitempty"`
	Status FeatureStatus `json:"status,omitempty"`
}

type FeatureSpec struct {
	Hidden      bool   `json:"hidden,omitempty"`
	Module      string `json:"module,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type FeatureStatus struct {
	Entitled      bool   `json:"entitled"`
	Enabled       bool   `json:"enabled"`
	BuyLink       string `json:"buy,omitempty"`
	TryLink       string `json:"try,omitempty"`
	LearnMoreLink string `json:"learnMore,omitempty"`
}
