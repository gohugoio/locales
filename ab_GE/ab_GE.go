package ab_GE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ab_GE struct {
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

// New returns a new instance of translator for the 'ab_GE' locale
func New() locales.Translator {
	return &ab_GE{
		locale:                 "ab_GE",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "Ажь", "Жəаб", "Хəажә", "Мш", "Лаҵ", "Рашә", "Ԥхынгә", "Нанҳә", "Цəыб", "Жьҭ", "Абҵ", "Ԥхынҷ"},
		monthsNarrow:           []string{"", "Жь", "Жə", "Хə", "М", "Л", "Р", "Гә", "Н", "Цə", "Ҭ", "Б", "Ҷ"},
		monthsWide:             []string{"", "Ажьырныҳəа", "Жəабран", "Хəажəкыра", "Мшаԥы", "Лаҵара", "Рашəара", "Ԥхынгəы", "Нанҳəа", "Цəыббра", "Жьҭаара", "Абҵара", "Ԥхынҷкәын"},
		daysAbbreviated:        []string{"Ам", "Ашә", "Аҩ", "Ах", "Аԥ", "Ахә", "Ас"},
		daysNarrow:             []string{"М", "Шә", "Ҩ", "Х", "Ԥ", "Хә", "С"},
		daysWide:               []string{"Амҽыша", "Ашәахьа", "Аҩаша", "Ахаша", "Аԥшьаша", "Ахәаша", "Асабша"},
		timezones:              map[string]string{"ACDT": "Агәҭантәи Австралиа, аԥхынтәи аамҭа", "ACST": "Акри аԥхынтәи аамҭа", "ACT": "Акри астандартә аамҭа", "ACWDT": "Агәҭантәи Австралиа, мраҭашәаратәи аԥхынтәи аамҭа", "ACWST": "Агәҭантәи Австралиа, мраҭашәаратәи астандартә аамҭа", "ADT": "Атлантикатәи аԥхынтәи аамҭа", "ADT Arabia": "Саудтәи Арабсҭан, аԥхынтәи аамҭа", "AEDT": "Мрагыларатәи Австралиа, аԥхынтәи аамҭа", "AEST": "Мрагыларатәи Австралиа, астандартә аамҭа", "AFT": "Афганистан", "AKDT": "Алиаска, аԥхынтәи аамҭа", "AKST": "Алиаска, астандартә аамҭа", "AMST": "Амазонка, аԥхынтәи аамҭа", "AMST Armenia": "Ермантәыла, аԥхынтәи аамҭа", "AMT": "Амазонка, астандартә аамҭа", "AMT Armenia": "Ермантәыла, астандартә аамҭа", "ANAST": "Анадыр аԥхынтәи аамҭа", "ANAT": "Анадыр астандартә аамҭа", "ARST": "Аргентина, аԥхынтәи аамҭа", "ART": "Аргентина, астандартә аамҭа", "AST": "Атлантикатәи астандартә аамҭа", "AST Arabia": "Саудтәи Арабсҭан, астандартә аамҭа", "AWDT": "Мраҭашәаратәи Австралиа, аԥхынтәи аамҭа", "AWST": "Мраҭашәаратәи Австралиа, астандартә аамҭа", "AZST": "Азербаиџьан, аԥхынтәи аамҭа", "AZT": "Азербаиџьан, астандартә аамҭа", "BDT Bangladesh": "Бангладеш, аԥхынтәи аамҭа", "BNT": "Брунеи-Даруссалам", "BOT": "Боливиа", "BRST": "Бразилиа, аԥхынтәи аамҭа", "BRT": "Бразилиа, астандартә аамҭа", "BST Bangladesh": "Бангладеш, астандартә аамҭа", "BT": "Бутан", "CAST": "Кеиси", "CAT": "Агәҭантәи Африка", "CCT": "Кокостәи ад-хақәа", "CDT": "Агәҭантәи Америка, аԥхынтәи аамҭа", "CHADT": "Чатем, аԥхынтәи аамҭа", "CHAST": "Чатем, астандартә аамҭа", "CHUT": "Трук", "CKT": "Кук идгьылбжьахақәа, астандартә аамҭа", "CKT DST": "Кук идгьылбжьахақәа, полуаԥхынтәи аамҭа", "CLST": "Чили, аԥхынтәи аамҭа", "CLT": "Чили, астандартә аамҭа", "COST": "Колумбиа, аԥхынтәи аамҭа", "COT": "Колумбиа, астандартә аамҭа", "CST": "Агәҭантәи Америка, астандартә аамҭа", "CST China": "Китаи, астандартә аамҭа", "CST China DST": "Китаи, аԥхынтәи аамҭа", "CVST": "Кабо-Верде, аԥхынтәи аамҭа", "CVT": "Кабо-Верде, астандартә аамҭа", "CXT": "ад-ха Қьырса", "ChST": "Чаморро", "ChST NMI": "Аҩадатәи Мариантәи ад-хақәа", "CuDT": "Куба, аԥхынтәи аамҭа", "CuST": "Куба, астандартә аамҭа", "DAVT": "Деивис", "DDUT": "Диумон-д’Иурвил", "EASST": "ад-ха Пасхи, аԥхынтәи аамҭа", "EAST": "ад-ха Пасхи, астандартә аамҭа", "EAT": "Мрагыларатәи Африка", "ECT": "Еквадор", "EDT": "Мрагыларатәи Америка, аԥхынтәи аамҭа", "EGDT": "Мрагыларатәи Гренландиа, аԥхынтәи аамҭа", "EGST": "Мрагыларатәи Гренландиа, стандарное времиа", "EST": "Мрагыларатәи Америка, астандартә аамҭа", "FEET": "Московатәи аамҭа", "FJT": "Фиџи, астандартә аамҭа", "FJT Summer": "Фиџи, аԥхынтәи аамҭа", "FKST": "Фолклендтәи ад-хақәа, аԥхынтәи аамҭа", "FKT": "Фолклендтәи ад-хақәа, астандартә аамҭа", "FNST": "Фернанду-ди-Норониа, аԥхынтәи аамҭа", "FNT": "Фернанду-ди-Норониа, астандартә аамҭа", "GALT": "Галапагостәи ад-хақәа", "GAMT": "Гамбе", "GEST": "Қырҭтәыла, аԥхынтәи аамҭа", "GET": "Қырҭтәыла, астандартә аамҭа", "GFT": "Францызтәи Гвиана", "GIT": "ад-хақәа Гилберта", "GMT": "Агринвич Ибжьаратәу Аамҭа", "GNSST": "GNSST", "GNST": "GNST", "GST": "Аџьамтә аӡыбжьахала", "GST Guam": "Гуам", "GYT": "Гаиана", "HADT": "Ҳаваи-алеуттәи астандартә аамҭа", "HAST": "Ҳаваи-алеуттәи астандартә аамҭа", "HKST": "Ҳонконг, аԥхынтәи аамҭа", "HKT": "Ҳонконг, астандартә аамҭа", "HOVST": "Ховд, аԥхынтәи аамҭа", "HOVT": "Ховд, астандартә аамҭа", "ICT": "Индокитаи", "IDT": "Израиль, аԥхынтәи аамҭа", "IOT": "Индиатәи аокеан", "IRKST": "Иркутск, аԥхынтәи аамҭа", "IRKT": "Иркутск, астандартә аамҭа", "IRST": "Иран, астандартә аамҭа", "IRST DST": "Иран, аԥхынтәи аамҭа", "IST": "Индиа", "IST Israel": "Израиль, астандартә аамҭа", "JDT": "Иапониа, аԥхынтәи аамҭа", "JST": "Иапониа, астандартә аамҭа", "KOST": "Косрае", "KRAST": "Красноиарск, аԥхынтәи аамҭа", "KRAT": "Красноиарск, астандартә аамҭа", "KST": "Кореиа, астандартә аамҭа", "KST DST": "Кореиа, аԥхынтәи аамҭа", "LHDT": "Лорд-Ҳау, аԥхынтәи аамҭа", "LHST": "Лорд-Ҳау, астандартә аамҭа", "LINT": "ад-хақәа Лаин", "MAGST": "Магадан, аԥхынтәи аамҭа", "MAGT": "Магадан, астандартә аамҭа", "MART": "Маркизтәи ад-хақәа", "MAWT": "Моусон", "MDT": "Аԥхынтәи ашьхатә аамҭа (Аҩадатәи Америка)", "MESZ": "Агәҭантәи Европа, аԥхынтәи аамҭа", "MEZ": "Агәҭантәи Европа, астандартә аамҭа", "MHT": "Маршаллтәи Адгьылбжьахақәа", "MMT": "Мианма", "MSD": "Москва, аԥхынтәи аамҭа", "MST": "Астандартә ашьхатә аамҭа (Аҩадатәи Америка)", "MUST": "Маврики, аԥхынтәи аамҭа", "MUT": "Маврики, астандартә аамҭа", "MVT": "Мальдив", "MYT": "Малаизиа", "NCT": "Каледониа ҿыц, астандартә аамҭа", "NDT": "Ниуфаундленд, аԥхынтәи аамҭа", "NDT New Caledonia": "Каледониа ҿыц, аԥхынтәи аамҭа", "NFDT": "Норфолк, аԥхынтәи аамҭа", "NFT": "Норфолк, астандартә аамҭа", "NOVST": "Новосибирск, аԥхынтәи аамҭа", "NOVT": "Новосибирск, астандартә аамҭа", "NPT": "Непал", "NRT": "Науру", "NST": "Ниуфаундленд, астандартә аамҭа", "NUT": "Ниуе", "NZDT": "Зеландиа ҿыц, аԥхынтәи аамҭа", "NZST": "Зеландиа ҿыц, астандартә аамҭа", "OESZ": "Мрагыларатәи Европа, аԥхынтәи аамҭа", "OEZ": "Мрагыларатәи Европа, астандартә аамҭа", "OMSST": "Омск, аԥхынтәи аамҭа", "OMST": "Омск, астандартә аамҭа", "PDT": "Аокеанҭынчтәи аԥхынтәи аамҭа", "PDTM": "Аокеанҭынчтәи амексикатә аԥхынтәи аамҭа", "PETDT": "Петропавловск-Камчаткатәи, аԥхынтәи аамҭа", "PETST": "Петропавловск-Камчаткатәи, астандартә аамҭа", "PGT": "Папуа — Гвинеиа ҿыц", "PHOT": "ад-хақәа Феникс", "PKT": "Пакистан, астандартә аамҭа", "PKT DST": "Пакистан, аԥхынтәи аамҭа", "PMDT": "Сен-Пиери Микелони, аԥхынтәи аамҭа", "PMST": "Сен-Пиери Микелони, астандартә аамҭа", "PONT": "Понпеи", "PST": "Аокеанҭынчтәи астандартә аамҭа", "PST Philippine": "Филиппин, астандартә аамҭа", "PST Philippine DST": "Филиппин, аԥхынтәи аамҭа", "PST Pitcairn": "Питкерн", "PSTM": "Аокеанҭынчтәи амексикатә астандартә аамҭа", "PWT": "Палау", "PYST": "Парагваи, аԥхынтәи аамҭа", "PYT": "Парагваи, астандартә аамҭа", "PYT Korea": "Пхениан", "RET": "Реиунон", "ROTT": "Ротера", "SAKST": "Сахалин, аԥхынтәи аамҭа", "SAKT": "Сахалин, астандартә аамҭа", "SAMST": "Самартәи аԥхынтәи аамҭа", "SAMT": "Самартәи астандартә аамҭа", "SAST": "Аладатәи Африка", "SBT": "Соломонтәи адгьылбжьахақәа", "SCT": "Сеишелтәи адгьылбжьахақәа", "SGT": "Сингапур", "SLST": "Шри-Ланка", "SRT": "Суринам", "SST Samoa": "Самоа, астандартә аамҭа", "SST Samoa Apia": "Апиа, астандартә аамҭа", "SST Samoa Apia DST": "Апиа, аԥхынтәи аамҭа", "SST Samoa DST": "Самоа, аԥхынтәи аамҭа", "SYOT": "Сиова", "TAAF": "Францызтәи Аладатәеи Антарктикатәи рҵакырадгьылқәа", "TAHT": "Таити", "TJT": "Таџьықсҭан", "TKT": "Токелау", "TLT": "Мрагыларатәи Тимор", "TMST": "Туркменистан, аԥхынтәи аамҭа", "TMT": "Туркменистан, астандартә аамҭа", "TOST": "Тонга, аԥхынтәи аамҭа", "TOT": "Тонга, астандартә аамҭа", "TVT": "Тувалу", "TWT": "Таиван, астандартә аамҭа", "TWT DST": "Таиван, аԥхынтәи аамҭа", "ULAST": "Улан-Батор, аԥхынтәи аамҭа", "ULAT": "Улан-Батор, астандартә аамҭа", "UYST": "Уругваи, аԥхынтәи аамҭа", "UYT": "Уругваи, астандартә аамҭа", "UZT": "Узбеқьисҭан, астандартә аамҭа", "UZT DST": "Узбеқьисҭан, аԥхынтәи аамҭа", "VET": "Венесуела", "VLAST": "Владивосток, аԥхынтәи аамҭа", "VLAT": "Владивосток, астандартә аамҭа", "VOLST": "Волгоград, аԥхынтәи аамҭа", "VOLT": "Волгоград, астандартә аамҭа", "VOST": "Амрагылара", "VUT": "Вануату, астандартә аамҭа", "VUT DST": "Вануату, аԥхынтәи аамҭа", "WAKT": "Уеик", "WARST": "Мраҭашәаратәи Аргентина, аԥхынтәи аамҭа", "WART": "Мраҭашәаратәи Аргентина, астандартә аамҭа", "WAST": "Мраҭашәаратәи Африка", "WAT": "Мраҭашәаратәи Африка", "WESZ": "Мраҭашәаратәи Европа, аԥхынтәи аамҭа", "WEZ": "Мраҭашәаратәи Европа, астандартә аамҭа", "WFT": "Уоллиси Футунеи", "WGST": "Мраҭашәаратәи Гренландиа, аԥхынтәи аамҭа", "WGT": "Мраҭашәаратәи Гренландиа, астандартә аамҭа", "WIB": "Мраҭашәаратәи Индонезиа", "WIT": "Мрагыларатәи Индонезиа", "WITA": "Агәҭантәи Индонезиа", "YAKST": "Иакутск, аԥхынтәи аамҭа", "YAKT": "Иакутск, астандартә аамҭа", "YEKST": "Екатеринбург, аԥхынтәи аамҭа", "YEKT": "Екатеринбург, астандартә аамҭа", "YST": "Иукон", "МСК": "Москва, астандартә аамҭа", "اقتاۋ": "Актау, астандартә аамҭа", "اقتاۋ قالاسى": "Актау аԥхынтәи аамҭа", "اقتوبە": "Актобе астандартә аамҭа", "اقتوبە قالاسى": "Актобе аԥхынтәи аамҭа", "الماتى": "Алма-Ата астандартә аамҭа", "الماتى قالاسى": "Алма-Ата аԥхынтәи аамҭа", "باتىس قازاق ەلى": "Мраҭашәаратәи Ҟазахсҭан", "شىعىش قازاق ەلى": "Мрагыларатәи Ҟазахсҭан", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Киргизиа", "قىزىلوردا": "Кызылорда, астандартә аамҭа*", "قىزىلوردا قالاسى": "Кызылорда, аԥхынтәи аамҭа*", "∅∅∅": "Перу, аԥхынтәи аамҭа"},
	}
}

// Locale returns the current translators string locale
func (ab *ab_GE) Locale() string {
	return ab.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ab_GE'
func (ab *ab_GE) PluralsCardinal() []locales.PluralRule {
	return ab.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ab_GE'
func (ab *ab_GE) PluralsOrdinal() []locales.PluralRule {
	return ab.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ab_GE'
func (ab *ab_GE) PluralsRange() []locales.PluralRule {
	return ab.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ab_GE'
func (ab *ab_GE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ab_GE'
func (ab *ab_GE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ab_GE'
func (ab *ab_GE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ab *ab_GE) MonthAbbreviated(month time.Month) string {
	return ab.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ab *ab_GE) MonthsAbbreviated() []string {
	return ab.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ab *ab_GE) MonthNarrow(month time.Month) string {
	return ab.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ab *ab_GE) MonthsNarrow() []string {
	return ab.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ab *ab_GE) MonthWide(month time.Month) string {
	return ab.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ab *ab_GE) MonthsWide() []string {
	return ab.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ab *ab_GE) WeekdayAbbreviated(weekday time.Weekday) string {
	return ab.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ab *ab_GE) WeekdaysAbbreviated() []string {
	return ab.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ab *ab_GE) WeekdayNarrow(weekday time.Weekday) string {
	return ab.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ab *ab_GE) WeekdaysNarrow() []string {
	return ab.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ab *ab_GE) WeekdayShort(weekday time.Weekday) string {
	return ab.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ab *ab_GE) WeekdaysShort() []string {
	return ab.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ab *ab_GE) WeekdayWide(weekday time.Weekday) string {
	return ab.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ab *ab_GE) WeekdaysWide() []string {
	return ab.daysWide
}

// Decimal returns the decimal point of number
func (ab *ab_GE) Decimal() string {
	return ab.decimal
}

// Group returns the group of number
func (ab *ab_GE) Group() string {
	return ab.group
}

// Group returns the minus sign of number
func (ab *ab_GE) Minus() string {
	return ab.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ab_GE' and handles both Whole and Real numbers based on 'v'
func (ab *ab_GE) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ab.group) - 1; j >= 0; j-- {
					b = append(b, ab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ab_GE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ab *ab_GE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ab.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ab.percentSuffix...)

	b = append(b, ab.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ab_GE'
func (ab *ab_GE) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ab.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ab.group) - 1; j >= 0; j-- {
					b = append(b, ab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ab.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ab.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ab_GE'
// in accounting notation.
func (ab *ab_GE) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ab.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ab.group) - 1; j >= 0; j-- {
					b = append(b, ab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ab.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ab.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ab.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ab.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd1, 0x88}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ab.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd1, 0x88}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ab.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ab.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd1, 0x88}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ab_GE'
func (ab *ab_GE) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ab.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
