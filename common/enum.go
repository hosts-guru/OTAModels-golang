package common

import (
	"github.com/hosts-guru/OTAModels-golang"
)

// TextFormatType Indicates the format of text used in the description e.g. unformatted  or html.
const (
	TextFormatTypeHTML      otamodel.EnumString = "HTML"
	TextFormatTypePlainText otamodel.EnumString = "PlainText"
)

// OptInStatusEnum If present and set to opt out, a customer has explicitly opted out of marketing
// communication. When set to opt out it overrides the ShareMarketInd.
const (
	OptInStatusEnumUnknown otamodel.EnumString = "Unknown"
	OptedIn                otamodel.EnumString = "OptedIn"
	OptedOut               otamodel.EnumString = "OptedOut"
)

//  YesNoInherit Permission for sharing data for marketing purposes.
//
// Permission for sharing data for synchronization of information held by other travel
// service providers.
const (
	Inherit otamodel.EnumString = "Inherit"
	No      otamodel.EnumString = "No"
	Yes     otamodel.EnumString = "Yes"
)

const (
	AddressRoleBusiness    otamodel.EnumString = "Business"
	AddressRoleDestination otamodel.EnumString = "Destination"
	AddressRoleHome        otamodel.EnumString = "Home"
	AddressRoleOther       otamodel.EnumString = "Other"
)

// ProximityEnumBase Specifies if the item in question is onsite, offsite, nearby, etc.
// Todo: verify to move
type ProximityEnumBase string

const (
	InformationNotAvailable ProximityEnumBase = "Information not available"
	Nearby                  ProximityEnumBase = "Nearby"
	Offsite                 ProximityEnumBase = "Offsite"
	Onsite                  ProximityEnumBase = "Onsite"
	OnsiteAndOffsite        ProximityEnumBase = "Onsite and offsite"
	ProximityEnumBaseOther  ProximityEnumBase = "Other_"
)

// ChargeUnitEnumBase  Source: Charge Type (CHG) OpenTravel codelist.
type ChargeUnitEnumBase string

const (
	AdditionsPerStay         ChargeUnitEnumBase = "AdditionsPerStay"
	ChargeUnitEnumBaseDaily  ChargeUnitEnumBase = "Daily"
	ChargeUnitEnumBaseHourly ChargeUnitEnumBase = "Hourly"
	ChargeUnitEnumBaseOther  ChargeUnitEnumBase = "Other_"
	ChargeUnitEnumBaseStay   ChargeUnitEnumBase = "Stay"
	ChargeUnitEnumBaseWeekly ChargeUnitEnumBase = "Weekly"
	Complimentary            ChargeUnitEnumBase = "Complimentary"
	FirstUse                 ChargeUnitEnumBase = "First use"
	Gallon                   ChargeUnitEnumBase = "Gallon"
	HalfDay                  ChargeUnitEnumBase = "Half day"
	Item                     ChargeUnitEnumBase = "Item"
	MaximumCharge            ChargeUnitEnumBase = "MaximumCharge"
	MinimumCharge            ChargeUnitEnumBase = "MinimumCharge"
	OneTimeUse               ChargeUnitEnumBase = "OneTimeUse"
	OneWay                   ChargeUnitEnumBase = "OneWay"
	Order                    ChargeUnitEnumBase = "Order"
	OverMinuteCharge         ChargeUnitEnumBase = "OverMinuteCharge"
	PerDozen                 ChargeUnitEnumBase = "Per dozen"
	PerFunction              ChargeUnitEnumBase = "Per function"
	PerGallon                ChargeUnitEnumBase = "Per gallon"
	PerItem                  ChargeUnitEnumBase = "Per item"
	PerMinute                ChargeUnitEnumBase = "Per minute"
	PerOccurance             ChargeUnitEnumBase = "Per occurance"
	PerOrder                 ChargeUnitEnumBase = "Per order"
	PerPerson                ChargeUnitEnumBase = "Per person"
	PerPersonPerNight        ChargeUnitEnumBase = "Per person per night"
	PerPersonPerStay         ChargeUnitEnumBase = "Per person per stay"
	PerRental                ChargeUnitEnumBase = "Per rental"
	PerRoom                  ChargeUnitEnumBase = "Per room"
	PerRoomPerNight          ChargeUnitEnumBase = "Per room per night"
	PerRoomPerStay           ChargeUnitEnumBase = "Per room per stay"
	PerTray                  ChargeUnitEnumBase = "Per tray"
	PerUnit                  ChargeUnitEnumBase = "Per unit"
	Rental                   ChargeUnitEnumBase = "Rental"
	ReservationBooking       ChargeUnitEnumBase = "Reservation/Booking"
	Room                     ChargeUnitEnumBase = "Room"
	RoundTrip                ChargeUnitEnumBase = "Round trip"
	Tray                     ChargeUnitEnumBase = "Tray"
	Unit                     ChargeUnitEnumBase = "Unit"
)

// TargetEnvironmentEnum Test or Production target system indicator.
type TargetEnvironmentEnum string

const (
	Production TargetEnvironmentEnum = "Production"
	Test       TargetEnvironmentEnum = "Test"
)

// TransactionStatusCodeEnum This indicates where this message falls within a sequence of messages.
type TransactionStatusCodeEnum string

const (
	Continuation TransactionStatusCodeEnum = "Continuation"
	End          TransactionStatusCodeEnum = "End"
	InSeries     TransactionStatusCodeEnum = "InSeries"
	Rollback     TransactionStatusCodeEnum = "Rollback"
	Start        TransactionStatusCodeEnum = "Start"
	Subsequent   TransactionStatusCodeEnum = "Subsequent"
)

// BookingChannelEnumBase  a type of booking channel.
type BookingChannelEnumBase string

const (
	Agent                            BookingChannelEnumBase = "Agent"
	AlternativeDistributionSystemADS BookingChannelEnumBase = "Alternative distribution system (ADS)"
	BookingChannelEnumBaseOther      BookingChannelEnumBase = "Other_"
	CentralReservationSystemCRS      BookingChannelEnumBase = "Central reservation system (CRS)"
	GlobalDistributionSystemGDS      BookingChannelEnumBase = "Global distribution system (GDS)"
	Internet                         BookingChannelEnumBase = "Internet"
	Kiosk                            BookingChannelEnumBase = "Kiosk"
	PropertyManagementSystemPMS      BookingChannelEnumBase = "Property management system (PMS)"
	SalesAndCateringSystemSCS        BookingChannelEnumBase = "Sales and catering system (SCS)"
	TourOperatorSystemTOS            BookingChannelEnumBase = "Tour operator system (TOS)"
)

