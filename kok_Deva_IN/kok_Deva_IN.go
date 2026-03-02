package kok_Deva_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kok_Deva_IN struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'kok_Deva_IN' locale
func New() locales.Translator {
	return &kok_Deva_IN{
		locale:                 "kok_Deva_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: "LCr",
		currencyNegativeSuffix: "LCr",
		monthsAbbreviated:      []string{"", "जाने", "फेब्रु", "मार्च", "एप्री", "मे", "जून", "जुल", "ऑग", "सप्टें", "ऑक्टो", "नोव्हें", "डिसें"},
		monthsNarrow:           []string{"", "जा", "फे", "मा", "ए", "मे", "जू", "जु", "ऑ", "स", "ऑ", "नो", "डि"},
		monthsWide:             []string{"", "जानेवारी", "फेब्रुवारी", "मार्च", "एप्रील", "मे", "जून", "जुलय", "ऑगस्ट", "सप्टेंबर", "ऑक्टोबर", "नोव्हेंबर", "डिसेंबर"},
		daysNarrow:             []string{"आ", "सो", "मं", "बु", "बि", "शु", "शे"},
		daysShort:              []string{"आ", "सोम", "मंगळार", "बुधवार", "बिरे", "शुक्रार", "शेनवार"},
		daysWide:               []string{"आयतार", "सोमार", "मंगळार", "बुधवार", "बिरेस्तार", "शुक्रार", "शेनवार"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "ऑस्ट्रेलीयन मध्य डेलायट वेळ", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ऑस्ट्रेलीयन मध्य अस्तंत डेलायट वेळ", "ACWST": "ऑस्ट्रेलीयन मध्य अस्तंत प्रमाणित वेळ", "ADT": "अटलांटीक डेलायट वेळ", "ADT Arabia": "अरबी डेलायट वेळ", "AEDT": "ऑस्ट्रेलीयन उदेंत डेलायट वेळ", "AEST": "ऑस्ट्रेलीयन उदेंत प्रमाणित वेळ", "AFT": "अफगानिस्तान वेळ", "AKDT": "अलास्का डेलायट वेळ", "AKST": "अलास्का प्रमाणीत वेळ", "AMST": "अमेझोन ग्रीष्म वेळ", "AMST Armenia": "आर्मेनिया ग्रीष्म वेळ", "AMT": "अमेझोन प्रमाणित वेळ", "AMT Armenia": "आर्मेनिया प्रमाणित वेळ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "अर्जेंटिना ग्रीष्म वेळ", "ART": "अर्जेंटिना प्रमाणित वेळ", "AST": "अटलांटीक प्रमाणित वेळ", "AST Arabia": "अरबी प्रमाणित वेळ", "AWDT": "ऑस्ट्रेलीयन अस्तंत डेलायट वेळ", "AWST": "ऑस्ट्रेलीयन अस्तंत प्रमाणित वेळ", "AZST": "अजरबैजान ग्रीष्म वेळ", "AZT": "अजरबैजान प्रमाणित वेळ", "BDT Bangladesh": "बांगलादेश ग्रीष्म वेळ", "BNT": "ब्रुनेइ दारूस्सलाम वेळ", "BOT": "बोलिव्हिया वेळ", "BRST": "ब्राझिलिया ग्रीष्म वेळ", "BRT": "ब्राझिलिया प्रमाणित वेळ", "BST Bangladesh": "बांगलादेश प्रमाणित वेळ", "BT": "भूतान", "CAST": "CAST", "CAT": "मध्य आफ्रिका वेळ", "CCT": "कोकोस आयलँड वेळ", "CDT": "मध्य डेलायट वेळ", "CHADT": "चॅथम डेलायट वेळ", "CHAST": "चॅथम प्रमाणित वेळ", "CHUT": "चुक वेळ", "CKT": "कूक आयलँड प्रमाणित वेळ", "CKT DST": "कूक आयलँड अर्द ग्रीष्म वेळ", "CLST": "चिली ग्रीष्म वेळ", "CLT": "चिली प्रमाणित वेळ", "COST": "कोलंबिया ग्रीष्म वेळ", "COT": "कोलंबिया प्रमाणित वेळ", "CST": "मध्य प्रमाणित वेळ", "CST China": "चीन प्रमाणित वेळ", "CST China DST": "चीन डेलायट वेळ", "CVST": "केप वर्दे ग्रीष्म वेळ", "CVT": "केप वर्दे प्रमाणित वेळ", "CXT": "क्रिसमस आयलँड वेळ", "ChST": "कॅमोरा प्रमाणित वेळ", "ChST NMI": "ChST NMI", "CuDT": "क्युबा डेलायट वेळ", "CuST": "क्युबा प्रमाणीत वेळ", "DAVT": "डेव्हीस वेळ", "DDUT": "द्युमाँ दूरवीय वेळ", "EASST": "ईस्टर आयलँड ग्रीष्म वेळ", "EAST": "ईस्टर आयलँड प्रमाणित वेळ", "EAT": "उदेंत आफ्रिका वेळ", "ECT": "इक्वेडोर वेळ", "EDT": "उदेंत डेलायट वेळ", "EGDT": "उदेंत ग्रीनलँड ग्रीष्म वेळ", "EGST": "उदेंत ग्रीनलँड प्रमाणीत वेळ", "EST": "उदेंत प्रमाणित वेळ", "FEET": "आनीक-उदेंत युरोपियन वेळ", "FJT": "फिजी प्रमाणित वेळ", "FJT Summer": "फिजी ग्रीष्म वेळ", "FKST": "फॉकलँड आयलँड्स ग्रीष्म वेळ", "FKT": "फॉकलँड आयलँड्स प्रमाणित वेळ", "FNST": "फर्नांडो दी नोरोन्हा ग्रीष्म वेळ", "FNT": "फर्नांडो दी नोरोन्हा प्रमाणित वेळ", "GALT": "गालापागोस वेळ", "GAMT": "गाम्बियर वेळ", "GEST": "जॉर्जिया ग्रीष्म वेळ", "GET": "जॉर्जिया प्रमाणित वेळ", "GFT": "फ्रेंच गयाना वेळ", "GIT": "गिल्बर्ट आयलँड वेळ", "GMT": "ग्रीनविच मध्य वेळ", "GNSST": "GNSST", "GNST": "GNST", "GST": "गल्फ प्रमाणित वेळ", "GST Guam": "GST Guam", "GYT": "गुयाना वेळ", "HADT": "हवाई-अलेयुशिन प्रमाणीत वेळ", "HAST": "हवाई-अलेयुशिन प्रमाणीत वेळ", "HKST": "हाँग काँग ग्रीष्म वेळ", "HKT": "हाँग काँग प्रमाणित वेळ", "HOVST": "होव्हड ग्रीष्म वेळ", "HOVT": "होव्हड प्रमाणित वेळ", "ICT": "इंडोचीन वेळ", "IDT": "इज़राइल डेलायट वेळ", "IOT": "हिंद म्हासागर वेळ", "IRKST": "ईर्कुटस्क ग्रीष्म वेळ", "IRKT": "ईर्कुटस्क प्रमाणित वेळ", "IRST": "इरान प्रमाणित वेळ", "IRST DST": "इरान डेलायट वेळ", "IST": "भारतीय प्रमाणित वेळ", "IST Israel": "इज़राइल प्रमाणित वेळ", "JDT": "जपान डेलायट वेळ", "JST": "जपान प्रमाणित वेळ", "KOST": "कोसरे वेळ", "KRAST": "क्रास्नोयार्स्क ग्रीष्म वेळ", "KRAT": "क्रास्नोयार्स्क प्रमाणित वेळ", "KST": "कोरियन प्रमाणित वेळ", "KST DST": "कोरियन डेलायट वेळ", "LHDT": "लॉर्ड होवे डेलायट वेळ", "LHST": "लॉर्ड होवे प्रमाणित वेळ", "LINT": "लायन आयलँड वेळ", "MAGST": "मगादान ग्रीष्म वेळ", "MAGT": "मगादान प्रमाणित वेळ", "MART": "मार्किसस वेळ", "MAWT": "मॉसन वेळ", "MDT": "पर्वतीय डेलायट वेळ", "MESZ": "मध्य युरोपियन ग्रीष्म वेळ", "MEZ": "मध्य युरोपियन प्रमाणित वेळ", "MHT": "मार्शल आयलँड वेळ", "MMT": "म्यानमार वेळ", "MSD": "मॉस्को ग्रीष्म वेळ", "MST": "पर्वतीय प्रमाणित वेळ", "MUST": "मॉरिशस ग्रीष्म वेळ", "MUT": "मॉरिशस प्रमाणित वेळ", "MVT": "मालदीव वेळ", "MYT": "मलेशिया वेळ", "NCT": "न्यु कॅलेडोनिया प्रमाणित वेळ", "NDT": "न्युफावंडलँड डेलायट वेळ", "NDT New Caledonia": "न्यु कॅलेडोनिया ग्रीष्म वेळ", "NFDT": "नॉरफॉक आयलँड ग्रीष्म वेळ", "NFT": "नॉरफॉक आयलँड प्रमाणित वेळ", "NOVST": "नोवोसिबिर्स्क ग्रीष्म वेळ", "NOVT": "नोवोसिबिर्स्क प्रमाणित वेळ", "NPT": "नेपाळ वेळ", "NRT": "नरू वेळ", "NST": "न्युफावंडलँड प्रमाणीत वेळ", "NUT": "न्युए वेळ", "NZDT": "न्युझीलॅन्ड डेलायट वेळ", "NZST": "न्युझीलॅन्ड प्रमाणित वेळ", "OESZ": "उदेंत युरोपियन ग्रीष्म वेळ", "OEZ": "उदेंत युरोपियन प्रमाणित वेळ", "OMSST": "ओम्स्क ग्रीष्म वेळ", "OMST": "ओम्स्क प्रमाणित वेळ", "PDT": "प्रशांत डेलायट वेळ", "PDTM": "मेक्सिकन प्रशांत डेलायट वेळ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "पापुआ न्यु गिनी वेळ", "PHOT": "फिनिक्स आयलँड वेळ", "PKT": "पाकिस्तान प्रमाणित वेळ", "PKT DST": "पाकिस्तान ग्रीष्म वेळ", "PMDT": "सेंट पियर आनी मिकलान डेलायट वेळ", "PMST": "सेंट पियर आनी मिकलान प्रमाणीत वेळ", "PONT": "पोनेप वेळ", "PST": "प्रशांत प्रमाणित वेळ", "PST Philippine": "फिलिपायन प्रमाणित वेळ", "PST Philippine DST": "फिलिपायन ग्रीष्म वेळ", "PST Pitcairn": "पिटकॅरन वेळ", "PSTM": "मेक्सिकन प्रशांत प्रमाणीत वेळ", "PWT": "पलाऊ वेळ", "PYST": "परागुआ ग्रीष्म वेळ", "PYT": "परागुआ प्रमाणित वेळ", "PYT Korea": "प्योंगयांग वेळ", "RET": "रियुनियन वेळ", "ROTT": "रोथेरा वेळ", "SAKST": "सखलिन ग्रीष्म वेळ", "SAKT": "सखलिन प्रमाणित वेळ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "दक्षिण आफ्रिका प्रमाणित वेळ", "SBT": "सोलोमन आइलँड्स वेळ", "SCT": "सेशेल्स वेळ", "SGT": "सिंगापूर प्रमाणित वेळ", "SLST": "SLST", "SRT": "सुरिनाम वेळ", "SST Samoa": "सामोआ प्रमाणित वेळ", "SST Samoa Apia": "अपिया प्रमाणित वेळ", "SST Samoa Apia DST": "अपिया डेलायट वेळ", "SST Samoa DST": "सामोआ डेलायट वेळ", "SYOT": "स्योवा वेळ", "TAAF": "फ्रेन्च दक्षिण आनी अंटार्क्टिक वेळ", "TAHT": "ताहिती वेळ", "TJT": "तजीकिस्तान वेळ", "TKT": "टोकलाऊ वेळ", "TLT": "उदेंत तिमोर वेळ", "TMST": "तुर्कमेनिस्तान ग्रीष्म वेळ", "TMT": "तुर्कमेनिस्तान प्रमाणित वेळ", "TOST": "टोंगा ग्रीष्म वेळ", "TOT": "टोंगा प्रमाणित वेळ", "TVT": "टुवालू वेळ", "TWT": "तैपेई प्रमाणित वेळ", "TWT DST": "तैपेई डेलायट वेळ", "ULAST": "उलानबतार ग्रीष्म वेळ", "ULAT": "उलानबतार प्रमाणित वेळ", "UYST": "उरुग्वे ग्रीष्म वेळ", "UYT": "उरुग्वे प्रमाणित वेळ", "UZT": "उज़्बेकिस्तान प्रमाणित वेळ", "UZT DST": "उज़्बेकिस्तान ग्रीष्म वेळ", "VET": "वेनेझुएला वेळ", "VLAST": "व्लादिवोस्तोक ग्रीष्म वेळ", "VLAT": "व्लादिवोस्तोक प्रमाणित वेळ", "VOLST": "व्होल्गोग्राड ग्रीष्म वेळ", "VOLT": "व्होल्गोग्राड प्रमाणित वेळ", "VOST": "वोस्तोक वेळ", "VUT": "वनातू प्रमाणित वेळ", "VUT DST": "वनातू ग्रीष्म वेळ", "WAKT": "वैक आयलँड वेळ", "WARST": "अस्तंत अर्जेंटिना ग्रीष्म वेळ", "WART": "अस्तंत अर्जेंटिना प्रमाणित वेळ", "WAST": "अस्तंत आफ्रिका वेळ", "WAT": "अस्तंत आफ्रिका वेळ", "WESZ": "अस्तंत युरोपियन ग्रीष्म वेळ", "WEZ": "अस्तंत युरोपियन प्रमाणित वेळ", "WFT": "वालिस आनी फ्यूचूना वेळ", "WGST": "अस्तंत ग्रीनलँड ग्रीष्म वेळ", "WGT": "अस्तंत ग्रीनलँड प्रमाणीत वेळ", "WIB": "अस्तंत इंडोनेशिया वेळ", "WIT": "उदेंत इंडोनेशिया वेळ", "WITA": "मध्य इंडोनेशिया वेळ", "YAKST": "यकुत्स्क ग्रीष्म वेळ", "YAKT": "यकुत्स्क प्रमाणित वेळ", "YEKST": "येकातेरिनबर्ग ग्रीष्म वेळ", "YEKT": "येकातेरिनबर्ग प्रमाणित वेळ", "YST": "युकॉन वेळ", "МСК": "मॉस्को प्रमाणित वेळ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "अस्तंत कझाकस्तान वेळ", "شىعىش قازاق ەلى": "उदेंत कझाकस्तान वेळ", "قازاق ەلى": "कझाखस्तान वेळ", "قىرعىزستان": "किर्गिज़स्तान वेळ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "अझोरेस ग्रीष्म वेळ"},
	}
}

