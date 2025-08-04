/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
)

// Setup_containerengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		nodepool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
