
## ADC Use-Case supported through Terraform

| ADC Use Case | Configuration Examples |
| ----------- | ---------------------- |
| Load Balancing | lbvserver <br /> lbvserver\_cmppolicy\_binding <br /> lbvserver\_filterpolicy\_binding <br /> lbvserver\_responderpolicy\_binding <br /> lbvserver\_rewritepolicy\_binding <br /> lbvserver\_service\_binding <br /> lbvserver\_transformpolicy\_binding <br /> simple\_lb <br /> simple\_server |
| Content Switching | csvserver <br /> csvserver\_cmppolicy\_binding <br /> csvserver\_cspolicy\_binding <br /> csvserver\_filterpolicy\_binding <br /> csvserver\_lbvserver\_binding <br /> csvserver\_responderpolicy\_binding <br /> csvserver\_rewritepolicy\_binding <br /> csvserver\_transformpolicy\_binding <br /> content\_switch\_ssl\_lb\_mon |
| Responder/Rewrite Policies | csvserver\_responderpolicy\_binding <br /> lbvserver\_responderpolicy\_binding <br /> responder <br /> csvserver\_rewritepolicy\_binding <br /> lbvserver\_rewritepolicy\_binding <br /> rewrite |
| SSL | content\_switch\_ssl\_lb\_mon <br /> ssl\_lb\_monitors <br /> sslaction <br /> sslcipher <br /> ssldhparam <br /> sslparameter <br /> sslpolicy <br /> sslprofile <br /> sslprofile\_sslcipher\_binding <br /> sslvserver <br /> sslvserver\_sslcertkey\_binding <br /> sslvserver\_sslpolicy\_binding |
| Global Load Server Balancing (GSLB) | gslb <br /> gslbservice\_lbmonbindings  |
| Web Application Firewall (WAF) | appfwfieldtype <br /> appfwjsoncontenttype <br /> appfwpolicy <br /> appfwpolicylabel <br /> appfwprofile <br /> appfwprofile\_denyurl\_binding <br /> appfwprofile\_starturl\_binding <br /> appfwjsonerrorpage <br />  appfwprofile_fileuploadtype_binding <br />  appfwxmlschema <br />  appfwhtmlerrorpage <br />  appfwprofile_contenttype_binding <br />  appfwlearningsettings <br />  appfwglobal_appfwpolicy_binding <br />  appfwsettings <br />  appfwprofile_excluderescontenttype_binding <br />  appfwprofile_sqlinjection_binding <br />  appfwurlencodedformcontenttype <br />  appfwxmlcontenttype <br />  appfwxmlerrorpage <br />  appfwprofile_csrftag_binding <br />  appfwconfidfield <br />  appfwprofile_xmlwsiurl_binding <br />  appfwprofile_cookieconsistency_binding <br />  appfwprofile_jsoncmdurl_binding <br />  appfwprofile_xmldosurl_binding <br />  appfwprofile_xmlxss_binding <br />  appfwprofile_jsondosurl_binding <br />  appfwprofile_fieldformat_binding  <br />  appfwprofile_cmdinjection_binding <br />  appfwprofile_creditcardnumber_binding <br />  appfwpolicylabel_appfwpolicy_binding <br />  appfwprofile_jsonsqlurl_binding <br />  appfwsignatures <br />  appfwprofile_xmlsqlinjection_binding <br />  appfwglobal_auditnslogpolicy_binding <br />  appfwprofile_crosssitescripting_binding <br />  appfwmultipartformcontenttype <br />  appfwprofile_fieldconsistency_binding <br />  appfwprofile_jsonxssurl_binding <br />  appfwglobal_auditsyslogpolicy_binding <br />  appfwprofile_safeobject_binding <br />  appfwwsdl <br />  appfwprofile_xmlattachmenturl_binding <br />  appfwprofile_xmlvalidationurl_binding <br />    appfwprofile_logexpression_binding <br />  appfwprofile_trustedlearningclients_binding |
| Core ADC features | cluster <br /> clusterL3 <br /> installer <br /> interface <br /> ipset <br /> linkset <br /> netprofile <br /> network <br /> nsacl <br /> nsconfig <br /> nsfeature <br /> nshttprofile <br /> nsip <br /> nsip6 <br /> nsparam <br /> nsrpcnode <br /> nstcpprofile <br /> password\_resetter <br /> pinger <br /> rebooter <br /> route <br /> routerdynamicrouting <br /> saveconfig <br /> systemfile <br /> systemgroup <br /> systemuser |
| Pool Licensing | nscapacity\_cico <br /> nscapacity\_pooled <br /> nscapacity\_vcpu <br /> nslicenseserver <br /> nslicense |
| SSL VPN | vpnglobal_auditsyslogpolicy_binding <br /> vpnglobal_authenticationcertpolicy_binding <br /> vpnglobal_authenticationldappolicy_binding <br /> vpnglobal_authenticationlocalpolicy_binding <br /> vpnglobal_authenticationnegotiatepolicy_binding <br /> vpnglobal_authenticationradiuspolicy_binding <br /> vpnglobal_authenticationsamlpolicy_binding <br /> vpnglobal_authenticationtacacspolicy_binding <br /> vpnglobal_domain_binding <br /> vpnglobal_intranetip_binding <br /> vpnglobal_intranetip6_binding <br /> vpnglobal_sharefileserver_binding <br /> vpnglobal_sslcertkey_binding <br /> vpnglobal_staserver_binding <br /> vpnglobal_vpnclientlessaccesspolicy_binding <br /> vpnglobal_vpneula_binding <br /> vpnglobal_vpnintranetapplication_binding <br /> vpnglobal_vpnnexthopserver_binding <br /> vpnglobal_vpnportaltheme_binding <br /> vpnglobal_vpnsessionpolicy_binding <br /> vpnglobal_vpntrafficpolicy_binding <br /> vpnglobal_vpnurl_binding <br /> vpnglobal_vpnurlpolicy_binding <br /> vpnintranetapplication <br /> vpnnexthopserver <br /> vpnparameter <br /> vpnpcoipprofile <br /> vpnpcoipvserverprofile <br /> vpnportaltheme <br /> vpnsamlssoprofile <br />  vpnsessionaction <br /> vpnsessionpolicy <br /> vpntrafficaction <br /> vpntrafficpolicy <br /> vpnurl <br /> vpnurlaction <br /> vpnurlpolicy <br /> vpnvserver_aaapreauthenticationpolicy_binding <br /> vpnvserver_analyticsprofile_binding <br />  vpnvserver_appcontroller_binding <br /> vpnvserver_appflowpolicy_binding <br /> vpnvserver_auditnslogpolicy_binding <br /> vpnvserver_auditsyslogpolicy_binding <br /> vpnvserver_authenticationcertpolicy_binding <br /> vpnvserver_authenticationdfapolicy_binding <br /> vpnvserver_authenticationldappolicy_binding <br />  vpnvserver_authenticationlocalpolicy_binding <br /> vpnvserver_authenticationloginschemapolicy_binding <br />  vpnvserver_authenticationnegotiatepolicy_binding <br /> vpnvserver_authenticationoauthidppolicy_binding <br /> vpnvserver_authenticationradiuspolicy_binding <br /> vpnvserver_authenticationsamlidppolicy_binding <br /> vpnvserver_authenticationsamlpolicy_binding <br /> vpnvserver_authenticationtacacspolicy_binding <br /> vpnvserver_authenticationwebauthpolicy_binding <br /> vpnvserver_cachepolicy_binding <br /> vpnvserver_cspolicy_binding <br /> vpnvserver_feopolicy_binding <br /> vpnvserver_icapolicy_binding <br /> vpnvserver_intranetip_binding <br /> vpnvserver_intranetip6_binding <br /> vpnvserver_responderpolicy_binding <br /> vpnvserver_rewritepolicy_binding <br /> vpnvserver_sharefileserver_binding <br /> vpnvserver_staserver_binding <br /> vpnvserver_vpnclientlessaccesspolicy_binding <br /> vpnvserver_vpneula_binding <br /> vpnvserver_vpnintranetapplication_binding <br /> vpnvserver_vpnnexthopserver_binding <br /> vpnvserver_vpnportaltheme_binding <br /> vpnvserver_vpnsessionpolicy_binding <br /> vpnvserver_vpntrafficpolicy_binding <br /> vpnvserver_vpnurl_binding <br /> vpnvserver_vpnurlpolicy_binding |
| Authentication | authenticationvserver_authenticationcertpolicy_binding <br />  authenticationvserver_rewritepolicy_binding <br />  authenticationcertaction <br />  authenticationcitrixauthaction <br />  authenticationldapaction <br />  authenticationvserver_authenticationldappolicy_binding <br />  authenticationsamlidppolicy <br />  authenticationoauthaction <br />  authenticationvserver_vpnportaltheme_binding <br />  authenticationvserver_authenticationwebauthpolicy_binding <br />  authenticationldappolicy <br />  authenticationvserver_auditsyslogpolicy_binding <br />  authenticationvserver_authenticationoauthidppolicy_binding <br />  authenticationvserver_cachepolicy_binding <br />  authenticationcertpolicy <br />  authenticationauthnprofile <br />  authenticationsamlidpprofile <br />  authenticationnegotiatepolicy <br />  authenticationtacacsaction <br />  authenticationvserver <br />  authenticationvserver_authenticationnegotiatepolicy_binding <br />  authenticationlocalpolicy <br />  authenticationvserver_auditnslogpolicy_binding <br />  authenticationnoauthaction <br />  authenticationldappolicy_systemglobal_binding <br />  authenticationvserver_authenticationsamlpolicy_binding <br />  authenticationnegotiateaction <br />  authenticationvserver_authenticationtacacspolicy_binding <br />  authenticationtacacspolicy <br />  authenticationvserver_authenticationlocalpolicy_binding <br />  authenticationpushservice <br />  authenticationwebauthaction <br />  authenticationcaptchaaction <br />  authenticationvserver_authenticationpolicy_binding <br />  authenticationradiuspolicy <br />  authenticationvserver_tmsessionpolicy_binding <br />  authenticationradiusaction <br />  authenticationloginschemapolicy <br />  authenticationwebauthpolicy <br />  authenticationvserver_cspolicy_binding <br />  authenticationstorefrontauthaction <br />  authenticationvserver_responderpolicy_binding <br />  authenticationvserver_authenticationsamlidppolicy_binding  <br />  authenticationemailaction <br />  authenticationdfapolicy <br />  authenticationsamlpolicy <br />  authenticationloginschema <br />  authenticationoauthidpprofile <br />  authenticationepaaction <br />  authenticationvserver_authenticationloginschemapolicy_binding <br />  authenticationpolicylabel_authenticationpolicy_binding <br />  authenticationdfaaction <br />  authenticationpolicylabel <br />  authenticationsamlaction <br />  authenticationvserver_authenticationradiuspolicy_binding <br />  authenticationoauthidppolicy <br />  authenticationpolicy |
| BOT | botprofile_captcha_binding <br />  botpolicylabel <br />  botprofile <br />  botprofile_tps_binding <br />  botprofile_trapinsertionurl_binding <br />  botprofile_logexpression_binding <br />  botprofile_whitelist_binding <br />  botpolicylabel_botpolicy_binding <br />  botsettings <br />  botprofile_ratelimit_binding <br />  botsignature <br />  botprofile_ipreputation_binding <br />  botglobal_botpolicy_binding <br />  botpolicy <br />  botprofile_blacklist_binding |
| Basic | locationparameter <br /> radiusnode <br /> servicegroup_lbmonitor_binding <br /> service_lbmonitor_binding <br /> servicegroup_servicegroupmember_binding <br /> servicegroup <br /> service_dospolicy_binding <br /> extendedmemoryparam <br /> service <br /> location <br /> server |
| NS | nsconsoleloginprompt <br />  nsdiameter <br />  nsassignment <br />  nsspparams <br />  nstrafficdomain_vlan_binding <br />  nstimer <br />  nsratecontrol <br />  nsservicepath_nsservicefunction_binding <br />  nsicapprofile <br />  nshmackey <br />  nslimitidentifier <br />  nsacl6 <br />  nspartition_vxlan_binding <br />  nsdhcpparams <br />  nshttpparam <br />  nspbr6 <br />  nslicenseserver <br />  nspartition_bridgegroup_binding <br />  nssimpleacl6 <br />  nspartition <br />  nsacl <br />  nsxmlnamespace <br />  nstcpbufparam <br />  nstrafficdomain_bridgegroup_binding <br />  nsappflowcollector <br />  nshttpprofile <br />  nssimpleacl <br />  nsservicepath <br />  nspartition_vlan_binding <br />  nsservicefunction <br />  nsencryptionkey <br />  nsvpxparam <br />  nslicenseparameters <br />  nsvariable <br />  nstrafficdomain <br />  nstcpparam <br />  nstrafficdomain_vxlan_binding |
| Network | ptp  <br />  netbridge_iptunnel_binding  <br />  mapdomain_mapbmr_binding  <br />  ip6tunnel  <br />  netbridge_vlan_binding  <br />  bridgetable  <br />  mapdomain  <br />  bridgegroup_vlan_binding  <br />  nat64param  <br />  bridgegroup  <br />  rsskeytype  <br />  netprofile_natrule_binding  <br />  iptunnelparam  <br />  appalgparam  <br />  vxlanvlanmap_vxlan_binding  <br />  bridgegroup_nsip6_binding  <br />  mapbmr  <br />  vridparam  <br />  arpparam  <br />  vrid  <br />  bridgegroup_nsip_binding  <br />  vrid6  <br />  vxlan_nsip6_binding  <br />  mapbmr_bmrv4network_binding  <br />  mapdmr  <br />  vxlan  <br />  fis  <br />  netprofile_srcportset_binding  <br />  rnatparam  <br />  l4param  <br />  netbridge  |