// UnitOfMeasureEnum Provides the unit of measure for the altitude (e.g., feet, meters, miles, kilometers).
// Refer to OpenTravel Code List Unit of Measure Code (UOM).
//
// The unit of measure in a code format. Refer to OpenTravel Code List Unit of Measure Code
// (UOM).
//
// Defines the average precipitation for the time as designated in Period and is qualified
// by the UnitOfMeasure.
//
// Assigned Type: ota2-0400:UnitOfMeasure_Enum
type UnitOfMeasureEnum string

const (
	Acre                   UnitOfMeasureEnum = "Acre"
	Centimeters            UnitOfMeasureEnum = "Centimeters"
	CubicMeters            UnitOfMeasureEnum = "Cubic meters"
	Feet                   UnitOfMeasureEnum = "Feet"
	Gallons                UnitOfMeasureEnum = "Gallons"
	Gigabytes              UnitOfMeasureEnum = "Gigabytes"
	Gram                   UnitOfMeasureEnum = "Gram"
	Hectare                UnitOfMeasureEnum = "Hectare"
	Inches                 UnitOfMeasureEnum = "Inches"
	Kilograms              UnitOfMeasureEnum = "Kilograms"
	Kilometers             UnitOfMeasureEnum = "Kilometers"
	Kilowatts              UnitOfMeasureEnum = "Kilowatts"
	Liters                 UnitOfMeasureEnum = "Liters"
	Megabytes              UnitOfMeasureEnum = "Megabytes"
	Meters                 UnitOfMeasureEnum = "Meters"
	Miles                  UnitOfMeasureEnum = "Miles"
	Millimeters            UnitOfMeasureEnum = "Millimeters"
	Ounce                  UnitOfMeasureEnum = "Ounce"
	Pixels                 UnitOfMeasureEnum = "Pixels"
	Pounds                 UnitOfMeasureEnum = "Pounds"
	SquareCentimeter       UnitOfMeasureEnum = "Square centimeter"
	SquareFeet             UnitOfMeasureEnum = "Square feet"
	SquareInch             UnitOfMeasureEnum = "Square inch"
	SquareMeters           UnitOfMeasureEnum = "Square meters"
	SquareMillimeter       UnitOfMeasureEnum = "Square millimeter"
	SquareYard             UnitOfMeasureEnum = "Square yard"
	UnitOfMeasureEnumBlock UnitOfMeasureEnum = "Block"
	Yards                  UnitOfMeasureEnum = "Yards"
)

// StatusEnumBase The status condition.
type StatusEnumBase string

const (
	Complete              StatusEnumBase = "Complete"
	Incomplete            StatusEnumBase = "Incomplete"
	NotProcessed          StatusEnumBase = "Not processed"
	StatusEnumBaseOther   StatusEnumBase = "Other_"
	StatusEnumBaseUnknown StatusEnumBase = "Unknown"
)

// ErrorWarningTypeEnumBase Tthe category of error or warning.
type ErrorWarningTypeEnumBase string

const (
	ApplicationError                 ErrorWarningTypeEnumBase = "Application error"
	Authentication                   ErrorWarningTypeEnumBase = "Authentication"
	AuthenticationTimeout            ErrorWarningTypeEnumBase = "Authentication timeout"
	Authorization                    ErrorWarningTypeEnumBase = "Authorization"
	BusinessRule                     ErrorWarningTypeEnumBase = "Business rule"
	ErrorWarningTypeEnumBaseAdvisory ErrorWarningTypeEnumBase = "Advisory"
	ErrorWarningTypeEnumBaseOther    ErrorWarningTypeEnumBase = "Other_"
	ErrorWarningTypeEnumBaseUnknown  ErrorWarningTypeEnumBase = "Unknown"
	NoImplementation                 ErrorWarningTypeEnumBase = "No implementation"
	ProcessingException              ErrorWarningTypeEnumBase = "Processing exception"
	ProtocolViolation                ErrorWarningTypeEnumBase = "Protocol violation"
	RequiredFieldMissing             ErrorWarningTypeEnumBase = "Required field missing"
)

// BasisTypeBase Provides basis for how an amount was calculated.
type BasisTypeBase string

const (
	BasisTypeBaseDays  BasisTypeBase = "Days"
	BasisTypeBaseOther BasisTypeBase = "Other_"
	Entire             BasisTypeBase = "Entire"
	First              BasisTypeBase = "First"
	FirstAndLast       BasisTypeBase = "First and last"
	Last               BasisTypeBase = "Last"
	Nights             BasisTypeBase = "Nights"
)

// DayOfWeekEnum Day of the Week types.
type DayOfWeekEnum string

const (
	Friday    DayOfWeekEnum = "Friday"
	Monday    DayOfWeekEnum = "Monday"
	Saturday  DayOfWeekEnum = "Saturday"
	Sunday    DayOfWeekEnum = "Sunday"
	Thursday  DayOfWeekEnum = "Thursday"
	Tuesday   DayOfWeekEnum = "Tuesday"
	Wednesday DayOfWeekEnum = "Wednesday"
)

// DeadlineDropTimeEnum An enumerated list of deadline types e.g.Before Arrival.
//
// An enumerated type indicating when the deadline drop time goes into effect.
type DeadlineDropTimeEnum string

const (
	AfterArrival      DeadlineDropTimeEnum = "AfterArrival"
	AfterBooking      DeadlineDropTimeEnum = "AfterBooking"
	AfterConfirmation DeadlineDropTimeEnum = "AfterConfirmation"
	AfterDeparture    DeadlineDropTimeEnum = "AfterDeparture"
	BeforeArrival     DeadlineDropTimeEnum = "BeforeArrival"
	BeforeDeparture   DeadlineDropTimeEnum = "BeforeDeparture"
)

// Source: Charge Type (CHG) OpenTravel codelist.
type FrequencyEnumBase string

const (
	FrequencyEnumBaseDaily    FrequencyEnumBase = "Daily"
	FrequencyEnumBaseFirstUse FrequencyEnumBase = "FirstUse"
	FrequencyEnumBaseHalfDay  FrequencyEnumBase = "HalfDay"
	FrequencyEnumBaseHourly   FrequencyEnumBase = "Hourly"
	FrequencyEnumBaseMinute   FrequencyEnumBase = "Minute"
	FrequencyEnumBaseOther    FrequencyEnumBase = "Other_"
	FrequencyEnumBasePerDozen FrequencyEnumBase = "PerDozen"
	FrequencyEnumBaseStay     FrequencyEnumBase = "Stay"
	FrequencyEnumBaseWeekly   FrequencyEnumBase = "Weekly"
	Night                     FrequencyEnumBase = "Night"
	Occurrence                FrequencyEnumBase = "Occurrence"
)

// Describes the type of fee or tax.
type TaxTypeEnumBase string

