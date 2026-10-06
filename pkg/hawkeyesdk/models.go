package hawkeyesdk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ApiResponse struct {
	Filenumber int    `json:"filenumber"`
	Message    string `json:"message"`
	Error      int    `json:"error"`
	Success    bool   `json:"success"`
}

type DocType int

const (
	DEFAULT DocType = iota
	FIRST_REPORT
	SECOND_REPORT
	THIRD_REPORT
	ACKNOWLEDGEMENT
	AOB
	ASSIGNMENT_SHEET
	BILL
	BILL_OF_LADING
	CALL_RECORDING
	CASH_CALL
	CHECK_IN_VIDEO
	CHECK_OUT_VIDEO
	CONDITION_REPORT
	CLAIM_DETAILS_REPORT
	CLAIM_STATUS_REPORT
	DAMAGE_ASSESSMENT
	DEDUCTIBLE_REQUEST_FINAL_NOTICE
	DEDUCTIBLE_REQUEST_FIRST_NOTICE
	DELIVERY_CONFIRMATION
	DEMAND
	DEMAND_LETTER
	DENIAL_LETTER
	DRIVER_EXCHANGE
	DRIVERS_LICENSE
	DV_FORM
	EMAIL
	EXPENSE_RECEIPT
	HC_DAMAGE_APPRAISAL
	IMAGES
	INCIDENT_REPORT
	INSURANCE_CARD
	INVOICE
	LIENHOLDER_INFO
	MARKET_VALUATION
	MITIGATION_LETTER
	NON_HC_DAMAGE_APPRAISAL
	OTHER
	PAYMENT_ADVISORY_LETTER
	PAYMENT_CONFIRMATION
	POLICE_REPORT
	POLICY
	POA
	RECORDED_STATEMENT
	REGISTRATION
	RELEASE
	RENTAL_AGREEMENT
	RESERVE_REPORT
	SETTLEMENT_CHECK
	STATUS_REPORT
	TITLE
	TOW_BILL
	TRAILER_INTERCHANGE_AGREEMENT
	VEHICLE_HISTORY
	VEHICLE_SPECIFICATIONS
	VENDOR_INVOICE
	INTERIM_INVOICE
	FINAL_INVOICE
	INCIDENT_REPORT_ACORD

	// docTypeCount bounds name lookups; keep it last.
	docTypeCount
)

type DocFile struct {
	Doctype   DocType `json:"doctype"`
	DateAdded string  `json:"dateadded"`
	User      string  `json:"user"`
	Notes     *string `json:"notes,omitempty"`
	Filename  string  `json:"filename"`
}

type LogTrail struct {
	Date     string `json:"date"`
	Activity string `json:"activity"`
	User     string `json:"user"`
}

type LoggedTime struct {
	Date       string  `json:"date"`
	Filenumber int     `json:"filenumber"`
	User       string  `json:"user"`
	Email      string  `json:"email"`
	Time       float32 `json:"loggedtime"`
}

type InsCompany struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Probability int    `json:"probability,omitempty"`
}

