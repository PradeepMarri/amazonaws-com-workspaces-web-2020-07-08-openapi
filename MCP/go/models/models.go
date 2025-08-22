package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DisassociateBrowserSettingsRequest represents the DisassociateBrowserSettingsRequest schema from the OpenAPI specification
type DisassociateBrowserSettingsRequest struct {
}

// ListBrowserSettingsRequest represents the ListBrowserSettingsRequest schema from the OpenAPI specification
type ListBrowserSettingsRequest struct {
}

// AssociateTrustStoreRequest represents the AssociateTrustStoreRequest schema from the OpenAPI specification
type AssociateTrustStoreRequest struct {
}

// UpdateIpAccessSettingsResponse represents the UpdateIpAccessSettingsResponse schema from the OpenAPI specification
type UpdateIpAccessSettingsResponse struct {
	Ipaccesssettings interface{} `json:"ipAccessSettings"`
}

// AssociateUserAccessLoggingSettingsResponse represents the AssociateUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type AssociateUserAccessLoggingSettingsResponse struct {
	Portalarn interface{} `json:"portalArn"`
	Useraccessloggingsettingsarn interface{} `json:"userAccessLoggingSettingsArn"`
}

// AssociateNetworkSettingsResponse represents the AssociateNetworkSettingsResponse schema from the OpenAPI specification
type AssociateNetworkSettingsResponse struct {
	Networksettingsarn interface{} `json:"networkSettingsArn"`
	Portalarn interface{} `json:"portalArn"`
}

// CreateBrowserSettingsRequest represents the CreateBrowserSettingsRequest schema from the OpenAPI specification
type CreateBrowserSettingsRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Additionalencryptioncontext interface{} `json:"additionalEncryptionContext,omitempty"`
	Browserpolicy interface{} `json:"browserPolicy"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Customermanagedkey interface{} `json:"customerManagedKey,omitempty"`
}

// AssociateUserSettingsResponse represents the AssociateUserSettingsResponse schema from the OpenAPI specification
type AssociateUserSettingsResponse struct {
	Portalarn interface{} `json:"portalArn"`
	Usersettingsarn interface{} `json:"userSettingsArn"`
}

// DeleteUserAccessLoggingSettingsResponse represents the DeleteUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type DeleteUserAccessLoggingSettingsResponse struct {
}

// ListIpAccessSettingsResponse represents the ListIpAccessSettingsResponse schema from the OpenAPI specification
type ListIpAccessSettingsResponse struct {
	Ipaccesssettings interface{} `json:"ipAccessSettings,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdatePortalResponse represents the UpdatePortalResponse schema from the OpenAPI specification
type UpdatePortalResponse struct {
	Portal interface{} `json:"portal,omitempty"`
}

// CreateNetworkSettingsResponse represents the CreateNetworkSettingsResponse schema from the OpenAPI specification
type CreateNetworkSettingsResponse struct {
	Networksettingsarn interface{} `json:"networkSettingsArn"`
}

// BrowserSettingsSummary represents the BrowserSettingsSummary schema from the OpenAPI specification
type BrowserSettingsSummary struct {
	Browsersettingsarn interface{} `json:"browserSettingsArn,omitempty"`
}

// GetPortalRequest represents the GetPortalRequest schema from the OpenAPI specification
type GetPortalRequest struct {
}

// AssociateTrustStoreResponse represents the AssociateTrustStoreResponse schema from the OpenAPI specification
type AssociateTrustStoreResponse struct {
	Portalarn interface{} `json:"portalArn"`
	Truststorearn interface{} `json:"trustStoreArn"`
}

// DisassociateNetworkSettingsRequest represents the DisassociateNetworkSettingsRequest schema from the OpenAPI specification
type DisassociateNetworkSettingsRequest struct {
}

// GetUserSettingsResponse represents the GetUserSettingsResponse schema from the OpenAPI specification
type GetUserSettingsResponse struct {
	Usersettings interface{} `json:"userSettings,omitempty"`
}

// DisassociateIpAccessSettingsResponse represents the DisassociateIpAccessSettingsResponse schema from the OpenAPI specification
type DisassociateIpAccessSettingsResponse struct {
}

