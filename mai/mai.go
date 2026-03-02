package mai

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mai struct {
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

// New returns a new instance of translator for the 'mai' locale
func New() locales.Translator {
	return &mai{
		locale:             "mai",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "जन॰", "फ़र॰", "मार्च", "अप्रैल", "मई", "जून", "जुल॰", "अग॰", "सित॰", "अक्तू॰", "नव॰", "दिस॰"},
		monthsNarrow:       []string{"", "ज", "फ", "मा", "अ", "म", "जू", "जु", "अ", "सि", "अ", "न", "दि"},
		monthsWide:         []string{"", "जनवरी", "फरवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्तूबर", "नवंबर", "दिसंबर"},
		daysAbbreviated:    []string{"रवि", "सोम", "मंगल", "बुध", "गुरु", "शुक्र", "शनि"},
		daysNarrow:         []string{"र", "सो", "मं", "बु", "गु", "शु", "श"},
		daysWide:           []string{"रवि दिन", "सोम दिन", "मंगल दिन", "बुध दिन", "बृहस्पति दिन", "शुक्र दिन", "शनि दिन"},
		periodsAbbreviated: []string{"भोर", "सांझ"},
		erasAbbreviated:    []string{"", ""},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"", ""},
		timezones:          map[string]string{"ACDT": "आस्ट्रेलियाई मध्य डेलाइट टाइम", "ACST": "आस्ट्रेलियाई मध्य मानक टाइम", "ACT": "ACT", "ACWDT": "आस्ट्रेलियाई मध्य पश्चिमी डेलाइट टाइम", "ACWST": "आस्ट्रेलियाई मध्य पश्चिमी मानक टाइम", "ADT": "अटलांटिक डेलाइट समय", "ADT Arabia": "अरेबियन डेलाइट टाइम", "AEDT": "आस्ट्रेलियाई पूरबी डेलाइट टाइम", "AEST": "आस्ट्रेलियाई पूरबी मानक टाइम", "AFT": "अफगानिस्तान टाइम", "AKDT": "अलास्का डेलाइट समय", "AKST": "अलास्का मानक समय", "AMST": "अमेजन समर टाइम", "AMST Armenia": "आर्मेनिया समर टाइम", "AMT": "अमेजन मानक टाइम", "AMT Armenia": "आर्मेनिया मानक टाइम", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "अर्जेंटीना समर टाइम", "ART": "अर्जेंटीना मानक टाइम", "AST": "अटलांटिक मानक समय", "AST Arabia": "अरेबियन मानक टाइम", "AWDT": "आस्ट्रेलियाई पश्चिमी डेलाइट टाइम", "AWST": "आस्ट्रेलियाई पश्चिमी मानक टाइम", "AZST": "अजरबैजान समर टाइम", "AZT": "अजरबैजान मानक टाइम", "BDT Bangladesh": "बंगलादेश समर टाइम", "BNT": "ब्रूनेई दारेसलाम टाइम", "BOT": "बोलीबिया टाइम", "BRST": "ब्राजीलिया समर टाइम", "BRT": "ब्राजीलिया मानक टाइम", "BST Bangladesh": "बंगलादेश मानक टाइम", "BT": "भूटान टाइम", "CAST": "CAST", "CAT": "मध्य अफ्रीका टाइम", "CCT": "कोकोस द्वीप टाइम", "CDT": "उत्तरी अमेरिकी केंद्रीय डेलाइट समय", "CHADT": "चेथम डेलाइट टाइम", "CHAST": "चेथम मानक टाइम", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "चिली समर टाइम", "CLT": "चिली मानक टाइम", "COST": "कोलंबिया समर टाइम", "COT": "कोलंबिया मानक टाइम", "CST": "उत्तरी अमेरिकी केंद्रीय मानक समय", "CST China": "चीनी मानत टाइम", "CST China DST": "चीनी डेलाइट टाइम", "CVST": "केप बर्ड समर टाइम", "CVT": "केप बर्ड मानक टाइम", "CXT": "क्रिसमस द्वीप टाइम", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "क्यूबा डेलाइट टाइम", "CuST": "क्यूबा मानक टाइम", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "ईस्टर आइलैंड समर टाइम", "EAST": "ईस्टर आइलैंड मानक टाइम", "EAT": "पूरब अफ्रीका टाइम", "ECT": "इक्वाडोर टाइम", "EDT": "उत्तरी अमेरिकी पूर्वी डेलाइट समय", "EGDT": "पूरबी ग्रीनलैंड समर टाइम", "EGST": "पूरबी ग्रीनलैंड मानक टाइम", "EST": "उत्तरी अमेरिकी पूर्वी मानक समय", "FEET": "फरदर-ईस्टर्न यूरोपीयन टाइम", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "फाकलैंड आइलैंड समर टाइम", "FKT": "फाकलैंड आइलैंड मानक टाइम", "FNST": "फर्नेंडो डी नोरोन्हा समर टाइम", "FNT": "फर्नेंडो डी नोरोन्हा मानक टाइम", "GALT": "गैलापैगोस टाइम", "GAMT": "GAMT", "GEST": "जार्जिया समर टाइम", "GET": "जार्जिया मानक टाइम", "GFT": "फ्रेंच गुयाना टाइम", "GIT": "GIT", "GMT": "ग्रीनविच मीन टाइम", "GNSST": "GNSST", "GNST": "GNST", "GST": "दक्षिण जार्जिया टाइम", "GST Guam": "GST Guam", "GYT": "गुयाना टाइम", "HADT": "हवाई-एल्यूटियन डेलाइट टाइम", "HAST": "हवाई-एल्यूटियन मानक टाइम", "HKST": "हांग कांग समर टाइम", "HKT": "हांग कांग मानक टाइम", "HOVST": "होब्द समर टाइम", "HOVT": "होब्द मानक टाइम", "ICT": "इंडोचाइना टाइम", "IDT": "इजरायल डेलाइट टाइम", "IOT": "हिंद महासागर टाइम", "IRKST": "इरकूत्स्क समर टाइम", "IRKT": "इरकूत्स्क मानक टाइम", "IRST": "ईरानमानक टाइम", "IRST DST": "ईरान डेलाइट टाइम", "IST": "भारतीय मानक समय", "IST Israel": "इजरायल मानक टाइम", "JDT": "जापान डेलाइट टाइम", "JST": "जापान मानक टाइम", "KOST": "KOST", "KRAST": "क्रेस्नोयार्स्क समर टाइम", "KRAT": "क्रेस्नोयार्स्क मानक टाइम", "KST": "कोरियन मानक टाइम", "KST DST": "कोरियन डेलाइट टाइम", "LHDT": "लार्ड होबे डेलाइट टाइम", "LHST": "लार्ड होबे मानक टाइम", "LINT": "LINT", "MAGST": "मगादान समर टाइम", "MAGT": "मगादान मानक टाइम", "MART": "MART", "MAWT": "MAWT", "MDT": "उत्तरी अमेरिकी माउंटेन डेलाइट समय", "MESZ": "मध्\u200dय यूरोपीय ग्रीष्\u200dमकालीन समय", "MEZ": "मध्य यूरोपीय मानक समय", "MHT": "MHT", "MMT": "मयनमार टाइम", "MSD": "मास्को समर टाइम", "MST": "उत्तरी अमेरिकी माउंटेन मानक समय", "MUST": "मारीशस समर टाइम", "MUT": "मारीशस मानक टाइम", "MVT": "मालदीव टाइम", "MYT": "मलेशिया टाइम", "NCT": "न्यू केलैडोनिया मानक समय", "NDT": "न्यूफाउंडलैंड डेलाइट टाइम", "NDT New Caledonia": "न्यू केलैडोनिया डेलाइट समय", "NFDT": "फाकलैंड द्वीप डेलाइट टाइम", "NFT": "फाकलैंड द्वीप मानक टाइम", "NOVST": "नोबोसिबिर्स्क समर टाइम", "NOVT": "नोबोसिबिर्स्क मानक टाइम", "NPT": "नेपाल टाइम", "NRT": "NRT", "NST": "न्यूफाउंडलैंड मानक टाइम", "NUT": "NUT", "NZDT": "न्यूजीलैंड डेलाइट टाइम", "NZST": "न्यूजीलैंड मानक टाइम", "OESZ": "पूर्वी यूरोपीय ग्रीष्मकालीन समय", "OEZ": "पूर्वी यूरोपीय मानक समय", "OMSST": "ओम्स्क समर टाइम", "OMST": "ओम्स्क मानक टाइम", "PDT": "उत्तरी अमेरिकी प्रशांत डेलाइट समय", "PDTM": "मैक्सिकन पेसिफिक डेलाइट टाइम", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "पाकिस्तान मानक टाइम", "PKT DST": "पाकिस्तान समर टाइम", "PMDT": "सेंट पिएरे आ मिक्वेलान डेलाइट टाइम", "PMST": "सेंट पिएरे आ मिक्वेलान मानक टाइम", "PONT": "PONT", "PST": "उत्तरी अमेरिकी प्रशांत मानक समय", "PST Philippine": "फिलिपीन मानक टाइम", "PST Philippine DST": "फिलिपीन समर टाइम", "PST Pitcairn": "PST Pitcairn", "PSTM": "मैक्सिकन पेसिफिक मानक टाइम", "PWT": "PWT", "PYST": "पराग्वे समर टाइम", "PYT": "पराग्वे मानक टाइम", "PYT Korea": "प्योंगयांग टाइम", "RET": "रियूनियन टाइम", "ROTT": "ROTT", "SAKST": "सकालिन समर टाइम", "SAKT": "सकालिन मानक टाइम", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "दक्षिण अफ्रीका मानक टाइम", "SBT": "SBT", "SCT": "सेशेल्स टाइम", "SGT": "सिंगापुर मानक टाइम", "SLST": "SLST", "SRT": "सुरीनाम टाइम", "SST Samoa": "सामोआ मानक समय", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "सामोआ समर समय", "SYOT": "SYOT", "TAAF": "फ़्रांसीसी दक्षिणी क्षेत्र आ अंटार्कटिक टाइम", "TAHT": "TAHT", "TJT": "ताजिकिस्तान", "TKT": "TKT", "TLT": "पूरबी तिमोर टाइम", "TMST": "तुर्कमेनिस्तान समर टाइम", "TMT": "तुर्कमेनिस्तान मानक टाइम", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "ताइपेई मानक टाइम", "TWT DST": "ताइपेई डेलाइट टाइम", "ULAST": "उलनबटोर समर टाइम", "ULAT": "उलनबटोर मानक टाइम", "UYST": "उरुग्वे समर टाइम", "UYT": "उरुग्वे मानक टाइम", "UZT": "उजबेकिस्तान मानक टाइम", "UZT DST": "उजबेकिस्तान समर टाइम", "VET": "बेनेजुएला टाइम", "VLAST": "ब्लादिबोस्तोक समर टाइम", "VLAT": "ब्लादिबोस्तोक मानक टाइम", "VOLST": "बोल्गोग्राद समर टाइम", "VOLT": "बोल्गोग्राद मानक टाइम", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "पश्चिमी अर्जेंटीना समर टाइम", "WART": "पश्चिमी अर्जेंटीना मानक टाइम", "WAST": "पश्चिम अफ्रीका टाइम", "WAT": "पश्चिम अफ्रीका टाइम", "WESZ": "पश्चिमी यूरोपीय ग्रीष्\u200dमकालीन समय", "WEZ": "पश्चिमी यूरोपीय मानक समय", "WFT": "WFT", "WGST": "पश्चिमी ग्रीनलैंड समर टाइम", "WGT": "पश्चिमी ग्रीनलैंड मानक टाइम", "WIB": "पश्चिमी इंडोनेशिया टाइम", "WIT": "पूरबी इंडोनेशिया टाइम", "WITA": "मध्य इंडोनेशिया टाइम", "YAKST": "यकुत्स्क समर टाइम", "YAKT": "यकुत्स्क मानक टाइम", "YEKST": "येकाटैरिनबर्ग समर टाइम", "YEKT": "येकाटैरिनबर्ग मानक टाइम", "YST": "यूकोन टाइम", "МСК": "मास्को मानक टाइम", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "पश्चिम कजाखस्तान टाइम", "شىعىش قازاق ەلى": "पूरब कजाखस्तान टाइम", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "किर्गिस्तान", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "एजोर्स समर टाइम"},
	}
}

