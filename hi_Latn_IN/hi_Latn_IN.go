package hi_Latn_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type hi_Latn_IN struct {
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

// New returns a new instance of translator for the 'hi_Latn_IN' locale
func New() locales.Translator {
	return &hi_Latn_IN{
		locale:             "hi_Latn_IN",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsRange:       []locales.PluralRule{2, 6},
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "जन॰", "फ़र॰", "मार्च", "अप्रैल", "मई", "जून", "जुल॰", "अग॰", "सित॰", "अक्टू॰", "नव॰", "दिस॰"},
		monthsNarrow:       []string{"", "ज", "फ़", "मा", "अ", "म", "जू", "जु", "अ", "सि", "अ", "न", "दि"},
		monthsWide:         []string{"", "जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्टूबर", "नवंबर", "दिसंबर"},
		daysAbbreviated:    []string{"रवि", "सोम", "मंगल", "बुध", "गुरु", "शुक्र", "शनि"},
		daysNarrow:         []string{"र", "सो", "मं", "बु", "गु", "शु", "श"},
		daysShort:          []string{"र", "सो", "मं", "बु", "गु", "शु", "श"},
		daysWide:           []string{"रविवार", "सोमवार", "मंगलवार", "बुधवार", "गुरुवार", "शुक्रवार", "शनिवार"},
		periodsAbbreviated: []string{"am", "pm"},
		periodsNarrow:      []string{"", ""},
		erasAbbreviated:    []string{"ईसा-पूर्व", "ईस्वी"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"", ""},
		timezones:          map[string]string{"ACDT": "ऑस्\u200dट्रेलियाई केंद्रीय डेलाइट समय", "ACST": "ऑस्\u200dट्रेलियाई केंद्रीय मानक समय", "ACT": "ACT", "ACWDT": "ऑस्\u200dट्रेलियाई केंद्रीय पश्चिमी डेलाइट समय", "ACWST": "ऑस्\u200dट्रेलियाई केंद्रीय पश्चिमी मानक समय", "ADT": "अटलांटिक डेलाइट समय", "ADT Arabia": "अरब डेलाइट समय", "AEDT": "ऑस्\u200dट्रेलियाई पूर्वी डेलाइट समय", "AEST": "ऑस्\u200dट्रेलियाई पूर्वी मानक समय", "AFT": "अफ़गानिस्तान समय", "AKDT": "अलास्\u200dका डेलाइट समय", "AKST": "अलास्\u200dका मानक समय", "AMST": "अमेज़न ग्रीष्मकालीन समय", "AMST Armenia": "आर्मेनिया ग्रीष्मकालीन समय", "AMT": "अमेज़न मानक समय", "AMT Armenia": "आर्मेनिया मानक समय", "ANAST": "एनाडीयर ग्रीष्मकालीन समय", "ANAT": "एनाडीयर मानक समय", "ARST": "अर्जेंटीना ग्रीष्मकालीन समय", "ART": "अर्जेंटीना मानक समय", "AST": "अटलांटिक मानक समय", "AST Arabia": "अरब मानक समय", "AWDT": "ऑस्ट्रेलियाई पश्चिमी डेलाइट समय", "AWST": "ऑस्ट्रेलियाई पश्चिमी मानक समय", "AZST": "अज़रबैजान ग्रीष्मकालीन समय", "AZT": "अज़रबैजान मानक समय", "BDT Bangladesh": "बांग्लादेश ग्रीष्मकालीन समय", "BNT": "ब्रूनेई दारूस्सलम समय", "BOT": "बोलीविया समय", "BRST": "ब्राज़ीलिया ग्रीष्मकालीन समय", "BRT": "ब्राज़ीलिया मानक समय", "BST Bangladesh": "बांग्लादेश मानक समय", "BT": "भूटान समय", "CAST": "CAST", "CAT": "मध्य अफ़्रीका समय", "CCT": "कोकोस द्वीपसमूह समय", "CDT": "उत्तरी अमेरिकी केंद्रीय डेलाइट समय", "CHADT": "चैथम डेलाइट समय", "CHAST": "चैथम मानक समय", "CHUT": "चुक समय", "CKT": "कुक द्वीपसमूह मानक समय", "CKT DST": "कुक द्वीपसमूह अर्द्ध ग्रीष्मकालीन समय", "CLST": "चिली ग्रीष्मकालीन समय", "CLT": "चिली मानक समय", "COST": "कोलंबिया ग्रीष्मकालीन समय", "COT": "कोलंबिया मानक समय", "CST": "उत्तरी अमेरिकी केंद्रीय मानक समय", "CST China": "चीन मानक समय", "CST China DST": "चीन डेलाइट समय", "CVST": "केप वर्ड ग्रीष्मकालीन समय", "CVT": "केप वर्ड मानक समय", "CXT": "क्रिसमस द्वीप समय", "ChST": "चामोरो मानक समय", "ChST NMI": "ChST NMI", "CuDT": "क्यूबा डेलाइट समय", "CuST": "क्यूबा मानक समय", "DAVT": "डेविस समय", "DDUT": "ड्यूमोंट डी अर्विले समय", "EASST": "ईस्टर द्वीप ग्रीष्मकालीन समय", "EAST": "ईस्टर द्वीप मानक समय", "EAT": "पूर्वी अफ़्रीका समय", "ECT": "इक्वाडोर समय", "EDT": "उत्तरी अमेरिकी पूर्वी डेलाइट समय", "EGDT": "पूर्वी ग्रीनलैंड ग्रीष्मकालीन समय", "EGST": "पूर्वी ग्रीनलैंड मानक समय", "EST": "उत्तरी अमेरिकी पूर्वी मानक समय", "FEET": "अग्र पूर्वी यूरोपीय समय", "FJT": "फ़िजी मानक समय", "FJT Summer": "फ़िजी ग्रीष्मकालीन समय", "FKST": "फ़ॉकलैंड द्वीपसमूह ग्रीष्मकालीन समय", "FKT": "फ़ॉकलैंड द्वीपसमूह मानक समय", "FNST": "फ़र्नांर्डो डे नोरोन्हा ग्रीष्मकालीन समय", "FNT": "फ़र्नांर्डो डे नोरोन्हा मानक समय", "GALT": "गैलापेगोस का समय", "GAMT": "गैंबियर समय", "GEST": "जॉर्जिया ग्रीष्मकालीन समय", "GET": "जॉर्जिया मानक समय", "GFT": "फ़्रेंच गुयाना समय", "GIT": "गिल्बर्ट द्वीपसमूह समय", "GMT": "ग्रीनविच मीन टाइम", "GNSST": "GNSST", "GNST": "GNST", "GST": "खाड़ी मानक समय", "GST Guam": "GST Guam", "GYT": "गुयाना समय", "HADT": "हवाई–आल्यूशन मानक समय", "HAST": "हवाई–आल्यूशन मानक समय", "HKST": "हाँग काँग ग्रीष्मकालीन समय", "HKT": "हाँग काँग मानक समय", "HOVST": "होव्ड ग्रीष्मकालीन समय", "HOVT": "होव्ड मानक समय", "ICT": "इंडोचाइना समय", "IDT": "इज़राइल डेलाइट समय", "IOT": "हिंद महासागर समय", "IRKST": "इर्कुत्स्क ग्रीष्मकालीन समय", "IRKT": "इर्कुत्स्क मानक समय", "IRST": "ईरान मानक समय", "IRST DST": "ईरान डेलाइट समय", "IST": "भारतीय मानक समय", "IST Israel": "इज़राइल मानक समय", "JDT": "जापान डेलाइट समय", "JST": "जापान मानक समय", "KOST": "कोसराए समय", "KRAST": "क्रास्नोयार्स्क ग्रीष्मकालीन समय", "KRAT": "क्रास्नोयार्स्क मानक समय", "KST": "कोरियाई मानक समय", "KST DST": "कोरियाई डेलाइट समय", "LHDT": "लॉर्ड होवे डेलाइट समय", "LHST": "लॉर्ड होवे मानक समय", "LINT": "लाइन द्वीपसमूह समय", "MAGST": "मागादान ग्रीष्मकालीन समय", "MAGT": "मागादान मानक समय", "MART": "मार्केसस समय", "MAWT": "माव्सन समय", "MDT": "उत्तरी अमेरिकी माउंटेन डेलाइट समय", "MESZ": "मध्\u200dय यूरोपीय ग्रीष्\u200dमकालीन समय", "MEZ": "मध्य यूरोपीय मानक समय", "MHT": "मार्शल द्वीपसमूह समय", "MMT": "म्यांमार समय", "MSD": "मॉस्को ग्रीष्मकालीन समय", "MST": "उत्तरी अमेरिकी माउंटेन मानक समय", "MUST": "मॉरीशस ग्रीष्मकालीन समय", "MUT": "मॉरीशस मानक समय", "MVT": "मालदीव समय", "MYT": "मलेशिया समय", "NCT": "न्यू कैलेडोनिया मानक समय", "NDT": "न्यूफ़ाउंडलैंड डेलाइट समय", "NDT New Caledonia": "न्यू कैलेडोनिया ग्रीष्मकालीन समय", "NFDT": "नॉरफ़ॉक द्वीप डेलाइट समय", "NFT": "नॉरफ़ॉक द्वीप मानक समय", "NOVST": "नोवोसिबिर्स्क ग्रीष्मकालीन समय", "NOVT": "नोवोसिबिर्स्क मानक समय", "NPT": "नेपाल समय", "NRT": "नौरू समय", "NST": "न्यूफ़ाउंडलैंड मानक समय", "NUT": "नीयू समय", "NZDT": "न्यूज़ीलैंड डेलाइट समय", "NZST": "न्यूज़ीलैंड मानक समय", "OESZ": "पूर्वी यूरोपीय ग्रीष्मकालीन समय", "OEZ": "पूर्वी यूरोपीय मानक समय", "OMSST": "ओम्स्क ग्रीष्मकालीन समय", "OMST": "ओम्स्क मानक समय", "PDT": "उत्तरी अमेरिकी प्रशांत डेलाइट समय", "PDTM": "मेक्सिकन प्रशांत डेलाइट समय", "PETDT": "पेट्रोपेवलास्क-कैमचात्सकी ग्रीष्मकालीन समय", "PETST": "पेट्रोपेवलास्क-कैमचात्सकी मानक समय", "PGT": "पापुआ न्यू गिनी समय", "PHOT": "फ़ीनिक्स द्वीपसमूह समय", "PKT": "पाकिस्तान मानक समय", "PKT DST": "पाकिस्तान ग्रीष्मकालीन समय", "PMDT": "सेंट पिएरे और मिक्वेलान डेलाइट समय", "PMST": "सेंट पिएरे और मिक्वेलान मानक समय", "PONT": "पोनापे समय", "PST": "उत्तरी अमेरिकी प्रशांत मानक समय", "PST Philippine": "फ़िलिपीन मानक समय", "PST Philippine DST": "फ़िलिपीन ग्रीष्मकालीन समय", "PST Pitcairn": "पिटकैर्न समय", "PSTM": "मेक्सिकन प्रशांत मानक समय", "PWT": "पलाउ समय", "PYST": "पैराग्वे ग्रीष्मकालीन समय", "PYT": "पैराग्वे मानक समय", "PYT Korea": "प्योंगयांग समय", "RET": "रीयूनियन समय", "ROTT": "रोथेरा समय", "SAKST": "सखालिन ग्रीष्मकालीन समय", "SAKT": "सखालिन मानक समय", "SAMST": "समारा ग्रीष्मकालीन समय", "SAMT": "समारा मानक समय", "SAST": "दक्षिण अफ़्रीका मानक समय", "SBT": "सोलोमन द्वीपसमूह समय", "SCT": "सेशेल्स समय", "SGT": "सिंगापुर समय", "SLST": "SLST", "SRT": "सूरीनाम समय", "SST Samoa": "समोआ मानक समय", "SST Samoa Apia": "एपिआ मानक समय", "SST Samoa Apia DST": "एपिआ डेलाइट समय", "SST Samoa DST": "समोआ डेलाइट समय", "SYOT": "स्योवा समय", "TAAF": "दक्षिणी फ़्रांस और अंटार्कटिक समय", "TAHT": "ताहिती समय", "TJT": "ताजिकिस्तान समय", "TKT": "टोकेलाऊ समय", "TLT": "पूर्वी तिमोर समय", "TMST": "तुर्कमेनिस्तान ग्रीष्मकालीन समय", "TMT": "तुर्कमेनिस्तान मानक समय", "TOST": "टोंगा ग्रीष्मकालीन समय", "TOT": "टोंगा मानक समय", "TVT": "तुवालू समय", "TWT": "ताइपे मानक समय", "TWT DST": "ताइपे डेलाइट समय", "ULAST": "उलान बटोर ग्रीष्मकालीन समय", "ULAT": "उलान बटोर मानक समय", "UYST": "उरुग्वे ग्रीष्मकालीन समय", "UYT": "उरुग्वे मानक समय", "UZT": "उज़्बेकिस्तान मानक समय", "UZT DST": "उज़्बेकिस्तान ग्रीष्मकालीन समय", "VET": "वेनेज़ुएला समय", "VLAST": "व्लादिवोस्तोक ग्रीष्मकालीन समय", "VLAT": "व्लादिवोस्तोक मानक समय", "VOLST": "वोल्गोग्राड ग्रीष्मकालीन समय", "VOLT": "वोल्गोग्राड मानक समय", "VOST": "वोस्तोक समय", "VUT": "वनुआतू मानक समय", "VUT DST": "वनुआतू ग्रीष्मकालीन समय", "WAKT": "वेक द्वीप समय", "WARST": "पश्चिमी अर्जेंटीना ग्रीष्मकालीन समय", "WART": "पश्चिमी अर्जेंटीना मानक समय", "WAST": "पश्चिम अफ़्रीका समय", "WAT": "पश्चिम अफ़्रीका समय", "WESZ": "पश्चिमी यूरोपीय ग्रीष्\u200dमकालीन समय", "WEZ": "पश्चिमी यूरोपीय मानक समय", "WFT": "वालिस और फ़्यूचूना समय", "WGST": "पश्चिमी ग्रीनलैंड ग्रीष्मकालीन समय", "WGT": "पश्चिमी ग्रीनलैंड मानक समय", "WIB": "पश्चिमी इंडोनेशिया समय", "WIT": "पूर्वी इंडोनेशिया समय", "WITA": "मध्य इंडोनेशिया समय", "YAKST": "याकुत्स्क ग्रीष्मकालीन समय", "YAKT": "याकुत्स्क मानक समय", "YEKST": "येकातेरिनबर्ग ग्रीष्मकालीन समय", "YEKT": "येकातेरिनबर्ग मानक समय", "YST": "युकॉन समय", "МСК": "मॉस्को मानक समय", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "पश्चिम कज़ाखस्तान समय", "شىعىش قازاق ەلى": "पूर्व कज़ाखस्तान समय", "قازاق ەلى": "कज़ाखस्तान समय", "قىرعىزستان": "किर्गिस्\u200dतान समय", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "पेरू ग्रीष्मकालीन समय"},
	}
}

