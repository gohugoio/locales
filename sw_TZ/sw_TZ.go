package sw_TZ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sw_TZ struct {
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

// New returns a new instance of translator for the 'sw_TZ' locale
func New() locales.Translator {
	return &sw_TZ{
		locale:                 "sw_TZ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " elfu ",
		currencyNegativePrefix: " elfu ",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ago", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januari", "Februari", "Machi", "Aprili", "Mei", "Juni", "Julai", "Agosti", "Septemba", "Oktoba", "Novemba", "Desemba"},
		daysWide:               []string{"Jumapili", "Jumatatu", "Jumanne", "Jumatano", "Alhamisi", "Ijumaa", "Jumamosi"},
		timezones:              map[string]string{"ACDT": "Saa za Mchana za Australia ya Kati", "ACST": "Saa za Wastani za Australia ya Kati", "ACT": "ACT", "ACWDT": "Saa za Mchana za Magharibi ya Kati ya Australia", "ACWST": "Saa za Wastani za Magharibi ya Kati ya Australia", "ADT": "Saa za Mchana za Atlantiki", "ADT Arabia": "Saa za Mchana za Arabiani", "AEDT": "Saa za Mchana za Mashariki mwa Australia", "AEST": "Saa za Wastani za Mashariki mwa Australia", "AFT": "Saa za Afghanistan", "AKDT": "Saa za Mchana za Alaska", "AKST": "Saa za Wastani za Alaska", "AMST": "Saa za Majira ya joto za Amazon", "AMST Armenia": "Saa za Majira ya joto za Armenia", "AMT": "Saa za Wastani za Amazon", "AMT Armenia": "Saa za Wastani za Armenia", "ANAST": "Saa za Kiangazi za Anadyr", "ANAT": "Saa za Wastani za Anadyr", "ARST": "Saa za Majira ya joto za Argentina", "ART": "Saa za Wastani za Argentina", "AST": "Saa za Wastani za Atlantiki", "AST Arabia": "Saa za Wastani za Uarabuni", "AWDT": "Saa za Mchana za Australia Magharibi", "AWST": "Saa za Wastani za Australia Magharibi", "AZST": "Saa za Majira ya joto za Azerbaijan", "AZT": "Saa za Wastani za Azerbaijan", "BDT Bangladesh": "Saa za Majira ya joto za Bangladesh", "BNT": "Saa za Brunei Darussalam", "BOT": "Saa za Bolivia", "BRST": "Saa za Majira ya joto za Brasilia", "BRT": "Saa za Wastani za Brasilia", "BST Bangladesh": "Saa za Wastani za Bangladesh", "BT": "Saa za Bhutan", "CAST": "CAST", "CAT": "Saa za Afrika ya Kati", "CCT": "Saa za Visiwa vya Cocos", "CDT": "Saa za Mchana za Kati", "CHADT": "Saa za Mchana za Chatham", "CHAST": "Saa za Wastani za Chatham", "CHUT": "Saa za Chuuk", "CKT": "Saa za Wastani za Visiwa vya Cook", "CKT DST": "Saa za Majira nusu ya joto za Visiwa Cook", "CLST": "Saa za Majira ya joto za Chile", "CLT": "Saa za Wastani za Chile", "COST": "Saa za Majira ya joto za Kolombia", "COT": "Saa za Wastani za Kolombia", "CST": "Saa za Wastani za Kati", "CST China": "Saa za Wastani za Uchina", "CST China DST": "Saa za Mchana za Uchina", "CVST": "Saa za Majira ya joto za Cape Verde", "CVT": "Saa za Wastani za Cape Verde", "CXT": "Saa za Kisiwa cha Krismasi", "ChST": "Saa za Wastani za Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Saa za Mchana za Kuba", "CuST": "Saa za Wastani ya Kuba", "DAVT": "Saa za Davis", "DDUT": "Saa za Dumont-d’Urville", "EASST": "Saa za Majira ya joto za Kisiwa cha Easter", "EAST": "Saa za Wastani za Kisiwa cha Easter", "EAT": "Saa za Afrika Mashariki", "ECT": "Saa za Ekwado", "EDT": "Saa za Mchana za Mashariki", "EGDT": "Saa za Majira ya joto za Greenland Mashariki", "EGST": "Saa za Wastani za Greenland Mashariki", "EST": "Saa za Wastani za Mashariki", "FEET": "Saa za Mashariki zaidi mwa Ulaya", "FJT": "Saa za Wastani za Fiji", "FJT Summer": "Saa za Majira ya joto za Fiji", "FKST": "Saa za Majira ya joto za Visiwa vya Falkland", "FKT": "Saa za Wastani za Visiwa vya Falkland", "FNST": "Saa za Majira ya joto za Fernando de Noronha", "FNT": "Saa za Wastani za Fernando de Noronha", "GALT": "Saa za Galapagos", "GAMT": "Saa za Gambier", "GEST": "Saa za Majira ya joto za Jojia", "GET": "Saa za Wastani za Jojia", "GFT": "Saa za Guiana ya Ufaransa", "GIT": "Saa za Visiwa vya Gilbert", "GMT": "Saa za Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Saa za Georgia Kusini", "GST Guam": "GST Guam", "GYT": "Saa za Guyana", "HADT": "Saa za Mchana za Hawaii-Aleutian", "HAST": "Saa za Wastani za Hawaii-Aleutian", "HKST": "Saa za Majira ya joto za Hong Kong", "HKT": "Saa za Wastani za Hong Kong", "HOVST": "Saa za Majira ya joto za Hovd", "HOVT": "Saa za Wastani za Hovd", "ICT": "Saa za Indochina", "IDT": "Saa za Mchana za Israeli", "IOT": "Saa za Bahari Hindi", "IRKST": "Saa za Majira ya joto za Irkutsk", "IRKT": "Saa za Wastani za Irkutsk", "IRST": "Saa za Wastani za Iran", "IRST DST": "Saa za Mchana za Iran", "IST": "Saa za Wastani za India", "IST Israel": "Saa za Wastani za Israeli", "JDT": "Saa za Mchana za Japan", "JST": "Saa za Wastani za Japani", "KOST": "Saa za Kosrae", "KRAST": "Saa za Majira ya joto za Krasnoyarsk", "KRAT": "Saa za Wastani za Krasnoyask", "KST": "Saa za Wastani za Korea", "KST DST": "Saa za Mchana za Korea", "LHDT": "Saa za Mchana za Lord Howe", "LHST": "Saa za Wastani za Lord Howe", "LINT": "Saa za Visiwa vya Line", "MAGST": "Saa za Majira ya joto za Magadan", "MAGT": "Saa za Wastani za Magadan", "MART": "Saa za Marquesas", "MAWT": "Saa za Mawson", "MDT": "MDT", "MESZ": "Saa za Majira ya joto za Ulaya ya Kati", "MEZ": "Saa za Wastani za Ulaya ya Kati", "MHT": "Saa za Visiwa vya Marshall", "MMT": "Saa za Myanmar", "MSD": "Saa za Majira ya joto za Moscow", "MST": "MST", "MUST": "Saa za Majira ya joto za Morisi", "MUT": "Saa za Wastani za Morisi", "MVT": "Saa za Maldives", "MYT": "Saa za Malaysia", "NCT": "Saa za Wastani za New Caledonia", "NDT": "Saa za Mchana za Newfoundland", "NDT New Caledonia": "Saa za Majira ya joto za New Caledonia", "NFDT": "Saa za Majira ya joto za Kisiwa cha Norfolk", "NFT": "Saa za Wastani za Kisiwa cha Norfolk", "NOVST": "Saa za Majira ya joto za Novosibirsk", "NOVT": "Saa za Wastani za Novosibirsk", "NPT": "Saa za Nepal", "NRT": "Saa za Nauru", "NST": "Saa za Wastani za Newfoundland", "NUT": "Saa za Niue", "NZDT": "Saa za Mchana za New Zealand", "NZST": "Saa za Wastani za New Zealand", "OESZ": "Saa za Majira ya joto za Mashariki mwa Ulaya", "OEZ": "Saa za Wastani za Mashariki mwa Ulaya", "OMSST": "Saa za Majira ya joto za Omsk", "OMST": "Saa za Wastani za Omsk", "PDT": "Saa za Mchana za Pasifiki", "PDTM": "Saa za mchana za pasifiki za Meksiko", "PETDT": "Saa za Kiangazi za Petropavlovsk-Kamchatski", "PETST": "Saa za Wastani za Petropavlovsk-Kamchatski", "PGT": "Saa za Papua New Guinea", "PHOT": "Saa za Visiwa vya Phoenix", "PKT": "Saa za Wastani za Pakistan", "PKT DST": "Saa za Majira ya joto za Pakistan", "PMDT": "Saa za Mchana za Saint-Pierre na Miquelon", "PMST": "Saa za Wastani ya Saint-Pierre na Miquelon", "PONT": "Saa za Ponape", "PST": "Saa za Wastani za Pasifiki", "PST Philippine": "Saa za Wastani za Ufilipino", "PST Philippine DST": "Saa za Majira ya joto za Ufilipino", "PST Pitcairn": "Saa za Pitcairn", "PSTM": "Saa za wastani za pasifiki za Meksiko", "PWT": "Saa za Palau", "PYST": "Saa za Majira ya joto za Paragwai", "PYT": "Saa za Wastani za Paragwai", "PYT Korea": "Saa za Pyongyang", "RET": "Saa za Reunion", "ROTT": "Saa za Rothera", "SAKST": "Saa za Majira ya joto za Sakhalin", "SAKT": "Saa za Wastani za Sakhalin", "SAMST": "Saa za Kiangazi za Samara", "SAMT": "Saa za Wastani za Samara", "SAST": "Saa za Wastani za Afrika Kusini", "SBT": "Saa za Visiwa vya Solomon", "SCT": "Saa za Ushelisheli", "SGT": "Saa za Wastani za Singapore", "SLST": "SLST", "SRT": "Saa za Suriname", "SST Samoa": "Saa za Wastani za Samoa", "SST Samoa Apia": "Saa za Wastani za Apia", "SST Samoa Apia DST": "Saa za Mchana za Apia", "SST Samoa DST": "Saa za Majira ya joto za Samoa", "SYOT": "Saa za Syowa", "TAAF": "Saa za Kusini mwa Ufaransa na Antaktiki", "TAHT": "Saa za Tahiti", "TJT": "Saa za Tajikistan", "TKT": "Saa za Tokelau", "TLT": "Saa za Timor Mashariki", "TMST": "Saa za Majira ya joto za Turkmenistan", "TMT": "Saa za Wastani za Turkmenistan", "TOST": "Saa za Majira ya joto za Tonga", "TOT": "Saa za Wastani za Tonga", "TVT": "Saa za Tuvalu", "TWT": "Saa za Wastani za Taipei", "TWT DST": "Saa za Mchana za Taipei", "ULAST": "Saa za Majira ya joto za Ulan Bator", "ULAT": "Saa za Wastani za Ulan Bator", "UYST": "Saa za Majira ya joto za Urugwai", "UYT": "Saa za Wastani za Urugwai", "UZT": "Saa za Wastani za Uzbekistan", "UZT DST": "Saa za Majira ya joto za Uzbekistan", "VET": "Saa za Venezuela", "VLAST": "Saa za Majira ya joto za Vladivostok", "VLAT": "Saa za Wastani za Vladivostok", "VOLST": "Saa za Majira ya joto za Volgograd", "VOLT": "Saa za Wastani za Volgograd", "VOST": "Saa za Vostok", "VUT": "Saa za Wastani za Vanuatu", "VUT DST": "Saa za Majira ya joto za Vanuatu", "WAKT": "Saa za Kisiwa cha Wake", "WARST": "Saa za Majira ya joto za Magharibi mwa Argentina", "WART": "Saa za Wastani za Magharibi mwa Argentina", "WAST": "Saa za Afrika Magharibi", "WAT": "Saa za Afrika Magharibi", "WESZ": "Saa za Majira ya joto za Magharibi mwa Ulaya", "WEZ": "Saa za Wastani za Magharibi mwa Ulaya", "WFT": "Saa za Wallis na Futuna", "WGST": "Saa za Majira ya joto za Greenland Magharibi", "WGT": "Saa za Wastani za Greenland Magharibi", "WIB": "Saa za Magharibi mwa Indonesia", "WIT": "Saa za Mashariki mwa Indonesia", "WITA": "Saa za Indonesia ya Kati", "YAKST": "Saa za Majira ya joto za Yakutsk", "YAKT": "Saa za Wastani za Yakutsk", "YEKST": "Saa za Majira ya joto za Yekaterinburg", "YEKT": "Saa za Wastani za Yekaterinburg", "YST": "Saa za Yukon", "МСК": "Saa za Wastani za Moscow", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Saa za Kazakhstan Magharibi", "شىعىش قازاق ەلى": "Saa za Kazakhstan Mashariki", "قازاق ەلى": "Saa za Kazakhstan", "قىرعىزستان": "Saa za Kyrgystan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Saa za Majira ya joto za Azores"},
	}
}

// Locale returns the current translators string locale
func (sw *sw_TZ) Locale() string {
	return sw.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sw_TZ'
func (sw *sw_TZ) PluralsCardinal() []locales.PluralRule {
	return sw.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sw_TZ'
func (sw *sw_TZ) PluralsOrdinal() []locales.PluralRule {
	return sw.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sw_TZ'
func (sw *sw_TZ) PluralsRange() []locales.PluralRule {
	return sw.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sw_TZ'
func (sw *sw_TZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sw_TZ'
func (sw *sw_TZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sw_TZ'
func (sw *sw_TZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := sw.CardinalPluralRule(num1, v1)
	end := sw.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sw *sw_TZ) MonthAbbreviated(month time.Month) string {
	if len(sw.monthsAbbreviated) == 0 {
		return ""
	}
	return sw.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sw *sw_TZ) MonthsAbbreviated() []string {
	return sw.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sw *sw_TZ) MonthNarrow(month time.Month) string {
	if len(sw.monthsNarrow) == 0 {
		return ""
	}
	return sw.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sw *sw_TZ) MonthsNarrow() []string {
	return sw.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sw *sw_TZ) MonthWide(month time.Month) string {
	if len(sw.monthsWide) == 0 {
		return ""
	}
	return sw.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sw *sw_TZ) MonthsWide() []string {
	return sw.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sw *sw_TZ) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sw.daysAbbreviated) == 0 {
		return ""
	}
	return sw.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sw *sw_TZ) WeekdaysAbbreviated() []string {
	return sw.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sw *sw_TZ) WeekdayNarrow(weekday time.Weekday) string {
	if len(sw.daysNarrow) == 0 {
		return ""
	}
	return sw.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sw *sw_TZ) WeekdaysNarrow() []string {
	return sw.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sw *sw_TZ) WeekdayShort(weekday time.Weekday) string {
	if len(sw.daysShort) == 0 {
		return ""
	}
	return sw.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sw *sw_TZ) WeekdaysShort() []string {
	return sw.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sw *sw_TZ) WeekdayWide(weekday time.Weekday) string {
	if len(sw.daysWide) == 0 {
		return ""
	}
	return sw.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sw *sw_TZ) WeekdaysWide() []string {
	return sw.daysWide
}

// Decimal returns the decimal point of number
func (sw *sw_TZ) Decimal() string {
	return sw.decimal
}

// Group returns the group of number
func (sw *sw_TZ) Group() string {
	return sw.group
}

// Group returns the minus sign of number
func (sw *sw_TZ) Minus() string {
	return sw.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sw_TZ' and handles both Whole and Real numbers based on 'v'
func (sw *sw_TZ) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sw_TZ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sw *sw_TZ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sw.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sw_TZ'
func (sw *sw_TZ) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sw.currencies[currency]
	l := len(s) + len(symbol) + 10

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(sw.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, sw.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sw_TZ'
// in accounting notation.
func (sw *sw_TZ) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sw.currencies[currency]
	l := len(s) + len(symbol) + 10

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sw.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sw.currencyNegativePrefix[j])
		}

		b = append(b, sw.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sw.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, sw.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, sw.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sw_TZ'
func (sw *sw_TZ) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sw.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