// DisassociateIpAccessSettingsRequest represents the DisassociateIpAccessSettingsRequest schema from the OpenAPI specification
type DisassociateIpAccessSettingsRequest struct {
}

// TrustStoreSummary represents the TrustStoreSummary schema from the OpenAPI specification
type TrustStoreSummary struct {
	Truststorearn interface{} `json:"trustStoreArn,omitempty"`
}

// ListBrowserSettingsResponse represents the ListBrowserSettingsResponse schema from the OpenAPI specification
type ListBrowserSettingsResponse struct {
	Browsersettings interface{} `json:"browserSettings,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdatePortalRequest represents the UpdatePortalRequest schema from the OpenAPI specification
type UpdatePortalRequest struct {
	Authenticationtype interface{} `json:"authenticationType,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
}

// IpAccessSettings represents the IpAccessSettings schema from the OpenAPI specification
type IpAccessSettings struct {
	Associatedportalarns interface{} `json:"associatedPortalArns,omitempty"`
	Creationdate interface{} `json:"creationDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Ipaccesssettingsarn interface{} `json:"ipAccessSettingsArn"`
	Iprules interface{} `json:"ipRules,omitempty"`
}

// CreateBrowserSettingsResponse represents the CreateBrowserSettingsResponse schema from the OpenAPI specification
type CreateBrowserSettingsResponse struct {
	Browsersettingsarn interface{} `json:"browserSettingsArn"`
}

// IpAccessSettingsSummary represents the IpAccessSettingsSummary schema from the OpenAPI specification
type IpAccessSettingsSummary struct {
	Creationdate interface{} `json:"creationDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Ipaccesssettingsarn interface{} `json:"ipAccessSettingsArn,omitempty"`
}

// UpdateIdentityProviderResponse represents the UpdateIdentityProviderResponse schema from the OpenAPI specification
type UpdateIdentityProviderResponse struct {
	Identityprovider interface{} `json:"identityProvider"`
}

// UserAccessLoggingSettings represents the UserAccessLoggingSettings schema from the OpenAPI specification
type UserAccessLoggingSettings struct {
	Associatedportalarns interface{} `json:"associatedPortalArns,omitempty"`
	Kinesisstreamarn interface{} `json:"kinesisStreamArn,omitempty"`
	Useraccessloggingsettingsarn interface{} `json:"userAccessLoggingSettingsArn"`
}

// ListPortalsResponse represents the ListPortalsResponse schema from the OpenAPI specification
type ListPortalsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Portals interface{} `json:"portals,omitempty"`
}

// ListUserAccessLoggingSettingsResponse represents the ListUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type ListUserAccessLoggingSettingsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Useraccessloggingsettings interface{} `json:"userAccessLoggingSettings,omitempty"`
}

// CreateUserAccessLoggingSettingsResponse represents the CreateUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type CreateUserAccessLoggingSettingsResponse struct {
	Useraccessloggingsettingsarn interface{} `json:"userAccessLoggingSettingsArn"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ListTrustStoresRequest represents the ListTrustStoresRequest schema from the OpenAPI specification
type ListTrustStoresRequest struct {
}

// DisassociateUserSettingsRequest represents the DisassociateUserSettingsRequest schema from the OpenAPI specification
type DisassociateUserSettingsRequest struct {
}

// GetPortalServiceProviderMetadataResponse represents the GetPortalServiceProviderMetadataResponse schema from the OpenAPI specification
type GetPortalServiceProviderMetadataResponse struct {
	Portalarn interface{} `json:"portalArn"`
	Serviceprovidersamlmetadata interface{} `json:"serviceProviderSamlMetadata,omitempty"`
}

// DeleteUserSettingsRequest represents the DeleteUserSettingsRequest schema from the OpenAPI specification
type DeleteUserSettingsRequest struct {
}

// UpdateBrowserSettingsResponse represents the UpdateBrowserSettingsResponse schema from the OpenAPI specification
type UpdateBrowserSettingsResponse struct {
	Browsersettings interface{} `json:"browserSettings"`
}

// UpdateNetworkSettingsRequest represents the UpdateNetworkSettingsRequest schema from the OpenAPI specification
type UpdateNetworkSettingsRequest struct {
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Subnetids interface{} `json:"subnetIds,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// IpRule represents the IpRule schema from the OpenAPI specification
type IpRule struct {
	Description interface{} `json:"description,omitempty"`
	Iprange interface{} `json:"ipRange"`
}

