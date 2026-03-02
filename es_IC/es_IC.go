package es_IC

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type es_IC struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	timeSeparator          string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'es_IC' locale
func New() locales.Translator {
	return &es_IC{
		locale:                 "es_IC",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sept", "oct", "nov", "dic"},
		monthsNarrow:           []string{"", "E", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
		daysAbbreviated:        []string{"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
		daysNarrow:             []string{"D", "L", "M", "X", "J", "V", "S"},
		daysShort:              []string{"DO", "LU", "MA", "MI", "JU", "VI", "SA"},
		daysWide:               []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		timezones:              map[string]string{"ACDT": "hora de verano de Australia central", "ACST": "Hora de verano de Acre", "ACT": "Hora estándar de Acre", "ACWDT": "hora de verano de Australia centroccidental", "ACWST": "hora estándar de Australia centroccidental", "ADT": "hora de verano del Atlántico", "ADT Arabia": "hora de verano de Arabia", "AEDT": "hora de verano de Australia oriental", "AEST": "hora estándar de Australia oriental", "AFT": "hora de Afganistán", "AKDT": "hora de verano de Alaska", "AKST": "hora estándar de Alaska", "AMST": "hora de verano del Amazonas", "AMST Armenia": "hora de verano de Armenia", "AMT": "hora estándar del Amazonas", "AMT Armenia": "hora estándar de Armenia", "ANAST": "hora de verano de Anadyr", "ANAT": "hora estándar de Anadyr", "ARST": "hora de verano de Argentina", "ART": "hora estándar de Argentina", "AST": "hora estándar del Atlántico", "AST Arabia": "hora estándar de Arabia", "AWDT": "hora de verano de Australia occidental", "AWST": "hora estándar de Australia occidental", "AZST": "hora de verano de Azerbaiyán", "AZT": "hora estándar de Azerbaiyán", "BDT Bangladesh": "hora de verano de Bangladés", "BNT": "hora de Brunéi", "BOT": "hora de Bolivia", "BRST": "hora de verano de Brasilia", "BRT": "hora estándar de Brasilia", "BST Bangladesh": "hora estándar de Bangladés", "BT": "hora de Bután", "CAST": "CAST", "CAT": "hora de África central", "CCT": "hora de las Islas Cocos", "CDT": "hora de verano central", "CHADT": "hora de verano de Chatham", "CHAST": "hora estándar de Chatham", "CHUT": "hora de Chuuk", "CKT": "hora estándar de las Islas Cook", "CKT DST": "hora de verano media de las Islas Cook", "CLST": "hora de verano de Chile", "CLT": "hora estándar de Chile", "COST": "hora de verano de Colombia", "COT": "hora estándar de Colombia", "CST": "hora estándar central", "CST China": "hora estándar de China", "CST China DST": "hora de verano de China", "CVST": "hora de verano de Cabo Verde", "CVT": "hora estándar de Cabo Verde", "CXT": "hora de la Isla de Navidad", "ChST": "hora estándar de Chamorro", "ChST NMI": "Hora de las Islas Marianas del Norte", "CuDT": "hora de verano de Cuba", "CuST": "hora estándar de Cuba", "DAVT": "hora de Davis", "DDUT": "hora de Dumont-d’Urville", "EASST": "hora de verano de la isla de Pascua", "EAST": "hora estándar de la isla de Pascua", "EAT": "hora de África oriental", "ECT": "hora de Ecuador", "EDT": "hora de verano oriental", "EGDT": "hora de verano de Groenlandia oriental", "EGST": "hora estándar de Groenlandia oriental", "EST": "hora estándar oriental", "FEET": "hora del extremo oriental de Europa", "FJT": "hora estándar de Fiyi", "FJT Summer": "hora de verano de Fiyi", "FKST": "hora de verano de las islas Malvinas", "FKT": "hora estándar de las islas Malvinas", "FNST": "hora de verano de Fernando de Noronha", "FNT": "hora estándar de Fernando de Noronha", "GALT": "hora de Galápagos", "GAMT": "hora de Gambier", "GEST": "hora de verano de Georgia", "GET": "hora estándar de Georgia", "GFT": "hora de la Guayana Francesa", "GIT": "hora de las islas Gilbert", "GMT": "hora del meridiano de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "hora estándar del Golfo", "GST Guam": "Hora estándar de Guam", "GYT": "hora de Guyana", "HADT": "hora estándar de Hawái-Aleutianas", "HAST": "hora estándar de Hawái-Aleutianas", "HKST": "hora de verano de Hong Kong", "HKT": "hora estándar de Hong Kong", "HOVST": "hora de verano de Hovd", "HOVT": "hora estándar de Hovd", "ICT": "hora de Indochina", "IDT": "hora de verano de Israel", "IOT": "hora del océano Índico", "IRKST": "hora de verano de Irkutsk", "IRKT": "hora estándar de Irkutsk", "IRST": "hora estándar de Irán", "IRST DST": "hora de verano de Irán", "IST": "hora estándar de la India", "IST Israel": "hora estándar de Israel", "JDT": "hora de verano de Japón", "JST": "hora estándar de Japón", "KOST": "hora de Kosrae", "KRAST": "hora de verano de Krasnoyarsk", "KRAT": "hora estándar de Krasnoyarsk", "KST": "hora estándar de Corea", "KST DST": "hora de verano de Corea", "LHDT": "hora de verano de Lord Howe", "LHST": "hora estándar de Lord Howe", "LINT": "hora de las Espóradas Ecuatoriales", "MAGST": "hora de verano de Magadán", "MAGT": "hora estándar de Magadán", "MART": "hora de Marquesas", "MAWT": "hora de Mawson", "MDT": "hora de verano de las Montañas Rocosas", "MESZ": "hora de verano de Europa central", "MEZ": "hora estándar de Europa central", "MHT": "hora de las Islas Marshall", "MMT": "hora de Myanmar", "MSD": "hora de verano de Moscú", "MST": "hora estándar de las Montañas Rocosas", "MUST": "hora de verano de Mauricio", "MUT": "hora estándar de Mauricio", "MVT": "hora de Maldivas", "MYT": "hora de Malasia", "NCT": "hora estándar de Nueva Caledonia", "NDT": "hora de verano de Terranova", "NDT New Caledonia": "hora de verano de Nueva Caledonia", "NFDT": "hora de verano de la isla Norfolk", "NFT": "hora estándar de la isla Norfolk", "NOVST": "hora de verano de Novosibirsk", "NOVT": "hora estándar de Novosibirsk", "NPT": "hora de Nepal", "NRT": "hora de Nauru", "NST": "hora estándar de Terranova", "NUT": "hora de Niue", "NZDT": "hora de verano de Nueva Zelanda", "NZST": "hora estándar de Nueva Zelanda", "OESZ": "hora de verano de Europa oriental", "OEZ": "hora estándar de Europa oriental", "OMSST": "hora de verano de Omsk", "OMST": "hora estándar de Omsk", "PDT": "hora de verano del Pacífico", "PDTM": "hora de verano del Pacífico de México", "PETDT": "hora de verano de Kamchatka", "PETST": "hora estándar de Kamchatka", "PGT": "hora de Papúa Nueva Guinea", "PHOT": "hora de las Islas Fénix", "PKT": "hora estándar de Pakistán", "PKT DST": "hora de verano de Pakistán", "PMDT": "hora de verano de San Pedro y Miquelón", "PMST": "hora estándar de San Pedro y Miquelón", "PONT": "hora de Pohnpei", "PST": "hora estándar del Pacífico", "PST Philippine": "hora estándar de Filipinas", "PST Philippine DST": "hora de verano de Filipinas", "PST Pitcairn": "hora de Pitcairn", "PSTM": "hora estándar del Pacífico de México", "PWT": "hora de Palaos", "PYST": "hora de verano de Paraguay", "PYT": "hora estándar de Paraguay", "PYT Korea": "hora de Pyongyang", "RET": "hora de Reunión", "ROTT": "hora de Rothera", "SAKST": "hora de verano de Sajalín", "SAKT": "hora estándar de Sajalín", "SAMST": "hora de verano de Samara", "SAMT": "hora estándar de Samara", "SAST": "hora de Sudáfrica", "SBT": "hora de las Islas Salomón", "SCT": "hora de Seychelles", "SGT": "hora de Singapur", "SLST": "Hora de Sri Lanka", "SRT": "hora de Surinam", "SST Samoa": "hora estándar de Samoa", "SST Samoa Apia": "hora estándar de Apia", "SST Samoa Apia DST": "horario de verano de Apia", "SST Samoa DST": "hora de verano de Samoa", "SYOT": "hora de Syowa", "TAAF": "hora de Antártida y Territorios Australes Franceses", "TAHT": "hora de Tahití", "TJT": "hora de Tayikistán", "TKT": "hora de Tokelau", "TLT": "hora de Timor Oriental", "TMST": "hora de verano de Turkmenistán", "TMT": "hora estándar de Turkmenistán", "TOST": "hora de verano de Tonga", "TOT": "hora estándar de Tonga", "TVT": "hora de Tuvalu", "TWT": "hora estándar de Taipéi", "TWT DST": "hora de verano de Taipéi", "ULAST": "hora de verano de Ulán Bator", "ULAT": "hora estándar de Ulán Bator", "UYST": "hora de verano de Uruguay", "UYT": "hora estándar de Uruguay", "UZT": "hora estándar de Uzbekistán", "UZT DST": "hora de verano de Uzbekistán", "VET": "hora de Venezuela", "VLAST": "hora de verano de Vladivostok", "VLAT": "hora estándar de Vladivostok", "VOLST": "hora de verano de Volgogrado", "VOLT": "hora estándar de Volgogrado", "VOST": "hora de Vostok", "VUT": "hora estándar de Vanuatu", "VUT DST": "hora de verano de Vanuatu", "WAKT": "hora de la isla Wake", "WARST": "hora de verano de Argentina occidental", "WART": "hora estándar de Argentina occidental", "WAST": "hora de África occidental", "WAT": "hora de África occidental", "WESZ": "hora de verano de Europa occidental", "WEZ": "hora estándar de Europa occidental", "WFT": "hora de Wallis y Futuna", "WGST": "hora de verano de Groenlandia occidental", "WGT": "hora estándar de Groenlandia occidental", "WIB": "hora de Indonesia occidental", "WIT": "hora de Indonesia oriental", "WITA": "hora de Indonesia central", "YAKST": "hora de verano de Yakutsk", "YAKT": "hora estándar de Yakutsk", "YEKST": "hora de verano de Ekaterimburgo", "YEKT": "hora estándar de Ekaterimburgo", "YST": "hora de Yukón", "МСК": "hora estándar de Moscú", "اقتاۋ": "Hora estándar de Aktau", "اقتاۋ قالاسى": "Hora de verano de Aktau", "اقتوبە": "Hora estándar de Aktobe", "اقتوبە قالاسى": "Hora de verano de Aktobe", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "hora de Kazajistán occidental", "شىعىش قازاق ەلى": "hora de Kazajistán oriental", "قازاق ەلى": "hora de Kazajistán", "قىرعىزستان": "hora de Kirguistán", "قىزىلوردا": "Hora estándar de Qyzylorda", "قىزىلوردا قالاسى": "Hora de verano de Qyzylorda", "∅∅∅": "hora de verano de Perú"},
	}
}

// Locale returns the current translators string locale
func (es *es_IC) Locale() string {
	return es.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'es_IC'
func (es *es_IC) PluralsCardinal() []locales.PluralRule {
	return es.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'es_IC'
func (es *es_IC) PluralsOrdinal() []locales.PluralRule {
	return es.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'es_IC'
func (es *es_IC) PluralsRange() []locales.PluralRule {
	return es.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'es_IC'
func (es *es_IC) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	e := int64(0)
	iMod1000000 := i % 1000000

	if n == 1 {
		return locales.PluralRuleOne
	} else if (e == 0 && i != 0 && iMod1000000 == 0 && v == 0) || (e < 0 || e > 5) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'es_IC'
func (es *es_IC) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'es_IC'
func (es *es_IC) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (es *es_IC) MonthAbbreviated(month time.Month) string {
	return es.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (es *es_IC) MonthsAbbreviated() []string {
	return es.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (es *es_IC) MonthNarrow(month time.Month) string {
	return es.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (es *es_IC) MonthsNarrow() []string {
	return es.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (es *es_IC) MonthWide(month time.Month) string {
	return es.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (es *es_IC) MonthsWide() []string {
	return es.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (es *es_IC) WeekdayAbbreviated(weekday time.Weekday) string {
	return es.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (es *es_IC) WeekdaysAbbreviated() []string {
	return es.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (es *es_IC) WeekdayNarrow(weekday time.Weekday) string {
	return es.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (es *es_IC) WeekdaysNarrow() []string {
	return es.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (es *es_IC) WeekdayShort(weekday time.Weekday) string {
	return es.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (es *es_IC) WeekdaysShort() []string {
	return es.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (es *es_IC) WeekdayWide(weekday time.Weekday) string {
	return es.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (es *es_IC) WeekdaysWide() []string {
	return es.daysWide
}

// Decimal returns the decimal point of number
func (es *es_IC) Decimal() string {
	return es.decimal
}

// Group returns the group of number
func (es *es_IC) Group() string {
	return es.group
}

// Group returns the minus sign of number
func (es *es_IC) Minus() string {
	return es.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'es_IC' and handles both Whole and Real numbers based on 'v'
func (es *es_IC) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'es_IC' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (es *es_IC) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, es.percentSuffix...)

	b = append(b, es.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'es_IC'
func (es *es_IC) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, es.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'es_IC'
// in accounting notation.
func (es *es_IC) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, es.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, es.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'es_IC'
func (es *es_IC) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'es_IC'
func (es *es_IC) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'es_IC'
func (es *es_IC) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'es_IC'
func (es *es_IC) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, es.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'es_IC'
func (es *es_IC) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'es_IC'
func (es *es_IC) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'es_IC'
func (es *es_IC) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'es_IC'
func (es *es_IC) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := es.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
