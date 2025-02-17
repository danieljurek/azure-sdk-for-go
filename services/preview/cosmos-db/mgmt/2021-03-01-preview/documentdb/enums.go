package documentdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// APIType enumerates the values for api type.
type APIType string

const (
	// APITypeCassandra ...
	APITypeCassandra APIType = "Cassandra"
	// APITypeGremlin ...
	APITypeGremlin APIType = "Gremlin"
	// APITypeGremlinV2 ...
	APITypeGremlinV2 APIType = "GremlinV2"
	// APITypeMongoDB ...
	APITypeMongoDB APIType = "MongoDB"
	// APITypeSQL ...
	APITypeSQL APIType = "Sql"
	// APITypeTable ...
	APITypeTable APIType = "Table"
)

// PossibleAPITypeValues returns an array of possible values for the APIType const type.
func PossibleAPITypeValues() []APIType {
	return []APIType{APITypeCassandra, APITypeGremlin, APITypeGremlinV2, APITypeMongoDB, APITypeSQL, APITypeTable}
}

// AuthenticationMethod enumerates the values for authentication method.
type AuthenticationMethod string

const (
	// AuthenticationMethodCassandra ...
	AuthenticationMethodCassandra AuthenticationMethod = "Cassandra"
	// AuthenticationMethodNone ...
	AuthenticationMethodNone AuthenticationMethod = "None"
)

// PossibleAuthenticationMethodValues returns an array of possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{AuthenticationMethodCassandra, AuthenticationMethodNone}
}

// BackupPolicyType enumerates the values for backup policy type.
type BackupPolicyType string

const (
	// BackupPolicyTypeContinuous ...
	BackupPolicyTypeContinuous BackupPolicyType = "Continuous"
	// BackupPolicyTypePeriodic ...
	BackupPolicyTypePeriodic BackupPolicyType = "Periodic"
)

// PossibleBackupPolicyTypeValues returns an array of possible values for the BackupPolicyType const type.
func PossibleBackupPolicyTypeValues() []BackupPolicyType {
	return []BackupPolicyType{BackupPolicyTypeContinuous, BackupPolicyTypePeriodic}
}

// BackupStorageRedundancy enumerates the values for backup storage redundancy.
type BackupStorageRedundancy string

const (
	// BackupStorageRedundancyGeo ...
	BackupStorageRedundancyGeo BackupStorageRedundancy = "Geo"
	// BackupStorageRedundancyLocal ...
	BackupStorageRedundancyLocal BackupStorageRedundancy = "Local"
	// BackupStorageRedundancyZone ...
	BackupStorageRedundancyZone BackupStorageRedundancy = "Zone"
)

// PossibleBackupStorageRedundancyValues returns an array of possible values for the BackupStorageRedundancy const type.
func PossibleBackupStorageRedundancyValues() []BackupStorageRedundancy {
	return []BackupStorageRedundancy{BackupStorageRedundancyGeo, BackupStorageRedundancyLocal, BackupStorageRedundancyZone}
}

// CompositePathSortOrder enumerates the values for composite path sort order.
type CompositePathSortOrder string

const (
	// CompositePathSortOrderAscending ...
	CompositePathSortOrderAscending CompositePathSortOrder = "ascending"
	// CompositePathSortOrderDescending ...
	CompositePathSortOrderDescending CompositePathSortOrder = "descending"
)

// PossibleCompositePathSortOrderValues returns an array of possible values for the CompositePathSortOrder const type.
func PossibleCompositePathSortOrderValues() []CompositePathSortOrder {
	return []CompositePathSortOrder{CompositePathSortOrderAscending, CompositePathSortOrderDescending}
}

// ConflictResolutionMode enumerates the values for conflict resolution mode.
type ConflictResolutionMode string

const (
	// ConflictResolutionModeCustom ...
	ConflictResolutionModeCustom ConflictResolutionMode = "Custom"
	// ConflictResolutionModeLastWriterWins ...
	ConflictResolutionModeLastWriterWins ConflictResolutionMode = "LastWriterWins"
)

// PossibleConflictResolutionModeValues returns an array of possible values for the ConflictResolutionMode const type.
func PossibleConflictResolutionModeValues() []ConflictResolutionMode {
	return []ConflictResolutionMode{ConflictResolutionModeCustom, ConflictResolutionModeLastWriterWins}
}

// ConnectorOffer enumerates the values for connector offer.
type ConnectorOffer string

const (
	// ConnectorOfferSmall ...
	ConnectorOfferSmall ConnectorOffer = "Small"
)

