/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/compute/clusternetwork"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/compute/computecapacityreservation"
	computecluster "github.com/oracle/provider-oci/internal/controller/compute/computecluster"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/compute/instanceconsoleconnection"
	instancepool "github.com/oracle/provider-oci/internal/controller/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/compute/shapemanagement"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appcataloglistingresourceversionagreement.Setup,
		appcatalogsubscription.Setup,
		clusternetwork.Setup,
		computecapacityreservation.Setup,
		computecluster.Setup,
		computeimagecapabilityschema.Setup,
		consolehistory.Setup,
		dedicatedvmhost.Setup,
		image.Setup,
		instance.Setup,
		instanceconfiguration.Setup,
		instanceconsoleconnection.Setup,
		instancepool.Setup,
		instancepoolinstance.Setup,
		shapemanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
