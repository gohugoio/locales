package kxv_Deva

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kxv_Deva struct {
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

// New returns a new instance of translator for the 'kxv_Deva' locale
func New() locales.Translator {
	return &kxv_Deva{
		locale:                 "kxv_Deva",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "पुसु", "माहा", "पागु", "हिर्रे", "बेसे", "जाट्टा", "आसाड़ी", "स्राबाँ", "बाॅदो", "दासारा", "दिवी", "पान्डे"},
		monthsNarrow:           []string{"", "पु", "मा", "पा", "हि", "बे", "जा", "आ", "स्रा", "बाॅ", "दा", "दि", "पा"},
		monthsWide:             []string{"", "पुसु लेञ्जु", "माहाका लेञ्जु", "पागुणी लेञ्जु", "हिरे लेञ्जु", "बेसे लेञ्जु", "जाटा लेञ्जु", "आसाड़ी लेञ्जु", "स्राबाँ लेञ्जु", "बोदो लेञ्जु", "दसारा लेञ्जु", "दिवी लेञ्जु", "पान्डे लेञ्जु"},
		daysAbbreviated:        []string{"आदि", "साॅम्मा", "मान्गा", "पूदा", "लाक्की", "सुकुरु", "सान्नि"},
		daysNarrow:             []string{"आ", "साॅ", "मा", "पू", "ला", "सु", "सा"},
		daysShort:              []string{"आ", "साॅ", "मा", "पू", "ला", "सु", "सा"},
		daysWide:               []string{"आदि वारा", "साॅम्वारा", "मंगाड़ा", "पुद्दारा", "लाक्कि वारा", "सुकुरु वारा", "सान्नि वारा"},
		periodsAbbreviated:     []string{"ए\u202fएम", "पी\u202fएम"},
		timezones:              map[string]string{"ACDT": "ऑस्\u200dट्रेलियाति मादिनी डेलाइट बेला", "ACST": "ऑस्\u200dट्रेलियाति मादिनी मानांकॉ बेला", "ACT": "ACT", "ACWDT": "ऑस्\u200dट्रेलियाति मादिनी वेड़ा कुण्पु डेलाइट बेला", "ACWST": "ऑस्\u200dट्रेलियाति मादिनी वेड़ा कुण्पु मानांकॉ बेला", "ADT": "अटलांटिक डेलाइट बेला", "ADT Arabia": "अरब डेलाइट बेला", "AEDT": "ऑस्\u200dट्रेलियाति वेड़ा हॉपु डेलाइट बेला", "AEST": "ऑस्\u200dट्रेलियाति वेड़ा हॉपु मानांकॉ बेला", "AFT": "आपगानिस्तान बेला", "AKDT": "अलास्\u200dका डेलाइट बेला", "AKST": "अलास्\u200dका मानांकॉ बेला", "AMST": "अमेज़न काराँ मासा बेला", "AMST Armenia": "आर्मेनिया काराँ मासा बेला", "AMT": "अमेज़न मानांकॉ बेला", "AMT Armenia": "आर्मेनिया मानांकॉ बेला", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "अर्जेंटीना काराँ मासा बेला", "ART": "अर्जेंटीना मानांकॉ बेला", "AST": "अटलांटिक मानांकॉ बेला", "AST Arabia": "अरब मानांकॉ बेला", "AWDT": "ऑस्ट्रेलियाति वेड़ा कुण्पु डेलाइट बेला", "AWST": "ऑस्ट्रेलियाति वेड़ा कुण्पु मानांकॉ बेला", "AZST": "अजरबाइजान काराँ मासा बेला", "AZT": "अजरबाइजान मानांकॉ बेला", "BDT Bangladesh": "बांग्लादेस काराँ मासा बेला", "BNT": "ब्रूनेति दारूस्सलम बेला", "BOT": "बोलीविया बेला", "BRST": "ब्राज़ीलिया काराँ मासा बेला", "BRT": "ब्राज़ीलिया मानांकॉ बेला", "BST Bangladesh": "बांग्लादेस मानांकॉ बेला", "BT": "बुटान बेला", "CAST": "CAST", "CAT": "मादिनी आप्रिका बेला", "CCT": "कोकोस द्वीप बेला", "CDT": "मादिनी डेलाइट बेला", "CHADT": "च्याताम डेलाइट बेला", "CHAST": "च्याताम मानांकॉ बेला", "CHUT": "चुक बेला", "CKT": "कुक द्वीप मानांकॉ बेला", "CKT DST": "कुक द्वीप आधा काराँ मासा बेला", "CLST": "चिली काराँ मासा बेला", "CLT": "चिली मानांकॉ बेला", "COST": "कोलंबिया काराँ मासा बेला", "COT": "कोलंबिया मानांकॉ बेला", "CST": "मादिनी मानांकॉ बेला", "CST China": "चीन मानांकॉ बेला", "CST China DST": "चीन डेलाइट बेला", "CVST": "केप वर्ड काराँ मासा बेला", "CVT": "केप वर्ड मानांकॉ बेला", "CXT": "क्रिसमस द्वीप बेला", "ChST": "चामोरो मानांकॉ बेला", "ChST NMI": "ChST NMI", "CuDT": "क्यूबा डेलाइट बेला", "CuST": "क्यूबा मानांकॉ बेला", "DAVT": "डेविस बेला", "DDUT": "ड्यूमोंट डी अर्विले बेला", "EASST": "ईस्टर द्वीप काराँ मासा बेला", "EAST": "ईस्टर द्वीप मानांकॉ बेला", "EAT": "वेड़ा हॉपु आप्रिका बेला", "ECT": "इक्वाडोर बेला", "EDT": "वेड़ा हॉपु डेलाइट बेला", "EGDT": "वेड़ा हॉपु ग्रीनलेंड काराँ मासा बेला", "EGST": "वेड़ा हॉपु ग्रीनलेंड मानांकॉ बेला", "EST": "वेड़ा हॉपु मानांकॉ बेला", "FEET": "ऑरॉ वेड़ा हॉपु युरोप-ति बेला", "FJT": "प़िजी मानांकॉ बेला", "FJT Summer": "प़िजी काराँ मासा बेला", "FKST": "पाक -लेंड द्वीप काराँ मासा बेला", "FKT": "पाक -लेंड द्वीप मानांकॉ बेला", "FNST": "प़र्नांर्डो डे नोरोन्हा काराँ मासा बेला", "FNT": "प़र्नांर्डो डे नोरोन्हा मानांकॉ बेला", "GALT": "ग्यालापागोस ति बेला", "GAMT": "ग्याम्बीयर बेला", "GEST": "जॉर्जिया काराँ मासा बेला", "GET": "जॉर्जिया मानांकॉ बेला", "GFT": "प़्रेंच गुयाना बेला", "GIT": "गिल्बर्ट द्वीप बेला", "GMT": "ग्रीनविच मीन बेला", "GNSST": "GNSST", "GNST": "GNST", "GST": "गल्प मानांकॉ बेला", "GST Guam": "GST Guam", "GYT": "गुयाना बेला", "HADT": "हवाति–आल्यूसन डेलाइट बेला", "HAST": "हवाति–आल्यूसन मानांकॉ बेला", "HKST": "हाँग काँग काराँ मासा बेला", "HKT": "हाँग काँग मानांकॉ बेला", "HOVST": "होव्ड काराँ मासा बेला", "HOVT": "होव्ड मानांकॉ बेला", "ICT": "इंडोचाइना बेला", "IDT": "इज़राइल डेलाइट बेला", "IOT": "बारॉत काजा सामुद्री बेला", "IRKST": "इर्कुत्स्क काराँ मासा बेला", "IRKT": "इर्कुत्स्क मानांकॉ बेला", "IRST": "इरान मानांकॉ बेला", "IRST DST": "तिरान डेलाइट बेला", "IST": "बारतीय मानांकॉ बेला", "IST Israel": "इज़राइल मानांकॉ बेला", "JDT": "जापान डेलाइट बेला", "JST": "जापान मानांकॉ बेला", "KOST": "कोसराए बेला", "KRAST": "क्रास्नोयार्स्क काराँ मासा बेला", "KRAT": "क्रास्नोयार्स्क मानांकॉ बेला", "KST": "कोरियाति मानांकॉ बेला", "KST DST": "कोरियाति डेलाइट बेला", "LHDT": "लॉर्ड होवे डेलाइट बेला", "LHST": "लॉर्ड होवे मानांकॉ बेला", "LINT": "लाइन द्वीप बेला", "MAGST": "मागादान काराँ मासा बेला", "MAGT": "मागादान मानांकॉ बेला", "MART": "मार्केसस बेला", "MAWT": "माव्सन बेला", "MDT": "MDT", "MESZ": "मादिनी युरोप-ति काराँ मासा बेला", "MEZ": "मादिनी युरोप-ति मानांकॉ बेला", "MHT": "मार्सल द्वीप बेला", "MMT": "म्यांमार बेला", "MSD": "मॉस्को काराँ मासा बेला", "MST": "MST", "MUST": "मॉरीसस काराँ मासा बेला", "MUT": "मॉरीसस मानांकॉ बेला", "MVT": "मालदीव बेला", "MYT": "मलेसिया बेला", "NCT": "न्यू केलेडोनिया मानांकॉ बेला", "NDT": "न्यूप़ाउंडलेंड डेलाइट बेला", "NDT New Caledonia": "न्यू केलेडोनिया काराँ मासा बेला", "NFDT": "नॉरप़ॉक द्वीप डेलाइट बेला", "NFT": "नॉरप़ॉक द्वीप मानांकॉ बेला", "NOVST": "नोवोसिबिर्स्क काराँ मासा बेला", "NOVT": "नोवोसिबिर्स्क मानांकॉ बेला", "NPT": "नेपाल बेला", "NRT": "नॉउरु बेला", "NST": "न्यूप़ाउंडलेंड मानांकॉ बेला", "NUT": "नीयू बेला", "NZDT": "न्यूज़ीलेंड डेलाइट बेला", "NZST": "न्यूज़ीलेंड मानांकॉ बेला", "OESZ": "वेड़ा हॉपु युरोप-ति काराँ मासा बेला", "OEZ": "वेड़ा हॉपु युरोप-ति मानांकॉ बेला", "OMSST": "ओम्स्क काराँ मासा बेला", "OMST": "ओम्स्क मानांकॉ बेला", "PDT": "पेसिपिक डेलाइट बेला", "PDTM": "मेक्सिकोति पेसिपिक डेलाइट बेला", "PETDT": "PETDT", "PETST": "PETST", "PGT": "पापुआ न्यू गिनी बेला", "PHOT": "प़ीनिक्स द्वीप बेला", "PKT": "पाकिस्तान मानांकॉ बेला", "PKT DST": "पाकिस्तान काराँ मासा बेला", "PMDT": "सेंट पिएरे ऑड़े मिक्वेलान डेलाइट बेला", "PMST": "सेंट पिएरे ऑड़े मिक्वेलान मानांकॉ बेला", "PONT": "पोनापे बेला", "PST": "पेसिपिक मानांकॉ बेला", "PST Philippine": "प़िलिपीन मानांकॉ बेला", "PST Philippine DST": "प़िलिपीन काराँ मासा बेला", "PST Pitcairn": "पिटकेर्न बेला", "PSTM": "मेक्सिकोति पेसिपिक मानांकॉ बेला", "PWT": "पलाउ बेला", "PYST": "पैराग्वे काराँ मासा बेला", "PYT": "पैराग्वे मानांकॉ बेला", "PYT Korea": "प्योंगयांग बेला", "RET": "रीयूनियन बेला", "ROTT": "रोथेरा बेला", "SAKST": "सकालिन काराँ मासा बेला", "SAKT": "सकालिन मानांकॉ बेला", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "दक्षिण आप्रिका मानांकॉ बेला", "SBT": "सोलोमन द्वीप बेला", "SCT": "सेसेल्स बेला", "SGT": "सिंगापुर बेला", "SLST": "SLST", "SRT": "सूरीनाम बेला", "SST Samoa": "समोआ मानांकॉ बेला", "SST Samoa Apia": "एपिआ मानांकॉ बेला", "SST Samoa Apia DST": "एपिआ डेलाइट बेला", "SST Samoa DST": "समोआ डेलाइट बेला", "SYOT": "स्योवा बेला", "TAAF": "दॉकिणॉ प़्रांस ऑड़े अंटार्कटिक बेला", "TAHT": "ताहिती बेला", "TJT": "ताजिकिस्तान बेला", "TKT": "टोकेलाऊ बेला", "TLT": "वेड़ा हॉपु तिमोर बेला", "TMST": "तुर्कमेनिस्तान काराँ मासा बेला", "TMT": "तुर्कमेनिस्तान मानांकॉ बेला", "TOST": "टोंगा काराँ मासा बेला", "TOT": "टोंगा मानांकॉ बेला", "TVT": "तुवालू बेला", "TWT": "ताइपे मानांकॉ बेला", "TWT DST": "ताइपे डेलाइट बेला", "ULAST": "उलान बटोर काराँ मासा बेला", "ULAT": "उलान बटोर मानांकॉ बेला", "UYST": "उरुग्वे काराँ मासा बेला", "UYT": "उरुग्वे मानांकॉ बेला", "UZT": "उज़्बेकिस्तान मानांकॉ बेला", "UZT DST": "उज़्बेकिस्तान काराँ मासा बेला", "VET": "वेनेज़ुएला बेला", "VLAST": "व्लादिवोस्तोक काराँ मासा बेला", "VLAT": "व्लादिवोस्तोक मानांकॉ बेला", "VOLST": "वोल्गोग्राड काराँ मासा बेला", "VOLT": "वोल्गोग्राड मानांकॉ बेला", "VOST": "वोस्तोक बेला", "VUT": "वनुआतू मानांकॉ बेला", "VUT DST": "वनुआतू काराँ मासा बेला", "WAKT": "वेक द्वीप बेला", "WARST": "वेड़ा कुण्पु अर्जेंटीना काराँ मासा बेला", "WART": "वेड़ा कुण्पु अर्जेंटीना मानांकॉ बेला", "WAST": "पस्चिम आप्रिका बेला", "WAT": "पस्चिम आप्रिका बेला", "WESZ": "वेड़ा कुण्पु युरोप-ति काराँ मासा बेला", "WEZ": "वेड़ा कुण्पु युरोप-ति मानांकॉ बेला", "WFT": "वालिस ऑड़े प़्यूचूना बेला", "WGST": "वेड़ा कुण्पु ग्रीनलेंड काराँ मासा बेला", "WGT": "वेड़ा कुण्पु ग्रीनलेंड मानांकॉ बेला", "WIB": "वेड़ा कुण्पु इंडोनेसिया बेला", "WIT": "वेड़ा हॉपु इंडोनेसिया बेला", "WITA": "मध्य इंडोनेसिया बेला", "YAKST": "याकुत्स्क काराँ मासा बेला", "YAKT": "याकुत्स्क मानांकॉ बेला", "YEKST": "येकातेरिनबर्ग काराँ मासा बेला", "YEKT": "येकातेरिनबर्ग मानांकॉ बेला", "YST": "YST", "МСК": "मॉस्को मानांकॉ बेला", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "वेड़ा कुण्पु कज़ाकस्तान बेला", "شىعىش قازاق ەلى": "वेड़ा हॉपु कज़ाकस्तान बेला", "قازاق ەلى": "कज़ाकस्तान बेला", "قىرعىزستان": "किर्गिस्\u200dतान बेला", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "पेरू काराँ मासा बेला"},
	}
}

