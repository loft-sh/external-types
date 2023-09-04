package server

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var DefaultGetFeatureFunction = GetFeature
var GetFeature = func(name string) (Feature, error) {
	for _, w := range AllFeatures {
		if name == w.feature.Name {
			return w.feature, nil
		}
	}
	return Feature{}, nil
}

type FeatureWrapper struct {
	feature Feature

	get            *func(string, func(string) (Feature, error)) (Feature, error)
	LegacyVersions map[string]Feature
}

func NewFeature(name string, spec FeatureSpec) FeatureWrapper {
	w := FeatureWrapper{
		feature: Feature{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: spec,
		},
	}

	AllFeatures = append(AllFeatures, w)

	return w
}

func (w FeatureWrapper) WithCustomGetFunction(
	fn func(string, func(string) (Feature, error)) (Feature, error),
) FeatureWrapper {
	w.get = &fn
	return w
}

func (w FeatureWrapper) Get() (Feature, error) {
	if w.get != nil {
		fn := *w.get
		return fn(w.feature.ObjectMeta.Name, DefaultGetFeatureFunction)
	}
	return Feature{}, nil
}
