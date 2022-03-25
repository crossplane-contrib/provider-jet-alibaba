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

	projectimm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/imm/project"
managersharedresource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managersharedresource"
hostattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostattachment"
attachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/disk/attachment"
connectvirtualborderrouter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/express/connectvirtualborderrouter"
storagegatewaygatewayfileshare "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewayfileshare"
accessrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/accessrule"
mailreceivers "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/mailreceivers"
anycasteipaddressattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eipanycast/anycasteipaddressattachment"
applicationoos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/application"
accessstrategy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/accessstrategy"
domaingroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/domaingroup"
vbrhealthcheck "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/vbrhealthcheck"
vpnserver "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ssl/vpnserver"
accountprivilege "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/accountprivilege"
domainresource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/domainresource"
instanceeais "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eais/instance"
instanceecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/instance"
launchtemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/launchtemplate"
havip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/havip"
transitrouterroutetableassociation "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterroutetableassociation"
aggregatecompliancepack "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/aggregatecompliancepack"
keyversion "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/keyversion"
groupmembership "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/groupmembership"
backendserver "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/backendserver"
chartrepository "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/chartrepository"
additionalcertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/additionalcertificate"
endpointaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/endpointaddress"
gatewayapp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayapp"
datacenter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cassandra/datacenter"
faceconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloudauth/faceconfig"
ruleattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/ruleattachment"
parametergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/parametergroup"
domainattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/domainattachment"
topicmns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mns/topic"
shardingnetworkprivateaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/shardingnetworkprivateaddress"
execution "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fnf/execution"
endpointgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/endpointgroup"
accountalias "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/accountalias"
account "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/account"
exchange "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/exchange"
accountdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/account"
domainscdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/scdn/domain"
transitrouterroutetablepropagation "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterroutetablepropagation"
slbattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/slbattachment"
servicesetting "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/servicesetting"
cluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/cluster"
aclsag "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/acl"
instanceattachmentdbfs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/instanceattachment"
imagecache "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/imagecache"
sessionmanagerstatus "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/sessionmanagerstatus"
grouppolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/grouppolicyattachment"
disk "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/disk"
snapshot "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/snapshot"
instancekvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/instance"
aggregator "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/aggregator"
database "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/database"
networkinterfaceattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/networkinterfaceattachment"
addresspool "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/addresspool"
storagegatewayexpresssync "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewayexpresssync"
aggregateconfigrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/aggregateconfigrule"
dedicatedhost "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cddc/dedicatedhost"
instancedb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/instance"
managerservicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerservicelinkedrole"
queuemns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mns/queue"
parameter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/parameter"
auditpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/auditpolicy"
centerservicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/centerservicelinkedrole"
queue "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/queue"
routeservice "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/routeservice"
dynamictaggroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/dynamictaggroup"
aliasram "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/alias"
groupsecurity "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/group"
industrialpidproject "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/brain/industrialpidproject"
clustercontainer "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/container/cluster"
servicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/servicelinkedrole"
firewallinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/firewallinstance"
apptemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/imp/apptemplate"
providerconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/providerconfig"
managerpolicyversion "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerpolicyversion"
acl "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/acl"
topicdatahub "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/datahub/topic"
table "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ots/table"
alarmess "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/alarm"
bridgerule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgerule"
instancegpdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/instance"
interfaceattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/interfaceattachment"
stackinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/stackinstance"
metricruletemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/metricruletemplate"
autoscalingconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/autoscalingconfig"
clusteredas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/cluster"
interfacerouter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/router/interface"
domainvod "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vod/domain"
protectionmodule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/protectionmodule"
storagegatewayexpresssyncshareattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewayexpresssyncshareattachment"
accessgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/accessgroup"
anycasteipaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eipanycast/anycasteipaddress"
audit "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/audit"
managercontrolpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managercontrolpolicy"
loadbalancerslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/loadbalancer"
hostgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostgroup"
readonlyinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/readonlyinstance"
trigger "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/trigger"
k8scluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/k8scluster"
quotaalarm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quotas/quotaalarm"
deliverychannel "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/deliverychannel"
dedicatedhostcluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/dedicatedhostcluster"
deploygroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/deploygroup"
autosnapshotpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/autosnapshotpolicy"
role "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/role"
masterslaveservergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/masterslaveservergroup"
bandwidthlimit "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/bandwidthlimit"
trail "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/actiontrail/trail"
listener "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/listener"
topic "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/topic"
bandwidthpackagega "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/bandwidthpackage"
project "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/datahub/project"
connectionvpn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/connection"
schedule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/schedule"
instancehbase "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbase/instance"
import "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/import"
zonerecord "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/zonerecord"
vpnclientcert "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ssl/vpnclientcert"
connectnetworkgrant "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/connectnetworkgrant"
worksfolder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/data/worksfolder"
readwritesplittingconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/readwritesplittingconnection"
instanceddoscoo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/instance"
containergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/containergroup"
loginprofile "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/loginprofile"
applicationserverfirewallrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationserverfirewallrule"
transitroutervpcattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitroutervpcattachment"
storagegatewaygatewaycachedisk "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewaycachedisk"
swarm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/swarm"
pair "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/key/pair"
accesskey "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/accesskey"
backup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/backup"
clonedbinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/clonedbinstance"
vpc "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/vpc"
hostaccountusergroupattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostaccountusergroupattachment"
bridgeeventbus "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeeventbus"
ossbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/ossbackupplan"
consumerchannel "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/consumerchannel"
bridgeslr "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeslr"
instanceattachmentots "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ots/instanceattachment"
zoneattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/zoneattachment"
storagegatewaygateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygateway"
repo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/repo"
schedulerrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/schedulerrule"
service "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/service"
attachmentlogtail "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/logtail/attachment"
accountprivilegepolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/accountprivilege"
certificatewaf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/certificate"
bandwidthpackageattachmentcommon "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/common/bandwidthpackageattachment"
eeinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eeinstance"
kubernetesautoscaler "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetesautoscaler"
managerpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerpolicyattachment"
connectnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/connectnetwork"
subscriptionjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/subscriptionjob"
userpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/userpolicyattachment"
vpcendpointserviceresource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointserviceresource"
userram "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/user"
vbrha "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/vbrha"
aclattachmentnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/aclattachment"
instancebastionhost "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/instance"
k8sapplication "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/k8sapplication"
etl "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/etl"
mounttarget "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/mounttarget"
vpcendpointserviceuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointserviceuser"
dhcpoptionsset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/dhcpoptionsset"
flowlogvpc "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/flowlog"
instancealikafka "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/instance"
instanceelasticsearch "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/elasticsearch/instance"
template "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/launch/template"
instancedbfs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/instance"
openapiimagecache "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/openapiimagecache"
connectiongpdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/connection"
export "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/export"
instancelindorm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/lindorm/instance"
snapshotnas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/snapshot"
subcontact "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/msc/subcontact"
topicons "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ons/topic"
managerrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerrole"
applicationservercustomimage "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationservercustomimage"
applicationserversnapshot "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationserversnapshot"
managedkubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/managedkubernetes"
aclga "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/acl"
topicsubscription "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mns/topicsubscription"
attachmentslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/attachment"
gatewayapi "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayapi"
app "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mhub/app"
interfaceconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/router/interfaceconnection"
managerresourceshare "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerresourceshare"
templateros "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/template"
dispatchrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/dispatchrule"
aliaskms "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/alias"
instanceots "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ots/instance"
vpcendpointservice "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointservice"
stackgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/stackgroup"
meshservicemesh "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/service/meshservicemesh"
transitrouterroutetable "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterroutetable"
enterpriseinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dms/enterpriseinstance"
autosnapshotpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/autosnapshotpolicyattachment"
ruleslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/rule"
ruleconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/rule"
eerepo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eerepo"
sharepermission "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/sharepermission"
instanceamqp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/instance"
monitorgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/monitorgroup"
instancedns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/instance"
prefixlist "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/prefixlist"
serverbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/serverbackupplan"
secretparameter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/secretparameter"
ssodirectory "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssodirectory"
eesyncrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eesyncrule"
userecd "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/user"
ecsbackupclient "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/ecsbackupclient"
accessgroupnas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/accessgroup"
bucketobject "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oss/bucketobject"
configsddp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sddp/config"
certificatesservicecertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ssl/certificatesservicecertificate"
subnet "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/subnet"
transitrouter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouter"
scalingconfiguration "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalingconfiguration"
imageecd "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/image"
configmap "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/configmap"
namespacesae "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/namespace"
aclnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/acl"
zone "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/zone"
record "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/record"
instancecen "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/instance"
migrationjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/migrationjob"
simpleofficesite "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/simpleofficesite"
applicationinfo "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quotas/applicationinfo"
shardinginstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/shardinginstance"
instanceons "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ons/instance"
backuppolicypolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/backuppolicy"
rulesddp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sddp/rule"
ssoaccessassignment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoaccessassignment"
snapshotecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/snapshot"
aclattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/aclattachment"
instancereserved "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/reserved/instance"
commandecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/command"
groupons "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ons/group"
groupram "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/group"
accountrds "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/account"
snatentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/snatentry"
accountgpdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/account"
managerfolder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerfolder"
gatewayvpcaccess "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayvpcaccess"
kubernetesnodepool "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetesnodepool"
filesystem "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/filesystem"
accelerator "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/accelerator"
tableattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/route/tableattachment"
dashboard "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/dashboard"
lifecyclepolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/lifecyclepolicy"
routeentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/routeentry"
jobmonitorrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/jobmonitorrule"
notification "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/notification"
lifecyclehook "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/lifecyclehook"
provisioninggroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/auto/provisioninggroup"
migrationinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/migrationinstance"
nasfilesystem "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/nasfilesystem"
industrialpidorganization "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/brain/industrialpidorganization"
transitrouterpeerattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterpeerattachment"
maildomain "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/maildomain"
policygroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/policygroup"
keypairattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/keypairattachment"
patchbaseline "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/patchbaseline"
firewallcontrolpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/firewallcontrolpolicy"
ssogroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssogroup"
machinegroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/machinegroup"
bridgeservicelinkedrole "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeservicelinkedrole"
clusterpolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/cluster"
clientuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/clientuser"
realtimelogdelivery "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/realtimelogdelivery"
transitrouterrouteentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitrouterrouteentry"
monitorgroupinstances "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/monitorgroupinstances"
securitypolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/securitypolicy"
domainwaf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/domain"
compliancepack "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/compliancepack"
attachmentess "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/attachment"
customdomain "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/customdomain"
rule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/rule"
saslacl "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/saslacl"
sitemonitor "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/sitemonitor"
managercontrolpolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managercontrolpolicyattachment"
domainattachmentdns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/domainattachment"
ossshipper "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/ossshipper"
bucket "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oss/bucket"
customergateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/customergateway"
hostgroupaccountusergroupattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostgroupaccountusergroupattachment"
functionasyncinvokeconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/functionasyncinvokeconfig"
connectionkvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/connection"
instancetsdb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/tsdb/instance"
domaindns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/domain"
alias "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/alias"
vpcendpointconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointconnection"
ciphertext "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/ciphertext"
projectmaxcompute "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/maxcompute/project"
hostaccountuserattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostaccountuserattachment"
ssouserattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssouserattachment"
keypairens "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ens/keypair"
association "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eip/association"
gatewaynat "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nat/gateway"
ipv6gateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/ipv6gateway"
manageraccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/manageraccount"
centergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/centergroup"
cacertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/cacertificate"
policysnapshot "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/snapshot/policy"
configurationrecorder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/config/configurationrecorder"
mountpoint "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dfs/mountpoint"
recorddns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/record"
gatewayappattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewayappattachment"
storagecapacityunit "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/storagecapacityunit"
function "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fc/function"
scalingrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalingrule"
changeset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/changeset"
domaincdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/domain"
application "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/application"
address "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eip/address"
monitorconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/monitorconfig"
storagegatewaystoragebundle "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaystoragebundle"
roleattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/roleattachment"
domainextension "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/domainextension"
instanceclusterattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/instanceclusterattachment"
accountkvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/account"
order "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/market/order"
natip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/natip"
routeentryvpn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/routeentry"
applicationgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/applicationgroup"
grouprule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/security/grouprule"
customline "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/customline"
userattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/userattachment"
housebackuppolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/click/housebackuppolicy"
mailtag "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/mailtag"
instancemongodb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/instance"
aclslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/acl"
alertcontactgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/alertcontactgroup"
ssoaccessconfigurationprovisioning "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoaccessconfigurationprovisioning"
mailmailaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/direct/mailmailaddress"
diskecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/disk"
stateconfiguration "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/stateconfiguration"
accountpasswordpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/accountpasswordpolicy"
policy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/policy"
prometheusalertrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/prometheusalertrule"
scheduledtask "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scheduledtask"
quotaapplication "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quotas/quotaapplication"
dnatentry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/dnatentry"
instancegrant "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/instancegrant"
serverlesskubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/serverlesskubernetes"
aclrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/aclrule"
databasedbinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/graph/databasedbinstance"
attachmenthavip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/havip/attachment"
upgradedbinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rds/upgradedbinstance"
qos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/qos"
kubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetes"
scalinggroupvservergroups "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalinggroupvservergroups"
devicegroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/iot/devicegroup"
clustercassandra "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cassandra/cluster"
command "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/command"
ecsbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/ecsbackupplan"
domainconfigscdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/scdn/domainconfig"
virtualhost "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/virtualhost"
imagecopy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/copy/image"
synchronizationinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/synchronizationinstance"
ipset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/ipset"
key "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/key"
vpcendpoint "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpoint"
actiontrail "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/actiontrail"
keypair "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecp/keypair"
dedicatedhostecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/dedicatedhost"
dedicatedhostgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cddc/dedicatedhostgroup"
alarmcontact "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/alarmcontact"
restorejob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/restorejob"
image "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/image"
bandwidthpackagecommon "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/common/bandwidthpackage"
bandwidthpackageattachmentga "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/bandwidthpackageattachment"
networkpackage "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/networkpackage"
scalinggroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ess/scalinggroup"
searchappgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/open/searchappgroup"
biuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/quick/biuser"
instancesddp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sddp/instance"
slb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/slb"
instance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/instance"
chartnamespace "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/chartnamespace"
ipv6internetbandwidth "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/ipv6internetbandwidth"
databasepolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/database"
instancewaf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/waf/instance"
servergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/servergroup"
entry "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/forward/entry"
auditlogconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/auditlogconfig"
transitroutervbrattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/transitroutervbrattachment"
storagegatewaygatewayblockvolume "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewayblockvolume"
rulepvtz "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/rule"
entryroute "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/route/entry"
bandwidthpackageattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/bandwidthpackageattachment"
houseaccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/click/houseaccount"
fileset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/fileset"
trafficmirrorfilter "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorfilter"
desktop "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecd/desktop"
filesystemnas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/filesystem"
uservpcauthorization "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/uservpcauthorization"
ssoscimservercredential "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoscimservercredential"
alarmcontactgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/alarmcontactgroup"
networkinterface "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/networkinterface"
store "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/store"
stack "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/stack"
loadbalancer "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/loadbalancer"
consumergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/consumergroup"
firewallcontrolpolicyorder "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/firewallcontrolpolicyorder"
qoscar "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/qoscar"
sasluser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alikafka/sasluser"
hostaccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostaccount"
applicationserverinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/simple/applicationserverinstance"
executionoos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/execution"
dbauditinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/yundun/dbauditinstance"
industrialpidloop "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/brain/industrialpidloop"
certificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cas/certificate"
vault "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/vault"
projectlog "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/project"
endpointpvtz "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/pvtz/endpoint"
samlprovider "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/samlprovider"
ingress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/ingress"
hostgroupaccountuserattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/hostgroupaccountuserattachment"
usergroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/usergroup"
forwardingrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/forwardingrule"
bgppeer "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/bgppeer"
natipcidr "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/natipcidr"
config "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/logtail/config"
serverlessinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/serverlessinstance"
surveillancesystemgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/video/surveillancesystemgroup"
networkinterface "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/interface"
eip "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/eip"
listenerga "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ga/listener"
elasticinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/gpdb/elasticinstance"
namespace "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/namespace"
trafficmirrorsession "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorsession"
product "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mhub/product"
managersharedtarget "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managersharedtarget"
alertcontact "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/arms/alertcontact"
enterpriseuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dms/enterpriseuser"
storeindex "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/storeindex"
accessrulenas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/accessrule"
organization "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/rdc/organization"
routemap "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/routemap"
housedbcluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/click/housedbcluster"
port "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddoscoo/port"
ssouser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssouser"
backuppolicydb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/backuppolicy"
gateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mse/gateway"
bgpgroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/bgpgroup"
diskattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/diskattachment"
jobtemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ehpc/jobtemplate"
pairattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/key/pairattachment"
historydeliveryjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/actiontrail/historydeliveryjob"
gatewaygroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/api/gatewaygroup"
ssoaccessconfiguration "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/ssoaccessconfiguration"
instanceddosbgp "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ddosbgp/instance"
alert "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/log/alert"
vpcendpointzone "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/privatelink/vpcendpointzone"
applicationscale "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/applicationscale"
flow "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fnf/flow"
secret "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kms/secret"
healthchecktemplate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alb/healthchecktemplate"
binding "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/amqp/binding"
storagegatewaygatewaylogging "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewaylogging"
kubernetespermissions "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetespermissions"
tableroute "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/route/table"
entrysnat "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/snat/entry"
group "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dns/group"
domainconfig "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/domainconfig"
flowlog "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/flowlog"
connectiondb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/db/connection"
tlscipherpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/tlscipherpolicy"
trafficmirrorfilteringressrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorfilteringressrule"
gatewayvpn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpn/gateway"
subscription "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/datahub/subscription"
autosnapshotpolicynas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/autosnapshotpolicy"
templateoos "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/oos/template"
connectnetworkattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/connectnetworkattachment"
schedulefnf "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/fnf/schedule"
shardingnetworkpublicaddress "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/shardingnetworkpublicaddress"
subsubscription "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/msc/subsubscription"
qospolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sag/qospolicy"
connection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/connection"
deploymentset "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/deploymentset"
applicationedas "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/application"
user "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/user"
managerresourcedirectory "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerresourcedirectory"
dhcpoptionssetattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/dhcpoptionssetattachment"
endpointaclpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/endpointaclpolicy"
applicationdeployment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/edas/applicationdeployment"
accountmongodb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mongodb/account"
groupmetricrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/groupmetricrule"
domainconfigdcdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dcdn/domainconfig"
synchronizationjob "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dts/synchronizationjob"
eenamespace "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cr/eenamespace"
kubernetesaddon "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/kubernetesaddon"
hpccluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/hpccluster"
recyclebin "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/recyclebin"
aclentries "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/network/aclentries"
dbcluster "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/dbcluster"
vswitch "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/vswitch"
domain "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/domain"
bgpnetwork "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/bgpnetwork"
accountpolardb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/account"
endpoint "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/polardb/endpoint"
securitypreference "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/securitypreference"
managerpolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerpolicy"
managerresourcegroup "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerresourcegroup"
dns "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alicloud/dns"
dedicatedhostaccount "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cddc/dedicatedhostaccount"
virtualnode "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/eci/virtualnode"
listenerslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/listener"
privatezone "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/privatezone"
gatewaygateway "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/database/gatewaygateway"
managerhandshake "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/resource/managerhandshake"
ipv6egressrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/ipv6egressrule"
trafficmirrorfilteregressrule "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/vpc/trafficmirrorfilteregressrule"
snapshotdbfs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dbfs/snapshot"
domaindcdn "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/dcdn/domain"
bridgeeventsource "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/event/bridgeeventsource"
nasbackupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/nasbackupplan"
edgekubernetes "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cs/edgekubernetes"
keypairecs "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ecs/keypair"
connectphysicalconnection "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/express/connectphysicalconnection"
bandwidthpackage "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/bandwidthpackage"
alarm "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cms/alarm"
backuppolicykvstore "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/kvstore/backuppolicy"
domainnew "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cdn/domainnew"
clustermse "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/mse/cluster"
applicationsae "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/sae/application"
host "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/bastionhost/host"
storagegatewaygatewaysmbuser "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cloud/storagegatewaygatewaysmbuser"
instancedrds "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/drds/instance"
servergroupslb "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/servergroup"
gtminstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/alidns/gtminstance"
clusteremr "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/emr/cluster"
dataflow "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/nas/dataflow"
templatescratch "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ros/templatescratch"
servercertificate "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/slb/servercertificate"
replicationvault "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/hbr/replicationvault"
copy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/image/copy"
subwebhook "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/msc/subwebhook"
rolepolicyattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/ram/rolepolicyattachment"
bastionhostinstance "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/yundun/bastionhostinstance"
backuppolicy "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/adb/backuppolicy"
backupplan "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cassandra/backupplan"
instanceattachment "github.com/crossplane-contrib/provider-jet-alibaba/internal/controller/cen/instanceattachment"

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
		slb.Setup,
		snapshot.Setup,
		subnet.Setup,
		vswitch.Setup,
		accessstrategy.Setup,
		addresspool.Setup,
		customline.Setup,
		domain.Setup,
		domainattachment.Setup,
		domaingroup.Setup,
		gtminstance.Setup,
		instance.Setup,
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
		instanceecs.Setup,
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
		import.Setup,
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
		interface.Setup,
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
		interfacerouter.Setup,
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
		vpc.Setup,
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
