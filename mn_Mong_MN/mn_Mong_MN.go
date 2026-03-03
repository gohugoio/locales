package mn_Mong_MN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mn_Mong_MN struct {
	locale            string
	pluralsCardinal   []locales.PluralRule
	pluralsOrdinal    []locales.PluralRule
	pluralsRange      []locales.PluralRule
	decimal           string
	group             string
	minus             string
	percent           string
	timeSeparator     string
	currencies        []string // idx = enum of currency code
	monthsAbbreviated []string
	monthsNarrow      []string
	monthsWide        []string
	daysAbbreviated   []string
	daysNarrow        []string
	daysShort         []string
	daysWide          []string
	timezones         map[string]string
}

// New returns a new instance of translator for the 'mn_Mong_MN' locale
func New() locales.Translator {
	return &mn_Mong_MN{
		locale:            "mn_Mong_MN",
		pluralsCardinal:   []locales.PluralRule{2, 6},
		pluralsOrdinal:    []locales.PluralRule{6},
		pluralsRange:      []locales.PluralRule{2, 6},
		decimal:           ",",
		group:             " ",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "кр", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "1\u202f᠊ᠷ ᠰᠠᠷ\u180eᠠ", "2\u202f᠊ᠷ ᠰᠠᠷ\u180eᠠ", "3᠊ᠷ ᠰᠠᠷ\u180eᠠ", "4\u202f᠊ᠷ ᠰᠠᠷ\u180eᠠ", "5\u202f᠊ᠷ ᠰᠠᠷ\u180eᠠ", "6\u202f᠊ᠷ ᠰᠠᠷ\u180eᠠ", "7\u202f᠊ᠷ ᠰᠠᠷ\u180eᠠ", "8᠊ᠷ ᠰᠠᠷ\u180eᠠ", "9 ᠊ᠷ ᠰᠠᠷ\u180eᠠ", "10 ᠊ᠷ ᠰᠠᠷ\u180eᠠ", "11 ᠊ᠷ ᠰᠠᠷ\u180eᠠ", "12 ᠊ᠷ ᠰᠠᠷ\u180eᠠ"},
		monthsNarrow:      []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII"},
		monthsWide:        []string{"", "ᠨᠢᠭᠡᠳᠥᠭᠡᠷ ᠰᠠᠷ\u180eᠠ", "ᠬᠣᠶᠠᠳᠣᠭᠠᠷ ᠰᠠᠷ\u202fᠠ", "ᠭᠣᠷᠪᠡᠳᠣᠭᠠᠷ ᠰᠠᠷ\u202fᠠ", "ᠳᠥᠷᠪᠡᠳᠥᠭᠡᠷ ᠰᠠᠷ\u180eᠠ", "ᠲᠠᠪᠣᠳᠣᠭᠠᠷ ᠰᠠᠷ\u202fᠠ", "ᠵᠢᠷᠭᠣᠭᠠᠳᠣᠭᠠᠷ ᠰᠠᠷ\u180eᠠ", "ᠲᠣᠯᠣᠭᠠᠳᠣᠭᠠᠷ ᠰᠠᠷ\u180eᠠ", "ᠨᠠᠢᠮᠠᠳᠥᠭᠠᠷ ᠰᠠᠷ\u180eᠠ", "ᠶᠢᠰᠥᠳᠥᠭᠡᠷ ᠰᠠᠷ\u180eᠠ", "ᠠᠷᠪᠠᠳᠣᠭᠠᠷ ᠰᠠᠷ\u180eᠠ", "ᠠᠷᠪᠠᠨ ᠨᠢᠭᠡᠳᠥᠭᠡᠷ ᠰᠠᠷ\u180eᠠ", "ᠠᠷᠪᠠᠨ ᠬᠣᠶᠠᠳᠣᠭᠠᠷ ᠰᠠᠷ\u180eᠠ"},
		daysAbbreviated:   []string{"ᠨᠢ", "ᠲᠠ", "ᠮᠢᠭ", "ᡀᠠ", "ᠫᠥᠷ", "ᠪᠠ", "ᠪᠢᠮ"},
		daysNarrow:        []string{"ᠨᠢ", "ᠳᠠ", "ᠮᠢᠭ", "ᡀᠠ", "ᠫᠥᠷ", "ᠪᠠ", "ᠪᠢ"},
		daysWide:          []string{"ᠨᠢᠮ\u180eᠠ", "ᠳᠠᠸᠠ", "ᠮᠢᠠᠠᠮᠠᠷ", "ᡀᠠᠭᠪᠠ", "ᠫᠦᠷᠪᠦ", "ᠪᠠᠰᠠᠩ", "ᠪᠢᠮᠪᠠ"},
		timezones:         map[string]string{"ACDT": "Төв Австралийн зуны цаг", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Австралийн төв баруун эргийн зуны цаг", "ACWST": "Австралийн төв баруун эргийн стандарт цаг", "ADT": "ᠠᠲ᠋ᠯᠠᠨ᠋ᠲ \u180eᠤᠨ ᠵᠣᠨ \u180eᠪ ᠴᠠᠭ", "ADT Arabia": "Арабын зуны цаг", "AEDT": "Австралийн зүүн эргийн зуны цаг", "AEST": "Австралийн зүүн эргийн стандарт цаг", "AFT": "Афганистаны цаг", "AKDT": "Аляскийн зуны цаг", "AKST": "Аляскийн стандарт цаг", "AMST": "Амазоны зуны цаг", "AMST Armenia": "Арменийн зуны цаг", "AMT": "Амазоны стандарт цаг", "AMT Armenia": "Арменийн стандарт цаг", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентины зуны цаг", "ART": "Аргентины стандарт цаг", "AST": "ᠠᠲ᠋ᠯᠠᠨ᠋ᠲ \u180eᠤᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "AST Arabia": "Арабын стандарт цаг", "AWDT": "Австралийн баруун эргийн зуны цаг", "AWST": "Австралийн баруун эргийн стандарт цаг", "AZST": "Азербайжаны зуны цаг", "AZT": "Азербайжаны стандарт цаг", "BDT Bangladesh": "Бангладешийн зуны цаг", "BNT": "Бруней Даруссаламын цаг", "BOT": "Боливийн цаг", "BRST": "Бразилийн зуны цаг", "BRT": "Бразилийн стандарт цаг", "BST Bangladesh": "Бангладешийн стандарт цаг", "BT": "Бутаны цаг", "CAST": "CAST", "CAT": "Төв Африкийн цаг", "CCT": "Кокос арлын цаг", "CDT": "ᠲᠥᠪ ᠵᠣᠨ \u180e\u180e\u180eᠤ ᠴᠠᠭ", "CHADT": "Чатемын зуны цаг", "CHAST": "Чатемын стандарт цаг", "CHUT": "Чүүкийн цаг", "CKT": "Күүкийн арлуудын стандарт цаг", "CKT DST": "Күүкийн арлуудын хагас зуны цаг", "CLST": "Чилийн зуны цаг", "CLT": "Чилийн стандарт цаг", "COST": "Колумбын зуны цаг", "COT": "Колумбын стандарт цаг", "CST": "ᠳᠥᠪ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "CST China": "Хятадын стандарт цаг", "CST China DST": "Хятадын зуны цаг", "CVST": "Кабо-Вердийн зуны цаг", "CVT": "Кабо-Вердийн стандарт цаг", "CXT": "Крисмас арлын цаг", "ChST": "Чаморрогийн цаг", "ChST NMI": "ChST NMI", "CuDT": "Кубын зуны цаг", "CuST": "Кубын стандарт цаг", "DAVT": "Дэвисийн цаг", "DDUT": "Дюмон д’Юрвилийн цаг", "EASST": "Зүүн Исландын зуны цаг", "EAST": "Зүүн Исландын стандарт цаг", "EAT": "Зүүн Африкийн цаг", "ECT": "Эквадорын цаг", "EDT": "ᠵᠡᠭᠥᠨ ᠡᠷᠭᠡ \u180eᠢᠢᠨ ᠵᠣᠨ \u180e\u180e\u180eᠤ ᠴᠠᠭ", "EGDT": "Зүүн Гренландын зуны цаг", "EGST": "Зүүн Гренландын стандарт цаг", "EST": "ᠵᠡᠭᠥᠨ ᠡᠷᠭᠡ \u180eᠢᠢᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "FEET": "Алс дорнод Европын цаг", "FJT": "Фижигийн стандарт цаг", "FJT Summer": "Фижигийн зуны цаг", "FKST": "Фолклендийн арлуудын зуны цаг", "FKT": "Фолклендийн арлуудын стандарт цаг", "FNST": "Фернандо де Норонагийн зуны цаг", "FNT": "Фернандо де Норонагийн стандарт цаг", "GALT": "Галапагосын цаг", "GAMT": "Гамбьегийн цаг", "GEST": "Гүржийн зуны цаг", "GET": "Гүржийн стандарт цаг", "GFT": "Францын Гвианагийн цаг", "GIT": "Гильбертийн арлуудын цаг", "GMT": "ᠭᠷᠢᠨ᠋ᠸᠢᠴᠢ ᠢᠢᠨ ᠴᠠᠭ", "GNSST": "GNSST", "GNST": "GNST", "GST": "Персийн булангийн цаг", "GST Guam": "GST Guam", "GYT": "Гайанагийн цаг", "HADT": "Хавай-Алеутын зуны цаг", "HAST": "Хавай-Алеутын стандарт цаг", "HKST": "Хонг Конгийн зуны цаг", "HKT": "Хонг Конгийн стандарт цаг", "HOVST": "Ховдын зуны цаг", "HOVT": "Ховдын стандарт цаг", "ICT": "Энэтхэг-Хятадын хойгийн цаг", "IDT": "Израилийн зуны цаг", "IOT": "Энэтхэгийн далайн цаг", "IRKST": "Эрхүүгийн зуны цаг", "IRKT": "Эрхүүгийн стандарт цаг", "IRST": "Ираны стандарт цаг", "IRST DST": "Ираны зуны цаг", "IST": "Энэтхэгийн цаг", "IST Israel": "Израилийн стандарт цаг", "JDT": "Японы зуны цаг", "JST": "Японы стандарт цаг", "KOST": "Косрэгийн цаг", "KRAST": "Красноярскийн зуны цаг", "KRAT": "Красноярскийн стандарт цаг", "KST": "Солонгосын стандарт цаг", "KST DST": "Солонгосын зуны цаг", "LHDT": "Лорд Хоугийн зуны цаг", "LHST": "Лорд Хоугийн стандарт цаг", "LINT": "Лайн арлуудын цаг", "MAGST": "Магаданы зуны цаг", "MAGT": "Магаданы стандарт цаг", "MART": "Маркезын арлуудын цаг", "MAWT": "Моусоны цаг", "MDT": "ᠠᠭᠣᠯᠠ \u180eᠢᠢᠨ ᠵᠣᠨ \u180e\u180eᠤ ᠴᠠᠭ", "MESZ": "ᠲᠥᠪ ᠡᠸᠣᠢᠷᠤᠫᠠ ᠢᠢᠨ ᠵᠣᠨ \u180e\u180eᠤ ᠴᠠᠭ", "MEZ": "ᠲᠥᠪ ᠡᠸᠣᠢᠷᠤᠫᠠ ᠢᠢᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "MHT": "Маршаллын арлуудын цаг", "MMT": "Мьянмарын цаг", "MSD": "Москвагийн зуны цаг", "MST": "ᠠᠭᠣᠯᠠ \u180e\u180e\u180e\u180eᠢᠢᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "MUST": "Маврикийн зуны цаг", "MUT": "Маврикийн стандарт цаг", "MVT": "Мальдивийн цаг", "MYT": "Малайзын цаг", "NCT": "Шинэ Каледонийн стандарт цаг", "NDT": "Нью-Фаундлендын зуны цаг", "NDT New Caledonia": "Шинэ Каледонийн зуны цаг", "NFDT": "Норфолк арлын зуны цаг", "NFT": "Норфолк арлын стандарт цаг", "NOVST": "Новосибирскийн зуны цаг", "NOVT": "Новосибирскийн стандарт цаг", "NPT": "Балбын цаг", "NRT": "Науругийн цаг", "NST": "Нью-Фаундлендын стандарт цаг", "NUT": "Ниуэгийн цаг", "NZDT": "Шинэ Зеландын зуны цаг", "NZST": "Шинэ Зеландын стандарт цаг", "OESZ": "ᠵᠡᠭᠦᠨ ᠡᠸᠣᠢᠷᠤᠫᠠ ᠢᠢᠨ ᠵᠣᠨ \u180e\u180eᠤ ᠴᠠᠭ", "OEZ": "ᠵᠡᠭᠦᠨ ᠡᠸᠣᠢᠷᠤᠫᠠ ᠢᠢᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "OMSST": "Омскийн зуны цаг", "OMST": "Омскийн стандарт цаг", "PDT": "ᠨᠣᠮᠣᠬᠠᠨ ᠳᠠᠯᠠᠢ \u180eᠢᠢᠨ ᠵᠣᠨ \u180e\u180e\u180eᠪ ᠴᠠᠭ", "PDTM": "Мексик-Номхон далайн зуны цаг", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Папуа Шинэ Гвинейн цаг", "PHOT": "Феникс арлын цаг", "PKT": "Пакистаны стандарт цаг", "PKT DST": "Пакистаны зуны цаг", "PMDT": "Сент-Пьер ба Микелоны зуны цаг", "PMST": "Сент-Пьер ба Микелоны стандарт цаг", "PONT": "Понапегийн цаг", "PST": "ᠨᠣᠮᠣᠬᠠᠨ ᠳᠠᠯᠠᠢ \u180e\u180eᠢᠢᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "PST Philippine": "Филиппиний стандарт цаг", "PST Philippine DST": "Филиппиний зуны цаг", "PST Pitcairn": "Питкернийн цаг", "PSTM": "Мексик-Номхон далайн стандарт цаг", "PWT": "Палаугийн цаг", "PYST": "Парагвайн зуны цаг", "PYT": "Парагвайн стандарт цаг", "PYT Korea": "Пёньяны цаг", "RET": "Реюнионы цаг", "ROTT": "Ротерагийн цаг", "SAKST": "Сахалины зуны цаг", "SAKT": "Сахалины стандарт цаг", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Өмнөд Африкийн стандарт цаг", "SBT": "Соломоны арлуудын цаг", "SCT": "Сейшелийн арлуудын цаг", "SGT": "Сингапурын цаг", "SLST": "SLST", "SRT": "Суринамын цаг", "SST Samoa": "Самоагийн стандарт цаг", "SST Samoa Apia": "Апиагийн стандарт цаг", "SST Samoa Apia DST": "Апиагийн зуны цаг", "SST Samoa DST": "Самоагийн зуны цаг", "SYOT": "Сёвагийн цаг", "TAAF": "Францын Өмнөд ба Антарктидийн цаг", "TAHT": "Таитигийн цаг", "TJT": "Тажикистаны цаг", "TKT": "Токелаугийн цаг", "TLT": "Зүүн Тиморын цаг", "TMST": "Туркменистаны зуны цаг", "TMT": "Туркменистаны стандарт цаг", "TOST": "Тонгагийн зуны цаг", "TOT": "Тонгагийн стандарт цаг", "TVT": "Тувалугийн цаг", "TWT": "Тайпейн стандарт цаг", "TWT DST": "Тайпейн зуны цаг", "ULAST": "Улаанбаатарын зуны цаг", "ULAT": "Улаанбаатарын стандарт цаг", "UYST": "Уругвайн зуны цаг", "UYT": "Уругвайн стандарт цаг", "UZT": "Узбекистаны стандарт цаг", "UZT DST": "Узбекистаны зуны цаг", "VET": "Венесуэлийн цаг", "VLAST": "Владивостокийн зуны цаг", "VLAT": "Владивостокийн стандарт цаг", "VOLST": "Волгоградын зуны цаг", "VOLT": "Волгоградын стандарт цаг", "VOST": "Востокийн цаг", "VUT": "Вануатугийн стандарт цаг", "VUT DST": "Вануатугийн зуны цаг", "WAKT": "Уэкийн арлуудын цаг", "WARST": "Баруун Аргентины зуны цаг", "WART": "Баруун Аргентины стандарт цаг", "WAST": "Баруун Африкийн цаг", "WAT": "Баруун Африкийн цаг", "WESZ": "ᠪᠠᠷᠠᠭᠣᠨ ᠡᠸᠣᠢᠷᠤᠫᠠ ᠢᠢᠨ ᠵᠣᠨ \u180e\u180eᠤ ᠴᠠᠭ", "WEZ": "ᠪᠠᠷᠠᠭᠣᠨ ᠡᠸᠣᠢᠷᠤᠫᠠ ᠢᠢᠨ ᠰᠲ᠋ᠠᠨ᠋ᠳᠠᠷᠳ᠋ ᠴᠠᠭ", "WFT": "Уоллис ба Футунагийн цаг", "WGST": "Баруун Гренландын зуны цаг", "WGT": "Баруун Гренландын стандарт цаг", "WIB": "Баруун Индонезийн цаг", "WIT": "Зүүн Индонезийн цаг", "WITA": "Төв Индонезийн цаг", "YAKST": "Якутын зуны цаг", "YAKT": "Якутын стандарт цаг", "YEKST": "Екатеринбургийн зуны цаг", "YEKT": "Екатеринбургийн стандарт цаг", "YST": "Юкон цагийн бүс", "МСК": "Москвагийн стандарт цаг", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Баруун Казахстаны цаг", "شىعىش قازاق ەلى": "Зүүн Казахстаны цаг", "قازاق ەلى": "Казахстаны цаг", "قىرعىزستان": "Киргизийн цаг", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Азорын зуны цаг"},
	}
}

