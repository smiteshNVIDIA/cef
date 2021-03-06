// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "CompletionCallback_gen.h"
	"C"
)

// CompletionCallbackProxy defines methods required for using CompletionCallback.
type CompletionCallbackProxy interface {
	OnComplete(self *CompletionCallback)
}

// CompletionCallback (cef_completion_callback_t from include/capi/cef_callback_capi.h)
// Generic callback structure used for asynchronous completion.
type CompletionCallback C.cef_completion_callback_t

// NewCompletionCallback creates a new CompletionCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewCompletionCallback(proxy CompletionCallbackProxy) *CompletionCallback {
	result := (*CompletionCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_completion_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_completion_callback_proxy(result.toNative())
	}
	return result
}

func (d *CompletionCallback) toNative() *C.cef_completion_callback_t {
	return (*C.cef_completion_callback_t)(d)
}

func lookupCompletionCallbackProxy(obj *BaseRefCounted) CompletionCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(CompletionCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type CompletionCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *CompletionCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnComplete (on_complete)
// Method that will be called once the task is complete.
func (d *CompletionCallback) OnComplete() {
	lookupCompletionCallbackProxy(d.Base()).OnComplete(d)
}

//nolint:gocritic
//export gocef_completion_callback_on_complete
func gocef_completion_callback_on_complete(self *C.cef_completion_callback_t) {
	me__ := (*CompletionCallback)(self)
	proxy__ := lookupCompletionCallbackProxy(me__.Base())
	proxy__.OnComplete(me__)
}
