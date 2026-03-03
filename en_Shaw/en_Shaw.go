package en_Shaw

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type en_Shaw struct {
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

// New returns a new instance of translator for the 'en_Shaw' locale
func New() locales.Translator {
	return &en_Shaw{
		locale:                 "en_Shaw",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "֏", "ANG", "Kz", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "KM", "BAN", "$", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "$", "Bs", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$", "BTN", "BUK", "P", "BYB", "BYN", "BYR", "$", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$", "CNH", "CNX", "CN¥", "$", "COU", "₡", "CSD", "CSK", "$", "$", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "kr", "$", "DZD", "ECS", "ECV", "EEK", "E£", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "$", "£", "FRF", "£", "GEK", "₾", "GHC", "GH₵", "£", "GMD", "FG", "GNS", "GQE", "GRD", "Q", "GWE", "GWP", "$", "HK$", "L", "HRD", "kn", "HTG", "Ft", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "kr", "ITL", "$", "JOD", "JP¥", "KES", "⃀", "៛", "CF", "₩", "KRH", "KRO", "₩", "KWD", "$", "₸", "₭", "L£", "Rs", "$", "LSL", "Lt", "LTT", "LUC", "LUF", "LUL", "Ls", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "Ar", "MGF", "MKD", "MKN", "MLF", "K", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "Rs", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "$", "₦", "NIC", "C$", "NLG", "kr", "Rs", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "Rs", "zł", "PLZ", "PTE", "₲", "QAR", "RHD", "ROL", "lei", "RSD", "₽", "RUR", "RF", "\\u20c1", "$", "SCR", "SDD", "SDG", "SDP", "kr", "$", "£", "SIT", "SKK", "SLE", "SLL", "SOS", "$", "SRG", "£", "STD", "Db", "SUR", "SVC", "£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "₺", "$", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "$", "UYW", "UZS", "VEB", "VED", "Bs", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "Cg.", "XDR", "XEU", "XFO", "XFU", "F\\u202fCFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "¤", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "·𐑡𐑨", "·𐑓𐑧", "·𐑥𐑸", "·𐑱𐑐", "·𐑥𐑱", "·𐑡𐑵", "·𐑡𐑫", "·𐑪𐑜", "·𐑕𐑧", "·𐑷𐑒", "·𐑯𐑴", "·𐑛𐑭"},
		monthsNarrow:           []string{"", "𐑡", "𐑓", "𐑥", "𐑱", "𐑥", "𐑡", "𐑡", "𐑷", "𐑕", "𐑪", "𐑯", "𐑛"},
		monthsWide:             []string{"", "·𐑡𐑨𐑙𐑘𐑭𐑢𐑺𐑰", "·𐑓𐑧𐑚𐑘𐑵𐑢𐑺𐑰", "·𐑥𐑸𐑗", "·𐑱𐑐𐑮𐑭𐑤", "·𐑥𐑱", "·𐑡𐑵𐑯", "·𐑡𐑫𐑤𐑲", "·𐑪𐑜𐑭𐑕𐑑", "·𐑕𐑧𐑐𐑑𐑧𐑥𐑚𐑸", "·𐑷𐑒𐑑𐑴𐑚𐑸", "·𐑯𐑴𐑝𐑧𐑥𐑚𐑸", "·𐑛𐑭𐑕𐑧𐑥𐑚𐑸"},
		daysAbbreviated:        []string{"·𐑕𐑭", "·𐑥𐑭", "·𐑑𐑵", "·𐑢𐑧", "·𐑔𐑻", "·𐑓𐑮", "·𐑕𐑨"},
		daysNarrow:             []string{"𐑕", "𐑥", "𐑑", "𐑢", "𐑔", "𐑓", "𐑕"},
		daysShort:              []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
		daysWide:               []string{"·𐑕𐑭𐑙𐑛𐑱", "·𐑥𐑭𐑙𐑛𐑱", "·𐑑𐑵𐑟𐑛𐑱", "·𐑢𐑧𐑙𐑟𐑛𐑱", "·𐑔𐑻𐑟𐑛𐑱", "·𐑓𐑮𐑲𐑛𐑱", "·𐑕𐑨𐑛𐑻𐑛𐑱"},
		timezones:              map[string]string{"ACDT": "Australian Central Daylight Time", "ACST": "Acre Summer Time", "ACT": "Acre Standard Time", "ACWDT": "Australian Central Western Daylight Time", "ACWST": "Australian Central Western Standard Time", "ADT": "Atlantic Daylight Time", "ADT Arabia": "Arabian Daylight Time", "AEDT": "Australian Eastern Daylight Time", "AEST": "Australian Eastern Standard Time", "AFT": "Afghanistan Time", "AKDT": "Alaska Daylight Time", "AKST": "Alaska Standard Time", "AMST": "Amazon Summer Time", "AMST Armenia": "Armenia Summer Time", "AMT": "Amazon Standard Time", "AMT Armenia": "Armenia Standard Time", "ANAST": "Anadyr Summer Time", "ANAT": "Anadyr Standard Time", "ARST": "Argentina Summer Time", "ART": "Argentina Standard Time", "AST": "Atlantic Standard Time", "AST Arabia": "Arabian Standard Time", "AWDT": "Australian Western Daylight Time", "AWST": "Australian Western Standard Time", "AZST": "Azerbaijan Summer Time", "AZT": "Azerbaijan Standard Time", "BDT Bangladesh": "Bangladesh Summer Time", "BNT": "Brunei Time", "BOT": "Bolivia Time", "BRST": "Brasilia Summer Time", "BRT": "Brasilia Standard Time", "BST Bangladesh": "Bangladesh Standard Time", "BT": "Bhutan Time", "CAST": "Casey Time", "CAT": "Central Africa Time", "CCT": "Cocos Islands Time", "CDT": "Central Daylight Time", "CHADT": "Chatham Daylight Time", "CHAST": "Chatham Standard Time", "CHUT": "Chuuk Time", "CKT": "Cook Islands Standard Time", "CKT DST": "Cook Islands Summer Time", "CLST": "Chile Summer Time", "CLT": "Chile Standard Time", "COST": "Colombia Summer Time", "COT": "Colombia Standard Time", "CST": "Central Standard Time", "CST China": "China Standard Time", "CST China DST": "China Daylight Time", "CVST": "Cape Verde Summer Time", "CVT": "Cape Verde Standard Time", "CXT": "Christmas Island Time", "ChST": "Chamorro Standard Time", "ChST NMI": "Northern Mariana Islands Time", "CuDT": "Cuba Daylight Time", "CuST": "Cuba Standard Time", "DAVT": "Davis Time", "DDUT": "Dumont d’Urville Time", "EASST": "Easter Island Summer Time", "EAST": "Easter Island Standard Time", "EAT": "East Africa Time", "ECT": "Ecuador Time", "EDT": "Eastern Daylight Time", "EGDT": "East Greenland Summer Time", "EGST": "East Greenland Standard Time", "EST": "Eastern Standard Time", "FEET": "𐑓𐑻𐑞𐑼-𐑰𐑕𐑑𐑼𐑯 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑑𐑲𐑥", "FJT": "Fiji Standard Time", "FJT Summer": "Fiji Summer Time", "FKST": "Falkland Islands Summer Time", "FKT": "Falkland Islands Standard Time", "FNST": "Fernando de Noronha Summer Time", "FNT": "Fernando de Noronha Standard Time", "GALT": "Galapagos Time", "GAMT": "Gambier Time", "GEST": "Georgia Summer Time", "GET": "Georgia Standard Time", "GFT": "French Guiana Time", "GIT": "Gilbert Islands Time", "GMT": "𐑜𐑮𐑧𐑯𐑦𐑗 𐑥𐑰𐑯 𐑑𐑲𐑥", "GNSST": "Greenland Summer Time", "GNST": "Greenland Standard Time", "GST": "South Georgia Time", "GST Guam": "Guam Standard Time", "GYT": "Guyana Time", "HADT": "Hawaii-Aleutian Daylight Time", "HAST": "Hawaii-Aleutian Standard Time", "HKST": "Hong Kong Summer Time", "HKT": "Hong Kong Standard Time", "HOVST": "Khovd Summer Time", "HOVT": "Khovd Standard Time", "ICT": "Indochina Time", "IDT": "Israel Daylight Time", "IOT": "Indian Ocean Time", "IRKST": "Irkutsk Summer Time", "IRKT": "Irkutsk Standard Time", "IRST": "Iran Standard Time", "IRST DST": "Iran Daylight Time", "IST": "India Standard Time", "IST Israel": "Israel Standard Time", "JDT": "Japan Daylight Time", "JST": "Japan Standard Time", "KOST": "Kosrae Time", "KRAST": "Krasnoyarsk Summer Time", "KRAT": "Krasnoyarsk Standard Time", "KST": "Korean Standard Time", "KST DST": "Korean Daylight Time", "LHDT": "Lord Howe Daylight Time", "LHST": "Lord Howe Standard Time", "LINT": "Line Islands Time", "MAGST": "Magadan Summer Time", "MAGT": "Magadan Standard Time", "MART": "Marquesas Time", "MAWT": "Mawson Time", "MDT": "Mountain Daylight Time", "MESZ": "𐑕𐑧𐑯𐑑𐑮𐑩𐑤 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑕𐑳𐑥𐑼 𐑑𐑲𐑥", "MEZ": "𐑕𐑧𐑯𐑑𐑮𐑩𐑤 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑕𐑑𐑨𐑯𐑛𐑼𐑛 𐑑𐑲𐑥", "MHT": "Marshall Islands Time", "MMT": "Myanmar Time", "MSD": "Moscow Summer Time", "MST": "Mountain Standard Time", "MUST": "Mauritius Summer Time", "MUT": "Mauritius Standard Time", "MVT": "Maldives Time", "MYT": "Malaysia Time", "NCT": "New Caledonia Standard Time", "NDT": "Newfoundland Daylight Time", "NDT New Caledonia": "New Caledonia Summer Time", "NFDT": "Norfolk Island Daylight Time", "NFT": "Norfolk Island Standard Time", "NOVST": "Novosibirsk Summer Time", "NOVT": "Novosibirsk Standard Time", "NPT": "Nepal Time", "NRT": "Nauru Time", "NST": "Newfoundland Standard Time", "NUT": "Niue Time", "NZDT": "New Zealand Daylight Time", "NZST": "New Zealand Standard Time", "OESZ": "𐑰𐑕𐑑𐑼𐑯 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑕𐑳𐑥𐑼 𐑑𐑲𐑥", "OEZ": "𐑰𐑕𐑑𐑼𐑯 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑕𐑑𐑨𐑯𐑛𐑼𐑛 𐑑𐑲𐑥", "OMSST": "Omsk Summer Time", "OMST": "Omsk Standard Time", "PDT": "Pacific Daylight Time", "PDTM": "Mexican Pacific Daylight Time", "PETDT": "Kamchatka Summer Time", "PETST": "Kamchatka Standard Time", "PGT": "Papua New Guinea Time", "PHOT": "Phoenix Islands Time", "PKT": "Pakistan Standard Time", "PKT DST": "Pakistan Summer Time", "PMDT": "St. Pierre & Miquelon Daylight Time", "PMST": "St. Pierre & Miquelon Standard Time", "PONT": "Pohnpei Time", "PST": "Pacific Standard Time", "PST Philippine": "Philippine Standard Time", "PST Philippine DST": "Philippine Summer Time", "PST Pitcairn": "Pitcairn Time", "PSTM": "Mexican Pacific Standard Time", "PWT": "Palau Time", "PYST": "Paraguay Summer Time", "PYT": "Paraguay Standard Time", "PYT Korea": "North Korea Time", "RET": "Réunion Time", "ROTT": "Rothera Time", "SAKST": "Sakhalin Summer Time", "SAKT": "Sakhalin Standard Time", "SAMST": "Samara Summer Time", "SAMT": "Samara Standard Time", "SAST": "South Africa Standard Time", "SBT": "Solomon Islands Time", "SCT": "Seychelles Time", "SGT": "Singapore Standard Time", "SLST": "Lanka Time", "SRT": "Suriname Time", "SST Samoa": "American Samoa Standard Time", "SST Samoa Apia": "Samoa Standard Time", "SST Samoa Apia DST": "Samoa Daylight Time", "SST Samoa DST": "American Samoa Daylight Time", "SYOT": "Syowa Time", "TAAF": "French Southern & Antarctic Time", "TAHT": "Tahiti Time", "TJT": "Tajikistan Time", "TKT": "Tokelau Time", "TLT": "Timor-Leste Time", "TMST": "Turkmenistan Summer Time", "TMT": "Turkmenistan Standard Time", "TOST": "Tonga Summer Time", "TOT": "Tonga Standard Time", "TVT": "Tuvalu Time", "TWT": "Taiwan Standard Time", "TWT DST": "Taiwan Daylight Time", "ULAST": "Ulaanbaatar Summer Time", "ULAT": "Ulaanbaatar Standard Time", "UYST": "Uruguay Summer Time", "UYT": "Uruguay Standard Time", "UZT": "Uzbekistan Standard Time", "UZT DST": "Uzbekistan Summer Time", "VET": "Venezuela Time", "VLAST": "Vladivostok Summer Time", "VLAT": "Vladivostok Standard Time", "VOLST": "Volgograd Summer Time", "VOLT": "Volgograd Standard Time", "VOST": "Vostok Time", "VUT": "Vanuatu Standard Time", "VUT DST": "Vanuatu Summer Time", "WAKT": "Wake Island Time", "WARST": "Western Argentina Summer Time", "WART": "Western Argentina Standard Time", "WAST": "West Africa Time", "WAT": "West Africa Time", "WESZ": "𐑢𐑧𐑕𐑑𐑼𐑯 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑕𐑳𐑥𐑼 𐑑𐑲𐑥", "WEZ": "𐑢𐑧𐑕𐑑𐑼𐑯 𐑘𐑫𐑼𐑩𐑐𐑾𐑯 𐑕𐑑𐑨𐑯𐑛𐑼𐑛 𐑑𐑲𐑥", "WFT": "Wallis & Futuna Time", "WGST": "West Greenland Summer Time", "WGT": "West Greenland Standard Time", "WIB": "Western Indonesia Time", "WIT": "Eastern Indonesia Time", "WITA": "Central Indonesia Time", "YAKST": "Yakutsk Summer Time", "YAKT": "Yakutsk Standard Time", "YEKST": "Yekaterinburg Summer Time", "YEKT": "Yekaterinburg Standard Time", "YST": "Yukon Time", "МСК": "Moscow Standard Time", "اقتاۋ": "Aqtau Standard Time", "اقتاۋ قالاسى": "Aqtau Summer Time", "اقتوبە": "Aqtobe Standard Time", "اقتوبە قالاسى": "Aqtobe Summer Time", "الماتى": "Almaty Standard Time", "الماتى قالاسى": "Almaty Summer Time", "باتىس قازاق ەلى": "West Kazakhstan Time", "شىعىش قازاق ەلى": "East Kazakhstan Time", "قازاق ەلى": "Kazakhstan Time", "قىرعىزستان": "Kyrgyzstan Time", "قىزىلوردا": "Kyzylorda Standard Time", "قىزىلوردا قالاسى": "Kyzylorda Summer Time", "∅∅∅": "Peru Summer Time"},
	}
}

