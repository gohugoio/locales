package ku_Arab_IQ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ku_Arab_IQ struct {
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

// New returns a new instance of translator for the 'ku_Arab_IQ' locale
func New() locales.Translator {
	return &ku_Arab_IQ{
		locale:                 "ku_Arab_IQ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "rbn", "sbt", "adr", "nsn", "gln", "hzr", "trm", "tbx", "îln", "cot", "mjd", "brf"},
		monthsNarrow:           []string{"", "R", "S", "A", "N", "G", "H", "T", "T", "Î", "C", "M", "B"},
		monthsWide:             []string{"", "rêbendan", "sibat", "adar", "nîsan", "gulan", "hezîran", "tîrmeh", "tebax", "îlon", "cotmeh", "mijdar", "berfanbar"},
		daysAbbreviated:        []string{"yşm", "dşm", "sşm", "çşm", "pşm", "înî", "şem"},
		daysNarrow:             []string{"Y", "D", "S", "Ç", "P", "Î", "Ş"},
		daysShort:              []string{"yş", "dş", "sş", "çş", "pş", "în", "şm"},
		daysWide:               []string{"yekşem", "duşem", "sêşem", "çarşem", "pêncşem", "înî", "şemî"},
		periodsAbbreviated:     []string{"BN", "PN"},
		periodsNarrow:          []string{"bn", "pn"},
		erasAbbreviated:        []string{"BM", "PM"},
		erasNarrow:             []string{"BM", "PM"},
		erasWide:               []string{"Berî Mîladê", "Piştî Mîladê"},
		timezones:              map[string]string{"ACDT": "Saeta Havînê ya Awistralyaya Navîn", "ACST": "Saeta Havînê ya Acreyê", "ACT": "Saeta Acreyê ya Standard", "ACWDT": "Saeta Havînê ya Rojavaya Navîn a Awistralyayê", "ACWST": "Saeta Standard a Rojavaya Navîn a Awistralyayê", "ADT": "Saeta Havînê ya Atlantîkê", "ADT Arabia": "Saeta Havînê ya Erebistanê", "AEDT": "Saeta Havînê ya Awistralyaya Rojhilat", "AEST": "Saeta Standard a Awistralyaya Rojhilat", "AFT": "Saeta Efxanistanê", "AKDT": "Saeta Havînê ya Alaskayê", "AKST": "Saeta Standard a Alaskayê", "AMST": "Saeta Havînê ya Amazonê", "AMST Armenia": "Saeta Havînê ya Ermenistanê", "AMT": "Saeta Amazonê ya Standard", "AMT Armenia": "Saeta Ermenistanê ya Standard", "ANAST": "Saeta Havînê ya Anadyrê", "ANAT": "Saeta Anadyrê ya Standard", "ARST": "Saeta Havînê ya Arjantînê", "ART": "Saeta Arjantînê ya Standard", "AST": "Saeta Standard a Atlantîkê", "AST Arabia": "Saeta Erebistanê ya Standard", "AWDT": "Saeta Havînê ya Awistralyaya Rojava", "AWST": "Saeta Standard a Awistralyaya Rojava", "AZST": "Saeta Havînê ya Azerbeycanê", "AZT": "Saeta Azerbeycanê ya Standard", "BDT Bangladesh": "Saeta Havînê ya Bengladeşê", "BNT": "Saeta Brûneyê", "BOT": "Saeta Bolîvyayê", "BRST": "Saeta Havînê ya Brasîlyayê", "BRT": "Saeta Brasîlyayê ya Standard", "BST Bangladesh": "Saeta Bengladeşê ya Standard", "BT": "Saeta Bûtanê", "CAST": "Saeta Caseyê", "CAT": "Saeta Afrîkaya Navîn", "CCT": "Saeta Giravên Cocosê", "CDT": "Saeta Havînê ya Merkezî ya Amerîkaya Bakur", "CHADT": "Saeta Havînê ya Chathamê", "CHAST": "Saeta Standard a Chathamê", "CHUT": "Saeta Chuukê", "CKT": "Saeta Giravên Cookê ya Standard", "CKT DST": "Saeta Havînê ya Giravên Cookê", "CLST": "Saeta Havînê ya Şîlîyê", "CLT": "Saeta Şîlîyê ya Standard", "COST": "Saeta Havînê ya Kolombîyayê", "COT": "Saeta Kolombîyayê ya Standard", "CST": "Saeta Merkezî ya Standard a Amerîkaya Bakur", "CST China": "Saeta Çînê ya Standard", "CST China DST": "Saeta Havînê ya Çînê", "CVST": "Saeta Havînê ya Cape Verdeyê", "CVT": "Saeta Standard a Cape Verdeyê", "CXT": "Saeta Girava Christmasê", "ChST": "Saeta Chamorroyê ya Standard", "ChST NMI": "Saeta Giravên Marianaya Bakur", "CuDT": "Saeta Havînê ya Kubayê", "CuST": "Saeta Standard a Kubayê", "DAVT": "Saeta Davîsê", "DDUT": "Saeta Dumont-d’Urvilleyê", "EASST": "Saeta Havînê ya Girava Paskalyayê", "EAST": "Saeta Standard a Girava Paskalyayê", "EAT": "Saeta Afrîkaya Rojhilat", "ECT": "Saeta Ekwadorê", "EDT": "Saeta Havînê ya Rojhilat ya Amerîkaya Bakur", "EGDT": "Saeta Havînê ya Grînlanda Rojhilat", "EGST": "Saeta Standard a Grînlanda Rojhilat", "EST": "Saeta Standard a Rojhilat ya Amerîkaya Bakur", "FEET": "Saeta Ewropaya Rojhilat a Pêştir", "FJT": "Saeta Fîjîyê ya Standard", "FJT Summer": "Saeta Havînê ya Fîjîyê", "FKST": "Saeta Havînê ya Giravên Falklandê", "FKT": "Saeta Giravên Falklandê ya Standard", "FNST": "Saeta Havînê ya Fernando de Noronhayê", "FNT": "Saeta Fernando de Noronhayê ya Standard", "GALT": "Saeta Galapagosê", "GAMT": "Saeta Gambierê", "GEST": "Saeta Havînê ya Gurcistanê", "GET": "Saeta Gurcistanê ya Standard", "GFT": "Saeta Guiyanaya Fransî", "GIT": "Saeta Giravên Gilbertê", "GMT": "Saeta Navînî ya Greenwichê", "GNSST": "Saeta Havînê ya Grînlendayê", "GNST": "Saeta Grînlendayê ya Standard", "GST": "Saeta Kendavê ya Standard", "GST Guam": "Saeta Guamê ya Standard", "GYT": "Saeta Guyanayê", "HADT": "Saeta Havînê ya Hawaii-Aleutianê", "HAST": "Saeta Standard a Hawaii-Aleutianê", "HKST": "Saeta Havînê ya Hong Kongê", "HKT": "Saeta Hong Kongê ya Standard", "HOVST": "Saeta Havînê ya Hovdê", "HOVT": "Saeta Hovdê ya Standard", "ICT": "Saeta Hindiçînê", "IDT": "Saeta Havînê ya Îsraîlê", "IOT": "Saeta Okyanûsa Hindê", "IRKST": "Saeta Havînê ya Irkutskê", "IRKT": "Saeta Irkutskê ya Standard", "IRST": "Saeta Îranê ya Standard", "IRST DST": "Saeta Havînê ya Îranê", "IST": "Saeta Hindistanê ya Standard", "IST Israel": "Saeta Îsraîlê ya Standard", "JDT": "Saeta Havînê ya Japonyayê", "JST": "Saeta Japonyayê ya Standard", "KOST": "Saeta Kosraeyê", "KRAST": "Saeta Havînê ya Krasnoyarskê", "KRAT": "Saeta Krasnoyarskê ya Standard", "KST": "Saeta Koreyê ya Standard", "KST DST": "Saeta Havînê ya Koreyê", "LHDT": "Saeta Havînê ya Lord Howeyê", "LHST": "Saeta Standard a Lord Howeyê", "LINT": "Saeta Giravên Lîneyê", "MAGST": "Saeta Havînê ya Magadanê", "MAGT": "Saeta Magadanê ya Standard", "MART": "Saeta Marquesasê", "MAWT": "Saeta Mawsonê", "MDT": "Saeta Havînê ya Çîyayî ya Amerîkaya Bakur", "MESZ": "Saeta Havînê ya Ewropaya Navîn", "MEZ": "Saeta Standard a Ewropaya Navîn", "MHT": "Saeta Giravên Marşalê", "MMT": "Saeta Myanmarê", "MSD": "Saeta Havînê ya Moskovayê", "MST": "Saeta Standard a Çîyayî ya Amerîkaya Bakur", "MUST": "Saeta Havînê ya Mauritiusê", "MUT": "Saeta Mauritiusê ya Standard", "MVT": "Saeta Maldîvan", "MYT": "Saeta Malezyayê", "NCT": "Saeta Standard a Kaledonyaya Nû", "NDT": "Saeta Havînê ya Newfoundlandê", "NDT New Caledonia": "Saeta Havînê ya Kaledonyaya Nû", "NFDT": "Saeta Havînê ya Girava Norfolkê", "NFT": "Saeta Standard a Girava Norfolkê", "NOVST": "Saeta Havînê ya Novosibirskê", "NOVT": "Saeta Novosibirskê ya Standard", "NPT": "Saeta Nepalê", "NRT": "Saeta Naûrûyê", "NST": "Saeta Standard a Newfoundlandê", "NUT": "Saeta Niueyê", "NZDT": "Saeta Havînê ya Zelandaya Nû", "NZST": "Saeta Standard a Zelandaya Nû", "OESZ": "Saeta Havînê ya Ewropaya Rojhilat", "OEZ": "Saeta Standard a Ewropaya Rojhilat", "OMSST": "Saeta Havînê ya Omskê", "OMST": "Saeta Omskê ya Standard", "PDT": "Saeta Havînê ya Pasîfîkê ya Amerîkaya Bakur", "PDTM": "Saeta Havînê ya Pasîfîka Meksîkayê", "PETDT": "Saeta Havînê ya Kamchatkayê", "PETST": "Saeta Kamchatkayê ya Standard", "PGT": "Saeta Gîneya Nû ya Papûayê", "PHOT": "Saeta Giravên Phoenîks", "PKT": "Saeta Pakistanê ya Standard", "PKT DST": "Saeta Havînê ya Pakistanê", "PMDT": "Saeta Havînê ya Saint Pierre û Miquelonê", "PMST": "Saeta Standard a Saint Pierre û Miquelonê", "PONT": "Saeta Ponapeyê", "PST": "Saeta Standard a Pasîfîkê ya Amerîkaya Bakur", "PST Philippine": "Saeta Fîlîpînê ya Standard", "PST Philippine DST": "Saeta Havînê ya Fîlîpînê", "PST Pitcairn": "Saeta Pitcairnê", "PSTM": "Saeta Standard a Pasîfîka Meksîkayê", "PWT": "Saeta Palauyê", "PYST": "Saeta Havînê ya Paragûayê", "PYT": "Saeta Paragûayê ya Standard", "PYT Korea": "Saeta Pyongyangê", "RET": "Saeta Réunionê", "ROTT": "Saeta Rotherayê", "SAKST": "Saeta Havînê ya Saxalînê", "SAKT": "Saeta Saxalînê ya Standard", "SAMST": "Saeta Havînê ya Samarayê", "SAMT": "Saeta Samarayê ya Standard", "SAST": "Saeta Standard a Afrîkaya Başûr", "SBT": "Saeta Giravên Solomonê", "SCT": "Saeta Seyşelerê", "SGT": "Saeta Sîngapûrê ya Standard", "SLST": "Saeta Lankayê", "SRT": "Saeta Surînamê", "SST Samoa": "Saeta Samoayê ya Standard", "SST Samoa Apia": "Saeta Apiayê ya Standard", "SST Samoa Apia DST": "Saeta Havînê ya Apiayê", "SST Samoa DST": "Saeta Havînê ya Samoayê", "SYOT": "Saeta Syowayê", "TAAF": "Saeta Antarktîka û Başûrê Fransayê", "TAHT": "Saeta Tahîtîyê", "TJT": "Saeta Tacikistanê", "TKT": "Saeta Tokelauyê", "TLT": "Saeta Tîmûra Rojhilat", "TMST": "Saeta Havînê ya Tirkmenistanê", "TMT": "Saeta Tirkmenistanê ya Standard", "TOST": "Saeta Havînê ya Tongayê", "TOT": "Saeta Tongayê ya Standard", "TVT": "Saeta Tûvalûyê", "TWT": "Saeta Taîpeîyê ya Standard", "TWT DST": "Saeta Havînê ya Taîpeîyê", "ULAST": "Saeta Havînê ya Ûlanbatarê", "ULAT": "Saeta Ûlanbatarê ya Standard", "UYST": "Saeta Havînê ya Ûrûgûayê", "UYT": "Saeta Ûrûgûayê ya Standard", "UZT": "Saeta Ozbekistanê ya Standard", "UZT DST": "Saeta Havînê ya Ozbekistanê", "VET": "Saeta Venezûelayê", "VLAST": "Saeta Havînê ya Vladivostokê", "VLAT": "Saeta Vladivostokê ya Standard", "VOLST": "Saeta Havînê ya Volgogradê", "VOLT": "Saeta Volgogradê ya Standard", "VOST": "Saeta Vostokê", "VUT": "Saeta Vanûatûyê ya Standard", "VUT DST": "Saeta Havînê ya Vanûatûyê", "WAKT": "Saeta Girava Wakeyê", "WARST": "Saeta Havînê ya Arjantîna Rojava", "WART": "Saeta Standard a Arjantîna Rojava", "WAST": "Saeta Afrîkaya Rojava", "WAT": "Saeta Afrîkaya Rojava", "WESZ": "Saeta Havînê ya Ewropaya Rojava", "WEZ": "Saeta Standard a Ewropaya Rojava", "WFT": "Saeta Wallis û Futunayê", "WGST": "Saeta Havînê ya Grînlanda Rojava", "WGT": "Saeta Standard a Grînlanda Rojava", "WIB": "Saeta Endonezyaya Rojava", "WIT": "Saeta Endonezyaya Rojhilat", "WITA": "Saeta Endonezyaya Navîn", "YAKST": "Saeta Havînê ya Yakutskê", "YAKT": "Saeta Yakutskê ya Standard", "YEKST": "Saeta Havînê ya Yekaterinburgê", "YEKT": "Saeta Yekaterinburgê ya Standard", "YST": "Saeta Yukonê", "МСК": "Saeta Moskovayê ya Standard", "اقتاۋ": "Saeta Aqtauyê ya Standard", "اقتاۋ قالاسى": "Saeta Havînê ya Aqtauyê", "اقتوبە": "Saeta Aqtobeyê ya Standard", "اقتوبە قالاسى": "Saeta Havînê ya Aqtobeyê", "الماتى": "Saeta Almatîyê ya Standard", "الماتى قالاسى": "Saeta Havînê ya Almatîyê", "باتىس قازاق ەلى": "Saeta Qazaxistana Rojava", "شىعىش قازاق ەلى": "Saeta Qazaxistana Rojhilat", "قازاق ەلى": "Saeta Qazaxistanê", "قىرعىزستان": "Saeta Qirxizistanê", "قىزىلوردا": "Saeta Qizilordayê ya Standard", "قىزىلوردا قالاسى": "Saeta Havînê ya Qizilordayê", "∅∅∅": "Saeta Havînê ya Perûyê"},
	}
}

// Locale returns the current translators string locale
func (ku *ku_Arab_IQ) Locale() string {
	return ku.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) PluralsCardinal() []locales.PluralRule {
	return ku.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) PluralsOrdinal() []locales.PluralRule {
	return ku.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) PluralsRange() []locales.PluralRule {
	return ku.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ku *ku_Arab_IQ) MonthAbbreviated(month time.Month) string {
	return ku.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ku *ku_Arab_IQ) MonthsAbbreviated() []string {
	return ku.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ku *ku_Arab_IQ) MonthNarrow(month time.Month) string {
	return ku.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ku *ku_Arab_IQ) MonthsNarrow() []string {
	return ku.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ku *ku_Arab_IQ) MonthWide(month time.Month) string {
	return ku.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ku *ku_Arab_IQ) MonthsWide() []string {
	return ku.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ku *ku_Arab_IQ) WeekdayAbbreviated(weekday time.Weekday) string {
	return ku.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ku *ku_Arab_IQ) WeekdaysAbbreviated() []string {
	return ku.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ku *ku_Arab_IQ) WeekdayNarrow(weekday time.Weekday) string {
	return ku.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ku *ku_Arab_IQ) WeekdaysNarrow() []string {
	return ku.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ku *ku_Arab_IQ) WeekdayShort(weekday time.Weekday) string {
	return ku.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ku *ku_Arab_IQ) WeekdaysShort() []string {
	return ku.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ku *ku_Arab_IQ) WeekdayWide(weekday time.Weekday) string {
	return ku.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ku *ku_Arab_IQ) WeekdaysWide() []string {
	return ku.daysWide
}

// Decimal returns the decimal point of number
func (ku *ku_Arab_IQ) Decimal() string {
	return ku.decimal
}

// Group returns the group of number
func (ku *ku_Arab_IQ) Group() string {
	return ku.group
}

// Group returns the minus sign of number
func (ku *ku_Arab_IQ) Minus() string {
	return ku.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ku_Arab_IQ' and handles both Whole and Real numbers based on 'v'
func (ku *ku_Arab_IQ) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ku_Arab_IQ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ku *ku_Arab_IQ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.minus[0])
	}

	b = append(b, ku.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ku.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ku.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ku.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ku_Arab_IQ'
// in accounting notation.
func (ku *ku_Arab_IQ) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ku.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ku.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ku.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ku.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ku.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xc3, 0xaa}...)
	b = append(b, []byte{0x20}...)
	b = append(b, ku.monthsWide[t.Month()]...)
	b = append(b, []byte{0x61}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x61, 0x6e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ku.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xc3, 0xaa}...)
	b = append(b, []byte{0x20}...)
	b = append(b, ku.monthsWide[t.Month()]...)
	b = append(b, []byte{0x61}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x61, 0x6e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ku_Arab_IQ'
func (ku *ku_Arab_IQ) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ku.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ku.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
