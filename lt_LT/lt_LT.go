package lt_LT

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type lt_LT struct {
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

// New returns a new instance of translator for the 'lt_LT' locale
func New() locales.Translator {
	return &lt_LT{
		locale:                 "lt_LT",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "Br", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "zl", "PLZ", "PTE", "Gs", "QAR", "RHD", "ROL", "RON", "RSD", "rb", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "saus.", "vas.", "kov.", "bal.", "geg.", "birž.", "liep.", "rugp.", "rugs.", "spal.", "lapkr.", "gruod."},
		monthsNarrow:           []string{"", "S", "V", "K", "B", "G", "B", "L", "R", "R", "S", "L", "G"},
		monthsWide:             []string{"", "sausio", "vasario", "kovo", "balandžio", "gegužės", "birželio", "liepos", "rugpjūčio", "rugsėjo", "spalio", "lapkričio", "gruodžio"},
		daysAbbreviated:        []string{"sk", "pr", "an", "tr", "kt", "pn", "št"},
		daysNarrow:             []string{"S", "P", "A", "T", "K", "P", "Š"},
		daysShort:              []string{"Sk", "Pr", "An", "Tr", "Kt", "Pn", "Št"},
		daysWide:               []string{"sekmadienis", "pirmadienis", "antradienis", "trečiadienis", "ketvirtadienis", "penktadienis", "šeštadienis"},
		timezones:              map[string]string{"ACDT": "Centrinės Australijos vasaros laikas", "ACST": "Centrinės Australijos žiemos laikas", "ACT": "Ako standartinis laikas", "ACWDT": "Centrinės vakarų Australijos vasaros laikas", "ACWST": "Centrinės vakarų Australijos žiemos laikas", "ADT": "Atlanto vasaros laikas", "ADT Arabia": "Arabijos vasaros laikas", "AEDT": "Rytų Australijos vasaros laikas", "AEST": "Rytų Australijos žiemos laikas", "AFT": "Afganistano laikas", "AKDT": "Aliaskos vasaros laikas", "AKST": "Aliaskos žiemos laikas", "AMST": "Amazonės vasaros laikas", "AMST Armenia": "Armėnijos vasaros laikas", "AMT": "Amazonės žiemos laikas", "AMT Armenia": "Armėnijos žiemos laikas", "ANAST": "Anadyrės vasaros laikas", "ANAT": "Anadyrės žiemos laikas", "ARST": "Argentinos vasaros laikas", "ART": "Argentinos žiemos laikas", "AST": "Atlanto žiemos laikas", "AST Arabia": "Arabijos žiemos laikas", "AWDT": "Vakarų Australijos vasaros laikas", "AWST": "Vakarų Australijos žiemos laikas", "AZST": "Azerbaidžano vasaros laikas", "AZT": "Azerbaidžano žiemos laikas", "BDT Bangladesh": "Bangladešo vasaros laikas", "BNT": "Brunėjaus Darusalamo laikas", "BOT": "Bolivijos laikas", "BRST": "Brazilijos vasaros laikas", "BRT": "Brazilijos žiemos laikas", "BST Bangladesh": "Bangladešo žiemos laikas", "BT": "Butano laikas", "CAST": "Keisio laikas", "CAT": "Centrinės Afrikos laikas", "CCT": "Kokosų Salų laikas", "CDT": "Šiaurės Amerikos centro vasaros laikas", "CHADT": "Čatamo vasaros laikas", "CHAST": "Čatamo žiemos laikas", "CHUT": "Čuko laikas", "CKT": "Kuko Salų žiemos laikas", "CKT DST": "Kuko Salų pusės vasaros laikas", "CLST": "Čilės vasaros laikas", "CLT": "Čilės žiemos laikas", "COST": "Kolumbijos vasaros laikas", "COT": "Kolumbijos žiemos laikas", "CST": "Šiaurės Amerikos centro žiemos laikas", "CST China": "Kinijos žiemos laikas", "CST China DST": "Kinijos vasaros laikas", "CVST": "Žaliojo Kyšulio vasaros laikas", "CVT": "Žaliojo Kyšulio žiemos laikas", "CXT": "Kalėdų Salos laikas", "ChST": "Čamoro laikas", "ChST NMI": "Šiaurės Marianos Salų laikas", "CuDT": "Kubos vasaros laikas", "CuST": "Kubos žiemos laikas", "DAVT": "Deiviso laikas", "DDUT": "Diumono d’Urvilio laikas", "EASST": "Velykų Salos vasaros laikas", "EAST": "Velykų salos žiemos laikas", "EAT": "Rytų Afrikos laikas", "ECT": "Ekvadoro laikas", "EDT": "Šiaurės Amerikos rytų vasaros laikas", "EGDT": "Grenlandijos rytų vasaros laikas", "EGST": "Grenlandijos rytų žiemos laikas", "EST": "Šiaurės Amerikos rytų žiemos laikas", "FEET": "Tolimųjų rytų Europos laikas", "FJT": "Fidžio žiemos laikas", "FJT Summer": "Fidžio vasaros laikas", "FKST": "Folklando Salų vasaros laikas", "FKT": "Folklandų Salų žiemos laikas", "FNST": "Fernando de Noronjos vasaros laikas", "FNT": "Fernando de Noronjos žiemos laikas", "GALT": "Galapagų laikas", "GAMT": "Gambyro laikas", "GEST": "Gruzijos vasaros laikas", "GET": "Gruzijos žiemos laikas", "GFT": "Prancūzijos Gvianos laikas", "GIT": "Gilberto Salų laikas", "GMT": "Grinvičo laikas", "GNSST": "GNSST", "GNST": "GNST", "GST": "Pietų Džordžijos laikas", "GST Guam": "Guamo laikas", "GYT": "Gajanos laikas", "HADT": "Havajų–Aleutų vasaros laikas", "HAST": "Havajų–Aleutų žiemos laikas", "HKST": "Honkongo vasaros laikas", "HKT": "Honkongo žiemos laikas", "HOVST": "Hovdo vasaros laikas", "HOVT": "Hovdo žiemos laikas", "ICT": "Indokinijos laikas", "IDT": "Izraelio vasaros laikas", "IOT": "Indijos vandenyno laikas", "IRKST": "Irkutsko vasaros laikas", "IRKT": "Irkutsko žiemos laikas", "IRST": "Irano žiemos laikas", "IRST DST": "Irano vasaros laikas", "IST": "Indijos laikas", "IST Israel": "Izraelio žiemos laikas", "JDT": "Japonijos vasaros laikas", "JST": "Japonijos žiemos laikas", "KOST": "Kosrajė laikas", "KRAST": "Krasnojarsko vasaros laikas", "KRAT": "Krasnojarsko žiemos laikas", "KST": "Korėjos žiemos laikas", "KST DST": "Korėjos vasaros laikas", "LHDT": "Lordo Hau vasaros laikas", "LHST": "Lordo Hau žiemos laikas", "LINT": "Laino Salų laikas", "MAGST": "Magadano vasaros laikas", "MAGT": "Magadano žiemos laikas", "MART": "Markizo Salų laikas", "MAWT": "Mosono laikas", "MDT": "Šiaurės Amerikos kalnų vasaros laikas", "MESZ": "Vidurio Europos vasaros laikas", "MEZ": "Vidurio Europos žiemos laikas", "MHT": "Maršalo Salų laikas", "MMT": "Mianmaro laikas", "MSD": "Maskvos vasaros laikas", "MST": "Šiaurės Amerikos kalnų žiemos laikas", "MUST": "Mauricijaus vasaros laikas", "MUT": "Mauricijaus žiemos laikas", "MVT": "Maldyvų laikas", "MYT": "Malaizijos laikas", "NCT": "Naujosios Kaledonijos žiemos laikas", "NDT": "Niufaundlendo vasaros laikas", "NDT New Caledonia": "Naujosios Kaledonijos vasaros laikas", "NFDT": "Norfolko Salų vasaros laikas", "NFT": "Norfolko Salų žiemos laikas", "NOVST": "Novosibirsko vasaros laikas", "NOVT": "Novosibirsko žiemos laikas", "NPT": "Nepalo laikas", "NRT": "Nauru laikas", "NST": "Niufaundlendo žiemos laikas", "NUT": "Niujė laikas", "NZDT": "Naujosios Zelandijos vasaros laikas", "NZST": "Naujosios Zelandijos žiemos laikas", "OESZ": "Rytų Europos vasaros laikas", "OEZ": "Rytų Europos žiemos laikas", "OMSST": "Omsko vasaros laikas", "OMST": "Omsko žiemos laikas", "PDT": "Šiaurės Amerikos Ramiojo vandenyno vasaros laikas", "PDTM": "Meksikos Ramiojo vandenyno vasaros laikas", "PETDT": "Kamčiatkos Petropavlovsko vasaros laikas", "PETST": "Kamčiatkos Petropavlovsko žiemos laikas", "PGT": "Papua Naujosios Gvinėjos laikas", "PHOT": "Fenikso Salų laikas", "PKT": "Pakistano žiemos laikas", "PKT DST": "Pakistano vasaros laikas", "PMDT": "Sen Pjero ir Mikelono vasaros laikas", "PMST": "Sen Pjero ir Mikelono žiemos laikas", "PONT": "Ponapės laikas", "PST": "Šiaurės Amerikos Ramiojo vandenyno žiemos laikas", "PST Philippine": "Filipinų žiemos laikas", "PST Philippine DST": "Filipinų vasaros laikas", "PST Pitcairn": "Pitkerno laikas", "PSTM": "Meksikos Ramiojo vandenyno žiemos laikas", "PWT": "Palau laikas", "PYST": "Paragvajaus vasaros laikas", "PYT": "Paragvajaus žiemos laikas", "PYT Korea": "Pchenjano laikas", "RET": "Reunjono laikas", "ROTT": "Roteros laikas", "SAKST": "Sachalino vasaros laikas", "SAKT": "Sachalino žiemos laikas", "SAMST": "Samaros vasaros laikas", "SAMT": "Samaros žiemos laikas", "SAST": "Pietų Afrikos laikas", "SBT": "Saliamono Salų laikas", "SCT": "Seišelių laikas", "SGT": "Singapūro laikas", "SLST": "Lankos laikas", "SRT": "Surinamo laikas", "SST Samoa": "Samoa žiemos laikas", "SST Samoa Apia": "Apijos žiemos laikas", "SST Samoa Apia DST": "Apijos vasaros laikas", "SST Samoa DST": "Samoa vasaros laikas", "SYOT": "Siovos laikas", "TAAF": "Pietų Prancūzijos ir antarktinis laikas", "TAHT": "Tahičio laikas", "TJT": "Tadžikistano laikas", "TKT": "Tokelau laikas", "TLT": "Rytų Timoro laikas", "TMST": "Turkmėnistano vasaros laikas", "TMT": "Turkmėnistano žiemos laikas", "TOST": "Tongos vasaros laikas", "TOT": "Tongos žiemos laikas", "TVT": "Tuvalu laikas", "TWT": "Taipėjaus žiemos laikas", "TWT DST": "Taipėjaus vasaros laikas", "ULAST": "Ulan Batoro vasaros laikas", "ULAT": "Ulan Batoro žiemos laikas", "UYST": "Urugvajaus vasaros laikas", "UYT": "Urugvajaus žiemos laikas", "UZT": "Uzbekistano žiemos laikas", "UZT DST": "Uzbekistano vasaros laikas", "VET": "Venesuelos laikas", "VLAST": "Vladivostoko vasaros laikas", "VLAT": "Vladivostoko žiemos laikas", "VOLST": "Volgogrado vasaros laikas", "VOLT": "Volgogrado žiemos laikas", "VOST": "Vostoko laikas", "VUT": "Vanuatu žiemos laikas", "VUT DST": "Vanuatu vasaros laikas", "WAKT": "Veiko Salos laikas", "WARST": "Vakarų Argentinos vasaros laikas", "WART": "Vakarų Argentinos žiemos laikas", "WAST": "Vakarų Afrikos laikas", "WAT": "Vakarų Afrikos laikas", "WESZ": "Vakarų Europos vasaros laikas", "WEZ": "Vakarų Europos žiemos laikas", "WFT": "Voliso ir Futūnos laikas", "WGST": "Grenlandijos vakarų vasaros laikas", "WGT": "Grenlandijos vakarų žiemos laikas", "WIB": "Vakarų Indonezijos laikas", "WIT": "Rytų Indonezijos laikas", "WITA": "Centrinės Indonezijos laikas", "YAKST": "Jakutsko vasaros laikas", "YAKT": "Jakutsko žiemos laikas", "YEKST": "Jekaterinburgo vasaros laikas", "YEKT": "Jekaterinburgo žiemos laikas", "YST": "Jukono laikas", "МСК": "Maskvos žiemos laikas", "اقتاۋ": "Aktau žiemos laikas", "اقتاۋ قالاسى": "Aktau vasaros laikas", "اقتوبە": "Aktobės žiemos laikas", "اقتوبە قالاسى": "Aktobės vasaros laikas", "الماتى": "Almatos žiemos laikas", "الماتى قالاسى": "Almatos vasaros laikas", "باتىس قازاق ەلى": "Vakarų Kazachstano laikas", "شىعىش قازاق ەلى": "Rytų Kazachstano laikas", "قازاق ەلى": "Kazachstano laikas", "قىرعىزستان": "Kirgistano laikas", "قىزىلوردا": "Kyzylordos žiemos laikas", "قىزىلوردا قالاسى": "Kyzylordos vasaros laikas", "∅∅∅": "Peru vasaros laikas"},
	}
}

