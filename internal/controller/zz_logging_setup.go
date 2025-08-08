/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	log "github.com/oracle/provider-oci/internal/controller/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/logging/unifiedagentconfiguration"
)

// Setup_logging creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logging(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
