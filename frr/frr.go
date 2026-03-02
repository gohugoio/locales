package frr

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type frr struct {
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
	currencyPositivePrefix string
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'frr' locale
func New() locales.Translator {
	return &frr{
		locale:                 "frr",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "K",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "K",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mar", "Apr", "Mei", "Jün", "Jül", "Aug", "Sep", "Okt", "Nof", "Det"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Janewoore", "Febrewoore", "Maarts", "April", "Mei", "Jüüne", "Jüüle", "August", "September", "Oktuuber", "Nofember", "Detsember"},
		daysAbbreviated:        []string{"Sön", "Mun", "Tei", "Wed", "Tür", "Fre", "San"},
		daysShort:              []string{"Sö", "Mu", "Te", "We", "Tü", "Fr", "Sa"},
		daysWide:               []string{"Söndai", "Mundai", "Teisdai", "Weedensdai", "Tüürsdai", "Freidai", "Saninj"},
		periodsAbbreviated:     []string{"i/m", "e/m"},
		periodsNarrow:          []string{"i", "e"},
		periodsWide:            []string{"de iarmade", "de eftermade"},
		erasAbbreviated:        []string{"f.Kr.", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"föör Krast", "föör üüs tidjreegning"},
		timezones:              map[string]string{"ACDT": "Austraalisk Sentraal Somertidj", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Austraalisk Madelwaast Somertidj", "ACWST": "Austraalisk Madelwaast Standard Tidj", "ADT": "Ameerikoo Atlantik Somertidj", "ADT Arabia": "Araabisk Somertidj", "AEDT": "Uast Austraalisk Somertidj", "AEST": "Uast Austraalisk Standard Tidj", "AFT": "Afghaanistaan Tidj", "AKDT": "Alaska Somertidj", "AKST": "Alaska Standard Tidj", "AMST": "Amazonas Somertidj", "AMST Armenia": "Armeenien Somertidj", "AMT": "Amazonas Standard Tidj", "AMT Armenia": "Armeenien Standard Tidj", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Argentiinien Somertidj", "ART": "Argentiinien Standard Tidj", "AST": "Ameerikoo Atlantik Standard Tidj", "AST Arabia": "Araabisk Standard Tidj", "AWDT": "Waast Austraalisk Somertidj", "AWST": "Waast Austraalisk Standard Tidj", "AZST": "Aserbaidschaan Somertidj", "AZT": "Aserbaidschaan Standard Tidj", "BDT Bangladesh": "Bangladesch Somertidj", "BNT": "Brunei Tidj", "BOT": "Boliiwien Tidj", "BRST": "Brasiilien Somertidj", "BRT": "Brasiilien Standard Tidj", "BST Bangladesh": "Bangladesch Standard Tidj", "BT": "Bhuutaan Tidj", "CAST": "CAST", "CAT": "Sentraal Afrikoo Tidj", "CCT": "Cocos Eilunen Tidj", "CDT": "Sentraal Ameerikoo Somertidj", "CHADT": "Chatham Somertidj", "CHAST": "Chatham Standard Tidj", "CHUT": "Chuuk Tidj", "CKT": "Cook Eilunen Standard Tidj", "CKT DST": "Cook Eilunen Somertidj", "CLST": "Chiile Somertidj", "CLT": "Chiile Standard Tidj", "COST": "Kolumbien Somertidj", "COT": "Kolumbien Standard Tidj", "CST": "Sentraal Ameerikoo Standard Tidj", "CST China": "Schiina Standard Tidj", "CST China DST": "Schiina Somertidj", "CVST": "Kapwerden Sommertidj", "CVT": "Kapwerden Standard Tidj", "CXT": "Jul Eilun Tidj", "ChST": "Chamorro Tidj", "ChST NMI": "ChST NMI", "CuDT": "Kuuba Somertidj", "CuST": "Kuuba Standard Tidj", "DAVT": "Davis Tidj", "DDUT": "Dumont d’Urville Tidj", "EASST": "Puask Eilun Somertidj", "EAST": "Puask Eilun Standard Tidj", "EAT": "Uast Afrikoo Tidj", "ECT": "Ekwadoor Tidj", "EDT": "Uast Ameerikoo Somertidj", "EGDT": "Uast Greenlun Somertidj", "EGST": "Uast Greenlun Standard Tidj", "EST": "Uast Ameerikoo Standard Tidj", "FEET": "Faarder Uasteuropeesk Tidj", "FJT": "Fidschi Standard Tidj", "FJT Summer": "Fidschi Somertidj", "FKST": "Falklun Eilunen Somertidj", "FKT": "Falklun Eilunen Standard Tidj", "FNST": "Fernando de Noronha Somertidj", "FNT": "Fernando de Noronha Standard Tidj", "GALT": "Galapagos Tidj", "GAMT": "Gambier Tidj", "GEST": "Georgien Somertidj", "GET": "Georgien Standard Tidj", "GFT": "Fransöösk Guayaana Tidj", "GIT": "Gilbert Eilunen Tidj", "GMT": "Madel Greenwich Tidj", "GNSST": "GNSST", "GNST": "GNST", "GST": "Süüdgeorgien Tidj", "GST Guam": "GST Guam", "GYT": "Guyaana Tidj", "HADT": "Hawaii-Aleuten Somertidj", "HAST": "Hawaii-Aleuten Standard Tidj", "HKST": "Hongkong Somertidj", "HKT": "Hongkong Standard Tidj", "HOVST": "Chowd Somertidj", "HOVT": "Chowd Standard Tidj", "ICT": "Indoschiina Tidj", "IDT": "Israel Somertidj", "IOT": "Indisk Ooseaan Tidj", "IRKST": "Irkutsk Somertidj", "IRKT": "Irkutsk Standard Tidj", "IRST": "Iraan Standard Tidj", "IRST DST": "Iraan Somertidj", "IST": "Indisk Tidj", "IST Israel": "Israel Standard Tidj", "JDT": "Jaapan Somertidj", "JST": "Jaapan Standard Tidj", "KOST": "Kosrae Tidj", "KRAST": "Krasnojarsk Somertidj", "KRAT": "Krasnojarsk Standard Tidj", "KST": "Korea Standard Tidj", "KST DST": "Korea Somertidj", "LHDT": "Lord Howe Somertidj", "LHST": "Lord Howe Standard Tidj", "LINT": "Line Eilunen Tidj", "MAGST": "Magadan Somertidj", "MAGT": "Magadan Standard Tidj", "MART": "Marquesas Tidj", "MAWT": "Mawson Tidj", "MDT": "Waast Ameerikoo Somertidj", "MESZ": "Madeleuropeesk Somertidj", "MEZ": "Madeleuropeesk Standard Tidj", "MHT": "Marshall Eilunen Tidj", "MMT": "Mjanmaar Tidj", "MSD": "Moskau Somertidj", "MST": "Waast Ameerikoo Standard Tidj", "MUST": "Mauritius Somertidj", "MUT": "Mauritius Standard Tidj", "MVT": "Malediiwen Tidj", "MYT": "Malaysia Tidj", "NCT": "Neikaledoonien Standard Tidj", "NDT": "Neifundlun Somertidj", "NDT New Caledonia": "Neikaledoonien Somertidj", "NFDT": "Norfolk Eilun Somertidj", "NFT": "Norfolk Eilun Standard Tidj", "NOVST": "Nowosibirsk Somertidj", "NOVT": "Nowosibirsk Standard Tidj", "NPT": "Neepaal Tidj", "NRT": "Nauru Tidj", "NST": "Neifundlun Standard Tidj", "NUT": "Niue Tidj", "NZDT": "Neisialun Somertidj", "NZST": "Neisialun Standard Tidj", "OESZ": "Uasteuropeesk Somertidj", "OEZ": "Uasteuropeesk Standard Tidj", "OMSST": "Omsk Somertidj", "OMST": "Omsk Standard Tidj", "PDT": "Ameerikoo Pasiifik Somertidj", "PDTM": "Meksiko Pasiifik Somertidj", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papua Neiguinea Tidj", "PHOT": "Phoenix Eilunen Tidj", "PKT": "Pakistaan Standard Tidj", "PKT DST": "Pakistaan Somertidj", "PMDT": "St. Pierre an Miquelon Somertidj", "PMST": "St. Pierre an Miquelon Standard Tidj", "PONT": "Pohnpei Tidj", "PST": "Ameerikoo Pasiifik Standard Tidj", "PST Philippine": "Filipiinen Standard Tidj", "PST Philippine DST": "Filipiinen Somertidj", "PST Pitcairn": "Pitcairn Eilunen Tidj", "PSTM": "Meksiko Pasiifik Standard Tidj", "PWT": "Palau Tidj", "PYST": "Paraguay Somertidj", "PYT": "Paraguay Standard Tidj", "PYT Korea": "Pjöngjang Tidj", "RET": "Réunion Tidj", "ROTT": "Rothera Tidj", "SAKST": "Sachalin Somertidj", "SAKT": "Sachalin Standard Tidj", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Süüdelk Afrikoo Standard Tidj", "SBT": "Salomoonen Tidj", "SCT": "Seschelen Tidj", "SGT": "Singapuur Tidj", "SLST": "SLST", "SRT": "Suurinam Tidj", "SST Samoa": "Amerikoons Samoa Standard Tidj", "SST Samoa Apia": "Samoa Standard Tidj", "SST Samoa Apia DST": "Samoa Somertidj", "SST Samoa DST": "Amerikoons Samoa Somertidj", "SYOT": "Syowa Tidj", "TAAF": "Fransöösk Süüd an Antarktis Tidj", "TAHT": "Tahiti Tidj", "TJT": "Tadjikistaan Tidj", "TKT": "Tokelau Tidj", "TLT": "Uast Tiimor Tidj", "TMST": "Turkmeenistaan Somertidj", "TMT": "Turkmeenistaan Standard Tidj", "TOST": "Tonga Somertidj", "TOT": "Tonga Standard Tidj", "TVT": "Tuwaalu Tidj", "TWT": "Taiwan Standard Tidj", "TWT DST": "Taiwan Somertidj", "ULAST": "Ulaanbaatar Somertidj", "ULAT": "Ulaanbaatar Standard Tidj", "UYST": "Uruguay Somertidj", "UYT": "Uruguay Standard Tidj", "UZT": "Usbekistaan Standard Tidj", "UZT DST": "Usbekistaan Somertidj", "VET": "Weenesuela Tidj", "VLAST": "Wladiwostok Somertidj", "VLAT": "Wladiwostok Standard Tidj", "VOLST": "Wolgograd Somertidj", "VOLT": "Wolgograd Standard Tidj", "VOST": "Wostok Tidj", "VUT": "Vanuatu Standard Tidj", "VUT DST": "Vanuatu Somertidj", "WAKT": "Wake Eilun Tidj", "WARST": "Waast Argentiinien Somertidj", "WART": "Waast Argentiinien Standard Tidj", "WAST": "Waast Afrikoo Tidj", "WAT": "Waast Afrikoo Tidj", "WESZ": "Waasteuropeesk Somertidj", "WEZ": "Waasteuropeesk Standard Tidj", "WFT": "Wallis an Futuna Tidj", "WGST": "Waast Greenlun Somertidj", "WGT": "Waast Greenlun Standard Tidj", "WIB": "Waast Indoneesien Tidj", "WIT": "Uast Indoneesien Tidj", "WITA": "Sentraal Indoneesien Tidj", "YAKST": "Jakutsk Somertidj", "YAKT": "Jakutsk Standard Tidj", "YEKST": "Jekaterinburg Somertidj", "YEKT": "Jekaterinburg Standard Tidj", "YST": "Yukon Tidj", "МСК": "Moskau Standard Tidj", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Waast Kasachstaan Tidj", "شىعىش قازاق ەلى": "Uast Kasachstaan Tidj", "قازاق ەلى": "Kasachstaan Tidj", "قىرعىزستان": "Kirgistaan Tidj", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Peruu Somertidj"},
	}
}

