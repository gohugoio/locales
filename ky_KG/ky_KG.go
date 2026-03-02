package ky_KG

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ky_KG struct {
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

// New returns a new instance of translator for the 'ky_KG' locale
func New() locales.Translator {
	return &ky_KG{
		locale:                 "ky_KG",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "янв.", "фев.", "мар.", "апр.", "май", "июн.", "июл.", "авг.", "сен.", "окт.", "ноя.", "дек."},
		monthsNarrow:           []string{"", "Я", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
		daysAbbreviated:        []string{"жек.", "дүй.", "шейш.", "шарш.", "бейш.", "жума", "ишм."},
		daysNarrow:             []string{"Ж", "Д", "Ш", "Ш", "Б", "Ж", "И"},
		daysShort:              []string{"жш.", "дш.", "шш.", "шр.", "бш.", "жм.", "иш."},
		daysWide:               []string{"жекшемби", "дүйшөмбү", "шейшемби", "шаршемби", "бейшемби", "жума", "ишемби"},
		periodsAbbreviated:     []string{"тң", "тк"},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"таңкы", "түштөн кийинки"},
		erasAbbreviated:        []string{"б.з.ч.", "б.з."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"биздин заманга чейин", "биздин заман"},
		timezones:              map[string]string{"ACDT": "Австралия борбордук жайкы убактысы", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Австралия борбордук чыгыш жайкы убактысы", "ACWST": "Австралия борбордук батыш кышкы убакыты", "ADT": "Атлантика жайкы убактысы", "ADT Arabia": "Арабия жайкы убакыты", "AEDT": "Австралия чыгыш жайкы убактысы", "AEST": "Австралия чыгыш кышкы убакыты", "AFT": "Афганистан убактысы", "AKDT": "Аляска жайкы убактысы", "AKST": "Аляска кышкы убактысы", "AMST": "Амазон жайкы убактысы", "AMST Armenia": "Армения жайкы убактысы", "AMT": "Амазон кышкы убактысы", "AMT Armenia": "Армения кышкы убакыты", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентина жайкы убактысы", "ART": "Аргентина кышкы убактысы", "AST": "Атлантика кышкы убактысы", "AST Arabia": "Арабия кышкы убакыты", "AWDT": "Австралия батыш жайкы убактысы", "AWST": "Австралия батыш кышкы убакыты", "AZST": "Азербайжан жайкы убактысы", "AZT": "Азербайжан кышкы убакыты", "BDT Bangladesh": "Бангладеш жайкы убактысы", "BNT": "Бруней Даруссалам убактысы", "BOT": "Боливия убактысы", "BRST": "Бразилия жайкы убактысы", "BRT": "Бразилия кышкы убактысы", "BST Bangladesh": "Бангладеш кышкы убакыты", "BT": "Бутан убактысы", "CAST": "CAST", "CAT": "Борбордук Африка убактысы", "CCT": "Кокос аралдарынын убактысы", "CDT": "Түндүк Америка, борбордук жайкы убактысы", "CHADT": "Чатам жайкы убактысы", "CHAST": "Чатам кышкы убакыт", "CHUT": "Чуук убактысы", "CKT": "Кук аралдарынын кышкы убактысы", "CKT DST": "Кук аралдарынын жарым жайкы убактысы", "CLST": "Чили жайкы убактысы", "CLT": "Чили кышкы убактысы", "COST": "Колумбия жайкы убактысы", "COT": "Колумбия кышкы убактысы", "CST": "Түндүк Америка, борбордук кышкы убактысы", "CST China": "Кытай кышкы убакыты", "CST China DST": "Кытай жайкы убакыты", "CVST": "Капе Верде жайкы убактысы", "CVT": "Капе Верде кышкы убакыты", "CXT": "Крисмас аралынын убактысы", "ChST": "Чаморро убактысы", "ChST NMI": "ChST NMI", "CuDT": "Куба жайкы убактысы", "CuST": "Куба кышкы убактысы", "DAVT": "Дэвис убактысы", "DDUT": "Дюмон-д-Урвил убактысы", "EASST": "Истер аралынын жайкы убакыты", "EAST": "Истер аралынын кышкы убакыты", "EAT": "Чыгыш Африка убактысы", "ECT": "Экуадор убактысы", "EDT": "Түндүк Америка, чыгыш жайкы убактысы", "EGDT": "Чыгыш Гренландия жайкы убактысы", "EGST": "Чыгыш Гренландия кышкы убактысы", "EST": "Түндүк Америка, чыгыш кышкы убактысы", "FEET": "Калининград жана Минск убактысы", "FJT": "Фижи кышкы убакыты", "FJT Summer": "Фижи жайкы убактысы", "FKST": "Фолкленд аралдарынын жайкы убактысы", "FKT": "Фолкленд аралдарынын кышкы убакыты", "FNST": "Фернандо де Норонья жайкы убактысы", "FNT": "Фернандо де Норонья кышкы убактысы", "GALT": "Галапагос убактысы", "GAMT": "Гамбие убактысы", "GEST": "Грузия жайкы убактысы", "GET": "Грузия кышкы убакыты", "GFT": "Француз Гвиана убактысы", "GIT": "Гилберт убактысы", "GMT": "Гринвич боюнча орточо убакыт", "GNSST": "GNSST", "GNST": "GNST", "GST": "Булуңдун стандарттык убакыты", "GST Guam": "GST Guam", "GYT": "Гвиана убактысы", "HADT": "Гавайи-Алеут кышкы убактысы", "HAST": "Гавайи-Алеут кышкы убактысы", "HKST": "Гонконг жайкы убактысы", "HKT": "Гонконг кышкы убакыты", "HOVST": "Ховд жайкы убактысы", "HOVT": "Ховд кышкы убакыты", "ICT": "Индокытай убактысы", "IDT": "Израиль жайкы убакыты", "IOT": "Инди океан убактысы", "IRKST": "Иркутск жайкы убакыты", "IRKT": "Иркутск кышкы убакыты", "IRST": "Иран кышкы убакыты", "IRST DST": "Иран күндүзгү убактысы", "IST": "Индия убактысы", "IST Israel": "Израиль кышкы убакыты", "JDT": "Жапон жайкы убактысы", "JST": "Жапон кышкы убакыты", "KOST": "Косрае убактысы", "KRAST": "Красноярск жайкы убактысы", "KRAT": "Красноярск кышкы убакыты", "KST": "Корея кышкы убакыты", "KST DST": "Корея жайкы убактысы", "LHDT": "Лорд Хау жайкы убактысы", "LHST": "Лорд Хау кышкы убакыты", "LINT": "Лайн аралдарынын убактысы", "MAGST": "Магадан жайкы убактысы", "MAGT": "Магадан кышкы убакыты", "MART": "Маркезас убактысы", "MAWT": "Моусон убактысы", "MDT": "MDT", "MESZ": "Борбордук Европа жайкы убактысы", "MEZ": "Борбордук Европа кышкы убакыты", "MHT": "Маршалл аралдарынын убактысы", "MMT": "Мйанмар убактысы", "MSD": "Москва жайкы убактысы", "MST": "MST", "MUST": "Маврикий жайкы убактысы", "MUT": "Маврикий кышкы убакыты", "MVT": "Мальдив убактысы", "MYT": "Малайзия убактысы", "NCT": "Жаңы Каледония кышкы убактысы", "NDT": "Нюфаундлэнд жайкы убактысы", "NDT New Caledonia": "Жаңы Каледония жайкы убактысы", "NFDT": "Норфолк жайкы убактысы", "NFT": "Норфолк кышкы убактысы", "NOVST": "Новосибирск жайкы убактысы", "NOVT": "Новосибирск кышкы убакыты", "NPT": "Непал убактысы", "NRT": "Науру убактысы", "NST": "Нюфаундлэнд кышкы убактысы", "NUT": "Ниуэ убактысы", "NZDT": "Жаңы Зеландия жайкы убакыты", "NZST": "Жаңы Зеландия кышкы убакыты", "OESZ": "Чыгыш Европа жайкы убактысы", "OEZ": "Чыгыш Европа кышкы убакыты", "OMSST": "Омск жайкы убактысы", "OMST": "Омск кышкы убакыты", "PDT": "Түндүк Америка, Тынч океан жайкы убактысы", "PDTM": "Мексика, Тынч океан жайкы убактысы", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Папуа-Жаңы Гвинея убакыты", "PHOT": "Феникс аралдарынын убактысы", "PKT": "Пакистан кышкы убакыты", "PKT DST": "Пакистан жайкы убактысы", "PMDT": "Сен Пьер жана Микелон жайкы убактысы", "PMST": "Сен Пьер жана Микелон кышкы убактысы", "PONT": "Понапе убактысы", "PST": "Түндүк Америка, Тынч океан кышкы убактысы", "PST Philippine": "Филиппин аралдарынын кышкы убактысы", "PST Philippine DST": "Филиппин аралдарынын жайкы убактысы", "PST Pitcairn": "Питкэрнг убактысы", "PSTM": "Мексика, Тынч океан кышкы убактысы", "PWT": "Палау убактысы", "PYST": "Парагвай жайкы убактысы", "PYT": "Парагвай кышкы убактысы", "PYT Korea": "Пхеньян убакыты", "RET": "Реюнион убактысы", "ROTT": "Ротера убактысы", "SAKST": "Сахалин жайкы убактысы", "SAKT": "Сахалин кышкы убакыты", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Түштүк Африка убактысы", "SBT": "Соломон аралдарынын убактысы", "SCT": "Сейшел убактысы", "SGT": "Сингапур убактысы", "SLST": "SLST", "SRT": "Суринаме убактысы", "SST Samoa": "Самоа кышкы убактысы", "SST Samoa Apia": "Апиа кышкы убактысы", "SST Samoa Apia DST": "Апиа жайкы убактысы", "SST Samoa DST": "Самоа жайкы убактысы", "SYOT": "Саоа убактысы", "TAAF": "Француз Түштүгү жана Антарктика убактысы", "TAHT": "Таити убактысы", "TJT": "Тажикстан убактысы", "TKT": "Токелау убактысы", "TLT": "Чыгыш Тимор убактысы", "TMST": "Түркмөнстан жайкы убактысы", "TMT": "Түркмөнстан кышкы убакыты", "TOST": "Тонга жайкы убактысы", "TOT": "Тонга кышкы убактысы", "TVT": "Тувалу убактысы", "TWT": "Тайпей кышкы убакыты", "TWT DST": "Тайпей жайкы убакыты", "ULAST": "Улан Батор жайкы убактысы", "ULAT": "Улан Батор кышкы убакыты", "UYST": "Уругвай жайкы убактысы", "UYT": "Уругвай кышкы убактысы", "UZT": "Өзбекстан кышкы убакыты", "UZT DST": "Өзбекстан жайкы убактысы", "VET": "Венесуэла убактысы", "VLAST": "Владивосток жайкы убактысы", "VLAT": "Владивосток кышкы убакыты", "VOLST": "Волгоград жайкы убактысы", "VOLT": "Волгоград кышкы убакыты", "VOST": "Восток убактысы", "VUT": "Вануату кышкы убакыты", "VUT DST": "Вануату жайкы убактысы", "WAKT": "Уейк аралдарынын убактысы", "WARST": "Батыш Аргентина жайкы убактысы", "WART": "Батыш Аргентина кышкы убактысы", "WAST": "Батыш Африка убактысы", "WAT": "Батыш Африка убактысы", "WESZ": "Батыш Европа жайкы убактысы", "WEZ": "Батыш Европа кышкы убакыты", "WFT": "Уолис жана Футуна убактысы", "WGST": "Батыш Гренландия жайкы убактысы", "WGT": "Батыш Гренландия кышкы убактысы", "WIB": "Батыш Индонезия убактысы", "WIT": "Чыгыш Индонезия убактысы", "WITA": "Борбордук Индонезия убактысы", "YAKST": "Якутск жайкы убактысы", "YAKT": "Якутск кышкы убакыты", "YEKST": "Екатеринбург жайкы убактысы", "YEKT": "Екатеринбург кышкы убакыты", "YST": "Юкон убактысы", "МСК": "Москва кышкы убакыты", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Батыш Казакстан убактысы", "شىعىش قازاق ەلى": "Чыгыш Казакстан убактысы", "قازاق ەلى": "Казакстан убактысы", "قىرعىزستان": "Кыргызстан убактысы", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Перу жайкы убактысы"},
	}
}

