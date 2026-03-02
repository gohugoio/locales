package cs

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type cs struct {
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

// New returns a new instance of translator for the 'cs' locale
func New() locales.Translator {
	return &cs{
		locale:                 "cs",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "Kčs", "CUC", "CUP", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "ECU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "led", "úno", "bře", "dub", "kvě", "čvn", "čvc", "srp", "zář", "říj", "lis", "pro"},
		monthsWide:             []string{"", "ledna", "února", "března", "dubna", "května", "června", "července", "srpna", "září", "října", "listopadu", "prosince"},
		daysAbbreviated:        []string{"ne", "po", "út", "st", "čt", "pá", "so"},
		daysNarrow:             []string{"N", "P", "Ú", "S", "Č", "P", "S"},
		daysWide:               []string{"neděle", "pondělí", "úterý", "středa", "čtvrtek", "pátek", "sobota"},
		timezones:              map[string]string{"ACDT": "středoaustralský letní čas", "ACST": "acrejský letní čas", "ACT": "acrejský standardní čas", "ACWDT": "středozápadní australský letní čas", "ACWST": "středozápadní australský standardní čas", "ADT": "atlantický letní čas", "ADT Arabia": "arabský letní čas", "AEDT": "východoaustralský letní čas", "AEST": "východoaustralský standardní čas", "AFT": "afghánský čas", "AKDT": "aljašský letní čas", "AKST": "aljašský standardní čas", "AMST": "amazonský letní čas", "AMST Armenia": "arménský letní čas", "AMT": "amazonský standardní čas", "AMT Armenia": "arménský standardní čas", "ANAST": "anadyrský letní čas", "ANAT": "anadyrský standardní čas", "ARST": "argentinský letní čas", "ART": "argentinský standardní čas", "AST": "atlantický standardní čas", "AST Arabia": "arabský standardní čas", "AWDT": "západoaustralský letní čas", "AWST": "západoaustralský standardní čas", "AZST": "ázerbájdžánský letní čas", "AZT": "ázerbájdžánský standardní čas", "BDT Bangladesh": "bangladéšský letní čas", "BNT": "brunejský čas", "BOT": "bolivijský čas", "BRST": "brasilijský letní čas", "BRT": "brasilijský standardní čas", "BST Bangladesh": "bangladéšský standardní čas", "BT": "bhútánský čas", "CAST": "čas Caseyho stanice", "CAT": "středoafrický čas", "CCT": "čas Kokosových ostrovů", "CDT": "severoamerický centrální letní čas", "CHADT": "chathamský letní čas", "CHAST": "chathamský standardní čas", "CHUT": "chuukský čas", "CKT": "standardní čas Cookových ostrovů", "CKT DST": "letní čas Cookových ostrovů", "CLST": "chilský letní čas", "CLT": "chilský standardní čas", "COST": "kolumbijský letní čas", "COT": "kolumbijský standardní čas", "CST": "severoamerický centrální standardní čas", "CST China": "čínský standardní čas", "CST China DST": "čínský letní čas", "CVST": "kapverdský letní čas", "CVT": "kapverdský standardní čas", "CXT": "čas Vánočního ostrova", "ChST": "chamorrský čas", "ChST NMI": "Severomariánský čas", "CuDT": "kubánský letní čas", "CuST": "kubánský standardní čas", "DAVT": "čas Davisovy stanice", "DDUT": "čas stanice Dumonta d’Urvilla", "EASST": "letní čas Velikonočního ostrova", "EAST": "standardní čas Velikonočního ostrova", "EAT": "východoafrický čas", "ECT": "ekvádorský čas", "EDT": "severoamerický východní letní čas", "EGDT": "východogrónský letní čas", "EGST": "východogrónský standardní čas", "EST": "severoamerický východní standardní čas", "FEET": "dálněvýchodoevropský čas", "FJT": "fidžijský standardní čas", "FJT Summer": "fidžijský letní čas", "FKST": "falklandský letní čas", "FKT": "falklandský standardní čas", "FNST": "letní čas souostroví Fernando de Noronha", "FNT": "standardní čas souostroví Fernando de Noronha", "GALT": "galapážský čas", "GAMT": "gambierský čas", "GEST": "gruzínský letní čas", "GET": "gruzínský standardní čas", "GFT": "francouzskoguyanský čas", "GIT": "čas Gilbertových ostrovů", "GMT": "greenwichský střední čas", "GNSST": "GNSST", "GNST": "GNST", "GST": "standardní čas Perského zálivu", "GST Guam": "Guamský čas", "GYT": "guyanský čas", "HADT": "havajsko-aleutský standardní čas", "HAST": "havajsko-aleutský standardní čas", "HKST": "hongkongský letní čas", "HKT": "hongkongský standardní čas", "HOVST": "hovdský letní čas", "HOVT": "hovdský standardní čas", "ICT": "indočínský čas", "IDT": "izraelský letní čas", "IOT": "indickooceánský čas", "IRKST": "irkutský letní čas", "IRKT": "irkutský standardní čas", "IRST": "íránský standardní čas", "IRST DST": "íránský letní čas", "IST": "indický čas", "IST Israel": "izraelský standardní čas", "JDT": "japonský letní čas", "JST": "japonský standardní čas", "KOST": "kosrajský čas", "KRAST": "krasnojarský letní čas", "KRAT": "krasnojarský standardní čas", "KST": "korejský standardní čas", "KST DST": "korejský letní čas", "LHDT": "letní čas ostrova lorda Howa", "LHST": "standardní čas ostrova lorda Howa", "LINT": "čas Rovníkových ostrovů", "MAGST": "magadanský letní čas", "MAGT": "magadanský standardní čas", "MART": "markézský čas", "MAWT": "čas Mawsonovy stanice", "MDT": "Macajský letní čas", "MESZ": "středoevropský letní čas", "MEZ": "středoevropský standardní čas", "MHT": "čas Marshallových ostrovů", "MMT": "myanmarský čas", "MSD": "moskevský letní čas", "MST": "Macajský standardní čas", "MUST": "mauricijský letní čas", "MUT": "mauricijský standardní čas", "MVT": "maledivský čas", "MYT": "malajský čas", "NCT": "novokaledonský standardní čas", "NDT": "newfoundlandský letní čas", "NDT New Caledonia": "novokaledonský letní čas", "NFDT": "norfolkský letní čas", "NFT": "norfolkský standardní čas", "NOVST": "novosibirský letní čas", "NOVT": "novosibirský standardní čas", "NPT": "nepálský čas", "NRT": "naurský čas", "NST": "newfoundlandský standardní čas", "NUT": "niuejský čas", "NZDT": "novozélandský letní čas", "NZST": "novozélandský standardní čas", "OESZ": "východoevropský letní čas", "OEZ": "východoevropský standardní čas", "OMSST": "omský letní čas", "OMST": "omský standardní čas", "PDT": "severoamerický pacifický letní čas", "PDTM": "mexický pacifický letní čas", "PETDT": "petropavlovsko-kamčatský letní čas", "PETST": "petropavlovsko-kamčatský standardní čas", "PGT": "čas Papuy-Nové Guiney", "PHOT": "čas Fénixových ostrovů", "PKT": "pákistánský standardní čas", "PKT DST": "pákistánský letní čas", "PMDT": "pierre-miquelonský letní čas", "PMST": "pierre-miquelonský standardní čas", "PONT": "ponapský čas", "PST": "severoamerický pacifický standardní čas", "PST Philippine": "filipínský standardní čas", "PST Philippine DST": "filipínský letní čas", "PST Pitcairn": "čas Pitcairnových ostrovů", "PSTM": "mexický pacifický standardní čas", "PWT": "palauský čas", "PYST": "paraguayský letní čas", "PYT": "paraguayský standardní čas", "PYT Korea": "pchjongjangský čas", "RET": "réunionský čas", "ROTT": "čas Rotherovy stanice", "SAKST": "sachalinský letní čas", "SAKT": "sachalinský standardní čas", "SAMST": "samarský letní čas", "SAMT": "samarský standardní čas", "SAST": "jihoafrický čas", "SBT": "čas Šalamounových ostrovů", "SCT": "seychelský čas", "SGT": "singapurský čas", "SLST": "Srílanský čas", "SRT": "surinamský čas", "SST Samoa": "samojský standardní čas", "SST Samoa Apia": "apijský standardní čas", "SST Samoa Apia DST": "apijský letní čas", "SST Samoa DST": "samojský letní čas", "SYOT": "čas stanice Šówa", "TAAF": "čas Francouzských jižních a antarktických území", "TAHT": "tahitský čas", "TJT": "tádžický čas", "TKT": "tokelauský čas", "TLT": "východotimorský čas", "TMST": "turkmenský letní čas", "TMT": "turkmenský standardní čas", "TOST": "tonžský letní čas", "TOT": "tonžský standardní čas", "TVT": "tuvalský čas", "TWT": "tchajpejský standardní čas", "TWT DST": "tchajpejský letní čas", "ULAST": "ulánbátarský letní čas", "ULAT": "ulánbátarský standardní čas", "UYST": "uruguayský letní čas", "UYT": "uruguayský standardní čas", "UZT": "uzbecký standardní čas", "UZT DST": "uzbecký letní čas", "VET": "venezuelský čas", "VLAST": "vladivostocký letní čas", "VLAT": "vladivostocký standardní čas", "VOLST": "volgogradský letní čas", "VOLT": "volgogradský standardní čas", "VOST": "čas stanice Vostok", "VUT": "vanuatský standardní čas", "VUT DST": "vanuatský letní čas", "WAKT": "čas ostrova Wake", "WARST": "západoargentinský letní čas", "WART": "západoargentinský standardní čas", "WAST": "západoafrický čas", "WAT": "západoafrický čas", "WESZ": "západoevropský letní čas", "WEZ": "západoevropský standardní čas", "WFT": "čas ostrovů Wallis a Futuna", "WGST": "západogrónský letní čas", "WGT": "západogrónský standardní čas", "WIB": "západoindonéský čas", "WIT": "východoindonéský čas", "WITA": "středoindonéský čas", "YAKST": "jakutský letní čas", "YAKT": "jakutský standardní čas", "YEKST": "jekatěrinburský letní čas", "YEKT": "jekatěrinburský standardní čas", "YST": "yukonský čas", "МСК": "moskevský standardní čas", "اقتاۋ": "Aktauský standardní čas", "اقتاۋ قالاسى": "Aktauský letní čas", "اقتوبە": "Aktobský standardní čas", "اقتوبە قالاسى": "Aktobský letní čas", "الماتى": "Almatský standardní čas", "الماتى قالاسى": "Almatský letní čas", "باتىس قازاق ەلى": "západokazachstánský čas", "شىعىش قازاق ەلى": "východokazachstánský čas", "قازاق ەلى": "kazachstánský čas", "قىرعىزستان": "kyrgyzský čas", "قىزىلوردا": "Kyzylordský standardní čas", "قىزىلوردا قالاسى": "Kyzylordský letní čas", "∅∅∅": "peruánský letní čas"},
	}
}

// Locale returns the current translators string locale
func (cs *cs) Locale() string {
	return cs.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'cs'
func (cs *cs) PluralsCardinal() []locales.PluralRule {
	return cs.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'cs'
func (cs *cs) PluralsOrdinal() []locales.PluralRule {
	return cs.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'cs'
func (cs *cs) PluralsRange() []locales.PluralRule {
	return cs.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'cs'
func (cs *cs) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew
	} else if v != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'cs'
func (cs *cs) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'cs'
func (cs *cs) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := cs.CardinalPluralRule(num1, v1)
	end := cs.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (cs *cs) MonthAbbreviated(month time.Month) string {
	return cs.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (cs *cs) MonthsAbbreviated() []string {
	return cs.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (cs *cs) MonthNarrow(month time.Month) string {
	return cs.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (cs *cs) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (cs *cs) MonthWide(month time.Month) string {
	return cs.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (cs *cs) MonthsWide() []string {
	return cs.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (cs *cs) WeekdayAbbreviated(weekday time.Weekday) string {
	return cs.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (cs *cs) WeekdaysAbbreviated() []string {
	return cs.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (cs *cs) WeekdayNarrow(weekday time.Weekday) string {
	return cs.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (cs *cs) WeekdaysNarrow() []string {
	return cs.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (cs *cs) WeekdayShort(weekday time.Weekday) string {
	return cs.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (cs *cs) WeekdaysShort() []string {
	return cs.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (cs *cs) WeekdayWide(weekday time.Weekday) string {
	return cs.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (cs *cs) WeekdaysWide() []string {
	return cs.daysWide
}

// Decimal returns the decimal point of number
func (cs *cs) Decimal() string {
	return cs.decimal
}

// Group returns the group of number
func (cs *cs) Group() string {
	return cs.group
}

// Group returns the minus sign of number
func (cs *cs) Minus() string {
	return cs.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'cs' and handles both Whole and Real numbers based on 'v'
func (cs *cs) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(cs.group) - 1; j >= 0; j-- {
					b = append(b, cs.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'cs' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (cs *cs) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cs.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, cs.percentSuffix...)

	b = append(b, cs.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'cs'
func (cs *cs) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(cs.group) - 1; j >= 0; j-- {
					b = append(b, cs.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, cs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, cs.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'cs'
// in accounting notation.
func (cs *cs) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(cs.group) - 1; j >= 0; j-- {
					b = append(b, cs.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, cs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, cs.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, cs.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'cs'
func (cs *cs) FmtDateShort(t time.Time) string {
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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'cs'
func (cs *cs) FmtDateMedium(t time.Time) string {
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

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'cs'
func (cs *cs) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, cs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'cs'
func (cs *cs) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, cs.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, cs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'cs'
func (cs *cs) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'cs'
func (cs *cs) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'cs'
func (cs *cs) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'cs'
func (cs *cs) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := cs.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
