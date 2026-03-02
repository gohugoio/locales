package ce

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ce struct {
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

// New returns a new instance of translator for the 'ce' locale
func New() locales.Translator {
	return &ce{
		locale:                 "ce",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "лей", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
		monthsNarrow:           []string{"", "Я", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
		daysAbbreviated:        []string{"кӀи", "ор", "ши", "кха", "еа", "пӀе", "шуо"},
		daysNarrow:             []string{"кӀи", "ор", "ши", "кха", "еа", "пӀе", "шуо"},
		daysWide:               []string{"кӀира", "оршот", "шинара", "кхаара", "еара", "пӀераска", "шуот"},
		timezones:              map[string]string{"ACDT": "Юккъера Австрали, аьхкенан хан", "ACST": "Юккъера Австрали, стандартан хан", "ACT": "ACT", "ACWDT": "Юккъера Австрали, малхбузен аьхкенан хан", "ACWST": "Юккъера Австрали, малхбузен стандартан хан", "ADT": "Атлантикан аьхкенан хан", "ADT Arabia": "СаӀудийн Ӏаьрбийчоьнан, аьхкенан хан", "AEDT": "Малхбален Австрали, аьхкенан хан", "AEST": "Малхбален Австрали, стандартан хан", "AFT": "ОвхӀан", "AKDT": "Аляска, аьхкенан хан", "AKST": "Аляска, стандартан хан", "AMST": "Амазонка, аьхкенан хан", "AMST Armenia": "Эрмалойчоь, аьхкенан хан", "AMT": "Амазонка, стандартан хан", "AMT Armenia": "Эрмалойчоь, стандартан хан", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентина, аьхкенан хан", "ART": "Аргентина, стандартан хан", "AST": "Атлантикан стандартан хан", "AST Arabia": "СаӀудийн Ӏаьрбийчоьнан стандартан хан", "AWDT": "Малхбузен Австрали, аьхкенан хан", "AWST": "Малхбузен Австрали, стандартан хан", "AZST": "Азербайджан, аьхкенан хан", "AZT": "Азербайджан, стандартан хан", "BDT Bangladesh": "Бангладеш, аьхкенан хан", "BNT": "Бруней-Даруссалам", "BOT": "Боливи", "BRST": "Бразили, аьхкенан хан", "BRT": "Бразили, стандартан хан", "BST Bangladesh": "Бангладеш, стандартан хан", "BT": "Бутан", "CAST": "CAST", "CAT": "Юккъера Африка", "CCT": "Кокосийн, гӀ-наш", "CDT": "Юккъера Америка, аьхкенан хан", "CHADT": "Чатем, аьхкенан хан", "CHAST": "Чатем, стандартан хан", "CHUT": "Чуук", "CKT": "Кукан, гӀ-наш, стандартан хан", "CKT DST": "Кукан, гӀ-наш, аьхкенан хан", "CLST": "Чили, аьхкенан хан", "CLT": "Чили, стандартан хан", "COST": "Колумби, аьхкенан хан", "COT": "Колумби, стандартан хан", "CST": "Юккъера Америка, стандартан хан", "CST China": "Цийчоьнан, стандартан хан", "CST China DST": "Цийчоьнан, аьхкенан хан", "CVST": "Кабо-Верде, аьхкенан хан", "CVT": "Кабо-Верде, стандартан хан", "CXT": "Ӏийса пайхамар винчу ден гӀайре", "ChST": "Чаморро", "ChST NMI": "ChST NMI", "CuDT": "Куба, аьхкенан хан", "CuST": "Куба, стандартан хан", "DAVT": "Дейвис", "DDUT": "Дюмон-д’Юрвиль", "EASST": "Мархин гӀайре, аьхкенан хан", "EAST": "Мархин гӀайре, стандартан хан", "EAT": "Малхбален Африка", "ECT": "Эквадор", "EDT": "Малхбален Америка, аьхкенан хан", "EGDT": "Малхбален Гренланди, аьхкенан хан", "EGST": "Малхбален Гренланди, стандартан хан", "EST": "Малхбален Америка, стандартан хан", "FEET": "Калининградера а, Минскера а хан", "FJT": "Фиджи, стандартан хан", "FJT Summer": "Фиджи, аьхкенан хан", "FKST": "Фолклендан гӀайренаш, аьхкенан хан", "FKT": "Фолклендан гӀайренаш, стандартан хан", "FNST": "Фернанду-ди-Норонья, аьхкенан хан", "FNT": "Фернанду-ди-Норонья, стандартан хан", "GALT": "Галапагосан гӀайренаш", "GAMT": "Гамбье", "GEST": "Гуьржийчоь, аьхкенан хан", "GET": "Гуьржийчоь, стандартан хан", "GFT": "Французийн Гвиана", "GIT": "Гилбертан, гӀ-наш", "GMT": "Гринвичица юкъара хан", "GNSST": "GNSST", "GNST": "GNST", "GST": "ГӀажарийн айма", "GST Guam": "GST Guam", "GYT": "Гайана", "HADT": "Гавайн-алеутийн стандартан хан", "HAST": "Гавайн-алеутийн стандартан хан", "HKST": "Гонконг, аьхкенан хан", "HKT": "Гонконг, стандартан хан", "HOVST": "Ховд, аьхкенан хан", "HOVT": "Ховд, стандартан хан", "ICT": "Индокитай", "IDT": "Израиль, аьхкенан хан", "IOT": "Индин океан", "IRKST": "Иркутск, аьхкенан хан", "IRKT": "Иркутск, стандартан хан", "IRST": "ГӀажарийчоь, стандартан хан", "IRST DST": "ГӀажарийчоь, аьхкенан хан", "IST": "ХӀинди", "IST Israel": "Израиль, стандартан хан", "JDT": "Япон, аьхкенан хан", "JST": "Япон, стандартан хан", "KOST": "Косраэ", "KRAST": "Красноярск, аьхкенан хан", "KRAT": "Красноярск, стандартан хан", "KST": "Корей, стандартан хан", "KST DST": "Корей, аьхкенан хан", "LHDT": "Лорд-Хау, аьхкенан хан", "LHST": "Лорд-Хау, стандартан хан", "LINT": "Лайн, гӀ-наш", "MAGST": "Магадан, аьхкенан хан", "MAGT": "Магадан, стандартан хан", "MART": "Маркизан, гӀ-наш", "MAWT": "Моусон", "MDT": "MDT", "MESZ": "Юккъера Европа, аьхкенан хан", "MEZ": "Юккъера Европа, стандартан хан", "MHT": "Маршалан , гӀ-наш", "MMT": "Мьянма", "MSD": "Москва, аьхкенан хан", "MST": "MST", "MUST": "Маврики, аьхкенан хан", "MUT": "Маврики, стандартан хан", "MVT": "Мальдиваш", "MYT": "Малайзи", "NCT": "Керла Каледони, стандартан хан", "NDT": "Ньюфаундленд, аьхкенан хан", "NDT New Caledonia": "Керла Каледони, аьхкенан хан", "NFDT": "Норфолк, аьхкенан хан", "NFT": "Норфолк, стандартан хан", "NOVST": "Новосибирск, аьхкенан хан", "NOVT": "Новосибирск, стандартан хан", "NPT": "Непал", "NRT": "Науру", "NST": "Ньюфаундленд, стандартан хан", "NUT": "Ниуэ", "NZDT": "Керла Зеланди, аьхкенан хан", "NZST": "Керла Зеланди, стандартан хан", "OESZ": "Малхбален Европа, аьхкенан хан", "OEZ": "Малхбален Европа, стандартан хан", "OMSST": "Омск, аьхкенан хан", "OMST": "Омск, стандартан хан", "PDT": "Тийна океанан аьхкенан хан", "PDTM": "Тийна океанан Мексикан аьхкенан хан", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Папуа – Керла Гвиней", "PHOT": "Феникс, гӀ-наш", "PKT": "Пакистан, стандартан хан", "PKT DST": "Пакистан, аьхкенан хан", "PMDT": "Сен-Пьер а, Микелон а, аьхкенан хан", "PMST": "Сен-Пьер а, Микелон а, стандартан хан", "PONT": "Понапе, гӀ-наш", "PST": "Тийна океанан стандартан хан", "PST Philippine": "Филиппинаш, стандартан хан", "PST Philippine DST": "Филиппинаш, аьхкенан хан", "PST Pitcairn": "Питкэрн", "PSTM": "Тийна океанан Мексикан стандартан хан", "PWT": "Палау", "PYST": "Парагвай, аьхкенан хан", "PYT": "Парагвай, стандартан хан", "PYT Korea": "Пхеньян", "RET": "Реюньон", "ROTT": "Ротера", "SAKST": "Сахалин, аьхкенан хан", "SAKT": "Сахалин, стандартан хан", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Къилба Африка", "SBT": "Соломонан, гӀ-наш", "SCT": "Сейшелан гӀайренаш", "SGT": "Сингапур", "SLST": "SLST", "SRT": "Суринам", "SST Samoa": "Самоа, стандартан хан", "SST Samoa Apia": "стандартан хан Апиа, Самоа", "SST Samoa Apia DST": "аьхкенан хан Апиа, Самоа", "SST Samoa DST": "Самоа, аьхкенан хан", "SYOT": "Сёва", "TAAF": "Французийн къилба а, Антарктидан а хан", "TAHT": "Таити, гӀ-наш", "TJT": "Таджикистан", "TKT": "Токелау", "TLT": "Малхбален Тимор", "TMST": "Туркменин аьхкенан хан", "TMT": "Туркменин стандартан хан", "TOST": "Тонга, аьхкенан хан", "TOT": "Тонга, стандартан хан", "TVT": "Тувалу", "TWT": "Тайвань, стандартан хан", "TWT DST": "Тайвань, аьхкенан хан", "ULAST": "Улан-Батор, аьхкенан хан", "ULAT": "Улан-Батор, стандартан хан", "UYST": "Уругвай, аьхкенан хан", "UYT": "Уругвай, стандартан хан", "UZT": "Узбекистанан стандартан хан", "UZT DST": "Узбекистанан аьхкенан хан", "VET": "Венесуэла", "VLAST": "Владивосток, аьхкенан хан", "VLAT": "Владивосток, стандартан хан", "VOLST": "Волгоград, аьхкенан хан", "VOLT": "Волгоград, стандартан хан", "VOST": "Восток", "VUT": "Вануату, стандартан хан", "VUT DST": "Вануату, аьхкенан хан", "WAKT": "Уэйк, гӀ-е", "WARST": "Малхбузен Аргентина, аьхкенан хан", "WART": "Малхбузен Аргентина, стандартан хан", "WAST": "Малхбузен Африка", "WAT": "Малхбузен Африка", "WESZ": "Малхбузен Европа, аьхкенан хан", "WEZ": "Малхбузен Европа, стандартан хан", "WFT": "Уоллис а, Футуна а", "WGST": "Малхбузен Гренланди, аьхкенан хан", "WGT": "Малхбузен Гренланди, стандартан хан", "WIB": "Малхбузен Индонези", "WIT": "Малхбален Индонези", "WITA": "Юккъера Индонези", "YAKST": "Якутск, аьхкенан хан", "YAKT": "Якутск, стандартан хан", "YEKST": "Екатеринбург, аьхкенан хан", "YEKT": "Екатеринбург, стандартан хан", "YST": "YST", "МСК": "Москва, стандартан хан", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Малхбузен Казахстан", "شىعىش قازاق ەلى": "Малхбален Казахстан", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Киргизи", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Перу, аьхкенан хан"},
	}
}

// Locale returns the current translators string locale
func (ce *ce) Locale() string {
	return ce.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ce'
func (ce *ce) PluralsCardinal() []locales.PluralRule {
	return ce.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ce'
func (ce *ce) PluralsOrdinal() []locales.PluralRule {
	return ce.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ce'
func (ce *ce) PluralsRange() []locales.PluralRule {
	return ce.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ce'
func (ce *ce) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ce'
func (ce *ce) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ce'
func (ce *ce) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ce *ce) MonthAbbreviated(month time.Month) string {
	return ce.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ce *ce) MonthsAbbreviated() []string {
	return ce.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ce *ce) MonthNarrow(month time.Month) string {
	return ce.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ce *ce) MonthsNarrow() []string {
	return ce.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ce *ce) MonthWide(month time.Month) string {
	return ce.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ce *ce) MonthsWide() []string {
	return ce.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ce *ce) WeekdayAbbreviated(weekday time.Weekday) string {
	return ce.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ce *ce) WeekdaysAbbreviated() []string {
	return ce.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ce *ce) WeekdayNarrow(weekday time.Weekday) string {
	return ce.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ce *ce) WeekdaysNarrow() []string {
	return ce.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ce *ce) WeekdayShort(weekday time.Weekday) string {
	return ce.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ce *ce) WeekdaysShort() []string {
	return ce.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ce *ce) WeekdayWide(weekday time.Weekday) string {
	return ce.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ce *ce) WeekdaysWide() []string {
	return ce.daysWide
}

// Decimal returns the decimal point of number
func (ce *ce) Decimal() string {
	return ce.decimal
}

// Group returns the group of number
func (ce *ce) Group() string {
	return ce.group
}

// Group returns the minus sign of number
func (ce *ce) Minus() string {
	return ce.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ce' and handles both Whole and Real numbers based on 'v'
func (ce *ce) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ce.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ce.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ce.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ce' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ce *ce) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ce.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ce.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ce.percentSuffix...)

	b = append(b, ce.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ce'
func (ce *ce) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ce.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ce.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ce.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ce.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ce.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ce.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ce'
// in accounting notation.
func (ce *ce) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ce.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ce.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ce.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ce.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ce.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ce.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ce.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ce'
func (ce *ce) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ce'
func (ce *ce) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ce'
func (ce *ce) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ce'
func (ce *ce) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ce'
func (ce *ce) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ce'
func (ce *ce) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ce'
func (ce *ce) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ce'
func (ce *ce) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	return string(b)
}