// Locale returns the current translators string locale
func (kok *kok_Deva_IN) Locale() string {
	return kok.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kok_Deva_IN'
func (kok *kok_Deva_IN) PluralsCardinal() []locales.PluralRule {
	return kok.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kok_Deva_IN'
func (kok *kok_Deva_IN) PluralsOrdinal() []locales.PluralRule {
	return kok.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kok_Deva_IN'
func (kok *kok_Deva_IN) PluralsRange() []locales.PluralRule {
	return kok.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 3 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kok *kok_Deva_IN) MonthAbbreviated(month time.Month) string {
	return kok.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kok *kok_Deva_IN) MonthsAbbreviated() []string {
	return kok.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kok *kok_Deva_IN) MonthNarrow(month time.Month) string {
	return kok.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kok *kok_Deva_IN) MonthsNarrow() []string {
	return kok.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kok *kok_Deva_IN) MonthWide(month time.Month) string {
	return kok.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kok *kok_Deva_IN) MonthsWide() []string {
	return kok.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kok *kok_Deva_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return kok.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kok *kok_Deva_IN) WeekdaysAbbreviated() []string {
	return kok.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kok *kok_Deva_IN) WeekdayNarrow(weekday time.Weekday) string {
	return kok.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kok *kok_Deva_IN) WeekdaysNarrow() []string {
	return kok.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kok *kok_Deva_IN) WeekdayShort(weekday time.Weekday) string {
	return kok.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kok *kok_Deva_IN) WeekdaysShort() []string {
	return kok.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kok *kok_Deva_IN) WeekdayWide(weekday time.Weekday) string {
	return kok.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kok *kok_Deva_IN) WeekdaysWide() []string {
	return kok.daysWide
}

// Decimal returns the decimal point of number
func (kok *kok_Deva_IN) Decimal() string {
	return kok.decimal
}

// Group returns the group of number
func (kok *kok_Deva_IN) Group() string {
	return kok.group
}

// Group returns the minus sign of number
func (kok *kok_Deva_IN) Minus() string {
	return kok.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kok_Deva_IN' and handles both Whole and Real numbers based on 'v'
func (kok *kok_Deva_IN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kok.group[0])
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
		b = append(b, kok.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kok_Deva_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kok *kok_Deva_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kok.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kok.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kok.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, kok.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kok.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kok_Deva_IN'
// in accounting notation.
func (kok *kok_Deva_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kok.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, kok.minus[0])

	} else {
		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, kok.currencyNegativeSuffix...)
	} else {
		b = append(b, kok.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, kok.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kok.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, kok.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kok.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kok_Deva_IN'
func (kok *kok_Deva_IN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kok.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