// Locale returns the current translators string locale
func (kxv *kxv_Deva) Locale() string {
	return kxv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kxv_Deva'
func (kxv *kxv_Deva) PluralsCardinal() []locales.PluralRule {
	return kxv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kxv_Deva'
func (kxv *kxv_Deva) PluralsOrdinal() []locales.PluralRule {
	return kxv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kxv_Deva'
func (kxv *kxv_Deva) PluralsRange() []locales.PluralRule {
	return kxv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Deva'
func (kxv *kxv_Deva) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Deva'
func (kxv *kxv_Deva) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kxv_Deva'
func (kxv *kxv_Deva) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kxv *kxv_Deva) MonthAbbreviated(month time.Month) string {
	if len(kxv.monthsAbbreviated) == 0 {
		return ""
	}
	return kxv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kxv *kxv_Deva) MonthsAbbreviated() []string {
	return kxv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kxv *kxv_Deva) MonthNarrow(month time.Month) string {
	if len(kxv.monthsNarrow) == 0 {
		return ""
	}
	return kxv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kxv *kxv_Deva) MonthsNarrow() []string {
	return kxv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kxv *kxv_Deva) MonthWide(month time.Month) string {
	if len(kxv.monthsWide) == 0 {
		return ""
	}
	return kxv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kxv *kxv_Deva) MonthsWide() []string {
	return kxv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kxv *kxv_Deva) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(kxv.daysAbbreviated) == 0 {
		return ""
	}
	return kxv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kxv *kxv_Deva) WeekdaysAbbreviated() []string {
	return kxv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kxv *kxv_Deva) WeekdayNarrow(weekday time.Weekday) string {
	if len(kxv.daysNarrow) == 0 {
		return ""
	}
	return kxv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kxv *kxv_Deva) WeekdaysNarrow() []string {
	return kxv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kxv *kxv_Deva) WeekdayShort(weekday time.Weekday) string {
	if len(kxv.daysShort) == 0 {
		return ""
	}
	return kxv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kxv *kxv_Deva) WeekdaysShort() []string {
	return kxv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kxv *kxv_Deva) WeekdayWide(weekday time.Weekday) string {
	if len(kxv.daysWide) == 0 {
		return ""
	}
	return kxv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kxv *kxv_Deva) WeekdaysWide() []string {
	return kxv.daysWide
}

// Decimal returns the decimal point of number
func (kxv *kxv_Deva) Decimal() string {
	return kxv.decimal
}

// Group returns the group of number
func (kxv *kxv_Deva) Group() string {
	return kxv.group
}

// Group returns the minus sign of number
func (kxv *kxv_Deva) Minus() string {
	return kxv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kxv_Deva' and handles both Whole and Real numbers based on 'v'
func (kxv *kxv_Deva) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kxv_Deva' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kxv *kxv_Deva) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kxv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kxv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Deva'
// in accounting notation.
func (kxv *kxv_Deva) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(kxv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyNegativePrefix[j])
		}

		b = append(b, kxv.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, kxv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kxv_Deva'
func (kxv *kxv_Deva) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kxv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
