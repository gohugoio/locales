package ne

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ne struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ne' locale
func New() locales.Translator {
	return &ne{
		locale:                 "ne",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "नेरू", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsNarrow:           []string{"", "जन", "फेब", "मार्च", "अप्र", "मे", "जुन", "जुल", "अग", "सेप", "अक्टो", "नोभे", "डिसे"},
		monthsWide:             []string{"", "जनवरी", "फेब्रुअरी", "मार्च", "अप्रिल", "मे", "जुन", "जुलाई", "अगस्ट", "सेप्टेम्बर", "अक्टोबर", "नोभेम्बर", "डिसेम्बर"},
		daysAbbreviated:        []string{"आइत", "सोम", "मङ्गल", "बुध", "बिहि", "शुक्र", "शनि"},
		daysNarrow:             []string{"आ", "सो", "म", "बु", "बि", "शु", "श"},
		daysWide:               []string{"आइतबार", "सोमबार", "मङ्गलबार", "बुधबार", "बिहिबार", "शुक्रबार", "शनिबार"},
		timezones:              map[string]string{"ACDT": "केन्द्रीय अस्ट्रेलिया दिवा समय", "ACST": "ACST", "ACT": "ACT", "ACWDT": "केन्द्रीय पश्चिमी अस्ट्रेलिया दिवा समय", "ACWST": "केन्द्रीय पश्चिमी अस्ट्रेलिया मानक समय", "ADT": "एट्लान्टिक दिवा समय", "ADT Arabia": "अरबी दिवा समय", "AEDT": "पूर्वी अस्ट्रेलिया दिवा समय", "AEST": "पूर्वी अस्ट्रेलिया मानक समय", "AFT": "अफगानिस्तान समय", "AKDT": "अलस्काको दिवा समय", "AKST": "अलस्काको मानक समय", "AMST": "एमाजोन ग्रीष्मकालीन समय", "AMST Armenia": "अर्मेनिया ग्रीष्मकालीन समय", "AMT": "एमाजोन मानक समय", "AMT Armenia": "अर्मेनिया मानक समय", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "अर्जेनटिनी ग्रीष्मकालीन समय", "ART": "अर्जेनटिनी मानक समय", "AST": "एट्लान्टिक मानक समय", "AST Arabia": "अरबी मानक समय", "AWDT": "पश्चिमी अस्ट्रेलिया दिवा समय", "AWST": "पश्चिमी अस्ट्रेलिया मानक समय", "AZST": "अजरबैजान ग्रीष्मकालीन समय", "AZT": "अजरबैजान मानक समय", "BDT Bangladesh": "बंगलादेशी ग्रीष्मकालीन समय", "BNT": "ब्रुनाइ दारूस्सलम समय", "BOT": "बोलिभिया समय", "BRST": "ब्राजिलीया ग्रीष्मकालीन समय", "BRT": "ब्राजिलिया मानक समय", "BST Bangladesh": "बंगलादेशी मानक समय", "BT": "भुटानी समय", "CAST": "CAST", "CAT": "केन्द्रीय अफ्रिकी समय", "CCT": "कोकोस टापु समय", "CDT": "केन्द्रीय दिवा समय", "CHADT": "चाथाम दिवा समय", "CHAST": "चाथाम मानक समय", "CHUT": "चुउक समय", "CKT": "कुक टापु मानक समय", "CKT DST": "कुक टापु आधा ग्रीष्मकालीन समय", "CLST": "चिली ग्रीष्मकालीन समय", "CLT": "चिली मानक समय", "COST": "कोलम्बियाली ग्रीष्मकालीन समय", "COT": "कोलम्बियाली मानक समय", "CST": "केन्द्रीय मानक समय", "CST China": "चीन मानक समय", "CST China DST": "चीन दिवा समय", "CVST": "केप भर्दे ग्रीष्मकालीन समय", "CVT": "केप भर्दे मानक समय", "CXT": "क्रिस्मस टापु समय", "ChST": "चामोर्रो मानक समय", "ChST NMI": "ChST NMI", "CuDT": "क्यूबाको दिवा समय", "CuST": "क्यूबाको मानक समय", "DAVT": "डेभिस समय", "DDUT": "डुमोन्ट-डी‘ उर्भिले समय", "EASST": "इस्टर टापू ग्रीष्म समय", "EAST": "इस्टर टापू मानक समय", "EAT": "पूर्वी अफ्रिकी समय", "ECT": "ईक्वोडोर समय", "EDT": "पूर्वी दिवा समय", "EGDT": "पूर्वी ग्रीनल्यान्डको ग्रीष्मकालीन समय", "EGST": "पूर्वी ग्रीनल्यान्डको मानक समय", "EST": "पूर्वी मानक समय", "FEET": "थप-पूर्वी युरोपेली समय", "FJT": "फिजी मानक समय", "FJT Summer": "फिजी ग्रीष्मकालीन समय", "FKST": "फल्कल्यान्ड टापू ग्रीष्मकालीन समय", "FKT": "फल्कल्यान्ड टापू मानक समय", "FNST": "फर्नान्डो डे नोरोन्हा ग्रीष्मकालीन समय", "FNT": "फर्नान्डो डे नोरोन्हा मानक समय", "GALT": "गालापागोस् समय", "GAMT": "ग्याम्बियर समय", "GEST": "जर्जिया ग्रीष्मकालीन समय", "GET": "जर्जिया मानक समय", "GFT": "फ्रेन्च ग्वाना समय", "GIT": "गिल्बर्ट टापु समय", "GMT": "ग्रीनविच मिन समय", "GNSST": "GNSST", "GNST": "GNST", "GST": "दक्षिण जर्जिया समय", "GST Guam": "GST Guam", "GYT": "गुयाना समय", "HADT": "हवाई-एलुटियन दिवा समय", "HAST": "हवाई-एलुटियन मानक समय", "HKST": "हङकङ ग्रीष्मकालीन समय", "HKT": "हङकङ मानक समय", "HOVST": "होब्ड ग्रीष्मकालीन समय", "HOVT": "होब्ड मानक समय", "ICT": "इन्डोचाइना समय", "IDT": "इजरायल दिवा समय", "IOT": "हिन्द महासागर समय", "IRKST": "ईर्कुट्स्क ग्रीष्मकालीन समय", "IRKT": "ईर्कुट्स्क मानक समय", "IRST": "इरानी मानक समय", "IRST DST": "इरानी दिवा समय", "IST": "भारतीय मानक समय", "IST Israel": "इजरायल मानक समय", "JDT": "जापान दिवा समय", "JST": "जापान मानक समय", "KOST": "कोसराए समय", "KRAST": "क्रासनोयार्क ग्रीष्मकालीन समय", "KRAT": "क्रासनोयार्क मानक समय", "KST": "कोरियाली मानक समय", "KST DST": "कोरियाली दिवा समय", "LHDT": "लर्ड हावे दिवा समय", "LHST": "लर्ड हावे मानक समय", "LINT": "लाइन टापु समय", "MAGST": "मागादान ग्रीष्मकालीन समय", "MAGT": "मागादान मानक समय", "MART": "मार्किसस समय", "MAWT": "म्वसन समय", "MDT": "MDT", "MESZ": "केन्द्रीय युरोपेली ग्रीष्मकालीन समय", "MEZ": "केन्द्रीय युरोपेली मानक समय", "MHT": "मार्शल टापु समय", "MMT": "म्यानमार समय", "MSD": "मस्को ग्रीष्मकालीन समय", "MST": "MST", "MUST": "मउरिटस ग्रीष्मकालीन समय", "MUT": "मउरिटस मानक समय", "MVT": "माल्दिभ्स समय", "MYT": "मलेसिया समय", "NCT": "नयाँ कालेदोनिया मानक समय", "NDT": "न्यूफाउनल्यान्डको दिवा समय", "NDT New Caledonia": "नयाँ कालेदोनिया ग्रीष्मकालीन समय", "NFDT": "नोर्फोक टापूको ग्रीष्मकालीन समय", "NFT": "नोर्फोक टापूको मानक समय", "NOVST": "नोभोसिविर्स्क ग्रीष्मकालीन समय", "NOVT": "नोभोसिविर्स्क मानक समय", "NPT": "नेपाली समय", "NRT": "नाउरु समय", "NST": "न्यूफाउनडल्यान्डको मानक समय", "NUT": "निउए समय", "NZDT": "न्यूजिल्यान्ड दिवा समय", "NZST": "न्यूजिल्यान्ड मानक समय", "OESZ": "पूर्वी युरोपेली ग्रीष्मकालीन समय", "OEZ": "पूर्वी युरोपेली मानक समय", "OMSST": "ओम्स्क ग्रीष्मकालीन समय", "OMST": "ओम्स्क मानक समय", "PDT": "प्यासिफिक दिवा समय", "PDTM": "मेक्सिकन प्यासिफिक दिवा समय", "PETDT": "PETDT", "PETST": "PETST", "PGT": "पपूवा न्यू गिनी समय", "PHOT": "फिनिक्स टापु समय", "PKT": "पाकिस्तानी मानक समय", "PKT DST": "पाकिस्तानी ग्रीष्मकालीन समय", "PMDT": "सेन्ट पियर्रे र मिक्युलोनको दिवा समय", "PMST": "सेन्ट पियर्रे र मिक्युलोनको मानक समय", "PONT": "पोनापे समय", "PST": "प्यासिफिक मानक समय", "PST Philippine": "फिलिपिनी मानक समय", "PST Philippine DST": "फिलिपिनी ग्रीष्मकालीन समय", "PST Pitcairn": "पिटकैरण समय", "PSTM": "मेक्सिकन प्यासिफिक मानक समय", "PWT": "पालाउ समय", "PYST": "पाराग्वे ग्रीष्मकालीन समय", "PYT": "पाराग्वे मानक समय", "PYT Korea": "प्योङयाङ समय", "RET": "रियुनियन समय", "ROTT": "रोथेरा समय", "SAKST": "साखालिन ग्रीष्मकालीन समय", "SAKT": "साखालिन मानक समय", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "दक्षिण अफ्रिकी समय", "SBT": "सोलोमोन टापु समय", "SCT": "सेयेचेलास् समय", "SGT": "सिंगापुर मानक समय", "SLST": "SLST", "SRT": "सुरिनामा समय", "SST Samoa": "सामोअ मानक समय", "SST Samoa Apia": "आपिया मानक समय", "SST Samoa Apia DST": "आपिया दिवा समय", "SST Samoa DST": "सामोअ दिवा समय", "SYOT": "स्योवा समय", "TAAF": "फ्रेन्च दक्षिणी र अन्टार्टिक समय", "TAHT": "ताहिती समय", "TJT": "ताजिकस्तान समय", "TKT": "तोकेलाउ समय", "TLT": "पूर्वी टिमोर समय", "TMST": "तुर्कमेनिस्तान ग्रीष्मकालीन मानक समय", "TMT": "तुर्कमेनिस्तान मानक समय", "TOST": "टोंगा ग्रीष्मकालीन समय", "TOT": "टोंगा मानक समय", "TVT": "टुभालु समय", "TWT": "ताइपेइ मानक समय", "TWT DST": "ताइपेइ दिवा समय", "ULAST": "उलान बाटोर ग्रीष्मकालीन समय", "ULAT": "उलान बाटोर मानक समय", "UYST": "उरुग्वे ग्रीष्मकालीन समय", "UYT": "उरूग्वे मानक समय", "UZT": "उज्बेकिस्तान मानक समय", "UZT DST": "उज्बेकिस्तान ग्रीष्मकालीन समय", "VET": "भेनेज्युएला समय", "VLAST": "भ्लादिभास्टोक ग्रीष्मकालीन समय", "VLAT": "भ्लादिभास्टोक मानक समय", "VOLST": "भोल्गाग्राद ग्रीष्मकालीन समय", "VOLT": "भोल्गाग्राद मानक समय", "VOST": "भास्टोक समय", "VUT": "भानुआतु मानक समय", "VUT DST": "भानुआतु ग्रीष्मकालीन समय", "WAKT": "वेक टापु समय", "WARST": "पश्चिमी अर्जेनटिनी ग्रीष्मकालीन समय", "WART": "पश्चिमी अर्जेनटिनी मानक समय", "WAST": "पश्चिम अफ्रिकी समय", "WAT": "पश्चिम अफ्रिकी समय", "WESZ": "युरोपेली ग्रीष्मकालीन समय", "WEZ": "पश्चिमी युरोपेली मानक समय", "WFT": "वालिस् र फुटुना समय", "WGST": "पश्चिमी ग्रीनल्यान्डको ग्रीष्मकालीन समय", "WGT": "पश्चिमी ग्रीनल्यान्डको मानक समय", "WIB": "पश्चिमी इन्डोनेशिया समय", "WIT": "पूर्वी इन्डोनेशिया समय", "WITA": "केन्द्रीय इन्डोनेशिया समय", "YAKST": "याकुस्ट ग्रीष्मकालीन समय", "YAKT": "याकुस्ट मानक समय", "YEKST": "येकाटेरिनबर्ग ग्रीष्मकालीन समय", "YEKT": "येकाटेरिनबर्ग मानक समय", "YST": "युकोनको समय", "МСК": "मस्को मानक समय", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "पश्चिम काजकस्तान समय", "شىعىش قازاق ەلى": "पूर्वी काजकस्तान समय", "قازاق ەلى": "काजकस्तानको समय", "قىرعىزستان": "किर्गिस्तान समय", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "एजोरेस् ग्रीष्मकालीन समय"},
	}
}