type Claim struct {
	Filenumber           int        `json:"filenumber,omitempty"`
	CustomerName         string     `json:"customername,omitempty"`
	ClientClaimNo        string     `json:"clientclaimno,omitempty"`
	RenterName           string     `json:"rentername,omitempty"`
	RANumber             string     `json:"ranumber,omitempty"`
	InsuredName          string     `json:"insuredname,omitempty"`
	InsuranceCompany     string     `json:"insurancecompany,omitempty"`
	ClaimNumber          string     `json:"claimnumber,omitempty"`
	PolicyNumber         string     `json:"policynumber,omitempty"`
	DateOfLoss           string     `json:"dateofloss,omitempty"`
	Adjuster             string     `json:"adjuster,omitempty"`
	AdjusterPhone        string     `json:"adjusterphone,omitempty"`
	FirstParty           bool       `json:"firstparty,omitempty"`
	ThirdParty           bool       `json:"thirdparty,omitempty"`
	CDW                  bool       `json:"cdw,omitempty"`
	HCAdj                string     `json:"hc_adj,omitempty"`
	OfficePhone          string     `json:"officephone,omitempty"`
	Email                string     `json:"email,omitempty"`
	VIN                  string     `json:"vin,omitempty"`
	VehYear              int        `json:"vehyear,omitempty"`
	VehMake              string     `json:"vehmake,omitempty"`
	VehModel             string     `json:"vehmodel,omitempty"`
	VehEdition           string     `json:"vehedition,omitempty"`
	Color                string     `json:"color,omitempty"`
	PlateNumber          string     `json:"platenumber,omitempty"`
	UnitNumber           string     `json:"unitnumber,omitempty"`
	InspectionDate       string     `json:"inspectiondate,omitempty"`
	EstimateAmount       float32    `json:"estimateamount,omitempty"`
	TotalLoss            bool       `json:"totalloss,omitempty"`
	ContinuedRentalAmt   float32    `json:"continuedrentalamt,omitempty"`
	DVAmt                float32    `json:"dv_amt,omitempty"`
	LiabilityAccepted    string     `json:"liabilityaccepted,omitempty"`
	LiabilityDenied      string     `json:"liabilitydenied,omitempty"`
	SettlementPD         float32    `json:"settlement_pd,omitempty"`
	SettlementSalvage    float32    `json:"settlement_salvage,omitempty"`
	SettlementCR         float32    `json:"settlement_cr,omitempty"`
	SettlementDV         float32    `json:"settlement_dv,omitempty"`
	SettlementOther      float32    `json:"settlement_other,omitempty"`
	SettlementDeductible float32    `json:"settlement_deductable,omitempty"`
	AdministrativeFee    float32    `json:"administrativefee,omitempty"`
	AppraisalFee         float32    `json:"appraisalfee,omitempty"`
	DateFileClosed       string     `json:"datefileclosed,omitempty"`
	SettlementOffer      float32    `json:"settlementoffer,omitempty"`
	Supplement           float32    `json:"supplement,omitempty"`
	SettlementTowing     float32    `json:"settlementtowing,omitempty"`
	SettlementStorage    float32    `json:"settlementstorage,omitempty"`
	DemandAdminFee       float32    `json:"demand_admin_fee,omitempty"`
	DemandAppraisalFee   float32    `json:"demand_appraisal_fee,omitempty"`
	EstimatedDate        string     `json:"estimateddate,omitempty"`
	DemandDate           string     `json:"demandate,omitempty"`
	PolicyStartDate      string     `json:"policystartdate,omitempty"`
	PolicyEndDate        string     `json:"policyenddate,omitempty"`
	VehicleOwner         string     `json:"vehicleowner,omitempty"`
	DocFiles             []DocFile  `json:"docfiles,omitempty"`
	LogTrail             []LogTrail `json:"logtrail,omitempty"`
}

func (d DocType) String() string {
	switch d {
	case DEFAULT:
		return "Uncategorized API Document"
	case FIRST_REPORT:
		return "1st Report"
	case SECOND_REPORT:
		return "2nd Report"
	case THIRD_REPORT:
		return "3rd Report"
	case ACKNOWLEDGEMENT:
		return "Acknowledgement"
	case AOB:
		return "Assignment of Benefits"
	case ASSIGNMENT_SHEET:
		return "Assignment Sheet"
	case BILL:
		return "Bill"
	case BILL_OF_LADING:
		return "Bill of Lading"
	case CALL_RECORDING:
		return "Call Recording"
	case CASH_CALL:
		return "Cash Call"
	case CHECK_IN_VIDEO:
		return "Check-in Video (Drop-Off)"
	case CHECK_OUT_VIDEO:
		return "Check-out Video (Pick up)"
	case CONDITION_REPORT:
		return "Condition Report"
	case CLAIM_DETAILS_REPORT:
		return "Claim Details Report"
	case CLAIM_STATUS_REPORT:
		return "Claim Status Report"
	case DAMAGE_ASSESSMENT:
		return "Damage Assessment"
	case DEDUCTIBLE_REQUEST_FINAL_NOTICE:
		return "Deductible Request Final Notice"
	case DEDUCTIBLE_REQUEST_FIRST_NOTICE:
		return "Deductible Request First Notice"
	case DELIVERY_CONFIRMATION:
		return "Delivery Confirmation"
	case DEMAND:
		return "Demand"
	case DEMAND_LETTER:
		return "Demand Letter"
	case DENIAL_LETTER:
		return "Denial Letter"
	case DRIVER_EXCHANGE:
		return "Driver Exchange"
	case DRIVERS_LICENSE:
		return "Drivers License"
	case DV_FORM:
		return "DV Form"
	case EMAIL:
		return "Email"
	case EXPENSE_RECEIPT:
		return "Expense Receipt"
	case HC_DAMAGE_APPRAISAL:
		return "HC Damage Appraisal"
	case IMAGES:
		return "Images"
	case INCIDENT_REPORT:
		return "Incident Report"
	case INSURANCE_CARD:
		return "Insurance Card"
	case INVOICE:
		return "Invoice"
	case LIENHOLDER_INFO:
		return "Lienholder Info"
	case MARKET_VALUATION:
		return "Market Valuation"
	case MITIGATION_LETTER:
		return "Mitigation Letter"
	case NON_HC_DAMAGE_APPRAISAL:
		return "Non-HC Damage Appraisal"
	case OTHER:
		return "Other"
	case PAYMENT_ADVISORY_LETTER:
		return "Payment Advisory Letter"
	case PAYMENT_CONFIRMATION:
		return "Payment Confirmation"
	case POLICE_REPORT:
		return "Police Report"
	case POLICY:
		return "Policy"
	case POA:
		return "Power of Attorney"
	case RECORDED_STATEMENT:
		return "Recorded Statement"
	case REGISTRATION:
		return "Registration"
	case RELEASE:
		return "Release"
	case RENTAL_AGREEMENT:
		return "Rental Agreement"
	case RESERVE_REPORT:
		return "Reserve Report"
	case SETTLEMENT_CHECK:
		return "Settlement Check"
	case STATUS_REPORT:
		return "Status Report"
	case TITLE:
		return "Title"
	case TOW_BILL:
		return "Tow Bill"
	case TRAILER_INTERCHANGE_AGREEMENT:
		return "Trailer Interchange Agreement"
	case VEHICLE_HISTORY:
		return "Vehicle History"
	case VEHICLE_SPECIFICATIONS:
		return "Vehicle Specifications"
	case VENDOR_INVOICE:
		return "Vendor Inv"
	case INTERIM_INVOICE:
		return "Interim Invoice"
	case FINAL_INVOICE:
		return "Final Invoice"
	case INCIDENT_REPORT_ACORD:
		return "Incident Report (ACORD)"
	default:
		return "Uncategorized API Document"
	}
}