// Locale returns the current translators string locale
func (lt *lt_LT) Locale() string {
	return lt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lt_LT'
func (lt *lt_LT) PluralsCardinal() []locales.PluralRule {
	return lt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lt_LT'
func (lt *lt_LT) PluralsOrdinal() []locales.PluralRule {
	return lt.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lt_LT'
func (lt *lt_LT) PluralsRange() []locales.PluralRule {
	return lt.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lt_LT'
func (lt *lt_LT) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	f := locales.F(n, v)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 1 && (nMod100 < 11 || nMod100 > 19) {
		return locales.PluralRuleOne
	} else if nMod10 >= 2 && nMod10 <= 9 && (nMod100 < 11 || nMod100 > 19) {
		return locales.PluralRuleFew
	} else if f != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lt_LT'
func (lt *lt_LT) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lt_LT'
func (lt *lt_LT) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := lt.CardinalPluralRule(num1, v1)
	end := lt.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
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
func (lt *lt_LT) MonthAbbreviated(month time.Month) string {
	if len(lt.monthsAbbreviated) == 0 {
		return ""
	}
	return lt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lt *lt_LT) MonthsAbbreviated() []string {
	return lt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lt *lt_LT) MonthNarrow(month time.Month) string {
	if len(lt.monthsNarrow) == 0 {
		return ""
	}
	return lt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lt *lt_LT) MonthsNarrow() []string {
	return lt.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lt *lt_LT) MonthWide(month time.Month) string {
	if len(lt.monthsWide) == 0 {
		return ""
	}
	return lt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lt *lt_LT) MonthsWide() []string {
	return lt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lt *lt_LT) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(lt.daysAbbreviated) == 0 {
		return ""
	}
	return lt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lt *lt_LT) WeekdaysAbbreviated() []string {
	return lt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lt *lt_LT) WeekdayNarrow(weekday time.Weekday) string {
	if len(lt.daysNarrow) == 0 {
		return ""
	}
	return lt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lt *lt_LT) WeekdaysNarrow() []string {
	return lt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lt *lt_LT) WeekdayShort(weekday time.Weekday) string {
	if len(lt.daysShort) == 0 {
		return ""
	}
	return lt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lt *lt_LT) WeekdaysShort() []string {
	return lt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lt *lt_LT) WeekdayWide(weekday time.Weekday) string {
	if len(lt.daysWide) == 0 {
		return ""
	}
	return lt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lt *lt_LT) WeekdaysWide() []string {
	return lt.daysWide
}

// Decimal returns the decimal point of number
func (lt *lt_LT) Decimal() string {
	return lt.decimal
}

// Group returns the group of number
func (lt *lt_LT) Group() string {
	return lt.group
}

// Group returns the minus sign of number
func (lt *lt_LT) Minus() string {
	return lt.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lt_LT' and handles both Whole and Real numbers based on 'v'
func (lt *lt_LT) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lt.group) - 1; j >= 0; j-- {
					b = append(b, lt.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lt_LT' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lt *lt_LT) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lt.percentSuffix...)

	b = append(b, lt.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lt_LT'
func (lt *lt_LT) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lt.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lt.group) - 1; j >= 0; j-- {
					b = append(b, lt.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, lt.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lt_LT'
// in accounting notation.
func (lt *lt_LT) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lt.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lt.group) - 1; j >= 0; j-- {
					b = append(b, lt.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, lt.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, lt.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0x6d}...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0x6d}...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64}...)
	b = append(b, []byte{0x2e, 0x2c, 0x20}...)
	b = append(b, lt.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'lt_LT'
func (lt *lt_LT) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
