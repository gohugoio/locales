package ia_001

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ia_001 struct {
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

// New returns a new instance of translator for the 'ia_001' locale
func New() locales.Translator {
	return &ia_001{
		locale:                 "ia_001",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: ")",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "mai", "jun", "jul", "aug", "sep", "oct", "nov", "dec"},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januario", "februario", "martio", "april", "maio", "junio", "julio", "augusto", "septembre", "octobre", "novembre", "decembre"},
		daysAbbreviated:        []string{"dom", "lun", "mar", "mer", "jov", "ven", "sab"},
		daysNarrow:             []string{"d", "l", "m", "m", "j", "v", "s"},
		daysShort:              []string{"do", "lu", "ma", "me", "jo", "ve", "sa"},
		daysWide:               []string{"dominica", "lunedi", "martedi", "mercuridi", "jovedi", "venerdi", "sabbato"},
		erasAbbreviated:        []string{"a.Chr.", "p.Chr."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ante Christo", "post Christo"},
		timezones:              map[string]string{"ACDT": "hora estive de Australia central", "ACST": "hora estive de Acre", "ACT": "hora normal de Acre", "ACWDT": "hora estive de Australia centro-occidental", "ACWST": "hora normal de Australia centro-occidental", "ADT": "hora estive atlantic", "ADT Arabia": "hora estive arabe", "AEDT": "hora estive de Australia oriental", "AEST": "hora normal de Australia oriental", "AFT": "hora de Afghanistan", "AKDT": "hora estive de Alaska", "AKST": "hora normal de Alaska", "AMST": "hora estive de Amazonia", "AMST Armenia": "hora estive de Armenia", "AMT": "hora normal de Amazonia", "AMT Armenia": "hora normal de Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "hora estive de Argentina", "ART": "hora normal de Argentina", "AST": "hora normal atlantic", "AST Arabia": "hora normal arabe", "AWDT": "hora estive de Australia occidental", "AWST": "hora normal de Australia occidental", "AZST": "hora estive de Azerbeidzhan", "AZT": "hora normal de Azerbeidzhan", "BDT Bangladesh": "hora estive de Bangladesh", "BNT": "hora de Brunei", "BOT": "hora de Bolivia", "BRST": "hora estive de Brasilia", "BRT": "hora normal de Brasilia", "BST Bangladesh": "hora normal de Bangladesh", "BT": "hora de Bhutan", "CAST": "CAST", "CAT": "hora de Africa Central", "CCT": "hora del Insulas Cocos", "CDT": "hora estive central", "CHADT": "hora estive de Chatham", "CHAST": "hora normal de Chatham", "CHUT": "hora de Chuuk", "CKT": "hora normal del Insulas Cook", "CKT DST": "hora estive del Insulas Cook", "CLST": "hora estive de Chile", "CLT": "hora normal de Chile", "COST": "hora estive de Colombia", "COT": "hora normal de Colombia", "CST": "hora normal central", "CST China": "hora normal de China", "CST China DST": "hora estive de China", "CVST": "hora estive de Capo Verde", "CVT": "hora normal de Capo Verde", "CXT": "hora del Insula de Natal", "ChST": "hora normal de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "hora estive de Cuba", "CuST": "hora normal de Cuba", "DAVT": "hora de Davis", "DDUT": "hora de Dumont d’Urville", "EASST": "hora estive del Insula de Pascha", "EAST": "hora normal del Insula de Pascha", "EAT": "hora de Africa del Est", "ECT": "hora de Ecuador", "EDT": "hora estive del est", "EGDT": "hora estive de Groenlandia oriental", "EGST": "hora normal de Groenlandia oriental", "EST": "hora normal del est", "FEET": "hora de Europa ultra-oriental", "FJT": "hora normal de Fiji", "FJT Summer": "hora estive de Fiji", "FKST": "hora estive del Insulas Falkland", "FKT": "hora normal del Insulas Falkland", "FNST": "hora estive de Fernando de Noronha", "FNT": "hora normal de Fernando de Noronha", "GALT": "hora del Galápagos", "GAMT": "hora de Gambier", "GEST": "hora estive de Georgia", "GET": "hora normal de Georgia", "GFT": "hora de Guiana Francese", "GIT": "hora del Insulas Gilbert", "GMT": "hora medie de Greenwich", "GNSST": "hora estive de Groenlandia", "GNST": "hora normal de Groenlandia", "GST": "hora normal del Golfo", "GST Guam": "hora normal de Guam", "GYT": "hora de Guyana", "HADT": "hora normal de Hawaii-Aleutianas", "HAST": "hora normal de Hawaii-Aleutianas", "HKST": "hora estive de Hongkong", "HKT": "hora normal de Hongkong", "HOVST": "hora estive de Khovd", "HOVT": "hora normal de Khovd", "ICT": "hora de Indochina", "IDT": "hora estive de Israel", "IOT": "hora del Oceano Indian", "IRKST": "hora estive de Irkutsk", "IRKT": "hora normal de Irkutsk", "IRST": "hora normal de Iran", "IRST DST": "hora estive de Iran", "IST": "hora normal de India", "IST Israel": "hora normal de Israel", "JDT": "hora estive de Japon", "JST": "hora normal de Japon", "KOST": "hora de Kosrae", "KRAST": "hora estive de Krasnoyarsk", "KRAT": "hora normal de Krasnoyarsk", "KST": "hora normal de Corea", "KST DST": "hora estive de Corea", "LHDT": "hora estive de Lord Howe", "LHST": "hora normal de Lord Howe", "LINT": "hora del Insulas del Linea", "MAGST": "hora estive de Magadan", "MAGT": "hora normal de Magadan", "MART": "hora de Marquesas", "MAWT": "hora de Mawson", "MDT": "MDT", "MESZ": "hora estive de Europa central", "MEZ": "hora normal de Europa central", "MHT": "hora del Insulas Marshall", "MMT": "hora de Myanmar", "MSD": "hora estive de Moscova", "MST": "MST", "MUST": "hora estive de Mauritio", "MUT": "hora normal de Mauritio", "MVT": "hora del Maldivas", "MYT": "hora de Malaysia", "NCT": "hora normal de Nove Caledonia", "NDT": "hora estive de Terranova", "NDT New Caledonia": "hora estive de Nove Caledonia", "NFDT": "hora estive del Insula Norfolk", "NFT": "hora normal del Insula Norfolk", "NOVST": "hora estive de Novosibirsk", "NOVT": "hora normal de Novosibirsk", "NPT": "hora de Nepal", "NRT": "hora de Nauru", "NST": "hora normal de Terranova", "NUT": "hora de Niue", "NZDT": "hora estive de Nove Zelanda", "NZST": "hora normal de Nove Zelanda", "OESZ": "hora estive de Europa oriental", "OEZ": "hora normal de Europa oriental", "OMSST": "hora estive de Omsk", "OMST": "hora normal de Omsk", "PDT": "hora estive pacific", "PDTM": "hora estive del Pacifico mexican", "PETDT": "PETDT", "PETST": "PETST", "PGT": "hora de Papua Nove Guinea", "PHOT": "hora del Insulas Phenice", "PKT": "hora normal de Pakistan", "PKT DST": "hora estive de Pakistan", "PMDT": "hora estive de Saint-Pierre e Miquelon", "PMST": "hora normal de Saint-Pierre e Miquelon", "PONT": "hora de Pohnpei", "PST": "hora normal pacific", "PST Philippine": "hora normal del Philippinas", "PST Philippine DST": "hora estive del Philippinas", "PST Pitcairn": "hora de Pitcairn", "PSTM": "hora normal del Pacifico mexican", "PWT": "hora de Palau", "PYST": "hora estive de Paraguay", "PYT": "hora normal de Paraguay", "PYT Korea": "hora de Corea del Nord", "RET": "hora de Réunion", "ROTT": "hora de Rothera", "SAKST": "hora estive de Sachalin", "SAKT": "hora normal de Sachalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "hora normal de Africa del Sud", "SBT": "hora del Insulas Solomon", "SCT": "hora del Seychelles", "SGT": "hora normal de Singapore", "SLST": "SLST", "SRT": "hora de Suriname", "SST Samoa": "hora normal de Samoa american", "SST Samoa Apia": "hora normal de Samoa", "SST Samoa Apia DST": "hora estive de Samoa", "SST Samoa DST": "hora estive de Samoa american", "SYOT": "hora de Syowa", "TAAF": "hora francese meridional e antarctic", "TAHT": "hora de Tahiti", "TJT": "hora de Tajikistan", "TKT": "hora de Tokelau", "TLT": "hora de Timor Oriental", "TMST": "hora estive de Turkmenistan", "TMT": "hora normal de Turkmenistan", "TOST": "hora estive de Tonga", "TOT": "hora normal de Tonga", "TVT": "hora de Tuvalu", "TWT": "hora normal de Taiwan", "TWT DST": "hora estive de Taiwan", "ULAST": "hora estive de Ulan Bator", "ULAT": "hora normal de Ulan Bator", "UYST": "hora estive de Uruguay", "UYT": "hora normal de Uruguay", "UZT": "hora normal de Uzbekistan", "UZT DST": "hora estive de Uzbekistan", "VET": "hora de Venezuela", "VLAST": "hora estive de Vladivostok", "VLAT": "hora normal de Vladivostok", "VOLST": "hora estive de Volgograd", "VOLT": "hora normal de Volgograd", "VOST": "hora de Vostok", "VUT": "hora normal de Vanuatu", "VUT DST": "hora estive de Vanuatu", "WAKT": "hora del Insula Wake", "WARST": "hora estive de Argentina occidental", "WART": "hora normal de Argentina occidental", "WAST": "hora de Africa del West", "WAT": "hora de Africa del West", "WESZ": "hora estive de Europa occidental", "WEZ": "hora normal de Europa occidental", "WFT": "hora de Wallis e Futuna", "WGST": "hora estive de Groenlandia occidental", "WGT": "hora normal de Groenlandia occidental", "WIB": "hora de Indonesia del West", "WIT": "hora de Indonesia del Est", "WITA": "hora de Indonesia Central", "YAKST": "hora estive de Yakutsk", "YAKT": "hora normal de Yakutsk", "YEKST": "hora estive de Ekaterinburg", "YEKT": "hora normal de Ekaterinburg", "YST": "hora de Yukon", "МСК": "hora normal de Moscova", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "hora de Kazakhstan del West", "شىعىش قازاق ەلى": "hora de Kazakhstan del Est", "قازاق ەلى": "hora de Kazakhstan", "قىرعىزستان": "hora de Kirghizistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "hora estive de Peru"},
	}
}

// Locale returns the current translators string locale
func (ia *ia_001) Locale() string {
	return ia.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ia_001'
func (ia *ia_001) PluralsCardinal() []locales.PluralRule {
	return ia.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ia_001'
func (ia *ia_001) PluralsOrdinal() []locales.PluralRule {
	return ia.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ia_001'
func (ia *ia_001) PluralsRange() []locales.PluralRule {
	return ia.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ia_001'
func (ia *ia_001) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ia_001'
func (ia *ia_001) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ia_001'
func (ia *ia_001) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ia *ia_001) MonthAbbreviated(month time.Month) string {
	return ia.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ia *ia_001) MonthsAbbreviated() []string {
	return ia.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ia *ia_001) MonthNarrow(month time.Month) string {
	return ia.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ia *ia_001) MonthsNarrow() []string {
	return ia.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ia *ia_001) MonthWide(month time.Month) string {
	return ia.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ia *ia_001) MonthsWide() []string {
	return ia.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ia *ia_001) WeekdayAbbreviated(weekday time.Weekday) string {
	return ia.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ia *ia_001) WeekdaysAbbreviated() []string {
	return ia.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ia *ia_001) WeekdayNarrow(weekday time.Weekday) string {
	return ia.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ia *ia_001) WeekdaysNarrow() []string {
	return ia.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ia *ia_001) WeekdayShort(weekday time.Weekday) string {
	return ia.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ia *ia_001) WeekdaysShort() []string {
	return ia.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ia *ia_001) WeekdayWide(weekday time.Weekday) string {
	return ia.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ia *ia_001) WeekdaysWide() []string {
	return ia.daysWide
}

// Decimal returns the decimal point of number
func (ia *ia_001) Decimal() string {
	return ia.decimal
}

// Group returns the group of number
func (ia *ia_001) Group() string {
	return ia.group
}

// Group returns the minus sign of number
func (ia *ia_001) Minus() string {
	return ia.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ia_001' and handles both Whole and Real numbers based on 'v'
func (ia *ia_001) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ia.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ia.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ia_001' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ia *ia_001) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ia.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ia.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ia_001'
func (ia *ia_001) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ia.currencies[currency]
	l := len(s) + len(symbol) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ia.group[0])
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

	for j := len(ia.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ia.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ia.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ia.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ia.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ia_001'
// in accounting notation.
func (ia *ia_001) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ia.currencies[currency]
	l := len(s) + len(symbol) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ia.group[0])
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

		for j := len(ia.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ia.currencyNegativePrefix[j])
		}

		b = append(b, ia.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ia.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ia.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ia.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ia.currencyNegativeSuffix...)
	} else {
		b = append(b, ia.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ia.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, ia.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ia.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0x6c, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, ia.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ia.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
