package server

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Module string

var DefaultGetFeatureFunction = GetFeature

var GetFeature = func(name string) (Feature, error) {
	for _, feat := range AllFeatures {
		if name == feat.Name {
			return feat, nil
		}
	}
	return Feature{}, nil
}

type FeatureSpec struct {
	Hidden      bool   `json:"hidden,omitempty"`
	Module      Module `json:"module,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type FeatureStatus struct {
	Entitled      bool   `json:"entitled"`
	Enabled       bool   `json:"enabled"`
	BuyLink       string `json:"buy,omitempty"`
	TryLink       string `json:"try,omitempty"`
	LearnMoreLink string `json:"learnMore,omitempty"`
}

func NewFeature(
	name string, spec FeatureSpec) Feature {
	feat := Feature{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: spec,
	}

	AllFeatures = append(AllFeatures, feat)

	return feat
}

type Feature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FeatureSpec   `json:"spec,omitempty"`
	Status FeatureStatus `json:"status,omitempty"`

	get            *func(string, func(string) (Feature, error)) (Feature, error)
	LegacyVersions map[string]Feature
}

func (f Feature) WithCustomGetFunction(
	fn func(string, func(string) (Feature, error)) (Feature, error),
) Feature {
	f.get = &fn
	return f
}

func (f Feature) Get() (Feature, error) {
	if f.get != nil {
		fn := *f.get
		return fn(f.ObjectMeta.Name, DefaultGetFeatureFunction)
	}
	return Feature{}, nil
}

const (
	ModuleDevPod                Module = "DevPod"
	ModuleVirtualClusters       Module = "Virtual Clusters"
	ModuleKubernetes            Module = "Kubernetes Management"
	ModulePlatformAuth          Module = "Platform: Authentication"
	ModulePlatformTemplating    Module = "Platform: Templating & Secrets"
	ModulePlatformDeployment    Module = "Platform: Deployment"
	ModulePlatformCustomization Module = "Platform: Customization"
	ModulePlatformIntegrations  Module = "Platform: Integrations"
)

var (
	AllFeatures = []Feature{}

	// DevPod
	DevPodWorkspaces = NewFeature(
		"devpod-workspaces",
		FeatureSpec{
			Hidden: true,
			Module: ModuleDevPod,
		})

	// Runners
	Runners = NewFeature(
		"runners",
		FeatureSpec{
			Hidden: true,
			Module: ModuleDevPod,
		})

	// Virtual Clusters
	VirtualClusters = NewFeature(
		"vcluster",
		FeatureSpec{
			DisplayName: "Virtual Cluster CRD & Controller",
			Module:      ModuleVirtualClusters,
		})
	VirtualClusterSleepMode = NewFeature(
		"vcluster-sleep-mode",
		FeatureSpec{
			DisplayName: "Sleep Mode",
			Module:      ModuleVirtualClusters,
		})
	VirtualClusterBuildInCoreDNS = NewFeature(
		"vcluster-built-in-coredns",
		FeatureSpec{
			DisplayName: "Built-in CoreDNS",
		})
	VirtualClusterSyncPatches = NewFeature(
		"vcluster-sync-patches",
		FeatureSpec{
			DisplayName: "Sync Patches",
		})
	VirtualClusterAdmissionControl = NewFeature(
		"vcluster-admission-control",
		FeatureSpec{
			DisplayName: "Virtual Admission Control",
		})
	VirtualClusterIsolatedControlPlane = NewFeature(
		"vcluster-isolated-control-plane",
		FeatureSpec{
			DisplayName: "Isolated Control Plane",
		})
	VirtualClusterCentralHostPathMapper = NewFeature(
		"vcluster-host-path-mapper",
		FeatureSpec{
			DisplayName: "Central HostPath Mapper",
		})

	// Spaces & Clusters
	Spaces = NewFeature(
		"spaces",
		FeatureSpec{
			DisplayName: "Self-Service Namespaces",
			Module:      ModuleKubernetes,
		})
	SpaceSleepMode = NewFeature(
		"spaces-sleep-mode",
		FeatureSpec{
			DisplayName: "Sleep Mode For Namespaces",
			Module:      ModuleKubernetes,
		})
	ConnectedClusters = NewFeature(
		"clusters",
		FeatureSpec{
			DisplayName: "Connected Clusters",
			Module:      ModuleKubernetes,
		})
	ClusterAccess = NewFeature(
		"cluster-access",
		FeatureSpec{
			Hidden: true,
		})
	ClusterRoles = NewFeature(
		"cluster-roles",
		FeatureSpec{
			Hidden: true,
		})

	// Auth-Related Features
	AuditLogging = NewFeature(
		"audit-logging",
		FeatureSpec{
			DisplayName: "Audit Logging",
			Module:      ModulePlatformAuth,
		})
	SSOAuth = NewFeature(
		"sso-authentication",
		FeatureSpec{
			DisplayName: "Single Sign-On (SSO)",
			Module:      ModulePlatformAuth,
		})
	MultipleSSOProviders = NewFeature(
		"multiple-sso-providers",
		FeatureSpec{
			DisplayName: "Multiple SSO Providers",
			Module:      ModulePlatformAuth,
		})
	AutomaticIngressAuth = NewFeature(
		"auto-ingress-authentication",
		FeatureSpec{
			DisplayName: "Automatic Ingress Authentication",
			Module:      ModulePlatformAuth,
		})
	OIDCProvider = NewFeature(
		"oidc-provider",
		FeatureSpec{
			DisplayName: "OIDC Provider",
			Module:      ModulePlatformAuth,
		})

	// Templating Features
	TemplateVersioning = NewFeature(
		"template-versioning",
		FeatureSpec{
			DisplayName: "Template Versioning",
			Module:      ModulePlatformTemplating,
		})
	Apps = NewFeature(
		"apps",
		FeatureSpec{
			DisplayName: "Apps",
			Module:      ModulePlatformTemplating,
		})

	// Secrets
	Secrets = NewFeature(
		"secrets",
		FeatureSpec{
			Module: ModulePlatformTemplating,
		})
	SecretEncryption = NewFeature(
		"secret-encyrption",
		FeatureSpec{
			Module: ModulePlatformTemplating,
		})

	// Integrations
	VaultIntegration = NewFeature(
		"vault-integration",
		FeatureSpec{
			Module: ModulePlatformIntegrations,
		})
	ArgoIntegration = NewFeature(
		"argo-integration",
		FeatureSpec{
			Module: ModulePlatformIntegrations,
		})

	// HA & Other Advanced Deployment Features
	HighAvailabilityMode = NewFeature(
		"ha-mode",
		FeatureSpec{
			DisplayName: "High-Availability Mode",
			Module:      ModulePlatformDeployment,
		})
	MultiRegionMode = NewFeature(
		"multi-region-mode",
		FeatureSpec{
			DisplayName: "Multi-Region Mode",
			Module:      ModulePlatformDeployment,
		})
	AirGappedMode = NewFeature( // Purely For Display Purposes
		"air-gapped-mode",
		FeatureSpec{
			DisplayName: "Air-Gapped Mode",
			Module:      ModulePlatformDeployment,
		})

	// UI Customization Features
	CustomBranding = NewFeature(
		"custom-branding",
		FeatureSpec{
			DisplayName: "Custom Branding",
			Module:      ModulePlatformCustomization,
		})
	AdvancedUICustomizations = NewFeature(
		"advanced-ui-customizations",
		FeatureSpec{
			DisplayName: "Advanced UI Customizations",
			Module:      ModulePlatformCustomization,
		})
)