// PossibleConnectorOfferValues returns an array of possible values for the ConnectorOffer const type.
func PossibleConnectorOfferValues() []ConnectorOffer {
	return []ConnectorOffer{ConnectorOfferSmall}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// CreateModeDefault ...
	CreateModeDefault CreateMode = "Default"
	// CreateModeRestore ...
	CreateModeRestore CreateMode = "Restore"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{CreateModeDefault, CreateModeRestore}
}

// CreateModeBasicDatabaseAccountCreateUpdateProperties enumerates the values for create mode basic database
// account create update properties.
type CreateModeBasicDatabaseAccountCreateUpdateProperties string

const (
	// CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeDatabaseAccountCreateUpdateProperties ...
	CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeDatabaseAccountCreateUpdateProperties CreateModeBasicDatabaseAccountCreateUpdateProperties = "DatabaseAccountCreateUpdateProperties"
	// CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeDefault ...
	CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeDefault CreateModeBasicDatabaseAccountCreateUpdateProperties = "Default"
	// CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeRestore ...
	CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeRestore CreateModeBasicDatabaseAccountCreateUpdateProperties = "Restore"
)

// PossibleCreateModeBasicDatabaseAccountCreateUpdatePropertiesValues returns an array of possible values for the CreateModeBasicDatabaseAccountCreateUpdateProperties const type.
func PossibleCreateModeBasicDatabaseAccountCreateUpdatePropertiesValues() []CreateModeBasicDatabaseAccountCreateUpdateProperties {
	return []CreateModeBasicDatabaseAccountCreateUpdateProperties{CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeDatabaseAccountCreateUpdateProperties, CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeDefault, CreateModeBasicDatabaseAccountCreateUpdatePropertiesCreateModeRestore}
}

// DatabaseAccountKind enumerates the values for database account kind.
type DatabaseAccountKind string

const (
	// DatabaseAccountKindGlobalDocumentDB ...
	DatabaseAccountKindGlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
	// DatabaseAccountKindMongoDB ...
	DatabaseAccountKindMongoDB DatabaseAccountKind = "MongoDB"
	// DatabaseAccountKindParse ...
	DatabaseAccountKindParse DatabaseAccountKind = "Parse"
)

// PossibleDatabaseAccountKindValues returns an array of possible values for the DatabaseAccountKind const type.
func PossibleDatabaseAccountKindValues() []DatabaseAccountKind {
	return []DatabaseAccountKind{DatabaseAccountKindGlobalDocumentDB, DatabaseAccountKindMongoDB, DatabaseAccountKindParse}
}

// DatabaseAccountOfferType enumerates the values for database account offer type.
type DatabaseAccountOfferType string

const (
	// DatabaseAccountOfferTypeStandard ...
	DatabaseAccountOfferTypeStandard DatabaseAccountOfferType = "Standard"
)

// PossibleDatabaseAccountOfferTypeValues returns an array of possible values for the DatabaseAccountOfferType const type.
func PossibleDatabaseAccountOfferTypeValues() []DatabaseAccountOfferType {
	return []DatabaseAccountOfferType{DatabaseAccountOfferTypeStandard}
}

// DataType enumerates the values for data type.
type DataType string

const (
	// DataTypeLineString ...
	DataTypeLineString DataType = "LineString"
	// DataTypeMultiPolygon ...
	DataTypeMultiPolygon DataType = "MultiPolygon"
	// DataTypeNumber ...
	DataTypeNumber DataType = "Number"
	// DataTypePoint ...
	DataTypePoint DataType = "Point"
	// DataTypePolygon ...
	DataTypePolygon DataType = "Polygon"
	// DataTypeString ...
	DataTypeString DataType = "String"
)

// PossibleDataTypeValues returns an array of possible values for the DataType const type.
func PossibleDataTypeValues() []DataType {
	return []DataType{DataTypeLineString, DataTypeMultiPolygon, DataTypeNumber, DataTypePoint, DataTypePolygon, DataTypeString}
}

// DefaultConsistencyLevel enumerates the values for default consistency level.
type DefaultConsistencyLevel string

const (
	// DefaultConsistencyLevelBoundedStaleness ...
	DefaultConsistencyLevelBoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
	// DefaultConsistencyLevelConsistentPrefix ...
	DefaultConsistencyLevelConsistentPrefix DefaultConsistencyLevel = "ConsistentPrefix"
	// DefaultConsistencyLevelEventual ...
	DefaultConsistencyLevelEventual DefaultConsistencyLevel = "Eventual"
	// DefaultConsistencyLevelSession ...
	DefaultConsistencyLevelSession DefaultConsistencyLevel = "Session"
	// DefaultConsistencyLevelStrong ...
	DefaultConsistencyLevelStrong DefaultConsistencyLevel = "Strong"
)

