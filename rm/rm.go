package rm

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type rm struct {
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
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'rm' locale
func New() locales.Translator {
	return &rm{
		locale:                 "rm",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  "\u202f",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "schan.", "favr.", "mars", "avr.", "matg", "zercl.", "fan.", "avust", "sett.", "oct.", "nov.", "dec."},
		monthsNarrow:           []string{"", "S", "F", "M", "A", "M", "Z", "F", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "da schaner", "da favrer", "da mars", "d’avrigl", "da matg", "da zercladur", "da fanadur", "d’avust", "da settember", "d’october", "da november", "da december"},
		daysAbbreviated:        []string{"du", "gli", "ma", "me", "gie", "ve", "so"},
		daysNarrow:             []string{"D", "G", "M", "M", "G", "V", "S"},
		daysWide:               []string{"dumengia", "glindesdi", "mardi", "mesemna", "gievgia", "venderdi", "sonda"},
		timezones:              map[string]string{"ACDT": "temp da stad dal center da l’Australia", "ACST": "ACST", "ACT": "ACT", "ACWDT": "temp da stad dal center-vest da l’Australia", "ACWST": "temp normal dal center-vest da l’Australia", "ADT": "temp da stad da l’Atlantic USA", "ADT Arabia": "temp da stad arab", "AEDT": "temp da stad da l’ost da l’Australia", "AEST": "temp normal da l’ost da l’Australia", "AFT": "temp da l’Afganistan", "AKDT": "temp da stad da l’Alasca", "AKST": "temp normal da l’Alasca", "AMST": "temp da stad da l’Amazonas", "AMST Armenia": "temp da stad da l’Armenia", "AMT": "temp normal da l’Amazonas", "AMT Armenia": "temp normal da l’Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "temp da stad da l’Argentina", "ART": "temp normal da l’Argentina", "AST": "temp normal da l’Atlantic USA", "AST Arabia": "temp normal arab", "AWDT": "temp da stad dal vest da l’Australia", "AWST": "temp normal dal vest da l’Australia", "AZST": "temp da stad da l’Aserbaidschan", "AZT": "temp normal da l’Aserbaidschan", "BDT Bangladesh": "temp da stad dal Bangladesch", "BNT": "temp dal Brunei", "BOT": "temp da la Bolivia", "BRST": "temp da stad da la Brasilia", "BRT": "temp normal da la Brasilia", "BST Bangladesh": "temp normal dal Bangladesch", "BT": "temp dal Butan", "CAST": "CAST", "CAT": "temp da l’Africa Centrala", "CCT": "temp da las Inslas Cocos", "CDT": "temp da stad dal center USA", "CHADT": "temp da stad da las Inslas Chatham", "CHAST": "temp normal da las Inslas Chatham", "CHUT": "temp da Chuuk", "CKT": "temp normal da las Inslas Cook", "CKT DST": "temp da stad da las Inslas Cook", "CLST": "temp da stad dal Chile", "CLT": "temp normal dal Chile", "COST": "temp da stad da la Columbia", "COT": "temp normal da la Columbia", "CST": "temp normal dal center USA", "CST China": "temp normal da la China", "CST China DST": "temp da stad da la China", "CVST": "temp da stad dal Cap Verd", "CVT": "temp normal dal Cap Verd", "CXT": "temp da l’Insla da Nadal", "ChST": "temp dals Chamorro", "ChST NMI": "ChST NMI", "CuDT": "temp da stad da la Cuba", "CuST": "temp normal da la Cuba", "DAVT": "temp da Davis", "DDUT": "temp da Dumont d’Urville", "EASST": "temp da stad da l’Insla da Pasca", "EAST": "temp normal da l’Insla da Pasca", "EAT": "temp da l’Africa Orientala", "ECT": "temp da l’Ecuador", "EDT": "temp da stad da l’ost USA", "EGDT": "temp da stad da la Grönlanda orientala", "EGST": "temp normal da la Grönlanda orientala", "EST": "temp normal da l’ost USA", "FEET": "temp da l’extrem orient da l’Europa", "FJT": "temp normal dal Fidschi", "FJT Summer": "temp da stad dal Fidschi", "FKST": "temp da stad da las Inslas Falkland", "FKT": "temp normal da las Inslas Falkland", "FNST": "temp da stad da Fernando de Noronha", "FNT": "temp normal da Fernando de Noronha", "GALT": "temp da las Inslas Galápagos", "GAMT": "temp da las Inslas Gambier", "GEST": "temp da stad da la Georgia", "GET": "temp normal da la Georgia", "GFT": "temp da la Guyana Franzosa", "GIT": "temp da las Inslas Gilbert", "GMT": "temp dal meridian da Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "temp normal dal Golf", "GST Guam": "GST Guam", "GYT": "temp da la Guyana", "HADT": "temp normal dal Hawai e las Aleutinas", "HAST": "temp normal dal Hawai e las Aleutinas", "HKST": "temp da stad dal Hongkong", "HKT": "temp normal dal Hongkong", "HOVST": "temp da stad da Hovd", "HOVT": "temp normal da Hovd", "ICT": "temp da l’Indochina", "IDT": "temp da stad da l’Israel", "IOT": "temp da l’Ocean Indic", "IRKST": "temp da stad dad Irkutsk", "IRKT": "temp normal dad Irkutsk", "IRST": "temp normal da l’Iran", "IRST DST": "temp da stad da l’Iran", "IST": "temp normal da l’India", "IST Israel": "temp normal da l’Israel", "JDT": "temp da stad dal Giapun", "JST": "temp normal dal Giapun", "KOST": "temp da Kosrae", "KRAST": "temp da stad da Krasnojarsk", "KRAT": "temp normal da Krasnojarsk", "KST": "temp normal da la Corea", "KST DST": "temp da stad da la Corea", "LHDT": "temp da stad da Lord Howe", "LHST": "temp normal da Lord Howe", "LINT": "temp da las Inslas da la Lingia", "MAGST": "temp da stad da Magadan", "MAGT": "temp normal da Magadan", "MART": "temp da las Inslas Marquesas", "MAWT": "temp da Mawson", "MDT": "MDT", "MESZ": "temp da stad da l’Europa Centrala", "MEZ": "temp normal da l’Europa Centrala", "MHT": "temp da las Inslas da Marshall", "MMT": "temp dal Myanmar", "MSD": "temp da stad da Moscau", "MST": "MST", "MUST": "temp da stad dal Mauritius", "MUT": "temp normal dal Mauritius", "MVT": "temp da las Maledivas", "MYT": "temp da la Malaisia", "NCT": "temp normal da la Nova Caledonia", "NDT": "temp da stad da la Terranova", "NDT New Caledonia": "temp da stad da la Nova Caledonia", "NFDT": "temp da stad da l’Insla Norfolk", "NFT": "temp normal da l’Insla Norfolk", "NOVST": "temp da stad da Novosibirsk", "NOVT": "temp normal da Novosibirsk", "NPT": "temp dal Nepal", "NRT": "temp da Nauru", "NST": "temp normal da la Terranova", "NUT": "temp da Niue", "NZDT": "temp da stad da la Nova Zelanda", "NZST": "temp normal da la Nova Zelanda", "OESZ": "temp da stad da l’Europa Orientala", "OEZ": "temp normal da l’Europa Orientala", "OMSST": "temp da stad dad Omsk", "OMST": "temp normal dad Omsk", "PDT": "temp da stad dal Pacific USA", "PDTM": "temp da stad mexican dal Pacific", "PETDT": "PETDT", "PETST": "PETST", "PGT": "temp da la Papua Nova Guinea", "PHOT": "temp da las Inslas Fenix", "PKT": "temp normal dal Pakistan", "PKT DST": "temp da stad dal Pakistan", "PMDT": "temp da stad da Saint Pierre e Miquelon", "PMST": "temp normal da Saint Pierre e Miquelon", "PONT": "temp da Pohnpei", "PST": "temp normal dal Pacific USA", "PST Philippine": "temp normal da las Filippinas", "PST Philippine DST": "temp da stad da las Filippinas", "PST Pitcairn": "temp da las Inslas Pitcairn", "PSTM": "temp normal mexican dal Pacific", "PWT": "temp da Palau", "PYST": "temp da stad dal Paraguai", "PYT": "temp normal dal Paraguai", "PYT Korea": "temp da la Corea dal Nord", "RET": "temp da la Réunion", "ROTT": "temp da Rothera", "SAKST": "temp da stad da Sachalin", "SAKT": "temp normal da Sachalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "temp normal da l’Africa Meridiunala", "SBT": "temp da las Inslas da Salomon", "SCT": "temp da las Seychellas", "SGT": "temp normal dal Singapur", "SLST": "SLST", "SRT": "temp dal Surinam", "SST Samoa": "temp normal da la Samoa Americana", "SST Samoa Apia": "temp normal da la Samoa", "SST Samoa Apia DST": "temp da stad da la Samoa", "SST Samoa DST": "temp da stad da la Samoa Americana", "SYOT": "temp da Shōwa", "TAAF": "temp dals Territoris Franzos Meridiunals ed Antarctics", "TAHT": "temp da Tahiti", "TJT": "temp dal Tadschikistan", "TKT": "temp da Tokelau", "TLT": "temp dal Timor da l’Ost", "TMST": "temp da stad dal Turkmenistan", "TMT": "temp normal dal Turkmenistan", "TOST": "temp da stad da Tonga", "TOT": "temp normal da Tonga", "TVT": "temp da las Inslas da Tuvalu", "TWT": "temp normal dal Taiwan", "TWT DST": "temp da stad dal Taiwan", "ULAST": "temp da stad dad Ulaanbaatar", "ULAT": "temp normal dad Ulaanbaatar", "UYST": "temp da stad da l’Uruguai", "UYT": "temp normal da l’Uruguai", "UZT": "temp normal da l’Usbekistan", "UZT DST": "temp da stad da l’Usbekistan", "VET": "temp da la Venezuela", "VLAST": "temp da stad da Vladivostok", "VLAT": "temp normal da Vladivostok", "VOLST": "temp da stad da Volgograd", "VOLT": "temp normal da Volgograd", "VOST": "temp da Vostok", "VUT": "temp normal dal Vanuatu", "VUT DST": "temp da stad dal Vanuatu", "WAKT": "temp da l’Insla Wake", "WARST": "temp da stad da l’Argentina occidentala", "WART": "temp normal da l’Argentina occidentala", "WAST": "temp da l’Africa Occidentala", "WAT": "temp da l’Africa Occidentala", "WESZ": "temp da stad da l’Europa dal Vest", "WEZ": "temp normal da l’Europa dal Vest", "WFT": "temp da Wallis e Futuna", "WGST": "temp da stad da la Grönlanda occidentala", "WGT": "temp normal da la Grönlanda occidentala", "WIB": "temp dal vest da l’Indonesia", "WIT": "temp da l’ost da l’Indonesia", "WITA": "temp dal center da l’Indonesia", "YAKST": "temp da stad da Jakutsk", "YAKT": "temp normal da Jakutsk", "YEKST": "temp da stad da Jekaterinburg", "YEKT": "temp normal da Jekaterinburg", "YST": "temp dal Yukon", "МСК": "temp normal da Moscau", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "temp dal Kasachstan occidental", "شىعىش قازاق ەلى": "temp dal Kasachstan oriental", "قازاق ەلى": "temp dal Kasachstan", "قىرعىزستان": "temp dal Kirghistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "temp da stad dal Peru"},
	}
}

