package tt_RU

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type tt_RU struct {
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

// New returns a new instance of translator for the 'tt_RU' locale
func New() locales.Translator {
	return &tt_RU{
		locale:                 "tt_RU",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "гыйн.", "фев.", "мар.", "апр.", "май", "июнь", "июль", "авг.", "сент.", "окт.", "нояб.", "дек."},
		monthsWide:             []string{"", "гыйнвар", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
		daysAbbreviated:        []string{"якш.", "дүш.", "сиш.", "чәр.", "пәнҗ.", "җом.", "шим."},
		daysNarrow:             []string{"Я", "Д", "С", "Ч", "П", "Җ", "Ш"},
		daysWide:               []string{"якшәмбе", "дүшәмбе", "сишәмбе", "чәршәмбе", "пәнҗешәмбе", "җомга", "шимбә"},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "Австралия үзәк җәйге вакыты", "ACST": "Акр җәйге вакыты", "ACT": "Акр стандарт вакыты", "ACWDT": "Австралия үзәк көнбатыш җәйге вакыты", "ACWST": "Австралия үзәк көнбатыш стандарт вакыты", "ADT": "Төньяк Америка җәйге атлантик вакыты", "ADT Arabia": "Гарәп җәйге вакыты", "AEDT": "Көнчыгыш Австралия җәйге вакыты", "AEST": "Көнчыгыш Австралия стандарт вакыты", "AFT": "Әфганстан вакыты", "AKDT": "Аляска җәйге вакыты", "AKST": "Аляска стандарт вакыты", "AMST": "Амазонка җәйге вакыты", "AMST Armenia": "Армения җәйге вакыты", "AMT": "Амазонка стандарт вакыты", "AMT Armenia": "Армения стандарт вакыты", "ANAST": "Анадырь җәйге вакыты", "ANAT": "Анадырь стандарт вакыты", "ARST": "Аргентина җәйге вакыты", "ART": "Аргентина Стандарт вакыты", "AST": "Төньяк Америка гадәти атлантик вакыты", "AST Arabia": "Гарәп стандарт вакыты", "AWDT": "Көнбатыш Австралия җәйге вакыты", "AWST": "Көнбатыш Австралия стандарт вакыты", "AZST": "Әзербайҗан җәйге вакыты", "AZT": "Әзербайҗан стандарт вакыты", "BDT Bangladesh": "Бангладеш җәйге вакыты", "BNT": "Бруней-Даруссалам вакыты", "BOT": "Боливия вакыты", "BRST": "Бразилиа җәйге вакыты", "BRT": "Бразилиа стандарт вакыты", "BST Bangladesh": "Бангладеш стандарт вакыты", "BT": "Бутан вакыты", "CAST": "CAST", "CAT": "Үзәк Африка вакыты", "CCT": "Кокос утраулары вакыты", "CDT": "Төньяк Америка җәйге үзәк вакыты", "CHADT": "Чатем җәйге вакыты", "CHAST": "Чатем стандарт вакыты", "CHUT": "Чуук вакыты", "CKT": "Кук утраулары стандарт вакыты", "CKT DST": "Кук утраулары уртача җәйге вакыты", "CLST": "Чили җәйге вакыты", "CLT": "Чили стандарт вакыты", "COST": "Колумбия җәйге вакыты", "COT": "Колумбия стандарт вакыты", "CST": "Төньяк Америка гадәти үзәк вакыты", "CST China": "Кытай стандарт вакыты", "CST China DST": "Кытай җәйге вакыты", "CVST": "Кабо-Верде җәйге вакыты", "CVT": "Кабо-Верде стандарт вакыты", "CXT": "Раштуа утравы вакыты", "ChST": "Чаморро стандарт вакыты", "ChST NMI": "ChST NMI", "CuDT": "Куба җәйге вакыты", "CuST": "Куба стандарт вакыты", "DAVT": "Дэвис вакыты", "DDUT": "Дюмон д’Юрвиль вакыты", "EASST": "Пасха утравы җәйге вакыты", "EAST": "Пасха утравы стандарт вакыты", "EAT": "Көнчыгыш Африка вакыты", "ECT": "Эквадор вакыты", "EDT": "Төньяк Америка җәйге көнчыгыш вакыты", "EGDT": "Көнчыгыш Гренландия җәйге вакыты", "EGST": "Көнчыгыш Гренландия стандарт вакыты", "EST": "Төньяк Америка гадәти көнчыгыш вакыты", "FEET": "Ерак Көнчыгыш Европа вакыты", "FJT": "Фиджи стандарт вакыты", "FJT Summer": "Фиджи җәйге вакыты", "FKST": "Фолкленд утраулары җәйге вакыты", "FKT": "Фолкленд утраулары стандарт вакыты", "FNST": "Фернанду-ди-Норонья җәйге вакыты", "FNT": "Фернанду-ди-Норонья стандарт вакыты", "GALT": "Галапагос утраулары вакыты", "GAMT": "Гамбье вакыты", "GEST": "Грузия җәйге вакыты", "GET": "Грузия стандарт вакыты", "GFT": "Француз Гвиана вакыты", "GIT": "Гилберт утраулары вакыты", "GMT": "Гринвич уртача вакыты", "GNSST": "GNSST", "GNST": "GNST", "GST": "Көньяк Джорджия вакыты", "GST Guam": "GST Guam", "GYT": "Гайана вакыты", "HADT": "Гавай-Алеут җәйге вакыты", "HAST": "Гавай-Алеут стандарт вакыты", "HKST": "Гонконг җәйге вакыты", "HKT": "Гонконг стандарт вакыты", "HOVST": "Ховд җәйге вакыты", "HOVT": "Ховд стандарт вакыты", "ICT": "Һинд-кытай вакыты", "IDT": "Исраил җәйге вакыты", "IOT": "Һинд океаны вакыты", "IRKST": "Иркутск җәйге вакыты", "IRKT": "Иркутск стандарт вакыты", "IRST": "Иран стандарт вакыты", "IRST DST": "Иран җәйге вакыты", "IST": "Һинд стандарт вакыты", "IST Israel": "Исраил стандарт вакыты", "JDT": "Япон җәйге вакыты", "JST": "Япон стандарт вакыты", "KOST": "Косраэ вакыты", "KRAST": "Красноярск җәйге вакыты", "KRAT": "Красноярск стандарт вакыты", "KST": "Корея стандарт вакыты", "KST DST": "Корея җәйге вакыты", "LHDT": "Лорд Хау җәйге вакыты", "LHST": "Лорд Хау стандарт вакыты", "LINT": "Лайн утраулары вакыты", "MAGST": "Магадан җәйге вакыты", "MAGT": "Магадан стандарт вакыты", "MART": "Маркиз утраулары вакыты", "MAWT": "Моусон вакыты", "MDT": "Төньяк Америка җәйге тау вакыты", "MESZ": "җәйге Үзәк Европа вакыты", "MEZ": "гадәти Үзәк Европа вакыты", "MHT": "Маршалл утраулары вакыты", "MMT": "Мьянма вакыты", "MSD": "Мәскәү җәйге вакыты", "MST": "Төньяк Америка гадәти тау вакыты", "MUST": "Маврикий җәйге вакыты", "MUT": "Маврикий стандарт вакыты", "MVT": "Мальдив утраулары вакыты", "MYT": "Малайзия вакыты", "NCT": "Яңа Каледония стандарт вакыты", "NDT": "Ньюфаундленд җәйге вакыты", "NDT New Caledonia": "Яңа Каледония җәйге вакыты", "NFDT": "Норфолк утравы җәйге вакыты", "NFT": "Норфолк утравы стандарт вакыты", "NOVST": "Новосибирск җәйге вакыты", "NOVT": "Новосибирск стандарт вакыты", "NPT": "Непал вакыты", "NRT": "Науру вакыты", "NST": "Ньюфаундленд стандарт вакыты", "NUT": "Ниуэ вакыты", "NZDT": "Яңа Зеландия җәйге вакыты", "NZST": "Яңа Зеландия стандарт вакыты", "OESZ": "җәйге Көнчыгыш Европа вакыты", "OEZ": "гадәти Көнчыгыш Европа вакыты", "OMSST": "Омск җәйге вакыты", "OMST": "Омск стандарт вакыты", "PDT": "Төньяк Америка җәйге Тын океан вакыты", "PDTM": "Мексика Тыныч океан җәйге вакыты", "PETDT": "Петропавловск-Камчатский җәйге вакыты", "PETST": "Петропавловск-Камчатский стандарт вакыты", "PGT": "Папуа Яңа Гвинея вакыты", "PHOT": "Феникс утраулары вакыты", "PKT": "Пакистан стандарт вакыты", "PKT DST": "Пакистан җәйге вакыты", "PMDT": "Сен-Пьер һәм Микелон җәйге вакыты", "PMST": "Сен-Пьер һәм Микелон стандарт вакыты", "PONT": "Понапе вакыты", "PST": "Төньяк Америка гадәти Тын океан вакыты", "PST Philippine": "Филиппин стандарт вакыты", "PST Philippine DST": "Филиппин җәйге вакыты", "PST Pitcairn": "Питкэрн вакыты", "PSTM": "Мексика Тыныч океан стандартвакыты", "PWT": "Палау вакыты", "PYST": "Парагвай җәйге вакыты", "PYT": "Парагвай стандарт вакыты", "PYT Korea": "Пхеньян вакыты", "RET": "Реюньон вакыты", "ROTT": "Ротера вакыты", "SAKST": "Сахалин җәйге вакыты", "SAKT": "Сахалин стандарт вакыты", "SAMST": "Самара җәйге вакыты", "SAMT": "Самара стандарт вакыты", "SAST": "Көньяк Африка вакыты", "SBT": "Соломон утраулары вакыты", "SCT": "Сейшел утраулары вакыты", "SGT": "Сингапур стандарт вакыты", "SLST": "SLST", "SRT": "Суринам вакыты", "SST Samoa": "Самоа стандарт вакыты", "SST Samoa Apia": "Апиа стандарт вакыты", "SST Samoa Apia DST": "Апиа җәйге вакыты", "SST Samoa DST": "Самоа җәйге вакыты", "SYOT": "Сёва вакыты", "TAAF": "Француз көньяк һәм Антарктика вакыты", "TAHT": "Таити вакыты", "TJT": "Таҗикстан вакыты", "TKT": "Токелау вакыты", "TLT": "Көнчыгыш Тимор вакыты", "TMST": "Төркмәнстан җәйге вакыты", "TMT": "Төркмәнстан стандарт вакыты", "TOST": "Тонга җәйге вакыты", "TOT": "Тонга стандарт вакыты", "TVT": "Тувалу вакыты", "TWT": "Тайпей стандарт вакыты", "TWT DST": "Тайпей җәйге вакыты", "ULAST": "Улан-Батор җәйге вакыты", "ULAT": "Улан-Батор стандарт вакыты", "UYST": "Уругвай җәйге вакыты", "UYT": "Уругвай стандарт вакыты", "UZT": "Үзбәкстан стандарт вакыты", "UZT DST": "Үзбәкстан җәйге вакыты", "VET": "Венесуэла вакыты", "VLAST": "Владивосток җәйге вакыты", "VLAT": "Владивосток стандарт вакыты", "VOLST": "Волгоград җәйге вакыты", "VOLT": "Волгоград стандарт вакыты", "VOST": "Восток вакыты", "VUT": "Вануату стандарт вакыты", "VUT DST": "Вануату җәйге вакыты", "WAKT": "Уэйк утравы вакыты", "WARST": "Көнбатыш Аргентина җәйге вакыты", "WART": "Көнбатыш Аргентина стандарт вакыты", "WAST": "Көнбатыш Африка вакыты", "WAT": "Көнбатыш Африка вакыты", "WESZ": "җәйге Көнбатыш Европа вакыты", "WEZ": "гадәти Көнбатыш Европа вакыты", "WFT": "Уоллис һәм Футуна вакыты", "WGST": "Көнбатыш Гренландия җәйге вакыты", "WGT": "Көнбатыш Гренландия стандарт вакыты", "WIB": "Көнбатыш Индонезия вакыты", "WIT": "Көнчыгыш Индонезия вакыты", "WITA": "Үзәк Индонезия вакыты", "YAKST": "Якутск җәйге вакыты", "YAKT": "Якутск стандарт вакыты", "YEKST": "Екатеринбург җәйге вакыты", "YEKT": "Екатеринбург стандарт вакыты", "YST": "Юкон вакыты", "МСК": "Мәскәү стандарт вакыты", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Көнбатыш Казахстан вакыты", "شىعىش قازاق ەلى": "Көнчыгыш Казахстан вакыты", "قازاق ەلى": "Казахстан вакыты", "قىرعىزستان": "Кыргызстан вакыты", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Перу җәйге вакыты"},
	}
}

