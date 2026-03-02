package zu_ZA

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type zu_ZA struct {
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

// New returns a new instance of translator for the 'zu_ZA' locale
func New() locales.Translator {
	return &zu_ZA{
		locale:                 "zu_ZA",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mas", "Eph", "Mey", "Jun", "Jul", "Aga", "Sep", "Okt", "Nov", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "E", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januwari", "Februwari", "Mashi", "Ephreli", "Meyi", "Juni", "Julayi", "Agasti", "Septhemba", "Okthoba", "Novemba", "Disemba"},
		daysAbbreviated:        []string{"Son", "Mso", "Bil", "Tha", "Sin", "Hla", "Mgq"},
		daysNarrow:             []string{"S", "M", "B", "T", "S", "H", "M"},
		daysWide:               []string{"ISonto", "UMsombuluko", "ULwesibili", "ULwesithathu", "ULwesine", "ULwesihlanu", "UMgqibelo"},
		timezones:              map[string]string{"ACDT": "Isikhathi sase-Australian Central sasemini", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Isikhathi sasemini sase-Australian Central West", "ACWST": "Isikhathi sase-Australian Central West esivamile", "ADT": "Isikhathi sase-Atlantic sasemini", "ADT Arabia": "Isikhathi semini sase-Arabian", "AEDT": "Isikhathi sasemini sase-Australian East", "AEST": "Isikhathi esivamile sase-Australian East", "AFT": "Isikhathi sase-Afghanistan", "AKDT": "Isikhathi sase-Alaska sasemini", "AKST": "Isikhathi sase-Alaska esijwayelekile", "AMST": "Isikhathi sase-Amazon sasehlobo", "AMST Armenia": "Isikhathi sehlobo sase-Armenia", "AMT": "Isikhathi sase-Amazon esijwayelekile", "AMT Armenia": "Isikhathi esivamile sase-Armenia", "ANAST": "esase-Anadyr Summer Time", "ANAT": "esase-Anadyr Standard Time", "ARST": "Isikhathi sase-Argentina sasehlobo", "ART": "Isikhathi sase-Argentina esijwayelekile", "AST": "Isikhathi sase-Atlantic esijwayelekile", "AST Arabia": "Isikhathi esivamile sase-Arabian", "AWDT": "Isikhathi sase-Australian Western sasemini", "AWST": "Isikhathi sase-Australian Western esivamile", "AZST": "Isikhathi sehlobo sase-Azerbaijan", "AZT": "Isikhathi esivamile sase-Azerbaijan", "BDT Bangladesh": "Isikhathi sase-Bangladesh sasehlobo", "BNT": "Isikhathi sase-Brunei Darussalam", "BOT": "Isikhathi sase-Bolivia", "BRST": "Isikhathi sase-Brasilia sasehlobo", "BRT": "Isikhathi sase-Brasilia esijwayelekile", "BST Bangladesh": "Isikhathi sase-Bangladesh esivamile", "BT": "Isikhathi sase-Bhutan", "CAST": "CAST", "CAT": "Isikhathi sase-Central Africa", "CCT": "Isikhathi sase-Cocos Islands", "CDT": "Isikhathi sase-North American Central sasemini", "CHADT": "Isikhathi sasemini sase-Chatham", "CHAST": "Isikhathi esivamile sase-Chatham", "CHUT": "Isikhathi sase-Chuuk", "CKT": "Isikhathi esivamile sase-Cook Islands", "CKT DST": "Isikhathi esiyingxenye yasehlobo sase-Cook Islands", "CLST": "Isikhathi sase-Chile sasehlobo", "CLT": "Isikhathi sase-Chile esijwayelekile", "COST": "Isikhathi sase-Colombia sasehlobo", "COT": "Isikhathi sase-Colombia esijwayelekile", "CST": "Isikhathi sase-North American Central esijwayelekile", "CST China": "Isikhathi esivamile sase-China", "CST China DST": "Isikhathi semini sase-China", "CVST": "Isikhathi sehlobo sase-Cape Verde", "CVT": "Isikhathi esezingeni sase-Cape Verde", "CXT": "Isikhathi sase-Christmas Island", "ChST": "Isikhathi esivamile sase-Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Isikhathi sase-Cuba sasemini", "CuST": "Isikhathi sase-Cuba esijwayelekile", "DAVT": "Isikhathi sase-Davis", "DDUT": "Isikhathi sase-Dumont-d’Urville", "EASST": "Isikhathi sase-Easter Island sasehlobo", "EAST": "Isikhathi sase-Easter Island esijwayelekile", "EAT": "Isikhathi saseMpumalanga Afrika", "ECT": "Isikhathi sase-Ecuador", "EDT": "Isikhathi sase-North American East sasemini", "EGDT": "Isikhathi sase-East Greenland sasemini", "EGST": "Isikhathi sase-East Greenland esijwayelekile", "EST": "Isikhathi sase-North American East esijwayelekile", "FEET": "Isikhathi sase-Further-eastern Europe", "FJT": "Isikhathi esivamile sase-Fiji", "FJT Summer": "Isikhathi sehlobo sase-Fiji", "FKST": "Isikhathi sase-Falkland Islands sasehlobo", "FKT": "Isikhathi sase-Falkland Islands esijwayelekile", "FNST": "Isikhathi sase-Fernando de Noronha sasehlobo", "FNT": "Isikhathi sase-Fernando de Noronha esivamile", "GALT": "Isikhathi sase-Galapagos", "GAMT": "Isikhathi sase-Gambier", "GEST": "Isikhathi sehlobo sase-Georgia", "GET": "Isikhathi esivamile sase-Georgia", "GFT": "Isikhathi sase-French Guiana", "GIT": "Isikhathi sase-Gilbert Islands", "GMT": "Isikhathi sase-Greenwich Mean", "GNSST": "GNSST", "GNST": "GNST", "GST": "Isikhathi esivamile sase-Gulf", "GST Guam": "GST Guam", "GYT": "Isikhathi sase-Guyana", "HADT": "Isikhathi sase-Hawaii-Aleutia esijwayelekile", "HAST": "Isikhathi sase-Hawaii-Aleutia esijwayelekile", "HKST": "Isikhathi sehlobo sase-Hong Kong", "HKT": "Isikhathi esivamile sase-Hong Kong", "HOVST": "Isikhathi sehlobo e-Hovd", "HOVT": "Isikhathi Esimisiwe sase-Hovd", "ICT": "Isikhathi sase-Indochina", "IDT": "Isikhathi sasemini sakwa-Israel", "IOT": "Isikhathi sase-Indian Ocean", "IRKST": "Isikhathi sasehlobo e-Irkutsk", "IRKT": "Isikhathi Esivamile sase-Irkutsk", "IRST": "Isikhathi sase-Iran esivamile", "IRST DST": "Isikhathi sase-Iran sasemini", "IST": "Isikhathi sase-India esivamile", "IST Israel": "Isikhathi esivamile sase-Israel", "JDT": "Isikhathi semini sase-Japan", "JST": "Isikhathi esivamile sase-Japan", "KOST": "Isikhathi sase-Kosrae", "KRAST": "Isikhathi sasehlobo e-Krasnoyarsk", "KRAT": "Isikhathi Esivamile sase-Krasnoyarsk", "KST": "Isikhathi Esivamile sase-Korea", "KST DST": "Isikhathi semini sase-Korea", "LHDT": "Isikhathi sase-Lord Howe sasemini", "LHST": "Isikhathi sase-Lord Howe esivamile", "LINT": "Isikhathi sase-Line Islands", "MAGST": "Isikhathi sasehlobo e-Magadan", "MAGT": "Isikhathi Esivamile sase-Magadan", "MART": "Isikhathi sase-Marquesas", "MAWT": "Isikhathi sase-Mawson", "MDT": "MDT", "MESZ": "Isikhathi sasehlobo sase-Central Europe", "MEZ": "Isikhathi esijwayelekile sase-Central Europe", "MHT": "Isikhathi sase-Marshall Islands", "MMT": "Isikhathi sase-Myanmar", "MSD": "Isikhathi sasehlobo e-Moscow", "MST": "MST", "MUST": "Isikhathi sehlobo sase-Mauritius", "MUT": "Isikhathi esivamile sase-Mauritius", "MVT": "Isikhathi sase-Maldives", "MYT": "Isikhathi sase-Malaysia", "NCT": "Isikhathi sase-New Caledonia esivamile", "NDT": "Isikhathi sase-Newfoundland sasemini", "NDT New Caledonia": "Isikhathi sase-New Caledonia sasehlobo", "NFDT": "Isikhathi sase-Norfolk Islands sasehlobo", "NFT": "Isikhathi sase-Norfolk Islands esivamile", "NOVST": "Isikhathi sasehlobo sase-Novosibirsk", "NOVT": "Isikhathi Esivamile sase-Novosibirsk", "NPT": "Isikhathi sase-Nepal", "NRT": "Isikhathi sase-Nauru", "NST": "Isikhathi sase-Newfoundland esijwayelekile", "NUT": "Isikhathi sase-Niue", "NZDT": "Isikhathi sasemini sase-New Zealand", "NZST": "Isikhathi esivamile sase-New Zealand", "OESZ": "Isikhathi sasehlobo sase-Eastern Europe", "OEZ": "Isikhathi esijwayelekile sase-Eastern Europe", "OMSST": "Isikhathi sasehlobo sase-Omsk", "OMST": "Isikhathi Esivamile sase-Omsk", "PDT": "Isikhathi sase-North American Pacific sasemini", "PDTM": "Isikhathi sase-Mexican Pacific sasemini", "PETDT": "esase-Petropavlovsk-Kamchatski Summer Time", "PETST": "esase-Petropavlovsk-Kamchatski Standard Time", "PGT": "Isikhathi sase-Papua New Guinea", "PHOT": "Isikhathi sase-Phoenix Islands", "PKT": "Isikhathi sase-Pakistan esivamile", "PKT DST": "Isikhathi sase-Pakistan sasehlobo", "PMDT": "Isikhathi sase-Saint Pierre nase-Miquelon sasemini", "PMST": "Iikhathi sase-Saint Pierre nase-Miquelon esijwayelekile", "PONT": "Isikhathi sase-Ponape", "PST": "Isikhathi sase-North American Pacific esijwayelekile", "PST Philippine": "Isikhathi esivamile sase-Philippine", "PST Philippine DST": "Isikhathi sehlobo sase-Philippine", "PST Pitcairn": "Isikhathi sase-Pitcairn", "PSTM": "Isikhathi sase-Mexican Pacific esijwayelekile", "PWT": "Isikhathi sase-Palau", "PYST": "Isikhathi sase-Paraguay sasehlobo", "PYT": "Isikhathi sase-Paraguay esivamile", "PYT Korea": "Isikhathi sase-Pyongyang", "RET": "Isikhathi sase-Reunion", "ROTT": "Isikhathi sase-Rothera", "SAKST": "Isikhathi sasehlobo e-Sakhalin", "SAKT": "Isikhathi Esivamile sase-Sakhalin", "SAMST": "esase-Samara Summer Time", "SAMT": "esase-Samara Standard Time", "SAST": "Isikhathi esivamile saseNingizimu Afrika", "SBT": "Isikhathi sase-Solomon Islands", "SCT": "Isikhathi sase-Seychelles", "SGT": "Isikhathi esivamile sase-Singapore", "SLST": "SLST", "SRT": "Isikhathi sase-Suriname", "SST Samoa": "Isikhathi sase-Samoa esivamile", "SST Samoa Apia": "Isikhathi sase-Apia esivamile", "SST Samoa Apia DST": "Isikhathi sase-Apia sasemini", "SST Samoa DST": "Isikhathi sase-Samoa sasemini", "SYOT": "Isikhathi sase-Syowa", "TAAF": "Isikhathi sase-French Southern nase-Antarctic", "TAHT": "Isikhathi sase-Tahiti", "TJT": "Isikhathi sase-Tajikistan", "TKT": "Isikhathi sase-Tokelau", "TLT": "Isikhathi sase-East Timor", "TMST": "Isikhathi sehlobo sase-Turkmenistan", "TMT": "Isikhathi esivamile sase-Turkmenistan", "TOST": "Isikhathi sase-Tonga sasehlobo", "TOT": "Isikhathi sase-Tonga esivamile", "TVT": "Isikhathi sase-Tuvalu", "TWT": "Isikhathi esivamile sase-Taipei", "TWT DST": "Isikhathi semini sase-Taipei", "ULAST": "Isikhathi sehlobo e-Ulan Bator", "ULAT": "Isikhathi Esimisiwe sase-Ulan Bator", "UYST": "Isikhathi sase-Uruguay sasehlobo", "UYT": "Isikhathi sase-Uruguay esijwayelekile", "UZT": "Isikhathi esivamile sase-Uzbekistan", "UZT DST": "Isikhathi sehlobo sase-Uzbekistan", "VET": "Isikhathi sase-Venezuela", "VLAST": "Isikhathi sasehlobo e-Vladivostok", "VLAT": "Isikhathi Esivamile sase-Vladivostok", "VOLST": "Isikhathi sase-Volgograd sasehlobo", "VOLT": "Isikhathi Esivamile sase-Volgograd", "VOST": "Isikhathi sase-Vostok", "VUT": "Isikhathi sase-Vanuatu esijwayelekile", "VUT DST": "Isikhathi sase-Vanuatu sasehlobo", "WAKT": "Isikhathi sase-Wake Island", "WARST": "Isikhathi saseNyakatho ne-Argentina sasehlobo", "WART": "Isikhathi saseNyakatho ne-Argentina esijwayelekile", "WAST": "Isikhathi saseNtshonalanga Afrika", "WAT": "Isikhathi saseNtshonalanga Afrika", "WESZ": "Isikhathi sasehlobo sase-Western Europe", "WEZ": "Isikhathi esijwayelekile sase-Western Europe", "WFT": "Isikhathi sase-Wallis nase-Futuna", "WGST": "Isikhathi sase-West Greenland sasehlobo", "WGT": "Isikhathi sase-West Greenland esijwayelekile", "WIB": "Isikhathi sase-Western Indonesia", "WIT": "Isikhathi sase-Eastern Indonesia", "WITA": "Isikhathi sase-Central Indonesia", "YAKST": "Isikhathi sasehlobo e-Yakutsk", "YAKT": "Isikhathi Esivamile sase-Yakutsk", "YEKST": "Isikhathi sasehlobo e-Yekaterinburg", "YEKT": "Isikhathi Esivamile sase-Yekaterinburg", "YST": "Yukon Time", "МСК": "Isikhathi sase-Moscow esijwayelekile", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Isikhathi saseNtshonalanga ne-Kazakhstan", "شىعىش قازاق ەلى": "Isikhathi sase-Mpumalanga ne-Kazakhstan", "قازاق ەلى": "Isikhathi saseKazakhstan", "قىرعىزستان": "Isikhathi sase-Kyrgystan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Isikhathi sase-Peru sasehlobo"},
	}
}