func (d *DocType) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*d = DocType(i)
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	s = strings.TrimSpace(s)
	for dt := range docTypeCount {
		if strings.EqualFold(s, dt.String()) {
			*d = dt
			return nil
		}
	}
	// Hawk adds document categories without notice. An unrecognized name
	// must not prevent decoding the claim, so it falls back to DEFAULT.
	*d = DEFAULT
	return nil
}

func parseSanitizedInt(data []byte) (int, error) {
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		return i, nil
	}

	// A mistyped period can reach us either inside a string ("12.345.678")
	// or as a JSON number (45.000). Both are sanitized the same way.
	var s string
	var n json.Number
	if err := json.Unmarshal(data, &s); err != nil {
		if numberErr := json.Unmarshal(data, &n); numberErr != nil {
			return 0, err
		}
		s = n.String()
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}

	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("could not parse integer from %q", s)
	}

	return i, nil
}

func (a *AdminClaim) UnmarshalJSON(data []byte) error {
	type adminClaimAlias AdminClaim

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	vehMileageRaw, hasVehMileage := raw["vehmileage"]
	if hasVehMileage {
		delete(raw, "vehmileage")
	}

	normalizedData, err := json.Marshal(raw)
	if err != nil {
		return err
	}

	var alias adminClaimAlias
	if err := json.Unmarshal(normalizedData, &alias); err != nil {
		return err
	}

	*a = AdminClaim(alias)

	if hasVehMileage {
		vehMileage, err := parseSanitizedInt(vehMileageRaw)
		if err != nil {
			return fmt.Errorf("invalid vehmileage: %w", err)
		}
		a.VehMileage = vehMileage
	}

	return nil
}

// AdminClaimGeneralSection identifies the claim and tracks its workflow state.
type AdminClaimGeneralSection struct {
	CustomerID            int        `json:"customerid,omitempty"`
	CustomerName          string     `json:"customername,omitempty"`
	CustomerAddress       string     `json:"customeraddress,omitempty"`
	CustomerEmail1        string     `json:"customeremail1,omitempty"`
	CustomerEmail2        string     `json:"customeremail2,omitempty"`
	SalesRepName          string     `json:"salesrepname,omitempty"`
	ClientClaimNo         string     `json:"clientclaimno,omitempty"`
	ClaimType             string     `json:"claimtype,omitempty"`
	SearchInfo            string     `json:"searchinfo,omitempty"`
	ID                    int        `json:"id,omitempty"`
	Filenumber            int        `json:"filenumber,omitempty"`
	MasterClaimNumber     string     `json:"masterclaimnumber,omitempty"`
	ReceivedVia           string     `json:"receivedvia,omitempty"`
	DateReceived          string     `json:"datereceived,omitempty"`
	RecordDate            string     `json:"recorddate,omitempty"`
	DateFileClosed        string     `json:"datefileclosed,omitempty"`
	ClaimDuration         int        `json:"claimduration,omitempty"`
	ClaimStatusID         int        `json:"claimstatusid,omitempty"`
	ClaimStatusName       string     `json:"claimstatusname,omitempty"`
	InsClaim              string     `json:"insclaim,omitempty"`
	Status                string     `json:"gaugestat,omitempty"`
	GaugeStatus           string     `json:"gaugestatus,omitempty"`
	GaugeGoal             string     `json:"gaugegoal,omitempty"`
	HideByDefault         int        `json:"hidebydefault,omitempty"`
	Locked                bool       `json:"locked,omitempty"`
	Protocol              string     `json:"protocol,omitempty"`
	SubmittedBy           string     `json:"submittedby,omitempty"`
	AckEmails             string     `json:"ackemails,omitempty"`
	AckEmailDateSent      string     `json:"ackemaildatesent,omitempty"`
	BusinessPhone         string     `json:"businessphone,omitempty"`
	HomePhone             string     `json:"homephone,omitempty"`
	MobilePhone           string     `json:"mobilephone,omitempty"`
	FaxNumber             string     `json:"faxnumber,omitempty"`
	NumberOfStatusEntries int        `json:"numberofstatusentries,omitempty"`
	LastStatusDate        string     `json:"laststatusdate,omitempty"`
	DaysSinceLastStatus   int        `json:"dayssincelaststatus,omitempty"`
	StatusLog             string     `json:"statuslog,omitempty"`
	DocsMissing           int        `json:"docsmissing,omitempty"`
	DocFiles              []DocFile  `json:"docfiles,omitempty"`
	LogTrail              []LogTrail `json:"logtrail,omitempty"`
	PaymentMethod         string     `json:"paymentmethod,omitempty"`
}