const (
	AssessmentLicenseTax   TaxTypeEnumBase = "Assessment/license tax"
	BedTax                 TaxTypeEnumBase = "Bed tax"
	CityTax                TaxTypeEnumBase = "City tax"
	ConventionTax          TaxTypeEnumBase = "Convention tax"
	CountryTax             TaxTypeEnumBase = "Country tax"
	CountyTax              TaxTypeEnumBase = "County tax"
	EnergyTax              TaxTypeEnumBase = "Energy tax"
	FederalTax             TaxTypeEnumBase = "Federal tax"
	FoodBeverageTax        TaxTypeEnumBase = "Food & beverage tax"
	GoodsAndServicesTaxGST TaxTypeEnumBase = "Goods and services tax (GST)"
	InsurancePremiumTax    TaxTypeEnumBase = "Insurance Premium Tax"
	LodgingTax             TaxTypeEnumBase = "Lodging tax"
	NationalGovernmentTax  TaxTypeEnumBase = "National government tax"
	OccupancyTax           TaxTypeEnumBase = "Occupancy tax"
	RoomTax                TaxTypeEnumBase = "Room Tax"
	SalesTax               TaxTypeEnumBase = "Sales tax"
	StateTax               TaxTypeEnumBase = "State tax"
	SurplusLinesTax        TaxTypeEnumBase = "Surplus Lines Tax"
	TaxTypeEnumBaseOther   TaxTypeEnumBase = "Other_"
	TotalTax               TaxTypeEnumBase = "Total tax"
	TourismTax             TaxTypeEnumBase = "Tourism tax"
	ValueAddedTaxVAT       TaxTypeEnumBase = "Value Added Tax (VAT)"
)

// Specifies a type of email address.
type EmailAddressTypeEnumBase string

const (
	EmailAddressTypeEnumBaseBusiness        EmailAddressTypeEnumBase = "Business"
	EmailAddressTypeEnumBaseManagingCompany EmailAddressTypeEnumBase = "Managing company"
	EmailAddressTypeEnumBaseOther           EmailAddressTypeEnumBase = "Other_"
	EmailAddressTypeEnumBaseProperty        EmailAddressTypeEnumBase = "Property"
	EmailAddressTypeEnumBaseSalesOffice     EmailAddressTypeEnumBase = "Sales office"
	Listserve                               EmailAddressTypeEnumBase = "Listserve"
	Personal                                EmailAddressTypeEnumBase = "Personal"
	ReservationOffice                       EmailAddressTypeEnumBase = "Reservation office"
)

type EnumTelephoneRoleBase string

const (
	EnumTelephoneRoleBaseHome   EnumTelephoneRoleBase = "Home"
	EnumTelephoneRoleBaseMobile EnumTelephoneRoleBase = "Mobile"
	EnumTelephoneRoleBaseOffice EnumTelephoneRoleBase = "Office"
	EnumTelephoneRoleBaseOther  EnumTelephoneRoleBase = "Other"
)

// Specifies where the contact is located.
type ContactLocationEnumBase string

const (
	CentralReservationOffice           ContactLocationEnumBase = "Central reservation office"
	ContactLocationEnumBaseOther       ContactLocationEnumBase = "Other_"
	ContactLocationEnumBaseSalesOffice ContactLocationEnumBase = "Sales office"
	CorporateHeadquarters              ContactLocationEnumBase = "Corporate headquarters"
	CorporateOffice                    ContactLocationEnumBase = "Corporate office"
	CustomerServiceOffice              ContactLocationEnumBase = "Customer service office"
	DivisionalOffice                   ContactLocationEnumBase = "Divisional office"
	FranchiseCompany                   ContactLocationEnumBase = "Franchise company"
	GlobalSalesOffice                  ContactLocationEnumBase = "Global sales office"
	HomeResidence                      ContactLocationEnumBase = "Home/residence"
	HotelDirect                        ContactLocationEnumBase = "Hotel direct"
	LocalReservationOffice             ContactLocationEnumBase = "Local reservation office"
	ManagementCompany                  ContactLocationEnumBase = "Management company"
	OwnershipCompany                   ContactLocationEnumBase = "Ownership company"
	RegionalSalesOffice                ContactLocationEnumBase = "Regional sales office"
	TechnicalSupportOffice             ContactLocationEnumBase = "Technical support office"
)

// Defines the preferred method of communication or the method of distribution for the
// material.
type DistributionTypeEnumBase string

const (
	AirportCollection                 DistributionTypeEnumBase = "Airport collection"
	Any                               DistributionTypeEnumBase = "Any"
	CityOffice                        DistributionTypeEnumBase = "City office"
	Courier                           DistributionTypeEnumBase = "Courier"
	DistributionTypeEnumBaseEmail     DistributionTypeEnumBase = "Email"
	DistributionTypeEnumBaseFax       DistributionTypeEnumBase = "Fax"
	DistributionTypeEnumBaseOther     DistributionTypeEnumBase = "Other_"
	DistributionTypeEnumBaseTelephone DistributionTypeEnumBase = "Telephone"
	ExpressMail                       DistributionTypeEnumBase = "Express mail"
	FTP                               DistributionTypeEnumBase = "FTP"
	HTTP                              DistributionTypeEnumBase = "HTTP"
	HotelDesk                         DistributionTypeEnumBase = "Hotel desk"
	Mail                              DistributionTypeEnumBase = "Mail"
	NonXML                            DistributionTypeEnumBase = "Non-XML"
	Website                           DistributionTypeEnumBase = "Website"
	WillCall                          DistributionTypeEnumBase = "Will call"
	XML                               DistributionTypeEnumBase = "XML"
)

type LocationCategoryEnumBase string

const (
	Beachfront                                LocationCategoryEnumBase = "Beachfront"
	BuisnessDistrict                          LocationCategoryEnumBase = "Buisness district"
	Downtown                                  LocationCategoryEnumBase = "Downtown"
	EntertainmentDistrict                     LocationCategoryEnumBase = "Entertainment district"
	Expressway                                LocationCategoryEnumBase = "Expressway"
	Gulf                                      LocationCategoryEnumBase = "Gulf"
	LocationCategoryEnumBaseAirport           LocationCategoryEnumBase = "Airport"
	LocationCategoryEnumBaseBay               LocationCategoryEnumBase = "Bay"
	LocationCategoryEnumBaseBeach             LocationCategoryEnumBase = "Beach"
	LocationCategoryEnumBaseCity              LocationCategoryEnumBase = "City"
	LocationCategoryEnumBaseFinancialDistrict LocationCategoryEnumBase = "Financial district"
	LocationCategoryEnumBaseLake              LocationCategoryEnumBase = "Lake"
	LocationCategoryEnumBaseMarina            LocationCategoryEnumBase = "Marina"
	LocationCategoryEnumBaseOther             LocationCategoryEnumBase = "Other_"
	LocationCategoryEnumBasePark              LocationCategoryEnumBase = "Park"
	LocationCategoryEnumBaseRiver             LocationCategoryEnumBase = "River"
	Mountain                                  LocationCategoryEnumBase = "Mountain"
	Oceanfront                                LocationCategoryEnumBase = "Oceanfront"
	Resort                                    LocationCategoryEnumBase = "Resort"
	Rural                                     LocationCategoryEnumBase = "Rural"
	ShoppingDistrict                          LocationCategoryEnumBase = "Shopping district"
	Suburban                                  LocationCategoryEnumBase = "Suburban"
	TheatreDistrict                           LocationCategoryEnumBase = "Theatre district"
	Waterfront                                LocationCategoryEnumBase = "Waterfront"
)

