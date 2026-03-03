package he

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type he struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'he' locale
func New() locales.Translator {
	return &he{
		locale:                 "he",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "\u200e-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "\u200eCN¥\u200e", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ל״י", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: "\u200f",
		currencyPositiveSuffix: " \u200f",
		currencyNegativePrefix: "\u200f-",
		currencyNegativeSuffix: " \u200f",
		monthsAbbreviated:      []string{"", "ינו׳", "פבר׳", "מרץ", "אפר׳", "מאי", "יוני", "יולי", "אוג׳", "ספט׳", "אוק׳", "נוב׳", "דצמ׳"},
		monthsWide:             []string{"", "ינואר", "פברואר", "מרץ", "אפריל", "מאי", "יוני", "יולי", "אוגוסט", "ספטמבר", "אוקטובר", "נובמבר", "דצמבר"},
		daysAbbreviated:        []string{"יום א׳", "יום ב׳", "יום ג׳", "יום ד׳", "יום ה׳", "יום ו׳", "שבת"},
		daysNarrow:             []string{"א׳", "ב׳", "ג׳", "ד׳", "ה׳", "ו׳", "ש׳"},
		daysShort:              []string{"א׳", "ב׳", "ג׳", "ד׳", "ה׳", "ו׳", "ש׳"},
		daysWide:               []string{"יום ראשון", "יום שני", "יום שלישי", "יום רביעי", "יום חמישי", "יום שישי", "יום שבת"},
		timezones:              map[string]string{"ACDT": "שעון מרכז אוסטרליה (קיץ)", "ACST": "שעון מרכז אוסטרליה (חורף)", "ACT": "ACT", "ACWDT": "שעון מרכז-מערב אוסטרליה (קיץ)", "ACWST": "שעון מרכז-מערב אוסטרליה (חורף)", "ADT": "שעון האוקיינוס האטלנטי (קיץ)", "ADT Arabia": "שעון חצי האי ערב (קיץ)", "AEDT": "שעון מזרח אוסטרליה (קיץ)", "AEST": "שעון מזרח אוסטרליה (חורף)", "AFT": "שעון אפגניסטן", "AKDT": "שעון אלסקה (קיץ)", "AKST": "שעון אלסקה (חורף)", "AMST": "שעון אמזונס (קיץ)", "AMST Armenia": "שעון ארמניה (קיץ)", "AMT": "שעון אמזונס (חורף)", "AMT Armenia": "שעון ארמניה (חורף)", "ANAST": "שעון קיץ אנדיר", "ANAT": "שעון רגיל אנדיר", "ARST": "שעון ארגנטינה (קיץ)", "ART": "שעון ארגנטינה (חורף)", "AST": "שעון האוקיינוס האטלנטי (חורף)", "AST Arabia": "שעון חצי האי ערב (חורף)", "AWDT": "שעון מערב אוסטרליה (קיץ)", "AWST": "שעון מערב אוסטרליה (חורף)", "AZST": "שעון אזרבייג׳ן (קיץ)", "AZT": "שעון אזרבייג׳ן (חורף)", "BDT Bangladesh": "שעון בנגלדש (קיץ)", "BNT": "שעון ברוניי דארוסלאם", "BOT": "שעון בוליביה", "BRST": "שעון ברזיליה (קיץ)", "BRT": "שעון ברזיליה (חורף)", "BST Bangladesh": "שעון בנגלדש (חורף)", "BT": "שעון בהוטן", "CAST": "CAST", "CAT": "שעון מרכז אפריקה", "CCT": "שעון איי קוקוס", "CDT": "שעון מרכז ארה״ב (קיץ)", "CHADT": "שעון צ׳טהאם (קיץ)", "CHAST": "שעון צ׳טהאם (חורף)", "CHUT": "שעון צ׳וק", "CKT": "שעון איי קוק (חורף)", "CKT DST": "שעון איי קוק (מחצית הקיץ)", "CLST": "שעון צ׳ילה (קיץ)", "CLT": "שעון צ׳ילה (חורף)", "COST": "שעון קולומביה (קיץ)", "COT": "שעון קולומביה (חורף)", "CST": "שעון מרכז ארה״ב (חורף)", "CST China": "שעון סין (חורף)", "CST China DST": "שעון סין (קיץ)", "CVST": "שעון כף ורדה (קיץ)", "CVT": "שעון כף ורדה (חורף)", "CXT": "שעון האי כריסטמס", "ChST": "שעון צ׳אמורו", "ChST NMI": "שעון איי מריאנה הצפוניים", "CuDT": "שעון קובה (קיץ)", "CuST": "שעון קובה (חורף)", "DAVT": "שעון דיוויס", "DDUT": "שעון דומון ד׳אורוויל", "EASST": "שעון אי הפסחא (קיץ)", "EAST": "שעון אי הפסחא (חורף)", "EAT": "שעון מזרח אפריקה", "ECT": "שעון אקוודור", "EDT": "שעון החוף המזרחי (קיץ)", "EGDT": "שעון מזרח גרינלנד (קיץ)", "EGST": "שעון מזרח גרינלנד (חורף)", "EST": "שעון החוף המזרחי (חורף)", "FEET": "שעון מינסק", "FJT": "שעון פיג׳י (חורף)", "FJT Summer": "שעון פיג׳י (קיץ)", "FKST": "שעון איי פוקלנד (קיץ)", "FKT": "שעון איי פוקלנד (חורף)", "FNST": "שעון פרננדו די נורוניה (קיץ)", "FNT": "שעון פרננדו די נורוניה (חורף)", "GALT": "שעון איי גלאפגוס", "GAMT": "שעון איי גמבייה", "GEST": "שעון גאורגיה (קיץ)", "GET": "שעון גאורגיה (חורף)", "GFT": "שעון גיאנה הצרפתית", "GIT": "שעון איי גילברט", "GMT": "שעון גריניץ׳\u200f", "GNSST": "GNSST", "GNST": "GNST", "GST": "שעון דרום ג׳ורג׳יה", "GST Guam": "שעון גואם", "GYT": "שעון גיאנה", "HADT": "שעון האיים האלאוטיים הוואי (קיץ)", "HAST": "שעון האיים האלאוטיים הוואי (חורף)", "HKST": "שעון הונג קונג (קיץ)", "HKT": "שעון הונג קונג (חורף)", "HOVST": "שעון חובד (קיץ)", "HOVT": "שעון חובד (חורף)", "ICT": "שעון הודו-סין", "IDT": "שעון ישראל (קיץ)", "IOT": "שעון האוקיינוס ההודי", "IRKST": "שעון אירקוסטק (קיץ)", "IRKT": "שעון אירקוטסק (חורף)", "IRST": "שעון איראן (חורף)", "IRST DST": "שעון איראן (קיץ)", "IST": "שעון הודו", "IST Israel": "שעון ישראל (חורף)", "JDT": "שעון יפן (קיץ)", "JST": "שעון יפן (חורף)", "KOST": "שעון קוסראה", "KRAST": "שעון קרסנויארסק (קיץ)", "KRAT": "שעון קרסנויארסק (חורף)", "KST": "שעון קוריאה (חורף)", "KST DST": "שעון קוריאה (קיץ)", "LHDT": "שעון אי הלורד האו (קיץ)", "LHST": "שעון אי הלורד האו (חורף)", "LINT": "שעון איי ליין", "MAGST": "שעון מגדן (קיץ)", "MAGT": "שעון מגדן (חורף)", "MART": "שעון איי מרקיז", "MAWT": "שעון מאוסון", "MDT": "שעון קיץ מקאו", "MESZ": "שעון מרכז אירופה (קיץ)", "MEZ": "שעון מרכז אירופה (חורף)", "MHT": "שעון איי מרשל", "MMT": "שעון מיאנמר", "MSD": "שעון מוסקבה (קיץ)", "MST": "שעון חורף מקאו", "MUST": "שעון מאוריציוס (קיץ)", "MUT": "שעון מאוריציוס (חורף)", "MVT": "שעון האיים המלדיביים", "MYT": "שעון מלזיה", "NCT": "שעון קלדוניה החדשה (חורף)", "NDT": "שעון ניופאונדלנד (קיץ)", "NDT New Caledonia": "שעון קלדוניה החדשה (קיץ)", "NFDT": "שעון האי נורפוק (קיץ)", "NFT": "שעון האי נורפוק (חורף)", "NOVST": "שעון נובוסיבירסק (קיץ)", "NOVT": "שעון נובוסיבירסק (חורף)", "NPT": "שעון נפאל", "NRT": "שעון נאורו", "NST": "שעון ניופאונדלנד (חורף)", "NUT": "שעון ניואה", "NZDT": "שעון ניו זילנד (קיץ)", "NZST": "שעון ניו זילנד (חורף)", "OESZ": "שעון מזרח אירופה (קיץ)", "OEZ": "שעון מזרח אירופה (חורף)", "OMSST": "שעון אומסק (קיץ)", "OMST": "שעון אומסק (חורף)", "PDT": "שעון מערב ארה״ב (קיץ)", "PDTM": "שעון מערב מקסיקו (קיץ)", "PETDT": "שעון קיץ פטרופבלובסק-קמצ׳טסקי", "PETST": "שעון רגיל פטרופבלובסק-קמצ׳טסקי", "PGT": "שעון פפואה גיניאה החדשה", "PHOT": "שעון איי פיניקס", "PKT": "שעון פקיסטן (חורף)", "PKT DST": "שעון פקיסטן (קיץ)", "PMDT": "שעון סנט פייר ומיקלון (קיץ)", "PMST": "שעון סנט פייר ומיקלון (חורף)", "PONT": "שעון פונאפי", "PST": "שעון מערב ארה״ב (חורף)", "PST Philippine": "שעון הפיליפינים (חורף)", "PST Philippine DST": "שעון הפיליפינים (קיץ)", "PST Pitcairn": "שעון פיטקרן", "PSTM": "שעון מערב מקסיקו (חורף)", "PWT": "שעון פלאו", "PYST": "שעון פרגוואי (קיץ)", "PYT": "שעון פרגוואי (חורף)", "PYT Korea": "שעון פיונגיאנג", "RET": "שעון ראוניון", "ROTT": "שעון רות׳רה", "SAKST": "שעון סחלין (קיץ)", "SAKT": "שעון סחלין (חורף)", "SAMST": "שעון קיץ סמרה", "SAMT": "שעון רגיל סמרה", "SAST": "שעון דרום אפריקה", "SBT": "שעון איי שלמה", "SCT": "שעון איי סיישל", "SGT": "שעון סינגפור", "SLST": "SLST", "SRT": "שעון סורינאם", "SST Samoa": "שעון סמואה (חורף)", "SST Samoa Apia": "שעון אפיה (חורף)", "SST Samoa Apia DST": "שעון אפיה (קיץ)", "SST Samoa DST": "שעון סמואה (קיץ)", "SYOT": "שעון סייווה", "TAAF": "שעון הארצות הדרומיות והאנטארקטיות של צרפת", "TAHT": "שעון טהיטי", "TJT": "שעון טג׳יקיסטן", "TKT": "שעון טוקלאו", "TLT": "שעון מזרח טימור", "TMST": "שעון טורקמניסטן (קיץ)", "TMT": "שעון טורקמניסטן (חורף)", "TOST": "שעון טונגה (קיץ)", "TOT": "שעון טונגה (חורף)", "TVT": "שעון טובאלו", "TWT": "שעון טאיפיי (חורף)", "TWT DST": "שעון טאיפיי (קיץ)", "ULAST": "שעון אולאן באטור (קיץ)", "ULAT": "שעון אולאן באטור (חורף)", "UYST": "שעון אורוגוואי (קיץ)", "UYT": "שעון אורוגוואי (חורף)", "UZT": "שעון אוזבקיסטן (חורף)", "UZT DST": "שעון אוזבקיסטן (קיץ)", "VET": "שעון ונצואלה", "VLAST": "שעון ולדיווסטוק (קיץ)", "VLAT": "שעון ולדיווסטוק (חורף)", "VOLST": "שעון וולגוגרד (קיץ)", "VOLT": "שעון וולגוגרד (חורף)", "VOST": "שעון ווסטוק", "VUT": "שעון ונואטו (חורף)", "VUT DST": "שעון ונואטו (קיץ)", "WAKT": "שעון האי וייק", "WARST": "שעון מערב ארגנטינה (קיץ)", "WART": "שעון מערב ארגנטינה (חורף)", "WAST": "שעון מערב אפריקה", "WAT": "שעון מערב אפריקה", "WESZ": "שעון מערב אירופה (קיץ)", "WEZ": "שעון מערב אירופה (חורף)", "WFT": "שעון וואליס ופוטונה", "WGST": "שעון מערב גרינלנד (קיץ)", "WGT": "שעון מערב גרינלנד (חורף)", "WIB": "שעון מערב אינדונזיה", "WIT": "שעון מזרח אינדונזיה", "WITA": "שעון מרכז אינדונזיה", "YAKST": "שעון יקוטסק (קיץ)", "YAKT": "שעון יקוטסק (חורף)", "YEKST": "שעון יקטרינבורג (קיץ)", "YEKT": "שעון יקטרינבורג (חורף)", "YST": "שעון יוקון", "МСК": "שעון מוסקבה (חורף)", "اقتاۋ": "שעון אקטאו (חורף)", "اقتاۋ قالاسى": "שעון אקטאו (קיץ)", "اقتوبە": "שעון אוקטובה (חורף)", "اقتوبە قالاسى": "שעון אוקטובה (קיץ)", "الماتى": "שעון אלמטי (חורף)", "الماتى قالاسى": "שעון אלמטי (קיץ)", "باتىس قازاق ەلى": "שעון מערב קזחסטן", "شىعىش قازاق ەلى": "שעון מזרח קזחסטן", "قازاق ەلى": "שעון קזחסטן", "قىرعىزستان": "שעון קירגיזסטן", "قىزىلوردا": "שעון קיזילורדה (חורף)", "قىزىلوردا قالاسى": "שעון קיזילורדה (קיץ)", "∅∅∅": "שעון פרו (קיץ)"},
	}
}

