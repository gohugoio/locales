package xnr

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type xnr struct {
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

// New returns a new instance of translator for the 'xnr' locale
func New() locales.Translator {
	return &xnr{
		locale:                 "xnr",
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
		monthsAbbreviated:      []string{"", "जन॰", "फ़र॰", "मार्च", "अप्रैल", "मई", "जून", "जुल॰", "अग॰", "सित॰", "अक्तू॰", "नव॰", "दिस॰"},
		monthsNarrow:           []string{"", "ज", "फ़", "मा", "अ", "म", "जू", "जु", "अ", "सि", "अ", "न", "दि"},
		monthsWide:             []string{"", "जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्तूबर", "नवंबर", "दिसंबर"},
		daysAbbreviated:        []string{"तोआर", "सोआर", "मंगल", "बुध", "वीर", "शुक्कर", "शनि"},
		daysNarrow:             []string{"त", "सो", "मं", "बु", "वी", "शु", "श"},
		daysShort:              []string{"त", "सो", "मं", "बु", "वी", "शु", "श"},
		daysWide:               []string{"तोआर", "सोआर", "मंगलवार", "बुधवार", "वीरवार", "शुक्करवार", "शनिच्चरवार"},
		periodsAbbreviated:     []string{"भ्यागा", "दपेहरा/संजा"},
		timezones:              map[string]string{"ACDT": "अस्\u200dट्रेलिया दे केंद्रीय दे ध्याड़े दे उजाले दा टैम", "ACST": "अस्\u200dट्रेलिया दे केंद्रीय दा मानक टैम", "ACT": "ACT", "ACWDT": "अस्\u200dट्रेलियाई केंद्रीय पश्चिमी ध्याड़े दे उजाले दा टैम", "ACWST": "अस्\u200dट्रेलिया केंद्रीय पश्चिमी दा मानक टैम", "ADT": "अटलांटिक दे ध्याड़े दे उजाले दा टैम", "ADT Arabia": "अरब दे ध्याड़े दे उजाले दा टैम", "AEDT": "अस्\u200dट्रेलियाई पैलेी दे ध्याड़े दे उजाले दा टैम", "AEST": "अस्\u200dट्रेलियाई पैलेी दा मानक टैम", "AFT": "अफ़गानिस्ताने दा टैम", "AKDT": "अलास्\u200dका दे ध्याड़े दे उजाले दा टैम", "AKST": "अलास्\u200dका दा मानक टैम", "AMST": "एमेज़ोन दी तोंदिया दा टैम", "AMST Armenia": "आर्मेनिया दी तोंदिया दा टैम", "AMT": "एमेज़ोन दा मानक टैम", "AMT Armenia": "आर्मेनिया दा मानक टैम", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "अर्जेंटीना दी तोंदिया दा टैम", "ART": "अर्जेंटीना दा मानक टैम", "AST": "अटलांटिक दा मानक टैम", "AST Arabia": "अरब दा मानक टैम", "AWDT": "अस्ट्रेलिया पश्चिमी दे ध्याड़े दे उजाले दा टैम", "AWST": "अस्ट्रेलिया पश्चिमी दा मानक टैम", "AZST": "अज़रबैजान दी तोंदिया दा टैम", "AZT": "अज़रबैजान दा मानक टैम", "BDT Bangladesh": "बांग्लादेशे दी तोंदिया दा टैम", "BNT": "ब्रूनेई दारूस्सलम दा टैम", "BOT": "बोलीविया दा टैम", "BRST": "ब्राज़ीलिया दी तोंदिया दा टैम", "BRT": "ब्राज़ीलिया दा मानक टैम", "BST Bangladesh": "बांग्लादेशे दा मानक टैम", "BT": "भूटाने दा टैम", "CAST": "CAST", "CAT": "मध्य अफ़्रीका दा टैम", "CCT": "कोकोस टापुआं दे झुंडे दा टैम", "CDT": "उत्तरी अमरिकी केंद्रीय दे ध्याड़े दे उजाले दा टैम", "CHADT": "चैथम दे ध्याड़े दे उजाले दा टैम", "CHAST": "चैथम दा मानक टैम", "CHUT": "चुक दा टैम", "CKT": "कुक टापुआं दे झुंडे दा मानक टैम", "CKT DST": "कुक टापुआं दे झुंडे दा अद्धी तोंदिया दा टैम", "CLST": "चिली दी तोंदिया दा टैम", "CLT": "चिली दा मानक टैम", "COST": "कोलंबिया दी तोंदिया दा टैम", "COT": "कोलंबिया दा मानक टैम", "CST": "उत्तरी अमरिकी केंद्रीय दा मानक टैम", "CST China": "चीन दा मानक टैम", "CST China DST": "चीन दे ध्याड़े दे उजाले दा टैम", "CVST": "केप वर्ड दी तोंदिया दा टैम", "CVT": "केप वर्ड दा मानक टैम", "CXT": "क्रिसमस टापू दा टैम", "ChST": "चामोरो दा मानक टैम", "ChST NMI": "ChST NMI", "CuDT": "क्यूबा दे ध्याड़े दे उजाले दा टैम", "CuST": "क्यूबा दा मानक टैम", "DAVT": "डेविस दा टैम", "DDUT": "ड्यूमोंट डी अर्विले दा टैम", "EASST": "ईस्टर टापू दी तोंदिया दा टैम", "EAST": "ईस्टर टापू दा मानक टैम", "EAT": "पैलेी अफ़्रीका दा टैम", "ECT": "इक्वाडोर दा टैम", "EDT": "उत्तरी अमरिकी पैलेी दे ध्याड़े दे उजाले दा टैम", "EGDT": "पैलेी ग्रीनलैंड दी तोंदिया दा टैम", "EGST": "पैलेी ग्रीनलैंड दा मानक टैम", "EST": "उत्तरी अमरिकी पैलेी दा मानक टैम", "FEET": "अग्र पैलेी यूरोपे दा टैम", "FJT": "फ़िजी दा मानक टैम", "FJT Summer": "फ़िजी दी तोंदिया दा टैम", "FKST": "फ़ॉकलैंड टापुआं दे झुंडे दी तोंदिया दा टैम", "FKT": "फ़ॉकलैंड टापुआं दे झुंडे दा मानक टैम", "FNST": "फ़र्नांर्डो डे नोरोन्हा दे तोंदिया दा टैम", "FNT": "फ़र्नांर्डो डे नोरोन्हा मानक दा टैम", "GALT": "गैलापेगोस दा टैम", "GAMT": "गैंबियर दा टैम", "GEST": "जॉर्जिया दी तोंदिया दा टैम", "GET": "जॉर्जिया दा मानक टैम", "GFT": "फ़्रेंच गुयाना दा टैम", "GIT": "गिल्बर्ट टापुआं दे झुंडे दा टैम", "GMT": "ग्रीनविच मीन टैम", "GNSST": "GNSST", "GNST": "GNST", "GST": "खाड़ी दा मानक टैम", "GST Guam": "GST Guam", "GYT": "गुयाना दा टैम", "HADT": "हवाई–आल्यूशन दे ध्याड़े दे उजाले दा टैम", "HAST": "हवाई–आल्यूशन दा मानक टैम", "HKST": "हाँग काँग दी तोंदिया दा टैम", "HKT": "हाँग काँग दा मानक टैम", "HOVST": "होव्ड दी तोंदिया दा टैम", "HOVT": "होव्ड दा मानक टैम", "ICT": "इंडोचाइना दा टैम", "IDT": "इज़राइल दे ध्याड़े दे उजाले दा टैम", "IOT": "हिंद महासागर दा टैम", "IRKST": "इर्कुत्स्क दी तोंदिया दा टैम", "IRKT": "इर्कुत्स्क दा मानक टैम", "IRST": "ईरान दा मानक टैम", "IRST DST": "ईरान दे ध्याड़े दे उजाले दा टैम", "IST": "भारते दा मानक टैम", "IST Israel": "इज़राइल दा मानक टैम", "JDT": "जापाने दे ध्याड़े दे उजाले दा टैम", "JST": "जापाने दा मानक टैम", "KOST": "कोसराए दा टैम", "KRAST": "क्रास्नोयार्स्क दी तोंदिया दा टैम", "KRAT": "क्रास्नोयार्स्क दा मानक टैम", "KST": "कोरियाई दा मानक टैम", "KST DST": "कोरियाई दे ध्याड़े दे उजाले दा टैम", "LHDT": "लॉर्ड होवे दे ध्याड़े दे उजाले दा टैम", "LHST": "लॉर्ड होवे दा मानक टैम", "LINT": "लाइन टापुआं दे झुंडे दा टैम", "MAGST": "मागादान दी तोंदिया दा टैम", "MAGT": "मागादान दा मानक टैम", "MART": "मार्केसस दा टैम", "MAWT": "माव्सन दा टैम", "MDT": "MDT", "MESZ": "बिचले यूरोपे दी तोंदिया दा टैम", "MEZ": "बिचले यूरोपे दा मानक टैम", "MHT": "मार्शल टापुआं दे झुंडे दा टैम", "MMT": "म्यांमार दा टैम", "MSD": "मॉस्को दि तोंदिया दा टैम", "MST": "MST", "MUST": "मॉरीशस दी तोंदिया दा टैम", "MUT": "मॉरीशस दा मानक टैम", "MVT": "मालदीव दा टैम", "MYT": "मलेशिया दा टैम", "NCT": "न्यू कैलेडोनिया दा मानक टैम", "NDT": "न्यूफ़ाउंडलैंड दे ध्याड़े दे उजाले दा टैम", "NDT New Caledonia": "न्यू कैलेडोनिया दी तोंदिया दा टैम", "NFDT": "नॉरफ़ॉक टापू दे ध्याड़े दे उजाले दा टैम", "NFT": "नॉरफ़ॉक टापू दा मानक टैम", "NOVST": "नोवोसिबिर्स्क दी तोंदिया दा टैम", "NOVT": "नोवोसिबिर्स्क दा मानक टैम", "NPT": "नेपाल दा टैम", "NRT": "नौरू दा टैम", "NST": "न्यूफ़ाउंडलैंड दा मानक टैम", "NUT": "नीयू दा टैम", "NZDT": "न्यूज़ीलैंड दे ध्याड़े दे उजाले दा टैम", "NZST": "न्यूज़ीलैंड दा मानक टैम", "OESZ": "पैलेी यूरोपे दी तोंदिया दा टैम", "OEZ": "पैलेी यूरोपे दा मानक टैम", "OMSST": "ओम्स्क दी तोंदिया दा टैम", "OMST": "ओम्स्क दा मानक टैम", "PDT": "उत्तरी अमेरिकी प्रशांते दे ध्याड़े दे उजाले दा टैम", "PDTM": "मेक्सिकन प्रशांत दे ध्याड़े दे उजाले दा टैम", "PETDT": "PETDT", "PETST": "PETST", "PGT": "पापुआ न्यू गिनी दा टैम", "PHOT": "फ़ीनिक्स टापुआं दे झुंडे दा टैम", "PKT": "पाकिस्ताने दा मानक टैम", "PKT DST": "पाकिस्ताने दी तोंदिया दा टैम", "PMDT": "सेंट पिएरे कने मिक्वेलान दे ध्याड़े दे उजाले दा टैम", "PMST": "सेंट पिएरे कने मिक्वेलान दा मानक टैम", "PONT": "पोनापे दा टैम", "PST": "उत्तरी अमेरिकी प्रशांते दा मानक टैम", "PST Philippine": "फ़िलिपीन दा मानक टैम", "PST Philippine DST": "फ़िलिपीन दी तोंदिया दा टैम", "PST Pitcairn": "पिटकैर्न दा टैम", "PSTM": "मेक्सिकन प्रशांत दा मानक टैम", "PWT": "पलाउ दा टैम", "PYST": "पैराग्वे दी तोंदिया दा टैम", "PYT": "पैराग्वे दा मानक टैम", "PYT Korea": "प्योंगयांग दा टैम", "RET": "रीयूनियन दा टैम", "ROTT": "रोथेरा दा टैम", "SAKST": "सखालिन दी तोंदिया दा टैम", "SAKT": "सखालिन दा मानक टैम", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "दखण अफ़्रीका दा मानक टैम", "SBT": "सोलोमन टापुआं दे झुंडे दा टैम", "SCT": "सेशेल्स दा टैम", "SGT": "सिंगापुरे दा टैम", "SLST": "SLST", "SRT": "सूरीनामे दा टैम", "SST Samoa": "समोआ दा मानक टैम", "SST Samoa Apia": "एपिआ दा मानक टैम", "SST Samoa Apia DST": "एपिआ दे ध्याड़े दे उजाले दा टैम", "SST Samoa DST": "समोआ दे ध्याड़े दे उजाले दा टैम", "SYOT": "स्योवा दा टैम", "TAAF": "दखणे बखें फ़्रांस कने अंटार्कटिक दा टैम", "TAHT": "ताहिती दा टैम", "TJT": "ताजिकिस्ताने दा टैम", "TKT": "टोकेलाऊ दा टैम", "TLT": "पैलेी तिमोरे दा टैम", "TMST": "तुर्कमेनिस्ताने दी तोंदिया दा टैम", "TMT": "तुर्कमेनिस्ताने दा मानक टैम", "TOST": "टोंगा दी तोंदिया दा टैम", "TOT": "टोंगा दा मानक टैम", "TVT": "तुवालू दा टैम", "TWT": "ताइपे दा मानक टैम", "TWT DST": "ताइपे दे ध्याड़े दे उजाले दा टैम", "ULAST": "उलान बटोर दी तोंदिया दा टैम", "ULAT": "उलान बटोर दा मानक टैम", "UYST": "उरुग्वे दी तोंदिया दा टैम", "UYT": "उरुग्वे दा मानक टैम", "UZT": "उज़्बेकिस्तान दा मानक टैम", "UZT DST": "उज़्बेकिस्तान दी तोंदिया दा टैम", "VET": "वेनेज़ुएला दा टैम", "VLAST": "व्लादिवोस्तोक दी तोंदिया दा टैम", "VLAT": "व्लादिवोस्तोक दा मानक टैम", "VOLST": "वोल्गोग्राड दी तोंदिया दा टैम", "VOLT": "वोल्गोग्राड दा मानक टैम", "VOST": "वोस्तोक दा टैम", "VUT": "वनुआतू दा मानक टैम", "VUT DST": "वनुआतू दी तोंदिया दा टैम", "WAKT": "वेक टापू दा टैम", "WARST": "पश्चिमी अर्जेंटीना दी तोंदिया दा टैम", "WART": "पश्चिमी अर्जेंटीना दा मानक टैम", "WAST": "पश्चिम अफ़्रीका दा टैम", "WAT": "पश्चिम अफ़्रीका दा टैम", "WESZ": "पश्चिमी यूरोपे दी तोंदिया दा टैम", "WEZ": "पश्चिमी यूरोपे दा मानक टैम", "WFT": "वालिस कने फ़्यूचूना दा टैम", "WGST": "पश्चिमी ग्रीनलैंड दी तोंदिया दा टैम", "WGT": "पश्चिमी ग्रीनलैंड दा मानक टैम", "WIB": "पश्चिमी इंडोनेशिया दा टैम", "WIT": "पैलेी इंडोनेशिया दा टैम", "WITA": "बिचले इंडोनेशिया दा टैम", "YAKST": "याकुत्स्क दी तोंदिया दा टैम", "YAKT": "याकुत्स्क दा मानक टैम", "YEKST": "येकातेरिनबर्ग दी तोंदिया दा टैम", "YEKT": "येकातेरिनबर्ग दा मानक टैम", "YST": "युकॉन टैम", "МСК": "मॉस्को दा मानक टैम", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "पश्चिम कज़ाखस्तान दा टैम", "شىعىش قازاق ەلى": "पैले कज़ाखस्तान दा टैम", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "किर्गिस्\u200dतान दा टैम", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "पेरू दी तोंदिया दा टैम"},
	}
}

// Locale returns the current translators string locale
func (xnr *xnr) Locale() string {
	return xnr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'xnr'
func (xnr *xnr) PluralsCardinal() []locales.PluralRule {
	return xnr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'xnr'
func (xnr *xnr) PluralsOrdinal() []locales.PluralRule {
	return xnr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'xnr'
func (xnr *xnr) PluralsRange() []locales.PluralRule {
	return xnr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'xnr'
func (xnr *xnr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'xnr'
func (xnr *xnr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'xnr'
func (xnr *xnr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (xnr *xnr) MonthAbbreviated(month time.Month) string {
	if len(xnr.monthsAbbreviated) == 0 {
		return ""
	}
	return xnr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (xnr *xnr) MonthsAbbreviated() []string {
	return xnr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (xnr *xnr) MonthNarrow(month time.Month) string {
	if len(xnr.monthsNarrow) == 0 {
		return ""
	}
	return xnr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (xnr *xnr) MonthsNarrow() []string {
	return xnr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (xnr *xnr) MonthWide(month time.Month) string {
	if len(xnr.monthsWide) == 0 {
		return ""
	}
	return xnr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (xnr *xnr) MonthsWide() []string {
	return xnr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (xnr *xnr) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(xnr.daysAbbreviated) == 0 {
		return ""
	}
	return xnr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (xnr *xnr) WeekdaysAbbreviated() []string {
	return xnr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (xnr *xnr) WeekdayNarrow(weekday time.Weekday) string {
	if len(xnr.daysNarrow) == 0 {
		return ""
	}
	return xnr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (xnr *xnr) WeekdaysNarrow() []string {
	return xnr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (xnr *xnr) WeekdayShort(weekday time.Weekday) string {
	if len(xnr.daysShort) == 0 {
		return ""
	}
	return xnr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (xnr *xnr) WeekdaysShort() []string {
	return xnr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (xnr *xnr) WeekdayWide(weekday time.Weekday) string {
	if len(xnr.daysWide) == 0 {
		return ""
	}
	return xnr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (xnr *xnr) WeekdaysWide() []string {
	return xnr.daysWide
}

// Decimal returns the decimal point of number
func (xnr *xnr) Decimal() string {
	return xnr.decimal
}

// Group returns the group of number
func (xnr *xnr) Group() string {
	return xnr.group
}

// Group returns the minus sign of number
func (xnr *xnr) Minus() string {
	return xnr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'xnr' and handles both Whole and Real numbers based on 'v'
func (xnr *xnr) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, xnr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, xnr.group[0])
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
		b = append(b, xnr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'xnr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (xnr *xnr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, xnr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, xnr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, xnr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'xnr'
func (xnr *xnr) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := xnr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, xnr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, xnr.group[0])
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

	for j := len(xnr.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, xnr.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, xnr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, xnr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'xnr'
// in accounting notation.
func (xnr *xnr) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := xnr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, xnr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, xnr.group[0])
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

		for j := len(xnr.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, xnr.currencyNegativePrefix[j])
		}

		b = append(b, xnr.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(xnr.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, xnr.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, xnr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'xnr'
func (xnr *xnr) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'xnr'
func (xnr *xnr) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, xnr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'xnr'
func (xnr *xnr) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, xnr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'xnr'
func (xnr *xnr) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, xnr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, xnr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'xnr'
func (xnr *xnr) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, xnr.periodsAbbreviated[0]...)
	} else {
		b = append(b, xnr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'xnr'
func (xnr *xnr) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, xnr.periodsAbbreviated[0]...)
	} else {
		b = append(b, xnr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'xnr'
func (xnr *xnr) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, xnr.periodsAbbreviated[0]...)
	} else {
		b = append(b, xnr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'xnr'
func (xnr *xnr) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, xnr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, xnr.periodsAbbreviated[0]...)
	} else {
		b = append(b, xnr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := xnr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
