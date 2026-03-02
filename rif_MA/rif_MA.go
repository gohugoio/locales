package rif_MA

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type rif_MA struct {
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

// New returns a new instance of translator for the 'rif_MA' locale
func New() locales.Translator {
	return &rif_MA{
		locale:                 "rif_MA",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "yen", "feb", "mar", "yeb", "may", "yun", "yul", "ɣuc", "cut", "kṭu", "nuw", "duj"},
		monthsNarrow:           []string{"", "y", "f", "m", "y", "m", "y", "y", "ɣ", "c", "k", "n", "d"},
		monthsWide:             []string{"", "yennayer", "febrayer", "mars", "yebril", "mayyu", "yunyu", "yulyu", "ɣuccet", "cutembir", "kṭuber", "nuwembir", "dujembir"},
		daysAbbreviated:        []string{"lḥe", "let", "ttl", "lar", "lex", "jje", "sse"},
		daysNarrow:             []string{"l", "l", "t", "l", "l", "j", "s"},
		daysShort:              []string{"lḥ", "le", "tt", "la", "lx", "jj", "ss"},
		daysWide:               []string{"lḥed", "letnayen", "ttlat", "larbeɛ", "lexmis", "jjemɛa", "ssebt"},
		periodsNarrow:          []string{"a", "p"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"zzat i yeccu", "awarni yeccu"},
		timezones:              map[string]string{"ACDT": "akud suntral n uzil n wustralya", "ACST": "akud suntral anaway n wustralya", "ACT": "ACT", "ACWDT": "akud suntral n uzil n lɣerb n wustralya", "ACWST": "akud suntral anaway n lɣerb n wustralya", "ADT": "akud n uzil n atlantik", "ADT Arabia": "akud n uzil n waɛrab", "AEDT": "akud n uzil n ccerq n wustralya", "AEST": "akud anaway n ccerq n wustralya", "AFT": "akud n afɣanistan", "AKDT": "akud n uzil n alaska", "AKST": "akud anaway n alaska", "AMST": "akud n uzil n amazun", "AMST Armenia": "akud n uzil n arminya", "AMT": "akud anaway n amazun", "AMT Armenia": "akud anaway n arminya", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "akud n uzil n arjentina", "ART": "akud anaway n arjentina", "AST": "akud anaway n atlantik", "AST Arabia": "akud anaway n waɛrab", "AWDT": "akud n uzil n lɣerb n wustralya", "AWST": "akud anaway n lɣerb n wustralya", "AZST": "akud n uzil n azrabidjan", "AZT": "akud anaway n azrabidjan", "BDT Bangladesh": "akud n uzil n bangladic", "BNT": "akud n brunay", "BOT": "akud n bulibya", "BRST": "akud n uzil n brazilya", "BRT": "akud anaway n brazilya", "BST Bangladesh": "akud anaway n bangladic", "BT": "akud n butan", "CAST": "CAST", "CAT": "akud n tefriqt n lwesṭ", "CCT": "akud n tiyzirin n kukus", "CDT": "akud n uzil n lwesṭ", "CHADT": "akud n uzil n tcatem", "CHAST": "akud anaway n tcatem", "CHUT": "akud n tcuk", "CKT": "akud anaway n tigzirin n kuk", "CKT DST": "akud n uzil n tigzirin n kuk", "CLST": "akud n uzil n cili", "CLT": "akud anaway n cili", "COST": "akud n uzil n kulumbya", "COT": "akud anaway n kulumbya", "CST": "akud anaway n lwesṭ", "CST China": "akud anaway n tcina", "CST China DST": "akud n uzil n tcina", "CVST": "akud n uzil n qabubirdi", "CVT": "akud anaway n qabubirdi", "CXT": "akud n tayzirt n kristmas", "ChST": "akud anaway n tcamuru", "ChST NMI": "ChST NMI", "CuDT": "akud n uzil n kuba", "CuST": "akud anaway n kuba", "DAVT": "akud n deybis", "DDUT": "akud n dimun durbil", "EASST": "akud n uzil n isterayland", "EAST": "akud anaway n isterayland", "EAT": "akud n tefriqt n ccerq", "ECT": "akud n ikwadur", "EDT": "akud n uzil n lɣerb", "EGDT": "akud n uzil n ccerq n grinland", "EGST": "akud anaway n ccerq n grinland", "EST": "akud anaway n lɣerb", "FEET": "akud n wuruppa tacerqect qaɛ", "FJT": "akud anaway n fiji", "FJT Summer": "akud n uzil n fiji", "FKST": "akud n uzil n falkland", "FKT": "akud anaway n falkland", "FNST": "akud n uzil n firnardu dinurunha", "FNT": "akud anaway n firnardu dinurunha", "GALT": "akud n galappagus", "GAMT": "akud n gambyi", "GEST": "akud n uzil n jyurjya", "GET": "akud anaway n jyurjya", "GFT": "akud n ɣana tafransist", "GIT": "akud n tigzirin n gilbert", "GMT": "GMT", "GNSST": "GNSST", "GNST": "GNST", "GST": "akud n jyurjya n wadday", "GST Guam": "GST Guam", "GYT": "akud n guyana", "HADT": "akud n uzil n haway-alucyan", "HAST": "akud anaway n haway-alucyan", "HKST": "akud n uzil n hungkung", "HKT": "akud anaway n hungkung", "HOVST": "akud n uzil n xubd", "HOVT": "akud anaway n xubd", "ICT": "akud n indutcina", "IDT": "akud n uzil n yisrayil", "IOT": "akud n lebḥer ahindi", "IRKST": "akud n uzil n irkutsek", "IRKT": "akud anaway n irkutsek", "IRST": "akud anaway n iran", "IRST DST": "akud n uzil n iran", "IST": "akud anawy n lhend", "IST Israel": "akud anaway n yisrayil", "JDT": "akud n uzil n jjapun", "JST": "akud anaway n jjapun", "KOST": "akud n kursay", "KRAST": "akud n uzil n krasnuyarsek", "KRAT": "akud anaway n krasnuyarsek", "KST": "akud anaway n kurya", "KST DST": "akud n uzil n kurya", "LHDT": "akud n uzil n lurdhaw", "LHST": "akud anaway n lurdhaw", "LINT": "akud n tigzirin n layn", "MAGST": "akud n uzil n magadan", "MAGT": "akud anaway n magadan", "MART": "akud n markwisas", "MAWT": "akud n mawson", "MDT": "akud n uzil n idurar", "MESZ": "akud n uzil n wuruppa n lwesṭ", "MEZ": "akud anaway n wuruppa n lwesṭ", "MHT": "akud n tigzirin n marcal", "MMT": "akud n myanmar", "MSD": "akud n uzil n musku", "MST": "akud anaway n idurar", "MUST": "akud n uzil n mawritus", "MUT": "akud anaway n mawritus", "MVT": "akud n lmaldib", "MYT": "akud n malizya", "NCT": "akud anaway n nyewkalidunya", "NDT": "akud n uzil n nyuw fawemd land", "NDT New Caledonia": "akud n uzil n nyewkalidunya", "NFDT": "akud n uzil n tayzirt n nurfuk", "NFT": "akud anaway n tayzirt n nurfuk", "NOVST": "akud n uzil n nubusibirsek", "NOVT": "akud anaway n nubusibirsek", "NPT": "akud n nnipal", "NRT": "akud n nawru", "NST": "akud anaway n nyuw fawemd land", "NUT": "akud n nyu", "NZDT": "akud n uzil n nyuziland", "NZST": "akud anaway n nyuziland", "OESZ": "akud n uzil n wuruppa n ccerq", "OEZ": "akud anaway n wuruppa n ccerq", "OMSST": "akud n uzil n umsek", "OMST": "akud anaway n umsek", "PDT": "akud n uzil n pasifik", "PDTM": "akud n uzil n pasifik amiksikan", "PETDT": "PETDT", "PETST": "PETST", "PGT": "akud n papwa ɣinya tamaynut", "PHOT": "akud n tigzirin n finiks", "PKT": "akud anaway n pakistan", "PKT DST": "akud n uzil n pakistan", "PMDT": "akud n uzil n sant-pyiɣ d mikilun", "PMST": "akud anaway n sant-pyiɣ d mikilun", "PONT": "akud n punpi", "PST": "akud anaway n pasifik", "PST Philippine": "akud anaway n filippin", "PST Philippine DST": "akud n uzil n filippin", "PST Pitcairn": "akud n pitkirn", "PSTM": "akud anaway n pasifik amiksikan", "PWT": "akud n palaw", "PYST": "akud n uzil n pparagway", "PYT": "akud anaway n pparagway", "PYT Korea": "akud n kurya n sennej", "RET": "akud n riyunyun", "ROTT": "akud n rutira", "SAKST": "akud n uzil n saxalin", "SAKT": "akud anaway n saxalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "akud n tefriqt n wadday", "SBT": "akud n tigzirin n sulumun", "SCT": "akud n saycal", "SGT": "akud anaway n singappur", "SLST": "SLST", "SRT": "akud n surinam", "SST Samoa": "akud anaway n samwa tamarikant", "SST Samoa Apia": "akud anaway n samwa", "SST Samoa Apia DST": "akud n uzil n samwa", "SST Samoa DST": "akud n uzil n samwa tamarikant", "SYOT": "akud n syuwa", "TAAF": "akud n tiwaddayin tifransisin", "TAHT": "akud n tahiti", "TJT": "akud n tajikistan", "TKT": "akud n tuwkulaw", "TLT": "akud n ṭimur licti", "TMST": "akud n uzil n ṭurkmanistan", "TMT": "akud anaway n ṭurkmanistan", "TOST": "akud n uzil n ṭunga", "TOT": "akud anaway n ṭunga", "TVT": "akud n ṭubalu", "TWT": "akud anaway n ṭṭaywan", "TWT DST": "akud n uzil n ṭṭaywan", "ULAST": "akud n uzil n ulanbaṭar", "ULAT": "akud anaway n ulanbaṭar", "UYST": "akud n uzil n urugway", "UYT": "akud anaway n urugway", "UZT": "akud n anaway n uẓbakistan", "UZT DST": "akud n uzil n uẓbakistan", "VET": "akud n vinzwila", "VLAST": "akud n uzil n bladibustuk", "VLAT": "akud anaway n bladibustuk", "VOLST": "akud n uzil n bulgugrad", "VOLT": "akud anaway n bulgugrad", "VOST": "akud n bustuk", "VUT": "akud anaway n banwatu", "VUT DST": "akud n uzil n banwatu", "WAKT": "akud n tagzirt n wayk", "WARST": "akud n uzil n lɣerb n arjentina", "WART": "akud anaway n lɣerb n arjentina", "WAST": "akud n tefriqt n lɣerb", "WAT": "akud n tefriqt n lɣerb", "WESZ": "akud n uzil n wuruppa n lɣerb", "WEZ": "akud anaway n wuruppa n lɣerb", "WFT": "akud n walis d futuna", "WGST": "akud n uzil n lwesṭ n grinland", "WGT": "akud anaway n lwesṭ n grinland", "WIB": "akud n lɣerb n yindunisya", "WIT": "akud n ccerq n yindunisya", "WITA": "akud n yindunisya n lwesṭ", "YAKST": "akud n uzil n yakutsek", "YAKT": "akud anaway n yakutsek", "YEKST": "akud n uzil n yakatirinburg", "YEKT": "akud anaway n yakatirinburg", "YST": "akud n yukun", "МСК": "akud anaway n musku", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "akud n lɣerb n kazaxistan", "شىعىش قازاق ەلى": "akud n ccerq n kazaxistan", "قازاق ەلى": "akud n kazaxistan", "قىرعىزستان": "akud n krigistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "akud n uzil n azures"},
	}
}

