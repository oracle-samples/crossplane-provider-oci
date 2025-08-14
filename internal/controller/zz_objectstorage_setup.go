/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bucket "github.com/oracle/provider-oci/internal/controller/objectstorage/bucket"
	object "github.com/oracle/provider-oci/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/objectlifecyclepolicy"
)

// Setup_objectstorage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_objectstorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
