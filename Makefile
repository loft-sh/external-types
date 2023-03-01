generate:
	  deepcopy-gen \
	  	--go-header-file hack/boilerplate.go.txt \
	  	--input-dirs ./argoproj/argo-cd/v2/pkg/apis/,./argoproj/argo-cd/v2/pkg/apis/application,./argoproj/argo-cd/v2/pkg/apis/application/v1alpha1 \
	  	-O zz_generated.deepcopy \
	  	-o ./

	  deepcopy-gen \
		--go-header-file hack/boilerplate.go.txt \
		--input-dirs ./loft-sh/admin-services/pkg/server/ \
		-O zz_generated.deepcopy \
		-o ./