// Portal represents the Portal schema from the OpenAPI specification
type Portal struct {
	Usersettingsarn interface{} `json:"userSettingsArn,omitempty"`
	Renderertype interface{} `json:"rendererType,omitempty"`
	Creationdate interface{} `json:"creationDate,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Networksettingsarn interface{} `json:"networkSettingsArn,omitempty"`
	Portalendpoint interface{} `json:"portalEndpoint,omitempty"`
	Portalstatus interface{} `json:"portalStatus,omitempty"`
	Browsertype interface{} `json:"browserType,omitempty"`
	Authenticationtype interface{} `json:"authenticationType,omitempty"`
	Truststorearn interface{} `json:"trustStoreArn,omitempty"`
	Browsersettingsarn interface{} `json:"browserSettingsArn,omitempty"`
	Ipaccesssettingsarn interface{} `json:"ipAccessSettingsArn,omitempty"`
	Portalarn interface{} `json:"portalArn,omitempty"`
	Useraccessloggingsettingsarn interface{} `json:"userAccessLoggingSettingsArn,omitempty"`
	Statusreason interface{} `json:"statusReason,omitempty"`
}

// UpdateTrustStoreResponse represents the UpdateTrustStoreResponse schema from the OpenAPI specification
type UpdateTrustStoreResponse struct {
	Truststorearn interface{} `json:"trustStoreArn"`
}

// ListTrustStoreCertificatesResponse represents the ListTrustStoreCertificatesResponse schema from the OpenAPI specification
type ListTrustStoreCertificatesResponse struct {
	Certificatelist interface{} `json:"certificateList,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Truststorearn interface{} `json:"trustStoreArn,omitempty"`
}

// DeleteIpAccessSettingsRequest represents the DeleteIpAccessSettingsRequest schema from the OpenAPI specification
type DeleteIpAccessSettingsRequest struct {
}

// GetUserSettingsRequest represents the GetUserSettingsRequest schema from the OpenAPI specification
type GetUserSettingsRequest struct {
}

// ListIdentityProvidersRequest represents the ListIdentityProvidersRequest schema from the OpenAPI specification
type ListIdentityProvidersRequest struct {
}

// UpdateIpAccessSettingsRequest represents the UpdateIpAccessSettingsRequest schema from the OpenAPI specification
type UpdateIpAccessSettingsRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Iprules interface{} `json:"ipRules,omitempty"`
}

// UserAccessLoggingSettingsSummary represents the UserAccessLoggingSettingsSummary schema from the OpenAPI specification
type UserAccessLoggingSettingsSummary struct {
	Kinesisstreamarn interface{} `json:"kinesisStreamArn,omitempty"`
	Useraccessloggingsettingsarn interface{} `json:"userAccessLoggingSettingsArn,omitempty"`
}

// IdentityProvider represents the IdentityProvider schema from the OpenAPI specification
type IdentityProvider struct {
	Identityprovidertype interface{} `json:"identityProviderType,omitempty"`
	Identityproviderarn interface{} `json:"identityProviderArn"`
	Identityproviderdetails interface{} `json:"identityProviderDetails,omitempty"`
	Identityprovidername interface{} `json:"identityProviderName,omitempty"`
}

// UpdateBrowserSettingsRequest represents the UpdateBrowserSettingsRequest schema from the OpenAPI specification
type UpdateBrowserSettingsRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Browserpolicy interface{} `json:"browserPolicy,omitempty"`
}

