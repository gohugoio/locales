package gl_ES

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type gl_ES struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
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

// New returns a new instance of translator for the 'gl_ES' locale
func New() locales.Translator {
	return &gl_ES{
		locale:                 "gl_ES",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "xan.", "feb.", "mar.", "abr.", "maio", "xuño", "xul.", "ago.", "set.", "out.", "nov.", "dec."},
		monthsNarrow:           []string{"", "x.", "f.", "m.", "a.", "m.", "x.", "x.", "a.", "s.", "o.", "n.", "d."},
		monthsWide:             []string{"", "xaneiro", "febreiro", "marzo", "abril", "maio", "xuño", "xullo", "agosto", "setembro", "outubro", "novembro", "decembro"},
		daysAbbreviated:        []string{"dom.", "luns", "mar.", "mér.", "xov.", "ven.", "sáb."},
		daysNarrow:             []string{"d.", "l.", "m.", "m.", "x.", "v.", "s."},
		daysShort:              []string{"do.", "lu.", "ma.", "mé.", "xo.", "ve.", "sá."},
		daysWide:               []string{"domingo", "luns", "martes", "mércores", "xoves", "venres", "sábado"},
		periodsAbbreviated:     []string{"a.m.", "p.m."},
		erasAbbreviated:        []string{"a.C.", "d.C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"antes de Cristo", "despois de Cristo"},
		timezones:              map[string]string{"ACDT": "hora de verán de Australia Central", "ACST": "hora de verán de Acre", "ACT": "hora estándar de Acre", "ACWDT": "hora de verán de Australia Occidental Central", "ACWST": "hora estándar de Australia Occidental Central", "ADT": "hora de verán do Atlántico", "ADT Arabia": "hora de verán árabe", "AEDT": "hora de verán de Australia Oriental", "AEST": "hora estándar de Australia Oriental", "AFT": "hora de Afganistán", "AKDT": "hora de verán de Alasca", "AKST": "hora estándar de Alasca", "AMST": "hora de verán do Amazonas", "AMST Armenia": "hora de verán de Armenia", "AMT": "hora estándar do Amazonas", "AMT Armenia": "hora estándar de Armenia", "ANAST": "Horario de verán de Anadir", "ANAT": "Horario estándar de Anadir", "ARST": "hora de verán da Arxentina", "ART": "hora estándar da Arxentina", "AST": "hora estándar do Atlántico", "AST Arabia": "hora estándar árabe", "AWDT": "hora de verán de Australia Occidental", "AWST": "hora estándar de Australia Occidental", "AZST": "hora de verán de Acerbaixán", "AZT": "hora estándar de Acerbaixán", "BDT Bangladesh": "hora de verán de Bangladesh", "BNT": "hora de Brunei Darussalam", "BOT": "hora de Bolivia", "BRST": "hora de verán de Brasilia", "BRT": "hora estándar de Brasilia", "BST Bangladesh": "hora estándar de Bangladesh", "BT": "hora de Bután", "CAST": "CAST", "CAT": "hora de África Central", "CCT": "hora das Illas Cocos", "CDT": "hora de verán central, Norteamérica", "CHADT": "hora de verán de Chatham", "CHAST": "hora estándar de Chatham", "CHUT": "hora de Chuuk", "CKT": "hora estándar das Illas Cook", "CKT DST": "hora de verán medio das Illas Cook", "CLST": "hora de verán de Chile", "CLT": "hora estándar de Chile", "COST": "hora de verán de Colombia", "COT": "hora estándar de Colombia", "CST": "hora estándar central, Norteamérica", "CST China": "hora estándar da China", "CST China DST": "hora de verán da China", "CVST": "hora de verán de Cabo Verde", "CVT": "hora estándar de Cabo Verde", "CXT": "hora da Illa Christmas", "ChST": "hora estándar chamorro", "ChST NMI": "ChST NMI", "CuDT": "hora de verán de Cuba", "CuST": "hora estándar de Cuba", "DAVT": "hora de Davis", "DDUT": "hora de Dumont-d’Urville", "EASST": "hora de verán da Illa de Pascua", "EAST": "hora estándar da Illa de Pascua", "EAT": "hora de África Oriental", "ECT": "hora de Ecuador", "EDT": "hora de verán do leste, América do Norte", "EGDT": "hora de verán de Groenlandia Oriental", "EGST": "hora estándar de Groenlandia Oriental", "EST": "hora estándar do leste, América do Norte", "FEET": "hora do extremo leste europeo", "FJT": "hora estándar de Fixi", "FJT Summer": "hora de verán de Fixi", "FKST": "hora de verán das Illas Malvinas", "FKT": "hora estándar das Illas Malvinas", "FNST": "hora de verán de Fernando de Noronha", "FNT": "hora estándar de Fernando de Noronha", "GALT": "hora das Galápagos", "GAMT": "hora de Gambier", "GEST": "hora de verán de Xeorxia", "GET": "hora estándar de Xeorxia", "GFT": "hora da Güiana Francesa", "GIT": "hora das Illas Gilbert", "GMT": "hora do meridiano de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "hora do Golfo", "GST Guam": "GST Guam", "GYT": "hora da Güiana", "HADT": "hora estándar de Hawai-illas Aleutianas", "HAST": "hora estándar de Hawai-illas Aleutianas", "HKST": "hora de verán de Hong Kong", "HKT": "hora estándar de Hong Kong", "HOVST": "hora de verán de Hovd", "HOVT": "hora estándar de Hovd", "ICT": "hora de Indochina", "IDT": "hora de verán de Israel", "IOT": "hora do Océano Índico", "IRKST": "hora de verán de Irkutsk", "IRKT": "hora estándar de Irkutsk", "IRST": "hora estándar de Irán", "IRST DST": "hora de verán de Irán", "IST": "hora da India", "IST Israel": "hora estándar de Israel", "JDT": "hora de verán do Xapón", "JST": "hora estándar do Xapón", "KOST": "hora de Kosrae", "KRAST": "hora de verán de Krasnoiarsk", "KRAT": "hora estándar de Krasnoiarsk", "KST": "hora estándar de Corea", "KST DST": "hora de verán de Corea", "LHDT": "hora de verán de Lord Howe", "LHST": "hora estándar de Lord Howe", "LINT": "hora das Illas da Liña", "MAGST": "hora de verán de Magadan", "MAGT": "hora estándar de Magadan", "MART": "hora das Marquesas", "MAWT": "hora de Mawson", "MDT": "MDT", "MESZ": "hora de verán de Europa Central", "MEZ": "hora estándar de Europa Central", "MHT": "hora das Illas Marshall", "MMT": "hora de Myanmar", "MSD": "hora de verán de Moscova", "MST": "MST", "MUST": "hora de verán de Mauricio", "MUT": "hora estándar de Mauricio", "MVT": "hora das Maldivas", "MYT": "hora de Malaisia", "NCT": "hora estándar de Nova Caledonia", "NDT": "hora de verán de Terra Nova", "NDT New Caledonia": "hora de verán de Nova Caledonia", "NFDT": "hora de verán da Illa Norfolk", "NFT": "hora estándar da Illa Norfolk", "NOVST": "hora de verán de Novosibirsk", "NOVT": "hora estándar de Novosibirsk", "NPT": "hora de Nepal", "NRT": "hora de Nauru", "NST": "hora estándar de Terra Nova", "NUT": "hora de Niue", "NZDT": "hora de verán de Nova Zelandia", "NZST": "hora estándar de Nova Zelandia", "OESZ": "hora de verán de Europa Oriental", "OEZ": "hora estándar de Europa Oriental", "OMSST": "hora de verán de Omsk", "OMST": "hora estándar de Omsk", "PDT": "hora de verán do Pacífico, América do Norte", "PDTM": "hora de verán do Pacífico mexicano", "PETDT": "Horario de verán de Petropávlovsk-Kamchatski", "PETST": "Horario estándar de Petropávlovsk-Kamchatski", "PGT": "hora de Papúa-Nova Guinea", "PHOT": "hora das Illas Fénix", "PKT": "hora estándar de Paquistán", "PKT DST": "hora de verán de Paquistán", "PMDT": "hora de verán de Saint Pierre et Miquelon", "PMST": "hora estándar de Saint Pierre et Miquelon", "PONT": "hora de Pohnpei", "PST": "hora estándar do Pacífico, América do Norte", "PST Philippine": "hora estándar de Filipinas", "PST Philippine DST": "hora de verán de Filipinas", "PST Pitcairn": "hora de Pitcairn", "PSTM": "hora estándar do Pacífico mexicano", "PWT": "hora de Palau", "PYST": "hora de verán do Paraguai", "PYT": "hora estándar do Paraguai", "PYT Korea": "hora de Pyongyang", "RET": "hora de Reunión", "ROTT": "hora de Rothera", "SAKST": "hora de verán de Sakhalín", "SAKT": "hora estándar de Sakhalín", "SAMST": "Horario de verán de Samara", "SAMT": "Horario estándar de Samara", "SAST": "hora de África Meridional", "SBT": "hora das Illas Salomón", "SCT": "hora das Seychelles", "SGT": "hora de Singapur", "SLST": "SLST", "SRT": "hora de Suriname", "SST Samoa": "hora estándar de Samoa", "SST Samoa Apia": "hora estándar de Apia", "SST Samoa Apia DST": "hora de verán de Apia", "SST Samoa DST": "hora de verán de Samoa", "SYOT": "hora de Syowa", "TAAF": "hora das Terras Austrais e Antárticas Francesas", "TAHT": "hora de Tahití", "TJT": "hora de Taxiquistán", "TKT": "hora de Tokelau", "TLT": "hora de Timor Leste", "TMST": "hora de verán de Turkmenistán", "TMT": "hora estándar de Turkmenistán", "TOST": "hora de verán de Tonga", "TOT": "hora estándar de Tonga", "TVT": "hora de Tuvalu", "TWT": "hora estándar de Taipei", "TWT DST": "hora de verán de Taipei", "ULAST": "hora de verán de Ulaanbaatar", "ULAT": "hora estándar de Ulaanbaatar", "UYST": "hora de verán do Uruguai", "UYT": "hora estándar do Uruguai", "UZT": "hora estándar de Uzbekistán", "UZT DST": "hora de verán de Uzbekistán", "VET": "hora de Venezuela", "VLAST": "hora de verán de Vladivostok", "VLAT": "hora estándar de Vladivostok", "VOLST": "hora de verán de Volgogrado", "VOLT": "hora estándar de Volgogrado", "VOST": "hora de Vostok", "VUT": "hora estándar de Vanuatu", "VUT DST": "hora de verán de Vanuatu", "WAKT": "hora da Illa Wake", "WARST": "hora de verán da Arxentina Occidental", "WART": "hora estándar da Arxentina Occidental", "WAST": "hora de África Occidental", "WAT": "hora de África Occidental", "WESZ": "hora de verán de Europa Occidental", "WEZ": "hora estándar de Europa Occidental", "WFT": "hora de Wallis e Futuna", "WGST": "hora de verán de Groenlandia Occidental", "WGT": "hora estándar de Groenlandia Occidental", "WIB": "hora de Indonesia Occidental", "WIT": "hora de Indonesia Oriental", "WITA": "hora de Indonesia Central", "YAKST": "hora de verán de Iakutsk", "YAKT": "hora estándar de Iakutsk", "YEKST": "hora de verán de Ekaterimburgo", "YEKT": "hora estándar de Ekaterimburgo", "YST": "hora de Yukon", "МСК": "hora estándar de Moscova", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "hora de Kazakistán Occidental", "شىعىش قازاق ەلى": "hora de Kazakistán Oriental", "قازاق ەلى": "hora de Kazakistán", "قىرعىزستان": "hora de Kirguizistán", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "hora de verán do Perú"},
	}
}

// Locale returns the current translators string locale
func (gl *gl_ES) Locale() string {
	return gl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'gl_ES'
func (gl *gl_ES) PluralsCardinal() []locales.PluralRule {
	return gl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'gl_ES'
func (gl *gl_ES) PluralsOrdinal() []locales.PluralRule {
	return gl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'gl_ES'
func (gl *gl_ES) PluralsRange() []locales.PluralRule {
	return gl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'gl_ES'
func (gl *gl_ES) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'gl_ES'
func (gl *gl_ES) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'gl_ES'
func (gl *gl_ES) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := gl.CardinalPluralRule(num1, v1)
	end := gl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (gl *gl_ES) MonthAbbreviated(month time.Month) string {
	return gl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (gl *gl_ES) MonthsAbbreviated() []string {
	return gl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (gl *gl_ES) MonthNarrow(month time.Month) string {
	return gl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (gl *gl_ES) MonthsNarrow() []string {
	return gl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (gl *gl_ES) MonthWide(month time.Month) string {
	return gl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (gl *gl_ES) MonthsWide() []string {
	return gl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (gl *gl_ES) WeekdayAbbreviated(weekday time.Weekday) string {
	return gl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (gl *gl_ES) WeekdaysAbbreviated() []string {
	return gl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (gl *gl_ES) WeekdayNarrow(weekday time.Weekday) string {
	return gl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (gl *gl_ES) WeekdaysNarrow() []string {
	return gl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (gl *gl_ES) WeekdayShort(weekday time.Weekday) string {
	return gl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (gl *gl_ES) WeekdaysShort() []string {
	return gl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (gl *gl_ES) WeekdayWide(weekday time.Weekday) string {
	return gl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (gl *gl_ES) WeekdaysWide() []string {
	return gl.daysWide
}

// Decimal returns the decimal point of number
func (gl *gl_ES) Decimal() string {
	return gl.decimal
}

// Group returns the group of number
func (gl *gl_ES) Group() string {
	return gl.group
}

// Group returns the minus sign of number
func (gl *gl_ES) Minus() string {
	return gl.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'gl_ES' and handles both Whole and Real numbers based on 'v'
func (gl *gl_ES) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'gl_ES' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (gl *gl_ES) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, gl.percentSuffix...)

	b = append(b, gl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'gl_ES'
func (gl *gl_ES) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, gl.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'gl_ES'
// in accounting notation.
func (gl *gl_ES) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gl.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, gl.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, gl.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, gl.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, gl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'gl_ES'
func (gl *gl_ES) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := gl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
