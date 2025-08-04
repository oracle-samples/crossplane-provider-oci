/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	backend "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/backendset"
	listener "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/listener"
	networkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/networkloadbalancer"
	networkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/networkloadbalancersbackendsetsunified"
)

// Setup_networkloadbalancer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkloadbalancer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.Setup,
		backendset.Setup,
		listener.Setup,
		networkloadbalancer.Setup,
		networkloadbalancersbackendsetsunified.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
