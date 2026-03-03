package sk_SK

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sk_SK struct {
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

// New returns a new instance of translator for the 'sk_SK' locale
func New() locales.Translator {
	return &sk_SK{
		locale:                 "sk_SK",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "NIS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "Cg", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "máj", "jún", "júl", "aug", "sep", "okt", "nov", "dec"},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januára", "februára", "marca", "apríla", "mája", "júna", "júla", "augusta", "septembra", "októbra", "novembra", "decembra"},
		daysAbbreviated:        []string{"ne", "po", "ut", "st", "št", "pi", "so"},
		daysNarrow:             []string{"n", "p", "u", "s", "š", "p", "s"},
		daysWide:               []string{"nedeľa", "pondelok", "utorok", "streda", "štvrtok", "piatok", "sobota"},
		timezones:              map[string]string{"ACDT": "stredoaustrálsky letný čas", "ACST": "stredoaustrálsky štandardný čas", "ACT": "acrejský štandardný čas", "ACWDT": "stredozápadný austrálsky letný čas", "ACWST": "stredozápadný austrálsky štandardný čas", "ADT": "atlantický letný čas", "ADT Arabia": "arabský letný čas", "AEDT": "východoaustrálsky letný čas", "AEST": "východoaustrálsky štandardný čas", "AFT": "afganský čas", "AKDT": "aljašský letný čas", "AKST": "aljašský štandardný čas", "AMST": "amazonský letný čas", "AMST Armenia": "arménsky letný čas", "AMT": "amazonský štandardný čas", "AMT Armenia": "arménsky štandardný čas", "ANAST": "Anadyrský letný čas", "ANAT": "Anadyrský štandardný čas", "ARST": "argentínsky letný čas", "ART": "argentínsky štandardný čas", "AST": "atlantický štandardný čas", "AST Arabia": "arabský štandardný čas", "AWDT": "západoaustrálsky letný čas", "AWST": "západoaustrálsky štandardný čas", "AZST": "azerbajdžanský letný čas", "AZT": "azerbajdžanský štandardný čas", "BDT Bangladesh": "bangladéšsky letný čas", "BNT": "brunejský čas", "BOT": "bolívijský čas", "BRST": "brazílsky letný čas", "BRT": "brazílsky štandardný čas", "BST Bangladesh": "bangladéšsky štandardný čas", "BT": "bhutánsky čas", "CAST": "čas Caseyho stanice", "CAT": "stredoafrický čas", "CCT": "čas Kokosových ostrovov", "CDT": "severoamerický centrálny letný čas", "CHADT": "chathamský letný čas", "CHAST": "chathamský štandardný čas", "CHUT": "chuukský čas", "CKT": "štandardný čas Cookových ostrovov", "CKT DST": "letný čas Cookových ostrovov", "CLST": "čilský letný čas", "CLT": "čilský štandardný čas", "COST": "kolumbijský letný čas", "COT": "kolumbijský štandardný čas", "CST": "severoamerický centrálny štandardný čas", "CST China": "čínsky štandardný čas", "CST China DST": "čínsky letný čas", "CVST": "kapverdský letný čas", "CVT": "kapverdský štandardný čas", "CXT": "čas Vianočného ostrova", "ChST": "chamorrský čas", "ChST NMI": "severomariánsky čas", "CuDT": "kubánsky letný čas", "CuST": "kubánsky štandardný čas", "DAVT": "čas Davisovej stanice", "DDUT": "čas stanice Dumonta d’Urvillea", "EASST": "letný čas Veľkonočného ostrova", "EAST": "štandardný čas Veľkonočného ostrova", "EAT": "východoafrický čas", "ECT": "ekvádorský čas", "EDT": "severoamerický východný letný čas", "EGDT": "východogrónsky letný čas", "EGST": "východogrónsky štandardný čas", "EST": "severoamerický východný štandardný čas", "FEET": "minský čas", "FJT": "fidžijský štandardný čas", "FJT Summer": "fidžijský letný čas", "FKST": "falklandský letný čas", "FKT": "falklandský štandardný čas", "FNST": "letný čas súostrovia Fernando de Noronha", "FNT": "štandardný čas súostrovia Fernando de Noronha", "GALT": "galapágsky čas", "GAMT": "gambierský čas", "GEST": "gruzínsky letný čas", "GET": "gruzínsky štandardný čas", "GFT": "francúzskoguyanský čas", "GIT": "čas Gilbertových ostrovov", "GMT": "greenwichský čas", "GNSST": "GNSST", "GNST": "GNST", "GST": "čas Južnej Georgie", "GST Guam": "guamský čas", "GYT": "guyanský čas", "HADT": "havajsko-aleutský letný čas", "HAST": "havajsko-aleutský štandardný čas", "HKST": "hongkonský letný čas", "HKT": "hongkonský štandardný čas", "HOVST": "chovdský letný čas", "HOVT": "chovdský štandardný čas", "ICT": "indočínsky čas", "IDT": "izraelský letný čas", "IOT": "indickooceánsky čas", "IRKST": "irkutský letný čas", "IRKT": "irkutský štandardný čas", "IRST": "iránsky štandardný čas", "IRST DST": "iránsky letný čas", "IST": "indický čas", "IST Israel": "izraelský štandardný čas", "JDT": "japonský letný čas", "JST": "japonský štandardný čas", "KOST": "kosrajský čas", "KRAST": "krasnojarský letný čas", "KRAT": "krasnojarský štandardný čas", "KST": "kórejský štandardný čas", "KST DST": "kórejský letný čas", "LHDT": "letný čas ostrova lorda Howa", "LHST": "štandardný čas ostrova lorda Howa", "LINT": "čas Rovníkových ostrovov", "MAGST": "magadanský letný čas", "MAGT": "magadanský štandardný čas", "MART": "markézsky čas", "MAWT": "čas Mawsonovej stanice", "MDT": "macajský letný čas", "MESZ": "stredoeurópsky letný čas", "MEZ": "stredoeurópsky štandardný čas", "MHT": "čas Marshallových ostrovov", "MMT": "mjanmarský čas", "MSD": "moskovský letný čas", "MST": "macajský štandardný čas", "MUST": "maurícijský letný čas", "MUT": "maurícijský štandardný čas", "MVT": "maldivský čas", "MYT": "malajzijský čas", "NCT": "novokaledónsky štandardný čas", "NDT": "newfoundlandský letný čas", "NDT New Caledonia": "novokaledónsky letný čas", "NFDT": "norfolský letný čas", "NFT": "norfolský štandardný čas", "NOVST": "novosibirský letný čas", "NOVT": "novosibirský štandardný čas", "NPT": "nepálsky čas", "NRT": "nauruský čas", "NST": "newfoundlandský štandardný čas", "NUT": "niuejský čas", "NZDT": "novozélandský letný čas", "NZST": "novozélandský štandardný čas", "OESZ": "východoeurópsky letný čas", "OEZ": "východoeurópsky štandardný čas", "OMSST": "omský letný čas", "OMST": "omský štandardný čas", "PDT": "severoamerický tichomorský letný čas", "PDTM": "mexický tichomorský letný čas", "PETDT": "Petropavlovsk-Kamčatskijský letný čas", "PETST": "Petropavlovsk-Kamčatský štandardný čas", "PGT": "čas Papuy-Novej Guiney", "PHOT": "čas Fénixových ostrovov", "PKT": "pakistanský štandardný čas", "PKT DST": "pakistanský letný čas", "PMDT": "pierre-miquelonský letný čas", "PMST": "pierre-miquelonský štandardný čas", "PONT": "ponapský čas", "PST": "severoamerický tichomorský štandardný čas", "PST Philippine": "filipínsky štandardný čas", "PST Philippine DST": "filipínsky letný čas", "PST Pitcairn": "čas Pitcairnových ostrovov", "PSTM": "mexický tichomorský štandardný čas", "PWT": "palauský čas", "PYST": "paraguajský letný čas", "PYT": "paraguajský štandardný čas", "PYT Korea": "pchjongjanský čas", "RET": "réunionský čas", "ROTT": "čas Rotherovej stanice", "SAKST": "sachalinský letný čas", "SAKT": "sachalinský štandardný čas", "SAMST": "Samarský letný čas", "SAMT": "Samarský štandardný čas", "SAST": "juhoafrický čas", "SBT": "čas Šalamúnových ostrovov", "SCT": "seychelský čas", "SGT": "singapurský štandardný čas", "SLST": "srílanský čas", "SRT": "surinamský čas", "SST Samoa": "samojský štandardný čas", "SST Samoa Apia": "apijský štandardný čas", "SST Samoa Apia DST": "apijský letný čas", "SST Samoa DST": "samojský letný čas", "SYOT": "čas stanice Šówa", "TAAF": "čas Francúzskych južných a antarktických území", "TAHT": "tahitský čas", "TJT": "tadžický čas", "TKT": "tokelauský čas", "TLT": "východotimorský čas", "TMST": "turkménsky letný čas", "TMT": "turkménsky štandardný čas", "TOST": "tonžský letný čas", "TOT": "tonžský štandardný čas", "TVT": "tuvalský čas", "TWT": "tchajpejský štandardný čas", "TWT DST": "tchajpejský letný čas", "ULAST": "ulanbátarský letný čas", "ULAT": "ulanbátarský štandardný čas", "UYST": "uruguajský letný čas", "UYT": "uruguajský štandardný čas", "UZT": "uzbecký štandardný čas", "UZT DST": "uzbecký letný čas", "VET": "venezuelský čas", "VLAST": "vladivostocký letný čas", "VLAT": "vladivostocký štandardný čas", "VOLST": "volgogradský letný čas", "VOLT": "volgogradský štandardný čas", "VOST": "čas stanice Vostok", "VUT": "vanuatský štandardný čas", "VUT DST": "vanuatský letný čas", "WAKT": "čas ostrova Wake", "WARST": "západoargentínsky letný čas", "WART": "západoargentínsky štandardný čas", "WAST": "západoafrický čas", "WAT": "západoafrický čas", "WESZ": "západoeurópsky letný čas", "WEZ": "západoeurópsky štandardný čas", "WFT": "čas ostrovov Wallis a Futuna", "WGST": "západogrónsky letný čas", "WGT": "západogrónsky štandardný čas", "WIB": "západoindonézsky čas", "WIT": "východoindonézsky čas", "WITA": "stredoindonézsky čas", "YAKST": "jakutský letný čas", "YAKT": "jakutský štandardný čas", "YEKST": "jekaterinburský letný čas", "YEKT": "jekaterinburský štandardný čas", "YST": "yukonský čas", "МСК": "moskovský štandardný čas", "اقتاۋ": "aktauský štandardný čas", "اقتاۋ قالاسى": "aktauský letný čas", "اقتوبە": "aktobský štandardný čas", "اقتوبە قالاسى": "aktobský letný čas", "الماتى": "almaatský štandardný čas", "الماتى قالاسى": "almaatský letný čas", "باتىس قازاق ەلى": "západokazachstanský čas", "شىعىش قازاق ەلى": "východokazachstanský čas", "قازاق ەلى": "kazachstanský čas", "قىرعىزستان": "kirgizský čas", "قىزىلوردا": "kyzylordský štandardný čas", "قىزىلوردا قالاسى": "kyzylordský letný čas", "∅∅∅": "peruánsky letný čas"},
	}
}