// Locale returns the current translators string locale
func (mn *mn_Mong_MN) Locale() string {
	return mn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mn_Mong_MN'
func (mn *mn_Mong_MN) PluralsCardinal() []locales.PluralRule {
	return mn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mn_Mong_MN'
func (mn *mn_Mong_MN) PluralsOrdinal() []locales.PluralRule {
	return mn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mn_Mong_MN'
func (mn *mn_Mong_MN) PluralsRange() []locales.PluralRule {
	return mn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

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
func (mn *mn_Mong_MN) MonthAbbreviated(month time.Month) string {
	if len(mn.monthsAbbreviated) == 0 {
		return ""
	}
	return mn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mn *mn_Mong_MN) MonthsAbbreviated() []string {
	return mn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mn *mn_Mong_MN) MonthNarrow(month time.Month) string {
	if len(mn.monthsNarrow) == 0 {
		return ""
	}
	return mn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mn *mn_Mong_MN) MonthsNarrow() []string {
	return mn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mn *mn_Mong_MN) MonthWide(month time.Month) string {
	if len(mn.monthsWide) == 0 {
		return ""
	}
	return mn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mn *mn_Mong_MN) MonthsWide() []string {
	return mn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mn *mn_Mong_MN) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(mn.daysAbbreviated) == 0 {
		return ""
	}
	return mn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mn *mn_Mong_MN) WeekdaysAbbreviated() []string {
	return mn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mn *mn_Mong_MN) WeekdayNarrow(weekday time.Weekday) string {
	if len(mn.daysNarrow) == 0 {
		return ""
	}
	return mn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mn *mn_Mong_MN) WeekdaysNarrow() []string {
	return mn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mn *mn_Mong_MN) WeekdayShort(weekday time.Weekday) string {
	if len(mn.daysShort) == 0 {
		return ""
	}
	return mn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mn *mn_Mong_MN) WeekdaysShort() []string {
	return mn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mn *mn_Mong_MN) WeekdayWide(weekday time.Weekday) string {
	if len(mn.daysWide) == 0 {
		return ""
	}
	return mn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mn *mn_Mong_MN) WeekdaysWide() []string {
	return mn.daysWide
}

// Decimal returns the decimal point of number
func (mn *mn_Mong_MN) Decimal() string {
	return mn.decimal
}

// Group returns the group of number
func (mn *mn_Mong_MN) Group() string {
	return mn.group
}

// Group returns the minus sign of number
func (mn *mn_Mong_MN) Minus() string {
	return mn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mn_Mong_MN' and handles both Whole and Real numbers based on 'v'
func (mn *mn_Mong_MN) FmtNumber(num float64, v uint64) string {

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

// FmtPercent returns 'num' with digits/precision of 'v' for 'mn_Mong_MN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mn *mn_Mong_MN) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mn.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mn_Mong_MN'
// in accounting notation.
func (mn *mn_Mong_MN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mn.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
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

		b = append(b, mn.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
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

// FmtDateShort returns the short date representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xe1, 0xa0, 0x8b, 0xe1, 0xa0, 0xa3, 0xe1, 0xa0, 0xa8, 0x20, 0xe1, 0xa0, 0xa4}...)
	b = append(b, mn.monthsWide[t.Month()]...)
	b = append(b, []byte{0xe1, 0xa0, 0x8e, 0xe1, 0xa0, 0x8e, 0x20, 0xe1, 0xa0, 0xa4, 0xe1, 0xa0, 0xa9, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xe1, 0xa0, 0xa3, 0xe1, 0xa0, 0xa8, 0x20, 0xe1, 0xa0, 0x8e, 0xe1, 0xa0, 0x8e, 0xe1, 0xa0, 0x8e, 0xe1, 0xa0, 0xa4, 0x20}...)
	b = append(b, mn.monthsWide[t.Month()]...)
	b = append(b, []byte{0xe1, 0xa0, 0x8e, 0xe1, 0xa0, 0x8e, 0xe1, 0xa0, 0xa2, 0xe1, 0xa0, 0xa2, 0xe1, 0xa0, 0xa8}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, mn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0xe1, 0xa0, 0x8b, 0xe1, 0xa0, 0xad, 0xe1, 0xa0, 0xa0, 0xe1, 0xa0, 0xb7, 0xe1, 0xa0, 0xa0, 0xe1, 0xa0, 0xad}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtTimeShort(t time.Time) string {

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

// FmtTimeMedium returns the medium time representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtTimeMedium(t time.Time) string {

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

// FmtTimeLong returns the long time representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtTimeLong(t time.Time) string {

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

// FmtTimeFull returns the full time representation of 't' for 'mn_Mong_MN'
func (mn *mn_Mong_MN) FmtTimeFull(t time.Time) string {

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
