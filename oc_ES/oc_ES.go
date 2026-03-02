package oc_ES

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type oc_ES struct {
	locale            string
	pluralsCardinal   []locales.PluralRule
	pluralsOrdinal    []locales.PluralRule
	pluralsRange      []locales.PluralRule
	decimal           string
	group             string
	minus             string
	percent           string
	percentSuffix     string
	timeSeparator     string
	currencies        []string // idx = enum of currency code
	monthsAbbreviated []string
	monthsNarrow      []string
	monthsWide        []string
	daysAbbreviated   []string
	daysNarrow        []string
	daysShort         []string
	daysWide          []string
	timezones         map[string]string
}

// New returns a new instance of translator for the 'oc_ES' locale
func New() locales.Translator {
	return &oc_ES{
		locale:            "oc_ES",
		pluralsCardinal:   nil,
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           ".",
		group:             ",",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "CF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SDG", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "TL", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:     "\u202f",
		monthsAbbreviated: []string{"", "gèr", "her", "mar", "abr", "mai", "jun", "jur", "ago", "set", "oct", "nov", "dec"},
		monthsNarrow:      []string{"", "G", "H", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:        []string{"", "gèr", "hereuèr", "març", "abriu", "mai", "junh", "juriòl", "agost", "seteme", "octobre", "noveme", "deseme"},
		daysAbbreviated:   []string{"dim", "del", "dma", "dmè", "dij", "diu", "dis"},
		daysNarrow:        []string{"D", "L", "M", "X", "J", "U", "S"},
		daysShort:         []string{"di", "de", "da", "dm", "dj", "du", "ds"},
		daysWide:          []string{"dimenge", "deluns", "dimars", "dimèrcles", "dijaus", "diuendres", "dissabte"},
		timezones:         map[string]string{"ACDT": "ora d’estiu d’Austràlia centrau", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ora d’estiu d’Austràlia centreoccidentau", "ACWST": "ora estandarda d’Austràlia centreoccidentau", "ADT": "ora d’estiu der Atlantic", "ADT Arabia": "ora d’estiu d’arabia", "AEDT": "ora d’estiu d’Austràlia orientau", "AEST": "ora estandard d’Austràlia orientau", "AFT": "ora d’Afganistan", "AKDT": "ora d’estiu d’Alaska", "AKST": "ora estandard d’Alaska", "AMST": "ora d’estiu der Amazònes", "AMST Armenia": "ora d’estiu d’Armenia", "AMT": "ora estandard der Amazònes", "AMT Armenia": "ora estandard d’Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ora d’estiu d’Argentina", "ART": "ora estandard d’Argentina", "AST": "ora estandard der Atlantic", "AST Arabia": "ora estandard d’Arabia", "AWDT": "ora d’estiu estandard d’Australia occidentau", "AWST": "ora estandard d’Australia occidentau", "AZST": "ora d’estiu d’Azerbaijan", "AZT": "ora estandard d’Azerbaijan", "BDT Bangladesh": "ora d’estiu de Bangladesh", "BNT": "ora de Brunei", "BOT": "ora de Bolivia", "BRST": "ora d’estiu de Brasilia", "BRT": "ora estandard de Brasilia", "BST Bangladesh": "ora estandard de Bangladesh", "BT": "ora de Butan", "CAST": "CAST", "CAT": "ora d’Africa central", "CCT": "ora dera isla de Cocos", "CDT": "ora d’estiu centrau", "CHADT": "ora d’estiu de Chatham", "CHAST": "ora estandard de Chatham", "CHUT": "ora de Chuuk", "CKT": "ora estandard des Isles Cook", "CKT DST": "ora d’estiu mieja des Isles Cook", "CLST": "ora d’estiu de Chile", "CLT": "ora estandard de Chile", "COST": "ora d’estiu de Colòmbia", "COT": "ora estandard de Colòmbia", "CST": "ora estandard centrau", "CST China": "ora estandard de China", "CST China DST": "ora d’estiu de China", "CVST": "ora d’estiu de Cap-Verd", "CVT": "ora estandard de Cap-Verd", "CXT": "ora dera isla de Nadau", "ChST": "ora estandard de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "ora d’estiu de Cuba", "CuST": "ora estandard de Cuba", "DAVT": "ora de Davis", "DDUT": "ora de Dumont-d’Urville", "EASST": "ora d’estiu dera isla de Pasqua", "EAST": "ora estandard dera isla de Pasqua", "EAT": "ora d’Africa orientau", "ECT": "ora d’Equator", "EDT": "ora d’estiu orientau", "EGDT": "ora d’estiu de Groenlandia orientau", "EGST": "ora estandard de Groenlandia orientau", "EST": "ora estandard orientau", "FEET": "ora der extrem d’Euròpa orientau", "FJT": "ora estandard de Fiji", "FJT Summer": "ora d’estiu de Fiji", "FKST": "ora d’estiu des isles Maldives", "FKT": "ora estandard des isles Maldives", "FNST": "ora d’estiu de Fernando de Noronha", "FNT": "ora estandard de Fernando de Noronha", "GALT": "ora de Galápagos", "GAMT": "ora de Gambier", "GEST": "ora d’estiu de Geòrgia", "GET": "ora estandard de Geòrgia", "GFT": "ora dera Guayana Francesa", "GIT": "ora des Isles Gilbert", "GMT": "ora deth meridian de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "ora estandard deth Gòlf", "GST Guam": "GST Guam", "GYT": "ora de Guyana", "HADT": "ora estandard de Hawai-Aleutianes", "HAST": "ora estandard de Hawai-Aleutianes", "HKST": "ora d’estiu de Hong Kong", "HKT": "ora estandard de Hong Kong", "HOVST": "ora d’estiu de Hovd", "HOVT": "ora estandard de Hovd", "ICT": "ora d’Indochina", "IDT": "ora d’estiu d’Israèl", "IOT": "ora der Ocean Indic", "IRKST": "ora d’estiu d’Irkutsk", "IRKT": "ora estandard d’Irkutsk", "IRST": "ora estandard d’Iran", "IRST DST": "ora d’estiu d’Iran", "IST": "ora estandard dera India", "IST Israel": "ora estàndard d’Israèl", "JDT": "ora d’estiu de Japon", "JST": "ora estandard de Japon", "KOST": "ora de Kosrae", "KRAST": "ora d’estiu de Krasnoyarsk", "KRAT": "ora estandard de Krasnoyarsk", "KST": "ora estandard de Corèa", "KST DST": "ora d’estiu de Corèa", "LHDT": "ora d’estiu de Lord Howe", "LHST": "ora estandard de Lord Howe", "LINT": "ora des Espòrades Equatorials", "MAGST": "ora d’estiu de Magadan", "MAGT": "ora estandard de Magadan", "MART": "ora des Isles Marqueses", "MAWT": "ora de Mawson", "MDT": "ora d’estiu des Montanhes Rocoses", "MESZ": "ora d’estiu d’Euròpa centrau", "MEZ": "ora estandard d’Euròpa centrau", "MHT": "ora des Isles Marshall", "MMT": "ora de Myanmar", "MSD": "ora d’estiu de Moscòu", "MST": "ora estandard des Montanhes Rocoses", "MUST": "ora d’estiu de Maurici", "MUT": "ora estandard de Maurici", "MVT": "ora des Malvines", "MYT": "ora de Malàisia", "NCT": "ora estandard de Nòva Caledònia", "NDT": "ora d’estiu de Terra-Nòva", "NDT New Caledonia": "ora d’estiu de Nòva Caledònia", "NFDT": "ora d’estiu dera Isla Norfolk", "NFT": "ora estandard dera Isla Norfolk", "NOVST": "ora d’estiu de Novosibirsk", "NOVT": "ora estandard de Novosibirsk", "NPT": "ora deth Nepal", "NRT": "ora de Nauru", "NST": "ora estandard de Terra-Nòva", "NUT": "ora de Niue", "NZDT": "ora d’estiu de Nòva Zelanda", "NZST": "ora estandard de Nòva Zelanda", "OESZ": "ora d’estiu d’Euròpa de l’èst", "OEZ": "ora estandard d’Euròpa de l’èst", "OMSST": "ora d’estiu d’Omsk", "OMST": "ora estandard d’Omsk", "PDT": "ora d’estiu deth Pacific", "PDTM": "ora d’estiu deth Pacífic de Mexic", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ora de Papua Nòva Guinèa", "PHOT": "ora des Isles Fénix", "PKT": "ora estandard deth Pakistan", "PKT DST": "ora d’estiu deth Pakistan", "PMDT": "ora d’estiu de St. Pierre e Miquelon", "PMST": "ora estandard de St. Pierre e Miquelon", "PONT": "ora de Pohnpei", "PST": "ora estandard deth Pacific", "PST Philippine": "ora estandard de Filipines", "PST Philippine DST": "ora d’estiu de Filipines", "PST Pitcairn": "ora de Pitcairn", "PSTM": "ora estandard deth Pacífic de Mexic", "PWT": "ora de Palaos", "PYST": "ora d’estiu de Paraguay", "PYT": "ora estandard de Paraguay", "PYT Korea": "ora de Pyongyang", "RET": "ora de Reünion", "ROTT": "ora de Rothera", "SAKST": "ora d’estiu de Sajalin", "SAKT": "ora estandard de Sajalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ora de Sudafrica", "SBT": "ora des Isles Salomon", "SCT": "ora de Seychelles", "SGT": "ora de Singapor", "SLST": "SLST", "SRT": "ora de Surinam", "SST Samoa": "ora estandard de Samoa", "SST Samoa Apia": "ora estandard d’Apia", "SST Samoa Apia DST": "ora d’esiu d’Apia", "SST Samoa DST": "ora d’estiu de Samoa", "SYOT": "ora de Syowa", "TAAF": "hora d’Antartida e Territoris Australs Francesi", "TAHT": "ora de Tahiti", "TJT": "ora de Tajikistan", "TKT": "ora de Tokelau", "TLT": "ora de Timòr orientau", "TMST": "ora d’estiu de Turkmenistan", "TMT": "ora estandard de Turkmenistan", "TOST": "ora d’estiu de Tonga", "TOT": "ora estandard de Tonga", "TVT": "ora de Tuvalu", "TWT": "ora estandard de Taipei", "TWT DST": "ora d’estiu de Taipei", "ULAST": "ora d’estiu de Ulan Bator", "ULAT": "ora estandard de Ulan Bator", "UYST": "ora d’estiu d’Uruguay", "UYT": "ora estandard d’Uruguay", "UZT": "ora estandard de Uzbekistan", "UZT DST": "ora d’estiu de Uzbekistan", "VET": "ora de Veneçuela", "VLAST": "ora d’estiu de Vladivostok", "VLAT": "ora estandard de Vladivostok", "VOLST": "ora d’estiu de Volgograd", "VOLT": "ora estandard de Volgograd", "VOST": "ora de Vostok", "VUT": "ora estandard de Vanuatu", "VUT DST": "ora d’estiu de Vanuatu", "WAKT": "ora dera Isla Wake", "WARST": "ora d’estiu d’Argentina occidentau", "WART": "ora estandard d’Argentina occidentau", "WAST": "ora d’Africa occidentau", "WAT": "ora d’Africa occidentau", "WESZ": "ora d’estiu d’Euròpa occidentau", "WEZ": "ora estandard d’Euròpa occidentau", "WFT": "ora de Wallis e Futuna", "WGST": "ora d’estiu de Groenlandia occidentau", "WGT": "ora estandard de Groenlandia occidentau", "WIB": "hora d’Indonesia occidentau", "WIT": "hora d’Indonesia orientau", "WITA": "hora d’Indonesia centrau", "YAKST": "ora d’estiu de Yakutsk", "YAKT": "ora estandard de Yakutsk", "YEKST": "ora d’estiu d’Ekaterimburg", "YEKT": "ora estandard d’Ekaterimburg", "YST": "ora de Yukon", "МСК": "ora estandard de Moscòu", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ora de Kazajistan occidentau", "شىعىش قازاق ەلى": "ora de Kazajistan orientau", "قازاق ەلى": "ora de Kazajistan", "قىرعىزستان": "ora de Kirguistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ora d’estiu des Açòres"},
	}
}

// Locale returns the current translators string locale
func (oc *oc_ES) Locale() string {
	return oc.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'oc_ES'
func (oc *oc_ES) PluralsCardinal() []locales.PluralRule {
	return oc.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'oc_ES'
func (oc *oc_ES) PluralsOrdinal() []locales.PluralRule {
	return oc.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'oc_ES'
func (oc *oc_ES) PluralsRange() []locales.PluralRule {
	return oc.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'oc_ES'
func (oc *oc_ES) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'oc_ES'
func (oc *oc_ES) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'oc_ES'
func (oc *oc_ES) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (oc *oc_ES) MonthAbbreviated(month time.Month) string {
	return oc.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (oc *oc_ES) MonthsAbbreviated() []string {
	return oc.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (oc *oc_ES) MonthNarrow(month time.Month) string {
	return oc.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (oc *oc_ES) MonthsNarrow() []string {
	return oc.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (oc *oc_ES) MonthWide(month time.Month) string {
	return oc.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (oc *oc_ES) MonthsWide() []string {
	return oc.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (oc *oc_ES) WeekdayAbbreviated(weekday time.Weekday) string {
	return oc.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (oc *oc_ES) WeekdaysAbbreviated() []string {
	return oc.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (oc *oc_ES) WeekdayNarrow(weekday time.Weekday) string {
	return oc.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (oc *oc_ES) WeekdaysNarrow() []string {
	return oc.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (oc *oc_ES) WeekdayShort(weekday time.Weekday) string {
	return oc.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (oc *oc_ES) WeekdaysShort() []string {
	return oc.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (oc *oc_ES) WeekdayWide(weekday time.Weekday) string {
	return oc.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (oc *oc_ES) WeekdaysWide() []string {
	return oc.daysWide
}

// Decimal returns the decimal point of number
func (oc *oc_ES) Decimal() string {
	return oc.decimal
}

// Group returns the group of number
func (oc *oc_ES) Group() string {
	return oc.group
}

// Group returns the minus sign of number
func (oc *oc_ES) Minus() string {
	return oc.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'oc_ES' and handles both Whole and Real numbers based on 'v'
func (oc *oc_ES) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, oc.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'oc_ES' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (oc *oc_ES) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, oc.percentSuffix...)

	b = append(b, oc.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'oc_ES'
func (oc *oc_ES) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := oc.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, oc.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, oc.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'oc_ES'
// in accounting notation.
func (oc *oc_ES) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := oc.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, oc.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, oc.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, oc.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, oc.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, oc.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, oc.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'oc_ES'
func (oc *oc_ES) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := oc.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