// Defines the 'Units' that can be applied to Stay restrictions.
type DurationUnitEnumBase string

const (
	DurationUnitEnumBaseDays  DurationUnitEnumBase = "Days"
	DurationUnitEnumBaseOther DurationUnitEnumBase = "Other_"
	Fri                       DurationUnitEnumBase = "FRI"
	Hours                     DurationUnitEnumBase = "Hours"
	Minutes                   DurationUnitEnumBase = "Minutes"
	Mon                       DurationUnitEnumBase = "MON"
	Months                    DurationUnitEnumBase = "Months"
	Sat                       DurationUnitEnumBase = "SAT"
	Sun                       DurationUnitEnumBase = "SUN"
	Thu                       DurationUnitEnumBase = "THU"
	Tues                      DurationUnitEnumBase = "TUES"
	Wed                       DurationUnitEnumBase = "WED"
)

// This is the object referred to by the relative position (e.g. cross street, airport).
// Refer to OpenTravel Code List Index Point Code (IPC).
//
// Indicates the type of location being referenced (e.g., Airport, Hotel).
type IndexPointCodeEnum string

const (
	Attraction                       IndexPointCodeEnum = "Attraction"
	BusCoachStation                  IndexPointCodeEnum = "Bus/coach station"
	BusinessLocation                 IndexPointCodeEnum = "Business location"
	CarRentalLocation                IndexPointCodeEnum = "Car rental location"
	ConventionCenter                 IndexPointCodeEnum = "Convention center"
	CrossStreet                      IndexPointCodeEnum = "Cross street"
	District                         IndexPointCodeEnum = "District"
	Educational                      IndexPointCodeEnum = "Educational"
	Event                            IndexPointCodeEnum = "Event"
	GeoLocation                      IndexPointCodeEnum = "Geo location"
	GroundTransport                  IndexPointCodeEnum = "Ground transport"
	Heliport                         IndexPointCodeEnum = "Heliport"
	Highway                          IndexPointCodeEnum = "Highway"
	HighwayExit                      IndexPointCodeEnum = "Highway exit"
	IndexPointCodeEnumAirport        IndexPointCodeEnum = "Airport"
	IndexPointCodeEnumCity           IndexPointCodeEnum = "City"
	IndexPointCodeEnumHotel          IndexPointCodeEnum = "Hotel"
	IndexPointCodeEnumIntersection   IndexPointCodeEnum = "Intersection"
	IndexPointCodeEnumShoppingCenter IndexPointCodeEnum = "Shopping center"
	LocalLandmark                    IndexPointCodeEnum = "Local landmark"
	MetroArea                        IndexPointCodeEnum = "Metro area"
	MilitaryBases                    IndexPointCodeEnum = "Military bases"
	NearestMajorCity                 IndexPointCodeEnum = "Nearest major city"
	NeighbouringState                IndexPointCodeEnum = "Neighbouring state"
	ParkRecreationalArea             IndexPointCodeEnum = "Park/recreational area"
	Port                             IndexPointCodeEnum = "Port"
	RailStation                      IndexPointCodeEnum = "Rail station"
	RegionalExpressTrainStation      IndexPointCodeEnum = "Regional express train station"
	ResortSkiArea                    IndexPointCodeEnum = "Resort/ski area"
	Sports                           IndexPointCodeEnum = "Sports"
	SubwayStation                    IndexPointCodeEnum = "Subway station"
	SurroundingCity                  IndexPointCodeEnum = "Surrounding city"
	TransportationPoints             IndexPointCodeEnum = "Transportation points"
)

// Indicates the accuracy of the property's geo-coding, since the property's longitude and
// latitude may not always be exact. Refer to OpenTravel Code List Position Accuracy Code
// (PAC).
//
// Specifies the level of accuracy for the positon.
type PositionAccuracyEnum string

const (
	County                           PositionAccuracyEnum = "County"
	Exact                            PositionAccuracyEnum = "Exact"
	PositionAccuracyEnumBlock        PositionAccuracyEnum = "Block"
	PositionAccuracyEnumCity         PositionAccuracyEnum = "City"
	PositionAccuracyEnumCountry      PositionAccuracyEnum = "Country"
	PositionAccuracyEnumIntersection PositionAccuracyEnum = "Intersection"
	PositionAccuracyEnumProperty     PositionAccuracyEnum = "Property"
	State                            PositionAccuracyEnum = "State"
	Street                           PositionAccuracyEnum = "Street"
	Zip5Code                         PositionAccuracyEnum = "Zip5Code"
	Zip7Code                         PositionAccuracyEnum = "Zip7Code"
	Zip9Code                         PositionAccuracyEnum = "Zip9Code"
)

// Used to indicate whether the context is to a facility or from a facility.
type ToFromEnum string

const (
	FromFacility ToFromEnum = "FromFacility"
	ToFacility   ToFromEnum = "ToFacility"
)

// Types of power.
type PowerTypeEnumBase string

const (
	Diesel                 PowerTypeEnumBase = "Diesel"
	Electric               PowerTypeEnumBase = "Electric"
	Gas                    PowerTypeEnumBase = "Gas"
	Oil                    PowerTypeEnumBase = "Oil"
	PowerTypeEnumBaseOther PowerTypeEnumBase = "Other_"
	Solar                  PowerTypeEnumBase = "Solar"
	Wind                   PowerTypeEnumBase = "Wind"
)

// Specifies areas where recycling is implemented.
type RecyclingLocationsEnumBase string

const (
	CommonAreas                           RecyclingLocationsEnumBase = "Common areas"
	Garage                                RecyclingLocationsEnumBase = "Garage"
	RecyclingLocationsEnumBaseGuestRoom   RecyclingLocationsEnumBase = "Guest room"
	RecyclingLocationsEnumBaseMeetingRoom RecyclingLocationsEnumBase = "Meeting room"
	RecyclingLocationsEnumBaseOther       RecyclingLocationsEnumBase = "Other_"
	RecyclingLocationsEnumBaseRestaurant  RecyclingLocationsEnumBase = "Restaurant"
	StaffArea                             RecyclingLocationsEnumBase = "Staff area"
)

