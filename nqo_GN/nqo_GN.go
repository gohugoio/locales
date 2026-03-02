package nqo_GN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type nqo_GN struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'nqo_GN' locale
func New() locales.Translator {
	return &nqo_GN{
		locale:             "nqo_GN",
		pluralsCardinal:    []locales.PluralRule{6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              "،",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "ߓߌ߲ߠ", "ߞߏ߲ߞ", "ߕߙߊ", "ߞߏ߲ߘ", "ߘߓߊ߬ߕ", "ߥߊ߬ߛ", "ߞߊ߬ߙ", "ߘߓߊ߬ߓ", "ߕߎߟߊߝߌ߲", "ߞߏ߲ߓ", "ߣߍߣ", "ߞߏߟ"},
		monthsNarrow:       []string{"", "ߓ", "ߞ", "ߕ", "ߞ", "ߘ", "ߥ", "ߞ", "ߘ", "ߕ", "ߞ", "ߣ", "ߞ"},
		monthsWide:         []string{"", "ߓߌ߲ߠߊߥߎߟߋ߲", "ߞߏ߲ߞߏߜߍ", "ߕߙߊߓߊ", "ߞߏ߲ߞߏߘߌ߬ߓߌ", "ߘߓߊ߬ߕߊ", "ߥߊ߬ߛߌ߬ߥߙߊ", "ߞߊ߬ߙߌߝߐ߭", "ߘߓߊ߬ߓߌߟߊ", "ߕߎߟߊߝߌ߲", "ߞߏ߲ߓߌߕߌ߮", "ߣߍߣߍߓߊ", "ߞߏߟߌ߲ߞߏߟߌ߲"},
		daysAbbreviated:    []string{"ߞߊ߯ߙ", "ߞߐ߬ߓ", "ߞߐ߬ߟߏ߲", "ߞߎߣ", "ߓߌߟ", "ߛߌ߬ߣ", "ߞߍ߲ߘ"},
		daysNarrow:         []string{"ߞ", "ߞ", "ߞ", "ߞ", "ߓ", "ߛ", "ߞ"},
		daysShort:          []string{"ߞߊ߯", "ߞߐ߬", "ߞߐ߬ߟߏ߲", "ߞߎ", "ߓߌ", "ߛߌ߬", "ߞߍ߲"},
		daysWide:           []string{"ߞߊ߯ߙߌߟߏ߲", "ߞߐ߬ߓߊ߬ߟߏ߲", "ߞߐ߬ߟߏ߲", "ߞߎߣߎ߲ߟߏ߲", "ߓߌߟߏ߲", "ߛߌ߬ߣߌ߲߬ߟߏ߲", "ߞߍ߲ߘߍߟߏ߲"},
		periodsAbbreviated: []string{"ߛ", "ߥ"},
		erasAbbreviated:    []string{"ߌߛ. ߡ. ߢߍ߫", "ߌߛ. ߡ. ߞߐ߫"},
		erasNarrow:         []string{"ߌߛ. ߢߍ߫", "ߌߛ. ߞߐ߫"},
		erasWide:           []string{"ߌߛߊ߫ ߡߏߦߌ ߢߍ߫", "ߌߛߊ߫ ߡߏߦߌ ߞߐ߫"},
		timezones:          map[string]string{"ACDT": "ߕߍߡߟߊ ߏߛߑߕߙߊߟߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "ACST": "ߕߍߡߟߊ ߏߛߑߕߙߊߟߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "ACT": "ACT", "ACWDT": "ߏߛߑߕߙߊߟߌ߫ ߕߟߋ߬ߓߋ߬-ߕߍߡߟߊ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "ACWST": "ߏߛߑߕߙߊߟߌ߫ ߕߟߋ߬ߓߋ߬-ߕߍߡߟߊ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "ADT": "ߟߌ߲ߓߊ߲-ߡߊ߲ߞߊ߲ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "ADT Arabia": "ߊߙߊߓߎߟߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AEDT": "ߕߟߋ߬ߓߐ ߏߛߑߕߙߊߟߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "AEST": "ߕߟߋ߬ߓߐ ߏߛߑߕߙߊߟߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AFT": "ߊߝߑߜ߭ߊߣߌߛߑߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "AKDT": "ߊߟߊߛߑߞߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "AKST": "ߊߟߊߛߑߞߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AMST": "ߊߡߊߖ߭ߏ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "AMST Armenia": "ߊߙߑߡߋߣߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AMT": "ߊߡߊߖ߭ߏߣ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AMT Armenia": "ߊߙߑߡߋߣߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ߊߙߑߖ߭ߊ߲ߕߌ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "ART": "ߊߙߑߖ߭ߊ߲ߕߌ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AST": "ߟߌ߲ߓߊ߲-ߡߊ߲ߞߊ߲ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AST Arabia": "ߊߙߊߓߎߟߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AWDT": "ߕߟߋ߬ߓߋ ߏߛߑߕߙߊߟߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "AWST": "ߕߟߋ߬ߓߋ ߏߛߑߕߙߊߟߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AZST": "ߊߖ߭ߍߙߑߓߊߦߑߖߊ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ ߛߎߡߊ߲ߘߊ߲ߕߊ", "AZT": "ߊߖ߭ߍߙߑߓߊߦߑߖߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "BDT Bangladesh": "ߓߊ߲ߜ߭ߑߟߊߘߍߛ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "BNT": "ߓߙߎߣߋߦߌ߫ ߘߊߙߎߛߑߛߟߊߡ ߕߎ߬ߡߊ߬ߘߊ", "BOT": "ߓߏߟߝ߭ߌ߫ ߕߎ߬ߡߊ߬ߘߊ", "BRST": "ߓߙߋߖ߭ߟߌߟߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "BRT": "ߓߙߋߖ߭ߟߌߟߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "BST Bangladesh": "ߓߊ߲ߜ߭ߑߟߊߘߍߛ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "BT": "ߓߎߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "CAST": "CAST", "CAT": "ߝߘߊ߬ߝߌ߲߬ߠߊ߫ ߕߊ߲ߓߊ߲ ߕߎ߬ߡߊ߬ߙߋ߲ ߢߊߓߘߍ", "CCT": "ߞߏߞߏߛ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "CDT": "ߞߐ߰ߘߎ߮ ߊߡߋߙߌߞߌ߬ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "CHADT": "ߗߕߊߡ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "CHAST": "ߗߕߊߡ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "CHUT": "ߗߎߞ ߕߎ߬ߡߊ߬ߘߊ", "CKT": "ߞߏߞ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "CKT DST": "ߞߏߞ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "CLST": "ߛߟߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "CLT": "ߛߟߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "COST": "ߞߐߟߐ߲ߓߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "COT": "ߞߐߟߐ߲ߓߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "CST": "ߞߐ߰ߘߎ߮ ߊߡߋߙߌߞߌ߬ ߕߎ߬ߡߘߊ߬ ߕߊ߲ߓߊ߲߫ ߛߎߡߊ߲ߘߊ߲ߕߊ", "CST China": "ߛߌߣ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "CST China DST": "ߛߌߣ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "CVST": "ߜߙߋߞߎ߲߫-ߝߙߌߛߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߊߙߋ߲", "CVT": "ߜߙߋߞߎ߲߫-ߝߙߌߛߌ߫ ߕߎ߬ߡߊ߬ߙߋ߲ ߢߊߓߘߍ", "CXT": "ߞߑߙߌߛߑߡߊߛ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "ChST": "ߛ߭ߊߡߏߙߏ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "ChST NMI": "ChST NMI", "CuDT": "ߞߎ߱ߓߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "CuST": "ߞߎ߳ߓߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "DAVT": "ߘߊߝ߭ߌߛ ߕߎ߬ߡߊ߬ߘߊ", "DDUT": "ߘߎߡߐ߲-ߘߎߙߑߝ߭ߌߟ ߕߎ߬ߡߊ߬ߘߊ", "EASST": "ߌߛߑߟߊ߲ߘ ߕߌ߲ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "EAST": "ߌߛߑߟߊ߲ߘ ߕߌ߲ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "EAT": "ߝߘߊ߬ߌ߲߬ߠߊ߫ ߓߟߋ߬ߓߐ ߕߎ߬ߡߊ߬ߙߋ߲ ߢߊߓߘߍ", "ECT": "ߌߞߑߥߊߘߐߙ ߕߎ߬ߡߊ߬ߘߊ", "EDT": "ߞߐ߰ߘߎ߰ ߊߡߋߙߌߞߌ߬ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "EGDT": "ߕߟߋ߬ߓߐ ߜ߭ߌߙߌ߲ߟߊ߲ߘ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "EGST": "ߕߟߋ߬ߓߐ ߜ߭ߌߙߌ߲ߟߊ߲ߘ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "EST": "ߞߐ߰ߘߎ߰ ߊߡߋߙߌߞߌ߬ ߕߟߋ߬ߓߐ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "FEET": "ߞߊߟߌߣߌ߲ߜ߭ߙߊߘ ߕߎ߬ߡߊ߬ߘߊ", "FJT": "ߝߖߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "FJT Summer": "ߝߖߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "FKST": "ߝߊߟߞߑߟߊ߲ߘ ߕߌ߲ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "FKT": "ߝߊߟߞߑߟߊ߲ߘ ߕߌ߲ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "FNST": "ߝߍߙߑߣߊ߲ߘߏ߫ ߘߋ߬ ߣߏߙߏ߲ߤߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "FNT": "ߝߍߙߑߣߊ߲ߘߏ߫ ߘߋ߬ ߣߏߙߏ߲ߤߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "GALT": "ߜ߭ߟߊߔߊߜ߭ߐߛ ߕߎ߬ߡߊ߬ߘߊ߫", "GAMT": "ߜ߭ߊ߲ߓߌߦߍߙ ߕߎ߬ߡߊ߬ߘߊ", "GEST": "ߖ߭ߐߙߑߖ߭ߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ ߛߎߡߊ߲ߘߊ߲ߕߊ", "GET": "ߖ߭ߐߙߑߖ߭ߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "GFT": "ߝߊ߬ߙߊ߲߬ߛߌ߬ ߜ߭ߎߦߣߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "GIT": "ߖ߭ߌߟߑߓߍߙߕ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "GMT": "ߜ߭ߙߋߣߍߕ ߕߎ߬ߡߊ߬ߘߊ", "GNSST": "GNSST", "GNST": "GNST", "GST": "ߖߌߘߏ߮ ߕߎ߬ߡߊ߬ߘߊ", "GST Guam": "GST Guam", "GYT": "ߜ߭ߎߦߣߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "HADT": "ߤߥߊߦ - ߊߟߋߦߎߕߌߦߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "HAST": "ߤߥߊߦ - ߊߟߋߦߎߕߌߦߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "HKST": "ߤߐ߲ߜ߭ ߞߐ߲ߜ߭ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "HKT": "ߤߐ߲ߜ߭ ߞߐ߲ߜ߭ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "HOVST": "ߤߏߝ߭ߘ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "HOVT": "ߤߏߝ߭ߘ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "ICT": "ߤߌ߲ߘߏ߫ ߛߌߣ ߕߎ߬ߡߊ߬ߘߊ", "IDT": "ߌߛߌ߬ߙߊ߬ߦߌߟߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ ߛߎߡߊ߲ߘߊ߲ߕߊ", "IOT": "ߍ߲ߘߎ߫ ߟߌ߲ߓߊ߲ߘߎ߯ ߕߎ߬ߡߊ߬ߙߋ߲", "IRKST": "ߌߙߑߞߎߕߑߛߞ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "IRKT": "ߌߙߑߞߎߕߑߛߞ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "IRST": "ߌߙߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "IRST DST": "ߌߙߊ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "IST": "ߤߌ߲ߘߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "IST Israel": "ߌߛߌ߬ߙߊ߬ߦߌߟߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "JDT": "ߖ߭ߊߔߐ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "JST": "ߖ߭ߊߔߐ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "KOST": "ߞߏߛߑߙߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "KRAST": "ߞߙߊߛߑߣߏߦߌߦߊߙߑߛߞ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "KRAT": "ߞߙߊߛߑߣߏߦߌߦߊߙߑߛߞ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "KST": "ߞߏߙߋ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "KST DST": "ߞߏߙߋ߫ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "LHDT": "ߟߐߙߘ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "LHST": "ߟߐߙߘ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "LINT": "ߌߛߑߟߊ߲ߘ ߞߍ߬ߙߍ߲߬ߘߍ ߕߎ߬ߡߊ߬ߘߊ", "MAGST": "ߡߜ߭ߊߘߊ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "MAGT": "ߡߜ߭ߊߘߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "MART": "ߡߊߙߞߌߛߊߛ ߕߎ߬ߡߊ߬ߘߊ", "MAWT": "ߡߊߥߑߛߐ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "MDT": "MDT", "MESZ": "ߕߍߡߊ߫ ߋߙߐߔߎ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "MEZ": "ߕߍߡߊ߫ ߋߙߐߔߎ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "MHT": "ߡߊߙߑߛߊߟ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "MMT": "ߡߌߦߊߣߑߡߊߙ ߕߎ߬ߡߊ߬ߘߊ", "MSD": "ߡߏߛߑߞߎ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "MST": "MST", "MUST": "ߡߏߙߌߛ ߕߟߋ߬ߡߊ߬ ߕߎߡߊߙߋ߲", "MUT": "ߡߏߙߌߛ ߕߎ߬ߡߊ߬ߙߋ߲ ߢߊߓߘߍ", "MVT": "ߡߊߟߑߘߌߝ߭ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "MYT": "ߡߊߟߍߘߎ߯ ߕߎ߬ߡߊ߬ߘߊ", "NCT": "ߞߊߟߌߘߏߣߌ߫ ߞߎߘߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "NDT": "ߝߎ߲ߘߑߟߊ߲ߘ ߞߎߘߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "NDT New Caledonia": "ߞߊߟߌߘߏߣߌ߫ ߞߎߘߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "NFDT": "ߣߐߙ ߝߐߟߞ ߕߌ߲ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "NFT": "ߣߐߙ ߝߐߟߞ ߕߌ߲ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "NOVST": "ߣߝ߭ߏߛߓߌߙߑߛߞ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "NOVT": "ߣߝ߭ߏߛߓߌߙߑߛߞ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "NPT": "ߣߋߔߊߟ ߕߎ߬ߡߊ߬ߘߊ", "NRT": "ߣߊߥߙߎ߫ ߕߎ߬ߡߊ߬ߘߊ", "NST": "ߝߎ߲ߘߑߟߊ߲ߘ ߞߎߘߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "NUT": "ߣߌߦߋ߫ ߕߎ߬ߡߊ߬ߘߊ", "NZDT": "ߣߌߦߎ-ߖ߭ߋߟߊ߲ߘ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "NZST": "ߣߌߦߎ-ߖ߭ߋߟߊ߲ߘ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "OESZ": "ߕߟߋ߬ߓߐ ߋߙߐߔߎ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "OEZ": "ߕߟߋ߬ߓߐ ߋߙߐߔߎ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "OMSST": "ߏߡߑߛߞ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "OMST": "ߏߡߑߛߞ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PDT": "ߞߐ߰ߘߎ߰ ߊߡߋߙߌߞߌ߬ ߣߌ߫ ߖߐ߫ ߟߌ߲ߓߊ߲ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "PDTM": "ߖߐ߫ ߟߌ߲ߓߊ߲ ߡߍ߲ߞߛߌߞߌ߬ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ߔߊߓߎߥߊ߫ ߖߌ߬ߣߍ߬ ߞߎߘߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "PHOT": "ߝߏߣߌߞߛ ߕߌ߲ ߠߎ߫ ߕߎ߬ߡߊ߬ߘߊ", "PKT": "ߔߊߞߌߛߑߕߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PKT DST": "ߔߊߞߌߛߑߕߊ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "PMDT": "ߛߍ߲-ߔߌߦߍߙ-ߋ-ߡߌߞߋߟߐ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "PMST": "ߛߍ߲-ߔߌߦߍߙ-ߋ-ߡߌߞߋߟߐ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PONT": "ߔߏ߲ߔߍ߫ ߕߎ߬ߡߊ߬ߘߊ", "PST": "ߞߐ߰ߘߎ߰ ߊߡߋߙߌߞߌ߬ ߣߌ߫ ߖߐ߫ ߟߌ߲ߓߊ߲ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PST Philippine": "ߝߟߌߔߌ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PST Philippine DST": "ߝߟߌߔߌ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "PST Pitcairn": "ߔߌߕߑߞߍߙߣ ߕߎ߬ߡߊ߬ߘߊ", "PSTM": "ߖߐ߫ ߟߌ߲ߓߊ߲ ߡߍ߲ߞߛߌߞߌ߬ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PWT": "ߔߟߊߐߛ ߕߎ߬ߡߊ߬ߘߊ", "PYST": "ߔߙߊߜ߭ߎߥߋ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "PYT": "ߔߙߊߜ߭ߎߥߋ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "PYT Korea": "ߔߌߦߐ߲ߜ߭ߦߊ߲ߜ߭ ߕߎ߬ߡߊ߬ߘߊ", "RET": "ߙߋߣߌߦߐ߲߫ ߕߎ߬ߡߊ߬ߙߋ߲", "ROTT": "ߙߏߕߋߙߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "SAKST": "ߛߊߞߑߤߊߟߌߣ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "SAKT": "ߛߊߞߑߤߊߟߌߣ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ߝߘߌ߬ߝߌ߲߬ߠߊ߫ ߥߙߏ߬ߘߎ߮ ߕߎ߬ߡߊ߬ߙߋ߲ ߢߊߓߘߍ", "SBT": "ߛߟߏ߬ߡߣߊ߬ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "SCT": "ߛߋߦߌߛߌߟ ߕߎ߬ߡߊ߬ߙߋ߲", "SGT": "ߛߍ߲ߜ߭ߊߔߎߙ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "SLST": "SLST", "SRT": "ߛߎߙߌߣߊߡ ߕߎ߬ߡߊ߬ߘߊ", "SST Samoa": "ߛߊߡߏߥߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "SST Samoa Apia": "ߊߔߌߦߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "SST Samoa Apia DST": "ߊߔߌߦߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "SST Samoa DST": "ߛߊߡߏߥߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "SYOT": "ߛߦߏߥߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "TAAF": "ߐߛߑߕߙߊߟߌ߫ ߘߎ߰ߞߟߏ ߣߌ߫ ߝߊ߬ߙߊ߲߬ߛߌ߫ ߊ߲ߕߊߙߑߕߌߞ ߕߎ߬ߡߊ߬ߙߋ", "TAHT": "ߕߊߤߕߌ߫ ߕߎ߬ߡߊ߬ߘߊ", "TJT": "ߕߊߖߞߌߛߑߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "TKT": "ߕߏߞߋߟߏ߫ ߕߎ߬ߡߊ߬ߘߊ", "TLT": "ߕߟߋ߬ߓߐ ߕߌߡߐߙ ߕߎ߬ߡߊ߬ߘߊ", "TMST": "ߕߎߙߞߑߡߋߣߌߛߑߕߊ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "TMT": "ߕߎߙߞߌߡߋߣߌߛߑߕߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "TOST": "ߕߏ߲ߜ߭ߊ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "TOT": "ߕߏ߲ߜ߭ߊ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "TVT": "ߕߎߝ߭ߊߟߎ߫ ߕߎ߬ߡߊ߬ߘߊ", "TWT": "ߕߊߦߌߔߋ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "TWT DST": "ߕߊߦߌߔߋ߫ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "ULAST": "ߎߟߊ߲ߓߕߐߙ ߕߟߋ߬ߡߊ߬ ߕߎ߬ߡߊ߬ߘߊ", "ULAT": "ߎߟߊ߲ߓߕߐߙ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "UYST": "ߎ߳ߙߑߜ߭ߋߦߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "UYT": "ߎ߳ߙߑߜ߭ߋߦߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "UZT": "ߏߖ߭ߑߓߊߞߌߛߑߕߊ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "UZT DST": "ߏߖ߭ߑߓߊߞߌߛߑߕߊ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "VET": "ߝ߭ߣߋߖ߭ߎߦߋߟߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "VLAST": "ߝ߭ߟߊߘߑߝ߭ߏߛߑߕߏߞ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "VLAT": "ߝ߭ߟߊߘߑߝ߭ߏߛߑߕߏߞ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "VOLST": "ߝ߭ߏߟߑߜ߭ߏߜ߭ߙߊߘ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "VOLT": "ߝ߭ߏߟߑߜ߭ߏߜ߭ߙߊߘ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "VOST": "ߝ߭ߐߛߑߕߐߞ ߕߎ߬ߡߊ߬ߘߊ", "VUT": "ߝ߭ߊߣߎߥߊߕߎ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "VUT DST": "ߝ߭ߊߣߎߥߊߕߎ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "WAKT": "ߥߊߞ ߕߌ߲ ߕߎ߬ߡߊ߬ߘߊ", "WARST": "ߕߟߋ߬ߓߋ ߊߙߑߖ߭ߊ߲ߕߌ߲߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "WART": "ߕߟߋ߬ߓߋ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "WAST": "ߝߘߊ߬ߝߌ߲߬ߠߊ߫ ߕߟߋ߬ߓߋ ߕߎ߬ߡߊ߬ߙߋ߲", "WAT": "ߝߘߊ߬ߝߌ߲߬ߠߊ߫ ߕߟߋ߬ߓߋ ߕߎ߬ߡߊ߬ߙߋ߲", "WESZ": "ߕߟߋ߬ߓߋ ߋߙߐߔߎ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "WEZ": "ߕߟߋ߬ߓߋ ߋߙߐߔߎ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "WFT": "ߥߊߟߌߛ ߣߌ߫ ߝߕߎߣߊ߫ ߕߎ߬ߡߊ߬ߘߊ", "WGST": "ߕߟߋ߬ߓߋ ߜ߭ߌߙߌ߲ߟߊ߲ߘ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "WGT": "ߕߟߋ߬ߓߋ ߜ߭ߌߙߌ߲ߟߊ߲ߘ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "WIB": "ߕߟߋ߬ߓߋ ߍ߲ߘߋߣߏߛߌ߫ ߕߎ߬ߡߊ߬ߘߊ", "WIT": "ߕߟߋ߬ߓߐ ߍ߲ߘߋߣߏߛߌ߫ ߕߎ߬ߡߊ߬ߘߊ", "WITA": "ߕߍߡߊ߫ ߍ߲ߘߋߣߏߛߌ߫ ߕߎ߬ߡߊ߬ߘߊ", "YAKST": "ߦߊߞߎߕߑߛߞߌ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "YAKT": "ߦߊߞߎߕߑߛߞߌ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "YEKST": "ߦߋߞߊߕߋߙߌ߲ߓߎߙߜ߭ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ", "YEKT": "ߦߋߞߊߕߋߙߌ߲ߓߎߙߜ߭ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "YST": "ߦߎߞߐ߲߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "МСК": "ߡߏߛߑߞߎ߫ ߕߎ߬ߡߘߊ߬ ߛߎߡߊ߲ߘߊ߲ߕߊ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ߕߟߋ߬ߓߋ ߞߖ߭ߊߞߌߛߑߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "شىعىش قازاق ەلى": "ߕߟߋ߬ߓߐ ߞߖ߭ߊߞߌߛߑߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "قازاق ەلى": "ߞߖ߭ߊߞߌߛߑߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "قىرعىزستان": "ߜ߭ߌߙߑߜ߭ߌߛߑߕߊ߲߫ ߕߎ߬ߡߊ߬ߘߊ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ߔߋߙߎ߫ ߕߟߋ߬ߡߊ߬ ߕߎߡߘߊ"},
	}
}

// Locale returns the current translators string locale
func (nqo *nqo_GN) Locale() string {
	return nqo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nqo_GN'
func (nqo *nqo_GN) PluralsCardinal() []locales.PluralRule {
	return nqo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nqo_GN'
func (nqo *nqo_GN) PluralsOrdinal() []locales.PluralRule {
	return nqo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nqo_GN'
func (nqo *nqo_GN) PluralsRange() []locales.PluralRule {
	return nqo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nqo_GN'
func (nqo *nqo_GN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nqo_GN'
func (nqo *nqo_GN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nqo_GN'
func (nqo *nqo_GN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nqo *nqo_GN) MonthAbbreviated(month time.Month) string {
	return nqo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nqo *nqo_GN) MonthsAbbreviated() []string {
	return nqo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nqo *nqo_GN) MonthNarrow(month time.Month) string {
	return nqo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nqo *nqo_GN) MonthsNarrow() []string {
	return nqo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nqo *nqo_GN) MonthWide(month time.Month) string {
	return nqo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nqo *nqo_GN) MonthsWide() []string {
	return nqo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nqo *nqo_GN) WeekdayAbbreviated(weekday time.Weekday) string {
	return nqo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nqo *nqo_GN) WeekdaysAbbreviated() []string {
	return nqo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nqo *nqo_GN) WeekdayNarrow(weekday time.Weekday) string {
	return nqo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nqo *nqo_GN) WeekdaysNarrow() []string {
	return nqo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nqo *nqo_GN) WeekdayShort(weekday time.Weekday) string {
	return nqo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nqo *nqo_GN) WeekdaysShort() []string {
	return nqo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nqo *nqo_GN) WeekdayWide(weekday time.Weekday) string {
	return nqo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nqo *nqo_GN) WeekdaysWide() []string {
	return nqo.daysWide
}

// Decimal returns the decimal point of number
func (nqo *nqo_GN) Decimal() string {
	return nqo.decimal
}

// Group returns the group of number
func (nqo *nqo_GN) Group() string {
	return nqo.group
}

// Group returns the minus sign of number
func (nqo *nqo_GN) Minus() string {
	return nqo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nqo_GN' and handles both Whole and Real numbers based on 'v'
func (nqo *nqo_GN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nqo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nqo.group) - 1; j >= 0; j-- {
					b = append(b, nqo.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nqo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nqo_GN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nqo *nqo_GN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nqo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nqo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nqo.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nqo_GN'
func (nqo *nqo_GN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nqo.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nqo_GN'
// in accounting notation.
func (nqo *nqo_GN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nqo.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nqo_GN'
func (nqo *nqo_GN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}