// Locale returns the current translators string locale
func (mai *mai) Locale() string {
	return mai.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mai'
func (mai *mai) PluralsCardinal() []locales.PluralRule {
	return mai.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mai'
func (mai *mai) PluralsOrdinal() []locales.PluralRule {
	return mai.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mai'
func (mai *mai) PluralsRange() []locales.PluralRule {
	return mai.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mai'
func (mai *mai) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mai'
func (mai *mai) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mai'
func (mai *mai) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mai *mai) MonthAbbreviated(month time.Month) string {
	return mai.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mai *mai) MonthsAbbreviated() []string {
	return mai.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mai *mai) MonthNarrow(month time.Month) string {
	return mai.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mai *mai) MonthsNarrow() []string {
	return mai.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mai *mai) MonthWide(month time.Month) string {
	return mai.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mai *mai) MonthsWide() []string {
	return mai.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mai *mai) WeekdayAbbreviated(weekday time.Weekday) string {
	return mai.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mai *mai) WeekdaysAbbreviated() []string {
	return mai.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mai *mai) WeekdayNarrow(weekday time.Weekday) string {
	return mai.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mai *mai) WeekdaysNarrow() []string {
	return mai.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mai *mai) WeekdayShort(weekday time.Weekday) string {
	return mai.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mai *mai) WeekdaysShort() []string {
	return mai.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mai *mai) WeekdayWide(weekday time.Weekday) string {
	return mai.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mai *mai) WeekdaysWide() []string {
	return mai.daysWide
}

// Decimal returns the decimal point of number
func (mai *mai) Decimal() string {
	return mai.decimal
}

// Group returns the group of number
func (mai *mai) Group() string {
	return mai.group
}

// Group returns the minus sign of number
func (mai *mai) Minus() string {
	return mai.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mai' and handles both Whole and Real numbers based on 'v'
func (mai *mai) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mai' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mai *mai) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mai'
func (mai *mai) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mai.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mai'
// in accounting notation.
func (mai *mai) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mai.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'mai'
func (mai *mai) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'mai'
func (mai *mai) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mai.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mai'
func (mai *mai) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mai.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mai'
func (mai *mai) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mai.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mai.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mai'
func (mai *mai) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mai.periodsAbbreviated[0]...)
	} else {
		b = append(b, mai.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mai'
func (mai *mai) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mai.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mai.periodsAbbreviated[0]...)
	} else {
		b = append(b, mai.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mai'
func (mai *mai) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mai.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mai.periodsAbbreviated[0]...)
	} else {
		b = append(b, mai.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mai'
func (mai *mai) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mai.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mai.periodsAbbreviated[0]...)
	} else {
		b = append(b, mai.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mai.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