// PossibleDefaultConsistencyLevelValues returns an array of possible values for the DefaultConsistencyLevel const type.
func PossibleDefaultConsistencyLevelValues() []DefaultConsistencyLevel {
	return []DefaultConsistencyLevel{DefaultConsistencyLevelBoundedStaleness, DefaultConsistencyLevelConsistentPrefix, DefaultConsistencyLevelEventual, DefaultConsistencyLevelSession, DefaultConsistencyLevelStrong}
}

// IndexingMode enumerates the values for indexing mode.
type IndexingMode string

const (
	// IndexingModeConsistent ...
	IndexingModeConsistent IndexingMode = "consistent"
	// IndexingModeLazy ...
	IndexingModeLazy IndexingMode = "lazy"
	// IndexingModeNone ...
	IndexingModeNone IndexingMode = "none"
)

// PossibleIndexingModeValues returns an array of possible values for the IndexingMode const type.
func PossibleIndexingModeValues() []IndexingMode {
	return []IndexingMode{IndexingModeConsistent, IndexingModeLazy, IndexingModeNone}
}

// IndexKind enumerates the values for index kind.
type IndexKind string

const (
	// IndexKindHash ...
	IndexKindHash IndexKind = "Hash"
	// IndexKindRange ...
	IndexKindRange IndexKind = "Range"
	// IndexKindSpatial ...
	IndexKindSpatial IndexKind = "Spatial"
)

// PossibleIndexKindValues returns an array of possible values for the IndexKind const type.
func PossibleIndexKindValues() []IndexKind {
	return []IndexKind{IndexKindHash, IndexKindRange, IndexKindSpatial}
}

// KeyKind enumerates the values for key kind.
type KeyKind string

const (
	// KeyKindPrimary ...
	KeyKindPrimary KeyKind = "primary"
	// KeyKindPrimaryReadonly ...
	KeyKindPrimaryReadonly KeyKind = "primaryReadonly"
	// KeyKindSecondary ...
	KeyKindSecondary KeyKind = "secondary"
	// KeyKindSecondaryReadonly ...
	KeyKindSecondaryReadonly KeyKind = "secondaryReadonly"
)

// PossibleKeyKindValues returns an array of possible values for the KeyKind const type.
func PossibleKeyKindValues() []KeyKind {
	return []KeyKind{KeyKindPrimary, KeyKindPrimaryReadonly, KeyKindSecondary, KeyKindSecondaryReadonly}
}

// ManagedCassandraProvisioningState enumerates the values for managed cassandra provisioning state.
type ManagedCassandraProvisioningState string

const (
	// ManagedCassandraProvisioningStateCanceled ...
	ManagedCassandraProvisioningStateCanceled ManagedCassandraProvisioningState = "Canceled"
	// ManagedCassandraProvisioningStateCreating ...
	ManagedCassandraProvisioningStateCreating ManagedCassandraProvisioningState = "Creating"
	// ManagedCassandraProvisioningStateDeleting ...
	ManagedCassandraProvisioningStateDeleting ManagedCassandraProvisioningState = "Deleting"
	// ManagedCassandraProvisioningStateFailed ...
	ManagedCassandraProvisioningStateFailed ManagedCassandraProvisioningState = "Failed"
	// ManagedCassandraProvisioningStateSucceeded ...
	ManagedCassandraProvisioningStateSucceeded ManagedCassandraProvisioningState = "Succeeded"
	// ManagedCassandraProvisioningStateUpdating ...
	ManagedCassandraProvisioningStateUpdating ManagedCassandraProvisioningState = "Updating"
)

// PossibleManagedCassandraProvisioningStateValues returns an array of possible values for the ManagedCassandraProvisioningState const type.
func PossibleManagedCassandraProvisioningStateValues() []ManagedCassandraProvisioningState {
	return []ManagedCassandraProvisioningState{ManagedCassandraProvisioningStateCanceled, ManagedCassandraProvisioningStateCreating, ManagedCassandraProvisioningStateDeleting, ManagedCassandraProvisioningStateFailed, ManagedCassandraProvisioningStateSucceeded, ManagedCassandraProvisioningStateUpdating}
}

// NetworkACLBypass enumerates the values for network acl bypass.
type NetworkACLBypass string

const (
	// NetworkACLBypassAzureServices ...
	NetworkACLBypassAzureServices NetworkACLBypass = "AzureServices"
	// NetworkACLBypassNone ...
	NetworkACLBypassNone NetworkACLBypass = "None"
)

