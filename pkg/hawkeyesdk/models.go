package hawkeyesdk

import (
	"encoding/json"
	"fmt"
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
	for dt := DEFAULT; dt <= FINAL_INVOICE; dt++ {
		if strings.EqualFold(s, dt.String()) {
			*d = dt
			return nil
		}
	}
	return fmt.Errorf("invalid document type: %s", s)
}

// The below applies only to Admin API users
type AdminClaim struct {
	CustomerName             string     `json:"customername,omitempty"`
	CustomerAddress          string     `json:"customeraddress,omitempty"`
	SearchInfo               string     `json:"searchinfo,omitempty"`
	ID                       int        `json:"id,omitempty"`
	Filenumber               int        `json:"filenumber,omitempty"`
	CustomerID               int        `json:"customerid,omitempty"`
	ClientClaimNo            string     `json:"clientclaimno,omitempty"`
	VehYear                  int        `json:"vehyear,omitempty"`
	VehMake                  string     `json:"vehmake,omitempty"`
	VehModel                 string     `json:"vehmodel,omitempty"`
	VehEdition               string     `json:"vehedition,omitempty"`
	Color                    string     `json:"color,omitempty"`
	VIN                      string     `json:"vin,omitempty"`
	RenterName               string     `json:"rentername,omitempty"`
	RenterPhone              string     `json:"renterphone,omitempty"`
	RenterEmail              string     `json:"renteremail,omitempty"`
	RenterAddress1           string     `json:"renteraddress1,omitempty"`
	RenterAddress2           string     `json:"renteraddress2,omitempty"`
	RenterCity               string     `json:"rentercity,omitempty"`
	RenterState              string     `json:"renterstate,omitempty"`
	RenterZip                string     `json:"renterzip,omitempty"`
	RenterPhone2             string     `json:"renterphone2,omitempty"`
	PolicyNumber             string     `json:"policynumber,omitempty"`
	DriverName               string     `json:"drivername,omitempty"`
	InsuranceCompany         string     `json:"insurancecompany,omitempty"`
	ClaimNumber              string     `json:"claimnumber,omitempty"`
	Adjuster                 string     `json:"adjuster,omitempty"`
	AdjusterPhone            string     `json:"adjusterphone,omitempty"`
	AdjEmail                 string     `json:"adjemail,omitempty"`
	AdjFax                   string     `json:"adjfax,omitempty"`
	DateOfLoss               string     `json:"dateofloss,omitempty"`
	DriverPhone              string     `json:"driverphone,omitempty"`
	DriverEmail              string     `json:"driveremail,omitempty"`
	PhysDamPrice             float32    `json:"physdamprice,omitempty"`
	LossDescription          string     `json:"lossdescription,omitempty"`
	DamageDescription        string     `json:"damagedescription,omitempty"`
	TeamLeaderAdjID          int        `json:"teamleader_adjid,omitempty"`
	TODO                     string     `json:"todo,omitempty"`
	VehMileage               int        `json:"vehmileage,omitempty"`
	LaborHours               float32    `json:"laborhours,omitempty"`
	DamageModifier           float32    `json:"damagemodifier,omitempty"`
	DailyRent                float32    `json:"dailyrent,omitempty"`
	VirtualAssID             int        `json:"virtualassid,omitempty"`
	LossType                 string     `json:"losstype,omitempty"`
	ReserveCategory          string     `json:"reservecategory,omitempty"`
	CatastropheDesc          string     `json:"catastrophedesc,omitempty"`
	Catastrophe              string     `json:"catastrophe,omitempty"`
	ReserveAmount            float32    `json:"reserveamount,omitempty"`
	ReserveAmount2           float32    `json:"reserveamount2,omitempty"`
	ReserveAmount3           float32    `json:"reserveamount3,omitempty"`
	ReserveAmount4           float32    `json:"reserveamount4,omitempty"`
	PolicyRequested          bool       `json:"policyrequested,omitempty"`
	PolicyReceived           bool       `json:"policyreceived,omitempty"`
	AOB                      bool       `json:"aob,omitempty"`
	PoliceReportReceived     bool       `json:"policereportreceived,omitempty"`
	RentalAgreement          bool       `json:"rentalagreement,omitempty"`
	POA                      bool       `json:"poa,omitempty"`
	Photos                   bool       `json:"photos,omitempty"`
	Estimate                 bool       `json:"estimate,omitempty"`
	LOU                      bool       `json:"lou,omitempty"`
	DV                       bool       `json:"dv,omitempty"`
	Demand                   bool       `json:"demand,omitempty"`
	ClaimPaid                bool       `json:"claimpaid,omitempty"`
	InvoicePaid              bool       `json:"invoicepaid,omitempty"`
	InvSubmitted             bool       `json:"invsubmitted,omitempty"`
	TotalLoss                bool       `json:"totalloss,omitempty"`
	FirstParty               bool       `json:"firstparty,omitempty"`
	ThirdParty               bool       `json:"thirdparty,omitempty"`
	CDW                      bool       `json:"cdw,omitempty"`
	CashCheck                bool       `json:"cashcheck,omitempty"`
	PoliceReportNumber       string     `json:"policereportnumber,omitempty"`
	ReportingAgency          string     `json:"reportingagency,omitempty"`
	EstimateAmount           float32    `json:"estimateamount,omitempty"`
	LossOfUseAmount          float32    `json:"lossofuseamnt,omitempty"`
	DVAmount                 float32    `json:"dv_amnt,omitempty"`
	LiabilityAccepted        string     `json:"liabilityaccepted,omitempty"`
	LiabilityDenied          string     `json:"liabilitydenied,omitempty"`
	SettlementOffer          float32    `json:"settlementoffer,omitempty"`
	SettlementCalcPDSupD     float32    `json:"settlementcalcpdsupd,omitempty"`
	SettlementPD             float32    `json:"settlement_pd,omitempty"`
	SettlementSalvage        float32    `json:"settlement_salvage,omitempty"`
	SettlementLOU            float32    `json:"settlement_lou,omitempty"`
	SettlementDV             float32    `json:"settlement_dv,omitempty"`
	SettlementOther          float32    `json:"settlement_other,omitempty"`
	SettlementDeductible     float32    `json:"settlement_deductable,omitempty"`
	Supplement               float32    `json:"supplement,omitempty"`
	ReceivedVia              string     `json:"receivedvia,omitempty"`
	DateReceived             string     `json:"datereceived,omitempty"`
	RecordDate               string     `json:"recorddate,omitempty"`
	AppraiserID              int        `json:"appraiserid,omitempty"`
	InsuredName              string     `json:"insuredname,omitempty"`
	InsdAddress1             string     `json:"insdaddress1,omitempty"`
	InsdAddress2             string     `json:"insdaddress2,omitempty"`
	InsdCity                 string     `json:"insdcity,omitempty"`
	InsdState                string     `json:"insdstate,omitempty"`
	InsdZip                  string     `json:"insdzip,omitempty"`
	InsdPhone                string     `json:"insdphone,omitempty"`
	InsdPhone2               string     `json:"insdphone2,omitempty"`
	InsdEmail                string     `json:"insdemail,omitempty"`
	Risk                     string     `json:"risk,omitempty"`
	RiskLocName              string     `json:"risklocname,omitempty"`
	RiskAddress              string     `json:"riskaddress,omitempty"`
	RiskCity                 string     `json:"riskcity,omitempty"`
	RiskState                string     `json:"riskstate,omitempty"`
	RiskZip                  string     `json:"riskzip,omitempty"`
	RiskContact              string     `json:"riskcontact,omitempty"`
	RiskPhone                string     `json:"riskphone,omitempty"`
	ClmtName                 string     `json:"clmtname,omitempty"`
	ClmtAddress1             string     `json:"clmtaddress1,omitempty"`
	ClmtAddress2             string     `json:"clmtaddress2,omitempty"`
	ClmtCity                 string     `json:"clmtcity,omitempty"`
	ClmtState                string     `json:"clmtstate,omitempty"`
	ClmtZip                  string     `json:"clmtzip,omitempty"`
	ClmtPhone                string     `json:"clmtphone,omitempty"`
	ClmtPhone2               string     `json:"clmtphone2,omitempty"`
	ClmtEmail                string     `json:"clmtemail,omitempty"`
	DateRptDue               string     `json:"daterptdue,omitempty"`
	NextStatusDue            string     `json:"nextstatusdue,omitempty"`
	AdjusterDiaryDate        string     `json:"adjusterdiarydate,omitempty"`
	DaysUntilRptDue          int        `json:"daysuntilrptdue,omitempty"`
	HCAdjID                  int        `json:"hc_adjid,omitempty"`
	AssistAdjID              int        `json:"assist_adjid,omitempty"`
	AmtInv                   float32    `json:"amt_inv,omitempty"`
	HCAdjuster               string     `json:"hcadjuster,omitempty"`
	HCAjusterEmail           string     `json:"hcajusteremail,omitempty"`
	HCAssistantAdjuster      string     `json:"hcassistantadjuster,omitempty"`
	Appraiser                string     `json:"appraiser,omitempty"`
	AppraiserDeskStandardFee float32    `json:"appraiserdeskstandardfee,omitempty"`
	AppraiserDeskExoticFee   float32    `json:"appraiserdeskexoticfee,omitempty"`
	PlateNumber              string     `json:"platenumber,omitempty"`
	UnitNumber               string     `json:"unitnumber,omitempty"`
	AckEmailDateSent         string     `json:"ackemaildatesent,omitempty"`
	InterimSubmittedAmt      float32    `json:"interimsubmittedamt,omitempty"`
	InterimInvoiceAmt        float32    `json:"interiminvoiceamt,omitempty"`
	InspectionDate           string     `json:"inspectiondate,omitempty"`
	AdministrativeFee        float32    `json:"administrativefee,omitempty"`
	AppraisalFee             float32    `json:"appraisalfee,omitempty"`
	DateFileClosed           string     `json:"datefileclosed,omitempty"`
	SettDamageDeposit        float32    `json:"settdamagedeposit,omitempty"`
	DemandAdminFee           float32    `json:"demand_admin_fee,omitempty"`
	DemandAppraisalFee       float32    `json:"demand_appraisal_fee,omitempty"`
	BusinessPhone            string     `json:"businessphone,omitempty"`
	HomePhone                string     `json:"homephone,omitempty"`
	MobilePhone              string     `json:"mobilephone,omitempty"`
	FaxNumber                string     `json:"faxnumber,omitempty"`
	ClaimType                string     `json:"claimtype,omitempty"`
	CustomerEmail1           string     `json:"customeremail1,omitempty"`
	CustomerEmail2           string     `json:"customeremail2,omitempty"`
	AckEmails                string     `json:"ackemails,omitempty"`
	ClientHourlyRate         float32    `json:"clienthourlyrate,omitempty"`
	ClaimRate                float32    `json:"claimrate,omitempty"`
	HideByDefault            int        `json:"hidebydefault,omitempty"`
	IsFlat                   bool       `json:"isflat,omitempty"`
	InsCheckReceived         bool       `json:"inscheckreceived,omitempty"`
	ClaimDuration            int        `json:"claimduration,omitempty"`
	ClaimStatusID            int        `json:"claimstatusid,omitempty"`
	InsClaim                 string     `json:"insclaim,omitempty"`
	ClaimStatusName          string     `json:"claimstatusname,omitempty"`
	InspNotNeeded            bool       `json:"inspnotneeded,omitempty"`
	ACV                      float32    `json:"acv,omitempty"`
	SalvageQuote             float32    `json:"salvagequote,omitempty"`
	RANumber                 string     `json:"ranumber,omitempty"`
	SettlementTowing         float32    `json:"settlementtowing,omitempty"`
	SettlementStorage        float32    `json:"settlementstorage,omitempty"`
	Towing                   float32    `json:"towing,omitempty"`
	Storage                  float32    `json:"storage,omitempty"`
	SalesRepName             string     `json:"salesrepname,omitempty"`
	Locked                   bool       `json:"locked,omitempty"`
	SettlementTotalLoss      float32    `json:"settlement_totalloss,omitempty"`
	DmgDepCollected          float32    `json:"dmgdepcollected,omitempty"`
	Deductible               float32    `json:"deductible,omitempty"`
	InvNotes                 string     `json:"invnotes,omitempty"`
	TowingStorage            string     `json:"towingstorage,omitempty"`
	Ownership                string     `json:"ownership,omitempty"`
	PoliceFire               string     `json:"policefire,omitempty"`
	Salvage                  string     `json:"salvage,omitempty"`
	UseOfExpert              string     `json:"useofexpert,omitempty"`
	BodilyInjury             string     `json:"bodilyinjury,omitempty"`
	DocFiles                 []DocFile  `json:"docfiles,omitempty"`
	LogTrail                 []LogTrail `json:"logtrail,omitempty"`
	PaymentMethod            string     `json:"paymentmethod,omitempty"`
	OpenAgreementDate        string     `json:"openagreementdate,omitempty"`
	ClosedAgreementDate      string     `json:"closedagreementdate,omitempty"`
	ReportBeforeRentalDate   string     `json:"reportbeforerentaldate,omitempty"`
	ReportAfterRentalDate    string     `json:"reportafterrentaldate,omitempty"`
	ReportDate               string     `json:"reportdate,omitempty"`
	EndRentalPeriodDate      string     `json:"endrentalperioddate,omitempty"`
	StartRentalPeriodDate    string     `json:"startrentalperioddate,omitempty"`
	BirthYear                int        `json:"birthyear,omitempty"`
	HandlingStartDate        string     `json:"handlingstartdate,omitempty"`
	DemandDate               string     `json:"demanddate,omitempty"`
	PolicyStartDate          string     `json:"policystartdate,omitempty"`
	PolicyEndDate            string     `json:"policyenddate,omitempty"`
	VehicleOwner             string     `json:"vehicleowner,omitempty"`
}