type AdminClaimVehicleSection struct {
	VehYear      int    `json:"vehyear,omitempty"`
	VehMake      string `json:"vehmake,omitempty"`
	VehModel     string `json:"vehmodel,omitempty"`
	VehEdition   string `json:"vehedition,omitempty"`
	VehType      string `json:"vehtype,omitempty"`
	VehMileage   int    `json:"vehmileage,omitempty"`
	Color        string `json:"color,omitempty"`
	VIN          string `json:"vin,omitempty"`
	PlateNumber  string `json:"platenumber,omitempty"`
	UnitNumber   string `json:"unitnumber,omitempty"`
	VehicleOwner string `json:"vehicleowner,omitempty"`
}

// AdminClaimInsuranceSection covers the policy, the parties, and the coverage terms.
type AdminClaimInsuranceSection struct {
	InsuranceCompany           string  `json:"insurancecompany,omitempty"`
	InsCompaniesID             []int   `json:"inscompaniesid,omitempty"`
	ClaimNumber                string  `json:"claimnumber,omitempty"`
	PolicyNumber               string  `json:"policynumber,omitempty"`
	PolicyRequested            bool    `json:"policyrequested,omitempty"`
	PolicyReceived             bool    `json:"policyreceived,omitempty"`
	PolicyStartDate            string  `json:"policystartdate,omitempty"`
	PolicyEndDate              string  `json:"policyenddate,omitempty"`
	InitialPolicyStartDate     string  `json:"initialpolicystartdate,omitempty"`
	SubPolicyNumber            string  `json:"subpolicynumber,omitempty"`
	SubPolicyInceptionDate     string  `json:"subpolicyinceptiondate,omitempty"`
	SubPolicyExpiryDate        string  `json:"subpolicyexpirydate,omitempty"`
	IsTPA                      bool    `json:"istpa,omitempty"`
	PriorTPAName               string  `json:"priortpaname,omitempty"`
	PriorClaimReference        string  `json:"priorclaimreference,omitempty"`
	LiabilityAccepted          string  `json:"liabilityaccepted,omitempty"`
	LiabilityDenied            string  `json:"liabilitydenied,omitempty"`
	DamageDescription          string  `json:"damagedescription,omitempty"`
	LossDescription            string  `json:"lossdescription,omitempty"`
	CatastropheDesc            string  `json:"catastrophedesc,omitempty"`
	Catastrophe                string  `json:"catastrophe,omitempty"`
	SeverityOfLoss             string  `json:"severityofloss,omitempty"`
	AOB                        bool    `json:"aob,omitempty"`
	POA                        bool    `json:"poa,omitempty"`
	FirstParty                 bool    `json:"firstparty,omitempty"`
	ThirdParty                 bool    `json:"thirdparty,omitempty"`
	CDW                        bool    `json:"cdw,omitempty"`
	Adjuster                   string  `json:"adjuster,omitempty"`
	AdjusterPhone              string  `json:"adjusterphone,omitempty"`
	AdjEmail                   string  `json:"adjemail,omitempty"`
	AdjFax                     string  `json:"adjfax,omitempty"`
	InsuredName                string  `json:"insuredname,omitempty"`
	InsdAddress1               string  `json:"insdaddress1,omitempty"`
	InsdAddress2               string  `json:"insdaddress2,omitempty"`
	InsdCity                   string  `json:"insdcity,omitempty"`
	InsdState                  string  `json:"insdstate,omitempty"`
	InsdZip                    string  `json:"insdzip,omitempty"`
	InsdPhone                  string  `json:"insdphone,omitempty"`
	InsdPhone2                 string  `json:"insdphone2,omitempty"`
	InsdEmail                  string  `json:"insdemail,omitempty"`
	BirthYear                  int     `json:"birthyear,omitempty"`
	DriverName                 string  `json:"drivername,omitempty"`
	DriverPhone                string  `json:"driverphone,omitempty"`
	DriverEmail                string  `json:"driveremail,omitempty"`
	ClmtName                   string  `json:"clmtname,omitempty"`
	ClmtAddress1               string  `json:"clmtaddress1,omitempty"`
	ClmtAddress2               string  `json:"clmtaddress2,omitempty"`
	ClmtCity                   string  `json:"clmtcity,omitempty"`
	ClmtState                  string  `json:"clmtstate,omitempty"`
	ClmtZip                    string  `json:"clmtzip,omitempty"`
	ClmtPhone                  string  `json:"clmtphone,omitempty"`
	ClmtPhone2                 string  `json:"clmtphone2,omitempty"`
	ClmtEmail                  string  `json:"clmtemail,omitempty"`
	Risk                       string  `json:"risk,omitempty"`
	RiskLocName                string  `json:"risklocname,omitempty"`
	RiskAddress                string  `json:"riskaddress,omitempty"`
	RiskCity                   string  `json:"riskcity,omitempty"`
	RiskState                  string  `json:"riskstate,omitempty"`
	RiskZip                    string  `json:"riskzip,omitempty"`
	RiskContact                string  `json:"riskcontact,omitempty"`
	RickContact                string  `json:"rickcontact,omitempty"`
	RiskPhone                  string  `json:"riskphone,omitempty"`
	TypeOfInsurance            string  `json:"typeofinsurance,omitempty"`
	DeductibleBasis            string  `json:"deductiblebasis,omitempty"`
	InsuredAmount              float32 `json:"insuredamount,omitempty"`
	PercentCeded               float32 `json:"percentceded,omitempty"`
	InsuredCountry             string  `json:"insuredcountry,omitempty"`
	BodilyInjuryPerPerson      float32 `json:"bodily_injury_per_person,omitempty"`
	BodilyInjuryPerLoss        float32 `json:"bodily_injury_per_loss,omitempty"`
	LiabilityPD                float32 `json:"liability_pd,omitempty"`
	FirstPartyPD               float32 `json:"first_party_pd,omitempty"`
	PIPLimit                   float32 `json:"pip_limit,omitempty"`
	DisclaimerOfCoverageIssued bool    `json:"disclaimerofcoverageissued,omitempty"`
	DisclaimerOfCoverageReason string  `json:"disclaimerofcoveragereason,omitempty"`
	ReferToUnderwriters        string  `json:"refertounderwriters,omitempty"`
	ReferredToSubrogation      bool    `json:"referredtosubrogation,omitempty"`
	RiskCodeID                 int     `json:"riskcodeid,omitempty"`
	IBCCode                    string  `json:"ibccode,omitempty"`
	IBCSegment                 string  `json:"ibcsegment,omitempty"`
	OSFILOB                    string  `json:"osfilob,omitempty"`
	Product                    string  `json:"product,omitempty"`
	Trade                      string  `json:"trade,omitempty"`
	OriginalCurrency           string  `json:"originalcurrency,omitempty"`
	SettlementCurrency         string  `json:"settlementcurrency,omitempty"`
	ExchangeRate               float32 `json:"exchangerate,omitempty"`
	RiskInceptionDate          string  `json:"riskinceptiondate,omitempty"`
	RiskExpiryDate             string  `json:"riskexpirydate,omitempty"`
	LocationOfRiskLocationID   int     `json:"locationofrisklocationid,omitempty"`
	LocationOfRiskCountry      string  `json:"locationofriskcountry,omitempty"`
	StateOfFiling              string  `json:"stateoffiling,omitempty"`
	PeriodOfCoverNarrative     string  `json:"periodofcovernarrative,omitempty"`
	PCSCode                    string  `json:"pcscode,omitempty"`
}