// Locale returns the current translators string locale
func (en *en_Shaw) Locale() string {
	return en.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'en_Shaw'
func (en *en_Shaw) PluralsCardinal() []locales.PluralRule {
	return en.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'en_Shaw'
func (en *en_Shaw) PluralsOrdinal() []locales.PluralRule {
	return en.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'en_Shaw'
func (en *en_Shaw) PluralsRange() []locales.PluralRule {
	return en.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'en_Shaw'
func (en *en_Shaw) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'en_Shaw'
func (en *en_Shaw) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

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

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'en_Shaw'
func (en *en_Shaw) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (en *en_Shaw) MonthAbbreviated(month time.Month) string {
	if len(en.monthsAbbreviated) == 0 {
		return ""
	}
	return en.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (en *en_Shaw) MonthsAbbreviated() []string {
	return en.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (en *en_Shaw) MonthNarrow(month time.Month) string {
	if len(en.monthsNarrow) == 0 {
		return ""
	}
	return en.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (en *en_Shaw) MonthsNarrow() []string {
	return en.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (en *en_Shaw) MonthWide(month time.Month) string {
	if len(en.monthsWide) == 0 {
		return ""
	}
	return en.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (en *en_Shaw) MonthsWide() []string {
	return en.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (en *en_Shaw) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(en.daysAbbreviated) == 0 {
		return ""
	}
	return en.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (en *en_Shaw) WeekdaysAbbreviated() []string {
	return en.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (en *en_Shaw) WeekdayNarrow(weekday time.Weekday) string {
	if len(en.daysNarrow) == 0 {
		return ""
	}
	return en.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (en *en_Shaw) WeekdaysNarrow() []string {
	return en.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (en *en_Shaw) WeekdayShort(weekday time.Weekday) string {
	if len(en.daysShort) == 0 {
		return ""
	}
	return en.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (en *en_Shaw) WeekdaysShort() []string {
	return en.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (en *en_Shaw) WeekdayWide(weekday time.Weekday) string {
	if len(en.daysWide) == 0 {
		return ""
	}
	return en.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (en *en_Shaw) WeekdaysWide() []string {
	return en.daysWide
}

// Decimal returns the decimal point of number
func (en *en_Shaw) Decimal() string {
	return en.decimal
}

// Group returns the group of number
func (en *en_Shaw) Group() string {
	return en.group
}

// Group returns the minus sign of number
func (en *en_Shaw) Minus() string {
	return en.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'en_Shaw' and handles both Whole and Real numbers based on 'v'
func (en *en_Shaw) FmtNumber(num float64, v uint64) string {

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

// FmtPercent returns 'num' with digits/precision of 'v' for 'en_Shaw' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (en *en_Shaw) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'en_Shaw'
func (en *en_Shaw) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := en.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'en_Shaw'
// in accounting notation.
func (en *en_Shaw) FmtAccounting(num float64, v uint64, currency currency.Type) string {

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

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, en.currencyNegativePrefix[0])

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
			b = append(b, en.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, en.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, en.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, en.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, en.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, en.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtTimeShort(t time.Time) string {

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

// FmtTimeMedium returns the medium time representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtTimeMedium(t time.Time) string {

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

// FmtTimeLong returns the long time representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtTimeLong(t time.Time) string {

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

// FmtTimeFull returns the full time representation of 't' for 'en_Shaw'
func (en *en_Shaw) FmtTimeFull(t time.Time) string {

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
