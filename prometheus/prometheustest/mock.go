package prometheustest

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/mock"
)

// PromAPIMock for mocking Prometheus API
type PromAPIMock struct {
	mock.Mock
}

func (o *PromAPIMock) Query(ctx context.Context, query string, ts time.Time) (model.Value, error) {
	args := o.Called(ctx, query, ts)
	return args.Get(0).(model.Value), args.Error(1)
}

func (o *PromAPIMock) QueryRange(ctx context.Context, query string, r v1.Range) (model.Value, error) {
	args := o.Called(ctx, query, r)
	return args.Get(0).(model.Value), args.Error(1)
}

func (o *PromAPIMock) LabelValues(ctx context.Context, label string) (model.LabelValues, error) {
	args := o.Called(ctx, label)
	return args.Get(0).(model.LabelValues), args.Error(1)
}

func (o *PromAPIMock) Series(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]model.LabelSet, error) {
	args := o.Called(ctx, matches, startTime, endTime)
	return args.Get(0).([]model.LabelSet), args.Error(1)
}

// AlwaysReturnEmpty mocks all possible queries to return empty result
func (o *PromAPIMock) AlwaysReturnEmpty() {
	metric := model.Metric{
		"__name__": "whatever",
		"instance": "whatever",
		"job":      "whatever"}
	o.On(
		"Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("time.Time"),
	).Return(model.Vector{}, nil)
	matrix := model.Matrix{
		&model.SampleStream{
			Metric: metric,
			Values: []model.SamplePair{}}}
	o.On(
		"QueryRange",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("v1.Range"),
	).Return(matrix, nil)
}

// SpyArgumentsAndReturnEmpty mocks all possible queries to return empty result,
// allowing to spy arguments through input callback
func (o *PromAPIMock) SpyArgumentsAndReturnEmpty(fn func(args mock.Arguments)) {
	metric := model.Metric{
		"__name__": "whatever",
		"instance": "whatever",
		"job":      "whatever"}
	o.On(
		"Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("time.Time"),
	).Run(fn).Return(model.Vector{}, nil)
	matrix := model.Matrix{
		&model.SampleStream{
			Metric: metric,
			Values: []model.SamplePair{}}}
	o.On(
		"QueryRange",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("v1.Range"),
	).Run(fn).Return(matrix, nil)
}