// GetTrustStoreResponse represents the GetTrustStoreResponse schema from the OpenAPI specification
type GetTrustStoreResponse struct {
	Truststore interface{} `json:"trustStore,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// CreateTrustStoreRequest represents the CreateTrustStoreRequest schema from the OpenAPI specification
type CreateTrustStoreRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Certificatelist interface{} `json:"certificateList"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// GetTrustStoreCertificateRequest represents the GetTrustStoreCertificateRequest schema from the OpenAPI specification
type GetTrustStoreCertificateRequest struct {
}

// CreateIdentityProviderResponse represents the CreateIdentityProviderResponse schema from the OpenAPI specification
type CreateIdentityProviderResponse struct {
	Identityproviderarn interface{} `json:"identityProviderArn"`
}

// UpdateTrustStoreRequest represents the UpdateTrustStoreRequest schema from the OpenAPI specification
type UpdateTrustStoreRequest struct {
	Certificatestoadd interface{} `json:"certificatesToAdd,omitempty"`
	Certificatestodelete interface{} `json:"certificatesToDelete,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// NetworkSettings represents the NetworkSettings schema from the OpenAPI specification
type NetworkSettings struct {
	Associatedportalarns interface{} `json:"associatedPortalArns,omitempty"`
	Networksettingsarn interface{} `json:"networkSettingsArn"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Subnetids interface{} `json:"subnetIds,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
}

// GetBrowserSettingsResponse represents the GetBrowserSettingsResponse schema from the OpenAPI specification
type GetBrowserSettingsResponse struct {
	Browsersettings interface{} `json:"browserSettings,omitempty"`
}

// AssociateNetworkSettingsRequest represents the AssociateNetworkSettingsRequest schema from the OpenAPI specification
type AssociateNetworkSettingsRequest struct {
}

// NetworkSettingsSummary represents the NetworkSettingsSummary schema from the OpenAPI specification
type NetworkSettingsSummary struct {
	Networksettingsarn interface{} `json:"networkSettingsArn,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
}

// DeleteTrustStoreResponse represents the DeleteTrustStoreResponse schema from the OpenAPI specification
type DeleteTrustStoreResponse struct {
}

// DeleteIdentityProviderRequest represents the DeleteIdentityProviderRequest schema from the OpenAPI specification
type DeleteIdentityProviderRequest struct {
}

// UpdateIdentityProviderRequest represents the UpdateIdentityProviderRequest schema from the OpenAPI specification
type UpdateIdentityProviderRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Identityproviderdetails interface{} `json:"identityProviderDetails,omitempty"`
	Identityprovidername interface{} `json:"identityProviderName,omitempty"`
	Identityprovidertype interface{} `json:"identityProviderType,omitempty"`
}

// GetTrustStoreCertificateResponse represents the GetTrustStoreCertificateResponse schema from the OpenAPI specification
type GetTrustStoreCertificateResponse struct {
	Certificate interface{} `json:"certificate,omitempty"`
	Truststorearn interface{} `json:"trustStoreArn,omitempty"`
}

// DeleteUserSettingsResponse represents the DeleteUserSettingsResponse schema from the OpenAPI specification
type DeleteUserSettingsResponse struct {
}

// GetIdentityProviderRequest represents the GetIdentityProviderRequest schema from the OpenAPI specification
type GetIdentityProviderRequest struct {
}

// GetIpAccessSettingsRequest represents the GetIpAccessSettingsRequest schema from the OpenAPI specification
type GetIpAccessSettingsRequest struct {
}

// TrustStore represents the TrustStore schema from the OpenAPI specification
type TrustStore struct {
	Truststorearn interface{} `json:"trustStoreArn,omitempty"`
	Associatedportalarns interface{} `json:"associatedPortalArns,omitempty"`
}

// DisassociateUserAccessLoggingSettingsRequest represents the DisassociateUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type DisassociateUserAccessLoggingSettingsRequest struct {
}

// DeleteIpAccessSettingsResponse represents the DeleteIpAccessSettingsResponse schema from the OpenAPI specification
type DeleteIpAccessSettingsResponse struct {
}

// CreatePortalRequest represents the CreatePortalRequest schema from the OpenAPI specification
type CreatePortalRequest struct {
	Additionalencryptioncontext interface{} `json:"additionalEncryptionContext,omitempty"`
	Authenticationtype interface{} `json:"authenticationType,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Customermanagedkey interface{} `json:"customerManagedKey,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// DisassociateBrowserSettingsResponse represents the DisassociateBrowserSettingsResponse schema from the OpenAPI specification
type DisassociateBrowserSettingsResponse struct {
}

// ListUserAccessLoggingSettingsRequest represents the ListUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type ListUserAccessLoggingSettingsRequest struct {
}

