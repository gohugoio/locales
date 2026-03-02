package pt_GQ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type pt_GQ struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'pt_GQ' locale
func New() locales.Translator {
	return &pt_GQ{
		locale:                 "pt_GQ",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jan.", "fev.", "mar.", "abr.", "mai.", "jun.", "jul.", "ago.", "set.", "out.", "nov.", "dez."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
		daysAbbreviated:        []string{"domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado"},
		daysNarrow:             []string{"D", "S", "T", "Q", "Q", "S", "S"},
		daysShort:              []string{"dom.", "seg.", "ter.", "qua.", "qui.", "sex.", "sáb."},
		daysWide:               []string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"},
		timezones:              map[string]string{"ACDT": "Hora de verão da Austrália Central", "ACST": "Hora de verão do Acre", "ACT": "Hora padrão do Acre", "ACWDT": "Hora de verão da Austrália Central Ocidental", "ACWST": "Hora padrão da Austrália Central Ocidental", "ADT": "Hora de verão do Atlântico", "ADT Arabia": "Hora de verão da Arábia", "AEDT": "Hora de verão da Austrália Oriental", "AEST": "Hora padrão da Austrália Oriental", "AFT": "Hora do Afeganistão", "AKDT": "Hora de verão do Alasca", "AKST": "Hora padrão do Alasca", "AMST": "Hora de verão do Amazonas", "AMST Armenia": "Hora de verão da Arménia", "AMT": "Hora padrão do Amazonas", "AMT Armenia": "Hora padrão da Arménia", "ANAST": "Hora de verão de Anadyr", "ANAT": "Hora padrão de Anadyr", "ARST": "Hora de verão da Argentina", "ART": "Hora padrão da Argentina", "AST": "Hora padrão do Atlântico", "AST Arabia": "Hora padrão da Arábia", "AWDT": "Hora de verão da Austrália Ocidental", "AWST": "Hora padrão da Austrália Ocidental", "AZST": "Hora de verão do Azerbaijão", "AZT": "Hora padrão do Azerbaijão", "BDT Bangladesh": "Hora de verão do Bangladeche", "BNT": "Hora do Brunei", "BOT": "Hora da Bolívia", "BRST": "Hora de verão de Brasília", "BRT": "Hora padrão de Brasília", "BST Bangladesh": "Hora padrão do Bangladeche", "BT": "Hora do Butão", "CAST": "CAST", "CAT": "Hora da África Central", "CCT": "Hora das Ilhas Cocos", "CDT": "Hora de verão central norte-americana", "CHADT": "Hora de verão de Chatham", "CHAST": "Hora padrão de Chatham", "CHUT": "Hora de Chuuk", "CKT": "Hora padrão das Ilhas Cook", "CKT DST": "Hora de verão das Ilhas Cook", "CLST": "Hora de verão do Chile", "CLT": "Hora padrão do Chile", "COST": "Hora de verão da Colômbia", "COT": "Hora padrão da Colômbia", "CST": "Hora padrão central norte-americana", "CST China": "Hora padrão da China", "CST China DST": "Hora de verão da China", "CVST": "Hora de verão de Cabo Verde", "CVT": "Hora padrão de Cabo Verde", "CXT": "Hora da Ilha do Natal", "ChST": "Hora padrão de Chamorro", "ChST NMI": "Hora das Ilhas Mariana do Norte", "CuDT": "Hora de verão de Cuba", "CuST": "Hora padrão de Cuba", "DAVT": "Hora de Davis", "DDUT": "Hora de Dumont-d’Urville", "EASST": "Hora de verão da Ilha da Páscoa", "EAST": "Hora padrão da Ilha da Páscoa", "EAT": "Hora da África Oriental", "ECT": "Hora do Equador", "EDT": "Hora de verão oriental norte-americana", "EGDT": "Hora de verão da Gronelândia Oriental", "EGST": "Hora padrão da Gronelândia Oriental", "EST": "Hora padrão oriental norte-americana", "FEET": "Hora do Extremo Leste da Europa", "FJT": "Hora padrão de Fiji", "FJT Summer": "Hora de verão de Fiji", "FKST": "Hora de verão das Ilhas Falkland", "FKT": "Hora padrão das Ilhas Falkland", "FNST": "Hora de verão de Fernando de Noronha", "FNT": "Hora padrão de Fernando de Noronha", "GALT": "Hora das Galápagos", "GAMT": "Hora de Gambier", "GEST": "Hora de verão da Geórgia", "GET": "Hora padrão da Geórgia", "GFT": "Hora da Guiana Francesa", "GIT": "Hora das Ilhas Gilbert", "GMT": "Hora de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Hora padrão do Golfo", "GST Guam": "Hora padrão de Guam", "GYT": "Hora da Guiana", "HADT": "Hora padrão do Havai e Aleutas", "HAST": "Hora padrão do Havai e Aleutas", "HKST": "Hora de verão de Hong Kong", "HKT": "Hora padrão de Hong Kong", "HOVST": "Hora de verão de Hovd", "HOVT": "Hora padrão de Hovd", "ICT": "Hora da Indochina", "IDT": "Hora de verão de Israel", "IOT": "Hora do Oceano Índico", "IRKST": "Hora de verão de Irkutsk", "IRKT": "Hora padrão de Irkutsk", "IRST": "Hora padrão do Irão", "IRST DST": "Hora de verão do Irão", "IST": "Hora padrão da Índia", "IST Israel": "Hora padrão de Israel", "JDT": "Hora de verão do Japão", "JST": "Hora padrão do Japão", "KOST": "Hora de Kosrae", "KRAST": "Hora de verão de Krasnoyarsk", "KRAT": "Hora padrão de Krasnoyarsk", "KST": "Hora padrão da Coreia", "KST DST": "Hora de verão da Coreia", "LHDT": "Hora de verão de Lord Howe", "LHST": "Hora padrão de Lord Howe", "LINT": "Hora das Ilhas Line", "MAGST": "Hora de verão de Magadan", "MAGT": "Hora padrão de Magadan", "MART": "Hora das Ilhas Marquesas", "MAWT": "Hora de Mawson", "MDT": "Hora de verão de montanha norte-americana", "MESZ": "Hora de verão da Europa Central", "MEZ": "Hora padrão da Europa Central", "MHT": "Hora das Ilhas Marshall", "MMT": "Hora de Mianmar", "MSD": "Hora de verão de Moscovo", "MST": "Hora padrão de montanha norte-americana", "MUST": "Hora de verão da Maurícia", "MUT": "Hora padrão da Maurícia", "MVT": "Hora das Maldivas", "MYT": "Hora da Malásia", "NCT": "Hora padrão da Nova Caledónia", "NDT": "Hora de verão da Terra Nova", "NDT New Caledonia": "Hora de verão da Nova Caledónia", "NFDT": "Hora de verão da Ilha Norfolk", "NFT": "Hora padrão da Ilha Norfolk", "NOVST": "Hora de verão de Novosibirsk", "NOVT": "Hora padrão de Novosibirsk", "NPT": "Hora do Nepal", "NRT": "Hora de Nauru", "NST": "Hora padrão da Terra Nova", "NUT": "Hora de Niuê", "NZDT": "Hora de verão da Nova Zelândia", "NZST": "Hora padrão da Nova Zelândia", "OESZ": "Hora de verão da Europa Oriental", "OEZ": "Hora padrão da Europa Oriental", "OMSST": "Hora de verão de Omsk", "OMST": "Hora padrão de Omsk", "PDT": "Hora de verão do Pacífico norte-americana", "PDTM": "Hora de verão do Pacífico Mexicano", "PETDT": "Hora de verão de Petropavlovsk-Kamchatski", "PETST": "Hora padrão de Petropavlovsk-Kamchatski", "PGT": "Hora de Papua Nova Guiné", "PHOT": "Hora das Ilhas Fénix", "PKT": "Hora padrão do Paquistão", "PKT DST": "Hora de verão do Paquistão", "PMDT": "Hora de verão de São Pedro e Miquelão", "PMST": "Hora padrão de São Pedro e Miquelão", "PONT": "Hora de Pohnpei", "PST": "Hora padrão do Pacífico norte-americana", "PST Philippine": "Hora padrão das Filipinas", "PST Philippine DST": "Hora de verão das Filipinas", "PST Pitcairn": "Hora de Pitcairn", "PSTM": "Hora padrão do Pacífico Mexicano", "PWT": "Hora de Palau", "PYST": "Hora de verão do Paraguai", "PYT": "Hora padrão do Paraguai", "PYT Korea": "Hora da Coreia do Norte", "RET": "Hora de Reunião", "ROTT": "Hora de Rothera", "SAKST": "Hora de verão de Sacalina", "SAKT": "Hora padrão de Sacalina", "SAMST": "Hora de verão de Samara", "SAMT": "Hora padrão de Samara", "SAST": "Hora da África do Sul", "SBT": "Hora das Ilhas Salomão", "SCT": "Hora das Seicheles", "SGT": "Hora padrão de Singapura", "SLST": "Hora do Sri Lanka", "SRT": "Hora do Suriname", "SST Samoa": "Hora padrão de Samoa Americana", "SST Samoa Apia": "Hora padrão de Samoa", "SST Samoa Apia DST": "Hora de verão de Samoa", "SST Samoa DST": "Hora de verão de Samoa Americana", "SYOT": "Hora de Syowa", "TAAF": "Hora das Terras Austrais e Antárcticas Francesas", "TAHT": "Hora do Taiti", "TJT": "Hora do Tajiquistão", "TKT": "Hora de Tokelau", "TLT": "Hora de Timor Leste", "TMST": "Hora de verão do Turquemenistão", "TMT": "Hora padrão do Turquemenistão", "TOST": "Hora de verão de Tonga", "TOT": "Hora padrão de Tonga", "TVT": "Hora de Tuvalu", "TWT": "Hora padrão de Taiwan", "TWT DST": "Hora de verão de Taiwan", "ULAST": "Hora de verão de Ulan Bator", "ULAT": "Hora padrão de Ulan Bator", "UYST": "Hora de verão do Uruguai", "UYT": "Hora padrão do Uruguai", "UZT": "Hora padrão do Uzbequistão", "UZT DST": "Hora de verão do Uzbequistão", "VET": "Hora da Venezuela", "VLAST": "Hora de verão de Vladivostok", "VLAT": "Hora padrão de Vladivostok", "VOLST": "Hora de verão de Volgogrado", "VOLT": "Hora padrão de Volgogrado", "VOST": "Hora de Vostok", "VUT": "Hora padrão do Vanuatu", "VUT DST": "Hora de verão do Vanuatu", "WAKT": "Hora da Ilha Wake", "WARST": "Hora de verão da Argentina Ocidental", "WART": "Hora padrão da Argentina Ocidental", "WAST": "Hora da África Ocidental", "WAT": "Hora da África Ocidental", "WESZ": "Hora de verão da Europa Ocidental", "WEZ": "Hora padrão da Europa Ocidental", "WFT": "Hora de Wallis e Futuna", "WGST": "Hora de verão da Gronelândia Ocidental", "WGT": "Hora padrão da Gronelândia Ocidental", "WIB": "Hora da Indonésia Ocidental", "WIT": "Hora da Indonésia Oriental", "WITA": "Hora da Indonésia Central", "YAKST": "Hora de verão de Yakutsk", "YAKT": "Hora padrão de Yakutsk", "YEKST": "Hora de verão de Ecaterimburgo", "YEKT": "Hora padrão de Ecaterimburgo", "YST": "Hora do Yukon", "МСК": "Hora padrão de Moscovo", "اقتاۋ": "Hora padrão de Aqtau", "اقتاۋ قالاسى": "Hora de verão de Aqtau", "اقتوبە": "Hora padrão de Aqtobe", "اقتوبە قالاسى": "Hora de verão de Aqtobe", "الماتى": "Hora padrão de Almaty", "الماتى قالاسى": "Hora de verão de Almaty", "باتىس قازاق ەلى": "Hora do Cazaquistão Ocidental", "شىعىش قازاق ەلى": "Hora do Cazaquistão Oriental", "قازاق ەلى": "Hora do Cazaquistão", "قىرعىزستان": "Hora do Quirguistão", "قىزىلوردا": "Hora padrão de Qyzylorda", "قىزىلوردا قالاسى": "Hora de verão de Qyzylorda", "∅∅∅": "Hora de verão dos Açores"},
	}
}

