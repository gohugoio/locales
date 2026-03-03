package kk_Arab

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kk_Arab struct {
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
	currencyPositiveSuffix string
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'kk_Arab' locale
func New() locales.Translator {
	return &kk_Arab{
		locale:                 "kk_Arab",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "֏", "ANG", "Kz", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "KM", "BAN", "$", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "$", "Bs", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$", "BTN", "BUK", "P", "BYB", "р.", "BYR", "$", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$", "CNH", "CNX", "CN¥", "$", "COU", "₡", "CSD", "CSK", "$", "$", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "kr", "$", "DZD", "ECS", "ECV", "EEK", "E£", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "$", "£", "FRF", "£", "GEK", "₾", "GHC", "GH₵", "£", "GMD", "FG", "GNS", "GQE", "GRD", "Q", "GWE", "GWP", "$", "HK$", "L", "HRD", "kn", "HTG", "Ft", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "kr", "ITL", "$", "JOD", "JP¥", "KES", "⃀", "៛", "CF", "₩", "KRH", "KRO", "₩", "KWD", "$", "₸", "₭", "L£", "Rs", "$", "ЛСЛ", "Lt", "LTT", "LUC", "LUF", "LUL", "Ls", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "Ar", "MGF", "MKD", "MKN", "MLF", "K", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "Rs", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "$", "₦", "NIC", "C$", "NLG", "kr", "Rs", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "Rs", "zł", "PLZ", "PTE", "₲", "QAR", "RHD", "ROL", "lei", "RSD", "₽", "RUR", "RF", "\u20c1", "$", "SCR", "SDD", "SDG", "SDP", "kr", "$", "£", "SIT", "SKK", "SLE", "SLL", "SOS", "$", "SRG", "£", "STD", "Db", "SUR", "SVC", "£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "₺", "$", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "$", "UYW", "UZS", "VEB", "VED", "Bs", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "Cg.", "XDR", "XEU", "XFO", "XFU", "F\u202fCFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "¤", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: " مىڭ",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: " مىڭ",
		monthsAbbreviated:      []string{"", "қаң.", "ақп.", "нау.", "сәу.", "мам.", "мау.", "шіл.", "там.", "қыр.", "қаз.", "қар.", "жел."},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "قاڭتار", "اقپان", "ناۋرىز", "ءساۋىر", "مامىر", "ماۋسىم", "شىلدە", "تامىز", "قىركۇيەك", "قازان", "قاراشا", "جەلتوقسان"},
		daysAbbreviated:        []string{"جەك", "دۇي", "سەي", "سار", "بەي", "جۇم", "سەن"},
		daysNarrow:             []string{"ج", "د", "س", "س", "ب", "ج", "س"},
		daysShort:              []string{"جە", "دۇ", "سە", "سا", "بە", "جۇ", "سن"},
		daysWide:               []string{"جەكسەنبى", "دۇيسەنبى", "سەيسەنبى", "سارسەنبى", "بەيسەنبى", "جۇما", "سەنبى"},
		timezones:              map[string]string{"ACDT": "اۋستراليا جازعى ورتالىق ۋاقىتى", "ACST": "اۋستراليا ستاندارتتى ورتالىق ۋاقىتى", "ACT": "ACT", "ACWDT": "اۋستراليا جازعى ورتالىق-باتىس ۋاقىتى", "ACWST": "اۋستراليا ستاندارتتى ورتالىق-باتىس ۋاقىتى", "ADT": "اتلانتيكا جازعى ۋاقىتى", "ADT Arabia": "ساۋد ارابياسى جازعى ۋاقىتى", "AEDT": "اۋستراليا جازعى شىعىس ۋاقىتى", "AEST": "اۋستراليا ستاندارتتى شىعىس ۋاقىتى", "AFT": "اۋعانستان ۋاقىتى", "AKDT": "الاسكا جازعى ۋاقىتى", "AKST": "الاسكا ستاندارتتى ۋاقىتى", "AMST": "امازون جازعى ۋاقىتى", "AMST Armenia": "ارمەنيا جازعى ۋاقىتى", "AMT": "امازون ستاندارتتى ۋاقىتى", "AMT Armenia": "ارمەنيا ستاندارتتى ۋاقىتى", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ارگەنتينا جازعى ۋاقىتى", "ART": "ارگەنتينا ستاندارتتى ۋاقىتى", "AST": "اتلانتيكا ستاندارتتى ۋاقىتى", "AST Arabia": "ساۋد ارابياسى ستاندارتتى ۋاقىتى", "AWDT": "اۋستراليا جازعى باتىس ۋاقىتى", "AWST": "اۋستراليا ستاندارتتى باتىس ۋاقىتى", "AZST": "ءازىربايجان جازعى ۋاقىتى", "AZT": "ءازىربايجان ستاندارتتى ۋاقىتى", "BDT Bangladesh": "بانگلادەش جازعى ۋاقىتى", "BNT": "برۋنەي-دارۋسسالام ۋاقىتى", "BOT": "بوليۆيا ۋاقىتى", "BRST": "برازيليا جازعى ۋاقىتى", "BRT": "برازيليا ستاندارتتى ۋاقىتى", "BST Bangladesh": "بانگلادەش ستاندارتتى ۋاقىتى", "BT": "بۋتان ۋاقىتى", "CAST": "CAST", "CAT": "ورتالىق افريكا ۋاقىتى", "CCT": "كوكوس ارالدارىنىڭ ۋاقىتى", "CDT": "ولتۇستىك امەريكا جازعى ورتالىق ۋاقىتى", "CHADT": "چاتەم جازعى ۋاقىتى", "CHAST": "چاتەم ستاندارتتى ۋاقىتى", "CHUT": "ترۋك ۋاقىتى", "CKT": "كۋك ارالدارىنىڭ ستاندارتتى ۋاقىتى", "CKT DST": "كۋك ارالدارىنىڭ جازعى ۋاقىتى", "CLST": "چيلي جازعى ۋاقىتى", "CLT": "چيلي ستاندارتتى ۋاقىتى", "COST": "كولۋمبيا جازعى ۋاقىتى", "COT": "كولۋمبيا ستاندارتتى ۋاقىتى", "CST": "سولتۇستىك امەريكا ستاندارتتى ورتالىق ۋاقىتى", "CST China": "قىتاي ستاندارتتى ۋاقىتى", "CST China DST": "قىتاي جازعى ۋاقىتى", "CVST": "كابو-ۆەردە جازعى ۋاقىتى", "CVT": "كابو-ۆەردە ستاندارتتى ۋاقىتى", "CXT": "كريستماس ارالىنىڭ ۋاقىتى", "ChST": "چاموررو ستاندارتتى ۋاقىتى", "ChST NMI": "ChST NMI", "CuDT": "كۋبا جازعى ۋاقىتى", "CuST": "كۋبا ستاندارتتى ۋاقىتى", "DAVT": "دەيۆيس ۋاقىتى", "DDUT": "ديۋمون-ديۋرۆيل ۋاقىتى", "EASST": "پاسحا ارالى جازعى ۋاقىتى", "EAST": "پاسحا ارالى ستاندارتتى ۋاقىتى", "EAT": "شىعىس افريكا ۋاقىتى", "ECT": "ەكۆادور ۋاقىتى", "EDT": "سولتۇستىك امەريكا جازعى شىعىس ۋاقىتى", "EGDT": "شىعىس گرەنلانديا جازعى ۋاقىتى", "EGST": "شىعىس گرەنلانديا ستاندارتتى ۋاقىتى", "EST": "سولتۇستىك امەريكا ستاندارتتى شىعىس ۋاقىتى", "FEET": "قيىر شىعىس ەۋروپا ۋاقىتى", "FJT": "فيجي ستاندارتتى ۋاقىتى", "FJT Summer": "فيجي جازعى ۋاقىتى", "FKST": "فولكلەند ارالدارى جازعى ۋاقىتى", "FKT": "فولكلەند ارالدارى ستاندارتتى ۋاقىتى", "FNST": "فەرناندۋ-دي-نورونيا جازعى ۋاقىتى", "FNT": "فەرناندۋ-دي-نورونيا ستاندارتتى ۋاقىتى", "GALT": "گالاپاگوس ۋاقىتى", "GAMT": "گامبە ۋاقىتى", "GEST": "گرۋزيا جازعى ۋاقىتى", "GET": "گرۋزيا ستاندارتتى ۋاقىتى", "GFT": "فرانتسۋز گۆياناسى ۋاقىتى", "GIT": "گيلبەرت ارالدارىنىڭ ۋاقىتى", "GMT": "گرينۆيچ ۋاقىتى", "GNSST": "GNSST", "GNST": "GNST", "GST": "وڭتۇستىك گەورگيا ۋاقىتى", "GST Guam": "GST Guam", "GYT": "گايانا ۋاقىتى", "HADT": "گاۆاي جانە الەۋت ارالدارى جازعى ۋاقىتى", "HAST": "گاۆاي جانە الەۋت ارالدارى ستاندارتتى ۋاقىتى", "HKST": "حوڭكوڭ جازعى ۋاقىتى", "HKT": "حوڭكوڭ ستاندارتتى ۋاقىتى", "HOVST": "حوۆد جازعى ۋاقىتى", "HOVT": "حوۆد ستاندارتتى ۋاقىتى", "ICT": "ءۇندى-قىتاي ۋاقىتى", "IDT": "يزرايل جازعى ۋاقىتى", "IOT": "ءۇندى مۇحيتى ۋاقىتى", "IRKST": "يركۋتسك جازعى ۋاقىتى", "IRKT": "يركۋتسك ستاندارتتى ۋاقىتى", "IRST": "يران ستاندارتتى ۋاقىتى", "IRST DST": "يران جازعى ۋاقىتى", "IST": "ءۇندىستان ستاندارتتى ۋاقىتى", "IST Israel": "يزرايل ستاندارتتى ۋاقىتى", "JDT": "جاپونيا جازعى ۋاقىتى", "JST": "جاپونيا ستاندارتتى ۋاقىتى", "KOST": "كۋسايە ۋاقىتى", "KRAST": "كراسنويارسك جازعى ۋاقىتى", "KRAT": "كراسنويارسك ستاندارتتى ۋاقىتى", "KST": "كورەيا ستاندارتتى ۋاقىتى", "KST DST": "كورەيا جازعى ۋاقىتى", "LHDT": "لورد-حاۋ جازعى ۋاقىتى", "LHST": "لورد-حاۋ ستاندارتتى ۋاقىتى", "LINT": "لاين ارالدارى ۋاقىتى", "MAGST": "ماگادان جازعى ۋاقىتى", "MAGT": "ماگادان ستاندارتتى ۋاقىتى", "MART": "ماركيز ارالدارى ۋاقىتى", "MAWT": "موۋسون ۋاقىتى", "MDT": "MDT", "MESZ": "ورتالىق ەۋروپا جازعى ۋاقىتى", "MEZ": "ورتالىق ەۋروپا ستاندارتتى ۋاقىتى", "MHT": "مارشال ارالدارى ۋاقىتى", "MMT": "ميانما ۋاقىتى", "MSD": "ماسكەۋ جازعى ۋاقىتى", "MST": "MST", "MUST": "ماۆريكيي جازعى ۋاقىتى", "MUT": "ماۆريكيي ستاندارتتى ۋاقىتى", "MVT": "مالديۆ ارالدارى ۋاقىتى", "MYT": "مالايزيا ۋاقىتى", "NCT": "جاڭا كالەدونيا ستاندارتتى ۋاقىتى", "NDT": "نيۋفاۋندلەند جازعى ۋاقىتى", "NDT New Caledonia": "جاڭا كالەدونيا جازعى ۋاقىتى", "NFDT": "نورفولك ارالى جازعى ۋاقىتى", "NFT": "نورفولك ارالى ستاندارتتى ۋاقىتى", "NOVST": "جاڭاسىبىر جازعى ۋاقىتى", "NOVT": "جاڭاسىبىر ستاندارتتى ۋاقىتى", "NPT": "نەپال ۋاقىتى", "NRT": "ناۋرۋ ۋاقىتى", "NST": "نيۋفاۋندلەند ستاندارتتى ۋاقىتى", "NUT": "نيۋە ۋاقىتى", "NZDT": "جاڭا زەلانديا جازعى ۋاقىتى", "NZST": "جاڭا زەلانديا ستاندارتتى ۋاقىتى", "OESZ": "شىعىس ەۋروپا جازعى ۋاقىتى", "OEZ": "شىعىس ەۋروپا ستاندارتتى ۋاقىتى", "OMSST": "ومبى جازعى ۋاقىتى", "OMST": "ومبى ستاندارتتى ۋاقىتى", "PDT": "سولتۇستىك امەريكا جازعى تىنىق مۇحيتى ۋاقىتى", "PDTM": "مەكسيكا جازعى تىنىق مۇحيت ۋاقىتى", "PETDT": "PETDT", "PETST": "PETST", "PGT": "پاپۋا – جاڭا گۆينەيا ۋاقىتى", "PHOT": "فەنيكس ارالدارى ۋاقىتى", "PKT": "پاكىستان ستاندارتتى ۋاقىتى", "PKT DST": "پاكىستان جازعى ۋاقىتى", "PMDT": "سەن-پەر جانە ميكەلون جازعى ۋاقىتى", "PMST": "سەن-پەر جانە ميكەلون ستاندارتتى ۋاقىتى", "PONT": "پونپەي ۋاقىتى", "PST": "سولتۇستىك امەريكا ستاندارتتى تىنىق مۇحيتى ۋاقىتى", "PST Philippine": "فيليپين ارالدارى ستاندارتتى ۋاقىتى", "PST Philippine DST": "فيليپين ارالدارى جازعى ۋاقىتى", "PST Pitcairn": "پيتكەرن ۋاقىتى", "PSTM": "مەكسيكا ستاندارتتى تىنىق مۇحيت ۋاقىتى", "PWT": "پالاۋ ۋاقىتى", "PYST": "پاراگۆاي جازعى ۋاقىتى", "PYT": "پاراگۆاي ستاندارتتى ۋاقىتى", "PYT Korea": "پحەنيان ۋاقىتى", "RET": "رەيۋنون ۋاقىتى", "ROTT": "روتەرا ۋقىتى", "SAKST": "ساحالين جازعى ۋاقىتى", "SAKT": "ساحالين ستاندارتتى ۋاقىتى", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "وڭتۇستىك افريكا ستاندارتتى ۋاقىتى", "SBT": "سولومون ارالدارى ۋاقىتى", "SCT": "سەيشەل ارالدارى ۋاقىتى", "SGT": "سينگاپۋر ستاندارتتى ۋاقىتى", "SLST": "SLST", "SRT": "سۋرينام ۋاقىتى", "SST Samoa": "ساموا ستاندارتتى ۋاقىتى", "SST Samoa Apia": "اپيا ستاندارتتى ۋاقىتى", "SST Samoa Apia DST": "اپيا جازعى ۋاقىتى", "SST Samoa DST": "ساموا جازعى ۋاقىتى", "SYOT": "سەۆا ۋاقىتى", "TAAF": "فرانتسيانىڭ وڭتۇستىك ايماعى جانە انتاركتيكا ۋاقىتى", "TAHT": "تايتي ۋاقىتى", "TJT": "تاجىكستان ۋاقىتى", "TKT": "توكەلاۋ ۋاقىتى", "TLT": "شىعىس تيمور ۋاقىتى", "TMST": "تۇرىكمەنستان جازعى ۋاقىتى", "TMT": "تۇرىكمەنستان ستاندارتتى ۋاقىتى", "TOST": "تونگا كازعى ۋاقىتى", "TOT": "تونگا ستاندارتتى ۋاقىتى", "TVT": "تۋۆالۋ ۋاقىتى", "TWT": "تايبەي ستاندارتتى ۋاقىتى", "TWT DST": "تايبەي جازعى ۋاقىتى", "ULAST": "ۇلانباتىر جازعى ۋاقىتى", "ULAT": "ۇلانباتىر ستاندارتتى ۋاقىتى", "UYST": "ۋرۋگۆاي جازعى ۋاقىتى", "UYT": "ۋرۋگۆاي ستاندارتتى ۋاقىتى", "UZT": "وزبەكستان ستاندارتتى ۋاقىتى", "UZT DST": "وزبەكستان جازعى ۋاقىتى", "VET": "ۆەنەسۋەلا ۋاقىتى", "VLAST": "ۆلاديۆوستوك جازعى ۋاقىتى", "VLAT": "ۆلاديۆوستوك ستاندارتتى ۋاقىتى", "VOLST": "ۆولگوگراد جازعى ۋاقىتى", "VOLT": "ۆولگوگراد ستاندارتتى ۋاقىتى", "VOST": "ۆوستوك ۋاقىتى", "VUT": "ۆانۋاتۋ ستاندارتتى ۋاقىتى", "VUT DST": "ۆانۋاتۋ جازعى ۋاقىتى", "WAKT": "ۋەيك ارالى ۋاقىتى", "WARST": "باتىس ارگەنتينا جازعى ۋاقىتى", "WART": "باتىس ارگەنتينا ستاندارتتى ۋاقىتى", "WAST": "باتىس افريكا ۋاقىتى", "WAT": "باتىس افريكا ۋاقىتى", "WESZ": "باتىس ەۋروپا جازعى ۋاقىتى", "WEZ": "باتىس ەۋروپا ستاندارتتى ۋاقىتى", "WFT": "ۋولليس جانە فۋتۋنا ۋاقىتى", "WGST": "باتىس گرەنلانديا جازعى ۋاقىتى", "WGT": "باتىس گرەنلانديا ستاندارتتى ۋاقىتى", "WIB": "باتىس يندونەزيا ۋاقىتى", "WIT": "شىعىس يندونەزيا ۋاقىتى", "WITA": "ورتالىق يندونەزيا ۋاقىتى", "YAKST": "ياكۋتسك جازعى ۋاقىتى", "YAKT": "ياكۋتسك ستاندارتتى ۋاقىتى", "YEKST": "ەكاتەرينبۋرگ جازعى ۋاقىتى", "YEKT": "ەكاتەرينبۋرگ ستاندارتتى ۋاقىتى", "YST": "يۋكون ۋاقىتى", "МСК": "ماسكەۋ ستاندارتتى ۋاقىتى", "اقتاۋ": "اقتاۋ ستاندارتتى ۋاقىتى", "اقتاۋ قالاسى": "اقتاۋ جازعى ۋاقىتى", "اقتوبە": "اقتوبە ستاندارتتى ۋاقىتى", "اقتوبە قالاسى": "اقتوبە جازعى ۋاقىتى", "الماتى": "الماتى ستاندارتتى ۋاقىتى", "الماتى قالاسى": "الماتى جازعى ۋاقىتى", "باتىس قازاق ەلى": "باتىس قازاق ەلى ۋاقىتى", "شىعىش قازاق ەلى": "شىعىس قازاق ەلى ۋاقىتى", "قازاق ەلى": "قازاق ەلى ۋاقىتى", "قىرعىزستان": "قىرعىزستان ۋاقىتى", "قىزىلوردا": "قىزىلوردا ستاندارتتى ۋاقىتى", "قىزىلوردا قالاسى": "قىزىلوردا جازعى ۋاقىتى", "∅∅∅": "ازور ارالدارى جازعى ۋاقىتى"},
	}
}

// Locale returns the current translators string locale
func (kk *kk_Arab) Locale() string {
	return kk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kk_Arab'
func (kk *kk_Arab) PluralsCardinal() []locales.PluralRule {
	return kk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kk_Arab'
func (kk *kk_Arab) PluralsOrdinal() []locales.PluralRule {
	return kk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kk_Arab'
func (kk *kk_Arab) PluralsRange() []locales.PluralRule {
	return kk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kk_Arab'
func (kk *kk_Arab) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kk_Arab'
func (kk *kk_Arab) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)

	if (nMod10 == 6) || (nMod10 == 9) || (nMod10 == 0 && n != 0) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kk_Arab'
func (kk *kk_Arab) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := kk.CardinalPluralRule(num1, v1)
	end := kk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kk *kk_Arab) MonthAbbreviated(month time.Month) string {
	if len(kk.monthsAbbreviated) == 0 {
		return ""
	}
	return kk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kk *kk_Arab) MonthsAbbreviated() []string {
	return kk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kk *kk_Arab) MonthNarrow(month time.Month) string {
	if len(kk.monthsNarrow) == 0 {
		return ""
	}
	return kk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kk *kk_Arab) MonthsNarrow() []string {
	return kk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kk *kk_Arab) MonthWide(month time.Month) string {
	if len(kk.monthsWide) == 0 {
		return ""
	}
	return kk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kk *kk_Arab) MonthsWide() []string {
	return kk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kk *kk_Arab) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(kk.daysAbbreviated) == 0 {
		return ""
	}
	return kk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kk *kk_Arab) WeekdaysAbbreviated() []string {
	return kk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kk *kk_Arab) WeekdayNarrow(weekday time.Weekday) string {
	if len(kk.daysNarrow) == 0 {
		return ""
	}
	return kk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kk *kk_Arab) WeekdaysNarrow() []string {
	return kk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kk *kk_Arab) WeekdayShort(weekday time.Weekday) string {
	if len(kk.daysShort) == 0 {
		return ""
	}
	return kk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kk *kk_Arab) WeekdaysShort() []string {
	return kk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kk *kk_Arab) WeekdayWide(weekday time.Weekday) string {
	if len(kk.daysWide) == 0 {
		return ""
	}
	return kk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kk *kk_Arab) WeekdaysWide() []string {
	return kk.daysWide
}

// Decimal returns the decimal point of number
func (kk *kk_Arab) Decimal() string {
	return kk.decimal
}

// Group returns the group of number
func (kk *kk_Arab) Group() string {
	return kk.group
}

// Group returns the minus sign of number
func (kk *kk_Arab) Minus() string {
	return kk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kk_Arab' and handles both Whole and Real numbers based on 'v'
func (kk *kk_Arab) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kk_Arab' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kk *kk_Arab) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kk_Arab'
func (kk *kk_Arab) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kk.currencies[currency]
	l := len(s) + len(symbol) + 12

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(kk.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kk.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kk.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kk_Arab'
// in accounting notation.
func (kk *kk_Arab) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kk.currencies[currency]
	l := len(s) + len(symbol) + 12

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(kk.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kk.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, kk.minus[0])

	} else {

		for j := len(kk.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kk.currencyPositivePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, kk.currencyNegativeSuffix...)
	} else {

		b = append(b, kk.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d, 0x20}...)
	b = append(b, kk.monthsAbbreviated[t.Month()]...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d, 0x20}...)
	b = append(b, kk.monthsWide[t.Month()]...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d, 0x20}...)
	b = append(b, kk.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = append(b, kk.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kk_Arab'
func (kk *kk_Arab) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
