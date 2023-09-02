package server

type Module string

type Feature struct {
	Name           string             `json:"name"`
	Enabled        bool               `json:"enabled"`
	Hidden         bool               `json:"hidden,omitempty"`
	Module         Module             `json:"module,omitempty"`
	DisplayName    string             `json:"displayName,omitempty"`
	LearnMoreLink  string             `json:"learnMore,omitempty"`
	BuyLink        string             `json:"buy,omitempty"`
	TryLink        string             `json:"try,omitempty"`
	LegacyVersions map[string]Feature `json:"-"`
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
	AllFeatures = map[string]Feature{
		DevPodWorkspaces.Name: DevPodWorkspaces,
		Runners.Name:          Runners,

		VirtualClusters.Name:                    VirtualClusters,
		VirtualClusterSleepMode.Name:            VirtualClusterSleepMode,
		VirtualClusterBuildInCoreDNS.Name:       VirtualClusterBuildInCoreDNS,
		VirtualClusterSyncPatches.Name:          VirtualClusterSyncPatches,
		VirtualClusterAdmissionControl.Name:     VirtualClusterAdmissionControl,
		VirtualClusterIsolatedControlPlane.Name: VirtualClusterIsolatedControlPlane,

		AuditLogging.Name:             AuditLogging,
		SSOAuth.Name:                  SSOAuth,
		MultipleSSOProviders.Name:     MultipleSSOProviders,
		AutomaticIngressAuth.Name:     AutomaticIngressAuth,
		OIDCProvider.Name:             OIDCProvider,
		HighAvailabilityMode.Name:     HighAvailabilityMode,
		MultiRegionMode.Name:          MultiRegionMode,
		CustomBranding.Name:           CustomBranding,
		AdvancedUICustomizations.Name: AdvancedUICustomizations,
		TemplateVersioning.Name:       TemplateVersioning,
		Apps.Name:                     Apps,
		Runners.Name:                  Runners,
	}

	// DevPod
	DevPodWorkspaces = Feature{
		Name:   "devpod-workspaces",
		Hidden: true,
		Module: ModuleDevPod,
	}

	// Runners
	Runners = Feature{
		Name:   "runners",
		Hidden: true,
		Module: ModuleDevPod,
	}

	// Virtual Clusters
	VirtualClusters = Feature{
		Name:        "vcluster",
		DisplayName: "Virtual Cluster CRD & Controller",
		Module:      ModuleVirtualClusters,
	}
	VirtualClusterSleepMode = Feature{
		Name:        SpaceSleepMode.Name,
		DisplayName: "Sleep Mode",
		Module:      ModuleVirtualClusters,
	}
	VirtualClusterBuildInCoreDNS = Feature{
		Name:        "vcluster-built-in-coredns",
		DisplayName: "Built-in CoreDNS",
	}
	VirtualClusterSyncPatches = Feature{
		Name:        "vcluster-sync-patches",
		DisplayName: "Sync Patches",
	}
	VirtualClusterAdmissionControl = Feature{
		Name:        "vcluster-admission-control",
		DisplayName: "Virtual Admission Control",
	}
	VirtualClusterIsolatedControlPlane = Feature{
		Name:        "vcluster-isolated-control-plane",
		DisplayName: "Isolated Control Plane",
	}
	VirtualClusterCentralHostPathMapper = Feature{
		Name:        "vcluster-host-path-mapper",
		DisplayName: "Central HostPath Mapper",
	}

	// Spaces & Clusters
	Spaces Feature = Feature{
		Name:        "spaces",
		DisplayName: "Self-Service Namespaces",
		Module:      ModuleKubernetes,
	}
	ConnectClusters = Feature{
		Name:        "clusters",
		DisplayName: "Connect Clusters",
		Module:      ModuleKubernetes,
	}
	SpaceSleepMode = Feature{
		Name:        "sleep-mode",
		DisplayName: "Sleep Mode For Namespaces",
		Module:      ModuleKubernetes,
	}
	ClusterAccess = Feature{
		Name:   "cluster-access",
		Hidden: true,
	}
	ClusterRoles = Feature{
		Name:   "cluster-roles",
		Hidden: true,
	}

	// Auth-Related Features
	AuditLogging = Feature{
		Name:        "audit-logging",
		DisplayName: "Audit Logging",
		Module:      ModulePlatformAuth,
	}
	SSOAuth = Feature{
		Name:        "sso-authentication",
		DisplayName: "Single Sign-On (SSO)",
		Module:      ModulePlatformAuth,
	}
	MultipleSSOProviders = Feature{
		Name:        "multiple-sso-providers",
		DisplayName: "Multiple SSO Providers",
		Module:      ModulePlatformAuth,
	}
	AutomaticIngressAuth = Feature{
		Name:        "auto-ingress-authentication",
		DisplayName: "Automatic Ingress Authentication",
		Module:      ModulePlatformAuth,
	}
	OIDCProvider = Feature{
		Name:        "oidc-provider",
		DisplayName: "OIDC Provider",
		Module:      ModulePlatformAuth,
	}

	// Templating Features
	TemplateVersioning = Feature{ // TODO: Needs to be implemented
		Name:        "template-versioning",
		DisplayName: "Template Versioning",
		Module:      ModulePlatformTemplating,
	}
	Apps = Feature{
		Name:        "apps",
		DisplayName: "Apps",
		Module:      ModulePlatformTemplating,
	}

	// Secrets
	Secrets = Feature{
		Name:   "secrets",
		Module: ModulePlatformTemplating,
	}
	SecretEncryption = Feature{
		Name:   "secret-encyrption",
		Module: ModulePlatformTemplating,
	}

	// Integrations
	VaultIntegration = Feature{
		Name:   "vault-integration",
		Module: ModulePlatformIntegrations,
	}
	ArgoIntegration = Feature{
		Name:   "argo-integration",
		Module: ModulePlatformIntegrations,
	}

	// HA & Other Advanced Deployment Features
	HighAvailabilityMode = Feature{
		Name:        "ha-mode",
		DisplayName: "High-Availability Mode",
		Module:      ModulePlatformDeployment,
	}
	MultiRegionMode = Feature{
		Name:        "multi-region-mode",
		DisplayName: "Multi-Region Mode",
		Module:      ModulePlatformDeployment,
	}
	AirGappedMode = Feature{
		Name:        "air-gapped-mode",
		DisplayName: "Air-Gapped Mode",
		Module:      ModulePlatformDeployment,
	}

	// UI Customization Features
	CustomBranding = Feature{
		Name:        "custom-branding",
		DisplayName: "Custom Branding",
		Module:      ModulePlatformCustomization,
	}
	AdvancedUICustomizations = Feature{
		Name:        "advanced-ui-customizations",
		DisplayName: "Advanced UI Customizations",
		Module:      ModulePlatformCustomization,
	}
)