// Specifies products that are recycled.
type RecycledProductsEnumBase string

const (
	Aluminum                      RecycledProductsEnumBase = "Aluminum"
	Batteries                     RecycledProductsEnumBase = "Batteries"
	ComputerParts                 RecycledProductsEnumBase = "Computer parts"
	Glass                         RecycledProductsEnumBase = "Glass"
	Lightbulbs                    RecycledProductsEnumBase = "Lightbulbs"
	Paper                         RecycledProductsEnumBase = "Paper"
	Plastic                       RecycledProductsEnumBase = "Plastic"
	RecycledProductsEnumBaseOther RecycledProductsEnumBase = "Other_"
)

// Describes the type of fee or tax.
type FeeTypeEnumBase string

const (
	AdultRollawayFee                FeeTypeEnumBase = "Adult rollaway fee"
	ApplicationFee                  FeeTypeEnumBase = "Application Fee"
	BanquetServiceFee               FeeTypeEnumBase = "Banquet service fee"
	BeverageWithAlcohol             FeeTypeEnumBase = "Beverage with alcohol"
	BeverageWithoutAlcohol          FeeTypeEnumBase = "Beverage without alcohol"
	ChildRollawayCharge             FeeTypeEnumBase = "Child rollaway charge"
	CityHotelFee                    FeeTypeEnumBase = "City hotel fee"
	CribFee                         FeeTypeEnumBase = "Crib fee"
	DestinationAmenityFee           FeeTypeEnumBase = "Destination amenity fee"
	EarlyCheckoutFee                FeeTypeEnumBase = "Early checkout fee"
	Exempt                          FeeTypeEnumBase = "Exempt"
	ExpressHandlingFee              FeeTypeEnumBase = "Express Handling Fee"
	ExtraChildCharge                FeeTypeEnumBase = "Extra child charge"
	ExtraPersonCharge               FeeTypeEnumBase = "Extra person charge"
	FeeTypeEnumBaseMiscellaneous    FeeTypeEnumBase = "Miscellaneous"
	FeeTypeEnumBaseOther            FeeTypeEnumBase = "Other_"
	Food                            FeeTypeEnumBase = "Food"
	LocalFee                        FeeTypeEnumBase = "Local fee"
	MaintenanceFee                  FeeTypeEnumBase = "Maintenance fee"
	PackageFee                      FeeTypeEnumBase = "Package fee"
	PetSanitationFee                FeeTypeEnumBase = "Pet sanitation fee"
	ResortFee                       FeeTypeEnumBase = "Resort fee"
	RollawayFee                     FeeTypeEnumBase = "Rollaway fee"
	RoomServiceFee                  FeeTypeEnumBase = "Room service fee"
	ServiceCharge                   FeeTypeEnumBase = "Service charge"
	Standard                        FeeTypeEnumBase = "Standard"
	StandardFoodAndBeverageGratuity FeeTypeEnumBase = "Standard food and beverage gratuity"
	StateCostRecoveryFee            FeeTypeEnumBase = "State cost recovery fee"
	Surcharge                       FeeTypeEnumBase = "Surcharge"
	Tobacco                         FeeTypeEnumBase = "Tobacco"
	TotalSurcharges                 FeeTypeEnumBase = "Total surcharges"
)

// Specifies the travel sector.
type TravelSectorEnumBase string

const (
	Air                       TravelSectorEnumBase = "Air"
	Cruise                    TravelSectorEnumBase = "Cruise"
	Excursion                 TravelSectorEnumBase = "Excursion"
	Golf                      TravelSectorEnumBase = "Golf"
	Insurance                 TravelSectorEnumBase = "Insurance"
	PackageOption             TravelSectorEnumBase = "Package option"
	Rail                      TravelSectorEnumBase = "Rail"
	Tour                      TravelSectorEnumBase = "Tour"
	TravelSectorEnumBaseCar   TravelSectorEnumBase = "Car"
	TravelSectorEnumBaseFerry TravelSectorEnumBase = "Ferry"
	TravelSectorEnumBaseHotel TravelSectorEnumBase = "Hotel"
	TravelSectorEnumBaseOther TravelSectorEnumBase = "Other_"
)

// Identifies the type of detail being sent.
type AdditionalDetailTypeEnumBase string

