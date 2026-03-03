package syr

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type syr struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	timeSeparator      string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'syr' locale
func New() locales.Translator {
	return &syr{
		locale:             "syr",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "ل.س.\u200f", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "ܟܢܘܢ ܒ", "ܫܒܛ", "ܐܕܪ", "ܢܝܣܢ", "ܐܝܪ", "ܚܙܝܪܢ", "ܬܡܘܙ", "ܐܒ", "ܐܝܠܘܠ", "ܬܫܪܝܢ ܐ", "ܬܫܪܝܢ ܒ", "ܟܢܘܢ ܐ"},
		monthsNarrow:       []string{"", "ܟ", "ܫ", "ܐ", "ܢ", "ܐ", "ܚ", "ܬ", "ܐ", "ܐ", "ܬ", "ܬ", "ܟ"},
		monthsWide:         []string{"", "ܟܢܘܢ ܐܚܪܝܐ", "ܫܒܛ", "ܐܕܪ", "ܢܝܣܢ", "ܐܝܪ", "ܚܙܝܪܢ", "ܬܡܘܙ", "ܐܒ", "ܐܝܠܘܠ", "ܬܫܪܝܢ ܩܕܡܝܐ", "ܬܫܪܝܢ ܐܚܪܝܐ", "ܟܢܘܢ ܩܕܡܝܐ"},
		daysAbbreviated:    []string{"ܚܕ", "ܬܪܝܢ", "ܬܠܬ", "ܐܪܒܥ", "ܚܡܫ", "ܥܪܘ", "ܫܒܬܐ"},
		daysNarrow:         []string{"ܚ", "ܬ", "ܬ", "ܐ", "ܚ", "ܥ", "ܫ"},
		daysWide:           []string{"ܚܕܒܫܒܐ", "ܬܪܝܢܒܫܒܐ", "ܬܠܬܒܫܒܐ", "ܐܪܒܥܒܫܒܐ", "ܚܡܫܒܫܒܐ", "ܥܪܘܒܬܐ", "ܫܒܬܐ"},
		periodsAbbreviated: []string{"\u070fܩܛ\u200c", "\u070fܒܛ\u200c"},
		timezones:          map[string]string{"ACDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܣܬܪܠܝܐ ܡܨܥܝܬܐ", "ACST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܣܬܪܠܝܐ ܡܨܥܝܬܐ", "ACT": "ܥܕܢܘܬܐ ܫܪܫܝܬܐ ܥܕܢܘܬܐ ܫܪܫܝܬܐ ܕܐܝܟܝܪ", "ACWDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܡܥܪܒܝܬܐ ܡܨܥܝܬܐ ܕܐܘܣܬܪܠܝܐ", "ACWST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܡܥܪܒܝܬܐ ܡܨܥܝܬܐ ܕܐܘܣܬܪܠܝܐ", "ADT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܐܛܠܢܛܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "ADT Arabia": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܪܒܝܐ", "AEDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܣܬܪܠܝܐ ܡܕܢܚܝܬܐ", "AEST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܣܬܪܠܝܐ ܡܕܢܚܝܬܐ", "AFT": "ܥܕܢܐ ܕܐܦܓܐܢܣܬܐܢ", "AKDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܠܐܣܟܐ", "AKST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܠܐܣܟܐ", "AMST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܡܙܢ", "AMST Armenia": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܪܡܢܝܐ", "AMT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܡܙܢ", "AMT Armenia": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܪܡܢܝܐ", "ANAST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܢܐܕܝܪ", "ANAT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܢܐܕܝܪ", "ARST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܪܓܢܬܝܢܐ", "ART": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܪܓܢܬܝܢܐ", "AST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܐܛܠܢܛܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "AST Arabia": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܪܒܝܐ", "AWDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܣܬܪܠܝܐ ܡܥܪܒܝܬܐ", "AWST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܣܬܪܠܝܐ ܡܥܪܒܝܬܐ", "AZST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܙܪܒܝܓܐܢ", "AZT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܙܪܒܝܓܐܢ", "BDT Bangladesh": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܒܢܓܠܐܕܝܫ", "BNT": "ܥܕܢܐ ܕܒܪܘܢܐܝ ܕܐܪܘܣܐܠܡ", "BOT": "ܥܕܢܐ ܕܒܘܠܝܒܝܐ", "BRST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܒܪܐܣܝܠܝܐ", "BRT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܒܪܐܣܝܠܝܐ", "BST Bangladesh": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܒܢܓܠܐܕܝܫ", "BT": "ܥܕܢܐ ܕܒܘܬܐܢ", "CAST": "ܥܕܢܐ ܕܟܐܝܣܝ", "CAT": "ܥܕܢܐ ܕܡܨܥܝܬܐ ܐܦܪܝܩܐ", "CCT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܟܘܟܘܣ", "CDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܡܨܥܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "CHADT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܬܫܐܬܡ", "CHAST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܬܫܐܬܡ", "CHUT": "ܥܕܢܐ ܕܬܫܘܟ", "CKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܓܙܝܪ̈ܐ ܕܟܘܟ", "CKT DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܓܙܝܪ̈ܐ ܕܟܘܟ", "CLST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܬܫܝܠܝ", "CLT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܬܫܝܠܝ", "COST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܟܘܠܘܡܒܝܐ", "COT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܟܘܠܘܡܒܝܐ", "CST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܡܨܥܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "CST China": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܨܝܢ", "CST China DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܨܝܢ", "CVST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܟܐܦ ܒܝܪܕܝ", "CVT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܟܐܦ ܒܝܪܕܝ", "CXT": "ܥܕܢܐ ܕܓܙܪܬܐ ܕܟܪܝܣܬܡܣ", "ChST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܬܫܐܡܘܪܘ", "ChST NMI": "ܥܕܢܐ ܕܓܙܝܖ̈ܐ ܕܡܪܝܐܢܐ ܓܪܒܝܝܬܐ", "CuDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܩܘܒܐ", "CuST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܩܘܒܐ", "DAVT": "ܥܕܢܐ ܕܕܒܝܣ", "DDUT": "ܥܕܢܐ ܕܕܘܡܘܢܬ ܕܐܘܪܒܝܠ", "EASST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܓܙܪܬܐ ܦܨܚܐ", "EAST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܓܙܪܬܐ ܦܨܚܐ", "EAT": "ܥܕܢܐ ܕܡܕܢܚ ܐܦܪܝܩܐ", "ECT": "ܥܕܢܐ ܕܐܩܘܐܕܘܪ", "EDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܡܕܢܚܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "EGDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܓܪܝܢܠܢܕ ܡܕܢܚܝܬܐ", "EGST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܓܪܝܢܠܢܕ ܡܕܢܚܝܬܐ", "EST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܡܕܢܚܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "FEET": "ܥܕܢܐ ܕܐܘܪܘܦܐ (ܗܡ ܡܕܢܚܐ)", "FJT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܦܝܓܝ", "FJT Summer": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܦܝܓܝ", "FKST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܓܙܝܪ̈ܐ ܕܦܠܟܠܢܕ", "FKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܓܙܝܪ̈ܐ ܕܦܠܟܠܢܕ", "FNST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܦܪܢܢܕܘ ܕܢܘܪܘܢܗܐ", "FNT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܦܪܢܢܕܘ ܕܢܘܪܘܢܗܐ", "GALT": "ܥܕܢܐ ܕܓܐܠܦܐܓܘܣ", "GAMT": "ܥܕܢܐ ܕܓܡܒܝܪ", "GEST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܓܘܪܓܝܐ", "GET": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܓܘܪܓܝܐ", "GFT": "ܥܕܢܐ ܕܓܘܝܐܢܐ ܦܪܢܣܝܬܐ", "GIT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܓܝܠܒܪܬ", "GMT": "ܥܕܢܐ ܕܓܪܝܢܟ", "GNSST": "GNSST", "GNST": "GNST", "GST": "ܥܕܢܐ ܕܓܝܘܪܓܝܐ ܬܝܡܢܝܬܐ", "GST Guam": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܓܘܐܡ", "GYT": "ܥܕܢܐ ܕܓܘܝܐܢܐ", "HADT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܗܐܘܐܝܝ ܐܠܘܫܝܢ", "HAST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܗܐܘܐܝܝ ܐܠܘܫܝܢ", "HKST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܗܘܢܓ ܟܘܢܓ", "HKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܗܘܢܓ ܟܘܢܓ", "HOVST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܗܘܒܕ", "HOVT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܗܘܒܕ", "ICT": "ܥܕܢܐ ܕܗܢܕܘܨܝܢ", "IDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܝܣܪܐܝܠ", "IOT": "ܥܕܢܐ ܕܐܘܩܝܢܘܣ ܗܢܕܘܝܐ", "IRKST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܝܪܟܘܬܣܟ", "IRKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܝܪܟܘܬܣܟ", "IRST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܝܪܐܢ", "IRST DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܝܪܐܢ", "IST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܗܢܕܘ", "IST Israel": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܝܣܪܐܝܠ", "JDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܝܦܢ", "JST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܝܦܢ", "KOST": "ܥܕܢܐ ܟܘܣܪܐܝ", "KRAST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܟܪܐܣܢܘܝܪܣܟ", "KRAT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܟܪܐܣܢܘܝܪܣܟ", "KST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܟܘܪܝܝܐ", "KST DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܟܘܪܝܝܐ", "LHDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܠܘܪܕ ܗܐܘ", "LHST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܠܘܪܕ ܗܐܘ", "LINT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܠܐܝܢ", "MAGST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܡܐܓܕܐܢ", "MAGT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܡܐܓܕܐܢ", "MART": "ܥܕܢܐ ܕܡܐܪܟܐܘܣܐܣ", "MAWT": "ܥܕܢܐ ܕܡܐܘܣܘܢ", "MDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܛܘܪܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "MESZ": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܪܘܦܐ ܡܨܥܝܬܐ", "MEZ": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܪܘܦܐ ܡܨܥܝܬܐ", "MHT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܡܐܪܫܐܠ", "MMT": "ܥܕܢܐ ܕܡܝܐܢܡܐܪ", "MSD": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܡܘܣܟܘ", "MST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܛܘܪܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "MUST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܡܘܪܝܛܝܘܣ", "MUT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܡܘܪܝܛܝܘܣ", "MVT": "ܥܕܢܐ ܕܓܙܪܬܐ ܡܐܠܕܝܒܝܬܐ", "MYT": "ܥܕܢܐ ܕܡܠܝܙܝܐ", "NCT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܢܝܘ ܟܠܝܕܘܢܝܐ", "NDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܢܝܘܦܐܘܢܠܢܕ", "NDT New Caledonia": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܢܝܘ ܟܠܝܕܘܢܝܐ", "NFDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܓܙܪܬܐ ܕܢܘܪܦܠܟ", "NFT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܓܙܪܬܐ ܕܢܘܪܦܠܟ", "NOVST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܢܘܒܘܣܝܒܪܣܟ", "NOVT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܢܘܒܘܣܝܒܪܣܟ", "NPT": "ܥܕܢܐ ܕܢܝܦܐܠ", "NRT": "ܥܕܢܐ ܕܢܐܘܪܘ", "NST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܢܝܘܦܐܘܢܠܢܕ", "NUT": "ܥܕܢܐ ܕܢܝܘܝ", "NZDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܢܝܘ ܙܝܠܢܕ", "NZST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܢܝܘ ܙܝܠܢܕ", "OESZ": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܪܘܦܐ ܡܕܢܚܝܬܐ", "OEZ": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܪܘܦܐ ܡܕܢܚܝܬܐ", "OMSST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܘܡܣܟ", "OMST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܘܡܣܟ", "PDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܫܝܢܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "PDTM": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܡܟܣܝܩܘ ܫܝܢܝܬܐ", "PETDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܦܝܬܪܘܦܒܠܒܣܟܝ-ܟܐܡܟܬܣܟܝ", "PETST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܦܝܬܪܘܦܒܠܒܣܟܝ-ܟܐܡܟܬܣܟܝ", "PGT": "ܥܕܢܐ ܕܦܐܦܘܐ ܓܝܢܝܐ ܚܕܬܐ", "PHOT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܦܝܢܝܟܣ", "PKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܦܐܟܣܬܐܢ", "PKT DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܦܐܟܣܬܐܢ", "PMDT": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܣܐܢܬ ܦܝܥܪ ܘܡܩܘܠܘܢ", "PMST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܣܐܢܬ ܦܝܥܪ ܘܡܩܘܠܘܢ", "PONT": "ܥܕܢܐ ܕܦܘܢܐܦܝ", "PST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܫܝܢܝܬܐ ܕܐܡܪܝܟܐ ܓܪܒܝܝܬܐ", "PST Philippine": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܦܝܠܝܦܝܢܝܐ", "PST Philippine DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܦܝܠܝܦܝܢܝܐ", "PST Pitcairn": "ܥܕܢܐ ܕܦܝܬܟܐܝܪܢ", "PSTM": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܡܟܣܝܩܘ ܫܝܢܝܬܐ", "PWT": "ܥܕܢܐ ܕܦܠܐܘ", "PYST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܦܪܓܘܐܝ", "PYT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܦܪܓܘܐܝ", "PYT Korea": "ܥܕܢܐ ܕܦܝܘܢܓܝܢܓ", "RET": "ܥܕܢܐ ܕܪܝܘܢܝܘܢ", "ROTT": "ܥܕܢܐ ܕܪܘܬܝܪܐ", "SAKST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܣܐܚܐܠܝܢ", "SAKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܣܐܚܐܠܝܢ", "SAMST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܣܡܐܪܐ", "SAMT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܣܡܐܪܐ", "SAST": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܬܝܡܢ ܐܦܪܝܩܐ", "SBT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܫܠܝܡܘܢ", "SCT": "ܥܕܢܐ ܕܣܐܝܫܝܠ", "SGT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܣܝܢܓܐܦܘܪ", "SLST": "ܥܕܢܐ ܕܠܐܢܟܐ", "SRT": "ܥܕܢܐ ܕܣܘܪܝܢܐܡ", "SST Samoa": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܣܡܘܐ", "SST Samoa Apia": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܦܝܐ", "SST Samoa Apia DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܦܝܐ", "SST Samoa DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܣܡܘܐ", "SYOT": "ܥܕܢܐ ܕܣܝܘܐ", "TAAF": "ܥܕܢܐ ܕܦܪܢܣܐ ܬܝܡܢܝܬܐ ܘܐܢܬܪܬܝܟܐ", "TAHT": "ܥܕܢܐ ܕܬܗܝܬܝ", "TJT": "ܥܕܢܐ ܕܬܐܓܝܟܣܬܐܢ", "TKT": "ܥܕܢܐ ܕܬܘܟܝܠܐܘ", "TLT": "ܥܕܢܐ ܕܡܕܢܚ ܬܝܡܘܪ", "TMST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܬܘܪܟܡܢܣܬܐܢ", "TMT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܬܘܪܟܡܢܣܬܐܢ", "TOST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܬܘܢܓܐ", "TOT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܬܘܢܓܐ", "TVT": "ܥܕܢܐ ܕܬܘܒܐܠܘ", "TWT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܬܐܝܦܐܝ", "TWT DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܬܐܝܦܐܝ", "ULAST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܠܐܢܒܐܬܘܪ", "ULAT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܠܐܢܒܐܬܘܪ", "UYST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܪܘܓܘܐܝ", "UYT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܪܘܓܘܐܝ", "UZT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܙܒܟܣܬܐܢ", "UZT DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܙܒܟܣܬܐܢ", "VET": "ܥܕܢܐ ܕܒܢܙܘܝܠܐ", "VLAST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܒܠܐܕܝܒܘܣܬܘܟ", "VLAT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܒܠܐܕܝܒܘܣܬܘܟ", "VOLST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܒܘܠܓܘܓܪܐܕ", "VOLT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܒܘܠܓܘܓܪܐܕ", "VOST": "ܥܕܢܐ ܕܒܘܣܬܘܟ", "VUT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܒܐܢܘܐܛܘ", "VUT DST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܒܐܢܘܐܛܘ", "WAKT": "ܥܕܢܐ ܕܓܙܝܪ̈ܐ ܕܘܐܝܟ", "WARST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܪܓܢܬܝܢܐ ܡܥܪܒܝܬܐ", "WART": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܪܓܢܬܝܢܐ ܡܥܪܒܝܬܐ", "WAST": "ܥܕܢܐ ܕܡܥܪܒ ܐܦܪܝܩܐ", "WAT": "ܥܕܢܐ ܕܡܥܪܒ ܐܦܪܝܩܐ", "WESZ": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܘܪܘܦܐ ܡܥܪܒܝܬܐ", "WEZ": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܘܪܘܦܐ ܡܥܪܒܝܬܐ", "WFT": "ܥܕܢܐ ܕܘܝܠܝܣ ܘܦܘܬܘܢܐ", "WGST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܓܪܝܢܠܢܕ ܕܡܥܪܒܝܬܐ", "WGT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܓܪܝܢܠܢܕ ܕܡܥܪܒܝܬܐ", "WIB": "ܥܕܢܐ ܕܗܢܕܘܨܝܢ ܡܥܪܒܝܬܐ", "WIT": "ܥܕܢܐ ܕܗܢܕܘܨܝܢ ܡܕܢܚܝܬܐ", "WITA": "ܥܕܢܐ ܕܗܢܕܘܨܝܢ ܡܨܥܝܬܐ", "YAKST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܝܐܟܘܬܣܟ", "YAKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܝܐܟܘܬܣܟ", "YEKST": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܝܟܐܬܝܪܢܒܝܪܓ", "YEKT": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܝܟܐܬܝܪܢܒܝܪܓ", "YST": "ܥܕܢܐ ܕܝܘܩܘܢ", "МСК": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܡܘܣܟܘ", "اقتاۋ": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܟܬܐܘ", "اقتاۋ قالاسى": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܟܬܐܘ", "اقتوبە": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܟܬܘܒ", "اقتوبە قالاسى": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܟܬܘܒ", "الماتى": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܐܠܡܐܬܝ", "الماتى قالاسى": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܠܡܐܬܝ", "باتىس قازاق ەلى": "ܥܕܢܐ ܕܡܥܪܒ ܟܙܩܣܬܐܢ", "شىعىش قازاق ەلى": "ܥܕܢܐ ܕܡܕܢܚ ܟܙܩܣܬܐܢ", "قازاق ەلى": "ܥܕܢܐ ܕܟܙܩܣܬܐܢ", "قىرعىزستان": "ܥܕܢܐ ܕܩܝܪܓܝܙܣܬܐܢ", "قىزىلوردا": "ܥܕܢܐ ܡܫܘܚܬܢܝܬܐ ܕܟܝܙܝܠܘܪܕܐ", "قىزىلوردا قالاسى": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܟܝܙܝܠܘܪܕܐ", "∅∅∅": "ܥܕܢܐ ܕܒܗܪ ܝܘܡܐ ܕܐܙܘܪ"},
	}
}

