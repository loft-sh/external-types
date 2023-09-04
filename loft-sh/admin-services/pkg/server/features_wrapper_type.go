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

func NewFeature(name string, spec FeatureSpec) Feature {
	w := FeatureWrapper{
		feature: Feature{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: spec,
		},
	}

	AllFeatures[name] = w

	return w.feature
}

func (f Feature) WithCustomGetFunction(
	fn func(string, func(string) (Feature, error)) (Feature, error),
) Feature {
	w, ok := AllFeatures[f.ObjectMeta.Name]
	if ok {
		w.get = &fn
	}
	return w.feature
}

func (w FeatureWrapper) Get() (Feature, error) {
	if w.get != nil {
		fn := *w.get
		return fn(w.feature.ObjectMeta.Name, DefaultGetFeatureFunction)
	}
	return Feature{}, nil
}

func (w FeatureWrapper) GetName() string {
	return w.feature.Name
}
