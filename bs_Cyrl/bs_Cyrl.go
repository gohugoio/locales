package bs_Cyrl

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type bs_Cyrl struct {
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

// New returns a new instance of translator for the 'bs_Cyrl' locale
func New() locales.Translator {
	return &bs_Cyrl{
		locale:                 "bs_Cyrl",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "КМ", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "Кч", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "зл", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "дин.", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "Тл", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "јан", "феб", "мар", "апр", "мај", "јун", "јул", "ауг", "сеп", "окт", "нов", "дец"},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануар", "фебруар", "март", "април", "мај", "јуни", "јули", "аугуст", "септембар", "октобар", "новембар", "децембар"},
		daysAbbreviated:        []string{"нед", "пон", "уто", "сри", "чет", "пет", "суб"},
		daysNarrow:             []string{"н", "п", "у", "с", "ч", "п", "с"},
		daysWide:               []string{"недјеља", "понедјељак", "уторак", "сриједа", "четвртак", "петак", "субота"},
		timezones:              map[string]string{"ACDT": "Аустралијско централно љетње рачунање времена", "ACST": "Аустралијско централно стандардно вријеме", "ACT": "Акре стандардно време", "ACWDT": "Аустралијско централно западно љетње рачунање времена", "ACWST": "Аустралијско централно западно стандардно вријеме", "ADT": "Атланско љетње рачунање времена", "ADT Arabia": "Арабијско љетње рачунање времена", "AEDT": "Аустралијско источно љетње рачунање времена", "AEST": "Аустралијско источно стандардно вријеме", "AFT": "Авганистан вријеме", "AKDT": "Аљаска љетње вријеме", "AKST": "Аљаска стандардно вријеме", "AMST": "Амазон љетње рачунање времена", "AMST Armenia": "Арменија љетње рачунање времена", "AMT": "Амазон стандардно вријеме", "AMT Armenia": "Арменија стандардно вријеме", "ANAST": "Анадир летње рачунање времена", "ANAT": "Анадир стандардно време", "ARST": "Аргентина љетње рачунање времена", "ART": "Аргентина стандардно вријеме", "AST": "Атланско стандардно вријеме", "AST Arabia": "Арабијско стандардно вријеме", "AWDT": "Аустралијско западно љетње рачунање времена", "AWST": "Аустралијско западно стандардно вријеме", "AZST": "Азербејџан љетње рачунање времена", "AZT": "Азербејџан стандардно вријеме", "BDT Bangladesh": "Бангладеш љетње рачунање времена", "BNT": "Брунеј Дарусалам вријеме", "BOT": "Боливија вријеме", "BRST": "Бразилија љетње рачунање времена", "BRT": "Бразилија стандардно вријеме", "BST Bangladesh": "Бангладеш стандардно вријеме", "BT": "Бутан вријеме", "CAST": "CAST", "CAT": "Централно-афричко вријеме", "CCT": "Кокос (Келинг) Острва вријеме", "CDT": "Централно љетње рачунање времена", "CHADT": "Чатам љетње рачунање времена", "CHAST": "Чатам стандардно вријеме", "CHUT": "Трук вријеме", "CKT": "Кукова острва стандардно вријеме", "CKT DST": "Кукова острва полу-љетње рачунање времена", "CLST": "Чиле љетње рачунање времена", "CLT": "Чиле стандардно вријеме", "COST": "Колумбија љетње рачунање времена", "COT": "Колумбија стандардно вријеме", "CST": "Централно стандардно вријеме", "CST China": "Кина стандардно вријеме", "CST China DST": "Кина љетње рачунање времена", "CVST": "Зеленортско љетње рачунање времена", "CVT": "Зеленортско стандардно вријеме", "CXT": "Божићна острва вријеме", "ChST": "Чаморо вријеме", "ChST NMI": "Северна Маријанска Острва време", "CuDT": "Куба љетње рачунање времена", "CuST": "Куба стандардно вријеме", "DAVT": "Дејвис вријеме", "DDUT": "Димон д’Урвил вријеме", "EASST": "Ускршња острва љетње рачунање времена", "EAST": "Ускршња острва стандардно вријеме", "EAT": "Источно-афричко вријеме", "ECT": "Еквадор вријеме", "EDT": "Источно љетње рачунање времена", "EGDT": "Источни Гренланд љетње рачунање времена", "EGST": "Источни Гренланд стандардно вријеме", "EST": "Источно стандардно вријеме", "FEET": "Даље источноевропско вријеме", "FJT": "Фиџи стандардно вријеме", "FJT Summer": "Фиџи љетње рачунање времена", "FKST": "Фолкландска Острва љетње рачунање времена", "FKT": "Фолкландска Острва стандардно вријеме", "FNST": "Фернандо де Нороња љетње рачунање времена", "FNT": "Фернандо де Нороња стандардно вријеме", "GALT": "Галапагос вријеме", "GAMT": "Гамбијер вријеме", "GEST": "Грузија љетње рачунање времена", "GET": "Грузија стандардно вријеме", "GFT": "Француска Гвајана вријеме", "GIT": "Гилберт острва вријеме", "GMT": "Гриничко средње вријеме", "GNSST": "GNSST", "GNST": "GNST", "GST": "Заливско вријеме", "GST Guam": "Гуам стандардно време", "GYT": "Гвајана вријеме", "HADT": "Хавајско-алеутско стандардно вријеме", "HAST": "Хавајско-алеутско стандардно вријеме", "HKST": "Хонгконшко љетње рачунање времена", "HKT": "Хонг Конг стандардно вријеме", "HOVST": "Ховд љетње рачунање времена", "HOVT": "Ховд стандардно вријеме", "ICT": "Индокина вријеме", "IDT": "Израелско љетње рачунање времена", "IOT": "Индијско океанско вријеме", "IRKST": "Иркуцк љетње рачунање времена", "IRKT": "Иркуцк стандардно вријеме", "IRST": "Иран стандардно вријеме", "IRST DST": "Иран љетње рачунање времена", "IST": "Индијско стандардно вријеме", "IST Israel": "Израелско стандардно вријеме", "JDT": "Јапанско љетње рачунање времена", "JST": "Јапанско стандардно вријеме", "KOST": "Кошре вријеме", "KRAST": "Краснојарск љетње рачунање времена", "KRAT": "Краснојарск стандардно вријеме", "KST": "Корејско стандардно вријеме", "KST DST": "Корејско љетње рачунање времена", "LHDT": "Лорд Хов љетње рачунање времена", "LHST": "Лорд Хов стандардно вријеме", "LINT": "Лине Острва вријеме", "MAGST": "Магадан љетње рачунање вемена", "MAGT": "Магадан стандардно вријеме", "MART": "Маркиз вријеме", "MAWT": "Мосон вријеме", "MDT": "Макао летње рачунање вемена", "MESZ": "Средњеевропско љетње рачунање времена", "MEZ": "Средњеевропско стандардно вријеме", "MHT": "Маршалска Острва вријеме", "MMT": "Мијанмар вријеме", "MSD": "Москва љетње рачунање времена", "MST": "Макао стандардно време", "MUST": "Маурицијус љетње рачунање времена", "MUT": "Маурицијус стандардно вријеме", "MVT": "Малдиви вријеме", "MYT": "Малезија вријеме", "NCT": "Нова Каледонија стандардно вријеме", "NDT": "Њуфаундленд љетње рачунање времена", "NDT New Caledonia": "Нова Каледонија љетње рачунање времена", "NFDT": "Норфолк Острво љетње рачунање вријеме", "NFT": "Норфолк Острво стандардно вријеме", "NOVST": "Новосибирск љетње рачунање времена", "NOVT": "Новосибирск стандардно вријеме", "NPT": "Непал вријеме", "NRT": "Науру вријеме", "NST": "Њуфаундленд стандардно вријеме", "NUT": "Ниуе вријеме", "NZDT": "Нови Зеланд љетње рачунање времена", "NZST": "Нови Зеланд стандардно вријеме", "OESZ": "Источноевропско љетње рачунање времена", "OEZ": "Источноевропско стандардно вријеме", "OMSST": "Омск љетње рачунање времена", "OMST": "Омск стандардно вријеме", "PDT": "Пацифичко љетње рачунање времена", "PDTM": "Мексичко пацифичко љетње рачунање времена", "PETDT": "Петропавловско-камчатско летње рачунање вемена", "PETST": "Петропавловско-камчатско стандардно време", "PGT": "Папуа Нова Гвинеја вријеме", "PHOT": "Феникс острва вријеме", "PKT": "Пакистан стандардно вријеме", "PKT DST": "Пакистан љетње рачунање времена", "PMDT": "Сен Пјер и Микелон љетње рачунање вемена", "PMST": "Сен Пјер и Микелон стандардно вријеме", "PONT": "Понапе вријеме", "PST": "Пацифичко стандардно вријеме", "PST Philippine": "Филипини стандардно вријеме", "PST Philippine DST": "Филипини љетње рачунање времена", "PST Pitcairn": "Питкерн вријеме", "PSTM": "Мексичко пацифичко стандардно вријеме", "PWT": "Палау вријеме", "PYST": "Парагвај љетње рачунање времена", "PYT": "Парагвај стандардно вријеме", "PYT Korea": "Пјонгјанг вријеме", "RET": "Реинион вријеме", "ROTT": "Ротера вријеме", "SAKST": "Сахалин љетње рачунање времена", "SAKT": "Сахалин стандардно вријеме", "SAMST": "Самара летње рачунање времена", "SAMT": "Самара стандардно време", "SAST": "Јужно-афричко вријеме", "SBT": "Соломонска Острва вријеме", "SCT": "Сејшели вријеме", "SGT": "Сингапур стандардно вријеме", "SLST": "Шри Ланка време", "SRT": "Суринам вријеме", "SST Samoa": "Самоа стандардно вријеме", "SST Samoa Apia": "Апија стандардно вријеме", "SST Samoa Apia DST": "Апија љетње рачунање времена", "SST Samoa DST": "Самоа љетње рачунање времена", "SYOT": "Шова вријеме", "TAAF": "Француско јужно и антарктичко вријеме", "TAHT": "Тахити вријеме", "TJT": "Таџикистан вријеме", "TKT": "Токелау вријеме", "TLT": "Тимор-Лесте вријеме", "TMST": "Туркменистан љетње рачунање времена", "TMT": "Туркменистан стандардно вријеме", "TOST": "Тонга љетње рачунање времена", "TOT": "Тонга стандардно вријеме", "TVT": "Тувалу вријеме", "TWT": "Тајпеј стандардно вријеме", "TWT DST": "Тајпеј љетње рачунање времена", "ULAST": "Улан Батор љетње рачунање времена", "ULAT": "Улан Батор стандардно вријеме", "UYST": "Уругвај љетње рачунање времена", "UYT": "Уругвај стандардно вријеме", "UZT": "Узбекистан стандардно вријеме", "UZT DST": "Узбекистан љетње рачунање времена", "VET": "Венецуела вријеме", "VLAST": "Владивосток љетње рачунање времена", "VLAT": "Владивосток стандардно вријеме", "VOLST": "Волгоград љетње рачунање времена", "VOLT": "Волгоград стандардно вријеме", "VOST": "Восток вријеме", "VUT": "Вануату стандардно вријеме", "VUT DST": "Вануату љетње рачунање времена", "WAKT": "Вејк острво вријеме", "WARST": "Западна Аргентина љетње рачунање времена", "WART": "Западна Аргентина стандардно вријеме", "WAST": "Западно-афричко вријеме", "WAT": "Западно-афричко вријеме", "WESZ": "Западноевропско љетње рачунање времена", "WEZ": "Западноевропско стандардно вријеме", "WFT": "Валис и Футуна Острва вријеме", "WGST": "Западни Гренланд љетње рачунање времена", "WGT": "Западни Гренланд стандардно вријеме", "WIB": "Западно-индонезијско вријеме", "WIT": "Источно-индонезијско вријеме", "WITA": "Централно-индонезијско вријеме", "YAKST": "Јакутск љетње рачунање времена", "YAKT": "Јакутск стандардно вријеме", "YEKST": "Јекатеринбург љетње рачунање времена", "YEKT": "Јекатеринбург стандардно вријеме", "YST": "Yukon Time", "МСК": "Москва стандардно вријеме", "اقتاۋ": "Акватау стандардно време", "اقتاۋ قالاسى": "Акватау летње рачунање времена", "اقتوبە": "Акутобе стандардно време", "اقتوبە قالاسى": "Акутобе летње рачунање времена", "الماتى": "Алмати стандардно време", "الماتى قالاسى": "Алмати летње рачунање времена", "باتىس قازاق ەلى": "Западно-казахстанско вријеме", "شىعىش قازاق ەلى": "Источно-казахстанско вријеме", "قازاق ەلى": "Kazakhstan Time", "قىرعىزستان": "Киргистан вријеме", "قىزىلوردا": "Кизилорда стандардно време", "قىزىلوردا قالاسى": "Кизилорда летње рачунање времена", "∅∅∅": "Перу љетње рачунање времена"},
	}
}