// Locale returns the current translators string locale
func (rif *rif_MA) Locale() string {
	return rif.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'rif_MA'
func (rif *rif_MA) PluralsCardinal() []locales.PluralRule {
	return rif.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'rif_MA'
func (rif *rif_MA) PluralsOrdinal() []locales.PluralRule {
	return rif.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'rif_MA'
func (rif *rif_MA) PluralsRange() []locales.PluralRule {
	return rif.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'rif_MA'
func (rif *rif_MA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'rif_MA'
func (rif *rif_MA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'rif_MA'
func (rif *rif_MA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (rif *rif_MA) MonthAbbreviated(month time.Month) string {
	return rif.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (rif *rif_MA) MonthsAbbreviated() []string {
	return rif.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (rif *rif_MA) MonthNarrow(month time.Month) string {
	return rif.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (rif *rif_MA) MonthsNarrow() []string {
	return rif.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (rif *rif_MA) MonthWide(month time.Month) string {
	return rif.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (rif *rif_MA) MonthsWide() []string {
	return rif.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (rif *rif_MA) WeekdayAbbreviated(weekday time.Weekday) string {
	return rif.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (rif *rif_MA) WeekdaysAbbreviated() []string {
	return rif.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (rif *rif_MA) WeekdayNarrow(weekday time.Weekday) string {
	return rif.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (rif *rif_MA) WeekdaysNarrow() []string {
	return rif.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (rif *rif_MA) WeekdayShort(weekday time.Weekday) string {
	return rif.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (rif *rif_MA) WeekdaysShort() []string {
	return rif.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (rif *rif_MA) WeekdayWide(weekday time.Weekday) string {
	return rif.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (rif *rif_MA) WeekdaysWide() []string {
	return rif.daysWide
}

// Decimal returns the decimal point of number
func (rif *rif_MA) Decimal() string {
	return rif.decimal
}

// Group returns the group of number
func (rif *rif_MA) Group() string {
	return rif.group
}

// Group returns the minus sign of number
func (rif *rif_MA) Minus() string {
	return rif.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'rif_MA' and handles both Whole and Real numbers based on 'v'
func (rif *rif_MA) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'rif_MA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (rif *rif_MA) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 1
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rif.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, rif.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, rif.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'rif_MA'
func (rif *rif_MA) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := rif.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rif.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(rif.group) - 1; j >= 0; j-- {
					b = append(b, rif.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, rif.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, rif.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, rif.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'rif_MA'
// in accounting notation.
func (rif *rif_MA) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := rif.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, rif.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(rif.group) - 1; j >= 0; j-- {
					b = append(b, rif.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, rif.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, rif.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, rif.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, rif.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, rif.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, rif.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, rif.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, rif.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'rif_MA'
func (rif *rif_MA) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, rif.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := rif.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