// Locale returns the current translators string locale
func (hi *hi_Latn_IN) Locale() string {
	return hi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hi_Latn_IN'
func (hi *hi_Latn_IN) PluralsCardinal() []locales.PluralRule {
	return hi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hi_Latn_IN'
func (hi *hi_Latn_IN) PluralsOrdinal() []locales.PluralRule {
	return hi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'hi_Latn_IN'
func (hi *hi_Latn_IN) PluralsRange() []locales.PluralRule {
	return hi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 3 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := hi.CardinalPluralRule(num1, v1)
	end := hi.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (hi *hi_Latn_IN) MonthAbbreviated(month time.Month) string {
	return hi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (hi *hi_Latn_IN) MonthsAbbreviated() []string {
	return hi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (hi *hi_Latn_IN) MonthNarrow(month time.Month) string {
	return hi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (hi *hi_Latn_IN) MonthsNarrow() []string {
	return hi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (hi *hi_Latn_IN) MonthWide(month time.Month) string {
	return hi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (hi *hi_Latn_IN) MonthsWide() []string {
	return hi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (hi *hi_Latn_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return hi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (hi *hi_Latn_IN) WeekdaysAbbreviated() []string {
	return hi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (hi *hi_Latn_IN) WeekdayNarrow(weekday time.Weekday) string {
	return hi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (hi *hi_Latn_IN) WeekdaysNarrow() []string {
	return hi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (hi *hi_Latn_IN) WeekdayShort(weekday time.Weekday) string {
	return hi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (hi *hi_Latn_IN) WeekdaysShort() []string {
	return hi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (hi *hi_Latn_IN) WeekdayWide(weekday time.Weekday) string {
	return hi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (hi *hi_Latn_IN) WeekdaysWide() []string {
	return hi.daysWide
}

// Decimal returns the decimal point of number
func (hi *hi_Latn_IN) Decimal() string {
	return hi.decimal
}

// Group returns the group of number
func (hi *hi_Latn_IN) Group() string {
	return hi.group
}

// Group returns the minus sign of number
func (hi *hi_Latn_IN) Minus() string {
	return hi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hi_Latn_IN' and handles both Whole and Real numbers based on 'v'
func (hi *hi_Latn_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 0
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, hi.group[0])
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
		b = append(b, hi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hi_Latn_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hi *hi_Latn_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, hi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hi.currencies[currency]
	l := len(s) + len(symbol) + 0
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, hi.group[0])
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

	if num < 0 {
		b = append(b, hi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hi_Latn_IN'
// in accounting notation.
func (hi *hi_Latn_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hi.currencies[currency]
	l := len(s) + len(symbol) + 0
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, hi.group[0])
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

		b = append(b, hi.minus[0])

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
			b = append(b, hi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, hi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, hi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, hi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, hi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, hi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, hi.periodsAbbreviated[0]...)
	} else {
		b = append(b, hi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, hi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, hi.periodsAbbreviated[0]...)
	} else {
		b = append(b, hi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, hi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, hi.periodsAbbreviated[0]...)
	} else {
		b = append(b, hi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'hi_Latn_IN'
func (hi *hi_Latn_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, hi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, hi.periodsAbbreviated[0]...)
	} else {
		b = append(b, hi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