// PossibleNetworkACLBypassValues returns an array of possible values for the NetworkACLBypass const type.
func PossibleNetworkACLBypassValues() []NetworkACLBypass {
	return []NetworkACLBypass{NetworkACLBypassAzureServices, NetworkACLBypassNone}
}

// NodeState enumerates the values for node state.
type NodeState string

const (
	// NodeStateJoining ...
	NodeStateJoining NodeState = "Joining"
	// NodeStateLeaving ...
	NodeStateLeaving NodeState = "Leaving"
	// NodeStateMoving ...
	NodeStateMoving NodeState = "Moving"
	// NodeStateNormal ...
	NodeStateNormal NodeState = "Normal"
	// NodeStateStopped ...
	NodeStateStopped NodeState = "Stopped"
)

// PossibleNodeStateValues returns an array of possible values for the NodeState const type.
func PossibleNodeStateValues() []NodeState {
	return []NodeState{NodeStateJoining, NodeStateLeaving, NodeStateMoving, NodeStateNormal, NodeStateStopped}
}

// NodeStatus enumerates the values for node status.
type NodeStatus string

const (
	// NodeStatusDown ...
	NodeStatusDown NodeStatus = "Down"
	// NodeStatusUp ...
	NodeStatusUp NodeStatus = "Up"
)

// PossibleNodeStatusValues returns an array of possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{NodeStatusDown, NodeStatusUp}
}

// OperationType enumerates the values for operation type.
type OperationType string

const (
	// OperationTypeCreate ...
	OperationTypeCreate OperationType = "Create"
	// OperationTypeDelete ...
	OperationTypeDelete OperationType = "Delete"
	// OperationTypeReplace ...
	OperationTypeReplace OperationType = "Replace"
	// OperationTypeSystemOperation ...
	OperationTypeSystemOperation OperationType = "SystemOperation"
)

// PossibleOperationTypeValues returns an array of possible values for the OperationType const type.
func PossibleOperationTypeValues() []OperationType {
	return []OperationType{OperationTypeCreate, OperationTypeDelete, OperationTypeReplace, OperationTypeSystemOperation}
}

// PartitionKind enumerates the values for partition kind.
type PartitionKind string

const (
	// PartitionKindHash ...
	PartitionKindHash PartitionKind = "Hash"
	// PartitionKindMultiHash ...
	PartitionKindMultiHash PartitionKind = "MultiHash"
	// PartitionKindRange ...
	PartitionKindRange PartitionKind = "Range"
)

// PossiblePartitionKindValues returns an array of possible values for the PartitionKind const type.
func PossiblePartitionKindValues() []PartitionKind {
	return []PartitionKind{PartitionKindHash, PartitionKindMultiHash, PartitionKindRange}
}

// PrimaryAggregationType enumerates the values for primary aggregation type.
type PrimaryAggregationType string

const (
	// PrimaryAggregationTypeAverage ...
	PrimaryAggregationTypeAverage PrimaryAggregationType = "Average"
	// PrimaryAggregationTypeLast ...
	PrimaryAggregationTypeLast PrimaryAggregationType = "Last"
	// PrimaryAggregationTypeMaximum ...
	PrimaryAggregationTypeMaximum PrimaryAggregationType = "Maximum"
	// PrimaryAggregationTypeMinimum ...
	PrimaryAggregationTypeMinimum PrimaryAggregationType = "Minimum"
	// PrimaryAggregationTypeNone ...
	PrimaryAggregationTypeNone PrimaryAggregationType = "None"
	// PrimaryAggregationTypeTotal ...
	PrimaryAggregationTypeTotal PrimaryAggregationType = "Total"
)

// PossiblePrimaryAggregationTypeValues returns an array of possible values for the PrimaryAggregationType const type.
func PossiblePrimaryAggregationTypeValues() []PrimaryAggregationType {
	return []PrimaryAggregationType{PrimaryAggregationTypeAverage, PrimaryAggregationTypeLast, PrimaryAggregationTypeMaximum, PrimaryAggregationTypeMinimum, PrimaryAggregationTypeNone, PrimaryAggregationTypeTotal}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled ...
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	// PublicNetworkAccessEnabled ...
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{PublicNetworkAccessDisabled, PublicNetworkAccessEnabled}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// RestoreMode enumerates the values for restore mode.
type RestoreMode string

const (
	// RestoreModePointInTime ...
	RestoreModePointInTime RestoreMode = "PointInTime"
)

// PossibleRestoreModeValues returns an array of possible values for the RestoreMode const type.
func PossibleRestoreModeValues() []RestoreMode {
	return []RestoreMode{RestoreModePointInTime}
}

// RoleDefinitionType enumerates the values for role definition type.
type RoleDefinitionType string

const (
	// RoleDefinitionTypeBuiltInRole ...
	RoleDefinitionTypeBuiltInRole RoleDefinitionType = "BuiltInRole"
	// RoleDefinitionTypeCustomRole ...
	RoleDefinitionTypeCustomRole RoleDefinitionType = "CustomRole"
)

// PossibleRoleDefinitionTypeValues returns an array of possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{RoleDefinitionTypeBuiltInRole, RoleDefinitionTypeCustomRole}
}

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// ServerVersionFourFullStopZero ...
	ServerVersionFourFullStopZero ServerVersion = "4.0"
	// ServerVersionThreeFullStopSix ...
	ServerVersionThreeFullStopSix ServerVersion = "3.6"
	// ServerVersionThreeFullStopTwo ...
	ServerVersionThreeFullStopTwo ServerVersion = "3.2"
)