const (
	AdditionalDetailTypeEnumBaseDrivingDirections AdditionalDetailTypeEnumBase = "Driving directions"
	AdditionalDetailTypeEnumBaseOther             AdditionalDetailTypeEnumBase = "Other_"
	AdvancedBookingInformation                    AdditionalDetailTypeEnumBase = "Advanced booking information"
	AmenityInformation                            AdditionalDetailTypeEnumBase = "Amenity information"
	AreasServed                                   AdditionalDetailTypeEnumBase = "Areas served"
	BookingGuidelines                             AdditionalDetailTypeEnumBase = "Booking guidelines"
	CancellationInformation                       AdditionalDetailTypeEnumBase = "Cancellation information"
	CateringDescription                           AdditionalDetailTypeEnumBase = "Catering description"
	CheckInCheckOutInformation                    AdditionalDetailTypeEnumBase = "Check in check out information"
	CheckInPolicy                                 AdditionalDetailTypeEnumBase = "Check-in policy"
	CheckOutPolicy                                AdditionalDetailTypeEnumBase = "Check-out policy"
	ChildrenInformation                           AdditionalDetailTypeEnumBase = "Children information"
	CommissionInformation                         AdditionalDetailTypeEnumBase = "Commission information"
	ContractNegotiatedBookingInformation          AdditionalDetailTypeEnumBase = "Contract/negotiated booking information"
	CorporateBookingInformation                   AdditionalDetailTypeEnumBase = "Corporate booking information"
	CuisineDescription                            AdditionalDetailTypeEnumBase = "Cuisine description"
	CustomsInformationForMaterial                 AdditionalDetailTypeEnumBase = "Customs information for material"
	DepositInformation                            AdditionalDetailTypeEnumBase = "Deposit information"
	DepositPolicyForMasterAccount                 AdditionalDetailTypeEnumBase = "Deposit policy for master account"
	DepositPolicyForReservations                  AdditionalDetailTypeEnumBase = "Deposit policy for reservations"
	DrivingDirectionsFromTheEast                  AdditionalDetailTypeEnumBase = "Driving directions from the east"
	DrivingDirectionsFromTheNorth                 AdditionalDetailTypeEnumBase = "Driving directions from the north"
	DrivingDirectionsFromTheSouth                 AdditionalDetailTypeEnumBase = "Driving directions from the south"
	DrivingDirectionsFromTheWest                  AdditionalDetailTypeEnumBase = "Driving directions from the west"
	EarlyCheckoutDescription                      AdditionalDetailTypeEnumBase = "Early checkout description"
	ExpressCheckInPolicy                          AdditionalDetailTypeEnumBase = "Express check-in policy"
	ExpressCheckOutPolicy                         AdditionalDetailTypeEnumBase = "Express check-out policy"
	ExtendedStayInformation                       AdditionalDetailTypeEnumBase = "Extended stay information"
	ExtraChargeInformation                        AdditionalDetailTypeEnumBase = "Extra charge information"
	ExtraPersonInformation                        AdditionalDetailTypeEnumBase = "Extra person information"
	FacilityRestrictions                          AdditionalDetailTypeEnumBase = "Facility restrictions"
	FamilyPlanDescription                         AdditionalDetailTypeEnumBase = "Family plan description"
	FoodAndBeverageMinimumsForGroups              AdditionalDetailTypeEnumBase = "Food and beverage minimums for groups"
	GeneralMeetingPlanningInformation             AdditionalDetailTypeEnumBase = "General meeting planning information"
	GovernmentBookingPolicy                       AdditionalDetailTypeEnumBase = "Government booking policy"
	GroupBookingInformation                       AdditionalDetailTypeEnumBase = "Group booking information"
	GroupMeetingPlanningInformation               AdditionalDetailTypeEnumBase = "Group meeting planning information"
	GuaranteeInformation                          AdditionalDetailTypeEnumBase = "Guarantee information"
	InclusionInformation                          AdditionalDetailTypeEnumBase = "Inclusion information"
	LastRoomAvailabilityDescription               AdditionalDetailTypeEnumBase = "Last room availability description"
	LateArrivalInformation                        AdditionalDetailTypeEnumBase = "Late arrival information"
	LateDepartureInformation                      AdditionalDetailTypeEnumBase = "Late departure information"
	MaximumStayInformation                        AdditionalDetailTypeEnumBase = "Maximum stay information"
	MealPlanDescription                           AdditionalDetailTypeEnumBase = "Meal plan description"
	MeetingRoomDescription                        AdditionalDetailTypeEnumBase = "Meeting room description"
	MinimumStayInformation                        AdditionalDetailTypeEnumBase = "Minimum stay information"
	MiscellaneousInformation                      AdditionalDetailTypeEnumBase = "Miscellaneous information"
	OffsiteFacilitiesInformation                  AdditionalDetailTypeEnumBase = "Offsite facilities information"
	OffsiteRecreationalActivitiesInformation      AdditionalDetailTypeEnumBase = "Offsite recreational activities information"
	OffsiteServicesInformation                    AdditionalDetailTypeEnumBase = "Offsite services information"
	OnsiteFacilitiesInformation                   AdditionalDetailTypeEnumBase = "Onsite facilities information"
	OnsiteRecreationalActivitiesInformation       AdditionalDetailTypeEnumBase = "Onsite recreational activities information"
	OnsiteServicesInformation                     AdditionalDetailTypeEnumBase = "Onsite services information"
	OversoldPolicyDescription                     AdditionalDetailTypeEnumBase = "Oversold policy description"
	PackageInformation                            AdditionalDetailTypeEnumBase = "Package information"
	PetPolicyDescription                          AdditionalDetailTypeEnumBase = "Pet policy description"
	PromotionalInformation                        AdditionalDetailTypeEnumBase = "Promotional information"
	PropertyDescription                           AdditionalDetailTypeEnumBase = "Property description"
	PropertyLocation                              AdditionalDetailTypeEnumBase = "Property location"
	RateDescription                               AdditionalDetailTypeEnumBase = "Rate description"
	RateDisclaimerInformation                     AdditionalDetailTypeEnumBase = "Rate disclaimer information"
	RestaurantServices                            AdditionalDetailTypeEnumBase = "Restaurant services"
	RoomDecorDescription                          AdditionalDetailTypeEnumBase = "Room decor description"
	RoomInformation                               AdditionalDetailTypeEnumBase = "Room information"
	RoomTypeUpgradeDescription                    AdditionalDetailTypeEnumBase = "Room type upgrade description"
	Seasons                                       AdditionalDetailTypeEnumBase = "Seasons"
	SecurityInformation                           AdditionalDetailTypeEnumBase = "Security information"
	ServiceChargeInformation                      AdditionalDetailTypeEnumBase = "Service charge information"
	SpecialEvents                                 AdditionalDetailTypeEnumBase = "Special events"
	SpecialOffersDescription                      AdditionalDetailTypeEnumBase = "Special offers description"
	SurchargeInformation                          AdditionalDetailTypeEnumBase = "Surcharge information"
	TaxInformation                                AdditionalDetailTypeEnumBase = "Tax information"
	TravelIndustryBookingInformation              AdditionalDetailTypeEnumBase = "Travel industry booking information"
	VisaTravelRequirementInformation              AdditionalDetailTypeEnumBase = "Visa/travel requirement information"
)

// Specifies the category of the mulitmedia item.
type PictureCategoryEnumBase string

const (
	Ballroom                             PictureCategoryEnumBase = "Ballroom"
	BarLounge                            PictureCategoryEnumBase = "Bar/Lounge"
	Basics                               PictureCategoryEnumBase = "Basics"
	BusinessCenter                       PictureCategoryEnumBase = "Business center"
	ExteriorView                         PictureCategoryEnumBase = "Exterior view"
	GolfCourse                           PictureCategoryEnumBase = "Golf course"
	GuestRoomAmenity                     PictureCategoryEnumBase = "Guest room amenity"
	HealthClub                           PictureCategoryEnumBase = "Health club"
	HotNews                              PictureCategoryEnumBase = "Hot news"
	LobbyView                            PictureCategoryEnumBase = "Lobby view"
	Logo                                 PictureCategoryEnumBase = "Logo"
	Map                                  PictureCategoryEnumBase = "Map"
	PictureCategoryEnumBaseBeach         PictureCategoryEnumBase = "Beach"
	PictureCategoryEnumBaseGuestRoom     PictureCategoryEnumBase = "Guest room"
	PictureCategoryEnumBaseMeetingRoom   PictureCategoryEnumBase = "Meeting room"
	PictureCategoryEnumBaseMiscellaneous PictureCategoryEnumBase = "Miscellaneous"
	PictureCategoryEnumBaseOther         PictureCategoryEnumBase = "Other_"
	PictureCategoryEnumBaseRestaurant    PictureCategoryEnumBase = "Restaurant"
	PoolView                             PictureCategoryEnumBase = "Pool view"
	Promotional                          PictureCategoryEnumBase = "Promotional"
	PropertyAmenity                      PictureCategoryEnumBase = "Property amenity"
	RecreationalFacility                 PictureCategoryEnumBase = "Recreational facility"
	SPA                                  PictureCategoryEnumBase = "Spa"
	Suite                                PictureCategoryEnumBase = "Suite"
)