// Locale returns the current translators string locale
func (frr *frr) Locale() string {
	return frr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'frr'
func (frr *frr) PluralsCardinal() []locales.PluralRule {
	return frr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'frr'
func (frr *frr) PluralsOrdinal() []locales.PluralRule {
	return frr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'frr'
func (frr *frr) PluralsRange() []locales.PluralRule {
	return frr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'frr'
func (frr *frr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'frr'
func (frr *frr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'frr'
func (frr *frr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (frr *frr) MonthAbbreviated(month time.Month) string {
	return frr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (frr *frr) MonthsAbbreviated() []string {
	return frr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (frr *frr) MonthNarrow(month time.Month) string {
	return frr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (frr *frr) MonthsNarrow() []string {
	return frr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (frr *frr) MonthWide(month time.Month) string {
	return frr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (frr *frr) MonthsWide() []string {
	return frr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (frr *frr) WeekdayAbbreviated(weekday time.Weekday) string {
	return frr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (frr *frr) WeekdaysAbbreviated() []string {
	return frr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (frr *frr) WeekdayNarrow(weekday time.Weekday) string {
	return frr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (frr *frr) WeekdaysNarrow() []string {
	return frr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (frr *frr) WeekdayShort(weekday time.Weekday) string {
	return frr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (frr *frr) WeekdaysShort() []string {
	return frr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (frr *frr) WeekdayWide(weekday time.Weekday) string {
	return frr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (frr *frr) WeekdaysWide() []string {
	return frr.daysWide
}

// Decimal returns the decimal point of number
func (frr *frr) Decimal() string {
	return frr.decimal
}

// Group returns the group of number
func (frr *frr) Group() string {
	return frr.group
}

// Group returns the minus sign of number
func (frr *frr) Minus() string {
	return frr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'frr' and handles both Whole and Real numbers based on 'v'
func (frr *frr) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'frr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (frr *frr) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'frr'
func (frr *frr) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := frr.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, frr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(frr.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, frr.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, frr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, frr.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'frr'
// in accounting notation.
func (frr *frr) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := frr.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, frr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(frr.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, frr.currencyNegativePrefix[j])
		}

		b = append(b, frr.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(frr.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, frr.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, frr.currencyNegativeSuffix...)
	} else {

		b = append(b, frr.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'frr'
func (frr *frr) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'frr'
func (frr *frr) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, frr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'frr'
func (frr *frr) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, frr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'frr'
func (frr *frr) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, frr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, frr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'frr'
func (frr *frr) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'frr'
func (frr *frr) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'frr'
func (frr *frr) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'frr'
func (frr *frr) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, frr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := frr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
