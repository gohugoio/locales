package sr_Latn

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sr_Latn struct {
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

// New returns a new instance of translator for the 'sr_Latn' locale
func New() locales.Translator {
	return &sr_Latn{
		locale:                 "sr_Latn",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "KM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "r.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "ლ", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "maj", "jun", "jul", "avg", "sep", "okt", "nov", "dec"},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
		daysAbbreviated:        []string{"ned", "pon", "uto", "sre", "čet", "pet", "sub"},
		daysNarrow:             []string{"n", "p", "u", "s", "č", "p", "s"},
		daysShort:              []string{"ne", "po", "ut", "sr", "če", "pe", "su"},
		daysWide:               []string{"nedelja", "ponedeljak", "utorak", "sreda", "četvrtak", "petak", "subota"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"p. n. e.", "n. e."},
		erasNarrow:             []string{"p.n.e.", "n.e."},
		erasWide:               []string{"pre nove ere", "nove ere"},
		timezones:              map[string]string{"ACDT": "Australijsko centralno letnje vreme", "ACST": "Akre letnje računanje vremena", "ACT": "Akre standardno vreme", "ACWDT": "Australijsko centralno zapadno letnje vreme", "ACWST": "Australijsko centralno zapadno standardno vreme", "ADT": "Atlantsko letnje vreme", "ADT Arabia": "Arabijsko letnje vreme", "AEDT": "Australijsko istočno letnje vreme", "AEST": "Australijsko istočno standardno vreme", "AFT": "Avganistan vreme", "AKDT": "Aljaska, letnje vreme", "AKST": "Aljaska, standardno vreme", "AMST": "Amazon, letnje vreme", "AMST Armenia": "Jermenija, letnje vreme", "AMT": "Amazon, standardno vreme", "AMT Armenia": "Jermenija, standardno vreme", "ANAST": "Anadir letnje računanje vremena", "ANAT": "Anadir standardno vreme", "ARST": "Argentina, letnje vreme", "ART": "Argentina, standardno vreme", "AST": "Atlantsko standardno vreme", "AST Arabia": "Arabijsko standardno vreme", "AWDT": "Australijsko zapadno letnje vreme", "AWST": "Australijsko zapadno standardno vreme", "AZST": "Azerbejdžan, letnje vreme", "AZT": "Azerbejdžan, standardno vreme", "BDT Bangladesh": "Bangladeš, letnje vreme", "BNT": "Brunej Darusalum vreme", "BOT": "Bolivija vreme", "BRST": "Brazilija, letnje vreme", "BRT": "Brazilija, standardno vreme", "BST Bangladesh": "Bangladeš, standardno vreme", "BT": "Butan vreme", "CAST": "CAST", "CAT": "Centralno-afričko vreme", "CCT": "Kokos (Keling) Ostrva vreme", "CDT": "Severnoameričko centralno letnje vreme", "CHADT": "Čatam, letnje vreme", "CHAST": "Čatam, standardno vreme", "CHUT": "Čuuk vreme", "CKT": "Kukova ostrva, standardno vreme", "CKT DST": "Kukova ostrva, polu-letnje vreme", "CLST": "Čile, letnje vreme", "CLT": "Čile, standardno vreme", "COST": "Kolumbija, letnje vreme", "COT": "Kolumbija, standardno vreme", "CST": "Severnoameričko centralno standardno vreme", "CST China": "Kinesko standardno vreme", "CST China DST": "Kina, letnje vreme", "CVST": "Zelenortska Ostrva, letnje vreme", "CVT": "Zelenortska Ostrva, standardno vreme", "CXT": "Božićno ostrvo vreme", "ChST": "Čamoro vreme", "ChST NMI": "Severna Marijanska Ostrva vreme", "CuDT": "Kuba, letnje vreme", "CuST": "Kuba, standardno vreme", "DAVT": "Dejvis vreme", "DDUT": "Dimon d’Urvil vreme", "EASST": "Uskršnja ostrva, letnje vreme", "EAST": "Uskršnja ostrva, standardno vreme", "EAT": "Istočno-afričko vreme", "ECT": "Ekvador vreme", "EDT": "Severnoameričko istočno letnje vreme", "EGDT": "Istočni Grenland, letnje vreme", "EGST": "Istočni Grenland, standardno vreme", "EST": "Severnoameričko istočno standardno vreme", "FEET": "Vreme daljeg istoka Evrope", "FJT": "Fidži, standardno vreme", "FJT Summer": "Fidži, letnje vreme", "FKST": "Folklandska Ostrva, letnje vreme", "FKT": "Folklandska Ostrva, standardno vreme", "FNST": "Fernando de Noronja, letnje vreme", "FNT": "Fernando de Noronja, standardno vreme", "GALT": "Galapagos vreme", "GAMT": "Gambije vreme", "GEST": "Gruzija, letnje vreme", "GET": "Gruzija, standardno vreme", "GFT": "Francuska Gvajana vreme", "GIT": "Gilbert ostrva vreme", "GMT": "Srednje vreme po Griniču", "GNSST": "GNSST", "GNST": "GNST", "GST": "Zalivsko vreme", "GST Guam": "Guam standardno vreme", "GYT": "Gvajana vreme", "HADT": "Havajsko-aleutsko standardno vreme", "HAST": "Havajsko-aleutsko standardno vreme", "HKST": "Hong Kong, letnje vreme", "HKT": "Hong Kong, standardno vreme", "HOVST": "Hovd, letnje vreme", "HOVT": "Hovd, standardno vreme", "ICT": "Indokina vreme", "IDT": "Izraelsko letnje vreme", "IOT": "Indijsko okeansko vreme", "IRKST": "Irkuck, letnje vreme", "IRKT": "Irkuck, standardno vreme", "IRST": "Iran, standardno vreme", "IRST DST": "Iran, letnje vreme", "IST": "Indijsko standardno vreme", "IST Israel": "Izraelsko standardno vreme", "JDT": "Japansko letnje vreme", "JST": "Japansko standardno vreme", "KOST": "Košre vreme", "KRAST": "Krasnojarsk, letnje vreme", "KRAT": "Krasnojarsk, standardno vreme", "KST": "Korejsko standardno vreme", "KST DST": "Korejsko letnje vreme", "LHDT": "Lord Hov, letnje vreme", "LHST": "Lord Hov, standardno vreme", "LINT": "Ostrva Lajn vreme", "MAGST": "Magadan, letnje vreme", "MAGT": "Magadan, standardno vreme", "MART": "Markiz vreme", "MAWT": "Moson vreme", "MDT": "Makao letnje računanje vremena", "MESZ": "Srednjeevropsko letnje vreme", "MEZ": "Srednjeevropsko standardno vreme", "MHT": "Maršalska Ostrva vreme", "MMT": "Mijanmar vreme", "MSD": "Moskva, letnje vreme", "MST": "Makao standardno vreme", "MUST": "Mauricijus, letnje vreme", "MUT": "Mauricijus, standardno vreme", "MVT": "Maldivi vreme", "MYT": "Malezija vreme", "NCT": "Nova Kaledonija, standardno vreme", "NDT": "Njufaundlend, letnje vreme", "NDT New Caledonia": "Nova Kaledonija, letnje vreme", "NFDT": "Norfolk Ostrvo, letnje vreme", "NFT": "Norfolk Ostrvo, standardno vreme", "NOVST": "Novosibirsk, letnje vreme", "NOVT": "Novosibirsk, standardno vreme", "NPT": "Nepal vreme", "NRT": "Nauru vreme", "NST": "Njufaundlend, standardno vreme", "NUT": "Niue vreme", "NZDT": "Novi Zeland, letnje vreme", "NZST": "Novi Zeland, standardno vreme", "OESZ": "Istočnoevropsko letnje vreme", "OEZ": "Istočnoevropsko standardno vreme", "OMSST": "Omsk, letnje vreme", "OMST": "Omsk, standardno vreme", "PDT": "Severnoameričko pacifičko letnje vreme", "PDTM": "Meksički Pacifik, letnje vreme", "PETDT": "Petropavlovsko-kamčatsko letnje računanje vremena", "PETST": "Petropavlovsko-kamčatsko standardno vreme", "PGT": "Papua Nova Gvineja vreme", "PHOT": "Feniks ostrva vreme", "PKT": "Pakistan, standardno vreme", "PKT DST": "Pakistan, letnje vreme", "PMDT": "Sen Pjer i Mikelon, letnje vreme", "PMST": "Sen Pjer i Mikelon, standardno vreme", "PONT": "Ponpej vreme", "PST": "Severnoameričko pacifičko standardno vreme", "PST Philippine": "Filipini, standardno vreme", "PST Philippine DST": "Filipini, letnje vreme", "PST Pitcairn": "Pitkern vreme", "PSTM": "Meksički Pacifik, standardno vreme", "PWT": "Palau vreme", "PYST": "Paragvaj, letnje vreme", "PYT": "Paragvaj, standardno vreme", "PYT Korea": "Pjongjanško vreme", "RET": "Reinion vreme", "ROTT": "Rotera vreme", "SAKST": "Sahalin, letnje vreme", "SAKT": "Sahalin, standardno vreme", "SAMST": "Samara letnje računanje vremena", "SAMT": "Samara standardno vreme", "SAST": "Južno-afričko vreme", "SBT": "Solomonska Ostrva vreme", "SCT": "Sejšeli vreme", "SGT": "Singapur, standardno vreme", "SLST": "Šri Lanka vreme", "SRT": "Surinam vreme", "SST Samoa": "Samoa, standardno vreme", "SST Samoa Apia": "Apija, standardno vreme", "SST Samoa Apia DST": "Apija, letnje vreme", "SST Samoa DST": "Samoa, letnje vreme", "SYOT": "Šova vreme", "TAAF": "Francusko južno i antarktičko vreme", "TAHT": "Tahiti vreme", "TJT": "Tadžikistan vreme", "TKT": "Tokelau vreme", "TLT": "Istočni timor vreme", "TMST": "Turkmenistan, letnje vreme", "TMT": "Turkmenistan, standardno vreme", "TOST": "Tonga, letnje vreme", "TOT": "Tonga, standardno vreme", "TVT": "Tuvalu vreme", "TWT": "Tajpej, standardno vreme", "TWT DST": "Tajpej, letnje vreme", "ULAST": "Ulan Bator, letnje vreme", "ULAT": "Ulan Bator, standardno vreme", "UYST": "Urugvaj, letnje vreme", "UYT": "Urugvaj, standardno vreme", "UZT": "Uzbekistan, standardno vreme", "UZT DST": "Uzbekistan, letnje vreme", "VET": "Venecuela vreme", "VLAST": "Vladivostok, letnje vreme", "VLAT": "Vladivostok, standardno vreme", "VOLST": "Volgograd, letnje vreme", "VOLT": "Volgograd, standardno vreme", "VOST": "Vostok vreme", "VUT": "Vanuatu, standardno vreme", "VUT DST": "Vanuatu, letnje vreme", "WAKT": "Vejk ostrvo vreme", "WARST": "Zapadna Argentina, letnje vreme", "WART": "Zapadna Argentina, standardno vreme", "WAST": "Zapadno-afričko vreme", "WAT": "Zapadno-afričko vreme", "WESZ": "Zapadnoevropsko letnje vreme", "WEZ": "Zapadnoevropsko standardno vreme", "WFT": "Valis i Futuna Ostrva vreme", "WGST": "Zapadni Grenland, letnje vreme", "WGT": "Zapadni Grenland, standardno vreme", "WIB": "Zapadno-indonezijsko vreme", "WIT": "Istočno-indonezijsko vreme", "WITA": "Centralno-indonezijsko vreme", "YAKST": "Jakutsk, letnje vreme", "YAKT": "Jakutsk, standardno vreme", "YEKST": "Jekaterinburg, letnje vreme", "YEKT": "Jekaterinburg, standardno vreme", "YST": "Jukon", "МСК": "Moskva, standardno vreme", "اقتاۋ": "Akvatau standardno vreme", "اقتاۋ قالاسى": "Akvatau letnje računanje vremena", "اقتوبە": "Akutobe standardno vreme", "اقتوبە قالاسى": "Akutobe letnje računanje vremena", "الماتى": "Almati standardno vreme", "الماتى قالاسى": "Almati letnje računanje vremena", "باتىس قازاق ەلى": "Zapadno-kazahstansko vreme", "شىعىش قازاق ەلى": "Istočno-kazahstansko vreme", "قازاق ەلى": "Kazahstansko vreme", "قىرعىزستان": "Kirgistan vreme", "قىزىلوردا": "Kizilorda standardno vreme", "قىزىلوردا قالاسى": "Kizilorda letnje računanje vremena", "∅∅∅": "Peru, letnje vreme"},
	}
}

