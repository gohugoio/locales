package kn_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kn_IN struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	timeSeparator          string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'kn_IN' locale
func New() locales.Translator {
	return &kn_IN{
		locale:                 "kn_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "ಜನ", "ಫೆಬ್ರ", "ಮಾರ್ಚ್", "ಏಪ್ರಿ", "ಮೇ", "ಜೂನ್", "ಜುಲೈ", "ಆಗ", "ಸೆಪ್ಟೆಂ", "ಅಕ್ಟೋ", "ನವೆಂ", "ಡಿಸೆಂ"},
		monthsNarrow:           []string{"", "ಜ", "ಫೆ", "ಮಾ", "ಏ", "ಮೇ", "ಜೂ", "ಜು", "ಆ", "ಸೆ", "ಅ", "ನ", "ಡಿ"},
		monthsWide:             []string{"", "ಜನವರಿ", "ಫೆಬ್ರವರಿ", "ಮಾರ್ಚ್", "ಏಪ್ರಿಲ್", "ಮೇ", "ಜೂನ್", "ಜುಲೈ", "ಆಗಸ್ಟ್", "ಸೆಪ್ಟೆಂಬರ್", "ಅಕ್ಟೋಬರ್", "ನವೆಂಬರ್", "ಡಿಸೆಂಬರ್"},
		daysAbbreviated:        []string{"ಭಾನು", "ಸೋಮ", "ಮಂಗಳ", "ಬುಧ", "ಗುರು", "ಶುಕ್ರ", "ಶನಿ"},
		daysNarrow:             []string{"ಭಾ", "ಸೋ", "ಮಂ", "ಬು", "ಗು", "ಶು", "ಶ"},
		daysWide:               []string{"ಭಾನುವಾರ", "ಸೋಮವಾರ", "ಮಂಗಳವಾರ", "ಬುಧವಾರ", "ಗುರುವಾರ", "ಶುಕ್ರವಾರ", "ಶನಿವಾರ"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಕೇಂದ್ರ ಹಗಲು ಸಮಯ", "ACST": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಕೇಂದ್ರ ಪ್ರಮಾಣಿತ ಸಮಯ", "ACT": "ಏಕರ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "ACWDT": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಕೇಂದ್ರ ಪಶ್ಚಿಮ ಹಗಲು ಸಮಯ", "ACWST": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಕೇಂದ್ರ ಪಶ್ಚಿಮ ಪ್ರಮಾಣಿತ ಸಮಯ", "ADT": "ಅಟ್ಲಾಂಟಿಕ್ ದಿನದ ಸಮಯ", "ADT Arabia": "ಅರೇಬಿಯನ್ ಹಗಲು ಸಮಯ", "AEDT": "ಪೂರ್ವ ಆಸ್ಟ್ರೇಲಿಯಾದ ಹಗಲು ಸಮಯ", "AEST": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಪೂರ್ವ ಪ್ರಮಾಣಿತ ಸಮಯ", "AFT": "ಅಫಘಾನಿಸ್ತಾನ ಸಮಯ", "AKDT": "\u200cಅಲಾಸ್ಕಾ ಹಗಲು ಸಮಯ", "AKST": "ಅಲಾಸ್ಕಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "AMST": "ಅಮೆಜಾನ್ ಬೇಸಿಗೆ ಸಮಯ", "AMST Armenia": "ಅರ್ಮೇನಿಯ ಬೇಸಿಗೆ ಸಮಯ", "AMT": "ಅಮೆಜಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "AMT Armenia": "ಅರ್ಮೇನಿಯ ಪ್ರಮಾಣಿತ ಸಮಯ", "ANAST": "ಅನಡೀರ್\u200c ಹಗಲು ಸಮಯ", "ANAT": "ಅನಡೀರ್\u200c ಪ್ರಮಾಣಿತ ಸಮಯ", "ARST": "ಅರ್ಜೆಂಟಿನಾ ಬೇಸಿಗೆ ಸಮಯ", "ART": "ಅರ್ಜೆಂಟೀನಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "AST": "ಅಟ್ಲಾಂಟಿಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "AST Arabia": "ಅರೇಬಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "AWDT": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಪಶ್ಚಿಮ ಹಗಲು ಸಮಯ", "AWST": "ಆಸ್ಟ್ರೇಲಿಯಾದ ಪಶ್ಚಿಮ ಪ್ರಮಾಣಿತ ಸಮಯ", "AZST": "ಅಜರ್ಬೈಜಾನ್ ಬೇಸಿಗೆ ಸಮಯ", "AZT": "ಅಜರ್ಬೈಜಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "BDT Bangladesh": "ಬಾಂಗ್ಲಾದೇಶ ಬೇಸಿಗೆ ಸಮಯ", "BNT": "ಬ್ರೂನಿ ದಾರುಸಲೆಮ್ ಸಮಯ", "BOT": "ಬೊಲಿವಿಯಾ ಸಮಯ", "BRST": "ಬ್ರೆಸಿಲಿಯಾ ಬೇಸಿಗೆ ಸಮಯ", "BRT": "ಬ್ರೆಸಿಲಿಯಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "BST Bangladesh": "ಬಾಂಗ್ಲಾದೇಶ ಪ್ರಮಾಣಿತ ಸಮಯ", "BT": "ಭೂತಾನ್ ಸಮಯ", "CAST": "CAST", "CAT": "ಮಧ್ಯ ಆಫ್ರಿಕಾ ಸಮಯ", "CCT": "ಕೋಕೋಸ್ ದ್ವೀಪಗಳ ಸಮಯ", "CDT": "ಉತ್ತರ ಅಮೆರಿಕದ ಕೇಂದ್ರೀಯ ದಿನದ ಸಮಯ", "CHADT": "ಚಥಾಮ್ ಹಗಲು ಸಮಯ", "CHAST": "ಚಥಾಮ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "CHUT": "ಚುಕ್ ಸಮಯ", "CKT": "ಕುಕ್ ದ್ವೀಪಗಳ ಪ್ರಮಾಣಿತ ಸಮಯ", "CKT DST": "ಕುಕ್ ದ್ವೀಪಗಳ ಮಧ್ಯಕಾಲೀನ ಬೇಸಿಗೆ ಸಮಯ", "CLST": "ಚಿಲಿ ಬೇಸಿಗೆ ಸಮಯ", "CLT": "ಚಿಲಿ ಪ್ರಮಾಣಿತ ಸಮಯ", "COST": "ಕೊಲಂಬಿಯಾ ಬೇಸಿಗೆ ಸಮಯ", "COT": "ಕೊಲಂಬಿಯಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "CST": "ಉತ್ತರ ಅಮೆರಿಕದ ಕೇಂದ್ರ ಪ್ರಮಾಣಿತ ಸಮಯ", "CST China": "ಚೀನಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "CST China DST": "ಚೀನಾ ಹಗಲು ಸಮಯ", "CVST": "ಕೇಪ್ ವರ್ಡ್ ಬೇಸಿಗೆ ಸಮಯ", "CVT": "ಕೇಪ್ ವರ್ಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "CXT": "ಕ್ರಿಸ್ಮಸ್ ದ್ವೀಪ ಸಮಯ", "ChST": "ಚಮೋರೋ ಪ್ರಮಾಣಿತ ಸಮಯ", "ChST NMI": "ಉತ್ತರ ಮರಿಯಾನಾ ದ್ವೀಪಗಳ ಸಮಯ", "CuDT": "ಕ್ಯೂಬಾ ದಿನದ ಸಮಯ", "CuST": "ಕ್ಯೂಬಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "DAVT": "ಡೇವಿಸ್ ಸಮಯ", "DDUT": "ಡುಮಂಟ್-ಡಿ ಉರ್ವಿಲೆ ಸಮಯ", "EASST": "ಈಸ್ಟರ್ ದ್ವೀಪ ಬೇಸಿಗೆ ಸಮಯ", "EAST": "ಈಸ್ಟರ್ ದ್ವೀಪ ಪ್ರಮಾಣಿತ ಸಮಯ", "EAT": "ಪೂರ್ವ ಆಫ್ರಿಕಾ ಸಮಯ", "ECT": "ಈಕ್ವೆಡಾರ್ ಸಮಯ", "EDT": "ಉತ್ತರ ಅಮೆರಿಕದ ಪೂರ್ವದ ದಿನದ ಸಮಯ", "EGDT": "ಪೂರ್ವ ಗ್ರೀನ್\u200cಲ್ಯಾಂಡ್ ಬೇಸಿಗೆ ಸಮಯ", "EGST": "ಪೂರ್ವ ಗ್ರೀನ್\u200cಲ್ಯಾಂಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "EST": "ಉತ್ತರ ಅಮೆರಿಕದ ಪೂರ್ವದ ಪ್ರಮಾಣಿತ ಸಮಯ", "FEET": "ಮತ್ತಷ್ಟು-ಪೂರ್ವ ಯುರೋಪಿಯನ್ ಸಮಯ", "FJT": "ಫಿಜಿ ಪ್ರಮಾಣಿತ ಸಮಯ", "FJT Summer": "ಫಿಜಿ ಬೇಸಿಗೆ ಸಮಯ", "FKST": "ಫಾಕ್\u200cಲ್ಯಾಂಡ್ ದ್ವೀಪಗಳ ಬೇಸಿಗೆ ಸಮಯ", "FKT": "ಫಾಕ್\u200cಲ್ಯಾಂಡ್ ದ್ವೀಪಗಳ ಪ್ರಮಾಣಿತ ಸಮಯ", "FNST": "ಫರ್ನಾಂಡೋ ದೆ ನೊರೊನ್ಹಾ ಬೇಸಿಗೆ ಸಮಯ", "FNT": "ಫೆರ್ನಾಂಡೋ ಡೆ ನೊರೊನ್ಹಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "GALT": "ಗಾಲಾಪಾಗೋಸ್ ಸಮಯ", "GAMT": "ಗ್ಯಾಂಬಿಯರ್ ಸಮಯ", "GEST": "ಜಾರ್ಜಿಯಾ ಬೇಸಿಗೆ ಸಮಯ", "GET": "ಜಾರ್ಜಿಯಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "GFT": "ಫ್ರೆಂಚ್ ಗಯಾನಾ ಸಮಯ", "GIT": "ಗಿಲ್ಬರ್ಟ್ ದ್ವೀಪಗಳ ಸಮಯ", "GMT": "ಗ್ರೀನ್\u200cವಿಚ್ ಸರಾಸರಿ ಕಾಲಮಾನ", "GNSST": "GNSST", "GNST": "GNST", "GST": "ದಕ್ಷಿಣ ಜಾರ್ಜಿಯಾ ಸಮಯ", "GST Guam": "ಗುವಾಮ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "GYT": "ಗಯಾನಾ ಸಮಯ", "HADT": "ಹವಾಯಿ-ಅಲ್ಯುಟಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "HAST": "ಹವಾಯಿ-ಅಲ್ಯುಟಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "HKST": "ಹಾಂಗ್ ಕಾಂಗ್ ಬೇಸಿಗೆ ಸಮಯ", "HKT": "ಹಾಂಗ್ ಕಾಂಗ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "HOVST": "ಹವ್ಡ್ ಬೇಸಿಗೆ ಸಮಯ", "HOVT": "ಹವ್ಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "ICT": "ಇಂಡೊಚೈನಾ ಸಮಯ", "IDT": "ಇಸ್ರೇಲ್ ಹಗಲು ಸಮಯ", "IOT": "ಹಿಂದೂ ಮಹಾಸಾಗರ ಸಮಯ", "IRKST": "ಇರ್\u200cಕುಟಸ್ಕ್ ಬೇಸಿಗೆ ಸಮಯ", "IRKT": "ಇರ್\u200cಕುಟಸ್ಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "IRST": "ಇರಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "IRST DST": "ಇರಾನ್ ಹಗಲು ಸಮಯ", "IST": "ಭಾರತೀಯ ಪ್ರಮಾಣಿತ ಸಮಯ", "IST Israel": "ಇಸ್ರೇಲ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "JDT": "ಜಪಾನ್ ಹಗಲು ಸಮಯ", "JST": "ಜಪಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "KOST": "ಕೊಸರೆ ಸಮಯ", "KRAST": "ಕ್ರಾಸ್\u200cನೊಯಾರ್ಸ್ಕ್ ಬೇಸಿಗೆ ಸಮಯ", "KRAT": "ಕ್ರಾಸ್\u200cನೊಯಾರ್ಸ್ಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "KST": "ಕೊರಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "KST DST": "ಕೊರಿಯನ್ ಹಗಲು ಸಮಯ", "LHDT": "ಲಾರ್ಡ್ ಹೋವ್ ಬೆಳಗಿನ ಸಮಯ", "LHST": "ಲಾರ್ಡ್ ಹೋವ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "LINT": "ಲೈನ್ ದ್ವೀಪಗಳ ಸಮಯ", "MAGST": "ಮಗಡಾನ್ ಬೇಸಿಗೆ ಸಮಯ", "MAGT": "ಮಗಡಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "MART": "ಮಾರ್ಕ್ಯುಸಸ್ ಸಮಯ", "MAWT": "ಮಾವ್\u200cಸನ್ ಸಮಯ", "MDT": "ಮಕಾವ್ ಬೇಸಿಗೆ ಸಮಯ", "MESZ": "ಮಧ್ಯ ಯುರೋಪಿಯನ್ ಬೇಸಿಗೆ ಸಮಯ", "MEZ": "ಮಧ್ಯ ಯುರೋಪಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "MHT": "ಮಾರ್ಷಲ್ ದ್ವೀಪಗಳ ಸಮಯ", "MMT": "ಮ್ಯಾನ್ಮಾರ್ ಸಮಯ", "MSD": "ಮಾಸ್ಕೋ ಬೇಸಿಗೆ ಸಮಯ", "MST": "ಮಕಾವ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "MUST": "ಮಾರಿಷಸ್ ಬೇಸಿಗೆ ಸಮಯ", "MUT": "ಮಾರಿಷಸ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "MVT": "ಮಾಲ್ಡೀವ್ಸ್ ಸಮಯ", "MYT": "ಮಲೇಷಿಯಾ ಸಮಯ", "NCT": "ಹೊಸ ಕ್ಯಾಲೆಡೋನಿಯಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "NDT": "ನ್ಯೂಫೌಂಡ್\u200cಲ್ಯಾಂಡ್ ದಿನದ ಸಮಯ", "NDT New Caledonia": "ಹೊಸ ಕ್ಯಾಲೆಡೋನಿಯಾ ಬೇಸಿಗೆಯ ಸಮಯ", "NFDT": "ನಾರ್ಫೋಕ್ ದ್ವೀಪ ಬೇಸಿಗೆ ಸಮಯ", "NFT": "ನಾರ್ಫೋಕ್ ದ್ವೀಪ ಪ್ರಮಾಣಿತ ಸಮಯ", "NOVST": "ನೊವೊಸಿಬಿರ್\u200cಸ್ಕ್ ಬೇಸಿಗೆ ಸಮಯ", "NOVT": "ನೊವೊಸಿಬಿರ್\u200cಸ್ಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "NPT": "ನೇಪಾಳ ಸಮಯ", "NRT": "ನೌರು ಸಮಯ", "NST": "ನ್ಯೂಫೌಂಡ್\u200cಲ್ಯಾಂಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "NUT": "ನಿಯು ಸಮಯ", "NZDT": "ನ್ಯೂಜಿಲ್ಯಾಂಡ್ ಹಗಲು ಸಮಯ", "NZST": "ನ್ಯೂಜಿಲ್ಯಾಂಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "OESZ": "ಪೂರ್ವ ಯುರೋಪಿಯನ್ ಬೇಸಿಗೆ ಸಮಯ", "OEZ": "ಪೂರ್ವ ಯುರೋಪಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "OMSST": "ಒಮಾಸ್ಕ್ ಬೇಸಿಗೆ ಸಮಯ", "OMST": "ಒಮಾಸ್ಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "PDT": "ಉತ್ತರ ಅಮೆರಿಕದ ಪೆಸಿಫಿಕ್ ದಿನದ ಸಮಯ", "PDTM": "ಮೆಕ್ಸಿಕನ್ ಪೆಸಿಫಿಕ್ ಹಗಲು ಸಮಯ", "PETDT": "ಪೆತ್ರೋಪಾವ್ಲೋಸ್ಕ್\u200c\u200c-ಕಮ್ಚತ್ಸ್\u200cಕೀ ಬೇಸಿಗೆ ಸಮಯ", "PETST": "ಪೆತ್ರೋಪಾವ್ಲೋಸ್ಕ್\u200c\u200c-ಕಮ್ಚತ್ಸ್\u200cಕೀ ಪ್ರಮಾಣಿತ ಸಮಯ", "PGT": "ಪಪುವಾ ನ್ಯೂ ಗಿನಿಯಾ ಸಮಯ", "PHOT": "ಫಿನಿಕ್ಸ್ ದ್ವೀಪಗಳ ಸಮಯ", "PKT": "ಪಾಕಿಸ್ತಾನ ಪ್ರಮಾಣಿತ ಸಮಯ", "PKT DST": "ಪಾಕಿಸ್ತಾನ ಬೇಸಿಗೆ ಸಮಯ", "PMDT": "ಸೇಂಟ್ ಪಿಯರ್ ಮತ್ತು ಮಿಕ್ವೆಲನ್ ಹಗಲು ಸಮಯ", "PMST": "ಸೇಂಟ್ ಪಿಯರ್ ಮತ್ತು ಮಿಕ್ವೆಲನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "PONT": "ಪೊನಾಪೆ ಸಮಯ", "PST": "ಉತ್ತರ ಅಮೆರಿಕದ ಪೆಸಿಫಿಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "PST Philippine": "ಫಿಲಿಫೈನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "PST Philippine DST": "ಫಿಲಿಫೈನ್ ಬೇಸಿಗೆ ಸಮಯ", "PST Pitcairn": "ಪಿಟ್\u200cಕೈರ್ನ್ ಸಮಯ", "PSTM": "ಮೆಕ್ಸಿಕನ್ ಪೆಸಿಫಿಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "PWT": "ಪಾಲಾವ್ ಸಮಯ", "PYST": "ಪರಾಗ್ವೇ ಬೇಸಿಗೆ ಸಮಯ", "PYT": "ಪರಾಗ್ವೇ ಪ್ರಮಾಣಿತ ಸಮಯ", "PYT Korea": "ಪ್ಯೊಂಗ್ಯಾಂಗ್ ಸಮಯ", "RET": "ರಿಯೂನಿಯನ್ ಸಮಯ", "ROTT": "ರೊತೇರಾ ಸಮಯ", "SAKST": "ಸ್ಯಾಕ್\u200cಹಲಿನ್ ಬೇಸಿಗೆ ಸಮಯ", "SAKT": "ಸ್ಯಾಕ್\u200cಹಲಿನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "SAMST": "ಸಮರ ಬೇಸಿಗೆ ಸಮಯ", "SAMT": "ಸಮರ ಪ್ರಮಾಣಿತ ಸಮಯ", "SAST": "ದಕ್ಷಿಣ ಆಫ್ರಿಕಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "SBT": "ಸಾಲಮನ್ ದ್ವೀಪಗಳ ಸಮಯ", "SCT": "ಸೀಷೆಲ್ಸ್ ಸಮಯ", "SGT": "ಸಿಂಗಪುರ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "SLST": "ಲಂಕಾ ಸಮಯ", "SRT": "ಸುರಿನೇಮ್ ಸಮಯ", "SST Samoa": "ಸಮೋವಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "SST Samoa Apia": "ಅಪಿಯಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "SST Samoa Apia DST": "ಅಪಿಯಾ ಹಗಲು ಸಮಯ", "SST Samoa DST": "ಸಮೋವಾ ಬೇಸಿಗೆ ಸಮಯ", "SYOT": "ಸ್ಯೊವಾ ಸಮಯ", "TAAF": "ದಕ್ಷಿಣ ಫ್ರೆಂಚ್ ಮತ್ತು ಅಂಟಾರ್ಟಿಕಾ ಸಮಯ", "TAHT": "ತಾಹಿತಿ ಸಮಯ", "TJT": "ತಝಕಿಸ್ತಾನ್ ಸಮಯ", "TKT": "ಟೊಕೆಲಾವ್ ಸಮಯ", "TLT": "ಪೂರ್ವ ಟಿಮೋರ್ ಸಮಯ", "TMST": "ತುರ್ಕ್\u200cಮೇನಿಸ್ತಾನ್ ಬೇಸಿಗೆ ಸಮಯ", "TMT": "ತುರ್ಕ್\u200cಮೇನಿಸ್ತಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "TOST": "ಟೊಂಗಾ ಬೇಸಿಗೆ ಸಮಯ", "TOT": "ಟೊಂಗಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "TVT": "ತುವಾಲು ಸಮಯ", "TWT": "ತೈಪೆ ಪ್ರಮಾಣಿತ ಸಮಯ", "TWT DST": "ತೈಪೆ ಹಗಲು ಸಮಯ", "ULAST": "ಉಲಾನ್ ಬತೊರ್ ಬೇಸಿಗೆ ಸಮಯ", "ULAT": "ಉಲಾನ್ ಬತೊರ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "UYST": "ಉರುಗ್ವೇ ಬೇಸಿಗೆ ಸಮಯ", "UYT": "ಉರುಗ್ವೇ ಪ್ರಮಾಣಿತ ಸಮಯ", "UZT": "ಉಜ್ಬೇಕಿಸ್ತಾನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "UZT DST": "ಉಜ್ಬೇಕಿಸ್ತಾನ್ ಬೇಸಿಗೆ ಸಮಯ", "VET": "ವೆನಿಜುವೆಲಾ ಸಮಯ", "VLAST": "ವ್ಲಾಡಿವೋಸ್ಟೋಕ್ ಬೇಸಿಗೆ ಸಮಯ", "VLAT": "ವ್ಲಾಡಿವೋಸ್ಟೋಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "VOLST": "ವೋಲ್ಗೋಗಾರ್ಡ್ ಬೇಸಿಗೆ ಸಮಯ", "VOLT": "ವೋಲ್ಗೋಗಾರ್ಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "VOST": "ವೋಸ್ಟೊಕ್ ಸಮಯ", "VUT": "ವನೌತು ಪ್ರಮಾಣಿತ ಸಮಯ", "VUT DST": "ವನೌತು ಬೇಸಿಗೆ ಸಮಯ", "WAKT": "ವೇಕ್ ದ್ವೀಪ ಸಮಯ", "WARST": "ಪಶ್ಚಿಮ ಅರ್ಜೆಂಟೀನಾ ಬೇಸಿಗೆ ಸಮಯ", "WART": "ಪಶ್ಚಿಮ ಅರ್ಜೆಂಟೀನಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "WAST": "ಪಶ್ಚಿಮ ಆಫ್ರಿಕಾ ಸಮಯ", "WAT": "ಪಶ್ಚಿಮ ಆಫ್ರಿಕಾ ಸಮಯ", "WESZ": "ಪಶ್ಚಿಮ ಯುರೋಪಿಯನ್ ಬೇಸಿಗೆ ಸಮಯ", "WEZ": "ಪಶ್ಚಿಮ ಯುರೋಪಿಯನ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "WFT": "ವ್ಯಾಲೀಸ್ ಮತ್ತು ಫ್ಯುಟುನಾ ಸಮಯ", "WGST": "ಪಶ್ಚಿಮ ಗ್ರೀನ್\u200cಲ್ಯಾಂಡ್ ಬೇಸಿಗೆ ಸಮಯ", "WGT": "ಪಶ್ಚಿಮ ಗ್ರೀನ್\u200cಲ್ಯಾಂಡ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "WIB": "ಪಶ್ಚಿಮ ಇಂಡೋನೇಷಿಯ ಸಮಯ", "WIT": "ಪೂರ್ವ ಇಂಡೋನೇಷಿಯಾ ಸಮಯ", "WITA": "ಮಧ್ಯ ಇಂಡೋನೇಷಿಯಾ ಸಮಯ", "YAKST": "ಯಾಕುಟ್ಸಕ್ ಬೇಸಿಗೆ ಸಮಯ", "YAKT": "ಯಾಕುಟ್ಸಕ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "YEKST": "ಯೇಕೇಟರಿನ್\u200dಬರ್ಗ್ ಬೇಸಿಗೆ ಸಮಯ", "YEKT": "ಯೇಕಟರಿನ್\u200cಬರ್ಗ್ ಪ್ರಮಾಣಿತ ಸಮಯ", "YST": "ಯುಕಾನ್ ಸಮಯ", "МСК": "ಮಾಸ್ಕೋ ಪ್ರಮಾಣಿತ ಸಮಯ", "اقتاۋ": "ಅಕ್ಟೌ ಪ್ರಮಾಣಿತ ಸಮಯ", "اقتاۋ قالاسى": "ಅಕ್ಟೌ ಬೇಸಿಗೆ ಸಮಯ", "اقتوبە": "ಅಕ್ಟೋಬೆ ಪ್ರಮಾಣಿತ ಸಮಯ", "اقتوبە قالاسى": "ಅಕ್ಟೋಬೆ ಬೇಸಿಗೆ ಸಮಯ", "الماتى": "ಆಲ್ಮೇಟಿ ಪ್ರಮಾಣಿತ ಸಮಯ", "الماتى قالاسى": "ಆಲ್ಮೇಟಿ ಬೇಸಿಗೆ ಸಮಯ", "باتىس قازاق ەلى": "ಪಶ್ಚಿಮ ಕಝಾಖ್\u200cಸ್ತಾನ್ ಸಮಯ", "شىعىش قازاق ەلى": "ಪೂರ್ವ ಕಝಾಖ್\u200cಸ್ತಾನ್ ಸಮಯ", "قازاق ەلى": "ಕಝಾಖ್\u200cಸ್ತಾನ್ ಸಮಯ", "قىرعىزستان": "ಕಿರ್ಗಿಸ್ತಾನ್ ಸಮಯ", "قىزىلوردا": "ಕೋಜಿಲೋರ್ಡಾ ಪ್ರಮಾಣಿತ ಸಮಯ", "قىزىلوردا قالاسى": "ಕೋಜಿಲೋರ್ಡಾ ಬೇಸಿಗೆ ಸಮಯ", "∅∅∅": "ಪೆರು ಬೇಸಿಗೆ ಸಮಯ"},
	}
}

// Locale returns the current translators string locale
func (kn *kn_IN) Locale() string {
	return kn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kn_IN'
func (kn *kn_IN) PluralsCardinal() []locales.PluralRule {
	return kn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kn_IN'
func (kn *kn_IN) PluralsOrdinal() []locales.PluralRule {
	return kn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kn_IN'
func (kn *kn_IN) PluralsRange() []locales.PluralRule {
	return kn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kn_IN'
func (kn *kn_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kn_IN'
func (kn *kn_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kn_IN'
func (kn *kn_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := kn.CardinalPluralRule(num1, v1)
	end := kn.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kn *kn_IN) MonthAbbreviated(month time.Month) string {
	return kn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kn *kn_IN) MonthsAbbreviated() []string {
	return kn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kn *kn_IN) MonthNarrow(month time.Month) string {
	return kn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kn *kn_IN) MonthsNarrow() []string {
	return kn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kn *kn_IN) MonthWide(month time.Month) string {
	return kn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kn *kn_IN) MonthsWide() []string {
	return kn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kn *kn_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return kn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kn *kn_IN) WeekdaysAbbreviated() []string {
	return kn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kn *kn_IN) WeekdayNarrow(weekday time.Weekday) string {
	return kn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kn *kn_IN) WeekdaysNarrow() []string {
	return kn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kn *kn_IN) WeekdayShort(weekday time.Weekday) string {
	return kn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kn *kn_IN) WeekdaysShort() []string {
	return kn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kn *kn_IN) WeekdayWide(weekday time.Weekday) string {
	return kn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kn *kn_IN) WeekdaysWide() []string {
	return kn.daysWide
}

// Decimal returns the decimal point of number
func (kn *kn_IN) Decimal() string {
	return kn.decimal
}

// Group returns the group of number
func (kn *kn_IN) Group() string {
	return kn.group
}

// Group returns the minus sign of number
func (kn *kn_IN) Minus() string {
	return kn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kn_IN' and handles both Whole and Real numbers based on 'v'
func (kn *kn_IN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kn.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kn_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kn *kn_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kn.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kn_IN'
func (kn *kn_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kn.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kn.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, kn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kn_IN'
// in accounting notation.
func (kn *kn_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kn.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kn.group[0])
				count = 1
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

		b = append(b, kn.currencyNegativePrefix[0])

	} else {
		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, kn.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, kn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, kn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, kn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, kn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kn.periodsAbbreviated[0]...)
	} else {
		b = append(b, kn.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kn.periodsAbbreviated[0]...)
	} else {
		b = append(b, kn.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kn.periodsAbbreviated[0]...)
	} else {
		b = append(b, kn.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kn_IN'
func (kn *kn_IN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kn.periodsAbbreviated[0]...)
	} else {
		b = append(b, kn.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