// Locale returns the current translators string locale
func (sk *sk_SK) Locale() string {
	return sk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sk_SK'
func (sk *sk_SK) PluralsCardinal() []locales.PluralRule {
	return sk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sk_SK'
func (sk *sk_SK) PluralsOrdinal() []locales.PluralRule {
	return sk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sk_SK'
func (sk *sk_SK) PluralsRange() []locales.PluralRule {
	return sk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sk_SK'
func (sk *sk_SK) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sk_SK'
func (sk *sk_SK) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sk_SK'
func (sk *sk_SK) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sk.CardinalPluralRule(num1, v1)
	end := sk.CardinalPluralRule(num2, v2)

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
func (sk *sk_SK) MonthAbbreviated(month time.Month) string {
	if len(sk.monthsAbbreviated) == 0 {
		return ""
	}
	return sk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sk *sk_SK) MonthsAbbreviated() []string {
	return sk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sk *sk_SK) MonthNarrow(month time.Month) string {
	if len(sk.monthsNarrow) == 0 {
		return ""
	}
	return sk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sk *sk_SK) MonthsNarrow() []string {
	return sk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sk *sk_SK) MonthWide(month time.Month) string {
	if len(sk.monthsWide) == 0 {
		return ""
	}
	return sk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sk *sk_SK) MonthsWide() []string {
	return sk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sk *sk_SK) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sk.daysAbbreviated) == 0 {
		return ""
	}
	return sk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sk *sk_SK) WeekdaysAbbreviated() []string {
	return sk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sk *sk_SK) WeekdayNarrow(weekday time.Weekday) string {
	if len(sk.daysNarrow) == 0 {
		return ""
	}
	return sk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sk *sk_SK) WeekdaysNarrow() []string {
	return sk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sk *sk_SK) WeekdayShort(weekday time.Weekday) string {
	if len(sk.daysShort) == 0 {
		return ""
	}
	return sk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sk *sk_SK) WeekdaysShort() []string {
	return sk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sk *sk_SK) WeekdayWide(weekday time.Weekday) string {
	if len(sk.daysWide) == 0 {
		return ""
	}
	return sk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sk *sk_SK) WeekdaysWide() []string {
	return sk.daysWide
}

// Decimal returns the decimal point of number
func (sk *sk_SK) Decimal() string {
	return sk.decimal
}

// Group returns the group of number
func (sk *sk_SK) Group() string {
	return sk.group
}

// Group returns the minus sign of number
func (sk *sk_SK) Minus() string {
	return sk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sk_SK' and handles both Whole and Real numbers based on 'v'
func (sk *sk_SK) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sk.group) - 1; j >= 0; j-- {
					b = append(b, sk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sk_SK' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sk *sk_SK) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sk.percentSuffix...)

	b = append(b, sk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sk_SK'
func (sk *sk_SK) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sk.group) - 1; j >= 0; j-- {
					b = append(b, sk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sk_SK'
// in accounting notation.
func (sk *sk_SK) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sk.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sk.group) - 1; j >= 0; j-- {
					b = append(b, sk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, sk.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sk.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sk_SK'
func (sk *sk_SK) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