// Locale returns the current translators string locale
func (zu *zu_ZA) Locale() string {
	return zu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'zu_ZA'
func (zu *zu_ZA) PluralsCardinal() []locales.PluralRule {
	return zu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'zu_ZA'
func (zu *zu_ZA) PluralsOrdinal() []locales.PluralRule {
	return zu.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'zu_ZA'
func (zu *zu_ZA) PluralsRange() []locales.PluralRule {
	return zu.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'zu_ZA'
func (zu *zu_ZA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'zu_ZA'
func (zu *zu_ZA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'zu_ZA'
func (zu *zu_ZA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := zu.CardinalPluralRule(num1, v1)
	end := zu.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (zu *zu_ZA) MonthAbbreviated(month time.Month) string {
	return zu.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (zu *zu_ZA) MonthsAbbreviated() []string {
	return zu.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (zu *zu_ZA) MonthNarrow(month time.Month) string {
	return zu.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (zu *zu_ZA) MonthsNarrow() []string {
	return zu.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (zu *zu_ZA) MonthWide(month time.Month) string {
	return zu.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (zu *zu_ZA) MonthsWide() []string {
	return zu.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (zu *zu_ZA) WeekdayAbbreviated(weekday time.Weekday) string {
	return zu.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (zu *zu_ZA) WeekdaysAbbreviated() []string {
	return zu.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (zu *zu_ZA) WeekdayNarrow(weekday time.Weekday) string {
	return zu.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (zu *zu_ZA) WeekdaysNarrow() []string {
	return zu.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (zu *zu_ZA) WeekdayShort(weekday time.Weekday) string {
	return zu.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (zu *zu_ZA) WeekdaysShort() []string {
	return zu.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (zu *zu_ZA) WeekdayWide(weekday time.Weekday) string {
	return zu.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (zu *zu_ZA) WeekdaysWide() []string {
	return zu.daysWide
}

// Decimal returns the decimal point of number
func (zu *zu_ZA) Decimal() string {
	return zu.decimal
}

// Group returns the group of number
func (zu *zu_ZA) Group() string {
	return zu.group
}

// Group returns the minus sign of number
func (zu *zu_ZA) Minus() string {
	return zu.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'zu_ZA' and handles both Whole and Real numbers based on 'v'
func (zu *zu_ZA) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, zu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'zu_ZA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (zu *zu_ZA) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, zu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, zu.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'zu_ZA'
func (zu *zu_ZA) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := zu.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zu.group[0])
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
		b = append(b, zu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, zu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'zu_ZA'
// in accounting notation.
func (zu *zu_ZA) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := zu.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zu.group[0])
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

		b = append(b, zu.currencyNegativePrefix[0])

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
			b = append(b, zu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, zu.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, zu.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, zu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, zu.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, zu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'zu_ZA'
func (zu *zu_ZA) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := zu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
