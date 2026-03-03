package qu_PE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type qu_PE struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	percentSuffix      string
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

// New returns a new instance of translator for the 'qu_PE' locale
func New() locales.Translator {
	return &qu_PE{
		locale:             "qu_PE",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBG", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "DBM", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "DBZ", "$CA", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHC", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "S/", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$US", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:      " ",
		monthsAbbreviated:  []string{"", "Ene", "Feb", "Mar", "Abr", "May", "Jun", "Jul", "Ago", "Set", "Oct", "Nov", "Dic"},
		monthsWide:         []string{"", "Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Setiembre", "Octubre", "Noviembre", "Diciembre"},
		daysAbbreviated:    []string{"Dom", "Lun", "Mar", "Mié", "Jue", "Vie", "Sab"},
		daysNarrow:         []string{"D", "L", "M", "X", "J", "V", "S"},
		daysWide:           []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"},
		periodsAbbreviated: []string{"a.m.", "p.m."},
		timezones:          map[string]string{"ACDT": "Hora de Verano de Australia Central", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Hora de Verano de Australia Central Occidental", "ACWST": "Hora Estandar de Australia Central Occidental", "ADT": "Hora De Verano del Atlántico", "ADT Arabia": "Hora de Verano de Arabia", "AEDT": "Hora de Verano de Australia Oriental", "AEST": "Hora Estandar de Australia Oriental", "AFT": "Hora de Afganistán", "AKDT": "Hora de Verano de Alaska", "AKST": "Hora Estandar de Alaska", "AMST": "Hora de Verano de Amazonas", "AMST Armenia": "Hora de Verano de Armenia", "AMT": "Hora Estandar de Amazonas", "AMT Armenia": "Hora Estandar de Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Hora de Verano de Argentina", "ART": "Hora Estandar de Argentina", "AST": "Hora Estandar del Atlántico", "AST Arabia": "Hora Estandar de Arabia", "AWDT": "Hora de Verano de Australia Occidental", "AWST": "Hora Estandar de Australia Occidental", "AZST": "Hora de Verano de Azerbaiyán", "AZT": "Hora Estandar de Azerbaiyán", "BDT Bangladesh": "Hora de Verano de Bangladesh", "BNT": "Hora de Brunei Darussalam", "BOT": "Bolivia Time", "BRST": "Hora de Verano de Brasilia", "BRT": "Hora Estandar de Brasilia", "BST Bangladesh": "Hora Estandar de Bangladesh", "BT": "Hora de Bután", "CAST": "CAST", "CAT": "Hora de Africa Central", "CCT": "Hora de Islas Cocos", "CDT": "Hora Central de Verano", "CHADT": "Hora de Verano de Chatham", "CHAST": "Hora Estandar de Chatham", "CHUT": "Hora de Chuuk", "CKT": "Hora Estandar de Isla Cocos", "CKT DST": "Hora de Verano de Isla Cocos", "CLST": "Hora de Verano de Chile", "CLT": "Hora Estandar de Chile", "COST": "Hora de Verano de Colombia", "COT": "Hora Estandar de Colombia", "CST": "Estandard Hora Central", "CST China": "Hora Estandar de China", "CST China DST": "Hora de Verano de China", "CVST": "Hora de Verano de Cabo Verde", "CVT": "Hora Estandar de Cabo Verde", "CXT": "Hora de Isla Christmas", "ChST": "Hora Estandar de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Hora de Verano de Cuba", "CuST": "Hora Estandar de Cuba", "DAVT": "Hora de Davis", "DDUT": "Hora de Dumont-d’Urville", "EASST": "Hora de Verano de Isla de Pascua", "EAST": "Hora Estandar de Isla de Pascua", "EAT": "Hora de Africa Oriental", "ECT": "Hora de Ecuador", "EDT": "Hora de Verano del Este", "EGDT": "Hora de Verano de Groenlandia", "EGST": "Hora Estandar de Groenlandia", "EST": "Hora Estandar del Este", "FEET": "Hora de Verano más Oriental de Europa", "FJT": "Hora Estandar de Fiji", "FJT Summer": "Hora de Verano de Fiji", "FKST": "Hora de Verano de Islas Malvinas", "FKT": "Hora Estandar de Islas Malvinas", "FNST": "Hora de Verano de Fernando de Noronha", "FNT": "Hora Estandar de Fernando de Noronha", "GALT": "Hora de Islas Galápagos", "GAMT": "Hora de Gambier", "GEST": "Hora de Verano de Georgia", "GET": "Hora Estandar de Georgia", "GFT": "Hora de Guayana Francesa", "GIT": "Hora de Islas Gilbert", "GMT": "Hora del Meridiano de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Hora Estandar del Golfo", "GST Guam": "GST Guam", "GYT": "Hora de Guyana", "HADT": "Hora de Verano de Hawai-Aleutiano", "HAST": "Hora Estandar de Hawai-Aleutiano", "HKST": "Hora de Verano de Hong Kong", "HKT": "Hora Estandar de Hong Kong", "HOVST": "Hora de Verano de Hovd", "HOVT": "Hora Estandar de Hovd", "ICT": "Hora de Indochina", "IDT": "Hora de Verano de Israel", "IOT": "Hora del Oceano Índico", "IRKST": "Hora de Verano de Irkutsk", "IRKT": "Hora Estandar de Irkutsk", "IRST": "Hora Estandar de Irán", "IRST DST": "Hora de Verano de Irán", "IST": "Hora Estandar de India", "IST Israel": "Hora Estandar de Israel", "JDT": "Hora de Verano de Japón", "JST": "Hora Estandar de Japón", "KOST": "Hora de Kosrae", "KRAST": "Hora de Verano de Krasnoyarsk", "KRAT": "Hora Estandar de Krasnoyarsk", "KST": "Hora Estandar de Corea", "KST DST": "Hora de Verano de Corea", "LHDT": "Hora de Verano de Lord Howe", "LHST": "Hora Estandar de Lord Howe", "LINT": "Hora de Espóradas Ecuatoriales", "MAGST": "Hora de Verano de Magadan", "MAGT": "Hora Estandar de Magadan", "MART": "Hora de Marquesas", "MAWT": "Hora de Mawson", "MDT": "Hora de Verano de la Montaña", "MESZ": "Hora de Verano de Europa Central", "MEZ": "Hora Estandar de Europa Central", "MHT": "Hora de Islas Marshall", "MMT": "Hora de Myanmar", "MSD": "Hora de Verano de Moscú", "MST": "Hora Estandar de la Montaña", "MUST": "Hora de Verano de Mauricio", "MUT": "Hora Estandar de Mauricio", "MVT": "Hora de Maldivas", "MYT": "Hora de Malasia", "NCT": "Hora Estandar de Nueva Caledonia", "NDT": "Hora de Verano de Terranova", "NDT New Caledonia": "Hora de Verano de Nueva Caledonia", "NFDT": "Hora de Verano de la Isla Norfolk", "NFT": "Hora Estandar de la Isla Norfolk", "NOVST": "Hora de Verano de Novosibirsk", "NOVT": "Hora Estandar de Novosibirsk", "NPT": "Hora de Nepal", "NRT": "Hora de Nauru", "NST": "Hora Estandar de Terranova", "NUT": "Hora de Niue", "NZDT": "Hora de Verano de Nueva Zelanda", "NZST": "Hora Estandar de Nueva Zelanda", "OESZ": "Hora de Verano de Europa Oriental", "OEZ": "Hora Estandar de Europa Oriental", "OMSST": "Hora de Verano de Omsk", "OMST": "Hora Estandar de Omsk", "PDT": "Hora de Verano del Pacífico", "PDTM": "Hora de Verano del Pacífico Mexicano", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Hora de Papua Nueva Guinea", "PHOT": "Hora de Islas Phoenix", "PKT": "Hora Estandar de Pakistán", "PKT DST": "Hora de Verano de Pakistán", "PMDT": "Hora de Verano de San Pedro y Miquelón", "PMST": "Hora Estandar de San Pedro y Miquelón", "PONT": "Hora de Pohnpei", "PST": "Hora Estandar del Pacífico", "PST Philippine": "Hora Estandar de Filipinas", "PST Philippine DST": "Hora de Verano de Filipinas", "PST Pitcairn": "Hora de Pitcairn", "PSTM": "Hora Estandar del Pacífico Mexicano", "PWT": "Hora de Palau", "PYST": "Hora de Verano de Paraguay", "PYT": "Hora Estandar de Paraguay", "PYT Korea": "Hora de Pionyang", "RET": "Hora de Réunion", "ROTT": "Hora de Rothera", "SAKST": "Hora de Verano de Sakhalin", "SAKT": "Hora Estandar de Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Hora de Sudafrica", "SBT": "Hora de Islas Salomón", "SCT": "Hora de Seychelles", "SGT": "Hora Estandar de Singapur", "SLST": "SLST", "SRT": "Hora de Surinam", "SST Samoa": "Hora Estandar de Samoa", "SST Samoa Apia": "Hora Estandar de Apia", "SST Samoa Apia DST": "Hora de Verano de Apia", "SST Samoa DST": "Hora de Verano de Samoa", "SYOT": "Hora de Syowa", "TAAF": "Hora Francés meridional y antártico", "TAHT": "Hora de Tahiti", "TJT": "Hora de Tayikistán", "TKT": "Hora de Tokelau", "TLT": "Hora de Timor Oriental", "TMST": "Hora de Verano de Turkmenistán", "TMT": "Hora Estandar de Turkmenistán", "TOST": "Hora de Verano de Tonga", "TOT": "Hora Estandar de Tonga", "TVT": "Hora de Tuvalu", "TWT": "Hora Estandar de Taipéi", "TWT DST": "Hora de Verano de Taipéi", "ULAST": "Hora de Verano de Ulán Bator", "ULAT": "Hora Estandar de Ulán Bator", "UYST": "Hora de Verano de Uruguay", "UYT": "Hora Estandar de Uruguay", "UZT": "Hora Estandar de Uzbekistán", "UZT DST": "Hora de Verano de Uzbekistán", "VET": "Hora de Venezuela", "VLAST": "Hora de Verano de Vladivostok", "VLAT": "Hora Estandar de Vladivostok", "VOLST": "Hora de Verano de Volgogrado", "VOLT": "Hora Estandar de Volgogrado", "VOST": "Hora de Vostok", "VUT": "Hora Estandar de Vanuatu", "VUT DST": "Hora de Verano de Vanuatu", "WAKT": "Hora de Isla Wake", "WARST": "Hora de Verano del Oeste de Argentina", "WART": "Hora Estandar del Oeste de Argentina", "WAST": "Hora de Africa Occidental", "WAT": "Hora de Africa Occidental", "WESZ": "Hora de Verano de Europa Occidental", "WEZ": "Hora Estandar de Europa Occidental", "WFT": "Hora de Wallis y Futuna", "WGST": "Hora de Verano de Groenlandia Occidental", "WGT": "Hora Estandar de Groenlandia Occidental", "WIB": "Hora de Indonesia Occidental", "WIT": "Hora de Indonesia Oriental", "WITA": "Hora de Indonesia Central", "YAKST": "Hora de Verano de Yakutsk", "YAKT": "Hora Estandar de Yakutsk", "YEKST": "Hora de Verano de Ekaterinburgo", "YEKT": "Hora Estandar de Ekaterinburgo", "YST": "Yukon Ura", "МСК": "Hora Estandar de Moscú", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Hora de Kazajistán del Oeste", "شىعىش قازاق ەلى": "Hora de Kazajistán Oriental", "قازاق ەلى": "Hora de Kazajistán", "قىرعىزستان": "Hora de Kirguistán", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Hora de Verano de las Azores"},
	}
}

