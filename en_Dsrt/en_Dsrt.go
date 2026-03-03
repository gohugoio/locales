package en_Dsrt

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type en_Dsrt struct {
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

// New returns a new instance of translator for the 'en_Dsrt' locale
func New() locales.Translator {
	return &en_Dsrt{
		locale:                 "en_Dsrt",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "֏", "ANG", "Kz", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "KM", "BAN", "$", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "$", "Bs", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$", "BTN", "BUK", "P", "BYB", "BYN", "BYR", "$", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$", "CNH", "CNX", "CN¥", "$", "COU", "₡", "CSD", "CSK", "$", "$", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "kr", "$", "DZD", "ECS", "ECV", "EEK", "E£", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "$", "£", "FRF", "£", "GEK", "₾", "GHC", "GH₵", "£", "GMD", "FG", "GNS", "GQE", "GRD", "Q", "GWE", "GWP", "$", "HK$", "L", "HRD", "kn", "HTG", "Ft", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "kr", "ITL", "$", "JOD", "JP¥", "KES", "⃀", "៛", "CF", "₩", "KRH", "KRO", "₩", "KWD", "$", "₸", "₭", "L£", "Rs", "$", "LSL", "Lt", "LTT", "LUC", "LUF", "LUL", "Ls", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "Ar", "MGF", "MKD", "MKN", "MLF", "K", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "Rs", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "$", "₦", "NIC", "C$", "NLG", "kr", "Rs", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "Rs", "zł", "PLZ", "PTE", "₲", "QAR", "RHD", "ROL", "lei", "RSD", "₽", "RUR", "RF", "\\u20c1", "$", "SCR", "SDD", "SDG", "SDP", "kr", "$", "£", "SIT", "SKK", "SLE", "SLL", "SOS", "$", "SRG", "£", "STD", "Db", "SUR", "SVC", "£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "₺", "$", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "$", "UYW", "UZS", "VEB", "VED", "Bs", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "Cg.", "XDR", "XEU", "XFO", "XFU", "F\\u202fCFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "¤", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "𐐖𐐰𐑌", "𐐙𐐯𐐺", "𐐣𐐪𐑉", "𐐁𐐹𐑉", "𐐣𐐩", "𐐖𐐭𐑌", "𐐖𐐭𐑊", "𐐂𐑀", "𐐝𐐯𐐹", "𐐉𐐿𐐻", "𐐤𐐬𐑂", "𐐔𐐨𐑅"},
		monthsNarrow:           []string{"", "𐐖", "𐐙", "𐐣", "𐐁", "𐐣", "𐐖", "𐐖", "𐐂", "𐐝", "𐐉", "𐐤", "𐐔"},
		monthsWide:             []string{"", "𐐖𐐰𐑌𐐷𐐭𐐯𐑉𐐨", "𐐙𐐯𐐺𐑉𐐭𐐯𐑉𐐨", "𐐣𐐪𐑉𐐽", "𐐁𐐹𐑉𐐮𐑊", "𐐣𐐩", "𐐖𐐭𐑌", "𐐖𐐭𐑊𐐴", "𐐂𐑀𐐲𐑅𐐻", "𐐝𐐯𐐹𐐻𐐯𐑋𐐺𐐲𐑉", "𐐉𐐿𐐻𐐬𐐺𐐲𐑉", "𐐤𐐬𐑂𐐯𐑋𐐺𐐲𐑉", "𐐔𐐨𐑅𐐯𐑋𐐺𐐲𐑉"},
		daysAbbreviated:        []string{"𐐝𐐲𐑌", "𐐣𐐲𐑌", "𐐓𐐭𐑆", "𐐎𐐯𐑌", "𐐛𐐲𐑉", "𐐙𐑉𐐴", "𐐝𐐰𐐻"},
		daysNarrow:             []string{"𐐝", "𐐣", "𐐓", "𐐎", "𐐛", "𐐙", "𐐝"},
		daysShort:              []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
		daysWide:               []string{"𐐝𐐲𐑌𐐼𐐩", "𐐣𐐲𐑌𐐼𐐩", "𐐓𐐭𐑆𐐼𐐩", "𐐎𐐯𐑌𐑆𐐼𐐩", "𐐛𐐲𐑉𐑆𐐼𐐩", "𐐙𐑉𐐴𐐼𐐩", "𐐝𐐰𐐻𐐲𐑉𐐼𐐩"},
		timezones:              map[string]string{"ACDT": "Australian Central Daylight Time", "ACST": "Australian Central Standard Time", "ACT": "Acre Standard Time", "ACWDT": "Australian Central Western Daylight Time", "ACWST": "Australian Central Western Standard Time", "ADT": "𐐈𐐻𐑊𐐰𐑌𐐻𐐮𐐿 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "ADT Arabia": "Arabian Daylight Time", "AEDT": "Australian Eastern Daylight Time", "AEST": "Australian Eastern Standard Time", "AFT": "Afghanistan Time", "AKDT": "𐐊𐑊𐐰𐑅𐐿𐐲 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "AKST": "𐐊𐑊𐐰𐑅𐐿𐐲 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "AMST": "Amazon Summer Time", "AMST Armenia": "Armenia Summer Time", "AMT": "Amazon Standard Time", "AMT Armenia": "Armenia Standard Time", "ANAST": "Anadyr Summer Time", "ANAT": "Anadyr Standard Time", "ARST": "Argentina Summer Time", "ART": "Argentina Standard Time", "AST": "𐐈𐐻𐑊𐐰𐑌𐐻𐐮𐐿 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "AST Arabia": "Arabian Standard Time", "AWDT": "Australian Western Daylight Time", "AWST": "Australian Western Standard Time", "AZST": "Azerbaijan Summer Time", "AZT": "Azerbaijan Standard Time", "BDT Bangladesh": "Bangladesh Summer Time", "BNT": "Brunei Time", "BOT": "Bolivia Time", "BRST": "Brasilia Summer Time", "BRT": "Brasilia Standard Time", "BST Bangladesh": "Bangladesh Standard Time", "BT": "Bhutan Time", "CAST": "Casey Time", "CAT": "Central Africa Time", "CCT": "Cocos Islands Time", "CDT": "𐐝𐐯𐑌𐐻𐑉𐐲𐑊 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "CHADT": "Chatham Daylight Time", "CHAST": "Chatham Standard Time", "CHUT": "Chuuk Time", "CKT": "Cook Islands Standard Time", "CKT DST": "Cook Islands Summer Time", "CLST": "Chile Summer Time", "CLT": "Chile Standard Time", "COST": "Colombia Summer Time", "COT": "Colombia Standard Time", "CST": "𐐝𐐯𐑌𐐻𐑉𐐲𐑊 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "CST China": "China Standard Time", "CST China DST": "China Daylight Time", "CVST": "Cape Verde Summer Time", "CVT": "Cape Verde Standard Time", "CXT": "Christmas Island Time", "ChST": "Chamorro Standard Time", "ChST NMI": "Northern Mariana Islands Time", "CuDT": "Cuba Daylight Time", "CuST": "Cuba Standard Time", "DAVT": "Davis Time", "DDUT": "Dumont d’Urville Time", "EASST": "Easter Island Summer Time", "EAST": "Easter Island Standard Time", "EAT": "East Africa Time", "ECT": "Ecuador Time", "EDT": "𐐀𐑅𐐻𐐲𐑉𐑌 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "EGDT": "East Greenland Summer Time", "EGST": "East Greenland Standard Time", "EST": "𐐀𐑅𐐻𐐲𐑉𐑌 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "FEET": "Further-eastern European Time", "FJT": "Fiji Standard Time", "FJT Summer": "Fiji Summer Time", "FKST": "Falkland Islands Summer Time", "FKT": "Falkland Islands Standard Time", "FNST": "Fernando de Noronha Summer Time", "FNT": "Fernando de Noronha Standard Time", "GALT": "Galapagos Time", "GAMT": "Gambier Time", "GEST": "Georgia Summer Time", "GET": "Georgia Standard Time", "GFT": "French Guiana Time", "GIT": "Gilbert Islands Time", "GMT": "Greenwich Mean Time", "GNSST": "Greenland Summer Time", "GNST": "Greenland Standard Time", "GST": "South Georgia Time", "GST Guam": "Guam Standard Time", "GYT": "Guyana Time", "HADT": "Hawaii-Aleutian Daylight Time", "HAST": "Hawaii-Aleutian Standard Time", "HKST": "𐐐𐐱𐑍 𐐗𐐱𐑍 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "HKT": "𐐐𐐱𐑍 𐐗𐐱𐑍 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "HOVST": "Khovd Summer Time", "HOVT": "Khovd Standard Time", "ICT": "Indochina Time", "IDT": "Israel Daylight Time", "IOT": "Indian Ocean Time", "IRKST": "Irkutsk Summer Time", "IRKT": "Irkutsk Standard Time", "IRST": "Iran Standard Time", "IRST DST": "Iran Daylight Time", "IST": "India Standard Time", "IST Israel": "Israel Standard Time", "JDT": "Japan Daylight Time", "JST": "Japan Standard Time", "KOST": "Kosrae Time", "KRAST": "Krasnoyarsk Summer Time", "KRAT": "Krasnoyarsk Standard Time", "KST": "Korean Standard Time", "KST DST": "Korean Daylight Time", "LHDT": "Lord Howe Daylight Time", "LHST": "Lord Howe Standard Time", "LINT": "Line Islands Time", "MAGST": "Magadan Summer Time", "MAGT": "Magadan Standard Time", "MART": "Marquesas Time", "MAWT": "Mawson Time", "MDT": "𐐣𐐵𐑌𐐻𐐲𐑌 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "MESZ": "Central European Summer Time", "MEZ": "Central European Standard Time", "MHT": "Marshall Islands Time", "MMT": "Myanmar Time", "MSD": "Moscow Summer Time", "MST": "𐐣𐐵𐑌𐐻𐐲𐑌 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "MUST": "Mauritius Summer Time", "MUT": "Mauritius Standard Time", "MVT": "Maldives Time", "MYT": "Malaysia Time", "NCT": "New Caledonia Standard Time", "NDT": "𐐤𐐭𐑁𐐲𐑌𐐼𐑊𐐲𐑌𐐼 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "NDT New Caledonia": "New Caledonia Summer Time", "NFDT": "Norfolk Island Daylight Time", "NFT": "Norfolk Island Standard Time", "NOVST": "Novosibirsk Summer Time", "NOVT": "Novosibirsk Standard Time", "NPT": "Nepal Time", "NRT": "Nauru Time", "NST": "𐐤𐐭𐑁𐐲𐑌𐐼𐑊𐐲𐑌𐐼 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "NUT": "Niue Time", "NZDT": "New Zealand Daylight Time", "NZST": "New Zealand Standard Time", "OESZ": "Eastern European Summer Time", "OEZ": "Eastern European Standard Time", "OMSST": "Omsk Summer Time", "OMST": "Omsk Standard Time", "PDT": "𐐑𐐲𐑅𐐮𐑁𐐮𐐿 𐐔𐐩𐑊𐐴𐐻 𐐓𐐴𐑋", "PDTM": "Mexican Pacific Daylight Time", "PETDT": "Kamchatka Summer Time", "PETST": "Kamchatka Standard Time", "PGT": "Papua New Guinea Time", "PHOT": "Phoenix Islands Time", "PKT": "Pakistan Standard Time", "PKT DST": "Pakistan Summer Time", "PMDT": "St. Pierre & Miquelon Daylight Time", "PMST": "St. Pierre & Miquelon Standard Time", "PONT": "Pohnpei Time", "PST": "𐐑𐐲𐑅𐐮𐑁𐐮𐐿 𐐝𐐻𐐰𐑌𐐼𐐲𐑉𐐼 𐐓𐐴𐑋", "PST Philippine": "Philippine Standard Time", "PST Philippine DST": "Philippine Summer Time", "PST Pitcairn": "Pitcairn Time", "PSTM": "Mexican Pacific Standard Time", "PWT": "Palau Time", "PYST": "Paraguay Summer Time", "PYT": "Paraguay Standard Time", "PYT Korea": "North Korea Time", "RET": "Réunion Time", "ROTT": "Rothera Time", "SAKST": "Sakhalin Summer Time", "SAKT": "Sakhalin Standard Time", "SAMST": "Samara Summer Time", "SAMT": "Samara Standard Time", "SAST": "South Africa Standard Time", "SBT": "Solomon Islands Time", "SCT": "Seychelles Time", "SGT": "Singapore Standard Time", "SLST": "Lanka Time", "SRT": "Suriname Time", "SST Samoa": "American Samoa Standard Time", "SST Samoa Apia": "Samoa Standard Time", "SST Samoa Apia DST": "Samoa Daylight Time", "SST Samoa DST": "American Samoa Daylight Time", "SYOT": "Syowa Time", "TAAF": "French Southern & Antarctic Time", "TAHT": "Tahiti Time", "TJT": "Tajikistan Time", "TKT": "Tokelau Time", "TLT": "Timor-Leste Time", "TMST": "Turkmenistan Summer Time", "TMT": "Turkmenistan Standard Time", "TOST": "Tonga Summer Time", "TOT": "Tonga Standard Time", "TVT": "Tuvalu Time", "TWT": "Taiwan Standard Time", "TWT DST": "Taiwan Daylight Time", "ULAST": "Ulaanbaatar Summer Time", "ULAT": "Ulaanbaatar Standard Time", "UYST": "Uruguay Summer Time", "UYT": "Uruguay Standard Time", "UZT": "Uzbekistan Standard Time", "UZT DST": "Uzbekistan Summer Time", "VET": "Venezuela Time", "VLAST": "Vladivostok Summer Time", "VLAT": "Vladivostok Standard Time", "VOLST": "Volgograd Summer Time", "VOLT": "Volgograd Standard Time", "VOST": "Vostok Time", "VUT": "Vanuatu Standard Time", "VUT DST": "Vanuatu Summer Time", "WAKT": "Wake Island Time", "WARST": "Western Argentina Summer Time", "WART": "Western Argentina Standard Time", "WAST": "West Africa Time", "WAT": "West Africa Time", "WESZ": "Western European Summer Time", "WEZ": "Western European Standard Time", "WFT": "Wallis & Futuna Time", "WGST": "West Greenland Summer Time", "WGT": "West Greenland Standard Time", "WIB": "Western Indonesia Time", "WIT": "Eastern Indonesia Time", "WITA": "Central Indonesia Time", "YAKST": "Yakutsk Summer Time", "YAKT": "Yakutsk Standard Time", "YEKST": "Yekaterinburg Summer Time", "YEKT": "Yekaterinburg Standard Time", "YST": "Yukon Time", "МСК": "Moscow Standard Time", "اقتاۋ": "Aqtau Standard Time", "اقتاۋ قالاسى": "Aqtau Summer Time", "اقتوبە": "Aqtobe Standard Time", "اقتوبە قالاسى": "Aqtobe Summer Time", "الماتى": "Almaty Standard Time", "الماتى قالاسى": "Almaty Summer Time", "باتىس قازاق ەلى": "West Kazakhstan Time", "شىعىش قازاق ەلى": "East Kazakhstan Time", "قازاق ەلى": "Kazakhstan Time", "قىرعىزستان": "Kyrgyzstan Time", "قىزىلوردا": "Kyzylorda Standard Time", "قىزىلوردا قالاسى": "Kyzylorda Summer Time", "∅∅∅": "Azores Summer Time"},
	}
}

