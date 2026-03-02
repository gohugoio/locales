package sid

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sid struct {
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

// New returns a new instance of translator for the 'sid' locale
func New() locales.Translator {
	return &sid{
		locale:             "sid",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "Br", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		daysAbbreviated:    []string{"Sam", "San", "Mak", "Row", "Ham", "Arb", "Qid"},
		daysNarrow:         []string{"S", "S", "M", "R", "H", "A", "Q"},
		daysWide:           []string{"Sambata", "Sanyo", "Maakisanyo", "Roowe", "Hamuse", "Arbe", "Qidaame"},
		periodsAbbreviated: []string{"soodo", "hawwaro"},
		timezones:          map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "ADT", "ADT Arabia": "ADT Arabia", "AEDT": "AEDT", "AEST": "AEST", "AFT": "AFT", "AKDT": "AKDT", "AKST": "AKST", "AMST": "AMST", "AMST Armenia": "AMST Armenia", "AMT": "AMT", "AMT Armenia": "AMT Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ARST", "ART": "ART", "AST": "AST", "AST Arabia": "AST Arabia", "AWDT": "AWDT", "AWST": "AWST", "AZST": "AZST", "AZT": "AZT", "BDT Bangladesh": "BDT Bangladesh", "BNT": "BNT", "BOT": "BOT", "BRST": "BRST", "BRT": "BRT", "BST Bangladesh": "BST Bangladesh", "BT": "BT", "CAST": "CAST", "CAT": "CAT", "CCT": "CCT", "CDT": "CDT", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "CST", "CST China": "CST China", "CST China DST": "CST China DST", "CVST": "CVST", "CVT": "CVT", "CXT": "CXT", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "CuDT", "CuST": "CuST", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "EASST", "EAST": "EAST", "EAT": "EAT", "ECT": "ECT", "EDT": "EDT", "EGDT": "EGDT", "EGST": "EGST", "EST": "EST", "FEET": "FEET", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "FKST", "FKT": "FKT", "FNST": "FNST", "FNT": "FNT", "GALT": "GALT", "GAMT": "GAMT", "GEST": "GEST", "GET": "GET", "GFT": "GFT", "GIT": "GIT", "GMT": "GMT", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HKST": "HKST", "HKT": "HKT", "HOVST": "HOVST", "HOVT": "HOVT", "ICT": "ICT", "IDT": "IDT", "IOT": "IOT", "IRKST": "IRKST", "IRKT": "IRKT", "IRST": "IRST", "IRST DST": "IRST DST", "IST": "IST", "IST Israel": "IST Israel", "JDT": "JDT", "JST": "JST", "KOST": "KOST", "KRAST": "KRAST", "KRAT": "KRAT", "KST": "KST", "KST DST": "KST DST", "LHDT": "LHDT", "LHST": "LHST", "LINT": "LINT", "MAGST": "MAGST", "MAGT": "MAGT", "MART": "MART", "MAWT": "MAWT", "MDT": "MDT", "MESZ": "MESZ", "MEZ": "MEZ", "MHT": "MHT", "MMT": "MMT", "MSD": "MSD", "MST": "MST", "MUST": "MUST", "MUT": "MUT", "MVT": "MVT", "MYT": "MYT", "NCT": "NCT", "NDT": "NDT", "NDT New Caledonia": "NDT New Caledonia", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "NOVST", "NOVT": "NOVT", "NPT": "NPT", "NRT": "NRT", "NST": "NST", "NUT": "NUT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "OESZ", "OEZ": "OEZ", "OMSST": "OMSST", "OMST": "OMST", "PDT": "PDT", "PDTM": "PDTM", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "PKT", "PKT DST": "PKT DST", "PMDT": "PMDT", "PMST": "PMST", "PONT": "PONT", "PST": "PST", "PST Philippine": "PST Philippine", "PST Philippine DST": "PST Philippine DST", "PST Pitcairn": "PST Pitcairn", "PSTM": "PSTM", "PWT": "PWT", "PYST": "PYST", "PYT": "PYT", "PYT Korea": "PYT Korea", "RET": "RET", "ROTT": "ROTT", "SAKST": "SAKST", "SAKT": "SAKT", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "SAST", "SBT": "SBT", "SCT": "SCT", "SGT": "SGT", "SLST": "SLST", "SRT": "SRT", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "TAAF", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "TLT", "TMST": "TMST", "TMT": "TMT", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "TWT", "TWT DST": "TWT DST", "ULAST": "ULAST", "ULAT": "ULAT", "UYST": "UYST", "UYT": "UYT", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "VET", "VLAST": "VLAST", "VLAT": "VLAT", "VOLST": "VOLST", "VOLT": "VOLT", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "WESZ", "WEZ": "WEZ", "WFT": "WFT", "WGST": "WGST", "WGT": "WGT", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "YAKST": "YAKST", "YAKT": "YAKT", "YEKST": "YEKST", "YEKT": "YEKT", "YST": "YST", "МСК": "МСК", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "شىعىش قازاق ەلى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (sid *sid) Locale() string {
	return sid.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sid'
func (sid *sid) PluralsCardinal() []locales.PluralRule {
	return sid.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sid'
func (sid *sid) PluralsOrdinal() []locales.PluralRule {
	return sid.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sid'
func (sid *sid) PluralsRange() []locales.PluralRule {
	return sid.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sid'
func (sid *sid) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sid'
func (sid *sid) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sid'
func (sid *sid) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sid *sid) MonthAbbreviated(month time.Month) string {
	return sid.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sid *sid) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sid *sid) MonthNarrow(month time.Month) string {
	return sid.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sid *sid) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (sid *sid) MonthWide(month time.Month) string {
	return sid.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sid *sid) MonthsWide() []string {
	return nil
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sid *sid) WeekdayAbbreviated(weekday time.Weekday) string {
	return sid.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sid *sid) WeekdaysAbbreviated() []string {
	return sid.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sid *sid) WeekdayNarrow(weekday time.Weekday) string {
	return sid.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sid *sid) WeekdaysNarrow() []string {
	return sid.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sid *sid) WeekdayShort(weekday time.Weekday) string {
	return sid.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sid *sid) WeekdaysShort() []string {
	return sid.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sid *sid) WeekdayWide(weekday time.Weekday) string {
	return sid.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sid *sid) WeekdaysWide() []string {
	return sid.daysWide
}

// Decimal returns the decimal point of number
func (sid *sid) Decimal() string {
	return sid.decimal
}

// Group returns the group of number
func (sid *sid) Group() string {
	return sid.group
}

// Group returns the minus sign of number
func (sid *sid) Minus() string {
	return sid.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sid' and handles both Whole and Real numbers based on 'v'
func (sid *sid) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sid.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sid.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sid.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sid' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sid *sid) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sid.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sid.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sid.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sid'
func (sid *sid) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sid.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sid.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sid.group[0])
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
		b = append(b, sid.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sid.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sid'
// in accounting notation.
func (sid *sid) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sid.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sid.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sid.group[0])
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

		b = append(b, sid.minus[0])

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
			b = append(b, sid.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sid'
func (sid *sid) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sid'
func (sid *sid) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, sid.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sid'
func (sid *sid) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sid.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sid'
func (sid *sid) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, sid.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, sid.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sid'
func (sid *sid) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sid.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sid.periodsAbbreviated[0]...)
	} else {
		b = append(b, sid.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sid'
func (sid *sid) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sid.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sid.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sid.periodsAbbreviated[0]...)
	} else {
		b = append(b, sid.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sid'
func (sid *sid) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sid.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sid.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sid.periodsAbbreviated[0]...)
	} else {
		b = append(b, sid.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sid'
func (sid *sid) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sid.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sid.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sid.periodsAbbreviated[0]...)
	} else {
		b = append(b, sid.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sid.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
