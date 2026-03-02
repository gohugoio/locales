package bs_Latn_BA

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type bs_Latn_BA struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
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
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'bs_Latn_BA' locale
func New() locales.Translator {
	return &bs_Latn_BA{
		locale:                 "bs_Latn_BA",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		monthsNarrow:           []string{"", "jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		monthsWide:             []string{"", "januar", "februar", "mart", "april", "maj", "juni", "juli", "august", "septembar", "oktobar", "novembar", "decembar"},
		daysAbbreviated:        []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysNarrow:             []string{"N", "P", "U", "S", "Č", "P", "S"},
		daysWide:               []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
		periodsAbbreviated:     []string{"a.\u202fm.", "p.\u202fm."},
		periodsNarrow:          []string{"prijepodne", "popodne"},
		periodsWide:            []string{"prijepodne", "popodne"},
		erasAbbreviated:        []string{"p. n. e.", "n. e."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "Centralnoaustralijsko ljetno vrijeme", "ACST": "Acre letnje računanje vremena", "ACT": "Acre standardno vreme", "ACWDT": "Australijsko centralnozapadno ljetno vrijeme", "ACWST": "Australijsko centralnozapadno standardno vrijeme", "ADT": "Sjevernoameričko atlantsko ljetno vrijeme", "ADT Arabia": "Arabijsko ljetno vrijeme", "AEDT": "Istočnoaustralijsko ljetno vrijeme", "AEST": "Istočnoaustralijsko standardno vrijeme", "AFT": "Afganistansko vrijeme", "AKDT": "Aljaskansko ljetno vrijeme", "AKST": "Aljaskansko standardno vrijeme", "AMST": "Amazonsko ljetno vrijeme", "AMST Armenia": "Armensko ljetno vrijeme", "AMT": "Amazonsko standardno vrijeme", "AMT Armenia": "Armensko standardno vrijeme", "ANAST": "Anadir letnje računanje vremena", "ANAT": "Anadir standardno vreme", "ARST": "Argentinsko ljetno vrijeme", "ART": "Argentinsko standardno vrijeme", "AST": "Sjevernoameričko atlantsko standardno vrijeme", "AST Arabia": "Arabijsko standardno vrijeme", "AWDT": "Zapadnoaustralijsko ljetno vrijeme", "AWST": "Zapadnoaustralijsko standardno vrijeme", "AZST": "Azerbejdžansko ljetno vrijeme", "AZT": "Azerbejdžansko standardno vrijeme", "BDT Bangladesh": "Bangladeško ljetno vrijeme", "BNT": "Brunejsko vrijeme", "BOT": "Bolivijsko vrijeme", "BRST": "Brazilijsko ljetno vrijeme", "BRT": "Brazilijsko standardno vrijeme", "BST Bangladesh": "Bangladeško standardno vrijeme", "BT": "Butansko vrijeme", "CAST": "CAST", "CAT": "Centralnoafričko vrijeme", "CCT": "Vrijeme na Ostrvima Kokos", "CDT": "Sjevernoameričko centralno ljetno vrijeme", "CHADT": "Čatamsko ljetno vrijeme", "CHAST": "Čatamsko standardno vrijeme", "CHUT": "Čučko vrijeme", "CKT": "Standardno vrijeme na Kukovim ostrvima", "CKT DST": "Poluljetno vrijeme na Kukovim ostrvima", "CLST": "Čileansko ljetno vrijeme", "CLT": "Čileansko standardno vrijeme", "COST": "Kolumbijsko ljetno vrijeme", "COT": "Kolumbijsko standardno vrijeme", "CST": "Sjevernoameričko centralno standardno vrijeme", "CST China": "Kinesko standardno vrijeme", "CST China DST": "Kinesko ljetno vrijeme", "CVST": "Zelenortsko ljetno vrijeme", "CVT": "Zelenortsko standardno vrijeme", "CXT": "Vrijeme na Božićnom Ostrvu", "ChST": "Čamorsko standardno vrijeme", "ChST NMI": "Severna Marijanska Ostrva vreme", "CuDT": "Kubansko ljetno vrijeme", "CuST": "Kubansko standardno vrijeme", "DAVT": "Vrijeme stanice Davis", "DDUT": "Vrijeme stanice Dumont-d’Urville", "EASST": "Uskršnjeostrvsko ljetno vrijeme", "EAST": "Uskršnjeostrvsko standardno vrijeme", "EAT": "Istočnoafričko vrijeme", "ECT": "Ekvadorsko vrijeme", "EDT": "Sjevernoameričko istočno ljetno vrijeme", "EGDT": "Istočnogrenlandsko ljetno vrijeme", "EGST": "Istočnogrenlandsko standardno vrijeme", "EST": "Sjevernoameričko istočno standardno vrijeme", "FEET": "Dalekoistočnoevropsko vrijeme", "FJT": "Standardno vrijeme na Fidžiju", "FJT Summer": "Fidžijsko ljetno vrijeme", "FKST": "Folklandsko ljetno vrijeme", "FKT": "Folklandsko standardno vrijeme", "FNST": "Ljetno vrijeme na ostrvu Fernando di Noronja", "FNT": "Standardno vrijeme na ostrvu Fernando di Noronja", "GALT": "Galapagosko vrijeme", "GAMT": "Gambijersko vrijeme", "GEST": "Gruzijsko ljetno vrijeme", "GET": "Gruzijsko standardno vrijeme", "GFT": "Francuskogvajansko vrijeme", "GIT": "Vrijeme na Gilbertovim ostrvima", "GMT": "Griničko vrijeme", "GNSST": "GNSST", "GNST": "GNST", "GST": "Zalivsko standardno vrijeme", "GST Guam": "Guam standardno vreme", "GYT": "Gvajansko vrijeme", "HADT": "Havajsko-aleućansko ljetno vrijeme", "HAST": "Havajsko-aleućansko standardno vrijeme", "HKST": "Hongkonško ljetno vrijeme", "HKT": "Hongkonško standardno vrijeme", "HOVST": "Hovdsko ljetno vrijeme", "HOVT": "Hovdsko standardno vrijeme", "ICT": "Indokinesko vrijeme", "IDT": "Izraelsko ljetno vrijeme", "IOT": "Vrijeme na Indijskom okeanu", "IRKST": "Irkutsko ljetno vrijeme", "IRKT": "Irkutsko standardno vrijeme", "IRST": "Iransko standardno vrijeme", "IRST DST": "Iransko ljetno vrijeme", "IST": "Indijsko standardno vrijeme", "IST Israel": "Izraelsko standardno vrijeme", "JDT": "Japansko ljetno vrijeme", "JST": "Japansko standardno vrijeme", "KOST": "Vrijeme na Ostrvu Kosrae", "KRAST": "Krasnojarsko ljetno vrijeme", "KRAT": "Krasnojarsko standardno vrijeme", "KST": "Korejsko standardno vrijeme", "KST DST": "Korejsko ljetno vrijeme", "LHDT": "Ljetno vrijeme na Ostrvu Lord Hau", "LHST": "Standardno vrijeme na Ostrvu Lord Hau", "LINT": "Vrijeme na Ostrvima Lajn", "MAGST": "Magadansko ljetno vrijeme", "MAGT": "Magadansko standardno vrijeme", "MART": "Vrijeme na Ostrvima Markiz", "MAWT": "Vrijeme stanice Mawson", "MDT": "Makao letnje računanje vremena", "MESZ": "Centralnoevropsko ljetno vrijeme", "MEZ": "Centralnoevropsko standardno vrijeme", "MHT": "Vrijeme na Maršalovim ostrvima", "MMT": "Mijanmarsko vrijeme", "MSD": "Moskovsko ljetno vrijeme", "MST": "Makao standardno vreme", "MUST": "Mauricijsko ljetno vrijeme", "MUT": "Mauricijsko standardno vrijeme", "MVT": "Maldivsko vrijeme", "MYT": "Malezijsko vrijeme", "NCT": "Novokaledonijsko standardno vrijeme", "NDT": "Njufaundlendsko ljetno vrijeme", "NDT New Caledonia": "Novokaledonijsko ljetno vrijeme", "NFDT": "Norfolško ljetno vrijeme", "NFT": "Norfolško standardno vrijeme", "NOVST": "Novosibirsko ljetno vrijeme", "NOVT": "Novosibirsko standardno vrijeme", "NPT": "Nepalsko vrijeme", "NRT": "Vrijeme na Ostrvu Nauru", "NST": "Njufaundlendsko standardno vrijeme", "NUT": "Vrijeme na Ostrvu Niue", "NZDT": "Novozelandsko ljetno vrijeme", "NZST": "Novozelandsko standardno vrijeme", "OESZ": "Istočnoevropsko ljetno vrijeme", "OEZ": "Istočnoevropsko standardno vrijeme", "OMSST": "Omsko ljetno vrijeme", "OMST": "Omsko standardno vrijeme", "PDT": "Sjevernoameričko pacifičko ljetno vrijeme", "PDTM": "Meksičko pacifičko ljetno vrijeme", "PETDT": "Petropavlovsk-Kamčatski letnje računanje vremena", "PETST": "Petropavlovsk-Kamčatski standardno vreme", "PGT": "Vrijeme na Papui Novoj Gvineji", "PHOT": "Vrijeme na Ostrvima Finiks", "PKT": "Pakistansko standardno vrijeme", "PKT DST": "Pakistansko ljetno vrijeme", "PMDT": "Ljetno vrijeme na Ostrvima Sveti Petar i Mikelon", "PMST": "Standardno vrijeme na Ostrvima Sveti Petar i Mikelon", "PONT": "Vrijeme na Ostrvu Ponape", "PST": "Sjevernoameričko pacifičko standardno vrijeme", "PST Philippine": "Filipinsko standardno vrijeme", "PST Philippine DST": "Filipinsko ljetno vrijeme", "PST Pitcairn": "Vrijeme na Ostrvima Pitkern", "PSTM": "Meksičko pacifičko standardno vrijeme", "PWT": "Vrijeme na Ostrvu Palau", "PYST": "Paragvajsko ljetno vrijeme", "PYT": "Paragvajsko standardno vrijeme", "PYT Korea": "Pjongjanško vrijeme", "RET": "Reunionsko vrijeme", "ROTT": "Vrijeme stanice Rothera", "SAKST": "Sahalinsko ljetno vrijeme", "SAKT": "Sahalinsko standardno vrijeme", "SAMST": "Samara letnje računanje vremena", "SAMT": "Samara standardno vreme", "SAST": "Južnoafričko standardno vrijeme", "SBT": "Vrijeme na Solomonskim ostrvima", "SCT": "Sejšelsko vrijeme", "SGT": "Singapursko standardno vrijeme", "SLST": "Lanka vreme", "SRT": "Surinamsko vrijeme", "SST Samoa": "Samoansko standardno vrijeme", "SST Samoa Apia": "Apijsko standardno vrijeme", "SST Samoa Apia DST": "Apijsko ljetno vrijeme", "SST Samoa DST": "Samoansko ljetno vrijeme", "SYOT": "Vrijeme stanice Syowa", "TAAF": "Vrijeme na Francuskoj Južnoj Teritoriji i Antarktiku", "TAHT": "Tahićansko vrijeme", "TJT": "tadžikistansko vrijeme", "TKT": "Vrijeme na Ostrvu Tokelau", "TLT": "Istočnotimorsko vrijeme", "TMST": "turkmenistansko ljetno vrijeme", "TMT": "turkmenistansko standardno vrijeme", "TOST": "Tongansko ljetno vrijeme", "TOT": "Tongansko standardno vrijeme", "TVT": "Tuvaluansko vrijeme", "TWT": "Tajpejsko standardno vrijeme", "TWT DST": "Tajpejsko ljetno vrijeme", "ULAST": "Ulanbatorsko ljetno vrijeme", "ULAT": "Ulanbatorsko standardno vrijeme", "UYST": "Urugvajsko ljetno vrijeme", "UYT": "Urugvajsko standardno vrijeme", "UZT": "uzbekistansko standardno vrijeme", "UZT DST": "uzbekistansko ljetno vrijeme", "VET": "Venecuelansko vrijeme", "VLAST": "Vladivostočko ljetno vrijeme", "VLAT": "Vladivostočko standardno vrijeme", "VOLST": "Volgogradsko ljetno vrijeme", "VOLT": "Volgogradsko standardno vrijeme", "VOST": "Vrijeme stanice Vostok", "VUT": "Vanuatuansko standardno vrijeme", "VUT DST": "Vanuatuansko ljetno vrijeme", "WAKT": "Vrijeme na Ostrvu Vejk", "WARST": "Zapadnoargentinsko ljetno vrijeme", "WART": "Zapadnoargentinsko standardno vrijeme", "WAST": "Zapadnoafričko vrijeme", "WAT": "Zapadnoafričko vrijeme", "WESZ": "Zapadnoevropsko ljetno vrijeme", "WEZ": "Zapadnoevropsko standardno vrijeme", "WFT": "Vrijeme na Ostrvima Valis i Futuna", "WGST": "Zapadnogrenlandsko ljetno vrijeme", "WGT": "Zapadnogrenlandsko standardno vrijeme", "WIB": "Zapadnoindonezijsko vrijeme", "WIT": "Istočnoindonezijsko vrijeme", "WITA": "Centralnoindonezijsko vrijeme", "YAKST": "Jakutsko ljetno vrijeme", "YAKT": "Jakutsko standardno vrijeme", "YEKST": "Jekaterinburško ljetno vrijeme", "YEKT": "Jekaterinburško standardno vrijeme", "YST": "Jukonsko vrijeme", "МСК": "Moskovsko standardno vrijeme", "اقتاۋ": "Akvtau standardno vreme", "اقتاۋ قالاسى": "Akvtau letnje računanje vremena", "اقتوبە": "Akvtobe standardno vreme", "اقتوبە قالاسى": "Akvtobe letnje računanje vremena", "الماتى": "Almatu standardno vreme", "الماتى قالاسى": "Almatu letnje računanje vremena", "باتىس قازاق ەلى": "zapadnokazahstansko vrijeme", "شىعىش قازاق ەلى": "istočnokazahstansko vrijeme", "قازاق ەلى": "kazahstansko vrijeme", "قىرعىزستان": "kirgistansko vrijeme", "قىزىلوردا": "Kizilorda standardno vreme", "قىزىلوردا قالاسى": "Kizilorda letnje računanje vremena", "∅∅∅": "Azorsko ljetno vrijeme"},
	}
}

// Locale returns the current translators string locale
func (bs *bs_Latn_BA) Locale() string {
	return bs.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bs_Latn_BA'
func (bs *bs_Latn_BA) PluralsCardinal() []locales.PluralRule {
	return bs.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bs_Latn_BA'
func (bs *bs_Latn_BA) PluralsOrdinal() []locales.PluralRule {
	return bs.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bs_Latn_BA'
func (bs *bs_Latn_BA) PluralsRange() []locales.PluralRule {
	return bs.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := bs.CardinalPluralRule(num1, v1)
	end := bs.CardinalPluralRule(num2, v2)

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
func (bs *bs_Latn_BA) MonthAbbreviated(month time.Month) string {
	return bs.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bs *bs_Latn_BA) MonthsAbbreviated() []string {
	return bs.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bs *bs_Latn_BA) MonthNarrow(month time.Month) string {
	return bs.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bs *bs_Latn_BA) MonthsNarrow() []string {
	return bs.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bs *bs_Latn_BA) MonthWide(month time.Month) string {
	return bs.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bs *bs_Latn_BA) MonthsWide() []string {
	return bs.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bs *bs_Latn_BA) WeekdayAbbreviated(weekday time.Weekday) string {
	return bs.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bs *bs_Latn_BA) WeekdaysAbbreviated() []string {
	return bs.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bs *bs_Latn_BA) WeekdayNarrow(weekday time.Weekday) string {
	return bs.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bs *bs_Latn_BA) WeekdaysNarrow() []string {
	return bs.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bs *bs_Latn_BA) WeekdayShort(weekday time.Weekday) string {
	return bs.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bs *bs_Latn_BA) WeekdaysShort() []string {
	return bs.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bs *bs_Latn_BA) WeekdayWide(weekday time.Weekday) string {
	return bs.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bs *bs_Latn_BA) WeekdaysWide() []string {
	return bs.daysWide
}

// Decimal returns the decimal point of number
func (bs *bs_Latn_BA) Decimal() string {
	return bs.decimal
}

// Group returns the group of number
func (bs *bs_Latn_BA) Group() string {
	return bs.group
}

// Group returns the minus sign of number
func (bs *bs_Latn_BA) Minus() string {
	return bs.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bs_Latn_BA' and handles both Whole and Real numbers based on 'v'
func (bs *bs_Latn_BA) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bs_Latn_BA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bs *bs_Latn_BA) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bs.percentSuffix...)

	b = append(b, bs.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bs.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bs_Latn_BA'
// in accounting notation.
func (bs *bs_Latn_BA) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, bs.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bs.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bs.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bs.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bs_Latn_BA'
func (bs *bs_Latn_BA) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bs.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
