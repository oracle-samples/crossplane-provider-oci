/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/identity/identityprovider"
	policy "github.com/oracle/provider-oci/internal/controller/identity/policy"
	tag "github.com/oracle/provider-oci/internal/controller/identity/tag"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/identity/tagnamespace"
)

// Setup_identity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authenticationpolicy.Setup,
		compartment.Setup,
		group.Setup,
		identityprovider.Setup,
		policy.Setup,
		tag.Setup,
		tagnamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
