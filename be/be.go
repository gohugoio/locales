package be

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type be struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
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

// New returns a new instance of translator for the 'be' locale
func New() locales.Translator {
	return &be{
		locale:                 "be",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{4, 6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "A$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "Bds$", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BD$", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "B$", "BTN", "BUK", "BWP", "BYB", "Br", "BYR", "BZ$", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC$", "$MN", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "RD$", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJ$", "FK£", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "G$", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "Íkr", "ITL", "J$", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "CI$", "KZT", "LAK", "LBP", "LKR", "L$", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "N$", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SI$", "SCR", "SDD", "SDG", "SDP", "SEK", "S$", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TT$", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "$U", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "сту", "лют", "сак", "кра", "мая", "чэр", "ліп", "жні", "вер", "кас", "ліс", "сне"},
		monthsNarrow:           []string{"", "с", "л", "с", "к", "м", "ч", "л", "ж", "в", "к", "л", "с"},
		monthsWide:             []string{"", "студзеня", "лютага", "сакавіка", "красавіка", "мая", "чэрвеня", "ліпеня", "жніўня", "верасня", "кастрычніка", "лістапада", "снежня"},
		daysAbbreviated:        []string{"нд", "пн", "аў", "ср", "чц", "пт", "сб"},
		daysNarrow:             []string{"н", "п", "а", "с", "ч", "п", "с"},
		daysWide:               []string{"нядзеля", "панядзелак", "аўторак", "серада", "чацвер", "пятніца", "субота"},
		periodsNarrow:          []string{"am", "pm"},
		erasAbbreviated:        []string{"да н.э.", "н.э."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"да нараджэння Хрыстова", "ад нараджэння Хрыстова"},
		timezones:              map[string]string{"ACDT": "Летні час цэнтральнай Аўстраліі", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Летні цэнтральна-заходні час Аўстраліі", "ACWST": "Стандартны цэнтральна-заходні час Аўстраліі", "ADT": "Атлантычны летні час", "ADT Arabia": "Летні час Саудаўскай Аравіі", "AEDT": "Летні час усходняй Аўстраліі", "AEST": "Стандартны час усходняй Аўстраліі", "AFT": "Афганістанскі час", "AKDT": "Летні час Аляскі", "AKST": "Стандартны час Аляскі", "AMST": "Амазонскі летні час", "AMST Armenia": "Летні час Арменіі", "AMT": "Амазонскі стандартны час", "AMT Armenia": "Стандартны час Арменіі", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргенцінскі летні час", "ART": "Аргенцінскі стандартны час", "AST": "Атлантычны стандартны час", "AST Arabia": "Стандартны час Саудаўскай Аравіі", "AWDT": "Летні час заходняй Аўстраліі", "AWST": "Стандартны час заходняй Аўстраліі", "AZST": "Летні час Азербайджана", "AZT": "Стандартны час Азербайджана", "BDT Bangladesh": "Летні час Бангладэш", "BNT": "Час Брунея", "BOT": "Балівійскі час", "BRST": "Бразільскі летні час", "BRT": "Бразільскі стандартны час", "BST Bangladesh": "Стандартны час Бангладэш", "BT": "Час Бутана", "CAST": "CAST", "CAT": "Цэнтральнаафрыканскі час", "CCT": "Час Какосавых астравоў", "CDT": "Паўночнаамерыканскі цэнтральны летні час", "CHADT": "Летні час Чатэма", "CHAST": "Стандартны час Чатэма", "CHUT": "Час Трука", "CKT": "Стандартны час астравоў Кука", "CKT DST": "Паўлетні час астравоў Кука", "CLST": "Чылійскі летні час", "CLT": "Чылійскі стандартны час", "COST": "Калумбійскі летні час", "COT": "Калумбійскі стандартны час", "CST": "Паўночнаамерыканскі цэнтральны стандартны час", "CST China": "Стандартны час Кітая", "CST China DST": "Летні час Кітая", "CVST": "Летні час Каба-Вердэ", "CVT": "Стандартны час Каба-Вердэ", "CXT": "Час вострава Каляд", "ChST": "Час Чамора", "ChST NMI": "ChST NMI", "CuDT": "Летні час Кубы", "CuST": "Стандартны час Кубы", "DAVT": "Час станцыі Дэйвіс", "DDUT": "Час станцыі Дзюмон-Дзюрвіль", "EASST": "Летні час вострава Вялікадня", "EAST": "Стандартны час вострава Вялікадня", "EAT": "Усходнеафрыканскі час", "ECT": "Эквадорскі час", "EDT": "Паўночнаамерыканскі ўсходні летні час", "EGDT": "Летні час Усходняй Грэнландыі", "EGST": "Стандартны час Усходняй Грэнландыі", "EST": "Паўночнаамерыканскі ўсходні стандартны час", "FEET": "Далёкаўсходнееўрапейскі час", "FJT": "Стандартны час Фіджы", "FJT Summer": "Летні час Фіджы", "FKST": "Летні час Фалклендскіх астравоў", "FKT": "Стандартны час Фалклендскіх астравоў", "FNST": "Летні час Фернанду-дзі-Наронья", "FNT": "Стандартны час Фернанду-дзі-Наронья", "GALT": "Стандартны час Галапагоскіх астравоў", "GAMT": "Час астравоў Гамб’е", "GEST": "Грузінскі летні час", "GET": "Грузінскі стандартны час", "GFT": "Час Французскай Гвіяны", "GIT": "Час астравоў Гілберта", "GMT": "Час па Грынвічы", "GNSST": "GNSST", "GNST": "GNST", "GST": "Час Персідскага заліва", "GST Guam": "GST Guam", "GYT": "Час Гаяны", "HADT": "Гавайска-Алеуцкі стандартны час", "HAST": "Гавайска-Алеуцкі стандартны час", "HKST": "Летні час Ганконга", "HKT": "Стандартны час Ганконга", "HOVST": "Летні час Хоўда", "HOVT": "Стандартны час Хоўда", "ICT": "Індакітайскі час", "IDT": "Ізраільскі летні час", "IOT": "Час Індыйскага акіяна", "IRKST": "Іркуцкі летні час", "IRKT": "Іркуцкі стандартны час", "IRST": "Іранскі стандартны час", "IRST DST": "Іранскі летні час", "IST": "Час Індыі", "IST Israel": "Ізраільскі стандартны час", "JDT": "Летні час Японіі", "JST": "Стандартны час Японіі", "KOST": "Час астравоў Кусаіе", "KRAST": "Краснаярскі летні час", "KRAT": "Краснаярскі стандартны час", "KST": "Стандартны час Карэі", "KST DST": "Летні час Карэі", "LHDT": "Летні час Лорд-Хау", "LHST": "Стандартны час Лорд-Хау", "LINT": "Час астравоў Лайн", "MAGST": "Магаданскі летні час", "MAGT": "Магаданскі стандартны час", "MART": "Час Маркізскіх астравоў", "MAWT": "Час станцыі Моўсан", "MDT": "MDT", "MESZ": "Цэнтральнаеўрапейскі летні час", "MEZ": "Цэнтральнаеўрапейскі стандартны час", "MHT": "Час Маршалавых астравоў", "MMT": "Час М’янмы", "MSD": "Маскоўскі летні час", "MST": "MST", "MUST": "Летні час Маўрыкія", "MUT": "Стандартны час Маўрыкія", "MVT": "Час Мальдываў", "MYT": "Час Малайзіі", "NCT": "Стандартны час Новай Каледоніі", "NDT": "Ньюфаўндлендскі летні час", "NDT New Caledonia": "Летні час Новай Каледоніі", "NFDT": "Летні час вострава Норфалк", "NFT": "Стандартны час вострава Норфалк", "NOVST": "Новасібірскі летні час", "NOVT": "Новасібірскі стандартны час", "NPT": "Непальскі час", "NRT": "Час Науру", "NST": "Ньюфаўндлендскі стандартны час", "NUT": "Час Ніуэ", "NZDT": "Летні час Новай Зеландыі", "NZST": "Стандартны час Новай Зеландыі", "OESZ": "Усходнееўрапейскі летні час", "OEZ": "Усходнееўрапейскі стандартны час", "OMSST": "Омскі летні час", "OMST": "Омскі стандартны час", "PDT": "Ціхаакіянскі летні час", "PDTM": "Мексіканскі ціхаакіянскі летні час", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Час Папуа-Новай Гвінеі", "PHOT": "Час астравоў Фенікс", "PKT": "Пакістанскі стандартны час", "PKT DST": "Пакістанскі летні час", "PMDT": "Стандартны летні час Сен-П’ер і Мікелон", "PMST": "Стандартны час Сен-П’ер і Мікелон", "PONT": "Час вострава Панпеі", "PST": "Ціхаакіянскі стандартны час", "PST Philippine": "Філіпінскі стандартны час", "PST Philippine DST": "Філіпінскі летні час", "PST Pitcairn": "Час вострава Піткэрн", "PSTM": "Мексіканскі ціхаакіянскі стандатны час", "PWT": "Час Палау", "PYST": "Летні час Парагвая", "PYT": "Стандартны час Парагвая", "PYT Korea": "Пхеньянскі час", "RET": "Час Рэюньёна", "ROTT": "Час станцыі Ротэра", "SAKST": "Сахалінскі летні час", "SAKT": "Сахалінскі стандартны час", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Паўднёваафрыканскі час", "SBT": "Час Саламонавых астравоў", "SCT": "Час Сейшэльскіх астравоў", "SGT": "Сінгапурскі час", "SLST": "SLST", "SRT": "Час Сурынама", "SST Samoa": "Стандартны час Самоа", "SST Samoa Apia": "Стандартны час Апіі", "SST Samoa Apia DST": "Летні час Апіі", "SST Samoa DST": "Летні час Самоа", "SYOT": "Час станцыі Сёва", "TAAF": "Час Французскай паўднёва-антарктычнай тэрыторыі", "TAHT": "Час Таіці", "TJT": "Час Таджыкістана", "TKT": "Час Такелау", "TLT": "Час Усходняга Тымора", "TMST": "Летні час Туркменістана", "TMT": "Стандартны час Туркменістана", "TOST": "Летні час Тонга", "TOT": "Стандартны час Тонга", "TVT": "Час Тувалу", "TWT": "Стандартны час Тайбэя", "TWT DST": "Летні час Тайбэя", "ULAST": "Летні час Улан-Батара", "ULAT": "Стандартны час Улан-Батара", "UYST": "Уругвайскі летні час", "UYT": "Уругвайскі стандартны час", "UZT": "Стандартны час Узбекістана", "UZT DST": "Летні час Узбекістана", "VET": "Венесуэльскі час", "VLAST": "Уладзівастоцкі летні час", "VLAT": "Уладзівастоцкі стандартны час", "VOLST": "Валгаградскі летні час", "VOLT": "Валгаградскі стандартны час", "VOST": "Час станцыі Васток", "VUT": "Стандартны час Вануату", "VUT DST": "Летні час Вануату", "WAKT": "Час вострава Уэйк", "WARST": "Летні час Заходняй Аргенціны", "WART": "Стандартны час Заходняй Аргенціны", "WAST": "Заходнеафрыканскі час", "WAT": "Заходнеафрыканскі час", "WESZ": "Заходнееўрапейскі летні час", "WEZ": "Заходнееўрапейскі стандартны час", "WFT": "Час астравоў Уоліс і Футуна", "WGST": "Летні час Заходняй Грэнландыі", "WGT": "Стандартны час Заходняй Грэнландыі", "WIB": "Заходнеінданезійскі час", "WIT": "Усходнеінданезійскі час", "WITA": "Цэнтральнаінданезійскі час", "YAKST": "Якуцкі летні час", "YAKT": "Якуцкі стандартны час", "YEKST": "Екацярынбургскі летні час", "YEKT": "Екацярынбургскі стандартны час", "YST": "Час Юкана", "МСК": "Маскоўскі стандартны час", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Заходнеказахстанскі час", "شىعىش قازاق ەلى": "Усходнеказахстанскі час", "قازاق ەلى": "Казахстанскі час", "قىرعىزستان": "Час Кыргызстана", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Перуанскі летні час"},
	}
}