// IdentityProviderSummary represents the IdentityProviderSummary schema from the OpenAPI specification
type IdentityProviderSummary struct {
	Identityproviderarn interface{} `json:"identityProviderArn,omitempty"`
	Identityprovidername interface{} `json:"identityProviderName,omitempty"`
	Identityprovidertype interface{} `json:"identityProviderType,omitempty"`
}

// UpdateNetworkSettingsResponse represents the UpdateNetworkSettingsResponse schema from the OpenAPI specification
type UpdateNetworkSettingsResponse struct {
	Networksettings interface{} `json:"networkSettings"`
}

// DisassociateNetworkSettingsResponse represents the DisassociateNetworkSettingsResponse schema from the OpenAPI specification
type DisassociateNetworkSettingsResponse struct {
}

// UserSettings represents the UserSettings schema from the OpenAPI specification
type UserSettings struct {
	Printallowed interface{} `json:"printAllowed,omitempty"`
	Idledisconnecttimeoutinminutes interface{} `json:"idleDisconnectTimeoutInMinutes,omitempty"`
	Pasteallowed interface{} `json:"pasteAllowed,omitempty"`
	Uploadallowed interface{} `json:"uploadAllowed,omitempty"`
	Downloadallowed interface{} `json:"downloadAllowed,omitempty"`
	Usersettingsarn interface{} `json:"userSettingsArn"`
	Associatedportalarns interface{} `json:"associatedPortalArns,omitempty"`
	Copyallowed interface{} `json:"copyAllowed,omitempty"`
	Disconnecttimeoutinminutes interface{} `json:"disconnectTimeoutInMinutes,omitempty"`
}

// PortalSummary represents the PortalSummary schema from the OpenAPI specification
type PortalSummary struct {
	Ipaccesssettingsarn interface{} `json:"ipAccessSettingsArn,omitempty"`
	Networksettingsarn interface{} `json:"networkSettingsArn,omitempty"`
	Renderertype interface{} `json:"rendererType,omitempty"`
	Useraccessloggingsettingsarn interface{} `json:"userAccessLoggingSettingsArn,omitempty"`
	Usersettingsarn interface{} `json:"userSettingsArn,omitempty"`
	Authenticationtype interface{} `json:"authenticationType,omitempty"`
	Creationdate interface{} `json:"creationDate,omitempty"`
	Portalstatus interface{} `json:"portalStatus,omitempty"`
	Browsersettingsarn interface{} `json:"browserSettingsArn,omitempty"`
	Portalendpoint interface{} `json:"portalEndpoint,omitempty"`
	Truststorearn interface{} `json:"trustStoreArn,omitempty"`
	Browsertype interface{} `json:"browserType,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Portalarn interface{} `json:"portalArn,omitempty"`
}

// GetUserAccessLoggingSettingsRequest represents the GetUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type GetUserAccessLoggingSettingsRequest struct {
}

// CreateUserSettingsRequest represents the CreateUserSettingsRequest schema from the OpenAPI specification
type CreateUserSettingsRequest struct {
	Idledisconnecttimeoutinminutes interface{} `json:"idleDisconnectTimeoutInMinutes,omitempty"`
	Downloadallowed interface{} `json:"downloadAllowed"`
	Pasteallowed interface{} `json:"pasteAllowed"`
	Tags interface{} `json:"tags,omitempty"`
	Uploadallowed interface{} `json:"uploadAllowed"`
	Copyallowed interface{} `json:"copyAllowed"`
	Printallowed interface{} `json:"printAllowed"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Disconnecttimeoutinminutes interface{} `json:"disconnectTimeoutInMinutes,omitempty"`
}

// GetPortalServiceProviderMetadataRequest represents the GetPortalServiceProviderMetadataRequest schema from the OpenAPI specification
type GetPortalServiceProviderMetadataRequest struct {
}