// Locale returns the current translators string locale
func (bs *bs_Cyrl) Locale() string {
	return bs.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bs_Cyrl'
func (bs *bs_Cyrl) PluralsCardinal() []locales.PluralRule {
	return bs.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bs_Cyrl'
func (bs *bs_Cyrl) PluralsOrdinal() []locales.PluralRule {
	return bs.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bs_Cyrl'
func (bs *bs_Cyrl) PluralsRange() []locales.PluralRule {
	return bs.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bs_Cyrl'
func (bs *bs_Cyrl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bs_Cyrl'
func (bs *bs_Cyrl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bs_Cyrl'
func (bs *bs_Cyrl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := bs.CardinalPluralRule(num1, v1)
	end := bs.CardinalPluralRule(num2, v2)

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
func (bs *bs_Cyrl) MonthAbbreviated(month time.Month) string {
	return bs.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bs *bs_Cyrl) MonthsAbbreviated() []string {
	return bs.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bs *bs_Cyrl) MonthNarrow(month time.Month) string {
	return bs.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bs *bs_Cyrl) MonthsNarrow() []string {
	return bs.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bs *bs_Cyrl) MonthWide(month time.Month) string {
	return bs.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bs *bs_Cyrl) MonthsWide() []string {
	return bs.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bs *bs_Cyrl) WeekdayAbbreviated(weekday time.Weekday) string {
	return bs.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bs *bs_Cyrl) WeekdaysAbbreviated() []string {
	return bs.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bs *bs_Cyrl) WeekdayNarrow(weekday time.Weekday) string {
	return bs.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bs *bs_Cyrl) WeekdaysNarrow() []string {
	return bs.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bs *bs_Cyrl) WeekdayShort(weekday time.Weekday) string {
	return bs.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bs *bs_Cyrl) WeekdaysShort() []string {
	return bs.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bs *bs_Cyrl) WeekdayWide(weekday time.Weekday) string {
	return bs.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bs *bs_Cyrl) WeekdaysWide() []string {
	return bs.daysWide
}

// Decimal returns the decimal point of number
func (bs *bs_Cyrl) Decimal() string {
	return bs.decimal
}

// Group returns the group of number
func (bs *bs_Cyrl) Group() string {
	return bs.group
}

// Group returns the minus sign of number
func (bs *bs_Cyrl) Minus() string {
	return bs.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bs_Cyrl' and handles both Whole and Real numbers based on 'v'
func (bs *bs_Cyrl) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bs_Cyrl' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bs *bs_Cyrl) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bs.percentSuffix...)

	b = append(b, bs.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bs.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bs_Cyrl'
// in accounting notation.
func (bs *bs_Cyrl) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bs.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bs.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtDateMedium(t time.Time) string {
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

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, bs.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bs_Cyrl'
func (bs *bs_Cyrl) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bs.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
