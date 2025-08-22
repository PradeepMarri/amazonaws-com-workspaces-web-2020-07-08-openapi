package main

import (
	"github.com/amazon-workspaces-web/mcp-server/config"
	"github.com/amazon-workspaces-web/mcp-server/models"
	tools_portalidp "github.com/amazon-workspaces-web/mcp-server/tools/portalidp"
	tools_portals "github.com/amazon-workspaces-web/mcp-server/tools/portals"
	tools_identityproviders "github.com/amazon-workspaces-web/mcp-server/tools/identityproviders"
	tools_useraccessloggingsettings "github.com/amazon-workspaces-web/mcp-server/tools/useraccessloggingsettings"
	tools_browsersettings "github.com/amazon-workspaces-web/mcp-server/tools/browsersettings"
	tools_tags "github.com/amazon-workspaces-web/mcp-server/tools/tags"
	tools_truststores "github.com/amazon-workspaces-web/mcp-server/tools/truststores"
	tools_ipaccesssettings "github.com/amazon-workspaces-web/mcp-server/tools/ipaccesssettings"
	tools_networksettings "github.com/amazon-workspaces-web/mcp-server/tools/networksettings"
	tools_usersettings "github.com/amazon-workspaces-web/mcp-server/tools/usersettings"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_portalidp.CreateGetportalserviceprovidermetadataTool(cfg),
		tools_portals.CreateListportalsTool(cfg),
		tools_portals.CreateCreateportalTool(cfg),
		tools_portals.CreateAssociateuseraccessloggingsettingsTool(cfg),
		tools_portals.CreateDisassociateipaccesssettingsTool(cfg),
		tools_portals.CreateDisassociatenetworksettingsTool(cfg),
		tools_identityproviders.CreateUpdateidentityproviderTool(cfg),
		tools_identityproviders.CreateDeleteidentityproviderTool(cfg),
		tools_identityproviders.CreateGetidentityproviderTool(cfg),
		tools_useraccessloggingsettings.CreateListuseraccessloggingsettingsTool(cfg),
		tools_useraccessloggingsettings.CreateCreateuseraccessloggingsettingsTool(cfg),
		tools_portals.CreateAssociateusersettingsTool(cfg),
		tools_browsersettings.CreateCreatebrowsersettingsTool(cfg),
		tools_browsersettings.CreateListbrowsersettingsTool(cfg),
		tools_portals.CreateAssociatebrowsersettingsTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_truststores.CreateListtruststorecertificatesTool(cfg),
		tools_portals.CreateDisassociateuseraccessloggingsettingsTool(cfg),
		tools_ipaccesssettings.CreateDeleteipaccesssettingsTool(cfg),
		tools_ipaccesssettings.CreateGetipaccesssettingsTool(cfg),
		tools_ipaccesssettings.CreateUpdateipaccesssettingsTool(cfg),
		tools_ipaccesssettings.CreateListipaccesssettingsTool(cfg),
		tools_ipaccesssettings.CreateCreateipaccesssettingsTool(cfg),
		tools_portals.CreateAssociatetruststoreTool(cfg),
		tools_networksettings.CreateListnetworksettingsTool(cfg),
		tools_networksettings.CreateCreatenetworksettingsTool(cfg),
		tools_portals.CreateAssociateipaccesssettingsTool(cfg),
		tools_truststores.CreateDeletetruststoreTool(cfg),
		tools_truststores.CreateGettruststoreTool(cfg),
		tools_truststores.CreateUpdatetruststoreTool(cfg),
		tools_portals.CreateDisassociateusersettingsTool(cfg),
		tools_portals.CreateDeleteportalTool(cfg),
		tools_portals.CreateGetportalTool(cfg),
		tools_portals.CreateUpdateportalTool(cfg),
		tools_usersettings.CreateListusersettingsTool(cfg),
		tools_usersettings.CreateCreateusersettingsTool(cfg),
		tools_useraccessloggingsettings.CreateDeleteuseraccessloggingsettingsTool(cfg),
		tools_useraccessloggingsettings.CreateGetuseraccessloggingsettingsTool(cfg),
		tools_useraccessloggingsettings.CreateUpdateuseraccessloggingsettingsTool(cfg),
		tools_usersettings.CreateUpdateusersettingsTool(cfg),
		tools_usersettings.CreateDeleteusersettingsTool(cfg),
		tools_usersettings.CreateGetusersettingsTool(cfg),
		tools_networksettings.CreateUpdatenetworksettingsTool(cfg),
		tools_networksettings.CreateDeletenetworksettingsTool(cfg),
		tools_networksettings.CreateGetnetworksettingsTool(cfg),
		tools_portals.CreateDisassociatetruststoreTool(cfg),
		tools_portals.CreateAssociatenetworksettingsTool(cfg),
		tools_identityproviders.CreateCreateidentityproviderTool(cfg),
		tools_browsersettings.CreateDeletebrowsersettingsTool(cfg),
		tools_browsersettings.CreateGetbrowsersettingsTool(cfg),
		tools_browsersettings.CreateUpdatebrowsersettingsTool(cfg),
		tools_portals.CreateListidentityprovidersTool(cfg),
		tools_truststores.CreateListtruststoresTool(cfg),
		tools_truststores.CreateCreatetruststoreTool(cfg),
		tools_portals.CreateDisassociatebrowsersettingsTool(cfg),
		tools_truststores.CreateGettruststorecertificateTool(cfg),
	}
}
