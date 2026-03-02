package uz_Cyrl

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type uz_Cyrl struct {
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

// New returns a new instance of translator for the 'uz_Cyrl' locale
func New() locales.Translator {
	return &uz_Cyrl{
		locale:                 "uz_Cyrl",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "сўм", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
		monthsNarrow:           []string{"", "Я", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "январ", "феврал", "март", "апрел", "май", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"},
		daysAbbreviated:        []string{"якш", "душ", "сеш", "чор", "пай", "жум", "шан"},
		daysNarrow:             []string{"Я", "Д", "С", "Ч", "П", "Ж", "Ш"},
		daysShort:              []string{"як", "ду", "се", "чо", "па", "жу", "ша"},
		daysWide:               []string{"якшанба", "душанба", "сешанба", "чоршанба", "пайшанба", "жума", "шанба"},
		periodsAbbreviated:     []string{"ТО", "ТК"},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"м.а.", "милодий"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"милоддан аввалги", "эрамиздан аввалги"},
		timezones:              map[string]string{"ACDT": "Марказий Австралия кундузги вақти", "ACST": "Марказий Австралия стандарт вақти", "ACT": "ACT", "ACWDT": "Марказий Австралия Ғарбий кундузги вақти", "ACWST": "Марказий Австралия Ғарбий стандарт вақти", "ADT": "Атлантика кундузги вақти", "ADT Arabia": "Арабистон кундузги вақти", "AEDT": "Шарқий Австралия кундузги вақти", "AEST": "Шарқий Австралия стандарт вақти", "AFT": "Афғонистон вақти", "AKDT": "Аляска кундузги вақти", "AKST": "Аляска стандарт вақти", "AMST": "Амазонка ёзги вақти", "AMST Armenia": "Арманистон ёзги вақти", "AMT": "Амазонка стандарт вақти", "AMT Armenia": "Арманистон стандарт вақти", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентина ёзги вақти", "ART": "Аргентина стандарт вақти", "AST": "Атлантика стандарт вақти", "AST Arabia": "Арабистон стандарт вақти", "AWDT": "Ғарбий Австралия кундузги вақти", "AWST": "Ғарбий Австралия стандарт вақти", "AZST": "Озарбайжон ёзги вақти", "AZT": "Озарбайжон стандарт вақти", "BDT Bangladesh": "Бангладеш ёзги вақти", "BNT": "Бруней Даруссалом вақти", "BOT": "Боливия вақти", "BRST": "Бразилия ёзги вақти", "BRT": "Бразилия стандарт вақти", "BST Bangladesh": "Бангладеш стандарт вақти", "BT": "Бутан вақти", "CAST": "CAST", "CAT": "Марказий Африка вақти", "CCT": "Кокос ороллари вақти", "CDT": "Шимолий Америка марказий кундузги вақти", "CHADT": "Чатхам кундузги вақти", "CHAST": "Чатхам стандарт вақти", "CHUT": "Чуук вақти", "CKT": "Кук ороллари стандарт вақти", "CKT DST": "Кук ороллари ярим ёзги вақти", "CLST": "Чили ёзги вақти", "CLT": "Чили стандарт вақти", "COST": "Колумбия ёзги вақти", "COT": "Колумбия стандарт вақти", "CST": "Шимолий Америка марказий стандарт вақти", "CST China": "Хитой стандарт вақти", "CST China DST": "Хитой кундузги вақти", "CVST": "Кабо-Верде ёзги вақти", "CVT": "Кабо-Верде стандарт вақти", "CXT": "Рождество ороли вақти", "ChST": "Каморро вақти", "ChST NMI": "ChST NMI", "CuDT": "Куба кундузги вақти", "CuST": "Куба стандарт вақти", "DAVT": "Дэвис вақти", "DDUT": "Думонт-д-Урвил вақти", "EASST": "Пасхи ороли ёзги вақти", "EAST": "Пасхи ороли стандарт вақти", "EAT": "Шарқий Африка вақти", "ECT": "Эквадор вақти", "EDT": "Шимолий Америка шарқий кундузги вақти", "EGDT": "Шарқий Гренландия ёзги вақти", "EGST": "Шарқий Гренландия стандарт вақти", "EST": "Шимолий Америка шарқий стандарт вақти", "FEET": "Kaliningrad va Minsk vaqti", "FJT": "Фижи стандарт вақти", "FJT Summer": "Фижи ёзги вақти", "FKST": "Фолькленд ороллари ёзги вақти", "FKT": "Фолькленд ороллари стандарт вақти", "FNST": "Фернандо де Норонья ёзги вақти", "FNT": "Фернандо де Норонья стандарт вақти", "GALT": "Галапагос вақти", "GAMT": "Гамбиер вақти", "GEST": "Грузия ёзги вақти", "GET": "Грузия стандарт вақти", "GFT": "Француз Гвианаси вақти", "GIT": "Гилберт ороллари вақти", "GMT": "Гринвич вақти", "GNSST": "GNSST", "GNST": "GNST", "GST": "Кўрфаз вақти", "GST Guam": "GST Guam", "GYT": "Гайана вақти", "HADT": "Гавайи-алеут кундузги вақти", "HAST": "Гавайи-алеут стандарт вақти", "HKST": "Гонконг ёзги вақти", "HKT": "Гонконг стандарт вақти", "HOVST": "Ховд ёзги вақти", "HOVT": "Ховд стандарт вақти", "ICT": "Ҳинд-Хитой вақти", "IDT": "Исроил кундузги вақти", "IOT": "Ҳинд океани вақти", "IRKST": "Иркутск ёзги вақти", "IRKT": "Иркутск стандарт вақти", "IRST": "Эрон стандарт вақти", "IRST DST": "Эрон кундузги вақти", "IST": "Ҳиндистон вақти", "IST Israel": "Исроил стандарт вақти", "JDT": "Япония кундузги вақти", "JST": "Япония стандарт вақти", "KOST": "Косрае вақти", "KRAST": "Красноярск ёзги вақти", "KRAT": "Красноярск стандарт вақти", "KST": "Корея стандарт вақти", "KST DST": "Корея кундузги вақти", "LHDT": "Лорд Хове кундузги вақти", "LHST": "Лорд Хове стандарт вақти", "LINT": "Лайн ороллари вақти", "MAGST": "Магадан ёзги вақти", "MAGT": "Магадан стандарт вақти", "MART": "Маркезас вақти", "MAWT": "Моувсон вақти", "MDT": "MDT", "MESZ": "Марказий Европа ёзги вақти", "MEZ": "Марказий Европа стандарт вақти", "MHT": "Маршалл ороллари вақти", "MMT": "Мьянма вақти", "MSD": "Москва ёзги вақти", "MST": "MST", "MUST": "Маврикий ёзги вақти", "MUT": "Маврикий стандарт вақти", "MVT": "Мальдив ороллар", "MYT": "Малайзия вақти", "NCT": "Янги Каледония стандарт вақти", "NDT": "Ньюфаундленд кундузги вақти", "NDT New Caledonia": "Янги Каледония ёзги вақти", "NFDT": "Норфолк ороли ёзги вақти", "NFT": "Норфолк ороли стандарт вақти", "NOVST": "Новосибирск ёзги вақти", "NOVT": "Новосибирск стандарт вақти", "NPT": "Непал вақти", "NRT": "Науру вақти", "NST": "Ньюфаундленд стандарт вақти", "NUT": "Ниуе вақти", "NZDT": "Янги Зеландия кундузги вақти", "NZST": "Янги Зеландия стандарт вақти", "OESZ": "Шарқий Европа ёзги вақти", "OEZ": "Шарқий Европа стандарт вақти", "OMSST": "Омск ёзги вақти", "OMST": "Омск стандарт вақти", "PDT": "Шимолий Америка тинч океани кундузги вақти", "PDTM": "Meksika Tinch okeani yozgi vaqti", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Папуа-Янги Гвинея вақти", "PHOT": "Феникс ороллари вақти", "PKT": "Покистон стандарт вақти", "PKT DST": "Покистон ёзги вақти", "PMDT": "Сент-Пьер ва Микелон кундузги вақти", "PMST": "Сент-Пьер ва Микелон стандарт вақти", "PONT": "Понапе вақти", "PST": "Шимолий Америка тинч океани стандарт вақти", "PST Philippine": "Филиппин стандарт вақти", "PST Philippine DST": "Филиппин ёзги вақти", "PST Pitcairn": "Питкерн вақти", "PSTM": "Meksika Tinch okeani standart vaqti", "PWT": "Палау вақти", "PYST": "Парагвай ёзги вақти", "PYT": "Парагвай стандарт вақти", "PYT Korea": "Pxenyan vaqti", "RET": "Реюньон вақти", "ROTT": "Ротера вақти", "SAKST": "Сахалин ёзги вақти", "SAKT": "Сахалин стандарт вақти", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Жанубий Африка вақти", "SBT": "Соломон ороллари вақти", "SCT": "Сейшел ороллари вақти", "SGT": "Сингапур вақти", "SLST": "SLST", "SRT": "Суринам вақти", "SST Samoa": "Самоа стандарт вақти", "SST Samoa Apia": "Apia standart vaqti", "SST Samoa Apia DST": "Apia yozgi vaqti", "SST Samoa DST": "Самоа кундузги вақти", "SYOT": "Сьова вақти", "TAAF": "Француз жанубий ва Антарктика вақти", "TAHT": "Таити вақти", "TJT": "Тожикистон вақти", "TKT": "Токелау вақти", "TLT": "Шарқий Тимор вақти", "TMST": "Туркманистон ёзги вақти", "TMT": "Туркманистон стандарт вақти", "TOST": "Тонга ёзги вақти", "TOT": "Тонга стандарт вақти", "TVT": "Тувалу вақти", "TWT": "Тайпей стандарт вақти", "TWT DST": "Тайпей кундузги вақти", "ULAST": "Улан-Батор ёзги вақти", "ULAT": "Улан-Батор стандарт вақти", "UYST": "Уругвай ёзги вақти", "UYT": "Уругвай стандарт вақти", "UZT": "Ўзбекистон стандарт вақти", "UZT DST": "Ўзбекистон ёзги вақти", "VET": "Венесуэла вақти", "VLAST": "Владивосток ёзги вақти", "VLAT": "Владивосток стандарт вақти", "VOLST": "Волгоград ёзги вақти", "VOLT": "Волгоград стандарт вақти", "VOST": "Восток вақти", "VUT": "Вануату стандарт вақти", "VUT DST": "Вануату ёзги вақти", "WAKT": "Уэйк ороли вақти", "WARST": "Ғарбий Аргентина ёзги вақти", "WART": "Ғарбий Аргентина стандарт вақти", "WAST": "Ғарбий Африка вақти", "WAT": "Ғарбий Африка вақти", "WESZ": "Ғарбий Европа ёзги вақти", "WEZ": "Ғарбий Европа стандарт вақти", "WFT": "Уэллис ва Футуна вақти", "WGST": "Ғарбий Гренландия ёзги вақти", "WGT": "Ғарбий Гренландия стандарт вақти", "WIB": "Ғарбий Индонезия вақти", "WIT": "Шарқий Индонезия вақти", "WITA": "Марказий Индонезия вақти", "YAKST": "Якутск ёзги вақти", "YAKT": "Якутск стандарт вақти", "YEKST": "Екатеринбург ёзги вақти", "YEKT": "Екатеринбург стандарт вақти", "YST": "Yukon vaqti", "МСК": "Москва стандарт вақти", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Ғарбий Қозоғистон вақти", "شىعىش قازاق ەلى": "Шарқий Қозоғистон вақти", "قازاق ەلى": "Qozogʻiston vaqti", "قىرعىزستان": "Қирғизистон вақти", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Азор ёзги вақти"},
	}
}

// Locale returns the current translators string locale
func (uz *uz_Cyrl) Locale() string {
	return uz.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uz_Cyrl'
func (uz *uz_Cyrl) PluralsCardinal() []locales.PluralRule {
	return uz.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uz_Cyrl'
func (uz *uz_Cyrl) PluralsOrdinal() []locales.PluralRule {
	return uz.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uz_Cyrl'
func (uz *uz_Cyrl) PluralsRange() []locales.PluralRule {
	return uz.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Cyrl'
func (uz *uz_Cyrl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Cyrl'
func (uz *uz_Cyrl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uz_Cyrl'
func (uz *uz_Cyrl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uz.CardinalPluralRule(num1, v1)
	end := uz.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uz *uz_Cyrl) MonthAbbreviated(month time.Month) string {
	return uz.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uz *uz_Cyrl) MonthsAbbreviated() []string {
	return uz.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uz *uz_Cyrl) MonthNarrow(month time.Month) string {
	return uz.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uz *uz_Cyrl) MonthsNarrow() []string {
	return uz.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uz *uz_Cyrl) MonthWide(month time.Month) string {
	return uz.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uz *uz_Cyrl) MonthsWide() []string {
	return uz.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayAbbreviated(weekday time.Weekday) string {
	return uz.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uz *uz_Cyrl) WeekdaysAbbreviated() []string {
	return uz.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayNarrow(weekday time.Weekday) string {
	return uz.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uz *uz_Cyrl) WeekdaysNarrow() []string {
	return uz.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayShort(weekday time.Weekday) string {
	return uz.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uz *uz_Cyrl) WeekdaysShort() []string {
	return uz.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayWide(weekday time.Weekday) string {
	return uz.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uz *uz_Cyrl) WeekdaysWide() []string {
	return uz.daysWide
}

// Decimal returns the decimal point of number
func (uz *uz_Cyrl) Decimal() string {
	return uz.decimal
}

// Group returns the group of number
func (uz *uz_Cyrl) Group() string {
	return uz.group
}

// Group returns the minus sign of number
func (uz *uz_Cyrl) Minus() string {
	return uz.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uz_Cyrl' and handles both Whole and Real numbers based on 'v'
func (uz *uz_Cyrl) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uz_Cyrl' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uz *uz_Cyrl) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uz.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, uz.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Cyrl'
// in accounting notation.
func (uz *uz_Cyrl) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, uz.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, uz.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, uz.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, uz.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := uz.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
