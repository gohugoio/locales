package ff_Adlm

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ff_Adlm struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'ff_Adlm' locale
func New() locales.Translator {
	return &ff_Adlm{
		locale:                 "ff_Adlm",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "𞤀𞤊𞤀", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "FG", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "𞤐𞤐𞤘", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "𞤑𞤆𞤘", "𞤆𞤆𞤖", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "𞤊𞤅𞤊𞤀", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "𞤅𞤊𞤀", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "𞤅𞤭𞥅𞤤", "𞤕𞤮𞤤", "𞤐𞤦𞤮𞥅𞤴", "𞤅𞤫𞥅𞤼", "𞤁𞤵𞥅𞤶", "𞤑𞤮𞤪", "𞤃𞤮𞤪", "𞤔𞤵𞤳", "𞤅𞤭𞤤", "𞤒𞤢𞤪", "𞤔𞤮𞤤", "𞤄𞤮𞤱"},
		monthsNarrow:           []string{"", "𞤅", "𞤕", "𞤄", "𞤅", "𞤁", "𞤑", "𞤃", "𞤔", "𞤅", "𞤒", "𞤔", "𞤄"},
		monthsWide:             []string{"", "𞤅𞤭𞥅𞤤𞤮", "𞤕𞤮𞤤𞤼𞤮", "𞤐𞤦𞤮𞥅𞤴𞤮", "𞤅𞤫𞥅𞤼𞤮", "𞤁𞤵𞥅𞤶𞤮", "𞤑𞤮𞤪𞤧𞤮", "𞤃𞤮𞤪𞤧𞤮", "𞤔𞤵𞤳𞤮", "𞤅𞤭𞤤𞤼𞤮", "𞤒𞤢𞤪𞤳𞤮", "𞤔𞤮𞤤𞤮", "𞤄𞤮𞤱𞤼𞤮"},
		daysAbbreviated:        []string{"𞤈𞤫𞤬", "𞤀𞥄𞤩𞤵", "𞤃𞤢𞤦", "𞤔𞤫𞤧", "𞤐𞤢𞥄𞤧", "𞤃𞤢𞤣", "𞤖𞤮𞤪"},
		daysNarrow:             []string{"𞤈", "𞤀𞥄", "𞤃", "𞤔", "𞤐", "𞤃", "𞤖"},
		daysWide:               []string{"𞤈𞤫𞤬𞤦𞤭𞤪𞥆𞤫", "𞤀𞥄𞤩𞤵𞤲𞥋𞤣𞤫", "𞤃𞤢𞤱𞤦𞤢𞥄𞤪𞤫", "𞤐𞤶𞤫𞤧𞤤𞤢𞥄𞤪𞤫", "𞤐𞤢𞥄𞤧𞤢𞥄𞤲𞤣𞤫", "𞤃𞤢𞤱𞤲𞤣𞤫", "𞤖𞤮𞤪𞤦𞤭𞤪𞥆𞤫"},
		periodsAbbreviated:     []string{"𞤀𞤎", "𞤇𞤎"},
		periodsNarrow:          []string{"𞤢", "𞤩"},
		erasAbbreviated:        []string{"𞤀𞤀𞤋", "𞤇𞤀𞤋"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"𞤀𞤣𞤮 𞤀𞤲𞥆𞤢𞤦𞤭 𞤋𞥅𞤧𞤢𞥄", "𞤇𞤢𞥄𞤱𞤮 𞤀𞤲𞥆𞤢𞤦𞤭 𞤋𞥅𞤧𞤢𞥄"},
		timezones:              map[string]string{"ACDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ACST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞥄𞤳𞤭𞤪", "ACT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞥄𞤳𞤭𞤪", "ACWDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ACWST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ADT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤫𞤳𞤵", "ADT Arabia": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞥄𞤪𞤢𞤦𞤭𞤴𞤢", "AEDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AEST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AFT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤀𞤬𞤺𞤢𞤲𞤭𞤧𞤼𞤢𞥄𞤲", "AKDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤤𞤢𞤧𞤳𞤢𞥄 𞤲𞤣𞤫𞤲", "AKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤤𞤢𞤧𞤳𞤢𞥄 𞤲𞤣𞤫𞤲", "AMST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤥𞤢𞥁𞤮𞥅𞤲", "AMST Armenia": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤪𞤥𞤫𞤲𞤭𞤴𞤢𞥄", "AMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤥𞤢𞥁𞤮𞥅𞤲", "AMT Armenia": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤪𞤥𞤫𞤲𞤭𞤴𞤢𞥄", "ANAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤲𞤢𞤣𞤭𞥅𞤪", "ANAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤲𞤢𞤣𞤭𞥅𞤪", "ARST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "ART": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "AST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤫𞤳𞤵 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫", "AST Arabia": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞥄𞤪𞤢𞤦𞤭𞤴𞤢", "AWDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AWST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AZST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤶𞤫𞤪𞤦𞤢𞤴𞤶𞤢𞤲", "AZT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤶𞤫𞤪𞤦𞤢𞤴𞤶𞤢𞤲", "BDT Bangladesh": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤄𞤢𞤲𞤺𞤭𞤤𞤢𞤣𞤫𞥅𞤧", "BNT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤄𞤵𞤪𞤲𞤢𞤴", "BOT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤄𞤮𞤤𞤭𞤾𞤭𞤴𞤢𞥄", "BRST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤄𞤪𞤢𞤧𞤭𞤤𞤭𞤴𞤢𞥄", "BRT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤄𞤪𞤢𞤧𞤭𞤤𞤭𞤴𞤢𞥄", "BST Bangladesh": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤄𞤢𞤲𞤺𞤭𞤤𞤢𞤣𞤫𞥅𞤧", "BT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤄𞤵𞤼𞤢𞥄𞤲", "CAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤑𞤢𞥄𞤧𞤫𞤴", "CAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "CCT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤑𞤮𞥅𞤳𞤮𞤧", "CDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "CHADT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤕𞤢𞥄𞤼𞤢𞤥", "CHAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤢𞤼𞤢𞥄𞤥", "CHUT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤵𞥅𞤳𞤵", "CKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤵𞤪𞤭𞥅𞤶𞤫 𞤑𞤵𞥅𞤳", "CKT DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤆𞤫𞤷𞥆𞤵 𞤕𞤫𞥅𞤯𞤵 𞤕𞤵𞤪𞤭𞥅𞤶𞤫 𞤑𞤵𞥅𞤳", "CLST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤕𞤭𞤤𞤫𞥅", "CLT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤭𞤤𞤫𞥅", "COST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤮𞤤𞤮𞤥𞤦𞤭𞤴𞤢𞥄", "COT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤮𞤤𞤮𞤥𞤦𞤭𞤴𞤢𞥄", "CST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "CST China": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤢𞤴𞤲𞤢", "CST China DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤕𞤢𞤴𞤲𞤢", "CVST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤢𞥄𞤦𞤮-𞤜𞤫𞤪𞤣𞤫", "CVT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤢𞥄𞤦𞤮 𞤜𞤫𞤪𞤣𞤫", "CXT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤵𞤪𞤭𞥅𞤪𞤫 𞤑𞤭𞤪𞤧𞤭𞤥𞤢𞥄𞤧", "ChST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤮𞤥𞤮𞥅𞤪𞤮", "ChST NMI": "ChST NMI", "CuDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤵𞥅𞤦𞤢𞥄", "CuST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤵𞥅𞤦𞤢𞥄", "DAVT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤁𞤫𞥅𞤾𞤭𞤧", "DDUT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤁𞤭𞤥𞤮𞤲𞤼𞤵-𞤁𞤵𞤪𞤾𞤭𞤤", "EASST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤋𞤧𞤼𞤮𞥅-𞤀𞤴𞤤𞤢𞤲𞤣", "EAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤋𞤧𞤼𞤮𞥅-𞤀𞤴𞤤𞤢𞤲𞤣", "EAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "ECT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤋𞤳𞤵𞤱𞤢𞤣𞤮𞥅𞤪", "EDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "EGDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "EGST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤬𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "EST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "FEET": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫𞥅𞤪𞤭 𞤀𞤪𞤮𞥅𞤦𞤢", "FJT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤭𞤶𞤭𞥅", "FJT Summer": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤭𞤶𞤭𞥅", "FKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤮𞤤𞤳𞤤𞤢𞤲𞤣-𞤀𞤴𞤤𞤢𞤲𞤣", "FKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤮𞤤𞤳𞤤𞤢𞤲𞤣-𞤀𞤴𞤤𞤢𞤲𞤣", "FNST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤫𞤪𞤲𞤢𞤲𞤣𞤮𞥅 𞤣𞤫 𞤐𞤮𞤪𞤮𞤲𞤽𞤢𞥄", "FNT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤫𞤪𞤲𞤢𞤲𞤣𞤮𞥅 𞤣𞤫 𞤐𞤮𞤪𞤮𞤲𞤽𞤢𞥄", "GALT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤘𞤢𞤤𞤢𞤨𞤢𞤺𞤮𞥅𞤧", "GAMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤘𞤢𞤥𞤦𞤭𞤴𞤫𞤪", "GEST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤔𞤮𞤪𞤶𞤭𞤴𞤢", "GET": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤔𞤮𞤪𞤶𞤭𞤴𞤢", "GFT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤘𞤢𞤴𞤢𞤲𞤢𞥄-𞤊𞤪𞤢𞤲𞤧𞤭", "GIT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤵𞤪𞤭𞥅𞤶𞤫 𞤘𞤭𞤤𞤦𞤫𞤪𞤼𞤵", "GMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤳𞤭𞤲𞤭𞥅𞤲𞥋𞤣𞤫 𞤘𞤪𞤭𞤲𞤱𞤭𞥅𞤧", "GNSST": "GNSST", "GNST": "GNST", "GST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤢𞤱𞤬-𞤔𞤮𞤪𞤶𞤭𞤴𞤢𞥄", "GST Guam": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤘𞤵𞤱𞤢𞥄𞤥", "GYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤘𞤢𞤴𞤢𞤲𞤢𞥄", "HADT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤢𞤱𞤢𞥄𞤴𞤭𞥅-𞤀𞤤𞤮𞤧𞤭𞤴𞤢𞤲", "HAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤢𞤱𞤢𞥄𞤴𞤭𞥅-𞤀𞤤𞤮𞤧𞤭𞤴𞤢𞤲", "HKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤮𞤲𞤺 𞤑𞤮𞤲𞤺", "HKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤮𞤲𞤺 𞤑𞤮𞤲𞤺", "HOVST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤮𞤬𞤣𞤵", "HOVT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤮𞤬𞤣𞤵", "ICT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤋𞤲𞤣𞤮𞤧𞤭𞥅𞤲", "IDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤋𞤧𞤪𞤢𞥄𞤭𞥅𞤤𞤵", "IOT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤢𞥄𞤴𞤮 𞤋𞤲𞤣𞤭𞤴𞤢𞤱𞤮", "IRKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤋𞤪𞤳𞤵𞤼𞤭𞤧𞤳𞤵", "IRKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤋𞤪𞤳𞤵𞤼𞤭𞤧𞤳𞤵", "IRST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤋𞤪𞤢𞥄𞤲", "IRST DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤋𞤪𞤢𞥄𞤲", "IST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤋𞤲𞤣𞤭𞤴𞤢", "IST Israel": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤋𞤧𞤪𞤢𞥄𞤭𞥅𞤤𞤵", "JDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤭𞤨𞥆𞤮𞤲", "JST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤭𞤨𞥆𞤮𞤲", "KOST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤑𞤮𞤧𞤪𞤢𞤴", "KRAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤢𞤪𞤢𞤧𞤲𞤮𞤴𞤢𞤪𞤧𞤭𞤳", "KRAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤪𞤢𞤧𞤲𞤮𞤴𞤢𞤪𞤧𞤭𞤳", "KST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤮𞥅𞤪𞤫𞤴𞤢𞥄", "KST DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤮𞥅𞤪𞤫𞤴𞤢𞥄", "LHDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤂𞤮𞤪𞤣𞤵-𞤖𞤮𞤱𞤫", "LHST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤂𞤮𞤪𞤣𞤵-𞤖𞤮𞤱𞤫", "LINT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤵𞤪𞤭𞥅𞤶𞤫 𞤂𞤢𞤴𞤲𞤵", "MAGST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤃𞤢𞤺𞤢𞤣𞤢𞤲", "MAGT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤃𞤢𞤺𞤢𞤣𞤢𞤲", "MART": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤢𞤪𞤳𞤫𞤧𞤢𞤧", "MAWT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤢𞤱𞤧𞤮𞤲", "MDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤃𞤢𞤳𞤢𞤱𞤮𞥅", "MESZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤀𞤪𞤮𞥅𞤦𞤢", "MEZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤀𞤪𞤮𞥅𞤦𞤢", "MHT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤵𞤪𞤭𞥅𞤶𞤫 𞤃𞤢𞤪𞤧𞤢𞤤", "MMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤭𞤴𞤢𞤥𞤢𞥄𞤪", "MSD": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤃𞤮𞤧𞤳𞤮", "MST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤃𞤢𞤳𞤢𞤱𞤮𞥅", "MUST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤃𞤮𞤪𞤭𞥅𞤧𞤭", "MUT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤃𞤮𞤪𞤭𞥅𞤧𞤭", "MVT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤢𞤤𞤣𞤢𞥄𞤴𞤭𞤧", "MYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤢𞤤𞤫𞥅𞤧𞤭𞤴𞤢", "NCT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤢𞤤𞤭𞤣𞤮𞤲𞤭𞤴𞤢𞥄 𞤖𞤫𞤧𞤮", "NDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤫𞤱-𞤊𞤵𞤲𞤣𞤵𞤤𞤢𞤲𞤣", "NDT New Caledonia": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤢𞤤𞤭𞤣𞤮𞤲𞤭𞤴𞤢𞥄 𞤖𞤫𞤧𞤮", "NFDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤅𞤵𞤪𞤭𞥅𞤪𞤫 𞤐𞤮𞤪𞤬𞤮𞤤𞤳𞤵", "NFT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤵𞤪𞤭𞥅𞤪𞤫 𞤐𞤮𞤪𞤬𞤮𞤤𞤳𞤵", "NOVST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤮𞤾𞤮𞤧𞤦𞤭𞤪𞤧𞤭𞤳", "NOVT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤮𞤾𞤮𞤧𞤦𞤭𞤪𞤧𞤭𞤳", "NPT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤐𞤫𞤨𞤢𞤤", "NRT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤐𞤵𞥅𞤪𞤵", "NST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤫𞤱-𞤊𞤵𞤲𞤣𞤵𞤤𞤢𞤲𞤣", "NUT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤐𞤵𞥅𞤱𞤭", "NZDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤫𞤱-𞤟𞤫𞤤𞤢𞤲𞤣𞤭", "NZST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤫𞤱-𞤟𞤫𞤤𞤢𞤲𞤣𞤭", "OESZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "OEZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "OMSST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤌𞤥𞤧𞤵𞤳𞤵", "OMST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤌𞤥𞤧𞤵𞤳𞤵", "PDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤁𞤫𞤰𞥆𞤮 𞤕𞤫𞥅𞤯𞤵 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "PDTM": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤁𞤫𞤰𞥆𞤮 𞤃𞤫𞤳𞤧𞤭𞤳𞤮𞥅", "PETDT": "PETDT", "PETST": "PETST", "PGT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤆𞤢𞤨𞤵𞤱𞤢 𞤘𞤭𞤲𞤫 𞤖𞤫𞤧𞤮", "PHOT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤵𞤪𞤭𞥅𞤶𞤫 𞤊𞤫𞤲𞤭𞤳𞤧𞤭", "PKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤆𞤢𞤳𞤭𞤧𞤼𞤢𞥄𞤲", "PKT DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤆𞤢𞤳𞤭𞤧𞤼𞤢𞥄𞤲", "PMDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤅𞤼. 𞤆𞤭𞤴𞤫𞥅𞤪 & 𞤃𞤭𞤳𞤫𞤤𞤮𞤲", "PMST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤼. 𞤆𞤭𞤴𞤫𞥅𞤪 & 𞤃𞤭𞤳𞤫𞤤𞤮𞤲", "PONT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤆𞤮𞤲𞤢𞥄𞤨𞤫", "PST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤁𞤫𞤰𞥆𞤮 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "PST Philippine": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤭𞤤𞤭𞤨𞤭𞥅𞤲", "PST Philippine DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤭𞤤𞤭𞤨𞤭𞥅𞤲", "PST Pitcairn": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤆𞤭𞤼𞤳𞤭𞥅𞤪𞤲𞤵", "PSTM": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤁𞤫𞤰𞥆𞤮 𞤃𞤫𞤳𞤧𞤭𞤳𞤮𞥅", "PWT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤆𞤮𞤤𞤢𞥄𞤱𞤮", "PYST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤆𞤢𞥄𞤪𞤢𞤺𞤮𞤴", "PYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤆𞤢𞥄𞤪𞤢𞤺𞤮𞤴", "PYT Korea": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤆𞤭𞤴𞤮𞤲𞤴𞤢𞤲", "RET": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤈𞤫𞤲𞤭𞤴𞤮𞤲", "ROTT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤈𞤮𞤼𞤫𞤪𞤢", "SAKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤅𞤢𞤿𞤢𞤤𞤭𞥅𞤲", "SAKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤢𞤿𞤢𞤤𞤭𞥅𞤲", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "SBT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤵𞤪𞤭𞥅𞤶𞤫 𞤅𞤵𞤤𞤢𞤴𞤥𞤢𞥄𞤲", "SCT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤫𞤴𞤭𞤧𞤫𞤤", "SGT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤭𞤲𞤺𞤢𞤨𞤵𞥅𞤪", "SLST": "SLST", "SRT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤭𞤪𞤭𞤲𞤢𞤥", "SST Samoa": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤢𞤥𞤵𞤱𞤢", "SST Samoa Apia": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞥄𞤨𞤭𞤴𞤢", "SST Samoa Apia DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞥄𞤨𞤭𞤴𞤢", "SST Samoa DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤅𞤢𞤥𞤵𞤱𞤢", "SYOT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤢𞥄𞤴𞤵𞤱𞤢", "TAAF": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤂𞤫𞤴𞤪𞤭 𞤊𞤪𞤢𞤲𞤧𞤭 & 𞤀𞤪𞤼𞤢𞤲𞤼𞤭𞤳𞤢", "TAHT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤢𞤸𞤭𞤼𞤭𞥅", "TJT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤢𞤶𞤭𞤳𞤭𞤧𞤼𞤢𞥄𞤲", "TKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤮𞥅𞤳𞤮𞤤𞤢𞥄𞤱𞤵", "TLT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤚𞤭𞥅𞤥𞤮𞤪", "TMST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤪𞤳𞤵𞤥𞤫𞤲𞤭𞤧𞤼𞤢𞥄𞤲", "TMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤪𞤳𞤵𞤥𞤫𞤲𞤭𞤧𞤼𞤢𞥄𞤲", "TOST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤮𞤲𞤺𞤢", "TOT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤮𞤲𞤺𞤢 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫", "TVT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤵𞥅𞤾𞤢𞤤𞤵", "TWT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤢𞤴𞤨𞤫𞥅", "TWT DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤢𞤴𞤨𞤫𞥅", "ULAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤓𞤤𞤢𞤲𞤦𞤢𞤼𞤢𞤪", "ULAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤓𞤤𞤢𞤲𞤦𞤢𞤼𞤢𞤪", "UYST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤒𞤵𞥅𞤪𞤺𞤮𞤴", "UYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤒𞤵𞥅𞤪𞤺𞤮𞤴", "UZT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤓𞥁𞤦𞤫𞤳𞤭𞤧𞤼𞤢𞥄𞤲", "UZT DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤓𞥁𞤦𞤫𞤳𞤭𞤧𞤼𞤢𞥄𞤲", "VET": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤜𞤫𞤲𞤭𞥅𞥁𞤮𞥅𞤤𞤢", "VLAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤜𞤭𞤤𞤢𞤾𞤮𞤧𞤼𞤮𞤳", "VLAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤜𞤭𞤤𞤢𞤾𞤮𞤧𞤼𞤮𞤳", "VOLST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤜𞤮𞤤𞤺𞤮𞤺𞤢𞤪𞤢𞥄𞤣", "VOLT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤜𞤮𞤤𞤺𞤮𞤺𞤢𞤪𞤢𞥄𞤣", "VOST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤜𞤮𞤧𞤼𞤮𞤳", "VUT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤜𞤢𞤲𞤵𞤱𞤢𞥄𞤼𞤵", "VUT DST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤜𞤢𞤲𞤵𞤱𞤢𞥄𞤼𞤵", "WAKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤵𞤪𞤭𞥅𞤪𞤫 𞤏𞤫𞥅𞤳𞤵", "WARST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "WART": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "WAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "WAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "WESZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "WEZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "WFT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤏𞤢𞤤𞥆𞤭𞥅𞤧 & 𞤊𞤵𞤼𞤵𞤲𞤢", "WGST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "WGT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "WIB": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤋𞤲𞤣𞤮𞤲𞤭𞥅𞤧𞤭𞤴𞤢", "WIT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤋𞤲𞤣𞤮𞤲𞤭𞥅𞤧𞤭𞤴𞤢", "WITA": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤮𞤥𞤦𞤮𞥅𞤪𞤭 𞤋𞤲𞤣𞤮𞤲𞤭𞥅𞤧𞤭𞤴𞤢", "YAKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤒𞤢𞤳𞤢𞤼𞤭𞤧𞤳𞤵", "YAKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤒𞤢𞤳𞤢𞤼𞤭𞤧𞤳𞤵", "YEKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤒𞤫𞤳𞤢𞤼𞤫𞤪𞤭𞤲𞤦𞤵𞤪𞤺𞤵", "YEKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤒𞤫𞤳𞤢𞤼𞤫𞤪𞤭𞤲𞤦𞤵𞤪𞤺𞤵", "YST": "YST", "МСК": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤃𞤮𞤧𞤳𞤮", "اقتاۋ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤳𞤼𞤢𞤱", "اقتاۋ قالاسى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤳𞤼𞤢𞤱", "اقتوبە": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤳𞤼𞤮𞤦𞤭𞥅", "اقتوبە قالاسى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤳𞤼𞤮𞤦𞤭𞥅", "الماتى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤤𞤥𞤢𞤼𞤭", "الماتى قالاسى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤤𞤥𞤢𞤼𞤭", "باتىس قازاق ەلى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤭𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤑𞤢𞥁𞤢𞤿𞤢𞤧𞤼𞤢𞥄𞤲", "شىعىش قازاق ەلى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤑𞤢𞥁𞤢𞤿𞤢𞤧𞤼𞤢𞥄𞤲", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤑𞤭𞤪𞤺𞤭𞤧𞤼𞤢𞥄𞤲", "قىزىلوردا": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤭𞤶𞤤𞤮𞤪𞤣𞤢𞥄", "قىزىلوردا قالاسى": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤭𞤶𞤤𞤮𞤪𞤣𞤢𞥄", "∅∅∅": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞥁𞤮𞤪𞤫𞤧"},
	}
}

