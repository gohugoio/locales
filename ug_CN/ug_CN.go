package ug_CN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ug_CN struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ug_CN' locale
func New() locales.Translator {
	return &ug_CN{
		locale:                 "ug_CN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "￥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsWide:             []string{"", "يانۋار", "فېۋرال", "مارت", "ئاپرېل", "ماي", "ئىيۇن", "ئىيۇل", "ئاۋغۇست", "سېنتەبىر", "ئۆكتەبىر", "نويابىر", "دېكابىر"},
		daysAbbreviated:        []string{"يە", "دۈ", "سە", "چا", "پە", "جۈ", "شە"},
		daysNarrow:             []string{"ي", "د", "س", "چ", "پ", "ج", "ش"},
		daysShort:              []string{"ي", "د", "س", "چ", "پ", "ج", "ش"},
		daysWide:               []string{"يەكشەنبە", "دۈشەنبە", "سەيشەنبە", "چارشەنبە", "پەيشەنبە", "جۈمە", "شەنبە"},
		timezones:              map[string]string{"ACDT": "ئاۋسترالىيە ئوتتۇرا قىسىم يازلىق ۋاقتى", "ACST": "ئاۋسترالىيە ئوتتۇرا قىسىم ئۆلچەملىك ۋاقتى", "ACT": "ئاكرې ئۆلچەملىك ۋاقتى", "ACWDT": "ئاۋسترالىيە ئوتتۇرا غەربىي قىسىم يازلىق ۋاقتى", "ACWST": "ئاۋستىرالىيە ئوتتۇرا غەربىي قىسىم ئۆلچەملىك ۋاقتى", "ADT": "ئاتلانتىك ئوكيان يازلىق ۋاقتى", "ADT Arabia": "ئەرەب يازلىق ۋاقتى", "AEDT": "ئاۋسترالىيە شەرقىي قىسىم يازلىق ۋاقتى", "AEST": "ئاۋسترالىيە شەرقىي قىسىم ئۆلچەملىك ۋاقتى", "AFT": "ئافغانىستان ۋاقتى", "AKDT": "ئالياسكا يازلىق ۋاقتى", "AKST": "ئالياسكا ئۆلچەملىك ۋاقتى", "AMST": "ئامازون يازلىق ۋاقتى", "AMST Armenia": "ئەرمېنىيە يازلىق ۋاقتى", "AMT": "ئامازون ئۆلچەملىك ۋاقتى", "AMT Armenia": "ئەرمېنىيە ئۆلچەملىك ۋاقتى", "ANAST": "ئانادىر يازلىق ۋاقتى", "ANAT": "ئانادىر ئۆلچەملىك ۋاقتى", "ARST": "ئارگېنتىنا يازلىق ۋاقتى", "ART": "ئارگېنتىنا ئۆلچەملىك ۋاقتى", "AST": "ئاتلانتىك ئوكيان ئۆلچەملىك ۋاقتى", "AST Arabia": "ئەرەب ئۆلچەملىك ۋاقتى", "AWDT": "ئاۋسترالىيە غەربىي قىسىم يازلىق ۋاقتى", "AWST": "ئاۋسترالىيە غەربىي قىسىم ئۆلچەملىك ۋاقتى", "AZST": "ئەزەربەيجان يازلىق ۋاقتى", "AZT": "ئەزەربەيجان ئۆلچەملىك ۋاقتى", "BDT Bangladesh": "باڭلادىش يازلىق ۋاقتى", "BNT": "بىرۇنىي دارۇسسالام ۋاقتى", "BOT": "بولىۋىيە ۋاقتى", "BRST": "بىرازىلىيە يازلىق ۋاقتى", "BRT": "بىرازىلىيە ئۆلچەملىك ۋاقتى", "BST Bangladesh": "باڭلادىش ئۆلچەملىك ۋاقتى", "BT": "بۇتان ۋاقتى", "CAST": "كاسېي ۋاقتى", "CAT": "ئوتتۇرا ئافرىقا ۋاقتى", "CCT": "كوكۇس ئارىلى ۋاقتى", "CDT": "ئوتتۇرا قىسىم يازلىق ۋاقتى", "CHADT": "چاتام يازلىق ۋاقتى", "CHAST": "چاتام ئۆلچەملىك ۋاقتى", "CHUT": "چۇك ۋاقتى", "CKT": "كۇك ئاراللىرى ئۆلچەملىك ۋاقتى", "CKT DST": "كۇك ئاراللىرى يېرىم يازلىق ۋاقتى", "CLST": "چىلى يازلىق ۋاقتى", "CLT": "چىلى ئۆلچەملىك ۋاقتى", "COST": "كولومبىيە يازلىق ۋاقتى", "COT": "كولومبىيە ئۆلچەملىك ۋاقتى", "CST": "ئوتتۇرا قىسىم ئۆلچەملىك ۋاقتى", "CST China": "جۇڭگو ئۆلچەملىك ۋاقتى", "CST China DST": "جۇڭگو يازلىق ۋاقتى", "CVST": "يېشىل تۇمشۇق يازلىق ۋاقتى", "CVT": "يېشىل تۇمشۇق ئۆلچەملىك ۋاقتى", "CXT": "روژدېستۋو ئارىلى ۋاقتى", "ChST": "چاموررو ئۆلچەملىك ۋاقتى", "ChST NMI": "شىمالىي مارىيانا ئاراللىرى ۋاقتى", "CuDT": "كۇبا يازلىق ۋاقتى", "CuST": "كۇبا ئۆلچەملىك ۋاقتى", "DAVT": "داۋىس ۋاقتى", "DDUT": "دۇمونت-دۇرۋىل ۋاقتى", "EASST": "ئېستېر ئارىلى يازلىق ۋاقتى", "EAST": "پاسكاليا ئارىلى ئۆلچەملىك ۋاقتى", "EAT": "شەرقىي ئافرىقا ۋاقتى", "ECT": "ئېكۋادور ۋاقتى", "EDT": "شەرقىي قىسىم يازلىق ۋاقتى", "EGDT": "شەرقىي گىرېنلاند يازلىق ۋاقتى", "EGST": "شەرقىي گىرېنلاند ئۆلچەملىك ۋاقتى", "EST": "شەرقىي قىسىم ئۆلچەملىك ۋاقتى", "FEET": "FEET", "FJT": "فىجى ئۆلچەملىك ۋاقتى", "FJT Summer": "فىجى يازلىق ۋاقتى", "FKST": "فالكلاند ئاراللىرى يازلىق ۋاقتى", "FKT": "فالكلاند ئاراللىرى ئۆلچەملىك ۋاقتى", "FNST": "فېرناندو-نورونخا يازلىق ۋاقتى", "FNT": "فېرناندو-نورونخا ئۆلچەملىك ۋاقتى", "GALT": "گالاپاگوس ۋاقتى", "GAMT": "گامبىيېر ۋاقتى", "GEST": "گىرۇزىيە يازلىق ۋاقتى", "GET": "گىرۇزىيە ئۆلچەملىك ۋاقتى", "GFT": "فىرانسىيەگە قاراشلىق گىۋىيانا ۋاقتى", "GIT": "گىلبېرت ئاراللىرى ۋاقتى", "GMT": "گىرىنۋىچ ۋاقتى", "GNSST": "GNSST", "GNST": "GNST", "GST": "جەنۇبىي جورجىيە ۋاقتى", "GST Guam": "گۇئام ئۆلچەملىك ۋاقتى", "GYT": "گىۋىيانا ۋاقتى", "HADT": "ھاۋاي-ئالېيۇت يازلىق ۋاقتى", "HAST": "ھاۋاي-ئالېيۇت ئۆلچەملىك ۋاقتى", "HKST": "شياڭگاڭ يازلىق ۋاقتى", "HKT": "شياڭگاڭ ئۆلچەملىك ۋاقتى", "HOVST": "خوۋد يازلىق ۋاقتى", "HOVT": "خوۋد ئۆلچەملىك ۋاقتى", "ICT": "ھىندى چىنى ۋاقتى", "IDT": "ئىسرائىلىيە يازلىق ۋاقتى", "IOT": "ھىندى ئوكيان ۋاقتى", "IRKST": "ئىركۇتسك يازلىق ۋاقتى", "IRKT": "ئىركۇتسك ئۆلچەملىك ۋاقتى", "IRST": "ئىران ئۆلچەملىك ۋاقتى", "IRST DST": "ئىران يازلىق ۋاقتى", "IST": "ھىندىستان ئۆلچەملىك ۋاقتى", "IST Israel": "ئىسرائىلىيە ئۆلچەملىك ۋاقتى", "JDT": "ياپونىيە يازلىق ۋاقتى", "JST": "ياپونىيە ئۆلچەملىك ۋاقتى", "KOST": "كوسرائې ۋاقتى", "KRAST": "كىراسنويارسك يازلىق ۋاقتى", "KRAT": "كىراسنويارسك ئۆلچەملىك ۋاقتى", "KST": "كورىيە ئۆلچەملىك ۋاقتى", "KST DST": "كورىيە يازلىق ۋاقتى", "LHDT": "لورد-خاي يازلىق ۋاقتى", "LHST": "لورد-خاي ئۆلچەملىك ۋاقتى", "LINT": "لاين ئاراللىرى ۋاقتى", "MAGST": "ماگادان يازلىق ۋاقتى", "MAGT": "ماگادان ئۆلچەملىك ۋاقتى", "MART": "ماركىز ۋاقتى", "MAWT": "ماۋسون ۋاقتى", "MDT": "ئاۋمېن يازلىق ۋاقتى", "MESZ": "ئوتتۇرا ياۋروپا يازلىق ۋاقتى", "MEZ": "ئوتتۇرا ياۋروپا ئۆلچەملىك ۋاقتى", "MHT": "مارشال ئاراللىرى ۋاقتى", "MMT": "بىرما ۋاقتى", "MSD": "موسكۋا يازلىق ۋاقتى", "MST": "ئاۋمېن ئۆلچەملىك ۋاقتى", "MUST": "ماۋرىتىئۇس يازلىق ۋاقتى", "MUT": "ماۋرىتىئۇس ئۆلچەملىك ۋاقتى", "MVT": "مالدىۋې ۋاقتى", "MYT": "مالايشىيا ۋاقتى", "NCT": "يېڭى كالېدونىيە ئۆلچەملىك ۋاقتى", "NDT": "نىۋفوئۇنلاند يازلىق ۋاقتى", "NDT New Caledonia": "يېڭى كالېدونىيە يازلىق ۋاقتى", "NFDT": "نورفولك ئاراللىرى يازلىق ۋاقتى", "NFT": "نورفولك ئاراللىرى ئۆلچەملىك ۋاقتى", "NOVST": "نوۋوسىبىرسك يازلىق ۋاقتى", "NOVT": "نوۋوسىبىرسك ئۆلچەملىك ۋاقتى", "NPT": "نېپال ۋاقتى", "NRT": "ناۋرۇ ۋاقتى", "NST": "نىۋفوئۇنلاند ئۆلچەملىك ۋاقتى", "NUT": "نىيۇئې ۋاقتى", "NZDT": "يېڭى زېلاندىيە يازلىق ۋاقتى", "NZST": "يېڭى زېلاندىيە ئۆلچەملىك ۋاقتى", "OESZ": "شەرقىي ياۋروپا يازلىق ۋاقتى", "OEZ": "شەرقىي ياۋروپا ئۆلچەملىك ۋاقتى", "OMSST": "ئومسك يازلىق ۋاقتى", "OMST": "ئومسك ئۆلچەملىك ۋاقتى", "PDT": "تىنچ ئوكيان يازلىق ۋاقتى", "PDTM": "مېكسىكا تىنچ ئوكيان يازلىق ۋاقتى", "PETDT": "پېتروپاۋلوۋسك-كامچاتكسكى يازلىق ۋاقتى", "PETST": "پېتروپاۋلوۋسك-كامچاتكسكى ئۆلچەملىك ۋاقتى", "PGT": "پاپۇئا يېڭى گىۋىنېيەسى ۋاقتى", "PHOT": "فېنىكس ئاراللىرى ۋاقتى", "PKT": "پاكىستان ئۆلچەملىك ۋاقتى", "PKT DST": "پاكىستان يازلىق ۋاقتى", "PMDT": "ساينىت پىئېر ۋە مىكېلون يازلىق ۋاقتى", "PMST": "ساينىت پىئېر ۋە مىكېلون ئۆلچەملىك ۋاقتى", "PONT": "پونپېي ۋاقتى", "PST": "تىنچ ئوكيان ئۆلچەملىك ۋاقتى", "PST Philippine": "فىلىپپىن ئۆلچەملىك ۋاقتى", "PST Philippine DST": "فىلىپپىن يازلىق ۋاقتى", "PST Pitcairn": "پىتكاير ۋاقتى", "PSTM": "مېكسىكا تىنچ ئوكيان ئۆلچەملىك ۋاقتى", "PWT": "پالاۋ ۋاقتى", "PYST": "پاراگۋاي يازلىق ۋاقتى", "PYT": "پاراگۋاي ئۆلچەملىك ۋاقتى", "PYT Korea": "PYT Korea", "RET": "رېئونىيون ۋاقتى", "ROTT": "روتېرا ۋاقتى", "SAKST": "ساخارىن يازلىق ۋاقتى", "SAKT": "ساخارىن ئۆلچەملىك ۋاقتى", "SAMST": "سامارا يازلىق ۋاقتى", "SAMT": "سامارا ئۆلچەملىك ۋاقتى", "SAST": "جەنۇبىي ئافرىقا ئۆلچەملىك ۋاقتى", "SBT": "سولومون ئاراللىرى ۋاقتى", "SCT": "سېيشېل ۋاقتى", "SGT": "سىنگاپور ۋاقتى", "SLST": "سىرى لانكا ۋاقتى", "SRT": "سۇرىنام ۋاقتى", "SST Samoa": "ساموئا ئۆلچەملىك ۋاقتى", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "سەمەرقەنت يازلىق ۋاقتى", "SYOT": "شوۋا ۋاقتى", "TAAF": "فىرانسىيەگە قاراشلىق جەنۇبىي ۋە ئانتاركتىكا ۋاقتى", "TAHT": "تايتى ۋاقتى", "TJT": "تاجىكىستان ۋاقتى", "TKT": "توكېلاۋ ۋاقتى", "TLT": "شەرقىي تىمور ۋاقتى", "TMST": "تۈركمەنىستان يازلىق ۋاقتى", "TMT": "تۈركمەنىستان ئۆلچەملىك ۋاقتى", "TOST": "تونگا يازلىق ۋاقتى", "TOT": "تونگا ئۆلچەملىك ۋاقتى", "TVT": "تۇۋالۇ ۋاقتى", "TWT": "تەيبېي ئۆلچەملىك ۋاقتى", "TWT DST": "تەيبېي يازلىق ۋاقتى", "ULAST": "ئۇلانباتور يازلىق ۋاقتى", "ULAT": "ئۇلانباتور ئۆلچەملىك ۋاقتى", "UYST": "ئۇرۇگۋاي يازلىق ۋاقتى", "UYT": "ئۇرۇگۋاي ئۆلچەملىك ۋاقتى", "UZT": "ئۆزبېكىستان ئۆلچەملىك ۋاقتى", "UZT DST": "ئۆزبېكىستان يازلىق ۋاقتى", "VET": "ۋېنېزۇئېلا ۋاقتى", "VLAST": "ۋىلادىۋوستوك يازلىق ۋاقتى", "VLAT": "ۋىلادىۋوستوك ئۆلچەملىك ۋاقتى", "VOLST": "ۋولگاگراد يازلىق ۋاقتى", "VOLT": "ۋولگاگراد ئۆلچەملىك ۋاقتى", "VOST": "ۋوستوك ۋاقتى", "VUT": "ۋانۇئاتۇ ئۆلچەملىك ۋاقتى", "VUT DST": "ۋانۇئاتۇ يازلىق ۋاقتى", "WAKT": "ۋېيك ئارىلى ۋاقتى", "WARST": "غەربىي ئارگېنتىنا يازلىق ۋاقتى", "WART": "غەربىي ئارگېنتىنا ئۆلچەملىك ۋاقتى", "WAST": "غەربىي ئافرىقا ۋاقتى", "WAT": "غەربىي ئافرىقا ۋاقتى", "WESZ": "غەربىي ياۋروپا يازلىق ۋاقتى", "WEZ": "غەربىي ياۋروپا ئۆلچەملىك ۋاقتى", "WFT": "ۋاللىس ۋە فۇتۇنا ۋاقتى", "WGST": "غەربىي گىرېنلاند يازلىق ۋاقتى", "WGT": "غەربىي گىرېنلاند ئۆلچەملىك ۋاقتى", "WIB": "غەربىي ھىندونېزىيە ۋاقتى", "WIT": "شەرقىي ھىندونېزىيە ۋاقتى", "WITA": "ئوتتۇرا ھىندونېزىيە ۋاقتى", "YAKST": "ياكۇتسك يازلىق ۋاقتى", "YAKT": "ياكۇتسك ئۆلچەملىك ۋاقتى", "YEKST": "يېكاتېرىنبۇرگ يازلىق ۋاقتى", "YEKT": "يېكاتېرىنبۇرگ ئۆلچەملىك ۋاقتى", "YST": "YST", "МСК": "موسكۋا ئۆلچەملىك ۋاقتى", "اقتاۋ": "ئاقتاي ئۆلچەملىك ۋاقتى", "اقتاۋ قالاسى": "ئاقتاي يازلىق ۋاقتى", "اقتوبە": "ئاقتۆبە ئۆلچەملىك ۋاقتى", "اقتوبە قالاسى": "ئاقتۆبە يازلىق ۋاقتى", "الماتى": "ئالمۇتا ئۆلچەملىك ۋاقتى", "الماتى قالاسى": "ئالمۇتا يازلىق ۋاقتى", "باتىس قازاق ەلى": "غەربىي قازاقىستان ۋاقتى", "شىعىش قازاق ەلى": "شەرقىي قازاقىستان ۋاقتى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرغىزىستان ۋاقتى", "قىزىلوردا": "قىزىلئوردا ئۆلچەملىك ۋاقتى", "قىزىلوردا قالاسى": "قىزىلئوردا يازلىق ۋاقتى", "∅∅∅": "ئازور يازلىق ۋاقتى"},
	}
}