// Locale returns the current translators string locale
func (ky *ky_KG) Locale() string {
	return ky.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ky_KG'
func (ky *ky_KG) PluralsCardinal() []locales.PluralRule {
	return ky.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ky_KG'
func (ky *ky_KG) PluralsOrdinal() []locales.PluralRule {
	return ky.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ky_KG'
func (ky *ky_KG) PluralsRange() []locales.PluralRule {
	return ky.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ky_KG'
func (ky *ky_KG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ky_KG'
func (ky *ky_KG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ky_KG'
func (ky *ky_KG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ky.CardinalPluralRule(num1, v1)
	end := ky.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ky *ky_KG) MonthAbbreviated(month time.Month) string {
	return ky.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ky *ky_KG) MonthsAbbreviated() []string {
	return ky.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ky *ky_KG) MonthNarrow(month time.Month) string {
	return ky.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ky *ky_KG) MonthsNarrow() []string {
	return ky.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ky *ky_KG) MonthWide(month time.Month) string {
	return ky.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ky *ky_KG) MonthsWide() []string {
	return ky.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ky *ky_KG) WeekdayAbbreviated(weekday time.Weekday) string {
	return ky.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ky *ky_KG) WeekdaysAbbreviated() []string {
	return ky.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ky *ky_KG) WeekdayNarrow(weekday time.Weekday) string {
	return ky.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ky *ky_KG) WeekdaysNarrow() []string {
	return ky.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ky *ky_KG) WeekdayShort(weekday time.Weekday) string {
	return ky.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ky *ky_KG) WeekdaysShort() []string {
	return ky.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ky *ky_KG) WeekdayWide(weekday time.Weekday) string {
	return ky.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ky *ky_KG) WeekdaysWide() []string {
	return ky.daysWide
}

// Decimal returns the decimal point of number
func (ky *ky_KG) Decimal() string {
	return ky.decimal
}

// Group returns the group of number
func (ky *ky_KG) Group() string {
	return ky.group
}

// Group returns the minus sign of number
func (ky *ky_KG) Minus() string {
	return ky.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ky_KG' and handles both Whole and Real numbers based on 'v'
func (ky *ky_KG) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ky.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ky.group) - 1; j >= 0; j-- {
					b = append(b, ky.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ky.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ky_KG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ky *ky_KG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ky.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ky.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ky.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ky_KG'
func (ky *ky_KG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ky.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ky.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ky.group) - 1; j >= 0; j-- {
					b = append(b, ky.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ky.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ky.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ky.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ky_KG'
// in accounting notation.
func (ky *ky_KG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ky.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ky.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ky.group) - 1; j >= 0; j-- {
					b = append(b, ky.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ky.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ky.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ky.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ky.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, ky.monthsAbbreviated[t.Month()]...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, ky.monthsWide[t.Month()]...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, ky.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ky.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ky_KG'
func (ky *ky_KG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ky.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ky.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
