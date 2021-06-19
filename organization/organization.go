package organization

import (
	"github.com/hosts-guru/OTAModels-golang"
	"github.com/hosts-guru/OTAModels-golang/common"
)

type FeatureAccessibility struct {
	Type                 string                   `json:"@type,omitempty"`
	Description          *common.Description      `json:"Description"`          // Assigned Type: ota2-0500:Description
	Proximity            *otamodel.EnumStringBase `json:"Proximity"`            // Assigned Type: ota2-0400:Proximity_Enum
	Quantity             int64                    `json:"quantity,omitempty"`   // Assigned Type: ota2-0500:NonNegativeInteger
	AccessibilityFeature *otamodel.EnumStringBase `json:"AccessibilityFeature"` // Assigned Type: ota2-0400:DisabilityFeature_Enum
}

type FeatureSecurity struct {
	Type            string                   `json:"@type,omitempty"`
	Description     *common.Description      `json:"Description"`        // Assigned Type: ota2-0500:Description
	Proximity       *otamodel.EnumStringBase `json:"Proximity"`          // Assigned Type: ota2-0400:Proximity_Enum
	Quantity        int64                    `json:"quantity,omitempty"` // Assigned Type: ota2-0500:NonNegativeInteger
	SecurityFeature *otamodel.EnumStringBase `json:"SecurityFeature"`    // Assigned Type: ota2-0400:SecurityFeature_Enum
}
