// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backup "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/database/backup"
	backupschedule "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/database/backupschedule"
	cluster "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/database/cluster"
	instance "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/database/instance"
	user "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/database/user"
	dnsrecord "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/dns/dnsrecord"
	firewall "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/firewall/firewall"
	rule "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/firewall/rule"
	clusterk8s "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/k8s/cluster"
	nodegroup "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/k8s/nodegroup"
	project "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/project/project"
	providerconfig "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/providerconfig"
	bucket "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/s3/bucket"
	bucketsubdomain "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/s3/bucketsubdomain"
	disk "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/server/disk"
	diskbackupschedule "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/server/diskbackupschedule"
	server "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/server/server"
	serverip "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/server/serverip"
	sshkey "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/ssh/sshkey"
	drive "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/storage/drive"
	floatingip "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/vpc/floatingip"
	vpc "github.com/lebedevdsl/crossplane-upjet-provider-timeweb/internal/controller/cluster/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.Setup,
		backupschedule.Setup,
		cluster.Setup,
		instance.Setup,
		user.Setup,
		dnsrecord.Setup,
		firewall.Setup,
		rule.Setup,
		clusterk8s.Setup,
		nodegroup.Setup,
		project.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketsubdomain.Setup,
		disk.Setup,
		diskbackupschedule.Setup,
		server.Setup,
		serverip.Setup,
		sshkey.Setup,
		drive.Setup,
		floatingip.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.SetupGated,
		backupschedule.SetupGated,
		cluster.SetupGated,
		instance.SetupGated,
		user.SetupGated,
		dnsrecord.SetupGated,
		firewall.SetupGated,
		rule.SetupGated,
		clusterk8s.SetupGated,
		nodegroup.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
		bucket.SetupGated,
		bucketsubdomain.SetupGated,
		disk.SetupGated,
		diskbackupschedule.SetupGated,
		server.SetupGated,
		serverip.SetupGated,
		sshkey.SetupGated,
		drive.SetupGated,
		floatingip.SetupGated,
		vpc.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
