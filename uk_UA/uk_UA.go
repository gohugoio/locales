package uk_UA

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type uk_UA struct {
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

// New returns a new instance of translator for the 'uk_UA' locale
func New() locales.Translator {
	return &uk_UA{
		locale:                 "uk_UA",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{4, 6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "січ.", "лют.", "бер.", "квіт.", "трав.", "черв.", "лип.", "серп.", "вер.", "жовт.", "лист.", "груд."},
		monthsNarrow:           []string{"", "с", "л", "б", "к", "т", "ч", "л", "с", "в", "ж", "л", "г"},
		monthsWide:             []string{"", "січня", "лютого", "березня", "квітня", "травня", "червня", "липня", "серпня", "вересня", "жовтня", "листопада", "грудня"},
		daysAbbreviated:        []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysNarrow:             []string{"Н", "П", "В", "С", "Ч", "П", "С"},
		daysWide:               []string{"неділя", "понеділок", "вівторок", "середа", "четвер", "пʼятниця", "субота"},
		periodsAbbreviated:     []string{"дп", "пп"},
		erasAbbreviated:        []string{"до н. е.", "н. е."},
		erasNarrow:             []string{"до н.е.", "н.е."},
		erasWide:               []string{"до нашої ери", "нашої ери"},
		timezones:              map[string]string{"ACDT": "за літнім центральноавстралійським часом", "ACST": "за стандартним центральноавстралійським часом", "ACT": "час: Акрі, стандартний", "ACWDT": "за літнім центральнозахідним австралійським часом", "ACWST": "за стандартним центральнозахідним австралійським часом", "ADT": "за атлантичним літнім часом", "ADT Arabia": "за арабським літнім часом", "AEDT": "за літнім східноавстралійським часом", "AEST": "за стандартним східноавстралійським часом", "AFT": "за часом в Афганістані", "AKDT": "за літнім часом на Алясці", "AKST": "за стандартним часом на Алясці", "AMST": "за літнім часом на Амазонці", "AMST Armenia": "за вірменським літнім часом", "AMT": "за стандартним часом на Амазонці", "AMT Armenia": "за вірменським стандартним часом", "ANAST": "час: Анадир, літній", "ANAT": "час: Анадир, стандартний", "ARST": "за літнім аргентинським часом", "ART": "за стандартним аргентинським часом", "AST": "за атлантичним стандартним часом", "AST Arabia": "за арабським стандартним часом", "AWDT": "за літнім західноавстралійським часом", "AWST": "за стандартним західноавстралійським часом", "AZST": "за літнім азербайджанським часом", "AZT": "за стандартним азербайджанським часом", "BDT Bangladesh": "за літнім часом у Бангладеш", "BNT": "за часом у Брунеї", "BOT": "за болівійським часом", "BRST": "за літнім бразильським часом", "BRT": "за стандартним бразильським часом", "BST Bangladesh": "за стандартним часом у Бангладеш", "BT": "за часом у Бутані", "CAST": "CAST", "CAT": "за центральноафриканським часом", "CCT": "за часом на Кокосових островах", "CDT": "за північноамериканським центральним літнім часом", "CHADT": "за літнім часом на архіпелазі Чатем", "CHAST": "за стандартним часом на архіпелазі Чатем", "CHUT": "за часом на островах Чуук", "CKT": "за стандартним часом на Островах Кука", "CKT DST": "за літнім часом на Островах Кука", "CLST": "за літнім чилійським часом", "CLT": "за стандартним чилійським часом", "COST": "за літнім колумбійським часом", "COT": "за стандартним колумбійським часом", "CST": "за північноамериканським центральним стандартним часом", "CST China": "за китайським стандартним часом", "CST China DST": "за китайським літнім часом", "CVST": "за літнім часом на островах Кабо-Верде", "CVT": "за стандартним часом на островах Кабо-Верде", "CXT": "за часом на острові Різдва", "ChST": "за часом на Північних Маріанських островах", "ChST NMI": "ChST NMI", "CuDT": "за літнім часом на Кубі", "CuST": "за стандартним часом на Кубі", "DAVT": "за часом на станції Девіс", "DDUT": "за часом на станції Дюмон дʼЮрвіль", "EASST": "за літнім часом на острові Пасхи", "EAST": "за стандартним часом на острові Пасхи", "EAT": "за східноафриканським часом", "ECT": "за часом в Еквадорі", "EDT": "за північноамериканським східним літнім часом", "EGDT": "за літнім східним часом у Ґренландії", "EGST": "за стандартним східним часом у Ґренландії", "EST": "за північноамериканським східним стандартним часом", "FEET": "за далекосхідним європейським часом", "FJT": "за стандартним часом у Фіджі", "FJT Summer": "за літнім часом у Фіджі", "FKST": "за літнім часом на Фолклендських Островах", "FKT": "за стандартним часом на Фолклендських Островах", "FNST": "за літнім часом на архіпелазі Фернанду-ді-Норонья", "FNT": "за стандартним часом на архіпелазі Фернанду-ді-Норонья", "GALT": "за часом Ґалапаґосу", "GAMT": "за часом на острові Ґамбʼє", "GEST": "за літнім грузинським часом", "GET": "за стандартним грузинським часом", "GFT": "за часом Французької Ґвіани", "GIT": "за часом на островах Гілберта", "GMT": "за Гринвічем", "GNSST": "GNSST", "GNST": "GNST", "GST": "за часом Перської затоки", "GST Guam": "за часом на острові Гуам", "GYT": "за часом у Ґаяні", "HADT": "за стандартним гавайсько-алеутським часом", "HAST": "за стандартним гавайсько-алеутським часом", "HKST": "за літнім часом у Гонконзі", "HKT": "за стандартним часом у Гонконзі", "HOVST": "за літнім часом у Ховді", "HOVT": "за стандартним часом у Ховді", "ICT": "за часом в Індокитаї", "IDT": "за ізраїльським літнім часом", "IOT": "за часом в Індійському Океані", "IRKST": "за іркутським літнім часом", "IRKT": "за іркутським стандартним часом", "IRST": "за іранським стандартним часом", "IRST DST": "за іранським літнім часом", "IST": "за індійським стандартним часом", "IST Israel": "за ізраїльським стандартним часом", "JDT": "за японським літнім часом", "JST": "за японським стандартним часом", "KOST": "за часом на острові Косрае", "KRAST": "за красноярським літнім часом", "KRAT": "за красноярським стандартним часом", "KST": "за корейським стандартним часом", "KST DST": "за корейським літнім часом", "LHDT": "за літнім часом на острові Лорд-Хау", "LHST": "за стандартним часом на острові Лорд-Хау", "LINT": "за часом на острові Лайн", "MAGST": "за магаданським літнім часом", "MAGT": "за магаданським стандартним часом", "MART": "за часом на Маркізьких островах", "MAWT": "за часом на станції Моусон", "MDT": "MDT", "MESZ": "за центральноєвропейським літнім часом", "MEZ": "за центральноєвропейським стандартним часом", "MHT": "за часом на Маршаллових Островах", "MMT": "за часом у Мʼянмі", "MSD": "за московським літнім часом", "MST": "MST", "MUST": "за літнім часом на острові Маврикій", "MUT": "за стандартним часом на острові Маврикій", "MVT": "за часом на Мальдівах", "MYT": "за часом у Малайзії", "NCT": "за стандартним часом на островах Нової Каледонії", "NDT": "за літнім часом у Ньюфаундленд", "NDT New Caledonia": "за літнім часом на островах Нової Каледонії", "NFDT": "за літнім часом на острові Норфолк", "NFT": "за стандартним часом на острові Норфолк", "NOVST": "за новосибірським літнім часом", "NOVT": "за новосибірським стандартним часом", "NPT": "за часом у Непалі", "NRT": "за часом на острові Науру", "NST": "за стандартним часом на острові Ньюфаундленд", "NUT": "за часом на острові Ніуе", "NZDT": "за літнім часом у Новій Зеландії", "NZST": "за стандартним часом у Новій Зеландії", "OESZ": "за східноєвропейським літнім часом", "OEZ": "за східноєвропейським стандартним часом", "OMSST": "за омським літнім часом", "OMST": "за омським стандартним часом", "PDT": "за північноамериканським тихоокеанським літнім часом", "PDTM": "за літнім тихоокеанським часом у Мексиці", "PETDT": "за камчатським літнім часом", "PETST": "за камчатським стандартним часом", "PGT": "за часом на островах Папуа-Нова Ґвінея", "PHOT": "за часом на островах Фенікс", "PKT": "за стандартним часом у Пакистані", "PKT DST": "за літнім часом у Пакистані", "PMDT": "за літнім часом на островах Сен-П’єр і Мікелон", "PMST": "за стандартним часом на островах Сен-П’єр і Мікелон", "PONT": "за часом на острові Понапе", "PST": "за північноамериканським тихоокеанським стандартним часом", "PST Philippine": "за стандартним часом на Філіппінах", "PST Philippine DST": "за літнім часом на Філіппінах", "PST Pitcairn": "за часом на островах Піткерн", "PSTM": "за стандартним тихоокеанським часом у Мексиці", "PWT": "за часом на острові Палау", "PYST": "за літнім параґвайським часом", "PYT": "за стандартним параґвайським часом", "PYT Korea": "за часом у Пхеньяні", "RET": "за часом на острові Реюньйон", "ROTT": "за часом на станції Ротера", "SAKST": "за сахалінським літнім часом", "SAKT": "за сахалінським стандартним часом", "SAMST": "за самарським літнім часом", "SAMT": "за самарським стандартним часом", "SAST": "за південноафриканським часом", "SBT": "за часом на Соломонових Островах", "SCT": "за часом на Сейшельських Островах", "SGT": "за часом у Сінгапурі", "SLST": "час: Ланка", "SRT": "за часом у Суринамі", "SST Samoa": "за стандартним часом на острові Самоа", "SST Samoa Apia": "за стандартним часом в Апіа", "SST Samoa Apia DST": "за літнім часом в Апіа", "SST Samoa DST": "за літнім часом на острові Самоа", "SYOT": "за часом на станції Сева", "TAAF": "за часом на Французьких Південних і Антарктичних територіях", "TAHT": "за часом на острові Таїті", "TJT": "за часом у Таджикистані", "TKT": "за часом на островах Токелау", "TLT": "за часом у Східному Тиморі", "TMST": "за літнім часом у Туркменістані", "TMT": "за стандартним часом у Туркменістані", "TOST": "за літнім часом на островах Тонга", "TOT": "за стандартним часом на островах Тонга", "TVT": "за часом на островах Тувалу", "TWT": "за стандартним часом у Тайбеї", "TWT DST": "за літнім часом у Тайбеї", "ULAST": "за літнім часом в Улан-Баторі", "ULAT": "за стандартним часом в Улан-Баторі", "UYST": "за літнім часом в Уруґваї", "UYT": "за стандартним часом в Уруґваї", "UZT": "за стандартним часом в Узбекистані", "UZT DST": "за літнім часом в Узбекистані", "VET": "за часом у Венесуелі", "VLAST": "за владивостоцьким літнім часом", "VLAT": "за владивостоцьким стандартним часом", "VOLST": "за волгоградським літнім часом", "VOLT": "за волгоградським стандартним часом", "VOST": "за часом на станції Восток", "VUT": "за стандартним часом на островах Вануату", "VUT DST": "за літнім часом на островах Вануату", "WAKT": "за часом на острові Вейк", "WARST": "за літнім за західноаргентинським часом", "WART": "за стандартним західноаргентинським часом", "WAST": "за західноафриканським часом", "WAT": "за західноафриканським часом", "WESZ": "за західноєвропейським літнім часом", "WEZ": "за західноєвропейським стандартним часом", "WFT": "за часом на островах Уолліс і Футуна", "WGST": "за літнім західним часом у Ґренландії", "WGT": "за стандартним західним часом у Ґренландії", "WIB": "за західноіндонезійським часом", "WIT": "за східноіндонезійським часом", "WITA": "за центральноіндонезійським часом", "YAKST": "за якутським літнім часом", "YAKT": "за якутським стандартним часом", "YEKST": "за єкатеринбурзьким літнім часом", "YEKT": "за єкатеринбурзьким стандартним часом", "YST": "за стандартним часом на Юконі", "МСК": "за московським стандартним часом", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "за західним часом у Казахстані", "شىعىش قازاق ەلى": "за східним часом у Казахстані", "قازاق ەلى": "за часом у Казахстані", "قىرعىزستان": "за часом у Киргизстані", "قىزىلوردا": "час: Кизилорда, стандартний", "قىزىلوردا قالاسى": "час: Кизилорда, літній", "∅∅∅": "за літнім часом у Перу"},
	}
}

// Locale returns the current translators string locale
func (uk *uk_UA) Locale() string {
	return uk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uk_UA'
func (uk *uk_UA) PluralsCardinal() []locales.PluralRule {
	return uk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uk_UA'
func (uk *uk_UA) PluralsOrdinal() []locales.PluralRule {
	return uk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uk_UA'
func (uk *uk_UA) PluralsRange() []locales.PluralRule {
	return uk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uk_UA'
func (uk *uk_UA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod100 := i % 100
	iMod10 := i % 10

	if v == 0 && iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14) {
		return locales.PluralRuleFew
	} else if (v == 0 && iMod10 == 0) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 11 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uk_UA'
func (uk *uk_UA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)
	nMod10 := math.Mod(n, 10)

	if nMod10 == 3 && nMod100 != 13 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uk_UA'
func (uk *uk_UA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uk.CardinalPluralRule(num1, v1)
	end := uk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uk *uk_UA) MonthAbbreviated(month time.Month) string {
	return uk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uk *uk_UA) MonthsAbbreviated() []string {
	return uk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uk *uk_UA) MonthNarrow(month time.Month) string {
	return uk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uk *uk_UA) MonthsNarrow() []string {
	return uk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uk *uk_UA) MonthWide(month time.Month) string {
	return uk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uk *uk_UA) MonthsWide() []string {
	return uk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayAbbreviated(weekday time.Weekday) string {
	return uk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uk *uk_UA) WeekdaysAbbreviated() []string {
	return uk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayNarrow(weekday time.Weekday) string {
	return uk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uk *uk_UA) WeekdaysNarrow() []string {
	return uk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayShort(weekday time.Weekday) string {
	return uk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uk *uk_UA) WeekdaysShort() []string {
	return uk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayWide(weekday time.Weekday) string {
	return uk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uk *uk_UA) WeekdaysWide() []string {
	return uk.daysWide
}

// Decimal returns the decimal point of number
func (uk *uk_UA) Decimal() string {
	return uk.decimal
}

// Group returns the group of number
func (uk *uk_UA) Group() string {
	return uk.group
}

// Group returns the minus sign of number
func (uk *uk_UA) Minus() string {
	return uk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uk_UA' and handles both Whole and Real numbers based on 'v'
func (uk *uk_UA) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uk_UA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uk *uk_UA) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uk_UA'
func (uk *uk_UA) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, uk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uk_UA'
// in accounting notation.
func (uk *uk_UA) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, uk.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, uk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, uk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd1, 0x80}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd1, 0x80}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, uk.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe2, 0x80, 0xaf, 0xd1, 0x80}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := uk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