// Locale returns the current translators string locale
func (rm *rm) Locale() string {
	return rm.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'rm'
func (rm *rm) PluralsCardinal() []locales.PluralRule {
	return rm.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'rm'
func (rm *rm) PluralsOrdinal() []locales.PluralRule {
	return rm.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'rm'
func (rm *rm) PluralsRange() []locales.PluralRule {
	return rm.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'rm'
func (rm *rm) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'rm'
func (rm *rm) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'rm'
func (rm *rm) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (rm *rm) MonthAbbreviated(month time.Month) string {
	return rm.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (rm *rm) MonthsAbbreviated() []string {
	return rm.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (rm *rm) MonthNarrow(month time.Month) string {
	return rm.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (rm *rm) MonthsNarrow() []string {
	return rm.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (rm *rm) MonthWide(month time.Month) string {
	return rm.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (rm *rm) MonthsWide() []string {
	return rm.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (rm *rm) WeekdayAbbreviated(weekday time.Weekday) string {
	return rm.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (rm *rm) WeekdaysAbbreviated() []string {
	return rm.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (rm *rm) WeekdayNarrow(weekday time.Weekday) string {
	return rm.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (rm *rm) WeekdaysNarrow() []string {
	return rm.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (rm *rm) WeekdayShort(weekday time.Weekday) string {
	return rm.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (rm *rm) WeekdaysShort() []string {
	return rm.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (rm *rm) WeekdayWide(weekday time.Weekday) string {
	return rm.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (rm *rm) WeekdaysWide() []string {
	return rm.daysWide
}

// Decimal returns the decimal point of number
func (rm *rm) Decimal() string {
	return rm.decimal
}

// Group returns the group of number
func (rm *rm) Group() string {
	return rm.group
}

// Group returns the minus sign of number
func (rm *rm) Minus() string {
	return rm.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'rm' and handles both Whole and Real numbers based on 'v'
func (rm *rm) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(rm.group) - 1; j >= 0; j-- {
					b = append(b, rm.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, rm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'rm' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (rm *rm) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rm.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, rm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, rm.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'rm'
func (rm *rm) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := rm.currencies[currency]
	l := len(s) + len(symbol) + 2 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(rm.group) - 1; j >= 0; j-- {
					b = append(b, rm.group[j])
				}
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
		b = append(b, rm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, rm.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'rm'
// in accounting notation.
func (rm *rm) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := rm.currencies[currency]
	l := len(s) + len(symbol) + 4 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(rm.group) - 1; j >= 0; j-- {
					b = append(b, rm.group[j])
				}
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

		b = append(b, rm.currencyNegativePrefix[0])

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
			b = append(b, rm.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, rm.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'rm'
func (rm *rm) FmtDateShort(t time.Time) string {
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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'rm'
func (rm *rm) FmtDateMedium(t time.Time) string {
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

// FmtDateLong returns the long date representation of 't' for 'rm'
func (rm *rm) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, rm.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'rm'
func (rm *rm) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, rm.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20, 0x69, 0x6c, 0x73}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, rm.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'rm'
func (rm *rm) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'rm'
func (rm *rm) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'rm'
func (rm *rm) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'rm'
func (rm *rm) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, rm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := rm.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
