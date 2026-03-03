package bg

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type bg struct {
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

// New returns a new instance of translator for the 'bg' locale
func New() locales.Translator {
	return &bg{
		locale:                 "bg",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "Af", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "лв.", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "щ.д.", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "яну", "фев", "март", "апр", "май", "юни", "юли", "авг", "сеп", "окт", "ное", "дек"},
		monthsNarrow:           []string{"", "я", "ф", "м", "а", "м", "ю", "ю", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "януари", "февруари", "март", "април", "май", "юни", "юли", "август", "септември", "октомври", "ноември", "декември"},
		daysAbbreviated:        []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysNarrow:             []string{"н", "п", "в", "с", "ч", "п", "с"},
		daysWide:               []string{"неделя", "понеделник", "вторник", "сряда", "четвъртък", "петък", "събота"},
		timezones:              map[string]string{"ACDT": "Централноавстралийско лятно часово време", "ACST": "Централноавстралийско стандартно време", "ACT": "ACT", "ACWDT": "Австралия – западно централно лятно часово време", "ACWST": "Австралия – западно централно стандартно време", "ADT": "Северноамериканско атлантическо лятно часово време", "ADT Arabia": "Арабско лятно часово време", "AEDT": "Източноавстралийско лятно часово време", "AEST": "Източноавстралийско стандартно време", "AFT": "Афганистанско време", "AKDT": "Аляска – лятно часово време", "AKST": "Аляска – стандартно време", "AMST": "Амазонско лятно часово време", "AMST Armenia": "Арменско лятно часово време", "AMT": "Амазонско стандартно време", "AMT Armenia": "Арменско стандартно време", "ANAST": "Анадир – лятно часово време", "ANAT": "Анадир – стандартно време", "ARST": "Аржентинско лятно часово време", "ART": "Аржентинско стандартно време", "AST": "Северноамериканско атлантическо стандартно време", "AST Arabia": "Арабско стандартно време", "AWDT": "Западноавстралийско лятно часово време", "AWST": "Западноавстралийско стандартно време", "AZST": "Азербайджанско лятно часово време", "AZT": "Азербайджанско стандартно време", "BDT Bangladesh": "Бангладешко лятно часово време", "BNT": "Бруней Даруссалам", "BOT": "Боливийско време", "BRST": "Бразилско лятно часово време", "BRT": "Бразилско стандартно време", "BST Bangladesh": "Бангладешко стандартно време", "BT": "Бутанско време", "CAST": "CAST", "CAT": "Централноафриканско време", "CCT": "Кокосови острови", "CDT": "Северноамериканско централно лятно часово време", "CHADT": "Чатъмско лятно часово време", "CHAST": "Чатъмско стандартно време", "CHUT": "Чуюк", "CKT": "Острови Кук – стандартно време", "CKT DST": "Острови Кук – лятно часово време", "CLST": "Чилийско лятно часово време", "CLT": "Чилийско стандартно време", "COST": "Колумбийско лятно часово време", "COT": "Колумбийско стандартно време", "CST": "Северноамериканско централно стандартно време", "CST China": "Китайско стандартно време", "CST China DST": "Китайско лятно часово време", "CVST": "Кабо Верде – лятно часово време", "CVT": "Кабо Верде – стандартно време", "CXT": "Остров Рождество", "ChST": "Чаморско време", "ChST NMI": "ChST NMI", "CuDT": "Кубинско лятно часово време", "CuST": "Кубинско стандартно време", "DAVT": "Дейвис", "DDUT": "Дюмон Дюрвил", "EASST": "Великденски остров – лятно часово време", "EAST": "Великденски остров – стандартно време", "EAT": "Източноафриканско време", "ECT": "Еквадорско време", "EDT": "Северноамериканско източно лятно часово време", "EGDT": "Източногренландско лятно часово време", "EGST": "Източногренландско стандартно време", "EST": "Северноамериканско източно стандартно време", "FEET": "Далечно източноевропейско време", "FJT": "Фиджийско стандартно време", "FJT Summer": "Фиджийско лятно часово време", "FKST": "Фолклендски острови – лятно часово време", "FKT": "Фолклендски острови – стандартно време", "FNST": "Фернандо де Нороня – лятно часово време", "FNT": "Фернандо де Нороня – стандартно време", "GALT": "Галапагоско време", "GAMT": "Гамбие", "GEST": "Грузинско лятно часово време", "GET": "Грузинско стандартно време", "GFT": "Френска Гвиана", "GIT": "Острови Гилбърт", "GMT": "Средно гринуичко време", "GNSST": "GNSST", "GNST": "GNST", "GST": "Южна Джорджия", "GST Guam": "GST Guam", "GYT": "Гаяна", "HADT": "Хавайско-алеутско лятно часово време", "HAST": "Хавайско-алеутско стандартно време", "HKST": "Хонконгско лятно часово време", "HKT": "Хонконгско стандартно време", "HOVST": "Ховдско лятно часово време", "HOVT": "Ховдско стандартно време", "ICT": "Индокитайско време", "IDT": "Израелско лятно часово време", "IOT": "Индийски океан", "IRKST": "Иркутско лятно часово време", "IRKT": "Иркутско стандартно време", "IRST": "Иранско стандартно време", "IRST DST": "Иранско лятно часово време", "IST": "Индийско време", "IST Israel": "Израелско стандартно време", "JDT": "Японско лятно часово време", "JST": "Японско стандартно време", "KOST": "Кошрай", "KRAST": "Красноярско лятно часово време", "KRAT": "Красноярско стандартно време", "KST": "Корейско стандартно време", "KST DST": "Корейско лятно часово време", "LHDT": "Лорд Хау – лятно часово време", "LHST": "Лорд Хау – стандартно време", "LINT": "Екваториални острови", "MAGST": "Магаданско лятно часово време", "MAGT": "Магаданско стандартно време", "MART": "Маркизки острови", "MAWT": "Моусън", "MDT": "Северноамериканско планинско лятно часово време", "MESZ": "Централноевропейско лятно часово време", "MEZ": "Централноевропейско стандартно време", "MHT": "Маршалови острови", "MMT": "Мианмарско време", "MSD": "Московско лятно часово време", "MST": "Северноамериканско планинско стандартно време", "MUST": "Мавриций – лятно часово време", "MUT": "Мавриций – стандартно време", "MVT": "Малдивско време", "MYT": "Малайзийско време", "NCT": "Новокаледонско стандартно време", "NDT": "Нюфаундлендско лятно часово време", "NDT New Caledonia": "Новокаледонско лятно часово време", "NFDT": "Норфолкско лятно часово време", "NFT": "Норфолкско стандартно време", "NOVST": "Новосибирско лятно часово време", "NOVT": "Новосибирско стандартно време", "NPT": "Непалско време", "NRT": "Науру", "NST": "Нюфаундлендско стандартно време", "NUT": "Ниуе", "NZDT": "Новозеландско лятно часово време", "NZST": "Новозеландско стандартно време", "OESZ": "Източноевропейско лятно часово време", "OEZ": "Източноевропейско стандартно време", "OMSST": "Омско лятно часово време", "OMST": "Омско стандартно време", "PDT": "Северноамериканско тихоокеанско лятно часово време", "PDTM": "Мексиканско тихоокеанско лятно часово време", "PETDT": "Петропавловск-Камчатски – лятно часово време", "PETST": "Петропавловск-Камчатски стандартно време", "PGT": "Папуа Нова Гвинея", "PHOT": "Острови Феникс", "PKT": "Пакистанско стандартно време", "PKT DST": "Пакистанско лятно часово време", "PMDT": "Сен Пиер и Микелон – лятно часово време", "PMST": "Сен Пиер и Микелон – стандартно време", "PONT": "Понапе", "PST": "Северноамериканско тихоокеанско стандартно време", "PST Philippine": "Филипинско стандартно време", "PST Philippine DST": "Филипинско лятно часово време", "PST Pitcairn": "Питкерн", "PSTM": "Мексиканско тихоокеанско стандартно време", "PWT": "Палау", "PYST": "Парагвайско лятно часово време", "PYT": "Парагвайско стандартно време", "PYT Korea": "Пхенянско време", "RET": "Реюнион", "ROTT": "Ротера", "SAKST": "Сахалинско лятно часово време", "SAKT": "Сахалинско стандартно време", "SAMST": "Самара – лятно часово време", "SAMT": "Самара – стандартно време", "SAST": "Южноафриканско време", "SBT": "Соломонови острови", "SCT": "Сейшели", "SGT": "Сингапурско време", "SLST": "SLST", "SRT": "Суринамско време", "SST Samoa": "Самоанско стандартно време", "SST Samoa Apia": "Апия – стандартно време", "SST Samoa Apia DST": "Апия – лятно часово време", "SST Samoa DST": "Самоанско лятно часово време", "SYOT": "Шова", "TAAF": "Френски южни и антарктически територии", "TAHT": "Таитянско време", "TJT": "Таджикистанско време", "TKT": "Токелау", "TLT": "Източнотиморско време", "TMST": "Туркменистанско лятно часово време", "TMT": "Туркменистанско стандартно време", "TOST": "Тонга – лятно часово време", "TOT": "Тонга – стандартно време", "TVT": "Тувалу", "TWT": "Тайпе – стандартно време", "TWT DST": "Тайпе – лятно часово време", "ULAST": "Уланбаторско лятно часово време", "ULAT": "Уланбаторско стандартно време", "UYST": "Уругвайско лятно часово време", "UYT": "Уругвайско стандартно време", "UZT": "Узбекистанско стандартно време", "UZT DST": "Узбекистанско лятно часово време", "VET": "Венецуелско време", "VLAST": "Владивостокско лятно часово време", "VLAT": "Владивостокско стандартно време", "VOLST": "Волгоградско лятно часово време", "VOLT": "Волгоградско стандартно време", "VOST": "Восток", "VUT": "Вануату – стандартно време", "VUT DST": "Вануату – лятно часово време", "WAKT": "Остров Уейк", "WARST": "Западноаржентинско лятно часово време", "WART": "Западноаржентинско стандартно време", "WAST": "Западноафриканско време", "WAT": "Западноафриканско време", "WESZ": "Западноевропейско лятно време", "WEZ": "Западноевропейско стандартно време", "WFT": "Уолис и Футуна", "WGST": "Западногренландско лятно часово време", "WGT": "Западногренландско стандартно време", "WIB": "Западноиндонезийско време", "WIT": "Източноиндонезийско време", "WITA": "Централноиндонезийско време", "YAKST": "Якутскско лятно часово време", "YAKT": "Якутскско стандартно време", "YEKST": "Екатеринбургско лятно часово време", "YEKT": "Екатеринбургско стандартно време", "YST": "Юкон", "МСК": "Московско стандартно време", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Западноказахстанско време", "شىعىش قازاق ەلى": "Източноказахстанско време", "قازاق ەلى": "Казахстанско време", "قىرعىزستان": "Киргизстанско време", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Перуанско лятно часово време"},
	}
}