// Identifies the type of information being sent.
type InformationTypeEnumBase string

const (
	AliasName                                InformationTypeEnumBase = "Alias name"
	Amenities                                InformationTypeEnumBase = "Amenities"
	Attractions                              InformationTypeEnumBase = "Attractions"
	Awards                                   InformationTypeEnumBase = "Awards"
	CorporateLocations                       InformationTypeEnumBase = "Corporate locations"
	DescriptiveContent                       InformationTypeEnumBase = "Descriptive content"
	Dining                                   InformationTypeEnumBase = "Dining"
	Facilities                               InformationTypeEnumBase = "Facilities"
	Geocodes                                 InformationTypeEnumBase = "Geocodes"
	InformationTypeEnumBaseAddress           InformationTypeEnumBase = "Address"
	InformationTypeEnumBaseAdvisory          InformationTypeEnumBase = "Advisory"
	InformationTypeEnumBaseContact           InformationTypeEnumBase = "Contact"
	InformationTypeEnumBaseDescription       InformationTypeEnumBase = "Description"
	InformationTypeEnumBaseDrivingDirections InformationTypeEnumBase = "Driving directions"
	InformationTypeEnumBaseLocation          InformationTypeEnumBase = "Location"
	InformationTypeEnumBaseOther             InformationTypeEnumBase = "Other_"
	InformationTypeEnumBaseTransportation    InformationTypeEnumBase = "Transportation"
	LongName                                 InformationTypeEnumBase = "Long name"
	Marketing                                InformationTypeEnumBase = "Marketing"
	Pictures                                 InformationTypeEnumBase = "Pictures"
	Policy                                   InformationTypeEnumBase = "Policy"
	Recreation                               InformationTypeEnumBase = "Recreation"
	Safety                                   InformationTypeEnumBase = "Safety"
	Services                                 InformationTypeEnumBase = "Services"
	ShortDescription                         InformationTypeEnumBase = "Short description"
	SpecialInstructions                      InformationTypeEnumBase = "Special instructions"
)

// Defines the specific type of identification or certification document.
type DocumentTypeEnumBase string

const (
	AirNexusCard               DocumentTypeEnumBase = "AirNexusCard"
	AlienRegistrationNumber    DocumentTypeEnumBase = "AlienRegistrationNumber"
	BorderCrossingCard         DocumentTypeEnumBase = "BorderCrossingCard"
	CrewMemberCertificate      DocumentTypeEnumBase = "CrewMemberCertificate"
	DocumentTypeEnumBaseOther  DocumentTypeEnumBase = "Other_"
	DriversLicense             DocumentTypeEnumBase = "DriversLicense"
	InsurancePolicyNumber      DocumentTypeEnumBase = "InsurancePolicyNumber"
	KnownTravelerNumber        DocumentTypeEnumBase = "KnownTravelerNumber"
	MerchantMariner            DocumentTypeEnumBase = "MerchantMariner"
	MilitaryIdentification     DocumentTypeEnumBase = "MilitaryIdentification"
	NationalIdentityDocument   DocumentTypeEnumBase = "NationalIdentityDocument"
	NaturalizationCertificate  DocumentTypeEnumBase = "NaturalizationCertificate"
	NonStandard                DocumentTypeEnumBase = "Non-standard"
	Passport                   DocumentTypeEnumBase = "Passport"
	PassportCard               DocumentTypeEnumBase = "PassportCard"
	PermanentResidentCard      DocumentTypeEnumBase = "PermanentResidentCard"
	PilotsLicense              DocumentTypeEnumBase = "PilotsLicense"
	RedressNumber              DocumentTypeEnumBase = "RedressNumber"
	RefugeeTravelDocument      DocumentTypeEnumBase = "RefugeeTravelDocument"
	TaxexEmptionNumber         DocumentTypeEnumBase = "TaxexEmptionNumber"
	VaccinationCertificate     DocumentTypeEnumBase = "VaccinationCertificate"
	VerhicleRegistrationNumber DocumentTypeEnumBase = "VerhicleRegistrationNumber"
	Visa                       DocumentTypeEnumBase = "Visa"
)

// Type of name of the individual, such as former, nickname, alternate or alias name. Refer
// to OpenTravel Code List Name Type (NAM).
//
// OTA Code list: Name Type\tNAM\t\n\t1\tFormer\n\t2\tNickname\n\t3\tAlternate\n\t4\tMaiden
type NameType string

const (
	Alternate NameType = "Alternate"
	Former    NameType = "Former"
	Maiden    NameType = "Maiden"
	Nickname  NameType = "Nickname"
)

// The gender of the person.
type GenderEnum string

const (
	Female            GenderEnum = "Female"
	GenderEnumUnknown GenderEnum = "Unknown"
	Male              GenderEnum = "Male"
)

// Specifies the marital status of the person.
//
// The marital status of a person.
type MaritalStatusEnum string

const (
	Annulled                 MaritalStatusEnum = "Annulled"
	CoHabiting               MaritalStatusEnum = "Co-habiting"
	Divorced                 MaritalStatusEnum = "Divorced"
	Engaged                  MaritalStatusEnum = "Engaged"
	MaritalStatusEnumUnknown MaritalStatusEnum = "Unknown"
	Married                  MaritalStatusEnum = "Married"
	Separated                MaritalStatusEnum = "Separated"
	Single                   MaritalStatusEnum = "Single"
	Widowed                  MaritalStatusEnum = "Widowed"
)

// Used to specify the period of time to which the rates apply.
type TimeUnitEnum string

const (
	Day                TimeUnitEnum = "Day"
	FullDuration       TimeUnitEnum = "Full duration"
	Hour               TimeUnitEnum = "Hour"
	Month              TimeUnitEnum = "Month"
	Second             TimeUnitEnum = "Second"
	TimeUnitEnumMinute TimeUnitEnum = "Minute"
	Week               TimeUnitEnum = "Week"
	Year               TimeUnitEnum = "Year"
)

// Specifies a point of reference.
type ReferencePointEnumBase string

