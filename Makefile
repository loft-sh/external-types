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

