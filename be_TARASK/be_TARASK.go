package be_TARASK

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type be_TARASK struct {
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

// New returns a new instance of translator for the 'be_TARASK' locale
func New() locales.Translator {
	return &be_TARASK{
		locale:                 "be_TARASK",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{4, 6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "сту", "лют", "сак", "кра", "тра", "чэр", "ліп", "жні", "вер", "кас", "ліс", "сьн"},
		monthsNarrow:           []string{"", "с", "л", "с", "к", "т", "ч", "л", "ж", "в", "к", "л", "с"},
		monthsWide:             []string{"", "студзеня", "лютага", "сакавіка", "красавіка", "траўня", "чэрвеня", "ліпеня", "жніўня", "верасьня", "кастрычніка", "лістапада", "сьнежня"},
		daysAbbreviated:        []string{"няд", "пан", "аўт", "сер", "чац", "пят", "суб"},
		daysNarrow:             []string{"н", "п", "а", "с", "ч", "п", "с"},
		daysWide:               []string{"нядзеля", "панядзелак", "аўторак", "серада", "чацьвер", "пятніца", "субота"},
		timezones:              map[string]string{"ACDT": "Летні час цэнтральнай Аўстраліі", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Летні цэнтральна-заходні час Аўстраліі", "ACWST": "Змоўчны цэнтральна-заходні час Аўстраліі", "ADT": "Атлянтычны летні час", "ADT Arabia": "Летні час Саудаўскай Арабіі", "AEDT": "Летні час усходняй Аўстраліі", "AEST": "Змоўчны час усходняй Аўстраліі", "AFT": "Аўганістанскі час", "AKDT": "Летні час Аляскі", "AKST": "Змоўчны час Аляскі", "AMST": "Амазонскі летні час", "AMST Armenia": "Летні час Армэніі", "AMT": "Амазонскі змоўчны час", "AMT Armenia": "Змоўчны час Армэніі", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Аргентынскі летні час", "ART": "Аргентынскі змоўчны час", "AST": "Атлянтычны змоўчны час", "AST Arabia": "Змоўчны час Саудаўскай Арабіі", "AWDT": "Летні час заходняй Аўстраліі", "AWST": "змоўчны час заходняй Аўстраліі", "AZST": "Летні час Азэрбайджану", "AZT": "Змоўчны час Азэрбайджану", "BDT Bangladesh": "Летні час Банглядэшу", "BNT": "Час Брунэю", "BOT": "Балівійскі час", "BRST": "Бразыльскі летні час", "BRT": "Бразыльскі змоўчны час", "BST Bangladesh": "Змоўчны час Банглядэшу", "BT": "Час Бутану", "CAST": "CAST", "CAT": "Цэнтральнаафрыканскі час", "CCT": "Час Какосавых астравоў", "CDT": "Паўночнаамэрыканскі цэнтральны летні час", "CHADT": "Летні час Чатэму", "CHAST": "Змоўчны час Чатэму", "CHUT": "Час Трука", "CKT": "Змоўчны час астравоў Кука", "CKT DST": "Паўлетні час астравоў Кука", "CLST": "Чылійскі летні час", "CLT": "Чылійскі змоўчны час", "COST": "Калюмбійскі летні час", "COT": "Калюмбійскі змоўчны час", "CST": "Паўночнаамэрыканскі цэнтральны змоўчны час", "CST China": "Змоўчны час Кітаю", "CST China DST": "Летні час Кітаю", "CVST": "Летні час Каба-Вэрдэ", "CVT": "Змоўчны час Каба-Вэрдэ", "CXT": "Час вострава Раства", "ChST": "Час Чамора", "ChST NMI": "ChST NMI", "CuDT": "Летні час Кубы", "CuST": "Змоўчны час Кубы", "DAVT": "Час станцыі Дэйвіс", "DDUT": "Час станцыі Дзюмон-Дзюрвіль", "EASST": "Летні час вострава Раства", "EAST": "Змоўчны час вострава Раства", "EAT": "Усходнеафрыканскі час", "ECT": "Эквадорскі час", "EDT": "Паўночнаамэрыканскі ўсходні летні час", "EGDT": "Летні час Усходняй Грэнляндыі", "EGST": "Змоўчны час Усходняй Грэнляндыі", "EST": "Паўночнаамэрыканскі ўсходні змоўчны час", "FEET": "Далёкаўсходнеэўрапейскі час", "FJT": "Змоўчны час Фіджы", "FJT Summer": "Летні час Фіджы", "FKST": "Летні час Фолклэндзкіх астравоў", "FKT": "Змоўчны час Фолклэндзкіх астравоў", "FNST": "Летні час Фэрнанду-ды-Наронья", "FNT": "Змоўчны час Фэрнанду-ды-Наронья", "GALT": "Змоўчны час Галяпагоскіх астравоў", "GAMT": "Час астравоў Гамб’е", "GEST": "Грузінскі летні час", "GET": "Грузінскі змоўчны час", "GFT": "Час Францускай Гвіяны", "GIT": "Час астравоў Гілбэрта", "GMT": "Час па Грынвічы", "GNSST": "GNSST", "GNST": "GNST", "GST": "Час Паўднёвай Георгіі", "GST Guam": "GST Guam", "GYT": "Час Гаяны", "HADT": "Гавайска-Алевуцкі летні час", "HAST": "Гавайска-Алэвуцкі змоўчны час", "HKST": "Летні час Ганконгу", "HKT": "Змоўчны час Ганконга", "HOVST": "Летні час Хоўда", "HOVT": "Змоўчны час Хоўда", "ICT": "Індакітайскі час", "IDT": "Ізраільскі летні час", "IOT": "Час Індыйскага акіяну", "IRKST": "Іркуцкі летні час", "IRKT": "Іркуцкі змоўчны час", "IRST": "Іранскі змоўчны час", "IRST DST": "Іранскі летні час", "IST": "Час Індыі", "IST Israel": "Ізраільскі змоўчны час", "JDT": "Летні час Японіі", "JST": "Змоўчны час Японіі", "KOST": "Час астравоў Касраэ", "KRAST": "Краснаярскі летні час", "KRAT": "Краснаярскі змоўчны час", "KST": "Змоўчны час Карэі", "KST DST": "Летні час Карэі", "LHDT": "Летні час Лорд-Гаў", "LHST": "Змоўчны час Лорд-Гаў", "LINT": "Час астравоў Лайн", "MAGST": "Магаданскі летні час", "MAGT": "Магаданскі змоўчны час", "MART": "Час Маркіскіх астравоў", "MAWT": "Час станцыі Мосан", "MDT": "MDT", "MESZ": "Цэнтральнаэўрапейскі летні час", "MEZ": "Цэнтральнаэўрапейскі змоўчны час", "MHT": "Час Маршалавых астравоў", "MMT": "Час М’янмы", "MSD": "Маскоўскі летні час", "MST": "MST", "MUST": "Летні час Маўрыцыю", "MUT": "Змоўчны час Маўрыцыю", "MVT": "Час Мальдываў", "MYT": "Час Малайзіі", "NCT": "Змоўчны час Новай Каледоніі", "NDT": "Ньюфаўндлэндзкі летні час", "NDT New Caledonia": "Летні час Новай Каледоніі", "NFDT": "Летні час вострава Норфалк", "NFT": "Змоўчны час вострава Норфалк", "NOVST": "Новасыбірскі летні час", "NOVT": "Новасыбірскі змоўчны час", "NPT": "Нэпальскі час", "NRT": "Час Наўру", "NST": "Ньюфаўндлэндзкі змоўчны час", "NUT": "Час Ніуэ", "NZDT": "Летні час Новай Зэляндыі", "NZST": "Змоўчны час Новай Зэляндыі", "OESZ": "Усходнеэўрапейскі летні час", "OEZ": "Усходнеэўрапейскі змоўчны час", "OMSST": "Омскі летні час", "OMST": "Омскі змоўчны час", "PDT": "Ціхаакіянскі летні час", "PDTM": "Мэксыканскі ціхаакіянскі летні час", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Час Папуа-Новай Гвінэі", "PHOT": "Час астравоў Фінікс", "PKT": "Пакістанскі змоўчны час", "PKT DST": "Пакістанскі летні час", "PMDT": "Змоўчны летні час Сэн-П’еру і Мікелёну", "PMST": "Змоўчны час Сэн-П’еру і Мікелёну", "PONT": "Час вострава Панпэі", "PST": "Ціхаакіянскі змоўчны час", "PST Philippine": "Філіпінскі змоўчны час", "PST Philippine DST": "Філіпінскі летні час", "PST Pitcairn": "Час вострава Піткэрн", "PSTM": "Мэксыканскі ціхаакіянскі стандатны час", "PWT": "Час Палаў", "PYST": "Летні час Парагваю", "PYT": "Змоўчны час Парагваю", "PYT Korea": "Пхэньянскі час", "RET": "Час Рэюньёна", "ROTT": "Час станцыі Ротэра", "SAKST": "Сахалінскі летні час", "SAKT": "Сахалінскі змоўчны час", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Паўднёваафрыканскі час", "SBT": "Час Саламонавых астравоў", "SCT": "Час Сэйшэлаў", "SGT": "Сынгапурскі час", "SLST": "SLST", "SRT": "Час Сурынаму", "SST Samoa": "Змоўчны час Самоа", "SST Samoa Apia": "Змоўчны час Апіі", "SST Samoa Apia DST": "Летні час Апіі", "SST Samoa DST": "Летні час Самоа", "SYOT": "Час станцыі Сёва", "TAAF": "Час Францускай паўднёва-антарктычнай тэрыторыі", "TAHT": "Час Таіці", "TJT": "Час Таджыкістану", "TKT": "Час Такелаў", "TLT": "Час Усходняга Тымору", "TMST": "Летні час Туркмэністану", "TMT": "Змоўчны час Туркмэністана", "TOST": "Летні час Тонга", "TOT": "Змоўчны час Тонга", "TVT": "Час Тувалу", "TWT": "Змоўчны час Тайбэя", "TWT DST": "Летні час Тайбэю", "ULAST": "Летні час Улан-Батару", "ULAT": "Змоўчны час Улан-Батару", "UYST": "Уругвайскі летні час", "UYT": "Уругвайскі змоўчны час", "UZT": "Змоўчны час Узбэкістану", "UZT DST": "Летні час Узбэкістану", "VET": "Вэнэсуэльскі час", "VLAST": "Уладзівастоцкі летні час", "VLAT": "Уладзівастоцкі змоўчны час", "VOLST": "Валгаградзкі летні час", "VOLT": "Валгаградзкі змоўчны час", "VOST": "Час станцыі Ўсход", "VUT": "Змоўчны час Вануату", "VUT DST": "Летні час Вануату", "WAKT": "Час вострава Ўэйк", "WARST": "Летні час Заходняй Аргэнтыны", "WART": "Змоўчны час Заходняй Аргэнтыны", "WAST": "Заходнеафрыканскі час", "WAT": "Заходнеафрыканскі час", "WESZ": "Заходнеэўрапейскі летні час", "WEZ": "Заходнеэўрапейскі змоўчны час", "WFT": "Час астравоў Ўоліс і Футуна", "WGST": "Летні час Заходняй Грэнляндыі", "WGT": "Змоўчны час Заходняй Грэнляндыі", "WIB": "Заходнеінданэзійскі час", "WIT": "Усходнеінданэзійскі час", "WITA": "Цэнтральнаінданэзійскі час", "YAKST": "Якуцкі летні час", "YAKT": "Якуцкі змоўчны час", "YEKST": "Екацярынбурскі летні час", "YEKT": "Екацярынбурскі змоўчны час", "YST": "Час Юкана", "МСК": "Маскоўскі змоўчны час", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Заходнеказахстанскі час", "شىعىش قازاق ەلى": "Усходнеказахстанскі час", "قازاق ەلى": "Казахстанскі час", "قىرعىزستان": "Час Кыргыстану", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Летні час Азорскіх астравоў"},
	}
}

// Locale returns the current translators string locale
func (be *be_TARASK) Locale() string {
	return be.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'be_TARASK'
func (be *be_TARASK) PluralsCardinal() []locales.PluralRule {
	return be.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'be_TARASK'
func (be *be_TARASK) PluralsOrdinal() []locales.PluralRule {
	return be.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'be_TARASK'
func (be *be_TARASK) PluralsRange() []locales.PluralRule {
	return be.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'be_TARASK'
func (be *be_TARASK) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'be_TARASK'
func (be *be_TARASK) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if (nMod10 == 2 || nMod10 == 3) && (nMod100 != 12 && nMod100 != 13) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'be_TARASK'
func (be *be_TARASK) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
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
func (be *be_TARASK) MonthAbbreviated(month time.Month) string {
	return be.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (be *be_TARASK) MonthsAbbreviated() []string {
	return be.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (be *be_TARASK) MonthNarrow(month time.Month) string {
	return be.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (be *be_TARASK) MonthsNarrow() []string {
	return be.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (be *be_TARASK) MonthWide(month time.Month) string {
	return be.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (be *be_TARASK) MonthsWide() []string {
	return be.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (be *be_TARASK) WeekdayAbbreviated(weekday time.Weekday) string {
	return be.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (be *be_TARASK) WeekdaysAbbreviated() []string {
	return be.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (be *be_TARASK) WeekdayNarrow(weekday time.Weekday) string {
	return be.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (be *be_TARASK) WeekdaysNarrow() []string {
	return be.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (be *be_TARASK) WeekdayShort(weekday time.Weekday) string {
	return be.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (be *be_TARASK) WeekdaysShort() []string {
	return be.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (be *be_TARASK) WeekdayWide(weekday time.Weekday) string {
	return be.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (be *be_TARASK) WeekdaysWide() []string {
	return be.daysWide
}

// Decimal returns the decimal point of number
func (be *be_TARASK) Decimal() string {
	return be.decimal
}

// Group returns the group of number
func (be *be_TARASK) Group() string {
	return be.group
}

// Group returns the minus sign of number
func (be *be_TARASK) Minus() string {
	return be.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'be_TARASK' and handles both Whole and Real numbers based on 'v'
func (be *be_TARASK) FmtNumber(num float64, v uint64) string {
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

// FmtPercent returns 'num' with digits/precision of 'v' for 'be_TARASK' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (be *be_TARASK) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'be_TARASK'
func (be *be_TARASK) FmtCurrency(num float64, v uint64, currency currency.Type) string {
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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'be_TARASK'
// in accounting notation.
func (be *be_TARASK) FmtAccounting(num float64, v uint64, currency currency.Type) string {
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

// FmtDateShort returns the short date representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtDateMedium(t time.Time) string {
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

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtDateLong(t time.Time) string {
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

// FmtDateFull returns the full date representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtDateFull(t time.Time) string {
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

// FmtTimeShort returns the short time representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtTimeShort(t time.Time) string {
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

// FmtTimeMedium returns the medium time representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtTimeMedium(t time.Time) string {
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

// FmtTimeLong returns the long time representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtTimeLong(t time.Time) string {
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

// FmtTimeFull returns the full time representation of 't' for 'be_TARASK'
func (be *be_TARASK) FmtTimeFull(t time.Time) string {
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