type AdminClaimRentalSection struct {
	RenterName             string  `json:"rentername,omitempty"`
	RenterPhone            string  `json:"renterphone,omitempty"`
	RenterEmail            string  `json:"renteremail,omitempty"`
	RenterAddress1         string  `json:"renteraddress1,omitempty"`
	RenterAddress2         string  `json:"renteraddress2,omitempty"`
	RenterCity             string  `json:"rentercity,omitempty"`
	RenterState            string  `json:"renterstate,omitempty"`
	RenterZip              string  `json:"renterzip,omitempty"`
	RenterPhone2           string  `json:"renterphone2,omitempty"`
	RentalAgreement        bool    `json:"rentalagreement,omitempty"`
	RANumber               string  `json:"ranumber,omitempty"`
	DailyRent              float32 `json:"dailyrent,omitempty"`
	OpenAgreementDate      string  `json:"openagreementdate,omitempty"`
	ClosedAgreementDate    string  `json:"closedagreementdate,omitempty"`
	ReportBeforeRentalDate string  `json:"reportbeforerentaldate,omitempty"`
	ReportAfterRentalDate  string  `json:"reportafterrentaldate,omitempty"`
	ReportDate             string  `json:"reportdate,omitempty"`
	StartRentalPeriodDate  string  `json:"startrentalperioddate,omitempty"`
	EndRentalPeriodDate    string  `json:"endrentalperioddate,omitempty"`
}

