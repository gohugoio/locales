package ee

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ee struct {
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

// New returns a new instance of translator for the 'ee' locale
func New() locales.Translator {
	return &ee{
		locale:                 "ee",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GH₵", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "dzv", "dzd", "ted", "afɔ", "dam", "mas", "sia", "dea", "any", "kel", "ade", "dzm"},
		monthsNarrow:           []string{"", "d", "d", "t", "a", "d", "m", "s", "d", "a", "k", "a", "d"},
		monthsWide:             []string{"", "dzove", "dzodze", "tedoxe", "afɔfĩe", "dame", "masa", "siamlɔm", "deasiamime", "anyɔnyɔ", "kele", "adeɛmekpɔxe", "dzome"},
		daysAbbreviated:        []string{"kɔs", "dzo", "bla", "kuɖ", "yaw", "fiɖ", "mem"},
		daysNarrow:             []string{"k", "d", "b", "k", "y", "f", "m"},
		daysWide:               []string{"kɔsiɖa", "dzoɖa", "blaɖa", "kuɖa", "yawoɖa", "fiɖa", "memleɖa"},
		periodsAbbreviated:     []string{"ŋdi", "ɣetrɔ"},
		timezones:              map[string]string{"ACDT": "Australian Central dzomeli gaƒoƒo me", "ACST": "Eker dzomeŋɔli gaƒoƒome", "ACT": "Eker gaƒoƒoɖoanyime", "ACWDT": "Australian Central Western kele gaƒoƒo me", "ACWST": "Australian Central Western nutome gaƒoƒo me", "ADT": "Atlantic kele gaƒoƒome", "ADT Arabia": "Arabia dzomeŋɔli gaƒoƒo me", "AEDT": "Australian Eastern kele gaƒoƒo me", "AEST": "Australian Eastern nutome gaƒoƒo me", "AFT": "Afghanistan gaƒoƒo me", "AKDT": "Alaska kele gaƒoƒo me", "AKST": "Alaska nutome gaƒoƒo me", "AMST": "Amazon dzomeŋɔli gaƒoƒo me", "AMST Armenia": "Armenia dzomeŋɔli gaƒoƒo me", "AMT": "Amazon nutome gaƒoƒo me", "AMT Armenia": "Armenia nutome gaƒoƒo me", "ANAST": "Anadir ŋkekeme gaƒoƒome", "ANAT": "Anadir gaƒoƒoɖoanyime", "ARST": "Argentina dzomeŋɔli gaƒoƒo me", "ART": "Argentina nutome gaƒoƒo me", "AST": "Atlantic nutome gaƒoƒome", "AST Arabia": "Arabia nutome gaƒoƒo me", "AWDT": "Australian Western kele gaƒoƒo me", "AWST": "Australian Western nutome gaƒoƒo me", "AZST": "Azerbaijan dzomeŋɔli gaƒoƒo me", "AZT": "Azerbaijan nutome gaƒoƒo me", "BDT Bangladesh": "Bangladesh dzomeŋɔli gaƒoƒo me", "BNT": "Brunei Darussalam gaƒoƒo me", "BOT": "Bolivia gaƒoƒo me", "BRST": "Brasilia dzomeŋɔli gaƒoƒo me", "BRT": "Brasilia nutome gaƒoƒo me", "BST Bangladesh": "Bangladesh nutome gaƒoƒo me", "BT": "Bhutan gaƒoƒo me", "CAST": "CAST", "CAT": "Central Africa gaƒoƒo me", "CCT": "Cocos Islands gaƒoƒo me", "CDT": "Titina America kele gaƒoƒo me", "CHADT": "Chatham kele gaƒoƒo me", "CHAST": "Chatham nutome gaƒoƒo me", "CHUT": "Chuuk gaƒoƒo me", "CKT": "Cook Islands nutome gaƒoƒo me", "CKT DST": "Cook Islands dzomeŋɔli gaƒoƒo me", "CLST": "Chile dzomeŋɔli gaƒoƒo me", "CLT": "Chile nutome gaƒoƒo me", "COST": "Colombia dzomeŋɔli gaƒoƒo me", "COT": "Colombia nutome gaƒoƒo me", "CST": "Titina America nutome gaƒoƒo me", "CST China": "China nutome gaƒoƒo me", "CST China DST": "China kele gaƒoƒo me", "CVST": "Cape Verde dzomeŋɔli gaƒoƒo me", "CVT": "Cape Verde nutome gaƒoƒo me", "CXT": "Christmas Island gaƒoƒo me", "ChST": "Chamorro gaƒoƒo me", "ChST NMI": "ChST NMI", "CuDT": "Cuba kele gaƒoƒome", "CuST": "Cuba nutome gaƒoƒome", "DAVT": "Davis gaƒoƒo me", "DDUT": "Dumont-d’Urville gaƒoƒo me", "EASST": "Easter Island dzomeŋɔli gaƒoƒo me", "EAST": "Easter Island nutome gaƒoƒo me", "EAT": "East Africa gaƒoƒo me", "ECT": "Ecuador gaƒoƒo me", "EDT": "Eastern America kele gaƒoƒo me", "EGDT": "East Greenland dzomeŋɔli gaƒoƒo me", "EGST": "East Greenland nutome gaƒoƒo me", "EST": "Eastern America nutome gaƒoƒo me", "FEET": "FEET", "FJT": "Fiji nutome gaƒoƒo me", "FJT Summer": "Fiji dzomeŋɔli gaƒoƒo me", "FKST": "Falkland Islands dzomeŋɔli gaƒoƒo me", "FKT": "Falkland Islands nutome gaƒoƒo me", "FNST": "Fernando de Noronha dzomeŋɔli gaƒoƒo me", "FNT": "Fernando de Noronha nutome gaƒoƒo me", "GALT": "Galapagos gaƒoƒo me", "GAMT": "Gambier gaƒoƒo me", "GEST": "Georgia dzomeŋɔli gaƒoƒo me", "GET": "Georgia nutome gaƒoƒo me", "GFT": "French Guiana gaƒoƒo me", "GIT": "Gilbert Islands gaƒoƒo me", "GMT": "Greenwich gaƒoƒo", "GNSST": "GNSST", "GNST": "GNST", "GST": "Gulf nutome gaƒoƒo me", "GST Guam": "GST Guam", "GYT": "Guyana gaƒoƒo me", "HADT": "Hawaii-Aleutia nutome gaƒoƒo me", "HAST": "Hawaii-Aleutia nutome gaƒoƒo me", "HKST": "Hong Kong dzomeŋɔli gaƒoƒo me", "HKT": "Hong Kong nutome gaƒoƒo me", "HOVST": "Hovd dzomeŋɔli gaƒoƒo me", "HOVT": "Hovd nutome gaƒoƒo me", "ICT": "Indonesia gaƒoƒo me", "IDT": "Israel kele gaƒoƒo me", "IOT": "Indian Ocean gaƒoƒo me", "IRKST": "Irkutsk dzomeŋɔli gaƒoƒo me", "IRKT": "Irkutsk nutome gaƒoƒo me", "IRST": "Iran nutome gaƒoƒo me", "IRST DST": "Iran kele gaƒoƒo me", "IST": "India gaƒoƒo me", "IST Israel": "Israel nutome gaƒoƒo me", "JDT": "Japan dzomeŋɔli gaƒoƒo me", "JST": "Japan nutome gaƒoƒo me", "KOST": "Kosrae gaƒoƒo me", "KRAST": "Krasnoyarsk dzomeŋɔli gaƒoƒo me", "KRAT": "Krasnoyarsk nutome gaƒoƒo me", "KST": "Korea nutome gaƒoƒo me", "KST DST": "Korea dzomeŋɔli gaƒoƒo me", "LHDT": "Lord Howe kele gaƒoƒo me", "LHST": "Lord Howe nutome gaƒoƒo me", "LINT": "Line Islands gaƒoƒo me", "MAGST": "Magadan dzomeŋɔli gaƒoƒo me", "MAGT": "Magadan nutome gaƒoƒo me", "MART": "Marquesas gaƒoƒo me", "MAWT": "Mawson gaƒoƒo me", "MDT": "Makau ŋkekeme gaƒoƒome", "MESZ": "Central Europe dzomeŋɔli gaƒoƒo me", "MEZ": "Central Europe nutome gaƒoƒo me", "MHT": "Marshall Islands gaƒoƒo me", "MMT": "Myanmar gaƒoƒo me", "MSD": "Moscow dzomeŋɔli gaƒoƒo me", "MST": "Makau gaƒoƒoɖoanyime", "MUST": "Mauritius dzomeŋɔli gaƒoƒo me", "MUT": "Mauritius nutome gaƒoƒo me", "MVT": "Maldives gaƒoƒo me", "MYT": "Malaysia gaƒoƒo me", "NCT": "New Caledonia nutome gaƒoƒo me", "NDT": "Newfoundland kele gaƒoƒome", "NDT New Caledonia": "New Caledonia dzomeŋɔli gaƒoƒo me", "NFDT": "Norfolk Island dzomeŋɔli gaƒoƒo me", "NFT": "Norfolk Island nutome gaƒoƒo me", "NOVST": "Novosibirsk dzomeŋɔli gaƒoƒo me", "NOVT": "Novosibirsk nutome gaƒoƒo me", "NPT": "Nepal gaƒoƒo me", "NRT": "Nauru gaƒoƒo me", "NST": "Newfoundland nutome gaƒoƒome", "NUT": "Niue gaƒoƒo me", "NZDT": "New Zealand kele gaƒoƒo me", "NZST": "New Zealand nutome gaƒoƒo me", "OESZ": "Ɣedzeƒe Europe ŋkekeme gaƒoƒome", "OEZ": "Ɣedzeƒe Europe gaƒoƒoɖoanyime", "OMSST": "Omsk dzomeŋɔli gaƒoƒo me", "OMST": "Omsk nutome gaƒoƒo me", "PDT": "Pacific kele gaƒoƒo me", "PDTM": "Mexican Pacific kele gaƒoƒome", "PETDT": "Petropavlovsk-Kamtsatski ŋkekeme gaƒoƒome", "PETST": "Petropavlovsk-Kamtsatski gaƒoƒoɖoanyime", "PGT": "Papua New Guinea gaƒoƒo me", "PHOT": "Phoenix Islands gaƒoƒo me", "PKT": "Pakistan nutome gaƒoƒo me", "PKT DST": "Pakistan dzomeŋɔli gaƒoƒo me", "PMDT": "St. Pierre & Miquelon kele gaƒoƒome", "PMST": "St. Pierre & Miquelon nutome gaƒoƒome", "PONT": "Ponape gaƒoƒo me", "PST": "Pacific nutome gaƒoƒo me", "PST Philippine": "Philippine nutome gaƒoƒo me", "PST Philippine DST": "Philippine dzomeŋɔli gaƒoƒo me", "PST Pitcairn": "Pitcairn gaƒoƒo me", "PSTM": "Mexican Pacific nutome gaƒoƒo me", "PWT": "Palau gaƒoƒo me", "PYST": "Paraguay dzomeŋɔli gaƒoƒo me", "PYT": "Paraguay nutome gaƒoƒo me", "PYT Korea": "Pyongyang gaƒoƒo me", "RET": "Reunion gaƒoƒo me", "ROTT": "Rothera gaƒoƒo me", "SAKST": "Sakhalin dzomeŋɔli gaƒoƒo me", "SAKT": "Sakhalin nutome gaƒoƒo me", "SAMST": "Samara ŋkekeme gaƒoƒome", "SAMT": "Samara gaƒoƒoɖoanyime", "SAST": "South Africa nutome gaƒoƒo me", "SBT": "Solomon Islands gaƒoƒo me", "SCT": "Seychelles gaƒoƒo me", "SGT": "Singapore nutome gaƒoƒo me", "SLST": "SLST", "SRT": "Suriname gaƒoƒome", "SST Samoa": "Samoa nutome gaƒoƒo me", "SST Samoa Apia": "Apia nutome gaƒoƒo me", "SST Samoa Apia DST": "Apia kele gaƒoƒo me", "SST Samoa DST": "Samoa kele gaƒoƒo me", "SYOT": "Syowa gaƒoƒo me", "TAAF": "French Southern & Antarctic gaƒoƒo me", "TAHT": "Tahiti gaƒoƒo me", "TJT": "Tajikistan gaƒoƒo me", "TKT": "Tokelau gaƒoƒo me", "TLT": "East Timor gaƒoƒo me", "TMST": "Turkmenistan dzomeŋɔli gaƒoƒo me", "TMT": "Turkmenistan nutome gaƒoƒo me", "TOST": "Tonga dzomeŋɔli gaƒoƒo me", "TOT": "Tonga nutome gaƒoƒo me", "TVT": "Tuvalu gaƒoƒo me", "TWT": "Taipei nutome gaƒoƒo me", "TWT DST": "Taipei kele gaƒoƒo me", "ULAST": "Ulan Bator dzomeŋɔli gaƒoƒo me", "ULAT": "Ulan Bator nutome gaƒoƒo me", "UYST": "Uruguay dzomeŋɔli gaƒoƒo me", "UYT": "Uruguay nutome gaƒoƒo me", "UZT": "Uzbekistan nutome gaƒoƒo me", "UZT DST": "Uzbekistan dzomeŋɔli gaƒoƒo me", "VET": "Venezuela gaƒoƒo me", "VLAST": "Vladivostok dzomeŋɔli gaƒoƒo me", "VLAT": "Vladivostok nutome gaƒoƒo me", "VOLST": "Vogograd dzomeŋɔli gaƒoƒo me", "VOLT": "Volgograd nutome gaƒoƒo me", "VOST": "Vostok gaƒoƒo me", "VUT": "Vanuatu nutome gaƒoƒo me", "VUT DST": "Vanuatu dzomeŋɔli gaƒoƒo me", "WAKT": "Wake Island gaƒoƒo me", "WARST": "Ɣetoɖoƒe Argentina dzomeŋɔli gaƒoƒo me", "WART": "Ɣetoɖoƒe Argentina nutome gaƒoƒo me", "WAST": "West Africa game", "WAT": "West Africa game", "WESZ": "Western Europe dzomeŋɔli gaƒoƒo me", "WEZ": "Western Europe nutome gaƒoƒo me", "WFT": "Wallis & Futuna gaƒoƒo me", "WGST": "West Greenland kele gaƒoƒo me", "WGT": "West Greenland nutome gaƒoƒo me", "WIB": "Western Indonesia gaƒoƒo me", "WIT": "Eastern Indonesia gaƒoƒo me", "WITA": "Central Indonesia gaƒoƒo me", "YAKST": "Yakutsk dzomeŋɔli gaƒoƒo me", "YAKT": "Yakutsk nutome gaƒoƒo me", "YEKST": "Yekaterinburg dzomeŋɔli gaƒoƒo me", "YEKT": "Yekaterinburg nutome gaƒoƒo me", "YST": "YST", "МСК": "Moscow nutome gaƒoƒo me", "اقتاۋ": "Aktau gaƒoƒoɖoanyime", "اقتاۋ قالاسى": "Aktau dzomeŋɔli gaƒoƒome", "اقتوبە": "Aktobe gaƒoƒoɖoanyime", "اقتوبە قالاسى": "Akttobe gaƒoƒome", "الماتى": "Almati gaƒoƒoɖoanyime", "الماتى قالاسى": "Almati dzomeŋɔli gaƒoƒome", "باتىس قازاق ەلى": "West Kazakhstan gaƒoƒo me", "شىعىش قازاق ەلى": "East Kazakhstan gaƒoƒo me", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Kyrgystan gaƒoƒo me", "قىزىلوردا": "Kizilɔrda gaƒoƒoɖoanyime", "قىزىلوردا قالاسى": "Kizilɔrda dzomeŋɔli gaƒoƒome", "∅∅∅": "Azores dzomeŋɔli gaƒoƒo me"},
	}
}

// Locale returns the current translators string locale
func (ee *ee) Locale() string {
	return ee.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ee'
func (ee *ee) PluralsCardinal() []locales.PluralRule {
	return ee.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ee'
func (ee *ee) PluralsOrdinal() []locales.PluralRule {
	return ee.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ee'
func (ee *ee) PluralsRange() []locales.PluralRule {
	return ee.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ee'
func (ee *ee) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ee'
func (ee *ee) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ee'
func (ee *ee) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ee *ee) MonthAbbreviated(month time.Month) string {
	if len(ee.monthsAbbreviated) == 0 {
		return ""
	}
	return ee.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ee *ee) MonthsAbbreviated() []string {
	return ee.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ee *ee) MonthNarrow(month time.Month) string {
	if len(ee.monthsNarrow) == 0 {
		return ""
	}
	return ee.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ee *ee) MonthsNarrow() []string {
	return ee.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ee *ee) MonthWide(month time.Month) string {
	if len(ee.monthsWide) == 0 {
		return ""
	}
	return ee.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ee *ee) MonthsWide() []string {
	return ee.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ee *ee) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ee.daysAbbreviated) == 0 {
		return ""
	}
	return ee.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ee *ee) WeekdaysAbbreviated() []string {
	return ee.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ee *ee) WeekdayNarrow(weekday time.Weekday) string {
	if len(ee.daysNarrow) == 0 {
		return ""
	}
	return ee.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ee *ee) WeekdaysNarrow() []string {
	return ee.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ee *ee) WeekdayShort(weekday time.Weekday) string {
	if len(ee.daysShort) == 0 {
		return ""
	}
	return ee.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ee *ee) WeekdaysShort() []string {
	return ee.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ee *ee) WeekdayWide(weekday time.Weekday) string {
	if len(ee.daysWide) == 0 {
		return ""
	}
	return ee.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ee *ee) WeekdaysWide() []string {
	return ee.daysWide
}

// Decimal returns the decimal point of number
func (ee *ee) Decimal() string {
	return ee.decimal
}

// Group returns the group of number
func (ee *ee) Group() string {
	return ee.group
}

// Group returns the minus sign of number
func (ee *ee) Minus() string {
	return ee.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ee' and handles both Whole and Real numbers based on 'v'
func (ee *ee) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ee.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ee.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ee.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ee' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ee *ee) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ee.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ee.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ee.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ee'
func (ee *ee) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ee.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ee.group[0])
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
		b = append(b, ee.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ee'
// in accounting notation.
func (ee *ee) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ee.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ee.group[0])
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

		b = append(b, ee.currencyNegativePrefix[0])

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
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ee.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ee'
func (ee *ee) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ee'
func (ee *ee) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6c, 0x69, 0x61}...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ee'
func (ee *ee) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6c, 0x69, 0x61}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ee'
func (ee *ee) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6c, 0x69, 0x61}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ee'
func (ee *ee) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ee'
func (ee *ee) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ee'
func (ee *ee) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ee'
func (ee *ee) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ee.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
