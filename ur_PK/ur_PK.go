package ur_PK

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ur_PK struct {
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ur_PK' locale
func New() locales.Translator {
	return &ur_PK{
		locale:                 "ur_PK",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "‎-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysWide:               []string{"اتوار", "پیر", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"قبل مسیح", "عیسوی"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"عام دور سے قبل", "عام دور"},
		timezones:              map[string]string{"ACDT": "آسٹریلین سنٹرل ڈے لائٹ ٹائم", "ACST": "ACST", "ACT": "ACT", "ACWDT": "آسٹریلین سنٹرل ویسٹرن ڈے لائٹ ٹائم", "ACWST": "آسٹریلین سنٹرل ویسٹرن اسٹینڈرڈ ٹائم", "ADT": "اٹلانٹک ڈے لائٹ ٹائم", "ADT Arabia": "عرب ڈے لائٹ ٹائم", "AEDT": "آسٹریلین ایسٹرن ڈے لائٹ ٹائم", "AEST": "آسٹریلین ایسٹرن اسٹینڈرڈ ٹائم", "AFT": "افغانستان کا وقت", "AKDT": "الاسکا ڈے لائٹ ٹائم", "AKST": "الاسکا اسٹینڈرڈ ٹائم", "AMST": "امیزون کا موسم گرما کا وقت", "AMST Armenia": "آرمینیا کا موسم گرما کا وقت", "AMT": "ایمیزون کا معیاری وقت", "AMT Armenia": "آرمینیا کا معیاری وقت", "ANAST": "انیدر سمر ٹائم", "ANAT": "انیدر اسٹینڈرڈ ٹائم", "ARST": "ارجنٹینا کا موسم گرما کا وقت", "ART": "ارجنٹینا کا معیاری وقت", "AST": "اٹلانٹک اسٹینڈرڈ ٹائم", "AST Arabia": "عرب کا معیاری وقت", "AWDT": "آسٹریلین ویسٹرن ڈے لائٹ ٹائم", "AWST": "آسٹریلیا ویسٹرن اسٹینڈرڈ ٹائم", "AZST": "آذربائیجان کا موسم گرما کا وقت", "AZT": "آذربائیجان کا معیاری وقت", "BDT Bangladesh": "بنگلہ دیش کا موسم گرما کا وقت", "BNT": "برونئی دارالسلام ٹائم", "BOT": "بولیویا کا وقت", "BRST": "برازیلیا سمر ٹائم", "BRT": "برازیلیا اسٹینڈرڈ ٹائم", "BST Bangladesh": "بنگلہ دیش کا معیاری وقت", "BT": "بھوٹان کا وقت", "CAST": "CAST", "CAT": "وسطی افریقہ ٹائم", "CCT": "کوکوس آئلینڈز ٹائم", "CDT": "سنٹرل ڈے لائٹ ٹائم", "CHADT": "چیتھم ڈے لائٹ ٹائم", "CHAST": "چیتھم اسٹینڈرڈ ٹائم", "CHUT": "چوک ٹائم", "CKT": "کک آئلینڈز سٹینڈرڈ ٹائم", "CKT DST": "کک آئلینڈز نصف سمر ٹائم", "CLST": "چلی کا موسم گرما کا وقت", "CLT": "چلی کا معیاری وقت", "COST": "کولمبیا کا موسم گرما کا وقت", "COT": "کولمبیا کا معیاری وقت", "CST": "سنٹرل اسٹینڈرڈ ٹائم", "CST China": "چین سٹینڈرڈ ٹائم", "CST China DST": "چینی ڈے لائٹ ٹائم", "CVST": "کیپ ورڈی سمر ٹائم", "CVT": "کیپ ورڈی سٹینڈرڈ ٹائم", "CXT": "کرسمس آئلینڈ ٹائم", "ChST": "چامورو سٹینڈرڈ ٹائم", "ChST NMI": "ChST NMI", "CuDT": "کیوبا ڈے لائٹ ٹائم", "CuST": "کیوبا اسٹینڈرڈ ٹائم", "DAVT": "ڈیوس ٹائم", "DDUT": "ڈومونٹ-ڈی’ارویلے ٹائم", "EASST": "ایسٹر آئلینڈ کا موسم گرما کا وقت", "EAST": "ایسٹر آئلینڈ کا معیاری وقت", "EAT": "مشرقی افریقہ ٹائم", "ECT": "ایکواڈور کا وقت", "EDT": "ایسٹرن ڈے لائٹ ٹائم", "EGDT": "مشرقی گرین لینڈ کا موسم گرما کا وقت", "EGST": "مشرقی گرین لینڈ اسٹینڈرڈ ٹائم", "EST": "ایسٹرن اسٹینڈرڈ ٹائم", "FEET": "بعید مشرقی یورپی وقت", "FJT": "فجی سٹینڈرڈ ٹائم", "FJT Summer": "فجی سمر ٹائم", "FKST": "فاک لینڈ آئلینڈز کا موسم گرما کا وقت", "FKT": "فاک لینڈ آئلینڈز کا معیاری وقت", "FNST": "فرنانڈو ڈی نورونہا سمر ٹائم", "FNT": "فرنانڈو ڈی نورنہا کا معیاری وقت", "GALT": "گالاپاگوز کا وقت", "GAMT": "گیمبیئر ٹائم", "GEST": "جارجیا کا موسم گرما کا وقت", "GET": "جارجیا کا معیاری وقت", "GFT": "فرینچ گیانا کا وقت", "GIT": "جلبرٹ آئلینڈز ٹائم", "GMT": "گرین وچ کا اصل وقت", "GNSST": "GNSST", "GNST": "GNST", "GST": "خلیج کا معیاری وقت", "GST Guam": "GST Guam", "GYT": "گیانا کا وقت", "HADT": "ہوائی الیوٹیئن ڈے لائٹ ٹائم", "HAST": "ہوائی الیوٹیئن اسٹینڈرڈ ٹائم", "HKST": "ہانگ کانگ سمر ٹائم", "HKT": "ہانگ کانگ سٹینڈرڈ ٹائم", "HOVST": "ہووڈ سمر ٹائم", "HOVT": "ہووڈ سٹینڈرڈ ٹائم", "ICT": "ہند چین ٹائم", "IDT": "اسرائیل ڈے لائٹ ٹائم", "IOT": "بحر ہند ٹائم", "IRKST": "ارکتسک سمر ٹائم", "IRKT": "ارکتسک سٹینڈرڈ ٹائم", "IRST": "ایران کا معیاری وقت", "IRST DST": "ایران ڈے لائٹ ٹائم", "IST": "ہندوستان کا معیاری وقت", "IST Israel": "اسرائیل کا معیاری وقت", "JDT": "جاپان ڈے لائٹ ٹائم", "JST": "جاپان سٹینڈرڈ ٹائم", "KOST": "کوسرے ٹائم", "KRAST": "کریسنویارسک سمر ٹائم", "KRAT": "کرسنویارسک سٹینڈرڈ ٹائم", "KST": "کوریا سٹینڈرڈ ٹائم", "KST DST": "کوریا ڈے لائٹ ٹائم", "LHDT": "لارڈ ہووے ڈے لائٹ ٹائم", "LHST": "لارڈ ہووے اسٹینڈرڈ ٹائم", "LINT": "لائن آئلینڈز ٹائم", "MAGST": "میگیدن سمر ٹائم", "MAGT": "مگادان اسٹینڈرڈ ٹائم", "MART": "مارکیسس ٹائم", "MAWT": "ماؤسن ٹائم", "MDT": "MDT", "MESZ": "وسطی یورپ کا موسم گرما کا وقت", "MEZ": "وسطی یورپ کا معیاری وقت", "MHT": "مارشل آئلینڈز ٹائم", "MMT": "میانمار ٹائم", "MSD": "ماسکو سمر ٹائم", "MST": "MST", "MUST": "ماریشس سمر ٹائم", "MUT": "ماریشس سٹینڈرڈ ٹائم", "MVT": "مالدیپ کا وقت", "MYT": "ملیشیا ٹائم", "NCT": "نیو کیلیڈونیا سٹینڈرڈ ٹائم", "NDT": "نیو فاؤنڈ لینڈ ڈے لائٹ ٹائم", "NDT New Caledonia": "نیو کیلیڈونیا سمر ٹائم", "NFDT": "نارفوک آئلینڈ ڈے لائٹ وقت", "NFT": "نارفوک آئلینڈ کا معیاری وقت", "NOVST": "نوووسیبرسک سمر ٹائم", "NOVT": "نوووسیبرسک سٹینڈرڈ ٹائم", "NPT": "نیپال کا وقت", "NRT": "ناؤرو ٹائم", "NST": "نیو فاؤنڈ لینڈ اسٹینڈرڈ ٹائم", "NUT": "نیئو ٹائم", "NZDT": "نیوزی لینڈ ڈے لائٹ ٹائم", "NZST": "نیوزی لینڈ کا معیاری وقت", "OESZ": "مشرقی یورپ کا موسم گرما کا وقت", "OEZ": "مشرقی یورپ کا معیاری وقت", "OMSST": "اومسک سمر ٹائم", "OMST": "اومسک سٹینڈرڈ ٹائم", "PDT": "پیسفک ڈے لائٹ ٹائم", "PDTM": "میکسیکن پیسفک ڈے لائٹ ٹائم", "PETDT": "پیٹروپاؤلووسک-کیمچسکی سمر ٹائم", "PETST": "پیٹروپاؤلووسک-کیمچسکی اسٹینڈرڈ ٹائم", "PGT": "پاپوآ نیو گنی ٹائم", "PHOT": "فینکس آئلینڈز ٹائم", "PKT": "پاکستان کا معیاری وقت", "PKT DST": "پاکستان کا موسم گرما کا وقت", "PMDT": "سینٹ پیئر اور مکلیئون ڈے لائٹ ٹائم", "PMST": "سینٹ پیئر اور مکلیئون اسٹینڈرڈ ٹائم", "PONT": "پوناپے ٹائم", "PST": "پیسفک اسٹینڈرڈ ٹائم", "PST Philippine": "فلپائن سٹینڈرڈ ٹائم", "PST Philippine DST": "فلپائن سمر ٹائم", "PST Pitcairn": "پٹکائرن ٹائم", "PSTM": "میکسیکن پیسفک اسٹینڈرڈ ٹائم", "PWT": "پلاؤ ٹائم", "PYST": "پیراگوئے کا موسم گرما کا وقت", "PYT": "پیراگوئے کا معیاری وقت", "PYT Korea": "پیانگ یانگ وقت", "RET": "ری یونین ٹائم", "ROTT": "روتھیرا کا وقت", "SAKST": "سخالین سمر ٹائم", "SAKT": "سخالین سٹینڈرڈ ٹائم", "SAMST": "سمارا سمر ٹائم", "SAMT": "سمارا اسٹینڈرڈ ٹائم", "SAST": "جنوبی افریقہ سٹینڈرڈ ٹائم", "SBT": "سولمن آئلینڈز ٹائم", "SCT": "سیشلیز ٹائم", "SGT": "سنگاپور سٹینڈرڈ ٹائم", "SLST": "SLST", "SRT": "سورینام کا وقت", "SST Samoa": "ساموآ سٹینڈرڈ ٹائم", "SST Samoa Apia": "ایپیا سٹینڈرڈ ٹائم", "SST Samoa Apia DST": "ایپیا ڈے لائٹ ٹائم", "SST Samoa DST": "ساموآ ڈے لائٹ ٹائم", "SYOT": "سیووا ٹائم", "TAAF": "فرینچ جنوبی اور انٹارکٹک ٹائم", "TAHT": "تاہیتی ٹائم", "TJT": "تاجکستان کا وقت", "TKT": "ٹوکیلاؤ ٹائم", "TLT": "مشرقی تیمور ٹائم", "TMST": "ترکمانستان کا موسم گرما کا وقت", "TMT": "ترکمانستان کا معیاری وقت", "TOST": "ٹونگا سمر ٹائم", "TOT": "ٹونگا سٹینڈرڈ ٹائم", "TVT": "ٹوالو ٹائم", "TWT": "تائی پیئی اسٹینڈرڈ ٹائم", "TWT DST": "تئی پیئی ڈے لائٹ ٹائم", "ULAST": "یولان بیتور سمر ٹائم", "ULAT": "یولان بیتور سٹینڈرڈ ٹائم", "UYST": "یوروگوئے کا موسم گرما کا وقت", "UYT": "یوروگوئے کا معیاری وقت", "UZT": "ازبکستان کا معیاری وقت", "UZT DST": "ازبکستان کا موسم گرما کا وقت", "VET": "وینزوئیلا کا وقت", "VLAST": "ولادی ووستک سمر ٹائم", "VLAT": "ولادی ووستک سٹینڈرڈ ٹائم", "VOLST": "وولگوگراد سمر ٹائم", "VOLT": "وولگوگراد اسٹینڈرڈ ٹائم", "VOST": "ووسٹاک کا وقت", "VUT": "وانوآٹو سٹینڈرڈ ٹائم", "VUT DST": "وانوآٹو سمر ٹائم", "WAKT": "ویک آئلینڈ ٹائم", "WARST": "مغربی ارجنٹینا کا موسم گرما کا وقت", "WART": "مغربی ارجنٹینا کا معیاری وقت", "WAST": "مغربی افریقہ ٹائم", "WAT": "مغربی افریقہ ٹائم", "WESZ": "مغربی یورپ کا موسم گرما کا وقت", "WEZ": "مغربی یورپ کا معیاری وقت", "WFT": "والیز اور فٹونا ٹائم", "WGST": "مغربی گرین لینڈ کا موسم گرما کا وقت", "WGT": "مغربی گرین لینڈ اسٹینڈرڈ ٹائم", "WIB": "مغربی انڈونیشیا ٹائم", "WIT": "مشرقی انڈونیشیا ٹائم", "WITA": "وسطی انڈونیشیا ٹائم", "YAKST": "یکوتسک سمر ٹائم", "YAKT": "یکوتسک اسٹینڈرڈ ٹائم", "YEKST": "یکاٹیرِنبرگ سمر ٹائم", "YEKT": "یکاٹیرِنبرگ اسٹینڈرڈ ٹائم", "YST": "یوکون ٹائم", "МСК": "ماسکو اسٹینڈرڈ ٹائم", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "مغربی قزاخستان کا وقت", "شىعىش قازاق ەلى": "مشرقی قزاخستان کا وقت", "قازاق ەلى": "قازقستان کا وقت", "قىرعىزستان": "کرغستان کا وقت", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ازوریس کا موسم گرما کا وقت"},
	}
}

// Locale returns the current translators string locale
func (ur *ur_PK) Locale() string {
	return ur.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ur_PK'
func (ur *ur_PK) PluralsCardinal() []locales.PluralRule {
	return ur.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ur_PK'
func (ur *ur_PK) PluralsOrdinal() []locales.PluralRule {
	return ur.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ur_PK'
func (ur *ur_PK) PluralsRange() []locales.PluralRule {
	return ur.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ur_PK'
func (ur *ur_PK) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ur_PK'
func (ur *ur_PK) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ur_PK'
func (ur *ur_PK) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ur *ur_PK) MonthAbbreviated(month time.Month) string {
	return ur.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ur *ur_PK) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ur *ur_PK) MonthNarrow(month time.Month) string {
	return ur.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ur *ur_PK) MonthsNarrow() []string {
	return ur.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ur *ur_PK) MonthWide(month time.Month) string {
	return ur.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ur *ur_PK) MonthsWide() []string {
	return ur.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ur *ur_PK) WeekdayAbbreviated(weekday time.Weekday) string {
	return ur.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ur *ur_PK) WeekdaysAbbreviated() []string {
	return ur.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ur *ur_PK) WeekdayNarrow(weekday time.Weekday) string {
	return ur.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ur *ur_PK) WeekdaysNarrow() []string {
	return ur.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ur *ur_PK) WeekdayShort(weekday time.Weekday) string {
	return ur.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ur *ur_PK) WeekdaysShort() []string {
	return ur.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ur *ur_PK) WeekdayWide(weekday time.Weekday) string {
	return ur.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ur *ur_PK) WeekdaysWide() []string {
	return ur.daysWide
}

// Decimal returns the decimal point of number
func (ur *ur_PK) Decimal() string {
	return ur.decimal
}

// Group returns the group of number
func (ur *ur_PK) Group() string {
	return ur.group
}

// Group returns the minus sign of number
func (ur *ur_PK) Minus() string {
	return ur.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ur_PK' and handles both Whole and Real numbers based on 'v'
func (ur *ur_PK) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ur.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ur_PK' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ur *ur_PK) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ur.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ur_PK'
func (ur *ur_PK) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ur.currencies[currency]
	l := len(s) + len(symbol) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ur.group[0])
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
		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ur.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ur_PK'
// in accounting notation.
func (ur *ur_PK) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ur.currencies[currency]
	l := len(s) + len(symbol) + 7 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ur.group[0])
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

		b = append(b, ur.currencyNegativePrefix[0])

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
			b = append(b, ur.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ur.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ur.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ur.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ur.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ur.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ur.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ur.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ur_PK'
func (ur *ur_PK) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ur.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ur.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