// Locale returns the current translators string locale
func (ug *ug_CN) Locale() string {
	return ug.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ug_CN'
func (ug *ug_CN) PluralsCardinal() []locales.PluralRule {
	return ug.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ug_CN'
func (ug *ug_CN) PluralsOrdinal() []locales.PluralRule {
	return ug.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ug_CN'
func (ug *ug_CN) PluralsRange() []locales.PluralRule {
	return ug.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ug_CN'
func (ug *ug_CN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ug_CN'
func (ug *ug_CN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ug_CN'
func (ug *ug_CN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ug.CardinalPluralRule(num1, v1)
	end := ug.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ug *ug_CN) MonthAbbreviated(month time.Month) string {
	if len(ug.monthsAbbreviated) == 0 {
		return ""
	}
	return ug.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ug *ug_CN) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ug *ug_CN) MonthNarrow(month time.Month) string {
	if len(ug.monthsNarrow) == 0 {
		return ""
	}
	return ug.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ug *ug_CN) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (ug *ug_CN) MonthWide(month time.Month) string {
	if len(ug.monthsWide) == 0 {
		return ""
	}
	return ug.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ug *ug_CN) MonthsWide() []string {
	return ug.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ug *ug_CN) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ug.daysAbbreviated) == 0 {
		return ""
	}
	return ug.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ug *ug_CN) WeekdaysAbbreviated() []string {
	return ug.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ug *ug_CN) WeekdayNarrow(weekday time.Weekday) string {
	if len(ug.daysNarrow) == 0 {
		return ""
	}
	return ug.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ug *ug_CN) WeekdaysNarrow() []string {
	return ug.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ug *ug_CN) WeekdayShort(weekday time.Weekday) string {
	if len(ug.daysShort) == 0 {
		return ""
	}
	return ug.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ug *ug_CN) WeekdaysShort() []string {
	return ug.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ug *ug_CN) WeekdayWide(weekday time.Weekday) string {
	if len(ug.daysWide) == 0 {
		return ""
	}
	return ug.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ug *ug_CN) WeekdaysWide() []string {
	return ug.daysWide
}

// Decimal returns the decimal point of number
func (ug *ug_CN) Decimal() string {
	return ug.decimal
}

// Group returns the group of number
func (ug *ug_CN) Group() string {
	return ug.group
}

// Group returns the minus sign of number
func (ug *ug_CN) Minus() string {
	return ug.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ug_CN' and handles both Whole and Real numbers based on 'v'
func (ug *ug_CN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ug.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ug.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ug.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ug_CN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ug *ug_CN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ug.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ug.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ug.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ug_CN'
func (ug *ug_CN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ug.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ug.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ug.group[0])
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
		b = append(b, ug.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ug.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ug_CN'
// in accounting notation.
func (ug *ug_CN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ug.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ug.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ug.group[0])
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

		b = append(b, ug.currencyNegativePrefix[0])

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
			b = append(b, ug.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ug.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, ug.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, ug.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, ug.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = append(b, ug.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ug_CN'
func (ug *ug_CN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ug.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ug.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
