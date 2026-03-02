package kxv_Orya

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kxv_Orya struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'kxv_Orya' locale
func New() locales.Translator {
	return &kxv_Orya{
		locale:                 "kxv_Orya",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ପୁସୁ", "ମାହା", "ପାଗୁ", "ହିରେ", "ବେସେ", "ଜାଟା", "ଆସାଡ଼ି", "ସ୍ରାବାଁ", "ବଦ", "ଦାସାରା", "ଦିୱି", "ପାଣ୍ଡେ"},
		monthsNarrow:           []string{"", "ପୁ", "ମା", "ପା", "ହି", "ବେ", "ଜା", "ଆ", "ସ୍ରା", "ବ", "ଦା", "ଦି", "ପା"},
		monthsWide:             []string{"", "ପୁସୁ ଲେଞ୍ଜୁ", "ମାହାକା ଲେଞ୍ଜୁ", "ପାଗୁଣି ଲେଞ୍ଜୁ", "ହିରେ ଲେଞ୍ଜୁ", "ବେସେ ଲେଞ୍ଜୁ", "ଜାଟା ଲେଞ୍ଜୁ", "ଆସାଡ଼ି ଲେଞ୍ଜୁ", "ସ୍ରାବାଁ ଲେଞ୍ଜୁ", "ବଦ ଲେଞ୍ଜୁ", "ଦାସାରା ଲେଞ୍ଜୁ", "ଦିୱିଡ଼ି ଲେଞ୍ଜୁ", "ପାଣ୍ଡେ ଲେଞ୍ଜୁ"},
		daysAbbreviated:        []string{"ଆଦି", "ସମ୍ବା", "ମାଙ୍ଗା", "ପୁଦା", "ଲାକି", "ସୁକ୍ରୁ", "ସାନି"},
		daysNarrow:             []string{"ଆ", "ସ", "ମା", "ପୁ", "ଲା", "ସୁ", "ସା"},
		daysShort:              []string{"ଆ", "ସ", "ମା", "ପୁ", "ଲା", "ସୁ", "ସାନି"},
		daysWide:               []string{"ଆଦି ୱାରା", "ସମ୍ବାରା", "ମାଙ୍ଗାଡ଼ା", "ପୁଦାରା", "ଲାକି ୱାରା", "ସୁକ୍ରୁ ୱାରା", "ସାନି ୱାରା"},
		periodsAbbreviated:     []string{"ଏ\u202fଏମ", "ପି\u202fଏମ"},
		periodsNarrow:          []string{"ଏ", "ପି"},
		periodsWide:            []string{"ଏ ଏମ", "ପି ଏମ"},
		erasAbbreviated:        []string{"ବିସି", "ଏଡି"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ବିଫୋର କ୍ରାଇଷ୍ଟ", "ଆନ୍ନା ଡୋମିନି"},
		timezones:              map[string]string{"ACDT": "ଅସ୍ଟ୍ରେଲିଆତି ମାଦିନି ଡେଲାଇଟ ବେଲା", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ଅସ୍ଟ୍ରେଲିଆତି ମାଦିନି ୱେଡ଼ାକୁଣ୍\u200dପୁ ଡେଲାଇଟ ବେଲା", "ACWST": "ଅସ୍ଟ୍ରେଲିଆତି ମାଦିନି ୱେଡ଼ାକୁଣ୍\u200dପୁ ମାନାଙ୍କ ବେଲା", "ADT": "ଆଟ୍ଲାଣ୍ଟିକ ଡେଲାଇଟ ବେଲା", "ADT Arabia": "ଆରବିଆତି ମେଦାଣା ବେଲା", "AEDT": "ଅସ୍ଟ୍ରେଲିଆତି ୱେଡ଼ାହପୁ ଡେଲାଇଟ ବେଲା", "AEST": "ଅସ୍ଟ୍ରେଲିଆତି ୱେଡ଼ାହପୁ ମାନାଙ୍କ ବେଲା", "AFT": "ଆପଗାନିସ୍ତାନ ବେଲା", "AKDT": "ଆଲାସ୍କା ଡେଲାଇଟ୍ ବେଲା", "AKST": "ଆଲାସ୍କା ମାନାଙ୍କ ବେଲା", "AMST": "ଆମେଜନ କାରାଁ ବେଲା", "AMST Armenia": "ଆରମେନିୟା କାରାଁ ବେଲା", "AMT": "ଆମେଜନ ମାନାଙ୍କ ବେଲା", "AMT Armenia": "ଆରମେନିୟା ମାନାଙ୍କ ବେଲା", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ଆର୍ଜେଣ୍ଟିନା କାରା ବେଲା", "ART": "ଆରଜେଣ୍ଟିନା ମାନାଙ୍କ ବେଲା", "AST": "ଆଟ୍ଲାଣ୍ଟିକ ମାନାଙ୍କ ବେଲା", "AST Arabia": "ଆରବିଆତି ମାନାଙ୍କ ବେଲା", "AWDT": "ଅସ୍ଟ୍ରେଲିଆତି ୱେଡ଼ାକୁଣ୍\u200dପୁ ଡେଲାଇଟ ବେଲା", "AWST": "ଅସ୍ଟ୍ରେଲିଆତି ୱେଡ଼ାକୁଣ୍\u200dପୁ ମାନାଙ୍କ ବେଲା", "AZST": "ଅଜେରବାଇଜାନ କାରାଁ ବେଲା", "AZT": "ଅଜେରବାଇଜାନ ମାନାଙ୍କ ବେଲା", "BDT Bangladesh": "ବାଙ୍ଗଲାଦେସ୍ କାରାଁ ବେଲା", "BNT": "ବ୍ରୁନେଇ ଦାରୁସାଲାମ ବେଲା", "BOT": "ବଲୱିଆ ବେଲା", "BRST": "ବ୍ରାଜିଲିୟା କାରାଁ ବେଲା", "BRT": "ବ୍ରାଜିଲିୟା ମାନକ ବେଲା", "BST Bangladesh": "ବାଙ୍ଗଲାଦେସ୍ ମାନାଙ୍କ ବେଲା", "BT": "ବୁଟାନ ବେଲା", "CAST": "CAST", "CAT": "ମାଦିନି ଆପ୍ରିକା ବେଲା", "CCT": "କକସ ଦିପ ବେଲା", "CDT": "ମାଦିନି ଡେଲାଇଟ୍ ବେଲା", "CHADT": "ଚାତାମ ଡେଲାଇଟ୍ ବେଲା", "CHAST": "ଚାତାମ ମାନାଙ୍କ ବେଲା", "CHUT": "ଚୁକ ବେଲା", "CKT": "କୁକ ଦିପ ମାନାଙ୍କ ବେଲା", "CKT DST": "କୁକ ଦିପ ଆଦ୍ଦା କାରାଁ ବେଲା", "CLST": "ଚିଲି କାରାଁ ବେଲା", "CLT": "ଚିଲି ମାନାଙ୍କ ବେଲା", "COST": "କଲମ୍ବିୟା କାରାଁ ବେଲା", "COT": "କଲମ୍ବିୟା ମାନାଙ୍କ ବେଲା", "CST": "ମାଦିନି ମାନାଙ୍କ ବେଲା", "CST China": "ଚିନ ମାନାଙ୍କ ବେଲା", "CST China DST": "ଚିନ ଡେଲାଇଟ ବେଲା", "CVST": "କେପ ବର୍ଡ କାରାଁ ବେଲା", "CVT": "କେପ ବର୍ଡ ମାନାଙ୍କ ବେଲା", "CXT": "କ୍ରିସମାସ ଦିପ ବେଲା", "ChST": "ଚମର ମାନାଙ୍କ ବେଲା", "ChST NMI": "ChST NMI", "CuDT": "କ୍ୱୁବା ଡେଲାଇଟ୍ ବେଲା", "CuST": "କ୍ୱୁବା ମାନାଙ୍କ ବେଲା", "DAVT": "ଡେବିସ୍ ବେଲା", "DDUT": "ଡ୍ୟୁମାଣ୍ଟ ଡି ଅରୱିଲେ ବେଲା", "EASST": "ଇସ୍ଟର ଦିପ କାରାଁ ବେଲା", "EAST": "ଇସ୍ଟର ଦିପ ମାନାଙ୍କ ବେଲା", "EAT": "ୱେଡ଼ା ହପୁ ଆପ୍ରିକା ବେଲା", "ECT": "ଇକ୍ୱାଡର ବେଲା", "EDT": "ୱେଡ଼ାହପୁ ଜାଗାତି ଡେଲାଇଟ୍ ବେଲା", "EGDT": "ୱେଡ଼ା ହପୁ ଗ୍ରିନଲାଣ୍ଡ କାରାଁ ବେଲା", "EGST": "ୱେଡ଼ା ହପୁ ଗ୍ରିନଲାଣ୍ଡ ମାନାଙ୍କ ବେଲା", "EST": "ୱେଡ଼ାହପୁ ଜାଗାତି ମାନାଙ୍କ ବେଲା", "FEET": "ଅର ୱେଡ଼ାହପୁ ୟୁରପତି ବେଲା", "FJT": "ପିଜି ମାନାଙ୍କ ବେଲା", "FJT Summer": "ପିଜି କାରାଁ ବେଲା", "FKST": "ପାକଲାଣ୍ଡ ଦିପତି କାରାଁ ବେଲା", "FKT": "ପାକଲାଣ୍ଡ ଦିପତି ମାନାଙ୍କ ବେଲା", "FNST": "ପର୍ନାଡ ଡେ ନରହ୍ନ କାରାଁ ବେଲା", "FNT": "ପର୍ନାଡ ଡେ ନରହ୍ନ ମାନାଙ୍କା ବେଲା", "GALT": "ଗଲାପଗସ ତି ବେଲା", "GAMT": "ଗମ୍ବିୟର ବେଲା", "GEST": "ଜର୍ଜିୟା କାରା ବେଲା", "GET": "ଜର୍ଜିୟା ମାନାଙ୍କ ବେଲା", "GFT": "ପ୍ରେଞ୍ଚ୍ ଗୁୟାନା ବେଲା", "GIT": "ଗିଲ୍\u200dବେର୍ଟ ଦିପତି ବେଲା", "GMT": "ଗ୍ରିନୱିଚ ମିନ ବେଲା", "GNSST": "GNSST", "GNST": "GNST", "GST": "ଗଲ୍ପ ମାନାଙ୍କ ବେଲା", "GST Guam": "GST Guam", "GYT": "ଗୁୟାନା ବେଲା", "HADT": "ହାୱାଇ-ଆଲ୍ୟୁସାନ ଡେ ଲାଇଟ ବେଲା", "HAST": "ହାୱାଇ-ଆଲ୍ୟୁସାନ ମାନାଙ୍କ ବେଲା", "HKST": "ହଙ୍ଗ କଙ୍କ କାରାଁ ବେଲା", "HKT": "ହଙ୍ଗ କଙ୍ଗ ମାନାଙ୍କ ବେଲା", "HOVST": "ହୱଡ କାରାଁ ବେଲା", "HOVT": "ହୱଡ ମାନାଙ୍କ ବେଲା", "ICT": "ଇଣ୍ଡଚିନା ବେଲା", "IDT": "ଇଜରାଇଲ ଡେଲାଇଟ ବେଲା", "IOT": "ବାରତ କାଜା ସାମୁଦ୍ରି ବେଲା", "IRKST": "ଇରକୁସ୍ତକ କାରାଁ ବେଲା", "IRKT": "ଇରକୁସ୍ତକ ମାନାଙ୍କ ବେଲା", "IRST": "ଇରାନ ମାନାଙ୍କ ବେଲା", "IRST DST": "ଇରାନ ଡେ ଲାଇଟ୍ ବେଲା", "IST": "ବାରତ ମାନାଙ୍କ ବେଲା", "IST Israel": "ଇଜରାଇଲ ମାନାଙ୍କ ବେଲା", "JDT": "ଜାପାନ ଡେଲାଇଟ ବେଲା", "JST": "ଜାପାନ ମାନାଙ୍କ ବେଲା", "KOST": "କସରାଏ ବେଲା", "KRAST": "କ୍ରାସ୍ନାର୍ସ୍କ କାରାଁ ବେଲା", "KRAT": "କ୍ରାସ୍ନୟାର୍ସ୍କ ମାନାଙ୍କ ବେଲା", "KST": "କରିୟା ମାନାଙ୍କ ବେଲା", "KST DST": "କରିୟା ଡେଲାଇଟ ବେଲା", "LHDT": "ଲର୍ଡ ହୱେ ଡେଲାଇଟ୍ ବେଲା", "LHST": "ଲର୍ଡ ହୱେ ମାନାଙ୍କ ବେଲା", "LINT": "ଲାଇନ ଦିପତି ବେଲା", "MAGST": "ମାଗାଦାନ କାରାଁ ବେଲା", "MAGT": "ମାଗାଦାନ ମାନଙ୍କ ବେଲା", "MART": "ମାର୍କସସ ବେଲା", "MAWT": "ମାୱସନ ବେଲା", "MDT": "MDT", "MESZ": "ମାଦିନା ୟୁରପିଆ ତି କାରାଁ ବେଲା", "MEZ": "ମାଦିନି ୟୁରପିଆ ତି ମାନାଙ୍କ ବେଲା", "MHT": "ମାର୍ସାଲ ଦିପ ବେଲା", "MMT": "ମ୍ୟାଁମାର ବେଲା", "MSD": "ମସ୍କ କାରାଁ ବେଲା", "MST": "MST", "MUST": "ମରିସସ୍ କାରାଁ ବେଲା", "MUT": "ମରିସସ୍ ମାନାଙ୍କ ବେଲା", "MVT": "ମାଲଡ୍ୱିସ୍ ବେଲା", "MYT": "ମାଲେସିୟା ବେଲା", "NCT": "ନ୍ୟୁ କେଲେଡନିୟା ମାନାଙ୍କ ବେଲା", "NDT": "ନ୍ୟୁପାଉଣ୍ଡଲାଣ୍ଡ ଡେଲାଇଟ ବେଲା", "NDT New Caledonia": "ନ୍ୟୁ କେଲେଡନିୟା କାରାଁ ବେଲା", "NFDT": "ନରପାଁକ ଦିପ ଡେଲାଇଟ ବେଲା", "NFT": "ନରପାଁକ ଦିପ ମାନାଙ୍କ ବେଲା", "NOVST": "ନୱସିବର୍ସ୍କ କାରାଁ ବଲା", "NOVT": "ନୱସିବିର୍ସ୍କ ମାନାଙ୍କ ବେଲା", "NPT": "ନେପାଲ ବେଲା", "NRT": "ନଉରୁ ବେଲା", "NST": "ନ୍ୟୁପାଉଣ୍ଡଲାଣ୍ଡ ମାନାଙ୍କ ବେଲା", "NUT": "ନିୟୁ ବେଲା", "NZDT": "ନ୍ୟୁଜିଲାଣ୍ଡ ଡେଲାଇଟ ବେଲା", "NZST": "ନ୍ୟୁଜିଲାଣ୍ଡ ମାନାଙ୍କ ବେଲା", "OESZ": "ୱେଡ଼ାହପୁ ୟୁରପିଆତି କାରାଁ ବେଲା", "OEZ": "ୱେଡ଼ାହପୁ ୟୁରପିଆତି ମାନାଙ୍କ ବେଲା", "OMSST": "ଅମସ୍କ କାରାଁ ବେଲା", "OMST": "ଅମସ୍କ ମାନାଙ୍କ ବେଲା", "PDT": "ପେସିପିକ୍ ଡେଲାଇଟ୍ ବେଲା", "PDTM": "ମେକ୍ସିକତି ପେସିପିକ ଡେଲାଇଟ ବେଲା", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ପାପୁୟା ପୁନି ଗିନି ବେଲା", "PHOT": "ପିନିକ୍ସ ଦିପତି ବେଲା", "PKT": "ପାକିସ୍ତାନ ମାନାଙ୍କ ବେଲା", "PKT DST": "ପାକିସ୍ତାନ କାରାଁ ବେଲା", "PMDT": "ସେଣ୍ଟ ପିଏରେ ଅଡ଼େ ମିକ୍ୱେଲାନ ଡେଲାଇଟ ବେଲା", "PMST": "ସେଣ୍ଟ ପିଏରେ ଅଡ଼େ ମିକ୍ୱେଲାନ ମାନାଙ୍କ ବେଲା", "PONT": "ପନାପେ ବେଲା", "PST": "ପେସିପିକ୍ ମାନାଙ୍କ ବେଲା", "PST Philippine": "ପିଲିପିନ ମାନାଙ୍କ ବେଲା", "PST Philippine DST": "ପିଲିପିନ କାରାଁ ବେଲା", "PST Pitcairn": "ପିଟକଇରନ ବେଲା", "PSTM": "ମେକ୍ସିକତି ପେସିପିକ ମାନାଙ୍କ ବେଲା", "PWT": "ପଲାଉ ବେଲା", "PYST": "ପେରାଗ୍ୱେ କାରାଁ ବେଲା", "PYT": "ପେରାଗ୍ୱେ ମାନାଙ୍କ ବେଲା", "PYT Korea": "ପ୍ୟଙ୍ଗୟାଙ୍କ ବେଲା", "RET": "ରିୟୁନିୟନ ବେଲା", "ROTT": "ରତେରା ବେଲା", "SAKST": "ସକାଲିନ କାରାଁ ବେଲା", "SAKT": "ସକାଲିନ ମାନାଙ୍କ ବେଲା", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ଦକିଣ ଆପ୍ରିକା ମାନାଙ୍କ ବେଲା", "SBT": "ସଲମନ ଦିପତି ବେଲା", "SCT": "ସେସେଲ୍ସ ବେଲା", "SGT": "ସାଙ୍ଗାପୁର ମାନାଙ୍କ ବେଲା", "SLST": "SLST", "SRT": "ସୁରିନାମ ବେଲା", "SST Samoa": "ସାମଆ ମାନାଙ୍କ ବେଲା", "SST Samoa Apia": "ଏପିଆ ମାନାଙ୍କ ବେଲା", "SST Samoa Apia DST": "ଏପିଆ ଡେଲାଇଟ୍ ବେଲା", "SST Samoa DST": "ସାମଆ ଡେଲାଇଟ ବେଲା", "SYOT": "ସୱା ବେଲା", "TAAF": "ପ୍ରେଞ୍ଚ ଦକିଣ ଅଡ଼େ ଆଣ୍ଟାରଟିକ ବେଲା", "TAHT": "ତାହିତି ବେଲା", "TJT": "ତାଜିକିସ୍ତାନ ବେଲା", "TKT": "ଟକେଲାଉ ବେଲା", "TLT": "ୱେଡ଼ାହପୁ ତିମର ବେଲା", "TMST": "ତୁର୍କମେନିସ୍ତାନ କାରାଁ ବେଲା", "TMT": "ତୁର୍କମେନିସ୍ତାନ ମାନାଙ୍କ ବେଲା", "TOST": "ଟଙ୍ଗା କାରାଁ ବେଲା", "TOT": "ଟଙ୍ଗା ମାନାଙ୍କ ବେଲା", "TVT": "ତୁୱାଲୁ ବେଲା", "TWT": "ତାଇପେ ମାନଙ୍କ ବେଲା", "TWT DST": "ତାଇପେ ଡେଲାଇଟ ବେଲା", "ULAST": "ଉଲାନ ବଟର କାରାଁ ବେଲା", "ULAT": "ଉଲାନ ବଟର ମାନାଙ୍କ ବେଲା", "UYST": "ଉରୁଗ୍ୱେ କାରାଁ ବେଲା", "UYT": "ଉରୁଗ୍ୱେ ମାନାଙ୍କ ବେଲା", "UZT": "ଉଜ୍ୱେକିସ୍ତାନ ମାନାଙ୍କ ବେଲା", "UZT DST": "ଉଜ୍ୱେକିସ୍ତାନ କାରାଁ ବେଲା", "VET": "ୱେନେଜୁଏଲା ବେଲା", "VLAST": "ୱ୍ଲାଦିୱସ୍ତକ କାରାଁ ବେଲା", "VLAT": "ୱ୍ଲାଦିୱସ୍ତକ ମାନାଙ୍କ ବେଲା", "VOLST": "ୱାଲଗଗ୍ରାଡ କାରାଁ ବେଲା", "VOLT": "ୱାଲଗଗ୍ରାଡ ମାନାଙ୍କ ବେଲା", "VOST": "ୱସ୍ତକ ବେଲା", "VUT": "ୱନୁଆତୁ ମାନାଙ୍କ ବେଲା", "VUT DST": "ୱନୁଆତୁ କାରାଁ ବେଲା", "WAKT": "ୱେକ ଦିପ ବେଲା", "WARST": "ୱେଡ଼ା କୁଣ୍\u200dପୁ ଆର୍ଜେଣ୍ଟିନା କାରାଁ ବେଲା", "WART": "ୱେଡ଼ା କୁଣ୍\u200dପୁ ଆର୍ଜେଣ୍ଟିନା ମାନାଙ୍କ ବେଲା", "WAST": "ୱେଡ଼ା କୁଣ୍\u200dପୁ ଆପ୍ରିକା ବେଲା", "WAT": "ୱେଡ଼ା କୁଣ୍\u200dପୁ ଆପ୍ରିକା ବେଲା", "WESZ": "ୱେଡ଼ାକୁଣ୍\u200dପୁ ୟୁରପତି କାରାଁ ବେଲା", "WEZ": "ୱେଡ଼ାକୁଣ୍\u200dପୁ ୟୁରପିଆତି ମାନାଙ୍କ ବେଲା", "WFT": "ୱାଲିସ ଅଡ଼େ ପୁଟୁନା ବେଲା", "WGST": "ୱେଡ଼ା କୁଣ୍\u200dପୁ ଗ୍ରିନଲାଣ୍ଡ କାରାଁ ବେଲା", "WGT": "ୱେଡ଼ା କୁଣ୍\u200dପୁ ଗ୍ରିନଲାଣ୍ଡ ମାନାଙ୍କ ବେଲା", "WIB": "ୱେଡ଼ାକୁଣ୍ପୁ ଇଣ୍ଡନେସିଆ ବେଲା", "WIT": "ୱେଡ଼ାହପୁ ଇଣ୍ଡନେସିଆ ବେଲା", "WITA": "ମାଦିନି ଇଣ୍ଡନେସିଆ ବେଲା", "YAKST": "ୟାକୁତ୍ସକ କାରାଁ ବେଲା", "YAKT": "ୟାକୁତ୍ସକ ମାନାଙ୍କ ବେଲା", "YEKST": "ୟେକାତେରିନବର୍ଗ କାରାଁ ବେଲା", "YEKT": "ୟେକାତେରିନବର୍ଗ ମାନାଙ୍କ ବେଲା", "YST": "YST", "МСК": "ମସ୍କ ମାନାଙ୍କ ବେଲା", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ୱେଡ଼ାକୁଣ୍ପୁ କଜାକାସ୍ତାନ ବେଲା", "شىعىش قازاق ەلى": "ୱେଡ଼ାହପୁ କଜାକାସ୍ତାନ ବେଲା", "قازاق ەلى": "କଜାକାସ୍ତାନ ବେଲା", "قىرعىزستان": "କିର୍ଗିସ୍ତାନ ବେଲା", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ଅଜରେସ କାରାଁ ବେଲା"},
	}
}

// Locale returns the current translators string locale
func (kxv *kxv_Orya) Locale() string {
	return kxv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kxv_Orya'
func (kxv *kxv_Orya) PluralsCardinal() []locales.PluralRule {
	return kxv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kxv_Orya'
func (kxv *kxv_Orya) PluralsOrdinal() []locales.PluralRule {
	return kxv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kxv_Orya'
func (kxv *kxv_Orya) PluralsRange() []locales.PluralRule {
	return kxv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Orya'
func (kxv *kxv_Orya) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Orya'
func (kxv *kxv_Orya) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kxv_Orya'
func (kxv *kxv_Orya) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kxv *kxv_Orya) MonthAbbreviated(month time.Month) string {
	return kxv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kxv *kxv_Orya) MonthsAbbreviated() []string {
	return kxv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kxv *kxv_Orya) MonthNarrow(month time.Month) string {
	return kxv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kxv *kxv_Orya) MonthsNarrow() []string {
	return kxv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kxv *kxv_Orya) MonthWide(month time.Month) string {
	return kxv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kxv *kxv_Orya) MonthsWide() []string {
	return kxv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kxv *kxv_Orya) WeekdayAbbreviated(weekday time.Weekday) string {
	return kxv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kxv *kxv_Orya) WeekdaysAbbreviated() []string {
	return kxv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kxv *kxv_Orya) WeekdayNarrow(weekday time.Weekday) string {
	return kxv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kxv *kxv_Orya) WeekdaysNarrow() []string {
	return kxv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kxv *kxv_Orya) WeekdayShort(weekday time.Weekday) string {
	return kxv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kxv *kxv_Orya) WeekdaysShort() []string {
	return kxv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kxv *kxv_Orya) WeekdayWide(weekday time.Weekday) string {
	return kxv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kxv *kxv_Orya) WeekdaysWide() []string {
	return kxv.daysWide
}

// Decimal returns the decimal point of number
func (kxv *kxv_Orya) Decimal() string {
	return kxv.decimal
}

// Group returns the group of number
func (kxv *kxv_Orya) Group() string {
	return kxv.group
}

// Group returns the minus sign of number
func (kxv *kxv_Orya) Minus() string {
	return kxv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kxv_Orya' and handles both Whole and Real numbers based on 'v'
func (kxv *kxv_Orya) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kxv_Orya' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kxv *kxv_Orya) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kxv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kxv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Orya'
// in accounting notation.
func (kxv *kxv_Orya) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kxv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyNegativePrefix[j])
		}

		b = append(b, kxv.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, kxv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kxv_Orya'
func (kxv *kxv_Orya) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kxv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
