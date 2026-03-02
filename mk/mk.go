package mk

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mk struct {
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

// New returns a new instance of translator for the 'mk' locale
func New() locales.Translator {
	return &mk{
		locale:                 "mk",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 5, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "ден.", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "јан.", "фев.", "мар.", "апр.", "мај", "јун.", "јул.", "авг.", "сеп.", "окт.", "ное.", "дек."},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануари", "февруари", "март", "април", "мај", "јуни", "јули", "август", "септември", "октомври", "ноември", "декември"},
		daysAbbreviated:        []string{"нед.", "пон.", "вто.", "сре.", "чет.", "пет.", "саб."},
		daysNarrow:             []string{"н", "п", "в", "с", "ч", "п", "с"},
		daysWide:               []string{"недела", "понеделник", "вторник", "среда", "четврток", "петок", "сабота"},
		timezones:              map[string]string{"ACDT": "Летно сметање на времето во Централна Австралија", "ACST": "Стандардно време во Централна Австралија", "ACT": "Акре стандардно време", "ACWDT": "Летно сметање на времето во Централна и Западна Австралија", "ACWST": "Стандардно време во Централна и Западна Австралија", "ADT": "Атлантско летно сметање на времето", "ADT Arabia": "Арапско летно сметање на времето", "AEDT": "Летно сметање на времето во Источна Австралија", "AEST": "Стандардно време во Источна Австралија", "AFT": "Време во Авганистан", "AKDT": "Летно сметање на времето во Алјаска", "AKST": "Стандардно време во Алјаска", "AMST": "Летно сметање на времето во Амазон", "AMST Armenia": "Летно време во Ерменија", "AMT": "Стандардно време во Амазон", "AMT Armenia": "Стандардно време во Ерменија", "ANAST": "Анадирско летно време", "ANAT": "Анадирско стандардно време", "ARST": "Летно сметање на времето во Аргентина", "ART": "Стандардно време во Аргентина", "AST": "Атлантско стандардно време", "AST Arabia": "Стандардно арапско време", "AWDT": "Летно сметање на времето во Западна Австралија", "AWST": "Стандардно време во Западна Австралија", "AZST": "Летно време во Азербејџан", "AZT": "Стандардно време во Азербејџан", "BDT Bangladesh": "Летно време во Бангладеш", "BNT": "Време во Брунеј Дарусалам", "BOT": "Време во Боливија", "BRST": "Летно сметање на времето во Бразилија", "BRT": "Стандардно време во Бразилија", "BST Bangladesh": "Стандардно време во Бангладеш", "BT": "Време во Бутан", "CAST": "CAST", "CAT": "Средноафриканско време", "CCT": "Време во Кокосови Острови", "CDT": "Централно летно сметање на времето во Северна Америка", "CHADT": "Летно сметање на времето во Чатам", "CHAST": "Стандардно време во Чатам", "CHUT": "Време во Чуук", "CKT": "Стандардно време во Кукови Острови", "CKT DST": "Летно време во Кукови Острови", "CLST": "Летно сметање на времето во Чиле", "CLT": "Стандардно време во Чиле", "COST": "Летно сметање на времето во Колумбија", "COT": "Стандардно време во Колумбија", "CST": "Централно стандардно време во Северна Америка", "CST China": "Стандардно време во Кина", "CST China DST": "Летно сметање на времето во Кина", "CVST": "Летно сметање на времето во Кабо Верде", "CVT": "Стандардно време во Кабо Верде", "CXT": "Време во Божиќен Остров", "ChST": "Време во Чаморо", "ChST NMI": "ChST NMI", "CuDT": "Летно сметање на времето во Куба", "CuST": "Стандардно време во Куба", "DAVT": "Време во Дејвис", "DDUT": "Време во Димон Дирвил", "EASST": "Летно време во Велигденски Остров", "EAST": "Стандардно време во Велигденски Остров", "EAT": "Источноафриканско време", "ECT": "Време во Еквадор", "EDT": "Источно летно сметање на времето во Северна Америка", "EGDT": "Летно сметање на времето во Источен Гренланд", "EGST": "Стандардно време во Источен Гренланд", "EST": "Источно стандардно време во Северна Америка", "FEET": "Калининградско време", "FJT": "Стандардно време во Фиџи", "FJT Summer": "Летно време во Фиџи", "FKST": "Летно сметање на времето во Фолкландски Острови", "FKT": "Стандардно време во Фолкландски Острови", "FNST": "Летно сметање на времето во Фернандо де Нороња", "FNT": "Стандардно време во Фернандо де Нороња", "GALT": "Време во Галапагос", "GAMT": "Време во Гамбе", "GEST": "Летно време во Грузија", "GET": "Стандардно време во Грузија", "GFT": "Време во Француска Гвајана", "GIT": "Време во Гилбертови Острови", "GMT": "Средно време по Гринич", "GNSST": "GNSST", "GNST": "GNST", "GST": "Време во Јужна Џорџија", "GST Guam": "GST Guam", "GYT": "Време во Гвајана", "HADT": "Стандардно време во Хаваи - Алеутски острови", "HAST": "Стандардно време во Хаваи - Алеутски острови", "HKST": "Летно време во Хонг Конг", "HKT": "Стандардно време во Хонг Конг", "HOVST": "Летно време во Ховд", "HOVT": "Стандардно време во Ховд", "ICT": "Време во Индокина", "IDT": "Летно сметање на времето во Израел", "IOT": "Време во Индиски океан", "IRKST": "Летно време во Иркутск", "IRKT": "Стандардно време во Иркутск", "IRST": "Стандардно време во Иран", "IRST DST": "Летно сметање на времето во Иран", "IST": "Време во Индија", "IST Israel": "Стандардно време во Израел", "JDT": "Летно сметање на времето во Јапонија", "JST": "Стандардно време во Јапонија", "KOST": "Време во Косра", "KRAST": "Летно време во Краснојарск", "KRAT": "Стандардно време во Краснојарск", "KST": "Стандардно време во Кореја", "KST DST": "Летно сметање на времето во Кореја", "LHDT": "Летно сметање на времето во Лорд Хау", "LHST": "Стандардно време во Лорд Хау", "LINT": "Време во Линиски Острови", "MAGST": "Летно време во Магадан", "MAGT": "Стандардно време во Магадан", "MART": "Време во Маркесас", "MAWT": "Време во Мосон", "MDT": "MDT", "MESZ": "Средноевропско летно време", "MEZ": "Средноевропско стандардно време", "MHT": "Време во Маршалски Острови", "MMT": "Време во Мјанмар", "MSD": "Летно сметање на времето во Москва", "MST": "MST", "MUST": "Летно сметање на времето во Маврициус", "MUT": "Стандардно време во Маврициус", "MVT": "Време во Малдиви", "MYT": "Време во Малезија", "NCT": "Стандардно време во Нова Каледонија", "NDT": "Летно сметање на времето во Њуфаундленд", "NDT New Caledonia": "Летно време во Нова Каледонија", "NFDT": "Летно сметање на времето во Норфолшки Остров", "NFT": "Стандардно време во Норфолшки Остров", "NOVST": "Летно време во Новосибирск", "NOVT": "Стандардно време во Новосибирск", "NPT": "Време во Непал", "NRT": "Време во Науру", "NST": "Стандардно време во Њуфаундленд", "NUT": "Време во Ниуе", "NZDT": "Летно сметање на времето во Нов Зеланд", "NZST": "Стандардно време во Нов Зеланд", "OESZ": "Источноевропско летно време", "OEZ": "Источноевропско стандардно време", "OMSST": "Летно време во Омск", "OMST": "Стандардно време во Омск", "PDT": "Пацифичко летно сметање на времето во Северна Америка", "PDTM": "Летно пацифичко време во Мексико", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Време во Папуа Нова Гвинеја", "PHOT": "Време во Островите Феникс", "PKT": "Стандардно време во Пакистан", "PKT DST": "Летно време во Пакистан", "PMDT": "Летно сметање на времето во Сент Пјер и Микелан", "PMST": "Стандардно време во Сент Пјер и Микелан", "PONT": "Време во Понапе", "PST": "Пацифичко стандардно време во Северна Америка", "PST Philippine": "Стандардно време во Филипини", "PST Philippine DST": "Летно време во Филипини", "PST Pitcairn": "Време во Питкерн", "PSTM": "Стандардно пацифичко време во Мексико", "PWT": "Време во Палау", "PYST": "Летно сметање на времето во Парагвај", "PYT": "Стандардно време во Парагвај", "PYT Korea": "Време во Пјонгјанг", "RET": "Време во Рејунион", "ROTT": "Време во Ротера", "SAKST": "Летно време во Сакалин", "SAKT": "Стандардно време во Сакалин", "SAMST": "Самара летно сметање на времето", "SAMT": "Самара стандардно време", "SAST": "Време во Јужноафриканска Република", "SBT": "Време во Соломонски Острови", "SCT": "Време во Сејшели", "SGT": "Време во Сингапур", "SLST": "SLST", "SRT": "Време во Суринам", "SST Samoa": "Стандардно време во Самоа", "SST Samoa Apia": "Стандардно време во Апија", "SST Samoa Apia DST": "Летно време во Апија", "SST Samoa DST": "Летно време во Самоа", "SYOT": "Време во Сајова", "TAAF": "Француско јужно и антарктичко време", "TAHT": "Време во Тахити", "TJT": "Време во Таџикистан", "TKT": "Време во Токелау", "TLT": "Време во Источен Тимор", "TMST": "Летно време во Туркменистан", "TMT": "Стандардно време во Туркменистан", "TOST": "Летно време во Тонга", "TOT": "Стандардно време во Тонга", "TVT": "Време во Тувалу", "TWT": "Стандардно време во Тајпеј", "TWT DST": "Летно сметање на времето во Тајпеј", "ULAST": "Летно време во Улан Батор", "ULAT": "Стандардно време во Улан Батор", "UYST": "Летно сметање на времето во Уругвај", "UYT": "Стандардно време во Уругвај", "UZT": "Стандардно време во Узбекистан", "UZT DST": "Летно време во Узбекистан", "VET": "Време во Венецуела", "VLAST": "Летно време во Владивосток", "VLAT": "Стандардно време во Владивосток", "VOLST": "Летно сметање на времето во Волгоград", "VOLT": "Стандардно време во Волгоград", "VOST": "Време во Восток", "VUT": "Стандардно време во Вануату", "VUT DST": "Летно време во Вануату", "WAKT": "Време во Островот Вејк", "WARST": "Летно сметање на времето во западна Аргентина", "WART": "Стандардно време во западна Аргентина", "WAST": "Западноафриканско време", "WAT": "Западноафриканско време", "WESZ": "Западноевропско летно време", "WEZ": "Западноевропско стандардно време", "WFT": "Време во Валис и Футуна", "WGST": "Летно сметање на времето во Западен Гренланд", "WGT": "Стандардно време во Западен Гренланд", "WIB": "Време во Западна Индонезија", "WIT": "Време во Источна Индонезија", "WITA": "Време во Централна Индонезија", "YAKST": "Летно време во Јакутск", "YAKT": "Стандардно време во Јакутск", "YEKST": "Летно време во Екатеринбург", "YEKT": "Стандардно време во Екатеринбург", "YST": "Време во Јукон", "МСК": "Стандардно време во Москва", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Време во Западен Казахстан", "شىعىش قازاق ەلى": "Време во Источен Казахстан", "قازاق ەلى": "Време во Казахстан", "قىرعىزستان": "Време во Киргистан", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Летно сметање на времето во Перу"},
	}
}

// Locale returns the current translators string locale
func (mk *mk) Locale() string {
	return mk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mk'
func (mk *mk) PluralsCardinal() []locales.PluralRule {
	return mk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mk'
func (mk *mk) PluralsOrdinal() []locales.PluralRule {
	return mk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mk'
func (mk *mk) PluralsRange() []locales.PluralRule {
	return mk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mk'
func (mk *mk) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mk'
func (mk *mk) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if iMod10 == 2 && iMod100 != 12 {
		return locales.PluralRuleTwo
	} else if (iMod10 == 7 || iMod10 == 8) && (iMod100 != 17 && iMod100 != 18) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mk'
func (mk *mk) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mk *mk) MonthAbbreviated(month time.Month) string {
	return mk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mk *mk) MonthsAbbreviated() []string {
	return mk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mk *mk) MonthNarrow(month time.Month) string {
	return mk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mk *mk) MonthsNarrow() []string {
	return mk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mk *mk) MonthWide(month time.Month) string {
	return mk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mk *mk) MonthsWide() []string {
	return mk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mk *mk) WeekdayAbbreviated(weekday time.Weekday) string {
	return mk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mk *mk) WeekdaysAbbreviated() []string {
	return mk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mk *mk) WeekdayNarrow(weekday time.Weekday) string {
	return mk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mk *mk) WeekdaysNarrow() []string {
	return mk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mk *mk) WeekdayShort(weekday time.Weekday) string {
	return mk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mk *mk) WeekdaysShort() []string {
	return mk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mk *mk) WeekdayWide(weekday time.Weekday) string {
	return mk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mk *mk) WeekdaysWide() []string {
	return mk.daysWide
}

// Decimal returns the decimal point of number
func (mk *mk) Decimal() string {
	return mk.decimal
}

// Group returns the group of number
func (mk *mk) Group() string {
	return mk.group
}

// Group returns the minus sign of number
func (mk *mk) Minus() string {
	return mk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mk' and handles both Whole and Real numbers based on 'v'
func (mk *mk) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mk' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mk *mk) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mk.percentSuffix...)

	b = append(b, mk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mk'
func (mk *mk) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, mk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mk'
// in accounting notation.
func (mk *mk) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, mk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, mk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'mk'
func (mk *mk) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'mk'
func (mk *mk) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mk.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mk'
func (mk *mk) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mk'
func (mk *mk) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, mk.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mk'
func (mk *mk) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mk'
func (mk *mk) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mk'
func (mk *mk) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mk'
func (mk *mk) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
