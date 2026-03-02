package wo

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type wo struct {
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
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'wo' locale
func New() locales.Translator {
	return &wo{
		locale:                 "wo",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "Vote $", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGPP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS.", "GIIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GT Q", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRKS", "HTG", "Vote Ft", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "Vote MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN.", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "Vote lei", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "K",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "K",
		monthsAbbreviated:      []string{"", "Sam", "Few", "Mar", "Awr", "Mee", "Suw", "Sul", "Ut", "Sàt", "Okt", "Now", "Des"},
		monthsWide:             []string{"", "Samwiyee", "Fewriyee", "Mars", "Awril", "Mee", "Suwe", "Sulet", "Ut", "Sàttumbar", "Oktoobar", "Nowàmbar", "Desàmbar"},
		daysAbbreviated:        []string{"Dib", "Alt", "Tal", "Àla", "Alx", "Àjj", "Ase"},
		daysNarrow:             []string{"Dib", "Alt", "Tal", "Àla", "Alx", "Àjj", "Ase"},
		daysWide:               []string{"Dibéer", "Altine", "Talaata", "Àlarba", "Alxamis", "Àjjuma", "Aseer"},
		periodsAbbreviated:     []string{"Sub", "Ngo"},
		erasAbbreviated:        []string{"JC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "Waxtu bëccëg ci diggu Australie", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Waxtu bëccëg ci diggu sowwu Australie", "ACWST": "Waxtu buñ miin ci diggu sowwu Australie", "ADT": "ADT (waxtu bëccëgu atlàntik)", "ADT Arabia": "Waxtu bëccëg ci Araab", "AEDT": "Waxtu buñ miin ci penku Australie", "AEST": "Waxtu penku bu Australie", "AFT": "waxtu Afganistan", "AKDT": "Waxtu bëccëg ci Alaska", "AKST": "Waxtu buñ miin ci Alaska", "AMST": "Waxtu ete bu Amazon", "AMST Armenia": "Waxtu ete bu Armeni", "AMT": "Waxtu buñ jagleel Amazon", "AMT Armenia": "Waxtu buñ miin ci Armeni", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Waxtu ete bu Argentine", "ART": "Waxtu buñ miin ci Arsantiin", "AST": "AST (waxtu estàndaaru penku)", "AST Arabia": "Waxtu buñ miin ci Araab", "AWDT": "Waxtu bëccëg bu sowwu Australie", "AWST": "Waxtu buñ miin ci sowwu Australie", "AZST": "Waxtu ete bu Azerbaïdjan", "AZT": "Waxtu Azerbaïdjan", "BDT Bangladesh": "Waxtu ete bu Bangladesh", "BNT": "Brunei Darussalam", "BOT": "Waxtu Bolivie", "BRST": "Brasilia summer time", "BRT": "Brasilia time", "BST Bangladesh": "Waxtu buñ miin ci Bangladesh", "BT": "waxtu Bhoutan", "CAST": "CAST", "CAT": "Waxtu Afrique Centrale", "CCT": "Waxtu ile Cocos", "CDT": "CDT (waxtu bëccëgu sàntaraal", "CHADT": "Chatham Daylight Time", "CHAST": "Chatham Standard Time", "CHUT": "Waxtu Chuuk", "CKT": "Waxtu buñ miin ci Ile Cook", "CKT DST": "Ile Cook xaaju ete", "CLST": "Waxtu ete bu Sili", "CLT": "Waxtu buñ miin ci Sili", "COST": "Jamonoy ete ci Kolombi", "COT": "Waxtu buñ miin ci Kolombi", "CST": "CST (waxtu estàndaaru sàntaraal)", "CST China": "Waxtu buñ miin ci Chine", "CST China DST": "Chine waxtu bëccëg", "CVST": "Cape Verde ci jamonoy ete", "CVT": "Cape Verde waxtu", "CXT": "waxtu ile bu noel", "ChST": "Chamorro Standard Time", "ChST NMI": "ChST NMI", "CuDT": "Cuba waxtu bëccëg", "CuST": "waxtu buñ miin ci Cuba", "DAVT": "Waxtu Davis", "DDUT": "Dumont-d’Urville", "EASST": "Jamonoy ete ci Ile de Pâques", "EAST": "Waxtu buñ miin ci Ile de Pâques", "EAT": "Waxtu Afrique sowwu jant", "ECT": "waxtu Ecuador", "EDT": "EDT (waxtu bëccëgu penku)", "EGDT": "Waxtu ete bu penku Greenland", "EGST": "Waxtu buñ miin ci penku Greenland", "EST": "EST (waxtu estàndaaru penku)", "FEET": "waxtu penku Europe", "FJT": "Fidji", "FJT Summer": "Jamonoy ete ci Fiji", "FKST": "Jamonoy ete ci ile Falkland", "FKT": "Falkland waxtu buñ miin", "FNST": "Fernando de noronha temps d’été", "FNT": "Vernando de Noronha", "GALT": "waxtu galapagos", "GAMT": "Waxtu Gambier", "GEST": "Georgie waxtu ete", "GET": "Georgie waxtu", "GFT": "Guyane française", "GIT": "waxtu ile Gilbert", "GMT": "GMT (waxtu Greenwich)", "GNSST": "GNSST", "GNST": "GNST", "GST": "Georgie du Sud", "GST Guam": "GST Guam", "GYT": "Waxtu Guyana", "HADT": "Waxtu bëccëg bu Hawaii-Aleutian", "HAST": "Waxtu buñ jagleel Hawaii-Aleutian", "HKST": "Jamonoy ete ci Hong Kong", "HKT": "waxtu buñ miin ci Hong Kong", "HOVST": "Hovd waxtu ete", "HOVT": "Hovd waxtu standard", "ICT": "waxtu Indochine", "IDT": "Israel waxtu bëccëg", "IOT": "Waxtu géeju Inde", "IRKST": "Waxtu ete bu Irkutsk", "IRKT": "waxtu Irkutsk time", "IRST": "Iran waxtu buñ miin", "IRST DST": "Waxtu bëccëg ci Iran", "IST": "Waxtu Inde", "IST Israel": "Waxtu buñ miin ci Israel", "JDT": "Japon waxtu bëccëg", "JST": "Waxtu japon", "KOST": "Waxtu Kosrae", "KRAST": "Krasnoyarsk ci jamonoy ete", "KRAT": "Krasnoyarsk waxtu", "KST": "waxtu buñ miin ci Kore", "KST DST": "waxtu bëccëg ci Kore", "LHDT": "ord Howe Daylight Time", "LHST": "Lord Howe waxtu buñ miin", "LINT": "Waxtu Ile Line", "MAGST": "Waxtu ete bu Magadan", "MAGT": "Magadan, waxtu", "MART": "Waxtu Marquesas", "MAWT": "waxtu Mawson", "MDT": "MDT", "MESZ": "CEST (waxtu ete wu ëroop sàntaraal)", "MEZ": "CEST (waxtu estàndaaru ëroop sàntaraal)", "MHT": "Waxtu Ile Marshall", "MMT": "waxtu Myanmar", "MSD": "Waxtu ete bu Moscou", "MST": "MST", "MUST": "Waxtu ete bu Maurice", "MUT": "Waxtu buñ miin ci Maurice", "MVT": "Waxtu Maldives", "MYT": "Malaysie", "NCT": "Waxtu buñ miin ci Caledonie bu bees", "NDT": "Terre-Neuve", "NDT New Caledonia": "Waxtu ete bu Nouvelle Caledonie", "NFDT": "waxtu bëccëg ci ile Norfolk", "NFT": "Waxtu buñ miin ci Ile Norfolk", "NOVST": "Novosibirsk ci jamonoy ete", "NOVT": "Novosibirsk waxtu", "NPT": "waxtu Nepal", "NRT": "waxtu Nauru", "NST": "Terre Neuve", "NUT": "Waxtu Niue", "NZDT": "Nouvelle-Zélande", "NZST": "Waxtu buñ miin ci Nouvelle-Zélande", "OESZ": "EEST (waxtu ete wu ëroop u penku)", "OEZ": "EEST (waxtu estàndaaru ëroop u penku)", "OMSST": "Omsk waxtu ete", "OMST": "Waxtu buñ miin ci Omsk", "PDT": "PDT (waxtu bëccëgu pasifik)", "PDTM": "Waxtu bëccëg ci Pacific Mexique", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papouasie-Nouvelle-Guiné", "PHOT": "waxtu ile Phoenix", "PKT": "Waxtu buñ miin ci Pakistan", "PKT DST": "Waxtu ete bu Pakistan", "PMDT": "Saint Pierre and Miquelon", "PMST": "Saint Pierre & Miquelon", "PONT": "Waxtu Ponape", "PST": "PST (waxtu estàndaaru pasifik)", "PST Philippine": "waxtu buñ miin ci filipiin", "PST Philippine DST": "Jamonoy ete ci Philippines", "PST Pitcairn": "Waxtu Pitcairn", "PSTM": "Waxtu buñ miin ci pasifik bu Mexico", "PWT": "waxtu Palau", "PYST": "Paraguay waxtu ete", "PYT": "paraguay waxtu", "PYT Korea": "waxtu Pyongyang", "RET": "waxtu ndaje", "ROTT": "Waxtu Rotera", "SAKST": "Sakhalin Sakhalin", "SAKT": "Saxalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Afrique du Sud", "SBT": "Waxtu Ile Solomon", "SCT": "Waxtu Seychelles", "SGT": "waxtu buñ miin ci Singapuur", "SLST": "SLST", "SRT": "Waxtu Surinam", "SST Samoa": "Samoa Standard Time", "SST Samoa Apia": "Waxtu buñ miin ci Apia", "SST Samoa Apia DST": "Apia waxtu bëccëg", "SST Samoa DST": "Samoa waxtu bëccëg", "SYOT": "waxtu syowa", "TAAF": "Waxtu Sud ak Antarctique bu Français", "TAHT": "waxtu Tahiti", "TJT": "Waxtu Tajikistaan", "TKT": "Tokelau time", "TLT": "Timor oriental", "TMST": "Waxtu ete bu Turkmenistan", "TMT": "Waxtu buñ miin", "TOST": "Jamonoy ete ci Tonga", "TOT": "Tonga waxtu buñ miin", "TVT": "Waxtu Tuvalu", "TWT": "Waxtu buñ miin ci Taipei", "TWT DST": "Taipei waxtu leeralu bis", "ULAST": "Ulaan Baatar time", "ULAT": "Ulaanbatar", "UYST": "Uruguay waxtu ete", "UYT": "Uruguay waxtu buñ miin", "UZT": "Waxtu buñ miin ci Ouzbékistan", "UZT DST": "Waxtu ete bu Ouzbékistan", "VET": "Waxtu Venezuela", "VLAST": "Vladivostok ci jamonoy ete", "VLAT": "Vladivostok ci waxtu", "VOLST": "Jamonoy ete bu Volgograd", "VOLT": "Volgograd waxtu buñ miin", "VOST": "Waxtu Vostok", "VUT": "Waxtu miin", "VUT DST": "Waxtu ete bu Vanuatu", "WAKT": "Waxtu Ile Wake", "WARST": "waxtu ete bu sowwu Argentine", "WART": "Waxtu buñ miin ci sowwu Argentine", "WAST": "Waxtu sowwu Afrique", "WAT": "Waxtu sowwu Afrique", "WESZ": "WEST (waxtu ete wu ëroop u sowwu-jant)", "WEZ": "WEST (waxtu estàndaaru ëroop u sowwu-jant)", "WFT": "Wallis & Futuna Time", "WGST": "waxtu ete bu sowwu Groenland", "WGT": "waxtu buñ miin ci sowwu Groenland", "WIB": "waxtu sowwu Enndonesi", "WIT": "waxtu penku Enndonesi", "WITA": "Waxtu Enndonesi bu diggi bi", "YAKST": "Waxtu ete bu Yakutsk", "YAKT": "Waxtu Yakutsk", "YEKST": "Jamonoy ete", "YEKT": "Yekatérinbourg", "YST": "Waxtu Yukon", "МСК": "Moscow Waxtu", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Kazakhstan bu sowwu jant", "شىعىش قازاق ەلى": "Kazakhstan penku", "قازاق ەلى": "Waxtu Kazakhstaan", "قىرعىزستان": "Waxtu Kirgistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Azores waxtu ete"},
	}
}

// Locale returns the current translators string locale
func (wo *wo) Locale() string {
	return wo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'wo'
func (wo *wo) PluralsCardinal() []locales.PluralRule {
	return wo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'wo'
func (wo *wo) PluralsOrdinal() []locales.PluralRule {
	return wo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'wo'
func (wo *wo) PluralsRange() []locales.PluralRule {
	return wo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'wo'
func (wo *wo) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'wo'
func (wo *wo) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'wo'
func (wo *wo) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (wo *wo) MonthAbbreviated(month time.Month) string {
	return wo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (wo *wo) MonthsAbbreviated() []string {
	return wo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (wo *wo) MonthNarrow(month time.Month) string {
	return wo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (wo *wo) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (wo *wo) MonthWide(month time.Month) string {
	return wo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (wo *wo) MonthsWide() []string {
	return wo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (wo *wo) WeekdayAbbreviated(weekday time.Weekday) string {
	return wo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (wo *wo) WeekdaysAbbreviated() []string {
	return wo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (wo *wo) WeekdayNarrow(weekday time.Weekday) string {
	return wo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (wo *wo) WeekdaysNarrow() []string {
	return wo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (wo *wo) WeekdayShort(weekday time.Weekday) string {
	return wo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (wo *wo) WeekdaysShort() []string {
	return wo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (wo *wo) WeekdayWide(weekday time.Weekday) string {
	return wo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (wo *wo) WeekdaysWide() []string {
	return wo.daysWide
}

// Decimal returns the decimal point of number
func (wo *wo) Decimal() string {
	return wo.decimal
}

// Group returns the group of number
func (wo *wo) Group() string {
	return wo.group
}

// Group returns the minus sign of number
func (wo *wo) Minus() string {
	return wo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'wo' and handles both Whole and Real numbers based on 'v'
func (wo *wo) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'wo' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (wo *wo) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'wo'
func (wo *wo) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := wo.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, wo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(wo.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, wo.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, wo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, wo.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'wo'
// in accounting notation.
func (wo *wo) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := wo.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, wo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(wo.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, wo.currencyNegativePrefix[j])
		}

		b = append(b, wo.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(wo.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, wo.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, wo.currencyNegativeSuffix...)
	} else {

		b = append(b, wo.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'wo'
func (wo *wo) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'wo'
func (wo *wo) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, wo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'wo'
func (wo *wo) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, wo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'wo'
func (wo *wo) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, wo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, wo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'wo'
func (wo *wo) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'wo'
func (wo *wo) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'wo'
func (wo *wo) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'wo'
func (wo *wo) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := wo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
