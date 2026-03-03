package ceb_PH

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ceb_PH struct {
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

// New returns a new instance of translator for the 'ceb_PH' locale
func New() locales.Translator {
	return &ceb_PH{
		locale:                 "ceb_PH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "US $", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ene", "Peb", "Mar", "Abr", "May", "Hun", "Hul", "Ago", "Sep", "Okt", "Nob", "Dis"},
		monthsNarrow:           []string{"", "E", "P", "M", "A", "M", "H", "H", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Enero", "Pebrero", "Marso", "Abril", "Mayo", "Hunyo", "Hulyo", "Agosto", "Septiyembre", "Oktubre", "Nobyembre", "Disyembre"},
		daysAbbreviated:        []string{"Dom", "Lun", "Mar", "Miy", "Huw", "Biy", "Sab"},
		daysNarrow:             []string{"D", "L", "M", "M", "H", "B", "S"},
		daysWide:               []string{"Domingo", "Lunes", "Martes", "Miyerkules", "Huwebes", "Biyernes", "Sabado"},
		timezones:              map[string]string{"ACDT": "Australian Central Daylight Time", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Australian Central Western Daylight Time", "ACWST": "Australian Central Western Standard Time", "ADT": "Daylight nga Oras sa Atlantic", "ADT Arabia": "Arabian Daylight Time", "AEDT": "Australian Eastern Daylight Time", "AEST": "Australian Eastern Standard Time", "AFT": "Oras sa Afghanistan", "AKDT": "Daylight nga Oras sa Alaska", "AKST": "Standard nga Oras sa Alaska", "AMST": "Amazon Summer Time", "AMST Armenia": "Armenia Summer Time", "AMT": "Amazon Standard Time", "AMT Armenia": "Armenia Standard Time", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Argentina Summer Time", "ART": "Argentina Standard Time", "AST": "Standard nga Oras sa Atlantic", "AST Arabia": "Arabian Standard Time", "AWDT": "Australian Western Daylight Time", "AWST": "Australian Western Standard Time", "AZST": "Azerbaijan Summer Time", "AZT": "Azerbaijan Standard Time", "BDT Bangladesh": "Bangladesh Summer Time", "BNT": "Oras sa Brunei Darussalam", "BOT": "Oras sa Bolivia", "BRST": "Brasilia Summer Time", "BRT": "Brasilia Standard Time", "BST Bangladesh": "Bangladesh Standard Time", "BT": "Oras sa Bhutan", "CAST": "CAST", "CAT": "Oras sa Central Africa", "CCT": "Oras sa Cocos Islands", "CDT": "Central Daylight Time", "CHADT": "Chatham Daylight Time", "CHAST": "Chatham Standard Time", "CHUT": "Oras sa Chuuk", "CKT": "Cook Islands Standard Time", "CKT DST": "Cook Islands Summer Time", "CLST": "Chile Summer Time", "CLT": "Chile Standard Time", "COST": "Colombia Summer Time", "COT": "Colombia Standard Time", "CST": "Central Standard Time", "CST China": "China Standard Time", "CST China DST": "China Daylight Time", "CVST": "Cape Verde Summer Time", "CVT": "Cape Verde Standard Time", "CXT": "Oras sa Christmas Island", "ChST": "Chamorro Standard Time", "ChST NMI": "ChST NMI", "CuDT": "Cuba Daylight Time", "CuST": "Cuba Standard Time", "DAVT": "Oras sa Davis", "DDUT": "Oras sa Dumont-d’Urville", "EASST": "Easter Island Summer Time", "EAST": "Easter Island Standard Time", "EAT": "Oras sa East Africa", "ECT": "Oras sa Ecuador", "EDT": "Eastern Daylight Time", "EGDT": "East Greenland Summer Time", "EGST": "East Greenland Standard Time", "EST": "Eastern Standard Time", "FEET": "Oras sa Further-eastern Europe", "FJT": "Fiji Standard Time", "FJT Summer": "Fiji Summer Time", "FKST": "Falkland Islands Summer Time", "FKT": "Falkland Islands Standard Time", "FNST": "Fernando de Noronha Summer Time", "FNT": "Fernando de Noronha Standard Time", "GALT": "Oras sa Galapagos", "GAMT": "Oras sa Gambier", "GEST": "Georgia Summer Time", "GET": "Georgia Standard Time", "GFT": "Oras sa French Guiana", "GIT": "Oras sa Gilbert Islands", "GMT": "Oras sa Greenwich Mean", "GNSST": "GNSST", "GNST": "GNST", "GST": "Gulf Standard Time", "GST Guam": "GST Guam", "GYT": "Oras sa Guyana", "HADT": "Hawaii-Aleutian Standard Time", "HAST": "Hawaii-Aleutian Standard Time", "HKST": "Hong Kong Summer Time", "HKT": "Hong Kong Standard Time", "HOVST": "Khovd Summer Time", "HOVT": "Khovd Standard Time", "ICT": "Oras sa Indochina", "IDT": "Israel Daylight Time", "IOT": "Oras sa Indian Ocean", "IRKST": "Irkutsk Summer Time", "IRKT": "Irkutsk Standard Time", "IRST": "Iran Standard Time", "IRST DST": "Iran Daylight Time", "IST": "India Standard Time", "IST Israel": "Israel Standard Time", "JDT": "Japan Daylight Time", "JST": "Japan Standard Time", "KOST": "Oras sa Kosrae", "KRAST": "Krasnoyarsk Summer Time", "KRAT": "Krasnoyarsk Standard Time", "KST": "Korean Standard Time", "KST DST": "Korean Daylight Time", "LHDT": "Lord Howe Daylight Time", "LHST": "Lord Howe Standard Time", "LINT": "Oras sa Line Islands", "MAGST": "Magadan Summer Time", "MAGT": "Magadan Standard Time", "MART": "Oras sa Marquesas", "MAWT": "Oras sa Mawson", "MDT": "Mountain Daylight Time", "MESZ": "Central European Summer Time", "MEZ": "Central European Standard Time", "MHT": "Oras sa Marshall Islands", "MMT": "Oras sa Myanmar", "MSD": "Moscow Summer Time", "MST": "Mountain Standard Time", "MUST": "Mauritius Summer Time", "MUT": "Mauritius Standard Time", "MVT": "Oras sa Maldives", "MYT": "Oras sa Malaysia", "NCT": "New Caledonia Standard Time", "NDT": "Newfoundland Daylight Time", "NDT New Caledonia": "New Caledonia Summer Time", "NFDT": "Norfolk Island Daylight Time", "NFT": "Norfolk Island Standard Time", "NOVST": "Novosibirsk Summer Time", "NOVT": "Novosibirsk Standard Time", "NPT": "Oras sa Nepal", "NRT": "Oras sa Nauru", "NST": "Newfoundland Standard Time", "NUT": "Oras sa Niue", "NZDT": "New Zealand Daylight Time", "NZST": "New Zealand Standard Time", "OESZ": "Eastern European Summer Time", "OEZ": "Eastern European Standard Time", "OMSST": "Omsk Summer Time", "OMST": "Omsk Standard Time", "PDT": "Daylight nga Oras sa Pasipiko", "PDTM": "Mexican Pacific Daylight Time", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Oras sa Papua New Guinea", "PHOT": "Oras sa Phoenix Islands", "PKT": "Pakistan Standard Time", "PKT DST": "Pakistan Summer Time", "PMDT": "St. Pierre & Miquelon Daylight Time", "PMST": "St. Pierre & Miquelon Standard Time", "PONT": "Oras sa Ponape", "PST": "Standard nga Oras sa Pasipiko", "PST Philippine": "Philippine Standard Time", "PST Philippine DST": "Philippine Summer Time", "PST Pitcairn": "Oras sa Pitcairn", "PSTM": "Mexican Pacific Standard Time", "PWT": "Oras sa Palau", "PYST": "Paraguay Summer Time", "PYT": "Paraguay Standard Time", "PYT Korea": "Oras sa Pyongyang", "RET": "Oras sa Reunion", "ROTT": "Oras sa Rothera", "SAKST": "Sakhalin Summer Time", "SAKT": "Sakhalin Standard Time", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "South Africa Standard Time", "SBT": "Oras sa Solomon Islands", "SCT": "Oras sa Seychelles", "SGT": "Singapore Standard Time", "SLST": "SLST", "SRT": "Oras sa Suriname", "SST Samoa": "American Samoa Standard Time", "SST Samoa Apia": "Samoa Standard Time", "SST Samoa Apia DST": "Samoa Daylight Time", "SST Samoa DST": "American Samoa Daylight Time", "SYOT": "Oras sa Syowa", "TAAF": "Oras sa French Southern ug Antarctic", "TAHT": "Oras sa Tahiti", "TJT": "Oras sa Tajikistan", "TKT": "Oras sa Tokelau", "TLT": "Oras sa East Timor", "TMST": "Turkmenistan Summer Time", "TMT": "Turkmenistan Standard Time", "TOST": "Tonga Summer Time", "TOT": "Tonga Standard Time", "TVT": "Oras sa Tuvalu", "TWT": "Taiwan Standard Time", "TWT DST": "Taiwan Daylight Time", "ULAST": "Ulaanbaatar Summer Time", "ULAT": "Ulaanbaatar Standard Time", "UYST": "Uruguay Summer Time", "UYT": "Uruguay Standard Time", "UZT": "Uzbekistan Standard Time", "UZT DST": "Uzbekistan Summer Time", "VET": "Oras sa Venezuela", "VLAST": "Vladivostok Summer Time", "VLAT": "Vladivostok Standard Time", "VOLST": "Volgograd Summer Time", "VOLT": "Volgograd Standard Time", "VOST": "Oras sa Vostok", "VUT": "Vanuatu Standard Time", "VUT DST": "Vanuatu Summer Time", "WAKT": "Oras sa Wake Island", "WARST": "Western Argentina Summer Time", "WART": "Western Argentina Standard Time", "WAST": "Oras sa West Africa", "WAT": "Oras sa West Africa", "WESZ": "Western European Summer Time", "WEZ": "Western European Standard Time", "WFT": "Oras sa Wallis & Futuna", "WGST": "West Greenland Summer Time", "WGT": "West Greenland Standard Time", "WIB": "Oras sa Western Indonesia", "WIT": "Oras sa Eastern Indonesia", "WITA": "Oras sa Central Indonesia", "YAKST": "Yakutsk Summer Time", "YAKT": "Yakutsk Standard Time", "YEKST": "Yekaterinburg Summer Time", "YEKT": "Yekaterinburg Standard Time", "YST": "Oras sa Yukon", "МСК": "Moscow Standard Time", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Oras sa West Kazakhstan", "شىعىش قازاق ەلى": "Oras sa East Kazakhstan", "قازاق ەلى": "Oras sa Kazakhstan", "قىرعىزستان": "Oras sa Kyrgyzstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Azores Summer Time"},
	}
}

// Locale returns the current translators string locale
func (ceb *ceb_PH) Locale() string {
	return ceb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ceb_PH'
func (ceb *ceb_PH) PluralsCardinal() []locales.PluralRule {
	return ceb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ceb_PH'
func (ceb *ceb_PH) PluralsOrdinal() []locales.PluralRule {
	return ceb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ceb_PH'
func (ceb *ceb_PH) PluralsRange() []locales.PluralRule {
	return ceb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ceb_PH'
func (ceb *ceb_PH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	fMod10 := f % 10

	if (v == 0 && (i == 1 || i == 2 || i == 3)) || (v == 0 && (iMod10 != 4 && iMod10 != 6 && iMod10 != 9)) || (v != 0 && (fMod10 != 4 && fMod10 != 6 && fMod10 != 9)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ceb_PH'
func (ceb *ceb_PH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ceb_PH'
func (ceb *ceb_PH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ceb *ceb_PH) MonthAbbreviated(month time.Month) string {
	if len(ceb.monthsAbbreviated) == 0 {
		return ""
	}
	return ceb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ceb *ceb_PH) MonthsAbbreviated() []string {
	return ceb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ceb *ceb_PH) MonthNarrow(month time.Month) string {
	if len(ceb.monthsNarrow) == 0 {
		return ""
	}
	return ceb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ceb *ceb_PH) MonthsNarrow() []string {
	return ceb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ceb *ceb_PH) MonthWide(month time.Month) string {
	if len(ceb.monthsWide) == 0 {
		return ""
	}
	return ceb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ceb *ceb_PH) MonthsWide() []string {
	return ceb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ceb *ceb_PH) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ceb.daysAbbreviated) == 0 {
		return ""
	}
	return ceb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ceb *ceb_PH) WeekdaysAbbreviated() []string {
	return ceb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ceb *ceb_PH) WeekdayNarrow(weekday time.Weekday) string {
	if len(ceb.daysNarrow) == 0 {
		return ""
	}
	return ceb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ceb *ceb_PH) WeekdaysNarrow() []string {
	return ceb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ceb *ceb_PH) WeekdayShort(weekday time.Weekday) string {
	if len(ceb.daysShort) == 0 {
		return ""
	}
	return ceb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ceb *ceb_PH) WeekdaysShort() []string {
	return ceb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ceb *ceb_PH) WeekdayWide(weekday time.Weekday) string {
	if len(ceb.daysWide) == 0 {
		return ""
	}
	return ceb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ceb *ceb_PH) WeekdaysWide() []string {
	return ceb.daysWide
}

// Decimal returns the decimal point of number
func (ceb *ceb_PH) Decimal() string {
	return ceb.decimal
}

// Group returns the group of number
func (ceb *ceb_PH) Group() string {
	return ceb.group
}

// Group returns the minus sign of number
func (ceb *ceb_PH) Minus() string {
	return ceb.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ceb_PH' and handles both Whole and Real numbers based on 'v'
func (ceb *ceb_PH) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ceb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ceb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ceb_PH' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ceb *ceb_PH) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ceb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ceb.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ceb_PH'
func (ceb *ceb_PH) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ceb.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ceb.group[0])
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
		b = append(b, ceb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ceb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ceb_PH'
// in accounting notation.
func (ceb *ceb_PH) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ceb.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ceb.group[0])
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

		b = append(b, ceb.currencyNegativePrefix[0])

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
			b = append(b, ceb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ceb.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ceb.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ceb.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ceb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ceb.monthsWide[t.Month()]...)
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

// FmtTimeShort returns the short time representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ceb_PH'
func (ceb *ceb_PH) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ceb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