// Locale returns the current translators string locale
func (sr *sr_Latn) Locale() string {
	return sr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sr_Latn'
func (sr *sr_Latn) PluralsCardinal() []locales.PluralRule {
	return sr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sr_Latn'
func (sr *sr_Latn) PluralsOrdinal() []locales.PluralRule {
	return sr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sr_Latn'
func (sr *sr_Latn) PluralsRange() []locales.PluralRule {
	return sr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Latn'
func (sr *sr_Latn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14)) || (fMod10 >= 2 && fMod10 <= 4 && (fMod100 < 12 || fMod100 > 14)) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Latn'
func (sr *sr_Latn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sr_Latn'
func (sr *sr_Latn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sr.CardinalPluralRule(num1, v1)
	end := sr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sr *sr_Latn) MonthAbbreviated(month time.Month) string {
	return sr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sr *sr_Latn) MonthsAbbreviated() []string {
	return sr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sr *sr_Latn) MonthNarrow(month time.Month) string {
	return sr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sr *sr_Latn) MonthsNarrow() []string {
	return sr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sr *sr_Latn) MonthWide(month time.Month) string {
	return sr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sr *sr_Latn) MonthsWide() []string {
	return sr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sr *sr_Latn) WeekdayAbbreviated(weekday time.Weekday) string {
	return sr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sr *sr_Latn) WeekdaysAbbreviated() []string {
	return sr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sr *sr_Latn) WeekdayNarrow(weekday time.Weekday) string {
	return sr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sr *sr_Latn) WeekdaysNarrow() []string {
	return sr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sr *sr_Latn) WeekdayShort(weekday time.Weekday) string {
	return sr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sr *sr_Latn) WeekdaysShort() []string {
	return sr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sr *sr_Latn) WeekdayWide(weekday time.Weekday) string {
	return sr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sr *sr_Latn) WeekdaysWide() []string {
	return sr.daysWide
}

// Decimal returns the decimal point of number
func (sr *sr_Latn) Decimal() string {
	return sr.decimal
}

// Group returns the group of number
func (sr *sr_Latn) Group() string {
	return sr.group
}

// Group returns the minus sign of number
func (sr *sr_Latn) Minus() string {
	return sr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sr_Latn' and handles both Whole and Real numbers based on 'v'
func (sr *sr_Latn) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sr_Latn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sr *sr_Latn) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Latn'
func (sr *sr_Latn) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Latn'
// in accounting notation.
func (sr *sr_Latn) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sr.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, sr.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sr_Latn'
func (sr *sr_Latn) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
