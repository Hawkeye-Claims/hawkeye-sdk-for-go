package hawkeyesdk

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
	LiabilityAccepted    bool       `json:"liabilityaccepted,omitempty"`
	LiabilityDenied      bool       `json:"liabilitydenied,omitempty"`
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
		return "Driver's License"
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
		return "Vendor Invoice"
	case INTERIM_INVOICE:
		return "Interim Invoice"
	case FINAL_INVOICE:
		return "Final Invoice"
	default:
		return "Uncategorized API Document"
	}
}