type AdminClaimHCSection struct {
	HCAdjuster               string  `json:"hcadjuster,omitempty"`
	HCAdjID                  int     `json:"hc_adjid,omitempty"`
	HCAjusterEmail           string  `json:"hcadjusteremail,omitempty"`
	TeamLeaderAdjID          int     `json:"teamleader_adjid,omitempty"`
	HCAssistantAdjuster      string  `json:"hcassistantadjuster,omitempty"`
	AssistAdjID              int     `json:"assist_adjid,omitempty"`
	VirtualAsst              string  `json:"virtualasst,omitempty"`
	VirtualAssID             int     `json:"virtualasstid,omitempty"`
	Appraiser                string  `json:"appraiser,omitempty"`
	AppraiserID              int     `json:"appraiserid,omitempty"`
	AppraiserDeskStandardFee float32 `json:"appraiserdeskstandardfee,omitempty"`
	AppraiserDeskExoticFee   float32 `json:"appraiserdeskexoticfee,omitempty"`
	AdjusterDiaryDate        string  `json:"adjusterdiarydate,omitempty"`
	DateRptDue               string  `json:"daterptdue,omitempty"`
	NextStatusDue            string  `json:"nextstatusdue,omitempty"`
	DiaryDate                string  `json:"diarydate,omitempty"`
	AssistAdjusterDiaryDate  string  `json:"assistadjusterdiarydate,omitempty"`
	DaysUntilRptDue          int     `json:"daysuntilrptdue,omitempty"`
	PeerReviewDate           string  `json:"peerreviewdate,omitempty"`
	PeerReviewBy             string  `json:"peerreviewby,omitempty"`
	HandlingStartDate        string  `json:"handlingstartdate,omitempty"`
	TODO                     string  `json:"todo,omitempty"`
}

type AdminClaimDamagesSection struct {
	InspectionDate  string  `json:"inspectiondate,omitempty"`
	EstimateDate    string  `json:"estimatedate,omitempty"`
	TotalLoss       bool    `json:"totalloss,omitempty"`
	InspNotNeeded   bool    `json:"inspnotneeded,omitempty"`
	PhysDamPrice    float32 `json:"physdamprice,omitempty"`
	ACV             float32 `json:"acv,omitempty"`
	DamageModifier  float32 `json:"damagemodifier,omitempty"`
	EstimateAmount  float32 `json:"estimateamount,omitempty"`
	LaborHours      float32 `json:"laborhours,omitempty"`
	LossOfUseAmount float32 `json:"lossofuseamt,omitempty"`
	DVAmount        float32 `json:"dv_amt,omitempty"`
	SalvageQuote    float32 `json:"salvagequote,omitempty"`
	Towing          float32 `json:"towing,omitempty"`
	Storage         float32 `json:"storage,omitempty"`
	TowingStorage   string  `json:"towingstorage,omitempty"`
	DmgDepCollected float32 `json:"dmgdepcollected,omitempty"`
	BodilyInjury    float32 `json:"bodilyinjury,omitempty"`
	Deductible      float32 `json:"deductible,omitempty"`
	Ownership       string  `json:"ownership,omitempty"`
	PoliceFire      string  `json:"policefire,omitempty"`
	Salvage         string  `json:"salvage,omitempty"`
	UseOfExpert     string  `json:"useofexpert,omitempty"`
	Estimate        bool    `json:"estimate,omitempty"`
	LOU             bool    `json:"lou,omitempty"`
	DV              bool    `json:"dv,omitempty"`
	Demand          bool    `json:"demand,omitempty"`
	Photos          bool    `json:"photos,omitempty"`
	IsFlat          bool    `json:"isflat,omitempty"`
}

