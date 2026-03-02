package ccp_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ccp_IN struct {
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ccp_IN' locale
func New() locales.Translator {
	return &ccp_IN{
		locale:                 "ccp_IN",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "𑄎𑄚𑄪", "𑄜𑄬𑄛𑄴", "𑄟𑄢𑄴𑄌𑄧", "𑄃𑄬𑄛𑄳𑄢𑄨𑄣𑄴", "𑄟𑄬", "𑄎𑄪𑄚𑄴", "𑄎𑄪𑄣𑄭", "𑄃𑄉𑄧𑄌𑄴𑄑𑄴", "𑄥𑄬𑄛𑄴𑄑𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄃𑄧𑄇𑄴𑄑𑄮𑄝𑄧𑄢𑄴", "𑄚𑄧𑄞𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄓𑄨𑄥𑄬𑄟𑄴𑄝𑄢𑄴"},
		monthsNarrow:           []string{"", "𑄎", "𑄜𑄬", "𑄟", "𑄃𑄬", "𑄟𑄬", "𑄎𑄪𑄚𑄴", "𑄎𑄪", "𑄃", "𑄥𑄬", "𑄃𑄧", "𑄚𑄧", "𑄓𑄨"},
		monthsWide:             []string{"", "𑄎𑄚𑄪𑄠𑄢𑄨", "𑄜𑄬𑄛𑄴𑄝𑄳𑄢𑄪𑄠𑄢𑄨", "𑄟𑄢𑄴𑄌𑄧", "𑄃𑄬𑄛𑄳𑄢𑄨𑄣𑄴", "𑄟𑄬", "𑄎𑄪𑄚𑄴", "𑄎𑄪𑄣𑄭", "𑄃𑄉𑄧𑄌𑄴𑄑𑄴", "𑄥𑄬𑄛𑄴𑄑𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄃𑄧𑄇𑄴𑄑𑄬𑄝𑄧𑄢𑄴", "𑄚𑄧𑄞𑄬𑄟𑄴𑄝𑄧𑄢𑄴", "𑄓𑄨𑄥𑄬𑄟𑄴𑄝𑄧𑄢𑄴"},
		daysAbbreviated:        []string{"𑄢𑄧𑄝𑄨", "𑄥𑄧𑄟𑄴", "𑄟𑄧𑄁𑄉𑄧𑄣𑄴", "𑄝𑄪𑄖𑄴", "𑄝𑄳𑄢𑄨𑄥𑄪𑄛𑄴", "𑄥𑄪𑄇𑄴𑄇𑄮𑄢𑄴", "𑄥𑄧𑄚𑄨"},
		daysNarrow:             []string{"𑄢𑄧", "𑄥𑄧", "𑄟𑄧", "𑄝𑄪", "𑄝𑄳𑄢𑄨", "𑄥𑄪", "𑄥𑄧"},
		daysWide:               []string{"𑄢𑄧𑄝𑄨𑄝𑄢𑄴", "𑄥𑄧𑄟𑄴𑄝𑄢𑄴", "𑄟𑄧𑄁𑄉𑄧𑄣𑄴𑄝𑄢𑄴", "𑄝𑄪𑄖𑄴𑄝𑄢𑄴", "𑄝𑄳𑄢𑄨𑄥𑄪𑄛𑄴𑄝𑄢𑄴", "𑄥𑄪𑄇𑄴𑄇𑄮𑄢𑄴𑄝𑄢𑄴", "𑄥𑄧𑄚𑄨𑄝𑄢𑄴"},
		periodsAbbreviated:     []string{"", ""},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACST": "𑄃𑄬𑄉𑄧𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACT": "𑄃𑄬𑄉𑄧𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "𑄃𑄑𑄴𑄣𑄚𑄴𑄖𑄨𑄉𑄮𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ADT Arabia": "𑄃𑄢𑄧𑄝𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AEDT": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄪𑄉𑄬𑄘𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AEST": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄪𑄉𑄬𑄘𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AFT": "𑄃𑄛𑄴𑄉𑄚𑄨𑄌𑄴𑄖𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AKDT": "𑄃𑄣𑄌𑄴𑄇 𑄘𑄨𑄚𑄮𑄃𑄣𑄮 𑄃𑄧𑄇𑄴𑄖𑄧", "AKST": "𑄃𑄣𑄌𑄴𑄇 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AMST": "𑄃𑄳𑄠𑄟𑄎𑄧𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧 𑄃𑄧𑄇𑄴𑄖𑄧", "AMST Armenia": "𑄃𑄢𑄴𑄟𑄬𑄚𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AMT": "𑄃𑄳𑄠𑄟𑄎𑄧𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AMT Armenia": "𑄃𑄢𑄴𑄟𑄬𑄚𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ANAST": "𑄃𑄧𑄚𑄴𑄘𑄠𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ANAT": "𑄃𑄧𑄚𑄴𑄘𑄠𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ARST": "𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ART": "𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AST": "𑄃𑄑𑄴𑄣𑄚𑄴𑄖𑄨𑄉𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AST Arabia": "𑄃𑄢𑄧𑄝𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AWDT": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AWST": "𑄃𑄧𑄌𑄴𑄑𑄳𑄢𑄬𑄣𑄨𑄠𑄧 𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AZST": "𑄃𑄎𑄢𑄴𑄝𑄭𑄎𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "AZT": "𑄃𑄎𑄢𑄴𑄝𑄭𑄎𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BDT Bangladesh": "𑄝𑄁𑄣𑄘𑄬𑄌𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BNT": "𑄝𑄳𑄢𑄪𑄚𑄬𑄭 𑄘𑄢𑄪𑄌𑄴𑄥𑄣𑄟𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BOT": "𑄝𑄮𑄣𑄨𑄞𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "BRST": "𑄝𑄳𑄢𑄥𑄨𑄣𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BRT": "𑄝𑄳𑄢𑄥𑄨𑄣𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BST Bangladesh": "𑄝𑄁𑄣𑄘𑄬𑄌𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "BT": "𑄞𑄪𑄑𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CAST": "CAST", "CAT": "𑄟𑄧𑄖𑄴𑄙𑄳𑄠 𑄃𑄜𑄳𑄢𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "CCT": "𑄇𑄮𑄇𑄮𑄌𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CDT": "𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦 𑄘𑄨𑄚𑄮𑄃𑄣𑄮 𑄃𑄧𑄇𑄴𑄖𑄧", "CHADT": "𑄌𑄳𑄠𑄗𑄟𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CHAST": "𑄌𑄳𑄠𑄗𑄟𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CHUT": "𑄌𑄪𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CKT": "𑄇𑄪𑄇𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CKT DST": "𑄇𑄪𑄇𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄖𑄴𑄙𑄬𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CLST": "𑄌𑄨𑄣𑄨 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CLT": "𑄌𑄨𑄣𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "COST": "𑄇𑄧𑄣𑄧𑄟𑄴𑄝𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "COT": "𑄇𑄧𑄣𑄧𑄟𑄴𑄝𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CST": "𑄃𑄏𑄧𑄣𑄴 𑄉𑄧𑄢𑄳𑄦 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CST China": "𑄌𑄨𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CST China DST": "𑄌𑄨𑄚𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CVST": "𑄇𑄬𑄛𑄴 𑄞𑄢𑄴𑄘𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CVT": "𑄇𑄬𑄛𑄴 𑄞𑄢𑄴𑄘𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CXT": "𑄇𑄳𑄢𑄨𑄌𑄴𑄟𑄥𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄃𑄧𑄇𑄴𑄖𑄧", "ChST": "𑄌𑄟𑄬𑄢𑄮 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ChST NMI": "𑄅𑄪𑄖𑄴𑄖𑄮𑄢𑄴 𑄟𑄬𑄢𑄨𑄚𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CuDT": "𑄇𑄨𑄃𑄪𑄝 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "CuST": "𑄇𑄨𑄃𑄪𑄝 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "DAVT": "𑄓𑄬𑄞𑄨𑄥𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "DDUT": "𑄓𑄪𑄟𑄧𑄚𑄴𑄑𑄴-𑄘𑄳𑄠𑄧 𑄅𑄪𑄢𑄴𑄞𑄨𑄣𑄬 𑄃𑄧𑄇𑄴𑄖𑄧", "EASST": "𑄃𑄨𑄥𑄴𑄑𑄢𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EAST": "𑄃𑄨𑄥𑄴𑄑𑄢𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EAT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄜𑄳𑄢𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "ECT": "𑄃𑄨𑄇𑄪𑄠𑄬𑄓𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EDT": "𑄛𑄪𑄉𑄮 𑄞𑄨𑄘𑄬𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EGDT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EGST": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "EST": "𑄛𑄪𑄉𑄮 𑄞𑄨𑄘𑄬𑄢𑄴 𑄛𑄳𑄢𑄧𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FEET": "𑄃𑄬𑄇𑄴𑄉𑄮𑄙𑄳𑄆𑄬𑄣𑄴 𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FJT": "𑄜𑄨𑄎𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FJT Summer": "𑄜𑄨𑄎𑄨 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FKST": "𑄜𑄧𑄇𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FKT": "𑄜𑄧𑄇𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FNST": "𑄜𑄢𑄴𑄚𑄚𑄴𑄘𑄮𑄓𑄨 𑄚𑄮𑄢𑄮𑄚𑄴𑄦 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "FNT": "𑄜𑄢𑄴𑄚𑄚𑄴𑄘𑄮𑄓𑄨 𑄚𑄮𑄢𑄮𑄚𑄴𑄦 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GALT": "𑄉𑄣𑄛𑄉𑄮𑄌𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GAMT": "𑄉𑄳𑄠𑄟𑄴𑄝𑄨𑄠𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GEST": "𑄎𑄧𑄢𑄴𑄎𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GET": "𑄎𑄧𑄢𑄴𑄎𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GFT": "𑄜𑄧𑄢𑄥𑄨 𑄉𑄠𑄚 𑄃𑄧𑄇𑄴𑄖𑄧", "GIT": "𑄉𑄨𑄣𑄴𑄝𑄢𑄴𑄑𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GMT": "𑄉𑄳𑄢𑄨𑄚𑄨𑄌𑄴 𑄟𑄨𑄚𑄴 𑄑𑄬𑄟𑄴", "GNSST": "GNSST", "GNST": "GNST", "GST": "𑄘𑄧𑄉𑄨𑄚𑄴 𑄎𑄧𑄢𑄴𑄎𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "GST Guam": "𑄉𑄪𑄠𑄟𑄴 𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "GYT": "𑄉𑄪𑄠𑄚 𑄃𑄧𑄇𑄴𑄖𑄧", "HADT": "𑄦𑄧𑄃𑄮𑄠𑄭-𑄃𑄣𑄬𑄃𑄪𑄖𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HAST": "𑄦𑄧𑄃𑄮𑄠𑄭-𑄃𑄣𑄬𑄃𑄪𑄖𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HKST": "𑄦𑄧𑄁 𑄇𑄧𑄁 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HKT": "𑄦𑄧𑄁 𑄇𑄧𑄁 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HOVST": "𑄦𑄮𑄞𑄧𑄓𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "HOVT": "𑄦𑄮𑄞𑄧𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ICT": "𑄃𑄨𑄚𑄴𑄘𑄮𑄌𑄩𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IDT": "𑄃𑄨𑄎𑄴𑄢𑄠𑄬𑄣𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IOT": "𑄃𑄨𑄚𑄴𑄘𑄨𑄠𑄬 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IRKST": "𑄃𑄨𑄢𑄴𑄖𑄪𑄑𑄴𑄥𑄴𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IRKT": "𑄃𑄨𑄢𑄴𑄖𑄪𑄑𑄴𑄥𑄴𑄇𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IRST": "𑄃𑄨𑄢𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IRST DST": "𑄃𑄨𑄢𑄚𑄴 𑄘𑄨𑄚𑄬𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IST": "𑄃𑄨𑄚𑄴𑄘𑄨𑄠𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "IST Israel": "𑄃𑄨𑄎𑄴𑄢𑄠𑄬𑄣𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "JDT": "𑄎𑄛𑄚𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "JST": "𑄎𑄛𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "KOST": "𑄇𑄮𑄌𑄴𑄢𑄳𑄆𑄬 𑄃𑄧𑄇𑄴𑄖𑄧", "KRAST": "𑄇𑄳𑄢𑄥𑄴𑄚𑄮𑄠𑄢𑄴𑄥𑄴𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "KRAT": "𑄇𑄳𑄢𑄥𑄴𑄚𑄮𑄠𑄢𑄴𑄥𑄴𑄇𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "KST": "𑄇𑄮𑄢𑄨𑄠𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "KST DST": "𑄇𑄮𑄢𑄨𑄠𑄚𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "LHDT": "𑄣𑄧𑄢𑄴𑄓𑄴 𑄦𑄤𑄬 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "LHST": "𑄣𑄧𑄢𑄴𑄓𑄴 𑄦𑄤𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "LINT": "𑄣𑄭𑄚𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MAGST": "𑄟𑄳𑄠𑄉𑄓𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MAGT": "𑄟𑄳𑄠𑄉𑄓𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MART": "𑄟𑄢𑄴𑄇𑄬𑄥𑄥𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MAWT": "𑄟𑄧𑄥𑄳𑄦𑄧𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MDT": "𑄦𑄨𑄣𑄧𑄧𑄱 𑄞𑄨𑄘𑄬𑄢𑄴 𑄘𑄨𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MESZ": "𑄟𑄧𑄖𑄴𑄙𑄳𑄠 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MEZ": "𑄟𑄧𑄖𑄴𑄙𑄳𑄠 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MHT": "𑄟𑄢𑄴𑄥𑄣𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MMT": "𑄟𑄠𑄚𑄴𑄟𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MSD": "𑄟𑄧𑄥𑄴𑄇𑄮 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MST": "𑄦𑄨𑄣𑄧𑄧𑄱 𑄞𑄨𑄘𑄬𑄢𑄴 𑄛𑄳𑄢𑄧𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MUST": "𑄟𑄧𑄢𑄨𑄥𑄥𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MUT": "𑄟𑄧𑄢𑄨𑄥𑄥𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MVT": "𑄟𑄣𑄴𑄘𑄨𑄛𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "MYT": "𑄟𑄣𑄴𑄠𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "NCT": "𑄚𑄨𑄃𑄪 𑄇𑄳𑄠𑄣𑄬𑄓𑄮𑄚𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NDT": "𑄚𑄨𑄃𑄪𑄜𑄃𑄪𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄨 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NDT New Caledonia": "𑄚𑄨𑄃𑄪 𑄇𑄳𑄠𑄣𑄬𑄓𑄮𑄚𑄨𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NFDT": "𑄚𑄧𑄢𑄴𑄜𑄮𑄇𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NFT": "𑄚𑄧𑄢𑄴𑄜𑄮𑄇𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NOVST": "𑄚𑄮𑄞𑄮𑄥𑄨𑄝𑄨𑄢𑄴𑄥𑄴𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NOVT": "𑄚𑄮𑄞𑄮𑄥𑄨𑄝𑄨𑄢𑄴𑄥𑄴𑄇𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NPT": "𑄚𑄬𑄛𑄣𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NRT": "𑄚𑄃𑄪𑄢𑄪 𑄃𑄧𑄇𑄴𑄖𑄧", "NST": "𑄚𑄨𑄃𑄪𑄜𑄃𑄪𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NUT": "𑄚𑄨𑄃𑄪𑄃𑄨 𑄃𑄧𑄇𑄴𑄖𑄧", "NZDT": "𑄚𑄨𑄃𑄪𑄎𑄨𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "NZST": "𑄚𑄨𑄃𑄪𑄎𑄨𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "OESZ": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "OEZ": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "OMSST": "𑄃𑄮𑄟𑄧𑄥𑄴𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "OMST": "𑄃𑄮𑄟𑄧𑄥𑄴𑄇𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PDT": "𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄞𑄨𑄘𑄬𑄢𑄴 𑄘𑄨𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PDTM": "𑄟𑄬𑄇𑄴𑄥𑄨𑄇𑄚𑄴 𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PETDT": "𑄛𑄨𑄖𑄳𑄢𑄬𑄛𑄳𑄠𑄞𑄧𑄣𑄧𑄥𑄴𑄇𑄴-𑄇𑄳𑄠𑄟𑄴𑄌𑄳𑄠𑄑𑄴𑄃𑄨𑄥𑄴𑄇𑄨 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PETST": "𑄛𑄨𑄖𑄳𑄢𑄬𑄛𑄳𑄠𑄞𑄧𑄣𑄧𑄥𑄴𑄇𑄴-𑄇𑄳𑄠𑄟𑄴𑄌𑄳𑄠𑄑𑄴𑄃𑄨𑄥𑄴𑄇𑄨 𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PGT": "𑄛𑄛𑄪𑄠 𑄚𑄨𑄃𑄪 𑄉𑄨𑄚𑄨 𑄃𑄧𑄇𑄴𑄖𑄧", "PHOT": "𑄜𑄮𑄚𑄨𑄇𑄴𑄥𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PKT": "𑄛𑄇𑄨𑄥𑄴𑄖𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PKT DST": "𑄛𑄇𑄨𑄥𑄴𑄖𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PMDT": "𑄥𑄬𑄚𑄴𑄑𑄴 𑄛𑄨𑄠𑄬𑄢𑄴 𑄃𑄮 𑄟𑄨𑄇𑄬𑄣𑄧𑄚𑄴 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PMST": "𑄥𑄬𑄚𑄴𑄑𑄴 𑄛𑄨𑄠𑄬𑄢𑄴 𑄃𑄮 𑄟𑄨𑄇𑄬𑄣𑄧𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PONT": "𑄛𑄮𑄚𑄛𑄬 𑄃𑄧𑄇𑄴𑄖𑄧", "PST": "𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄞𑄨𑄘𑄬𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PST Philippine": "𑄜𑄨𑄣𑄨𑄛𑄭𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PST Philippine DST": "𑄜𑄨𑄣𑄨𑄛𑄭𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PST Pitcairn": "𑄛𑄨𑄖𑄴𑄇𑄬𑄠𑄢𑄴𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PSTM": "𑄟𑄬𑄇𑄴𑄥𑄨𑄇𑄚𑄴 𑄛𑄳𑄢𑄧𑄥𑄚𑄴𑄖𑄧 𑄟𑄧𑄦𑄥𑄉𑄧𑄢𑄧𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PWT": "𑄛𑄣𑄅𑄪 𑄃𑄧𑄇𑄴𑄖𑄧", "PYST": "𑄛𑄳𑄠𑄢𑄉𑄪𑄠𑄬 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PYT": "𑄛𑄳𑄠𑄢𑄉𑄪𑄠𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "PYT Korea": "𑄛𑄨𑄠𑄧𑄁 𑄃𑄨𑄠𑄁 𑄃𑄧𑄇𑄴𑄖𑄧", "RET": "𑄢𑄨𑄃𑄨𑄃𑄪𑄚𑄨𑄠𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ROTT": "𑄢𑄧𑄗𑄬𑄢 𑄃𑄧𑄇𑄴𑄖𑄧", "SAKST": "𑄥𑄈𑄣𑄨𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SAKT": "𑄥𑄈𑄣𑄨𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SAMST": "𑄥𑄟𑄢 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SAMT": "𑄥𑄟𑄢 𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SAST": "𑄘𑄧𑄉𑄨𑄚𑄴 𑄃𑄜𑄳𑄢𑄨𑄇 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SBT": "𑄥𑄧𑄣𑄮𑄟𑄧𑄚𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬𑄉𑄪𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SCT": "𑄥𑄬𑄥𑄬𑄣𑄧𑄥𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SGT": "𑄥𑄨𑄁𑄉𑄛𑄪𑄢 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SLST": "𑄣𑄧𑄁𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "SRT": "𑄥𑄪𑄢𑄨𑄚𑄟𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SST Samoa": "𑄥𑄟𑄮𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SST Samoa Apia": "𑄃𑄧𑄛𑄨𑄠 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SST Samoa Apia DST": "𑄃𑄧𑄛𑄨𑄠 𑄘𑄨𑄚𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SST Samoa DST": "𑄥𑄟𑄮𑄠 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "SYOT": "𑄥𑄠𑄮𑄤 𑄃𑄧𑄇𑄴𑄖𑄧", "TAAF": "𑄜𑄧𑄢𑄥𑄨 𑄘𑄧𑄉𑄨𑄚𑄴 𑄃𑄳𑄃 𑄃𑄚𑄴𑄑𑄢𑄴𑄇𑄧𑄑𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "TAHT": "𑄖𑄦𑄨𑄖𑄨 𑄃𑄧𑄇𑄴𑄖𑄧", "TJT": "𑄖𑄎𑄈𑄥𑄴𑄖𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TKT": "𑄑𑄮𑄇𑄬𑄣𑄃𑄪 𑄃𑄧𑄇𑄴𑄖𑄧", "TLT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄑𑄨𑄟𑄪𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TMST": "𑄖𑄪𑄢𑄴𑄇𑄴𑄟𑄬𑄚𑄨𑄌𑄴𑄖𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TMT": "𑄖𑄪𑄢𑄴𑄇𑄴𑄟𑄬𑄚𑄨𑄌𑄴𑄖𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TOST": "𑄑𑄮𑄋𑄴𑄉 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TOT": "𑄑𑄮𑄋𑄴𑄉 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TVT": "𑄑𑄪𑄞𑄣𑄪 𑄃𑄧𑄇𑄴𑄖𑄧", "TWT": "𑄖𑄭𑄛𑄳𑄆𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "TWT DST": "𑄖𑄭𑄛𑄳𑄆𑄬 𑄘𑄨𑄚𑄮𑄃𑄣𑄮𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ULAST": "𑄃𑄪𑄣𑄚𑄴 𑄝𑄖𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "ULAT": "𑄃𑄪𑄣𑄚𑄴 𑄝𑄖𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "UYST": "𑄃𑄪𑄢𑄪𑄉𑄪𑄠𑄬 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "UYT": "𑄃𑄪𑄢𑄪𑄉𑄪𑄠𑄬 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "UZT": "𑄃𑄪𑄎𑄴𑄝𑄬𑄇𑄨𑄖𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "UZT DST": "𑄃𑄪𑄎𑄴𑄝𑄬𑄇𑄨𑄖𑄚𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VET": "𑄞𑄬𑄚𑄬𑄎𑄪𑄠𑄬𑄣 𑄃𑄧𑄇𑄴𑄖𑄧", "VLAST": "𑄝𑄳𑄣𑄘𑄨𑄝𑄮𑄥𑄴𑄖𑄮𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VLAT": "𑄝𑄳𑄣𑄘𑄨𑄝𑄮𑄥𑄴𑄖𑄮𑄇𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VOLST": "𑄞𑄧𑄣𑄴𑄉𑄮𑄉𑄳𑄢𑄓𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VOLT": "𑄞𑄧𑄣𑄴𑄉𑄮𑄉𑄳𑄢𑄓𑄴 𑄟𑄧𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VOST": "𑄞𑄧𑄥𑄴𑄑𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VUT": "𑄞𑄚𑄪𑄠𑄖𑄪 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "VUT DST": "𑄞𑄚𑄪𑄠𑄖𑄪 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WAKT": "𑄃𑄮𑄠𑄬𑄇𑄴 𑄉𑄭 𑄉𑄭 𑄞𑄨𑄘𑄬 𑄃𑄧𑄇𑄴𑄖𑄧", "WARST": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WART": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄢𑄴𑄎𑄬𑄚𑄴𑄑𑄨𑄚 𑄛𑄳𑄢𑄧𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WAST": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄜𑄳𑄢𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "WAT": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄜𑄳𑄢𑄨𑄇 𑄃𑄧𑄇𑄴𑄖𑄧", "WESZ": "𑄛𑄧𑄏𑄬𑄟𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WEZ": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄨𑄃𑄪𑄢𑄮𑄝𑄮𑄢𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WFT": "𑄤𑄣𑄨𑄥𑄴 𑄃𑄳𑄃 𑄜𑄪𑄑𑄪𑄚 𑄃𑄧𑄇𑄴𑄖𑄧", "WGST": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WGT": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄉𑄳𑄢𑄨𑄚𑄴𑄣𑄳𑄠𑄚𑄴𑄓𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "WIB": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄃𑄨𑄚𑄴𑄘𑄮𑄚𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "WIT": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄃𑄨𑄚𑄴𑄘𑄮𑄚𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "WITA": "𑄃𑄏𑄧𑄣𑄴 𑄉𑄢𑄳𑄦 𑄃𑄨𑄚𑄴𑄘𑄮𑄚𑄬𑄥𑄨𑄠 𑄃𑄧𑄇𑄴𑄖𑄧", "YAKST": "𑄠𑄇𑄪𑄖𑄴𑄥𑄴𑄇𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "YAKT": "𑄠𑄇𑄪𑄖𑄴𑄥𑄴𑄇𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "YEKST": "𑄃𑄨𑄠𑄬𑄇𑄖𑄬𑄢𑄨𑄚𑄴𑄝𑄪𑄢𑄴𑄉𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "YEKT": "𑄃𑄨𑄠𑄬𑄇𑄖𑄬𑄢𑄨𑄚𑄴𑄝𑄪𑄢𑄴𑄉𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "YST": "YST", "МСК": "𑄟𑄧𑄥𑄴𑄇𑄮 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "اقتاۋ": "𑄃𑄇𑄴𑄑𑄃𑄮 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "اقتاۋ قالاسى": "𑄃𑄇𑄴𑄑𑄃𑄮 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "اقتوبە": "𑄃𑄇𑄴𑄑𑄮𑄝𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "اقتوبە قالاسى": "𑄃𑄇𑄴𑄑𑄮𑄝𑄴 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "الماتى": "𑄃𑄣𑄴𑄟𑄑𑄨 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "الماتى قالاسى": "𑄃𑄣𑄴𑄟𑄑𑄨 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧 𑄃𑄧𑄇𑄴𑄖𑄧", "باتىس قازاق ەلى": "𑄛𑄧𑄏𑄨𑄟𑄬𑄘𑄨 𑄇𑄎𑄇𑄴𑄥𑄳𑄖𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "شىعىش قازاق ەلى": "𑄛𑄪𑄉𑄬𑄘𑄨 𑄇𑄎𑄇𑄴𑄥𑄳𑄖𑄚𑄴 𑄟𑄚𑄧𑄇𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "𑄇𑄨𑄢𑄴𑄉𑄨𑄥𑄴𑄖𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "قىزىلوردا": "𑄇𑄨𑄎𑄨𑄣𑄮𑄢𑄴𑄓 𑄟𑄚𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "قىزىلوردا قالاسى": "𑄇𑄨𑄎𑄨𑄣𑄮𑄢𑄴𑄓 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧", "∅∅∅": "𑄛𑄬𑄢𑄪 𑄉𑄧𑄢𑄧𑄟𑄴𑄇𑄣𑄧𑄢𑄴 𑄃𑄧𑄇𑄴𑄖𑄧"},
	}
}

