package sr

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sr struct {
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
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'sr' locale
func New() locales.Translator {
	return &sr{
		locale:                 "sr",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "КМ", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "ლ", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "јан", "феб", "мар", "апр", "мај", "јун", "јул", "авг", "сеп", "окт", "нов", "дец"},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
		daysAbbreviated:        []string{"нед", "пон", "уто", "сре", "чет", "пет", "суб"},
		daysNarrow:             []string{"н", "п", "у", "с", "ч", "п", "с"},
		daysShort:              []string{"не", "по", "ут", "ср", "че", "пе", "су"},
		daysWide:               []string{"недеља", "понедељак", "уторак", "среда", "четвртак", "петак", "субота"},
		periodsAbbreviated:     []string{"", ""},
		erasAbbreviated:        []string{"п. н. е.", "н. е."},
		erasNarrow:             []string{"п.н.е.", "н.е."},
		erasWide:               []string{"пре нове ере", "нове ере"},
		timezones:              map[string]string{"ACDT": "Аустралијско централно летње време", "ACST": "Аустралијско централно стандардно време", "ACT": "Акре стандардно време", "ACWDT": "Аустралијско централно западно летње време", "ACWST": "Аустралијско централно западно стандардно време", "ADT": "Атлантско летње време", "ADT Arabia": "Арабијско летње време", "AEDT": "Аустралијско источно летње време", "AEST": "Аустралијско источно стандардно време", "AFT": "Авганистан време", "AKDT": "Аљаска, летње време", "AKST": "Аљаска, стандардно време", "AMST": "Амазон, летње време", "AMST Armenia": "Јерменија, летње време", "AMT": "Амазон, стандардно време", "AMT Armenia": "Јерменија, стандардно време", "ANAST": "Анадир летње рачунање времена", "ANAT": "Анадир стандардно време", "ARST": "Аргентина, летње време", "ART": "Аргентина, стандардно време", "AST": "Атлантско стандардно време", "AST Arabia": "Арабијско стандардно време", "AWDT": "Аустралијско западно летње време", "AWST": "Аустралијско западно стандардно време", "AZST": "Азербејџан, летње време", "AZT": "Азербејџан, стандардно време", "BDT Bangladesh": "Бангладеш, летње време", "BNT": "Брунеј Дарусалум време", "BOT": "Боливија време", "BRST": "Бразилија, летње време", "BRT": "Бразилија, стандардно време", "BST Bangladesh": "Бангладеш, стандардно време", "BT": "Бутан време", "CAST": "CAST", "CAT": "Централно-афричко време", "CCT": "Кокос (Келинг) Острва време", "CDT": "Северноамеричко централно летње време", "CHADT": "Чатам, летње време", "CHAST": "Чатам, стандардно време", "CHUT": "Чуук време", "CKT": "Кукова острва, стандардно време", "CKT DST": "Кукова острва, полу-летње време", "CLST": "Чиле, летње време", "CLT": "Чиле, стандардно време", "COST": "Колумбија, летње време", "COT": "Колумбија, стандардно време", "CST": "Северноамеричко централно стандардно време", "CST China": "Кинеско стандардно време", "CST China DST": "Кина, летње време", "CVST": "Зеленортска Острва, летње време", "CVT": "Зеленортска Острва, стандардно време", "CXT": "Божићно острво време", "ChST": "Чаморо време", "ChST NMI": "Северна Маријанска Острва време", "CuDT": "Куба, летње време", "CuST": "Куба, стандардно време", "DAVT": "Дејвис време", "DDUT": "Димон д’Урвил време", "EASST": "Ускршња острва, летње време", "EAST": "Ускршња острва, стандардно време", "EAT": "Источно-афричко време", "ECT": "Еквадор време", "EDT": "Северноамеричко источно летње време", "EGDT": "Источни Гренланд, летње време", "EGST": "Источни Гренланд, стандардно време", "EST": "Северноамеричко источно стандардно време", "FEET": "Време даљег истока Европе", "FJT": "Фиџи, стандардно време", "FJT Summer": "Фиџи, летње време", "FKST": "Фолкландска Острва, летње време", "FKT": "Фолкландска Острва, стандардно време", "FNST": "Фернандо де Нороња, летње време", "FNT": "Фернандо де Нороња, стандардно време", "GALT": "Галапагос време", "GAMT": "Гамбије време", "GEST": "Грузија, летње време", "GET": "Грузија, стандардно време", "GFT": "Француска Гвајана време", "GIT": "Гилберт острва време", "GMT": "Средње време по Гриничу", "GNSST": "GNSST", "GNST": "GNST", "GST": "Заливско време", "GST Guam": "Гуам стандардно време", "GYT": "Гвајана време", "HADT": "Хавајско-алеутско стандардно време", "HAST": "Хавајско-алеутско стандардно време", "HKST": "Хонг Конг, летње време", "HKT": "Хонг Конг, стандардно време", "HOVST": "Ховд, летње време", "HOVT": "Ховд, стандардно време", "ICT": "Индокина време", "IDT": "Израелско летње време", "IOT": "Индијско океанско време", "IRKST": "Иркуцк, летње време", "IRKT": "Иркуцк, стандардно време", "IRST": "Иран, стандардно време", "IRST DST": "Иран, летње време", "IST": "Индијско стандардно време", "IST Israel": "Израелско стандардно време", "JDT": "Јапанско летње време", "JST": "Јапанско стандардно време", "KOST": "Кошре време", "KRAST": "Краснојарск, летње време", "KRAT": "Краснојарск, стандардно време", "KST": "Корејско стандардно време", "KST DST": "Корејско летње време", "LHDT": "Лорд Хов, летње време", "LHST": "Лорд Хов, стандардно време", "LINT": "Острва Лајн време", "MAGST": "Магадан, летње време", "MAGT": "Магадан, стандардно време", "MART": "Маркиз време", "MAWT": "Мосон време", "MDT": "Макао летње рачунање времена", "MESZ": "Средњеевропско летње време", "MEZ": "Средњеевропско стандардно време", "MHT": "Маршалска Острва време", "MMT": "Мијанмар време", "MSD": "Москва, летње време", "MST": "Макао стандардно време", "MUST": "Маурицијус, летње време", "MUT": "Маурицијус, стандардно време", "MVT": "Малдиви време", "MYT": "Малезија време", "NCT": "Нова Каледонија, стандардно време", "NDT": "Њуфаундленд, летње време", "NDT New Caledonia": "Нова Каледонија, летње време", "NFDT": "Норфолк Острво, летње време", "NFT": "Норфолк Острво, стандардно време", "NOVST": "Новосибирск, летње време", "NOVT": "Новосибирск, стандардно време", "NPT": "Непал време", "NRT": "Науру време", "NST": "Њуфаундленд, стандардно време", "NUT": "Ниуе време", "NZDT": "Нови Зеланд, летње време", "NZST": "Нови Зеланд, стандардно време", "OESZ": "Источноевропско летње време", "OEZ": "Источноевропско стандардно време", "OMSST": "Омск, летње време", "OMST": "Омск, стандардно време", "PDT": "Северноамеричко пацифичко летње време", "PDTM": "Мексички Пацифик, летње време", "PETDT": "Петропавловско-камчатско летње рачунање времена", "PETST": "Петропавловско-камчатско стандардно време", "PGT": "Папуа Нова Гвинеја време", "PHOT": "Феникс острва време", "PKT": "Пакистан, стандардно време", "PKT DST": "Пакистан, летње време", "PMDT": "Сен Пјер и Микелон, летње време", "PMST": "Сен Пјер и Микелон, стандардно време", "PONT": "Понпеј време", "PST": "Северноамеричко пацифичко стандардно време", "PST Philippine": "Филипини, стандардно време", "PST Philippine DST": "Филипини, летње време", "PST Pitcairn": "Питкерн време", "PSTM": "Мексички Пацифик, стандардно време", "PWT": "Палау време", "PYST": "Парагвај, летње време", "PYT": "Парагвај, стандардно време", "PYT Korea": "Пјонгјаншко време", "RET": "Реинион време", "ROTT": "Ротера време", "SAKST": "Сахалин, летње време", "SAKT": "Сахалин, стандардно време", "SAMST": "Самара летње рачунање времена", "SAMT": "Самара стандардно време", "SAST": "Јужно-афричко време", "SBT": "Соломонска Острва време", "SCT": "Сејшели време", "SGT": "Сингапур, стандардно време", "SLST": "Шри Ланка време", "SRT": "Суринам време", "SST Samoa": "Самоа, стандардно време", "SST Samoa Apia": "Апија, стандардно време", "SST Samoa Apia DST": "Апија, летње време", "SST Samoa DST": "Самоа, летње време", "SYOT": "Шова време", "TAAF": "Француско јужно и антарктичко време", "TAHT": "Тахити време", "TJT": "Таџикистан време", "TKT": "Токелау време", "TLT": "Источни тимор време", "TMST": "Туркменистан, летње време", "TMT": "Туркменистан, стандардно време", "TOST": "Тонга, летње време", "TOT": "Тонга, стандардно време", "TVT": "Тувалу време", "TWT": "Тајпеј, стандардно време", "TWT DST": "Тајпеј, летње време", "ULAST": "Улан Батор, летње време", "ULAT": "Улан Батор, стандардно време", "UYST": "Уругвај, летње време", "UYT": "Уругвај, стандардно време", "UZT": "Узбекистан, стандардно време", "UZT DST": "Узбекистан, летње време", "VET": "Венецуела време", "VLAST": "Владивосток, летње време", "VLAT": "Владивосток, стандардно време", "VOLST": "Волгоград, летње време", "VOLT": "Волгоград, стандардно време", "VOST": "Восток време", "VUT": "Вануату, стандардно време", "VUT DST": "Вануату, летње време", "WAKT": "Вејк острво време", "WARST": "Западна Аргентина, летње време", "WART": "Западна Аргентина, стандардно време", "WAST": "Западно-афричко време", "WAT": "Западно-афричко време", "WESZ": "Западноевропско летње време", "WEZ": "Западноевропско стандардно време", "WFT": "Валис и Футуна Острва време", "WGST": "Западни Гренланд, летње време", "WGT": "Западни Гренланд, стандардно време", "WIB": "Западно-индонезијско време", "WIT": "Источно-индонезијско време", "WITA": "Централно-индонезијско време", "YAKST": "Јакутск, летње време", "YAKT": "Јакутск, стандардно време", "YEKST": "Јекатеринбург, летње време", "YEKT": "Јекатеринбург, стандардно време", "YST": "Јукон", "МСК": "Москва, стандардно време", "اقتاۋ": "Акватау стандардно време", "اقتاۋ قالاسى": "Акватау летње рачунање времена", "اقتوبە": "Акутобе стандардно време", "اقتوبە قالاسى": "Акутобе летње рачунање времена", "الماتى": "Алмати стандардно време", "الماتى قالاسى": "Алмати летње рачунање времена", "باتىس قازاق ەلى": "Западно-казахстанско време", "شىعىش قازاق ەلى": "Источно-казахстанско време", "قازاق ەلى": "Казахстанско време", "قىرعىزستان": "Киргистан време", "قىزىلوردا": "Кизилорда стандардно време", "قىزىلوردا قالاسى": "Кизилорда летње рачунање времена", "∅∅∅": "Перу, летње време"},
	}
}

