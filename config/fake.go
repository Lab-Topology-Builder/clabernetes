package config

import (
	k8scorev1 "k8s.io/api/core/v1"
)

// GetFakeManager returns a fake config manager -- eventually this should have some options to load
// it with data for unit tests. That is a future me problem.
func GetFakeManager() Manager {
	return fakeManager{}
}

type fakeManager struct{}

func (f fakeManager) Start() error {
	return nil
}

func (f fakeManager) GetGlobalAnnotations() map[string]string {
	return make(map[string]string)
}

func (f fakeManager) GetGlobalLabels() map[string]string {
	return make(map[string]string)
}

func (f fakeManager) GetAllMetadata() (annotations, labels map[string]string) {
	return f.GetGlobalAnnotations(), f.GetGlobalLabels()
}

func (f fakeManager) GetResourcesForContainerlabKind(
	containerlabKind string,
	containerlabType string,
) *k8scorev1.ResourceRequirements {
	_, _ = containerlabKind, containerlabType

	return nil
}

func (f fakeManager) GetInClusterDNSSuffix() string {
	return "svc.cluster.local"
}

func (f fakeManager) GetImagePullThroughOverride() string {
	return "auto"
}
