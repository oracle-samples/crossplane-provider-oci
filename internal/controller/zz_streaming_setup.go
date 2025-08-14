/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	connectharness "github.com/oracle/provider-oci/internal/controller/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/streaming/streampool"
)

// Setup_streaming creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_streaming(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
