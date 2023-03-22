/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	containerconfiguration "github.com/oracle/provider-oci/internal/controller/artifacts/containerconfiguration"
	containerrepository "github.com/oracle/provider-oci/internal/controller/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/artifacts/repository"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificateauthority"
	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
	capturefilter "github.com/oracle/provider-oci/internal/controller/core/capturefilter"
	cpe "github.com/oracle/provider-oci/internal/controller/core/cpe"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/core/dedicatedvmhost"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/core/dhcpoptions"
	drg "github.com/oracle/provider-oci/internal/controller/core/drg"
	drgattachment "github.com/oracle/provider-oci/internal/controller/core/drgattachment"
	drgattachmentmanagement "github.com/oracle/provider-oci/internal/controller/core/drgattachmentmanagement"
	drgattachmentslist "github.com/oracle/provider-oci/internal/controller/core/drgattachmentslist"
	drgroutedistribution "github.com/oracle/provider-oci/internal/controller/core/drgroutedistribution"
	drgroutedistributionstatement "github.com/oracle/provider-oci/internal/controller/core/drgroutedistributionstatement"
	drgroutetable "github.com/oracle/provider-oci/internal/controller/core/drgroutetable"
	drgroutetablerouterule "github.com/oracle/provider-oci/internal/controller/core/drgroutetablerouterule"
	image "github.com/oracle/provider-oci/internal/controller/core/image"
	instance "github.com/oracle/provider-oci/internal/controller/core/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/core/instanceconfiguration"
	internetgateway "github.com/oracle/provider-oci/internal/controller/core/internetgateway"
	ipsec "github.com/oracle/provider-oci/internal/controller/core/ipsec"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/core/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/core/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/core/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/core/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/core/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/core/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/core/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/core/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/core/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/core/routetable"
	securitylist "github.com/oracle/provider-oci/internal/controller/core/securitylist"
	subnet "github.com/oracle/provider-oci/internal/controller/core/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/core/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/core/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/core/vnicattachment"
	volume "github.com/oracle/provider-oci/internal/controller/core/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/core/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/core/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/core/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/core/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/core/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/core/volumegroupbackup"
	record "github.com/oracle/provider-oci/internal/controller/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/dns/zone"
	rule "github.com/oracle/provider-oci/internal/controller/events/rule"
	storageexport "github.com/oracle/provider-oci/internal/controller/file/storageexport"
	storageexportset "github.com/oracle/provider-oci/internal/controller/file/storageexportset"
	storagefilesystem "github.com/oracle/provider-oci/internal/controller/file/storagefilesystem"
	storagemounttarget "github.com/oracle/provider-oci/internal/controller/file/storagemounttarget"
	storagereplication "github.com/oracle/provider-oci/internal/controller/file/storagereplication"
	storagesnapshot "github.com/oracle/provider-oci/internal/controller/file/storagesnapshot"
	application "github.com/oracle/provider-oci/internal/controller/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/functions/invokefunction"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/identity/identityprovider"
	policy "github.com/oracle/provider-oci/internal/controller/identity/policy"
	tag "github.com/oracle/provider-oci/internal/controller/identity/tag"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/identity/tagnamespace"
	key "github.com/oracle/provider-oci/internal/controller/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/kms/keyversion"
	vault "github.com/oracle/provider-oci/internal/controller/kms/vault"
	balancerbackend "github.com/oracle/provider-oci/internal/controller/load/balancerbackend"
	balancerbackendset "github.com/oracle/provider-oci/internal/controller/load/balancerbackendset"
	balancercertificate "github.com/oracle/provider-oci/internal/controller/load/balancercertificate"
	balancerhostname "github.com/oracle/provider-oci/internal/controller/load/balancerhostname"
	balancerlistener "github.com/oracle/provider-oci/internal/controller/load/balancerlistener"
	balancerloadbalancer "github.com/oracle/provider-oci/internal/controller/load/balancerloadbalancer"
	balancerloadbalancerroutingpolicy "github.com/oracle/provider-oci/internal/controller/load/balancerloadbalancerroutingpolicy"
	balancerpathrouteset "github.com/oracle/provider-oci/internal/controller/load/balancerpathrouteset"
	balancerruleset "github.com/oracle/provider-oci/internal/controller/load/balancerruleset"
	balancersslciphersuite "github.com/oracle/provider-oci/internal/controller/load/balancersslciphersuite"
	log "github.com/oracle/provider-oci/internal/controller/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/logging/unifiedagentconfiguration"
	alarm "github.com/oracle/provider-oci/internal/controller/monitoring/alarm"
	firewallnetworkfirewall "github.com/oracle/provider-oci/internal/controller/network/firewallnetworkfirewall"
	firewallnetworkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/network/firewallnetworkfirewallpolicy"
	loadbalancerbackend "github.com/oracle/provider-oci/internal/controller/network/loadbalancerbackend"
	loadbalancerbackendset "github.com/oracle/provider-oci/internal/controller/network/loadbalancerbackendset"
	loadbalancerlistener "github.com/oracle/provider-oci/internal/controller/network/loadbalancerlistener"
	loadbalancernetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/network/loadbalancernetworkloadbalancer"
	loadbalancernetworkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/network/loadbalancernetworkloadbalancersbackendsetsunified"
	bucket "github.com/oracle/provider-oci/internal/controller/objectstorage/bucket"
	object "github.com/oracle/provider-oci/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/objectlifecyclepolicy"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/ons/subscription"
	providerconfig "github.com/oracle/provider-oci/internal/controller/providerconfig"
	accesspolicy "github.com/oracle/provider-oci/internal/controller/servicemesh/accesspolicy"
	ingressgateway "github.com/oracle/provider-oci/internal/controller/servicemesh/ingressgateway"
	ingressgatewayroutetable "github.com/oracle/provider-oci/internal/controller/servicemesh/ingressgatewayroutetable"
	mesh "github.com/oracle/provider-oci/internal/controller/servicemesh/mesh"
	virtualdeployment "github.com/oracle/provider-oci/internal/controller/servicemesh/virtualdeployment"
	virtualservice "github.com/oracle/provider-oci/internal/controller/servicemesh/virtualservice"
	virtualserviceroutetable "github.com/oracle/provider-oci/internal/controller/servicemesh/virtualserviceroutetable"
	connectharness "github.com/oracle/provider-oci/internal/controller/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/streaming/streampool"
	secret "github.com/oracle/provider-oci/internal/controller/vault/secret"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerconfiguration.Setup,
		containerrepository.Setup,
		genericartifact.Setup,
		repository.Setup,
		certificateauthority.Setup,
		cluster.Setup,
		nodepool.Setup,
		capturefilter.Setup,
		cpe.Setup,
		dedicatedvmhost.Setup,
		dhcpoptions.Setup,
		drg.Setup,
		drgattachment.Setup,
		drgattachmentmanagement.Setup,
		drgattachmentslist.Setup,
		drgroutedistribution.Setup,
		drgroutedistributionstatement.Setup,
		drgroutetable.Setup,
		drgroutetablerouterule.Setup,
		image.Setup,
		instance.Setup,
		instanceconfiguration.Setup,
		internetgateway.Setup,
		ipsec.Setup,
		localpeeringgateway.Setup,
		natgateway.Setup,
		networksecuritygroup.Setup,
		networksecuritygroupsecurityrule.Setup,
		privateip.Setup,
		publicip.Setup,
		publicippool.Setup,
		publicippoolcapacity.Setup,
		remotepeeringconnection.Setup,
		routetable.Setup,
		securitylist.Setup,
		subnet.Setup,
		vcn.Setup,
		vlan.Setup,
		vnicattachment.Setup,
		volume.Setup,
		volumeattachment.Setup,
		volumebackup.Setup,
		volumebackuppolicy.Setup,
		volumebackuppolicyassignment.Setup,
		volumegroup.Setup,
		volumegroupbackup.Setup,
		record.Setup,
		resolver.Setup,
		resolverendpoint.Setup,
		rrset.Setup,
		steeringpolicy.Setup,
		steeringpolicyattachment.Setup,
		tsigkey.Setup,
		view.Setup,
		zone.Setup,
		rule.Setup,
		storageexport.Setup,
		storageexportset.Setup,
		storagefilesystem.Setup,
		storagemounttarget.Setup,
		storagereplication.Setup,
		storagesnapshot.Setup,
		application.Setup,
		function.Setup,
		invokefunction.Setup,
		httpmonitor.Setup,
		pingmonitor.Setup,
		authenticationpolicy.Setup,
		compartment.Setup,
		group.Setup,
		identityprovider.Setup,
		policy.Setup,
		tag.Setup,
		tagnamespace.Setup,
		key.Setup,
		keyversion.Setup,
		vault.Setup,
		balancerbackend.Setup,
		balancerbackendset.Setup,
		balancercertificate.Setup,
		balancerhostname.Setup,
		balancerlistener.Setup,
		balancerloadbalancer.Setup,
		balancerloadbalancerroutingpolicy.Setup,
		balancerpathrouteset.Setup,
		balancerruleset.Setup,
		balancersslciphersuite.Setup,
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
		alarm.Setup,
		firewallnetworkfirewall.Setup,
		firewallnetworkfirewallpolicy.Setup,
		loadbalancerbackend.Setup,
		loadbalancerbackendset.Setup,
		loadbalancerlistener.Setup,
		loadbalancernetworkloadbalancer.Setup,
		loadbalancernetworkloadbalancersbackendsetsunified.Setup,
		bucket.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		notificationtopic.Setup,
		subscription.Setup,
		providerconfig.Setup,
		accesspolicy.Setup,
		ingressgateway.Setup,
		ingressgatewayroutetable.Setup,
		mesh.Setup,
		virtualdeployment.Setup,
		virtualservice.Setup,
		virtualserviceroutetable.Setup,
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
		secret.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