// Locale returns the current translators string locale
func (sr *sr) Locale() string {
	return sr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sr'
func (sr *sr) PluralsCardinal() []locales.PluralRule {
	return sr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sr'
func (sr *sr) PluralsOrdinal() []locales.PluralRule {
	return sr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sr'
func (sr *sr) PluralsRange() []locales.PluralRule {
	return sr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sr'
func (sr *sr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14)) || (fMod10 >= 2 && fMod10 <= 4 && (fMod100 < 12 || fMod100 > 14)) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sr'
func (sr *sr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sr'
func (sr *sr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := sr.CardinalPluralRule(num1, v1)
	end := sr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sr *sr) MonthAbbreviated(month time.Month) string {
	return sr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sr *sr) MonthsAbbreviated() []string {
	return sr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sr *sr) MonthNarrow(month time.Month) string {
	return sr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sr *sr) MonthsNarrow() []string {
	return sr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sr *sr) MonthWide(month time.Month) string {
	return sr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sr *sr) MonthsWide() []string {
	return sr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sr *sr) WeekdayAbbreviated(weekday time.Weekday) string {
	return sr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sr *sr) WeekdaysAbbreviated() []string {
	return sr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sr *sr) WeekdayNarrow(weekday time.Weekday) string {
	return sr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sr *sr) WeekdaysNarrow() []string {
	return sr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sr *sr) WeekdayShort(weekday time.Weekday) string {
	return sr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sr *sr) WeekdaysShort() []string {
	return sr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sr *sr) WeekdayWide(weekday time.Weekday) string {
	return sr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sr *sr) WeekdaysWide() []string {
	return sr.daysWide
}

// Decimal returns the decimal point of number
func (sr *sr) Decimal() string {
	return sr.decimal
}

// Group returns the group of number
func (sr *sr) Group() string {
	return sr.group
}

// Group returns the minus sign of number
func (sr *sr) Minus() string {
	return sr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sr' and handles both Whole and Real numbers based on 'v'
func (sr *sr) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sr *sr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sr'
func (sr *sr) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sr'
// in accounting notation.
func (sr *sr) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sr.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sr'
func (sr *sr) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sr'
func (sr *sr) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sr'
func (sr *sr) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sr'
func (sr *sr) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, sr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sr'
func (sr *sr) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sr'
func (sr *sr) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sr'
func (sr *sr) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sr'
func (sr *sr) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
