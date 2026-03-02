package fil

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type fil struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'fil' locale
func New() locales.Translator {
	return &fil{
		locale:                 "fil",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ene", "Peb", "Mar", "Abr", "May", "Hun", "Hul", "Ago", "Set", "Okt", "Nob", "Dis"},
		monthsNarrow:           []string{"", "Ene", "Peb", "Mar", "Abr", "May", "Hun", "Hul", "Ago", "Set", "Okt", "Nob", "Dis"},
		monthsWide:             []string{"", "Enero", "Pebrero", "Marso", "Abril", "Mayo", "Hunyo", "Hulyo", "Agosto", "Setyembre", "Oktubre", "Nobyembre", "Disyembre"},
		daysAbbreviated:        []string{"Lin", "Lun", "Mar", "Miy", "Huw", "Biy", "Sab"},
		daysNarrow:             []string{"Lin", "Lun", "Mar", "Miy", "Huw", "Biy", "Sab"},
		daysWide:               []string{"Linggo", "Lunes", "Martes", "Miyerkules", "Huwebes", "Biyernes", "Sabado"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"am", "pm"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Before Christ", "Anno Domini"},
		timezones:              map[string]string{"ACDT": "Daylight Time sa Gitnang Australya", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Daylight Time sa Gitnang Kanlurang Australia", "ACWST": "Standard Time ng Gitnang Kanluran ng Australia", "ADT": "Daylight na Oras sa Atlantiko", "ADT Arabia": "Daylight Time sa Arabia", "AEDT": "Daylight Time sa Silangang Australia", "AEST": "Standard na Oras sa Silangang Australia", "AFT": "Oras sa Afghanistan", "AKDT": "Daylight Time sa Alaska", "AKST": "Standard na Oras sa Alaska", "AMST": "Oras sa Tag-init ng Amazon", "AMST Armenia": "Oras sa Tag-init ng Armenia", "AMT": "Standard na Oras sa Amazon", "AMT Armenia": "Standard na Oras sa Armenia", "ANAST": "Summer Time sa Anadyr", "ANAT": "Standard Time sa Anadyr", "ARST": "Oras sa Tag-init ng Argentina", "ART": "Standard na Oras sa Argentina", "AST": "Standard na Oras sa Atlantiko", "AST Arabia": "Standard na Oras sa Arabia", "AWDT": "Daylight Time sa Kanlurang Australia", "AWST": "Standard na Oras sa Kanlurang Australia", "AZST": "Oras sa Tag-init ng Azerbaijan", "AZT": "Standard na Oras sa Azerbaijan", "BDT Bangladesh": "Oras sa Tag-init ng Bangladesh", "BNT": "Oras sa Brunei Darussalam", "BOT": "Oras sa Bolivia", "BRST": "Oras sa Tag-init ng Brasilia", "BRT": "Standard na Oras sa Brasilia", "BST Bangladesh": "Standard na Oras sa Bangladesh", "BT": "Oras sa Bhutan", "CAST": "CAST", "CAT": "Oras sa Gitnang Africa", "CCT": "Oras sa Cocos Islands", "CDT": "Sentral na Daylight na Oras sa North America", "CHADT": "Daylight Time sa Chatham", "CHAST": "Standard na Oras sa Chatham", "CHUT": "Oras sa Chuuk", "CKT": "Standard na Oras sa Cook Islands", "CKT DST": "Oras sa Kalahati ng Tag-init ng Cook Islands", "CLST": "Oras sa Tag-init ng Chile", "CLT": "Standard na Oras sa Chile", "COST": "Oras sa Tag-init ng Colombia", "COT": "Standard na Oras sa Colombia", "CST": "Sentral na Standard na Oras sa North America", "CST China": "Standard na Oras sa China", "CST China DST": "Daylight Time sa China", "CVST": "Oras sa Tag-init ng Cape Verde", "CVT": "Standard na Oras sa Cape Verde", "CXT": "Oras sa Christmas Island", "ChST": "Standard na Oras sa Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Daylight na Oras sa Cuba", "CuST": "Standard na Oras sa Cuba", "DAVT": "Oras sa Davis", "DDUT": "Oras sa Dumont-d’Urville", "EASST": "Oras sa Tag-init ng Easter Island", "EAST": "Standard na Oras sa Easter Island", "EAT": "Oras sa Silangang Africa", "ECT": "Oras sa Ecuador", "EDT": "Daylight na Oras sa Silangan ng Hilagang Amerika", "EGDT": "Oras sa Tag-init ng Silangang Greenland", "EGST": "Standard na Oras sa Silangang Greenland", "EST": "Standard na Oras sa Silangan ng Hilangang Amerika", "FEET": "Oras sa Pinaka-silangang Europe", "FJT": "Standard na Oras sa Fiji", "FJT Summer": "Oras sa Tag-init ng Fiji", "FKST": "Oras sa Tag-init ng Falkland Islands", "FKT": "Standard na Oras sa Falkland Islands", "FNST": "Oras sa Tag-init ng Fernando de Noronha", "FNT": "Standard na Oras sa Fernando de Noronha", "GALT": "Oras sa Galapagos", "GAMT": "Oras sa Gambier", "GEST": "Oras sa Tag-init ng Georgia", "GET": "Standard na Oras sa Georgia", "GFT": "Oras sa French Guiana", "GIT": "Oras sa Gilbert Islands", "GMT": "Greenwich Mean Time", "GNSST": "GNSST", "GNST": "GNST", "GST": "Oras sa Gulf", "GST Guam": "GST Guam", "GYT": "Oras sa Guyana", "HADT": "Standard na Oras sa Hawaii-Aleutian", "HAST": "Standard na Oras sa Hawaii-Aleutian", "HKST": "Oras sa Tag-init ng Hong Kong", "HKT": "Standard na Oras sa Hong Kong", "HOVST": "Oras sa Tag-init ng Hovd", "HOVT": "Standard na Oras sa Hovd", "ICT": "Oras sa Indochina", "IDT": "Daylight Time sa Israel", "IOT": "Oras sa Indian Ocean", "IRKST": "Oras sa Tag-init ng Irkutsk", "IRKT": "Standard na Oras sa Irkutsk", "IRST": "Standard na Oras sa Iran", "IRST DST": "Daylight Time sa Iran", "IST": "Standard na Oras sa India", "IST Israel": "Standard na Oras sa Israel", "JDT": "Daylight Time sa Japan", "JST": "Standard na Oras sa Japan", "KOST": "Oras sa Kosrae", "KRAST": "Oras sa Tag-init ng Krasnoyarsk", "KRAT": "Standard na Oras sa Krasnoyarsk", "KST": "Standard na Oras sa Korea", "KST DST": "Daylight Time sa Korea", "LHDT": "Daylight Time sa Lorde Howe", "LHST": "Standard na Oras sa Lord Howe", "LINT": "Oras sa Line Islands", "MAGST": "Oras sa Tag-init ng Magadan", "MAGT": "Standard na Oras sa Magadan", "MART": "Oras sa Marquesas", "MAWT": "Oras sa Mawson", "MDT": "MDT", "MESZ": "Oras sa Tag-init ng Gitnang Europe", "MEZ": "Standard na Oras sa Gitnang Europe", "MHT": "Oras sa Marshall Islands", "MMT": "Oras sa Myanmar", "MSD": "Oras sa Tag-init ng Moscow", "MST": "MST", "MUST": "Oras sa Tag-init ng Mauritius", "MUT": "Standard na Oras sa Mauritius", "MVT": "Oras sa Maldives", "MYT": "Oras sa Malaysia", "NCT": "Standard na Oras sa New Caledonia", "NDT": "Daylight na Oras sa Newfoundland", "NDT New Caledonia": "Oras sa Tag-init ng New Caledonia", "NFDT": "Daylight Time sa Norfolk Island", "NFT": "Standard na Oras sa Norfolk Island", "NOVST": "Oras sa Tag-init ng Novosibirsk", "NOVT": "Standard na Oras sa Novosibirsk", "NPT": "Oras sa Nepal", "NRT": "Oras sa Nauru", "NST": "Standard na Oras sa Newfoundland", "NUT": "Oras sa Niue", "NZDT": "Daylight Time sa New Zealand", "NZST": "Standard na Oras sa New Zealand", "OESZ": "Oras sa Tag-init ng Silangang Europe", "OEZ": "Standard na Oras sa Silangang Europe", "OMSST": "Oras sa Tag-init ng Omsk", "OMST": "Standard na Oras sa Omsk", "PDT": "Daylight na Oras sa Pasipiko sa Hilagang Amerika", "PDTM": "Daylight na Oras sa Pasipiko ng Mexico", "PETDT": "Summer Time sa Petropavlovsk-Kamchatski", "PETST": "Standard Time sa Petropavlovsk-Kamchatski", "PGT": "Oras sa Papua New Guinea", "PHOT": "Oras sa Phoenix Islands", "PKT": "Standard na Oras sa Pakistan", "PKT DST": "Oras sa Tag-init ng Pakistan", "PMDT": "Daylight na Oras sa Saint Pierre & Miquelon", "PMST": "Standard na Oras sa Saint Pierre & Miquelon", "PONT": "Oras sa Ponape", "PST": "Standard na Oras sa Pasipiko sa Hilagang Amerika", "PST Philippine": "Standard na Oras sa Pilipinas", "PST Philippine DST": "Oras sa Tag-init ng Pilipinas", "PST Pitcairn": "Oras sa Pitcairn", "PSTM": "Standard na Oras sa Pasipiko ng Mexico", "PWT": "Oras sa Palau", "PYST": "Oras sa Tag-init ng Paraguay", "PYT": "Standard na Oras sa Paraguay", "PYT Korea": "Oras sa Pyongyang", "RET": "Oras sa Reunion", "ROTT": "Oras sa Rothera", "SAKST": "Oras sa Tag-init ng Sakhalin", "SAKT": "Standard na Oras sa Sakhalin", "SAMST": "Samara Daylight", "SAMT": "Standard Time sa Samara", "SAST": "Oras sa Timog Africa", "SBT": "Oras sa Solomon Islands", "SCT": "Oras sa Seychelles", "SGT": "Standard na Oras sa Singapore", "SLST": "SLST", "SRT": "Oras sa Suriname", "SST Samoa": "Standard na Oras sa Samoa", "SST Samoa Apia": "Standard na Oras sa Apia", "SST Samoa Apia DST": "Daylight Time sa Apia", "SST Samoa DST": "Daylight Time sa Samoa", "SYOT": "Oras sa Syowa", "TAAF": "Oras sa Katimugang France at Antartiko", "TAHT": "Oras sa Tahiti", "TJT": "Oras sa Tajikistan", "TKT": "Oras sa Tokelau", "TLT": "Oras sa East Timor", "TMST": "Oras sa Tag-init ng Turkmenistan", "TMT": "Standard na Oras sa Turkmenistan", "TOST": "Oras sa Tag-init ng Tonga", "TOT": "Standard na Oras sa Tonga", "TVT": "Oras sa Tuvalu", "TWT": "Standard na Oras sa Taipei", "TWT DST": "Daylight Time sa Taipei", "ULAST": "Oras sa Tag-init ng Ulan Bator", "ULAT": "Standard na Oras sa Ulan Bator", "UYST": "Oras sa Tag-init ng Uruguay", "UYT": "Standard na Oras sa Uruguay", "UZT": "Standard na Oras sa Uzbekistan", "UZT DST": "Oras sa Tag-init ng Uzbekistan", "VET": "Oras sa Venezuela", "VLAST": "Oras sa Tag-init ng Vladivostok", "VLAT": "Standard na Oras sa Vladivostok", "VOLST": "Oras sa Tag-init ng Volgograd", "VOLT": "Standard na Oras sa Volgograd", "VOST": "Oras sa Vostok", "VUT": "Standard na Oras sa Vanuatu", "VUT DST": "Oras sa Tag-init ng Vanuatu", "WAKT": "Oras sa Wake Island", "WARST": "Oras sa Tag-init ng Kanlurang Argentina", "WART": "Standard na Oras sa Kanlurang Argentina", "WAST": "Oras sa Kanlurang Africa", "WAT": "Oras sa Kanlurang Africa", "WESZ": "Oras sa Tag-init ng Kanlurang Europe", "WEZ": "Standard na Oras sa Kanlurang Europe", "WFT": "Oras sa Wallis & Futuna", "WGST": "Oras sa Tag-init ng Kanlurang Greenland", "WGT": "Standard na Oras sa Kanlurang Greenland", "WIB": "Oras sa Kanlurang Indonesia", "WIT": "Oras sa Silangang Indonesia", "WITA": "Oras sa Gitnang Indonesia", "YAKST": "Oras sa Tag-init ng Yakutsk", "YAKT": "Standard na Oras sa Yakutsk", "YEKST": "Oras sa Tag-init ng Yekaterinburg", "YEKT": "Standard na Oras sa Yekaterinburg", "YST": "Oras sa Yukon", "МСК": "Standard na Oras sa Moscow", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Oras sa Kanlurang Kazakhstan", "شىعىش قازاق ەلى": "Oras sa Silangang Kazakhstan", "قازاق ەلى": "Oras ng Kazakhstan", "قىرعىزستان": "Oras sa Kyrgystan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Oras sa Tag-init ng Peru"},
	}
}

// Locale returns the current translators string locale
func (fil *fil) Locale() string {
	return fil.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fil'
func (fil *fil) PluralsCardinal() []locales.PluralRule {
	return fil.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fil'
func (fil *fil) PluralsOrdinal() []locales.PluralRule {
	return fil.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fil'
func (fil *fil) PluralsRange() []locales.PluralRule {
	return fil.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fil'
func (fil *fil) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fil'
func (fil *fil) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fil'
func (fil *fil) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := fil.CardinalPluralRule(num1, v1)
	end := fil.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fil *fil) MonthAbbreviated(month time.Month) string {
	return fil.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fil *fil) MonthsAbbreviated() []string {
	return fil.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fil *fil) MonthNarrow(month time.Month) string {
	return fil.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fil *fil) MonthsNarrow() []string {
	return fil.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fil *fil) MonthWide(month time.Month) string {
	return fil.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fil *fil) MonthsWide() []string {
	return fil.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fil *fil) WeekdayAbbreviated(weekday time.Weekday) string {
	return fil.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fil *fil) WeekdaysAbbreviated() []string {
	return fil.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fil *fil) WeekdayNarrow(weekday time.Weekday) string {
	return fil.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fil *fil) WeekdaysNarrow() []string {
	return fil.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fil *fil) WeekdayShort(weekday time.Weekday) string {
	return fil.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fil *fil) WeekdaysShort() []string {
	return fil.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fil *fil) WeekdayWide(weekday time.Weekday) string {
	return fil.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fil *fil) WeekdaysWide() []string {
	return fil.daysWide
}

// Decimal returns the decimal point of number
func (fil *fil) Decimal() string {
	return fil.decimal
}

// Group returns the group of number
func (fil *fil) Group() string {
	return fil.group
}

// Group returns the minus sign of number
func (fil *fil) Minus() string {
	return fil.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fil' and handles both Whole and Real numbers based on 'v'
func (fil *fil) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fil.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fil.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fil' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fil *fil) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fil.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fil.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fil'
func (fil *fil) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fil.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fil.group[0])
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
		b = append(b, fil.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fil.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fil'
// in accounting notation.
func (fil *fil) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fil.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fil.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fil.group[0])
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

		b = append(b, fil.currencyNegativePrefix[0])

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
			b = append(b, fil.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fil.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fil'
func (fil *fil) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'fil'
func (fil *fil) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, fil.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'fil'
func (fil *fil) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, fil.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'fil'
func (fil *fil) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, fil.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, fil.monthsWide[t.Month()]...)
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

// FmtTimeShort returns the short time representation of 't' for 'fil'
func (fil *fil) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fil'
func (fil *fil) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fil.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fil'
func (fil *fil) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fil.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fil'
func (fil *fil) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, fil.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fil.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, fil.periodsAbbreviated[0]...)
	} else {
		b = append(b, fil.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fil.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
