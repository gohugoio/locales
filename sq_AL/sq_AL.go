package sq_AL

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sq_AL struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'sq_AL' locale
func New() locales.Translator {
	return &sq_AL{
		locale:                 "sq_AL",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "Lekë", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKS", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "ANG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jan", "shk", "mar", "pri", "maj", "qer", "korr", "gush", "sht", "tet", "nën", "dhj"},
		monthsNarrow:           []string{"", "j", "sh", "m", "p", "m", "q", "k", "g", "sh", "t", "n", "dh"},
		monthsWide:             []string{"", "janar", "shkurt", "mars", "prill", "maj", "qershor", "korrik", "gusht", "shtator", "tetor", "nëntor", "dhjetor"},
		daysAbbreviated:        []string{"die", "hën", "mar", "mër", "enj", "pre", "sht"},
		daysNarrow:             []string{"d", "h", "m", "m", "e", "p", "sh"},
		daysWide:               []string{"e diel", "e hënë", "e martë", "e mërkurë", "e enjte", "e premte", "e shtunë"},
		periodsAbbreviated:     []string{"p.d.", "m.d."},
		timezones:              map[string]string{"ACDT": "Ora verore e Australisë Qendrore", "ACST": "Ora standarde e Australisë Qendrore", "ACT": "Ora standarde e Ejkrit [Ako]", "ACWDT": "Ora verore e Australisë Qendroro-Perëndimore", "ACWST": "Ora standarde e Australisë Qendroro-Perëndimore", "ADT": "Ora verore e Atlantikut", "ADT Arabia": "Ora verore arabe", "AEDT": "Ora verore e Australisë Lindore", "AEST": "Ora standarde e Australisë Lindore", "AFT": "Ora e Afganistanit", "AKDT": "Ora verore e Alaskës", "AKST": "Ora standarde e Alaskës", "AMST": "Ora verore e Amazonës", "AMST Armenia": "Ora verore e Armenisë", "AMT": "Ora standarde e Amazonës", "AMT Armenia": "Ora standarde e Armenisë", "ANAST": "Ora verore e Anadirit", "ANAT": "Ora standarde e Anadirit", "ARST": "Ora verore e Argjentinës", "ART": "Ora standarde e Argjentinës", "AST": "Ora standarde e Atlantikut", "AST Arabia": "Ora standarde arabe", "AWDT": "Ora verore e Australisë Perëndimore", "AWST": "Ora standarde e Australisë Perëndimore", "AZST": "Ora verore e Azerbajxhanit", "AZT": "Ora standarde e Azerbajxhanit", "BDT Bangladesh": "Ora verore e Bangladeshit", "BNT": "Ora e Brunei-Durasalamit", "BOT": "Ora e Bolivisë", "BRST": "Ora verore e Brazilisë", "BRT": "Ora standarde e Brazilisë", "BST Bangladesh": "Ora standarde e Bangladeshit", "BT": "Ora e Butanit", "CAST": "Ora e Kejsit", "CAT": "Ora e Afrikës Qendrore", "CCT": "Ora e Ishujve Kokos", "CDT": "Ora verore e SHBA-së Qendrore", "CHADT": "Ora verore e Katamit", "CHAST": "Ora standarde e Katamit", "CHUT": "Ora e Çukut", "CKT": "Ora standarde e Ishujve Kuk", "CKT DST": "Ora verore e Ishujve Kuk", "CLST": "Ora verore e Kilit", "CLT": "Ora standarde e Kilit", "COST": "Ora verore e Kolumbisë", "COT": "Ora standarde e Kolumbisë", "CST": "Ora standarde e SHBA-së Qendrore", "CST China": "Ora standarde e Kinës", "CST China DST": "Ora verore e Kinës", "CVST": "Ora verore e Kepit të Gjelbër", "CVT": "Ora standarde e Kepit të Gjelbër", "CXT": "Ora e Ishullit të Krishtlindjeve", "ChST": "Ora e Kamorros", "ChST NMI": "Ora e Ishujve të Marianës së Veriut", "CuDT": "Ora verore e Kubës", "CuST": "Ora standarde e Kubës", "DAVT": "Ora e Dejvisit", "DDUT": "Ora e Dumont-d’Urvilës", "EASST": "Ora verore e Ishullit të Pashkës", "EAST": "Ora standarde e Ishullit të Pashkës", "EAT": "Ora e Afrikës Lindore", "ECT": "Ora e Ekuadorit", "EDT": "Ora verore e SHBA-së Lindore", "EGDT": "Ora verore e Grenlandës Lindore", "EGST": "Ora standarde e Grenlandës Lindore", "EST": "Ora standarde e SHBA-së Lindore", "FEET": "Ora e Evropës së Largët Lindore", "FJT": "Ora standarde e Fixhit", "FJT Summer": "Ora verore e Fixhit", "FKST": "Ora verore e Ishujve Falkland", "FKT": "Ora standarde e Ishujve Falkland", "FNST": "Ora verore e Fernando-de-Noronjës", "FNT": "Ora standarde e Fernando-de-Noronjës", "GALT": "Ora e Galapagosit", "GAMT": "Ora e Gambierit", "GEST": "Ora verore e Gjeorgjisë", "GET": "Ora standarde e Gjeorgjisë", "GFT": "Ora e Guajanës Franceze", "GIT": "Ora e Ishujve Gilbert", "GMT": "Ora e Grinuiçit", "GNSST": "GNSST", "GNST": "GNST", "GST": "Ora e Xhorxhas të Jugut", "GST Guam": "Ora e Guamit", "GYT": "Ora e Guajanës", "HADT": "Ora verore e Ishujve Hauai-Aleutian", "HAST": "Ora standarde e Ishujve Hauai-Aleutian", "HKST": "Ora verore e Hong-Kongut", "HKT": "Ora standarde e Hong-Kongut", "HOVST": "Ora verore e Hovdit", "HOVT": "Ora standarde e Hovdit", "ICT": "Ora e Indokinës", "IDT": "Ora verore e Izraelit", "IOT": "Ora e Oqeanit Indian", "IRKST": "Ora verore e Irkutskut", "IRKT": "Ora standarde e Irkutskut", "IRST": "Ora standarde e Iranit", "IRST DST": "Ora verore e Iranit", "IST": "Ora standarde e Indisë", "IST Israel": "Ora standarde e Izraelit", "JDT": "Ora verore e Japonisë", "JST": "Ora standarde e Japonisë", "KOST": "Ora e Kosrës", "KRAST": "Ora verore e Krasnojarskut", "KRAT": "Ora standarde e Krasnojarskut", "KST": "Ora standarde koreane", "KST DST": "Ora verore koreane", "LHDT": "Ora verore e Lord-Houit", "LHST": "Ora standarde e Lord-Houit", "LINT": "Ora e Ishujve Sporadikë Ekuatorialë", "MAGST": "Ora verore e Magadanit", "MAGT": "Ora standarde e Magadanit", "MART": "Ora e Ishujve Markezë", "MAWT": "Ora e Mausonit", "MDT": "Ora verore e Makaos", "MESZ": "Ora verore e Evropës Qendrore", "MEZ": "Ora standarde e Evropës Qendrore", "MHT": "Ora e Ishujve Marshall", "MMT": "Ora e Mianmarit", "MSD": "Ora verore e Moskës", "MST": "Ora standarde e Makaos", "MUST": "Ora verore e Mauritiusit", "MUT": "Ora standarde e Mauritiusit", "MVT": "Ora e Maldiveve", "MYT": "Ora e Malajzisë", "NCT": "Ora standarde e Kaledonisë së Re", "NDT": "Ora verore e Njufaundlendit [Tokës së Re]", "NDT New Caledonia": "Ora verore e Kaledonisë së Re", "NFDT": "Ora verore e Ishullit Norfolk", "NFT": "Ora standarde e Ishullit Norfolk", "NOVST": "Ora verore e Novosibirskut", "NOVT": "Ora standarde e Novosibirskut", "NPT": "Ora e Nepalit", "NRT": "Ora e Naurusë", "NST": "Ora standarde e Njufaundlendit [Tokës së Re]", "NUT": "Ora e Niuesë", "NZDT": "Ora verore e Zelandës së Re", "NZST": "Ora standarde e Zelandës së Re", "OESZ": "Ora verore e Evropës Lindore", "OEZ": "Ora standarde e Evropës Lindore", "OMSST": "Ora verore e Omskut", "OMST": "Ora standarde e Omskut", "PDT": "Ora verore e Territoreve Amerikane të Bregut të Paqësorit", "PDTM": "Ora verore e Territoreve Meksikane të Bregut të Paqësorit", "PETDT": "Ora verore e Petropavllovsk-Kamçatkës", "PETST": "Ora standarde e Petropavllovsk-Kamçatkës", "PGT": "Ora e Guinesë së Re-Papua", "PHOT": "Ora e Ishujve Feniks", "PKT": "Ora standarde e Pakistanit", "PKT DST": "Ora verore e Pakistanit", "PMDT": "Ora verore e Shën-Pier dhe Mikelon", "PMST": "Ora standarde e Shën-Pier dhe Mikelon", "PONT": "Ora e Ponapeit", "PST": "Ora standarde e Territoreve Amerikane të Bregut të Paqësorit", "PST Philippine": "Ora standarde e Filipineve", "PST Philippine DST": "Ora verore e Filipineve", "PST Pitcairn": "Ora e Pitkernit", "PSTM": "Ora standarde e Territoreve Meksikane të Bregut të Paqësorit", "PWT": "Ora e Palaut", "PYST": "Ora Verore e Paraguait", "PYT": "Ora standarde e Paraguait", "PYT Korea": "Ora e Penianit", "RET": "Ora e Reunionit", "ROTT": "Ora e Rodherës", "SAKST": "Ora verore e Sakalinit", "SAKT": "Ora standarde e Sakalinit", "SAMST": "Ora verore e Samarës", "SAMT": "Ora standarde e Samarës", "SAST": "Ora standarde e Afrikës Jugore", "SBT": "Ora e Ishujve Solomon", "SCT": "Ora e Sejshelleve", "SGT": "Ora e Singaporit", "SLST": "Ora e Lankasë", "SRT": "Ora e Surinamit", "SST Samoa": "Ora standarde e Samoas", "SST Samoa Apia": "Ora standarde e Apias", "SST Samoa Apia DST": "Ora verore e Apias", "SST Samoa DST": "Ora verore e Samoas", "SYOT": "Ora e Sjouit", "TAAF": "Ora e Territoreve Jugore dhe Antarktike Franceze", "TAHT": "Ora e Tahitit", "TJT": "Ora e Taxhikistanit", "TKT": "Ora e Tokelaut", "TLT": "Ora e Timorit Lindor", "TMST": "Ora verore e Turkmenistanit", "TMT": "Ora standarde e Turkmenistanit", "TOST": "Ora verore e Tongës", "TOT": "Ora standarde e Tongës", "TVT": "Ora e Tuvalusë", "TWT": "Ora standarde e Tajpeit", "TWT DST": "Ora verore e Tajpeit", "ULAST": "Ora verore e Ulan-Batorit", "ULAT": "Ora standarde e Ulan-Batorit", "UYST": "Ora verore e Uruguait", "UYT": "Ora standarde e Uruguait", "UZT": "Ora standarde e Uzbekistanit", "UZT DST": "Ora verore e Uzbekistanit", "VET": "Ora e Venezuelës", "VLAST": "Ora verore e Vladivostokut", "VLAT": "Ora standarde e Vladivostokut", "VOLST": "Ora verore e Volgogradit", "VOLT": "Ora standarde e Volgogradit", "VOST": "Ora e Vostokut", "VUT": "Ora standarde e Vanuatusë", "VUT DST": "Ora verore e Vanuatusë", "WAKT": "Ora e Ishullit Uejk", "WARST": "Ora verore e Argjentinës Perëndimore", "WART": "Ora standarde e Argjentinës Perëndimore", "WAST": "Ora e Afrikës Perëndimore", "WAT": "Ora e Afrikës Perëndimore", "WESZ": "Ora verore e Evropës Perëndimore", "WEZ": "Ora standarde e Evropës Perëndimore", "WFT": "Ora e Uollisit dhe Futunës", "WGST": "Ora verore e Grënlandës Perëndimore", "WGT": "Ora standarde e Grënlandës Perëndimore", "WIB": "Ora e Indonezisë Perëndimore", "WIT": "Ora e Indonezisë Lindore", "WITA": "Ora e Indonezisë Qendrore", "YAKST": "Ora verore e Jakutskut", "YAKT": "Ora standarde e Jakutskut", "YEKST": "Ora verore e Ekaterinburgut", "YEKT": "Ora standarde e Ekaterinburgut", "YST": "Ora e Jukonit", "МСК": "Ora standarde e Moskës", "اقتاۋ": "Ora standarde e Aktaut", "اقتاۋ قالاسى": "Ora verore e Aktaut", "اقتوبە": "Ora standarde e Aktobit", "اقتوبە قالاسى": "Ora verore e Aktobit", "الماتى": "Ora standarde e Almatit", "الماتى قالاسى": "Ora verore e Almatit", "باتىس قازاق ەلى": "Ora e Kazakistanit Perëndimor", "شىعىش قازاق ەلى": "Ora e Kazakistanit Lindor", "قازاق ەلى": "Ora e Kazakistanit", "قىرعىزستان": "Ora e Kirgizisë", "قىزىلوردا": "Ora standarde e Kizilordit", "قىزىلوردا قالاسى": "Ora verore e Kizilordit", "∅∅∅": "Ora verore e Perusë"},
	}
}