// PossibleServerVersionValues returns an array of possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{ServerVersionFourFullStopZero, ServerVersionThreeFullStopSix, ServerVersionThreeFullStopTwo}
}

// SpatialType enumerates the values for spatial type.
type SpatialType string

const (
	// SpatialTypeLineString ...
	SpatialTypeLineString SpatialType = "LineString"
	// SpatialTypeMultiPolygon ...
	SpatialTypeMultiPolygon SpatialType = "MultiPolygon"
	// SpatialTypePoint ...
	SpatialTypePoint SpatialType = "Point"
	// SpatialTypePolygon ...
	SpatialTypePolygon SpatialType = "Polygon"
)

// PossibleSpatialTypeValues returns an array of possible values for the SpatialType const type.
func PossibleSpatialTypeValues() []SpatialType {
	return []SpatialType{SpatialTypeLineString, SpatialTypeMultiPolygon, SpatialTypePoint, SpatialTypePolygon}
}

// TriggerOperation enumerates the values for trigger operation.
type TriggerOperation string

const (
	// TriggerOperationAll ...
	TriggerOperationAll TriggerOperation = "All"
	// TriggerOperationCreate ...
	TriggerOperationCreate TriggerOperation = "Create"
	// TriggerOperationDelete ...
	TriggerOperationDelete TriggerOperation = "Delete"
	// TriggerOperationReplace ...
	TriggerOperationReplace TriggerOperation = "Replace"
	// TriggerOperationUpdate ...
	TriggerOperationUpdate TriggerOperation = "Update"
)

// PossibleTriggerOperationValues returns an array of possible values for the TriggerOperation const type.
func PossibleTriggerOperationValues() []TriggerOperation {
	return []TriggerOperation{TriggerOperationAll, TriggerOperationCreate, TriggerOperationDelete, TriggerOperationReplace, TriggerOperationUpdate}
}

// TriggerType enumerates the values for trigger type.
type TriggerType string

const (
	// TriggerTypePost ...
	TriggerTypePost TriggerType = "Post"
	// TriggerTypePre ...
	TriggerTypePre TriggerType = "Pre"
)

// PossibleTriggerTypeValues returns an array of possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{TriggerTypePost, TriggerTypePre}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeBackupPolicy ...
	TypeBackupPolicy Type = "BackupPolicy"
	// TypeContinuous ...
	TypeContinuous Type = "Continuous"
	// TypePeriodic ...
	TypePeriodic Type = "Periodic"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeBackupPolicy, TypeContinuous, TypePeriodic}
}

// UnitType enumerates the values for unit type.
type UnitType string

const (
	// UnitTypeBytes ...
	UnitTypeBytes UnitType = "Bytes"
	// UnitTypeBytesPerSecond ...
	UnitTypeBytesPerSecond UnitType = "BytesPerSecond"
	// UnitTypeCount ...
	UnitTypeCount UnitType = "Count"
	// UnitTypeCountPerSecond ...
	UnitTypeCountPerSecond UnitType = "CountPerSecond"
	// UnitTypeMilliseconds ...
	UnitTypeMilliseconds UnitType = "Milliseconds"
	// UnitTypePercent ...
	UnitTypePercent UnitType = "Percent"
	// UnitTypeSeconds ...
	UnitTypeSeconds UnitType = "Seconds"
)

// PossibleUnitTypeValues returns an array of possible values for the UnitType const type.
func PossibleUnitTypeValues() []UnitType {
	return []UnitType{UnitTypeBytes, UnitTypeBytesPerSecond, UnitTypeCount, UnitTypeCountPerSecond, UnitTypeMilliseconds, UnitTypePercent, UnitTypeSeconds}
}
