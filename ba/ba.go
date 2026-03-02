package ba

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ba struct {
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

// New returns a new instance of translator for the 'ba' locale
func New() locales.Translator {
	return &ba{
		locale:                 "ba",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "F CFA", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "ғин.", "фев.", "мар.", "апр.", "май", "июн.", "июл.", "авг.", "сент.", "окт.", "нояб.", "дек."},
		monthsNarrow:           []string{"", "Ғ", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "ғинуар", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
		daysAbbreviated:        []string{"йәк.", "дүш.", "шиш.", "шар.", "кес.", "йом.", "шәм."},
		daysNarrow:             []string{"Й", "Д", "Ш", "Ш", "К", "Й", "Ш"},
		daysShort:              []string{"йш", "дш", "шш", "шр", "кс", "йм", "шб"},
		daysWide:               []string{"йәкшәмбе", "дүшәмбе", "шишәмбе", "шаршамбы", "кесаҙна", "йома", "шәмбе"},
		timezones:              map[string]string{"ACDT": "Үҙәк Австралия йәйге ваҡыты", "ACST": "Акри йәйге ваҡыты", "ACT": "Акри стандарт ваҡыты", "ACWDT": "Үҙәк Австралия көнбайыш йәйге ваҡыты", "ACWST": "Үҙәк Австралия көнбайыш стандарт ваҡыты", "ADT": "Атлантик йәйге ваҡыты", "ADT Arabia": "Сәғүд Ғәрәбстаны йәйге ваҡыты", "AEDT": "Көнсығыш Австралия йәйге ваҡыты", "AEST": "Көнсығыш Австралия стандарт ваҡыты", "AFT": "Афғанстан ваҡыты", "AKDT": "Аляска йәйге ваҡыты", "AKST": "Аляска стандарт ваҡыты", "AMST": "Амазон йәйге ваҡыты", "AMST Armenia": "Әрмәнстан йәйге ваҡыты", "AMT": "Амазон стандарт ваҡыты", "AMT Armenia": "Әрмәнстан стандарт ваҡыты", "ANAST": "Анадырь йәйге ваҡыты", "ANAT": "Анадырь стандарт ваҡыты", "ARST": "Аргентина йәйге ваҡыты", "ART": "Аргентина стандарт ваҡыты", "AST": "Атлантик стандарт ваҡыты", "AST Arabia": "Сәғүд Ғәрәбстаны стандарт ваҡыты", "AWDT": "Көнбайыш Австралия йәйге ваҡыты", "AWST": "Көнбайыш Австралия стандарт ваҡыты", "AZST": "Әзербайжан йәйге ваҡыты", "AZT": "Әзербайжан стандарт ваҡыты", "BDT Bangladesh": "Бангладеш йәйге ваҡыты", "BNT": "Бруней ваҡыты", "BOT": "Боливия ваҡыты", "BRST": "Бразилия йәйге ваҡыты", "BRT": "Бразилия стандарт ваҡыты", "BST Bangladesh": "Бангладеш стандарт ваҡыты", "BT": "Бутан ваҡыты", "CAST": "Кейси ваҡыты", "CAT": "Үҙәк Африка ваҡыты", "CCT": "Кокос утрауҙары ваҡыты", "CDT": "Үҙәк Америка йәйге ваҡыты", "CHADT": "Чатем йәйге ваҡыты", "CHAST": "Чатем стандарт ваҡыты", "CHUT": "Чуук ваҡыты", "CKT": "Кук утрауҙары стандарт ваҡыты", "CKT DST": "Кук утрауҙары йәйге ваҡыты", "CLST": "Чили йәйге ваҡыты", "CLT": "Чили стандарт ваҡыты", "COST": "Колумбия йәйге ваҡыты", "COT": "Колумбия стандарт ваҡыты", "CST": "Үҙәк Америка стандарт ваҡыты", "CST China": "Ҡытай стандарт ваҡыты", "CST China DST": "Ҡытай йәйге ваҡыты", "CVST": "Кабо-Верде йәйге ваҡыты", "CVT": "Кабо-Верде стандарт ваҡыты", "CXT": "Раштыуа утрауы ваҡыты", "ChST": "Чаморро стандарт ваҡыты", "ChST NMI": "ChST NMI", "CuDT": "Куба йәйге ваҡыты", "CuST": "Куба стандарт ваҡыты", "DAVT": "Дейвис ваҡыты", "DDUT": "Дюмон-д’Юрвиль ваҡыты", "EASST": "Пасха утрауы йәйге ваҡыты", "EAST": "Пасха утрауы стандарт ваҡыты", "EAT": "Көнсығыш Африка ваҡыты", "ECT": "Эквадор ваҡыты", "EDT": "Көнсығыш Америка йәйге ваҡыты", "EGDT": "Көнсығыш Гренландия йәйге ваҡыты", "EGST": "Көнсығыш Гренландия стандарт ваҡыты", "EST": "Көнсығыш Америка стандарт ваҡыты", "FEET": "Мәскәү ваҡыты", "FJT": "Фиджи стандарт ваҡыты", "FJT Summer": "Фиджи йәйге ваҡыты", "FKST": "Фолкленд утрауҙары йәйге ваҡыты", "FKT": "Фолкленд утрауҙары стандарт ваҡыты", "FNST": "Фернанду-ди-Норонья йәйге ваҡыты", "FNT": "Фернанду-ди-Норонья стандарт ваҡыты", "GALT": "Галапагос ваҡыты", "GAMT": "Гамбье ваҡыты", "GEST": "Грузия йәйге ваҡыты", "GET": "Грузия стандарт ваҡыты", "GFT": "Француз Гвианаһы ваҡыты", "GIT": "Гилберт утрауҙары ваҡыты", "GMT": "Гринвич буйынса уртаса ваҡыт", "GNSST": "Гренландия йәйге ваҡыты", "GNST": "Гренландия стандарт ваҡыты", "GST": "Көньяҡ Георгия ваҡыты", "GST Guam": "Гуам стандарт ваҡыты", "GYT": "Гайана ваҡыты", "HADT": "Гавай-Алеут йәйге ваҡыты", "HAST": "Гавай-Алеут стандарт ваҡыты", "HKST": "Гонконг йәйге ваҡыты", "HKT": "Гонконг стандарт ваҡыты", "HOVST": "Ховд йәйге ваҡыты", "HOVT": "Ховд стандарт ваҡыты", "ICT": "Һинд-Ҡытай ваҡыты", "IDT": "Израиль йәйге ваҡыты", "IOT": "Һинд океаны ваҡыты", "IRKST": "Иркутск йәйге ваҡыты", "IRKT": "Иркутск стандарт ваҡыты", "IRST": "Иран стандарт ваҡыты", "IRST DST": "Иран йәйге ваҡыты", "IST": "Һиндостан стандарт ваҡыты", "IST Israel": "Израиль стандарт ваҡыты", "JDT": "Япония йәйге ваҡыты", "JST": "Япония стандарт ваҡыты", "KOST": "Косрае ваҡыты", "KRAST": "Красноярск йәйге ваҡыты", "KRAT": "Красноярск стандарт ваҡыты", "KST": "Корея стандарт ваҡыты", "KST DST": "Корея йәйге ваҡыты", "LHDT": "Лорд-Хау йәйге ваҡыты", "LHST": "Лорд-Хау стандарт ваҡыты", "LINT": "Лайн утрауҙары ваҡыты", "MAGST": "Магадан йәйге ваҡыты", "MAGT": "Магадан стандарт ваҡыты", "MART": "Маркиз утрауҙары ваҡыты", "MAWT": "Моусон ваҡыты", "MDT": "Макао йәйге ваҡыты", "MESZ": "Үҙәк Европа йәйге ваҡыты", "MEZ": "Үҙәк Европа стандарт ваҡыты", "MHT": "Маршалл Утрауҙары ваҡыты", "MMT": "Мьянма ваҡыты", "MSD": "Мәскәү йәйге ваҡыты", "MST": "Макао стандарт ваҡыты", "MUST": "Маврикий йәйге ваҡыты", "MUT": "Маврикий стандарт ваҡыты", "MVT": "Мальдив утрауҙары ваҡыты", "MYT": "Малайзия ваҡыты", "NCT": "Яңы Каледония стандарт ваҡыты", "NDT": "Ньюфаундленд йәйге ваҡыты", "NDT New Caledonia": "Яңы Каледония йәйге ваҡыты", "NFDT": "Норфолк йәйге ваҡыты", "NFT": "Норфолк стандарт ваҡыты", "NOVST": "Новосибирск йәйге ваҡыты", "NOVT": "Новосибирск стандарт ваҡыты", "NPT": "Непал ваҡыты", "NRT": "Науру ваҡыты", "NST": "Ньюфаундленд стандарт ваҡыты", "NUT": "Ниуэ ваҡыты", "NZDT": "Яңы Зеландия йәйге ваҡыты", "NZST": "Яңы Зеландия стандарт ваҡыты", "OESZ": "Көнсығыш Европа йәйге ваҡыты", "OEZ": "Көнсығыш Европа стандарт ваҡыты", "OMSST": "Омск йәйге ваҡыты", "OMST": "Омск стандарт ваҡыты", "PDT": "Тымыҡ океан йәйге ваҡыты", "PDTM": "Мексика Тымыҡ океан йәйге ваҡыты", "PETDT": "Камчатка йәйге ваҡыты", "PETST": "Камчатка стандарт ваҡыты", "PGT": "Папуа – Яңы Гвинея", "PHOT": "Феникс утрауҙары ваҡыты", "PKT": "Пакистан стандарт ваҡыты", "PKT DST": "Пакистан йәйге ваҡыты", "PMDT": "Сен-Пьер һәм Микелон йәйге ваҡыты", "PMST": "Сен-Пьер һәм Микелон стандарт ваҡыты", "PONT": "Понпеи ваҡыты", "PST": "Тымыҡ океан стандарт ваҡыты", "PST Philippine": "Филиппин стандарт ваҡыты", "PST Philippine DST": "Филиппин йәйге ваҡыты", "PST Pitcairn": "Питкэрн ваҡыты", "PSTM": "Мексика Тымыҡ океан стандарт ваҡыты", "PWT": "Палау ваҡыты", "PYST": "Парагвай йәйге ваҡыты", "PYT": "Парагвай стандарт ваҡыты", "PYT Korea": "Төньяҡ Корея ваҡыты", "RET": "Реюньон ваҡыты", "ROTT": "Ротера ваҡыты", "SAKST": "Сахалин йәйге ваҡыты", "SAKT": "Сахалин стандарт ваҡыты", "SAMST": "Һамар йәйге ваҡыты", "SAMT": "Һамар стандарт ваҡыты", "SAST": "Көньяҡ Африка стандарт ваҡыты", "SBT": "Соломон утрауҙары ваҡыты", "SCT": "Сейшел утрауҙары ваҡыты", "SGT": "Сингапур стандарт ваҡыты", "SLST": "Ланка ваҡыты", "SRT": "Суринам ваҡыты", "SST Samoa": "Америка Самоаһы стандарт ваҡыты", "SST Samoa Apia": "Самоа стандарт ваҡыты", "SST Samoa Apia DST": "Самоа йәйге ваҡыты", "SST Samoa DST": "Америка Самоаһы йәйге ваҡыты", "SYOT": "Сёва ваҡыты", "TAAF": "Француз Көньяҡ һәм Антарктика ваҡыты", "TAHT": "Таити ваҡыты", "TJT": "Тажикстан ваҡыты", "TKT": "Токелау ваҡыты", "TLT": "Көнсығыш Тимор ваҡыты", "TMST": "Төркмәнстан йәйге ваҡыты", "TMT": "Төркмәнстан стандарт ваҡыты", "TOST": "Тонга йәйге ваҡыты", "TOT": "Тонга стандарт ваҡыты", "TVT": "Тувалу ваҡыты", "TWT": "Тайвань стандарт ваҡыты", "TWT DST": "Тайвань йәйге ваҡыты", "ULAST": "Улан-Батор йәйге ваҡыты", "ULAT": "Улан-Батор стандарт ваҡыты", "UYST": "Уругвай йәйге ваҡыты", "UYT": "Уругвай стандарт ваҡыты", "UZT": "Үзбәкстан стандарт ваҡыты", "UZT DST": "Үзбәкстан йәйге ваҡыты", "VET": "Венесуэла ваҡыты", "VLAST": "Владивосток йәйге ваҡыты", "VLAT": "Владивосток стандарт ваҡыты", "VOLST": "Волгоград йәйге ваҡыты", "VOLT": "Волгоград стандарт ваҡыты", "VOST": "Восток ваҡыты", "VUT": "Вануату стандарт ваҡыты", "VUT DST": "Вануату йәйге ваҡыты", "WAKT": "Уэйк утрауы ваҡыты", "WARST": "Көнбайыш Аргентина йәйге ваҡыты", "WART": "Көнбайыш Аргентина стандарт ваҡыты", "WAST": "Көнбайыш Африка ваҡыты", "WAT": "Көнбайыш Африка ваҡыты", "WESZ": "Көнбайыш Европа йәйге ваҡыты", "WEZ": "Көнбайыш Европа стандарт ваҡыты", "WFT": "Уоллис һәм Футуна ваҡыты", "WGST": "Көнбайыш Гренландия йәйге ваҡыты", "WGT": "Көнбайыш Гренландия стандарт ваҡыты", "WIB": "Көнбайыш Индонезия", "WIT": "Көнсығыш Индонезия", "WITA": "Үҙәк Индонезия ваҡыты", "YAKST": "Якутск йәйге ваҡыты", "YAKT": "Якутск стандарт ваҡыты", "YEKST": "Екатеринбург йәйге ваҡыты", "YEKT": "Екатеринбург стандарт ваҡыты", "YST": "Юкон ваҡыты", "МСК": "Мәскәү стандарт ваҡыты", "اقتاۋ": "Аҡтау стандарт ваҡыты", "اقتاۋ قالاسى": "Аҡтау йәйге ваҡыты", "اقتوبە": "Аҡтүбә стандарт ваҡыты", "اقتوبە قالاسى": "Аҡтүбә йәйге ваҡыты", "الماتى": "Алматы стандарт ваҡыты", "الماتى قالاسى": "Алматы йәйге ваҡыты", "باتىس قازاق ەلى": "Көнбайыш Ҡаҙағстан ваҡыты", "شىعىش قازاق ەلى": "Көнсығыш Ҡаҙағстан ваҡыты", "قازاق ەلى": "Ҡаҙағстан ваҡыты", "قىرعىزستان": "Ҡырғыҙстан ваҡыты", "قىزىلوردا": "Ҡыҙылурҙа стандарт ваҡыты", "قىزىلوردا قالاسى": "Ҡыҙылурҙа йәйге ваҡыты", "∅∅∅": "Азор утрауҙары йәйге ваҡыты"},
	}
}

// Locale returns the current translators string locale
func (ba *ba) Locale() string {
	return ba.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ba'
func (ba *ba) PluralsCardinal() []locales.PluralRule {
	return ba.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ba'
func (ba *ba) PluralsOrdinal() []locales.PluralRule {
	return ba.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ba'
func (ba *ba) PluralsRange() []locales.PluralRule {
	return ba.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ba'
func (ba *ba) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ba'
func (ba *ba) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ba'
func (ba *ba) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ba *ba) MonthAbbreviated(month time.Month) string {
	return ba.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ba *ba) MonthsAbbreviated() []string {
	return ba.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ba *ba) MonthNarrow(month time.Month) string {
	return ba.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ba *ba) MonthsNarrow() []string {
	return ba.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ba *ba) MonthWide(month time.Month) string {
	return ba.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ba *ba) MonthsWide() []string {
	return ba.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ba *ba) WeekdayAbbreviated(weekday time.Weekday) string {
	return ba.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ba *ba) WeekdaysAbbreviated() []string {
	return ba.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ba *ba) WeekdayNarrow(weekday time.Weekday) string {
	return ba.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ba *ba) WeekdaysNarrow() []string {
	return ba.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ba *ba) WeekdayShort(weekday time.Weekday) string {
	return ba.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ba *ba) WeekdaysShort() []string {
	return ba.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ba *ba) WeekdayWide(weekday time.Weekday) string {
	return ba.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ba *ba) WeekdaysWide() []string {
	return ba.daysWide
}

// Decimal returns the decimal point of number
func (ba *ba) Decimal() string {
	return ba.decimal
}

// Group returns the group of number
func (ba *ba) Group() string {
	return ba.group
}

// Group returns the minus sign of number
func (ba *ba) Minus() string {
	return ba.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ba' and handles both Whole and Real numbers based on 'v'
func (ba *ba) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ba.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ba.group) - 1; j >= 0; j-- {
					b = append(b, ba.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ba.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ba' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ba *ba) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ba.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ba.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ba.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ba'
func (ba *ba) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ba.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ba.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ba.group) - 1; j >= 0; j-- {
					b = append(b, ba.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ba.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ba.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ba.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ba'
// in accounting notation.
func (ba *ba) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ba.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ba.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ba.group) - 1; j >= 0; j-- {
					b = append(b, ba.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ba.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ba.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ba.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ba.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ba'
func (ba *ba) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ba'
func (ba *ba) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ba.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd0, 0xb9}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ba'
func (ba *ba) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ba.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd0, 0xb9}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ba'
func (ba *ba) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ba.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ba.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd0, 0xb9}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ba'
func (ba *ba) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ba'
func (ba *ba) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ba'
func (ba *ba) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ba'
func (ba *ba) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ba.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ba.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
