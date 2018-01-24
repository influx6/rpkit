package rp

import (
	"context"
)

// ServerHooks implements a structure to be used in a servers hook.
type ServerHooks struct {
	ResponseSentFunc     func(context.Context)
	ResponsePreparedFunc func(context.Context)
	RequestRejectedFunc  func(context.Context)
	RequestAcceptedFunc  func(context.Context)
	RequestReceivedFunc  func(context.Context)
	RequestProcessedFunc func(context.Context)
}

func (sh ServerHooks) ResponseSent(ctx context.Context) {
	if sh.ResponseSentFunc != nil {
		sh.ResponseSentFunc(ctx)
	}
}

func (sh ServerHooks) ResponsePrepared(ctx context.Context) {
	if sh.ResponsePreparedFunc != nil {
		sh.ResponsePreparedFunc(ctx)
	}
}

func (sh ServerHooks) RequestRejected(ctx context.Context) {
	if sh.RequestRejectedFunc != nil {
		sh.RequestRejectedFunc(ctx)
	}
}

func (sh ServerHooks) RequestAccepted(ctx context.Context) {
	if sh.RequestAcceptedFunc != nil {
		sh.RequestAcceptedFunc(ctx)
	}
}

func (sh ServerHooks) RequestReceived(ctx context.Context) {
	if sh.RequestReceivedFunc != nil {
		sh.RequestReceivedFunc(ctx)
	}
}

func (sh ServerHooks) RequestProcessed(ctx context.Context) {
	if sh.RequestProcessedFunc != nil {
		sh.RequestProcessedFunc(ctx)
	}
}
