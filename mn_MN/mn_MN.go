package mn_MN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mn_MN struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'mn_MN' locale
func New() locales.Translator {
	return &mn_MN{
		locale:                 "mn_MN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "1-р сар", "2-р сар", "3-р сар", "4-р сар", "5-р сар", "6-р сар", "7-р сар", "8-р сар", "9-р сар", "10-р сар", "11-р сар", "12-р сар"},
		monthsNarrow:           []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII"},
		monthsWide:             []string{"", "нэгдүгээр сар", "хоёрдугаар сар", "гуравдугаар сар", "дөрөвдүгээр сар", "тавдугаар сар", "зургаадугаар сар", "долоодугаар сар", "наймдугаар сар", "есдүгээр сар", "аравдугаар сар", "арван нэгдүгээр сар", "арван хоёрдугаар сар"},
		daysAbbreviated:        []string{"Ня", "Да", "Мя", "Лх", "Пү", "Ба", "Бя"},
		daysNarrow:             []string{"Ня", "Да", "Мя", "Лх", "Пү", "Ба", "Бя"},
		daysWide:               []string{"ням", "даваа", "мягмар", "лхагва", "пүрэв", "баасан", "бямба"},
		timezones:              map[string]string{"ACDT": "Төв Австралийн зуны цаг", "ACST": "Төв Австралийн стандарт цаг", "ACT": "ACT", "ACWDT": "Австралийн төв баруун эргийн зуны цаг", "ACWST": "Австралийн төв баруун эргийн стандарт цаг", "ADT": "Атлантын зуны цаг", "ADT Arabia": "Арабын зуны цаг", "AEDT": "Австралийн зүүн эргийн зуны цаг", "AEST": "Австралийн зүүн эргийн стандарт цаг", "AFT": "Афганистаны цаг", "AKDT": "Аляскийн зуны цаг", "AKST": "Аляскийн стандарт цаг", "AMST": "Амазоны зуны цаг", "AMST Armenia": "Арменийн зуны цаг", "AMT": "Амазоны стандарт цаг", "AMT Armenia": "Арменийн стандарт цаг", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентины зуны цаг", "ART": "Аргентины стандарт цаг", "AST": "Атлантын стандарт цаг", "AST Arabia": "Арабын стандарт цаг", "AWDT": "Австралийн баруун эргийн зуны цаг", "AWST": "Австралийн баруун эргийн стандарт цаг", "AZST": "Азербайжаны зуны цаг", "AZT": "Азербайжаны стандарт цаг", "BDT Bangladesh": "Бангладешийн зуны цаг", "BNT": "Бруней Даруссаламын цаг", "BOT": "Боливийн цаг", "BRST": "Бразилийн зуны цаг", "BRT": "Бразилийн стандарт цаг", "BST Bangladesh": "Бангладешийн стандарт цаг", "BT": "Бутаны цаг", "CAST": "CAST", "CAT": "Төв Африкийн цаг", "CCT": "Кокос арлын цаг", "CDT": "Төв зуны цаг", "CHADT": "Чатемын зуны цаг", "CHAST": "Чатемын стандарт цаг", "CHUT": "Чүүкийн цаг", "CKT": "Күүкийн арлуудын стандарт цаг", "CKT DST": "Күүкийн арлуудын хагас зуны цаг", "CLST": "Чилийн зуны цаг", "CLT": "Чилийн стандарт цаг", "COST": "Колумбын зуны цаг", "COT": "Колумбын стандарт цаг", "CST": "Төв стандарт цаг", "CST China": "Хятадын стандарт цаг", "CST China DST": "Хятадын зуны цаг", "CVST": "Кабо-Вердийн зуны цаг", "CVT": "Кабо-Вердийн стандарт цаг", "CXT": "Крисмас арлын цаг", "ChST": "Чаморрогийн цаг", "ChST NMI": "ChST NMI", "CuDT": "Кубын зуны цаг", "CuST": "Кубын стандарт цаг", "DAVT": "Дэвисийн цаг", "DDUT": "Дюмон д’Юрвилийн цаг", "EASST": "Зүүн Исландын зуны цаг", "EAST": "Зүүн Исландын стандарт цаг", "EAT": "Зүүн Африкийн цаг", "ECT": "Эквадорын цаг", "EDT": "Зүүн эргийн зуны цаг", "EGDT": "Зүүн Гренландын зуны цаг", "EGST": "Зүүн Гренландын стандарт цаг", "EST": "Зүүн эргийн стандарт цаг", "FEET": "Алс дорнод Европын цаг", "FJT": "Фижигийн стандарт цаг", "FJT Summer": "Фижигийн зуны цаг", "FKST": "Фолклендийн арлуудын зуны цаг", "FKT": "Фолклендийн арлуудын стандарт цаг", "FNST": "Фернандо де Норонагийн зуны цаг", "FNT": "Фернандо де Норонагийн стандарт цаг", "GALT": "Галапагосын цаг", "GAMT": "Гамбьегийн цаг", "GEST": "Гүржийн зуны цаг", "GET": "Гүржийн стандарт цаг", "GFT": "Францын Гвианагийн цаг", "GIT": "Гильбертийн арлуудын цаг", "GMT": "Гринвичийн цаг", "GNSST": "GNSST", "GNST": "GNST", "GST": "Персийн булангийн цаг", "GST Guam": "GST Guam", "GYT": "Гайанагийн цаг", "HADT": "Хавай-Алеутын зуны цаг", "HAST": "Хавай-Алеутын стандарт цаг", "HKST": "Хонг Конгийн зуны цаг", "HKT": "Хонг Конгийн стандарт цаг", "HOVST": "Ховдын зуны цаг", "HOVT": "Ховдын стандарт цаг", "ICT": "Энэтхэг-Хятадын хойгийн цаг", "IDT": "Израилийн зуны цаг", "IOT": "Энэтхэгийн далайн цаг", "IRKST": "Эрхүүгийн зуны цаг", "IRKT": "Эрхүүгийн стандарт цаг", "IRST": "Ираны стандарт цаг", "IRST DST": "Ираны зуны цаг", "IST": "Энэтхэгийн цаг", "IST Israel": "Израилийн стандарт цаг", "JDT": "Японы зуны цаг", "JST": "Японы стандарт цаг", "KOST": "Косрэгийн цаг", "KRAST": "Красноярскийн зуны цаг", "KRAT": "Красноярскийн стандарт цаг", "KST": "Солонгосын стандарт цаг", "KST DST": "Солонгосын зуны цаг", "LHDT": "Лорд Хоугийн зуны цаг", "LHST": "Лорд Хоугийн стандарт цаг", "LINT": "Лайн арлуудын цаг", "MAGST": "Магаданы зуны цаг", "MAGT": "Магаданы стандарт цаг", "MART": "Маркезын арлуудын цаг", "MAWT": "Моусоны цаг", "MDT": "MDT", "MESZ": "Төв Европын зуны цаг", "MEZ": "Төв Европын стандарт цаг", "MHT": "Маршаллын арлуудын цаг", "MMT": "Мьянмарын цаг", "MSD": "Москвагийн зуны цаг", "MST": "MST", "MUST": "Маврикийн зуны цаг", "MUT": "Маврикийн стандарт цаг", "MVT": "Мальдивийн цаг", "MYT": "Малайзын цаг", "NCT": "Шинэ Каледонийн стандарт цаг", "NDT": "Нью-Фаундлендын зуны цаг", "NDT New Caledonia": "Шинэ Каледонийн зуны цаг", "NFDT": "Норфолк арлын зуны цаг", "NFT": "Норфолк арлын стандарт цаг", "NOVST": "Новосибирскийн зуны цаг", "NOVT": "Новосибирскийн стандарт цаг", "NPT": "Балбын цаг", "NRT": "Науругийн цаг", "NST": "Нью-Фаундлендын стандарт цаг", "NUT": "Ниуэгийн цаг", "NZDT": "Шинэ Зеландын зуны цаг", "NZST": "Шинэ Зеландын стандарт цаг", "OESZ": "Зүүн Европын зуны цаг", "OEZ": "Зүүн Европын стандарт цаг", "OMSST": "Омскийн зуны цаг", "OMST": "Омскийн стандарт цаг", "PDT": "Номхон далайн зуны цаг", "PDTM": "Мексик-Номхон далайн зуны цаг", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Папуа Шинэ Гвинейн цаг", "PHOT": "Феникс арлын цаг", "PKT": "Пакистаны стандарт цаг", "PKT DST": "Пакистаны зуны цаг", "PMDT": "Сент-Пьер ба Микелоны зуны цаг", "PMST": "Сент-Пьер ба Микелоны стандарт цаг", "PONT": "Понапегийн цаг", "PST": "Номхон далайн стандарт цаг", "PST Philippine": "Филиппиний стандарт цаг", "PST Philippine DST": "Филиппиний зуны цаг", "PST Pitcairn": "Питкернийн цаг", "PSTM": "Мексик-Номхон далайн стандарт цаг", "PWT": "Палаугийн цаг", "PYST": "Парагвайн зуны цаг", "PYT": "Парагвайн стандарт цаг", "PYT Korea": "Пёньяны цаг", "RET": "Реюнионы цаг", "ROTT": "Ротерагийн цаг", "SAKST": "Сахалины зуны цаг", "SAKT": "Сахалины стандарт цаг", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Өмнөд Африкийн стандарт цаг", "SBT": "Соломоны арлуудын цаг", "SCT": "Сейшелийн арлуудын цаг", "SGT": "Сингапурын цаг", "SLST": "SLST", "SRT": "Суринамын цаг", "SST Samoa": "Самоагийн стандарт цаг", "SST Samoa Apia": "Апиагийн стандарт цаг", "SST Samoa Apia DST": "Апиагийн зуны цаг", "SST Samoa DST": "Самоагийн зуны цаг", "SYOT": "Сёвагийн цаг", "TAAF": "Францын Өмнөд ба Антарктидийн цаг", "TAHT": "Таитигийн цаг", "TJT": "Тажикистаны цаг", "TKT": "Токелаугийн цаг", "TLT": "Зүүн Тиморын цаг", "TMST": "Туркменистаны зуны цаг", "TMT": "Туркменистаны стандарт цаг", "TOST": "Тонгагийн зуны цаг", "TOT": "Тонгагийн стандарт цаг", "TVT": "Тувалугийн цаг", "TWT": "Тайпейн стандарт цаг", "TWT DST": "Тайпейн зуны цаг", "ULAST": "Улаанбаатарын зуны цаг", "ULAT": "Улаанбаатарын стандарт цаг", "UYST": "Уругвайн зуны цаг", "UYT": "Уругвайн стандарт цаг", "UZT": "Узбекистаны стандарт цаг", "UZT DST": "Узбекистаны зуны цаг", "VET": "Венесуэлийн цаг", "VLAST": "Владивостокийн зуны цаг", "VLAT": "Владивостокийн стандарт цаг", "VOLST": "Волгоградын зуны цаг", "VOLT": "Волгоградын стандарт цаг", "VOST": "Востокийн цаг", "VUT": "Вануатугийн стандарт цаг", "VUT DST": "Вануатугийн зуны цаг", "WAKT": "Уэкийн арлуудын цаг", "WARST": "Баруун Аргентины зуны цаг", "WART": "Баруун Аргентины стандарт цаг", "WAST": "Баруун Африкийн цаг", "WAT": "Баруун Африкийн цаг", "WESZ": "Баруун Европын зуны цаг", "WEZ": "Баруун Европын стандарт цаг", "WFT": "Уоллис ба Футунагийн цаг", "WGST": "Баруун Гренландын зуны цаг", "WGT": "Баруун Гренландын стандарт цаг", "WIB": "Баруун Индонезийн цаг", "WIT": "Зүүн Индонезийн цаг", "WITA": "Төв Индонезийн цаг", "YAKST": "Якутын зуны цаг", "YAKT": "Якутын стандарт цаг", "YEKST": "Екатеринбургийн зуны цаг", "YEKT": "Екатеринбургийн стандарт цаг", "YST": "Юкон цагийн бүс", "МСК": "Москвагийн стандарт цаг", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Баруун Казахстаны цаг", "شىعىش قازاق ەلى": "Зүүн Казахстаны цаг", "قازاق ەلى": "Казахстаны цаг", "قىرعىزستان": "Киргизийн цаг", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Азорын зуны цаг"},
	}
}

// Locale returns the current translators string locale
func (mn *mn_MN) Locale() string {
	return mn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mn_MN'
func (mn *mn_MN) PluralsCardinal() []locales.PluralRule {
	return mn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mn_MN'
func (mn *mn_MN) PluralsOrdinal() []locales.PluralRule {
	return mn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mn_MN'
func (mn *mn_MN) PluralsRange() []locales.PluralRule {
	return mn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mn_MN'
func (mn *mn_MN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mn_MN'
func (mn *mn_MN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mn_MN'
func (mn *mn_MN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := mn.CardinalPluralRule(num1, v1)
	end := mn.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mn *mn_MN) MonthAbbreviated(month time.Month) string {
	if len(mn.monthsAbbreviated) == 0 {
		return ""
	}
	return mn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mn *mn_MN) MonthsAbbreviated() []string {
	return mn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mn *mn_MN) MonthNarrow(month time.Month) string {
	if len(mn.monthsNarrow) == 0 {
		return ""
	}
	return mn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mn *mn_MN) MonthsNarrow() []string {
	return mn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mn *mn_MN) MonthWide(month time.Month) string {
	if len(mn.monthsWide) == 0 {
		return ""
	}
	return mn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mn *mn_MN) MonthsWide() []string {
	return mn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mn *mn_MN) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(mn.daysAbbreviated) == 0 {
		return ""
	}
	return mn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mn *mn_MN) WeekdaysAbbreviated() []string {
	return mn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mn *mn_MN) WeekdayNarrow(weekday time.Weekday) string {
	if len(mn.daysNarrow) == 0 {
		return ""
	}
	return mn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mn *mn_MN) WeekdaysNarrow() []string {
	return mn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mn *mn_MN) WeekdayShort(weekday time.Weekday) string {
	if len(mn.daysShort) == 0 {
		return ""
	}
	return mn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mn *mn_MN) WeekdaysShort() []string {
	return mn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mn *mn_MN) WeekdayWide(weekday time.Weekday) string {
	if len(mn.daysWide) == 0 {
		return ""
	}
	return mn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mn *mn_MN) WeekdaysWide() []string {
	return mn.daysWide
}

// Decimal returns the decimal point of number
func (mn *mn_MN) Decimal() string {
	return mn.decimal
}

// Group returns the group of number
func (mn *mn_MN) Group() string {
	return mn.group
}

// Group returns the minus sign of number
func (mn *mn_MN) Minus() string {
	return mn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mn_MN' and handles both Whole and Real numbers based on 'v'
func (mn *mn_MN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(mn.group) - 1; j >= 0; j-- {
					b = append(b, mn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mn_MN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mn *mn_MN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mn.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mn_MN'
func (mn *mn_MN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mn.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(mn.group) - 1; j >= 0; j-- {
					b = append(b, mn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(mn.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, mn.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, mn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mn_MN'
// in accounting notation.
func (mn *mn_MN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mn.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(mn.group) - 1; j >= 0; j-- {
					b = append(b, mn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(mn.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, mn.currencyNegativePrefix[j])
		}

		b = append(b, mn.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(mn.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, mn.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xbe, 0xd0, 0xbd, 0xd1, 0x8b}...)
	b = append(b, []byte{0x20}...)
	b = append(b, mn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xd1, 0x8b, 0xd0, 0xbd}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xbe, 0xd0, 0xbd, 0xd1, 0x8b}...)
	b = append(b, []byte{0x20}...)
	b = append(b, mn.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd1, 0x8b, 0xd0, 0xbd}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xbe, 0xd0, 0xbd, 0xd1, 0x8b}...)
	b = append(b, []byte{0x20}...)
	b = append(b, mn.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd1, 0x8b, 0xd0, 0xbd}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, mn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0xd0, 0xb3, 0xd0, 0xb0, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb3}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mn.timeSeparator...)

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

// FmtTimeFull returns the full time representation of 't' for 'mn_MN'
func (mn *mn_MN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := mn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
