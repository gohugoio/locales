package trw_PK

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type trw_PK struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'trw_PK' locale
func New() locales.Translator {
	return &trw_PK{
		locale:          "trw_PK",
		pluralsCardinal: nil,
		pluralsOrdinal:  nil,
		pluralsRange:    nil,
		decimal:         ".",
		group:           ",",
		minus:           "-",
		percent:         "%",
		timeSeparator:   ":",
		currencies:      []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsNarrow:    []string{"", "ج", "ف", "م", "ا", "م", "ج", "ج", "ا", "س", "ا", "ن", "د"},
		monthsWide:      []string{"", "جنوری", "فروری", "مارچ", "اپریل", "مئ", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysNarrow:      []string{"ا", "د", "گ", "چ", "پ", "ش", "ل"},
		daysWide:        []string{"ایکشیمے", "دُوشیمے", "گھن آنگا", "چارشیمے", "پَئ شیمے", "شُوگار", "لَو آنگا"},
		periodsNarrow:   []string{"a", "p"},
		erasAbbreviated: []string{"", ""},
		erasNarrow:      []string{"", ""},
		erasWide:        []string{"", ""},
		timezones:       map[string]string{"ACDT": "آسٹریلین سنٹرل ڈے لائٹ ٹائم", "ACST": "آسٹریلین سنٹرل اسٹینڈرڈ ٹائم", "ACT": "ACT", "ACWDT": "آسٹریلین سنٹرل ویسٹرن ڈے لائٹ ٹائم", "ACWST": "آسٹریلین سنٹرل ویسٹرن اسٹینڈرڈ ٹائم", "ADT": "اٹلانٹک ڈے لائٹ ٹائم", "ADT Arabia": "عرب ڈے لائٹ ٹائم", "AEDT": "آسٹریلین ایسٹرن ڈے لائٹ ٹائم", "AEST": "آسٹریلین ایسٹرن اسٹینڈرڈ ٹائم", "AFT": "افغانستان سی وَخ", "AKDT": "الاسکا ڈے لائٹ ٹائم", "AKST": "الاسکا اسٹینڈرڈ ٹائم", "AMST": "امیزون سی موسم گرما سی وَخ", "AMST Armenia": "آرمینیا سی موسم گرما سی وَخ", "AMT": "ایمیزون سی معیاری وَخ", "AMT Armenia": "آرمینیا سی معیاری وَخ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ارجنٹینا سی موسم گرما سی وَخ", "ART": "ارجنٹینا سی معیاری وَخ", "AST": "اٹلانٹک اسٹینڈرڈ ٹائم", "AST Arabia": "عرب سی معیاری وَخ", "AWDT": "آسٹریلین ویسٹرن ڈے لائٹ ٹائم", "AWST": "سٹریلیا ویسٹرن اسٹینڈرڈ ٹائم", "AZST": "آذربائیجان سی موسم گرما سی وَخ", "AZT": "آذربائیجان سی معیاری وَخ", "BDT Bangladesh": "بنگلہ دیش سی موسم گرما سی وَخ", "BNT": "برونئی دارالسلام ٹائم", "BOT": "بولیویا سی وَخ", "BRST": "برازیلیا سمر ٹائم", "BRT": "برازیلیا اسٹینڈرڈ ٹائم", "BST Bangladesh": "بنگلہ دیش سی معیاری وَخ", "BT": "بھوٹان سی وَخ", "CAST": "CAST", "CAT": "وسطی افریقہ ٹائم", "CCT": "کوکوس آئلینڈز ٹائم", "CDT": "سنٹرل ڈے لائٹ ٹائم", "CHADT": "چیتھم ڈے لائٹ ٹائم", "CHAST": "چیتھم اسٹینڈرڈ ٹائم", "CHUT": "چوک ٹائم", "CKT": "کک آئلینڈز سٹینڈرڈ ٹائم", "CKT DST": "کک آئلینڈز نصف سمر ٹائم", "CLST": "چلی سی موسم گرما سی وَخ", "CLT": "چلی سی معیاری وَخ", "COST": "کولمبیا سی موسم گرما سی وَخ", "COT": "کولمبیا سی معیاری وَخ", "CST": "سنٹرل اسٹینڈرڈ ٹائم", "CST China": "چین سٹینڈرڈ ٹائم", "CST China DST": "چینی ڈے لائٹ ٹائم", "CVST": "کیپ ورڈی سمر ٹائم", "CVT": "کیپ ورڈی سٹینڈرڈ ٹائم", "CXT": "کرسمس آئلینڈ ٹائم", "ChST": "چامورو سٹینڈرڈ ٹائم", "ChST NMI": "ChST NMI", "CuDT": "کیوبا ڈے لائٹ ٹائم", "CuST": "کیوبا اسٹینڈرڈ ٹائم", "DAVT": "ڈیوس ٹائم", "DDUT": "ڈومونٹ-ڈی’ارویلے ٹائم", "EASST": "ایسٹر آئلینڈ سی موسم گرما سی وَخ", "EAST": "ایسٹر آئلینڈ سی معیاری وَخ", "EAT": "مشرقی افریقہ ٹائم", "ECT": "ایکواڈور سی وَخ", "EDT": "ایسٹرن ڈے لائٹ ٹائم", "EGDT": "مشرقی گرین لینڈ سی موسم گرما سی وَخ", "EGST": "مشرقی گرین لینڈ اسٹینڈرڈ ٹائم", "EST": "ایسٹرن اسٹینڈرڈ ٹائم", "FEET": "بعید مشرقی یورپی وَخ", "FJT": "فجی سٹینڈرڈ ٹائم", "FJT Summer": "فجی سمر ٹائم", "FKST": "فاک لینڈ آئلینڈز سی موسم گرما سی وَخ", "FKT": "فاک لینڈ آئلینڈز سی معیاری وَخ", "FNST": "فرنانڈو ڈی نورونہا سمر ٹائم", "FNT": "فرنانڈو ڈی نورنہا سی معیاری وَخ", "GALT": "گالاپاگوز سی وَخ", "GAMT": "گیمبیئر ٹائم", "GEST": "جارجیا سی موسم گرما سی وَخ", "GET": "جارجیا سی معیاری وَخ", "GFT": "فرینچ گیانا سی وَخ", "GIT": "جلبرٹ آئلینڈز ٹائم", "GMT": "گرین وچ سی اصل وَخ", "GNSST": "GNSST", "GNST": "GNST", "GST": "خلیج سی معیاری وَخ", "GST Guam": "GST Guam", "GYT": "گیانا سی وَخ", "HADT": "ہوائی الیوٹیئن اسٹینڈرڈ ٹائم", "HAST": "ہوائی الیوٹیئن اسٹینڈرڈ ٹائم", "HKST": "ہانگ سینگ سمر ٹائم", "HKT": "ہانگ سینگ سٹینڈرڈ ٹائم", "HOVST": "ہووڈ سمر ٹائم", "HOVT": "ہووڈ سٹینڈرڈ ٹائم", "ICT": "ہند چین ٹائم", "IDT": "اسرائیل ڈے لائٹ ٹائم", "IOT": "بحر ہند ٹائم", "IRKST": "ارکتسک سمر ٹائم", "IRKT": "ارکتسک سٹینڈرڈ ٹائم", "IRST": "ایران سی معیاری وَخ", "IRST DST": "ایران ڈے لائٹ ٹائم", "IST": "ہندوستان سی معیاری وَخ", "IST Israel": "اسرائیل سی معیاری وَخ", "JDT": "جاپان ڈے لائٹ ٹائم", "JST": "جاپان سٹینڈرڈ ٹائم", "KOST": "کوسرے ٹائم", "KRAST": "کریسنویارسک سمر ٹائم", "KRAT": "کرسنویارسک سٹینڈرڈ ٹائم", "KST": "کوریا سٹینڈرڈ ٹائم", "KST DST": "کوریا ڈے لائٹ ٹائم", "LHDT": "لارڈ ہووے ڈے لائٹ ٹائم", "LHST": "لارڈ ہووے اسٹینڈرڈ ٹائم", "LINT": "لائن آئلینڈز ٹائم", "MAGST": "میگیدن سمر ٹائم", "MAGT": "مگادان اسٹینڈرڈ ٹائم", "MART": "مارکیسس ٹائم", "MAWT": "ماؤسن ٹائم", "MDT": "MDT", "MESZ": "وسطی یورپ سی موسم گرما سی وَخ", "MEZ": "وسطی یورپ سی معیاری وَخ", "MHT": "مارشل آئلینڈز ٹائم", "MMT": "میانمار ٹائم", "MSD": "ماسکو سمر ٹائم", "MST": "MST", "MUST": "ماریشس سمر ٹائم", "MUT": "ماریشس سٹینڈرڈ ٹائم", "MVT": "مالدیپ سی وَخ", "MYT": "ملیشیا ٹائم", "NCT": "نیو کیلیڈونیا سٹینڈرڈ ٹائم", "NDT": "نیو فاؤنڈ لینڈ ڈے لائٹ ٹائم", "NDT New Caledonia": "نیو کیلیڈونیا سمر ٹائم", "NFDT": "نارفوک آئلینڈ سی موسم گرما سی وَخ", "NFT": "نارفوک آئلینڈ سی معیاری وَخ", "NOVST": "نوووسیبرسک سمر ٹائم", "NOVT": "نوووسیبرسک سٹینڈرڈ ٹائم", "NPT": "نیپال سی وَخ", "NRT": "ناؤرو ٹائم", "NST": "نیو فاؤنڈ لینڈ اسٹینڈرڈ ٹائم", "NUT": "نیئو ٹائم", "NZDT": "نیوزی لینڈ ڈے لائٹ ٹائم", "NZST": "نیوزی لینڈ سی معیاری وَخ", "OESZ": "مشرقی یورپ سی موسم گرما سی وَخ", "OEZ": "مشرقی یورپ سی معیاری وَخ", "OMSST": "اومسک سمر ٹائم", "OMST": "اومسک سٹینڈرڈ ٹائم", "PDT": "پیسفک ڈے لائٹ ٹائم", "PDTM": "میکسیکن پیسفک ڈے لائٹ ٹائم", "PETDT": "PETDT", "PETST": "PETST", "PGT": "پاپوآ نیو گنی ٹائم", "PHOT": "فینکس آئلینڈز ٹائم", "PKT": "پاکستان سی معیاری وَخ", "PKT DST": "پاکستان سی موسم گرما سی وَخ", "PMDT": "سینٹ پیئر آں مکلیئون ڈے لائٹ ٹائم", "PMST": "سینٹ پیئر آں مکلیئون اسٹینڈرڈ ٹائم", "PONT": "پوناپے ٹائم", "PST": "پیسفک اسٹینڈرڈ ٹائم", "PST Philippine": "فلپائن سٹینڈرڈ ٹائم", "PST Philippine DST": "فلپائن سمر ٹائم", "PST Pitcairn": "پٹکائرن ٹائم", "PSTM": "میکسیکن پیسفک اسٹینڈرڈ ٹائم", "PWT": "پلاؤ ٹائم", "PYST": "پیراگوئے سی موسم گرما سی وَخ", "PYT": "پیراگوئے سی معیاری وَخ", "PYT Korea": "پیانگ یانگ وَخ", "RET": "ری یونین ٹائم", "ROTT": "روتھیرا سی وَخ", "SAKST": "سخالین سمر ٹائم", "SAKT": "سخالین سٹینڈرڈ ٹائم", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "جنوبی افریقہ سٹینڈرڈ ٹائم", "SBT": "سولمن آئلینڈز ٹائم", "SCT": "سیشلیز ٹائم", "SGT": "سنگاپور سٹینڈرڈ ٹائم", "SLST": "SLST", "SRT": "سورینام سی وَخ", "SST Samoa": "ساموآ سٹینڈرڈ ٹائم", "SST Samoa Apia": "ایپیا سٹینڈرڈ ٹائم", "SST Samoa Apia DST": "ایپیا ڈے لائٹ ٹائم", "SST Samoa DST": "ساموآ ڈے لائٹ ٹائم", "SYOT": "سیووا ٹائم", "TAAF": "فرینچ جنوبی آں انٹارکٹک ٹائم", "TAHT": "تاہیتی ٹائم", "TJT": "تاجکستان سی وَخ", "TKT": "ٹوکیلاؤ ٹائم", "TLT": "مشرقی تیمور ٹائم", "TMST": "ترکمانستان سی موسم گرما سی وَخ", "TMT": "ترکمانستان سی معیاری وَخ", "TOST": "ٹونگا سمر ٹائم", "TOT": "ٹونگا سٹینڈرڈ ٹائم", "TVT": "ٹوالو ٹائم", "TWT": "تائی پیئی اسٹینڈرڈ ٹائم", "TWT DST": "تئی پیئی ڈے لائٹ ٹائم", "ULAST": "یولان بیتور سمر ٹائم", "ULAT": "یولان بیتور سٹینڈرڈ ٹائم", "UYST": "یوروگوئے سی موسم گرما سی وَخ", "UYT": "یوروگوئے سی معیاری وَخ", "UZT": "ازبکستان سی معیاری وَخ", "UZT DST": "ازبکستان سی موسم گرما سی وَخ", "VET": "وینزوئیلا سی وَخ", "VLAST": "ولادی ووستک سمر ٹائم", "VLAT": "ولادی ووستک سٹینڈرڈ ٹائم", "VOLST": "وولگوگراد سمر ٹائم", "VOLT": "وولگوگراد اسٹینڈرڈ ٹائم", "VOST": "ووسٹاک سی وَخ", "VUT": "وانوآٹو سٹینڈرڈ ٹائم", "VUT DST": "وانوآٹو سمر ٹائم", "WAKT": "ویک آئلینڈ ٹائم", "WARST": "مغربی ارجنٹینا سی موسم گرما سی وَخ", "WART": "مغربی ارجنٹینا سی معیاری وَخ", "WAST": "مغربی افریقہ ٹائم", "WAT": "مغربی افریقہ ٹائم", "WESZ": "مغربی یورپ سی موسم گرما سی وَخ", "WEZ": "مغربی یورپ سی معیاری وَخ", "WFT": "والیز اور فٹونا ٹائم", "WGST": "مغربی گرین لینڈ سی موسم گرما سی وَخ", "WGT": "مغربی گرین لینڈ اسٹینڈرڈ ٹائم", "WIB": "مغربی انڈونیشیا ٹائم", "WIT": "مشرقی انڈونیشیا ٹائم", "WITA": "وسطی انڈونیشیا ٹائم", "YAKST": "یکوتسک سمر ٹائم", "YAKT": "یکوتسک اسٹینڈرڈ ٹائم", "YEKST": "یکاٹیرِنبرگ سمر ٹائم", "YEKT": "یکاٹیرِنبرگ اسٹینڈرڈ ٹائم", "YST": "یوکون ٹائم", "МСК": "ماسکو اسٹینڈرڈ ٹائم", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "مغربی قزاخستان سی وَخ", "شىعىش قازاق ەلى": "مشرقی قزاخستان سی وَخ", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "کرغستان سی وَخ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "پیرو سی موسم گرما سی وَخ"},
	}
}

// Locale returns the current translators string locale
func (trw *trw_PK) Locale() string {
	return trw.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'trw_PK'
func (trw *trw_PK) PluralsCardinal() []locales.PluralRule {
	return trw.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'trw_PK'
func (trw *trw_PK) PluralsOrdinal() []locales.PluralRule {
	return trw.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'trw_PK'
func (trw *trw_PK) PluralsRange() []locales.PluralRule {
	return trw.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'trw_PK'
func (trw *trw_PK) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'trw_PK'
func (trw *trw_PK) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'trw_PK'
func (trw *trw_PK) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (trw *trw_PK) MonthAbbreviated(month time.Month) string {
	return trw.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (trw *trw_PK) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (trw *trw_PK) MonthNarrow(month time.Month) string {
	return trw.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (trw *trw_PK) MonthsNarrow() []string {
	return trw.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (trw *trw_PK) MonthWide(month time.Month) string {
	return trw.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (trw *trw_PK) MonthsWide() []string {
	return trw.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (trw *trw_PK) WeekdayAbbreviated(weekday time.Weekday) string {
	return trw.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (trw *trw_PK) WeekdaysAbbreviated() []string {
	return trw.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (trw *trw_PK) WeekdayNarrow(weekday time.Weekday) string {
	return trw.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (trw *trw_PK) WeekdaysNarrow() []string {
	return trw.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (trw *trw_PK) WeekdayShort(weekday time.Weekday) string {
	return trw.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (trw *trw_PK) WeekdaysShort() []string {
	return trw.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (trw *trw_PK) WeekdayWide(weekday time.Weekday) string {
	return trw.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (trw *trw_PK) WeekdaysWide() []string {
	return trw.daysWide
}

// Decimal returns the decimal point of number
func (trw *trw_PK) Decimal() string {
	return trw.decimal
}

// Group returns the group of number
func (trw *trw_PK) Group() string {
	return trw.group
}

// Group returns the minus sign of number
func (trw *trw_PK) Minus() string {
	return trw.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'trw_PK' and handles both Whole and Real numbers based on 'v'
func (trw *trw_PK) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, trw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, trw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, trw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'trw_PK' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (trw *trw_PK) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, trw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, trw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, trw.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'trw_PK'
func (trw *trw_PK) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := trw.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'trw_PK'
// in accounting notation.
func (trw *trw_PK) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := trw.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, trw.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, trw.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, trw.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, trw.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, trw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, trw.periodsAbbreviated[0]...)
	} else {
		b = append(b, trw.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, trw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, trw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, trw.periodsAbbreviated[0]...)
	} else {
		b = append(b, trw.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, trw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, trw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, trw.periodsAbbreviated[0]...)
	} else {
		b = append(b, trw.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'trw_PK'
func (trw *trw_PK) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, trw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, trw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, trw.periodsAbbreviated[0]...)
	} else {
		b = append(b, trw.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := trw.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