// ListUserSettingsRequest represents the ListUserSettingsRequest schema from the OpenAPI specification
type ListUserSettingsRequest struct {
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// ListIpAccessSettingsRequest represents the ListIpAccessSettingsRequest schema from the OpenAPI specification
type ListIpAccessSettingsRequest struct {
}

// UpdateUserAccessLoggingSettingsRequest represents the UpdateUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type UpdateUserAccessLoggingSettingsRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Kinesisstreamarn interface{} `json:"kinesisStreamArn,omitempty"`
}

// CreateIdentityProviderRequest represents the CreateIdentityProviderRequest schema from the OpenAPI specification
type CreateIdentityProviderRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Identityproviderdetails interface{} `json:"identityProviderDetails"`
	Identityprovidername interface{} `json:"identityProviderName"`
	Identityprovidertype interface{} `json:"identityProviderType"`
	Portalarn interface{} `json:"portalArn"`
}

// UpdateUserSettingsRequest represents the UpdateUserSettingsRequest schema from the OpenAPI specification
type UpdateUserSettingsRequest struct {
	Downloadallowed interface{} `json:"downloadAllowed,omitempty"`
	Idledisconnecttimeoutinminutes interface{} `json:"idleDisconnectTimeoutInMinutes,omitempty"`
	Pasteallowed interface{} `json:"pasteAllowed,omitempty"`
	Printallowed interface{} `json:"printAllowed,omitempty"`
	Uploadallowed interface{} `json:"uploadAllowed,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Copyallowed interface{} `json:"copyAllowed,omitempty"`
	Disconnecttimeoutinminutes interface{} `json:"disconnectTimeoutInMinutes,omitempty"`
}

// GetIdentityProviderResponse represents the GetIdentityProviderResponse schema from the OpenAPI specification
type GetIdentityProviderResponse struct {
	Identityprovider interface{} `json:"identityProvider,omitempty"`
}

// GetUserAccessLoggingSettingsResponse represents the GetUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type GetUserAccessLoggingSettingsResponse struct {
	Useraccessloggingsettings interface{} `json:"userAccessLoggingSettings,omitempty"`
}

// CreatePortalResponse represents the CreatePortalResponse schema from the OpenAPI specification
type CreatePortalResponse struct {
	Portalendpoint interface{} `json:"portalEndpoint"`
	Portalarn interface{} `json:"portalArn"`
}

// CreateNetworkSettingsRequest represents the CreateNetworkSettingsRequest schema from the OpenAPI specification
type CreateNetworkSettingsRequest struct {
	Vpcid interface{} `json:"vpcId"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds"`
	Subnetids interface{} `json:"subnetIds"`
	Tags interface{} `json:"tags,omitempty"`
}

// DeleteBrowserSettingsResponse represents the DeleteBrowserSettingsResponse schema from the OpenAPI specification
type DeleteBrowserSettingsResponse struct {
}

// DisassociateTrustStoreResponse represents the DisassociateTrustStoreResponse schema from the OpenAPI specification
type DisassociateTrustStoreResponse struct {
}

// AssociateIpAccessSettingsResponse represents the AssociateIpAccessSettingsResponse schema from the OpenAPI specification
type AssociateIpAccessSettingsResponse struct {
	Ipaccesssettingsarn interface{} `json:"ipAccessSettingsArn"`
	Portalarn interface{} `json:"portalArn"`
}

// DeleteNetworkSettingsResponse represents the DeleteNetworkSettingsResponse schema from the OpenAPI specification
type DeleteNetworkSettingsResponse struct {
}

// DeletePortalRequest represents the DeletePortalRequest schema from the OpenAPI specification
type DeletePortalRequest struct {
}

// UserSettingsSummary represents the UserSettingsSummary schema from the OpenAPI specification
type UserSettingsSummary struct {
	Uploadallowed interface{} `json:"uploadAllowed,omitempty"`
	Usersettingsarn interface{} `json:"userSettingsArn,omitempty"`
	Copyallowed interface{} `json:"copyAllowed,omitempty"`
	Disconnecttimeoutinminutes interface{} `json:"disconnectTimeoutInMinutes,omitempty"`
	Downloadallowed interface{} `json:"downloadAllowed,omitempty"`
	Idledisconnecttimeoutinminutes interface{} `json:"idleDisconnectTimeoutInMinutes,omitempty"`
	Pasteallowed interface{} `json:"pasteAllowed,omitempty"`
	Printallowed interface{} `json:"printAllowed,omitempty"`
}

