generate:
	go run k8s.io/code-generator/cmd/deepcopy-gen@v0.35.0 \
		--go-header-file hack/boilerplate.go.txt \
		--output-file zz_generated.deepcopy.go \
		./argoproj/argo-cd/v2/pkg/apis \
		./argoproj/argo-cd/v2/pkg/apis/application \
		./argoproj/argo-cd/v2/pkg/apis/application/v1alpha1

	go run k8s.io/code-generator/cmd/deepcopy-gen@v0.35.0 \
		--go-header-file hack/boilerplate.go.txt \
		--output-file zz_generated.deepcopy.go \
		./loft-sh/admin-services/pkg/server

	go run k8s.io/kube-openapi/cmd/openapi-gen@v0.0.0-20260330154417-16be699c7b31 \
		--go-header-file hack/boilerplate.go.txt \
		--output-file zz_generated.openapi.go \
		--output-model-name-file zz_generated.model_name.go \
		--output-dir ./argoproj/argo-cd/v2/pkg/apis/application/v1alpha1 \
		--output-pkg github.com/loft-sh/external-types/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1 \
		--report-filename /dev/null \
		./argoproj/argo-cd/v2/pkg/apis/application/v1alpha1

	go run k8s.io/kube-openapi/cmd/openapi-gen@v0.0.0-20260330154417-16be699c7b31 \
		--go-header-file hack/boilerplate.go.txt \
		--output-file zz_generated.openapi.go \
		--output-dir ./loft-sh/admin-services/pkg/server \
		--output-pkg github.com/loft-sh/external-types/loft-sh/admin-services/pkg/server \
		--report-filename /dev/null \
		./loft-sh/admin-services/pkg/server

