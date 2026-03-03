package sat

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sat struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'sat' locale
func New() locales.Translator {
	return &sat{
		locale:                 "sat",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ᱜᱮᱞᱥᱟᱭ",
		currencyNegativeSuffix: " ᱜᱮᱞᱥᱟᱭ",
		monthsAbbreviated:      []string{"", "ᱡᱟᱱ", "ᱯᱷᱟ", "ᱢᱟᱨ", "ᱟᱯᱨ", "ᱢᱮ", "ᱡᱩᱱ", "ᱡᱩᱞ", "ᱟᱜᱟ", "ᱥᱮᱯ", "ᱚᱠᱴ", "ᱱᱟᱣ", "ᱫᱤᱥ"},
		monthsNarrow:           []string{"", "ᱡ", "ᱯ", "ᱢ", "ᱟ", "ᱢ", "ᱡ", "ᱡ", "ᱟ", "ᱥ", "ᱚ", "ᱱ", "ᱫ"},
		monthsWide:             []string{"", "ᱡᱟᱱᱣᱟᱨᱤ", "ᱯᱷᱟᱨᱣᱟᱨᱤ", "ᱢᱟᱨᱪ", "ᱟᱯᱨᱮᱞ", "ᱢᱮ", "ᱡᱩᱱ", "ᱡᱩᱞᱟᱭ", "ᱟᱜᱟᱥᱛ", "ᱥᱮᱯᱴᱮᱢᱵᱟᱨ", "ᱚᱠᱴᱚᱵᱟᱨ", "ᱱᱟᱣᱟᱢᱵᱟᱨ", "ᱫᱤᱥᱟᱢᱵᱟᱨ"},
		daysAbbreviated:        []string{"ᱥᱤᱸ", "ᱚᱛ", "ᱵᱟ", "ᱥᱟᱹ", "ᱥᱟᱹᱨ", "ᱡᱟᱹ", "ᱧᱩ"},
		daysNarrow:             []string{"ᱥ", "ᱚ", "ᱵ", "ᱥ", "ᱥ", "ᱡ", "ᱧ"},
		daysWide:               []string{"ᱥᱤᱸᱜᱮ", "ᱚᱛᱮ", "ᱵᱟᱞᱮ", "ᱥᱟᱹᱜᱩᱱ", "ᱥᱟᱹᱨᱫᱤ", "ᱡᱟᱹᱨᱩᱢ", "ᱧᱩᱦᱩᱢ"},
		timezones:              map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "ᱮᱴᱞᱟᱱᱴᱤᱠ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "ADT Arabia": "ADT Arabia", "AEDT": "AEDT", "AEST": "AEST", "AFT": "AFT", "AKDT": "ᱟᱹᱞᱟᱥᱠᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "AKST": "ᱟᱹᱞᱟᱥᱠᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "AMST": "ᱟᱢᱟᱡᱚᱱ ᱥᱤᱛᱩᱝ ᱚᱠᱴᱚ", "AMST Armenia": "AMST Armenia", "AMT": "ᱟᱢᱟᱡᱚᱱ ᱮᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱴᱚ", "AMT Armenia": "AMT Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱥᱤᱛᱩᱝ ᱚᱠᱴᱚ", "ART": "ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱮᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱴᱚ", "AST": "ᱮᱴᱞᱟᱱᱴᱤᱠ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "AST Arabia": "AST Arabia", "AWDT": "AWDT", "AWST": "AWST", "AZST": "AZST", "AZT": "AZT", "BDT Bangladesh": "BDT Bangladesh", "BNT": "BNT", "BOT": "BOT", "BRST": "BRST", "BRT": "BRT", "BST Bangladesh": "BST Bangladesh", "BT": "BT", "CAST": "CAST", "CAT": "CAT", "CCT": "CCT", "CDT": "ᱛᱟᱱᱟᱞᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "ᱛᱟᱱᱟᱞᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "CST China": "CST China", "CST China DST": "CST China DST", "CVST": "CVST", "CVT": "CVT", "CXT": "CXT", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "ᱠᱭᱩᱵᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "CuST": "ᱠᱭᱩᱵᱟ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "EASST", "EAST": "EAST", "EAT": "EAT", "ECT": "ECT", "EDT": "ᱤᱥᱴᱟᱨᱱ ᱥᱤᱧᱟᱜ ᱵᱚᱠᱛᱚ", "EGDT": "ᱥᱟᱢᱟᱝ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱥᱤᱛᱩᱝ ᱚᱠᱛᱚ", "EGST": "ᱥᱟᱢᱟᱝ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "EST": "ᱤᱥᱴᱟᱨᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "FEET": "FEET", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "FKST", "FKT": "FKT", "FNST": "FNST", "FNT": "FNT", "GALT": "GALT", "GAMT": "GAMT", "GEST": "GEST", "GET": "GET", "GFT": "GFT", "GIT": "GIT", "GMT": "ᱜᱨᱤᱱᱣᱤᱪ ᱢᱤᱱ ᱚᱠᱛᱚ", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "GYT", "HADT": "ᱦᱟᱣᱟᱭᱤᱼᱟᱞᱮᱣᱴᱤᱭᱟᱱ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "HAST": "ᱦᱟᱣᱟᱭᱤᱼᱟᱞᱮᱣᱴᱤᱭᱟᱱ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "HKST": "HKST", "HKT": "HKT", "HOVST": "HOVST", "HOVT": "HOVT", "ICT": "ICT", "IDT": "IDT", "IOT": "IOT", "IRKST": "IRKST", "IRKT": "IRKT", "IRST": "IRST", "IRST DST": "IRST DST", "IST": "IST", "IST Israel": "IST Israel", "JDT": "JDT", "JST": "JST", "KOST": "KOST", "KRAST": "KRAST", "KRAT": "KRAT", "KST": "KST", "KST DST": "KST DST", "LHDT": "LHDT", "LHST": "LHST", "LINT": "LINT", "MAGST": "MAGST", "MAGT": "MAGT", "MART": "MART", "MAWT": "MAWT", "MDT": "ᱢᱟᱩᱱᱴᱮᱱ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "MESZ": "ᱥᱮᱱᱴᱨᱟᱞ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "MEZ": "ᱥᱮᱱᱴᱨᱟᱞ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "MHT": "MHT", "MMT": "MMT", "MSD": "MSD", "MST": "ᱢᱟᱩᱱᱴᱮᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "MUST": "MUST", "MUT": "MUT", "MVT": "MVT", "MYT": "MYT", "NCT": "NCT", "NDT": "ᱱᱤᱭᱩᱯᱷᱟᱩᱱᱰᱞᱮᱸᱰ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "NDT New Caledonia": "NDT New Caledonia", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "NOVST", "NOVT": "NOVT", "NPT": "NPT", "NRT": "NRT", "NST": "ᱱᱤᱭᱩᱯᱷᱟᱩᱱᱰᱞᱮᱸᱰ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "NUT": "NUT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "ᱤᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "OEZ": "ᱤᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "OMSST": "OMSST", "OMST": "OMST", "PDT": "ᱯᱮᱥᱤᱯᱷᱤᱠ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PDTM": "ᱢᱮᱠᱥᱤᱠᱟᱱ ᱯᱨᱚᱥᱟᱱᱛ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "PKT", "PKT DST": "PKT DST", "PMDT": "ᱮᱥ ᱴᱤ ᱯᱤᱭᱮᱨᱮ ᱟᱨ ᱢᱤᱠᱮᱞᱚᱱ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PMST": "ᱮᱥ ᱴᱤ ᱯᱤᱭᱮᱨᱮ ᱟᱨ ᱢᱤᱠᱮᱞᱚᱱ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "PONT": "PONT", "PST": "ᱯᱮᱥᱤᱯᱷᱤᱠ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "PST Philippine": "PST Philippine", "PST Philippine DST": "PST Philippine DST", "PST Pitcairn": "PST Pitcairn", "PSTM": "ᱢᱮᱠᱥᱤᱠᱟᱱ ᱯᱨᱚᱥᱟᱱᱛ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "PWT": "PWT", "PYST": "PYST", "PYT": "PYT", "PYT Korea": "PYT Korea", "RET": "RET", "ROTT": "ROTT", "SAKST": "SAKST", "SAKT": "SAKT", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "SAST", "SBT": "SBT", "SCT": "SCT", "SGT": "SGT", "SLST": "SLST", "SRT": "SRT", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "TAAF", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "TLT", "TMST": "TMST", "TMT": "TMT", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "TWT", "TWT DST": "TWT DST", "ULAST": "ULAST", "ULAT": "ULAT", "UYST": "UYST", "UYT": "UYT", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "VET", "VLAST": "VLAST", "VLAT": "VLAT", "VOLST": "VOLST", "VOLT": "VOLT", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "ᱯᱟᱪᱮ ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱥᱤᱛᱩᱝ ᱚᱠᱴᱚ", "WART": "ᱯᱟᱪᱮ ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱴᱚ", "WAST": "WAST", "WAT": "WAT", "WESZ": "ᱣᱮᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "WEZ": "ᱣᱮᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "WFT": "WFT", "WGST": "ᱯᱟᱪᱮ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱥᱤᱛᱩᱝ ᱚᱠᱛᱚ", "WGT": "ᱯᱟᱪᱮ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "YAKST": "YAKST", "YAKT": "YAKT", "YEKST": "YEKST", "YEKT": "YEKT", "YST": "ᱭᱩᱠᱚᱱ ᱚᱠᱛᱚ", "МСК": "МСК", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "شىعىش قازاق ەلى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (sat *sat) Locale() string {
	return sat.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sat'
func (sat *sat) PluralsCardinal() []locales.PluralRule {
	return sat.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sat'
func (sat *sat) PluralsOrdinal() []locales.PluralRule {
	return sat.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sat'
func (sat *sat) PluralsRange() []locales.PluralRule {
	return sat.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sat'
func (sat *sat) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sat'
func (sat *sat) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sat'
func (sat *sat) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sat *sat) MonthAbbreviated(month time.Month) string {
	if len(sat.monthsAbbreviated) == 0 {
		return ""
	}
	return sat.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sat *sat) MonthsAbbreviated() []string {
	return sat.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sat *sat) MonthNarrow(month time.Month) string {
	if len(sat.monthsNarrow) == 0 {
		return ""
	}
	return sat.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sat *sat) MonthsNarrow() []string {
	return sat.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sat *sat) MonthWide(month time.Month) string {
	if len(sat.monthsWide) == 0 {
		return ""
	}
	return sat.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sat *sat) MonthsWide() []string {
	return sat.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sat *sat) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sat.daysAbbreviated) == 0 {
		return ""
	}
	return sat.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sat *sat) WeekdaysAbbreviated() []string {
	return sat.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sat *sat) WeekdayNarrow(weekday time.Weekday) string {
	if len(sat.daysNarrow) == 0 {
		return ""
	}
	return sat.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sat *sat) WeekdaysNarrow() []string {
	return sat.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sat *sat) WeekdayShort(weekday time.Weekday) string {
	if len(sat.daysShort) == 0 {
		return ""
	}
	return sat.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sat *sat) WeekdaysShort() []string {
	return sat.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sat *sat) WeekdayWide(weekday time.Weekday) string {
	if len(sat.daysWide) == 0 {
		return ""
	}
	return sat.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sat *sat) WeekdaysWide() []string {
	return sat.daysWide
}

// Decimal returns the decimal point of number
func (sat *sat) Decimal() string {
	return sat.decimal
}

// Group returns the group of number
func (sat *sat) Group() string {
	return sat.group
}

// Group returns the minus sign of number
func (sat *sat) Minus() string {
	return sat.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sat' and handles both Whole and Real numbers based on 'v'
func (sat *sat) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sat.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sat.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sat.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sat' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sat *sat) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sat.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sat.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sat.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sat'
func (sat *sat) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sat.currencies[currency]
	l := len(s) + len(symbol) + 22

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sat.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, sat.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sat.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sat'
// in accounting notation.
func (sat *sat) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sat.currencies[currency]
	l := len(s) + len(symbol) + 22

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sat.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, sat.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, sat.currencyNegativeSuffix...)
	} else {

		b = append(b, sat.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sat'
func (sat *sat) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'sat'
func (sat *sat) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sat.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sat'
func (sat *sat) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sat.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sat'
func (sat *sat) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sat.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sat.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sat'
func (sat *sat) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sat'
func (sat *sat) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sat.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sat'
func (sat *sat) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sat.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sat'
func (sat *sat) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sat.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sat.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
