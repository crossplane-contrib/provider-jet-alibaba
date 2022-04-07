/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	historydeliveryjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/actiontrail/historydeliveryjob"
	trail "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/actiontrail/trail"
	account "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/account"
	backuppolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/backuppolicy"
	cluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/cluster"
	connection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/connection"
	dbcluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/dbcluster"
	acl "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/acl"
	healthchecktemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/healthchecktemplate"
	listener "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/listener"
	loadbalancer "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/loadbalancer"
	rule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/rule"
	securitypolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/securitypolicy"
	servergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/servergroup"
	actiontrail "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/actiontrail"
	disk "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/disk"
	dns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/dns"
	eip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/eip"
	havip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/havip"
	image "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/image"
	instance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/instance"
	slb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/slb"
	snapshot "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/snapshot"
	subnet "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/subnet"
	vpc "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/vpc"
	vswitch "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/vswitch"
	accessstrategy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/accessstrategy"
	addresspool "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/addresspool"
	customline "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/customline"
	domain "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/domain"
	domainattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/domainattachment"
	domaingroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/domaingroup"
	gtminstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/gtminstance"
	instancealidns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/instance"
	monitorconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/monitorconfig"
	record "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/record"
	consumergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/consumergroup"
	instancealikafka "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/instance"
	saslacl "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/saslacl"
	sasluser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/sasluser"
	topic "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/topic"
	binding "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/binding"
	exchange "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/exchange"
	instanceamqp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/instance"
	queue "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/queue"
	virtualhost "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/virtualhost"
	gatewayapi "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayapi"
	gatewayapp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayapp"
	gatewayappattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayappattachment"
	gatewaygroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewaygroup"
	gatewayvpcaccess "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayvpcaccess"
	alertcontact "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/alertcontact"
	alertcontactgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/alertcontactgroup"
	dispatchrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/dispatchrule"
	prometheusalertrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/prometheusalertrule"
	provisioninggroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/auto/provisioninggroup"
	host "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/host"
	hostaccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostaccount"
	hostaccountuserattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostaccountuserattachment"
	hostaccountusergroupattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostaccountusergroupattachment"
	hostattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostattachment"
	hostgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostgroup"
	hostgroupaccountuserattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostgroupaccountuserattachment"
	hostgroupaccountusergroupattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostgroupaccountusergroupattachment"
	instancebastionhost "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/instance"
	user "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/user"
	userattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/userattachment"
	usergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/usergroup"
	industrialpidloop "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/brain/industrialpidloop"
	industrialpidorganization "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/brain/industrialpidorganization"
	industrialpidproject "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/brain/industrialpidproject"
	certificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cas/certificate"
	backupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cassandra/backupplan"
	clustercassandra "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cassandra/cluster"
	datacenter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cassandra/datacenter"
	dedicatedhost "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cddc/dedicatedhost"
	dedicatedhostaccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cddc/dedicatedhostaccount"
	dedicatedhostgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cddc/dedicatedhostgroup"
	domaincdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/domain"
	domainconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/domainconfig"
	domainnew "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/domainnew"
	realtimelogdelivery "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/realtimelogdelivery"
	bandwidthlimit "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/bandwidthlimit"
	bandwidthpackage "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/bandwidthpackage"
	bandwidthpackageattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/bandwidthpackageattachment"
	flowlog "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/flowlog"
	instancecen "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/instance"
	instanceattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/instanceattachment"
	instancegrant "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/instancegrant"
	privatezone "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/privatezone"
	routeentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/routeentry"
	routemap "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/routemap"
	routeservice "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/routeservice"
	transitrouter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouter"
	transitrouterpeerattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterpeerattachment"
	transitrouterrouteentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterrouteentry"
	transitrouterroutetable "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterroutetable"
	transitrouterroutetableassociation "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterroutetableassociation"
	transitrouterroutetablepropagation "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterroutetablepropagation"
	transitroutervbrattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitroutervbrattachment"
	transitroutervpcattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitroutervpcattachment"
	vbrhealthcheck "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/vbrhealthcheck"
	houseaccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/click/houseaccount"
	housebackuppolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/click/housebackuppolicy"
	housedbcluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/click/housedbcluster"
	connectnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/connectnetwork"
	connectnetworkattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/connectnetworkattachment"
	connectnetworkgrant "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/connectnetworkgrant"
	firewallcontrolpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/firewallcontrolpolicy"
	firewallcontrolpolicyorder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/firewallcontrolpolicyorder"
	firewallinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/firewallinstance"
	ssoaccessassignment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoaccessassignment"
	ssoaccessconfiguration "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoaccessconfiguration"
	ssoaccessconfigurationprovisioning "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoaccessconfigurationprovisioning"
	ssodirectory "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssodirectory"
	ssogroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssogroup"
	ssoscimservercredential "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoscimservercredential"
	ssouser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssouser"
	ssouserattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssouserattachment"
	storagegatewayexpresssync "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewayexpresssync"
	storagegatewayexpresssyncshareattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewayexpresssyncshareattachment"
	storagegatewaygateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygateway"
	storagegatewaygatewayblockvolume "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewayblockvolume"
	storagegatewaygatewaycachedisk "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewaycachedisk"
	storagegatewaygatewayfileshare "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewayfileshare"
	storagegatewaygatewaylogging "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewaylogging"
	storagegatewaygatewaysmbuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewaysmbuser"
	storagegatewaystoragebundle "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaystoragebundle"
	faceconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloudauth/faceconfig"
	alarm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/alarm"
	alarmcontact "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/alarmcontact"
	alarmcontactgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/alarmcontactgroup"
	dynamictaggroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/dynamictaggroup"
	groupmetricrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/groupmetricrule"
	metricruletemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/metricruletemplate"
	monitorgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/monitorgroup"
	monitorgroupinstances "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/monitorgroupinstances"
	sitemonitor "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/sitemonitor"
	bandwidthpackagecommon "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/common/bandwidthpackage"
	bandwidthpackageattachmentcommon "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/common/bandwidthpackageattachment"
	aggregatecompliancepack "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/aggregatecompliancepack"
	aggregateconfigrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/aggregateconfigrule"
	aggregator "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/aggregator"
	compliancepack "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/compliancepack"
	configurationrecorder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/configurationrecorder"
	deliverychannel "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/deliverychannel"
	ruleconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/rule"
	clustercontainer "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/container/cluster"
	imagecopy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/copy/image"
	chartnamespace "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/chartnamespace"
	chartrepository "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/chartrepository"
	eeinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eeinstance"
	eenamespace "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eenamespace"
	eerepo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eerepo"
	eesyncrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eesyncrule"
	endpointaclpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/endpointaclpolicy"
	namespace "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/namespace"
	repo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/repo"
	application "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/application"
	autoscalingconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/autoscalingconfig"
	edgekubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/edgekubernetes"
	kubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetes"
	kubernetesaddon "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetesaddon"
	kubernetesautoscaler "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetesautoscaler"
	kubernetesnodepool "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetesnodepool"
	kubernetespermissions "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetespermissions"
	managedkubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/managedkubernetes"
	serverlesskubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/serverlesskubernetes"
	swarm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/swarm"
	worksfolder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/data/worksfolder"
	gatewaygateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/database/gatewaygateway"
	project "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/datahub/project"
	subscription "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/datahub/subscription"
	topicdatahub "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/datahub/topic"
	accountdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/account"
	accountprivilege "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/accountprivilege"
	backuppolicydb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/backuppolicy"
	connectiondb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/connection"
	database "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/database"
	instancedb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/instance"
	readonlyinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/readonlyinstance"
	readwritesplittingconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/readwritesplittingconnection"
	instancedbfs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/instance"
	instanceattachmentdbfs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/instanceattachment"
	servicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/servicelinkedrole"
	snapshotdbfs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/snapshot"
	domaindcdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dcdn/domain"
	domainconfigdcdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dcdn/domainconfig"
	instanceddosbgp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddosbgp/instance"
	domainresource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/domainresource"
	instanceddoscoo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/instance"
	port "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/port"
	schedulerrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/schedulerrule"
	accessgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/accessgroup"
	accessrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/accessrule"
	filesystem "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/filesystem"
	mountpoint "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/mountpoint"
	maildomain "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/maildomain"
	mailmailaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/mailmailaddress"
	mailreceivers "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/mailreceivers"
	mailtag "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/mailtag"
	attachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/disk/attachment"
	enterpriseinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dms/enterpriseinstance"
	enterpriseuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dms/enterpriseuser"
	domaindns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/domain"
	domainattachmentdns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/domainattachment"
	group "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/group"
	instancedns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/instance"
	recorddns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/record"
	instancedrds "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/drds/instance"
	consumerchannel "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/consumerchannel"
	jobmonitorrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/jobmonitorrule"
	migrationinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/migrationinstance"
	migrationjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/migrationjob"
	subscriptionjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/subscriptionjob"
	synchronizationinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/synchronizationinstance"
	synchronizationjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/synchronizationjob"
	instanceeais "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eais/instance"
	command "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/command"
	desktop "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/desktop"
	imageecd "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/image"
	nasfilesystem "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/nasfilesystem"
	networkpackage "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/networkpackage"
	policygroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/policygroup"
	simpleofficesite "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/simpleofficesite"
	userecd "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/user"
	containergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/containergroup"
	imagecache "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/imagecache"
	openapiimagecache "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/openapiimagecache"
	virtualnode "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/virtualnode"
	keypair "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecp/keypair"
	autosnapshotpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/autosnapshotpolicy"
	autosnapshotpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/autosnapshotpolicyattachment"
	commandecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/command"
	dedicatedhostecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/dedicatedhost"
	dedicatedhostcluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/dedicatedhostcluster"
	deploymentset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/deploymentset"
	diskecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/disk"
	diskattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/diskattachment"
	hpccluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/hpccluster"
	keypairecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/keypair"
	keypairattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/keypairattachment"
	launchtemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/networkinterfaceattachment"
	prefixlist "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/prefixlist"
	sessionmanagerstatus "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/sessionmanagerstatus"
	snapshotecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/snapshot"
	storagecapacityunit "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/storagecapacityunit"
	applicationedas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/application"
	applicationdeployment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/applicationdeployment"
	applicationscale "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/applicationscale"
	clusteredas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/cluster"
	deploygroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/deploygroup"
	instanceclusterattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/instanceclusterattachment"
	k8sapplication "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/k8sapplication"
	k8scluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/k8scluster"
	slbattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/slbattachment"
	jobtemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ehpc/jobtemplate"
	address "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eip/address"
	association "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eip/association"
	anycasteipaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eipanycast/anycasteipaddress"
	anycasteipaddressattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eipanycast/anycasteipaddressattachment"
	instanceelasticsearch "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/elasticsearch/instance"
	clusteremr "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/emr/cluster"
	keypairens "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ens/keypair"
	alarmess "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/alarm"
	attachmentess "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/attachment"
	lifecyclehook "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/notification"
	scalingconfiguration "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalingconfiguration"
	scalinggroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalinggroup"
	scalinggroupvservergroups "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalinggroupvservergroups"
	scalingrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalingrule"
	schedule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/schedule"
	scheduledtask "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scheduledtask"
	bridgeeventbus "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeeventbus"
	bridgeeventsource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeeventsource"
	bridgerule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgerule"
	bridgeservicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeservicelinkedrole"
	bridgeslr "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeslr"
	connectphysicalconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/express/connectphysicalconnection"
	connectvirtualborderrouter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/express/connectvirtualborderrouter"
	alias "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/alias"
	customdomain "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/customdomain"
	function "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/function"
	functionasyncinvokeconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/functionasyncinvokeconfig"
	service "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/service"
	trigger "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/trigger"
	execution "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fnf/execution"
	flow "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fnf/flow"
	schedulefnf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fnf/schedule"
	entry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/forward/entry"
	accelerator "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/accelerator"
	aclga "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/acl"
	aclattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/aclattachment"
	additionalcertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/additionalcertificate"
	bandwidthpackagega "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/bandwidthpackage"
	bandwidthpackageattachmentga "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/bandwidthpackageattachment"
	endpointgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/endpointgroup"
	forwardingrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/forwardingrule"
	ipset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/ipset"
	listenerga "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/listener"
	accountgpdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/account"
	connectiongpdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/connection"
	elasticinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/elasticinstance"
	instancegpdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/instance"
	databasedbinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/graph/databasedbinstance"
	attachmenthavip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/havip/attachment"
	instancehbase "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbase/instance"
	ecsbackupclient "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/ecsbackupclient"
	ecsbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/ecsbackupplan"
	nasbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/nasbackupplan"
	ossbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/ossbackupplan"
	replicationvault "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/replicationvault"
	restorejob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/restorejob"
	serverbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/serverbackupplan"
	vault "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/vault"
	copy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/copy"
	export "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/export"
	sharepermission "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/sharepermission"
	projectimm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/imm/project"
	apptemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/imp/apptemplate"
	devicegroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/iot/devicegroup"
	pair "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/key/pair"
	pairattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/key/pairattachment"
	aliaskms "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/alias"
	ciphertext "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/ciphertext"
	key "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/key"
	keyversion "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/keyversion"
	secret "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/secret"
	accountkvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/account"
	auditlogconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/auditlogconfig"
	backuppolicykvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/backuppolicy"
	connectionkvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/connection"
	instancekvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/instance"
	template "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/launch/template"
	instancelindorm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/lindorm/instance"
	alert "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/alert"
	audit "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/audit"
	dashboard "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/dashboard"
	etl "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/etl"
	machinegroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/machinegroup"
	ossshipper "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/ossshipper"
	projectlog "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/project"
	store "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/store"
	storeindex "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/storeindex"
	attachmentlogtail "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/logtail/attachment"
	config "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/logtail/config"
	order "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/market/order"
	projectmaxcompute "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/maxcompute/project"
	app "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mhub/app"
	product "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mhub/product"
	queuemns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mns/queue"
	topicmns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mns/topic"
	topicsubscription "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mns/topicsubscription"
	accountmongodb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/account"
	auditpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/auditpolicy"
	instancemongodb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/instance"
	serverlessinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/serverlessinstance"
	shardinginstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/shardinginstance"
	shardingnetworkprivateaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/shardingnetworkprivateaddress"
	shardingnetworkpublicaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/shardingnetworkpublicaddress"
	subcontact "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/msc/subcontact"
	subsubscription "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/msc/subsubscription"
	subwebhook "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/msc/subwebhook"
	clustermse "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mse/cluster"
	gateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mse/gateway"
	accessgroupnas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/accessgroup"
	accessrulenas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/accessrule"
	autosnapshotpolicynas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/autosnapshotpolicy"
	dataflow "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/dataflow"
	fileset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/fileset"
	filesystemnas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/filesystem"
	lifecyclepolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/lifecyclepolicy"
	mounttarget "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/mounttarget"
	recyclebin "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/recyclebin"
	snapshotnas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/snapshot"
	gatewaynat "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nat/gateway"
	aclnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/acl"
	aclattachmentnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/aclattachment"
	aclentries "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/aclentries"
	interfaceattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/interfaceattachment"
	groupons "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ons/group"
	instanceons "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ons/instance"
	topicons "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ons/topic"
	applicationoos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/application"
	applicationgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/applicationgroup"
	executionoos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/execution"
	parameter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/parameter"
	patchbaseline "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/patchbaseline"
	secretparameter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/secretparameter"
	servicesetting "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/servicesetting"
	stateconfiguration "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/stateconfiguration"
	templateoos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/template"
	searchappgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/open/searchappgroup"
	bucket "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oss/bucket"
	bucketobject "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oss/bucketobject"
	instanceots "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ots/instance"
	instanceattachmentots "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ots/instanceattachment"
	table "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ots/table"
	accountpolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/account"
	accountprivilegepolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/accountprivilege"
	backuppolicypolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/backuppolicy"
	clusterpolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/cluster"
	databasepolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/database"
	endpoint "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/endpoint"
	endpointaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/endpointaddress"
	vpcendpoint "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpoint"
	vpcendpointconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointconnection"
	vpcendpointservice "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointservice"
	vpcendpointserviceresource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointserviceresource"
	vpcendpointserviceuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointserviceuser"
	vpcendpointzone "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointzone"
	providerconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/providerconfig"
	endpointpvtz "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/endpoint"
	rulepvtz "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/rule"
	ruleattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/ruleattachment"
	uservpcauthorization "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/uservpcauthorization"
	zone "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/zone"
	zoneattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/zoneattachment"
	zonerecord "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/zonerecord"
	biuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quick/biuser"
	applicationinfo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quotas/applicationinfo"
	quotaalarm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quotas/quotaalarm"
	quotaapplication "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quotas/quotaapplication"
	accesskey "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/accesskey"
	accountalias "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/accountpasswordpolicy"
	aliasram "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/alias"
	groupram "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/group"
	groupmembership "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/grouppolicyattachment"
	loginprofile "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/loginprofile"
	policy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/policy"
	role "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/role"
	roleattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/roleattachment"
	rolepolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/samlprovider"
	securitypreference "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/securitypreference"
	userram "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/user"
	userpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/userpolicyattachment"
	organization "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rdc/organization"
	accountrds "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/account"
	backup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/backup"
	clonedbinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/clonedbinstance"
	parametergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/parametergroup"
	upgradedbinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/upgradedbinstance"
	instancereserved "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/reserved/instance"
	manageraccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/manageraccount"
	managercontrolpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managercontrolpolicy"
	managercontrolpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managercontrolpolicyattachment"
	managerfolder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerfolder"
	managerhandshake "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerhandshake"
	managerpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerpolicy"
	managerpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerpolicyattachment"
	managerpolicyversion "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerpolicyversion"
	managerresourcedirectory "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerresourcedirectory"
	managerresourcegroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerresourcegroup"
	managerresourceshare "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerresourceshare"
	managerrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerrole"
	managerservicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerservicelinkedrole"
	managersharedresource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managersharedresource"
	managersharedtarget "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managersharedtarget"
	changeset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/changeset"
	stack "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/stack"
	stackgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/stackgroup"
	stackinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/stackinstance"
	templateros "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/template"
	templatescratch "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/templatescratch"
	entryroute "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/route/entry"
	tableroute "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/route/table"
	tableattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/route/tableattachment"
	interfaceconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/router/interfaceconnection"
	applicationsae "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/application"
	configmap "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/configmap"
	ingress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/ingress"
	namespacesae "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/namespace"
	aclsag "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/acl"
	aclrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/aclrule"
	clientuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/clientuser"
	dnatentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/dnatentry"
	qos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/qos"
	qoscar "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/qoscar"
	qospolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/qospolicy"
	snatentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/snatentry"
	domainscdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/scdn/domain"
	domainconfigscdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/scdn/domainconfig"
	configsddp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sddp/config"
	instancesddp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sddp/instance"
	rulesddp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sddp/rule"
	centergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/centergroup"
	centerservicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/centerservicelinkedrole"
	groupsecurity "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/group"
	grouprule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/grouprule"
	meshservicemesh "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/service/meshservicemesh"
	applicationservercustomimage "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationservercustomimage"
	applicationserverfirewallrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationserverfirewallrule"
	applicationserverinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationserverinstance"
	applicationserversnapshot "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationserversnapshot"
	aclslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/acl"
	attachmentslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/attachment"
	backendserver "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/backendserver"
	cacertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/cacertificate"
	domainextension "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/domainextension"
	listenerslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/listener"
	loadbalancerslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/loadbalancer"
	masterslaveservergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/masterslaveservergroup"
	ruleslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/rule"
	servercertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/servercertificate"
	servergroupslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/servergroup"
	tlscipherpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/tlscipherpolicy"
	policysnapshot "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/snapshot/policy"
	entrysnat "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/snat/entry"
	certificatesservicecertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ssl/certificatesservicecertificate"
	vpnclientcert "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ssl/vpnclientcert"
	vpnserver "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ssl/vpnserver"
	instancetsdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/tsdb/instance"
	surveillancesystemgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/video/surveillancesystemgroup"
	domainvod "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vod/domain"
	bgpgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/bgpgroup"
	bgpnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/bgpnetwork"
	bgppeer "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/bgppeer"
	dhcpoptionsset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/dhcpoptionsset"
	dhcpoptionssetattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/dhcpoptionssetattachment"
	flowlogvpc "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/flowlog"
	ipv6egressrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/ipv6egressrule"
	ipv6gateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/ipv6gateway"
	ipv6internetbandwidth "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/ipv6internetbandwidth"
	natip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/natip"
	natipcidr "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/natipcidr"
	trafficmirrorfilter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorfilter"
	trafficmirrorfilteregressrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorfilteregressrule"
	trafficmirrorfilteringressrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorfilteringressrule"
	trafficmirrorsession "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorsession"
	vbrha "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/vbrha"
	connectionvpn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/connection"
	customergateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/customergateway"
	gatewayvpn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/gateway"
	routeentryvpn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/routeentry"
	certificatewaf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/certificate"
	domainwaf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/domain"
	instancewaf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/instance"
	protectionmodule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/protectionmodule"
	bastionhostinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/yundun/bastionhostinstance"
	dbauditinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/yundun/dbauditinstance"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		historydeliveryjob.Setup,
		trail.Setup,
		account.Setup,
		backuppolicy.Setup,
		cluster.Setup,
		connection.Setup,
		dbcluster.Setup,
		acl.Setup,
		healthchecktemplate.Setup,
		listener.Setup,
		loadbalancer.Setup,
		rule.Setup,
		securitypolicy.Setup,
		servergroup.Setup,
		actiontrail.Setup,
		disk.Setup,
		dns.Setup,
		eip.Setup,
		havip.Setup,
		image.Setup,
		instance.Setup,
		slb.Setup,
		snapshot.Setup,
		subnet.Setup,
		vpc.Setup,
		vswitch.Setup,
		accessstrategy.Setup,
		addresspool.Setup,
		customline.Setup,
		domain.Setup,
		domainattachment.Setup,
		domaingroup.Setup,
		gtminstance.Setup,
		instancealidns.Setup,
		monitorconfig.Setup,
		record.Setup,
		consumergroup.Setup,
		instancealikafka.Setup,
		saslacl.Setup,
		sasluser.Setup,
		topic.Setup,
		binding.Setup,
		exchange.Setup,
		instanceamqp.Setup,
		queue.Setup,
		virtualhost.Setup,
		gatewayapi.Setup,
		gatewayapp.Setup,
		gatewayappattachment.Setup,
		gatewaygroup.Setup,
		gatewayvpcaccess.Setup,
		alertcontact.Setup,
		alertcontactgroup.Setup,
		dispatchrule.Setup,
		prometheusalertrule.Setup,
		provisioninggroup.Setup,
		host.Setup,
		hostaccount.Setup,
		hostaccountuserattachment.Setup,
		hostaccountusergroupattachment.Setup,
		hostattachment.Setup,
		hostgroup.Setup,
		hostgroupaccountuserattachment.Setup,
		hostgroupaccountusergroupattachment.Setup,
		instancebastionhost.Setup,
		user.Setup,
		userattachment.Setup,
		usergroup.Setup,
		industrialpidloop.Setup,
		industrialpidorganization.Setup,
		industrialpidproject.Setup,
		certificate.Setup,
		backupplan.Setup,
		clustercassandra.Setup,
		datacenter.Setup,
		dedicatedhost.Setup,
		dedicatedhostaccount.Setup,
		dedicatedhostgroup.Setup,
		domaincdn.Setup,
		domainconfig.Setup,
		domainnew.Setup,
		realtimelogdelivery.Setup,
		bandwidthlimit.Setup,
		bandwidthpackage.Setup,
		bandwidthpackageattachment.Setup,
		flowlog.Setup,
		instancecen.Setup,
		instanceattachment.Setup,
		instancegrant.Setup,
		privatezone.Setup,
		routeentry.Setup,
		routemap.Setup,
		routeservice.Setup,
		transitrouter.Setup,
		transitrouterpeerattachment.Setup,
		transitrouterrouteentry.Setup,
		transitrouterroutetable.Setup,
		transitrouterroutetableassociation.Setup,
		transitrouterroutetablepropagation.Setup,
		transitroutervbrattachment.Setup,
		transitroutervpcattachment.Setup,
		vbrhealthcheck.Setup,
		houseaccount.Setup,
		housebackuppolicy.Setup,
		housedbcluster.Setup,
		connectnetwork.Setup,
		connectnetworkattachment.Setup,
		connectnetworkgrant.Setup,
		firewallcontrolpolicy.Setup,
		firewallcontrolpolicyorder.Setup,
		firewallinstance.Setup,
		ssoaccessassignment.Setup,
		ssoaccessconfiguration.Setup,
		ssoaccessconfigurationprovisioning.Setup,
		ssodirectory.Setup,
		ssogroup.Setup,
		ssoscimservercredential.Setup,
		ssouser.Setup,
		ssouserattachment.Setup,
		storagegatewayexpresssync.Setup,
		storagegatewayexpresssyncshareattachment.Setup,
		storagegatewaygateway.Setup,
		storagegatewaygatewayblockvolume.Setup,
		storagegatewaygatewaycachedisk.Setup,
		storagegatewaygatewayfileshare.Setup,
		storagegatewaygatewaylogging.Setup,
		storagegatewaygatewaysmbuser.Setup,
		storagegatewaystoragebundle.Setup,
		faceconfig.Setup,
		alarm.Setup,
		alarmcontact.Setup,
		alarmcontactgroup.Setup,
		dynamictaggroup.Setup,
		groupmetricrule.Setup,
		metricruletemplate.Setup,
		monitorgroup.Setup,
		monitorgroupinstances.Setup,
		sitemonitor.Setup,
		bandwidthpackagecommon.Setup,
		bandwidthpackageattachmentcommon.Setup,
		aggregatecompliancepack.Setup,
		aggregateconfigrule.Setup,
		aggregator.Setup,
		compliancepack.Setup,
		configurationrecorder.Setup,
		deliverychannel.Setup,
		ruleconfig.Setup,
		clustercontainer.Setup,
		imagecopy.Setup,
		chartnamespace.Setup,
		chartrepository.Setup,
		eeinstance.Setup,
		eenamespace.Setup,
		eerepo.Setup,
		eesyncrule.Setup,
		endpointaclpolicy.Setup,
		namespace.Setup,
		repo.Setup,
		application.Setup,
		autoscalingconfig.Setup,
		edgekubernetes.Setup,
		kubernetes.Setup,
		kubernetesaddon.Setup,
		kubernetesautoscaler.Setup,
		kubernetesnodepool.Setup,
		kubernetespermissions.Setup,
		managedkubernetes.Setup,
		serverlesskubernetes.Setup,
		swarm.Setup,
		worksfolder.Setup,
		gatewaygateway.Setup,
		project.Setup,
		subscription.Setup,
		topicdatahub.Setup,
		accountdb.Setup,
		accountprivilege.Setup,
		backuppolicydb.Setup,
		connectiondb.Setup,
		database.Setup,
		instancedb.Setup,
		readonlyinstance.Setup,
		readwritesplittingconnection.Setup,
		instancedbfs.Setup,
		instanceattachmentdbfs.Setup,
		servicelinkedrole.Setup,
		snapshotdbfs.Setup,
		domaindcdn.Setup,
		domainconfigdcdn.Setup,
		instanceddosbgp.Setup,
		domainresource.Setup,
		instanceddoscoo.Setup,
		port.Setup,
		schedulerrule.Setup,
		accessgroup.Setup,
		accessrule.Setup,
		filesystem.Setup,
		mountpoint.Setup,
		maildomain.Setup,
		mailmailaddress.Setup,
		mailreceivers.Setup,
		mailtag.Setup,
		attachment.Setup,
		enterpriseinstance.Setup,
		enterpriseuser.Setup,
		domaindns.Setup,
		domainattachmentdns.Setup,
		group.Setup,
		instancedns.Setup,
		recorddns.Setup,
		instancedrds.Setup,
		consumerchannel.Setup,
		jobmonitorrule.Setup,
		migrationinstance.Setup,
		migrationjob.Setup,
		subscriptionjob.Setup,
		synchronizationinstance.Setup,
		synchronizationjob.Setup,
		instanceeais.Setup,
		command.Setup,
		desktop.Setup,
		imageecd.Setup,
		nasfilesystem.Setup,
		networkpackage.Setup,
		policygroup.Setup,
		simpleofficesite.Setup,
		userecd.Setup,
		containergroup.Setup,
		imagecache.Setup,
		openapiimagecache.Setup,
		virtualnode.Setup,
		keypair.Setup,
		autosnapshotpolicy.Setup,
		autosnapshotpolicyattachment.Setup,
		commandecs.Setup,
		dedicatedhostecs.Setup,
		dedicatedhostcluster.Setup,
		deploymentset.Setup,
		diskecs.Setup,
		diskattachment.Setup,
		hpccluster.Setup,
		keypairecs.Setup,
		keypairattachment.Setup,
		launchtemplate.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		prefixlist.Setup,
		sessionmanagerstatus.Setup,
		snapshotecs.Setup,
		storagecapacityunit.Setup,
		applicationedas.Setup,
		applicationdeployment.Setup,
		applicationscale.Setup,
		clusteredas.Setup,
		deploygroup.Setup,
		instanceclusterattachment.Setup,
		k8sapplication.Setup,
		k8scluster.Setup,
		slbattachment.Setup,
		jobtemplate.Setup,
		address.Setup,
		association.Setup,
		anycasteipaddress.Setup,
		anycasteipaddressattachment.Setup,
		instanceelasticsearch.Setup,
		clusteremr.Setup,
		keypairens.Setup,
		alarmess.Setup,
		attachmentess.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		scalingconfiguration.Setup,
		scalinggroup.Setup,
		scalinggroupvservergroups.Setup,
		scalingrule.Setup,
		schedule.Setup,
		scheduledtask.Setup,
		bridgeeventbus.Setup,
		bridgeeventsource.Setup,
		bridgerule.Setup,
		bridgeservicelinkedrole.Setup,
		bridgeslr.Setup,
		connectphysicalconnection.Setup,
		connectvirtualborderrouter.Setup,
		alias.Setup,
		customdomain.Setup,
		function.Setup,
		functionasyncinvokeconfig.Setup,
		service.Setup,
		trigger.Setup,
		execution.Setup,
		flow.Setup,
		schedulefnf.Setup,
		entry.Setup,
		accelerator.Setup,
		aclga.Setup,
		aclattachment.Setup,
		additionalcertificate.Setup,
		bandwidthpackagega.Setup,
		bandwidthpackageattachmentga.Setup,
		endpointgroup.Setup,
		forwardingrule.Setup,
		ipset.Setup,
		listenerga.Setup,
		accountgpdb.Setup,
		connectiongpdb.Setup,
		elasticinstance.Setup,
		instancegpdb.Setup,
		databasedbinstance.Setup,
		attachmenthavip.Setup,
		instancehbase.Setup,
		ecsbackupclient.Setup,
		ecsbackupplan.Setup,
		nasbackupplan.Setup,
		ossbackupplan.Setup,
		replicationvault.Setup,
		restorejob.Setup,
		serverbackupplan.Setup,
		vault.Setup,
		copy.Setup,
		export.Setup,
		sharepermission.Setup,
		projectimm.Setup,
		apptemplate.Setup,
		devicegroup.Setup,
		pair.Setup,
		pairattachment.Setup,
		aliaskms.Setup,
		ciphertext.Setup,
		key.Setup,
		keyversion.Setup,
		secret.Setup,
		accountkvstore.Setup,
		auditlogconfig.Setup,
		backuppolicykvstore.Setup,
		connectionkvstore.Setup,
		instancekvstore.Setup,
		template.Setup,
		instancelindorm.Setup,
		alert.Setup,
		audit.Setup,
		dashboard.Setup,
		etl.Setup,
		machinegroup.Setup,
		ossshipper.Setup,
		projectlog.Setup,
		store.Setup,
		storeindex.Setup,
		attachmentlogtail.Setup,
		config.Setup,
		order.Setup,
		projectmaxcompute.Setup,
		app.Setup,
		product.Setup,
		queuemns.Setup,
		topicmns.Setup,
		topicsubscription.Setup,
		accountmongodb.Setup,
		auditpolicy.Setup,
		instancemongodb.Setup,
		serverlessinstance.Setup,
		shardinginstance.Setup,
		shardingnetworkprivateaddress.Setup,
		shardingnetworkpublicaddress.Setup,
		subcontact.Setup,
		subsubscription.Setup,
		subwebhook.Setup,
		clustermse.Setup,
		gateway.Setup,
		accessgroupnas.Setup,
		accessrulenas.Setup,
		autosnapshotpolicynas.Setup,
		dataflow.Setup,
		fileset.Setup,
		filesystemnas.Setup,
		lifecyclepolicy.Setup,
		mounttarget.Setup,
		recyclebin.Setup,
		snapshotnas.Setup,
		gatewaynat.Setup,
		aclnetwork.Setup,
		aclattachmentnetwork.Setup,
		aclentries.Setup,
		interfaceattachment.Setup,
		groupons.Setup,
		instanceons.Setup,
		topicons.Setup,
		applicationoos.Setup,
		applicationgroup.Setup,
		executionoos.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		secretparameter.Setup,
		servicesetting.Setup,
		stateconfiguration.Setup,
		templateoos.Setup,
		searchappgroup.Setup,
		bucket.Setup,
		bucketobject.Setup,
		instanceots.Setup,
		instanceattachmentots.Setup,
		table.Setup,
		accountpolardb.Setup,
		accountprivilegepolardb.Setup,
		backuppolicypolardb.Setup,
		clusterpolardb.Setup,
		databasepolardb.Setup,
		endpoint.Setup,
		endpointaddress.Setup,
		vpcendpoint.Setup,
		vpcendpointconnection.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceresource.Setup,
		vpcendpointserviceuser.Setup,
		vpcendpointzone.Setup,
		providerconfig.Setup,
		endpointpvtz.Setup,
		rulepvtz.Setup,
		ruleattachment.Setup,
		uservpcauthorization.Setup,
		zone.Setup,
		zoneattachment.Setup,
		zonerecord.Setup,
		biuser.Setup,
		applicationinfo.Setup,
		quotaalarm.Setup,
		quotaapplication.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		aliasram.Setup,
		groupram.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		loginprofile.Setup,
		policy.Setup,
		role.Setup,
		roleattachment.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		securitypreference.Setup,
		userram.Setup,
		userpolicyattachment.Setup,
		organization.Setup,
		accountrds.Setup,
		backup.Setup,
		clonedbinstance.Setup,
		parametergroup.Setup,
		upgradedbinstance.Setup,
		instancereserved.Setup,
		manageraccount.Setup,
		managercontrolpolicy.Setup,
		managercontrolpolicyattachment.Setup,
		managerfolder.Setup,
		managerhandshake.Setup,
		managerpolicy.Setup,
		managerpolicyattachment.Setup,
		managerpolicyversion.Setup,
		managerresourcedirectory.Setup,
		managerresourcegroup.Setup,
		managerresourceshare.Setup,
		managerrole.Setup,
		managerservicelinkedrole.Setup,
		managersharedresource.Setup,
		managersharedtarget.Setup,
		changeset.Setup,
		stack.Setup,
		stackgroup.Setup,
		stackinstance.Setup,
		templateros.Setup,
		templatescratch.Setup,
		entryroute.Setup,
		tableroute.Setup,
		tableattachment.Setup,
		interfaceconnection.Setup,
		applicationsae.Setup,
		configmap.Setup,
		ingress.Setup,
		namespacesae.Setup,
		aclsag.Setup,
		aclrule.Setup,
		clientuser.Setup,
		dnatentry.Setup,
		qos.Setup,
		qoscar.Setup,
		qospolicy.Setup,
		snatentry.Setup,
		domainscdn.Setup,
		domainconfigscdn.Setup,
		configsddp.Setup,
		instancesddp.Setup,
		rulesddp.Setup,
		centergroup.Setup,
		centerservicelinkedrole.Setup,
		groupsecurity.Setup,
		grouprule.Setup,
		meshservicemesh.Setup,
		applicationservercustomimage.Setup,
		applicationserverfirewallrule.Setup,
		applicationserverinstance.Setup,
		applicationserversnapshot.Setup,
		aclslb.Setup,
		attachmentslb.Setup,
		backendserver.Setup,
		cacertificate.Setup,
		domainextension.Setup,
		listenerslb.Setup,
		loadbalancerslb.Setup,
		masterslaveservergroup.Setup,
		ruleslb.Setup,
		servercertificate.Setup,
		servergroupslb.Setup,
		tlscipherpolicy.Setup,
		policysnapshot.Setup,
		entrysnat.Setup,
		certificatesservicecertificate.Setup,
		vpnclientcert.Setup,
		vpnserver.Setup,
		instancetsdb.Setup,
		surveillancesystemgroup.Setup,
		domainvod.Setup,
		bgpgroup.Setup,
		bgpnetwork.Setup,
		bgppeer.Setup,
		dhcpoptionsset.Setup,
		dhcpoptionssetattachment.Setup,
		flowlogvpc.Setup,
		ipv6egressrule.Setup,
		ipv6gateway.Setup,
		ipv6internetbandwidth.Setup,
		natip.Setup,
		natipcidr.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilteregressrule.Setup,
		trafficmirrorfilteringressrule.Setup,
		trafficmirrorsession.Setup,
		vbrha.Setup,
		connectionvpn.Setup,
		customergateway.Setup,
		gatewayvpn.Setup,
		routeentryvpn.Setup,
		certificatewaf.Setup,
		domainwaf.Setup,
		instancewaf.Setup,
		protectionmodule.Setup,
		bastionhostinstance.Setup,
		dbauditinstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