// Locale returns the current translators string locale
func (pt *pt_GQ) Locale() string {
	return pt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pt_GQ'
func (pt *pt_GQ) PluralsCardinal() []locales.PluralRule {
	return pt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pt_GQ'
func (pt *pt_GQ) PluralsOrdinal() []locales.PluralRule {
	return pt.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pt_GQ'
func (pt *pt_GQ) PluralsRange() []locales.PluralRule {
	return pt.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pt_GQ'
func (pt *pt_GQ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pt_GQ'
func (pt *pt_GQ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pt_GQ'
func (pt *pt_GQ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
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
func (pt *pt_GQ) MonthAbbreviated(month time.Month) string {
	return pt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pt *pt_GQ) MonthsAbbreviated() []string {
	return pt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pt *pt_GQ) MonthNarrow(month time.Month) string {
	return pt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pt *pt_GQ) MonthsNarrow() []string {
	return pt.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pt *pt_GQ) MonthWide(month time.Month) string {
	return pt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pt *pt_GQ) MonthsWide() []string {
	return pt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pt *pt_GQ) WeekdayAbbreviated(weekday time.Weekday) string {
	return pt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pt *pt_GQ) WeekdaysAbbreviated() []string {
	return pt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pt *pt_GQ) WeekdayNarrow(weekday time.Weekday) string {
	return pt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pt *pt_GQ) WeekdaysNarrow() []string {
	return pt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pt *pt_GQ) WeekdayShort(weekday time.Weekday) string {
	return pt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pt *pt_GQ) WeekdaysShort() []string {
	return pt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pt *pt_GQ) WeekdayWide(weekday time.Weekday) string {
	return pt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pt *pt_GQ) WeekdaysWide() []string {
	return pt.daysWide
}

// Decimal returns the decimal point of number
func (pt *pt_GQ) Decimal() string {
	return pt.decimal
}

// Group returns the group of number
func (pt *pt_GQ) Group() string {
	return pt.group
}

// Group returns the minus sign of number
func (pt *pt_GQ) Minus() string {
	return pt.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pt_GQ' and handles both Whole and Real numbers based on 'v'
func (pt *pt_GQ) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
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
				for j := len(pt.group) - 1; j >= 0; j-- {
					b = append(b, pt.group[j])
				}
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

// FmtPercent returns 'num' with digits/precision of 'v' for 'pt_GQ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pt *pt_GQ) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pt_GQ'
func (pt *pt_GQ) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pt.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
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
				for j := len(pt.group) - 1; j >= 0; j-- {
					b = append(b, pt.group[j])
				}
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

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, pt.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pt_GQ'
// in accounting notation.
func (pt *pt_GQ) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pt.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
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
				for j := len(pt.group) - 1; j >= 0; j-- {
					b = append(b, pt.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pt.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, pt.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, pt.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtDateMedium(t time.Time) string {
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

// FmtDateLong returns the long date representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtDateLong(t time.Time) string {
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

// FmtDateFull returns the full date representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtDateFull(t time.Time) string {
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

// FmtTimeShort returns the short time representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtTimeShort(t time.Time) string {
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

// FmtTimeMedium returns the medium time representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtTimeMedium(t time.Time) string {
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

// FmtTimeLong returns the long time representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtTimeLong(t time.Time) string {
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

// FmtTimeFull returns the full time representation of 't' for 'pt_GQ'
func (pt *pt_GQ) FmtTimeFull(t time.Time) string {
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
