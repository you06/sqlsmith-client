package concurrency

import (
	"github.com/you06/doppelganger/executor"
	"github.com/juju/errors"
)

// Executor struct
type Executor struct {
	opt         *executor.Option
	concurrency int
	executors   []*executor.Executor
}

// New create Executor
func New(dsn string, opt *executor.Option, concurrency int) (*Executor, error) {
	e := Executor{
		opt: opt,
		concurrency: concurrency,
	}
	for i := 0; i < concurrency; i++ {
		exec, err := executor.New(dsn, opt.Clone())
		if err != nil {
			return nil, errors.Trace(err)
		}
		e.executors = append(e.executors, exec)
	}
	return &e, nil
}

// NewABTest create abtest Executor
func NewABTest(dsn1, dsn2 string, opt *executor.Option, concurrency int) (*Executor, error) {
	e := Executor{
		opt: opt,
		concurrency: concurrency,
	}
	for i := 0; i < concurrency; i++ {
		exec, err := executor.NewABTest(dsn1, dsn2, opt.Clone())
		if err != nil {
			return nil, errors.Trace(err)
		}
		e.executors = append(e.executors, exec)
	}
	return &e, nil
}
