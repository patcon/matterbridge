// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementSettingCategory Entity representing a setting category
type DeviceManagementSettingCategory struct {
	// Entity is the base model of DeviceManagementSettingCategory
	Entity
	// DisplayName The category name
	DisplayName *string `json:"displayName,omitempty"`
	// SettingDefinitions undocumented
	SettingDefinitions []DeviceManagementSettingDefinition `json:"settingDefinitions,omitempty"`
}