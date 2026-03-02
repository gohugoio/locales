package sr_Cyrl_BA

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sr_Cyrl_BA struct {
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

// New returns a new instance of translator for the 'sr_Cyrl_BA' locale
func New() locales.Translator {
	return &sr_Cyrl_BA{
		locale:                 "sr_Cyrl_BA",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "јан", "феб", "мар", "апр", "мај", "јун", "јул", "авг", "сеп", "окт", "нов", "дец"},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
		daysAbbreviated:        []string{"нед", "пон", "уто", "сри", "чет", "пет", "суб"},
		daysNarrow:             []string{"н", "п", "у", "с", "ч", "п", "с"},
		daysShort:              []string{"не", "по", "ут", "ср", "че", "пе", "су"},
		daysWide:               []string{"недјеља", "понедјељак", "уторак", "сриједа", "четвртак", "петак", "субота"},
		periodsAbbreviated:     []string{"прије\u202fподне", "по\u202fподне"},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"прије подне", "по подне"},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"прије нове ере", "нове ере"},
		timezones:              map[string]string{"ACDT": "Аустралијско централно љетње вријеме", "ACST": "Акре летње рачунање времена", "ACT": "Акре стандардно време", "ACWDT": "Аустралијско централно западно љетње вријеме", "ACWST": "Аустралијско централно западно стандардно вријеме", "ADT": "Атлантско љетње вријеме", "ADT Arabia": "Арабијско љетње вријеме", "AEDT": "Аустралијско источно љетње вријеме", "AEST": "Аустралијско источно стандардно вријеме", "AFT": "Авганистан вријеме", "AKDT": "Аљаска, љетње вријеме", "AKST": "Аљаска, стандардно вријеме", "AMST": "Амазон, љетње вријеме", "AMST Armenia": "Јерменија, љетње вријеме", "AMT": "Амазон, стандардно вријеме", "AMT Armenia": "Јерменија, стандардно вријеме", "ANAST": "Анадир летње рачунање времена", "ANAT": "Анадир стандардно време", "ARST": "Аргентина, љетње вријеме", "ART": "Аргентина, стандардно вријеме", "AST": "Атлантско стандардно вријеме", "AST Arabia": "Арабијско стандардно вријеме", "AWDT": "Аустралијско западно љетње вријеме", "AWST": "Аустралијско западно стандардно вријеме", "AZST": "Азербејџан, љетње вријеме", "AZT": "Азербејџан, стандардно вријеме", "BDT Bangladesh": "Бангладеш, љетње вријеме", "BNT": "Брунеј Дарусалум вријеме", "BOT": "Боливија вријеме", "BRST": "Бразилија, љетње вријеме", "BRT": "Бразилија, стандардно вријеме", "BST Bangladesh": "Бангладеш, стандардно вријеме", "BT": "Бутан вријеме", "CAST": "CAST", "CAT": "Централно-афричко вријеме", "CCT": "Кокосова (Килинг) острва вријеме", "CDT": "Сјеверноамеричко централно љетње вријеме", "CHADT": "Чатам, љетње вријеме", "CHAST": "Чатам, стандардно вријеме", "CHUT": "Чук вријеме", "CKT": "Кукова Острва, стандардно вријеме", "CKT DST": "Кукова Острва, полуљетње вријеме", "CLST": "Чиле, љетње вријеме", "CLT": "Чиле, стандардно вријеме", "COST": "Колумбија, љетње вријеме", "COT": "Колумбија, стандардно вријеме", "CST": "Сјеверноамеричко централно стандардно вријеме", "CST China": "Кинеско стандардно вријеме", "CST China DST": "Кина, љетње вријеме", "CVST": "Зеленортска Острва, љетње вријеме", "CVT": "Зеленортска Острва, стандардно вријеме", "CXT": "Божићно острво вријеме", "ChST": "Чаморо вријеме", "ChST NMI": "Северна Маријанска Острва време", "CuDT": "Куба, љетње вријеме", "CuST": "Куба, стандардно вријеме", "DAVT": "Дејвис вријеме", "DDUT": "Димон д’Ирвил вријеме", "EASST": "Ускршња острва, љетње вријеме", "EAST": "Ускршња острва, стандардно вријеме", "EAT": "Источно-афричко вријеме", "ECT": "Еквадор вријеме", "EDT": "Сјеверноамеричко источно љетње вријеме", "EGDT": "Источни Гренланд, љетње вријеме", "EGST": "Источни Гренланд, стандардно вријеме", "EST": "Сјеверноамеричко источно стандардно вријеме", "FEET": "Време даљег истока Европе", "FJT": "Фиџи, стандардно вријеме", "FJT Summer": "Фиџи, љетње вријеме", "FKST": "Фолкландска Острва, љетње вријеме", "FKT": "Фолкландска Острва, стандардно вријеме", "FNST": "Фернандо де Нороња, љетње вријеме", "FNT": "Фернандо де Нороња, стандардно вријеме", "GALT": "Галапагос вријеме", "GAMT": "Гамбије вријеме", "GEST": "Грузија, љетње вријеме", "GET": "Грузија, стандардно вријеме", "GFT": "Француска Гвајана вријеме", "GIT": "Гилбертова острва вријеме", "GMT": "Средње вријеме по Гриничу", "GNSST": "GNSST", "GNST": "GNST", "GST": "Заливско вријеме", "GST Guam": "Гуам стандардно време", "GYT": "Гвајана вријеме", "HADT": "Хавајско-алеутско стандардно вријеме", "HAST": "Хавајско-алеутско стандардно вријеме", "HKST": "Хонг Конг, љетње вријеме", "HKT": "Хонг Конг, стандардно вријеме", "HOVST": "Ховд, љетње вријеме", "HOVT": "Ховд, стандардно вријеме", "ICT": "Индокина вријеме", "IDT": "Израелско љетње вријеме", "IOT": "Индијско океанско вријеме", "IRKST": "Иркуцк, љетње вријеме", "IRKT": "Иркуцк, стандардно вријеме", "IRST": "Иран, стандардно вријеме", "IRST DST": "Иран, љетње вријеме", "IST": "Индијско стандардно вријеме", "IST Israel": "Израелско стандардно вријеме", "JDT": "Јапанско љетње вријеме", "JST": "Јапанско стандардно вријеме", "KOST": "Кошре вријеме", "KRAST": "Краснојарск, љетње вријеме", "KRAT": "Краснојарск, стандардно вријеме", "KST": "Корејско стандардно вријеме", "KST DST": "Корејско љетње вријеме", "LHDT": "Лорд Хов, љетње вријеме", "LHST": "Лорд Хов, стандардно вријеме", "LINT": "Линијска острва вријеме", "MAGST": "Магадан, љетње вријеме", "MAGT": "Магадан, стандардно вријеме", "MART": "Маркиз вријеме", "MAWT": "Мосон вријеме", "MDT": "Макао летње рачунање времена", "MESZ": "Средњоевропско љетње вријеме", "MEZ": "Средњоевропско стандардно вријеме", "MHT": "Маршалска Острва вријеме", "MMT": "Мјанмар вријеме", "MSD": "Москва, љетње вријеме", "MST": "Макао стандардно време", "MUST": "Маурицијус, љетње вријеме", "MUT": "Маурицијус, стандардно вријеме", "MVT": "Малдиви вријеме", "MYT": "Малезија вријеме", "NCT": "Нова Каледонија, стандардно вријеме", "NDT": "Њуфаундленд, љетње вријеме", "NDT New Caledonia": "Нова Каледонија, љетње вријеме", "NFDT": "острво Норфолк, љетње вријеме", "NFT": "острво Норфолк, стандардно вријеме", "NOVST": "Новосибирск, љетње вријеме", "NOVT": "Новосибирск, стандардно вријеме", "NPT": "Непал вријеме", "NRT": "Науру вријеме", "NST": "Њуфаундленд, стандардно вријеме", "NUT": "Нијуе вријеме", "NZDT": "Нови Зеланд, љетње вријеме", "NZST": "Нови Зеланд, стандардно вријеме", "OESZ": "Источноевропско љетње вријеме", "OEZ": "Источноевропско стандардно вријеме", "OMSST": "Омск, љетње вријеме", "OMST": "Омск, стандардно вријеме", "PDT": "Сјеверноамеричко пацифичко летње вријеме", "PDTM": "Мексички Пацифик, љетње вријеме", "PETDT": "Петропавловско-камчатско летње рачунање времена", "PETST": "Петропавловско-камчатско стандардно време", "PGT": "Папуа Нова Гвинеја вријеме", "PHOT": "Феникс острва вријеме", "PKT": "Пакистан, стандардно вријеме", "PKT DST": "Пакистан, љетње вријеме", "PMDT": "Сен Пјер и Микелон, љетње вријеме", "PMST": "Сен Пјер и Микелон, стандардно вријеме", "PONT": "Понпеј вријеме", "PST": "Сјеверноамеричко пацифичко стандардно вријеме", "PST Philippine": "Филипини, стандардно вријеме", "PST Philippine DST": "Филипини, љетње вријеме", "PST Pitcairn": "Питкерн вријеме", "PSTM": "Мексички Пацифик, стандардно вријеме", "PWT": "Палау вријеме", "PYST": "Парагвај, љетње вријеме", "PYT": "Парагвај, стандардно вријеме", "PYT Korea": "Пјонгјаншко вријеме", "RET": "Реунион вријеме", "ROTT": "Ротера вријеме", "SAKST": "Сахалин, љетње вријеме", "SAKT": "Сахалин, стандардно вријеме", "SAMST": "Самара летње рачунање времена", "SAMT": "Самара стандардно време", "SAST": "Јужно-афричко вријеме", "SBT": "Соломонска Острва вријеме", "SCT": "Сејшели вријеме", "SGT": "Сингапур, стандардно вријеме", "SLST": "Шри Ланка време", "SRT": "Суринам вријеме", "SST Samoa": "Самоа, стандардно вријеме", "SST Samoa Apia": "Апија, стандардно вријеме", "SST Samoa Apia DST": "Апија, љетње вријеме", "SST Samoa DST": "Самоа, љетње вријеме", "SYOT": "Шова вријеме", "TAAF": "Француско јужно и антарктичко вријеме", "TAHT": "Тахити вријеме", "TJT": "Таџикистан вријеме", "TKT": "Токелау вријеме", "TLT": "Источни Тимор вријеме", "TMST": "Туркменистан, љетње вријеме", "TMT": "Туркменистан, стандардно вријеме", "TOST": "Тонга, љетње вријеме", "TOT": "Тонга, стандардно вријеме", "TVT": "Тувалу вријеме", "TWT": "Тајпеј, стандардно вријеме", "TWT DST": "Тајпеј, љетње вријеме", "ULAST": "Улан Батор, љетње вријееме", "ULAT": "Улан Батор, стандардно вријеме", "UYST": "Уругвај, љетње вријеме", "UYT": "Уругвај, стандардно вријеме", "UZT": "Узбекистан, стандардно вријеме", "UZT DST": "Узбекистан, љетње вријеме", "VET": "Венецуела вријеме", "VLAST": "Владивосток, љетње вријеме", "VLAT": "Владивосток, стандардно вријеме", "VOLST": "Волгоград, љетње вријеме", "VOLT": "Волгоград, стандардно вријеме", "VOST": "Восток вријеме", "VUT": "Вануату, стандардно вријеме", "VUT DST": "Вануату, љетње вријеме", "WAKT": "острво Вејк вријеме", "WARST": "Западна Аргентина, љетње вријеме", "WART": "Западна Аргентина, стандардно вријеме", "WAST": "Западно-афричко вријеме", "WAT": "Западно-афричко вријеме", "WESZ": "Западноевропско љетње вријеме", "WEZ": "Западноевропско стандардно вријеме", "WFT": "острва Валис и Футуна вријеме", "WGST": "Западни Гренланд, љетње вријеме", "WGT": "Западни Гренланд, стандардно вријеме", "WIB": "Западно-индонезијско вријеме", "WIT": "Источно-индонезијско вријеме", "WITA": "Централно-индонезијско вријеме", "YAKST": "Јакутск, љетње вријеме", "YAKT": "Јакутск, стандардно вријеме", "YEKST": "Јекатеринбург, љетње вријеме", "YEKT": "Јекатеринбург, стандардно вријеме", "YST": "Јукон", "МСК": "Москва, стандардно вријеме", "اقتاۋ": "Акватау стандардно време", "اقتاۋ قالاسى": "Акватау летње рачунање времена", "اقتوبە": "Акутобе стандардно време", "اقتوبە قالاسى": "Акутобе летње рачунање времена", "الماتى": "Алмати стандардно време", "الماتى قالاسى": "Алмати летње рачунање времена", "باتىس قازاق ەلى": "Западно-казахстанско вријеме", "شىعىش قازاق ەلى": "Источно-казахстанско вријеме", "قازاق ەلى": "Казахстанско вријеме", "قىرعىزستان": "Киргистан вријеме", "قىزىلوردا": "Кизилорда стандардно време", "قىزىلوردا قالاسى": "Кизилорда летње рачунање времена", "∅∅∅": "Перу, љетње вријеме"},
	}
}

