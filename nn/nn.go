package nn

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type nn struct {
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

// New returns a new instance of translator for the 'nn' locale
func New() locales.Translator {
	return &nn{
		locale:             "nn",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "lei", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		daysAbbreviated:    []string{"sø.", "må.", "ty.", "on.", "to.", "fr.", "la."},
		daysShort:          []string{"sø.", "må.", "ty.", "on.", "to.", "fr.", "la."},
		daysWide:           []string{"søndag", "måndag", "tysdag", "onsdag", "torsdag", "fredag", "laurdag"},
		periodsAbbreviated: []string{"f.m.", "e.m."},
		periodsWide:        []string{"", ""},
		erasAbbreviated:    []string{"", ""},
		erasNarrow:         []string{"fvt", "vt"},
		erasWide:           []string{"før vår tidsrekning", "etter vår tidsrekning"},
		timezones:          map[string]string{"ACDT": "sentralaustralsk sommartid", "ACST": "sentralaustralsk standardtid", "ACT": "ACT", "ACWDT": "vest-sentralaustralsk sommartid", "ACWST": "vest-sentralaustralsk standardtid", "ADT": "sommartid for den nordamerikanske atlanterhavskysten", "ADT Arabia": "arabisk sommartid", "AEDT": "austaustralsk sommartid", "AEST": "austaustralsk standardtid", "AFT": "AFT", "AKDT": "alaskisk sommartid", "AKST": "alaskisk normaltid", "AMST": "sommartid for Amazonas", "AMST Armenia": "armensk sommartid", "AMT": "normaltid for Amazonas", "AMT Armenia": "armensk normaltid", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "argentinsk sommartid", "ART": "argentinsk normaltid", "AST": "normaltid for den nordamerikanske atlanterhavskysten", "AST Arabia": "arabisk normaltid", "AWDT": "vestaustralsk sommartid", "AWST": "vestaustralsk standardtid", "AZST": "aserbajdsjansk sommartid", "AZT": "aserbajdsjansk normaltid", "BDT Bangladesh": "bangladeshisk sommartid", "BNT": "BNT", "BOT": "BOT", "BRST": "sommartid for Brasilia", "BRT": "normaltid for Brasilia", "BST Bangladesh": "bangladeshisk normaltid", "BT": "BT", "CAST": "CAST", "CAT": "CAT", "CCT": "tidssone for Kokosøyane", "CDT": "sommartid for sentrale Nord-Amerika", "CHADT": "sommartid for Chatham", "CHAST": "normaltid for Chatham", "CHUT": "tidssone for Chuukøyane", "CKT": "normaltid for Cookøyane", "CKT DST": "sommartid for Cookøyane", "CLST": "chilensk sommartid", "CLT": "chilensk normaltid", "COST": "kolombiansk sommartid", "COT": "kolombiansk normaltid", "CST": "normaltid for sentrale Nord-Amerika", "CST China": "kinesisk normaltid", "CST China DST": "kinesisk sommartid", "CVST": "kappverdisk sommartid", "CVT": "kappverdisk normaltid", "CXT": "CXT", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "kubansk sommartid", "CuST": "kubansk normaltid", "DAVT": "DAVT", "DDUT": "tidssone for Dumont-d’Urville", "EASST": "sommartid for Påskeøya", "EAST": "normaltid for Påskeøya", "EAT": "austafrikansk tid", "ECT": "ECT", "EDT": "sommartid for den nordamerikanske austkysten", "EGDT": "austgrønlandsk sommartid", "EGST": "austgrønlandsk normaltid", "EST": "normaltid for den nordamerikanske austkysten", "FEET": "fjern-austeuropeisk tid", "FJT": "fijiansk normaltid", "FJT Summer": "fijiansk sommartid", "FKST": "sommartid for Falklandsøyane", "FKT": "normaltid for Falklandsøyane", "FNST": "sommartid for Fernando de Noronha", "FNT": "normaltid for Fernando de Noronha", "GALT": "tidssone for Galápagosøyane", "GAMT": "GAMT", "GEST": "georgisk sommartid", "GET": "georgisk normaltid", "GFT": "GFT", "GIT": "tidssone for Gilbertøyane", "GMT": "GMT", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "GYT", "HADT": "sommartid for Hawaii og Aleutene", "HAST": "normaltid for Hawaii og Aleutene", "HKST": "hongkongkinesisk sommartid", "HKT": "hongkongkinesisk normaltid", "HOVST": "sommartid for Khovd", "HOVT": "normaltid for Khovd", "ICT": "ICT", "IDT": "israelsk sommartid", "IOT": "IOT", "IRKST": "sommartid for Irkutsk", "IRKT": "normaltid for Irkutsk", "IRST": "iransk normaltid", "IRST DST": "iransk sommartid", "IST": "IST", "IST Israel": "israelsk normaltid", "JDT": "japansk sommartid", "JST": "japansk normaltid", "KOST": "KOST", "KRAST": "sommartid for Krasnojarsk", "KRAT": "normaltid for Krasnojarsk", "KST": "koreansk normaltid", "KST DST": "koreansk sommartid", "LHDT": "sommartid for Lord Howe-øya", "LHST": "normaltid for Lord Howe-øya", "LINT": "tidssone for Lineøyane", "MAGST": "sommartid for Magadan", "MAGT": "normaltid for Magadan", "MART": "tidssone for Marquesasøyane", "MAWT": "MAWT", "MDT": "sommartid for Rocky Mountains (USA)", "MESZ": "sentraleuropeisk sommartid", "MEZ": "sentraleuropeisk standardtid", "MHT": "MHT", "MMT": "MMT", "MSD": "sommartid for Moskva", "MST": "normaltid for Rocky Mountains (USA)", "MUST": "mauritisk sommartid", "MUT": "mauritisk normaltid", "MVT": "MVT", "MYT": "MYT", "NCT": "kaledonsk normaltid", "NDT": "sommartid for Newfoundland", "NDT New Caledonia": "kaledonsk sommartid", "NFDT": "sommartid for Norfolkøya", "NFT": "normaltid for Norfolkøya", "NOVST": "sommartid for Novosibirsk", "NOVT": "normaltid for Novosibirsk", "NPT": "NPT", "NRT": "NRT", "NST": "normaltid for Newfoundland", "NUT": "NUT", "NZDT": "nyzealandsk sommartid", "NZST": "nyzealandsk normaltid", "OESZ": "austeuropeisk sommartid", "OEZ": "austeuropeisk standardtid", "OMSST": "sommartid for Omsk", "OMST": "normaltid for Omsk", "PDT": "sommartid for den nordamerikanske stillehavskysten", "PDTM": "sommartid for den meksikanske stillehavskysten", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "tidssone for Phoenixøyane", "PKT": "pakistansk normaltid", "PKT DST": "pakistansk sommartid", "PMDT": "sommartid for Saint-Pierre-et-Miquelon", "PMST": "normaltid for Saint-Pierre-et-Miquelon", "PONT": "PONT", "PST": "normaltid for den nordamerikanske stillehavskysten", "PST Philippine": "filippinsk normaltid", "PST Philippine DST": "filippinsk sommartid", "PST Pitcairn": "PST Pitcairn", "PSTM": "normaltid for den meksikanske stillehavskysten", "PWT": "PWT", "PYST": "paraguayansk sommartid", "PYT": "paraguayansk normaltid", "PYT Korea": "PYT Korea", "RET": "RET", "ROTT": "ROTT", "SAKST": "sommartid for Sakhalin", "SAKT": "normaltid for Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "SAST", "SBT": "SBT", "SCT": "SCT", "SGT": "SGT", "SLST": "SLST", "SRT": "SRT", "SST Samoa": "samoansk normaltid", "SST Samoa Apia": "normaltid for Apia", "SST Samoa Apia DST": "sommartid for Apia", "SST Samoa DST": "samoansk sommartid", "SYOT": "SYOT", "TAAF": "tidssone for Dei franske sørterritoria", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "austtimoresisk tid", "TMST": "turkmensk sommartid", "TMT": "turkmensk normaltid", "TOST": "tongansk sommartid", "TOT": "tongansk normaltid", "TVT": "TVT", "TWT": "normaltid for Taipei", "TWT DST": "sommartid for Taipei", "ULAST": "sommartid for Ulan Bator", "ULAT": "normaltid for Ulan Bator", "UYST": "uruguayansk sommartid", "UYT": "uruguayansk normaltid", "UZT": "usbekisk normaltid", "UZT DST": "usbekisk sommartid", "VET": "VET", "VLAST": "sommartid for Vladivostok", "VLAT": "normaltid for Vladivostok", "VOLST": "sommartid for Volgograd", "VOLT": "normaltid for Volgograd", "VOST": "VOST", "VUT": "vanuatisk normaltid", "VUT DST": "vanuatisk sommartid", "WAKT": "WAKT", "WARST": "vestargentinsk sommartid", "WART": "vestargentinsk normaltid", "WAST": "WAST", "WAT": "WAT", "WESZ": "vesteuropeisk sommartid", "WEZ": "vesteuropeisk standardtid", "WFT": "tidssone for Wallis- og Futunaøyane", "WGST": "vestgrønlandsk sommartid", "WGT": "vestgrønlandsk normaltid", "WIB": "WIB", "WIT": "austindonesisk tid", "WITA": "WITA", "YAKST": "sommartid for Jakutsk", "YAKT": "normaltid for Jakutsk", "YEKST": "sommartid for Jekaterinburg", "YEKT": "normaltid for Jekaterinburg", "YST": "YST", "МСК": "normaltid for Moskva", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "austkasakhstansk tid", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "asorisk sommartid"},
	}
}

// Locale returns the current translators string locale
func (nn *nn) Locale() string {
	return nn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nn'
func (nn *nn) PluralsCardinal() []locales.PluralRule {
	return nn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nn'
func (nn *nn) PluralsOrdinal() []locales.PluralRule {
	return nn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nn'
func (nn *nn) PluralsRange() []locales.PluralRule {
	return nn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nn'
func (nn *nn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nn'
func (nn *nn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nn'
func (nn *nn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nn *nn) MonthAbbreviated(month time.Month) string {
	return nn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nn *nn) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nn *nn) MonthNarrow(month time.Month) string {
	return nn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nn *nn) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (nn *nn) MonthWide(month time.Month) string {
	return nn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nn *nn) MonthsWide() []string {
	return nil
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nn *nn) WeekdayAbbreviated(weekday time.Weekday) string {
	return nn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nn *nn) WeekdaysAbbreviated() []string {
	return nn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nn *nn) WeekdayNarrow(weekday time.Weekday) string {
	return nn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nn *nn) WeekdaysNarrow() []string {
	return nn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nn *nn) WeekdayShort(weekday time.Weekday) string {
	return nn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nn *nn) WeekdaysShort() []string {
	return nn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nn *nn) WeekdayWide(weekday time.Weekday) string {
	return nn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nn *nn) WeekdaysWide() []string {
	return nn.daysWide
}

// Decimal returns the decimal point of number
func (nn *nn) Decimal() string {
	return nn.decimal
}

// Group returns the group of number
func (nn *nn) Group() string {
	return nn.group
}

// Group returns the minus sign of number
func (nn *nn) Minus() string {
	return nn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nn' and handles both Whole and Real numbers based on 'v'
func (nn *nn) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nn *nn) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nn'
func (nn *nn) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nn.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nn'
// in accounting notation.
func (nn *nn) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nn.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'nn'
func (nn *nn) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'nn'
func (nn *nn) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nn'
func (nn *nn) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nn'
func (nn *nn) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nn'
func (nn *nn) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nn'
func (nn *nn) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nn'
func (nn *nn) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nn'
func (nn *nn) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
