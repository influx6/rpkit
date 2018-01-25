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
	RequestErrorFunc     func(context.Context, error)
	RequestPanicFunc     func(context.Context, interface{})
}

// ResponseSent calls the ResponseSentFunc function field.
func (sh ServerHooks) ResponseSent(ctx context.Context) {
	if sh.ResponseSentFunc != nil {
		sh.ResponseSentFunc(ctx)
	}
}

// ResponsePrepared calls the ResponsePreparedFunc function field.
func (sh ServerHooks) ResponsePrepared(ctx context.Context) {
	if sh.ResponsePreparedFunc != nil {
		sh.ResponsePreparedFunc(ctx)
	}
}

// RequestRejected calls the RequestRejectedFunc function field.
func (sh ServerHooks) RequestRejected(ctx context.Context) {
	if sh.RequestRejectedFunc != nil {
		sh.RequestRejectedFunc(ctx)
	}
}

// RequestAccepted calls the RequestAcceptedFunc function field.
func (sh ServerHooks) RequestAccepted(ctx context.Context) {
	if sh.RequestAcceptedFunc != nil {
		sh.RequestAcceptedFunc(ctx)
	}
}

// RequestReceived calls the RequestReceivedFunc function field.
func (sh ServerHooks) RequestReceived(ctx context.Context) {
	if sh.RequestReceivedFunc != nil {
		sh.RequestReceivedFunc(ctx)
	}
}

// RequestProcessed calls the RequestProcessedFunc function field.
func (sh ServerHooks) RequestProcessed(ctx context.Context) {
	if sh.RequestProcessedFunc != nil {
		sh.RequestProcessedFunc(ctx)
	}
}

// RequestError calls the RequestErrorFunc function field.
func (sh ServerHooks) RequestError(ctx context.Context, err error) {
	if sh.RequestErrorFunc != nil {
		sh.RequestErrorFunc(ctx, err)
	}
}

// RequestPanic calls the RequestPanicFunc function field.
func (sh ServerHooks) RequestPanic(ctx context.Context, err interface{}) {
	if sh.RequestPanicFunc != nil {
		sh.RequestPanicFunc(ctx, err)
	}
}