// Locale returns the current translators string locale
func (tt *tt_RU) Locale() string {
	return tt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'tt_RU'
func (tt *tt_RU) PluralsCardinal() []locales.PluralRule {
	return tt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'tt_RU'
func (tt *tt_RU) PluralsOrdinal() []locales.PluralRule {
	return tt.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'tt_RU'
func (tt *tt_RU) PluralsRange() []locales.PluralRule {
	return tt.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'tt_RU'
func (tt *tt_RU) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'tt_RU'
func (tt *tt_RU) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'tt_RU'
func (tt *tt_RU) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (tt *tt_RU) MonthAbbreviated(month time.Month) string {
	return tt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (tt *tt_RU) MonthsAbbreviated() []string {
	return tt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (tt *tt_RU) MonthNarrow(month time.Month) string {
	return tt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (tt *tt_RU) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (tt *tt_RU) MonthWide(month time.Month) string {
	return tt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (tt *tt_RU) MonthsWide() []string {
	return tt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (tt *tt_RU) WeekdayAbbreviated(weekday time.Weekday) string {
	return tt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (tt *tt_RU) WeekdaysAbbreviated() []string {
	return tt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (tt *tt_RU) WeekdayNarrow(weekday time.Weekday) string {
	return tt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (tt *tt_RU) WeekdaysNarrow() []string {
	return tt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (tt *tt_RU) WeekdayShort(weekday time.Weekday) string {
	return tt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (tt *tt_RU) WeekdaysShort() []string {
	return tt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (tt *tt_RU) WeekdayWide(weekday time.Weekday) string {
	return tt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (tt *tt_RU) WeekdaysWide() []string {
	return tt.daysWide
}

// Decimal returns the decimal point of number
func (tt *tt_RU) Decimal() string {
	return tt.decimal
}

// Group returns the group of number
func (tt *tt_RU) Group() string {
	return tt.group
}

// Group returns the minus sign of number
func (tt *tt_RU) Minus() string {
	return tt.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'tt_RU' and handles both Whole and Real numbers based on 'v'
func (tt *tt_RU) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'tt_RU' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (tt *tt_RU) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, tt.percentSuffix...)

	b = append(b, tt.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'tt_RU'
func (tt *tt_RU) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tt.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tt.group) - 1; j >= 0; j-- {
					b = append(b, tt.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, tt.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'tt_RU'
// in accounting notation.
func (tt *tt_RU) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tt.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tt.group) - 1; j >= 0; j-- {
					b = append(b, tt.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, tt.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, tt.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, tt.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tt.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb5, 0xd0, 0xbb}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb5, 0xd0, 0xbb}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb5, 0xd0, 0xbb}...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, tt.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'tt_RU'
func (tt *tt_RU) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := tt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
