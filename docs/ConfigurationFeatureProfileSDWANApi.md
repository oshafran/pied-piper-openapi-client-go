# \ConfigurationFeatureProfileSDWANApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAaaProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateAaaProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/aaa | 
[**CreateBannerProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateBannerProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/banner | 
[**CreateBasicProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateBasicProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/basic | 
[**CreateBfdProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateBfdProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/bfd | 
[**CreateCellularControllerAndCellularProfileParcelAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateCellularControllerAndCellularProfileParcelAssociationForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId}/cellular-profile | 
[**CreateCellularControllerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateCellularControllerProfileParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller | 
[**CreateCellularProfileProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateCellularProfileProfileParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/cellular-profile | 
[**CreateGlobalProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateGlobalProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/global | 
[**CreateLanVpnAndRoutingBgpParcelAssociationForService**](ConfigurationFeatureProfileSDWANApi.md#CreateLanVpnAndRoutingBgpParcelAssociationForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/bgp | 
[**CreateLanVpnAndRoutingOspfParcelAssociationForService**](ConfigurationFeatureProfileSDWANApi.md#CreateLanVpnAndRoutingOspfParcelAssociationForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/ospf | 
[**CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnParcelId}/interface/ethernet/{ethernetId}/tracker | 
[**CreateLanVpnInterfaceEthernetParcelForService**](ConfigurationFeatureProfileSDWANApi.md#CreateLanVpnInterfaceEthernetParcelForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet | 
[**CreateLanVpnProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#CreateLanVpnProfileParcelForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn | 
[**CreateLoggingProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateLoggingProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/logging | 
[**CreateManagementVpnInterfaceEthernetParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateManagementVpnInterfaceEthernetParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId}/interface/ethernet | 
[**CreateManagementVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateManagementVpnProfileParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn | 
[**CreateNtpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateNtpProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/ntp | 
[**CreateOmpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#CreateOmpProfileParcelForSystem) | **Post** /v1/feature-profile/sdwan/system/{systemId}/omp | 
[**CreateRoutingBgpProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#CreateRoutingBgpProfileParcelForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/routing/bgp | 
[**CreateRoutingOspfProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#CreateRoutingOspfProfileParcelForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/routing/ospf | 
[**CreateSdwanConfigProfileParcelForCli**](ConfigurationFeatureProfileSDWANApi.md#CreateSdwanConfigProfileParcelForCli) | **Post** /v1/feature-profile/sdwan/cli/{cliId}/config | 
[**CreateSdwanFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#CreateSdwanFeatureProfile) | **Post** /v1/feature-profile/sdwan/cli | 
[**CreateSdwanServiceFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#CreateSdwanServiceFeatureProfile) | **Post** /v1/feature-profile/sdwan/service | 
[**CreateSdwanSystemFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#CreateSdwanSystemFeatureProfile) | **Post** /v1/feature-profile/sdwan/system | 
[**CreateSdwanTransportFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#CreateSdwanTransportFeatureProfile) | **Post** /v1/feature-profile/sdwan/transport | 
[**CreateTrackerProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#CreateTrackerProfileParcelForService) | **Post** /v1/feature-profile/sdwan/service/{serviceId}/tracker | 
[**CreateTrackerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateTrackerProfileParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/tracker | 
[**CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnParcelId}/interface/ethernet/{ethernetId}/tracker | 
[**CreateWanVpnInterfaceEthernetParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateWanVpnInterfaceEthernetParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet | 
[**CreateWanVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#CreateWanVpnProfileParcelForTransport) | **Post** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn | 
[**DeleteAaaProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteAaaProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/aaa/{aaaId} | 
[**DeleteBannerProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteBannerProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/banner/{bannerId} | 
[**DeleteBasicProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteBasicProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/basic/{basicId} | 
[**DeleteBfdProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteBfdProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/bfd/{bfdId} | 
[**DeleteCellularControllerAndCellularProfileAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteCellularControllerAndCellularProfileAssociationForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId}/cellular-profile/{cellularProfileId} | 
[**DeleteCellularControllerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteCellularControllerProfileParcelForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId} | 
[**DeleteCellularProfileProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteCellularProfileProfileParcelForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/cellular-profile/{cellularProfileId} | 
[**DeleteConfigProfileParcelForCLI**](ConfigurationFeatureProfileSDWANApi.md#DeleteConfigProfileParcelForCLI) | **Delete** /v1/feature-profile/sdwan/cli/{cliId}/config/{configId} | 
[**DeleteGlobalProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteGlobalProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/global/{globalId} | 
[**DeleteLanVpnAndRoutingBgpAssociationForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteLanVpnAndRoutingBgpAssociationForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/bgp/{bgpId} | 
[**DeleteLanVpnAndRoutingOspfAssociationForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteLanVpnAndRoutingOspfAssociationForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/ospf/{ospfId} | 
[**DeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransport) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker/{trackerId} | 
[**DeleteLanVpnInterfaceEthernetForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteLanVpnInterfaceEthernetForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**DeleteLanVpnProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteLanVpnProfileParcelForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId} | 
[**DeleteLoggingProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteLoggingProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/logging/{loggingId} | 
[**DeleteManagementVpnInterfaceEthernetForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteManagementVpnInterfaceEthernetForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**DeleteManagementVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteManagementVpnProfileParcelForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId} | 
[**DeleteNtpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteNtpProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/ntp/{ntpId} | 
[**DeleteOmpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#DeleteOmpProfileParcelForSystem) | **Delete** /v1/feature-profile/sdwan/system/{systemId}/omp/{ompId} | 
[**DeleteRoutingBgpProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteRoutingBgpProfileParcelForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/routing/bgp/{bgpId} | 
[**DeleteRoutingOspfProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteRoutingOspfProfileParcelForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/routing/ospf/{ospfId} | 
[**DeleteSdwanFeatureProfileForCli**](ConfigurationFeatureProfileSDWANApi.md#DeleteSdwanFeatureProfileForCli) | **Delete** /v1/feature-profile/sdwan/cli/{cliId} | 
[**DeleteSdwanServiceFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#DeleteSdwanServiceFeatureProfile) | **Delete** /v1/feature-profile/sdwan/service/{serviceId} | 
[**DeleteSdwanSystemFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#DeleteSdwanSystemFeatureProfile) | **Delete** /v1/feature-profile/sdwan/system/{systemId} | 
[**DeleteSdwanTransportFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#DeleteSdwanTransportFeatureProfile) | **Delete** /v1/feature-profile/sdwan/transport/{transportId} | 
[**DeleteTrackerProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#DeleteTrackerProfileParcelForService) | **Delete** /v1/feature-profile/sdwan/service/{serviceId}/tracker/{trackerId} | 
[**DeleteTrackerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteTrackerProfileParcelForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/tracker/{trackerId} | 
[**DeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker/{trackerId} | 
[**DeleteWanVpnInterfaceEthernetForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteWanVpnInterfaceEthernetForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**DeleteWanVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#DeleteWanVpnProfileParcelForTransport) | **Delete** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId} | 
[**EditAaaProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditAaaProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/aaa/{aaaId} | 
[**EditBannerProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditBannerProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/banner/{bannerId} | 
[**EditBasicProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditBasicProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/basic/{basicId} | 
[**EditBfdProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditBfdProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/bfd/{bfdId} | 
[**EditCellularControllerAndCellularProfileParcelAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditCellularControllerAndCellularProfileParcelAssociationForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId}/cellular-profile/{cellularProfileId} | 
[**EditCellularControllerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditCellularControllerProfileParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId} | 
[**EditCellularProfileProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditCellularProfileProfileParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/cellular-profile/{cellularProfileId} | 
[**EditConfigProfileParcelForCLI**](ConfigurationFeatureProfileSDWANApi.md#EditConfigProfileParcelForCLI) | **Put** /v1/feature-profile/sdwan/cli/{cliId}/config/{configId} | 
[**EditGlobalProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditGlobalProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/global/{globalId} | 
[**EditLanVpnAndRoutingBgpParcelAssociationForService**](ConfigurationFeatureProfileSDWANApi.md#EditLanVpnAndRoutingBgpParcelAssociationForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/bgp/{bgpId} | 
[**EditLanVpnAndRoutingOspfParcelAssociationForService**](ConfigurationFeatureProfileSDWANApi.md#EditLanVpnAndRoutingOspfParcelAssociationForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/ospf/{ospfId} | 
[**EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker/{trackerId} | 
[**EditLanVpnInterfaceEthernetParcelForService**](ConfigurationFeatureProfileSDWANApi.md#EditLanVpnInterfaceEthernetParcelForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**EditLanVpnProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#EditLanVpnProfileParcelForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId} | 
[**EditLoggingProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditLoggingProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/logging/{loggingId} | 
[**EditManagementVpnInterfaceEthernetParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditManagementVpnInterfaceEthernetParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**EditManagementVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditManagementVpnProfileParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId} | 
[**EditNtpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditNtpProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/ntp/{ntpId} | 
[**EditOmpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#EditOmpProfileParcelForSystem) | **Put** /v1/feature-profile/sdwan/system/{systemId}/omp/{ompId} | 
[**EditRoutingBgpProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#EditRoutingBgpProfileParcelForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/routing/bgp/{bgpId} | 
[**EditRoutingOspfProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#EditRoutingOspfProfileParcelForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/routing/ospf/{ospfId} | 
[**EditSdwanFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#EditSdwanFeatureProfile) | **Put** /v1/feature-profile/sdwan/cli/{cliId} | 
[**EditSdwanServiceFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#EditSdwanServiceFeatureProfile) | **Put** /v1/feature-profile/sdwan/service/{serviceId} | 
[**EditSdwanSystemFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#EditSdwanSystemFeatureProfile) | **Put** /v1/feature-profile/sdwan/system/{systemId} | 
[**EditSdwanTransportFeatureProfile**](ConfigurationFeatureProfileSDWANApi.md#EditSdwanTransportFeatureProfile) | **Put** /v1/feature-profile/sdwan/transport/{transportId} | 
[**EditTrackerProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#EditTrackerProfileParcelForService) | **Put** /v1/feature-profile/sdwan/service/{serviceId}/tracker/{trackerId} | 
[**EditTrackerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditTrackerProfileParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/tracker/{trackerId} | 
[**EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker/{trackerId} | 
[**EditWanVpnInterfaceEthernetParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditWanVpnInterfaceEthernetParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**EditWanVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#EditWanVpnProfileParcelForTransport) | **Put** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId} | 
[**GetAaaProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetAaaProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/aaa/{aaaId} | 
[**GetAaaProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetAaaProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/aaa | 
[**GetBannerProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetBannerProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/banner/{bannerId} | 
[**GetBannerProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetBannerProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/banner | 
[**GetBasicProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetBasicProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/basic/{basicId} | 
[**GetBasicProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetBasicProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/basic | 
[**GetBfdProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetBfdProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/bfd/{bfdId} | 
[**GetBfdProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetBfdProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/bfd | 
[**GetCedgeSystemGlobalParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetCedgeSystemGlobalParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/global/schema | 
[**GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId}/cellular-profile/{cellularProfileId} | 
[**GetCellularControllerAssociatedCellularProfileParcelsForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetCellularControllerAssociatedCellularProfileParcelsForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId}/cellular-profile | 
[**GetCellularControllerProfileParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetCellularControllerProfileParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller/{cellularControllerId} | 
[**GetCellularControllerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetCellularControllerProfileParcelForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/cellular-controller | 
[**GetCellularProfileProfileParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetCellularProfileProfileParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/cellular-profile/{cellularProfileId} | 
[**GetCellularProfileProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetCellularProfileProfileParcelForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/cellular-profile | 
[**GetConfigProfileParcelByParcelIdForCLI**](ConfigurationFeatureProfileSDWANApi.md#GetConfigProfileParcelByParcelIdForCLI) | **Get** /v1/feature-profile/sdwan/cli/{cliId}/config/{configId} | 
[**GetConfigProfileParcelForCLI**](ConfigurationFeatureProfileSDWANApi.md#GetConfigProfileParcelForCLI) | **Get** /v1/feature-profile/sdwan/cli/{cliId}/config | 
[**GetGlobalProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetGlobalProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/global/{globalId} | 
[**GetGlobalProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetGlobalProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/global | 
[**GetInterfaceEthernetParcelsForServiceLanVpn**](ConfigurationFeatureProfileSDWANApi.md#GetInterfaceEthernetParcelsForServiceLanVpn) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet | 
[**GetInterfaceEthernetParcelsForTransportManagementVpn**](ConfigurationFeatureProfileSDWANApi.md#GetInterfaceEthernetParcelsForTransportManagementVpn) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId}/interface/ethernet | 
[**GetInterfaceEthernetParcelsForTransportWanVpn**](ConfigurationFeatureProfileSDWANApi.md#GetInterfaceEthernetParcelsForTransportWanVpn) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet | 
[**GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/bgp/{bgpId} | 
[**GetLanVpnAssociatedRoutingBgpParcelsForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnAssociatedRoutingBgpParcelsForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/bgp | 
[**GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/ospf/{ospfId} | 
[**GetLanVpnAssociatedRoutingOspfParcelsForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnAssociatedRoutingOspfParcelsForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/routing/ospf | 
[**GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker/{trackerId} | 
[**GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker | 
[**GetLanVpnInterfaceEthernetParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnInterfaceEthernetParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**GetLanVpnProfileParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnProfileParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn/{vpnId} | 
[**GetLanVpnProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#GetLanVpnProfileParcelForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/lan/vpn | 
[**GetLoggingProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetLoggingProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/logging/{loggingId} | 
[**GetLoggingProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetLoggingProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/logging | 
[**GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**GetManagementVpnProfileParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetManagementVpnProfileParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn/{vpnId} | 
[**GetManagementVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetManagementVpnProfileParcelForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/management/vpn | 
[**GetNtpProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetNtpProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/ntp/{ntpId} | 
[**GetNtpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetNtpProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/ntp | 
[**GetOmpProfileParcelByParcelIdForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetOmpProfileParcelByParcelIdForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/omp/{ompId} | 
[**GetOmpProfileParcelForSystem**](ConfigurationFeatureProfileSDWANApi.md#GetOmpProfileParcelForSystem) | **Get** /v1/feature-profile/sdwan/system/{systemId}/omp | 
[**GetRoutingBgpProfileParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetRoutingBgpProfileParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/routing/bgp/{bgpId} | 
[**GetRoutingBgpProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#GetRoutingBgpProfileParcelForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/routing/bgp | 
[**GetRoutingOspfProfileParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetRoutingOspfProfileParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/routing/ospf/{ospfId} | 
[**GetRoutingOspfProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#GetRoutingOspfProfileParcelForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/routing/ospf | 
[**GetSdwanFeatureProfileByProfileId**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanFeatureProfileByProfileId) | **Get** /v1/feature-profile/sdwan/cli/{cliId} | 
[**GetSdwanFeatureProfileBySdwanFamily**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanFeatureProfileBySdwanFamily) | **Get** /v1/feature-profile/sdwan | 
[**GetSdwanFeatureProfilesByFamilyAndType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanFeatureProfilesByFamilyAndType) | **Get** /v1/feature-profile/sdwan/cli | 
[**GetSdwanServiceFeatureProfileByProfileId**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceFeatureProfileByProfileId) | **Get** /v1/feature-profile/sdwan/service/{serviceId} | 
[**GetSdwanServiceFeatureProfiles**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceFeatureProfiles) | **Get** /v1/feature-profile/sdwan/service | 
[**GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema) | **Get** /v1/feature-profile/sdwan/service/lan/vpn/interface/ethernet/schema | 
[**GetSdwanServiceLanVpnParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceLanVpnParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/service/lan/vpn/schema | 
[**GetSdwanServiceRoutingBgpParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceRoutingBgpParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/service/routing/bgp/schema | 
[**GetSdwanServiceRoutingOspfParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceRoutingOspfParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/service/routing/ospf/schema | 
[**GetSdwanServiceTrackerParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanServiceTrackerParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/service/tracker/schema | 
[**GetSdwanSystemAaaParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemAaaParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/aaa/schema | 
[**GetSdwanSystemBannerParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemBannerParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/banner/schema | 
[**GetSdwanSystemBasicParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemBasicParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/basic/schema | 
[**GetSdwanSystemBfdParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemBfdParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/bfd/schema | 
[**GetSdwanSystemFeatureProfileByProfileId**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemFeatureProfileByProfileId) | **Get** /v1/feature-profile/sdwan/system/{systemId} | 
[**GetSdwanSystemFeatureProfiles**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemFeatureProfiles) | **Get** /v1/feature-profile/sdwan/system | 
[**GetSdwanSystemLoggingParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemLoggingParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/logging/schema | 
[**GetSdwanSystemNtpParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemNtpParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/ntp/schema | 
[**GetSdwanSystemOmpParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanSystemOmpParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/system/omp/schema | 
[**GetSdwanTransportCellularControllerParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportCellularControllerParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/transport/cellular-controller/schema | 
[**GetSdwanTransportCellularProfileParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportCellularProfileParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/transport/cellular-profile/schema | 
[**GetSdwanTransportFeatureProfileByProfileId**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportFeatureProfileByProfileId) | **Get** /v1/feature-profile/sdwan/transport/{transportId} | 
[**GetSdwanTransportFeatureProfiles**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportFeatureProfiles) | **Get** /v1/feature-profile/sdwan/transport | 
[**GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema) | **Get** /v1/feature-profile/sdwan/transport/management/vpn/interface/ethernet/schema | 
[**GetSdwanTransportManagementVpnParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportManagementVpnParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/transport/management/vpn/schema | 
[**GetSdwanTransportTrackerParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportTrackerParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/transport/tracker/schema | 
[**GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema) | **Get** /v1/feature-profile/sdwan/transport/wan/vpn/interface/ethernet/schema | 
[**GetSdwanTransportWanVpnParcelSchemaBySchemaType**](ConfigurationFeatureProfileSDWANApi.md#GetSdwanTransportWanVpnParcelSchemaBySchemaType) | **Get** /v1/feature-profile/sdwan/transport/wan/vpn/schema | 
[**GetTrackerProfileParcelByParcelIdForService**](ConfigurationFeatureProfileSDWANApi.md#GetTrackerProfileParcelByParcelIdForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/tracker/{trackerId} | 
[**GetTrackerProfileParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetTrackerProfileParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/tracker/{trackerId} | 
[**GetTrackerProfileParcelForService**](ConfigurationFeatureProfileSDWANApi.md#GetTrackerProfileParcelForService) | **Get** /v1/feature-profile/sdwan/service/{serviceId}/tracker | 
[**GetTrackerProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetTrackerProfileParcelForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/tracker | 
[**GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker/{trackerId} | 
[**GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId}/tracker | 
[**GetWanVpnInterfaceEthernetParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetWanVpnInterfaceEthernetParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId}/interface/ethernet/{ethernetId} | 
[**GetWanVpnProfileParcelByParcelIdForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetWanVpnProfileParcelByParcelIdForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn/{vpnId} | 
[**GetWanVpnProfileParcelForTransport**](ConfigurationFeatureProfileSDWANApi.md#GetWanVpnProfileParcelForTransport) | **Get** /v1/feature-profile/sdwan/transport/{transportId}/wan/vpn | 



## CreateAaaProfileParcelForSystem

> string CreateAaaProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Aaa Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateAaaProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateAaaProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAaaProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateAaaProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAaaProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Aaa Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBannerProfileParcelForSystem

> string CreateBannerProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Banner Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateBannerProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateBannerProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBannerProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateBannerProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBannerProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Banner Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBasicProfileParcelForSystem

> string CreateBasicProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Basic Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateBasicProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateBasicProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBasicProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateBasicProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBasicProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Basic Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBfdProfileParcelForSystem

> string CreateBfdProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Bfd Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateBfdProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateBfdProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBfdProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateBfdProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBfdProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Bfd Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCellularControllerAndCellularProfileParcelAssociationForTransport

> string CreateCellularControllerAndCellularProfileParcelAssociationForTransport(ctx, transportId, cellularControllerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Cellular Controller Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Cellular Profile Profile Parcel Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateCellularControllerAndCellularProfileParcelAssociationForTransport(context.Background(), transportId, cellularControllerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateCellularControllerAndCellularProfileParcelAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCellularControllerAndCellularProfileParcelAssociationForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateCellularControllerAndCellularProfileParcelAssociationForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Cellular Controller Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCellularControllerAndCellularProfileParcelAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Cellular Profile Profile Parcel Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCellularControllerProfileParcelForTransport

> string CreateCellularControllerProfileParcelForTransport(ctx, transportId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Cellular Controller Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateCellularControllerProfileParcelForTransport(context.Background(), transportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateCellularControllerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCellularControllerProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateCellularControllerProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCellularControllerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Cellular Controller Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCellularProfileProfileParcelForTransport

> string CreateCellularProfileProfileParcelForTransport(ctx, transportId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Cellular Profile Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateCellularProfileProfileParcelForTransport(context.Background(), transportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateCellularProfileProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCellularProfileProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateCellularProfileProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCellularProfileProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Cellular Profile Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalProfileParcelForSystem

> string CreateGlobalProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Global Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateGlobalProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateGlobalProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateGlobalProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Global Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanVpnAndRoutingBgpParcelAssociationForService

> string CreateLanVpnAndRoutingBgpParcelAssociationForService(ctx, serviceId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Lan Vpn Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Routing Bgp Profile Parcel Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateLanVpnAndRoutingBgpParcelAssociationForService(context.Background(), serviceId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateLanVpnAndRoutingBgpParcelAssociationForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanVpnAndRoutingBgpParcelAssociationForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateLanVpnAndRoutingBgpParcelAssociationForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Lan Vpn Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanVpnAndRoutingBgpParcelAssociationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Routing Bgp Profile Parcel Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanVpnAndRoutingOspfParcelAssociationForService

> string CreateLanVpnAndRoutingOspfParcelAssociationForService(ctx, serviceId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Lan Vpn Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Routing Ospf Profile Parcel Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateLanVpnAndRoutingOspfParcelAssociationForService(context.Background(), serviceId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateLanVpnAndRoutingOspfParcelAssociationForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanVpnAndRoutingOspfParcelAssociationForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateLanVpnAndRoutingOspfParcelAssociationForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Lan Vpn Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanVpnAndRoutingOspfParcelAssociationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Routing Ospf Profile Parcel Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport

> string CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(ctx, serviceId, vpnParcelId, ethernetId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnParcelId := "vpnParcelId_example" // string | VPN Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(context.Background(), serviceId, vpnParcelId, ethernetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnParcelId** | **string** | VPN Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Tracker Profile Parcel Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanVpnInterfaceEthernetParcelForService

> string CreateLanVpnInterfaceEthernetParcelForService(ctx, serviceId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Lan Vpn Interface Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateLanVpnInterfaceEthernetParcelForService(context.Background(), serviceId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateLanVpnInterfaceEthernetParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanVpnInterfaceEthernetParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateLanVpnInterfaceEthernetParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanVpnInterfaceEthernetParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Lan Vpn Interface Ethernet Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanVpnProfileParcelForService

> string CreateLanVpnProfileParcelForService(ctx, serviceId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Lan Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateLanVpnProfileParcelForService(context.Background(), serviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateLanVpnProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanVpnProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateLanVpnProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanVpnProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Lan Vpn Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoggingProfileParcelForSystem

> string CreateLoggingProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Logging Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateLoggingProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateLoggingProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoggingProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateLoggingProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoggingProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Logging Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManagementVpnInterfaceEthernetParcelForTransport

> string CreateManagementVpnInterfaceEthernetParcelForTransport(ctx, transportId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Management Vpn Interface Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateManagementVpnInterfaceEthernetParcelForTransport(context.Background(), transportId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateManagementVpnInterfaceEthernetParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagementVpnInterfaceEthernetParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateManagementVpnInterfaceEthernetParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagementVpnInterfaceEthernetParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Management Vpn Interface Ethernet Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManagementVpnProfileParcelForTransport

> string CreateManagementVpnProfileParcelForTransport(ctx, transportId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Management Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateManagementVpnProfileParcelForTransport(context.Background(), transportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateManagementVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagementVpnProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateManagementVpnProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagementVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Management Vpn Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNtpProfileParcelForSystem

> string CreateNtpProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Ntp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateNtpProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateNtpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNtpProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateNtpProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNtpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Ntp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOmpProfileParcelForSystem

> string CreateOmpProfileParcelForSystem(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Omp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateOmpProfileParcelForSystem(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateOmpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOmpProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateOmpProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOmpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Omp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoutingBgpProfileParcelForService

> string CreateRoutingBgpProfileParcelForService(ctx, serviceId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Routing Bgp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateRoutingBgpProfileParcelForService(context.Background(), serviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateRoutingBgpProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoutingBgpProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateRoutingBgpProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutingBgpProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Routing Bgp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoutingOspfProfileParcelForService

> string CreateRoutingOspfProfileParcelForService(ctx, serviceId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Routing Ospf Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateRoutingOspfProfileParcelForService(context.Background(), serviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateRoutingOspfProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoutingOspfProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateRoutingOspfProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutingOspfProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Routing Ospf Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdwanConfigProfileParcelForCli

> string CreateSdwanConfigProfileParcelForCli(ctx, cliId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | cli config Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateSdwanConfigProfileParcelForCli(context.Background(), cliId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateSdwanConfigProfileParcelForCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdwanConfigProfileParcelForCli`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateSdwanConfigProfileParcelForCli`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdwanConfigProfileParcelForCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | cli config Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdwanFeatureProfile

> string CreateSdwanFeatureProfile(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "{"name":"Cli profile name","description":"my Cli profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateSdwanFeatureProfile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateSdwanFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdwanFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateSdwanFeatureProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdwanFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdwanServiceFeatureProfile

> string CreateSdwanServiceFeatureProfile(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "{"name":"service profile name","description":"my service profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateSdwanServiceFeatureProfile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateSdwanServiceFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdwanServiceFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateSdwanServiceFeatureProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdwanServiceFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdwanSystemFeatureProfile

> string CreateSdwanSystemFeatureProfile(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "{"name":"system profile name","description":"my system profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateSdwanSystemFeatureProfile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateSdwanSystemFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdwanSystemFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateSdwanSystemFeatureProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdwanSystemFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdwanTransportFeatureProfile

> string CreateSdwanTransportFeatureProfile(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "{"name":"transport profile name","description":"my transport profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateSdwanTransportFeatureProfile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateSdwanTransportFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdwanTransportFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateSdwanTransportFeatureProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdwanTransportFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackerProfileParcelForService

> string CreateTrackerProfileParcelForService(ctx, serviceId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateTrackerProfileParcelForService(context.Background(), serviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateTrackerProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrackerProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateTrackerProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackerProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Tracker Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackerProfileParcelForTransport

> string CreateTrackerProfileParcelForTransport(ctx, transportId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateTrackerProfileParcelForTransport(context.Background(), transportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateTrackerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrackerProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateTrackerProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Tracker Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport

> string CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(ctx, transportId, vpnParcelId, ethernetId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnParcelId := "vpnParcelId_example" // string | VPN Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(context.Background(), transportId, vpnParcelId, ethernetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnParcelId** | **string** | VPN Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Tracker Profile Parcel Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWanVpnInterfaceEthernetParcelForTransport

> string CreateWanVpnInterfaceEthernetParcelForTransport(ctx, transportId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Wan Vpn Interface Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateWanVpnInterfaceEthernetParcelForTransport(context.Background(), transportId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateWanVpnInterfaceEthernetParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWanVpnInterfaceEthernetParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateWanVpnInterfaceEthernetParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWanVpnInterfaceEthernetParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Wan Vpn Interface Ethernet Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWanVpnProfileParcelForTransport

> string CreateWanVpnProfileParcelForTransport(ctx, transportId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | Wan Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.CreateWanVpnProfileParcelForTransport(context.Background(), transportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.CreateWanVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWanVpnProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.CreateWanVpnProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWanVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Wan Vpn Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAaaProfileParcelForSystem

> DeleteAaaProfileParcelForSystem(ctx, systemId, aaaId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    aaaId := "aaaId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteAaaProfileParcelForSystem(context.Background(), systemId, aaaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteAaaProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**aaaId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAaaProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBannerProfileParcelForSystem

> DeleteBannerProfileParcelForSystem(ctx, systemId, bannerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    bannerId := "bannerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteBannerProfileParcelForSystem(context.Background(), systemId, bannerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteBannerProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**bannerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBannerProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBasicProfileParcelForSystem

> DeleteBasicProfileParcelForSystem(ctx, systemId, basicId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    basicId := "basicId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteBasicProfileParcelForSystem(context.Background(), systemId, basicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteBasicProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**basicId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBasicProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBfdProfileParcelForSystem

> DeleteBfdProfileParcelForSystem(ctx, systemId, bfdId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    bfdId := "bfdId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteBfdProfileParcelForSystem(context.Background(), systemId, bfdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteBfdProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**bfdId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBfdProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCellularControllerAndCellularProfileAssociationForTransport

> DeleteCellularControllerAndCellularProfileAssociationForTransport(ctx, transportId, cellularControllerId, cellularProfileId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Profile Parcel ID
    cellularProfileId := "cellularProfileId_example" // string | Cellular Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteCellularControllerAndCellularProfileAssociationForTransport(context.Background(), transportId, cellularControllerId, cellularProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteCellularControllerAndCellularProfileAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Profile Parcel ID | 
**cellularProfileId** | **string** | Cellular Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCellularControllerAndCellularProfileAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCellularControllerProfileParcelForTransport

> DeleteCellularControllerProfileParcelForTransport(ctx, transportId, cellularControllerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteCellularControllerProfileParcelForTransport(context.Background(), transportId, cellularControllerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteCellularControllerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCellularControllerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCellularProfileProfileParcelForTransport

> DeleteCellularProfileProfileParcelForTransport(ctx, transportId, cellularProfileId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularProfileId := "cellularProfileId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteCellularProfileProfileParcelForTransport(context.Background(), transportId, cellularProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteCellularProfileProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularProfileId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCellularProfileProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigProfileParcelForCLI

> DeleteConfigProfileParcelForCLI(ctx, cliId, configId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile ID
    configId := "configId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteConfigProfileParcelForCLI(context.Background(), cliId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteConfigProfileParcelForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 
**configId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigProfileParcelForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalProfileParcelForSystem

> DeleteGlobalProfileParcelForSystem(ctx, systemId, globalId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    globalId := "globalId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteGlobalProfileParcelForSystem(context.Background(), systemId, globalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteGlobalProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**globalId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanVpnAndRoutingBgpAssociationForService

> DeleteLanVpnAndRoutingBgpAssociationForService(ctx, serviceId, vpnId, bgpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    bgpId := "bgpId_example" // string | Routing Bgp Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteLanVpnAndRoutingBgpAssociationForService(context.Background(), serviceId, vpnId, bgpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteLanVpnAndRoutingBgpAssociationForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**bgpId** | **string** | Routing Bgp Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanVpnAndRoutingBgpAssociationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanVpnAndRoutingOspfAssociationForService

> DeleteLanVpnAndRoutingOspfAssociationForService(ctx, serviceId, vpnId, ospfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ospfId := "ospfId_example" // string | Routing Ospf Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteLanVpnAndRoutingOspfAssociationForService(context.Background(), serviceId, vpnId, ospfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteLanVpnAndRoutingOspfAssociationForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ospfId** | **string** | Routing Ospf Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanVpnAndRoutingOspfAssociationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransport

> DeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransport(ctx, serviceId, vpnId, ethernetId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    trackerId := "trackerId_example" // string | Tracker Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransport(context.Background(), serviceId, vpnId, ethernetId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 
**trackerId** | **string** | Tracker Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanVpnInterfaceEthernetAndTrackerAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanVpnInterfaceEthernetForService

> DeleteLanVpnInterfaceEthernetForService(ctx, serviceId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteLanVpnInterfaceEthernetForService(context.Background(), serviceId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteLanVpnInterfaceEthernetForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanVpnInterfaceEthernetForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanVpnProfileParcelForService

> DeleteLanVpnProfileParcelForService(ctx, serviceId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteLanVpnProfileParcelForService(context.Background(), serviceId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteLanVpnProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanVpnProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoggingProfileParcelForSystem

> DeleteLoggingProfileParcelForSystem(ctx, systemId, loggingId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    loggingId := "loggingId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteLoggingProfileParcelForSystem(context.Background(), systemId, loggingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteLoggingProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**loggingId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoggingProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagementVpnInterfaceEthernetForTransport

> DeleteManagementVpnInterfaceEthernetForTransport(ctx, transportId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteManagementVpnInterfaceEthernetForTransport(context.Background(), transportId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteManagementVpnInterfaceEthernetForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementVpnInterfaceEthernetForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagementVpnProfileParcelForTransport

> DeleteManagementVpnProfileParcelForTransport(ctx, transportId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteManagementVpnProfileParcelForTransport(context.Background(), transportId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteManagementVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNtpProfileParcelForSystem

> DeleteNtpProfileParcelForSystem(ctx, systemId, ntpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    ntpId := "ntpId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteNtpProfileParcelForSystem(context.Background(), systemId, ntpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteNtpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**ntpId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNtpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOmpProfileParcelForSystem

> DeleteOmpProfileParcelForSystem(ctx, systemId, ompId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    ompId := "ompId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteOmpProfileParcelForSystem(context.Background(), systemId, ompId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteOmpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**ompId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOmpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutingBgpProfileParcelForService

> DeleteRoutingBgpProfileParcelForService(ctx, serviceId, bgpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    bgpId := "bgpId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteRoutingBgpProfileParcelForService(context.Background(), serviceId, bgpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteRoutingBgpProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**bgpId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutingBgpProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutingOspfProfileParcelForService

> DeleteRoutingOspfProfileParcelForService(ctx, serviceId, ospfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    ospfId := "ospfId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteRoutingOspfProfileParcelForService(context.Background(), serviceId, ospfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteRoutingOspfProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**ospfId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutingOspfProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdwanFeatureProfileForCli

> DeleteSdwanFeatureProfileForCli(ctx, cliId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteSdwanFeatureProfileForCli(context.Background(), cliId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteSdwanFeatureProfileForCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdwanFeatureProfileForCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdwanServiceFeatureProfile

> DeleteSdwanServiceFeatureProfile(ctx, serviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteSdwanServiceFeatureProfile(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteSdwanServiceFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdwanServiceFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdwanSystemFeatureProfile

> DeleteSdwanSystemFeatureProfile(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteSdwanSystemFeatureProfile(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteSdwanSystemFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdwanSystemFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdwanTransportFeatureProfile

> DeleteSdwanTransportFeatureProfile(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteSdwanTransportFeatureProfile(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteSdwanTransportFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdwanTransportFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrackerProfileParcelForService

> DeleteTrackerProfileParcelForService(ctx, serviceId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    trackerId := "trackerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteTrackerProfileParcelForService(context.Background(), serviceId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteTrackerProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**trackerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackerProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrackerProfileParcelForTransport

> DeleteTrackerProfileParcelForTransport(ctx, transportId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    trackerId := "trackerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteTrackerProfileParcelForTransport(context.Background(), transportId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteTrackerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**trackerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransport

> DeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransport(ctx, transportId, vpnId, ethernetId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    trackerId := "trackerId_example" // string | Tracker Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransport(context.Background(), transportId, vpnId, ethernetId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 
**trackerId** | **string** | Tracker Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWanVpnInterfaceEthernetAndTrackerAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWanVpnInterfaceEthernetForTransport

> DeleteWanVpnInterfaceEthernetForTransport(ctx, transportId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteWanVpnInterfaceEthernetForTransport(context.Background(), transportId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteWanVpnInterfaceEthernetForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWanVpnInterfaceEthernetForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWanVpnProfileParcelForTransport

> DeleteWanVpnProfileParcelForTransport(ctx, transportId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.DeleteWanVpnProfileParcelForTransport(context.Background(), transportId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.DeleteWanVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWanVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAaaProfileParcelForSystem

> string EditAaaProfileParcelForSystem(ctx, systemId, aaaId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    aaaId := "aaaId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Aaa Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditAaaProfileParcelForSystem(context.Background(), systemId, aaaId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditAaaProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAaaProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditAaaProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**aaaId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAaaProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Aaa Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBannerProfileParcelForSystem

> string EditBannerProfileParcelForSystem(ctx, systemId, bannerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    bannerId := "bannerId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Banner Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditBannerProfileParcelForSystem(context.Background(), systemId, bannerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditBannerProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBannerProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditBannerProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**bannerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBannerProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Banner Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBasicProfileParcelForSystem

> string EditBasicProfileParcelForSystem(ctx, systemId, basicId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    basicId := "basicId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Basic Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditBasicProfileParcelForSystem(context.Background(), systemId, basicId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditBasicProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBasicProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditBasicProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**basicId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBasicProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Basic Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBfdProfileParcelForSystem

> string EditBfdProfileParcelForSystem(ctx, systemId, bfdId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    bfdId := "bfdId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Bfd Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditBfdProfileParcelForSystem(context.Background(), systemId, bfdId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditBfdProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBfdProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditBfdProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**bfdId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBfdProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Bfd Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCellularControllerAndCellularProfileParcelAssociationForTransport

> string EditCellularControllerAndCellularProfileParcelAssociationForTransport(ctx, transportId, cellularControllerId, cellularProfileId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Profile Parcel ID
    cellularProfileId := "cellularProfileId_example" // string | Cellular Profile ID
    body := "refer to schema for documentation and example" // string | Cellular Profile Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditCellularControllerAndCellularProfileParcelAssociationForTransport(context.Background(), transportId, cellularControllerId, cellularProfileId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditCellularControllerAndCellularProfileParcelAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCellularControllerAndCellularProfileParcelAssociationForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditCellularControllerAndCellularProfileParcelAssociationForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Profile Parcel ID | 
**cellularProfileId** | **string** | Cellular Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCellularControllerAndCellularProfileParcelAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Cellular Profile Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCellularControllerProfileParcelForTransport

> string EditCellularControllerProfileParcelForTransport(ctx, transportId, cellularControllerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Cellular Controller Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditCellularControllerProfileParcelForTransport(context.Background(), transportId, cellularControllerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditCellularControllerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCellularControllerProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditCellularControllerProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCellularControllerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Cellular Controller Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCellularProfileProfileParcelForTransport

> string EditCellularProfileProfileParcelForTransport(ctx, transportId, cellularProfileId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularProfileId := "cellularProfileId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Cellular Profile Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditCellularProfileProfileParcelForTransport(context.Background(), transportId, cellularProfileId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditCellularProfileProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCellularProfileProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditCellularProfileProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularProfileId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCellularProfileProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Cellular Profile Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditConfigProfileParcelForCLI

> string EditConfigProfileParcelForCLI(ctx, cliId, configId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile ID
    configId := "configId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | cli config Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditConfigProfileParcelForCLI(context.Background(), cliId, configId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditConfigProfileParcelForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditConfigProfileParcelForCLI`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditConfigProfileParcelForCLI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 
**configId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditConfigProfileParcelForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | cli config Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditGlobalProfileParcelForSystem

> string EditGlobalProfileParcelForSystem(ctx, systemId, globalId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    globalId := "globalId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Global Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditGlobalProfileParcelForSystem(context.Background(), systemId, globalId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditGlobalProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditGlobalProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditGlobalProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**globalId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditGlobalProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Global Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLanVpnAndRoutingBgpParcelAssociationForService

> string EditLanVpnAndRoutingBgpParcelAssociationForService(ctx, serviceId, vpnId, bgpId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    bgpId := "bgpId_example" // string | Routing Bgp ID
    body := "refer to schema for documentation and example" // string | Routing Bgp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditLanVpnAndRoutingBgpParcelAssociationForService(context.Background(), serviceId, vpnId, bgpId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditLanVpnAndRoutingBgpParcelAssociationForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLanVpnAndRoutingBgpParcelAssociationForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditLanVpnAndRoutingBgpParcelAssociationForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**bgpId** | **string** | Routing Bgp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLanVpnAndRoutingBgpParcelAssociationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Routing Bgp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLanVpnAndRoutingOspfParcelAssociationForService

> string EditLanVpnAndRoutingOspfParcelAssociationForService(ctx, serviceId, vpnId, ospfId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ospfId := "ospfId_example" // string | Routing Ospf ID
    body := "refer to schema for documentation and example" // string | Routing Ospf Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditLanVpnAndRoutingOspfParcelAssociationForService(context.Background(), serviceId, vpnId, ospfId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditLanVpnAndRoutingOspfParcelAssociationForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLanVpnAndRoutingOspfParcelAssociationForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditLanVpnAndRoutingOspfParcelAssociationForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ospfId** | **string** | Routing Ospf ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLanVpnAndRoutingOspfParcelAssociationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Routing Ospf Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport

> string EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(ctx, serviceId, vpnId, ethernetId, trackerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    trackerId := "trackerId_example" // string | Tracker ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(context.Background(), serviceId, vpnId, ethernetId, trackerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 
**trackerId** | **string** | Tracker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLanVpnInterfaceEthernetAndTrackerParcelAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** | Tracker Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLanVpnInterfaceEthernetParcelForService

> string EditLanVpnInterfaceEthernetParcelForService(ctx, serviceId, vpnId, ethernetId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface ID
    body := "refer to schema for documentation and example" // string | Lan Vpn Interface Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditLanVpnInterfaceEthernetParcelForService(context.Background(), serviceId, vpnId, ethernetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditLanVpnInterfaceEthernetParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLanVpnInterfaceEthernetParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditLanVpnInterfaceEthernetParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLanVpnInterfaceEthernetParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Lan Vpn Interface Ethernet Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLanVpnProfileParcelForService

> string EditLanVpnProfileParcelForService(ctx, serviceId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Lan Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditLanVpnProfileParcelForService(context.Background(), serviceId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditLanVpnProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLanVpnProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditLanVpnProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLanVpnProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Lan Vpn Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLoggingProfileParcelForSystem

> string EditLoggingProfileParcelForSystem(ctx, systemId, loggingId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    loggingId := "loggingId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Logging Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditLoggingProfileParcelForSystem(context.Background(), systemId, loggingId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditLoggingProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLoggingProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditLoggingProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**loggingId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLoggingProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Logging Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditManagementVpnInterfaceEthernetParcelForTransport

> string EditManagementVpnInterfaceEthernetParcelForTransport(ctx, transportId, vpnId, ethernetId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface ID
    body := "refer to schema for documentation and example" // string | Management Vpn Interface Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditManagementVpnInterfaceEthernetParcelForTransport(context.Background(), transportId, vpnId, ethernetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditManagementVpnInterfaceEthernetParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditManagementVpnInterfaceEthernetParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditManagementVpnInterfaceEthernetParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditManagementVpnInterfaceEthernetParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Management Vpn Interface Ethernet Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditManagementVpnProfileParcelForTransport

> string EditManagementVpnProfileParcelForTransport(ctx, transportId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Management Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditManagementVpnProfileParcelForTransport(context.Background(), transportId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditManagementVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditManagementVpnProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditManagementVpnProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditManagementVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Management Vpn Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditNtpProfileParcelForSystem

> string EditNtpProfileParcelForSystem(ctx, systemId, ntpId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    ntpId := "ntpId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Ntp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditNtpProfileParcelForSystem(context.Background(), systemId, ntpId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditNtpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditNtpProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditNtpProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**ntpId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditNtpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Ntp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOmpProfileParcelForSystem

> string EditOmpProfileParcelForSystem(ctx, systemId, ompId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    ompId := "ompId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Omp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditOmpProfileParcelForSystem(context.Background(), systemId, ompId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditOmpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOmpProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditOmpProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**ompId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOmpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Omp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRoutingBgpProfileParcelForService

> string EditRoutingBgpProfileParcelForService(ctx, serviceId, bgpId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    bgpId := "bgpId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Routing Bgp Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditRoutingBgpProfileParcelForService(context.Background(), serviceId, bgpId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditRoutingBgpProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditRoutingBgpProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditRoutingBgpProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**bgpId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRoutingBgpProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Routing Bgp Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRoutingOspfProfileParcelForService

> string EditRoutingOspfProfileParcelForService(ctx, serviceId, ospfId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    ospfId := "ospfId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Routing Ospf Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditRoutingOspfProfileParcelForService(context.Background(), serviceId, ospfId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditRoutingOspfProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditRoutingOspfProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditRoutingOspfProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**ospfId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRoutingOspfProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Routing Ospf Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSdwanFeatureProfile

> string EditSdwanFeatureProfile(ctx, cliId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile Id
    body := "{"name":"my Cli profile","description":"my Cli profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditSdwanFeatureProfile(context.Background(), cliId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditSdwanFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSdwanFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditSdwanFeatureProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSdwanFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSdwanServiceFeatureProfile

> string EditSdwanServiceFeatureProfile(ctx, serviceId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile Id
    body := "{"name":"my service profile","description":"my service profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditSdwanServiceFeatureProfile(context.Background(), serviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditSdwanServiceFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSdwanServiceFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditSdwanServiceFeatureProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSdwanServiceFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSdwanSystemFeatureProfile

> string EditSdwanSystemFeatureProfile(ctx, systemId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile Id
    body := "{"name":"my system profile","description":"my system profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditSdwanSystemFeatureProfile(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditSdwanSystemFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSdwanSystemFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditSdwanSystemFeatureProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSdwanSystemFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSdwanTransportFeatureProfile

> string EditSdwanTransportFeatureProfile(ctx, transportId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile Id
    body := "{"name":"my transport profile","description":"my transport profile"}" // string | SDWAN Feature profile (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditSdwanTransportFeatureProfile(context.Background(), transportId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditSdwanTransportFeatureProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSdwanTransportFeatureProfile`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditSdwanTransportFeatureProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSdwanTransportFeatureProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | SDWAN Feature profile | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTrackerProfileParcelForService

> string EditTrackerProfileParcelForService(ctx, serviceId, trackerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    trackerId := "trackerId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditTrackerProfileParcelForService(context.Background(), serviceId, trackerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditTrackerProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTrackerProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditTrackerProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**trackerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTrackerProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Tracker Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTrackerProfileParcelForTransport

> string EditTrackerProfileParcelForTransport(ctx, transportId, trackerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    trackerId := "trackerId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditTrackerProfileParcelForTransport(context.Background(), transportId, trackerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditTrackerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTrackerProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditTrackerProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**trackerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTrackerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Tracker Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport

> string EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(ctx, transportId, vpnId, ethernetId, trackerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    trackerId := "trackerId_example" // string | Tracker ID
    body := "refer to schema for documentation and example" // string | Tracker Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport(context.Background(), transportId, vpnId, ethernetId, trackerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 
**trackerId** | **string** | Tracker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWanVpnInterfaceEthernetAndTrackerParcelAssociationForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** | Tracker Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWanVpnInterfaceEthernetParcelForTransport

> string EditWanVpnInterfaceEthernetParcelForTransport(ctx, transportId, vpnId, ethernetId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface ID
    body := "refer to schema for documentation and example" // string | Wan Vpn Interface Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditWanVpnInterfaceEthernetParcelForTransport(context.Background(), transportId, vpnId, ethernetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditWanVpnInterfaceEthernetParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWanVpnInterfaceEthernetParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditWanVpnInterfaceEthernetParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWanVpnInterfaceEthernetParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | Wan Vpn Interface Ethernet Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWanVpnProfileParcelForTransport

> string EditWanVpnProfileParcelForTransport(ctx, transportId, vpnId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | Wan Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.EditWanVpnProfileParcelForTransport(context.Background(), transportId, vpnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.EditWanVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWanVpnProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.EditWanVpnProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWanVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Wan Vpn Profile Parcel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAaaProfileParcelByParcelIdForSystem

> string GetAaaProfileParcelByParcelIdForSystem(ctx, systemId, aaaId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    aaaId := "aaaId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetAaaProfileParcelByParcelIdForSystem(context.Background(), systemId, aaaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetAaaProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAaaProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetAaaProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**aaaId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAaaProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAaaProfileParcelForSystem

> string GetAaaProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetAaaProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetAaaProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAaaProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetAaaProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAaaProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBannerProfileParcelByParcelIdForSystem

> string GetBannerProfileParcelByParcelIdForSystem(ctx, systemId, bannerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    bannerId := "bannerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetBannerProfileParcelByParcelIdForSystem(context.Background(), systemId, bannerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetBannerProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBannerProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetBannerProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**bannerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannerProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBannerProfileParcelForSystem

> string GetBannerProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetBannerProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetBannerProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBannerProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetBannerProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannerProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBasicProfileParcelByParcelIdForSystem

> string GetBasicProfileParcelByParcelIdForSystem(ctx, systemId, basicId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    basicId := "basicId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetBasicProfileParcelByParcelIdForSystem(context.Background(), systemId, basicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetBasicProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetBasicProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**basicId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBasicProfileParcelForSystem

> string GetBasicProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetBasicProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetBasicProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetBasicProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBfdProfileParcelByParcelIdForSystem

> string GetBfdProfileParcelByParcelIdForSystem(ctx, systemId, bfdId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    bfdId := "bfdId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetBfdProfileParcelByParcelIdForSystem(context.Background(), systemId, bfdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetBfdProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBfdProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetBfdProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**bfdId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBfdProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBfdProfileParcelForSystem

> string GetBfdProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetBfdProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetBfdProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBfdProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetBfdProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBfdProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCedgeSystemGlobalParcelSchemaBySchemaType

> string GetCedgeSystemGlobalParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCedgeSystemGlobalParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCedgeSystemGlobalParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCedgeSystemGlobalParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCedgeSystemGlobalParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCedgeSystemGlobalParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport

> string GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport(ctx, transportId, cellularControllerId, cellularProfileId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Profile Parcel ID
    cellularProfileId := "cellularProfileId_example" // string | Cellular Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport(context.Background(), transportId, cellularControllerId, cellularProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Profile Parcel ID | 
**cellularProfileId** | **string** | Cellular Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularControllerAssociatedCellularProfileParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellularControllerAssociatedCellularProfileParcelsForTransport

> string GetCellularControllerAssociatedCellularProfileParcelsForTransport(ctx, transportId, cellularControllerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Feature Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCellularControllerAssociatedCellularProfileParcelsForTransport(context.Background(), transportId, cellularControllerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCellularControllerAssociatedCellularProfileParcelsForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularControllerAssociatedCellularProfileParcelsForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCellularControllerAssociatedCellularProfileParcelsForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Feature Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularControllerAssociatedCellularProfileParcelsForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellularControllerProfileParcelByParcelIdForTransport

> string GetCellularControllerProfileParcelByParcelIdForTransport(ctx, transportId, cellularControllerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularControllerId := "cellularControllerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCellularControllerProfileParcelByParcelIdForTransport(context.Background(), transportId, cellularControllerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCellularControllerProfileParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularControllerProfileParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCellularControllerProfileParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularControllerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularControllerProfileParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellularControllerProfileParcelForTransport

> string GetCellularControllerProfileParcelForTransport(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCellularControllerProfileParcelForTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCellularControllerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularControllerProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCellularControllerProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularControllerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellularProfileProfileParcelByParcelIdForTransport

> string GetCellularProfileProfileParcelByParcelIdForTransport(ctx, transportId, cellularProfileId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    cellularProfileId := "cellularProfileId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCellularProfileProfileParcelByParcelIdForTransport(context.Background(), transportId, cellularProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCellularProfileProfileParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularProfileProfileParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCellularProfileProfileParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**cellularProfileId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularProfileProfileParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellularProfileProfileParcelForTransport

> string GetCellularProfileProfileParcelForTransport(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetCellularProfileProfileParcelForTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetCellularProfileProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularProfileProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetCellularProfileProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularProfileProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigProfileParcelByParcelIdForCLI

> string GetConfigProfileParcelByParcelIdForCLI(ctx, cliId, configId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile ID
    configId := "configId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetConfigProfileParcelByParcelIdForCLI(context.Background(), cliId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetConfigProfileParcelByParcelIdForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigProfileParcelByParcelIdForCLI`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetConfigProfileParcelByParcelIdForCLI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 
**configId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigProfileParcelByParcelIdForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigProfileParcelForCLI

> string GetConfigProfileParcelForCLI(ctx, cliId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetConfigProfileParcelForCLI(context.Background(), cliId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetConfigProfileParcelForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigProfileParcelForCLI`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetConfigProfileParcelForCLI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigProfileParcelForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalProfileParcelByParcelIdForSystem

> string GetGlobalProfileParcelByParcelIdForSystem(ctx, systemId, globalId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    globalId := "globalId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetGlobalProfileParcelByParcelIdForSystem(context.Background(), systemId, globalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetGlobalProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetGlobalProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**globalId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalProfileParcelForSystem

> string GetGlobalProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetGlobalProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetGlobalProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetGlobalProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceEthernetParcelsForServiceLanVpn

> string GetInterfaceEthernetParcelsForServiceLanVpn(ctx, serviceId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForServiceLanVpn(context.Background(), serviceId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForServiceLanVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceEthernetParcelsForServiceLanVpn`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForServiceLanVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceEthernetParcelsForServiceLanVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceEthernetParcelsForTransportManagementVpn

> string GetInterfaceEthernetParcelsForTransportManagementVpn(ctx, transportId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForTransportManagementVpn(context.Background(), transportId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForTransportManagementVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceEthernetParcelsForTransportManagementVpn`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForTransportManagementVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceEthernetParcelsForTransportManagementVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceEthernetParcelsForTransportWanVpn

> string GetInterfaceEthernetParcelsForTransportWanVpn(ctx, transportId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForTransportWanVpn(context.Background(), transportId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForTransportWanVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceEthernetParcelsForTransportWanVpn`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetInterfaceEthernetParcelsForTransportWanVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceEthernetParcelsForTransportWanVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService

> string GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService(ctx, serviceId, vpnId, bgpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    bgpId := "bgpId_example" // string | Routing Bgp Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService(context.Background(), serviceId, vpnId, bgpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingBgpParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**bgpId** | **string** | Routing Bgp Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnAssociatedRoutingBgpParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnAssociatedRoutingBgpParcelsForService

> string GetLanVpnAssociatedRoutingBgpParcelsForService(ctx, serviceId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingBgpParcelsForService(context.Background(), serviceId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingBgpParcelsForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnAssociatedRoutingBgpParcelsForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingBgpParcelsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnAssociatedRoutingBgpParcelsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService

> string GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService(ctx, serviceId, vpnId, ospfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ospfId := "ospfId_example" // string | Routing Ospf Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService(context.Background(), serviceId, vpnId, ospfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingOspfParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ospfId** | **string** | Routing Ospf Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnAssociatedRoutingOspfParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnAssociatedRoutingOspfParcelsForService

> string GetLanVpnAssociatedRoutingOspfParcelsForService(ctx, serviceId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingOspfParcelsForService(context.Background(), serviceId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingOspfParcelsForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnAssociatedRoutingOspfParcelsForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnAssociatedRoutingOspfParcelsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnAssociatedRoutingOspfParcelsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport

> string GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport(ctx, serviceId, vpnId, ethernetId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    trackerId := "trackerId_example" // string | Tracker Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport(context.Background(), serviceId, vpnId, ethernetId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 
**trackerId** | **string** | Tracker Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport

> string GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport(ctx, serviceId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport(context.Background(), serviceId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnInterfaceEthernetAssociatedTrackerParcelsForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnInterfaceEthernetParcelByParcelIdForService

> string GetLanVpnInterfaceEthernetParcelByParcelIdForService(ctx, serviceId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetParcelByParcelIdForService(context.Background(), serviceId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnInterfaceEthernetParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnInterfaceEthernetParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnInterfaceEthernetParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnProfileParcelByParcelIdForService

> string GetLanVpnProfileParcelByParcelIdForService(ctx, serviceId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnProfileParcelByParcelIdForService(context.Background(), serviceId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnProfileParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnProfileParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnProfileParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnProfileParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanVpnProfileParcelForService

> string GetLanVpnProfileParcelForService(ctx, serviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLanVpnProfileParcelForService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLanVpnProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanVpnProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLanVpnProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanVpnProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggingProfileParcelByParcelIdForSystem

> string GetLoggingProfileParcelByParcelIdForSystem(ctx, systemId, loggingId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    loggingId := "loggingId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLoggingProfileParcelByParcelIdForSystem(context.Background(), systemId, loggingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLoggingProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoggingProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLoggingProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**loggingId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggingProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggingProfileParcelForSystem

> string GetLoggingProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetLoggingProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetLoggingProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoggingProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetLoggingProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggingProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport

> string GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport(ctx, transportId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport(context.Background(), transportId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetManagementVpnInterfaceEthernetParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementVpnInterfaceEthernetParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementVpnProfileParcelByParcelIdForTransport

> string GetManagementVpnProfileParcelByParcelIdForTransport(ctx, transportId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetManagementVpnProfileParcelByParcelIdForTransport(context.Background(), transportId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetManagementVpnProfileParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementVpnProfileParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetManagementVpnProfileParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementVpnProfileParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementVpnProfileParcelForTransport

> string GetManagementVpnProfileParcelForTransport(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetManagementVpnProfileParcelForTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetManagementVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementVpnProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetManagementVpnProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNtpProfileParcelByParcelIdForSystem

> string GetNtpProfileParcelByParcelIdForSystem(ctx, systemId, ntpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    ntpId := "ntpId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetNtpProfileParcelByParcelIdForSystem(context.Background(), systemId, ntpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetNtpProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNtpProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetNtpProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**ntpId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNtpProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNtpProfileParcelForSystem

> string GetNtpProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetNtpProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetNtpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNtpProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetNtpProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNtpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOmpProfileParcelByParcelIdForSystem

> string GetOmpProfileParcelByParcelIdForSystem(ctx, systemId, ompId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID
    ompId := "ompId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetOmpProfileParcelByParcelIdForSystem(context.Background(), systemId, ompId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetOmpProfileParcelByParcelIdForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOmpProfileParcelByParcelIdForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetOmpProfileParcelByParcelIdForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 
**ompId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOmpProfileParcelByParcelIdForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOmpProfileParcelForSystem

> string GetOmpProfileParcelForSystem(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetOmpProfileParcelForSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetOmpProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOmpProfileParcelForSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetOmpProfileParcelForSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOmpProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingBgpProfileParcelByParcelIdForService

> string GetRoutingBgpProfileParcelByParcelIdForService(ctx, serviceId, bgpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    bgpId := "bgpId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetRoutingBgpProfileParcelByParcelIdForService(context.Background(), serviceId, bgpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetRoutingBgpProfileParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutingBgpProfileParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetRoutingBgpProfileParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**bgpId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingBgpProfileParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingBgpProfileParcelForService

> string GetRoutingBgpProfileParcelForService(ctx, serviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetRoutingBgpProfileParcelForService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetRoutingBgpProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutingBgpProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetRoutingBgpProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingBgpProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingOspfProfileParcelByParcelIdForService

> string GetRoutingOspfProfileParcelByParcelIdForService(ctx, serviceId, ospfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    ospfId := "ospfId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetRoutingOspfProfileParcelByParcelIdForService(context.Background(), serviceId, ospfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetRoutingOspfProfileParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutingOspfProfileParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetRoutingOspfProfileParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**ospfId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingOspfProfileParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingOspfProfileParcelForService

> string GetRoutingOspfProfileParcelForService(ctx, serviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetRoutingOspfProfileParcelForService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetRoutingOspfProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutingOspfProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetRoutingOspfProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingOspfProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanFeatureProfileByProfileId

> map[string]interface{} GetSdwanFeatureProfileByProfileId(ctx, cliId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cliId := "cliId_example" // string | Feature Profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfileByProfileId(context.Background(), cliId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfileByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanFeatureProfileByProfileId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfileByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanFeatureProfileByProfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanFeatureProfileBySdwanFamily

> map[string]interface{} GetSdwanFeatureProfileBySdwanFamily(ctx).Offset(offset).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    offset := int32(56) // int32 | Pagination offset (optional)
    limit := int32(56) // int32 | Pagination limit (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfileBySdwanFamily(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfileBySdwanFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanFeatureProfileBySdwanFamily`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfileBySdwanFamily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanFeatureProfileBySdwanFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Pagination offset | 
 **limit** | **int32** | Pagination limit | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanFeatureProfilesByFamilyAndType

> map[string]interface{} GetSdwanFeatureProfilesByFamilyAndType(ctx).Offset(offset).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    offset := int32(56) // int32 | Pagination offset (optional)
    limit := int32(56) // int32 | Pagination limit (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfilesByFamilyAndType(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfilesByFamilyAndType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanFeatureProfilesByFamilyAndType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanFeatureProfilesByFamilyAndType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanFeatureProfilesByFamilyAndTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Pagination offset | 
 **limit** | **int32** | Pagination limit | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceFeatureProfileByProfileId

> map[string]interface{} GetSdwanServiceFeatureProfileByProfileId(ctx, serviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceFeatureProfileByProfileId(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceFeatureProfileByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceFeatureProfileByProfileId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceFeatureProfileByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceFeatureProfileByProfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceFeatureProfiles

> map[string]interface{} GetSdwanServiceFeatureProfiles(ctx).Offset(offset).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    offset := int32(56) // int32 | Pagination offset (optional)
    limit := int32(56) // int32 | Pagination limit (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceFeatureProfiles(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceFeatureProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceFeatureProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceFeatureProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceFeatureProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Pagination offset | 
 **limit** | **int32** | Pagination limit | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema

> string GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceLanVpnInterfaceEthernetParcelSchemaBySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceLanVpnParcelSchemaBySchemaType

> string GetSdwanServiceLanVpnParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceLanVpnParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceLanVpnParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceLanVpnParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceLanVpnParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceLanVpnParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceRoutingBgpParcelSchemaBySchemaType

> string GetSdwanServiceRoutingBgpParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceRoutingBgpParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceRoutingBgpParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceRoutingBgpParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceRoutingBgpParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceRoutingBgpParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceRoutingOspfParcelSchemaBySchemaType

> string GetSdwanServiceRoutingOspfParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceRoutingOspfParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceRoutingOspfParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceRoutingOspfParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceRoutingOspfParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceRoutingOspfParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanServiceTrackerParcelSchemaBySchemaType

> string GetSdwanServiceTrackerParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanServiceTrackerParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceTrackerParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanServiceTrackerParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanServiceTrackerParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanServiceTrackerParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemAaaParcelSchemaBySchemaType

> string GetSdwanSystemAaaParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemAaaParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemAaaParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemAaaParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemAaaParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemAaaParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemBannerParcelSchemaBySchemaType

> string GetSdwanSystemBannerParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBannerParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBannerParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemBannerParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBannerParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemBannerParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemBasicParcelSchemaBySchemaType

> string GetSdwanSystemBasicParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBasicParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBasicParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemBasicParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBasicParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemBasicParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemBfdParcelSchemaBySchemaType

> string GetSdwanSystemBfdParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBfdParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBfdParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemBfdParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemBfdParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemBfdParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemFeatureProfileByProfileId

> map[string]interface{} GetSdwanSystemFeatureProfileByProfileId(ctx, systemId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | Feature Profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemFeatureProfileByProfileId(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemFeatureProfileByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemFeatureProfileByProfileId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemFeatureProfileByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemFeatureProfileByProfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemFeatureProfiles

> map[string]interface{} GetSdwanSystemFeatureProfiles(ctx).Offset(offset).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    offset := int32(56) // int32 | Pagination offset (optional)
    limit := int32(56) // int32 | Pagination limit (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemFeatureProfiles(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemFeatureProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemFeatureProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemFeatureProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemFeatureProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Pagination offset | 
 **limit** | **int32** | Pagination limit | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemLoggingParcelSchemaBySchemaType

> string GetSdwanSystemLoggingParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemLoggingParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemLoggingParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemLoggingParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemLoggingParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemLoggingParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemNtpParcelSchemaBySchemaType

> string GetSdwanSystemNtpParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemNtpParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemNtpParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemNtpParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemNtpParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemNtpParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanSystemOmpParcelSchemaBySchemaType

> string GetSdwanSystemOmpParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanSystemOmpParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemOmpParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanSystemOmpParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanSystemOmpParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanSystemOmpParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportCellularControllerParcelSchemaBySchemaType

> string GetSdwanTransportCellularControllerParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportCellularControllerParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportCellularControllerParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportCellularControllerParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportCellularControllerParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportCellularControllerParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportCellularProfileParcelSchemaBySchemaType

> string GetSdwanTransportCellularProfileParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportCellularProfileParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportCellularProfileParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportCellularProfileParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportCellularProfileParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportCellularProfileParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportFeatureProfileByProfileId

> map[string]interface{} GetSdwanTransportFeatureProfileByProfileId(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportFeatureProfileByProfileId(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportFeatureProfileByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportFeatureProfileByProfileId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportFeatureProfileByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportFeatureProfileByProfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportFeatureProfiles

> map[string]interface{} GetSdwanTransportFeatureProfiles(ctx).Offset(offset).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    offset := int32(56) // int32 | Pagination offset (optional)
    limit := int32(56) // int32 | Pagination limit (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportFeatureProfiles(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportFeatureProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportFeatureProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportFeatureProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportFeatureProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Pagination offset | 
 **limit** | **int32** | Pagination limit | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema

> string GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportManagementVpnInterfaceEthernetParcelSchemaBySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportManagementVpnParcelSchemaBySchemaType

> string GetSdwanTransportManagementVpnParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportManagementVpnParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportManagementVpnParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportManagementVpnParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportManagementVpnParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportManagementVpnParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportTrackerParcelSchemaBySchemaType

> string GetSdwanTransportTrackerParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportTrackerParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportTrackerParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportTrackerParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportTrackerParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportTrackerParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema

> string GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportWanVpnInterfaceEthernetParcelSchemaBySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdwanTransportWanVpnParcelSchemaBySchemaType

> string GetSdwanTransportWanVpnParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetSdwanTransportWanVpnParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportWanVpnParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSdwanTransportWanVpnParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetSdwanTransportWanVpnParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSdwanTransportWanVpnParcelSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackerProfileParcelByParcelIdForService

> string GetTrackerProfileParcelByParcelIdForService(ctx, serviceId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID
    trackerId := "trackerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelByParcelIdForService(context.Background(), serviceId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelByParcelIdForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackerProfileParcelByParcelIdForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelByParcelIdForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 
**trackerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackerProfileParcelByParcelIdForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackerProfileParcelByParcelIdForTransport

> string GetTrackerProfileParcelByParcelIdForTransport(ctx, transportId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    trackerId := "trackerId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelByParcelIdForTransport(context.Background(), transportId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackerProfileParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**trackerId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackerProfileParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackerProfileParcelForService

> string GetTrackerProfileParcelForService(ctx, serviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelForService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackerProfileParcelForService`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackerProfileParcelForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackerProfileParcelForTransport

> string GetTrackerProfileParcelForTransport(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelForTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackerProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetTrackerProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackerProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport

> string GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport(ctx, transportId, vpnId, ethernetId, trackerId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID
    trackerId := "trackerId_example" // string | Tracker Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport(context.Background(), transportId, vpnId, ethernetId, trackerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 
**trackerId** | **string** | Tracker Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanVpnInterfaceEthernetAssociatedTrackerParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport

> string GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport(ctx, transportId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Feature Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport(context.Background(), transportId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Feature Parcel ID | 
**ethernetId** | **string** | Interface Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanVpnInterfaceEthernetAssociatedTrackerParcelsForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanVpnInterfaceEthernetParcelByParcelIdForTransport

> string GetWanVpnInterfaceEthernetParcelByParcelIdForTransport(ctx, transportId, vpnId, ethernetId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    ethernetId := "ethernetId_example" // string | Interface Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetParcelByParcelIdForTransport(context.Background(), transportId, vpnId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanVpnInterfaceEthernetParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetWanVpnInterfaceEthernetParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 
**ethernetId** | **string** | Interface Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanVpnInterfaceEthernetParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanVpnProfileParcelByParcelIdForTransport

> string GetWanVpnProfileParcelByParcelIdForTransport(ctx, transportId, vpnId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetWanVpnProfileParcelByParcelIdForTransport(context.Background(), transportId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetWanVpnProfileParcelByParcelIdForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanVpnProfileParcelByParcelIdForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetWanVpnProfileParcelByParcelIdForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanVpnProfileParcelByParcelIdForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanVpnProfileParcelForTransport

> string GetWanVpnProfileParcelForTransport(ctx, transportId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transportId := "transportId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileSDWANApi.GetWanVpnProfileParcelForTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileSDWANApi.GetWanVpnProfileParcelForTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanVpnProfileParcelForTransport`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileSDWANApi.GetWanVpnProfileParcelForTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanVpnProfileParcelForTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

