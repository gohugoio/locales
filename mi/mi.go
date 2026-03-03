package mi

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mi struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	timeSeparator      string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'mi' locale
func New() locales.Translator {
	return &mi{
		locale:            "mi",
		pluralsCardinal:   nil,
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           ".",
		group:             ",",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "Hān", "Pēp", "Māe", "Āpe", "Mei", "Hun", "Hūr", "Āku", "Hep", "Oke", "Noe", "Tīh"},
		monthsNarrow:      []string{"", "H", "P", "M", "Ā", "M", "H", "H", "Ā", "H", "O", "N", "T"},
		monthsWide:        []string{"", "Hānuere", "Pēpuere", "Māehe", "Āpereira", "Mei", "Hune", "Hūrae", "Ākuhata", "Hepetema", "Oketopa", "Noema", "Tīhema"},
		daysAbbreviated:   []string{"Rāt", "Man", "Tūr", "Wen", "Tāi", "Par", "Rāh"},
		daysNarrow:        []string{"Rt", "M", "T", "W", "T", "P", "Rh"},
		daysShort:         []string{"Rāt", "Man", "Tū", "Wen", "Tāi", "Par", "Rāh"},
		daysWide:          []string{"Rātapu", "Mane", "Tūrei", "Wenerei", "Tāite", "Paraire", "Rāhoroi"},
		timezones:         map[string]string{"ACDT": "Wā Ahitereiria Waenga Awatea", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Wā Ahitereiria Waenga-Uru Awatea", "ACWST": "Wā Ahitereiria Waenga-Uru Arowhānui", "ADT": "Wā Awatea Ranatiki", "ADT Arabia": "Wā Arāpia Awatea", "AEDT": "Wā Ahitereiria ki te Rāwhiti Awatea", "AEST": "Wā Ahitereiria ki te Rāwhiti Arowhānui", "AFT": "Wā Awhekenetāna", "AKDT": "Wā Awatea Alaska", "AKST": "Wā Arowhānui Alaska", "AMST": "Wā Amahona Raumati", "AMST Armenia": "Wā Āmenia Raumati", "AMT": "Wā Amahona Arowhānui", "AMT Armenia": "Wā Āmenia Arowhānui", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Wā Āketina Raumati", "ART": "Wā Āketina Arowhānui", "AST": "Wā Arowhānui Ranatiki", "AST Arabia": "Wā Arāpia Arowhānui", "AWDT": "Wā Ahitereiria ki te Uru Awatea", "AWST": "Wā Ahitereiria ki te Uru Arowhānui", "AZST": "Wā Atepaihānia Raumati", "AZT": "Wā Atepaihānia Arowhānui", "BDT Bangladesh": "Wā Pākaratēhi Raumati", "BNT": "Wā Poronai Darussalam", "BOT": "Wā Poriwia", "BRST": "Wā Parīhia Raumati", "BRT": "Wā Parīhia Arowhānui", "BST Bangladesh": "Wā Pākaratēhi Arowhānui", "BT": "Wā Pūtana", "CAST": "CAST", "CAT": "Wā o Te Puku o Āwherika", "CCT": "Wā o Ngā Moutere Kokohi", "CDT": "Wā Awatea Waenga", "CHADT": "Wā Rēkohu Awatea", "CHAST": "Wā Rēkohu Arowhānui", "CHUT": "Wā Chuuk", "CKT": "Wā Kuki Airani Arowhānui", "CKT DST": "Wā Kuki Airani Raumati Haurua", "CLST": "Wā Hiri Raumati", "CLT": "Wā Hiri Arowhānui", "COST": "Wā Koromōpia Raumati", "COT": "Wā Koromōpia Arowhānui", "CST": "Wā Arowhānui Waenga", "CST China": "Wā Haina Arowhānui", "CST China DST": "Wā Haina Awatea", "CVST": "Wā Raumati o Te Kūrae Matomato", "CVT": "Wā Arowhānui o Te Kūrae Matomato", "CXT": "Wā o Te Moutere Kirihimete", "ChST": "Wā Chamorro Arowhānui", "ChST NMI": "ChST NMI", "CuDT": "Wā Awatea Kiupa", "CuST": "Wā Arowhānui Kiupa", "DAVT": "Wā Rēweti", "DDUT": "Wā Dumont-d’Urville", "EASST": "Wā ki te Moutere Aranga Raumati", "EAST": "Wā ki te Moutere Aranga Arowhānui", "EAT": "Wā o Āwherika ki te rāwhiti", "ECT": "Wā Ekuatoa", "EDT": "Wā Awatea Rāwhiti", "EGDT": "Wā Raumati o Whenuakāriki ki te rāwhiti", "EGST": "Wā Arowhānui o Whenuakāriki ki te rāwhiti", "EST": "Wā Arowhānui Rāwhiti", "FEET": "Wā Ūropi ki te rāwhiti rawa", "FJT": "Wā Whītī Arowhānui", "FJT Summer": "Wā Whītī Raumati", "FKST": "Wā ki Ngā Motu Whākana Raumati", "FKT": "Wā ki Ngā Motu Whākana Arowhānui", "FNST": "Wā Fernando de Noronha Raumati", "FNT": "Wā Fernando de Noronha Arowhānui", "GALT": "Wā Galapagos", "GAMT": "Wā Gambier", "GEST": "Wā Hōria Raumati", "GET": "Wā Hōria Arowhānui", "GFT": "Wā Kiāna Wīwī", "GIT": "Wā Kiripati", "GMT": "Wā Toharite Kiriwīti", "GNSST": "GNSST", "GNST": "GNST", "GST": "Wā Hōria ki te Tonga", "GST Guam": "GST Guam", "GYT": "Wā Kaiana", "HADT": "Wā Arowhānui Hawaii-Aleutian", "HAST": "Wā Arowhānui Hawaii-Aleutian", "HKST": "Wā Hongipua Raumati", "HKT": "Wā Hongipua Arowhānui", "HOVST": "Wā Hovd Raumati", "HOVT": "Wā Hovd Arowhānui", "ICT": "Wā Īniahaina", "IDT": "Wā Iharaira Awatea", "IOT": "Wā o Te Moana Īnia", "IRKST": "Wā Irkutsk Raumati", "IRKT": "Wā Irkutsk Arowhānui", "IRST": "Wā Irāna Arowhānui", "IRST DST": "Wā Irāna Awatea", "IST": "Wā Īnia", "IST Israel": "Wā Iharaira Arowhānui", "JDT": "Wā Hapani Awatea", "JST": "Wā Hapani Arowhānui", "KOST": "Wā Kosrae", "KRAST": "Wā Krasnoyarsk Raumati", "KRAT": "Wā Krasnoyarsk Arowhānui", "KST": "Wā Kōrea Arowhānui", "KST DST": "Wā Kōrea Awatea", "LHDT": "Wā Lord Howe Awatea", "LHST": "Wā Lord Howe Arowhānui", "LINT": "Wā o Ngā Mouter o Te Raina", "MAGST": "Wā Magadan Raumati", "MAGT": "Wā Magadan Arowhānui", "MART": "Wā Marquesas", "MAWT": "Wā Mawson", "MDT": "Wā Awatea Maunga", "MESZ": "Wā Raumati Uropi Waenga", "MEZ": "Wā Arowhānui Uropi Waenga", "MHT": "Wā o Ngā Motu Māhara", "MMT": "Wā Pēma", "MSD": "Wā Mohikau Raumati", "MST": "Wā Arowhānui Maunga", "MUST": "Wā Marihi Raumati", "MUT": "Wā Marihi Arowhānui", "MVT": "Wā Māratiri", "MYT": "Wā Mareia", "NCT": "Wā Whenua Kanaki Arowhānui", "NDT": "Wā Awatea Newfoundland", "NDT New Caledonia": "Wā Whenua Kanaki Raumati", "NFDT": "Wā o Te Moutere Nōpoke Awatea", "NFT": "Wā o Te Moutere Nōpoke Arowhānui", "NOVST": "Wā Novosibirsk Raumati", "NOVT": "Wā Novosibirsk Arowhānui", "NPT": "Wā Nepōra", "NRT": "Wā Nauru", "NST": "Wā Arowhānui Newfoundland", "NUT": "Wā Niue", "NZDT": "Wā Aotearoa Awatea", "NZST": "Wā Aotearoa Arowhānui", "OESZ": "Wā Raumati Uropi Rāwhiti", "OEZ": "Wā Arowhānui Uropi Rāwhiti", "OMSST": "Wā Omsk Raumati", "OMST": "Wā Omsk Arowhānui", "PDT": "Wā Awatea Kiwa", "PDTM": "Wā Awatea Mēhiko Kiwa", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Wā Papua Nūkini", "PHOT": "Wā o Ngā Moutere Phoenix", "PKT": "Wā Pakitāne Arowhānui", "PKT DST": "Wā Pakitāne Raumati", "PMDT": "Wā Awatea o St. Pierre me Miquelon", "PMST": "Wā Arowhānui o St. Pierre me Miquelon", "PONT": "Wā Ponape", "PST": "Wā Arowhānui Kiwa", "PST Philippine": "Wā Piripīni Arowhānui", "PST Philippine DST": "Wā Piripīni Raumati", "PST Pitcairn": "Wā Pitcairn", "PSTM": "Wā Arowhānui Mēhiko Kiwa", "PWT": "Wā Pārau", "PYST": "Wā Parakai Raumati", "PYT": "Wā Parakai Arowhānui", "PYT Korea": "Wā Pyongyang", "RET": "Wā Reunion", "ROTT": "Wā Rothera", "SAKST": "Wā Sakhalin Raumati", "SAKT": "Wā Sakhalin Arowhānui", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Wā Arowhānui o Āwherika ki te tonga", "SBT": "Wā o Ngā Motu Horomona", "SCT": "Wā Heikere", "SGT": "Wā Hingapoa Arowhānui", "SLST": "SLST", "SRT": "Wā Huriname", "SST Samoa": "Wā Hāmoa Arowhānui", "SST Samoa Apia": "Wā Āpia Arowhānui", "SST Samoa Apia DST": "Wā Āpia Awatea", "SST Samoa DST": "Wā Hāmoa Awatea", "SYOT": "Wā Syowa", "TAAF": "Wā Wīwī o Te Tonga me te Kōpakatanga ki te Tonga", "TAHT": "Wā Tahiti", "TJT": "Wā Takiritānga", "TKT": "Wā Tokerau", "TLT": "Wā o Timoa ki te Rāwhiti", "TMST": "Wā Tukumanatānga Raumati", "TMT": "Wā Tukumanatānga Arowhānui", "TOST": "Wā Tonga Raumati", "TOT": "Wā Tonga Arowhānui", "TVT": "Wā Tūwaru", "TWT": "Wā Taipei Arowhānui", "TWT DST": "Wā Taipei Awatea", "ULAST": "Wā Ulaanbaatar Raumati", "ULAT": "Wā Ulaanbaatar Arowhānui", "UYST": "Wā Urukoi Raumati", "UYT": "Wā Urukoi Arowhānui", "UZT": "Wā Uhipeketāne Arowhānui", "UZT DST": "Wā Uhipeketāne Raumati", "VET": "Wā Penehūera", "VLAST": "Wā Vladivostok Raumati", "VLAT": "Wā Vladivostok Arowhānui", "VOLST": "Wā Volgograd Raumati", "VOLT": "Wā Volgograd Arowhānui", "VOST": "Wā Vostok", "VUT": "Wā Whenuatū Arowhānui", "VUT DST": "Wā Whenuatū Raumati", "WAKT": "Wā o Te Motu Wake", "WARST": "Wā Āketina ki te uru Raumati", "WART": "Wā Āketina ki te uru Arowhānui", "WAST": "Wā o Āwherika ki te uru", "WAT": "Wā o Āwherika ki te uru", "WESZ": "Wā Raumati Uropi Uru", "WEZ": "Wā Arowhānui Uropi Uru", "WFT": "Wā Wārihi me Whutuna", "WGST": "Wā Raumati o Whenuakāriki ki te uru", "WGT": "Wā Arowhānui o Whenuakāriki ki te uru", "WIB": "Wā Initonīhia ki te uru", "WIT": "Wā Initonīhia ki te rāwhiti", "WITA": "Wā Initonīhia Waenga", "YAKST": "Wā Yakutsk Raumati", "YAKT": "Wā Yakutsk Arowhānui", "YEKST": "Wā Yekaterinburg Raumati", "YEKT": "Wā Yekaterinburg Arowhānui", "YST": "Wā Yukon", "МСК": "Wā Mohikau Arowhānui", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Wā Katatānga ki te Uru", "شىعىش قازاق ەلى": "Wā Katatānga ki te Rāwhiti", "قازاق ەلى": "Wā Katatānga", "قىرعىزستان": "Wā Kikitānga", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Wā Azores Raumati"},
	}
}

