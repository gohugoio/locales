package szl_PL

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type szl_PL struct {
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

// New returns a new instance of translator for the 'szl_PL' locale
func New() locales.Translator {
	return &szl_PL{
		locale:                 "szl_PL",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "sty", "lut", "mar", "kwi", "moj", "czy", "lip", "siy", "wrz", "paź", "lis", "gru"},
		monthsNarrow:           []string{"", "S", "L", "M", "K", "M", "C", "L", "S", "W", "P", "L", "G"},
		monthsWide:             []string{"", "stycznia", "lutego", "marca", "kwietnia", "moja", "czyrwca", "lipca", "siyrpnia", "września", "października", "listopada", "grudnia"},
		daysAbbreviated:        []string{"niy", "pyń", "wto", "str", "szt", "piō", "sob"},
		daysNarrow:             []string{"n", "p", "w", "s", "s", "p", "s"},
		daysShort:              []string{"nd", "pń", "wt", "st", "sz", "pt", "sb"},
		daysWide:               []string{"niydziela", "pyńdziałek", "wtorek", "strzoda", "sztwortek", "piōntek", "sobota"},
		timezones:              map[string]string{"ACDT": "postrzodkowoaustralijski latowy czas", "ACST": "postrzodkowoaustralijski sztandardowy czas", "ACT": "ACT", "ACWDT": "postrzodkowo-zachodnioaustralijski latowy czas", "ACWST": "postrzodkowo-zachodnioaustralijski sztandardowy czas", "ADT": "czas atlantycki latowy", "ADT Arabia": "Pōłwysep Arabski (latowy czas)", "AEDT": "wschodnioaustralijski latowy czas", "AEST": "wschodnioaustralijski sztandardowy czas", "AFT": "Afganistan", "AKDT": "alaskański czas latowy", "AKST": "alaskański czas sztandardowy", "AMST": "amazōński latowy czas", "AMST Armenia": "Armynijo (latowy czas)", "AMT": "amazōński sztandardowy czas", "AMT Armenia": "Armynijo (sztandardowy czas)", "ANAST": "latowy czas Anadyr", "ANAT": "sztandardowy czas Anadyr", "ARST": "Argyntyna (latowy czas)", "ART": "Argyntyna (sztandardowy czas)", "AST": "czas atlantycki sztandardowy", "AST Arabia": "Pōłwysep Arabski (sztandardowy czas)", "AWDT": "zachodnioaustralijski latowy czas", "AWST": "zachodnioaustralijski sztandardowy czas", "AZST": "Azerbejdżan (latowy czas)", "AZT": "Azerbejdżan (sztandardowy czas)", "BDT Bangladesh": "Bangladesz (latowy czas)", "BNT": "BNT", "BOT": "Boliwijo", "BRST": "Brasília (latowy czas)", "BRT": "Brasília (sztandardowy czas)", "BST Bangladesh": "Bangladesz (sztandardowy czas)", "BT": "BT", "CAST": "CAST", "CAT": "postrzodkowoafrykański czas", "CCT": "Wyspy Kokosowe", "CDT": "czas postrzodkowoamerykański latowy", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "Chuuk", "CKT": "Wyspy Cooka (sztandardowy czas)", "CKT DST": "Wyspy Cooka (latowy czas)", "CLST": "Chile (latowy czas)", "CLT": "Chile (sztandardowy czas)", "COST": "Kolumbijo (latowy czas)", "COT": "Kolumbijo (sztandardowy czas)", "CST": "czas postrzodkowoamerykański sztandardowy", "CST China": "Chiny (sztandardowy czas)", "CST China DST": "Chiny (latowy czas)", "CVST": "Wyspy Zielōnego Przilōndka (latowy czas)", "CVT": "Wyspy Zielōnego Przilōndka (sztandardowy czas)", "CXT": "Godnio Wyspa", "ChST": "Czamorro", "ChST NMI": "ChST NMI", "CuDT": "Kuba (latowy czas)", "CuST": "Kuba (sztandardowy czas)", "DAVT": "DAVT", "DDUT": "Dumont-d’Urville", "EASST": "Wyspa Wielkanocno (latowy czas)", "EAST": "Wyspa Wielkanocno (sztandardowy czas)", "EAT": "wschodnioafrykański czas", "ECT": "Ekwador", "EDT": "czas wschodnioamerykański latowy", "EGDT": "Grynlandyjo Wschodnia (latowy czas)", "EGST": "Grynlandyjo Wschodnia (sztandardowy czas)", "EST": "czas wschodnioamerykański sztandardowy", "FEET": "wschodnioeuropejski dalszy czas", "FJT": "Fidżi (sztandardowy czas)", "FJT Summer": "Fidżi (latowy czas)", "FKST": "Falklandy (latowy czas)", "FKT": "Falklandy (sztandardowy czas)", "FNST": "Fernando de Noronha (latowy czas)", "FNT": "Fernando de Noronha (sztandardowy czas)", "GALT": "czas Galapagos", "GAMT": "Wyspy Gambiera", "GEST": "Gruzyjo (latowy czas)", "GET": "Gruzyjo (sztandardowy czas)", "GFT": "Gujana Francusko", "GIT": "Gilbertowe Wyspy", "GMT": "uniwersalnego czasu", "GNSST": "GNSST", "GNST": "GNST", "GST": "Georgia Połedniowo", "GST Guam": "GST Guam", "GYT": "Gujana", "HADT": "Hawaje-Aleuty (sztandardowy czas)", "HAST": "Hawaje-Aleuty (sztandardowy czas)", "HKST": "Hōngkōng (latowy czas)", "HKT": "Hōngkōng (sztandardowy czas)", "HOVST": "Kobdo (latowy czas)", "HOVT": "Kobdo (sztandardowy czas)", "ICT": "indochiński czas", "IDT": "Izrael (latowy czas)", "IOT": "Ôcean Indyjski", "IRKST": "Irkuck (latowy czas)", "IRKT": "Irkuck (sztandardowy czas)", "IRST": "IRST", "IRST DST": "IRST DST", "IST": "indyjski sztandardowy czas", "IST Israel": "Izrael (sztandardowy czas)", "JDT": "Japōnijo (latowy czas)", "JST": "Japōnijo (sztandardowy czas)", "KOST": "KOST", "KRAST": "Krasnojarsk (latowy czas)", "KRAT": "Krasnojarsk (sztandardowy czas)", "KST": "KST", "KST DST": "KST DST", "LHDT": "Lord Howe (latowy czas)", "LHST": "Lord Howe (sztandardowy czas)", "LINT": "Line Islands", "MAGST": "MAGST", "MAGT": "MAGT", "MART": "Markizy", "MAWT": "MAWT", "MDT": "MDT", "MESZ": "postrzodkowoeuropejski latowy czas", "MEZ": "postrzodkowoeuropejski sztandardowy czas", "MHT": "Wyspy Marshalla", "MMT": "Mjanma", "MSD": "Moskwa (latowy)", "MST": "MST", "MUST": "MUST", "MUT": "MUT", "MVT": "Malediwy", "MYT": "Malezyjo", "NCT": "Nowo Kaledōnijo (sztandardowy czas)", "NDT": "Nowo Fundlandyjo (latowy czas)", "NDT New Caledonia": "Nowo Kaledōnijo (latowy czas)", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "Nowosybirsk (latowy czas)", "NOVT": "Nowosybirsk (sztandardowy czas)", "NPT": "NPT", "NRT": "NRT", "NST": "Nowo Fundlandyjo (sztandardowy czas)", "NUT": "NUT", "NZDT": "Nowo Zelandyjo (latowy czas)", "NZST": "Nowo Zelandyjo (sztandardowy czas)", "OESZ": "wschodnioeuropejski latowy czas", "OEZ": "wschodnioeuropejski sztandardowy czas", "OMSST": "Ômsk (latowy czas)", "OMST": "Ômsk (sztandardowy czas)", "PDT": "czas pacyficzny latowy", "PDTM": "Meksyk (czas pacyficzny latowy)", "PETDT": "czas Pietropawłowsk Kamczacki latowy", "PETST": "sztandardowy czas Pietropawłowsk Kamczacki", "PGT": "Papua-Nowo Gwinea", "PHOT": "Fyniks", "PKT": "PKT", "PKT DST": "PKT DST", "PMDT": "Saint-Pierre i Miquelon (latowy czas)", "PMST": "Saint-Pierre i Miquelon (sztandardowy czas)", "PONT": "Pohnpei", "PST": "czas pacyficzny sztandardowy", "PST Philippine": "Filipiny (sztandardowy czas)", "PST Philippine DST": "Filipiny (latowy czas)", "PST Pitcairn": "PST Pitcairn", "PSTM": "Meksyk (czas pacyficzny sztandardowy)", "PWT": "PWT", "PYST": "Paragwaj (latowy czas)", "PYT": "Paragwaj (sztandardowy czas)", "PYT Korea": "Pjōngjang", "RET": "RET", "ROTT": "ROTT", "SAKST": "Sachalin (latowy czas)", "SAKT": "Sachalin (sztandardowy czas)", "SAMST": "czas Samara latowy", "SAMT": "sztandardowy czas Samara", "SAST": "połedniowoafrykański czas", "SBT": "Wyspy Salōmōna", "SCT": "Seszele", "SGT": "Singapur", "SLST": "SLST", "SRT": "Surinam", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "Francuske Terytoria Połedniowe i Antarktyczne", "TAHT": "TAHT", "TJT": "Tadżykistan", "TKT": "TKT", "TLT": "Timor Wschodni", "TMST": "Turkmynistan (latowy czas)", "TMT": "Turkmynistan (sztandardowy czas)", "TOST": "Tōnga (latowy czas)", "TOT": "Tōnga (sztandardowy czas)", "TVT": "TVT", "TWT": "Tajpej (sztandardowy czas)", "TWT DST": "Tajpej (latowy czas)", "ULAST": "Ułan Bator (latowy czas)", "ULAT": "Ułan Bator (sztandardowy czas)", "UYST": "Urugwaj (latowy czas)", "UYT": "Urugwaj (sztandardowy czas)", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "Wynezuela", "VLAST": "Władywostok (latowy czas)", "VLAT": "Władywostok (sztandardowy czas)", "VOLST": "Wołgograd (latowy czas)", "VOLT": "Wołgograd (sztandardowy czas)", "VOST": "Wostok", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "Argyntyna Zachodnio (latowy czas)", "WART": "Argyntyna Zachodnio (sztandardowy czas)", "WAST": "zachodnioafrykański czas", "WAT": "zachodnioafrykański czas", "WESZ": "zachodnioeuropejski latowy czas", "WEZ": "zachodnioeuropejski sztandardowy czas", "WFT": "Wallis i Futuna", "WGST": "Grynlandyjo Zachodnio (latowy czas)", "WGT": "Grynlandyjo Zachodnio (sztandardowy czas)", "WIB": "Indōnezyjo Zachodnio", "WIT": "Indōnezyjo Wschodnio", "WITA": "Indōnezyjo Postrzodkowo", "YAKST": "Jakuck (latowy czas)", "YAKT": "Jakuck (sztandardowy czas)", "YEKST": "Jekaterynburg (latowy czas)", "YEKT": "Jekaterynburg (sztandardowy czas)", "YST": "czas jukōński", "МСК": "Moskwa (sztandardowy)", "اقتاۋ": "czas auktaucki sztandardowy", "اقتاۋ قالاسى": "czas auktaucki latowy", "اقتوبە": "czas aktiubiński sztandardowy", "اقتوبە قالاسى": "czas aktiubiński latowy", "الماتى": "czas ałmacki sztandardowy", "الماتى قالاسى": "czas ałmacki latowy", "باتىس قازاق ەلى": "Kazachstan Zachodni", "شىعىش قازاق ەلى": "Kazachstan Wschodni", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Kirgistan", "قىزىلوردا": "czas kyzyłordzki sztandardowy", "قىزىلوردا قالاسى": "czas kyzyłordzki latowy", "∅∅∅": "Peru (latowy czas)"},
	}
}

// Locale returns the current translators string locale
func (szl *szl_PL) Locale() string {
	return szl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'szl_PL'
func (szl *szl_PL) PluralsCardinal() []locales.PluralRule {
	return szl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'szl_PL'
func (szl *szl_PL) PluralsOrdinal() []locales.PluralRule {
	return szl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'szl_PL'
func (szl *szl_PL) PluralsRange() []locales.PluralRule {
	return szl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'szl_PL'
func (szl *szl_PL) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'szl_PL'
func (szl *szl_PL) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'szl_PL'
func (szl *szl_PL) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (szl *szl_PL) MonthAbbreviated(month time.Month) string {
	return szl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (szl *szl_PL) MonthsAbbreviated() []string {
	return szl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (szl *szl_PL) MonthNarrow(month time.Month) string {
	return szl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (szl *szl_PL) MonthsNarrow() []string {
	return szl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (szl *szl_PL) MonthWide(month time.Month) string {
	return szl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (szl *szl_PL) MonthsWide() []string {
	return szl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (szl *szl_PL) WeekdayAbbreviated(weekday time.Weekday) string {
	return szl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (szl *szl_PL) WeekdaysAbbreviated() []string {
	return szl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (szl *szl_PL) WeekdayNarrow(weekday time.Weekday) string {
	return szl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (szl *szl_PL) WeekdaysNarrow() []string {
	return szl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (szl *szl_PL) WeekdayShort(weekday time.Weekday) string {
	return szl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (szl *szl_PL) WeekdaysShort() []string {
	return szl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (szl *szl_PL) WeekdayWide(weekday time.Weekday) string {
	return szl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (szl *szl_PL) WeekdaysWide() []string {
	return szl.daysWide
}

// Decimal returns the decimal point of number
func (szl *szl_PL) Decimal() string {
	return szl.decimal
}

// Group returns the group of number
func (szl *szl_PL) Group() string {
	return szl.group
}

// Group returns the minus sign of number
func (szl *szl_PL) Minus() string {
	return szl.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'szl_PL' and handles both Whole and Real numbers based on 'v'
func (szl *szl_PL) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, szl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(szl.group) - 1; j >= 0; j-- {
					b = append(b, szl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, szl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'szl_PL' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (szl *szl_PL) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, szl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, szl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, szl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'szl_PL'
func (szl *szl_PL) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := szl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, szl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(szl.group) - 1; j >= 0; j-- {
					b = append(b, szl.group[j])
				}
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

	for j := len(szl.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, szl.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, szl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, szl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'szl_PL'
// in accounting notation.
func (szl *szl_PL) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := szl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, szl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(szl.group) - 1; j >= 0; j-- {
					b = append(b, szl.group[j])
				}
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

		for j := len(szl.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, szl.currencyNegativePrefix[j])
		}

		b = append(b, szl.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(szl.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, szl.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, szl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, szl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, szl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, szl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, szl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'szl_PL'
func (szl *szl_PL) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, szl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := szl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
