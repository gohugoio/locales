package cop

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type cop struct {
	locale            string
	pluralsCardinal   []locales.PluralRule
	pluralsOrdinal    []locales.PluralRule
	pluralsRange      []locales.PluralRule
	decimal           string
	group             string
	minus             string
	percent           string
	timeSeparator     string
	currencies        []string // idx = enum of currency code
	monthsAbbreviated []string
	monthsNarrow      []string
	monthsWide        []string
	daysAbbreviated   []string
	daysNarrow        []string
	daysShort         []string
	daysWide          []string
	timezones         map[string]string
}

// New returns a new instance of translator for the 'cop' locale
func New() locales.Translator {
	return &cop{
		locale:            "cop",
		pluralsCardinal:   nil,
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           ".",
		group:             ",",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "ⲓⲁⲛⲟⲩⲁⲣⲓ", "ⲫⲉⲃⲣⲟⲩⲁⲣⲓ", "ⲙⲁⲣⲧⲓ", "ⲁⲡⲣⲓⲗⲓ", "ⲙⲁⲓⲟ", "ⲓⲟⲩⲛⲓⲟⲩ", "ⲓⲟⲩⲗⲓⲟⲩ", "ⲁⲩⲅⲟⲩⲥⲑⲟⲩ", "ⲥⲉⲡⲧⲉⲙⲃⲣⲓ", "ⲟⲕⲧⲱⲃⲣⲓ", "ⲛⲟⲉⲙⲃⲣⲓ", "ⲇⲉⲕⲉⲙⲃⲣⲓ"},
		monthsNarrow:      []string{"", "Ⲓ", "Ⲫ", "Ⲙ", "Ⲁ", "Ⲙ", "Ⲓ", "Ⲓ", "Ⲁ", "Ⲥ", "Ⲟ", "Ⲛ", "Ⲇ"},
		monthsWide:        []string{"", "ⲓⲁⲛⲟⲩⲁⲣⲓ", "ⲫⲉⲃⲣⲟⲩⲁⲣⲓ", "ⲙⲁⲣⲧⲓ", "ⲁⲡⲣⲓⲗⲓ", "ⲙⲁⲓⲟ", "ⲓⲟⲩⲛⲓⲟⲩ", "ⲓⲟⲩⲗⲓⲟⲩ", "ⲁⲩⲅⲟⲩⲥⲑⲟⲩ", "ⲥⲉⲡⲧⲉⲙⲃⲣⲓ", "ⲟⲕⲧⲱⲃⲣⲓ", "ⲛⲟⲉⲙⲃⲣⲓ", "ⲇⲉⲕⲉⲙⲃⲣⲓ"},
		daysAbbreviated:   []string{"ⲡⲓⲁ̅", "ⲡⲓⲃ̅", "ⲡⲓⲅ̅", "ⲡⲓⲇ̅", "ⲡⲓⲉ̅", "ⲡⲓⲋ̅", "ⲡⲓⲍ̅"},
		daysNarrow:        []string{"ⲡⲓⲁ̅", "ⲡⲓⲃ̅", "ⲡⲓⲅ̅", "ⲡⲓⲇ̅", "ⲡⲓⲉ̅", "ⲡⲓⲋ̅", "ⲡⲓⲍ̅"},
		daysShort:         []string{"ⲡⲓⲁ̅", "ⲡⲓⲃ̅", "ⲡⲓⲅ̅", "ⲡⲓⲇ̅", "ⲡⲓⲉ̅", "ⲡⲓⲋ̅", "ⲡⲓⲍ̅"},
		daysWide:          []string{"ⲡⲓⲟⲩⲁⲓ", "ⲡⲓⲥ̀ⲛⲁⲩ", "ⲡⲓϣⲟⲙⲧ", "ⲡⲓϥ̀ⲧⲟⲩ", "ⲡⲓⲧ̀ⲓⲟⲩ", "ⲡⲓⲥⲟⲟⲩ", "ⲡⲓϣⲁϣϥ"},
		timezones:         map[string]string{"ACDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀̀ⲑⲙⲏϯ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀̀ⲑⲙⲏϯ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "ACWST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀̀ⲑⲙⲏϯ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "ADT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲁⲑⲗⲁⲛⲑⲓⲕ", "ADT Arabia": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲁⲣⲁⲃⲟⲥ", "AEDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "AEST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "AFT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲁϥⲅⲁⲛⲓⲥⲑⲁⲛ", "AKDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲁⲗⲁⲥⲕⲁ", "AKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲗⲁⲥⲕⲁ", "AMST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲁⲙⲁⲍⲟⲛ", "AMST Armenia": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲁⲣⲙⲉⲛⲓⲁ", "AMT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲙⲁⲍⲟⲛ", "AMT Armenia": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲣⲙⲉⲛⲓⲁ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲁⲣϫⲉⲛⲑⲓⲛ", "ART": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲣϫⲉⲛⲑⲓⲛ", "AST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲑⲗⲁⲛⲑⲓⲕ", "AST Arabia": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲣⲁⲃⲟⲥ", "AWDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "AWST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲁⲩⲥⲧⲣⲁⲗⲓⲁ", "AZST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲁⲍⲉⲣⲃⲁⲓϫⲁⲛ", "AZT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲁⲍⲉⲣⲃⲁⲓϫⲁⲛ", "BDT Bangladesh": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲃⲁⲛⲅⲗⲁⲇⲉϣ", "BNT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲃⲣⲟⲩⲛⲉⲓ", "BOT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲃⲟⲗⲓⲃⲓⲁ", "BRST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲃⲣⲁⲍⲓⲗ", "BRT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲃⲣⲁⲍⲓⲗ", "BST Bangladesh": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲃⲁⲛⲅⲗⲁⲇⲉϣ", "BT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲃⲟⲩⲑⲁⲛ", "CAST": "CAST", "CAT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϯⲫⲣⲓⲕⲓⲁ ⲛ̀̀ⲑⲙⲏϯ", "CCT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲕⲟⲕⲟⲥ", "CDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲑⲙⲏϯ", "CHADT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ϣⲁⲑⲁⲙ", "CHAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϣⲁⲑⲁⲙ", "CHUT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϭⲟⲩⲕ", "CKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲕⲟⲩⲕ", "CKT DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲕⲟⲩⲕ", "CLST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ϭⲓⲗⲉ", "CLT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϭⲓⲗⲉ", "COST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲕⲟⲗⲟⲙⲃⲓⲁ", "COT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲕⲟⲗⲟⲙⲃⲓⲁ", "CST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲑⲙⲏϯ", "CST China": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϭⲓⲛⲁ", "CST China DST": "ⲡ̀ⲥⲏⲟⲩ ⲛⲧⲟⲟⲩⲓ ⲛ̀ϭⲓⲛⲁ", "CVST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲕⲁⲡ ⲃⲉⲣⲇⲉ", "CVT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲕⲁⲡ ⲃⲉⲣⲇⲉ", "CXT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲓⲙⲟⲩⲓ ⲛ̀ⲭⲣⲓⲥⲧⲙⲁⲥ", "ChST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϭⲁⲙⲟⲣⲟ", "ChST NMI": "ChST NMI", "CuDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲧⲉ ⲕⲟⲩⲃⲁ", "CuST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲧⲉ ⲕⲟⲩⲃⲁ", "DAVT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲇⲁⲩⲓⲥ", "DDUT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲇⲟⲩⲣⲙⲟⲛⲧ ⲇⲟⲩⲣⲃⲓⲗ", "EASST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲓⲥⲑⲉⲣ ⲁⲓⲥⲗⲁⲛⲇ", "EAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲓⲥⲑⲉⲣ ⲁⲓⲥⲗⲁⲛⲇ", "EAT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϯⲫⲣⲓⲕⲓⲁ ⲙ̀ⲡⲉⲓⲉⲃⲧ", "ECT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲉⲕⲟⲩⲁⲇⲟⲣ", "EDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲉⲓⲉⲃⲧ", "EGDT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡϣⲱⲙ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲧⲉ ⲅⲣⲓⲛⲗⲁⲛⲇ", "EGST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲧⲉ ⲅⲣⲓⲛⲗⲁⲛⲇ", "EST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲓⲉⲃⲧ", "FEET": "ⲟⲩⲕⲉ ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "FJT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϥⲓⲇϫⲓ", "FJT Summer": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ϥⲓⲇϫⲓ", "FKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ϥⲟⲕⲗⲁⲛⲇ", "FKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ϥⲟⲕⲗⲁⲛⲇ", "FNST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ϥⲉⲣⲛⲁⲛⲇⲟ ⲇⲉ ⲛⲟⲣⲟⲛϩⲁ", "FNT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϥⲉⲣⲛⲁⲛⲇⲟ ⲇⲉ ⲛⲟⲣⲟⲛϩⲁ", "GALT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲅⲁⲗⲁⲡⲁⲅⲟⲥ", "GAMT": "̀ⲡⲥⲟⲏⲩ ⲛ̀ⲅⲁⲙⲃⲓⲉⲣ", "GEST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲅⲉⲟⲣⲅⲓⲁ", "GET": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲅⲉⲟⲣⲅⲓⲁ", "GFT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲅⲩⲁⲛⲁ ⲛ̀ϥⲉⲣⲉⲛⲥⲟⲥ", "GIT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ϫⲓⲗⲃⲉⲣⲧ", "GMT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲑⲙⲏϯ ⲛ̀ⲅⲣⲓⲛⲓϭ", "GNSST": "GNSST", "GNST": "GNST", "GST": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲓⲅⲟⲗϥ", "GST Guam": "GST Guam", "GYT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲅⲩⲁⲛⲁ", "HADT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϩⲁⲟⲩⲁⲓ-ⲁⲗⲉⲩⲑⲓⲁⲛ", "HAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϩⲁⲟⲩⲁⲓ-ⲁⲗⲉⲩⲑⲓⲁⲛ", "HKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ϩⲟⲛⲅ ⲕⲟⲛⲅ", "HKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϩⲟⲛⲅ ⲕⲟⲛⲅ", "HOVST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ϧⲟⲃⲇ", "HOVT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϧⲟⲃⲇ", "ICT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲓⲛⲇⲟϭⲓⲛⲁ", "IDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲓⲥⲣⲁⲏⲗ", "IOT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲓⲱⲕⲉⲁⲛⲟⲥ ⲛ̀ϩⲉⲛⲧⲟⲩ", "IRKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲓⲣⲕⲟⲩⲧⲥⲕ", "IRKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲓⲣⲕⲟⲩⲧⲥⲕ", "IRST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲓⲣⲁⲛ", "IRST DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲓⲣⲁⲛ", "IST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ϩⲉⲛⲧⲟⲩ", "IST Israel": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲓⲥⲣⲁⲏⲗ", "JDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲓⲁⲡⲁⲛ", "JST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲓⲁⲡⲁⲛ", "KOST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲕⲟⲥⲣⲁⲉ", "KRAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲕⲣⲁⲥⲛⲟⲓⲁⲣⲥⲕ", "KRAT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲕⲣⲁⲥⲛⲟⲓⲁⲣⲥⲕ", "KST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲕⲟⲣⲉⲁ", "KST DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲕⲟⲣⲉⲁ", "LHDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲗⲟⲣⲇ ϩⲟⲩⲉ", "LHST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲗⲟⲣⲇ ϩⲟⲩⲉ", "LINT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲗⲓⲛ", "MAGST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲙⲁⲅⲁⲇⲁⲛ", "MAGT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲙⲁⲅⲁⲇⲁⲛ", "MART": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲙⲁⲣⲕⲉⲥⲁⲥ", "MAWT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲙⲁⲩⲥⲟⲛ", "MDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲓⲧⲱⲟⲩ", "MESZ": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀̀ⲑⲙⲏϯ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "MEZ": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀̀ⲑⲙⲏϯ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "MHT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲙ̀ⲙⲁⲣϣⲁⲗ", "MMT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲙⲓⲁⲛⲙⲁⲣ", "MSD": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲙⲟⲥⲕⲟ", "MST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲓⲧⲱⲟⲩ", "MUST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲙⲁⲩⲣⲓⲑⲓⲟⲩⲥ", "MUT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲙⲁⲩⲣⲓⲑⲓⲟⲩⲥ", "MVT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲙⲁⲗⲇⲓⲃ", "MYT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲙⲁⲗⲉⲍⲓⲁ", "NCT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲛⲓⲟⲩⲕⲁⲗⲉⲇⲟⲛⲓⲁ", "NDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛⲛ̀ⲓⲟⲩϥⲁⲩⲛⲇⲗⲁⲛⲇ", "NDT New Caledonia": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲛⲓⲟⲩⲕⲁⲗⲉⲇⲟⲛⲓⲁ", "NFDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲛⲟⲣϥⲟⲕ", "NFT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲛⲟⲣϥⲟⲕ", "NOVST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲛⲟⲃⲟⲥⲓⲃⲓⲣⲥⲕ", "NOVT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲛⲟⲃⲟⲥⲓⲃⲓⲣⲥⲕ", "NPT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲉⲡⲁⲗ", "NRT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲁⲩⲣⲟⲩ", "NST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛⲛ̀ⲓⲟⲩϥⲁⲩⲛⲇⲗⲁⲛⲇ", "NUT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲟⲩⲉ", "NZDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲛⲓⲟⲩⲍⲓⲗⲁⲛⲇⲁ", "NZST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲛⲓⲟⲩⲍⲓⲗⲁⲛⲇⲁ", "OESZ": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "OEZ": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "OMSST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲟⲙⲥⲕ", "OMST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲟⲙⲥⲕ", "PDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲡⲁⲥⲓϥⲓⲕ", "PDTM": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲙ̀ⲙⲉⲝⲓⲕⲟ ⲙ̀ⲡⲁⲥⲓϥⲓⲕ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲁⲡⲟⲩⲁ ⲛⲓⲟⲩⲅⲓⲛⲉⲁ", "PHOT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲙ̀ⲫⲓⲛⲓⲝ", "PKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲁⲕⲓⲥⲑⲁⲛ", "PKT DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲡⲁⲕⲓⲥⲑⲁⲛ", "PMDT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲥⲁⲛⲧ ⲡⲓⲉⲣ & ⲙⲓⲕⲉⲗⲟⲛ", "PMST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲥⲁⲛⲧ ⲡⲓⲉⲣ & ⲙⲓⲕⲉⲗⲟⲛ", "PONT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲟϩⲛⲡⲉⲓ", "PST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲁⲥⲓϥⲓⲕ", "PST Philippine": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲫⲓⲗⲓⲡⲡⲓⲛ", "PST Philippine DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲫⲓⲗⲓⲡⲡⲓⲛ", "PST Pitcairn": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲓⲑⲕⲁⲓⲣⲛ", "PSTM": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲙⲉⲝⲓⲕⲟ ⲙ̀ⲡⲁⲥⲓϥⲓⲕ", "PWT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲙⲡⲁⲗⲁⲩ", "PYST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲡⲁⲣⲁⲅⲟⲩⲁⲓ", "PYT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲁⲣⲁⲅⲟⲩⲁⲓ", "PYT Korea": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲕⲟⲣⲉⲁ ⲙ̀ⲡⲉⲙϩⲓⲧ", "RET": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲣⲉⲓⲟⲩⲛⲓⲟⲛ", "ROTT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲣⲟⲑⲉⲣⲁ", "SAKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲥⲁϧⲁⲗⲓⲛ", "SAKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲥⲁϧⲁⲗⲓⲛ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲟϩⲣⲟⲥ ⲛ̀ϯⲫⲣⲓⲕⲓⲁ ⲙ̀ⲡⲓⲣⲏⲥ", "SBT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲛⲓⲙⲟⲩⲓ ⲛ̀ⲥⲟⲗⲟⲙⲟⲛ", "SCT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲥⲩϣⲉⲗ", "SGT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲥⲓⲛⲅⲁⲡⲟⲣ", "SLST": "SLST", "SRT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲥⲟⲩⲣⲓⲛⲁⲙ", "SST Samoa": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲥⲁⲙⲟⲣⲁ ⲛ̀ⲁⲙⲉⲣⲓⲕⲁⲛⲟⲥ", "SST Samoa Apia": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲥⲁⲙⲟⲁ", "SST Samoa Apia DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲥⲁⲙⲟⲁ", "SST Samoa DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲥⲁⲙⲟⲣⲁ ⲛ̀ⲁⲙⲉⲣⲓⲕⲁⲛⲟⲥ", "SYOT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲥⲩⲟⲩⲁ", "TAAF": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϥⲉⲣⲉⲛⲥⲟⲥ ⲙ̀ⲡⲓⲣⲏⲥ & ⲁⲛⲧⲁⲣⲕⲑⲓⲕⲁ", "TAHT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲑⲁϩⲓⲑⲓ", "TJT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲑⲁϫⲓⲕⲓⲥⲑⲁⲛ", "TKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲑⲟⲕⲉⲗⲁⲩ", "TLT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲑⲓⲙⲟⲣ-ⲗⲉⲥⲧ", "TMST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲑⲟⲩⲣⲕⲙⲉⲛⲓⲥⲑⲁⲛ", "TMT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲑⲟⲩⲣⲕⲙⲉⲛⲓⲥⲑⲁⲛ", "TOST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲑⲟⲛⲅⲁ", "TOT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲑⲟⲛⲅⲁ", "TVT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲑⲟⲩⲃⲁⲗⲟⲩ", "TWT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲑⲁⲓⲟⲩⲁⲛ", "TWT DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲧⲟⲟⲩⲓ ⲛ̀ⲑⲁⲓⲟⲩⲁⲛ", "ULAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲟⲩⲗⲁⲛⲃⲁⲧⲁⲣ", "ULAT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲟⲩⲗⲁⲛⲃⲁⲧⲁⲣ", "UYST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲟⲩⲣⲟⲩⲅⲟⲩⲁⲓ", "UYT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲟⲩⲣⲟⲩⲅⲟⲩⲁⲓ", "UZT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲟⲩⲍⲃⲉⲕⲓⲥⲑⲁⲛ", "UZT DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲟⲩⲍⲃⲉⲕⲓⲥⲑⲁⲛ", "VET": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲃⲉⲛⲉⲍⲟⲩⲉⲗⲁ", "VLAST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲃⲗⲁⲇⲓⲃⲟⲥⲑⲟⲕ", "VLAT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲃⲗⲁⲇⲓⲃⲟⲥⲑⲟⲕ", "VOLST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲃⲟⲗⲅⲟⲅⲣⲁⲇ", "VOLT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲃⲟⲗⲅⲟⲅⲣⲁⲇ", "VOST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲃⲟⲥⲑⲟⲕ", "VUT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲃⲁⲛⲁⲑⲟⲩ", "VUT DST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲃⲁⲛⲁⲑⲟⲩ", "WAKT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲓⲙⲟⲩⲓ ⲛ̀ⲟⲩⲁⲕ", "WARST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲁⲣϫⲉⲛⲑⲓⲛ", "WART": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲁⲣϫⲉⲛⲑⲓⲛ", "WAST": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ϯⲫⲣⲓⲕⲓⲁ", "WAT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ϯⲫⲣⲓⲕⲓⲁ", "WESZ": "ⲡⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "WEZ": "ⲡⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲉⲩⲣⲟⲡⲁ", "WFT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲟⲩⲁⲗⲓⲥ & ϥⲟⲩⲑⲟⲩⲛⲁ", "WGST": "̀ⲡⲥⲏⲟⲩ ⲙ̀̀ⲡϣⲱⲙ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲧⲉ ⲅⲣⲓⲛⲗⲁⲛⲇ", "WGT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲧⲉ ⲅⲣⲓⲛⲗⲁⲛⲇ", "WIB": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲓⲛⲇⲟⲛⲉⲥⲓⲁ", "WIT": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲓⲛⲇⲟⲛⲉⲥⲓⲁ", "WITA": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀̀ⲑⲙⲏϯ ⲛ̀ⲓⲛⲇⲟⲛⲉⲥⲓⲁ", "YAKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲓⲁⲕⲟⲩⲧⲥⲕ", "YAKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲓⲁⲕⲟⲩⲧⲥⲕ", "YEKST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲛ̀ⲓⲉⲕⲁⲑⲉⲣⲓⲛⲃⲟⲩⲣⲅ", "YEKT": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲛ̀ⲓⲉⲕⲁⲑⲉⲣⲓⲛⲃⲟⲩⲣⲅ", "YST": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲓⲟⲩⲕⲟⲛ", "МСК": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϩⲟⲣⲟⲥ ⲙ̀ⲙⲟⲥⲕⲟ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲙⲉⲛⲧ ⲛ̀ⲕⲁⲍⲁϧⲉⲥⲑⲁⲛ", "شىعىش قازاق ەلى": "ⲡ̀ⲥⲏⲟⲩ ⲙ̀ⲡⲉⲓⲉⲃⲧ ⲛ̀ⲕⲁⲍⲁϧⲉⲥⲑⲁⲛ", "قازاق ەلى": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲕⲁⲍⲁϧⲉⲥⲑⲁⲛ", "قىرعىزستان": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ⲕⲩⲣⲅⲩⲥⲑⲁⲛ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ⲡ̀ⲥⲏⲟⲩ ⲛ̀ϣⲱⲙ ⲙ̀ⲡⲉⲣⲟⲩ"},
	}
}

