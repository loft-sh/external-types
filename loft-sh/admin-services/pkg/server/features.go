package server

type Module string

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
	AllFeatures = []FeatureWrapper{}

	// DevPod
	DevPodWorkspaces = NewFeature(
		"devpod-workspaces",
		FeatureSpec{
			Hidden: true,
			Module: string(ModuleDevPod),
		})

	// Runners
	Runners = NewFeature(
		"runners",
		FeatureSpec{
			Hidden: true,
			Module: string(ModuleDevPod),
		})

	// Virtual Clusters
	VirtualClusters = NewFeature(
		"vcluster",
		FeatureSpec{
			DisplayName: "Virtual Cluster CRD & Controller",
			Module:      string(ModuleVirtualClusters),
		})
	VirtualClusterSleepMode = NewFeature(
		"vcluster-sleep-mode",
		FeatureSpec{
			DisplayName: "Sleep Mode",
			Module:      string(ModuleVirtualClusters),
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
			Module:      string(ModuleKubernetes),
		})
	SpaceSleepMode = NewFeature(
		"spaces-sleep-mode",
		FeatureSpec{
			DisplayName: "Sleep Mode For Namespaces",
			Module:      string(ModuleKubernetes),
		})
	ConnectedClusters = NewFeature(
		"clusters",
		FeatureSpec{
			DisplayName: "Connected Clusters",
			Module:      string(ModuleKubernetes),
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
			Module:      string(ModulePlatformAuth),
		})
	SSOAuth = NewFeature(
		"sso-authentication",
		FeatureSpec{
			DisplayName: "Single Sign-On (SSO)",
			Module:      string(ModulePlatformAuth),
		})
	MultipleSSOProviders = NewFeature(
		"multiple-sso-providers",
		FeatureSpec{
			DisplayName: "Multiple SSO Providers",
			Module:      string(ModulePlatformAuth),
		})
	AutomaticIngressAuth = NewFeature(
		"auto-ingress-authentication",
		FeatureSpec{
			DisplayName: "Automatic Ingress Authentication",
			Module:      string(ModulePlatformAuth),
		})
	OIDCProvider = NewFeature(
		"oidc-provider",
		FeatureSpec{
			DisplayName: "OIDC Provider",
			Module:      string(ModulePlatformAuth),
		})

	// Templating Features
	TemplateVersioning = NewFeature(
		"template-versioning",
		FeatureSpec{
			DisplayName: "Template Versioning",
			Module:      string(ModulePlatformTemplating),
		})
	Apps = NewFeature(
		"apps",
		FeatureSpec{
			DisplayName: "Apps",
			Module:      string(ModulePlatformTemplating),
		})

	// Secrets
	Secrets = NewFeature(
		"secrets",
		FeatureSpec{
			Module: string(ModulePlatformTemplating),
		})
	SecretEncryption = NewFeature(
		"secret-encyrption",
		FeatureSpec{
			Module: string(ModulePlatformTemplating),
		})

	// Integrations
	VaultIntegration = NewFeature(
		"vault-integration",
		FeatureSpec{
			Module: string(ModulePlatformIntegrations),
		})
	ArgoIntegration = NewFeature(
		"argo-integration",
		FeatureSpec{
			Module: string(ModulePlatformIntegrations),
		})

	// HA & Other Advanced Deployment Features
	HighAvailabilityMode = NewFeature(
		"ha-mode",
		FeatureSpec{
			DisplayName: "High-Availability Mode",
			Module:      string(ModulePlatformDeployment),
		})
	MultiRegionMode = NewFeature(
		"multi-region-mode",
		FeatureSpec{
			DisplayName: "Multi-Region Mode",
			Module:      string(ModulePlatformDeployment),
		})
	AirGappedMode = NewFeature( // Purely For Display Purposes
		"air-gapped-mode",
		FeatureSpec{
			DisplayName: "Air-Gapped Mode",
			Module:      string(ModulePlatformDeployment),
		})

	// UI Customization Features
	CustomBranding = NewFeature(
		"custom-branding",
		FeatureSpec{
			DisplayName: "Custom Branding",
			Module:      string(ModulePlatformCustomization),
		})
	AdvancedUICustomizations = NewFeature(
		"advanced-ui-customizations",
		FeatureSpec{
			DisplayName: "Advanced UI Customizations",
			Module:      string(ModulePlatformCustomization),
		})
)
