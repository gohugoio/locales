package my

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type my struct {
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'my' locale
func New() locales.Translator {
	return &my{
		locale:                 "my",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "NAf", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "Afl", "AZM", "AZN", "BAD", "BAM", "BAN", "Bds$", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "B$", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "G", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "K", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "B/.", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TT$", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ဇန်", "ဖေ", "မတ်", "ဧ", "မေ", "ဇွန်", "ဇူ", "ဩ", "စက်", "အောက်", "နို", "ဒီ"},
		monthsNarrow:           []string{"", "ဇ", "ဖ", "မ", "ဧ", "မ", "ဇ", "ဇ", "ဩ", "စ", "အ", "န", "ဒ"},
		monthsWide:             []string{"", "ဇန်နဝါရီ", "ဖေဖော်ဝါရီ", "မတ်", "ဧပြီ", "မေ", "ဇွန်", "ဇူလိုင်", "ဩဂုတ်", "စက်တင်ဘာ", "အောက်တိုဘာ", "နိုဝင်ဘာ", "ဒီဇင်ဘာ"},
		daysNarrow:             []string{"တ", "တ", "အ", "ဗ", "က", "သ", "စ"},
		daysShort:              []string{"နွေ", "လာ", "ဂါ", "ဟူး", "တေး", "ကြာ", "နေ"},
		daysWide:               []string{"တနင်္ဂနွေ", "တနင်္လာ", "အင်္ဂါ", "ဗုဒ္ဓဟူး", "ကြာသပတေး", "သောကြာ", "စနေ"},
		periodsAbbreviated:     []string{"နံနက်", "ညနေ"},
		erasAbbreviated:        []string{"ဘီစီ", "အဒေီ"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ခရစ်တော် မပေါ်မီနှစ်", "ခရစ်နှစ်"},
		timezones:              map[string]string{"ACDT": "ဩစတြေးလျ အလယ်ပိုင်း နွေရာသီ အချိန်", "ACST": "ACST", "ACT": "ACT", "ACWDT": "သြစတြေးလျား အနောက်အလယ်ပိုင်း နွေရာသီ အချိန်", "ACWST": "သြစတြေးလျား အနောက်အလယ်ပိုင်း စံတော်ချိန်", "ADT": "အတ္တလန်တစ် နွေရာသီ စံတော်ချိန်", "ADT Arabia": "အာရေဗျ နွေရာသီ အချိန်", "AEDT": "အရှေ့ဩစတြေးလျ နွေရာသီ အချိန်", "AEST": "အရှေ့ဩစတြေးလျ စံတော်ချိန်", "AFT": "အာဖဂန်နစ္စတန် အချိန်", "AKDT": "အလာစကာ နွေရာသီ စံတော်ချိန်", "AKST": "အလာစကာ စံတော်ချိန်", "AMST": "အမေဇုံ နွေရာသီအချိန်", "AMST Armenia": "အာမေးနီးယား နွေရာသီ အချိန်", "AMT": "အမေဇုံ စံတော်ချိန်", "AMT Armenia": "အာမေးနီးယား စံတော်ချိန်", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "အာဂျင်တီးနား နွေရာသီအချိန်", "ART": "အာဂျင်တီးနား စံတော်ချိန်", "AST": "အတ္တလန်တစ် စံတော်ချိန်", "AST Arabia": "အာရေဗျ စံတော်ချိန်", "AWDT": "ဩစတြေးလျ နွေရာသီ အချိန်", "AWST": "အနောက်ဩစတြေးလျ စံတော်ချိန်", "AZST": "အဇာဘိုင်ဂျန် နွေရာသီ အချိန်", "AZT": "အဇာဘိုင်ဂျန် စံတော်ချိန်", "BDT Bangladesh": "ဘင်္ဂလားဒေ့ရှ် နွေရာသီ အချိန်", "BNT": "ဘရူနိုင်း စံတော်ချိန်", "BOT": "ဘိုလီးဘီးယား အချိန်", "BRST": "ဘရာဇီး နွေရာသီ အချိန်", "BRT": "ဘရာဇီး စံတော်ချိန်", "BST Bangladesh": "ဘင်္ဂလားဒေ့ရှ် စံတော်ချိန်", "BT": "ဘူတန် အချိန်", "CAST": "CAST", "CAT": "အလယ်အာဖရိက အချိန်", "CCT": "ကိုကိုးကျွန်း အချိန်", "CDT": "အလယ်ပိုင်း နွေရာသီစံတော်ချိန်", "CHADT": "ချားသမ် နွေရာသီ အချိန်", "CHAST": "ချားသမ်စံတော်ချိန်", "CHUT": "ချုခ် အချိန်", "CKT": "ကွတ်ကျွန်းစု စံတော်ချိန်", "CKT DST": "ကွတ်ကျွန်းစု နွေရာသီ အချိန်", "CLST": "ချီလီ နွေရာသီ အချိန်", "CLT": "ချီလီ စံတော်ချိန်", "COST": "ကိုလံဘီယာ နွေရာသီ အချိန်", "COT": "ကိုလံဘီယာ စံတော်ချိန်", "CST": "အလယ်ပိုင်းစံတော်ချိန်", "CST China": "တရုတ် စံတော်ချိန်", "CST China DST": "တရုတ် နွေရာသီ အချိန်", "CVST": "ကိတ်ဗာဒီ နွေရာသီ အချိန်", "CVT": "ကိတ်ဗာဒီ စံတော်ချိန်", "CXT": "ခရစ်စမတ်ကျွန်း အချိန်", "ChST": "ချာမိုရို အချိန်", "ChST NMI": "ChST NMI", "CuDT": "ကျူးဘား နွေရာသီ စံတော်ချိန်", "CuST": "ကျူးဘား စံတော်ချိန်", "DAVT": "ဒေးဗစ် အချိန်", "DDUT": "ဒူးမော့တ် ဒါရ်ဗီးလ် အချိန်", "EASST": "အီစတာကျွန်း နွေရာသီ အချိန်", "EAST": "အီစတာကျွန်း စံတော်ချိန်", "EAT": "အရှေ့အာဖရိက အချိန်", "ECT": "အီကွေဒေါ အချိန်", "EDT": "အရှေ့ပိုင်း နွေရာသီစံတော်ချိန်", "EGDT": "အရှေ့ဂရင်းလန် နွေရာသီ စံတော်ချိန်", "EGST": "အရှေ့ဂရင်းလန်း စံတော်ချိန်", "EST": "အရှေ့ပိုင်းစံတော်ချိန်", "FEET": "အရှေ့ဖျား ဥရောပဒေသ အချိန်", "FJT": "ဖီဂျီ စံတော်ချိန်", "FJT Summer": "ဖီဂျီ နွေရာသီ အချိန်", "FKST": "ဖော့ကလန်ကျွန်းစု နွေရာသီ အချိန်", "FKT": "ဖော့ကလန်ကျွန်းစု စံတော်ချိန်", "FNST": "ဖာနန်ဒိုးဒီနိုးရိုးညာ နွေရာသီ အချိန်", "FNT": "ဖာနန်ဒိုးဒီနိုးရိုးညာ စံတော်ချိန်", "GALT": "ဂါလားပါဂိုးစ် အချိန်", "GAMT": "ဂမ်ဘီယာ အချိန်", "GEST": "ဂျော်ဂျီယာ နွေရာသီ အချိန်", "GET": "ဂျော်ဂျီယာ စံတော်ချိန်", "GFT": "ပြင်သစ် ဂီအားနား အချိန်", "GIT": "ဂီလ်ဘတ်ကျွန်းစု အချိန်", "GMT": "ဂရင်းနစ် စံတော်ချိန်", "GNSST": "GNSST", "GNST": "GNST", "GST": "တောင်ဂျော်ဂျီယာ အချိန်", "GST Guam": "GST Guam", "GYT": "ဂိုင်ယာနာ အချိန်", "HADT": "ဟာဝိုင်ယီ အယ်လူးရှန်း နွေရာသီ စံတော်ချိန်", "HAST": "ဟာဝိုင်ယီ အယ်လူးရှန်း စံတော်ချိန်", "HKST": "ဟောင်ကောင် နွေရာသီ အချိန်", "HKT": "ဟောင်ကောင် စံတော်ချိန်", "HOVST": "ဟိုးဗ် နွေရာသီ အချိန်", "HOVT": "ဟိုးဗ် စံတော်ချိန်", "ICT": "အင်ဒိုချိုင်းနား အချိန်", "IDT": "အစ္စရေး နွေရာသီ အချိန်", "IOT": "အိန္ဒိယသမုဒ္ဒရာ အချိန်", "IRKST": "အီရူခူတ် နွေရာသီ အချိန်", "IRKT": "အီရူခူတ် စံတော်ချိန်", "IRST": "အီရန် စံတော်ချိန်", "IRST DST": "အီရန် နွေရာသီ အချိန်", "IST": "အိန္ဒိယ စံတော်ချိန်", "IST Israel": "အစ္စရေး စံတော်ချိန်", "JDT": "ဂျပန် နွေရာသီ အချိန်", "JST": "ဂျပန် စံတော်ချိန်", "KOST": "ခိုစ်ရိုင် အချိန်", "KRAST": "ခရာ့စ်နိုရာစ် နွေရာသီ အချိန်", "KRAT": "ခရာ့စ်နိုရာစ် စံတော်ချိန်", "KST": "ကိုရီးယား စံတော်ချိန်", "KST DST": "ကိုရီးယား နွေရာသီ အချိန်", "LHDT": "လော့ဒ်ဟောင် နွေရာသီ အချိန်", "LHST": "လော့ဒ်ဟောင် စံတော်ချိန်", "LINT": "လိုင်းကျွန်းစု အချိန်", "MAGST": "မာဂါဒန်း နွေရာသီ အချိန်", "MAGT": "မာဂါဒန်း စံတော်ချိန်", "MART": "မာခေးအပ်စ် အချိန်", "MAWT": "မော်စွန် အချိန်", "MDT": "တောင်တန်း နွေရာသီစံတော်ချိန်", "MESZ": "ဥရောပအလယ်ပိုင်း နွေရာသီ အချိန်", "MEZ": "ဥရောပအလယ်ပိုင်း စံတော်ချိန်", "MHT": "မာရှယ်ကျွန်းစု အချိန်", "MMT": "မြန်မာ အချိန်", "MSD": "မော်စကို နွေရာသီ အချိန်", "MST": "တောင်တန်းစံတော်ချိန်", "MUST": "မောရစ်ရှ နွေရာသီ အချိန်", "MUT": "မောရစ်ရှ စံတော်ချိန်", "MVT": "မော်လဒိုက် အချိန်", "MYT": "မလေးရှား အချိန်", "NCT": "နယူးကာလီဒိုးနီးယား စံတော်ချိန်", "NDT": "နယူးဖောင်လန် နွေရာသီ စံတော်ချိန်", "NDT New Caledonia": "နယူးကာလီဒိုးနီးယား နွေရာသီ အချိန်", "NFDT": "နောဖော့ခ်ကျွန်း နွေရာသီ စံတော်ချိန်", "NFT": "နောဖော့ခ်ကျွန်း စံတော်ချိန်", "NOVST": "နိုဗိုစဲဘီအဲယ်စ် နွေရာသီ အချိန်", "NOVT": "နိုဗိုစဲဘီအဲယ်စ် စံတော်ချိန်", "NPT": "နီပေါ အချိန်", "NRT": "နာဥူရူ အချိန်", "NST": "နယူးဖောင်လန် စံတော်ချိန်", "NUT": "နီဥူအေ အချိန်", "NZDT": "နယူးဇီလန် နွေရာသီ အချိန်", "NZST": "နယူးဇီလန် စံတော်ချိန်", "OESZ": "အရှေ့ဥရောပ နွေရာသီ အချိန်", "OEZ": "အရှေ့ဥရောပ စံတော်ချိန်", "OMSST": "အွမ်းစ်ခ် နွေရာသီ အချိန်", "OMST": "အွမ်းစ်ခ် စံတော်ချိန်", "PDT": "ပစိဖိတ် နွေရာသီစံတော်ချိန်", "PDTM": "မက္ကစီကန် ပစိဖိတ် နွေရာသီ စံတော်ချိန်", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ပါပူအာနယူးဂီနီ အချိန်", "PHOT": "ဖီးနစ်ကျွန်းစု အချိန်", "PKT": "ပါကစ္စတန် စံတော်ချိန်", "PKT DST": "ပါကစ္စတန် နွေရာသီ အချိန်", "PMDT": "စိန့်ပီအဲနှင့် မီခွီလွန် နွေရာသီ စံတော်ချိန်", "PMST": "စိန့်ပီအဲနှင့်မီခွီလွန်စံတော်ချိန်", "PONT": "ဖိုနာဖဲအ် အချိန်", "PST": "ပစိဖိတ်စံတော်ချိန်", "PST Philippine": "ဖိလစ်ပိုင် စံတော်ချိန်", "PST Philippine DST": "ဖိလစ်ပိုင် နွေရာသီ အချိန်", "PST Pitcairn": "ပါတ်ကယ်ရင် အချိန်", "PSTM": "မက္ကဆီကန် ပစိဖိတ် စံတော်ချိန်", "PWT": "ပလာအို အချိန်", "PYST": "ပါရာဂွေး နွေရာသီ အချိန်", "PYT": "ပါရာဂွေး စံတော်ချိန်", "PYT Korea": "ပြုံယန်း အချိန်", "RET": "ရီယူနီယံ အချိန်", "ROTT": "ရိုသီရာ အချိန်", "SAKST": "ဆာခါလင် နွေရာသီ အချိန်", "SAKT": "ဆာခါလင် စံတော်ချိန်", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "တောင်အာဖရိက အချိန်", "SBT": "ဆော်လမွန်ကျွန်းစု အချိန်", "SCT": "ဆေးရှဲ အချိန်", "SGT": "စင်္ကာပူ အချိန်", "SLST": "သီရိလင်္ကာ အချိန်", "SRT": "စူးရီနာမ်အချိန်", "SST Samoa": "ဆမိုအာ စံတော်ချိန်", "SST Samoa Apia": "အပီယာ စံတော်ချိန်", "SST Samoa Apia DST": "အပီယာ နွေရာသီ အချိန်", "SST Samoa DST": "ဆမိုးအား နွေရာသီ အချိန်", "SYOT": "ရှိုဝါ အချိန်", "TAAF": "ပြင်သစ်တောင်ပိုင်းနှင့် အန္တာတိတ် အချိန်", "TAHT": "တဟီတီ အချိန်", "TJT": "တာဂျစ်ကစ္စတန် အချိန်", "TKT": "တိုကီလာဥ အချိန်", "TLT": "အရှေ့တီမော အချိန်", "TMST": "တာ့ခ်မင်နစ္စတန် နွေရာသီ အချိန်", "TMT": "တာ့ခ်မင်နစ္စတန် စံတော်ချိန်", "TOST": "တွန်ဂါ နွေရာသီ အချိန်", "TOT": "တွန်ဂါ စံတော်ချိန်", "TVT": "တူဗားလူ အချိန်", "TWT": "ထိုင်ပေ စံတော်ချိန်", "TWT DST": "ထိုင်ပေ နွေရာသီ အချိန်", "ULAST": "ဥလန်ဘာတော နွေရာသီ အချိန်", "ULAT": "ဥလန်ဘာတော စံတော်ချိန်", "UYST": "ဥရုဂွေး နွေရာသီ အချိန်", "UYT": "ဥရုဂွေး စံတော်ချိန်", "UZT": "ဥဇဘက်ကစ္စတန် စံတော်ချိန်", "UZT DST": "ဥဇဘက်ကစ္စတန် နွေရာသီ အချိန်", "VET": "ဗင်နီဇွဲလား အချိန်", "VLAST": "ဗလာဒီဗော့စတော့ခ် နွေရာသီ အချိန်", "VLAT": "ဗလာဒီဗော့စတော့ခ် စံတော်ချိန်", "VOLST": "ဗိုဂိုဂရက် နွေရာသီ အချိန်", "VOLT": "ဗိုလ်ဂိုဂရက် စံတော်ချိန်", "VOST": "ဗိုစ်တိုခ် အချိန်", "VUT": "ဗနွားတူ စံတော်ချိန်", "VUT DST": "ဗနွားတူ နွေရာသီ အချိန်", "WAKT": "ဝိတ်ခ်ကျွန်း အချိန်", "WARST": "အနောက် အာဂျင်တီးနား နွေရာသီ အချိန်", "WART": "အနောက် အာဂျင်တီးနား စံတော်ချိန်", "WAST": "အနောက်အာဖရိက အချိန်", "WAT": "အနောက်အာဖရိက အချိန်", "WESZ": "အနောက်ဥရောပ နွေရာသီ အချိန်", "WEZ": "အနောက်ဥရောပ စံတော်ချိန်", "WFT": "ဝေါလီစ်နှင့် ဖူကျူနာ အချိန်", "WGST": "အနောက် ဂရင်းလန် နွေရာသီ စံတော်ချိန်", "WGT": "အနောက် ဂရင်းလန်း စံတော်ချိန်", "WIB": "အနောက်ပိုင်း အင်ဒိုနီးရှား အချိန်", "WIT": "အရှေ့ပိုင်း အင်ဒိုနီးရှား အချိန်", "WITA": "အလယ်ပိုင်း အင်ဒိုနီးရှား အချိန်", "YAKST": "ယူခူးတ်စ် နွေရာသီ အချိန်", "YAKT": "ယူခူးတ်စ် စံတော်ချိန်", "YEKST": "ရယ်ခါးတီရင်ဘာခ် နွေရာသီ အချိန်", "YEKT": "ရယ်ခါးတီရင်ဘားခ် စံတော်ချိန်", "YST": "ယူကွန်း အချိန်", "МСК": "မော်စကို စံတော်ချိန်", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "အနောက်ကာဇက်စတန် အချိန်", "شىعىش قازاق ەلى": "အရှေ့ကာဇက်စတန် အချိန်", "قازاق ەلى": "ကာဇက်စတန် အချိန်", "قىرعىزستان": "ကာဂျစ္စတန် အချိန်", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ပီရူး နွေရာသီ အချိန်"},
	}
}

// Locale returns the current translators string locale
func (my *my) Locale() string {
	return my.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'my'
func (my *my) PluralsCardinal() []locales.PluralRule {
	return my.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'my'
func (my *my) PluralsOrdinal() []locales.PluralRule {
	return my.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'my'
func (my *my) PluralsRange() []locales.PluralRule {
	return my.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'my'
func (my *my) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'my'
func (my *my) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'my'
func (my *my) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (my *my) MonthAbbreviated(month time.Month) string {
	return my.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (my *my) MonthsAbbreviated() []string {
	return my.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (my *my) MonthNarrow(month time.Month) string {
	return my.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (my *my) MonthsNarrow() []string {
	return my.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (my *my) MonthWide(month time.Month) string {
	return my.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (my *my) MonthsWide() []string {
	return my.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (my *my) WeekdayAbbreviated(weekday time.Weekday) string {
	return my.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (my *my) WeekdaysAbbreviated() []string {
	return my.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (my *my) WeekdayNarrow(weekday time.Weekday) string {
	return my.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (my *my) WeekdaysNarrow() []string {
	return my.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (my *my) WeekdayShort(weekday time.Weekday) string {
	return my.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (my *my) WeekdaysShort() []string {
	return my.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (my *my) WeekdayWide(weekday time.Weekday) string {
	return my.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (my *my) WeekdaysWide() []string {
	return my.daysWide
}

// Decimal returns the decimal point of number
func (my *my) Decimal() string {
	return my.decimal
}

// Group returns the group of number
func (my *my) Group() string {
	return my.group
}

// Group returns the minus sign of number
func (my *my) Minus() string {
	return my.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'my' and handles both Whole and Real numbers based on 'v'
func (my *my) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'my' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (my *my) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'my'
func (my *my) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := my.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, my.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, my.group[0])
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

	for j := len(my.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, my.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, my.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, my.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'my'
// in accounting notation.
func (my *my) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := my.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, my.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, my.group[0])
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

		for j := len(my.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, my.currencyNegativePrefix[j])
		}

		b = append(b, my.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(my.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, my.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, my.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'my'
func (my *my) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'my'
func (my *my) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, my.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'my'
func (my *my) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, my.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'my'
func (my *my) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, my.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, my.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'my'
func (my *my) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'my'
func (my *my) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, my.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'my'
func (my *my) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x20}...)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, my.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'my'
func (my *my) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	tz, _ := t.Zone()

	if btz, ok := my.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x20}...)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, my.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}