// Locale returns the current translators string locale
func (cop *cop) Locale() string {
	return cop.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'cop'
func (cop *cop) PluralsCardinal() []locales.PluralRule {
	return cop.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'cop'
func (cop *cop) PluralsOrdinal() []locales.PluralRule {
	return cop.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'cop'
func (cop *cop) PluralsRange() []locales.PluralRule {
	return cop.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'cop'
func (cop *cop) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'cop'
func (cop *cop) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'cop'
func (cop *cop) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (cop *cop) MonthAbbreviated(month time.Month) string {
	return cop.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (cop *cop) MonthsAbbreviated() []string {
	return cop.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (cop *cop) MonthNarrow(month time.Month) string {
	return cop.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (cop *cop) MonthsNarrow() []string {
	return cop.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (cop *cop) MonthWide(month time.Month) string {
	return cop.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (cop *cop) MonthsWide() []string {
	return cop.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (cop *cop) WeekdayAbbreviated(weekday time.Weekday) string {
	return cop.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (cop *cop) WeekdaysAbbreviated() []string {
	return cop.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (cop *cop) WeekdayNarrow(weekday time.Weekday) string {
	return cop.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (cop *cop) WeekdaysNarrow() []string {
	return cop.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (cop *cop) WeekdayShort(weekday time.Weekday) string {
	return cop.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (cop *cop) WeekdaysShort() []string {
	return cop.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (cop *cop) WeekdayWide(weekday time.Weekday) string {
	return cop.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (cop *cop) WeekdaysWide() []string {
	return cop.daysWide
}

// Decimal returns the decimal point of number
func (cop *cop) Decimal() string {
	return cop.decimal
}

// Group returns the group of number
func (cop *cop) Group() string {
	return cop.group
}

// Group returns the minus sign of number
func (cop *cop) Minus() string {
	return cop.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'cop' and handles both Whole and Real numbers based on 'v'
func (cop *cop) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cop.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cop.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cop.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'cop' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (cop *cop) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cop.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cop.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, cop.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'cop'
func (cop *cop) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cop.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'cop'
// in accounting notation.
func (cop *cop) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cop.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'cop'
func (cop *cop) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'cop'
func (cop *cop) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'cop'
func (cop *cop) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'cop'
func (cop *cop) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'cop'
func (cop *cop) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'cop'
func (cop *cop) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'cop'
func (cop *cop) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'cop'
func (cop *cop) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}