// Locale returns the current translators string locale
func (en *en_Dsrt) Locale() string {
	return en.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'en_Dsrt'
func (en *en_Dsrt) PluralsCardinal() []locales.PluralRule {
	return en.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'en_Dsrt'
func (en *en_Dsrt) PluralsOrdinal() []locales.PluralRule {
	return en.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'en_Dsrt'
func (en *en_Dsrt) PluralsRange() []locales.PluralRule {
	return en.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'en_Dsrt'
func (en *en_Dsrt) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'en_Dsrt'
func (en *en_Dsrt) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 1 && nMod100 != 11 {
		return locales.PluralRuleOne
	} else if nMod10 == 2 && nMod100 != 12 {
		return locales.PluralRuleTwo
	} else if nMod10 == 3 && nMod100 != 13 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'en_Dsrt'
func (en *en_Dsrt) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (en *en_Dsrt) MonthAbbreviated(month time.Month) string {
	if len(en.monthsAbbreviated) == 0 {
		return ""
	}
	return en.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (en *en_Dsrt) MonthsAbbreviated() []string {
	return en.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (en *en_Dsrt) MonthNarrow(month time.Month) string {
	if len(en.monthsNarrow) == 0 {
		return ""
	}
	return en.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (en *en_Dsrt) MonthsNarrow() []string {
	return en.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (en *en_Dsrt) MonthWide(month time.Month) string {
	if len(en.monthsWide) == 0 {
		return ""
	}
	return en.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (en *en_Dsrt) MonthsWide() []string {
	return en.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (en *en_Dsrt) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(en.daysAbbreviated) == 0 {
		return ""
	}
	return en.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (en *en_Dsrt) WeekdaysAbbreviated() []string {
	return en.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (en *en_Dsrt) WeekdayNarrow(weekday time.Weekday) string {
	if len(en.daysNarrow) == 0 {
		return ""
	}
	return en.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (en *en_Dsrt) WeekdaysNarrow() []string {
	return en.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (en *en_Dsrt) WeekdayShort(weekday time.Weekday) string {
	if len(en.daysShort) == 0 {
		return ""
	}
	return en.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (en *en_Dsrt) WeekdaysShort() []string {
	return en.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (en *en_Dsrt) WeekdayWide(weekday time.Weekday) string {
	if len(en.daysWide) == 0 {
		return ""
	}
	return en.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (en *en_Dsrt) WeekdaysWide() []string {
	return en.daysWide
}

// Decimal returns the decimal point of number
func (en *en_Dsrt) Decimal() string {
	return en.decimal
}

// Group returns the group of number
func (en *en_Dsrt) Group() string {
	return en.group
}

// Group returns the minus sign of number
func (en *en_Dsrt) Minus() string {
	return en.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'en_Dsrt' and handles both Whole and Real numbers based on 'v'
func (en *en_Dsrt) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, en.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, en.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'en_Dsrt' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (en *en_Dsrt) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, en.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, en.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'en_Dsrt'
func (en *en_Dsrt) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := en.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, en.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(en.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, en.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, en.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, en.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'en_Dsrt'
// in accounting notation.
func (en *en_Dsrt) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := en.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, en.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, en.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(en.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, en.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, en.minus[0])

	} else {

		for j := len(en.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, en.currencyPositivePrefix[j])
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
			b = append(b, en.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, en.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, en.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, en.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, en.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, en.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, en.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'en_Dsrt'
func (en *en_Dsrt) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, en.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, en.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := en.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
