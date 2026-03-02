package ar_001

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ar_001 struct {
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

// New returns a new instance of translator for the 'ar_001' locale
func New() locales.Translator {
	return &ar_001{
		locale:                 "ar_001",
		pluralsCardinal:        []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{1, 4, 5, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "\u200e-",
		percent:                "\u200e%\u200e",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: "\u200f",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "\u200f",
		currencyNegativeSuffix: " ",
		monthsNarrow:           []string{"", "ي", "ف", "م", "أ", "و", "ن", "ل", "غ", "س", "ك", "ب", "د"},
		monthsWide:             []string{"", "يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		daysNarrow:             []string{"ح", "ن", "ث", "ر", "خ", "ج", "س"},
		daysShort:              []string{"أحد", "إثنين", "ثلاثاء", "أربعاء", "خميس", "جمعة", "سبت"},
		daysWide:               []string{"الأحد", "الاثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
		periodsAbbreviated:     []string{"ص", "م"},
		timezones:              map[string]string{"ACDT": "توقيت وسط أستراليا الصيفي", "ACST": "توقيت وسط أستراليا الرسمي", "ACT": "ACT", "ACWDT": "توقيت غرب وسط أستراليا الصيفي", "ACWST": "توقيت غرب وسط أستراليا الرسمي", "ADT": "التوقيت الصيفي الأطلسي", "ADT Arabia": "التوقيت العربي الصيفي", "AEDT": "توقيت شرق أستراليا الصيفي", "AEST": "توقيت شرق أستراليا الرسمي", "AFT": "توقيت أفغانستان", "AKDT": "توقيت ألاسكا الصيفي", "AKST": "التوقيت الرسمي لألاسكا", "AMST": "توقيت الأمازون الصيفي", "AMST Armenia": "توقيت أرمينيا الصيفي", "AMT": "توقيت الأمازون الرسمي", "AMT Armenia": "توقيت أرمينيا الرسمي", "ANAST": "التوقيت الصيفي لأنادير", "ANAT": "توقيت أنادير الرسمي", "ARST": "توقيت الأرجنتين الصيفي", "ART": "توقيت الأرجنتين الرسمي", "AST": "التوقيت الرسمي الأطلسي", "AST Arabia": "التوقيت العربي الرسمي", "AWDT": "توقيت غرب أستراليا الصيفي", "AWST": "توقيت غرب أستراليا الرسمي", "AZST": "توقيت أذربيجان الصيفي", "AZT": "توقيت أذربيجان الرسمي", "BDT Bangladesh": "توقيت بنغلاديش الصيفي", "BNT": "توقيت بروناي", "BOT": "توقيت بوليفيا", "BRST": "توقيت برازيليا الصيفي", "BRT": "توقيت برازيليا الرسمي", "BST Bangladesh": "توقيت بنغلاديش الرسمي", "BT": "توقيت بوتان", "CAST": "CAST", "CAT": "توقيت وسط أفريقيا", "CCT": "توقيت جزر كوكوس", "CDT": "التوقيت الصيفي المركزي لأمريكا الشمالية", "CHADT": "توقيت تشاتام الصيفي", "CHAST": "توقيت تشاتام الرسمي", "CHUT": "توقيت شوك", "CKT": "توقيت جزر كوك الرسمي", "CKT DST": "توقيت جزر كوك الصيفي", "CLST": "توقيت تشيلي الصيفي", "CLT": "توقيت تشيلي الرسمي", "COST": "توقيت كولومبيا الصيفي", "COT": "توقيت كولومبيا الرسمي", "CST": "التوقيت الرسمي المركزي لأمريكا الشمالية", "CST China": "توقيت الصين الرسمي", "CST China DST": "توقيت الصين الصيفي", "CVST": "توقيت الرأس الأخضر الصيفي", "CVT": "توقيت الرأس الأخضر الرسمي", "CXT": "توقيت جزر الكريسماس", "ChST": "توقيت تشامورو", "ChST NMI": "توقيت جزر ماريانا الشمالية", "CuDT": "توقيت كوبا الصيفي", "CuST": "توقيت كوبا الرسمي", "DAVT": "توقيت دافيز", "DDUT": "توقيت دي مونت دو روفيل", "EASST": "توقيت جزيرة استر الصيفي", "EAST": "توقيت جزيرة استر الرسمي", "EAT": "توقيت شرق أفريقيا", "ECT": "توقيت الإكوادور", "EDT": "التوقيت الصيفي الشرقي لأمريكا الشمالية", "EGDT": "توقيت شرق غرينلاند الصيفي", "EGST": "توقيت شرق غرينلاند الرسمي", "EST": "التوقيت الرسمي الشرقي لأمريكا الشمالية", "FEET": "التوقيت الأوروبي (أكثر شرقًا)", "FJT": "توقيت فيجي الرسمي", "FJT Summer": "توقيت فيجي الصيفي", "FKST": "توقيت جزر فوكلاند الصيفي", "FKT": "توقيت جزر فوكلاند الرسمي", "FNST": "توقيت فرناندو دي نورونها الصيفي", "FNT": "توقيت فرناندو دي نورونها الرسمي", "GALT": "توقيت غلاباغوس", "GAMT": "توقيت جامبير", "GEST": "توقيت جورجيا الصيفي", "GET": "توقيت جورجيا الرسمي", "GFT": "توقيت غويانا الفرنسية", "GIT": "توقيت جزر جيلبرت", "GMT": "توقيت غرينتش", "GNSST": "GNSST", "GNST": "GNST", "GST": "توقيت الخليج", "GST Guam": "توقيت غوام", "GYT": "توقيت غيانا", "HADT": "توقيت هاواي ألوتيان الرسمي", "HAST": "توقيت هاواي ألوتيان الرسمي", "HKST": "توقيت هونغ كونغ الصيفي", "HKT": "توقيت هونغ كونغ الرسمي", "HOVST": "توقيت هوفد الصيفي", "HOVT": "توقيت هوفد الرسمي", "ICT": "توقيت الهند الصينية", "IDT": "توقيت إسرائيل الصيفي", "IOT": "توقيت المحيط الهندي", "IRKST": "توقيت إركوتسك الصيفي", "IRKT": "توقيت إركوتسك الرسمي", "IRST": "توقيت إيران الرسمي", "IRST DST": "توقيت إيران الصيفي", "IST": "توقيت الهند", "IST Israel": "توقيت إسرائيل الرسمي", "JDT": "توقيت اليابان الصيفي", "JST": "توقيت اليابان الرسمي", "KOST": "توقيت كوسرا", "KRAST": "التوقيت الصيفي لكراسنويارسك", "KRAT": "توقيت كراسنويارسك الرسمي", "KST": "توقيت كوريا الرسمي", "KST DST": "توقيت كوريا الصيفي", "LHDT": "التوقيت الصيفي للورد هاو", "LHST": "توقيت لورد هاو الرسمي", "LINT": "توقيت جزر لاين", "MAGST": "توقيت ماغادان الصيفي", "MAGT": "توقيت ماغادان الرسمي", "MART": "توقيت ماركيساس", "MAWT": "توقيت ماوسون", "MDT": "MDT", "MESZ": "توقيت وسط أوروبا الصيفي", "MEZ": "توقيت وسط أوروبا الرسمي", "MHT": "توقيت جزر مارشال", "MMT": "توقيت ميانمار", "MSD": "توقيت موسكو الصيفي", "MST": "MST", "MUST": "توقيت موريشيوس الصيفي", "MUT": "توقيت موريشيوس الرسمي", "MVT": "توقيت جزر المالديف", "MYT": "توقيت ماليزيا", "NCT": "توقيت كاليدونيا الجديدة الرسمي", "NDT": "توقيت نيوفاوندلاند الصيفي", "NDT New Caledonia": "توقيت كاليدونيا الجديدة الصيفي", "NFDT": "توقيت جزيرة نورفولك الصيفي", "NFT": "توقيت جزيرة نورفولك الرسمي", "NOVST": "توقيت نوفوسيبيرسك الصيفي", "NOVT": "توقيت نوفوسيبيرسك الرسمي", "NPT": "توقيت نيبال", "NRT": "توقيت ناورو", "NST": "توقيت نيوفاوندلاند الرسمي", "NUT": "توقيت نيوي", "NZDT": "توقيت نيوزيلندا الصيفي", "NZST": "توقيت نيوزيلندا الرسمي", "OESZ": "توقيت شرق أوروبا الصيفي", "OEZ": "توقيت شرق أوروبا الرسمي", "OMSST": "توقيت أومسك الصيفي", "OMST": "توقيت أومسك الرسمي", "PDT": "توقيت المحيط الهادي الصيفي", "PDTM": "توقيت المحيط الهادي الصيفي للمكسيك", "PETDT": "توقيت بيتروبافلوفسك-كامتشاتسكي الصيفي", "PETST": "توقيت بيتروبافلوفسك-كامتشاتسكي", "PGT": "توقيت بابوا غينيا الجديدة", "PHOT": "توقيت جزر فينكس", "PKT": "توقيت باكستان الرسمي", "PKT DST": "توقيت باكستان الصيفي", "PMDT": "توقيت سانت بيير وميكولون الصيفي", "PMST": "توقيت سانت بيير وميكولون الرسمي", "PONT": "توقيت بونابي", "PST": "توقيت المحيط الهادي الرسمي", "PST Philippine": "توقيت الفيلبين الرسمي", "PST Philippine DST": "توقيت الفيلبين الصيفي", "PST Pitcairn": "توقيت بيتكيرن", "PSTM": "توقيت المحيط الهادي الرسمي للمكسيك", "PWT": "توقيت بالاو", "PYST": "توقيت باراغواي الصيفي", "PYT": "توقيت باراغواي الرسمي", "PYT Korea": "توقيت بيونغ يانغ", "RET": "توقيت روينيون", "ROTT": "توقيت روثيرا", "SAKST": "توقيت ساخالين الصيفي", "SAKT": "توقيت ساخالين الرسمي", "SAMST": "توقيت سمارا الصيفي", "SAMT": "توقيت سمارا", "SAST": "توقيت جنوب أفريقيا", "SBT": "توقيت جزر سليمان", "SCT": "توقيت سيشل", "SGT": "توقيت سنغافورة", "SLST": "SLST", "SRT": "توقيت سورينام", "SST Samoa": "توقيت ساموا الرسمي", "SST Samoa Apia": "التوقيت الرسمي لآبيا", "SST Samoa Apia DST": "التوقيت الصيفي لأبيا", "SST Samoa DST": "توقيت ساموا الصيفي", "SYOT": "توقيت سايووا", "TAAF": "توقيت المقاطعات الفرنسية الجنوبية والأنتارتيكية", "TAHT": "توقيت تاهيتي", "TJT": "توقيت طاجكستان", "TKT": "توقيت توكيلاو", "TLT": "توقيت تيمور الشرقية", "TMST": "توقيت تركمانستان الصيفي", "TMT": "توقيت تركمانستان الرسمي", "TOST": "توقيت تونغا الصيفي", "TOT": "توقيت تونغا الرسمي", "TVT": "توقيت توفالو", "TWT": "توقيت تايبيه الرسمي", "TWT DST": "توقيت تايبيه الصيفي", "ULAST": "توقيت أولان باتور الصيفي", "ULAT": "توقيت أولان باتور الرسمي", "UYST": "توقيت أوروغواي الصيفي", "UYT": "توقيت أوروغواي الرسمي", "UZT": "توقيت أوزبكستان الرسمي", "UZT DST": "توقيت أوزبكستان الصيفي", "VET": "توقيت فنزويلا", "VLAST": "توقيت فلاديفوستوك الصيفي", "VLAT": "توقيت فلاديفوستوك الرسمي", "VOLST": "توقيت فولغوغراد الصيفي", "VOLT": "توقيت فولغوغراد الرسمي", "VOST": "توقيت فوستوك", "VUT": "توقيت فانواتو الرسمي", "VUT DST": "توقيت فانواتو الصيفي", "WAKT": "توقيت جزيرة ويك", "WARST": "توقيت غرب الأرجنتين الصيفي", "WART": "توقيت غرب الأرجنتين الرسمي", "WAST": "توقيت غرب أفريقيا", "WAT": "توقيت غرب أفريقيا", "WESZ": "توقيت غرب أوروبا الصيفي", "WEZ": "توقيت غرب أوروبا الرسمي", "WFT": "توقيت واليس و فوتونا", "WGST": "توقيت غرب غرينلاند الصيفي", "WGT": "توقيت غرب غرينلاند الرسمي", "WIB": "توقيت غرب إندونيسيا", "WIT": "توقيت شرق إندونيسيا", "WITA": "توقيت وسط إندونيسيا", "YAKST": "توقيت ياكوتسك الصيفي", "YAKT": "توقيت ياكوتسك الرسمي", "YEKST": "توقيت يكاترينبورغ الصيفي", "YEKT": "توقيت يكاترينبورغ الرسمي", "YST": "توقيت يوكون", "МСК": "توقيت موسكو الرسمي", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "توقيت غرب كازاخستان", "شىعىش قازاق ەلى": "توقيت شرق كازاخستان", "قازاق ەلى": "توقيت كازاخستان", "قىرعىزستان": "توقيت قيرغيزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "توقيت بيرو الصيفي"},
	}
}

// Locale returns the current translators string locale
func (ar *ar_001) Locale() string {
	return ar.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ar_001'
func (ar *ar_001) PluralsCardinal() []locales.PluralRule {
	return ar.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ar_001'
func (ar *ar_001) PluralsOrdinal() []locales.PluralRule {
	return ar.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ar_001'
func (ar *ar_001) PluralsRange() []locales.PluralRule {
	return ar.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ar_001'
func (ar *ar_001) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if nMod100 >= 3 && nMod100 <= 10 {
		return locales.PluralRuleFew
	} else if nMod100 >= 11 && nMod100 <= 99 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ar_001'
func (ar *ar_001) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ar_001'
func (ar *ar_001) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := ar.CardinalPluralRule(num1, v1)
	end := ar.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleZero
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleTwo {
		return locales.PluralRuleZero
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleTwo {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ar *ar_001) MonthAbbreviated(month time.Month) string {
	return ar.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ar *ar_001) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ar *ar_001) MonthNarrow(month time.Month) string {
	return ar.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ar *ar_001) MonthsNarrow() []string {
	return ar.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ar *ar_001) MonthWide(month time.Month) string {
	return ar.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ar *ar_001) MonthsWide() []string {
	return ar.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ar *ar_001) WeekdayAbbreviated(weekday time.Weekday) string {
	return ar.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ar *ar_001) WeekdaysAbbreviated() []string {
	return ar.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ar *ar_001) WeekdayNarrow(weekday time.Weekday) string {
	return ar.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ar *ar_001) WeekdaysNarrow() []string {
	return ar.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ar *ar_001) WeekdayShort(weekday time.Weekday) string {
	return ar.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ar *ar_001) WeekdaysShort() []string {
	return ar.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ar *ar_001) WeekdayWide(weekday time.Weekday) string {
	return ar.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ar *ar_001) WeekdaysWide() []string {
	return ar.daysWide
}

// Decimal returns the decimal point of number
func (ar *ar_001) Decimal() string {
	return ar.decimal
}

// Group returns the group of number
func (ar *ar_001) Group() string {
	return ar.group
}

// Group returns the minus sign of number
func (ar *ar_001) Minus() string {
	return ar.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ar_001' and handles both Whole and Real numbers based on 'v'
func (ar *ar_001) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ar.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ar.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ar_001' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ar *ar_001) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 12
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ar.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ar.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ar_001'
func (ar *ar_001) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ar.currencies[currency]
	l := len(s) + len(symbol) + 10 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ar.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ar.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(ar.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ar.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ar.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ar.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ar_001'
// in accounting notation.
func (ar *ar_001) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ar.currencies[currency]
	l := len(s) + len(symbol) + 10 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ar.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ar.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ar.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ar.currencyNegativePrefix[j])
		}

		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}

	} else {
		for j := len(ar.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ar.currencyPositivePrefix[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ar.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ar.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ar.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ar_001'
func (ar *ar_001) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ar_001'
func (ar *ar_001) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ar_001'
func (ar *ar_001) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ar.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ar_001'
func (ar *ar_001) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ar.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ar.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ar_001'
func (ar *ar_001) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ar_001'
func (ar *ar_001) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ar_001'
func (ar *ar_001) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ar_001'
func (ar *ar_001) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ar.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