// Locale returns the current translators string locale
func (ccp *ccp_IN) Locale() string {
	return ccp.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ccp_IN'
func (ccp *ccp_IN) PluralsCardinal() []locales.PluralRule {
	return ccp.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ccp_IN'
func (ccp *ccp_IN) PluralsOrdinal() []locales.PluralRule {
	return ccp.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ccp_IN'
func (ccp *ccp_IN) PluralsRange() []locales.PluralRule {
	return ccp.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ccp_IN'
func (ccp *ccp_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ccp_IN'
func (ccp *ccp_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ccp_IN'
func (ccp *ccp_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ccp *ccp_IN) MonthAbbreviated(month time.Month) string {
	return ccp.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ccp *ccp_IN) MonthsAbbreviated() []string {
	return ccp.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ccp *ccp_IN) MonthNarrow(month time.Month) string {
	return ccp.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ccp *ccp_IN) MonthsNarrow() []string {
	return ccp.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ccp *ccp_IN) MonthWide(month time.Month) string {
	return ccp.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ccp *ccp_IN) MonthsWide() []string {
	return ccp.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ccp *ccp_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return ccp.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ccp *ccp_IN) WeekdaysAbbreviated() []string {
	return ccp.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ccp *ccp_IN) WeekdayNarrow(weekday time.Weekday) string {
	return ccp.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ccp *ccp_IN) WeekdaysNarrow() []string {
	return ccp.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ccp *ccp_IN) WeekdayShort(weekday time.Weekday) string {
	return ccp.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ccp *ccp_IN) WeekdaysShort() []string {
	return ccp.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ccp *ccp_IN) WeekdayWide(weekday time.Weekday) string {
	return ccp.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ccp *ccp_IN) WeekdaysWide() []string {
	return ccp.daysWide
}

// Decimal returns the decimal point of number
func (ccp *ccp_IN) Decimal() string {
	return ccp.decimal
}

// Group returns the group of number
func (ccp *ccp_IN) Group() string {
	return ccp.group
}

// Group returns the minus sign of number
func (ccp *ccp_IN) Minus() string {
	return ccp.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ccp_IN' and handles both Whole and Real numbers based on 'v'
func (ccp *ccp_IN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ccp.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, ccp.group[0])
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
		b = append(b, ccp.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ccp_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ccp *ccp_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ccp.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ccp.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ccp.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ccp_IN'
func (ccp *ccp_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ccp.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ccp.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, ccp.group[0])
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
		b = append(b, ccp.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ccp.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ccp_IN'
// in accounting notation.
func (ccp *ccp_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ccp.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ccp.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, ccp.group[0])
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
		b = append(b, ccp.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ccp.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ccp.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ccp.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ccp.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ccp.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ccp.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ccp.periodsAbbreviated[0]...)
	} else {
		b = append(b, ccp.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ccp.periodsAbbreviated[0]...)
	} else {
		b = append(b, ccp.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ccp.periodsAbbreviated[0]...)
	} else {
		b = append(b, ccp.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ccp_IN'
func (ccp *ccp_IN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ccp.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ccp.periodsAbbreviated[0]...)
	} else {
		b = append(b, ccp.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ccp.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
