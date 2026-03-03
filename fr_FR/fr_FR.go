package fr_FR

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type fr_FR struct {
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

// New returns a new instance of translator for the 'fr_FR' locale
func New() locales.Translator {
	return &fr_FR{
		locale:                 "fr_FR",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  "\u202f",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$AR", "ATS", "$AU", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "FB", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$BM", "$BN", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "$BZ", "$CA", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$CL", "CNH", "CNX", "CNY", "$CO", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "£CY", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "£E", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "$FJ", "£FK", "F", "£GB", "GEK", "GEL", "GHC", "GHS", "£GI", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "£IE", "£IL", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "₤IT", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "FC", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "£LB", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "fMA", "MCF", "MDC", "MDL", "MGA", "Fmg", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "£MT", "MUR", "MVP", "MVR", "MWK", "$MX", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "$NA", "NGN", "NIC", "$C", "NLG", "NOK", "NPR", "$NZ", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "$RH", "ROL", "L", "RSD", "RUB", "р.", "FR", "SAR", "$SB", "SCR", "SDD", "SDG", "SDP", "SEK", "$SG", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "$SR", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "$T", "TPE", "TRL", "LT", "$TT", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$US", "USN", "USS", "UYI", "UYP", "$UY", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "$WS", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "DTS", "XEU", "XFO", "XFU", "XOF", "XPD", "FCFP", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "Kw", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
		daysAbbreviated:        []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		daysNarrow:             []string{"D", "L", "M", "M", "J", "V", "S"},
		daysShort:              []string{"di", "lu", "ma", "me", "je", "ve", "sa"},
		daysWide:               []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		timezones:              map[string]string{"ACDT": "heure d’été du centre de l’Australie", "ACST": "heure normale du centre de l’Australie", "ACT": "heure normale de l’Acre", "ACWDT": "heure d’été du centre-ouest de l’Australie", "ACWST": "heure normale du centre-ouest de l’Australie", "ADT": "heure d’été de l’Atlantique", "ADT Arabia": "heure d’été de l’Arabie", "AEDT": "heure d’été de l’Est de l’Australie", "AEST": "heure normale de l’Est de l’Australie", "AFT": "heure de l’Afghanistan", "AKDT": "heure d’été de l’Alaska", "AKST": "heure normale de l’Alaska", "AMST": "heure d’été de l’Amazonie", "AMST Armenia": "heure d’été d’Arménie", "AMT": "heure normale de l’Amazonie", "AMT Armenia": "heure normale de l’Arménie", "ANAST": "heure d’été d’Anadyr", "ANAT": "heure normale d’Anadyr", "ARST": "heure d’été de l’Argentine", "ART": "heure normale d’Argentine", "AST": "heure normale de l’Atlantique", "AST Arabia": "heure normale de l’Arabie", "AWDT": "heure d’été de l’Ouest de l’Australie", "AWST": "heure normale de l’Ouest de l’Australie", "AZST": "heure d’été d’Azerbaïdjan", "AZT": "heure normale de l’Azerbaïdjan", "BDT Bangladesh": "heure d’été du Bangladesh", "BNT": "heure du Brunei", "BOT": "heure de Bolivie", "BRST": "heure d’été de Brasilia", "BRT": "heure normale de Brasilia", "BST Bangladesh": "heure normale du Bangladesh", "BT": "heure du Bhoutan", "CAST": "heure de Casey", "CAT": "heure normale d’Afrique centrale", "CCT": "heure des îles Cocos", "CDT": "heure d’été du centre nord-américain", "CHADT": "heure d’été des îles Chatham", "CHAST": "heure normale des îles Chatham", "CHUT": "heure de Chuuk", "CKT": "heure normale des îles Cook", "CKT DST": "heure d’été des îles Cook", "CLST": "heure d’été du Chili", "CLT": "heure normale du Chili", "COST": "heure d’été de Colombie", "COT": "heure normale de Colombie", "CST": "heure normale du centre nord-américain", "CST China": "heure normale de la Chine", "CST China DST": "heure d’été de Chine", "CVST": "heure d’été du Cap-Vert", "CVT": "heure normale du Cap-Vert", "CXT": "heure de l’île Christmas", "ChST": "heure des Chamorro", "ChST NMI": "heure des îles Mariannes du Nord", "CuDT": "heure d’été de Cuba", "CuST": "heure normale de Cuba", "DAVT": "heure de Davis", "DDUT": "heure de Dumont-d’Urville", "EASST": "heure d’été de l’île de Pâques", "EAST": "heure normale de l’île de Pâques", "EAT": "heure normale d’Afrique de l’Est", "ECT": "heure de l’Équateur", "EDT": "heure d’été de l’Est nord-américain", "EGDT": "heure d’été de l’Est du Groenland", "EGST": "heure normale de l’Est du Groenland", "EST": "heure normale de l’Est nord-américain", "FEET": "heure de Kaliningrad", "FJT": "heure normale des îles Fidji", "FJT Summer": "heure d’été des îles Fidji", "FKST": "heure d’été des îles Malouines", "FKT": "heure normale des îles Malouines", "FNST": "heure d’été de Fernando de Noronha", "FNT": "heure normale de Fernando de Noronha", "GALT": "heure des îles Galápagos", "GAMT": "heure des îles Gambier", "GEST": "heure d’été de Géorgie", "GET": "heure normale de la Géorgie", "GFT": "heure de la Guyane française", "GIT": "heure des îles Gilbert", "GMT": "heure moyenne de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "heure du Golfe", "GST Guam": "heure de Guam", "GYT": "heure du Guyana", "HADT": "heure d’été d’Hawaï - Aléoutiennes", "HAST": "heure normale d’Hawaï - Aléoutiennes", "HKST": "heure d’été de Hong Kong", "HKT": "heure normale de Hong Kong", "HOVST": "heure d’été de Hovd", "HOVT": "heure normale de Hovd", "ICT": "heure d’Indochine", "IDT": "heure d’été d’Israël", "IOT": "heure de l’Océan Indien", "IRKST": "heure d’été d’Irkoutsk", "IRKT": "heure normale d’Irkoutsk", "IRST": "heure normale d’Iran", "IRST DST": "heure d’été d’Iran", "IST": "heure de l’Inde", "IST Israel": "heure normale d’Israël", "JDT": "heure d’été du Japon", "JST": "heure normale du Japon", "KOST": "heure de Kosrae", "KRAST": "heure d’été de Krasnoïarsk", "KRAT": "heure normale de Krasnoïarsk", "KST": "heure normale de la Corée", "KST DST": "heure d’été de Corée", "LHDT": "heure d’été de Lord Howe", "LHST": "heure normale de Lord Howe", "LINT": "heure des îles de la Ligne", "MAGST": "heure d’été de Magadan", "MAGT": "heure normale de Magadan", "MART": "heure des îles Marquises", "MAWT": "heure de Mawson", "MDT": "heure d’été des Rocheuses", "MESZ": "heure d’été d’Europe centrale", "MEZ": "heure normale d’Europe centrale", "MHT": "heure des îles Marshall", "MMT": "heure du Myanmar", "MSD": "heure d’été de Moscou", "MST": "heure normale des Rocheuses", "MUST": "heure d’été de Maurice", "MUT": "heure normale de Maurice", "MVT": "heure des Maldives", "MYT": "heure de la Malaisie", "NCT": "heure normale de la Nouvelle-Calédonie", "NDT": "heure d’été de Terre-Neuve", "NDT New Caledonia": "heure d’été de Nouvelle-Calédonie", "NFDT": "heure d’été de l’île Norfolk", "NFT": "heure normale de l’île Norfolk", "NOVST": "heure d’été de Novossibirsk", "NOVT": "heure normale de Novossibirsk", "NPT": "heure du Népal", "NRT": "heure de Nauru", "NST": "heure normale de Terre-Neuve", "NUT": "heure de Niue", "NZDT": "heure d’été de la Nouvelle-Zélande", "NZST": "heure normale de la Nouvelle-Zélande", "OESZ": "heure d’été d’Europe de l’Est", "OEZ": "heure normale d’Europe de l’Est", "OMSST": "heure d’été de Omsk", "OMST": "heure normale de Omsk", "PDT": "heure d’été du Pacifique nord-américain", "PDTM": "heure d’été du Pacifique mexicain", "PETDT": "heure d’été de Petropavlovsk-Kamchatski", "PETST": "heure normale de Petropavlovsk-Kamchatski", "PGT": "heure de la Papouasie-Nouvelle-Guinée", "PHOT": "heure des îles Phoenix", "PKT": "heure normale du Pakistan", "PKT DST": "heure d’été du Pakistan", "PMDT": "heure d’été de Saint-Pierre-et-Miquelon", "PMST": "heure normale de Saint-Pierre-et-Miquelon", "PONT": "heure de l’île de Pohnpei", "PST": "heure normale du Pacifique nord-américain", "PST Philippine": "heure normale des Philippines", "PST Philippine DST": "heure d’été des Philippines", "PST Pitcairn": "heure des îles Pitcairn", "PSTM": "heure normale du Pacifique mexicain", "PWT": "heure des Palaos", "PYST": "heure d’été du Paraguay", "PYT": "heure normale du Paraguay", "PYT Korea": "heure de Pyongyang", "RET": "heure de La Réunion", "ROTT": "heure de Rothera", "SAKST": "heure d’été de Sakhaline", "SAKT": "heure normale de Sakhaline", "SAMST": "heure d’été de Samara", "SAMT": "heure normale de Samara", "SAST": "heure normale d’Afrique méridionale", "SBT": "heure des îles Salomon", "SCT": "heure des Seychelles", "SGT": "heure de Singapour", "SLST": "heure de Lanka", "SRT": "heure du Suriname", "SST Samoa": "heure normale des Samoa", "SST Samoa Apia": "heure normale d’Apia", "SST Samoa Apia DST": "heure d’été d’Apia", "SST Samoa DST": "heure d’été des Samoa", "SYOT": "heure de Syowa", "TAAF": "heure des Terres australes et antarctiques françaises", "TAHT": "heure de Tahiti", "TJT": "heure du Tadjikistan", "TKT": "heure de Tokelau", "TLT": "heure du Timor oriental", "TMST": "heure d’été du Turkménistan", "TMT": "heure normale du Turkménistan", "TOST": "heure d’été de Tonga", "TOT": "heure normale des Tonga", "TVT": "heure des Tuvalu", "TWT": "heure normale de Taipei", "TWT DST": "heure d’été de Taipei", "ULAST": "heure d’été d’Oulan-Bator", "ULAT": "heure normale d’Oulan-Bator", "UYST": "heure d’été de l’Uruguay", "UYT": "heure normale de l’Uruguay", "UZT": "heure normale de l’Ouzbékistan", "UZT DST": "heure d’été de l’Ouzbékistan", "VET": "heure du Venezuela", "VLAST": "heure d’été de Vladivostok", "VLAT": "heure normale de Vladivostok", "VOLST": "heure d’été de Volgograd", "VOLT": "heure normale de Volgograd", "VOST": "heure de Vostok", "VUT": "heure normale du Vanuatu", "VUT DST": "heure d’été de Vanuatu", "WAKT": "heure de l’île Wake", "WARST": "heure d’été de l’Ouest argentin", "WART": "heure normale de l’Ouest argentin", "WAST": "heure d’Afrique de l’Ouest", "WAT": "heure d’Afrique de l’Ouest", "WESZ": "heure d’été d’Europe de l’Ouest", "WEZ": "heure normale d’Europe de l’Ouest", "WFT": "heure de Wallis-et-Futuna", "WGST": "heure d’été de l’Ouest du Groenland", "WGT": "heure normale de l’Ouest du Groenland", "WIB": "heure de l’Ouest indonésien", "WIT": "heure de l’Est indonésien", "WITA": "heure du Centre indonésien", "YAKST": "heure d’été de Iakoutsk", "YAKT": "heure normale de Iakoutsk", "YEKST": "heure d’été d’Ekaterinbourg", "YEKT": "heure normale d’Ekaterinbourg", "YST": "heure normale du Yukon", "МСК": "heure normale de Moscou", "اقتاۋ": "heure normale d’Aktaou", "اقتاۋ قالاسى": "heure d’été d’Aktaou", "اقتوبە": "heure normale d’Aqtöbe", "اقتوبە قالاسى": "heure d’été d’Aqtöbe", "الماتى": "heure normale d’Alma Ata", "الماتى قالاسى": "heure d’été d’Alma Ata", "باتىس قازاق ەلى": "heure de l’Ouest du Kazakhstan", "شىعىش قازاق ەلى": "heure de l’Est du Kazakhstan", "قازاق ەلى": "heure du Kazakhstan", "قىرعىزستان": "heure du Kirghizistan", "قىزىلوردا": "heure normale de Kyzylorda", "قىزىلوردا قالاسى": "heure d’été de Kyzylorda", "∅∅∅": "heure d’été des Açores"},
	}
}

// Locale returns the current translators string locale
func (fr *fr_FR) Locale() string {
	return fr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fr_FR'
func (fr *fr_FR) PluralsCardinal() []locales.PluralRule {
	return fr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fr_FR'
func (fr *fr_FR) PluralsOrdinal() []locales.PluralRule {
	return fr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fr_FR'
func (fr *fr_FR) PluralsRange() []locales.PluralRule {
	return fr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fr_FR'
func (fr *fr_FR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	e := int64(0)
	iMod1000000 := i % 1000000

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	} else if (e == 0 && i != 0 && iMod1000000 == 0 && v == 0) || (e < 0 || e > 5) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fr_FR'
func (fr *fr_FR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fr_FR'
func (fr *fr_FR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := fr.CardinalPluralRule(num1, v1)
	end := fr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fr *fr_FR) MonthAbbreviated(month time.Month) string {
	if len(fr.monthsAbbreviated) == 0 {
		return ""
	}
	return fr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fr *fr_FR) MonthsAbbreviated() []string {
	return fr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fr *fr_FR) MonthNarrow(month time.Month) string {
	if len(fr.monthsNarrow) == 0 {
		return ""
	}
	return fr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fr *fr_FR) MonthsNarrow() []string {
	return fr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fr *fr_FR) MonthWide(month time.Month) string {
	if len(fr.monthsWide) == 0 {
		return ""
	}
	return fr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fr *fr_FR) MonthsWide() []string {
	return fr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fr *fr_FR) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(fr.daysAbbreviated) == 0 {
		return ""
	}
	return fr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fr *fr_FR) WeekdaysAbbreviated() []string {
	return fr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fr *fr_FR) WeekdayNarrow(weekday time.Weekday) string {
	if len(fr.daysNarrow) == 0 {
		return ""
	}
	return fr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fr *fr_FR) WeekdaysNarrow() []string {
	return fr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fr *fr_FR) WeekdayShort(weekday time.Weekday) string {
	if len(fr.daysShort) == 0 {
		return ""
	}
	return fr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fr *fr_FR) WeekdaysShort() []string {
	return fr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fr *fr_FR) WeekdayWide(weekday time.Weekday) string {
	if len(fr.daysWide) == 0 {
		return ""
	}
	return fr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fr *fr_FR) WeekdaysWide() []string {
	return fr.daysWide
}

// Decimal returns the decimal point of number
func (fr *fr_FR) Decimal() string {
	return fr.decimal
}

// Group returns the group of number
func (fr *fr_FR) Group() string {
	return fr.group
}

// Group returns the minus sign of number
func (fr *fr_FR) Minus() string {
	return fr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fr_FR' and handles both Whole and Real numbers based on 'v'
func (fr *fr_FR) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fr.group) - 1; j >= 0; j-- {
					b = append(b, fr.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fr_FR' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fr *fr_FR) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fr.percentSuffix...)

	b = append(b, fr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fr_FR'
func (fr *fr_FR) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fr.group) - 1; j >= 0; j-- {
					b = append(b, fr.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, fr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fr_FR'
// in accounting notation.
func (fr *fr_FR) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fr.currencies[currency]
	l := len(s) + len(symbol) + 6 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fr.group) - 1; j >= 0; j-- {
					b = append(b, fr.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, fr.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, fr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, fr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fr_FR'
func (fr *fr_FR) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
