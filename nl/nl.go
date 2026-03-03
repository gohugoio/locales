package nl

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type nl struct {
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
	currencyPositivePrefix string
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

// New returns a new instance of translator for the 'nl' locale
func New() locales.Translator {
	return &nl{
		locale:                 "nl",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "C$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJ$", "FKP", "FRF", "GBP", "GEK", "ლ", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "р.", "RWF", "SAR", "SI$", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "Cg", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZiG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "jan", "feb", "mrt", "apr", "mei", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
		daysAbbreviated:        []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		daysNarrow:             []string{"Z", "M", "D", "W", "D", "V", "Z"},
		daysWide:               []string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
		timezones:              map[string]string{"ACDT": "Midden-Australische zomertijd", "ACST": "Midden-Australische standaardtijd", "ACT": "Acre-standaardtijd", "ACWDT": "Midden-Australische westelijke zomertijd", "ACWST": "Midden-Australische westelijke standaardtijd", "ADT": "Atlantic-zomertijd", "ADT Arabia": "Arabische zomertijd", "AEDT": "Oost-Australische zomertijd", "AEST": "Oost-Australische standaardtijd", "AFT": "Afghaanse tijd", "AKDT": "Alaska-zomertijd", "AKST": "Alaska-standaardtijd", "AMST": "Amazone-zomertijd", "AMST Armenia": "Armeense zomertijd", "AMT": "Amazone-standaardtijd", "AMT Armenia": "Armeense standaardtijd", "ANAST": "Anadyr-zomertijd", "ANAT": "Anadyr-standaardtijd", "ARST": "Argentijnse zomertijd", "ART": "Argentijnse standaardtijd", "AST": "Atlantic-standaardtijd", "AST Arabia": "Arabische standaardtijd", "AWDT": "West-Australische zomertijd", "AWST": "West-Australische standaardtijd", "AZST": "Azerbeidzjaanse zomertijd", "AZT": "Azerbeidzjaanse standaardtijd", "BDT Bangladesh": "Bengalese zomertijd", "BNT": "Bruneise tijd", "BOT": "Boliviaanse tijd", "BRST": "Braziliaanse zomertijd", "BRT": "Braziliaanse standaardtijd", "BST Bangladesh": "Bengalese standaardtijd", "BT": "Bhutaanse tijd", "CAST": "Casey tijd", "CAT": "Centraal-Afrikaanse tijd", "CCT": "Cocoseilandse tijd", "CDT": "Central-zomertijd", "CHADT": "Chatham-zomertijd", "CHAST": "Chatham-standaardtijd", "CHUT": "Chuukse tijd", "CKT": "Cookeilandse standaardtijd", "CKT DST": "Cookeilandse halve zomertijd", "CLST": "Chileense zomertijd", "CLT": "Chileense standaardtijd", "COST": "Colombiaanse zomertijd", "COT": "Colombiaanse standaardtijd", "CST": "Central-standaardtijd", "CST China": "Chinese standaardtijd", "CST China DST": "Chinese zomertijd", "CVST": "Kaapverdische zomertijd", "CVT": "Kaapverdische standaardtijd", "CXT": "Christmaseilandse tijd", "ChST": "Chamorro-tijd", "ChST NMI": "Noordelijk Mariaanse tijd", "CuDT": "Cubaanse zomertijd", "CuST": "Cubaanse standaardtijd", "DAVT": "Davis-tijd", "DDUT": "Dumont-d’Urville-tijd", "EASST": "Paaseilandse zomertijd", "EAST": "Paaseilandse standaardtijd", "EAT": "Oost-Afrikaanse tijd", "ECT": "Ecuadoraanse tijd", "EDT": "Eastern-zomertijd", "EGDT": "Oost-Groenlandse zomertijd", "EGST": "Oost-Groenlandse standaardtijd", "EST": "Eastern-standaardtijd", "FEET": "Verder-oostelijk-Europese tijd", "FJT": "Fijische standaardtijd", "FJT Summer": "Fijische zomertijd", "FKST": "Falklandeilandse zomertijd", "FKT": "Falklandeilandse standaardtijd", "FNST": "Fernando de Noronha-zomertijd", "FNT": "Fernando de Noronha-standaardtijd", "GALT": "Galapagoseilandse tijd", "GAMT": "Gambiereilandse tijd", "GEST": "Georgische zomertijd", "GET": "Georgische standaardtijd", "GFT": "Frans-Guyaanse tijd", "GIT": "Gilberteilandse tijd", "GMT": "Greenwich Mean Time", "GNSST": "Groenlandse zomertijd", "GNST": "Groenlandse standaardtijd", "GST": "Zuid-Georgische tijd", "GST Guam": "Guamese standaardtijd", "GYT": "Guyaanse tijd", "HADT": "Hawaii-Aleoetische standaardtijd", "HAST": "Hawaii-Aleoetische standaardtijd", "HKST": "Hongkongse zomertijd", "HKT": "Hongkongse standaardtijd", "HOVST": "Hovd-zomertijd", "HOVT": "Hovd-standaardtijd", "ICT": "Indochinese tijd", "IDT": "Israëlische zomertijd", "IOT": "Indische Oceaan-tijd", "IRKST": "Irkoetsk-zomertijd", "IRKT": "Irkoetsk-standaardtijd", "IRST": "Iraanse standaardtijd", "IRST DST": "Iraanse zomertijd", "IST": "Indiase tijd", "IST Israel": "Israëlische standaardtijd", "JDT": "Japanse zomertijd", "JST": "Japanse standaardtijd", "KOST": "Kosraese tijd", "KRAST": "Krasnojarsk-zomertijd", "KRAT": "Krasnojarsk-standaardtijd", "KST": "Koreaanse standaardtijd", "KST DST": "Koreaanse zomertijd", "LHDT": "Lord Howe-eilandse zomertijd", "LHST": "Lord Howe-eilandse standaardtijd", "LINT": "Line-eilandse tijd", "MAGST": "Magadan-zomertijd", "MAGT": "Magadan-standaardtijd", "MART": "Marquesaseilandse tijd", "MAWT": "Mawson-tijd", "MDT": "Mountain-zomertijd", "MESZ": "Midden-Europese zomertijd", "MEZ": "Midden-Europese standaardtijd", "MHT": "Marshalleilandse tijd", "MMT": "Myanmarese tijd", "MSD": "Moskou-zomertijd", "MST": "Mountain-standaardtijd", "MUST": "Mauritiaanse zomertijd", "MUT": "Mauritiaanse standaardtijd", "MVT": "Maldivische tijd", "MYT": "Maleisische tijd", "NCT": "Nieuw-Caledonische standaardtijd", "NDT": "Newfoundland-zomertijd", "NDT New Caledonia": "Nieuw-Caledonische zomertijd", "NFDT": "Norfolkeilandse zomertijd", "NFT": "Norfolkeilandse standaardtijd", "NOVST": "Novosibirsk-zomertijd", "NOVT": "Novosibirsk-standaardtijd", "NPT": "Nepalese tijd", "NRT": "Nauruaanse tijd", "NST": "Newfoundland-standaardtijd", "NUT": "Niuese tijd", "NZDT": "Nieuw-Zeelandse zomertijd", "NZST": "Nieuw-Zeelandse standaardtijd", "OESZ": "Oost-Europese zomertijd", "OEZ": "Oost-Europese standaardtijd", "OMSST": "Omsk-zomertijd", "OMST": "Omsk-standaardtijd", "PDT": "Pacific-zomertijd", "PDTM": "Mexicaanse Pacific-zomertijd", "PETDT": "Petropavlovsk-Kamtsjatski-zomertijd", "PETST": "Petropavlovsk-Kamtsjatski-standaardtijd", "PGT": "Papoea-Nieuw-Guineese tijd", "PHOT": "Phoenixeilandse tijd", "PKT": "Pakistaanse standaardtijd", "PKT DST": "Pakistaanse zomertijd", "PMDT": "Saint Pierre en Miquelon-zomertijd", "PMST": "Saint Pierre en Miquelon-standaardtijd", "PONT": "Pohnpei-tijd", "PST": "Pacific-standaardtijd", "PST Philippine": "Filipijnse standaardtijd", "PST Philippine DST": "Filipijnse zomertijd", "PST Pitcairn": "Pitcairneilandse tijd", "PSTM": "Mexicaanse Pacific-standaardtijd", "PWT": "Belause tijd", "PYST": "Paraguayaanse zomertijd", "PYT": "Paraguayaanse standaardtijd", "PYT Korea": "Pyongyang-tijd", "RET": "Réunionse tijd", "ROTT": "Rothera-tijd", "SAKST": "Sachalin-zomertijd", "SAKT": "Sachalin-standaardtijd", "SAMST": "Samara-zomertijd", "SAMT": "Samara-standaardtijd", "SAST": "Zuid-Afrikaanse tijd", "SBT": "Salomonseilandse tijd", "SCT": "Seychelse tijd", "SGT": "Singaporese standaardtijd", "SLST": "Lanka-tijd", "SRT": "Surinaamse tijd", "SST Samoa": "Samoaanse standaardtijd", "SST Samoa Apia": "Apia-standaardtijd", "SST Samoa Apia DST": "Apia-zomertijd", "SST Samoa DST": "Samoaanse zomertijd", "SYOT": "Syowa-tijd", "TAAF": "Franse zuidelijke en Antarctische tijd", "TAHT": "Tahitiaanse tijd", "TJT": "Tadzjiekse tijd", "TKT": "Tokelau-eilandse tijd", "TLT": "Oost-Timorese tijd", "TMST": "Turkmeense zomertijd", "TMT": "Turkmeense standaardtijd", "TOST": "Tongaanse zomertijd", "TOT": "Tongaanse standaardtijd", "TVT": "Tuvaluaanse tijd", "TWT": "Taipei-standaardtijd", "TWT DST": "Taipei-zomertijd", "ULAST": "Ulaanbaatar-zomertijd", "ULAT": "Ulaanbaatar-standaardtijd", "UYST": "Uruguayaanse zomertijd", "UYT": "Uruguayaanse standaardtijd", "UZT": "Oezbeekse standaardtijd", "UZT DST": "Oezbeekse zomertijd", "VET": "Venezolaanse tijd", "VLAST": "Vladivostok-zomertijd", "VLAT": "Vladivostok-standaardtijd", "VOLST": "Wolgograd-zomertijd", "VOLT": "Wolgograd-standaardtijd", "VOST": "Vostok-tijd", "VUT": "Vanuatuaanse standaardtijd", "VUT DST": "Vanuatuaanse zomertijd", "WAKT": "Wake-eilandse tijd", "WARST": "West-Argentijnse zomertijd", "WART": "West-Argentijnse standaardtijd", "WAST": "West-Afrikaanse tijd", "WAT": "West-Afrikaanse tijd", "WESZ": "West-Europese zomertijd", "WEZ": "West-Europese standaardtijd", "WFT": "Wallis en Futunase tijd", "WGST": "West-Groenlandse zomertijd", "WGT": "West-Groenlandse standaardtijd", "WIB": "West-Indonesische tijd", "WIT": "Oost-Indonesische tijd", "WITA": "Centraal-Indonesische tijd", "YAKST": "Jakoetsk-zomertijd", "YAKT": "Jakoetsk-standaardtijd", "YEKST": "Jekaterinenburg-zomertijd", "YEKT": "Jekaterinenburg-standaardtijd", "YST": "Yukon-tijd", "МСК": "Moskou-standaardtijd", "اقتاۋ": "Aqtau-standaardtijd", "اقتاۋ قالاسى": "Aqtau-zomertijd", "اقتوبە": "Aqtöbe-standaardtijd", "اقتوبە قالاسى": "Aqtöbe-zomertijd", "الماتى": "Alma-Ata-standaardtijd", "الماتى قالاسى": "Alma-Ata-zomertijd", "باتىس قازاق ەلى": "West-Kazachse tijd", "شىعىش قازاق ەلى": "Oost-Kazachse tijd", "قازاق ەلى": "Kazachse tijd", "قىرعىزستان": "Kirgizische tijd", "قىزىلوردا": "Qyzylorda-standaardtijd", "قىزىلوردا قالاسى": "Qyzylorda-zomertijd", "∅∅∅": "Peruaanse zomertijd"},
	}
}

// Locale returns the current translators string locale
func (nl *nl) Locale() string {
	return nl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nl'
func (nl *nl) PluralsCardinal() []locales.PluralRule {
	return nl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nl'
func (nl *nl) PluralsOrdinal() []locales.PluralRule {
	return nl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nl'
func (nl *nl) PluralsRange() []locales.PluralRule {
	return nl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nl'
func (nl *nl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nl'
func (nl *nl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nl'
func (nl *nl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := nl.CardinalPluralRule(num1, v1)
	end := nl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nl *nl) MonthAbbreviated(month time.Month) string {
	if len(nl.monthsAbbreviated) == 0 {
		return ""
	}
	return nl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nl *nl) MonthsAbbreviated() []string {
	return nl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nl *nl) MonthNarrow(month time.Month) string {
	if len(nl.monthsNarrow) == 0 {
		return ""
	}
	return nl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nl *nl) MonthsNarrow() []string {
	return nl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nl *nl) MonthWide(month time.Month) string {
	if len(nl.monthsWide) == 0 {
		return ""
	}
	return nl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nl *nl) MonthsWide() []string {
	return nl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nl *nl) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(nl.daysAbbreviated) == 0 {
		return ""
	}
	return nl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nl *nl) WeekdaysAbbreviated() []string {
	return nl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nl *nl) WeekdayNarrow(weekday time.Weekday) string {
	if len(nl.daysNarrow) == 0 {
		return ""
	}
	return nl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nl *nl) WeekdaysNarrow() []string {
	return nl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nl *nl) WeekdayShort(weekday time.Weekday) string {
	if len(nl.daysShort) == 0 {
		return ""
	}
	return nl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nl *nl) WeekdaysShort() []string {
	return nl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nl *nl) WeekdayWide(weekday time.Weekday) string {
	if len(nl.daysWide) == 0 {
		return ""
	}
	return nl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nl *nl) WeekdaysWide() []string {
	return nl.daysWide
}

// Decimal returns the decimal point of number
func (nl *nl) Decimal() string {
	return nl.decimal
}

// Group returns the group of number
func (nl *nl) Group() string {
	return nl.group
}

// Group returns the minus sign of number
func (nl *nl) Minus() string {
	return nl.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nl' and handles both Whole and Real numbers based on 'v'
func (nl *nl) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nl' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nl *nl) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nl.percentSuffix...)

	b = append(b, nl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nl'
func (nl *nl) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(nl.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, nl.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nl'
// in accounting notation.
func (nl *nl) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nl.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
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

		for j := len(nl.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, nl.currencyNegativePrefix[j])
		}

	} else {

		for j := len(nl.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, nl.currencyPositivePrefix[j])
		}

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
			b = append(b, nl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, nl.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'nl'
func (nl *nl) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'nl'
func (nl *nl) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nl'
func (nl *nl) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nl'
func (nl *nl) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nl'
func (nl *nl) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nl'
func (nl *nl) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nl'
func (nl *nl) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nl'
func (nl *nl) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