// Locale returns the current translators string locale
func (mi *mi) Locale() string {
	return mi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mi'
func (mi *mi) PluralsCardinal() []locales.PluralRule {
	return mi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mi'
func (mi *mi) PluralsOrdinal() []locales.PluralRule {
	return mi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mi'
func (mi *mi) PluralsRange() []locales.PluralRule {
	return mi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mi'
func (mi *mi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mi'
func (mi *mi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mi'
func (mi *mi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mi *mi) MonthAbbreviated(month time.Month) string {
	if len(mi.monthsAbbreviated) == 0 {
		return ""
	}
	return mi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mi *mi) MonthsAbbreviated() []string {
	return mi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mi *mi) MonthNarrow(month time.Month) string {
	if len(mi.monthsNarrow) == 0 {
		return ""
	}
	return mi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mi *mi) MonthsNarrow() []string {
	return mi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mi *mi) MonthWide(month time.Month) string {
	if len(mi.monthsWide) == 0 {
		return ""
	}
	return mi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mi *mi) MonthsWide() []string {
	return mi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mi *mi) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(mi.daysAbbreviated) == 0 {
		return ""
	}
	return mi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mi *mi) WeekdaysAbbreviated() []string {
	return mi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mi *mi) WeekdayNarrow(weekday time.Weekday) string {
	if len(mi.daysNarrow) == 0 {
		return ""
	}
	return mi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mi *mi) WeekdaysNarrow() []string {
	return mi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mi *mi) WeekdayShort(weekday time.Weekday) string {
	if len(mi.daysShort) == 0 {
		return ""
	}
	return mi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mi *mi) WeekdaysShort() []string {
	return mi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mi *mi) WeekdayWide(weekday time.Weekday) string {
	if len(mi.daysWide) == 0 {
		return ""
	}
	return mi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mi *mi) WeekdaysWide() []string {
	return mi.daysWide
}

// Decimal returns the decimal point of number
func (mi *mi) Decimal() string {
	return mi.decimal
}

// Group returns the group of number
func (mi *mi) Group() string {
	return mi.group
}

// Group returns the minus sign of number
func (mi *mi) Minus() string {
	return mi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mi' and handles both Whole and Real numbers based on 'v'
func (mi *mi) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mi' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mi *mi) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mi'
func (mi *mi) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mi.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mi'
// in accounting notation.
func (mi *mi) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mi.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'mi'
func (mi *mi) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'mi'
func (mi *mi) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mi'
func (mi *mi) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mi'
func (mi *mi) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mi'
func (mi *mi) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mi'
func (mi *mi) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mi'
func (mi *mi) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mi'
func (mi *mi) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
