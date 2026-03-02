package ha_NE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ha_NE struct {
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

// New returns a new instance of translator for the 'ha_NE' locale
func New() locales.Translator {
	return &ha_NE{
		locale:                 "ha_NE",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "K",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "K",
		monthsAbbreviated:      []string{"", "Jan", "Fab", "Mar", "Afi", "May", "Yun", "Yul", "Agu", "Sat", "Okt", "Nuw", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "Y", "Y", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Janairu", "Faburairu", "Maris", "Afirilu", "Mayu", "Yuni", "Yuli", "Agusta", "Satumba", "Oktoba", "Nuwamba", "Disamba"},
		daysAbbreviated:        []string{"Lah", "Lit", "Tal", "Lar", "Alh", "Jum", "Asa"},
		daysNarrow:             []string{"L", "L", "T", "L", "A", "J", "A"},
		daysShort:              []string{"Lh", "Li", "Ta", "Lr", "Al", "Ju", "As"},
		daysWide:               []string{"Lahadi", "Litinin", "Talata", "Laraba", "Alhamis", "Jummaʼa", "Asabar"},
		timezones:              map[string]string{"ACDT": "Lokacin Rana na Tsakiyar Austiraliya", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Lokacin Rana na Yammacin Tsakiyar Austiraliya", "ACWST": "Tsayayyen Lokacin Yammacin Tsakiyar Austiraliya", "ADT": "Lokacin Rana na Kanada, Puerto Rico da Virgin Islands", "ADT Arabia": "Lokacin Rana na Arebiya", "AEDT": "Lokacin Rana na Gabashin Austiraliya", "AEST": "Tsayayyen lokacin Gabashin Australia", "AFT": "Lokacin Afghanistan", "AKDT": "Lokacin Rana na Alaska", "AKST": "Tsayayyen Lokacin Alaska", "AMST": "Lokacin Bazara na Amazon", "AMST Armenia": "Lokacin Bazara na Armenia", "AMT": "Tsayayyen Lokacin Amazon", "AMT Armenia": "Tsayayyen Lokacin Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Lokacin Bazara na Argentina", "ART": "Tsayayyen Lokacin Argentina", "AST": "Tsayayyen Lokacin Kanada, Puerto Rico da Virgin Islands", "AST Arabia": "Tsayayyen lokacin Arebiya", "AWDT": "Lokacin Rana na Yammacin Austiralia", "AWST": "Tsayayyen Lokacin Yammacin Austiralia", "AZST": "Lokacin Bazara na Azerbaijan", "AZT": "Tsayayyen Lokacin Azerbaijan", "BDT Bangladesh": "Lokacin Bazara na Bangladesh", "BNT": "Lokacin Brunei Darussalam", "BOT": "Lokacin Bolivia", "BRST": "Lokacin Bazara na Brasillia", "BRT": "Tsayayyen Lokacin Brasillia", "BST Bangladesh": "Tsayayyen Lokacin Bangladesh", "BT": "Lokacin Bhutan", "CAST": "CAST", "CAT": "Lokacin Afirka ta Tsakiya", "CCT": "Lokacin Cocos Islands", "CDT": "Lokacin Rana dake Arewacin Amurika ta Tsakiya", "CHADT": "Lokacin Rana na Chatham", "CHAST": "Tsayayyen Lokacin Chatham", "CHUT": "Lokacin Chuuk", "CKT": "Tsayayyen Lokacin Cook Islands", "CKT DST": "Rabin Lokacin Bazara na Cook Islands", "CLST": "Lokacin Bazara na Chile", "CLT": "Tsayayyen Lokacin Chile", "COST": "Lokacin Bazara na Colombia", "COT": "Tsayayyen Lokacin Colombia", "CST": "Tsayayyen Lokaci dake Arewacin Amurika ta Tsakiya", "CST China": "Tsayayyen Lokacin Sin", "CST China DST": "Lokacin Rana na Sin", "CVST": "Lokacin Bazara na Cape Verde", "CVT": "Tsayayyen lokacin Cape Verde", "CXT": "Lokacin Christmas Island", "ChST": "Tsayayyen Lokacin Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Lokacin Rana na Kuba", "CuST": "Tsayayyen Lokacin Kuba", "DAVT": "Lokacin Davis", "DDUT": "Lokacin Dumont-d’Urville", "EASST": "Lokacin Bazara na Easter Island", "EAST": "Tsayayyen Lokacin Easter Island", "EAT": "Lokacin Gabashin Afirka", "ECT": "Lokacin Ecuador", "EDT": "Lokacin Rana ta Gabas dake Arewacin Amurika", "EGDT": "Lokacin Rana na Gabashin Greenland", "EGST": "Tsayayyen Lokacin Gabashin Greenland", "EST": "Tsayayyen Lokacin Gabas dake Arewacin Amurika", "FEET": "Lokacin Gabashin Turai mai Nisa", "FJT": "Tsayayyen Lokacin Fiji", "FJT Summer": "Lokacin Bazara na Fiji", "FKST": "Lokacin Bazara na Falkland Islands", "FKT": "Tsayayyen Lokacin Falkland Islands", "FNST": "Lokacin Bazara na Fernando de Noronha", "FNT": "Tsayayyen Lokacin Fernando de Noronha", "GALT": "Lokacin Galapagos", "GAMT": "Lokacin Gambier", "GEST": "Lokacin Bazara na Georgia", "GET": "Tsayayyen Lokacin Georgia", "GFT": "Lokacin French Guiana", "GIT": "Lokacin Gilbert Islands", "GMT": "Lokacin Greenwich a Ingila", "GNSST": "GNSST", "GNST": "GNST", "GST": "Tsayayyen lokacin Gulf", "GST Guam": "GST Guam", "GYT": "Lokacin Guyana", "HADT": "Tsayayyen Lokacin Hawaii-Aleutian", "HAST": "Tsayayyen Lokacin Hawaii-Aleutian", "HKST": "Lokacin Bazara na Hong Kong", "HKT": "Tsayayyen Lokacin Hong Kong", "HOVST": "Lokacin Bazara na Hovd", "HOVT": "Tsayayyen Lokacin Hovd", "ICT": "Lokacin Indochina", "IDT": "Lokacin Hasken Rana na Israʼila", "IOT": "Lokacin Tekun Indiya", "IRKST": "Lokacin Bazara na Irkutsk", "IRKT": "Tsayayyen Lokacin Irkutsk", "IRST": "Tsayayyen Lokacin Iran", "IRST DST": "Lokacin Rana na Iran", "IST": "Tsayayyen lokacin Indiya", "IST Israel": "Tsayayyen lokacin Israʼila", "JDT": "Lokacin Hasken Rana na Japan", "JST": "Tsayayyen lokacin Japan", "KOST": "Lokacin Kosrae", "KRAST": "Lokacin Bazara na Krasnoyarsk", "KRAT": "Tsayayyen Lokacin Krasnoyarsk", "KST": "Tsayayyen Lokacin Koriya", "KST DST": "Lokacin Rana na Koriya", "LHDT": "Lokacin Rana na Vote Lord Howe", "LHST": "Tsayayyen Lokacin Lord Howe", "LINT": "Lokacin Line Islands", "MAGST": "Lokacin Bazara na Magadan", "MAGT": "Tsayayyen Lokacin Magadan", "MART": "Lokacin Marquesas", "MAWT": "Lokacin Mawson", "MDT": "Lokacin Rana na Tsaunin Arewacin Amurka", "MESZ": "Tsakiyar bazara a lokaci turai", "MEZ": "Ida Tsakiyar a Lokaci Turai", "MHT": "Lokacin Marshall Islands", "MMT": "Lokacin Myanmar", "MSD": "Lokacin Bazara na Moscow", "MST": "Lokaci tsayayye na tsauni a Arewacin Amurica", "MUST": "Lokacin Bazara na Mauritius", "MUT": "Tsayayyen Lokacin Mauritius", "MVT": "Lokacin Maldives", "MYT": "Lokacin Malaysia", "NCT": "Tsayayyen Lokacin New Caledonia", "NDT": "Lokaci rana ta Newfoundland", "NDT New Caledonia": "Lokacin Bazara na New Caledonia", "NFDT": "Lokacin Rana na Norfolk Island", "NFT": "Tsayayyen Lokacin Norfolk Island", "NOVST": "Lokacin Bazara na Novosibirsk", "NOVT": "Novosibirsk Standard Time", "NPT": "Lokacin Nepal", "NRT": "Lokacin Nauru", "NST": "Lokaci Tsayayye ta Newfoundland", "NUT": "Lokacin Niue", "NZDT": "Lokacin Rana na New Zealand", "NZST": "Tsayayyen Lokacin New Zealand", "OESZ": "Gabas a lokaci turai da bazara", "OEZ": "Ida lokaci a turai gabas", "OMSST": "Lokacin Bazara na Omsk", "OMST": "Tsayayyen Lokacin Omsk", "PDT": "Lokacin Rana na Arewacin Amurka", "PDTM": "Lokacin Rana na Mekziko Pacific", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Lokacin Papua New Guinea", "PHOT": "Lokacin Phoenix Islands", "PKT": "Tsayayyen Lokacin Pakistan", "PKT DST": "Lokacin Bazara na Pakistan", "PMDT": "Lokacin Rana na St. Pierre da Miquelon", "PMST": "Tsayayyen Lokacin St. Pierre da Miquelon", "PONT": "Lokacin Ponape", "PST": "Lokaci Tsayayye na Arewacin Amurika", "PST Philippine": "Tsayayyen Lokacin Philippine", "PST Philippine DST": "Lokacin Bazara na Philippine", "PST Pitcairn": "Lokacin Pitcairn", "PSTM": "Tsayayyen Lokacin Mekziko Pacific", "PWT": "Lokacin Palau", "PYST": "Lokacin Bazara na Paraguay", "PYT": "Tsayayyen Lokacin Paraguay", "PYT Korea": "Lokacin Pyongyang", "RET": "Lokacin Réunion", "ROTT": "Lokacin Rothera", "SAKST": "Lokacin Bazara na Sakhalin", "SAKT": "Tsayayyen Lokacin Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Tsayayyen Lokacin Afirka ta Kudu", "SBT": "Lokacin Rana na Solomon", "SCT": "Lokacin Seychelles", "SGT": "Tsayayyen Lokacin Singapore", "SLST": "SLST", "SRT": "Lokacin Suriname", "SST Samoa": "Tsayayyen Lokacin Samoa", "SST Samoa Apia": "Tsayayyen Lokacin Apia", "SST Samoa Apia DST": "Lokacin Rana na Apia", "SST Samoa DST": "Lokacin Rana na Vote Samoa", "SYOT": "Lokacin Syowa", "TAAF": "Lokacin Kudancin Faransa da Antarctic", "TAHT": "Lokacin Tahiti", "TJT": "Lokacin Tajikistan", "TKT": "Lokacin Tokelau", "TLT": "Lokacin East Timor", "TMST": "Lokacin Bazara na Turkmenistan", "TMT": "Tsayayyen Lokacin Turkmenistan", "TOST": "Lokacin Bazara na Tonga", "TOT": "Tsayayyen Lokacin Tonga", "TVT": "Lokacin Tuvalu", "TWT": "Tsayayyen Lokacin Taipei", "TWT DST": "Lokacin Rana na Taipei", "ULAST": "Lokacin Bazara na Ulaanbaatar", "ULAT": "Tsayayyen Lokacin Ulaanbaatar", "UYST": "Lokacin Bazara na Uruguay", "UYT": "Tsayayyen Lokacin Uruguay", "UZT": "Tsayayyen Lokacin Uzbekistan", "UZT DST": "Lokacin Bazara na Uzbekistan", "VET": "Lokacin Venezuela", "VLAST": "Lokacin Bazara na Vladivostok", "VLAT": "Tsayayyen Lokacin Vladivostok", "VOLST": "Lokacin Bazara na Volgograd", "VOLT": "Tsayayyen Lokacin Volgograd", "VOST": "Lokacin Vostok", "VUT": "Tsayayyen Lokacin Vanuatu", "VUT DST": "Lokacin Bazara na Vanuatu", "WAKT": "Lokacin Wake Island", "WARST": "Lokacin Bazara na Yammacin Argentina", "WART": "Tsayayyen Lokacin Yammacin Argentina", "WAST": "Lokacin Afirka ta Yamma", "WAT": "Lokacin Afirka ta Yamma", "WESZ": "Ida lokaci ta yammacin turai da bazara", "WEZ": "Ida lokaci ta yammacin turai", "WFT": "Lokacin Wallis da Futuna", "WGST": "Lokacin Rana na Yammacin Greenland", "WGT": "Tsayayyen Lokacin Yammacin Greenland", "WIB": "Lokacin Yammacin Indonesia", "WIT": "Lokacin Gabashin Indonesia", "WITA": "Lokacin Indonesia ta Tsakiya", "YAKST": "Lokacin Bazara na Yakutsk", "YAKT": "Tsayayyen Lokacin Yakutsk", "YEKST": "Lokacin Bazara na Yekaterinburg", "YEKT": "Tsayayyen Lokacin Yekaterinburg", "YST": "Lokacin Yukon", "МСК": "Tsayayyen Lokacin Moscow", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Lokacin Yammacin Kazakhstan", "شىعىش قازاق ەلى": "Lokacin Gabashin Kazakhstan", "قازاق ەلى": "Lokacin Kazakhstan", "قىرعىزستان": "Lokacin Kyrgyzstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Lokacin Bazara na Peru"},
	}
}

// Locale returns the current translators string locale
func (ha *ha_NE) Locale() string {
	return ha.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ha_NE'
func (ha *ha_NE) PluralsCardinal() []locales.PluralRule {
	return ha.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ha_NE'
func (ha *ha_NE) PluralsOrdinal() []locales.PluralRule {
	return ha.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ha_NE'
func (ha *ha_NE) PluralsRange() []locales.PluralRule {
	return ha.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ha_NE'
func (ha *ha_NE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ha_NE'
func (ha *ha_NE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ha_NE'
func (ha *ha_NE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ha *ha_NE) MonthAbbreviated(month time.Month) string {
	return ha.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ha *ha_NE) MonthsAbbreviated() []string {
	return ha.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ha *ha_NE) MonthNarrow(month time.Month) string {
	return ha.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ha *ha_NE) MonthsNarrow() []string {
	return ha.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ha *ha_NE) MonthWide(month time.Month) string {
	return ha.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ha *ha_NE) MonthsWide() []string {
	return ha.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ha *ha_NE) WeekdayAbbreviated(weekday time.Weekday) string {
	return ha.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ha *ha_NE) WeekdaysAbbreviated() []string {
	return ha.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ha *ha_NE) WeekdayNarrow(weekday time.Weekday) string {
	return ha.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ha *ha_NE) WeekdaysNarrow() []string {
	return ha.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ha *ha_NE) WeekdayShort(weekday time.Weekday) string {
	return ha.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ha *ha_NE) WeekdaysShort() []string {
	return ha.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ha *ha_NE) WeekdayWide(weekday time.Weekday) string {
	return ha.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ha *ha_NE) WeekdaysWide() []string {
	return ha.daysWide
}

// Decimal returns the decimal point of number
func (ha *ha_NE) Decimal() string {
	return ha.decimal
}

// Group returns the group of number
func (ha *ha_NE) Group() string {
	return ha.group
}

// Group returns the minus sign of number
func (ha *ha_NE) Minus() string {
	return ha.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ha_NE' and handles both Whole and Real numbers based on 'v'
func (ha *ha_NE) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ha.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ha_NE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ha *ha_NE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ha.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ha_NE'
func (ha *ha_NE) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ha.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(ha.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ha.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ha.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ha_NE'
// in accounting notation.
func (ha *ha_NE) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ha.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ha.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ha.currencyNegativePrefix[j])
		}

		b = append(b, ha.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ha.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ha.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, ha.currencyNegativeSuffix...)
	} else {
		b = append(b, ha.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ha.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ha_NE'
func (ha *ha_NE) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ha.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