type AdminClaimSettlementSection struct {
	SettlementOffer           float32 `json:"settlementoffer,omitempty"`
	SettlementCalcPDSupD      float32 `json:"settlementcalcpdsupd,omitempty"`
	SettlementPD              float32 `json:"settlement_pd,omitempty"`
	SettlementSalvage         float32 `json:"settlement_salvage,omitempty"`
	SettlementLOU             float32 `json:"settlement_lou,omitempty"`
	SettlementDV              float32 `json:"settlement_dv,omitempty"`
	SettlementOther           float32 `json:"settlement_other,omitempty"`
	SettlementDeductible      float32 `json:"settlement_deductable,omitempty"`
	SettlementTotalLoss       float32 `json:"settlement_totalloss,omitempty"`
	SettlementConfirmed       bool    `json:"settlementconfirmed,omitempty"`
	DateClaimAmountAgreedWith string  `json:"dateclaimamountagreedwith,omitempty"`
	DateClaimAmountAgreed     string  `json:"dateclaimamountagreed,omitempty"`
	Supplement                float32 `json:"supplement,omitempty"`
	Supplement2               float32 `json:"supplement2,omitempty"`
	DemandOffer               float32 `json:"demandoffer,omitempty"`
	DemandAdminFee            float32 `json:"demand_admin_fee,omitempty"`
	DemandAppraisalFee        float32 `json:"demand_appraisal_fee,omitempty"`
	DemandDate                string  `json:"demanddate,omitempty"`
	SettlementTowing          float32 `json:"settlementtowing,omitempty"`
	SettlementStorage         float32 `json:"settlementstorage,omitempty"`
	SettDamageDeposit         float32 `json:"settdamagedeposit,omitempty"`
	ExGratisPayment           float32 `json:"exgratispayment,omitempty"`
	ExcessPayment             float32 `json:"excesspayment,omitempty"`
	AmtInv                    float32 `json:"amt_inv,omitempty"`
	InvoiceAmount             float32 `json:"invoiceamount,omitempty"`
	InterimSubmittedAmt       float32 `json:"interimsubmittedamt,omitempty"`
	InterimInvoiceAmt         float32 `json:"interiminvoiceamt,omitempty"`
	AdministrativeFee         float32 `json:"administrativefee,omitempty"`
	AppraisalFee              float32 `json:"appraisalfee,omitempty"`
	ClientHourlyRate          float32 `json:"clienthourlyrate,omitempty"`
	ClaimRate                 float32 `json:"claimrate,omitempty"`
	InvNotes                  string  `json:"invnotes,omitempty"`
	ClaimPaid                 bool    `json:"claimpaid,omitempty"`
	InvoicePaid               bool    `json:"invoicepaid,omitempty"`
	InvSubmitted              bool    `json:"invsubmitted,omitempty"`
	InsCheckReceived          bool    `json:"inscheckreceived,omitempty"`
	CashCheck                 bool    `json:"cashcheck,omitempty"`
}

// AdminClaimLossLocationSection describes where and when the loss happened,
// including the incident evidence collected there.
type AdminClaimLossLocationSection struct {
	LossType               string `json:"losstype,omitempty"`
	DateOfLoss             string `json:"dateofloss,omitempty"`
	DateOfLossToDate       string `json:"dateoflosstodate,omitempty"`
	AccidentDate           string `json:"accidentdate,omitempty"`
	LocationOfLossAddress  string `json:"locationoflossaddress,omitempty"`
	LocationOfLossState    string `json:"locationoflossstate,omitempty"`
	LocationOfLossZip      string `json:"locationoflosszip,omitempty"`
	LocationOfLossCountry  string `json:"locationoflosscountry,omitempty"`
	LocationOfLossCounty   string `json:"locationoflosscounty,omitempty"`
	JurisdictionOfTheClaim string `json:"jurisdictionoftheclaim,omitempty"`
	CauseOfLossCode1       string `json:"causeoflosscode1,omitempty"`
	CauseOfLossCode2       string `json:"causeoflosscode2,omitempty"`
	PoliceReportReceived   bool   `json:"policereportreceived,omitempty"`
	PoliceReportNumber     string `json:"policereportnumber,omitempty"`
	ReportingAgency        string `json:"reportingagency,omitempty"`
	FireReportNumber       string `json:"firereportnumber,omitempty"`
	FireReportReceived     bool   `json:"firereportreceived,omitempty"`
	FireReportingAgency    string `json:"firereportingagency,omitempty"`
	FireReportDate         string `json:"firereportdate,omitempty"`
}

