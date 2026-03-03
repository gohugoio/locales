package ks_Deva_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ks_Deva_IN struct {
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

// New returns a new instance of translator for the 'ks_Deva_IN' locale
func New() locales.Translator {
	return &ks_Deva_IN{
		locale:            "ks_Deva_IN",
		pluralsCardinal:   []locales.PluralRule{2, 6},
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           ".",
		group:             ",",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "جنؤری", "فرؤری", "مارٕچ", "اپریل", "مئی", "جوٗن", "جُلَے", "اگست", "ستمبر", "اکتوٗبر", "نومبر", "دسمبر"},
		monthsNarrow:      []string{"", "ج", "ف", "م", "ا", "م", "ج", "ج", "ا", "س", "س", "ا", "ن"},
		monthsWide:        []string{"", "جنؤری", "فرؤری", "مارٕچ", "اپریل", "مئی", "جوٗن", "جُلَے", "اگست", "ستمبر", "اکتوٗبر", "نومبر", "دَسَمبَر"},
		daysAbbreviated:   []string{"آتھوار", "ژٔندٕروار", "بۆموار", "بودوار", "برؠسوار", "جُمہ", "بٹوار"},
		daysNarrow:        []string{"ا", "ژ", "ب", "ب", "ب", "ج", "ب"},
		daysWide:          []string{"اَتھوار", "ژٔندرٕروار", "بۆموار", "بودوار", "برؠسوار", "جُمہ", "بٹوار"},
		timezones:         map[string]string{"ACDT": "آسٹریلِیَن مرکزی ڈےلایِٔٹ ٹایِم", "ACST": "آسٹریلِیَن مرکزی سٹینڑاڑ ٹایِم", "ACT": "اؠکرے سٹینڑاڑ ٹایِم", "ACWDT": "آسٹریلِیَن مرکزی مغربی ڈےلایِٔٹ ٹایِم", "ACWST": "آسٹریلِیَن مرکزی مغربی سٹینڑاڑ ٹایِم", "ADT": "اؠٹلانٹِک ڈےلایِٔٹ ٹایِم", "ADT Arabia": "ارؠبِیَن ڈےلایِٔٹ ٹایِم", "AEDT": "آسٹریلِیَن مشرقی ڈےلایِٔٹ ٹایِم", "AEST": "آسٹریلِیَن مشرقی سٹینڑاڑ ٹایِم", "AFT": "افغانِستان ٹایِم", "AKDT": "اؠلاسکا ڈےلایِٔٹ ٹایِم", "AKST": "اؠلاسکا سٹینڑاڑ ٹایِم", "AMST": "اؠمَزَن سَمَر ٹایِم", "AMST Armenia": "ارمیٖنِیا سَمَر ٹایِم", "AMT": "اؠمَزَن سٹینڑاڑ ٹایِم", "AMT Armenia": "ارمیٖنِیا سٹینڈرڈ ٹایِم", "ANAST": "اؠنڑیٖر سَمَر ٹایِم", "ANAT": "اؠنَڑیٖر سٹینڑاڑ ٹایِم", "ARST": "ارجؠنٹیٖنا سَمَر ٹایِم", "ART": "ارجؠنٹیٖنا سٹینڑاڑ ٹایِم", "AST": "اؠٹلانٹِک سٹینڑاڑ ٹایِم", "AST Arabia": "ارؠبِیَن سٹینڈرڈ ٹایِم", "AWDT": "آسٹریلِیَن مغرِبیٖ ڈےلایٔٹ ٹایِم", "AWST": "آسٹریلِیَن مغرِبی سٹینڑاڑ ٹایِم", "AZST": "ازربائیجان سَمَر ٹائم", "AZT": "ازربائیجان سٹینڈرڈ ٹائم", "BDT Bangladesh": "بَنگلادیش سَمَر ٹایِم", "BNT": "بروٗنَے دَروٗسَلَم ٹایِم", "BOT": "بولِوِیا ٹایِم", "BRST": "برؠسِلِیا سَمَر ٹایِم", "BRT": "برؠسِلِیا سٹینڑاڑ ٹایِم", "BST Bangladesh": "بَنگلادیش سٹینڑاڑ ٹایِم", "BT": "بوٗٹان ٹایِم", "CAST": "CAST", "CAT": "مرکزی افریٖقا ٹایِم", "CCT": "کوکوز اَیلینڑز ٹایِم", "CDT": "مرکزی ڈےلایِٔٹ ٹایِم", "CHADT": "چؠتھَم سَمَر ٹایِم", "CHAST": "کؠتھَم سٹینڑاڑ ٹایِم", "CHUT": "ٹٔرک ٹایِم", "CKT": "کُک اَیلینڑز سٹینڑاڑ ٹایِم", "CKT DST": "کُک اَیلینڑز حاف سَمَر ٹایِم", "CLST": "چِلی سَمَر ٹایِم", "CLT": "چِلی سٹینڑاڑ ٹایِم", "COST": "کولومبِیا سَمَر ٹایِم", "COT": "کولومبِیا سٹینڑاڑ ٹایِم", "CST": "مرکزی سٹینڑاڑ ٹایِم", "CST China": "چَینا سٹینڈرڈ ٹایِم", "CST China DST": "چَینا ڈےلایِٔٹ ٹایِم", "CVST": "کیپ سَمَر ٹایِم", "CVT": "کیپ ؤرڑو سٹینڈرڈ ٹایِم", "CXT": "کرسمَس ٹایِم", "ChST": "کؠمورو سٹینڑاڑ ٹایِم", "ChST NMI": "شُمٲلی مَرِیانا ٹایِم", "CuDT": "کیوٗبا ڈےلایِٔٹ ٹایِم", "CuST": "کیوٗبا سٹینڑاڑ ٹایِم", "DAVT": "ڑیوِس ٹایِم", "DDUT": "ڑمانٹ ڈی اُرویٖل ٹایِم", "EASST": "ایٖسٹَر جزیرٕ سَمَر ٹایِم", "EAST": "ایٖسٹَر جزیرٕ سٹینڈرڈ ٹایِم", "EAT": "مشرقی افریٖقا ٹایِم", "ECT": "اِکویڑَر ٹایِم", "EDT": "مشرقی ڈےلایِٔٹ ٹایِم", "EGDT": "مشرِقی گریٖن لینڈُک سَمَر ٹایِم", "EGST": "مشرِقی گریٖن لینڈُک سٹینڑاڑ ٹایِم", "EST": "مشرقی سٹینڑاڑ ٹایِم", "FEET": "مزید مشرقی یورپی ٹائم", "FJT": "فیٖجی سٹینڑاڑ ٹایِم", "FJT Summer": "فیٖجی سَمَر ٹایِم", "FKST": "فالک لینڈ جزیرٕ سَمَر ٹائم", "FKT": "فالک لینڈ جزیرٕ سٹینڈرڈ ٹائم", "FNST": "فرنینڈو ڈی نورونہا سَمَر ٹائم", "FNT": "فرنینڈو ڈی نورونہا سٹینڈرڈ ٹائم", "GALT": "گؠلؠپیگوز ٹایِم", "GAMT": "گؠمبِیَر ٹایِم", "GEST": "جورجِیاہُک سَمَر ٹایِم", "GET": "جورجِیاہُک سٹینڈرڈ ٹایِم", "GFT": "فرؠنچ گیوٗؠنا ٹایِم", "GIT": "گِلبٲٹ ججیٖرُک ٹایِم", "GMT": "گریٖن وِچ میٖن ٹایِم", "GNSST": "GNSST", "GNST": "GNST", "GST": "شُمٲلی جورجِیا ٹایِم", "GST Guam": "گُوؠم ٹایِم", "GYT": "گُیَنا ٹایِم", "HADT": "حَواے اؠلیوٗٹِیَن سٹینڑاڑ ٹایِم", "HAST": "حَواے اؠلیوٗٹِیَن سٹینڑاڑ ٹایِم", "HKST": "ہانگ کانگ سَمر ٹائم", "HKT": "ہانگ کانگ سٹینڈرڈ ٹائم", "HOVST": "حووڑ سَمَر ٹایِم", "HOVT": "حووڑ سٹینڈرڈ ٹایِم", "ICT": "اِنڑوچَینا ٹایِم", "IDT": "اِسرٲیِلی ڑےلایِٔٹ ٹایِم", "IOT": "ہِندوستٲنؠ اوشَن ٹائم", "IRKST": "اِرکُٹسک سَمَر ٹایِم", "IRKT": "اِرکُٹسک سٹینڈرڈ ٹایِم", "IRST": "اِیٖرٲنؠ سٹینڑاڑ ٹایِم", "IRST DST": "اِیٖرٲنی سَمَر ٹایِم", "IST": "ہِندوستان", "IST Israel": "اِسرٲیِلی سٹینڈرڈ ٹایِم", "JDT": "جاپٲنؠ ڑےلایِٔٹ ٹایِم", "JST": "جاپٲنؠ سٹینڈرڈ ٹایِم", "KOST": "کورسَے ٹایِم", "KRAST": "کرؠسنوےیارسک سَمَر ٹایِم", "KRAT": "کرؠسنوےیارسک سٹینڈرڈ ٹایِم", "KST": "کورِیا سٹینڈرڈ ٹایِم", "KST DST": "کورِیا ڑےلایِٔٹ ٹایِم", "LHDT": "لعاڑ ڑےلایٔٹ ٹایِم", "LHST": "لعاڑ حووے سٹینڑاڑ ٹایِم", "LINT": "لایِٔن ججیٖرُک ٹایِم", "MAGST": "مَگَدَن سَمَر ٹایِم", "MAGT": "مَگَدَن سٹینڈرڈ ٹایِم", "MART": "مارقیوٗسَس ٹایِم", "MAWT": "ماسَن ٹایِم", "MDT": "ماونٹین ڈےلایِٔٹ ٹایِم", "MESZ": "مرکزی یوٗرپی سَمَر ٹایِم", "MEZ": "مرکزی یوٗرپی سٹینڑاڑ ٹایِم", "MHT": "مارشَل ججیٖرُک ٹایِم", "MMT": "مِیانمَر ٹایِم", "MSD": "ماسکو سَمَر ٹایِم", "MST": "ماونٹین سٹینڑاڑ ٹایِم", "MUST": "مورِشَس سَمَر ٹایِم", "MUT": "مورِشَس سٹینڈرڈ ٹایِم", "MVT": "مالدیٖوٕز ٹایِم", "MYT": "مَلیشِیا ٹایِم", "NCT": "نِو کیلؠڑونِیا سٹینڑاڑ ٹایِم", "NDT": "نیو فاؤنڈ لینڈ ڈے لائٹ ٹائم", "NDT New Caledonia": "نِو کیلؠڑونِیس سَمَر ٹایِم", "NFDT": "نورفعاک سَمَر ٹایِم", "NFT": "نورفعاک سٹینڑاڑ ٹایِم", "NOVST": "نۄوۄسِبٔرسک سَمَر ٹایِم", "NOVT": "نۄوۄسِبٔرسک سٹینڈرڈ ٹایِم", "NPT": "نؠپٲلؠ ٹایِم", "NRT": "نَعوٗروٗ ٹایِم", "NST": "نیو فاؤنڈ لینڈ سٹینڈرڈ ٹائم", "NUT": "نِیوٗ ٹایِم", "NZDT": "نِوزِلینڑ ڑےلایٔٹ ٹایِم", "NZST": "نِوزِلینڑ سٹینڑاڑ ٹایِم", "OESZ": "مشرقی یوٗرپی سَمَر ٹایِم", "OEZ": "مشرقی یوٗرپی سٹینڑاڑ ٹایِم", "OMSST": "اۄمسک سَمَر ٹایِم", "OMST": "اۄمسک سٹینڈرڈ ٹایِم", "PDT": "پیسِفِک ڈےلایِٔٹ ٹایِم", "PDTM": "میکسیکن پیسیفک ڈے لائٹ ٹائم", "PETDT": "کَمچَٹکا سَمَر ٹایِم", "PETST": "کَمچَٹکا سٹینڑاڑ ٹایِم", "PGT": "پاپُعا نیوٗ گؠنی ٹایِم", "PHOT": "پھونِکس ججیٖرُک ٹایِم", "PKT": "پاکِستان سٹینڑاڑ ٹایِم", "PKT DST": "پاکِستان سَمَر ٹایِم", "PMDT": "سینٹ پَیری مِقیوٗلَن ڑےلایِٔٹ ٹایِم", "PMST": "سینٹ پَیری مِقیوٗلَن سٹینڑاڑ ٹایِم", "PONT": "پونیپ ٹایِم", "PST": "پیسِفِک سٹینڑاڑ ٹایِم", "PST Philippine": "پھِلِپایِن سٹینڑاڑ ٹایِم", "PST Philippine DST": "پھِلِپایِن سَمَر ٹایِم", "PST Pitcairn": "پِٹکیرٕن ٹایِم", "PSTM": "میکسیکن پیسیفک سٹینڈرڈ ٹائم", "PWT": "پَلاو ٹایِم", "PYST": "پیرؠگوے سَمَر ٹایِم", "PYT": "پیرؠگوے سٹینڈرڈ ٹایِم", "PYT Korea": "یونگ یانگ ٹائم", "RET": "رِیوٗنِیَن ٹایِم", "ROTT": "روتھؠرا ٹایِم", "SAKST": "سَکھؠلِن سَمَر ٹایِم", "SAKT": "سَکھؠلِن سٹینڈرڈ ٹایِم", "SAMST": "سمؠرا سَمَر ٹایِم", "SAMT": "سمؠرا سٹینڑاڑ ٹایِم", "SAST": "جنوٗبی افریقا ٹایِم", "SBT": "سولومَن ججیٖرَن ہُند ٹایِم", "SCT": "سیشؠلٕز ٹایِم", "SGT": "سِنگاپوٗر ٹایِم", "SLST": "لَنکا ٹایِم", "SRT": "سُرِنام ٹایِم", "SST Samoa": "سؠموآ سٹینڑاڑ ٹایِم", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "سؠموآ سَمَر ٹایِم", "SYOT": "سیووا ٹایِم", "TAAF": "فرینچ جنوبی تٕہ انٹارکٹِک ٹائم", "TAHT": "ٹاہِٹی ٹایِم", "TJT": "تاجکستان ٹائم", "TKT": "ٹوکؠلو ٹایِم", "TLT": "ایٖسٹ ٹیٖمَر ٹایِم", "TMST": "تُرکمؠنِستان سَمَر ٹایِم", "TMT": "ترکمانستان سٹینڈرڈ ٹائم", "TOST": "ٹعانگا سَمَر ٹایِم", "TOT": "ٹعانگا سٹینڑاڑ ٹایِم", "TVT": "ٹوٗوَلوٗ ٹایِم", "TWT": "ٹے پے سٹینڈرڈ ٹائم", "TWT DST": "ٹے پے ڈے لائٹ ٹائم", "ULAST": "اولن باٹر سَمَر ٹایِم", "ULAT": "اولن باٹر سٹینڈرڈ ٹائم", "UYST": "یوٗرؠگوَے سَمَر ٹایِم", "UYT": "یوٗرؠگوَے سٹینڈرڈ ٹایِم", "UZT": "اُزبیکِستان سٹینڈرڈ ٹایِم", "UZT DST": "اُزبیکِستانُک سَمَر ٹایِم", "VET": "وؠنؠزیوٗلا ٹایِم", "VLAST": "ولاڑِووسٹوک سَمَر ٹایِم", "VLAT": "ولاڑِووسٹوک سٹینڈرڈ ٹایِم", "VOLST": "وولگوگریڑ سَمَر ٹایِم", "VOLT": "وولگوگریڑ سٹینڈرڈ ٹایِم", "VOST": "ووسٹوک ٹایِم", "VUT": "وَنوٗاَٹوٗ سٹینڑاڑ ٹایِم", "VUT DST": "وَنوٗاَٹوٗ سَمَر ٹایِم", "WAKT": "ویک ججیٖرُک ٹایِم", "WARST": "مغربی ارجؠنٹیٖنا سَمَر ٹایِم", "WART": "مغربی ارجؠنٹیٖنا سٹینڑاڑ ٹایِم", "WAST": "مغربی افریٖقا ٹایِم", "WAT": "مغربی افریٖقا ٹایِم", "WESZ": "مغرِبی یوٗرِپی سَمَر ٹایِم", "WEZ": "مغرِبی یوٗرپی سٹینڈرڈ ٹایِم", "WFT": "والِس تہٕ فیوٗٹیوٗنا ٹایِم", "WGST": "مغرِبی گریٖن لینڈُک سَمَر ٹایِم", "WGT": "مغرِبی گریٖن لینڈُک سٹینڑاڑ ٹایِم", "WIB": "مغرِبی اِنڑونیشِیا ٹایِم", "WIT": "مشرِقی اِنڑونیشِیا ٹایِم", "WITA": "مرکزی اِنڑونیشِیا ٹایِم", "YAKST": "یَکُٹُسک سَمَر ٹایِم", "YAKT": "یَکُٹسک سٹینڈرڈ ٹایِم", "YEKST": "یؠکَٹرِنبٔرگ سَمَر ٹایِم", "YEKT": "یؠکَٹٔرِنبٔرگ سٹینڈرڈ ٹایِم", "YST": "یوکون ٹائم", "МСК": "ماسکو سٹینڈرڈ ٹایِم", "اقتاۋ": "اؠکٹاؤ سٹینڑاڑ ٹایِم", "اقتاۋ قالاسى": "اؠکٹاؤ سَمَر ٹایِم", "اقتوبە": "اؠکٹوب سٹینڑاڑ ٹایِم", "اقتوبە قالاسى": "اؠکٹوب سَمَر ٹایِم", "الماتى": "اؠلمؠٹی سٹینڑاڑ ٹایِم", "الماتى قالاسى": "اؠلمؠٹی سَمَر ٹایِم", "باتىس قازاق ەلى": "مغربی قازقستان ٹائم", "شىعىش قازاق ەلى": "مشرقی قازقستان ٹائم", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "کرغزستان ٹائم", "قىزىلوردا": "قِزلوڑا سٹینڑاڑ ٹایِم", "قىزىلوردا قالاسى": "قِزلوڑا سَمَر ٹایِم", "∅∅∅": "پٔروٗ سَمَر ٹایِم"},
	}
}