// Locale returns the current translators string locale
func (ne *ne) Locale() string {
	return ne.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ne'
func (ne *ne) PluralsCardinal() []locales.PluralRule {
	return ne.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ne'
func (ne *ne) PluralsOrdinal() []locales.PluralRule {
	return ne.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ne'
func (ne *ne) PluralsRange() []locales.PluralRule {
	return ne.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ne'
func (ne *ne) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ne'
func (ne *ne) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 1 && n <= 4 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ne'
func (ne *ne) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ne.CardinalPluralRule(num1, v1)
	end := ne.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ne *ne) MonthAbbreviated(month time.Month) string {
	if len(ne.monthsAbbreviated) == 0 {
		return ""
	}
	return ne.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ne *ne) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ne *ne) MonthNarrow(month time.Month) string {
	if len(ne.monthsNarrow) == 0 {
		return ""
	}
	return ne.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ne *ne) MonthsNarrow() []string {
	return ne.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ne *ne) MonthWide(month time.Month) string {
	if len(ne.monthsWide) == 0 {
		return ""
	}
	return ne.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ne *ne) MonthsWide() []string {
	return ne.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ne *ne) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ne.daysAbbreviated) == 0 {
		return ""
	}
	return ne.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ne *ne) WeekdaysAbbreviated() []string {
	return ne.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ne *ne) WeekdayNarrow(weekday time.Weekday) string {
	if len(ne.daysNarrow) == 0 {
		return ""
	}
	return ne.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ne *ne) WeekdaysNarrow() []string {
	return ne.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ne *ne) WeekdayShort(weekday time.Weekday) string {
	if len(ne.daysShort) == 0 {
		return ""
	}
	return ne.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ne *ne) WeekdaysShort() []string {
	return ne.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ne *ne) WeekdayWide(weekday time.Weekday) string {
	if len(ne.daysWide) == 0 {
		return ""
	}
	return ne.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ne *ne) WeekdaysWide() []string {
	return ne.daysWide
}

// Decimal returns the decimal point of number
func (ne *ne) Decimal() string {
	return ne.decimal
}

// Group returns the group of number
func (ne *ne) Group() string {
	return ne.group
}

// Group returns the minus sign of number
func (ne *ne) Minus() string {
	return ne.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ne' and handles both Whole and Real numbers based on 'v'
func (ne *ne) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, ne.group[0])
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
		b = append(b, ne.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ne' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ne *ne) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ne.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ne.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ne'
func (ne *ne) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ne.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, ne.group[0])
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

	for j := len(ne.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ne.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, ne.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ne.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ne'
// in accounting notation.
func (ne *ne) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ne.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, ne.group[0])
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

		for j := len(ne.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ne.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, ne.minus[0])

	} else {

		for j := len(ne.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ne.currencyPositivePrefix[j])
		}

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
			b = append(b, ne.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ne'
func (ne *ne) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ne'
func (ne *ne) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ne.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ne'
func (ne *ne) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ne.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ne'
func (ne *ne) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ne.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ne.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ne'
func (ne *ne) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ne'
func (ne *ne) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ne'
func (ne *ne) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ne'
func (ne *ne) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ne.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ne.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