// AdminClaimSubroDefenseSection tracks subrogation recovery and litigation defense.
type AdminClaimSubroDefenseSection struct {
	SubroReceivedDemand     bool    `json:"subro_receiveddemand,omitempty"`
	SubroEstimateAmount     float32 `json:"subro_estimateamount,omitempty"`
	SubroTotalLoss          bool    `json:"subro_totalloss,omitempty"`
	SubroACV                float32 `json:"subro_acv,omitempty"`
	SubroContinuedRentalAmt float32 `json:"subro_continuedrentalamt,omitempty"`
	SubroDVAmt              float32 `json:"subro_dv_amt,omitempty"`
	SubroTowing             float32 `json:"subro_towing,omitempty"`
	SubroStorage            float32 `json:"subro_storage,omitempty"`
	SubroSalvageQuote       float32 `json:"subro_salvagequote,omitempty"`
	SubroBodilyInjury       float32 `json:"subro_bodilyinjury,omitempty"`
	SubroDemandAdminFee     float32 `json:"subro_demand_admin_fee,omitempty"`
	SubroDemandAppraisalFee float32 `json:"subro_demand_appraisal_fee,omitempty"`
	SubroTotalDemand        float32 `json:"subro_totaldemand,omitempty"`
	DateSubrogation         string  `json:"datesubrogation,omitempty"`
	ProjectedRecovery       float32 `json:"projectedrecovery,omitempty"`
	LitigationStatus        string  `json:"litigationstatus,omitempty"`
}

// AdminClaimReserveSection holds reserve amounts by category, including fee reserves.
type AdminClaimReserveSection struct {
	ReserveCategory              string  `json:"reservecategory,omitempty"`
	ReserveAmount                float32 `json:"reserveamount,omitempty"`
	ReserveAmount2               float32 `json:"reserveamount2,omitempty"`
	ReserveAmount3               float32 `json:"reserveamount3,omitempty"`
	ReserveAmount4               float32 `json:"reserveamount4,omitempty"`
	DefenceFeeReserve            float32 `json:"defencefeereserve,omitempty"`
	AdjusterFeeReserve           float32 `json:"adjusterfeereserve,omitempty"`
	TPAFeeReserve                float32 `json:"tpafeereserve,omitempty"`
	AttorneyCoverageFeeReserve   float32 `json:"attorneycoveragefeereserve,omitempty"`
	AttorneyMonitoringFeeReserve float32 `json:"attorneymonitoringfeereserve,omitempty"`
	OtherFeeReserve              float32 `json:"otherfeereserve,omitempty"`
}

// AdminClaimMedicareSection covers medical, treatment, and heads-of-damage details.
type AdminClaimMedicareSection struct {
	MedicareUSABodilyInjury                 bool    `json:"medicareusbodilyinjury,omitempty"`
	MedicareEligibilityCheckPerformance     bool    `json:"medicareeligibilitycheckperformance,omitempty"`
	MedicareOutcomeOfEligibilityStatusCheck bool    `json:"medicareoutcomeofeligibilitystatuscheck,omitempty"`
	MedicareConditionalPayments             bool    `json:"medicareconditionalpayments,omitempty"`
	MedicareMSPComplianceServices           bool    `json:"medicaremspcomplianceservices,omitempty"`
	Plan                                    string  `json:"plan,omitempty"`
	PatientName                             string  `json:"patientname,omitempty"`
	TreatmentType                           string  `json:"treatmenttype,omitempty"`
	CountryOfTreatment                      string  `json:"countryoftreatment,omitempty"`
	DateOfTreatment                         string  `json:"dateoftreatment,omitempty"`
	DateFeesPaid                            string  `json:"datefeespaid,omitempty"`
	BodyFunctionsOrStructuresAffected       string  `json:"bodyfunctionsorstructuresaffected,omitempty"`
	HeadsOfDamagePastEconomicLoss           float32 `json:"headsofdamage_pasteconomicloss,omitempty"`
	HeadsOfDamageFutureEconomicLoss         float32 `json:"headsofdamage_futureeconomicloss,omitempty"`
	HeadsOfDamagePastMedicalHospital        float32 `json:"headsofdamage_pastmedicalhospital,omitempty"`
	HeadsOfDamageFutureMedicalHospital      float32 `json:"headsofdamage_futuremedicalhospital,omitempty"`
	HeadsOfDamageFutureCaringServices       float32 `json:"headsofdamage_futurecaringservices,omitempty"`
	HeadsOfDamageGeneralDamages             float32 `json:"headsofdamage_generaldamages,omitempty"`
	HeadsOfDamageInterest                   float32 `json:"headsofdamage_interest,omitempty"`
}

// AdminClaim is the full admin claim, composed from embedded sections.
// It decodes from the flat admin claims endpoint payload.
type AdminClaim struct {
	AdminClaimGeneralSection
	AdminClaimVehicleSection
	AdminClaimInsuranceSection
	AdminClaimRentalSection
	AdminClaimHCSection
	AdminClaimDamagesSection
	AdminClaimSettlementSection
	AdminClaimLossLocationSection
	AdminClaimSubroDefenseSection
	AdminClaimReserveSection
	AdminClaimMedicareSection
}