const (
	AmusementPark                           ReferencePointEnumBase = "Amusement park"
	Arena                                   ReferencePointEnumBase = "Arena"
	Bar                                     ReferencePointEnumBase = "Bar"
	BusStation                              ReferencePointEnumBase = "Bus station"
	Church                                  ReferencePointEnumBase = "Church"
	CityCenter                              ReferencePointEnumBase = "City center"
	Corporation                             ReferencePointEnumBase = "Corporation"
	EducationalInstitution                  ReferencePointEnumBase = "Educational Institution"
	FerryStation                            ReferencePointEnumBase = "Ferry station"
	FinancialInstitution                    ReferencePointEnumBase = "Financial institution"
	Landmark                                ReferencePointEnumBase = "Landmark"
	Library                                 ReferencePointEnumBase = "Library"
	Market                                  ReferencePointEnumBase = "Market"
	MedicalFacility                         ReferencePointEnumBase = "Medical facility"
	MetroSubwayStation                      ReferencePointEnumBase = "Metro/subway station"
	Monument                                ReferencePointEnumBase = "Monument"
	Museum                                  ReferencePointEnumBase = "Museum"
	Racetrack                               ReferencePointEnumBase = "Racetrack"
	ReferencePointEnumBaseAirport           ReferencePointEnumBase = "Airport"
	ReferencePointEnumBaseBay               ReferencePointEnumBase = "Bay"
	ReferencePointEnumBaseBeach             ReferencePointEnumBase = "Beach"
	ReferencePointEnumBaseFinancialDistrict ReferencePointEnumBase = "Financial district"
	ReferencePointEnumBaseLake              ReferencePointEnumBase = "Lake"
	ReferencePointEnumBaseMarina            ReferencePointEnumBase = "Marina"
	ReferencePointEnumBaseOther             ReferencePointEnumBase = "Other_"
	ReferencePointEnumBasePark              ReferencePointEnumBase = "Park"
	ReferencePointEnumBaseRestaurant        ReferencePointEnumBase = "Restaurant"
	ReferencePointEnumBaseRiver             ReferencePointEnumBase = "River"
	ReferencePointEnumBaseShoppingCenter    ReferencePointEnumBase = "Shopping center"
	SportsFacility                          ReferencePointEnumBase = "Sports facility"
	Synagogue                               ReferencePointEnumBase = "Synagogue"
	TrainStation                            ReferencePointEnumBase = "Train station"
	Zoo                                     ReferencePointEnumBase = "Zoo"
)

// Where the phone is located.
type PhoneLocationEnumBase string

const (
	BrandReservationsOffice              PhoneLocationEnumBase = "Brand reservations office"
	CentralReservationsOffice            PhoneLocationEnumBase = "Central reservations office"
	PhoneLocationEnumBaseHome            PhoneLocationEnumBase = "Home"
	PhoneLocationEnumBaseManagingCompany PhoneLocationEnumBase = "Managing company"
	PhoneLocationEnumBaseMobile          PhoneLocationEnumBase = "Mobile"
	PhoneLocationEnumBaseOffice          PhoneLocationEnumBase = "Office"
	PhoneLocationEnumBaseOther           PhoneLocationEnumBase = "Other_"
	PhoneLocationEnumBaseSalesOffice     PhoneLocationEnumBase = "Sales office"
	PropertyDirect                       PhoneLocationEnumBase = "Property direct"
	PropertyReservationOffice            PhoneLocationEnumBase = "Property reservation Office"
)

// Defines the technology associated with this phone number.
type PhoneTechTypeEnumBase string

const (
	Pager                      PhoneTechTypeEnumBase = "Pager"
	PhoneTechTypeEnumBaseFax   PhoneTechTypeEnumBase = "Fax"
	PhoneTechTypeEnumBaseOther PhoneTechTypeEnumBase = "Other_"
	TTY                        PhoneTechTypeEnumBase = "TTY"
	Voice                      PhoneTechTypeEnumBase = "Voice"
)

type PhoneUseTypeEnumBase string

const (
	DaytimeContact              PhoneUseTypeEnumBase = "Daytime contact"
	EmergencyContact            PhoneUseTypeEnumBase = "Emergency contact"
	EveningContact              PhoneUseTypeEnumBase = "Evening contact"
	GuestUse                    PhoneUseTypeEnumBase = "Guest use"
	PhoneUseTypeEnumBaseContact PhoneUseTypeEnumBase = "Contact"
	PhoneUseTypeEnumBaseMobile  PhoneUseTypeEnumBase = "Mobile"
	PhoneUseTypeEnumBaseOther   PhoneUseTypeEnumBase = "Other_"
	PickupContact               PhoneUseTypeEnumBase = "Pickup contact"
	TollFreeNumber              PhoneUseTypeEnumBase = "Toll free number"
	TravelArranger              PhoneUseTypeEnumBase = "Travel arranger"
)

// Type of tranportation: Air, Rail, Bus, Boat, Private Auto, Other.
type TransportationEnumBase string

const (
	Bicycle                     TransportationEnumBase = "Bicycle"
	Boat                        TransportationEnumBase = "Boat"
	Bus                         TransportationEnumBase = "Bus"
	CableCar                    TransportationEnumBase = "Cable car"
	Carriage                    TransportationEnumBase = "Carriage"
	CourtesyCar                 TransportationEnumBase = "Courtesy car"
	ExpressTrain                TransportationEnumBase = "Express train"
	Helicopter                  TransportationEnumBase = "Helicopter"
	Limousine                   TransportationEnumBase = "Limousine"
	Metro                       TransportationEnumBase = "Metro"
	Monorail                    TransportationEnumBase = "Monorail"
	Motorbike                   TransportationEnumBase = "Motorbike"
	PackAnimal                  TransportationEnumBase = "Pack animal"
	Plane                       TransportationEnumBase = "Plane"
	Public                      TransportationEnumBase = "Public"
	RentalCar                   TransportationEnumBase = "Rental car"
	Rickshaw                    TransportationEnumBase = "Rickshaw"
	SedanChair                  TransportationEnumBase = "Sedan chair"
	Shuttle                     TransportationEnumBase = "Shuttle"
	Subway                      TransportationEnumBase = "Subway"
	Taxi                        TransportationEnumBase = "Taxi"
	Train                       TransportationEnumBase = "Train"
	TransportationEnumBaseCar   TransportationEnumBase = "Car"
	TransportationEnumBaseFerry TransportationEnumBase = "Ferry"
	TransportationEnumBaseOther TransportationEnumBase = "Other_"
	Trolley                     TransportationEnumBase = "Trolley"
	Tube                        TransportationEnumBase = "Tube"
	Walk                        TransportationEnumBase = "Walk"
	WaterTaxi                   TransportationEnumBase = "Water taxi"
)

// Provides the units in which the AverageHighTemp and AverageLowTemp are defined (i.e.
// Celsius or Fahrenheit).
type TemperatureUnit string

const (
	Celsuis    TemperatureUnit = "Celsuis"
	Fahrenheit TemperatureUnit = "Fahrenheit"
)
