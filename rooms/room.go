package rooms

import (
	"github.com/hosts-guru/OTAModels-golang"
	"github.com/hosts-guru/OTAModels-golang/common"
	"github.com/hosts-guru/OTAModels-golang/organization"
)

// RoomTypeID Assigned Type: ph-0200:RoomType
type RoomTypeID struct {
	Type  string  `json:"@type"`
	ObjID *string `json:"objID,omitempty"` // A unique identifier within this document for this object.
}

// RoomType Assigned Type: ph-0200:RoomType
type RoomType struct {
	RoomTypeID
	Amenity             []common.Amenity           `json:"Amenity,omitempty"`
	Code                *common.Code               `json:"Code"`        // Assigned Type: ota2-0500:Code
	Description         *common.Description        `json:"Description"` // Assigned Type: ota2-0500:Description
	ExtensionPoint      interface{}                `json:"ExtensionPoint"`
	RoomAmenities       []common.Code              `json:"RoomAmenities,omitempty"`
	RoomCharacteristics *RoomCharacteristicsDetail `json:"RoomCharacteristics"` // Assigned Type: ph-0200:RoomCharacteristics
	RoomOccupancy       []RoomOccupancy            `json:"RoomOccupancy,omitempty"`
	Tier                *string                    `json:"Tier,omitempty"` // TODO:  Create enumerated list.  Indicates the category of the room. Typical values would; be Moderate, Standard, or Deluxe. Refer to OpenTravel Code List Segment Category Code; (SEG).
}

type RoomOccupancy struct {
	Type              string                    `json:"@type"`
	AgeQualifying     []otamodel.EnumStringBase `json:"AgeQualifying,omitempty"`
	ExtensionPoint    interface{}               `json:"ExtensionPoint"`
	MaxOccupancy      *int64                    `json:"maxOccupancy,omitempty"`      // Assigned Type: ota2-0500:NonNegativeInteger
	MinOccupancy      *int64                    `json:"minOccupancy,omitempty"`      // Assigned Type: ota2-0500:NonNegativeInteger
	StandardOccupancy *int64                    `json:"standardOccupancy,omitempty"` // Assigned Type: ota2-0500:NonNegativeInteger
}

// RoomCharacteristics Assigned Type: ph-0200:RoomCharacteristics
type RoomCharacteristics struct {
	Type                  string                              `json:"@type"`
	BedTypeCode           *otamodel.EnumStringBase            `json:"BedTypeCode"`            // Assigned Type: ota2-0400:BedType_Enum
	Category              *otamodel.EnumStringBase            `json:"Category"`               // Assigned Type: ota2-0400:SegmentCategory_Enum
	ClassificationCode    *otamodel.EnumStringBase            `json:"ClassificationCode"`     // Assigned Type: ota2-0400:GuestRoomInfo_Enum
	CompositeInd          *bool                               `json:"compositeInd,omitempty"` // Indicates that the room (suite) is a composite of smaller units.
	ExtensionPointSummary interface{}                         `json:"ExtensionPoint_Summary"`
	FeatureAccessibility  []organization.FeatureAccessibility `json:"FeatureAccessibility,omitempty"`
	FeatureSecurity       []organization.FeatureSecurity      `json:"FeatureSecurity,omitempty"`
	LocationCode          *otamodel.EnumStringBase            `json:"LocationCode"`            // Assigned Type: ota2-0400:RoomLocation_Enum
	NonSmokingInd         *bool                               `json:"nonSmokingInd,omitempty"` // Non-smoking indicator.
	ViewCode              *otamodel.EnumStringBase            `json:"ViewCode"`                // Assigned Type: ota2-0400:RoomViewType_Enum
}

// RoomCharacteristicsDetail Assigned Type: ph-0200:RoomCharacteristics
type RoomCharacteristicsDetail struct {
	RoomCharacteristics
	ArchitectureCode     *string                  `json:"architectureCode,omitempty"` // Assigned Type: ota2-0500:OTA_Code
	ExtensionPointDetail interface{}              `json:"ExtensionPoint_Detail"`
	Floor                *int64                   `json:"floor,omitempty"`             // Floor on which the room is located.
	Gender               *otamodel.EnumStringBase `json:"gender,omitempty"`            // Used to request or specify a gender assignment for a room. Note: Typically used by; Hosteliers.
	InvBlockCode         *string                  `json:"invBlockCode,omitempty"`      // Assigned Type: ota2-0500:StringTiny
	NumberOfBathrooms    *int64                   `json:"numberOfBathrooms,omitempty"` // The number of bathrooms in the unit.
	NumberOfBedrooms     *int64                   `json:"numberOfBedrooms,omitempty"`  // The number of bedrooms in the unit.
	SharedRoomInd        *bool                    `json:"sharedRoomInd,omitempty"`     // If TRUE, the room requires or has sharing available. Note: Typically used by Hosteliers.
	SizeMeasurement      *string                  `json:"sizeMeasurement,omitempty"`   // Textual description of room dimensions.
	TypeCode             *string                  `json:"typeCode,omitempty"`          // Assigned Type: ota2-0500:StringTiny
	UnitNumber           *string                  `json:"unitNumber,omitempty"`        // Assigned Type: ota2-0500:StringTiny
}
