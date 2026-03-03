package sat_Deva

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sat_Deva struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'sat_Deva' locale
func New() locales.Translator {
	return &sat_Deva{
		locale:                 "sat_Deva",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "֏", "ANG", "Kz", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "KM", "BAN", "$", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "$", "Bs", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$", "BTN", "BUK", "P", "BYB", "BYN", "BYR", "$", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$", "CNH", "CNX", "CN¥", "$", "COU", "₡", "CSD", "CSK", "$", "$", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "kr", "$", "DZD", "ECS", "ECV", "EEK", "E£", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "$", "£", "FRF", "£", "GEK", "₾", "GHC", "GH₵", "£", "GMD", "FG", "GNS", "GQE", "GRD", "Q", "GWE", "GWP", "$", "HK$", "L", "HRD", "kn", "HTG", "Ft", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "kr", "ITL", "$", "JOD", "JP¥", "KES", "⃀", "៛", "CF", "₩", "KRH", "KRO", "₩", "KWD", "$", "₸", "₭", "L£", "Rs", "$", "LSL", "Lt", "LTT", "LUC", "LUF", "LUL", "Ls", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "Ar", "MGF", "MKD", "MKN", "MLF", "K", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "Rs", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "$", "₦", "NIC", "C$", "NLG", "kr", "Rs", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "Rs", "zł", "PLZ", "PTE", "₲", "QAR", "RHD", "ROL", "lei", "RSD", "₽", "RUR", "RF", "\u20c1", "$", "SCR", "SDD", "SDG", "SDP", "kr", "$", "£", "SIT", "SKK", "SLE", "SLL", "SOS", "$", "SRG", "£", "STD", "Db", "SUR", "SVC", "£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "₺", "$", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "$", "UYW", "UZS", "VEB", "VED", "Bs", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "Cg.", "XDR", "XEU", "XFO", "XFU", "F\u202fCFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "¤", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ᱡᱟᱱ", "ᱯᱷᱟ", "ᱢᱟᱨ", "ᱟᱯᱨ", "ᱢᱮ", "ᱡᱩᱱ", "ᱡᱩᱞ", "ᱟᱜᱟ", "ᱥᱮᱯ", "ᱚᱠᱴ", "ᱱᱟᱣ", "ᱫᱤᱥ"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "M01", "M02", "M03", "M04", "M05", "M06", "M07", "M08", "M09", "M10", "M11", "M12"},
		daysAbbreviated:        []string{"ᱥᱤᱸ", "ᱚᱛ", "ᱵᱟ", "ᱥᱟᱹ", "ᱥᱟᱹᱨ", "ᱡᱟᱹ", "ᱧᱩ"},
		daysNarrow:             []string{"S", "M", "T", "W", "T", "F", "S"},
		daysWide:               []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		periodsAbbreviated:     []string{"AM", "PM"},
		timezones:              map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "ᱮᱴᱞᱟᱱᱴᱤᱠ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "ADT Arabia": "ADT Arabia", "AEDT": "AEDT", "AEST": "AEST", "AFT": "AFT", "AKDT": "ᱟᱹᱞᱟᱥᱠᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "AKST": "ᱟᱹᱞᱟᱥᱠᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "AMST": "ᱟᱢᱟᱡᱚᱱ ᱥᱤᱛᱩᱝ ᱚᱠᱴᱚ", "AMST Armenia": "AMST Armenia", "AMT": "ᱟᱢᱟᱡᱚᱱ ᱮᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱴᱚ", "AMT Armenia": "AMT Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱥᱤᱛᱩᱝ ᱚᱠᱴᱚ", "ART": "ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱮᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱴᱚ", "AST": "ᱮᱴᱞᱟᱱᱴᱤᱠ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "AST Arabia": "AST Arabia", "AWDT": "AWDT", "AWST": "AWST", "AZST": "AZST", "AZT": "AZT", "BDT Bangladesh": "BDT Bangladesh", "BNT": "BNT", "BOT": "BOT", "BRST": "BRST", "BRT": "BRT", "BST Bangladesh": "BST Bangladesh", "BT": "BT", "CAST": "CAST", "CAT": "CAT", "CCT": "CCT", "CDT": "ᱛᱟᱱᱟᱞᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "ᱛᱟᱱᱟᱞᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "CST China": "CST China", "CST China DST": "CST China DST", "CVST": "CVST", "CVT": "CVT", "CXT": "CXT", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "ᱠᱭᱩᱵᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "CuST": "ᱠᱭᱩᱵᱟ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "EASST", "EAST": "EAST", "EAT": "EAT", "ECT": "ECT", "EDT": "ᱤᱥᱴᱟᱨᱱ ᱥᱤᱧᱟᱜ ᱵᱚᱠᱛᱚ", "EGDT": "ᱥᱟᱢᱟᱝ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱥᱤᱛᱩᱝ ᱚᱠᱛᱚ", "EGST": "ᱥᱟᱢᱟᱝ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "EST": "ᱤᱥᱴᱟᱨᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "FEET": "FEET", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "FKST", "FKT": "FKT", "FNST": "FNST", "FNT": "FNT", "GALT": "GALT", "GAMT": "GAMT", "GEST": "GEST", "GET": "GET", "GFT": "GFT", "GIT": "GIT", "GMT": "ᱜᱨᱤᱱᱣᱤᱪ ᱢᱤᱱ ᱚᱠᱛᱚ", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "GYT", "HADT": "ᱦᱟᱣᱟᱭᱤᱼᱟᱞᱮᱣᱴᱤᱭᱟᱱ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "HAST": "ᱦᱟᱣᱟᱭᱤᱼᱟᱞᱮᱣᱴᱤᱭᱟᱱ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "HKST": "HKST", "HKT": "HKT", "HOVST": "HOVST", "HOVT": "HOVT", "ICT": "ICT", "IDT": "IDT", "IOT": "IOT", "IRKST": "IRKST", "IRKT": "IRKT", "IRST": "IRST", "IRST DST": "IRST DST", "IST": "IST", "IST Israel": "IST Israel", "JDT": "JDT", "JST": "JST", "KOST": "KOST", "KRAST": "KRAST", "KRAT": "KRAT", "KST": "KST", "KST DST": "KST DST", "LHDT": "LHDT", "LHST": "LHST", "LINT": "LINT", "MAGST": "MAGST", "MAGT": "MAGT", "MART": "MART", "MAWT": "MAWT", "MDT": "MDT", "MESZ": "ᱥᱮᱱᱴᱨᱟᱞ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "MEZ": "ᱥᱮᱱᱴᱨᱟᱞ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "MHT": "MHT", "MMT": "MMT", "MSD": "MSD", "MST": "MST", "MUST": "MUST", "MUT": "MUT", "MVT": "MVT", "MYT": "MYT", "NCT": "NCT", "NDT": "ᱱᱤᱭᱩᱯᱷᱟᱩᱱᱰᱞᱮᱸᱰ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "NDT New Caledonia": "NDT New Caledonia", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "NOVST", "NOVT": "NOVT", "NPT": "NPT", "NRT": "NRT", "NST": "ᱱᱤᱭᱩᱯᱷᱟᱩᱱᱰᱞᱮᱸᱰ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "NUT": "NUT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "ᱤᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "OEZ": "ᱤᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "OMSST": "OMSST", "OMST": "OMST", "PDT": "ᱯᱮᱥᱤᱯᱷᱤᱠ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PDTM": "ᱢᱮᱠᱥᱤᱠᱟᱱ ᱯᱨᱚᱥᱟᱱᱛ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "PKT", "PKT DST": "PKT DST", "PMDT": "ᱮᱥ ᱴᱤ ᱯᱤᱭᱮᱨᱮ ᱟᱨ ᱢᱤᱠᱮᱞᱚᱱ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PMST": "ᱮᱥ ᱴᱤ ᱯᱤᱭᱮᱨᱮ ᱟᱨ ᱢᱤᱠᱮᱞᱚᱱ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "PONT": "PONT", "PST": "ᱯᱮᱥᱤᱯᱷᱤᱠ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "PST Philippine": "PST Philippine", "PST Philippine DST": "PST Philippine DST", "PST Pitcairn": "PST Pitcairn", "PSTM": "ᱢᱮᱠᱥᱤᱠᱟᱱ ᱯᱨᱚᱥᱟᱱᱛ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "PWT": "PWT", "PYST": "PYST", "PYT": "PYT", "PYT Korea": "PYT Korea", "RET": "RET", "ROTT": "ROTT", "SAKST": "SAKST", "SAKT": "SAKT", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "SAST", "SBT": "SBT", "SCT": "SCT", "SGT": "SGT", "SLST": "SLST", "SRT": "SRT", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "TAAF", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "TLT", "TMST": "TMST", "TMT": "TMT", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "TWT", "TWT DST": "TWT DST", "ULAST": "ULAST", "ULAT": "ULAT", "UYST": "UYST", "UYT": "UYT", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "VET", "VLAST": "VLAST", "VLAT": "VLAT", "VOLST": "VOLST", "VOLT": "VOLT", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "ᱯᱟᱪᱮ ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱥᱤᱛᱩᱝ ᱚᱠᱴᱚ", "WART": "ᱯᱟᱪᱮ ᱟᱹᱨᱡᱮᱱᱴᱤᱱᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱴᱚ", "WAST": "WAST", "WAT": "WAT", "WESZ": "ᱣᱮᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "WEZ": "ᱣᱮᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "WFT": "WFT", "WGST": "ᱯᱟᱪᱮ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱥᱤᱛᱩᱝ ᱚᱠᱛᱚ", "WGT": "ᱯᱟᱪᱮ ᱜᱽᱨᱤᱱᱞᱮᱱᱰ ᱢᱟᱱᱚᱠ ᱚᱠᱛᱚ", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "YAKST": "YAKST", "YAKT": "YAKT", "YEKST": "YEKST", "YEKT": "YEKT", "YST": "ᱭᱩᱠᱚᱱ ᱚᱠᱛᱚ", "МСК": "МСК", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "شىعىش قازاق ەلى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (sat *sat_Deva) Locale() string {
	return sat.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sat_Deva'
func (sat *sat_Deva) PluralsCardinal() []locales.PluralRule {
	return sat.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sat_Deva'
func (sat *sat_Deva) PluralsOrdinal() []locales.PluralRule {
	return sat.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sat_Deva'
func (sat *sat_Deva) PluralsRange() []locales.PluralRule {
	return sat.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sat_Deva'
func (sat *sat_Deva) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sat_Deva'
func (sat *sat_Deva) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sat_Deva'
func (sat *sat_Deva) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sat *sat_Deva) MonthAbbreviated(month time.Month) string {
	if len(sat.monthsAbbreviated) == 0 {
		return ""
	}
	return sat.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sat *sat_Deva) MonthsAbbreviated() []string {
	return sat.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sat *sat_Deva) MonthNarrow(month time.Month) string {
	if len(sat.monthsNarrow) == 0 {
		return ""
	}
	return sat.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sat *sat_Deva) MonthsNarrow() []string {
	return sat.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sat *sat_Deva) MonthWide(month time.Month) string {
	if len(sat.monthsWide) == 0 {
		return ""
	}
	return sat.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sat *sat_Deva) MonthsWide() []string {
	return sat.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sat *sat_Deva) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sat.daysAbbreviated) == 0 {
		return ""
	}
	return sat.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sat *sat_Deva) WeekdaysAbbreviated() []string {
	return sat.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sat *sat_Deva) WeekdayNarrow(weekday time.Weekday) string {
	if len(sat.daysNarrow) == 0 {
		return ""
	}
	return sat.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sat *sat_Deva) WeekdaysNarrow() []string {
	return sat.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sat *sat_Deva) WeekdayShort(weekday time.Weekday) string {
	if len(sat.daysShort) == 0 {
		return ""
	}
	return sat.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sat *sat_Deva) WeekdaysShort() []string {
	return sat.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sat *sat_Deva) WeekdayWide(weekday time.Weekday) string {
	if len(sat.daysWide) == 0 {
		return ""
	}
	return sat.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sat *sat_Deva) WeekdaysWide() []string {
	return sat.daysWide
}

// Decimal returns the decimal point of number
func (sat *sat_Deva) Decimal() string {
	return sat.decimal
}

// Group returns the group of number
func (sat *sat_Deva) Group() string {
	return sat.group
}

// Group returns the minus sign of number
func (sat *sat_Deva) Minus() string {
	return sat.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sat_Deva' and handles both Whole and Real numbers based on 'v'
func (sat *sat_Deva) FmtNumber(num float64, v uint64) string {

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

// FmtPercent returns 'num' with digits/precision of 'v' for 'sat_Deva' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sat *sat_Deva) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sat_Deva'
func (sat *sat_Deva) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sat.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
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

	for j := len(sat.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, sat.currencyPositivePrefix[j])
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

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sat.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sat_Deva'
// in accounting notation.
func (sat *sat_Deva) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sat.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
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

		for j := len(sat.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sat.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, sat.minus[0])

	} else {

		for j := len(sat.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, sat.currencyPositivePrefix[j])
		}

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
			b = append(b, sat.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtDateLong(t time.Time) string {

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

// FmtDateFull returns the full date representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtDateFull(t time.Time) string {

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

// FmtTimeShort returns the short time representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtTimeShort(t time.Time) string {

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

// FmtTimeMedium returns the medium time representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtTimeMedium(t time.Time) string {

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

// FmtTimeLong returns the long time representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtTimeLong(t time.Time) string {

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

// FmtTimeFull returns the full time representation of 't' for 'sat_Deva'
func (sat *sat_Deva) FmtTimeFull(t time.Time) string {

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
