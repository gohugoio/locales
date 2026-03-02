package cv_RU

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type cv_RU struct {
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

// New returns a new instance of translator for the 'cv_RU' locale
func New() locales.Translator {
	return &cv_RU{
		locale:                 "cv_RU",
		pluralsCardinal:        []locales.PluralRule{1, 2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "кӑр.", "нар.", "пуш", "ака", "ҫу", "ҫӗрт.", "утӑ", "ҫур.", "авӑн", "юпа", "чӳк", "раш."},
		monthsNarrow:           []string{"", "К", "Н", "П", "А", "Ҫ", "Ҫ", "У", "Ҫ", "А", "Ю", "Ч", "Р"},
		monthsWide:             []string{"", "кӑрлач", "нарӑс", "пуш", "ака", "ҫу", "ҫӗртме", "утӑ", "ҫурла", "авӑн", "юпа", "чӳк", "раштав"},
		daysAbbreviated:        []string{"вр", "тн", "ыт", "юн", "кҫ", "эр", "шм"},
		daysNarrow:             []string{"В", "Т", "Ы", "Ю", "К", "Э", "Ш"},
		daysShort:              []string{"вр", "тн", "ыт", "юн", "кҫ", "эр", "шм"},
		daysWide:               []string{"вырсарни кун", "тунти кун", "ытлари кун", "юн кун", "кӗҫнерни кун", "эрне кун", "шӑмат кун"},
		timezones:              map[string]string{"ACDT": "Вӑта Австрали ҫуллахи вӑхӑчӗ", "ACST": "Акри ҫуллахи вӑхӑчӗ", "ACT": "Акри стандартлӑ вӑхӑчӗ", "ACWDT": "Анӑҫ Вӑта Австрали ҫуллахи вӑхӑчӗ", "ACWST": "Анӑҫ Вӑта Австрали стандартлӑ вӑхӑчӗ", "ADT": "Атлантика ҫуллахи вӑхӑчӗ", "ADT Arabia": "Сауд Аравийӗ ҫуллахи вӑхӑчӗ", "AEDT": "Тухӑҫ Австрали ҫуллахи вӑхӑчӗ", "AEST": "Тухӑҫ Австрали стандартлӑ вӑхӑчӗ", "AFT": "Афганистан вӑхӑчӗ", "AKDT": "Аляска ҫуллахи вӑхӑчӗ", "AKST": "Аляска стандартлӑ вӑхӑчӗ", "AMST": "Амазонка ҫуллахи вӑхӑчӗ", "AMST Armenia": "Армени ҫуллахи вӑхӑчӗ", "AMT": "Амазонка стандартлӑ вӑхӑчӗ", "AMT Armenia": "Армени стандартлӑ вӑхӑчӗ", "ANAST": "Анадырь ҫуллахи вӑхӑчӗ", "ANAT": "Анадырь стандартлӑ вӑхӑчӗ", "ARST": "Аргентина ҫуллахи вӑхӑчӗ", "ART": "Аргентина стандартлӑ вӑхӑчӗ", "AST": "Атлантика стандартлӑ вӑхӑчӗ", "AST Arabia": "Сауд Аравийӗ стандартлӑ вӑхӑчӗ", "AWDT": "Анӑҫ Австрали ҫуллахи вӑхӑчӗ", "AWST": "Анӑҫ Австрали стандартлӑ вӑхӑчӗ", "AZST": "Азербайджан ҫуллахи вӑхӑчӗ", "AZT": "Азербайджан стандартлӑ вӑхӑчӗ", "BDT Bangladesh": "Бангладеш ҫуллахи вӑхӑчӗ", "BNT": "Бруней вӑхӑчӗ", "BOT": "Боливи вӑхӑчӗ", "BRST": "Бразили ҫуллахи вӑхӑчӗ", "BRT": "Бразили стандартлӑ вӑхӑчӗ", "BST Bangladesh": "Бангладеш стандартлӑ вӑхӑчӗ", "BT": "Бутан вӑхӑчӗ", "CAST": "Кейси вӑхӑчӗ", "CAT": "Вӑта Африка вӑхӑчӗ", "CCT": "Кокос утравӗсен вӑхӑчӗ", "CDT": "Вӑта Америка ҫуллахи вӑхӑчӗ", "CHADT": "Чатем ҫуллахи вӑхӑчӗ", "CHAST": "Чатем стандартлӑ вӑхӑчӗ", "CHUT": "Чуук вӑхӑчӗ", "CKT": "Кук утравӗсен стандартлӑ вӑхӑчӗ", "CKT DST": "Кук утравӗсен ҫуллахи вӑхӑчӗ", "CLST": "Чили ҫуллахи вӑхӑчӗ", "CLT": "Чили стандартлӑ вӑхӑчӗ", "COST": "Колумби ҫуллахи вӑхӑчӗ", "COT": "Колумби стандартлӑ вӑхӑчӗ", "CST": "Вӑта Америка стандартлӑ вӑхӑчӗ", "CST China": "Китай стандартлӑ вӑхӑчӗ", "CST China DST": "Китай ҫуллахи вӑхӑчӗ", "CVST": "Кабо-Верде ҫуллахи вӑхӑчӗ", "CVT": "Кабо-Верде стандартлӑ вӑхӑчӗ", "CXT": "Раштав утравӗ вӑхӑчӗ", "ChST": "Чаморро вӑхӑчӗ", "ChST NMI": "Ҫур ҫӗр Мариана утравӗсен вӑхӑчӗ", "CuDT": "Куба ҫуллахи вӑхӑчӗ", "CuST": "Куба стандартлӑ вӑхӑчӗ", "DAVT": "Дейвис вӑхӑчӗ", "DDUT": "Дюмон-д’Юрвиль вӑхӑчӗ", "EASST": "Мӑн кун утравӗн ҫуллахи вӑхӑчӗ", "EAST": "Мӑн кун утравӗн стандартлӑ вӑхӑчӗ", "EAT": "Тухӑҫ Африка вӑхӑчӗ", "ECT": "Эквадор вӑхӑчӗ", "EDT": "Тухӑҫ Америка ҫуллахи вӑхӑчӗ", "EGDT": "Тухӑҫ Гренланди ҫуллахи вӑхӑчӗ", "EGST": "Тухӑҫ Гренланди стандартлӑ вӑхӑчӗ", "EST": "Тухӑҫ Америка стандартлӑ вӑхӑчӗ", "FEET": "Инҫет Тухӑҫ Европа вӑхӑчӗ", "FJT": "Фиджи стандартлӑ вӑхӑчӗ", "FJT Summer": "Фиджи ҫуллахи вӑхӑчӗ", "FKST": "Фолкленд утравӗсен ҫуллахи вӑхӑчӗ", "FKT": "Фолкленд утравӗсен стандартлӑ вӑхӑчӗ", "FNST": "Фернанду-ди-Норонья ҫуллахи вӑхӑчӗ", "FNT": "Фернанду-ди-Норонья стандартлӑ вӑхӑчӗ", "GALT": "Галапагос вӑхӑчӗ", "GAMT": "Гамбье вӑхӑчӗ", "GEST": "Грузи ҫуллахи вӑхӑчӗ", "GET": "Грузи стандартлӑ вӑхӑчӗ", "GFT": "Франци Гвианин вӑхӑчӗ", "GIT": "Гилберт утравӗсен вӑхӑчӗ", "GMT": "Гринвич вӑхӑчӗ", "GNSST": "Гренланди ҫуллахи вӑхӑчӗ", "GNST": "Гренланди стандартлӑ вӑхӑчӗ", "GST": "Перс кӳлмекӗн вӑхӑчӗ", "GST Guam": "Гуам стандартлӑ вӑхӑчӗ", "GYT": "Гайана вӑхӑчӗ", "HADT": "Гавайи-Алеут стандартлӑ вӑхӑчӗ", "HAST": "Гавайи-Алеут стандартлӑ вӑхӑчӗ", "HKST": "Гонконг ҫуллахи вӑхӑчӗ", "HKT": "Гонконг стандартлӑ вӑхӑчӗ", "HOVST": "Ховд ҫуллахи вӑхӑчӗ", "HOVT": "Ховд стандартлӑ вӑхӑчӗ", "ICT": "Индокитай вӑхӑчӗ", "IDT": "Израиль ҫуллахи вӑхӑчӗ", "IOT": "Инди океанӗ вӑхӑчӗ", "IRKST": "Иркутск ҫуллахи вӑхӑчӗ", "IRKT": "Иркутск стандартлӑ вӑхӑчӗ", "IRST": "Иран стандартлӑ вӑхӑчӗ", "IRST DST": "Иран ҫуллахи вӑхӑчӗ", "IST": "Инди вӑхӑчӗ", "IST Israel": "Израиль стандартлӑ вӑхӑчӗ", "JDT": "Япони ҫуллахи вӑхӑчӗ", "JST": "Япони стандартлӑ вӑхӑчӗ", "KOST": "Косрае вӑхӑчӗ", "KRAST": "Красноярск ҫуллахи вӑхӑчӗ", "KRAT": "Красноярск стандартлӑ вӑхӑчӗ", "KST": "Корея стандартлӑ вӑхӑчӗ", "KST DST": "Корея ҫуллахи вӑхӑчӗ", "LHDT": "Лорд-Хау ҫуллахи вӑхӑчӗ", "LHST": "Лорд-Хау стандартлӑ вӑхӑчӗ", "LINT": "Лайн утравӗсен вӑхӑчӗ", "MAGST": "Магадан ҫуллахи вӑхӑчӗ", "MAGT": "Магадан стандартлӑ вӑхӑчӗ", "MART": "Маркизас утравӗсен вӑхӑчӗ", "MAWT": "Моусон вӑхӑчӗ", "MDT": "Ту ҫуллахи вӑхӑчӗ", "MESZ": "Вӑта Европа ҫуллахи вӑхӑчӗ", "MEZ": "Вӑта Европа стандартлӑ вӑхӑчӗ", "MHT": "Маршалл утравӗсен вӑхӑчӗ", "MMT": "Мьянма вӑхӑчӗ", "MSD": "Мускав ҫуллахи вӑхӑчӗ", "MST": "Ту стандартлӑ вӑхӑчӗ", "MUST": "Маврики ҫуллахи вӑхӑчӗ", "MUT": "Маврики стандартлӑ вӑхӑчӗ", "MVT": "Мальдива вӑхӑчӗ", "MYT": "Малайзи вӑхӑчӗ", "NCT": "Ҫӗнӗ Каледони стандартлӑ вӑхӑчӗ", "NDT": "Ньюфаундленд ҫуллахи вӑхӑчӗ", "NDT New Caledonia": "Ҫӗнӗ Каледони ҫуллахи вӑхӑчӗ", "NFDT": "Норфолк ҫуллахи вӑхӑчӗ", "NFT": "Норфолк стандартлӑ вӑхӑчӗ", "NOVST": "Новосибирск ҫуллахи вӑхӑчӗ", "NOVT": "Новосибирск стандартлӑ вӑхӑчӗ", "NPT": "Непал вӑхӑчӗ", "NRT": "Науру вӑхӑчӗ", "NST": "Ньюфаундленд стандартлӑ вӑхӑчӗ", "NUT": "Ниуэ вӑхӑчӗ", "NZDT": "Ҫӗнӗ Зеланди ҫуллахи вӑхӑчӗ", "NZST": "Ҫӗнӗ Зеланди стандартлӑ вӑхӑчӗ", "OESZ": "Тухӑҫ Европа ҫуллахи вӑхӑчӗ", "OEZ": "Тухӑҫ Европа стандартлӑ вӑхӑчӗ", "OMSST": "Омск ҫуллахи вӑхӑчӗ", "OMST": "Омск стандартлӑ вӑхӑчӗ", "PDT": "Лӑпкӑ океан ҫуллахи вӑхӑчӗ", "PDTM": "Мексика Лӑпкӑ океан ҫуллахи вӑхӑчӗ", "PETDT": "Петропавловск-Камчатский ҫуллахи вӑхӑчӗ", "PETST": "Петропавловск-Камчатский стандартлӑ вӑхӑчӗ", "PGT": "Папуа — Ҫӗнӗ Гвинея вӑхӑчӗ", "PHOT": "Феникс утравӗсен вӑхӑчӗ", "PKT": "Пакистан стандартлӑ вӑхӑчӗ", "PKT DST": "Пакистан ҫуллахи вӑхӑчӗ", "PMDT": "Сен-Пьерпа Микелон ҫуллахи вӑхӑчӗ", "PMST": "Сен-Пьерпа Микелон стандартлӑ вӑхӑчӗ", "PONT": "Понпеи вӑхӑчӗ", "PST": "Лӑпкӑ океан стандартлӑ вӑхӑчӗ", "PST Philippine": "Филиппин стандартлӑ вӑхӑчӗ", "PST Philippine DST": "Филиппин ҫуллахи вӑхӑчӗ", "PST Pitcairn": "Питкэрн вӑхӑчӗ", "PSTM": "Мексика Лӑпкӑ океан стандартлӑ вӑхӑчӗ", "PWT": "Палау вӑхӑчӗ", "PYST": "Парагвай ҫуллахи вӑхӑчӗ", "PYT": "Парагвай стандартлӑ вӑхӑчӗ", "PYT Korea": "Ҫур ҫӗр Корея вӑхӑчӗ", "RET": "Реюньон вӑхӑчӗ", "ROTT": "Ротера вӑхӑчӗ", "SAKST": "Сахалин ҫуллахи вӑхӑчӗ", "SAKT": "Сахалин стандартлӑ вӑхӑчӗ", "SAMST": "Самар ҫуллахи вӑхӑчӗ", "SAMT": "Самар стандартлӑ вӑхӑчӗ", "SAST": "Кӑнтӑр Африка вӑхӑчӗ", "SBT": "Соломон вӑхӑчӗ", "SCT": "Сейшел утравӗсен вӑхӑчӗ", "SGT": "Сингапур вӑхӑчӗ", "SLST": "Шри-Ланка вӑхӑчӗ", "SRT": "Суринам вӑхӑчӗ", "SST Samoa": "Самоа стандартлӑ вӑхӑчӗ", "SST Samoa Apia": "Апиа стандартлӑ вӑхӑчӗ", "SST Samoa Apia DST": "Апиа ҫуллахи вӑхӑчӗ", "SST Samoa DST": "Самоа ҫуллахи вӑхӑчӗ", "SYOT": "Сёва вӑхӑчӗ", "TAAF": "Франци Кӑнтӑрпа Антарктика территорийӗсен вӑхӑчӗ", "TAHT": "Таити вӑхӑчӗ", "TJT": "Таджикистан вӑхӑчӗ", "TKT": "Токелау вӑхӑчӗ", "TLT": "Тухӑҫ Тимор вӑхӑчӗ", "TMST": "Туркменистан ҫуллахи вӑхӑчӗ", "TMT": "Туркменистан стандартлӑ вӑхӑчӗ", "TOST": "Тонга ҫуллахи вӑхӑчӗ", "TOT": "Тонга стандартлӑ вӑхӑчӗ", "TVT": "Тувалу вӑхӑчӗ", "TWT": "Тайвань стандартлӑ вӑхӑчӗ", "TWT DST": "Тайвань ҫуллахи вӑхӑчӗ", "ULAST": "Улан-Батор ҫуллахи вӑхӑчӗ", "ULAT": "Улан-Батор стандартлӑ вӑхӑчӗ", "UYST": "Уругвай ҫуллахи вӑхӑчӗ", "UYT": "Уругвай стандартлӑ вӑхӑчӗ", "UZT": "Узбекистан стандартлӑ вӑхӑчӗ", "UZT DST": "Узбекистан ҫуллахи вӑхӑчӗ", "VET": "Венесуэла вӑхӑчӗ", "VLAST": "Владивосток ҫуллахи вӑхӑчӗ", "VLAT": "Владивосток стандартлӑ вӑхӑчӗ", "VOLST": "Волгоград ҫуллахи вӑхӑчӗ", "VOLT": "Волгоград стандартлӑ вӑхӑчӗ", "VOST": "Восток вӑхӑчӗ", "VUT": "Вануату стандартлӑ вӑхӑчӗ", "VUT DST": "Вануату ҫуллахи вӑхӑчӗ", "WAKT": "Уэйк утравӗ вӑхӑчӗ", "WARST": "Анӑҫ Аргентина ҫуллахи вӑхӑчӗ", "WART": "Анӑҫ Аргентина стандартлӑ вӑхӑчӗ", "WAST": "Анӑҫ Африка вӑхӑчӗ", "WAT": "Анӑҫ Африка вӑхӑчӗ", "WESZ": "Анӑҫ Европа ҫуллахи вӑхӑчӗ", "WEZ": "Анӑҫ Европа стандартлӑ вӑхӑчӗ", "WFT": "Уоллиспа Футуна вӑхӑчӗ", "WGST": "Анӑҫ ҫуллахи вӑхӑчӗ", "WGT": "Анӑҫ Гринланди стандартлӑ вӑхӑчӗ", "WIB": "Анӑҫ Индонези вӑхӑчӗ", "WIT": "Тухӑҫ Индонези вӑхӑчӗ", "WITA": "Вӑта Индонези вӑхӑчӗ", "YAKST": "Якутск ҫуллахи вӑхӑчӗ", "YAKT": "Якутск стандартлӑ вӑхӑчӗ", "YEKST": "Екатеринбург ҫуллахи вӑхӑчӗ", "YEKT": "Екатеринбург стандартлӑ вӑхӑчӗ", "YST": "Юкон вӑхӑчӗ", "МСК": "Мускав стандартлӑ вӑхӑчӗ", "اقتاۋ": "Актау стандартлӑ вӑхӑчӗ", "اقتاۋ قالاسى": "Актау ҫуллахи вӑхӑчӗ", "اقتوبە": "Актобе стандартлӑ вӑхӑчӗ", "اقتوبە قالاسى": "Актобе ҫуллахи вӑхӑчӗ", "الماتى": "Алматы стандартлӑ вӑхӑчӗ", "الماتى قالاسى": "Алматы ҫуллахи вӑхӑчӗ", "باتىس قازاق ەلى": "Анӑҫ Казахстан вӑхӑчӗ", "شىعىش قازاق ەلى": "Тухӑҫ Казахстан вӑхӑчӗ", "قازاق ەلى": "Казахстан вӑхӑчӗ", "قىرعىزستان": "Кӑркӑстан вӑхӑчӗ", "قىزىلوردا": "Кызылорда стандартлӑ вӑхӑчӗ", "قىزىلوردا قالاسى": "Кызылорда ҫуллахи вӑхӑчӗ", "∅∅∅": "Азор утравӗсен ҫуллахи вӑхӑчӗ"},
	}
}

// Locale returns the current translators string locale
func (cv *cv_RU) Locale() string {
	return cv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'cv_RU'
func (cv *cv_RU) PluralsCardinal() []locales.PluralRule {
	return cv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'cv_RU'
func (cv *cv_RU) PluralsOrdinal() []locales.PluralRule {
	return cv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'cv_RU'
func (cv *cv_RU) PluralsRange() []locales.PluralRule {
	return cv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'cv_RU'
func (cv *cv_RU) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'cv_RU'
func (cv *cv_RU) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'cv_RU'
func (cv *cv_RU) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (cv *cv_RU) MonthAbbreviated(month time.Month) string {
	return cv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (cv *cv_RU) MonthsAbbreviated() []string {
	return cv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (cv *cv_RU) MonthNarrow(month time.Month) string {
	return cv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (cv *cv_RU) MonthsNarrow() []string {
	return cv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (cv *cv_RU) MonthWide(month time.Month) string {
	return cv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (cv *cv_RU) MonthsWide() []string {
	return cv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (cv *cv_RU) WeekdayAbbreviated(weekday time.Weekday) string {
	return cv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (cv *cv_RU) WeekdaysAbbreviated() []string {
	return cv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (cv *cv_RU) WeekdayNarrow(weekday time.Weekday) string {
	return cv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (cv *cv_RU) WeekdaysNarrow() []string {
	return cv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (cv *cv_RU) WeekdayShort(weekday time.Weekday) string {
	return cv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (cv *cv_RU) WeekdaysShort() []string {
	return cv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (cv *cv_RU) WeekdayWide(weekday time.Weekday) string {
	return cv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (cv *cv_RU) WeekdaysWide() []string {
	return cv.daysWide
}

// Decimal returns the decimal point of number
func (cv *cv_RU) Decimal() string {
	return cv.decimal
}

// Group returns the group of number
func (cv *cv_RU) Group() string {
	return cv.group
}

// Group returns the minus sign of number
func (cv *cv_RU) Minus() string {
	return cv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'cv_RU' and handles both Whole and Real numbers based on 'v'
func (cv *cv_RU) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(cv.group) - 1; j >= 0; j-- {
					b = append(b, cv.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'cv_RU' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (cv *cv_RU) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, cv.percentSuffix...)

	b = append(b, cv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'cv_RU'
func (cv *cv_RU) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(cv.group) - 1; j >= 0; j-- {
					b = append(b, cv.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, cv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, cv.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'cv_RU'
// in accounting notation.
func (cv *cv_RU) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cv.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(cv.group) - 1; j >= 0; j-- {
					b = append(b, cv.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cv.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, cv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, cv.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, cv.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, cv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, cv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, cv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, cv.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cv.timeSeparator...)

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

// FmtTimeFull returns the full time representation of 't' for 'cv_RU'
func (cv *cv_RU) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := cv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