// CreateIpAccessSettingsRequest represents the CreateIpAccessSettingsRequest schema from the OpenAPI specification
type CreateIpAccessSettingsRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Customermanagedkey interface{} `json:"customerManagedKey,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Iprules interface{} `json:"ipRules"`
	Tags interface{} `json:"tags,omitempty"`
	Additionalencryptioncontext interface{} `json:"additionalEncryptionContext,omitempty"`
}

// DisassociateUserAccessLoggingSettingsResponse represents the DisassociateUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type DisassociateUserAccessLoggingSettingsResponse struct {
}

// AssociateBrowserSettingsResponse represents the AssociateBrowserSettingsResponse schema from the OpenAPI specification
type AssociateBrowserSettingsResponse struct {
	Browsersettingsarn interface{} `json:"browserSettingsArn"`
	Portalarn interface{} `json:"portalArn"`
}

// ListIdentityProvidersResponse represents the ListIdentityProvidersResponse schema from the OpenAPI specification
type ListIdentityProvidersResponse struct {
	Identityproviders interface{} `json:"identityProviders,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteIdentityProviderResponse represents the DeleteIdentityProviderResponse schema from the OpenAPI specification
type DeleteIdentityProviderResponse struct {
}

// ListNetworkSettingsResponse represents the ListNetworkSettingsResponse schema from the OpenAPI specification
type ListNetworkSettingsResponse struct {
	Networksettings interface{} `json:"networkSettings,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateUserAccessLoggingSettingsRequest represents the CreateUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type CreateUserAccessLoggingSettingsRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Kinesisstreamarn interface{} `json:"kinesisStreamArn"`
}

// ListUserSettingsResponse represents the ListUserSettingsResponse schema from the OpenAPI specification
type ListUserSettingsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Usersettings interface{} `json:"userSettings,omitempty"`
}

// CertificateSummary represents the CertificateSummary schema from the OpenAPI specification
type CertificateSummary struct {
	Notvalidbefore interface{} `json:"notValidBefore,omitempty"`
	Subject interface{} `json:"subject,omitempty"`
	Thumbprint interface{} `json:"thumbprint,omitempty"`
	Issuer interface{} `json:"issuer,omitempty"`
	Notvalidafter interface{} `json:"notValidAfter,omitempty"`
}

// CreateTrustStoreResponse represents the CreateTrustStoreResponse schema from the OpenAPI specification
type CreateTrustStoreResponse struct {
	Truststorearn interface{} `json:"trustStoreArn"`
}

// Certificate represents the Certificate schema from the OpenAPI specification
type Certificate struct {
	Notvalidbefore interface{} `json:"notValidBefore,omitempty"`
	Subject interface{} `json:"subject,omitempty"`
	Thumbprint interface{} `json:"thumbprint,omitempty"`
	Body interface{} `json:"body,omitempty"`
	Issuer interface{} `json:"issuer,omitempty"`
	Notvalidafter interface{} `json:"notValidAfter,omitempty"`
}

// ListNetworkSettingsRequest represents the ListNetworkSettingsRequest schema from the OpenAPI specification
type ListNetworkSettingsRequest struct {
}

// ListTrustStoresResponse represents the ListTrustStoresResponse schema from the OpenAPI specification
type ListTrustStoresResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Truststores interface{} `json:"trustStores,omitempty"`
}

// ListTrustStoreCertificatesRequest represents the ListTrustStoreCertificatesRequest schema from the OpenAPI specification
type ListTrustStoreCertificatesRequest struct {
}

// ListPortalsRequest represents the ListPortalsRequest schema from the OpenAPI specification
type ListPortalsRequest struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Tags interface{} `json:"tags"`
}

// GetTrustStoreRequest represents the GetTrustStoreRequest schema from the OpenAPI specification
type GetTrustStoreRequest struct {
}

// DeletePortalResponse represents the DeletePortalResponse schema from the OpenAPI specification
type DeletePortalResponse struct {
}

// IdentityProviderDetails represents the IdentityProviderDetails schema from the OpenAPI specification
type IdentityProviderDetails struct {
}

// EncryptionContextMap represents the EncryptionContextMap schema from the OpenAPI specification
type EncryptionContextMap struct {
}

// GetNetworkSettingsResponse represents the GetNetworkSettingsResponse schema from the OpenAPI specification
type GetNetworkSettingsResponse struct {
	Networksettings interface{} `json:"networkSettings,omitempty"`
}

// CreateIpAccessSettingsResponse represents the CreateIpAccessSettingsResponse schema from the OpenAPI specification
type CreateIpAccessSettingsResponse struct {
	Ipaccesssettingsarn interface{} `json:"ipAccessSettingsArn"`
}

// DeleteTrustStoreRequest represents the DeleteTrustStoreRequest schema from the OpenAPI specification
type DeleteTrustStoreRequest struct {
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// AssociateUserAccessLoggingSettingsRequest represents the AssociateUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type AssociateUserAccessLoggingSettingsRequest struct {
}

// GetIpAccessSettingsResponse represents the GetIpAccessSettingsResponse schema from the OpenAPI specification
type GetIpAccessSettingsResponse struct {
	Ipaccesssettings interface{} `json:"ipAccessSettings,omitempty"`
}

// UpdateUserSettingsResponse represents the UpdateUserSettingsResponse schema from the OpenAPI specification
type UpdateUserSettingsResponse struct {
	Usersettings interface{} `json:"userSettings"`
}

// UpdateUserAccessLoggingSettingsResponse represents the UpdateUserAccessLoggingSettingsResponse schema from the OpenAPI specification
type UpdateUserAccessLoggingSettingsResponse struct {
	Useraccessloggingsettings interface{} `json:"userAccessLoggingSettings"`
}

// CreateUserSettingsResponse represents the CreateUserSettingsResponse schema from the OpenAPI specification
type CreateUserSettingsResponse struct {
	Usersettingsarn interface{} `json:"userSettingsArn"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// AssociateIpAccessSettingsRequest represents the AssociateIpAccessSettingsRequest schema from the OpenAPI specification
type AssociateIpAccessSettingsRequest struct {
}

// DeleteBrowserSettingsRequest represents the DeleteBrowserSettingsRequest schema from the OpenAPI specification
type DeleteBrowserSettingsRequest struct {
}

// GetPortalResponse represents the GetPortalResponse schema from the OpenAPI specification
type GetPortalResponse struct {
	Portal interface{} `json:"portal,omitempty"`
}

// DeleteUserAccessLoggingSettingsRequest represents the DeleteUserAccessLoggingSettingsRequest schema from the OpenAPI specification
type DeleteUserAccessLoggingSettingsRequest struct {
}

// DeleteNetworkSettingsRequest represents the DeleteNetworkSettingsRequest schema from the OpenAPI specification
type DeleteNetworkSettingsRequest struct {
}

// DisassociateTrustStoreRequest represents the DisassociateTrustStoreRequest schema from the OpenAPI specification
type DisassociateTrustStoreRequest struct {
}

// AssociateUserSettingsRequest represents the AssociateUserSettingsRequest schema from the OpenAPI specification
type AssociateUserSettingsRequest struct {
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// AssociateBrowserSettingsRequest represents the AssociateBrowserSettingsRequest schema from the OpenAPI specification
type AssociateBrowserSettingsRequest struct {
}

// DisassociateUserSettingsResponse represents the DisassociateUserSettingsResponse schema from the OpenAPI specification
type DisassociateUserSettingsResponse struct {
}

// GetBrowserSettingsRequest represents the GetBrowserSettingsRequest schema from the OpenAPI specification
type GetBrowserSettingsRequest struct {
}

// BrowserSettings represents the BrowserSettings schema from the OpenAPI specification
type BrowserSettings struct {
	Browserpolicy interface{} `json:"browserPolicy,omitempty"`
	Browsersettingsarn interface{} `json:"browserSettingsArn"`
	Associatedportalarns interface{} `json:"associatedPortalArns,omitempty"`
}

// GetNetworkSettingsRequest represents the GetNetworkSettingsRequest schema from the OpenAPI specification
type GetNetworkSettingsRequest struct {
}