// Locale returns the current translators string locale
func (sr *sr_Cyrl_BA) Locale() string {
	return sr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) PluralsCardinal() []locales.PluralRule {
	return sr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) PluralsOrdinal() []locales.PluralRule {
	return sr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) PluralsRange() []locales.PluralRule {
	return sr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
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
func (sr *sr_Cyrl_BA) MonthAbbreviated(month time.Month) string {
	return sr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sr *sr_Cyrl_BA) MonthsAbbreviated() []string {
	return sr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sr *sr_Cyrl_BA) MonthNarrow(month time.Month) string {
	return sr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sr *sr_Cyrl_BA) MonthsNarrow() []string {
	return sr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sr *sr_Cyrl_BA) MonthWide(month time.Month) string {
	return sr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sr *sr_Cyrl_BA) MonthsWide() []string {
	return sr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sr *sr_Cyrl_BA) WeekdayAbbreviated(weekday time.Weekday) string {
	return sr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sr *sr_Cyrl_BA) WeekdaysAbbreviated() []string {
	return sr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sr *sr_Cyrl_BA) WeekdayNarrow(weekday time.Weekday) string {
	return sr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sr *sr_Cyrl_BA) WeekdaysNarrow() []string {
	return sr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sr *sr_Cyrl_BA) WeekdayShort(weekday time.Weekday) string {
	return sr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sr *sr_Cyrl_BA) WeekdaysShort() []string {
	return sr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sr *sr_Cyrl_BA) WeekdayWide(weekday time.Weekday) string {
	return sr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sr *sr_Cyrl_BA) WeekdaysWide() []string {
	return sr.daysWide
}

// Decimal returns the decimal point of number
func (sr *sr_Cyrl_BA) Decimal() string {
	return sr.decimal
}

// Group returns the group of number
func (sr *sr_Cyrl_BA) Group() string {
	return sr.group
}

// Group returns the minus sign of number
func (sr *sr_Cyrl_BA) Minus() string {
	return sr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sr_Cyrl_BA' and handles both Whole and Real numbers based on 'v'
func (sr *sr_Cyrl_BA) FmtNumber(num float64, v uint64) string {
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

// FmtPercent returns 'num' with digits/precision of 'v' for 'sr_Cyrl_BA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sr *sr_Cyrl_BA) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtCurrency(num float64, v uint64, currency currency.Type) string {
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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Cyrl_BA'
// in accounting notation.
func (sr *sr_Cyrl_BA) FmtAccounting(num float64, v uint64, currency currency.Type) string {
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

// FmtDateShort returns the short date representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtDateMedium(t time.Time) string {
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

// FmtDateLong returns the long date representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtDateLong(t time.Time) string {
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

// FmtDateFull returns the full date representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtDateFull(t time.Time) string {
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

// FmtTimeShort returns the short time representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtTimeShort(t time.Time) string {
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

// FmtTimeMedium returns the medium time representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtTimeMedium(t time.Time) string {
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

// FmtTimeLong returns the long time representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtTimeLong(t time.Time) string {
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

// FmtTimeFull returns the full time representation of 't' for 'sr_Cyrl_BA'
func (sr *sr_Cyrl_BA) FmtTimeFull(t time.Time) string {
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