// Locale returns the current translators string locale
func (he *he) Locale() string {
	return he.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'he'
func (he *he) PluralsCardinal() []locales.PluralRule {
	return he.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'he'
func (he *he) PluralsOrdinal() []locales.PluralRule {
	return he.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'he'
func (he *he) PluralsRange() []locales.PluralRule {
	return he.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'he'
func (he *he) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 1 && v == 0) || (i == 0 && v != 0) {
		return locales.PluralRuleOne
	} else if i == 2 && v == 0 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'he'
func (he *he) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'he'
func (he *he) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (he *he) MonthAbbreviated(month time.Month) string {
	if len(he.monthsAbbreviated) == 0 {
		return ""
	}
	return he.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (he *he) MonthsAbbreviated() []string {
	return he.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (he *he) MonthNarrow(month time.Month) string {
	if len(he.monthsNarrow) == 0 {
		return ""
	}
	return he.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (he *he) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (he *he) MonthWide(month time.Month) string {
	if len(he.monthsWide) == 0 {
		return ""
	}
	return he.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (he *he) MonthsWide() []string {
	return he.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (he *he) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(he.daysAbbreviated) == 0 {
		return ""
	}
	return he.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (he *he) WeekdaysAbbreviated() []string {
	return he.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (he *he) WeekdayNarrow(weekday time.Weekday) string {
	if len(he.daysNarrow) == 0 {
		return ""
	}
	return he.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (he *he) WeekdaysNarrow() []string {
	return he.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (he *he) WeekdayShort(weekday time.Weekday) string {
	if len(he.daysShort) == 0 {
		return ""
	}
	return he.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (he *he) WeekdaysShort() []string {
	return he.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (he *he) WeekdayWide(weekday time.Weekday) string {
	if len(he.daysWide) == 0 {
		return ""
	}
	return he.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (he *he) WeekdaysWide() []string {
	return he.daysWide
}

// Decimal returns the decimal point of number
func (he *he) Decimal() string {
	return he.decimal
}

// Group returns the group of number
func (he *he) Group() string {
	return he.group
}

// Group returns the minus sign of number
func (he *he) Minus() string {
	return he.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'he' and handles both Whole and Real numbers based on 'v'
func (he *he) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, he.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, he.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(he.minus) - 1; j >= 0; j-- {
			b = append(b, he.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'he' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (he *he) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, he.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(he.minus) - 1; j >= 0; j-- {
			b = append(b, he.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, he.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'he'
func (he *he) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := he.currencies[currency]
	l := len(s) + len(symbol) + 13 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, he.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, he.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(he.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, he.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(he.minus) - 1; j >= 0; j-- {
			b = append(b, he.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, he.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, he.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'he'
// in accounting notation.
func (he *he) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := he.currencies[currency]
	l := len(s) + len(symbol) + 14 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, he.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, he.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(he.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, he.currencyNegativePrefix[j])
		}

	} else {

		for j := len(he.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, he.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, he.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, he.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, he.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'he'
func (he *he) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'he'
func (he *he) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xd7, 0x91}...)
	b = append(b, he.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'he'
func (he *he) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xd7, 0x91}...)
	b = append(b, he.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'he'
func (he *he) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, he.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xd7, 0x91}...)
	b = append(b, he.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'he'
func (he *he) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, he.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'he'
func (he *he) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, he.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, he.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'he'
func (he *he) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, he.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, he.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'he'
func (he *he) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, he.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, he.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := he.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
