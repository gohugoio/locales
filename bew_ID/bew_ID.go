package bew_ID

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type bew_ID struct {
	locale            string
	pluralsCardinal   []locales.PluralRule
	pluralsOrdinal    []locales.PluralRule
	pluralsRange      []locales.PluralRule
	decimal           string
	group             string
	minus             string
	percent           string
	timeSeparator     string
	currencies        []string // idx = enum of currency code
	monthsAbbreviated []string
	monthsNarrow      []string
	monthsWide        []string
	daysAbbreviated   []string
	daysNarrow        []string
	daysShort         []string
	daysWide          []string
	timezones         map[string]string
}

// New returns a new instance of translator for the 'bew_ID' locale
func New() locales.Translator {
	return &bew_ID{
		locale:            "bew_ID",
		pluralsCardinal:   nil,
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           ",",
		group:             ".",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ".",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "Jan", "Pèb", "Mar", "Apr", "Méi", "Jun", "Jul", "Ags", "Sèp", "Okt", "Nop", "Dés"},
		monthsNarrow:      []string{"", "J", "P", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:        []string{"", "Januari", "Pèbruari", "Maret", "April", "Méi", "Juni", "Juli", "Agustus", "Sèptèmber", "Oktober", "Nopèmber", "Désèmber"},
		daysAbbreviated:   []string{"Min", "Sen", "Sel", "Reb", "Kem", "Jum", "Sap"},
		daysNarrow:        []string{"M", "S", "S", "R", "K", "J", "S"},
		daysWide:          []string{"Minggu", "Senèn", "Selasa", "Rebo", "Kemis", "Juma’at", "Saptu"},
		timezones:         map[string]string{"ACDT": "Waktu Musim Pentèr Ostrali Tena", "ACST": "Waktu Pakem Ostrali Tenga", "ACT": "Waktu Pakem Akre", "ACWDT": "Waktu Musim Pentèr Ostrali Kulon-tenga", "ACWST": "Waktu Pakem Ostrali Kulon-tenga", "ADT": "Waktu Musim Pentèr Atlantik", "ADT Arabia": "Waktu Musim Pentèr Arab", "AEDT": "Waktu Musim Pentèr Ostrali Wètan", "AEST": "Waktu Pakem Ostrali Wètan", "AFT": "Waktu Apganistan", "AKDT": "Waktu Musim Pentèr Alaska", "AKST": "Waktu Pakem Alaska", "AMST": "Waktu Musim Pentèr Amason", "AMST Armenia": "Waktu Musim Pentèr Lemènder", "AMT": "Waktu Pakem Amason", "AMT Armenia": "Waktu Pakem Lemènder", "ANAST": "Waktu Musim Pentèr Anadir", "ANAT": "Waktu Pakem Anadir", "ARST": "Waktu Musim Pentèr Argèntina", "ART": "Waktu Pakem Argèntina", "AST": "Waktu Pakem Atlantik", "AST Arabia": "Waktu Pakem Arab", "AWDT": "Waktu Musim Pentèr Ostrali Kulon", "AWST": "Waktu Pakem Ostrali Kulon", "AZST": "Waktu Musim Pentèr Asèrbaijan", "AZT": "Waktu Pakem Asèrbaijan", "BDT Bangladesh": "Waktu Musim Pentèr Benggaladésa", "BNT": "Waktu Bruné Darussalam", "BOT": "Waktu Boliwi", "BRST": "Waktu Musim Pentèr Brasilia", "BRT": "Waktu Pakem Brasilia", "BST Bangladesh": "Waktu Pakem Benggaladésa", "BT": "Waktu Butan", "CAST": "Waktu Casey", "CAT": "Waktu Aprika Tenga", "CCT": "Waktu Pulo Kokos", "CDT": "Waktu Musim Pentèr Tenga", "CHADT": "Waktu Musim Pentèr Catham", "CHAST": "Waktu Pakem Catham", "CHUT": "Waktu Cuk", "CKT": "Waktu Pakem Pulo Cook", "CKT DST": "Waktu Musim Pentèr Pulo Cook", "CLST": "Waktu Musim Pentèr Cili", "CLT": "Waktu Pakem Cili", "COST": "Waktu Musim Pentèr Kolombia", "COT": "Waktu Pakem Kolombia", "CST": "Waktu Pakem Tenga", "CST China": "Waktu Pakem Tiongkok", "CST China DST": "Waktu Musim Pentèr Tiongkok", "CVST": "Waktu Musim Pentèr Tanjung Ijo", "CVT": "Waktu Pakem Tanjung Ijo", "CXT": "Waktu Pulo Natal", "ChST": "Waktu Pakem Camoro", "ChST NMI": "ChST NMI", "CuDT": "Waktu Musim Pentèr Kuba", "CuST": "Waktu Pakem Kuba", "DAVT": "Waktu Davis", "DDUT": "Waktu Dumont-d’Urville", "EASST": "Waktu Musim Pentèr Pulo Paskah", "EAST": "Waktu Pakem Pulo Paskah", "EAT": "Waktu Aprika Wètan", "ECT": "Waktu Èkuador", "EDT": "Waktu Musim Pentèr Wètan", "EGDT": "Waktu Musim Pentèr Grunlan Wètan", "EGST": "Waktu Pakem Grunlan Wètan", "EST": "Waktu Pakem Wètan", "FEET": "Waktu Èropa Wètan-jau", "FJT": "Waktu Pakem Piji", "FJT Summer": "Waktu Musim Pentèr Piji", "FKST": "Waktu Musim Pentèr Pulo Paklan", "FKT": "Waktu Pakem Pulo Paklan", "FNST": "Waktu Musim Pentèr Fernando de Noronha", "FNT": "Waktu Pakem Fernando de Noronha", "GALT": "Waktu Galapagos", "GAMT": "Waktu Gambier", "GEST": "Waktu Musim Pentèr Géorgi", "GET": "Waktu Pakem Géorgi", "GFT": "Waktu Guyana Prasman", "GIT": "Waktu Pulo Gilbet", "GMT": "Waktu Rerata Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Waktu Géorgi Kidul", "GST Guam": "GST Guam", "GYT": "Waktu Guyana", "HADT": "Waktu Pakem Hawai-Aléut", "HAST": "Waktu Pakem Hawai-Aléut", "HKST": "Waktu Musim Pentèr Hongkong", "HKT": "Waktu Pakem Hongkong", "HOVST": "Waktu Musim Pentèr Houd", "HOVT": "Waktu Pakem Houd", "ICT": "Waktu Indocina", "IDT": "Waktu Musim Pentèr Israèl", "IOT": "Waktu Laotan Hindi", "IRKST": "Waktu Musim Pentèr Irkut", "IRKT": "Waktu Pakem Irkut", "IRST": "Waktu Pakem Iran", "IRST DST": "Waktu Musim Pentèr Iran", "IST": "Waktu Hindi Pakem", "IST Israel": "Waktu Pakem Israèl", "JDT": "Waktu Musim Pentèr Jepang", "JST": "Waktu Pakem Jepang", "KOST": "Waktu Kosaé", "KRAST": "Waktu Musim Pentèr Krasnoyar", "KRAT": "Waktu Pakem Krasnoyar", "KST": "Waktu Pakem Koréa", "KST DST": "Waktu Musim Pentèr Koréa", "LHDT": "Waktu Musim Pentèr Lord Howe", "LHST": "Waktu Pakem Lord Howe", "LINT": "Waktu Pulo Line", "MAGST": "Waktu Musim Pentèr Magadan", "MAGT": "Waktu Pakem Magadan", "MART": "Waktu Markésas", "MAWT": "Waktu Mawson", "MDT": "Waktu Musim Pentèr Gunung", "MESZ": "Waktu Musim Pentèr Èropa Tenga", "MEZ": "Waktu Pakem Èropa Tenga", "MHT": "Waktu Pulo Marsal", "MMT": "Waktu Mianmar", "MSD": "Waktu Musim Pentèr Mosko", "MST": "Waktu Pakem Gunung", "MUST": "Waktu Musim Pentèr Moritius", "MUT": "Waktu Pakem Moritius", "MVT": "Waktu Maladéwa", "MYT": "Waktu Malésia", "NCT": "Waktu Pakem Kalédoni Baru", "NDT": "Waktu Musim Pentèr Niuponlan", "NDT New Caledonia": "Waktu Musim Pentèr Kalédoni Baru", "NFDT": "Waktu Musim Pentèr Pulo Norpok", "NFT": "Waktu Pakem Pulo Norpok", "NOVST": "Waktu Musim Pentèr Nowosibir", "NOVT": "Waktu Pakem Nowosibir", "NPT": "Waktu Népal", "NRT": "Waktu Nauru", "NST": "Waktu Pakem Niuponlan", "NUT": "Waktu Niué", "NZDT": "Waktu Musim Pentèr Sélan Baru", "NZST": "Waktu Pakem Sélan Baru", "OESZ": "Waktu Musim Pentèr Èropa Wètan", "OEZ": "Waktu Pakem Èropa Wètan", "OMSST": "Waktu Musim Pentèr Om", "OMST": "Waktu Pakem Om", "PDT": "Waktu Musim Pentèr Teduh", "PDTM": "Waktu Musim Pentèr Mèksiko Teduh", "PETDT": "Waktu Musim Pentèr Pètropaulop-Kamcatka", "PETST": "Waktu Pakem Pètropaulop-Kamcatka", "PGT": "Waktu Papua Giné Baru", "PHOT": "Waktu Pulo Pènik", "PKT": "Waktu Pakem Pakistan", "PKT DST": "Waktu Musim Pentèr Pakistan", "PMDT": "Waktu Musim Pentèr Sint-Pièr èn Mikélon", "PMST": "Waktu Pakem Sint-Pièr èn Mikélon", "PONT": "Waktu Ponapé", "PST": "Waktu Pakem Teduh", "PST Philippine": "Waktu Pakem Pilipénen", "PST Philippine DST": "Waktu Musim Pentèr Pilipénen", "PST Pitcairn": "Waktu Pitkèren", "PSTM": "Waktu Pakem Mèksiko Teduh", "PWT": "Waktu Palau", "PYST": "Waktu Musim Pentèr Paragué", "PYT": "Waktu Pakem Paragué", "PYT Korea": "Waktu Piongyang", "RET": "Waktu Réunion", "ROTT": "Waktu Rothera", "SAKST": "Waktu Musim Pentèr Sahalin", "SAKT": "Waktu Pakem Sahalin", "SAMST": "Waktu Musim Pentèr Samara", "SAMT": "Waktu Pakem Samara", "SAST": "Waktu Pakem Aprika Kidul", "SBT": "Waktu Pulo Suléman", "SCT": "Waktu Sésèl", "SGT": "Waktu Pakem Singapur", "SLST": "Waktu Sri Langka", "SRT": "Waktu Suriname", "SST Samoa": "Waktu Pakem Samoa", "SST Samoa Apia": "Waktu Pakem Apia", "SST Samoa Apia DST": "Waktu Musim Pentèr Apia", "SST Samoa DST": "Waktu Musim Pentèr Samoa", "SYOT": "Waktu Syowa", "TAAF": "Waktu Wilayah Kulon èn Kutub Kidul Prasman", "TAHT": "Waktu Tahiti", "TJT": "Waktu Tajikistan", "TKT": "Waktu Tokelau", "TLT": "Waktu Timor Wètan", "TMST": "Waktu Musim Pentèr Turkmènistan", "TMT": "Waktu Pakem Turkmènistan", "TOST": "Waktu Musim Pentèr Tonga", "TOT": "Waktu Pakem Tonga", "TVT": "Waktu Tuwalu", "TWT": "Waktu Pakem Taipé", "TWT DST": "Waktu Musim Pentèr Taipé", "ULAST": "Waktu Musim Pentèr Ulanbator", "ULAT": "Waktu Pakem Ulanbator", "UYST": "Waktu Musim Pentèr Urugué", "UYT": "Waktu Pakem Urugué", "UZT": "Waktu Pakem Usbèkistan", "UZT DST": "Waktu Musim Pentèr Usbèkistan", "VET": "Waktu Bénésuèla", "VLAST": "Waktu Musim Pentèr Weladiwostok", "VLAT": "Waktu Pakem Weladiwostok", "VOLST": "Waktu Musim Pentèr Wolgograd", "VOLT": "Waktu Pake Wolgograd", "VOST": "Waktu Wostok", "VUT": "Waktu Pakem Wanuatu", "VUT DST": "Waktu Musim Pentèr Wanuatu", "WAKT": "Waktu Pulo Wék", "WARST": "Waktu Musim Pentèr Argèntina Kulon", "WART": "Waktu Pakem Argèntina Kulon", "WAST": "Waktu Aprika Kulon", "WAT": "Waktu Aprika Kulon", "WESZ": "Waktu Musim Pentèr Èropa Kulon", "WEZ": "Waktu Pakem Èropa Kulon", "WFT": "Waktu Walis èn Putuna", "WGST": "Waktu Musim Pentèr Grunlan Kulon", "WGT": "Waktu Pakem Grunlan Kulon", "WIB": "Waktu Indonésia Kulon", "WIT": "Waktu Indonésia Wètan", "WITA": "Waktu Indonésia Tenga", "YAKST": "Waktu Musim Pentèr Yakut", "YAKT": "Waktu Pakem Yakut", "YEKST": "Waktu Musim Pentèr Yékatèrinenbereh", "YEKT": "Waktu Pakem Yékatèrinenbereh", "YST": "Waktu Yukon", "МСК": "Waktu Pakem Mosko", "اقتاۋ": "Waktu Pakem Akto", "اقتاۋ قالاسى": "Waktu Musim Pentèr Akto", "اقتوبە": "Waktu Pakem Aktobé", "اقتوبە قالاسى": "Waktu Musim Pentèr Aktobé", "الماتى": "Waktu Pakem Almati", "الماتى قالاسى": "Waktu Musim Pentèr Almati", "باتىس قازاق ەلى": "Waktu Kasakstan Kulon", "شىعىش قازاق ەلى": "Waktu Kasakstan Wètan", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Waktu Kirgistan", "قىزىلوردا": "Waktu Pakem Kisilorda", "قىزىلوردا قالاسى": "Waktu Musim Pentèr Kisilorda", "∅∅∅": "Waktu Musim Pentèr Péru"},
	}
}

// Locale returns the current translators string locale
func (bew *bew_ID) Locale() string {
	return bew.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bew_ID'
func (bew *bew_ID) PluralsCardinal() []locales.PluralRule {
	return bew.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bew_ID'
func (bew *bew_ID) PluralsOrdinal() []locales.PluralRule {
	return bew.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bew_ID'
func (bew *bew_ID) PluralsRange() []locales.PluralRule {
	return bew.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bew_ID'
func (bew *bew_ID) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bew_ID'
func (bew *bew_ID) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bew_ID'
func (bew *bew_ID) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bew *bew_ID) MonthAbbreviated(month time.Month) string {
	if len(bew.monthsAbbreviated) == 0 {
		return ""
	}
	return bew.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bew *bew_ID) MonthsAbbreviated() []string {
	return bew.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bew *bew_ID) MonthNarrow(month time.Month) string {
	if len(bew.monthsNarrow) == 0 {
		return ""
	}
	return bew.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bew *bew_ID) MonthsNarrow() []string {
	return bew.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bew *bew_ID) MonthWide(month time.Month) string {
	if len(bew.monthsWide) == 0 {
		return ""
	}
	return bew.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bew *bew_ID) MonthsWide() []string {
	return bew.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bew *bew_ID) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(bew.daysAbbreviated) == 0 {
		return ""
	}
	return bew.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bew *bew_ID) WeekdaysAbbreviated() []string {
	return bew.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bew *bew_ID) WeekdayNarrow(weekday time.Weekday) string {
	if len(bew.daysNarrow) == 0 {
		return ""
	}
	return bew.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bew *bew_ID) WeekdaysNarrow() []string {
	return bew.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bew *bew_ID) WeekdayShort(weekday time.Weekday) string {
	if len(bew.daysShort) == 0 {
		return ""
	}
	return bew.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bew *bew_ID) WeekdaysShort() []string {
	return bew.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bew *bew_ID) WeekdayWide(weekday time.Weekday) string {
	if len(bew.daysWide) == 0 {
		return ""
	}
	return bew.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bew *bew_ID) WeekdaysWide() []string {
	return bew.daysWide
}

// Decimal returns the decimal point of number
func (bew *bew_ID) Decimal() string {
	return bew.decimal
}

// Group returns the group of number
func (bew *bew_ID) Group() string {
	return bew.group
}

// Group returns the minus sign of number
func (bew *bew_ID) Minus() string {
	return bew.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bew_ID' and handles both Whole and Real numbers based on 'v'
func (bew *bew_ID) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bew.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bew.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bew.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bew_ID' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bew *bew_ID) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bew.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bew.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bew.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bew_ID'
func (bew *bew_ID) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bew.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bew.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bew.group[0])
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
		b = append(b, bew.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bew.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bew_ID'
// in accounting notation.
func (bew *bew_ID) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bew.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bew.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bew.group[0])
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

		b = append(b, bew.minus[0])

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
			b = append(b, bew.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtDateShort(t time.Time) string {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bew.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bew.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bew.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bew.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bew_ID'
func (bew *bew_ID) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bew.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
