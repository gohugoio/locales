package sr_Latn_BA

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sr_Latn_BA struct {
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

// New returns a new instance of translator for the 'sr_Latn_BA' locale
func New() locales.Translator {
	return &sr_Latn_BA{
		locale:                 "sr_Latn_BA",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "јан", "феб", "мар", "апр", "мај", "јун", "јул", "авг", "сеп", "окт", "нов", "дец"},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
		daysAbbreviated:        []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysNarrow:             []string{"н", "п", "у", "с", "ч", "п", "с"},
		daysShort:              []string{"не", "по", "ут", "ср", "че", "пе", "су"},
		daysWide:               []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
		periodsAbbreviated:     []string{"prije\u202fpodne", "po\u202fpodne"},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"prije podne", "po podne"},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"prije nove ere", "nove ere"},
		timezones:              map[string]string{"ACDT": "Australijsko centralno ljetnje vrijeme", "ACST": "Australijsko centralno standardno vrijeme", "ACT": "Акре стандардно време", "ACWDT": "Australijsko centralno zapadno ljetnje vrijeme", "ACWST": "Australijsko centralno zapadno standardno vrijeme", "ADT": "Atlantsko ljetnje vrijeme", "ADT Arabia": "Arabijsko ljetnje vrijeme", "AEDT": "Australijsko istočno ljetnje vrijeme", "AEST": "Australijsko istočno standardno vrijeme", "AFT": "Avganistan vrijeme", "AKDT": "Aljaska, ljetnje vrijeme", "AKST": "Aljaska, standardno vrijeme", "AMST": "Amazon, ljetnje vrijeme", "AMST Armenia": "Jermenija, ljetnje vrijeme", "AMT": "Amazon, standardno vrijeme", "AMT Armenia": "Jermenija, standardno vrijeme", "ANAST": "Анадир летње рачунање времена", "ANAT": "Анадир стандардно време", "ARST": "Argentina, ljetnje vrijeme", "ART": "Argentina, standardno vrijeme", "AST": "Atlantsko standardno vrijeme", "AST Arabia": "Arabijsko standardno vrijeme", "AWDT": "Australijsko zapadno ljetnje vrijeme", "AWST": "Australijsko zapadno standardno vrijeme", "AZST": "Azerbejdžan, ljetnje vrijeme", "AZT": "Azerbejdžan, standardno vrijeme", "BDT Bangladesh": "Bangladeš, ljetnje vrijeme", "BNT": "Brunej Darusalum vrijeme", "BOT": "Bolivija vrijeme", "BRST": "Brazilija, ljetnje vrijeme", "BRT": "Brazilija, standardno vrijeme", "BST Bangladesh": "Bangladeš, standardno vrijeme", "BT": "Butan vrijeme", "CAST": "CAST", "CAT": "Centralno-afričko vrijeme", "CCT": "Kokosova (Kiling) ostrva vrijeme", "CDT": "Sjevernoameričko centralno ljetnje vrijeme", "CHADT": "Čatam, ljetnje vrijeme", "CHAST": "Čatam, standardno vrijeme", "CHUT": "Čuk vrijeme", "CKT": "Kukova Ostrva, standardno vrijeme", "CKT DST": "Kukova Ostrva, poluljetnje vrijeme", "CLST": "Čile, ljetnje vrijeme", "CLT": "Čile, standardno vrijeme", "COST": "Kolumbija, ljetnje vrijeme", "COT": "Kolumbija, standardno vrijeme", "CST": "Sjevernoameričko centralno standardno vrijeme", "CST China": "Kinesko standardno vrijeme", "CST China DST": "Kina, ljetnje vrijeme", "CVST": "Zelenortska Ostrva, ljetnje vrijeme", "CVT": "Zelenortska Ostrva, standardno vrijeme", "CXT": "Božićno ostrvo vrijeme", "ChST": "Čamoro vrijeme", "ChST NMI": "Северна Маријанска Острва време", "CuDT": "Kuba, ljetnje vrijeme", "CuST": "Kuba, standardno vrijeme", "DAVT": "Dejvis vrijeme", "DDUT": "Dimon d’Irvil vrijeme", "EASST": "Uskršnja ostrva, ljetnje vrijeme", "EAST": "Uskršnja ostrva, standardno vrijeme", "EAT": "Istočno-afričko vrijeme", "ECT": "Ekvador vrijeme", "EDT": "Sjevernoameričko istočno ljetnje vrijeme", "EGDT": "Istočni Grenland, ljetnje vrijeme", "EGST": "Istočni Grenland, standardno vrijeme", "EST": "Sjevernoameričko istočno standardno vrijeme", "FEET": "Време даљег истока Европе", "FJT": "Fidži, standardno vrijeme", "FJT Summer": "Fidži, ljetnje vrijeme", "FKST": "Folklandska Ostrva, ljetnje vrijeme", "FKT": "Folklandska Ostrva, standardno vrijeme", "FNST": "Fernando de Noronja, ljetnje vrijeme", "FNT": "Fernando de Noronja, standardno vrijeme", "GALT": "Galapagos vrijeme", "GAMT": "Gambije vrijeme", "GEST": "Gruzija, ljetnje vrijeme", "GET": "Gruzija, standardno vrijeme", "GFT": "Francuska Gvajana vrijeme", "GIT": "Gilbertova ostrva vrijeme", "GMT": "Srednje vrijeme po Griniču", "GNSST": "GNSST", "GNST": "GNST", "GST": "Zalivsko vrijeme", "GST Guam": "Гуам стандардно време", "GYT": "Gvajana vrijeme", "HADT": "Havajsko-aleutsko standardno vrijeme", "HAST": "Havajsko-aleutsko standardno vrijeme", "HKST": "Hong Kong, ljetnje vrijeme", "HKT": "Hong Kong, standardno vrijeme", "HOVST": "Hovd, ljetnje vrijeme", "HOVT": "Hovd, standardno vrijeme", "ICT": "Indokina vrijeme", "IDT": "Izraelsko ljetnje vrijeme", "IOT": "Indijsko okeansko vrijeme", "IRKST": "Irkuck, ljetnje vrijeme", "IRKT": "Irkuck, standardno vrijeme", "IRST": "Iran, standardno vrijeme", "IRST DST": "Iran, ljetnje vrijeme", "IST": "Indijsko standardno vrijeme", "IST Israel": "Izraelsko standardno vrijeme", "JDT": "Japansko ljetnje vrijeme", "JST": "Japansko standardno vrijeme", "KOST": "Košre vrijeme", "KRAST": "Krasnojarsk, ljetnje vrijeme", "KRAT": "Krasnojarsk, standardno vrijeme", "KST": "Korejsko standardno vrijeme", "KST DST": "Korejsko ljetnje vrijeme", "LHDT": "Lord Hov, ljetnje vrijeme", "LHST": "Lord Hov, standardno vrijeme", "LINT": "Linijska ostrva vrijeme", "MAGST": "Magadan, ljetnje vrijeme", "MAGT": "Magadan, standardno vrijeme", "MART": "Markiz vrijeme", "MAWT": "Moson vrijeme", "MDT": "Макао летње рачунање времена", "MESZ": "Srednjoevropsko ljetnje vrijeme", "MEZ": "Srednjoevropsko standardno vrijeme", "MHT": "Maršalska Ostrva vrijeme", "MMT": "Mjanmar vrijeme", "MSD": "Moskva, ljetnje vrijeme", "MST": "Макао стандардно време", "MUST": "Mauricijus, ljetnje vrijeme", "MUT": "Mauricijus, standardno vrijeme", "MVT": "Maldivi vrijeme", "MYT": "Malezija vrijeme", "NCT": "Nova Kaledonija, standardno vrijeme", "NDT": "Njufaundlend, ljetnje vrijeme", "NDT New Caledonia": "Nova Kaledonija, ljetnje vrijeme", "NFDT": "ostrvo Norfolk, ljetnje vrijeme", "NFT": "ostrvo Norfolk, standardno vrijeme", "NOVST": "Novosibirsk, ljetnje vrijeme", "NOVT": "Novosibirsk, standardno vrijeme", "NPT": "Nepal vrijeme", "NRT": "Nauru vrijeme", "NST": "Njufaundlend, standardno vrijeme", "NUT": "Nijue vrijeme", "NZDT": "Novi Zeland, ljetnje vrijeme", "NZST": "Novi Zeland, standardno vrijeme", "OESZ": "Istočnoevropsko ljetnje vrijeme", "OEZ": "Istočnoevropsko standardno vrijeme", "OMSST": "Omsk, ljetnje vrijeme", "OMST": "Omsk, standardno vrijeme", "PDT": "Sjevernoameričko pacifičko letnje vrijeme", "PDTM": "Meksički Pacifik, ljetnje vrijeme", "PETDT": "Петропавловско-камчатско летње рачунање времена", "PETST": "Петропавловско-камчатско стандардно време", "PGT": "Papua Nova Gvineja vrijeme", "PHOT": "Feniks ostrva vrijeme", "PKT": "Pakistan, standardno vrijeme", "PKT DST": "Pakistan, ljetnje vrijeme", "PMDT": "Sen Pjer i Mikelon, ljetnje vrijeme", "PMST": "Sen Pjer i Mikelon, standardno vrijeme", "PONT": "Ponpej vrijeme", "PST": "Sjevernoameričko pacifičko standardno vrijeme", "PST Philippine": "Filipini, standardno vrijeme", "PST Philippine DST": "Filipini, ljetnje vrijeme", "PST Pitcairn": "Pitkern vrijeme", "PSTM": "Meksički Pacifik, standardno vrijeme", "PWT": "Palau vrijeme", "PYST": "Paragvaj, ljetnje vrijeme", "PYT": "Paragvaj, standardno vrijeme", "PYT Korea": "Pjongjanško vrijeme", "RET": "Reunion vrijeme", "ROTT": "Rotera vrijeme", "SAKST": "Sahalin, ljetnje vrijeme", "SAKT": "Sahalin, standardno vrijeme", "SAMST": "Самара летње рачунање времена", "SAMT": "Самара стандардно време", "SAST": "Južno-afričko vrijeme", "SBT": "Solomonska Ostrva vrijeme", "SCT": "Sejšeli vrijeme", "SGT": "Singapur, standardno vrijeme", "SLST": "Шри Ланка време", "SRT": "Surinam vrijeme", "SST Samoa": "Samoa, standardno vrijeme", "SST Samoa Apia": "Apija, standardno vrijeme", "SST Samoa Apia DST": "Apija, ljetnje vrijeme", "SST Samoa DST": "Samoa, ljetnje vrijeme", "SYOT": "Šova vrijeme", "TAAF": "Francusko južno i antarktičko vrijeme", "TAHT": "Tahiti vrijeme", "TJT": "Tadžikistan vrijeme", "TKT": "Tokelau vrijeme", "TLT": "Istočni Timor vrijeme", "TMST": "Turkmenistan, ljetnje vrijeme", "TMT": "Turkmenistan, standardno vrijeme", "TOST": "Tonga, ljetnje vrijeme", "TOT": "Tonga, standardno vrijeme", "TVT": "Tuvalu vrijeme", "TWT": "Tajpej, standardno vrijeme", "TWT DST": "Tajpej, ljetnje vrijeme", "ULAST": "Ulan Bator, ljetnje vrijeeme", "ULAT": "Ulan Bator, standardno vrijeme", "UYST": "Urugvaj, ljetnje vrijeme", "UYT": "Urugvaj, standardno vrijeme", "UZT": "Uzbekistan, standardno vrijeme", "UZT DST": "Uzbekistan, ljetnje vrijeme", "VET": "Venecuela vrijeme", "VLAST": "Vladivostok, ljetnje vrijeme", "VLAT": "Vladivostok, standardno vrijeme", "VOLST": "Volgograd, ljetnje vrijeme", "VOLT": "Volgograd, standardno vrijeme", "VOST": "Vostok vrijeme", "VUT": "Vanuatu, standardno vrijeme", "VUT DST": "Vanuatu, ljetnje vrijeme", "WAKT": "ostrvo Vejk vrijeme", "WARST": "Zapadna Argentina, ljetnje vrijeme", "WART": "Zapadna Argentina, standardno vrijeme", "WAST": "Zapadno-afričko vrijeme", "WAT": "Zapadno-afričko vrijeme", "WESZ": "Zapadnoevropsko ljetnje vrijeme", "WEZ": "Zapadnoevropsko standardno vrijeme", "WFT": "ostrva Valis i Futuna vrijeme", "WGST": "Zapadni Grenland, ljetnje vrijeme", "WGT": "Zapadni Grenland, standardno vrijeme", "WIB": "Zapadno-indonezijsko vrijeme", "WIT": "Istočno-indonezijsko vrijeme", "WITA": "Centralno-indonezijsko vrijeme", "YAKST": "Jakutsk, ljetnje vrijeme", "YAKT": "Jakutsk, standardno vrijeme", "YEKST": "Jekaterinburg, ljetnje vrijeme", "YEKT": "Jekaterinburg, standardno vrijeme", "YST": "Јукон", "МСК": "Moskva, standardno vrijeme", "اقتاۋ": "Акватау стандардно време", "اقتاۋ قالاسى": "Акватау летње рачунање времена", "اقتوبە": "Акутобе стандардно време", "اقتوبە قالاسى": "Акутобе летње рачунање времена", "الماتى": "Алмати стандардно време", "الماتى قالاسى": "Алмати летње рачунање времена", "باتىس قازاق ەلى": "Zapadno-kazahstansko vrijeme", "شىعىش قازاق ەلى": "Istočno-kazahstansko vrijeme", "قازاق ەلى": "Kazahstansko vrijeme", "قىرعىزستان": "Kirgistan vrijeme", "قىزىلوردا": "Кизилорда стандардно време", "قىزىلوردا قالاسى": "Кизилорда летње рачунање времена", "∅∅∅": "Peru, ljetnje vrijeme"},
	}
}