// Locale returns the current translators string locale
func (sq *sq_AL) Locale() string {
	return sq.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sq_AL'
func (sq *sq_AL) PluralsCardinal() []locales.PluralRule {
	return sq.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sq_AL'
func (sq *sq_AL) PluralsOrdinal() []locales.PluralRule {
	return sq.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sq_AL'
func (sq *sq_AL) PluralsRange() []locales.PluralRule {
	return sq.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sq_AL'
func (sq *sq_AL) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sq_AL'
func (sq *sq_AL) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if n == 1 {
		return locales.PluralRuleOne
	} else if nMod10 == 4 && nMod100 != 14 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sq_AL'
func (sq *sq_AL) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sq.CardinalPluralRule(num1, v1)
	end := sq.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sq *sq_AL) MonthAbbreviated(month time.Month) string {
	if len(sq.monthsAbbreviated) == 0 {
		return ""
	}
	return sq.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sq *sq_AL) MonthsAbbreviated() []string {
	return sq.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sq *sq_AL) MonthNarrow(month time.Month) string {
	if len(sq.monthsNarrow) == 0 {
		return ""
	}
	return sq.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sq *sq_AL) MonthsNarrow() []string {
	return sq.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sq *sq_AL) MonthWide(month time.Month) string {
	if len(sq.monthsWide) == 0 {
		return ""
	}
	return sq.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sq *sq_AL) MonthsWide() []string {
	return sq.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sq.daysAbbreviated) == 0 {
		return ""
	}
	return sq.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sq *sq_AL) WeekdaysAbbreviated() []string {
	return sq.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayNarrow(weekday time.Weekday) string {
	if len(sq.daysNarrow) == 0 {
		return ""
	}
	return sq.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sq *sq_AL) WeekdaysNarrow() []string {
	return sq.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayShort(weekday time.Weekday) string {
	if len(sq.daysShort) == 0 {
		return ""
	}
	return sq.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sq *sq_AL) WeekdaysShort() []string {
	return sq.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayWide(weekday time.Weekday) string {
	if len(sq.daysWide) == 0 {
		return ""
	}
	return sq.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sq *sq_AL) WeekdaysWide() []string {
	return sq.daysWide
}

// Decimal returns the decimal point of number
func (sq *sq_AL) Decimal() string {
	return sq.decimal
}

// Group returns the group of number
func (sq *sq_AL) Group() string {
	return sq.group
}

// Group returns the minus sign of number
func (sq *sq_AL) Minus() string {
	return sq.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sq_AL' and handles both Whole and Real numbers based on 'v'
func (sq *sq_AL) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sq_AL' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sq *sq_AL) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sq.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sq_AL'
func (sq *sq_AL) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sq.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sq.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sq_AL'
// in accounting notation.
func (sq *sq_AL) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sq.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, sq.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sq.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sq.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sq.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sq.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
