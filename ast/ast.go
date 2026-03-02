package ast

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ast struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'ast' locale
func New() locales.Translator {
	return &ast{
		locale:                 "ast",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "xin", "feb", "mar", "abr", "may", "xun", "xnt", "ago", "set", "och", "pay", "avi"},
		monthsNarrow:           []string{"", "X", "F", "M", "A", "M", "X", "X", "A", "S", "O", "P", "A"},
		monthsWide:             []string{"", "de xineru", "de febreru", "de marzu", "d’abril", "de mayu", "de xunu", "de xunetu", "d’agostu", "de setiembre", "d’ochobre", "de payares", "d’avientu"},
		daysAbbreviated:        []string{"dom", "llu", "mar", "mié", "xue", "vie", "sáb"},
		daysNarrow:             []string{"D", "L", "M", "M", "X", "V", "S"},
		daysShort:              []string{"do", "ll", "ma", "mi", "xu", "vi", "sá"},
		daysWide:               []string{"domingu", "llunes", "martes", "miércoles", "xueves", "vienres", "sábadu"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"de la mañana", "de la tarde"},
		erasAbbreviated:        []string{"e.C.", "d.C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"enantes de Cristu", "después de Cristu"},
		timezones:              map[string]string{"ACDT": "Hora braniega d’Australia central", "ACST": "Hora estándar d’Australia central", "ACT": "hora estándar d’Acre", "ACWDT": "Hora braniega d’Australia central del oeste", "ACWST": "Hora estándar d’Australia central del oeste", "ADT": "Hora braniega del Atlánticu", "ADT Arabia": "Hora braniega d’Arabia", "AEDT": "Hora braniega d’Australia del este", "AEST": "Hora estándar d’Australia del este", "AFT": "Hora d’Afganistán", "AKDT": "Hora braniega d’Alaska", "AKST": "Hora estándar d’Alaska", "AMST": "Hora braniega del Amazonas", "AMST Armenia": "Hora braniega d’Armenia", "AMT": "Hora estándar del Amazonas", "AMT Armenia": "Hora estándar d’Armenia", "ANAST": "hora braniega d’Anadyr", "ANAT": "hora estándar d’Anadyr", "ARST": "Hora braniega d’Arxentina", "ART": "Hora estándar d’Arxentina", "AST": "Hora estándar del Atlánticu", "AST Arabia": "Hora estándar d’Arabia", "AWDT": "Hora braniega d’Australia del oeste", "AWST": "Hora estándar d’Australia del oeste", "AZST": "Hora braniega d’Azerbaixán", "AZT": "Hora estándar d’Azerbaixán", "BDT Bangladesh": "Hora braniega de Bangladex", "BNT": "Hora de Brunéi Darussalam", "BOT": "Hora de Bolivia", "BRST": "Hora braniega de Brasilia", "BRT": "Hora estándar de Brasilia", "BST Bangladesh": "Hora estándar de Bangladex", "BT": "Hora de Bután", "CAST": "Hora de Casey", "CAT": "Hora d’África central", "CCT": "Hora de les Islles Cocos", "CDT": "Hora braniega central norteamericana", "CHADT": "Hora braniega de Chatham", "CHAST": "Hora estándar de Chatham", "CHUT": "Hora de Chuuk", "CKT": "Hora estándar de les Islles Cook", "CKT DST": "Hora media braniega de les Islles Cook", "CLST": "Hora braniega de Chile", "CLT": "Hora estándar de Chile", "COST": "Hora braniega de Colombia", "COT": "Hora estándar de Colombia", "CST": "Hora estándar central norteamericana", "CST China": "Hora estándar de China", "CST China DST": "Hora braniega de China", "CVST": "Hora braniega de Cabu Verde", "CVT": "Hora estándar de Cabu Verde", "CXT": "Hora estándar de la Islla Christmas", "ChST": "Hora estándar de Chamorro", "ChST NMI": "Hora de les Islles Marianes del Norte", "CuDT": "Hora braniega de Cuba", "CuST": "Hora estándar de Cuba", "DAVT": "Hora de Davis", "DDUT": "Hora de Dumont-d’Urville", "EASST": "Hora braniega de la Islla de Pascua", "EAST": "Hora estándar de la Islla de Pascua", "EAT": "Hora d’África del este", "ECT": "Hora d’Ecuador", "EDT": "Hora braniega del este norteamericanu", "EGDT": "Hora braniega de Groenlandia oriental", "EGST": "Hora estándar de Groenlandia oriental", "EST": "Hora estándar del este norteamericanu", "FEET": "Hora d’Europa del estremu este", "FJT": "Hora estándar de Fixi", "FJT Summer": "Hora braniega de Fixi", "FKST": "Hora braniega de les Islles Falkland", "FKT": "Hora estándar de les Islles Falkland", "FNST": "Hora braniega de Fernando de Noronha", "FNT": "Hora estándar de Fernando de Noronha", "GALT": "Hora de Galápagos", "GAMT": "Hora de Gambier", "GEST": "Hora braniega de Xeorxa", "GET": "Hora estándar de Xeorxa", "GFT": "Hora de La Guyana Francesa", "GIT": "Hora de les Islles Gilbert", "GMT": "Hora media de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Hora estándar del Golfu", "GST Guam": "Hora estándar de Guam", "GYT": "Hora de La Guyana", "HADT": "Hora estándar de Hawaii-Aleutianes", "HAST": "Hora estándar de Hawaii-Aleutianes", "HKST": "Hora braniega de Ḥong Kong", "HKT": "Hora estándar de Ḥong Kong", "HOVST": "Hora braniega de Hovd", "HOVT": "Hora estándar de Hovd", "ICT": "Hora d’Indochina", "IDT": "Hora braniega d’Israel", "IOT": "Hora del Océanu Índicu", "IRKST": "Hora braniega d’Irkutsk", "IRKT": "Hora estándar d’Irkutsk", "IRST": "Hora estándar d’Irán", "IRST DST": "Hora braniega d’Irán", "IST": "Hora estándar de la India", "IST Israel": "Hora estándar d’Israel", "JDT": "Hora braniega de Xapón", "JST": "Hora estándar de Xapón", "KOST": "Hora de Kosrae", "KRAST": "Hora braniega de Krasnoyarsk", "KRAT": "Hora estándar de Krasnoyarsk", "KST": "Hora estándar de Corea", "KST DST": "Hora braniega de Corea", "LHDT": "Hora braniega de Lord Howe", "LHST": "Hora estándar de Lord Howe", "LINT": "Hora de les Islles Line", "MAGST": "Hora braniega de Magadán", "MAGT": "Hora estándar de Magadán", "MART": "Hora de les Marqueses", "MAWT": "Hora de Mawson", "MDT": "Hora braniega de Macáu", "MESZ": "Hora braniega d’Europa Central", "MEZ": "Hora estándar d’Europa Central", "MHT": "Hora de les Islles Marshall", "MMT": "Hora de Myanmar", "MSD": "Hora braniega de Moscú", "MST": "Hora estándar de Macáu", "MUST": "Hora braniega de Mauriciu", "MUT": "Hora estándar de Mauriciu", "MVT": "Hora de Les Maldives", "MYT": "Hora de Malasia", "NCT": "Hora estándar de Nueva Caledonia", "NDT": "Hora braniega de Newfoundland", "NDT New Caledonia": "Hora braniega de Nueva Caledonia", "NFDT": "Hora braniega de la Islla Norfolk", "NFT": "Hora estándar de la Islla Norfolk", "NOVST": "Hora braniega de Novosibirsk", "NOVT": "Hora estándar de Novosibirsk", "NPT": "Hora del Nepal", "NRT": "Hora de Nauru", "NST": "Hora estándar de Newfoundland", "NUT": "Hora de Niue", "NZDT": "Hora braniega de Nueva Zelanda", "NZST": "Hora estándar de Nueva Zelanda", "OESZ": "Hora braniega d’Europa del Este", "OEZ": "Hora estándar d’Europa del Este", "OMSST": "Hora braniega d’Omsk", "OMST": "Hora estándar d’Omsk", "PDT": "Hora braniega del Pacíficu norteamericanu", "PDTM": "Hora braniega del Pacíficu de Méxicu", "PETDT": "hora braniega de Petropavlovsk-Kamchatski", "PETST": "hora estandar de Petropavlovsk-Kamchatski", "PGT": "Hora de Papúa Nueva Guinea", "PHOT": "Hora de les Islles Phoenix", "PKT": "Hora estándar del Paquistán", "PKT DST": "Hora braniega del Paquistán", "PMDT": "Hora braniega de Saint Pierre y Miquelon", "PMST": "Hora estándar de Saint Pierre y Miquelon", "PONT": "Hora de Ponape", "PST": "Hora estándar del Pacíficu norteamericanu", "PST Philippine": "Hora estándar de Filipines", "PST Philippine DST": "Hora de branu de Filipines", "PST Pitcairn": "Hora de Pitcairn", "PSTM": "Hora estándar del Pacíficu de Méxicu", "PWT": "Hora de Palau", "PYST": "Hora braniega del Paraguái", "PYT": "Hora estándar del Paraguái", "PYT Korea": "hora de Pyongyang", "RET": "Hora de Reunión", "ROTT": "Hora de Rothera", "SAKST": "Hora braniega de Saxalín", "SAKT": "Hora estándar de Saxalín", "SAMST": "Hora braniega de Samara", "SAMT": "Hora estándar de Samara", "SAST": "Hora de Sudáfrica", "SBT": "Hora de les Islles Salomón", "SCT": "Hora de Les Seixeles", "SGT": "Hora estándar de Singapur", "SLST": "Hora de Lanka", "SRT": "Hora del Surinam", "SST Samoa": "Hora estándar de Samoa", "SST Samoa Apia": "Hora estándar d’Apia", "SST Samoa Apia DST": "Hora braniega d’Apia", "SST Samoa DST": "Hora braniega de Samoa", "SYOT": "Hora de Syowa", "TAAF": "Hora del sur y l’antárticu francés", "TAHT": "Hora de Tahiti", "TJT": "Hora del Taxiquistán", "TKT": "Hora de Tokelau", "TLT": "Hora de Timor Oriental", "TMST": "Hora braniega del Turkmenistán", "TMT": "Hora estándar del Turkmenistán", "TOST": "Hora braniega de Tonga", "TOT": "Hora estándar de Tonga", "TVT": "Hora de Tuvalu", "TWT": "Hora estándar de Taipéi", "TWT DST": "Hora braniega de Taipéi", "ULAST": "Hora braniega d’Ulán Bátor", "ULAT": "Hora estándar d’Ulán Bátor", "UYST": "Hora braniega del Uruguái", "UYT": "Hora estándar del Uruguái", "UZT": "Hora estándar del Uzbequistán", "UZT DST": "Hora braniega del Uzbequistán", "VET": "Hora de Venezuela", "VLAST": "Hora braniega de Vladivostok", "VLAT": "Hora estándar de Vladivostok", "VOLST": "Hora braniega de Volgográu", "VOLT": "Hora estándar de Volgográu", "VOST": "Hora de Vostok", "VUT": "Hora estándar de Vanuatu", "VUT DST": "Hora braniega de Vanuatu", "WAKT": "Hora de la Islla Wake", "WARST": "Hora braniega occidental d’Arxentina", "WART": "Hora estándar occidental d’Arxentina", "WAST": "Hora d’África del oeste", "WAT": "Hora d’África del oeste", "WESZ": "Hora braniega d’Europa Occidental", "WEZ": "Hora estándar d’Europa Occidental", "WFT": "Hora de Wallis y Futuna", "WGST": "Hora braniega de Groenlandia occidental", "WGT": "Hora estándar de Groenlandia occidental", "WIB": "Hora d’Indonesia del oeste", "WIT": "Hora d’Indonesia del este", "WITA": "Hora d’Indonesia central", "YAKST": "Hora braniega de Yakutsk", "YAKT": "Hora estándar de Yakutsk", "YEKST": "Hora braniega de Yekaterimburgu", "YEKT": "Hora estándar de Yekaterimburgu", "YST": "YST", "МСК": "Hora estándar de Moscú", "اقتاۋ": "Hora estándar d’Aqtau", "اقتاۋ قالاسى": "Hora braniega d’Aqtau", "اقتوبە": "Hora estándar d’Aqtobe", "اقتوبە قالاسى": "Hora braniega d’Aqtobe", "الماتى": "hora estándar d’Almaty", "الماتى قالاسى": "hora braniega d’Almaty", "باتىس قازاق ەلى": "Hora del Kazakstán occidental", "شىعىش قازاق ەلى": "Hora del Kazakstán oriental", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Hora del Kirguistán", "قىزىلوردا": "Hora estándar de Qyzylorda", "قىزىلوردا قالاسى": "Hora braniega de Qyzylorda", "∅∅∅": "Hora braniega del Perú"},
	}
}

// Locale returns the current translators string locale
func (ast *ast) Locale() string {
	return ast.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ast'
func (ast *ast) PluralsCardinal() []locales.PluralRule {
	return ast.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ast'
func (ast *ast) PluralsOrdinal() []locales.PluralRule {
	return ast.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ast'
func (ast *ast) PluralsRange() []locales.PluralRule {
	return ast.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ast'
func (ast *ast) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ast'
func (ast *ast) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ast'
func (ast *ast) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ast *ast) MonthAbbreviated(month time.Month) string {
	return ast.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ast *ast) MonthsAbbreviated() []string {
	return ast.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ast *ast) MonthNarrow(month time.Month) string {
	return ast.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ast *ast) MonthsNarrow() []string {
	return ast.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ast *ast) MonthWide(month time.Month) string {
	return ast.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ast *ast) MonthsWide() []string {
	return ast.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ast *ast) WeekdayAbbreviated(weekday time.Weekday) string {
	return ast.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ast *ast) WeekdaysAbbreviated() []string {
	return ast.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ast *ast) WeekdayNarrow(weekday time.Weekday) string {
	return ast.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ast *ast) WeekdaysNarrow() []string {
	return ast.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ast *ast) WeekdayShort(weekday time.Weekday) string {
	return ast.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ast *ast) WeekdaysShort() []string {
	return ast.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ast *ast) WeekdayWide(weekday time.Weekday) string {
	return ast.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ast *ast) WeekdaysWide() []string {
	return ast.daysWide
}

// Decimal returns the decimal point of number
func (ast *ast) Decimal() string {
	return ast.decimal
}

// Group returns the group of number
func (ast *ast) Group() string {
	return ast.group
}

// Group returns the minus sign of number
func (ast *ast) Minus() string {
	return ast.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ast' and handles both Whole and Real numbers based on 'v'
func (ast *ast) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ast.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ast.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ast.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ast' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ast *ast) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ast.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ast.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ast.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ast'
func (ast *ast) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ast.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ast.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ast.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ast.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ast.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ast.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ast'
// in accounting notation.
func (ast *ast) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ast.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ast.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ast.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ast.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ast.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ast.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ast.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ast'
func (ast *ast) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ast'
func (ast *ast) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ast'
func (ast *ast) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ast'
func (ast *ast) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ast.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ast'
func (ast *ast) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ast'
func (ast *ast) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ast'
func (ast *ast) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ast'
func (ast *ast) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ast.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
