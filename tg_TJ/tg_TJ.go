package tg_TJ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type tg_TJ struct {
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

// New returns a new instance of translator for the 'tg_TJ' locale
func New() locales.Translator {
	return &tg_TJ{
		locale:                 "tg_TJ",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "Янв", "Фев", "Мар", "Апр", "Май", "Июн", "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"},
		monthsNarrow:           []string{"", "Я", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "Январ", "Феврал", "Март", "Апрел", "Май", "Июн", "Июл", "Август", "Сентябр", "Октябр", "Ноябр", "Декабр"},
		daysAbbreviated:        []string{"Яшб", "Дшб", "Сшб", "Чшб", "Пшб", "Ҷмъ", "Шнб"},
		daysNarrow:             []string{"Я", "Д", "С", "Ч", "П", "Ҷ", "Ш"},
		daysWide:               []string{"Якшанбе", "Душанбе", "Сешанбе", "Чоршанбе", "Панҷшанбе", "Ҷумъа", "Шанбе"},
		timezones:              map[string]string{"ACDT": "Вақти рӯонаи Австралияи Марказӣ", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Вақти рӯзонаи Ғарбии Австралия Марказӣ", "ACWST": "Вақти стандартии Ғарбии Австралия Марказӣ", "ADT": "Вақти рӯзонаи атлантикӣ", "ADT Arabia": "Вақти рӯзонаи Арабистон", "AEDT": "Вақти рӯзонаи шарқии Австралия", "AEST": "Вақти стандартии шарқии Австралия", "AFT": "Вақти Афғонистон", "AKDT": "Вақти рӯзонаи Аляска", "AKST": "Вақти стандартии Аляска", "AMST": "Вақти тобистонаи Амазон", "AMST Armenia": "Вақти рӯзонаи Арманистон", "AMT": "Вақти стандартии Амазон", "AMT Armenia": "Вақти стандартии Арманистон", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Вақти тобистонаи Аргентина", "ART": "Вақти стандартии Аргентина", "AST": "Вақти стандартии атлантикӣ", "AST Arabia": "Вақти стандартии Арабистон", "AWDT": "Вақти рӯзонаи Австралияи Ғарбӣ", "AWST": "Вақти стандартии Австралияи Ғарбӣ", "AZST": "Вақти тобистонаи Озарбойҷон", "AZT": "Вақти стандартии Озарбойҷон", "BDT Bangladesh": "Вақти тобистонаи Бангладеш", "BNT": "Вақти Бруней Доруссалом", "BOT": "Вақти Боливия", "BRST": "Вақти тобистонаи Бразилия", "BRT": "Вақти стандартии Бразилия", "BST Bangladesh": "Вақти стандартии Бангладеш", "BT": "Вақти Бутан", "CAST": "CAST", "CAT": "Вақти Африқои Марказӣ", "CCT": "Вақти Ҷазираҳои Кокос", "CDT": "Вақти рӯзонаи марказӣ", "CHADT": "Вақти рӯзонаи Чатам", "CHAST": "Вақти стандартии Чатам", "CHUT": "Вақти Чук", "CKT": "Вақти стандартии Ҷазираҳои Кук", "CKT DST": "Вақти нимаи тобистонаи Ҷазираҳои Кук", "CLST": "Вақти тобистонаи Чили", "CLT": "Вақти стандартии Чили", "COST": "Вақти тобистонаи Колумбия", "COT": "Вақти стандартии Колумбия", "CST": "Вақти стандартии марказӣ", "CST China": "Вақти стандартии Чин", "CST China DST": "Вақти рӯзонаи Чин", "CVST": "Вақти тобистонаи Кабо-Верде", "CVT": "Вақти стандартии Кабо-Верде", "CXT": "Вақти ҷазираи Мавлуди Исо", "ChST": "Вақти стандартии Чаморро", "ChST NMI": "ChST NMI", "CuDT": "Вақти рӯзонаи Куба", "CuST": "Вақти стандартии Куба", "DAVT": "Вақти Давис", "DDUT": "Вақти Дюмон-д’Урвил", "EASST": "Вақти тобистонаи ҷазираи Пасха", "EAST": "Вақти стандартии ҷазираи Пасха", "EAT": "Вақти Африқои Шарқӣ", "ECT": "Вақти Эквадор", "EDT": "Вақти рӯзонаи шарқӣ", "EGDT": "Вақти тобистонаи Гренландияи Шарқӣ", "EGST": "Вақти стандартии Гренландияи Шарқӣ", "EST": "Вақти стандартии шарқӣ", "FEET": "Вақти Аврупои Шарқӣ", "FJT": "Вақти стандартии Фиҷи", "FJT Summer": "Вақти тобистонаи Фиҷи", "FKST": "Вақти тобистонаи Ҷазираҳои Фолкленд", "FKT": "Вақти стандартии Ҷазираҳои Фолкленд", "FNST": "Вақти тобистонаи Фернандо де Норонха", "FNT": "Вақти стандартии Фернандо де Норонха", "GALT": "Вақти Галапагос", "GAMT": "Вақти Гамбир", "GEST": "Вақти тобистонаи Гурҷистон", "GET": "Вақти стандартии Гурҷистон", "GFT": "Вақти Гвианаи Фаронса", "GIT": "Вақти Ҷазираҳои Гилберт", "GMT": "Вақти миёнаи Гринвич", "GNSST": "GNSST", "GNST": "GNST", "GST": "Вақти стандартии Халиҷи Форс", "GST Guam": "GST Guam", "GYT": "Вақти Гайана", "HADT": "Вақти стандартии Ҳавайӣ-Алеутӣ", "HAST": "Вақти стандартии Ҳавайӣ-Алеутӣ", "HKST": "Вақти тобистонаи Ҳонконг", "HKT": "Вақти стандартии Ҳонконг", "HOVST": "Вақти тобистонаи Ховд", "HOVT": "Вақти стандартии Ҳовд", "ICT": "Вақти Ҳиндучин", "IDT": "Вақти рӯзонаи Исроил", "IOT": "Вақти уқёнуси Ҳинд", "IRKST": "Вақти тобистонаи Иркутск", "IRKT": "Вақти стандартии Иркутск", "IRST": "Вақти стандартии Эрон", "IRST DST": "Вақти рӯзонаи Эрон", "IST": "Вақти стандартии Ҳиндустон", "IST Israel": "Вақти стандартии Исроил", "JDT": "Вақти равшании Ҷопон", "JST": "Вақти стандартии Ҷопон", "KOST": "Вақти Косрае", "KRAST": "Вақти тобистонаи Красноярск", "KRAT": "Вақти стандартии Красноярск", "KST": "Вақти стандартии Корея", "KST DST": "Вақти рӯзонаи Корея", "LHDT": "Вақти рӯзонаи Лорд Хоу", "LHST": "Вақти стандартии Лорд Хоу", "LINT": "Вақти Ҷазираҳои Лин", "MAGST": "Вақти тобистонаи Магадан", "MAGT": "Вақти стандартии Магадан", "MART": "Вақти Маркес", "MAWT": "Вақти Мавсон", "MDT": "Вақти рӯзонаи кӯҳӣ", "MESZ": "Вақти тобистонаи аврупоии марказӣ", "MEZ": "Вақти стандартии аврупоии марказӣ", "MHT": "Вақти Ҷазираҳои Маршалл", "MMT": "Вақти Мянма", "MSD": "Вақти тобистонаи Маскав", "MST": "Вақти стандартии кӯҳӣ", "MUST": "Вақти тобистонаи Маврикий", "MUT": "Вақти стандартии Маврикий", "MVT": "Вақти Малдив", "MYT": "Вақти Малайзия", "NCT": "Вақти стандартии Каледонияи Нав", "NDT": "Вақти рӯзонаи Нюфаундленд", "NDT New Caledonia": "Вақти тобистонаи Каледонияи Нав", "NFDT": "Вақти рӯзонаи ҷазираи Норфолк", "NFT": "Вақти стандартии ҷазираи Норфолк", "NOVST": "Вақти тобистонаи Новосибирск", "NOVT": "Вақти стандартии Новосибирск", "NPT": "Вақти Непал", "NRT": "Вақти Науру", "NST": "Вақти стандартии Нюфаундленд", "NUT": "Вақти Ниуэ", "NZDT": "Вақти рӯзонаи Зеландияи Нав", "NZST": "Вақти стандартии Зеландияи Нав", "OESZ": "Вақти тобистонаи аврупоии шарқӣ", "OEZ": "Вақти стандартии аврупоии шарқӣ", "OMSST": "Вақти тобистонаи Омск", "OMST": "Омск вақти стандартӣ", "PDT": "Вақти рӯзонаи Уқёнуси Ором", "PDTM": "Вақти рӯзонаи Уқёнуси Ором Мексика", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Вақти Папуа Гвинеяи Нав", "PHOT": "Вақти Ҷазираҳои Финикс", "PKT": "Вақти стандартии Покистон", "PKT DST": "Вақти тобистонаи Покистон", "PMDT": "Вақти рӯзонаи Сент-Пиер ва Микелон", "PMST": "Вақти стандартии Сент-Пиер ва Микелон", "PONT": "Ponape Time", "PST": "Вақти стандартии Уқёнуси Ором", "PST Philippine": "Вақти стандартии Филиппин", "PST Philippine DST": "Вақти тобистонаи Филиппин", "PST Pitcairn": "Вақти Питкэрн", "PSTM": "Вақти стандартии Уқёнуси Ором Мексика", "PWT": "Вақти Палау", "PYST": "Вақти тобистонаи Парагвай", "PYT": "Вақти стандартии Парагвай", "PYT Korea": "Вақти Пхенян", "RET": "Вақти Реюнион", "ROTT": "Вақти Ротера", "SAKST": "Вақти тобистонаи Сахалин", "SAKT": "Вақти стандартии Сахалин", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Вақти стандартии Африқои Ҷанубӣ", "SBT": "Вақти Ҷазираҳои Соломон", "SCT": "Вақти Сейшел", "SGT": "Вақти стандартии Сингапур", "SLST": "SLST", "SRT": "Вақти Суринам", "SST Samoa": "Вақти стандартии Самоа", "SST Samoa Apia": "Вақти стандартии Апиа", "SST Samoa Apia DST": "Вақти рӯзонаи Апиа", "SST Samoa DST": "Вақти рӯзонаи Самоа", "SYOT": "Вақти Сёва", "TAAF": "Вақти ҷанубӣ ва Антарктидаи Фаронса", "TAHT": "Вақти Таити", "TJT": "Вақти Тоҷикистон", "TKT": "Вақти Токелау", "TLT": "Вақти Тимори Шарқӣ", "TMST": "Вақти тобистонаи Туркманистон", "TMT": "Вақти стандартии Туркманистон", "TOST": "Вақти тобистонаи Тонга", "TOT": "Вақти стандартии Тонга", "TVT": "Вақти Тувалу", "TWT": "Вақти стандартии Тайбэй", "TWT DST": "Вақти рӯзонаи Тайбэй", "ULAST": "Вақти тобистонаи Улан-Батор", "ULAT": "Вақти стандартии Улан-Батор", "UYST": "Вақти тобистонаи Уругвай", "UYT": "Вақти стандартии Уругвай", "UZT": "Вақти стандартии Ӯзбекистон", "UZT DST": "Вақти тобистонаи Ӯзбекистон", "VET": "Вақти Венесуэла", "VLAST": "Вақти тобистонаи Владивосток", "VLAT": "Вақти стандартии Владивосток", "VOLST": "Вақти тобистонаи Волгоград", "VOLT": "Вақти стандартии Волгоград", "VOST": "Вақти Восток", "VUT": "Вақти стандартии Вануату", "VUT DST": "Вақти тобистонаи Вануату", "WAKT": "Вақти бедории ҷазира", "WARST": "Вақти тобистонаи Аргентинаи Ғарбӣ", "WART": "Вақти стандартии Аргентинаи Ғарбӣ", "WAST": "Вақти Африқои Ғарбӣ", "WAT": "Вақти Африқои Ғарбӣ", "WESZ": "Вақти тобистонаи аврупоии ғарбӣ", "WEZ": "Вақти стандартии аврупоии ғарбӣ", "WFT": "Вақти Уоллис ва Футуна", "WGST": "Вақти тобистонаи Гренландияи Ғарбӣ", "WGT": "Вақти стандартии Гренландияи Ғарбӣ", "WIB": "Вақти Индонезияи Ғарбӣ", "WIT": "Вақти шарқии Индонезия", "WITA": "Вақти Индонезияи Марказӣ", "YAKST": "Якутск вақти тобистона", "YAKT": "Якутск вақти стандартӣ", "YEKST": "Вақти тобистонаи Екатеринбург", "YEKT": "Вақти стандартии Екатеринбург", "YST": "Вақти Юкон", "МСК": "Вақти стандартии Маскав", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Вақти Қазоқистони Ғарбӣ", "شىعىش قازاق ەلى": "Вақти Қазоқистони Шарқӣ", "قازاق ەلى": "Вақти Қазоқистон", "قىرعىزستان": "Вақти Қирғизистон", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Вақти тобистонаи Перу"},
	}
}

