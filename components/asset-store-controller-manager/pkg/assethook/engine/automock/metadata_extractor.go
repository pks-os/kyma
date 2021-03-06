// Code generated by mockery v1.0.0
package automock

import context "context"
import engine "github.com/kyma-project/kyma/components/asset-store-controller-manager/pkg/assethook/engine"
import mock "github.com/stretchr/testify/mock"
import v1alpha2 "github.com/kyma-project/kyma/components/asset-store-controller-manager/pkg/apis/assetstore/v1alpha2"

// MetadataExtractor is an autogenerated mock type for the MetadataExtractor type
type MetadataExtractor struct {
	mock.Mock
}

// Extract provides a mock function with given fields: ctx, object, basePath, files, services
func (_m *MetadataExtractor) Extract(ctx context.Context, object engine.Accessor, basePath string, files []string, services []v1alpha2.WebhookService) ([]engine.File, error) {
	ret := _m.Called(ctx, object, basePath, files, services)

	var r0 []engine.File
	if rf, ok := ret.Get(0).(func(context.Context, engine.Accessor, string, []string, []v1alpha2.WebhookService) []engine.File); ok {
		r0 = rf(ctx, object, basePath, files, services)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]engine.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, engine.Accessor, string, []string, []v1alpha2.WebhookService) error); ok {
		r1 = rf(ctx, object, basePath, files, services)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
