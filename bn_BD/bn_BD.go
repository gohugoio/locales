package bn_BD

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type bn_BD struct {
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

// New returns a new instance of translator for the 'bn_BD' locale
func New() locales.Translator {
	return &bn_BD{
		locale:                 "bn_BD",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "জানু", "ফেব", "মার্চ", "এপ্রি", "মে", "জুন", "জুল", "আগ", "সেপ", "অক্টো", "নভে", "ডিসে"},
		monthsNarrow:           []string{"", "জা", "ফে", "মা", "এ", "মে", "জুন", "জু", "আ", "সে", "অ", "ন", "ডি"},
		monthsWide:             []string{"", "জানুয়ারি", "ফেব্রুয়ারি", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "অক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		daysAbbreviated:        []string{"রবি", "সোম", "মঙ্গল", "বুধ", "বৃহস্পতি", "শুক্র", "শনি"},
		daysNarrow:             []string{"র", "সো", "ম", "বু", "বৃ", "শু", "শ"},
		daysShort:              []string{"রঃ", "সোঃ", "মঃ", "বুঃ", "বৃঃ", "শুঃ", "শনি"},
		daysWide:               []string{"রবিবার", "সোমবার", "মঙ্গলবার", "বুধবার", "বৃহস্পতিবার", "শুক্রবার", "শনিবার"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "অস্ট্রেলীয় কেন্দ্রীয় দিবালোক সময়", "ACST": "অস্ট্রেলীয় কেন্দ্রীয় মানক সময়", "ACT": "একর মানক সময়", "ACWDT": "অস্ট্রেলীয় কেন্দ্রীয় পশ্চিমি দিবালোক সময়", "ACWST": "অস্ট্রেলীয় কেন্দ্রীয় পশ্চিমি মানক সময়", "ADT": "অতলান্তিক দিবালোক সময়", "ADT Arabia": "আরবি দিবালোক সময়", "AEDT": "অস্ট্রেলীয় পূর্ব দিবালোক সময়", "AEST": "অস্ট্রেলীয় পূর্ব মানক সময়", "AFT": "আফগানিস্তান সময়", "AKDT": "আলাস্কা দিবালোক সময়", "AKST": "আলাস্কা মানক সময়", "AMST": "আমাজন গ্রীষ্মকালীন সময়", "AMST Armenia": "আর্মেনিয়া গ্রীষ্মকালীন সময়", "AMT": "আমাজন মানক সময়", "AMT Armenia": "আর্মেনিয়া মানক সময়", "ANAST": "অনদ্য্র্ গ্রীষ্মকালীন সময়", "ANAT": "অনদ্য্র্ মানক সময়", "ARST": "আর্জেন্টিনা গ্রীষ্মকালীন সময়", "ART": "আর্জেন্টিনা মানক সময়", "AST": "অতলান্তিক মানক সময়", "AST Arabia": "আরবি মানক সময়", "AWDT": "অস্ট্রেলীয় পশ্চিমি দিবালোক সময়", "AWST": "অস্ট্রেলীয় পশ্চিমি মানক সময়", "AZST": "আজারবাইজান গ্রীষ্মকালীন সময়", "AZT": "আজারবাইজান মানক সময়", "BDT Bangladesh": "বাংলাদেশ গ্রীষ্মকালীন সময়", "BNT": "ব্রুনেই দারুসসালাম সময়", "BOT": "বোলিভিয়া সময়", "BRST": "ব্রাসিলিয়া গ্রীষ্মকালীন সময়", "BRT": "ব্রাসিলিয়া মানক সময়", "BST Bangladesh": "বাংলাদেশ মানক সময়", "BT": "ভুটান সময়", "CAST": "CAST", "CAT": "মধ্য আফ্রিকা সময়", "CCT": "কোকোস দ্বীপপুঞ্জ সময়", "CDT": "কেন্দ্রীয় দিবালোক সময়", "CHADT": "চ্যাথাম দিবালোক সময়", "CHAST": "চ্যাথাম মানক সময়", "CHUT": "চুক সময়", "CKT": "কুক দ্বীপপুঞ্জ মানক সময়", "CKT DST": "কুক দ্বীপপুঞ্জ অর্ধেক গ্রীষ্মকালীন সময়", "CLST": "চিলি গ্রীষ্মকালীন সময়", "CLT": "চিলি মানক সময়", "COST": "কলম্বিয়া গ্রীষ্মকালীন সময়", "COT": "কোলোম্বিয়া মানক সময়", "CST": "কেন্দ্রীয় মানক সময়", "CST China": "চীন মানক সময়", "CST China DST": "চীন দিবালোক সময়", "CVST": "কেপ ভার্দ গ্রীষ্মকালীন সময়", "CVT": "কেপ ভার্দ মানক সময়", "CXT": "ক্রিসমাস দ্বীপ সময়", "ChST": "চামেরো মানক সময়", "ChST NMI": "উত্তর মেরিন দ্বীপপুঞ্জ সময়", "CuDT": "কিউবা দিবালোক সময়", "CuST": "কিউবা মানক সময়", "DAVT": "ডেভিস সময়", "DDUT": "ডুমন্ট-দ্য’উরভিলে সময়", "EASST": "ইস্টার দ্বীপ গ্রীষ্মকালীন সময়", "EAST": "ইস্টার দ্বীপ মানক সময়", "EAT": "পূর্ব আফ্রিকা সময়", "ECT": "ইকুয়েডর সময়", "EDT": "পূর্বাঞ্চলের দিবালোক সময়", "EGDT": "পূর্ব গ্রীনল্যান্ড গ্রীষ্মকালীন সময়", "EGST": "পূর্ব গ্রীনল্যান্ড মানক সময়", "EST": "পূর্বাঞ্চলের মানক সময়", "FEET": "প্রান্তীয় পূর্ব ইউরোপীয় সময়", "FJT": "ফিজি মানক সময়", "FJT Summer": "ফিজি গ্রীষ্মকালীন সময়", "FKST": "ফকল্যান্ড দ্বীপপুঞ্জ গ্রীষ্মকালীন সময়", "FKT": "ফকল্যান্ড দ্বীপপুঞ্জ মানক সময়", "FNST": "ফার্নান্দো ডি নোরোনহা গ্রীষ্মকালীন সময়", "FNT": "ফার্নান্দো ডি নোরোনহা মানক সময়", "GALT": "গালাপাগোস সময়", "GAMT": "গ্যাম্বিয়ার সময়", "GEST": "জর্জিয়া গ্রীষ্মকালীন সময়", "GET": "জর্জিয়া মানক সময়", "GFT": "ফরাসি গায়ানা সময়", "GIT": "গিলবার্ট দ্বীপপুঞ্জ সময়", "GMT": "গ্রীনিচ মিন টাইম", "GNSST": "GNSST", "GNST": "GNST", "GST": "দক্ষিণ জর্জিয়া সময়", "GST Guam": "গুয়াম মান সময়", "GYT": "গুয়ানা সময়", "HADT": "হাওয়াই-আলেউত দিবালোক সময়", "HAST": "হাওয়াই-আলেউত মানক সময়", "HKST": "হং কং গ্রীষ্মকালীন সময়", "HKT": "হং কং মানক সময়", "HOVST": "হোভড গ্রীষ্মকালীন সময়", "HOVT": "হোভড মানক সময়", "ICT": "ইন্দোচীন সময়", "IDT": "ইজরায়েল দিবালোক সময়", "IOT": "ভারত মহাসাগরীয় সময়", "IRKST": "ইরকুটস্ক গ্রীষ্মকালীন সময়", "IRKT": "ইরকুটস্ক মানক সময়", "IRST": "ইরান মানক সময়", "IRST DST": "ইরান দিবালোক সময়", "IST": "ভারতীয় মানক সময়", "IST Israel": "ইজরায়েল মানক সময়", "JDT": "জাপান দিবালোক সময়", "JST": "জাপান মানক সময়", "KOST": "কোসরেই সময়", "KRAST": "ক্রাসনোয়ার্স্কি গ্রীষ্মকালীন সময়", "KRAT": "ক্রাসনোয়ার্স্কি মানক সময়", "KST": "কোরিয়ান মানক সময়", "KST DST": "কোরিয়ান দিবালোক সময়", "LHDT": "লর্ড হাওয়ে দিবালোক মসয়", "LHST": "লর্ড হাওয়ে মানক মসয়", "LINT": "লাইন দ্বীপপুঞ্জ সময়", "MAGST": "ম্যাগাডান গ্রীষ্মকালীন সময়", "MAGT": "ম্যাগাডান মানক সময়", "MART": "মার্কেসাস সময়", "MAWT": "মসন সময়", "MDT": "মাকাও গ্রীষ্মকাল সময়", "MESZ": "মধ্য ইউরোপীয় গ্রীষ্মকালীন সময়", "MEZ": "মধ্য ইউরোপীয় মানক সময়", "MHT": "মার্শাল দ্বীপপুঞ্জ সময়", "MMT": "মায়ানমার সময়", "MSD": "মস্কো গ্রীষ্মকালীন সময়", "MST": "মাকাও মান সময়", "MUST": "মরিশাস গ্রীষ্মকালীন সময়", "MUT": "মরিশাস মানক সময়", "MVT": "মালদ্বীপ সময়", "MYT": "মালয়েশিয়া সময়", "NCT": "নিউ ক্যালেডোনিয়া মানক সময়", "NDT": "নিউফাউন্ডল্যান্ড দিবালোক সময়", "NDT New Caledonia": "নিউ ক্যালেডোনিয়া গ্রীষ্মকালীন সময়", "NFDT": "নরফোক দ্বীপ দিবালোক সময়", "NFT": "নরফোক দ্বীপ মানক সময়", "NOVST": "নোভোসিবির্স্ক গ্রীষ্মকালীন সময়", "NOVT": "নোভোসিবির্স্ক মানক সময়", "NPT": "নেপাল সময়", "NRT": "নাউরু সময়", "NST": "নিউফাউন্ডল্যান্ড মানক সময়", "NUT": "নিউই সময়", "NZDT": "নিউজিল্যান্ড দিবালোক সময়", "NZST": "নিউজিল্যান্ড মানক সময়", "OESZ": "পূর্ব ইউরোপীয় গ্রীষ্মকালীন সময়", "OEZ": "পূর্ব ইউরোপীয় মানক সময়", "OMSST": "ওমস্ক গ্রীষ্মকালীন সময়", "OMST": "ওমস্ক মানক সময়", "PDT": "প্রশান্ত মহাসাগরীয় অঞ্চলের দিনের সময়", "PDTM": "মেক্সিকান প্রশান্ত মহাসাগরীয় দিবালোক সময়", "PETDT": "পিত্রেপ্যাভলস্ক- ক্যামচ্যাটস্কি গৃষ্মকালীন সময়", "PETST": "পিত্রেপ্যাভলস্ক- ক্যামচ্যাটস্কি মান সময়", "PGT": "পাপুয়া নিউ গিনি সময়", "PHOT": "ফোনিক্স দ্বীপপুঞ্জ সময়", "PKT": "পাকিস্তান মানক সময়", "PKT DST": "পাকিস্তান গ্রীষ্মকালীন সময়", "PMDT": "সেন্ট পিয়ের ও মিকেলন দিবালোক সময়", "PMST": "সেন্ট পিয়ের ও মিকেলন মানক সময়", "PONT": "পোনাপে সময়", "PST": "প্রশান্ত মহাসাগরীয় অঞ্চলের মানক সময়", "PST Philippine": "ফিলিপাইন মানক সময়", "PST Philippine DST": "ফিলিপাইন গ্রীষ্মকালীন সময়", "PST Pitcairn": "পিটকেয়ার্ন সময়", "PSTM": "মেক্সিকান প্রশান্ত মহসাগরীয় মানক সময়", "PWT": "পালাউ সময়", "PYST": "প্যারাগুয়ে গ্রীষ্মকালীন সময়", "PYT": "প্যারাগুয়ে মানক সময়", "PYT Korea": "পিয়ংইয়াং সময়", "RET": "রিইউনিয়ন সময়", "ROTT": "রথেরা সময়", "SAKST": "সাখালিন গ্রীষ্মকালীন সময়", "SAKT": "সাখালিন মানক সময়", "SAMST": "সামারা গৃষ্মকালীন সময়", "SAMT": "সামারা মান সময়", "SAST": "দক্ষিণ আফ্রিকা মানক সময়", "SBT": "সলোমন দ্বীপপুঞ্জ সময়", "SCT": "সেশেলস সময়", "SGT": "সিঙ্গাপুর মানক সময়", "SLST": "লঙ্কা সময়", "SRT": "সুরিনাম সময়", "SST Samoa": "সামোয়া মানক সময়", "SST Samoa Apia": "অপিয়া মানক সময়", "SST Samoa Apia DST": "অপিয়া দিনের সময়", "SST Samoa DST": "সামোয়া দিবালোক সময়", "SYOT": "সায়োওয়া সময়", "TAAF": "ফরাসি দক্ষিণ এবং আন্টার্কটিক সময়", "TAHT": "তাহিতি সময়", "TJT": "তাজাখাস্তান সময়", "TKT": "টোকেলাউ সময়", "TLT": "পূর্ব টিমর সময়", "TMST": "তুর্কমেনিস্তান গ্রীষ্মকালীন সময়", "TMT": "তুর্কমেনিস্তান মানক সময়", "TOST": "টোঙ্গা গ্রীষ্মকালীন সময়", "TOT": "টোঙ্গা মানক সময়", "TVT": "টুভালু সময়", "TWT": "তাইপেই মানক সময়", "TWT DST": "তাইপেই দিবালোক সময়", "ULAST": "উলান বাতোর গ্রীষ্মকালীন সময়", "ULAT": "উলান বাতোর মানক সময়", "UYST": "উরুগুয়ে গ্রীষ্মকালীন সময়", "UYT": "উরুগুয়ে মানক সময়", "UZT": "উজবেকিস্তান মানক সময়", "UZT DST": "উজবেকিস্তান গ্রীষ্মকালীন সময়", "VET": "ভেনেজুয়েলা সময়", "VLAST": "ভ্লাদিভস্তক গ্রীষ্মকালীন সময়", "VLAT": "ভ্লাদিভস্তক মানক সময়", "VOLST": "ভলগোগ্রাড গ্রীষ্মকালীন সময়", "VOLT": "ভলগোগ্রাড মানক সময়", "VOST": "ভসটক সময়", "VUT": "ভানুয়াতু মানক সময়", "VUT DST": "ভানুয়াতু গ্রীষ্মকালীন সময়", "WAKT": "ওয়েক দ্বীপ সময়", "WARST": "পশ্চিমি আর্জেনটিনা গ্রীষ্মকালীন সময়", "WART": "পশ্চিমি আর্জেনটিনার প্রমাণ সময়", "WAST": "পশ্চিম আফ্রিকা সময়", "WAT": "পশ্চিম আফ্রিকা সময়", "WESZ": "পশ্চিম ইউরোপীয় গ্রীষ্মকালীন সময়", "WEZ": "পশ্চিম ইউরোপীয় মানক সময়", "WFT": "ওয়ালিস এবং ফুটুনা সময়", "WGST": "পশ্চিম গ্রীনল্যান্ড গ্রীষ্মকালীন সময়", "WGT": "পশ্চিম গ্রীনল্যান্ড মানক সময়", "WIB": "পশ্চিমী ইন্দোনেশিয়া সময়", "WIT": "পূর্ব ইন্দোনেশিয়া সময়", "WITA": "কেন্দ্রীয় ইন্দোনেশিয়া সময়", "YAKST": "ইয়াকুটাস্ক গ্রীষ্মকালীন সময়", "YAKT": "ইয়াকুটাস্ক মানক সময়", "YEKST": "ইয়েকাতেরিনবুর্গ গ্রীষ্মকালীন সময়", "YEKT": "ইয়েকাতেরিনবুর্গ মানক সময়", "YST": "ইউকোন সময়", "МСК": "মস্কো মানক সময়", "اقتاۋ": "আকটাও মানক সময়", "اقتاۋ قالاسى": "আকটাও গ্রীষ্মকাল সময়", "اقتوبە": "আকটোব মানক সময়", "اقتوبە قالاسى": "আকটোব গ্রীষ্মকাল সময়", "الماتى": "আল্মাটি মানক সময়", "الماتى قالاسى": "আল্মাটি গ্রীষ্মকাল সময়", "باتىس قازاق ەلى": "পশ্চিম কাজাখাস্তান সময়", "شىعىش قازاق ەلى": "পূর্ব কাজাখাস্তান সময়", "قازاق ەلى": "কাজাখাস্তান সময়", "قىرعىزستان": "কিরগিস্তান সময়", "قىزىلوردا": "কিজিলোর্ডা মান সময়", "قىزىلوردا قالاسى": "কিজিলোর্ডা গ্রীষ্মকাল সময়", "∅∅∅": "এজোরেস গ্রীষ্মকালীন সময়"},
	}
}

// Locale returns the current translators string locale
func (bn *bn_BD) Locale() string {
	return bn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bn_BD'
func (bn *bn_BD) PluralsCardinal() []locales.PluralRule {
	return bn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bn_BD'
func (bn *bn_BD) PluralsOrdinal() []locales.PluralRule {
	return bn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bn_BD'
func (bn *bn_BD) PluralsRange() []locales.PluralRule {
	return bn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bn_BD'
func (bn *bn_BD) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bn_BD'
func (bn *bn_BD) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 5 || n == 7 || n == 8 || n == 9 || n == 10 {
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

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bn_BD'
func (bn *bn_BD) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := bn.CardinalPluralRule(num1, v1)
	end := bn.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bn *bn_BD) MonthAbbreviated(month time.Month) string {
	if len(bn.monthsAbbreviated) == 0 {
		return ""
	}
	return bn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bn *bn_BD) MonthsAbbreviated() []string {
	return bn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bn *bn_BD) MonthNarrow(month time.Month) string {
	if len(bn.monthsNarrow) == 0 {
		return ""
	}
	return bn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bn *bn_BD) MonthsNarrow() []string {
	return bn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bn *bn_BD) MonthWide(month time.Month) string {
	if len(bn.monthsWide) == 0 {
		return ""
	}
	return bn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bn *bn_BD) MonthsWide() []string {
	return bn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bn *bn_BD) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(bn.daysAbbreviated) == 0 {
		return ""
	}
	return bn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bn *bn_BD) WeekdaysAbbreviated() []string {
	return bn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bn *bn_BD) WeekdayNarrow(weekday time.Weekday) string {
	if len(bn.daysNarrow) == 0 {
		return ""
	}
	return bn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bn *bn_BD) WeekdaysNarrow() []string {
	return bn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bn *bn_BD) WeekdayShort(weekday time.Weekday) string {
	if len(bn.daysShort) == 0 {
		return ""
	}
	return bn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bn *bn_BD) WeekdaysShort() []string {
	return bn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bn *bn_BD) WeekdayWide(weekday time.Weekday) string {
	if len(bn.daysWide) == 0 {
		return ""
	}
	return bn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bn *bn_BD) WeekdaysWide() []string {
	return bn.daysWide
}

// Decimal returns the decimal point of number
func (bn *bn_BD) Decimal() string {
	return bn.decimal
}

// Group returns the group of number
func (bn *bn_BD) Group() string {
	return bn.group
}

// Group returns the minus sign of number
func (bn *bn_BD) Minus() string {
	return bn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bn_BD' and handles both Whole and Real numbers based on 'v'
func (bn *bn_BD) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, bn.group[0])
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
		b = append(b, bn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bn_BD' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bn *bn_BD) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bn.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bn_BD'
func (bn *bn_BD) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bn.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, bn.group[0])
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
		b = append(b, bn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bn_BD'
// in accounting notation.
func (bn *bn_BD) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bn.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, bn.group[0])
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

		b = append(b, bn.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
		b = append(b, bn.currencyNegativeSuffix...)
	} else {

		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bn.periodsAbbreviated[0]...)
	} else {
		b = append(b, bn.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bn.periodsAbbreviated[0]...)
	} else {
		b = append(b, bn.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bn.periodsAbbreviated[0]...)
	} else {
		b = append(b, bn.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bn_BD'
func (bn *bn_BD) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bn.periodsAbbreviated[0]...)
	} else {
		b = append(b, bn.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