// Locale returns the current translators string locale
func (ff *ff_Adlm) Locale() string {
	return ff.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ff_Adlm'
func (ff *ff_Adlm) PluralsCardinal() []locales.PluralRule {
	return ff.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ff_Adlm'
func (ff *ff_Adlm) PluralsOrdinal() []locales.PluralRule {
	return ff.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ff_Adlm'
func (ff *ff_Adlm) PluralsRange() []locales.PluralRule {
	return ff.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ff_Adlm'
func (ff *ff_Adlm) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ff_Adlm'
func (ff *ff_Adlm) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ff_Adlm'
func (ff *ff_Adlm) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ff *ff_Adlm) MonthAbbreviated(month time.Month) string {
	return ff.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ff *ff_Adlm) MonthsAbbreviated() []string {
	return ff.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ff *ff_Adlm) MonthNarrow(month time.Month) string {
	return ff.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ff *ff_Adlm) MonthsNarrow() []string {
	return ff.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ff *ff_Adlm) MonthWide(month time.Month) string {
	return ff.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ff *ff_Adlm) MonthsWide() []string {
	return ff.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayAbbreviated(weekday time.Weekday) string {
	return ff.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ff *ff_Adlm) WeekdaysAbbreviated() []string {
	return ff.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayNarrow(weekday time.Weekday) string {
	return ff.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ff *ff_Adlm) WeekdaysNarrow() []string {
	return ff.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayShort(weekday time.Weekday) string {
	return ff.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ff *ff_Adlm) WeekdaysShort() []string {
	return ff.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayWide(weekday time.Weekday) string {
	return ff.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ff *ff_Adlm) WeekdaysWide() []string {
	return ff.daysWide
}

// Decimal returns the decimal point of number
func (ff *ff_Adlm) Decimal() string {
	return ff.decimal
}

// Group returns the group of number
func (ff *ff_Adlm) Group() string {
	return ff.group
}

// Group returns the minus sign of number
func (ff *ff_Adlm) Minus() string {
	return ff.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ff_Adlm' and handles both Whole and Real numbers based on 'v'
func (ff *ff_Adlm) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ff_Adlm' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ff *ff_Adlm) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ff_Adlm'
func (ff *ff_Adlm) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ff.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ff.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ff.group) - 1; j >= 0; j-- {
					b = append(b, ff.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ff.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ff.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ff.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ff_Adlm'
// in accounting notation.
func (ff *ff_Adlm) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ff.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ff.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ff.group) - 1; j >= 0; j-- {
					b = append(b, ff.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ff.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ff.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ff.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ff.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ff.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xe2, 0xb9, 0x81, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ff.monthsWide[t.Month()]...)
	b = append(b, []byte{0xe2, 0xb9, 0x81, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ff.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ff.monthsWide[t.Month()]...)
	b = append(b, []byte{0xe2, 0xb9, 0x81, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ff.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
