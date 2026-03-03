package kk_Cyrl_KZ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kk_Cyrl_KZ struct {
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

// New returns a new instance of translator for the 'kk_Cyrl_KZ' locale
func New() locales.Translator {
	return &kk_Cyrl_KZ{
		locale:                 "kk_Cyrl_KZ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "₸", "LAK", "LBP", "LKR", "LRD", "ЛСЛ", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "қаң.", "ақп.", "нау.", "сәу.", "мам.", "мау.", "шіл.", "там.", "қыр.", "қаз.", "қар.", "жел."},
		monthsNarrow:           []string{"", "Қ", "А", "Н", "С", "М", "М", "Ш", "Т", "Қ", "Қ", "Қ", "Ж"},
		monthsWide:             []string{"", "қаңтар", "ақпан", "наурыз", "сәуір", "мамыр", "маусым", "шілде", "тамыз", "қыркүйек", "қазан", "қараша", "желтоқсан"},
		daysAbbreviated:        []string{"жс", "дс", "сс", "ср", "бс", "жм", "сб"},
		daysNarrow:             []string{"Ж", "Д", "С", "С", "Б", "Ж", "С"},
		daysWide:               []string{"жексенбі", "дүйсенбі", "сейсенбі", "сәрсенбі", "бейсенбі", "жұма", "сенбі"},
		timezones:              map[string]string{"ACDT": "Аустралия жазғы орталық уақыты", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Аустралия жазғы орталық-батыс уақыты", "ACWST": "Аустралия стандартты орталық-батыс уақыты", "ADT": "Атлантика жазғы уақыты", "ADT Arabia": "Сауд Арабиясы жазғы уақыты", "AEDT": "Аустралия жазғы шығыс уақыты", "AEST": "Аустралия стандартты шығыс уақыты", "AFT": "Ауғанстан уақыты", "AKDT": "Аляска жазғы уақыты", "AKST": "Аляска стандартты уақыты", "AMST": "Амазонка жазғы уақыты", "AMST Armenia": "Армения жазғы уақыты", "AMT": "Амазонка стандартты уақыты", "AMT Armenia": "Армения стандартты уақыты", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентина жазғы уақыты", "ART": "Аргентина стандартты уақыты", "AST": "Атлантика стандартты уақыты", "AST Arabia": "Сауд Арабиясы стандартты уақыты", "AWDT": "Аустралия жазғы батыс уақыты", "AWST": "Аустралия стандартты батыс уақыты", "AZST": "Әзірбайжан жазғы уақыты", "AZT": "Әзірбайжан стандартты уақыты", "BDT Bangladesh": "Бангладеш жазғы уақыты", "BNT": "Бруней-Даруссалам уақыты", "BOT": "Боливия уақыты", "BRST": "Бразилия жазғы уақыты", "BRT": "Бразилия стандартты уақыты", "BST Bangladesh": "Бангладеш стандартты уақыты", "BT": "Бутан уақыты", "CAST": "CAST", "CAT": "Орталық Африка уақыты", "CCT": "Кокос аралдарының уақыты", "CDT": "Солтүстік Америка жазғы орталық уақыты", "CHADT": "Чатем жазғы уақыты", "CHAST": "Чатем стандартты уақыты", "CHUT": "Трук уақыты", "CKT": "Кук аралдарының стандартты уақыты", "CKT DST": "Кук аралдарының жартылай жазғы уақыты", "CLST": "Чили жазғы уақыты", "CLT": "Чили стандартты уақыты", "COST": "Колумбия жазғы уақыты", "COT": "Колумбия стандартты уақыты", "CST": "Солтүстік Америка стандартты орталық уақыты", "CST China": "Қытай стандартты уақыты", "CST China DST": "Қытай жазғы уақыты", "CVST": "Кабо-Верде жазғы уақыты", "CVT": "Кабо-Верде стандартты уақыты", "CXT": "Рождество аралының уақыты", "ChST": "Чаморро стандартты уақыты", "ChST NMI": "ChST NMI", "CuDT": "Куба жазғы уақыты", "CuST": "Куба стандартты уақыты", "DAVT": "Дейвис уақыты", "DDUT": "Дюмон-д’Юрвиль уақыты", "EASST": "Пасха аралы жазғы уақыты", "EAST": "Пасха аралы стандартты уақыты", "EAT": "Шығыс Африка уақыты", "ECT": "Эквадор уақыты", "EDT": "Солтүстік Америка жазғы шығыс уақыты", "EGDT": "Шығыс Гренландия жазғы уақыты", "EGST": "Шығыс Гренландия стандартты уақыты", "EST": "Солтүстік Америка стандартты шығыс уақыты", "FEET": "Қиыр Шығыс Еуропа уақыты", "FJT": "Фиджи стандартты уақыты", "FJT Summer": "Фиджи жазғы уақыты", "FKST": "Фолкленд аралдары жазғы уақыты", "FKT": "Фолкленд аралдары стандартты уақыты", "FNST": "Фернанду-ди-Норонья жазғы уақыты", "FNT": "Фернанду-ди-Норонья стандартты уақыты", "GALT": "Галапагос уақыты", "GAMT": "Гамбье уақыты", "GEST": "Грузия жазғы уақыты", "GET": "Грузия стандартты уақыты", "GFT": "Француз Гвианасы уақыты", "GIT": "Гилберт аралдарының уақыты", "GMT": "Гринвич уақыты", "GNSST": "GNSST", "GNST": "GNST", "GST": "Оңтүстік Георгия уақыты", "GST Guam": "GST Guam", "GYT": "Гайана уақыты", "HADT": "Гавай және Алеут аралдары стандартты уақыты", "HAST": "Гавай және Алеут аралдары стандартты уақыты", "HKST": "Гонконг жазғы уақыты", "HKT": "Гонконг стандартты уақыты", "HOVST": "Ховд жазғы уақыты", "HOVT": "Ховд стандартты уақыты", "ICT": "Үндіқытай уақыты", "IDT": "Израиль жазғы уақыты", "IOT": "Үнді мұхиты уақыты", "IRKST": "Иркутск жазғы уақыты", "IRKT": "Иркутск стандартты уақыты", "IRST": "Иран стандартты уақыты", "IRST DST": "Иран жазғы уақыты", "IST": "Үндістан стандартты уақыты", "IST Israel": "Израиль стандартты уақыты", "JDT": "Жапония жазғы уақыты", "JST": "Жапония стандартты уақыты", "KOST": "Кусаие уақыты", "KRAST": "Красноярск жазғы уақыты", "KRAT": "Красноярск стандартты уақыты", "KST": "Корея стандартты уақыты", "KST DST": "Корея жазғы уақыты", "LHDT": "Лорд-Хау жазғы уақыты", "LHST": "Лорд-Хау стандартты уақыты", "LINT": "Лайн аралдары уақыты", "MAGST": "Магадан жазғы уақыты", "MAGT": "Магадан стандартты уақыты", "MART": "Маркиз аралдары уақыты", "MAWT": "Моусон уақыты", "MDT": "MDT", "MESZ": "Орталық Еуропа жазғы уақыты", "MEZ": "Орталық Еуропа стандартты уақыты", "MHT": "Маршалл аралдары уақыты", "MMT": "Мьянма уақыты", "MSD": "Мәскеу жазғы уақыты", "MST": "MST", "MUST": "Маврикий жазғы уақыты", "MUT": "Маврикий стандартты уақыты", "MVT": "Мальдив аралдары уақыты", "MYT": "Малайзия уақыты", "NCT": "Жаңа Каледония стандартты уақыты", "NDT": "Ньюфаундленд жазғы уақыты", "NDT New Caledonia": "Жаңа Каледония жазғы уақыты", "NFDT": "Норфолк аралы жазғы уақыты", "NFT": "Норфолк аралы стандартты уақыты", "NOVST": "Новосібір жазғы уақыты", "NOVT": "Новосібір стандартты уақыты", "NPT": "Непал уақыты", "NRT": "Науру уақыты", "NST": "Ньюфаундленд стандартты уақыты", "NUT": "Ниуэ уақыты", "NZDT": "Жаңа Зеландия жазғы уақыты", "NZST": "Жаңа Зеландия стандартты уақыты", "OESZ": "Шығыс Еуропа жазғы уақыты", "OEZ": "Шығыс Еуропа стандартты уақыты", "OMSST": "Омбы жазғы уақыты", "OMST": "Омбы стандартты уақыты", "PDT": "Солтүстік Америка жазғы Тынық мұхиты уақыты", "PDTM": "Мексика жазғы Тынық мұхит уақыты", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Папуа – Жаңа Гвинея уақыты", "PHOT": "Феникс аралдары уақыты", "PKT": "Пәкістан стандартты уақыты", "PKT DST": "Пәкістан жазғы уақыты", "PMDT": "Сен-Пьер және Микелон жазғы уақыты", "PMST": "Сен-Пьер және Микелон стандартты уақыты", "PONT": "Понпеи уақыты", "PST": "Солтүстік Америка стандартты Тынық мұхиты уақыты", "PST Philippine": "Филиппин аралдары стандартты уақыты", "PST Philippine DST": "Филиппин аралдары жазғы уақыты", "PST Pitcairn": "Питкэрн уақыты", "PSTM": "Мексика стандартты Тынық мұхит уақыты", "PWT": "Палау уақыты", "PYST": "Парагвай жазғы уақыты", "PYT": "Парагвай стандартты уақыты", "PYT Korea": "Пхеньян уақыты", "RET": "Реюньон уақыты", "ROTT": "Ротера уақыты", "SAKST": "Сахалин жазғы уақыты", "SAKT": "Сахалин стандартты уақыты", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Оңтүстік Африка стандартты уақыты", "SBT": "Соломон аралдары уақыты", "SCT": "Сейшель аралдары уақыты", "SGT": "Сингапур стандартты уақыты", "SLST": "SLST", "SRT": "Суринам уақыты", "SST Samoa": "Самоа стандартты уақыты", "SST Samoa Apia": "Апиа стандартты уақыты", "SST Samoa Apia DST": "Апиа жазғы уақыты", "SST Samoa DST": "Самоа жазғы уақыты", "SYOT": "Сёва уақыты", "TAAF": "Францияның оңтүстік аймағы және Антарктика уақыты", "TAHT": "Таити уақыты", "TJT": "Тәжікстан уақыты", "TKT": "Токелау уақыты", "TLT": "Шығыс Тимор уақыты", "TMST": "Түрікменстан жазғы уақыты", "TMT": "Түрікменстан стандартты уақыты", "TOST": "Тонга жазғы уақыты", "TOT": "Тонга стандартты уақыты", "TVT": "Тувалу уақыты", "TWT": "Тайбэй стандартты уақыты", "TWT DST": "Тайбэй жазғы уақыты", "ULAST": "Ұланбатыр жазғы уақыты", "ULAT": "Ұланбатыр стандартты уақыты", "UYST": "Уругвай жазғы уақыты", "UYT": "Уругвай стандартты уақыты", "UZT": "Өзбекстан стандартты уақыты", "UZT DST": "Өзбекстан жазғы уақыты", "VET": "Венесуэла уақыты", "VLAST": "Владивосток жазғы уақыты", "VLAT": "Владивосток стандартты уақыты", "VOLST": "Волгоград жазғы уақыты", "VOLT": "Волгоград стандартты уақыты", "VOST": "Восток уақыты", "VUT": "Вануату стандартты уақыты", "VUT DST": "Вануату жазғы уақыты", "WAKT": "Уэйк аралы уақыты", "WARST": "Батыс Аргентина жазғы уақыты", "WART": "Батыс Аргентина стандартты уақыты", "WAST": "Батыс Африка уақыты", "WAT": "Батыс Африка уақыты", "WESZ": "Батыс Еуропа жазғы уақыты", "WEZ": "Батыс Еуропа стандартты уақыты", "WFT": "Уоллис және Футуна уақыты", "WGST": "Батыс Гренландия жазғы уақыты", "WGT": "Батыс Гренландия стандартты уақыты", "WIB": "Батыс Индонезия уақыты", "WIT": "Шығыс Индонезия уақыты", "WITA": "Орталық Индонезия уақыты", "YAKST": "Якутск жазғы уақыты", "YAKT": "Якутск стандартты уақыты", "YEKST": "Екатеринбург жазғы уақыты", "YEKT": "Екатеринбург стандартты уақыты", "YST": "Юкон уақыты", "МСК": "Мәскеу стандартты уақыты", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Батыс Қазақстан уақыты", "شىعىش قازاق ەلى": "Шығыс Қазақстан уақыты", "قازاق ەلى": "Қазақстан уақыты", "قىرعىزستان": "Қырғызстан уақыты", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Азор аралдары жазғы уақыты"},
	}
}

// Locale returns the current translators string locale
func (kk *kk_Cyrl_KZ) Locale() string {
	return kk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) PluralsCardinal() []locales.PluralRule {
	return kk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) PluralsOrdinal() []locales.PluralRule {
	return kk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) PluralsRange() []locales.PluralRule {
	return kk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)

	if (nMod10 == 6) || (nMod10 == 9) || (nMod10 == 0 && n != 0) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := kk.CardinalPluralRule(num1, v1)
	end := kk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kk *kk_Cyrl_KZ) MonthAbbreviated(month time.Month) string {
	if len(kk.monthsAbbreviated) == 0 {
		return ""
	}
	return kk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kk *kk_Cyrl_KZ) MonthsAbbreviated() []string {
	return kk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kk *kk_Cyrl_KZ) MonthNarrow(month time.Month) string {
	if len(kk.monthsNarrow) == 0 {
		return ""
	}
	return kk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kk *kk_Cyrl_KZ) MonthsNarrow() []string {
	return kk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kk *kk_Cyrl_KZ) MonthWide(month time.Month) string {
	if len(kk.monthsWide) == 0 {
		return ""
	}
	return kk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kk *kk_Cyrl_KZ) MonthsWide() []string {
	return kk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kk *kk_Cyrl_KZ) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(kk.daysAbbreviated) == 0 {
		return ""
	}
	return kk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kk *kk_Cyrl_KZ) WeekdaysAbbreviated() []string {
	return kk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kk *kk_Cyrl_KZ) WeekdayNarrow(weekday time.Weekday) string {
	if len(kk.daysNarrow) == 0 {
		return ""
	}
	return kk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kk *kk_Cyrl_KZ) WeekdaysNarrow() []string {
	return kk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kk *kk_Cyrl_KZ) WeekdayShort(weekday time.Weekday) string {
	if len(kk.daysShort) == 0 {
		return ""
	}
	return kk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kk *kk_Cyrl_KZ) WeekdaysShort() []string {
	return kk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kk *kk_Cyrl_KZ) WeekdayWide(weekday time.Weekday) string {
	if len(kk.daysWide) == 0 {
		return ""
	}
	return kk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kk *kk_Cyrl_KZ) WeekdaysWide() []string {
	return kk.daysWide
}

// Decimal returns the decimal point of number
func (kk *kk_Cyrl_KZ) Decimal() string {
	return kk.decimal
}

// Group returns the group of number
func (kk *kk_Cyrl_KZ) Group() string {
	return kk.group
}

// Group returns the minus sign of number
func (kk *kk_Cyrl_KZ) Minus() string {
	return kk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kk_Cyrl_KZ' and handles both Whole and Real numbers based on 'v'
func (kk *kk_Cyrl_KZ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kk.group) - 1; j >= 0; j-- {
					b = append(b, kk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kk_Cyrl_KZ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kk *kk_Cyrl_KZ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kk.group) - 1; j >= 0; j-- {
					b = append(b, kk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, kk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kk_Cyrl_KZ'
// in accounting notation.
func (kk *kk_Cyrl_KZ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kk.group) - 1; j >= 0; j-- {
					b = append(b, kk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, kk.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, kk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, kk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtDateShort(t time.Time) string {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kk.monthsAbbreviated[t.Month()]...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kk.monthsWide[t.Month()]...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, kk.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kk_Cyrl_KZ'
func (kk *kk_Cyrl_KZ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
