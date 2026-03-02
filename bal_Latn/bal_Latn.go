package bal_Latn

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type bal_Latn struct {
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
	currencyPositivePrefix string
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'bal_Latn' locale
func New() locales.Translator {
	return &bal_Latn{
		locale:                 "bal_Latn",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AWGA", "ALK", "ALBL", "ARMD", "NLAG", "ANGK", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARJP", "ATS", "AUD", "ARBF", "AZM", "AZRM", "BAD", "BHBM", "BAN", "BRBD", "BGDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BLGL", "BGO", "BHD", "BRDF", "BRMD", "BRND", "BLWB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BHMD", "BTNN", "BUK", "BTSP", "BYB", "BLRR", "BYR", "BLZD", "KN$", "KNGF", "CHE", "SWZF", "CHW", "CLE", "CLF", "CLP", "CNYÁ", "CNX", "CHNY", "KLBP", "COU", "KRK", "CSD", "CSK", "KBKP", "CBAP", "KWRE", "CYP", "CHKK", "DDM", "DEM", "DJBF", "DNMK", "DOMP", "ALJD", "ECS", "ECV", "EEK", "MSRP", "ERTN", "ESA", "ESB", "ESP", "ETPB", "EUR", "FIM", "FJD", "FKP", "FRF", "BRTP", "GEK", "JRJL", "GHC", "GNAS", "GBRP", "GMBD", "GWNF", "GNS", "GQE", "GRD", "GTMK", "GWE", "GWP", "GYD", "HGKD", "HNDL", "HRD", "KRSK", "HTNG", "HGRF", "ENDR", "IEP", "ILP", "ILR", "ILS", "HNDR", "IQD", "ERNR", "ISJ", "ISLK", "ITL", "JMKD", "JOD", "¥", "KINS", "KGSS", "KMBR", "KMRF", "SHKW", "KRH", "KRO", "ZBKW", "KWD", "KMID", "KZKT", "LTKK", "LBNP", "SRLR", "LBRD", "LSTL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LBYD", "MRKD", "MAF", "MCF", "MDC", "MLDL", "MLGA", "MGF", "MKDD", "MKN", "MLF", "MNMK", "MNGT", "MKNP", "MRO", "MRTU", "MTL", "MTP", "MURR", "MVP", "MLDR", "MLWK", "MKS$", "MXP", "MXV", "MLRG", "MZE", "MZM", "MZBM", "NMBD", "NJRN", "NIC", "NKGC", "NLG", "NRWK", "NPLR", "NZD", "OMR", "PNMB", "PEI", "PRSL", "PES", "PGK", "PLPP", "PKRS", "PLNZ", "PLZ", "PTE", "PRGG", "QAR", "RHD", "ROL", "RMNL", "SRBD", "RUSR", "RUR", "RWDF", "SAR", "SBD", "SCLR", "SDD", "SDNP", "SDP", "SWDK", "SGPD", "SHLP", "SIT", "SKK", "SLNL", "SRLL", "SÓMS", "SRND", "SRG", "ZRSP", "STD", "STPD", "SUR", "SVC", "SURP", "SWZL", "TÁIB", "TJR", "TJS", "TMM", "TMT", "TNSD", "TOP", "TPE", "TRL", "TRKL", "TDTD", "NTWD", "TNZS", "YKNH", "UAK", "UGS", "YUGS", "$", "USN", "USS", "UYI", "UYP", "YRGP", "UYW", "OZBS", "VEB", "VED", "VEF", "WNZB", "WTND", "VNN", "VUV", "WST", "DARF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "RKB$", "XCG", "XDR", "XEU", "XFO", "XFU", "RACF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAPR", "ZMK", "ZMBK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "H",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "H",
		monthsAbbreviated:      []string{"", "Jan", "Par", "Már", "Apr", "Mai", "Jun", "Jól", "Aga", "Sat", "Akt", "Naw", "Das"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "Janwari", "Parwari", "Márch", "Aprél", "Mai", "Jun", "Jólái", "Agast", "Satambar", "Aktubar", "Nawambar", "Dasambar"},
		daysAbbreviated:        []string{"Yak", "Do", "Say", "Chá", "Pan", "Jom", "Sha"},
		daysNarrow:             []string{"Y", "D", "S", "Ch", "P", "J", "Sh"},
		daysWide:               []string{"Yakshambeh", "Doshambeh", "Sayshambeh", "Chárshambeh", "Panchshambeh", "Jomah", "Shambeh"},
		periodsAbbreviated:     []string{"am", "pm"},
		timezones:              map[string]string{"ACDT": "Delgáhi Ástréliáay garmági wahd", "ACST": "Delgáhi Ástréliáay anjári wahd", "ACT": "ACT", "ACWDT": "Delgáhirónendi Ástréliáay garmági wahd", "ACWST": "Delgáhirónendi Ástréliáay anjári wahd", "ADT": "Atlantáay róchi wahd", "ADT Arabia": "Arabi róchi wahd", "AEDT": "Ródarátki Ástréliáay garmági wahd", "AEST": "Ródarátki Ástréliáay anjári wahd", "AFT": "Awgánestánay wahd", "AKDT": "Aláskáay garmági wahd", "AKST": "Aláskáay anjári wahd", "AMST": "Amázónay garmági wahd", "AMST Armenia": "Árminiáay garmági wahd", "AMT": "Amázónay anjári wahd", "AMT Armenia": "Árminiáay anjári wahd", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Arjentináay garmági wahd", "ART": "Arjentináay anjári wahd", "AST": "Atlantáay anjári wahd", "AST Arabia": "Arabi anjári wahd", "AWDT": "Rónendi Ástréliáay garmági wahd", "AWST": "Rónendi Ástréliáay anjári wahd", "AZST": "Ázerbáijánay garmági wahd", "AZT": "Ázerbáijánay anjári wahd", "BDT Bangladesh": "Bangaladéshay garmági wahd", "BNT": "Brunáiay wahd", "BOT": "Boliwiáay wahd", "BRST": "Brázilay garmági wahd", "BRT": "Brázilay anjári wahd", "BST Bangladesh": "Bangaladéshay anjári wahd", "BT": "Buthánay wahd", "CAST": "CAST", "CAT": "Delgáhi Aprikáay wahd", "CCT": "Kukus Islánday wahd", "CDT": "Delgáhi Amrikáay garmági wahd", "CHADT": "Chatam róchi wahd", "CHAST": "Chatam anjári wahd", "CHUT": "Chukay wahd", "CKT": "Kuk Islánday anjári wahd", "CKT DST": "Kuk Islánday ném-garmági wahd", "CLST": "Chilayay garmági wahd", "CLT": "Chilayay anjári wahd", "COST": "Kolambiáay garmági wahd", "COT": "Kolambiáay anjári wahd", "CST": "Delgáhi Amrikáay anjári wahd", "CST China": "Chinay anjári wahd", "CST China DST": "Chinay róchi wahd", "CVST": "Kap Wardéay garmági wahd", "CVT": "Kap Wardéay anjári wahd", "CXT": "Kresmes Islánday wahd", "ChST": "Chamorróay wahd", "ChST NMI": "ChST NMI", "CuDT": "Kyubáay róchay wahd", "CuST": "Kyubáay anjári wahd", "DAVT": "Dawisay wahd", "DDUT": "Dumont Urwilay wahd", "EASST": "Isthar Islánday garmági wahd", "EAST": "Isthar Islánday anjári wahd", "EAT": "Ródarátki Aprikáay wahd", "ECT": "Ekwádóray wahd", "EDT": "Ródarátki Amrikáay garmági wahd", "EGDT": "Ródarátki Grinlánday garmági wahd", "EGST": "Ródarátki Grinlánday anjári wahd", "EST": "Ródarátki Amrikáay anjári wahd", "FEET": "Démterén Ródarátki Yuropay anjári wahd", "FJT": "Fijiay anjári wahd", "FJT Summer": "Fijiay garmági wahd", "FKST": "Palklánd Islánday garmági wahd", "FKT": "Palklánd Islánday anjári wahd", "FNST": "Noronáay garmági wahd", "FNT": "Noronáay anjári wahd", "GALT": "Galapagosay wahd", "GAMT": "Gambiray wahd", "GEST": "Járjiáay garmági wahd", "GET": "Járjiáay anjári wahd", "GFT": "Paránsi Gwináay wahd", "GIT": "Gelbart Islánday wahd", "GMT": "Grinwech Min Wahd", "GNSST": "GNSST", "GNST": "GNST", "GST": "Zerbári Járjiáay wahd", "GST Guam": "GST Guam", "GYT": "Guyánáay wahd", "HADT": "Hawái/Alushiay garmági wahd", "HAST": "Hawái/Alushiay anjári wahd", "HKST": "Háng Kángay garmági wahd", "HKT": "Háng Kángay anjári wahd", "HOVST": "Hówday garmági wahd", "HOVT": "Hówday anjári wahd", "ICT": "Hendóchinay wahd", "IDT": "Esráilay róchi wahd", "IOT": "Hendi zeray wahd", "IRKST": "Erkuskay garmági wahd", "IRKT": "Erkuskay anjári wahd", "IRST": "Éránay anjári wahd", "IRST DST": "Éránay róchi wahd", "IST": "Henday anjári wahd", "IST Israel": "Esráilay anjári wahd", "JDT": "Jápánay róchi wahd", "JST": "Jápánay anjári wahd", "KOST": "Kósraiay wahd", "KRAST": "Krasnóyáskay garmági wahd", "KRAT": "Krasnóyáskay anjári wahd", "KST": "Kóriáay anjári wahd", "KST DST": "Kóriáay róchi wahd", "LHDT": "Ástréliáay, Ládhaway garmági wahd", "LHST": "Ástréliáay, Ládhaway anjári wahd", "LINT": "Liné Islánday wahd", "MAGST": "Mágadánay garmági wahd", "MAGT": "Mágadánay anjári wahd", "MART": "Markésásay wahd", "MAWT": "Mawsonay wahd", "MDT": "MDT", "MESZ": "Delgáhi Yuropay garmági wahd", "MEZ": "Delgáhi Yuropay anjári wahd", "MHT": "Márshal Islánday wahd", "MMT": "Myanmáray wahd", "MSD": "Máskóay garmági wahd", "MST": "MST", "MUST": "Muritániáay garmági wahd", "MUT": "Muritániáay anjári wahd", "MVT": "Máldipay wahd", "MYT": "Malishiáay wahd", "NCT": "Nyu Kaledóniáay anjári wahd", "NDT": "Nipándlaynday garmági wahd", "NDT New Caledonia": "Nyu Kaledóniáay garmági wahd", "NFDT": "Nurpolk Islánday róchi wahd", "NFT": "Nurpolk Islánday anjári wahd", "NOVST": "Nawásibiskay garmági wahd", "NOVT": "Nawásibiskay anjári wahd", "NPT": "Népálay wahd", "NRT": "Nauruay wahd", "NST": "Nipándlaynday anjári wahd", "NUT": "Niuay wahd", "NZDT": "Niu Zilánday róchi wahd", "NZST": "Niu Zilánday anjári wahd", "OESZ": "Ródarátki Yuropay garmági wahd", "OEZ": "Ródarátki Yuropay anjári wahd", "OMSST": "Ómskay garmági wahd", "OMST": "Ómskay anjári wahd", "PDT": "Árámzeri Amrikáay garmági wahd", "PDTM": "Árámzeri Meksikóay garmági wahd", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Pápuá Niu Giniáay wahd", "PHOT": "Phoeneks Islánday wahd", "PKT": "Pákestánay anjári wahd", "PKT DST": "Pákestánay garmági wahd", "PMDT": "St. Péri o Mikwélin róchi wahd", "PMST": "St. Péri o Mikwélin ajári wahd", "PONT": "Pónpiay wahd", "PST": "Árámzeri Amrikáay anjári wahd", "PST Philippine": "Pelpinay anjári wahd", "PST Philippine DST": "Pelpinay garmági wahd", "PST Pitcairn": "Pitkarénay wahd", "PSTM": "Árámzeri Meksikóay anjári wahd", "PWT": "Paláuay wahd", "PYST": "Paragóayay garmági wahd", "PYT": "Paragóayay anjári wahd", "PYT Korea": "Pyongyángay wahd", "RET": "Réyunianay wahd", "ROTT": "Rothéráay wahd", "SAKST": "Sakhálinay garmági wahd", "SAKT": "Sakhálinay anjári wahd", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Ródarátki Aprikáay anjári wahd", "SBT": "Solomán Islánday wahd", "SCT": "Séchelesay wahd", "SGT": "Sengápuray anjári wahd", "SLST": "SLST", "SRT": "Surinaymay wahd", "SST Samoa": "Samóáway anjári wahd", "SST Samoa Apia": "Apiáay anjári wahd", "SST Samoa Apia DST": "Apiáay róchi wahd", "SST Samoa DST": "Samóáway róchi wahd", "SYOT": "Syówáay wahd", "TAAF": "Zerbári Paransi o Antárktikáay wahd", "TAHT": "Tahitiay wahd", "TJT": "Tájekestánay wahd", "TKT": "Tokeláuay wahd", "TLT": "Ródarátki Timuray wahd", "TMST": "Torkmenestánay garmági wahd", "TMT": "Torkmenestánay anjári wahd", "TOST": "Tongáay garmági wahd", "TOT": "Tongáay anjári wahd", "TVT": "Tuwáluay wahd", "TWT": "Táipiay anjári wahd", "TWT DST": "Táipiay róchi wahd", "ULAST": "Ulánbátaray garmági wahd", "ULAT": "Ulánbátaray anjári wahd", "UYST": "Yurógóayay garmági wahd", "UYT": "Yurógóayay anjári wahd", "UZT": "Ozbekestánay anjári wahd", "UZT DST": "Ozbekestánay garmági wahd", "VET": "Wenezwéláay wahd", "VLAST": "Waládiwástókay garmági wahd", "VLAT": "Waládiwástókay anjári wahd", "VOLST": "Wolgograday garmági wahd", "VOLT": "Wolgograday anjári wahd", "VOST": "Wostokay wahd", "VUT": "Wánuátuay anjári wahd", "VUT DST": "Wánuátuay garmági wahd", "WAKT": "Wayk Islánday wahd", "WARST": "Rónendi Arjentináay gramági wahd", "WART": "Rónendi Arjentináay anjári wahd", "WAST": "Rónendi Aprikáay wahd", "WAT": "Rónendi Aprikáay wahd", "WESZ": "Rónendi Yuropay garmági wahd", "WEZ": "Rónendi Yuropay anjári wahd", "WFT": "Wallis o Futunáay wahd", "WGST": "Rónendi Grinlánd Garmági Wahd", "WGT": "Rónendi Grinlánday anjári wahd", "WIB": "Rónendi Endhonishiáay anjári wahd", "WIT": "Ródarátki Endhonishiáay anjári wahd", "WITA": "Delgáhi Endhonishiáay anjári wahd", "YAKST": "Yákuskay garmági wahd", "YAKT": "Yákuskay anjári wahd", "YEKST": "Yakátrinborgay garmági wahd", "YEKT": "Yakátrinborgay anjári wahd", "YST": "Yukón wahd", "МСК": "Máskóay anjári wahd", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Rónendi Kázekestánay anjári wahd", "شىعىش قازاق ەلى": "Ródarátki Kázekestánay anjári wahd", "قازاق ەلى": "Kázakestánay wahd", "قىرعىزستان": "Kargezestánay wahd", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Azóresay garmági wahd"},
	}
}

// Locale returns the current translators string locale
func (bal *bal_Latn) Locale() string {
	return bal.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bal_Latn'
func (bal *bal_Latn) PluralsCardinal() []locales.PluralRule {
	return bal.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bal_Latn'
func (bal *bal_Latn) PluralsOrdinal() []locales.PluralRule {
	return bal.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bal_Latn'
func (bal *bal_Latn) PluralsRange() []locales.PluralRule {
	return bal.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bal_Latn'
func (bal *bal_Latn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bal_Latn'
func (bal *bal_Latn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bal_Latn'
func (bal *bal_Latn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bal *bal_Latn) MonthAbbreviated(month time.Month) string {
	if len(bal.monthsAbbreviated) == 0 {
		return ""
	}
	return bal.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bal *bal_Latn) MonthsAbbreviated() []string {
	return bal.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bal *bal_Latn) MonthNarrow(month time.Month) string {
	if len(bal.monthsNarrow) == 0 {
		return ""
	}
	return bal.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bal *bal_Latn) MonthsNarrow() []string {
	return bal.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bal *bal_Latn) MonthWide(month time.Month) string {
	if len(bal.monthsWide) == 0 {
		return ""
	}
	return bal.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bal *bal_Latn) MonthsWide() []string {
	return bal.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bal *bal_Latn) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(bal.daysAbbreviated) == 0 {
		return ""
	}
	return bal.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bal *bal_Latn) WeekdaysAbbreviated() []string {
	return bal.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bal *bal_Latn) WeekdayNarrow(weekday time.Weekday) string {
	if len(bal.daysNarrow) == 0 {
		return ""
	}
	return bal.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bal *bal_Latn) WeekdaysNarrow() []string {
	return bal.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bal *bal_Latn) WeekdayShort(weekday time.Weekday) string {
	if len(bal.daysShort) == 0 {
		return ""
	}
	return bal.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bal *bal_Latn) WeekdaysShort() []string {
	return bal.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bal *bal_Latn) WeekdayWide(weekday time.Weekday) string {
	if len(bal.daysWide) == 0 {
		return ""
	}
	return bal.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bal *bal_Latn) WeekdaysWide() []string {
	return bal.daysWide
}

// Decimal returns the decimal point of number
func (bal *bal_Latn) Decimal() string {
	return bal.decimal
}

// Group returns the group of number
func (bal *bal_Latn) Group() string {
	return bal.group
}

// Group returns the minus sign of number
func (bal *bal_Latn) Minus() string {
	return bal.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bal_Latn' and handles both Whole and Real numbers based on 'v'
func (bal *bal_Latn) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bal.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bal.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bal.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bal_Latn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bal *bal_Latn) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bal.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bal.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bal.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bal_Latn'
func (bal *bal_Latn) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bal.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bal.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(bal.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, bal.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, bal.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bal.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bal_Latn'
// in accounting notation.
func (bal *bal_Latn) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bal.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bal.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(bal.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, bal.currencyNegativePrefix[j])
		}

		b = append(b, bal.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(bal.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, bal.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, bal.currencyNegativeSuffix...)
	} else {
		b = append(b, bal.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bal.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bal.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2c}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bal.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, bal.periodsAbbreviated[0]...)
	} else {
		b = append(b, bal.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bal.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bal.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, bal.periodsAbbreviated[0]...)
	} else {
		b = append(b, bal.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bal.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bal.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, bal.periodsAbbreviated[0]...)
	} else {
		b = append(b, bal.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bal_Latn'
func (bal *bal_Latn) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	if h < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bal.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bal.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, bal.periodsAbbreviated[0]...)
	} else {
		b = append(b, bal.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bal.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