// Locale returns the current translators string locale
func (qu *qu_PE) Locale() string {
	return qu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'qu_PE'
func (qu *qu_PE) PluralsCardinal() []locales.PluralRule {
	return qu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'qu_PE'
func (qu *qu_PE) PluralsOrdinal() []locales.PluralRule {
	return qu.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'qu_PE'
func (qu *qu_PE) PluralsRange() []locales.PluralRule {
	return qu.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'qu_PE'
func (qu *qu_PE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'qu_PE'
func (qu *qu_PE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'qu_PE'
func (qu *qu_PE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (qu *qu_PE) MonthAbbreviated(month time.Month) string {
	if len(qu.monthsAbbreviated) == 0 {
		return ""
	}
	return qu.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (qu *qu_PE) MonthsAbbreviated() []string {
	return qu.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (qu *qu_PE) MonthNarrow(month time.Month) string {
	if len(qu.monthsNarrow) == 0 {
		return ""
	}
	return qu.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (qu *qu_PE) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (qu *qu_PE) MonthWide(month time.Month) string {
	if len(qu.monthsWide) == 0 {
		return ""
	}
	return qu.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (qu *qu_PE) MonthsWide() []string {
	return qu.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (qu *qu_PE) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(qu.daysAbbreviated) == 0 {
		return ""
	}
	return qu.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (qu *qu_PE) WeekdaysAbbreviated() []string {
	return qu.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (qu *qu_PE) WeekdayNarrow(weekday time.Weekday) string {
	if len(qu.daysNarrow) == 0 {
		return ""
	}
	return qu.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (qu *qu_PE) WeekdaysNarrow() []string {
	return qu.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (qu *qu_PE) WeekdayShort(weekday time.Weekday) string {
	if len(qu.daysShort) == 0 {
		return ""
	}
	return qu.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (qu *qu_PE) WeekdaysShort() []string {
	return qu.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (qu *qu_PE) WeekdayWide(weekday time.Weekday) string {
	if len(qu.daysWide) == 0 {
		return ""
	}
	return qu.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (qu *qu_PE) WeekdaysWide() []string {
	return qu.daysWide
}

// Decimal returns the decimal point of number
func (qu *qu_PE) Decimal() string {
	return qu.decimal
}

// Group returns the group of number
func (qu *qu_PE) Group() string {
	return qu.group
}

// Group returns the minus sign of number
func (qu *qu_PE) Minus() string {
	return qu.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'qu_PE' and handles both Whole and Real numbers based on 'v'
func (qu *qu_PE) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, qu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, qu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'qu_PE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (qu *qu_PE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, qu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, qu.percentSuffix...)

	b = append(b, qu.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'qu_PE'
func (qu *qu_PE) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := qu.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, qu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, qu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, qu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'qu_PE'
// in accounting notation.
func (qu *qu_PE) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := qu.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, qu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, qu.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, qu.decimal...)
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

// FmtDateShort returns the short date representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, qu.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, qu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, qu.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, qu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, qu.periodsAbbreviated[0]...)
	} else {
		b = append(b, qu.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, qu.periodsAbbreviated[0]...)
	} else {
		b = append(b, qu.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, qu.periodsAbbreviated[0]...)
	} else {
		b = append(b, qu.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'qu_PE'
func (qu *qu_PE) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, qu.periodsAbbreviated[0]...)
	} else {
		b = append(b, qu.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := qu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
