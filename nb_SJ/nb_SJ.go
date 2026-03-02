package nb_SJ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type nb_SJ struct {
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

// New returns a new instance of translator for the 'nb_SJ' locale
func New() locales.Translator {
	return &nb_SJ{
		locale:          "nb_SJ",
		pluralsCardinal: []locales.PluralRule{2, 6},
		pluralsOrdinal:  []locales.PluralRule{6},
		pluralsRange:    []locales.PluralRule{6},
		timeSeparator:   ":",
		currencies:      []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		timezones:       map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "ADT", "ADT Arabia": "ADT Arabia", "AEDT": "AEDT", "AEST": "AEST", "AFT": "AFT", "AKDT": "AKDT", "AKST": "AKST", "AMST": "AMST", "AMST Armenia": "AMST Armenia", "AMT": "AMT", "AMT Armenia": "AMT Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ARST", "ART": "ART", "AST": "AST", "AST Arabia": "AST Arabia", "AWDT": "AWDT", "AWST": "AWST", "AZST": "AZST", "AZT": "AZT", "BDT Bangladesh": "BDT Bangladesh", "BNT": "BNT", "BOT": "BOT", "BRST": "BRST", "BRT": "BRT", "BST Bangladesh": "BST Bangladesh", "BT": "BT", "CAST": "CAST", "CAT": "CAT", "CCT": "CCT", "CDT": "CDT", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "CST", "CST China": "CST China", "CST China DST": "CST China DST", "CVST": "CVST", "CVT": "CVT", "CXT": "CXT", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "CuDT", "CuST": "CuST", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "EASST", "EAST": "EAST", "EAT": "EAT", "ECT": "ECT", "EDT": "EDT", "EGDT": "EGDT", "EGST": "EGST", "EST": "EST", "FEET": "FEET", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "FKST", "FKT": "FKT", "FNST": "FNST", "FNT": "FNT", "GALT": "GALT", "GAMT": "GAMT", "GEST": "GEST", "GET": "GET", "GFT": "GFT", "GIT": "GIT", "GMT": "GMT", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HKST": "HKST", "HKT": "HKT", "HOVST": "HOVST", "HOVT": "HOVT", "ICT": "ICT", "IDT": "IDT", "IOT": "IOT", "IRKST": "IRKST", "IRKT": "IRKT", "IRST": "IRST", "IRST DST": "IRST DST", "IST": "IST", "IST Israel": "IST Israel", "JDT": "JDT", "JST": "JST", "KOST": "KOST", "KRAST": "KRAST", "KRAT": "KRAT", "KST": "KST", "KST DST": "KST DST", "LHDT": "LHDT", "LHST": "LHST", "LINT": "LINT", "MAGST": "MAGST", "MAGT": "MAGT", "MART": "MART", "MAWT": "MAWT", "MDT": "MDT", "MESZ": "MESZ", "MEZ": "MEZ", "MHT": "MHT", "MMT": "MMT", "MSD": "MSD", "MST": "MST", "MUST": "MUST", "MUT": "MUT", "MVT": "MVT", "MYT": "MYT", "NCT": "NCT", "NDT": "NDT", "NDT New Caledonia": "NDT New Caledonia", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "NOVST", "NOVT": "NOVT", "NPT": "NPT", "NRT": "NRT", "NST": "NST", "NUT": "NUT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "OESZ", "OEZ": "OEZ", "OMSST": "OMSST", "OMST": "OMST", "PDT": "PDT", "PDTM": "PDTM", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "PKT", "PKT DST": "PKT DST", "PMDT": "PMDT", "PMST": "PMST", "PONT": "PONT", "PST": "PST", "PST Philippine": "PST Philippine", "PST Philippine DST": "PST Philippine DST", "PST Pitcairn": "PST Pitcairn", "PSTM": "PSTM", "PWT": "PWT", "PYST": "PYST", "PYT": "PYT", "PYT Korea": "PYT Korea", "RET": "RET", "ROTT": "ROTT", "SAKST": "SAKST", "SAKT": "SAKT", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "SAST", "SBT": "SBT", "SCT": "SCT", "SGT": "SGT", "SLST": "SLST", "SRT": "SRT", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "TAAF", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "TLT", "TMST": "TMST", "TMT": "TMT", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "TWT", "TWT DST": "TWT DST", "ULAST": "ULAST", "ULAT": "ULAT", "UYST": "UYST", "UYT": "UYT", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "VET", "VLAST": "VLAST", "VLAT": "VLAT", "VOLST": "VOLST", "VOLT": "VOLT", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "WESZ", "WEZ": "WEZ", "WFT": "WFT", "WGST": "WGST", "WGT": "WGT", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "YAKST": "YAKST", "YAKT": "YAKT", "YEKST": "YEKST", "YEKT": "YEKT", "YST": "YST", "МСК": "МСК", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "شىعىش قازاق ەلى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (nb *nb_SJ) Locale() string {
	return nb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nb_SJ'
func (nb *nb_SJ) PluralsCardinal() []locales.PluralRule {
	return nb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nb_SJ'
func (nb *nb_SJ) PluralsOrdinal() []locales.PluralRule {
	return nb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nb_SJ'
func (nb *nb_SJ) PluralsRange() []locales.PluralRule {
	return nb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nb_SJ'
func (nb *nb_SJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nb_SJ'
func (nb *nb_SJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nb_SJ'
func (nb *nb_SJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nb *nb_SJ) MonthAbbreviated(month time.Month) string {
	return nb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nb *nb_SJ) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nb *nb_SJ) MonthNarrow(month time.Month) string {
	return nb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nb *nb_SJ) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (nb *nb_SJ) MonthWide(month time.Month) string {
	return nb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nb *nb_SJ) MonthsWide() []string {
	return nil
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nb *nb_SJ) WeekdayAbbreviated(weekday time.Weekday) string {
	return nb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nb *nb_SJ) WeekdaysAbbreviated() []string {
	return nb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nb *nb_SJ) WeekdayNarrow(weekday time.Weekday) string {
	return nb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nb *nb_SJ) WeekdaysNarrow() []string {
	return nb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nb *nb_SJ) WeekdayShort(weekday time.Weekday) string {
	return nb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nb *nb_SJ) WeekdaysShort() []string {
	return nb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nb *nb_SJ) WeekdayWide(weekday time.Weekday) string {
	return nb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nb *nb_SJ) WeekdaysWide() []string {
	return nb.daysWide
}

// Decimal returns the decimal point of number
func (nb *nb_SJ) Decimal() string {
	return nb.decimal
}

// Group returns the group of number
func (nb *nb_SJ) Group() string {
	return nb.group
}

// Group returns the minus sign of number
func (nb *nb_SJ) Minus() string {
	return nb.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nb_SJ' and handles both Whole and Real numbers based on 'v'
func (nb *nb_SJ) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nb_SJ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nb *nb_SJ) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nb_SJ'
func (nb *nb_SJ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nb.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nb_SJ'
// in accounting notation.
func (nb *nb_SJ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nb.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nb_SJ'
func (nb *nb_SJ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}
