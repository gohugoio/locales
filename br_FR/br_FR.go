package br_FR

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type br_FR struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	timeSeparator          string
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'br_FR' locale
func New() locales.Translator {
	return &br_FR{
		locale:                 "br_FR",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "$A", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "$CA", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "£ E", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "£ RU", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "$ HK", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "£L", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "$ ZN", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "$ T", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$ SU", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "Gen.", "Cʼhwe.", "Meur.", "Ebr.", "Mae", "Mezh.", "Goue.", "Eost", "Gwen.", "Here", "Du", "Kzu."},
		monthsNarrow:           []string{"", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"},
		monthsWide:             []string{"", "Genver", "Cʼhwevrer", "Meurzh", "Ebrel", "Mae", "Mezheven", "Gouere", "Eost", "Gwengolo", "Here", "Du", "Kerzu"},
		daysAbbreviated:        []string{"Sul", "Lun", "Meu.", "Mer.", "Yaou", "Gwe.", "Sad."},
		daysNarrow:             []string{"Su", "L", "Mz", "Mc", "Y", "G", "Sa"},
		daysWide:               []string{"Sul", "Lun", "Meurzh", "Mercʼher", "Yaou", "Gwener", "Sadorn"},
		timezones:              map[string]string{"ACDT": "eur hañv Kreizaostralia", "ACST": "eur cʼhoañv Kreizaostralia", "ACT": "ACT", "ACWDT": "eur hañv Kreizaostralia ar Cʼhornôg", "ACWST": "eur cʼhoañv Kreizaostralia ar Cʼhornôg", "ADT": "eur hañv an Atlantel", "ADT Arabia": "eur hañv Arabia", "AEDT": "eur hañv Aostralia ar Reter", "AEST": "eur cʼhoañv Aostralia ar Reter", "AFT": "eur Afghanistan", "AKDT": "eur hañv Alaska", "AKST": "eur cʼhoañv Alaska", "AMST": "eur hañv an Amazon", "AMST Armenia": "eur hañv Armenia", "AMT": "eur cʼhoañv an Amazon", "AMT Armenia": "eur cʼhoañv Armenia", "ANAST": "eur hañv Anadyrʼ", "ANAT": "eur cʼhoañv Anadyrʼ", "ARST": "eur hañv Arcʼhantina", "ART": "eur cʼhoañv Arcʼhantina", "AST": "eur cʼhoañv an Atlantel", "AST Arabia": "eur cʼhoañv Arabia", "AWDT": "eur hañv Aostralia ar Cʼhornôg", "AWST": "eur cʼhoañv Aostralia ar Cʼhornôg", "AZST": "eur hañv Azerbaidjan", "AZT": "eur cʼhoañv Azerbaidjan", "BDT Bangladesh": "eur hañv Bangladesh", "BNT": "eur Brunei", "BOT": "eur Bolivia", "BRST": "eur hañv Brasília", "BRT": "eur cʼhoañv Brasília", "BST Bangladesh": "eur cʼhoañv Bangladesh", "BT": "eur Bhoutan", "CAST": "CAST", "CAT": "eur Kreizafrika", "CCT": "eur Inizi Kokoz", "CDT": "eur hañv ar Cʼhreiz", "CHADT": "eur hañv Chatham", "CHAST": "eur cʼhoañv Chatham", "CHUT": "eur Chuuk", "CKT": "eur cʼhoañv Inizi Cook", "CKT DST": "eur hañv Inizi Cook", "CLST": "eur hañv Chile", "CLT": "eur cʼhoañv Chile", "COST": "eur hañv Kolombia", "COT": "eur cʼhoañv Kolombia", "CST": "eur cʼhoañv ar Cʼhreiz", "CST China": "eur cʼhoañv Sina", "CST China DST": "eur hañv Sina", "CVST": "eur hañv ar Cʼhab-Glas", "CVT": "eur cʼhoañv ar Cʼhab-Glas", "CXT": "eur Enez Christmas", "ChST": "eur Chamorro", "ChST NMI": "eur Mariana an Norzh", "CuDT": "eur hañv Kuba", "CuST": "eur cʼhoañv Kuba", "DAVT": "eur Davis", "DDUT": "eur Dumont-d’Urville", "EASST": "eur hañv Enez Pask", "EAST": "eur cʼhoañv Enez Pask", "EAT": "eur Afrika ar Reter", "ECT": "eur Ecuador", "EDT": "eur hañv ar Reter", "EGDT": "eur hañv Greunland ar Reter", "EGST": "eur cʼhoañv Greunland ar Reter", "EST": "eur cʼhoañv ar Reter", "FEET": "eur Kaliningrad", "FJT": "eur cʼhoañv Fidji", "FJT Summer": "eur hañv Fidji", "FKST": "eur hañv Inizi Falkland", "FKT": "eur cʼhoañv Inizi Falkland", "FNST": "eur hañv Fernando de Noronha", "FNT": "eur cʼhoañv Fernando de Noronha", "GALT": "eur Inizi Galápagos", "GAMT": "eur Inizi Gambier", "GEST": "eur hañv Jorjia", "GET": "eur cʼhoañv Jorjia", "GFT": "eur Gwiana cʼhall", "GIT": "eur Inizi Gilbert", "GMT": "amzer keitat Greenwich (AKG)", "GNSST": "GNSST", "GNST": "GNST", "GST": "eur cʼhoañv ar Pleg-mor Arab-ha-Pers", "GST Guam": "eur cʼhoañv Guam", "GYT": "eur Guyana", "HADT": "eur cʼhoañv Hawaii hag an Aleouted", "HAST": "eur cʼhoañv Hawaii hag an Aleouted", "HKST": "eur hañv Hong Kong", "HKT": "eur cʼhoañv Hong Kong", "HOVST": "eur hañv Khovd", "HOVT": "eur cʼhoañv Khovd", "ICT": "eur Indez-Sina", "IDT": "eur hañv Israel", "IOT": "eur Meurvor Indez", "IRKST": "eur hañv Irkutsk", "IRKT": "eur cʼhoañv Irkutsk", "IRST": "eur cʼhoañv Iran", "IRST DST": "eur hañv Iran", "IST": "eur cʼhoañv India", "IST Israel": "eur cʼhoañv Israel", "JDT": "eur hañv Japan", "JST": "eur cʼhoañv Japan", "KOST": "eur Kosrae", "KRAST": "eur hañv Krasnoyarsk", "KRAT": "eur cʼhoañv Krasnoyarsk", "KST": "eur cʼhoañv Korea", "KST DST": "eur hañv Korea", "LHDT": "eur hañv Lord Howe", "LHST": "eur cʼhoañv Lord Howe", "LINT": "eur Line Islands", "MAGST": "eur hañv Magadan", "MAGT": "eur cʼhoañv Magadan", "MART": "eur Inizi Markiz", "MAWT": "eur Mawson", "MDT": "eur hañv ar Menezioù", "MESZ": "eur hañv Kreizeuropa", "MEZ": "eur cʼhoañv Kreizeuropa", "MHT": "eur Inizi Marshall", "MMT": "eur Myanmar", "MSD": "eur hañv Moskov", "MST": "eur cʼhoañv ar Menezioù", "MUST": "eur hañv Moris", "MUT": "eur cʼhoañv Moris", "MVT": "eur ar Maldivez", "MYT": "eur Malaysia", "NCT": "eur cʼhoañv Kaledonia-Nevez", "NDT": "eur hañv Newfoundland", "NDT New Caledonia": "eur hañv Kaledonia-Nevez", "NFDT": "eur hañv Enez Norfolk", "NFT": "eur cʼhoañv Enez Norfolk", "NOVST": "eur hañv Novosibirsk", "NOVT": "eur cʼhoañv Novosibirsk", "NPT": "eur Nepal", "NRT": "eur Nauru", "NST": "eur cʼhoañv Newfoundland", "NUT": "eur Niue", "NZDT": "eur hañv Zeland-Nevez", "NZST": "eur cʼhoañv Zeland-Nevez", "OESZ": "eur hañv Europa ar Reter", "OEZ": "eur cʼhoañv Europa ar Reter", "OMSST": "eur hañv Omsk", "OMST": "eur cʼhoañv Omsk", "PDT": "eur hañv an Habask", "PDTM": "eur hañv an Habask mecʼhikan", "PETDT": "PETDT", "PETST": "PETST", "PGT": "eur Papoua Ginea-Nevez", "PHOT": "eur Inizi Phoenix", "PKT": "eur cʼhoañv Pakistan", "PKT DST": "eur hañv Pakistan", "PMDT": "eur hañv Sant-Pêr-ha-Mikelon", "PMST": "eur cʼhoañv Sant-Pêr-ha-Mikelon", "PONT": "eur Pohnpei", "PST": "eur cʼhoañv an Habask", "PST Philippine": "eur cʼhoañv ar Filipinez", "PST Philippine DST": "eur hañv ar Filipinez", "PST Pitcairn": "eur Pitcairn", "PSTM": "eur cʼhoañv an Habask mecʼhikan", "PWT": "eur Palau", "PYST": "eur hañv Paraguay", "PYT": "eur cʼhoañv Paraguay", "PYT Korea": "eur Norzhkorea", "RET": "eur ar Reünion", "ROTT": "eur Rothera", "SAKST": "eur hañv Sakhalin", "SAKT": "eur cʼhoañv Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "eur cʼhoañv Suafrika", "SBT": "eur Inizi Salomon", "SCT": "eur Sechelez", "SGT": "eur cʼhoañv Singapour", "SLST": "eur Sri Lanka", "SRT": "eur Surinam", "SST Samoa": "eur cʼhoañv Samoa Amerikan", "SST Samoa Apia": "eur cʼhoañv Samoa", "SST Samoa Apia DST": "eur hañv Samoa", "SST Samoa DST": "eur hañv Samoa Amerikan", "SYOT": "eur Syowa", "TAAF": "eur Douaroù aostral Frañs hag Antarktika", "TAHT": "eur Tahiti", "TJT": "eur Tadjikistan", "TKT": "eur Tokelau", "TLT": "eur Timor ar Reter", "TMST": "eur hañv Turkmenistan", "TMT": "eur cʼhoañv Turkmenistan", "TOST": "eur hañv Tonga", "TOT": "eur cʼhoañv Tonga", "TVT": "eur Tuvalu", "TWT": "eur cʼhoañv Taiwan", "TWT DST": "eur hañv Taiwan", "ULAST": "eur hañv Ulaanbaatar", "ULAT": "eur cʼhoañv Ulaanbaatar", "UYST": "eur hañv Uruguay", "UYT": "eur cʼhoañv Uruguay", "UZT": "eur cʼhoañv Ouzbekistan", "UZT DST": "eur hañv Ouzbekistan", "VET": "eur Venezuela", "VLAST": "eur hañv Vladivostok", "VLAT": "eur cʼhoañv Vladivostok", "VOLST": "eur hañv Volgograd", "VOLT": "eur cʼhoañv Volgograd", "VOST": "eur Vostok", "VUT": "eur cʼhoañv Vanuatu", "VUT DST": "eur hañv Vanuatu", "WAKT": "eur Enez Wake", "WARST": "eur hañv Arcʼhantina ar Cʼhornôg", "WART": "eur cʼhoañv Arcʼhantina ar Cʼhornôg", "WAST": "eur Afrika ar Cʼhornôg", "WAT": "eur Afrika ar Cʼhornôg", "WESZ": "eur hañv Europa ar Cʼhornôg", "WEZ": "eur cʼhoañv Europa ar Cʼhornôg", "WFT": "eur Wallis-ha-Futuna", "WGST": "eur hañv Greunland ar Cʼhornôg", "WGT": "eur cʼhoañv Greunland ar Cʼhornôg", "WIB": "eur Indonezia ar Cʼhornôg", "WIT": "eur Indonezia ar Reter", "WITA": "eur Kreiz Indonezia", "YAKST": "eur hañv Yakutsk", "YAKT": "eur cʼhoañv Yakutsk", "YEKST": "eur hañv Yekaterinbourg", "YEKT": "eur cʼhoañv Yekaterinbourg", "YST": "eur Yukon", "МСК": "eur cʼhoañv Moskov", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "eur cʼhoañv Almaty", "الماتى قالاسى": "eur hañv Almaty", "باتىس قازاق ەلى": "eur Kazakstan ar Cʼhornôg", "شىعىش قازاق ەلى": "eur Kazakstan ar Reter", "قازاق ەلى": "eur Kazakstan", "قىرعىزستان": "eur Kyrgyzstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "eur hañv an Azorez"},
	}
}

// Locale returns the current translators string locale
func (br *br_FR) Locale() string {
	return br.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'br_FR'
func (br *br_FR) PluralsCardinal() []locales.PluralRule {
	return br.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'br_FR'
func (br *br_FR) PluralsOrdinal() []locales.PluralRule {
	return br.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'br_FR'
func (br *br_FR) PluralsRange() []locales.PluralRule {
	return br.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'br_FR'
func (br *br_FR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)
	nMod1000000 := math.Mod(n, 1000000)

	if nMod10 == 1 && (nMod100 != 11 && nMod100 != 71 && nMod100 != 91) {
		return locales.PluralRuleOne
	} else if nMod10 == 2 && (nMod100 != 12 && nMod100 != 72 && nMod100 != 92) {
		return locales.PluralRuleTwo
	} else if nMod10 >= 3 && nMod10 <= 4 && (nMod10 == 9) && (nMod100 < 10 || nMod100 > 19) || (nMod100 < 70 || nMod100 > 79) || (nMod100 < 90 || nMod100 > 99) {
		return locales.PluralRuleFew
	} else if n != 0 && nMod1000000 == 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'br_FR'
func (br *br_FR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'br_FR'
func (br *br_FR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (br *br_FR) MonthAbbreviated(month time.Month) string {
	if len(br.monthsAbbreviated) == 0 {
		return ""
	}
	return br.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (br *br_FR) MonthsAbbreviated() []string {
	return br.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (br *br_FR) MonthNarrow(month time.Month) string {
	if len(br.monthsNarrow) == 0 {
		return ""
	}
	return br.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (br *br_FR) MonthsNarrow() []string {
	return br.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (br *br_FR) MonthWide(month time.Month) string {
	if len(br.monthsWide) == 0 {
		return ""
	}
	return br.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (br *br_FR) MonthsWide() []string {
	return br.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (br *br_FR) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(br.daysAbbreviated) == 0 {
		return ""
	}
	return br.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (br *br_FR) WeekdaysAbbreviated() []string {
	return br.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (br *br_FR) WeekdayNarrow(weekday time.Weekday) string {
	if len(br.daysNarrow) == 0 {
		return ""
	}
	return br.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (br *br_FR) WeekdaysNarrow() []string {
	return br.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (br *br_FR) WeekdayShort(weekday time.Weekday) string {
	if len(br.daysShort) == 0 {
		return ""
	}
	return br.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (br *br_FR) WeekdaysShort() []string {
	return br.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (br *br_FR) WeekdayWide(weekday time.Weekday) string {
	if len(br.daysWide) == 0 {
		return ""
	}
	return br.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (br *br_FR) WeekdaysWide() []string {
	return br.daysWide
}

// Decimal returns the decimal point of number
func (br *br_FR) Decimal() string {
	return br.decimal
}

// Group returns the group of number
func (br *br_FR) Group() string {
	return br.group
}

// Group returns the minus sign of number
func (br *br_FR) Minus() string {
	return br.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'br_FR' and handles both Whole and Real numbers based on 'v'
func (br *br_FR) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, br.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'br_FR' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (br *br_FR) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, br.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, br.percentSuffix...)

	b = append(b, br.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'br_FR'
func (br *br_FR) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := br.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, br.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, br.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, br.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'br_FR'
// in accounting notation.
func (br *br_FR) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := br.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, br.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, br.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, br.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, br.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'br_FR'
func (br *br_FR) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'br_FR'
func (br *br_FR) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, br.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'br_FR'
func (br *br_FR) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, br.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'br_FR'
func (br *br_FR) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, br.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, br.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'br_FR'
func (br *br_FR) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'br_FR'
func (br *br_FR) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, br.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'br_FR'
func (br *br_FR) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, br.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'br_FR'
func (br *br_FR) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, br.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := br.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
