package mni_Mtei

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mni_Mtei struct {
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

// New returns a new instance of translator for the 'mni_Mtei' locale
func New() locales.Translator {
	return &mni_Mtei{
		locale:             "mni_Mtei",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "জন", "ফেব্রুৱারি", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "ওগ", "সেপ্টেম্বর", "ওক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		monthsNarrow:       []string{"", "জা", "ফে", "মার", "এপ", "মে", "জুন", "জুল", "আ", "সে", "ওক", "নব", "ডি"},
		monthsWide:         []string{"", "জনুৱারী", "ফেব্রুৱারি", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "\u200cওগষ্ট", "সেপ্টেম্বর", "ওক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		daysNarrow:         []string{"নোং", "নিং", "লৈবা", "য়ুম", "শগো", "ইরা", "থাং"},
		daysWide:           []string{"নোংমাইজিং", "নিংথৌকাবা", "লৈবাকপোকপা", "য়ুমশকৈশা", "শগোলশেন", "ইরাই", "থাংজ"},
		periodsAbbreviated: []string{"নুমাং", "PM"},
		periodsWide:        []string{"এ এম", "পি এম"},
		erasAbbreviated:    []string{"খৃ: মমাং", "খৃ: মতুং"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"", ""},
		timezones:          map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "অটলান্টিক দেলাইট টাইম", "ADT Arabia": "ADT Arabia", "AEDT": "AEDT", "AEST": "AEST", "AFT": "AFT", "AKDT": "AKDT", "AKST": "AKST", "AMST": "AMST", "AMST Armenia": "AMST Armenia", "AMT": "AMT", "AMT Armenia": "AMT Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ARST", "ART": "ART", "AST": "অটলান্টিক ষ্টেন্দর্দ টাইম", "AST Arabia": "AST Arabia", "AWDT": "AWDT", "AWST": "AWST", "AZST": "AZST", "AZT": "AZT", "BDT Bangladesh": "BDT Bangladesh", "BNT": "BNT", "BOT": "BOT", "BRST": "BRST", "BRT": "BRT", "BST Bangladesh": "BST Bangladesh", "BT": "BT", "CAST": "CAST", "CAT": "CAT", "CCT": "CCT", "CDT": "নোর্থ অমেরিকান সেন্ত্রেল দেলাইট টাইম", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "নোর্থ অমেরিকান সেন্ত্রেল ষ্টেন্দর্দ টাইম", "CST China": "CST China", "CST China DST": "CST China DST", "CVST": "CVST", "CVT": "CVT", "CXT": "CXT", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "CuDT", "CuST": "CuST", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "EASST", "EAST": "EAST", "EAT": "EAT", "ECT": "ECT", "EDT": "নোর্থ অমেরিকান ইষ্টর্ন দেলাইট টাইম", "EGDT": "EGDT", "EGST": "EGST", "EST": "নোর্থ অমেরিকান ইষ্টর্ন ষ্টেন্দর্দ টাইম", "FEET": "FEET", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "FKST", "FKT": "FKT", "FNST": "FNST", "FNT": "FNT", "GALT": "GALT", "GAMT": "GAMT", "GEST": "GEST", "GET": "GET", "GFT": "GFT", "GIT": "GIT", "GMT": "গ্রিনৱিচ মিন টাইম", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HKST": "HKST", "HKT": "HKT", "HOVST": "HOVST", "HOVT": "HOVT", "ICT": "ICT", "IDT": "IDT", "IOT": "IOT", "IRKST": "IRKST", "IRKT": "IRKT", "IRST": "IRST", "IRST DST": "IRST DST", "IST": "IST", "IST Israel": "IST Israel", "JDT": "JDT", "JST": "JST", "KOST": "KOST", "KRAST": "KRAST", "KRAT": "KRAT", "KST": "KST", "KST DST": "KST DST", "LHDT": "LHDT", "LHST": "LHST", "LINT": "LINT", "MAGST": "MAGST", "MAGT": "MAGT", "MART": "MART", "MAWT": "MAWT", "MDT": "নোর্থ অমেরিকান মাউন্টেন দেলাইট টাইম", "MESZ": "সেন্ত্রেল য়ুরোপিয়ান সমর টাইম", "MEZ": "সেন্ত্রেল য়ুরোপিয়ান ষ্টেন্দর্দ টাইম", "MHT": "MHT", "MMT": "MMT", "MSD": "MSD", "MST": "নোর্থ অমেরিকান মাউন্টেন ষ্টেন্দর্দ টাইম", "MUST": "MUST", "MUT": "MUT", "MVT": "MVT", "MYT": "MYT", "NCT": "NCT", "NDT": "NDT", "NDT New Caledonia": "NDT New Caledonia", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "NOVST", "NOVT": "NOVT", "NPT": "NPT", "NRT": "NRT", "NST": "NST", "NUT": "NUT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "ইষ্টর্ন য়ুরোপিয়ান সমর টাইম", "OEZ": "ইষ্টর্ন য়ুরোপিয়ান ষ্টেন্দর্দ টাইম", "OMSST": "OMSST", "OMST": "OMST", "PDT": "নোর্থ অমেরিকান পেসিফিক দেলাইট টাইম", "PDTM": "PDTM", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "PKT", "PKT DST": "PKT DST", "PMDT": "PMDT", "PMST": "PMST", "PONT": "PONT", "PST": "নোর্থ অমেরিকান পেসিফিক ষ্টেন্দর্দ টাইম", "PST Philippine": "PST Philippine", "PST Philippine DST": "PST Philippine DST", "PST Pitcairn": "PST Pitcairn", "PSTM": "PSTM", "PWT": "PWT", "PYST": "PYST", "PYT": "PYT", "PYT Korea": "PYT Korea", "RET": "RET", "ROTT": "ROTT", "SAKST": "SAKST", "SAKT": "SAKT", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "SAST", "SBT": "SBT", "SCT": "SCT", "SGT": "SGT", "SLST": "SLST", "SRT": "SRT", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "TAAF", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "TLT", "TMST": "TMST", "TMT": "TMT", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "TWT", "TWT DST": "TWT DST", "ULAST": "ULAST", "ULAT": "ULAT", "UYST": "UYST", "UYT": "UYT", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "VET", "VLAST": "VLAST", "VLAT": "VLAT", "VOLST": "VOLST", "VOLT": "VOLT", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "ৱেষ্টর্ন য়ুরোপিয়ান সমর টাইম", "WEZ": "ৱেষ্টর্ন য়ুরোপিয়ান ষ্টেন্দর্দ টাইম", "WFT": "WFT", "WGST": "WGST", "WGT": "WGT", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "YAKST": "YAKST", "YAKT": "YAKT", "YEKST": "YEKST", "YEKT": "YEKT", "YST": "YST", "МСК": "МСК", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "شىعىش قازاق ەلى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (mni *mni_Mtei) Locale() string {
	return mni.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mni_Mtei'
func (mni *mni_Mtei) PluralsCardinal() []locales.PluralRule {
	return mni.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mni_Mtei'
func (mni *mni_Mtei) PluralsOrdinal() []locales.PluralRule {
	return mni.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mni_Mtei'
func (mni *mni_Mtei) PluralsRange() []locales.PluralRule {
	return mni.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mni_Mtei'
func (mni *mni_Mtei) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mni_Mtei'
func (mni *mni_Mtei) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mni_Mtei'
func (mni *mni_Mtei) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mni *mni_Mtei) MonthAbbreviated(month time.Month) string {
	return mni.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mni *mni_Mtei) MonthsAbbreviated() []string {
	return mni.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mni *mni_Mtei) MonthNarrow(month time.Month) string {
	return mni.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mni *mni_Mtei) MonthsNarrow() []string {
	return mni.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mni *mni_Mtei) MonthWide(month time.Month) string {
	return mni.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mni *mni_Mtei) MonthsWide() []string {
	return mni.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mni *mni_Mtei) WeekdayAbbreviated(weekday time.Weekday) string {
	return mni.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mni *mni_Mtei) WeekdaysAbbreviated() []string {
	return mni.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mni *mni_Mtei) WeekdayNarrow(weekday time.Weekday) string {
	return mni.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mni *mni_Mtei) WeekdaysNarrow() []string {
	return mni.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mni *mni_Mtei) WeekdayShort(weekday time.Weekday) string {
	return mni.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mni *mni_Mtei) WeekdaysShort() []string {
	return mni.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mni *mni_Mtei) WeekdayWide(weekday time.Weekday) string {
	return mni.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mni *mni_Mtei) WeekdaysWide() []string {
	return mni.daysWide
}

// Decimal returns the decimal point of number
func (mni *mni_Mtei) Decimal() string {
	return mni.decimal
}

// Group returns the group of number
func (mni *mni_Mtei) Group() string {
	return mni.group
}

// Group returns the minus sign of number
func (mni *mni_Mtei) Minus() string {
	return mni.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mni_Mtei' and handles both Whole and Real numbers based on 'v'
func (mni *mni_Mtei) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mni_Mtei' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mni *mni_Mtei) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mni_Mtei'
func (mni *mni_Mtei) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mni.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mni_Mtei'
// in accounting notation.
func (mni *mni_Mtei) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mni.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

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

// FmtDateLong returns the long date representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mni.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mni.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mni.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mni_Mtei'
func (mni *mni_Mtei) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mni.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