// Locale returns the current translators string locale
func (be *be) Locale() string {
	return be.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'be'
func (be *be) PluralsCardinal() []locales.PluralRule {
	return be.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'be'
func (be *be) PluralsOrdinal() []locales.PluralRule {
	return be.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'be'
func (be *be) PluralsRange() []locales.PluralRule {
	return be.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'be'
func (be *be) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 1 && nMod100 != 11 {
		return locales.PluralRuleOne
	} else if nMod10 >= 2 && nMod10 <= 4 && (nMod100 < 12 || nMod100 > 14) {
		return locales.PluralRuleFew
	} else if (nMod10 == 0) || (nMod10 >= 5 && nMod10 <= 9) || (nMod100 >= 11 && nMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'be'
func (be *be) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if (nMod10 == 2 || nMod10 == 3) && (nMod100 != 12 && nMod100 != 13) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'be'
func (be *be) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := be.CardinalPluralRule(num1, v1)
	end := be.CardinalPluralRule(num2, v2)

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
func (be *be) MonthAbbreviated(month time.Month) string {
	return be.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (be *be) MonthsAbbreviated() []string {
	return be.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (be *be) MonthNarrow(month time.Month) string {
	return be.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (be *be) MonthsNarrow() []string {
	return be.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (be *be) MonthWide(month time.Month) string {
	return be.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (be *be) MonthsWide() []string {
	return be.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (be *be) WeekdayAbbreviated(weekday time.Weekday) string {
	return be.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (be *be) WeekdaysAbbreviated() []string {
	return be.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (be *be) WeekdayNarrow(weekday time.Weekday) string {
	return be.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (be *be) WeekdaysNarrow() []string {
	return be.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (be *be) WeekdayShort(weekday time.Weekday) string {
	return be.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (be *be) WeekdaysShort() []string {
	return be.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (be *be) WeekdayWide(weekday time.Weekday) string {
	return be.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (be *be) WeekdaysWide() []string {
	return be.daysWide
}

// Decimal returns the decimal point of number
func (be *be) Decimal() string {
	return be.decimal
}

// Group returns the group of number
func (be *be) Group() string {
	return be.group
}

// Group returns the minus sign of number
func (be *be) Minus() string {
	return be.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'be' and handles both Whole and Real numbers based on 'v'
func (be *be) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(be.group) - 1; j >= 0; j-- {
					b = append(b, be.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'be' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (be *be) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, be.percentSuffix...)

	b = append(b, be.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'be'
func (be *be) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := be.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(be.group) - 1; j >= 0; j-- {
					b = append(b, be.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, be.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, be.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'be'
// in accounting notation.
func (be *be) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := be.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(be.group) - 1; j >= 0; j-- {
					b = append(b, be.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, be.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, be.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, be.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'be'
func (be *be) FmtDateShort(t time.Time) string {
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

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'be'
func (be *be) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, be.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'be'
func (be *be) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, be.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'be'
func (be *be) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, be.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, be.monthsWide[t.Month()]...)
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

// FmtTimeShort returns the short time representation of 't' for 'be'
func (be *be) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, be.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'be'
func (be *be) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, be.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, be.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'be'
func (be *be) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, be.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, be.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'be'
func (be *be) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, be.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, be.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := be.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
