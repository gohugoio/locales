package pt

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type pt struct {
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

// New returns a new instance of translator for the 'pt' locale
func New() locales.Translator {
	return &pt{
		locale:                 "pt",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "Esc.", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "S£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: " mil",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: " mil",
		monthsAbbreviated:      []string{"", "jan.", "fev.", "mar.", "abr.", "mai.", "jun.", "jul.", "ago.", "set.", "out.", "nov.", "dez."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
		daysAbbreviated:        []string{"dom.", "seg.", "ter.", "qua.", "qui.", "sex.", "sáb."},
		daysNarrow:             []string{"D", "S", "T", "Q", "Q", "S", "S"},
		daysWide:               []string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"},
		periodsAbbreviated:     []string{"", ""},
		erasAbbreviated:        []string{"a.C.", "d.C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"antes de Cristo", "depois de Cristo"},
		timezones:              map[string]string{"ACDT": "Horário de Verão da Austrália Central", "ACST": "Horário de Verão do Acre", "ACT": "Horário Padrão do Acre", "ACWDT": "Horário de Verão da Austrália Centro-Ocidental", "ACWST": "Horário Padrão da Austrália Centro-Ocidental", "ADT": "Horário de Verão do Atlântico", "ADT Arabia": "Horário de Verão da Arábia", "AEDT": "Horário de Verão da Austrália Oriental", "AEST": "Horário Padrão da Austrália Oriental", "AFT": "Horário do Afeganistão", "AKDT": "Horário de Verão do Alasca", "AKST": "Horário Padrão do Alasca", "AMST": "Horário de Verão do Amazonas", "AMST Armenia": "Horário de Verão da Armênia", "AMT": "Horário Padrão do Amazonas", "AMT Armenia": "Horário Padrão da Armênia", "ANAST": "Horário de Verão do Anadyr", "ANAT": "Horário Padrão do Anadyr", "ARST": "Horário de Verão da Argentina", "ART": "Horário Padrão da Argentina", "AST": "Horário Padrão do Atlântico", "AST Arabia": "Horário Padrão da Arábia", "AWDT": "Horário de Verão da Austrália Ocidental", "AWST": "Horário Padrão da Austrália Ocidental", "AZST": "Horário de Verão do Arzeibaijão", "AZT": "Horário Padrão do Arzeibaijão", "BDT Bangladesh": "Horário de Verão de Bangladesh", "BNT": "Horário de Brunei Darussalam", "BOT": "Horário da Bolívia", "BRST": "Horário de Verão de Brasília", "BRT": "Horário Padrão de Brasília", "BST Bangladesh": "Horário Padrão de Bangladesh", "BT": "Horário do Butão", "CAST": "CAST", "CAT": "Horário da África Central", "CCT": "Horário das Ilhas Coco", "CDT": "Horário de Verão Central", "CHADT": "Horário de Verão de Chatham", "CHAST": "Horário Padrão de Chatham", "CHUT": "Horário de Chuuk", "CKT": "Horário Padrão das Ilhas Cook", "CKT DST": "Meio Horário de Verão das Ilhas Cook", "CLST": "Horário de Verão do Chile", "CLT": "Horário Padrão do Chile", "COST": "Horário de Verão da Colômbia", "COT": "Horário Padrão da Colômbia", "CST": "Horário Padrão Central", "CST China": "Horário Padrão da China", "CST China DST": "Horário de Verão da China", "CVST": "Horário de Verão de Cabo Verde", "CVT": "Horário Padrão de Cabo Verde", "CXT": "Horário da Ilha Christmas", "ChST": "Horário de Chamorro", "ChST NMI": "Horário das Ilhas Mariana do Norte", "CuDT": "Horário de Verão de Cuba", "CuST": "Horário Padrão de Cuba", "DAVT": "Horário de Davis", "DDUT": "Horário de Dumont-d’Urville", "EASST": "Horário de Verão da Ilha de Páscoa", "EAST": "Horário Padrão da Ilha de Páscoa", "EAT": "Horário da África Oriental", "ECT": "Horário do Equador", "EDT": "Horário de Verão do Leste", "EGDT": "Horário de Verão da Groelândia Oriental", "EGST": "Horário Padrão da Groelândia Oriental", "EST": "Horário Padrão do Leste", "FEET": "Horário do Extremo Leste Europeu", "FJT": "Horário Padrão de Fiji", "FJT Summer": "Horário de Verão de Fiji", "FKST": "Horário de Verão das Ilhas Malvinas", "FKT": "Horário Padrão das Ilhas Malvinas", "FNST": "Horário de Verão de Fernando de Noronha", "FNT": "Horário Padrão de Fernando de Noronha", "GALT": "Horário de Galápagos", "GAMT": "Horário de Gambier", "GEST": "Horário de Verão da Geórgia", "GET": "Horário Padrão da Geórgia", "GFT": "Horário da Guiana Francesa", "GIT": "Horário das Ilhas Gilberto", "GMT": "Horário do Meridiano de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Horário do Golfo", "GST Guam": "Horário Padrão de Guam", "GYT": "Horário da Guiana", "HADT": "Horário de Verão do Havaí e Ilhas Aleutas", "HAST": "Horário Padrão do Havaí e Ilhas Aleutas", "HKST": "Horário de Verão de Hong Kong", "HKT": "Horário Padrão de Hong Kong", "HOVST": "Horário de Verão de Hovd", "HOVT": "Horário Padrão de Hovd", "ICT": "Horário da Indochina", "IDT": "Horário de Verão de Israel", "IOT": "Horário do Oceano Índico", "IRKST": "Horário de Verão de Irkutsk", "IRKT": "Horário Padrão de Irkutsk", "IRST": "Horário Padrão do Irã", "IRST DST": "Horário de Verão do Irã", "IST": "Horário Padrão da Índia", "IST Israel": "Horário Padrão de Israel", "JDT": "Horário de Verão do Japão", "JST": "Horário Padrão do Japão", "KOST": "Horário de Kosrae", "KRAST": "Horário de Verão de Krasnoyarsk", "KRAT": "Horário Padrão de Krasnoyarsk", "KST": "Horário Padrão da Coreia", "KST DST": "Horário de Verão da Coreia", "LHDT": "Horário de Verão de Lord Howe", "LHST": "Horário Padrão de Lord Howe", "LINT": "Horário das Ilhas da Linha", "MAGST": "Horário de Verão de Magadan", "MAGT": "Horário Padrão de Magadan", "MART": "Horário das Marquesas", "MAWT": "Horário de Mawson", "MDT": "Horário de Verão das Montanhas", "MESZ": "Horário de Verão da Europa Central", "MEZ": "Horário Padrão da Europa Central", "MHT": "Horário das Ilhas Marshall", "MMT": "Horário de Mianmar", "MSD": "Horário de Verão de Moscou", "MST": "Horário Padrão das Montanhas", "MUST": "Horário de Verão de Maurício", "MUT": "Horário Padrão de Maurício", "MVT": "Horário das Ilhas Maldivas", "MYT": "Horário da Malásia", "NCT": "Horário Padrão da Nova Caledônia", "NDT": "Horário de Verão da Terra Nova", "NDT New Caledonia": "Horário de Verão da Nova Caledônia", "NFDT": "Horário de Verão da Ilha Norfolk", "NFT": "Horário Padrão da Ilha Norfolk", "NOVST": "Horário de Verão de Novosibirsk", "NOVT": "Horário Padrão de Novosibirsk", "NPT": "Horário do Nepal", "NRT": "Horário de Nauru", "NST": "Horário Padrão da Terra Nova", "NUT": "Horário de Niue", "NZDT": "Horário de Verão da Nova Zelândia", "NZST": "Horário Padrão da Nova Zelândia", "OESZ": "Horário de Verão da Europa Oriental", "OEZ": "Horário Padrão da Europa Oriental", "OMSST": "Horário de Verão de Omsk", "OMST": "Horário Padrão de Omsk", "PDT": "Horário de Verão do Pacífico", "PDTM": "Horário de Verão do Pacífico Mexicano", "PETDT": "Horário de Verão de Petropavlovsk-Kamchatski", "PETST": "Horário Padrão de Petropavlovsk-Kamchatski", "PGT": "Horário de Papua-Nova Guiné", "PHOT": "Horário das Ilhas Fênix", "PKT": "Horário Padrão do Paquistão", "PKT DST": "Horário de Verão do Paquistão", "PMDT": "Horário Verão de São Pedro e Miquelão", "PMST": "Horário Padrão de São Pedro e Miquelão", "PONT": "Horário de Ponape", "PST": "Horário Padrão do Pacífico", "PST Philippine": "Horário Padrão das Filipinas", "PST Philippine DST": "Horário de Verão das Filipinas", "PST Pitcairn": "Horário de Pitcairn", "PSTM": "Horário Padrão do Pacífico Mexicano", "PWT": "Horário de Palau", "PYST": "Horário de Verão do Paraguai", "PYT": "Horário Padrão do Paraguai", "PYT Korea": "Horário de Pyongyang", "RET": "Horário de Reunião", "ROTT": "Horário de Rothera", "SAKST": "Horário de Verão de Sacalina", "SAKT": "Horário Padrão de Sacalina", "SAMST": "Horário de Verão de Samara", "SAMT": "Horário Padrão de Samara", "SAST": "Horário da África do Sul", "SBT": "Horário das Ilhas Salomão", "SCT": "Horário de Seicheles", "SGT": "Horário Padrão de Singapura", "SLST": "Horário de Lanka", "SRT": "Horário do Suriname", "SST Samoa": "Horário Padrão de Samoa", "SST Samoa Apia": "Horário Padrão de Apia", "SST Samoa Apia DST": "Horário de Verão de Apia", "SST Samoa DST": "Horário de Verão de Samoa", "SYOT": "Horário de Syowa", "TAAF": "Horário dos Territórios Franceses do Sul e Antártida", "TAHT": "Horário do Taiti", "TJT": "Horário do Tajiquistão", "TKT": "Horário de Tokelau", "TLT": "Horário do Timor-Leste", "TMST": "Horário de Verão do Turcomenistão", "TMT": "Horário Padrão do Turcomenistão", "TOST": "Horário de Verão de Tonga", "TOT": "Horário Padrão de Tonga", "TVT": "Horário de Tuvalu", "TWT": "Horário Padrão de Taipei", "TWT DST": "Horário de Verão de Taipei", "ULAST": "Horário de Verão de Ulan Bator", "ULAT": "Horário Padrão de Ulan Bator", "UYST": "Horário de Verão do Uruguai", "UYT": "Horário Padrão do Uruguai", "UZT": "Horário Padrão do Uzbequistão", "UZT DST": "Horário de Verão do Uzbequistão", "VET": "Horário da Venezuela", "VLAST": "Horário de Verão de Vladivostok", "VLAT": "Horário Padrão de Vladivostok", "VOLST": "Horário de Verão de Volgogrado", "VOLT": "Horário Padrão de Volgogrado", "VOST": "Horário de Vostok", "VUT": "Horário Padrão de Vanuatu", "VUT DST": "Horário de Verão de Vanuatu", "WAKT": "Horário das Ilhas Wake", "WARST": "Horário de Verão da Argentina Ocidental", "WART": "Horário Padrão da Argentina Ocidental", "WAST": "Horário da África Ocidental", "WAT": "Horário da África Ocidental", "WESZ": "Horário de Verão da Europa Ocidental", "WEZ": "Horário Padrão da Europa Ocidental", "WFT": "Horário de Wallis e Futuna", "WGST": "Horário de Verão da Groenlândia Ocidental", "WGT": "Horário Padrão da Groenlândia Ocidental", "WIB": "Horário da Indonésia Ocidental", "WIT": "Horário da Indonésia Oriental", "WITA": "Horário da Indonésia Central", "YAKST": "Horário de Verão de Yakutsk", "YAKT": "Horário Padrão de Yakutsk", "YEKST": "Horário de Verão de Ecaterimburgo", "YEKT": "Horário Padrão de Ecaterimburgo", "YST": "Horário do Yukon", "МСК": "Horário Padrão de Moscou", "اقتاۋ": "Horário Padrão do Aqtau", "اقتاۋ قالاسى": "Horário de Verão do Aqtau", "اقتوبە": "Horário Padrão do Aqtobe", "اقتوبە قالاسى": "Horário de Verão do Aqtobe", "الماتى": "Horário Padrão do Almaty", "الماتى قالاسى": "Horário de Verão do Almaty", "باتىس قازاق ەلى": "Horário do Cazaquistão Ocidental", "شىعىش قازاق ەلى": "Horário do Cazaquistão Oriental", "قازاق ەلى": "Horário do Cazaquistão", "قىرعىزستان": "Horário do Quirguistão", "قىزىلوردا": "Horário Padrão de Qyzylorda", "قىزىلوردا قالاسى": "Horário de Verão de Qyzylorda", "∅∅∅": "Horário de Verão dos Açores"},
	}
}

// Locale returns the current translators string locale
func (pt *pt) Locale() string {
	return pt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pt'
func (pt *pt) PluralsCardinal() []locales.PluralRule {
	return pt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pt'
func (pt *pt) PluralsOrdinal() []locales.PluralRule {
	return pt.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pt'
func (pt *pt) PluralsRange() []locales.PluralRule {
	return pt.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pt'
func (pt *pt) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	e := int64(0)
	iMod1000000 := i % 1000000

	if i >= 0 && i <= 1 {
		return locales.PluralRuleOne
	} else if (e == 0 && i != 0 && iMod1000000 == 0 && v == 0) || (e < 0 || e > 5) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pt'
func (pt *pt) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pt'
func (pt *pt) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := pt.CardinalPluralRule(num1, v1)
	end := pt.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pt *pt) MonthAbbreviated(month time.Month) string {
	return pt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pt *pt) MonthsAbbreviated() []string {
	return pt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pt *pt) MonthNarrow(month time.Month) string {
	return pt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pt *pt) MonthsNarrow() []string {
	return pt.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pt *pt) MonthWide(month time.Month) string {
	return pt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pt *pt) MonthsWide() []string {
	return pt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pt *pt) WeekdayAbbreviated(weekday time.Weekday) string {
	return pt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pt *pt) WeekdaysAbbreviated() []string {
	return pt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pt *pt) WeekdayNarrow(weekday time.Weekday) string {
	return pt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pt *pt) WeekdaysNarrow() []string {
	return pt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pt *pt) WeekdayShort(weekday time.Weekday) string {
	return pt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pt *pt) WeekdaysShort() []string {
	return pt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pt *pt) WeekdayWide(weekday time.Weekday) string {
	return pt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pt *pt) WeekdaysWide() []string {
	return pt.daysWide
}

// Decimal returns the decimal point of number
func (pt *pt) Decimal() string {
	return pt.decimal
}

// Group returns the group of number
func (pt *pt) Group() string {
	return pt.group
}

// Group returns the minus sign of number
func (pt *pt) Minus() string {
	return pt.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pt' and handles both Whole and Real numbers based on 'v'
func (pt *pt) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pt.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pt' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pt *pt) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pt.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pt'
func (pt *pt) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pt.currencies[currency]
	l := len(s) + len(symbol) + 9

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(pt.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, pt.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, pt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pt.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pt'
// in accounting notation.
func (pt *pt) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pt.currencies[currency]
	l := len(s) + len(symbol) + 9

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(pt.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, pt.currencyNegativePrefix[j])
		}

		b = append(b, pt.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(pt.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, pt.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, pt.currencyNegativeSuffix...)
	} else {

		b = append(b, pt.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pt'
func (pt *pt) FmtDateShort(t time.Time) string {

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

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'pt'
func (pt *pt) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, pt.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pt'
func (pt *pt) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, pt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pt'
func (pt *pt) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pt.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, pt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pt'
func (pt *pt) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pt'
func (pt *pt) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pt'
func (pt *pt) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pt'
func (pt *pt) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