// Locale returns the current translators string locale
func (tg *tg_TJ) Locale() string {
	return tg.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'tg_TJ'
func (tg *tg_TJ) PluralsCardinal() []locales.PluralRule {
	return tg.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'tg_TJ'
func (tg *tg_TJ) PluralsOrdinal() []locales.PluralRule {
	return tg.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'tg_TJ'
func (tg *tg_TJ) PluralsRange() []locales.PluralRule {
	return tg.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'tg_TJ'
func (tg *tg_TJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'tg_TJ'
func (tg *tg_TJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'tg_TJ'
func (tg *tg_TJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (tg *tg_TJ) MonthAbbreviated(month time.Month) string {
	return tg.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (tg *tg_TJ) MonthsAbbreviated() []string {
	return tg.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (tg *tg_TJ) MonthNarrow(month time.Month) string {
	return tg.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (tg *tg_TJ) MonthsNarrow() []string {
	return tg.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (tg *tg_TJ) MonthWide(month time.Month) string {
	return tg.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (tg *tg_TJ) MonthsWide() []string {
	return tg.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (tg *tg_TJ) WeekdayAbbreviated(weekday time.Weekday) string {
	return tg.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (tg *tg_TJ) WeekdaysAbbreviated() []string {
	return tg.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (tg *tg_TJ) WeekdayNarrow(weekday time.Weekday) string {
	return tg.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (tg *tg_TJ) WeekdaysNarrow() []string {
	return tg.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (tg *tg_TJ) WeekdayShort(weekday time.Weekday) string {
	return tg.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (tg *tg_TJ) WeekdaysShort() []string {
	return tg.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (tg *tg_TJ) WeekdayWide(weekday time.Weekday) string {
	return tg.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (tg *tg_TJ) WeekdaysWide() []string {
	return tg.daysWide
}

// Decimal returns the decimal point of number
func (tg *tg_TJ) Decimal() string {
	return tg.decimal
}

// Group returns the group of number
func (tg *tg_TJ) Group() string {
	return tg.group
}

// Group returns the minus sign of number
func (tg *tg_TJ) Minus() string {
	return tg.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'tg_TJ' and handles both Whole and Real numbers based on 'v'
func (tg *tg_TJ) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tg.group) - 1; j >= 0; j-- {
					b = append(b, tg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'tg_TJ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (tg *tg_TJ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tg.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, tg.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'tg_TJ'
func (tg *tg_TJ) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tg.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tg.group) - 1; j >= 0; j-- {
					b = append(b, tg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tg.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, tg.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'tg_TJ'
// in accounting notation.
func (tg *tg_TJ) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tg.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tg.group) - 1; j >= 0; j-- {
					b = append(b, tg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tg.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, tg.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, tg.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tg.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tg.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, tg.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tg.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'tg_TJ'
func (tg *tg_TJ) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := tg.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