// Locale returns the current translators string locale
func (sr *sr_Latn_BA) Locale() string {
	return sr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sr_Latn_BA'
func (sr *sr_Latn_BA) PluralsCardinal() []locales.PluralRule {
	return sr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sr_Latn_BA'
func (sr *sr_Latn_BA) PluralsOrdinal() []locales.PluralRule {
	return sr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sr_Latn_BA'
func (sr *sr_Latn_BA) PluralsRange() []locales.PluralRule {
	return sr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

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
func (sr *sr_Latn_BA) MonthAbbreviated(month time.Month) string {
	return sr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sr *sr_Latn_BA) MonthsAbbreviated() []string {
	return sr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sr *sr_Latn_BA) MonthNarrow(month time.Month) string {
	return sr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sr *sr_Latn_BA) MonthsNarrow() []string {
	return sr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sr *sr_Latn_BA) MonthWide(month time.Month) string {
	return sr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sr *sr_Latn_BA) MonthsWide() []string {
	return sr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sr *sr_Latn_BA) WeekdayAbbreviated(weekday time.Weekday) string {
	return sr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sr *sr_Latn_BA) WeekdaysAbbreviated() []string {
	return sr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sr *sr_Latn_BA) WeekdayNarrow(weekday time.Weekday) string {
	return sr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sr *sr_Latn_BA) WeekdaysNarrow() []string {
	return sr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sr *sr_Latn_BA) WeekdayShort(weekday time.Weekday) string {
	return sr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sr *sr_Latn_BA) WeekdaysShort() []string {
	return sr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sr *sr_Latn_BA) WeekdayWide(weekday time.Weekday) string {
	return sr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sr *sr_Latn_BA) WeekdaysWide() []string {
	return sr.daysWide
}

// Decimal returns the decimal point of number
func (sr *sr_Latn_BA) Decimal() string {
	return sr.decimal
}

// Group returns the group of number
func (sr *sr_Latn_BA) Group() string {
	return sr.group
}

// Group returns the minus sign of number
func (sr *sr_Latn_BA) Minus() string {
	return sr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sr_Latn_BA' and handles both Whole and Real numbers based on 'v'
func (sr *sr_Latn_BA) FmtNumber(num float64, v uint64) string {

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

// FmtPercent returns 'num' with digits/precision of 'v' for 'sr_Latn_BA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sr *sr_Latn_BA) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtCurrency(num float64, v uint64, currency currency.Type) string {

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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Latn_BA'
// in accounting notation.
func (sr *sr_Latn_BA) FmtAccounting(num float64, v uint64, currency currency.Type) string {

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

// FmtDateShort returns the short date representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtDateLong(t time.Time) string {

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

// FmtDateFull returns the full date representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtDateFull(t time.Time) string {

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

// FmtTimeShort returns the short time representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtTimeShort(t time.Time) string {

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

// FmtTimeMedium returns the medium time representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtTimeMedium(t time.Time) string {

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

// FmtTimeLong returns the long time representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtTimeLong(t time.Time) string {

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

// FmtTimeFull returns the full time representation of 't' for 'sr_Latn_BA'
func (sr *sr_Latn_BA) FmtTimeFull(t time.Time) string {

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