// Locale returns the current translators string locale
func (syr *syr) Locale() string {
	return syr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'syr'
func (syr *syr) PluralsCardinal() []locales.PluralRule {
	return syr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'syr'
func (syr *syr) PluralsOrdinal() []locales.PluralRule {
	return syr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'syr'
func (syr *syr) PluralsRange() []locales.PluralRule {
	return syr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'syr'
func (syr *syr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'syr'
func (syr *syr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'syr'
func (syr *syr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (syr *syr) MonthAbbreviated(month time.Month) string {
	if len(syr.monthsAbbreviated) == 0 {
		return ""
	}
	return syr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (syr *syr) MonthsAbbreviated() []string {
	return syr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (syr *syr) MonthNarrow(month time.Month) string {
	if len(syr.monthsNarrow) == 0 {
		return ""
	}
	return syr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (syr *syr) MonthsNarrow() []string {
	return syr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (syr *syr) MonthWide(month time.Month) string {
	if len(syr.monthsWide) == 0 {
		return ""
	}
	return syr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (syr *syr) MonthsWide() []string {
	return syr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (syr *syr) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(syr.daysAbbreviated) == 0 {
		return ""
	}
	return syr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (syr *syr) WeekdaysAbbreviated() []string {
	return syr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (syr *syr) WeekdayNarrow(weekday time.Weekday) string {
	if len(syr.daysNarrow) == 0 {
		return ""
	}
	return syr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (syr *syr) WeekdaysNarrow() []string {
	return syr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (syr *syr) WeekdayShort(weekday time.Weekday) string {
	if len(syr.daysShort) == 0 {
		return ""
	}
	return syr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (syr *syr) WeekdaysShort() []string {
	return syr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (syr *syr) WeekdayWide(weekday time.Weekday) string {
	if len(syr.daysWide) == 0 {
		return ""
	}
	return syr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (syr *syr) WeekdaysWide() []string {
	return syr.daysWide
}

// Decimal returns the decimal point of number
func (syr *syr) Decimal() string {
	return syr.decimal
}

// Group returns the group of number
func (syr *syr) Group() string {
	return syr.group
}

// Group returns the minus sign of number
func (syr *syr) Minus() string {
	return syr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'syr' and handles both Whole and Real numbers based on 'v'
func (syr *syr) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, syr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, syr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, syr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'syr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (syr *syr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, syr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, syr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, syr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'syr'
func (syr *syr) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := syr.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'syr'
// in accounting notation.
func (syr *syr) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := syr.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'syr'
func (syr *syr) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'syr'
func (syr *syr) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xdc, 0x92}...)
	b = append(b, syr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'syr'
func (syr *syr) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xdc, 0x92}...)
	b = append(b, syr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'syr'
func (syr *syr) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, syr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xdc, 0x92}...)
	b = append(b, syr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'syr'
func (syr *syr) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, syr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, syr.periodsAbbreviated[0]...)
	} else {
		b = append(b, syr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'syr'
func (syr *syr) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, syr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, syr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, syr.periodsAbbreviated[0]...)
	} else {
		b = append(b, syr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'syr'
func (syr *syr) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, syr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, syr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, syr.periodsAbbreviated[0]...)
	} else {
		b = append(b, syr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'syr'
func (syr *syr) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, syr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, syr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, syr.periodsAbbreviated[0]...)
	} else {
		b = append(b, syr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := syr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
