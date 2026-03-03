package kab

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kab struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	timeSeparator      string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'kab' locale
func New() locales.Translator {
	return &kab{
		locale:             "kab",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ",",
		group:              " ",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$AR", "ATS", "$AU", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "FB", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$BM", "$BN", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "$BS", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "$BZ", "$CA", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$CL", "CNH", "CNX", "CNY", "$CO", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "£CY", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DA", "ECS", "ECV", "EEK", "£E", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "$FJ", "£FK", "F", "GBP", "GEK", "GEL", "GHC", "GHS", "£GI", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "£IE", "£IL", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "₤IT", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "£LB", "LKR", "LRD", "lLS", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "fMA", "MCF", "MDC", "MDL", "MGA", "Fmg", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "£MT", "MUR", "MVP", "MVR", "MWK", "$MX", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "$NA", "NGN", "NIC", "$C", "NLG", "NOK", "NPR", "$NZ", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "$RH", "ROL", "L", "RSD", "RUB", "RUR", "RWF", "SAR", "$SB", "SCR", "SDD", "SDG", "SDP", "SEK", "$SG", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "$SR", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "LT", "$TT", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "$UY", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WS$", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "DTS", "XEU", "XFO", "XFU", "XOF", "XPD", "FCFP", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "Kw", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "Yen", "Fur", "Meɣ", "Yeb", "May", "Yun", "Yul", "Ɣuc", "Cte", "Tub", "Nun", "Duǧ"},
		monthsNarrow:       []string{"", "Y", "F", "M", "Y", "M", "Y", "Y", "Ɣ", "C", "T", "N", "D"},
		monthsWide:         []string{"", "Yennayer", "Fuṛar", "Meɣres", "Yebrir", "Mayyu", "Yunyu", "Yulyu", "Ɣuct", "Ctembeṛ", "Tubeṛ", "Nunembeṛ", "Duǧembeṛ"},
		daysAbbreviated:    []string{"Acer", "Arim", "Aram", "Ahad", "Amhad", "Sem", "Sed"},
		daysNarrow:         []string{"C", "R", "R", "H", "M", "S", "S"},
		daysShort:          []string{"Acer", "Arim", "Aram", "Ahad", "Amhad", "Sem", "Sed"},
		daysWide:           []string{"Acer", "Arim", "Aram", "Ahad", "Amhad", "Sem", "Sed"},
		periodsAbbreviated: []string{"n\u202ftufat", "n\u202ftmeddit"},
		timezones:          map[string]string{"ACDT": "Akud n unebdu n Ustralya talemmast", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Akud n unebdu n tlemmast n umalu n Ustralya", "ACWST": "Akud amezday n tlemmast n umalu n Ustralya", "ADT": "Akud aṭlasan n unebdu", "ADT Arabia": "Akud aɛrab n unebdu", "AEDT": "Akud n unebdu n Ustralya n usammar", "AEST": "Akud amezday n Ustralya n usammar", "AFT": "Akud n Afɣanistan", "AKDT": "Akud n unebdu n Alaska", "AKST": "Akud amezday n Alaska", "AMST": "Akud n unebdu n Amaẓun", "AMST Armenia": "Akud n unebdu n Arminya", "AMT": "Akud amezday n Amaẓun", "AMT Armenia": "Akud amezday n Arminya", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Akud n unebdu n Arjuntin", "ART": "Akud amezday n Arjuntin", "AST": "Akud amezday aṭlasan", "AST Arabia": "Akud aɛrab amezday", "AWDT": "Akud n unebdu Ustralya n umalu", "AWST": "Akud amezday n Ustralya n umalu", "AZST": "Akud n unebdu n Aẓerbiǧan", "AZT": "Akud amezday n Aẓerbiǧan", "BDT Bangladesh": "Akud n unebdu n Bangladac", "BNT": "Akud n Brunay Dar Salam", "BOT": "Akud n Bulivi", "BRST": "Akud n unebdu n Braẓilya", "BRT": "Akud amezday n Braẓilya", "BST Bangladesh": "Akud amezday n Bangladac", "BT": "Akud n Butan", "CAST": "CAST", "CAT": "Akud n Tefriqt talemmast", "CCT": "Akud n Tegzirin n Kuku", "CDT": "Akud n unebdu n tlemmast n Marikan", "CHADT": "Akud n unebdu Catham", "CHAST": "Akud amezday n Catham", "CHUT": "Akud n Chuuk", "CKT": "Akud amezday n Tegzirin n Kuk", "CKT DST": "Akud n unebdu n Tegzirin n Kuk", "CLST": "Akud n unebdu n Cili", "CLT": "Akud amezday n Cili", "COST": "Akud n unebdu n Kulumbya", "COT": "Akud amezday n Kulumbya", "CST": "Akud amezday n tlemmast n Marikan", "CST China": "Akud amezday n Cin", "CST China DST": "Akud n unebdu n Cin", "CVST": "Akud n unebdu n Kap Vir", "CVT": "Akud amezday n Kap Vir", "CXT": "Akud n Tegzirin n Krismas", "ChST": "Akud amezday n Camurru", "ChST NMI": "ChST NMI", "CuDT": "Akud n unebdu n Kuba", "CuST": "Akud amezday n Kuba", "DAVT": "Akud n Davis", "DDUT": "Akud n Dumont-d’Urville", "EASST": "Akud n unebdu n Island n usammar", "EAST": "Akud amezday n Island n usammar", "EAT": "Akud n Tefriqt n usammar", "ECT": "Akud n Ikwaṭur", "EDT": "Akud n unebdu n usammar n Marikan", "EGDT": "Akud n unebdu n Grinland n usammar", "EGST": "Akud amezday n Grinland n usammar", "EST": "Akud amezday n usammar n Marikan", "FEET": "Akud nniḍen n Turuft n Usammar", "FJT": "Akud amezday n Fiji", "FJT Summer": "Akud n unebdu n Fiji", "FKST": "Akud n unebdu Tegzirin n Falkland", "FKT": "Akud amezday n Tegzirin n Falkland", "FNST": "Akud n unebdu n Firnandu n Nurunha", "FNT": "Akud amezday n Firnandu n Nurunha", "GALT": "Akud n Galapagus", "GAMT": "Akud n Tegzirin Gambyer", "GEST": "Akud n unebdu n Jyurjya", "GET": "Akud amezday n Jyurjya", "GFT": "Akud n Gwiyan tafransist", "GIT": "Akud n Tegzirin Jilbar", "GMT": "Akud alemmas n Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Akud n Jyurjya n unẓul", "GST Guam": "GST Guam", "GYT": "Akud n Gwiyan", "HADT": "Akud amezday n Haway-Aliwsyan", "HAST": "Akud amezday n Haway-Aliwsyan", "HKST": "Akud n unebdu n Hung Kung", "HKT": "Akud amezday n Hung Kung", "HOVST": "Akud n unebdu n Huvd", "HOVT": "Akud amezday n Huvd", "ICT": "Akud n Inducin", "IDT": "Akud n unebdu n Izrayil", "IOT": "Akud n Ugaraw Ahendi", "IRKST": "Akud n unebdu n Irkutsk", "IRKT": "Akud amezday n Irkutsk", "IRST": "Akud amezday n Iran", "IRST DST": "Akud n unebdu Iran", "IST": "Akud amezday n Lhend", "IST Israel": "Akud amezday n Izrayil", "JDT": "Akud n unebdu n Japun", "JST": "Akud amezday n Japun", "KOST": "Akud n Kosrae", "KRAST": "Akud n unebdu n Krasnoyarsk", "KRAT": "Akud amagnu n Krasnoyarsk", "KST": "Akud amezday n Kurya", "KST DST": "Akud n unebdu n Kurya", "LHDT": "Akud n Unebdu n Lord Howe", "LHST": "Akud Amagnu n Lord Howe", "LINT": "Akud n Tegzirin Lin", "MAGST": "Akud n unebdu n Magadan", "MAGT": "Akud amezday n Magadan", "MART": "Akud n Tegzirin Markiz", "MAWT": "Akud n Mawsun", "MDT": "MDT", "MESZ": "Akud n unebdu n Turuft talemmast", "MEZ": "Akud amezday n Turuft talemmast", "MHT": "Akud n Tegzirin Marcal", "MMT": "Akud n Myanmar", "MSD": "Akud n unebdu n Muṣku", "MST": "MST", "MUST": "Akud n unebdu n Muris", "MUT": "Akud amezday n Muris", "MVT": "Akud n Maldiv", "MYT": "Akud n Malizya", "NCT": "Akud amezday n Kalidunya Tamaynut", "NDT": "Akud n unebdu n Wakal amaynut", "NDT New Caledonia": "Akud n unebdu n Kalidunya Tamaynut", "NFDT": "Akud n unebdu n Tegzirt n Nurfulk", "NFT": "Akud amezday n Tegzirt n Nurfulk", "NOVST": "Akud n Unebdu n Novosibirsk", "NOVT": "Akud Amagnu n Novosibirsk", "NPT": "Akud n Nipal", "NRT": "Akud n Nuru", "NST": "Akud amezday n Wakal amaynut", "NUT": "Akud n Niyu", "NZDT": "Akud n unebdu Ziland Tamaynut", "NZST": "Akud amezday n Ziland Tamaynut", "OESZ": "Akud n unebdu n Turuft n usammar", "OEZ": "Akud amezday n Turuft n usammar", "OMSST": "Akud n Unebdu n Omsk", "OMST": "Akud Amagnu n Omsk", "PDT": "Akud umelwi n unebdu", "PDTM": "Akud amelwi n unebdu n Miksik", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Akud n Papwa n Ɣinya Tamaynut", "PHOT": "Akud n Tegzirin n Finiks", "PKT": "Akud amezday n Pakistan", "PKT DST": "Akud n unebdu n Pakistan", "PMDT": "Akud n unebdu n San Pyir & Miklun", "PMST": "Akud amezday n San Pyir & Miklun", "PONT": "Akud n Ponape", "PST": "Akud amezday amelwi", "PST Philippine": "Akud amezday n Filipin", "PST Philippine DST": "Akud n unebdu n Filipin", "PST Pitcairn": "Akud n Pitkarn", "PSTM": "Akud amezday amelwi n Miksik", "PWT": "Akud n Palau", "PYST": "Akud n unebdu n Paragway", "PYT": "Akud amezday n Paragway", "PYT Korea": "Akud n Pyungyung", "RET": "Akud n Riyunyun", "ROTT": "Akud n Rothera", "SAKST": "Akud n Unebdu n Sakhalin", "SAKT": "Akud Amagnu n Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Akud amezday n Tefriqt n unẓul", "SBT": "Akud n Tegzirin Salumun", "SCT": "Akud n Saycal", "SGT": "Akud amezday n Sangapur", "SLST": "SLST", "SRT": "Akud n Surinam", "SST Samoa": "Akud amezday n Samwa", "SST Samoa Apia": "Akud amezday n Apya", "SST Samoa Apia DST": "Akud n unebdu n Apya", "SST Samoa DST": "Akud n unebdu n Samwa", "SYOT": "Akud n Syuwa", "TAAF": "Akud n wakal n unẓul d Antarktik n Fransa", "TAHT": "Akud n Tahiti", "TJT": "Akud n Ṭajikistan", "TKT": "Akud n Ṭukilaw", "TLT": "Akud n Timur n usammar", "TMST": "Akud n unebdu n Turkmanistan", "TMT": "Akud amezday n Turkmanistan", "TOST": "Akud n unebdu n Ṭunga", "TOT": "Akud amezday n Ṭunga", "TVT": "Akud n Tuvalu", "TWT": "Akud amezday n Taypay", "TWT DST": "Akud n unebdu n Taypay", "ULAST": "Akud n unebdu n Ulanbatar", "ULAT": "Akud amezday n Ulanbatar", "UYST": "Akud n unebdu n Urugway", "UYT": "Akud amezday n Urugway", "UZT": "Akud amezday n Uzbikistan", "UZT DST": "Akud n unebdu n Uzbikistan", "VET": "Akud n Vinizwila", "VLAST": "Akud n Unebdu n Vladivostok", "VLAT": "Akud Amagnu n Vladivostok", "VOLST": "Akud n Unebdu n Volgograd", "VOLT": "Akud Amagnu n Volgograd", "VOST": "Akud n Vostok", "VUT": "Akud Amagnu n Vanuyatu", "VUT DST": "Akud n Unebdu n Vanuyatu", "WAKT": "Akud n Tegzirin n Wake", "WARST": "Akud n unebdu n Arjuntin n usammar", "WART": "Akud amezday n Arjuntin n usammar", "WAST": "Akud n Tefriqt n umalu", "WAT": "Akud n Tefriqt n umalu", "WESZ": "Akud n unebdu Turuft n umalu", "WEZ": "Akud amezday n Turuft n umalu", "WFT": "Akud n Wallis akked Futuna", "WGST": "Akud n unebdu n Grinland n umalu", "WGT": "Akud amezday n Grinland n umalu", "WIB": "Akud n umalu n Andunisya", "WIT": "Akud n usammar n Andunisya", "WITA": "Akud n tlemmast n Andunisya", "YAKST": "Akud n unebdu n Yakutsk", "YAKT": "Akud amezday n Yakutsk", "YEKST": "Akud n Unebdu n Yekaterinburg", "YEKT": "Akud Amagnu n Yekaterinburg", "YST": "YST", "МСК": "Akud amezday n Muṣku", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Akud n n umalu n Kazaxistan", "شىعىش قازاق ەلى": "Akud n usammar n Kazaxistan", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Akud n Kirigistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Akud n unebdu n Aẓuris"},
	}
}

// Locale returns the current translators string locale
func (kab *kab) Locale() string {
	return kab.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kab'
func (kab *kab) PluralsCardinal() []locales.PluralRule {
	return kab.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kab'
func (kab *kab) PluralsOrdinal() []locales.PluralRule {
	return kab.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kab'
func (kab *kab) PluralsRange() []locales.PluralRule {
	return kab.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kab'
func (kab *kab) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kab'
func (kab *kab) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kab'
func (kab *kab) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kab *kab) MonthAbbreviated(month time.Month) string {
	if len(kab.monthsAbbreviated) == 0 {
		return ""
	}
	return kab.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kab *kab) MonthsAbbreviated() []string {
	return kab.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kab *kab) MonthNarrow(month time.Month) string {
	if len(kab.monthsNarrow) == 0 {
		return ""
	}
	return kab.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kab *kab) MonthsNarrow() []string {
	return kab.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kab *kab) MonthWide(month time.Month) string {
	if len(kab.monthsWide) == 0 {
		return ""
	}
	return kab.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kab *kab) MonthsWide() []string {
	return kab.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kab *kab) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(kab.daysAbbreviated) == 0 {
		return ""
	}
	return kab.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kab *kab) WeekdaysAbbreviated() []string {
	return kab.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kab *kab) WeekdayNarrow(weekday time.Weekday) string {
	if len(kab.daysNarrow) == 0 {
		return ""
	}
	return kab.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kab *kab) WeekdaysNarrow() []string {
	return kab.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kab *kab) WeekdayShort(weekday time.Weekday) string {
	if len(kab.daysShort) == 0 {
		return ""
	}
	return kab.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kab *kab) WeekdaysShort() []string {
	return kab.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kab *kab) WeekdayWide(weekday time.Weekday) string {
	if len(kab.daysWide) == 0 {
		return ""
	}
	return kab.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kab *kab) WeekdaysWide() []string {
	return kab.daysWide
}

// Decimal returns the decimal point of number
func (kab *kab) Decimal() string {
	return kab.decimal
}

// Group returns the group of number
func (kab *kab) Group() string {
	return kab.group
}

// Group returns the minus sign of number
func (kab *kab) Minus() string {
	return kab.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kab' and handles both Whole and Real numbers based on 'v'
func (kab *kab) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kab.group) - 1; j >= 0; j-- {
					b = append(b, kab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kab' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kab *kab) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kab.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kab'
func (kab *kab) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kab.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kab.group) - 1; j >= 0; j-- {
					b = append(b, kab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kab.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kab'
// in accounting notation.
func (kab *kab) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kab.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kab.group) - 1; j >= 0; j-- {
					b = append(b, kab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, kab.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kab.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kab'
func (kab *kab) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kab'
func (kab *kab) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kab.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kab'
func (kab *kab) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kab.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kab'
func (kab *kab) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, kab.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kab.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kab'
func (kab *kab) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kab'
func (kab *kab) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kab'
func (kab *kab) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kab'
func (kab *kab) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kab.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