// Locale returns the current translators string locale
func (ks *ks_Deva_IN) Locale() string {
	return ks.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ks_Deva_IN'
func (ks *ks_Deva_IN) PluralsCardinal() []locales.PluralRule {
	return ks.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ks_Deva_IN'
func (ks *ks_Deva_IN) PluralsOrdinal() []locales.PluralRule {
	return ks.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ks_Deva_IN'
func (ks *ks_Deva_IN) PluralsRange() []locales.PluralRule {
	return ks.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ks *ks_Deva_IN) MonthAbbreviated(month time.Month) string {
	if len(ks.monthsAbbreviated) == 0 {
		return ""
	}
	return ks.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ks *ks_Deva_IN) MonthsAbbreviated() []string {
	return ks.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ks *ks_Deva_IN) MonthNarrow(month time.Month) string {
	if len(ks.monthsNarrow) == 0 {
		return ""
	}
	return ks.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ks *ks_Deva_IN) MonthsNarrow() []string {
	return ks.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ks *ks_Deva_IN) MonthWide(month time.Month) string {
	if len(ks.monthsWide) == 0 {
		return ""
	}
	return ks.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ks *ks_Deva_IN) MonthsWide() []string {
	return ks.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ks *ks_Deva_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ks.daysAbbreviated) == 0 {
		return ""
	}
	return ks.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ks *ks_Deva_IN) WeekdaysAbbreviated() []string {
	return ks.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ks *ks_Deva_IN) WeekdayNarrow(weekday time.Weekday) string {
	if len(ks.daysNarrow) == 0 {
		return ""
	}
	return ks.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ks *ks_Deva_IN) WeekdaysNarrow() []string {
	return ks.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ks *ks_Deva_IN) WeekdayShort(weekday time.Weekday) string {
	if len(ks.daysShort) == 0 {
		return ""
	}
	return ks.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ks *ks_Deva_IN) WeekdaysShort() []string {
	return ks.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ks *ks_Deva_IN) WeekdayWide(weekday time.Weekday) string {
	if len(ks.daysWide) == 0 {
		return ""
	}
	return ks.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ks *ks_Deva_IN) WeekdaysWide() []string {
	return ks.daysWide
}

// Decimal returns the decimal point of number
func (ks *ks_Deva_IN) Decimal() string {
	return ks.decimal
}

// Group returns the group of number
func (ks *ks_Deva_IN) Group() string {
	return ks.group
}

// Group returns the minus sign of number
func (ks *ks_Deva_IN) Minus() string {
	return ks.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ks_Deva_IN' and handles both Whole and Real numbers based on 'v'
func (ks *ks_Deva_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ks.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ks.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ks.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ks_Deva_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ks *ks_Deva_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ks.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ks.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ks.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ks.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ks.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ks.group[0])
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
		b = append(b, ks.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ks.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ks_Deva_IN'
// in accounting notation.
func (ks *ks_Deva_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ks.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ks.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ks.group[0])
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

		b = append(b, ks.minus[0])

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
			b = append(b, ks.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ks.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ks.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ks.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ks.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ks.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ks.periodsAbbreviated[0]...)
	} else {
		b = append(b, ks.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ks.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ks.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ks.periodsAbbreviated[0]...)
	} else {
		b = append(b, ks.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ks.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ks.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ks.periodsAbbreviated[0]...)
	} else {
		b = append(b, ks.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ks_Deva_IN'
func (ks *ks_Deva_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ks.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ks.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ks.periodsAbbreviated[0]...)
	} else {
		b = append(b, ks.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ks.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
