package ru

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ru struct {
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

// New returns a new instance of translator for the 'ru' locale
func New() locales.Translator {
	return &ru{
		locale:                 "ru",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "ლ", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "₽", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "ТМТ", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "Cg", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "янв.", "февр.", "мар.", "апр.", "мая", "июн.", "июл.", "авг.", "сент.", "окт.", "нояб.", "дек."},
		monthsNarrow:           []string{"", "Я", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "ноября", "декабря"},
		daysAbbreviated:        []string{"вс", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysNarrow:             []string{"В", "П", "В", "С", "Ч", "П", "С"},
		daysWide:               []string{"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"до н.э.", "н.э."},
		erasWide:               []string{"до Рождества Христова", "от Рождества Христова"},
		timezones:              map[string]string{"ACDT": "Центральная Австралия, летнее время", "ACST": "Акри летнее время", "ACT": "Акри стандартное время", "ACWDT": "Центральная Австралия, западное летнее время", "ACWST": "Центральная Австралия, западное стандартное время", "ADT": "Атлантическое летнее время", "ADT Arabia": "Саудовская Аравия, летнее время", "AEDT": "Восточная Австралия, летнее время", "AEST": "Восточная Австралия, стандартное время", "AFT": "Афганистан", "AKDT": "Аляска, летнее время", "AKST": "Аляска, стандартное время", "AMST": "Амазонка, летнее время", "AMST Armenia": "Армения, летнее время", "AMT": "Амазонка, стандартное время", "AMT Armenia": "Армения, стандартное время", "ANAST": "Анадырь летнее время", "ANAT": "Анадырь стандартное время", "ARST": "Аргентина, летнее время", "ART": "Аргентина, стандартное время", "AST": "Атлантическое стандартное время", "AST Arabia": "Саудовская Аравия, стандартное время", "AWDT": "Западная Австралия, летнее время", "AWST": "Западная Австралия, стандартное время", "AZST": "Азербайджан, летнее время", "AZT": "Азербайджан, стандартное время", "BDT Bangladesh": "Бангладеш, летнее время", "BNT": "Бруней-Даруссалам", "BOT": "Боливия", "BRST": "Бразилия, летнее время", "BRT": "Бразилия, стандартное время", "BST Bangladesh": "Бангладеш, стандартное время", "BT": "Бутан", "CAST": "Кейси", "CAT": "Центральная Африка", "CCT": "Кокосовые о-ва", "CDT": "Центральная Америка, летнее время", "CHADT": "Чатем, летнее время", "CHAST": "Чатем, стандартное время", "CHUT": "Трук", "CKT": "Острова Кука, стандартное время", "CKT DST": "Острова Кука, полулетнее время", "CLST": "Чили, летнее время", "CLT": "Чили, стандартное время", "COST": "Колумбия, летнее время", "COT": "Колумбия, стандартное время", "CST": "Центральная Америка, стандартное время", "CST China": "Китай, стандартное время", "CST China DST": "Китай, летнее время", "CVST": "Кабо-Верде, летнее время", "CVT": "Кабо-Верде, стандартное время", "CXT": "о-в Рождества", "ChST": "Чаморро", "ChST NMI": "Северные Марианские о-ва", "CuDT": "Куба, летнее время", "CuST": "Куба, стандартное время", "DAVT": "Дейвис", "DDUT": "Дюмон-д’Юрвиль", "EASST": "О-в Пасхи, летнее время", "EAST": "О-в Пасхи, стандартное время", "EAT": "Восточная Африка", "ECT": "Эквадор", "EDT": "Восточная Америка, летнее время", "EGDT": "Восточная Гренландия, летнее время", "EGST": "Восточная Гренландия, стандарное время", "EST": "Восточная Америка, стандартное время", "FEET": "Московское время", "FJT": "Фиджи, стандартное время", "FJT Summer": "Фиджи, летнее время", "FKST": "Фолклендские о-ва, летнее время", "FKT": "Фолклендские о-ва, стандартное время", "FNST": "Фернанду-ди-Норонья, летнее время", "FNT": "Фернанду-ди-Норонья, стандартное время", "GALT": "Галапагосские о-ва", "GAMT": "Гамбье", "GEST": "Грузия, летнее время", "GET": "Грузия, стандартное время", "GFT": "Французская Гвиана", "GIT": "о-ва Гилберта", "GMT": "Среднее время по Гринвичу", "GNSST": "GNSST", "GNST": "GNST", "GST": "Южная Георгия", "GST Guam": "Гуам", "GYT": "Гайана", "HADT": "Гавайско-алеутское летнее время", "HAST": "Гавайско-алеутское стандартное время", "HKST": "Гонконг, летнее время", "HKT": "Гонконг, стандартное время", "HOVST": "Ховд, летнее время", "HOVT": "Ховд, стандартное время", "ICT": "Индокитай", "IDT": "Израиль, летнее время", "IOT": "Индийский океан", "IRKST": "Иркутск, летнее время", "IRKT": "Иркутск, стандартное время", "IRST": "Иран, стандартное время", "IRST DST": "Иран, летнее время", "IST": "Индия", "IST Israel": "Израиль, стандартное время", "JDT": "Япония, летнее время", "JST": "Япония, стандартное время", "KOST": "Косрае", "KRAST": "Красноярск, летнее время", "KRAT": "Красноярск, стандартное время", "KST": "Корея, стандартное время", "KST DST": "Корея, летнее время", "LHDT": "Лорд-Хау, летнее время", "LHST": "Лорд-Хау, стандартное время", "LINT": "о-ва Лайн", "MAGST": "Магадан, летнее время", "MAGT": "Магадан, стандартное время", "MART": "Маркизские о-ва", "MAWT": "Моусон", "MDT": "Летнее горное время (Северная Америка)", "MESZ": "Центральная Европа, летнее время", "MEZ": "Центральная Европа, стандартное время", "MHT": "Маршалловы Острова", "MMT": "Мьянма", "MSD": "Москва, летнее время", "MST": "Стандартное горное время (Северная Америка)", "MUST": "Маврикий, летнее время", "MUT": "Маврикий, стандартное время", "MVT": "Мальдивы", "MYT": "Малайзия", "NCT": "Новая Каледония, стандартное время", "NDT": "Ньюфаундленд, летнее время", "NDT New Caledonia": "Новая Каледония, летнее время", "NFDT": "Норфолк, летнее время", "NFT": "Норфолк, стандартное время", "NOVST": "Новосибирск, летнее время", "NOVT": "Новосибирск, стандартное время", "NPT": "Непал", "NRT": "Науру", "NST": "Ньюфаундленд, стандартное время", "NUT": "Ниуэ", "NZDT": "Новая Зеландия, летнее время", "NZST": "Новая Зеландия, стандартное время", "OESZ": "Восточная Европа, летнее время", "OEZ": "Восточная Европа, стандартное время", "OMSST": "Омск, летнее время", "OMST": "Омск, стандартное время", "PDT": "Тихоокеанское летнее время", "PDTM": "Тихоокеанское мексиканское летнее время", "PETDT": "Петропавловск-Камчатский, летнее время", "PETST": "Петропавловск-Камчатский, стандартное время", "PGT": "Папуа – Новая Гвинея", "PHOT": "о-ва Феникс", "PKT": "Пакистан, стандартное время", "PKT DST": "Пакистан, летнее время", "PMDT": "Сен-Пьер и Микелон, летнее время", "PMST": "Сен-Пьер и Микелон, стандартное время", "PONT": "Понпеи", "PST": "Тихоокеанское стандартное время", "PST Philippine": "Филиппины, стандартное время", "PST Philippine DST": "Филиппины, летнее время", "PST Pitcairn": "Питкэрн", "PSTM": "Тихоокеанское мексиканское стандартное время", "PWT": "Палау", "PYST": "Парагвай, летнее время", "PYT": "Парагвай, стандартное время", "PYT Korea": "Пхеньян", "RET": "Реюньон", "ROTT": "Ротера", "SAKST": "Сахалин, летнее время", "SAKT": "Сахалин, стандартное время", "SAMST": "Самарское летнее время", "SAMT": "Самарское стандартное время", "SAST": "Южная Африка", "SBT": "Соломоновы Острова", "SCT": "Сейшельские Острова", "SGT": "Сингапур", "SLST": "Шри-Ланка", "SRT": "Суринам", "SST Samoa": "Самоа, стандартное время", "SST Samoa Apia": "Апиа, стандартное время", "SST Samoa Apia DST": "Апиа, летнее время", "SST Samoa DST": "Самоа, летнее время", "SYOT": "Сёва", "TAAF": "Французские Южные и Антарктические территории", "TAHT": "Таити", "TJT": "Таджикистан", "TKT": "Токелау", "TLT": "Восточный Тимор", "TMST": "Туркменистан, летнее время", "TMT": "Туркменистан, стандартное время", "TOST": "Тонга, летнее время", "TOT": "Тонга, стандартное время", "TVT": "Тувалу", "TWT": "Тайвань, стандартное время", "TWT DST": "Тайвань, летнее время", "ULAST": "Улан-Батор, летнее время", "ULAT": "Улан-Батор, стандартное время", "UYST": "Уругвай, летнее время", "UYT": "Уругвай, стандартное время", "UZT": "Узбекистан, стандартное время", "UZT DST": "Узбекистан, летнее время", "VET": "Венесуэла", "VLAST": "Владивосток, летнее время", "VLAT": "Владивосток, стандартное время", "VOLST": "Волгоград, летнее время", "VOLT": "Волгоград, стандартное время", "VOST": "Восток", "VUT": "Вануату, стандартное время", "VUT DST": "Вануату, летнее время", "WAKT": "Уэйк", "WARST": "Западная Аргентина, летнее время", "WART": "Западная Аргентина, стандартное время", "WAST": "Западная Африка", "WAT": "Западная Африка", "WESZ": "Западная Европа, летнее время", "WEZ": "Западная Европа, стандартное время", "WFT": "Уоллис и Футуна", "WGST": "Западная Гренландия, летнее время", "WGT": "Западная Гренландия, стандартное время", "WIB": "Западная Индонезия", "WIT": "Восточная Индонезия", "WITA": "Центральная Индонезия", "YAKST": "Якутск, летнее время", "YAKT": "Якутск, стандартное время", "YEKST": "Екатеринбург, летнее время", "YEKT": "Екатеринбург, стандартное время", "YST": "Юкон", "МСК": "Москва, стандартное время", "اقتاۋ": "Актау, стандартное время", "اقتاۋ قالاسى": "Актау летнее время", "اقتوبە": "Актобе стандартное время", "اقتوبە قالاسى": "Актобе летнее время", "الماتى": "Алма-Ата стандартное время", "الماتى قالاسى": "Алма-Ата летнее время", "باتىس قازاق ەلى": "Западный Казахстан", "شىعىش قازاق ەلى": "Восточный Казахстан", "قازاق ەلى": "Казахстан", "قىرعىزستان": "Киргизия", "قىزىلوردا": "Кызылорда, стандартное время*", "قىزىلوردا قالاسى": "Кызылорда, летнее время*", "∅∅∅": "Перу, летнее время"},
	}
}

// Locale returns the current translators string locale
func (ru *ru) Locale() string {
	return ru.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ru'
func (ru *ru) PluralsCardinal() []locales.PluralRule {
	return ru.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ru'
func (ru *ru) PluralsOrdinal() []locales.PluralRule {
	return ru.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ru'
func (ru *ru) PluralsRange() []locales.PluralRule {
	return ru.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ru'
func (ru *ru) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if v == 0 && iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14) {
		return locales.PluralRuleFew
	} else if (v == 0 && iMod10 == 0) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 11 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ru'
func (ru *ru) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ru'
func (ru *ru) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ru.CardinalPluralRule(num1, v1)
	end := ru.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ru *ru) MonthAbbreviated(month time.Month) string {
	return ru.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ru *ru) MonthsAbbreviated() []string {
	return ru.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ru *ru) MonthNarrow(month time.Month) string {
	return ru.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ru *ru) MonthsNarrow() []string {
	return ru.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ru *ru) MonthWide(month time.Month) string {
	return ru.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ru *ru) MonthsWide() []string {
	return ru.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ru *ru) WeekdayAbbreviated(weekday time.Weekday) string {
	return ru.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ru *ru) WeekdaysAbbreviated() []string {
	return ru.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ru *ru) WeekdayNarrow(weekday time.Weekday) string {
	return ru.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ru *ru) WeekdaysNarrow() []string {
	return ru.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ru *ru) WeekdayShort(weekday time.Weekday) string {
	return ru.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ru *ru) WeekdaysShort() []string {
	return ru.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ru *ru) WeekdayWide(weekday time.Weekday) string {
	return ru.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ru *ru) WeekdaysWide() []string {
	return ru.daysWide
}

// Decimal returns the decimal point of number
func (ru *ru) Decimal() string {
	return ru.decimal
}

// Group returns the group of number
func (ru *ru) Group() string {
	return ru.group
}

// Group returns the minus sign of number
func (ru *ru) Minus() string {
	return ru.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ru' and handles both Whole and Real numbers based on 'v'
func (ru *ru) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ru' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ru *ru) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 1
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ru.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ru.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ru'
func (ru *ru) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ru.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ru.group) - 1; j >= 0; j-- {
					b = append(b, ru.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ru.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ru.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ru.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ru'
// in accounting notation.
func (ru *ru) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ru.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ru.group) - 1; j >= 0; j-- {
					b = append(b, ru.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ru.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ru.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ru.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ru.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ru'
func (ru *ru) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ru'
func (ru *ru) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ru.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'ru'
func (ru *ru) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ru.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'ru'
func (ru *ru) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ru.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ru.monthsWide[t.Month()]...)
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

// FmtTimeShort returns the short time representation of 't' for 'ru'
func (ru *ru) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ru'
func (ru *ru) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ru'
func (ru *ru) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ru'
func (ru *ru) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ru.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ru.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
