package fa_AF

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type fa_AF struct {
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

// New returns a new instance of translator for the 'fa_AF' locale
func New() locales.Translator {
	return &fa_AF{
		locale:                 "fa_AF",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "\u200e−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "\u200e( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "جنو", "فبروری", "مارچ", "اپریل", "می", "جون", "جول", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسم"},
		monthsNarrow:           []string{"", "ج", "ف", "م", "ا", "م", "ج", "ج", "ا", "س", "ا", "ن", "د"},
		monthsWide:             []string{"", "جنوری", "فبروری", "مارچ", "اپریل", "می", "جون", "جولای", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysNarrow:             []string{"ی", "د", "س", "چ", "پ", "ج", "ش"},
		daysShort:              []string{"۱ش", "۲ش", "۳ش", "۴ش", "۵ش", "ج", "ش"},
		daysWide:               []string{"یکشنبه", "دوشنبه", "سه\u200cشنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"},
		timezones:              map[string]string{"ACDT": "وقت تابستانی مرکز استرالیا", "ACST": "ACST", "ACT": "ACT", "ACWDT": "وقت تابستانی مرکز استرالیای غربی", "ACWST": "وقت عادی مرکز استرالیای غربی", "ADT": "وقت تابستانی آتلانتیک", "ADT Arabia": "وقت تابستانی عربستان", "AEDT": "وقت تابستانی استرالیای شرقی", "AEST": "وقت عادی استرالیای شرقی", "AFT": "وقت افغانستان", "AKDT": "وقت تابستانی آلاسکا", "AKST": "وقت عادی آلاسکا", "AMST": "وقت تابستانی آمازون", "AMST Armenia": "وقت تابستانی ارمنستان", "AMT": "وقت عادی آمازون", "AMT Armenia": "وقت عادی ارمنستان", "ANAST": "وقت تابستانی آنادیر", "ANAT": "وقت عادی آنادیر", "ARST": "وقت تابستانی آرژانتین", "ART": "وقت عادی آرژانتین", "AST": "وقت عادی آتلانتیک", "AST Arabia": "وقت عادی عربستان", "AWDT": "وقت تابستانی استرالیای غربی", "AWST": "وقت عادی استرالیای غربی", "AZST": "وقت تابستانی جمهوری آذربایجان", "AZT": "وقت عادی جمهوری آذربایجان", "BDT Bangladesh": "وقت تابستانی بنگلادش", "BNT": "وقت برونئی دارالسلام", "BOT": "وقت بولیوی", "BRST": "وقت تابستانی برازیلیا", "BRT": "وقت عادی برازیلیا", "BST Bangladesh": "وقت عادی بنگلادش", "BT": "وقت بوتان", "CAST": "CAST", "CAT": "وقت مرکز آفریقا", "CCT": "وقت جزایر کوکوس", "CDT": "وقت تابستانی مرکز امریکا", "CHADT": "وقت تابستانی چت\u200cهام", "CHAST": "وقت عادی چت\u200cهام", "CHUT": "وقت چوئوک", "CKT": "وقت عادی جزایر کوک", "CKT DST": "وقت تابستانی جزایر کوک", "CLST": "وقت تابستانی شیلی", "CLT": "وقت عادی شیلی", "COST": "وقت تابستانی کلمبیا", "COT": "وقت عادی کلمبیا", "CST": "وقت عادی مرکز امریکا", "CST China": "وقت عادی چین", "CST China DST": "وقت تابستانی چین", "CVST": "وقت تابستانی کیپ\u200cورد", "CVT": "وقت عادی کیپ\u200cورد", "CXT": "وقت جزیرهٔ کریسمس", "ChST": "وقت عادی چامورو", "ChST NMI": "وقت جزایر ماریانای شمالی", "CuDT": "وقت تابستانی کوبا", "CuST": "وقت عادی کوبا", "DAVT": "وقت دیویس", "DDUT": "وقت دومون دورویل", "EASST": "وقت تابستانی جزیرهٔ ایستر", "EAST": "وقت عادی جزیرهٔ ایستر", "EAT": "وقت شرق افریقا", "ECT": "وقت اکوادور", "EDT": "وقت تابستانی شرق امریکا", "EGDT": "وقت تابستانی شرق گرینلند", "EGST": "وقت عادی شرق گرینلند", "EST": "وقت عادی شرق امریکا", "FEET": "وقت تابستانی مکان\u200cهای دیگر شرق اروپا", "FJT": "وقت عادی فیجی", "FJT Summer": "وقت تابستانی فیجی", "FKST": "وقت تابستانی جزایر فالکلند", "FKT": "وقت عادی جزایر فالکلند", "FNST": "وقت تابستانی فرناندو دی نورونیا", "FNT": "وقت عادی فرناندو دی نورونیا", "GALT": "وقت گالاپاگوس", "GAMT": "وقت گامبیه", "GEST": "وقت تابستانی گرجستان", "GET": "وقت عادی گرجستان", "GFT": "وقت گویان فرانسه", "GIT": "وقت جزایر گیلبرت", "GMT": "وقت گرینویچ", "GNSST": "GNSST", "GNST": "GNST", "GST": "وقت عادی خلیج فارس", "GST Guam": "وقت عادی گوام", "GYT": "وقت گویان", "HADT": "وقت عادی هاوایی‐الوشن", "HAST": "وقت عادی هاوایی‐الوشن", "HKST": "وقت تابستانی هنگ\u200cکنگ", "HKT": "وقت عادی هنگ\u200cکنگ", "HOVST": "وقت تابستانی خوود", "HOVT": "وقت عادی خوود", "ICT": "وقت هندوچین", "IDT": "وقت تابستانی اسرائیل", "IOT": "وقت اقیانوس هند", "IRKST": "وقت تابستانی ایرکوتسک", "IRKT": "وقت عادی ایرکوتسک", "IRST": "وقت عادی ایران", "IRST DST": "وقت تابستانی ایران", "IST": "وقت هند", "IST Israel": "وقت عادی اسرائیل", "JDT": "وقت تابستانی ژاپن", "JST": "وقت عادی ژاپن", "KOST": "وقت کوسرای", "KRAST": "وقت تابستانی کراسنویارسک", "KRAT": "وقت عادی کراسنویارسک", "KST": "وقت عادی کره", "KST DST": "وقت تابستانی کره", "LHDT": "وقت تابستانی لردهو", "LHST": "وقت عادی لردهو", "LINT": "وقت جزایر لاین", "MAGST": "وقت تابستانی ماگادان", "MAGT": "وقت عادی ماگادان", "MART": "وقت مارکوئز", "MAWT": "وقت ماوسون", "MDT": "وقت تابستانی کوهستانی امریکا", "MESZ": "وقت تابستانی مرکز اروپا", "MEZ": "وقت عادی مرکز اروپا", "MHT": "وقت جزایر مارشال", "MMT": "وقت میانمار", "MSD": "وقت تابستانی مسکو", "MST": "وقت عادی کوهستانی امریکا", "MUST": "وقت تابستانی موریس", "MUT": "وقت عادی موریس", "MVT": "وقت مالدیو", "MYT": "وقت مالزی", "NCT": "وقت عادی کالدونیای جدید", "NDT": "وقت تابستانی نیوفاندلند", "NDT New Caledonia": "وقت تابستانی کالدونیای جدید", "NFDT": "وقت تابستانی جزیرهٔ نورفولک", "NFT": "وقت عادی جزیرهٔ نورفولک", "NOVST": "وقت تابستانی نووسیبیرسک", "NOVT": "وقت عادی نووسیبیرسک", "NPT": "وقت نپال", "NRT": "وقت نائورو", "NST": "وقت عادی نیوفاندلند", "NUT": "وقت نیوئه", "NZDT": "وقت تابستانی نیوزیلند", "NZST": "وقت عادی نیوزیلند", "OESZ": "وقت تابستانی شرق اروپا", "OEZ": "وقت عادی شرق اروپا", "OMSST": "وقت تابستانی اومسک", "OMST": "وقت عادی اومسک", "PDT": "وقت تابستانی غرب امریکا", "PDTM": "وقت تابستانی شرق مکزیک", "PETDT": "وقت تابستانی پتروپاولوسک‐کامچاتسکی", "PETST": "وقت عادی پتروپاولوسک‐کامچاتسکی", "PGT": "وقت پاپوا گینهٔ نو", "PHOT": "وقت جزایر فونیکس", "PKT": "وقت عادی پاکستان", "PKT DST": "وقت تابستانی پاکستان", "PMDT": "وقت تابستانی سنت\u200cپیر و میکلون", "PMST": "وقت عادی سنت\u200cپیر و میکلون", "PONT": "وقت پوناپه", "PST": "وقت عادی غرب امریکا", "PST Philippine": "وقت عادی فیلیپین", "PST Philippine DST": "وقت تابستانی فیلیپین", "PST Pitcairn": "وقت پیتکایرن", "PSTM": "وقت عادی شرق مکزیک", "PWT": "وقت پالائو", "PYST": "وقت تابستانی پاراگوئه", "PYT": "وقت عادی پاراگوئه", "PYT Korea": "وقت پیونگ\u200cیانگ", "RET": "وقت رئونیون", "ROTT": "وقت روترا", "SAKST": "وقت تابستانی ساخالین", "SAKT": "وقت عادی ساخالین", "SAMST": "وقت تابستانی سامارا", "SAMT": "وقت عادی سامارا", "SAST": "وقت عادی جنوب افریقا", "SBT": "وقت جزایر سلیمان", "SCT": "وقت سیشل", "SGT": "وقت سنگاپور", "SLST": "وقت لانکا", "SRT": "وقت سورینام", "SST Samoa": "وقت عادی ساموا", "SST Samoa Apia": "وقت عادی آپیا", "SST Samoa Apia DST": "وقت تابستانی آپیا", "SST Samoa DST": "وقت تابستانی ساموا", "SYOT": "وقت شووا", "TAAF": "وقت سرزمین\u200cهای جنوبی و جنوبگان فرانسه", "TAHT": "وقت تاهیتی", "TJT": "وقت تاجیکستان", "TKT": "وقت توکلائو", "TLT": "وقت تیمور شرقی", "TMST": "وقت تابستانی ترکمنستان", "TMT": "وقت عادی ترکمنستان", "TOST": "وقت تابستانی تونگا", "TOT": "وقت عادی تونگا", "TVT": "وقت تووالو", "TWT": "وقت عادی تایپه", "TWT DST": "وقت تابستانی تایپه", "ULAST": "وقت تابستانی اولان\u200cباتور", "ULAT": "وقت عادی اولان\u200cباتور", "UYST": "وقت تابستانی اروگوئه", "UYT": "وقت عادی اروگوئه", "UZT": "وقت عادی ازبکستان", "UZT DST": "وقت تابستانی ازبکستان", "VET": "وقت ونزوئلا", "VLAST": "وقت تابستانی ولادی\u200cوستوک", "VLAT": "وقت عادی ولادی\u200cوستوک", "VOLST": "وقت تابستانی ولگاگراد", "VOLT": "وقت عادی ولگاگراد", "VOST": "وقت وستوک", "VUT": "وقت عادی واناتو", "VUT DST": "وقت تابستانی واناتو", "WAKT": "وقت جزیرهٔ ویک", "WARST": "وقت تابستانی غرب آرژانتین", "WART": "وقت عادی غرب آرژانتین", "WAST": "وقت غرب افریقا", "WAT": "وقت غرب افریقا", "WESZ": "وقت تابستانی غرب اروپا", "WEZ": "وقت عادی غرب اروپا", "WFT": "وقت والیس و فوتونا", "WGST": "وقت تابستانی غرب گرینلند", "WGT": "وقت عادی غرب گرینلند", "WIB": "وقت غرب اندونزی", "WIT": "وقت شرق اندونزی", "WITA": "وقت مرکز اندونزی", "YAKST": "وقت تابستانی یاکوتسک", "YAKT": "وقت عادی یاکوتسک", "YEKST": "وقت تابستانی یکاترینبورگ", "YEKT": "وقت عادی یکاترینبورگ", "YST": "وقت یوکان", "МСК": "وقت عادی مسکو", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "وقت عادی آلماآتا", "الماتى قالاسى": "وقت تابستانی آلماآتا", "باتىس قازاق ەلى": "وقت غرب قزاقستان", "شىعىش قازاق ەلى": "وقت شرق قزاقستان", "قازاق ەلى": "وقت قزاقستان", "قىرعىزستان": "وقت قرقیزستان", "قىزىلوردا": "وقت عادی قیزیل\u200cاوردا", "قىزىلوردا قالاسى": "وقت تابستانی قیزیل\u200cاوردا", "∅∅∅": "وقت تابستانی آزور"},
	}
}

// Locale returns the current translators string locale
func (fa *fa_AF) Locale() string {
	return fa.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fa_AF'
func (fa *fa_AF) PluralsCardinal() []locales.PluralRule {
	return fa.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fa_AF'
func (fa *fa_AF) PluralsOrdinal() []locales.PluralRule {
	return fa.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fa_AF'
func (fa *fa_AF) PluralsRange() []locales.PluralRule {
	return fa.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fa_AF'
func (fa *fa_AF) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fa_AF'
func (fa *fa_AF) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fa_AF'
func (fa *fa_AF) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := fa.CardinalPluralRule(num1, v1)
	end := fa.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fa *fa_AF) MonthAbbreviated(month time.Month) string {
	return fa.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fa *fa_AF) MonthsAbbreviated() []string {
	return fa.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fa *fa_AF) MonthNarrow(month time.Month) string {
	return fa.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fa *fa_AF) MonthsNarrow() []string {
	return fa.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fa *fa_AF) MonthWide(month time.Month) string {
	return fa.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fa *fa_AF) MonthsWide() []string {
	return fa.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayAbbreviated(weekday time.Weekday) string {
	return fa.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fa *fa_AF) WeekdaysAbbreviated() []string {
	return fa.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayNarrow(weekday time.Weekday) string {
	return fa.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fa *fa_AF) WeekdaysNarrow() []string {
	return fa.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayShort(weekday time.Weekday) string {
	return fa.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fa *fa_AF) WeekdaysShort() []string {
	return fa.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fa *fa_AF) WeekdayWide(weekday time.Weekday) string {
	return fa.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fa *fa_AF) WeekdaysWide() []string {
	return fa.daysWide
}

// Decimal returns the decimal point of number
func (fa *fa_AF) Decimal() string {
	return fa.decimal
}

// Group returns the group of number
func (fa *fa_AF) Group() string {
	return fa.group
}

// Group returns the minus sign of number
func (fa *fa_AF) Minus() string {
	return fa.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fa_AF' and handles both Whole and Real numbers based on 'v'
func (fa *fa_AF) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fa.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fa.minus) - 1; j >= 0; j-- {
			b = append(b, fa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fa_AF' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fa *fa_AF) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 8
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fa.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fa.minus) - 1; j >= 0; j-- {
			b = append(b, fa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fa.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fa_AF'
func (fa *fa_AF) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fa.currencies[currency]
	l := len(s) + len(symbol) + 9 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fa.group[0])
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

	for j := len(fa.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, fa.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(fa.minus) - 1; j >= 0; j-- {
			b = append(b, fa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fa_AF'
// in accounting notation.
func (fa *fa_AF) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fa.currencies[currency]
	l := len(s) + len(symbol) + 14 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fa.group[0])
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

		for j := len(fa.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, fa.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(fa.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, fa.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fa.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, fa.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

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

// FmtTimeFull returns the full time representation of 't' for 'fa_AF'
func (fa *fa_AF) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := fa.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