// Locale returns the current translators string locale
func (bg *bg) Locale() string {
	return bg.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bg'
func (bg *bg) PluralsCardinal() []locales.PluralRule {
	return bg.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bg'
func (bg *bg) PluralsOrdinal() []locales.PluralRule {
	return bg.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bg'
func (bg *bg) PluralsRange() []locales.PluralRule {
	return bg.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bg'
func (bg *bg) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bg'
func (bg *bg) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bg'
func (bg *bg) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bg *bg) MonthAbbreviated(month time.Month) string {
	if len(bg.monthsAbbreviated) == 0 {
		return ""
	}
	return bg.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bg *bg) MonthsAbbreviated() []string {
	return bg.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bg *bg) MonthNarrow(month time.Month) string {
	if len(bg.monthsNarrow) == 0 {
		return ""
	}
	return bg.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bg *bg) MonthsNarrow() []string {
	return bg.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bg *bg) MonthWide(month time.Month) string {
	if len(bg.monthsWide) == 0 {
		return ""
	}
	return bg.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bg *bg) MonthsWide() []string {
	return bg.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bg *bg) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(bg.daysAbbreviated) == 0 {
		return ""
	}
	return bg.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bg *bg) WeekdaysAbbreviated() []string {
	return bg.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bg *bg) WeekdayNarrow(weekday time.Weekday) string {
	if len(bg.daysNarrow) == 0 {
		return ""
	}
	return bg.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bg *bg) WeekdaysNarrow() []string {
	return bg.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bg *bg) WeekdayShort(weekday time.Weekday) string {
	if len(bg.daysShort) == 0 {
		return ""
	}
	return bg.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bg *bg) WeekdaysShort() []string {
	return bg.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bg *bg) WeekdayWide(weekday time.Weekday) string {
	if len(bg.daysWide) == 0 {
		return ""
	}
	return bg.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bg *bg) WeekdaysWide() []string {
	return bg.daysWide
}

// Decimal returns the decimal point of number
func (bg *bg) Decimal() string {
	return bg.decimal
}

// Group returns the group of number
func (bg *bg) Group() string {
	return bg.group
}

// Group returns the minus sign of number
func (bg *bg) Minus() string {
	return bg.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bg' and handles both Whole and Real numbers based on 'v'
func (bg *bg) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(bg.group) - 1; j >= 0; j-- {
					b = append(b, bg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bg' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bg *bg) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bg.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bg'
func (bg *bg) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bg.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(bg.group) - 1; j >= 0; j-- {
					b = append(b, bg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bg.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bg.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bg'
// in accounting notation.
func (bg *bg) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bg.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(bg.group) - 1; j >= 0; j-- {
					b = append(b, bg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, bg.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bg.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bg.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bg.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bg'
func (bg *bg) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'bg'
func (bg *bg) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

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

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bg'
func (bg *bg) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bg.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'bg'
func (bg *bg) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bg.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bg.monthsWide[t.Month()]...)
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

// FmtTimeShort returns the short time representation of 't' for 'bg'
func (bg *bg) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bg'
func (bg *bg) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bg'
func (bg *bg) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xd1, 0x87}...)
	b = append(b, []byte{0x2e, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bg'
func (bg *bg) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xd1, 0x87}...)
	b = append(b, []byte{0x2e, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bg.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
