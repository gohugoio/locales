package is_IS

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type is_IS struct {
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

// New returns a new instance of translator for the 'is_IS' locale
func New() locales.Translator {
	return &is_IS{
		locale:                 "is_IS",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mar.", "apr.", "maí", "jún.", "júl.", "ágú.", "sep.", "okt.", "nóv.", "des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "Á", "S", "O", "N", "D"},
		monthsWide:             []string{"", "janúar", "febrúar", "mars", "apríl", "maí", "júní", "júlí", "ágúst", "september", "október", "nóvember", "desember"},
		daysAbbreviated:        []string{"sun.", "mán.", "þri.", "mið.", "fim.", "fös.", "lau."},
		daysNarrow:             []string{"S", "M", "Þ", "M", "F", "F", "L"},
		daysShort:              []string{"su.", "má.", "þr.", "mi.", "fi.", "fö.", "la."},
		daysWide:               []string{"sunnudagur", "mánudagur", "þriðjudagur", "miðvikudagur", "fimmtudagur", "föstudagur", "laugardagur"},
		timezones:              map[string]string{"ACDT": "Sumartími í Mið-Ástralíu", "ACST": "Staðaltími í Mið-Ástralíu", "ACT": "ACT", "ACWDT": "Sumartími í miðvesturhluta Ástralíu", "ACWST": "Staðaltími í miðvesturhluta Ástralíu", "ADT": "Sumartími á Atlantshafssvæðinu", "ADT Arabia": "Sumartími í Arabíu", "AEDT": "Sumartími í Austur-Ástralíu", "AEST": "Staðaltími í Austur-Ástralíu", "AFT": "Afganistantími", "AKDT": "Sumartími í Alaska", "AKST": "Staðaltími í Alaska", "AMST": "Sumartími á Amasónsvæðinu", "AMST Armenia": "Sumartími í Armeníu", "AMT": "Staðaltími á Amasónsvæðinu", "AMT Armenia": "Staðaltími í Armeníu", "ANAST": "Sumartími í Anadyr", "ANAT": "Staðaltími í Anadyr", "ARST": "Sumartími í Argentínu", "ART": "Staðaltími í Argentínu", "AST": "Staðaltími á Atlantshafssvæðinu", "AST Arabia": "Staðaltími í Arabíu", "AWDT": "Sumartími í Vestur-Ástralíu", "AWST": "Staðaltími í Vestur-Ástralíu", "AZST": "Sumartími í Aserbaídsjan", "AZT": "Staðaltími í Aserbaídsjan", "BDT Bangladesh": "Sumartími í Bangladess", "BNT": "Brúneitími", "BOT": "Bólivíutími", "BRST": "Sumartími í Brasilíu", "BRT": "Staðaltími í Brasilíu", "BST Bangladesh": "Staðaltími í Bangladess", "BT": "Bútantími", "CAST": "CAST", "CAT": "Mið-Afríkutími", "CCT": "Kókoseyjatími", "CDT": "Sumartími í miðhluta Bandaríkjanna og Kanada", "CHADT": "Sumartími í Chatham", "CHAST": "Staðaltími í Chatham", "CHUT": "Chuuk-tími", "CKT": "Staðaltími á Cooks-eyjum", "CKT DST": "Hálfsumartími á Cooks-eyjum", "CLST": "Sumartími í Síle", "CLT": "Staðaltími í Síle", "COST": "Sumartími í Kólumbíu", "COT": "Staðaltími í Kólumbíu", "CST": "Staðaltími í miðhluta Bandaríkjanna og Kanada", "CST China": "Staðaltími í Kína", "CST China DST": "Sumartími í Kína", "CVST": "Sumartími á Grænhöfðaeyjum", "CVT": "Staðaltími á Grænhöfðaeyjum", "CXT": "Jólaeyjartími", "ChST": "Chamorro-staðaltími", "ChST NMI": "ChST NMI", "CuDT": "Sumartími á Kúbu", "CuST": "Staðaltími á Kúbu", "DAVT": "Davis-tími", "DDUT": "Tími á Dumont-d’Urville", "EASST": "Sumartími á Páskaeyju", "EAST": "Staðaltími á Páskaeyju", "EAT": "Austur-Afríkutími", "ECT": "Ekvadortími", "EDT": "Sumartími í austurhluta Bandaríkjanna og Kanada", "EGDT": "Sumartími á Austur-Grænlandi", "EGST": "Staðaltími á Austur-Grænlandi", "EST": "Staðaltími í austurhluta Bandaríkjanna og Kanada", "FEET": "Staðartími Kalíníngrad", "FJT": "Staðaltími á Fídjíeyjum", "FJT Summer": "Sumartími á Fídjíeyjum", "FKST": "Sumartími á Falklandseyjum", "FKT": "Staðaltími á Falklandseyjum", "FNST": "Sumartími í Fernando de Noronha", "FNT": "Staðaltími í Fernando de Noronha", "GALT": "Galapagos-tími", "GAMT": "Gambier-tími", "GEST": "Sumartími í Georgíu", "GET": "Staðaltími í Georgíu", "GFT": "Tími í Frönsku Gvæjana", "GIT": "Tími á Gilbert-eyjum", "GMT": "Greenwich-staðaltími", "GNSST": "GNSST", "GNST": "GNST", "GST": "Suður-Georgíutími", "GST Guam": "GST Guam", "GYT": "Gvæjanatími", "HADT": "Sumartími á Havaí og Aleúta", "HAST": "Staðaltími á Havaí og Aleúta", "HKST": "Sumartími í Hong Kong", "HKT": "Staðaltími í Hong Kong", "HOVST": "Sumartími í Hovd", "HOVT": "Staðaltími í Hovd", "ICT": "Indókínatími", "IDT": "Sumartími í Ísrael", "IOT": "Indlandshafstími", "IRKST": "Sumartími í Irkutsk", "IRKT": "Staðaltími í Irkutsk", "IRST": "Staðaltími í Íran", "IRST DST": "Sumartími í Íran", "IST": "Indlandstími", "IST Israel": "Staðaltími í Ísrael", "JDT": "Sumartími í Japan", "JST": "Staðaltími í Japan", "KOST": "Kosrae-tími", "KRAST": "Sumartími í Krasnoyarsk", "KRAT": "Staðaltími í Krasnoyarsk", "KST": "Staðaltími í Kóreu", "KST DST": "Sumartími í Kóreu", "LHDT": "Sumartími á Lord Howe-eyju", "LHST": "Staðaltími á Lord Howe-eyju", "LINT": "Línueyja-tími", "MAGST": "Sumartími í Magadan", "MAGT": "Staðaltími í Magadan", "MART": "Tími á Markgreifafrúreyjum", "MAWT": "Mawson-tími", "MDT": "MDT", "MESZ": "Sumartími í Mið-Evrópu", "MEZ": "Staðaltími í Mið-Evrópu", "MHT": "Tími á Marshall-eyjum", "MMT": "Mjanmar-tími", "MSD": "Sumartími í Moskvu", "MST": "MST", "MUST": "Sumartími á Máritíus", "MUT": "Staðaltími á Máritíus", "MVT": "Maldíveyja-tími", "MYT": "Malasíutími", "NCT": "Staðaltími í Nýju-Kaledóníu", "NDT": "Sumartími á Nýfundnalandi", "NDT New Caledonia": "Sumartími í Nýju-Kaledóníu", "NFDT": "Sumartími á Norfolk-eyju", "NFT": "Staðaltími á Norfolk-eyju", "NOVST": "Sumartími í Novosibirsk", "NOVT": "Staðaltími í Novosibirsk", "NPT": "Nepaltími", "NRT": "Nárú-tími", "NST": "Staðaltími á Nýfundnalandi", "NUT": "Niue-tími", "NZDT": "Sumartími á Nýja-Sjálandi", "NZST": "Staðaltími á Nýja-Sjálandi", "OESZ": "Sumartími í Austur-Evrópu", "OEZ": "Staðaltími í Austur-Evrópu", "OMSST": "Sumartími í Omsk", "OMST": "Staðaltími í Omsk", "PDT": "Sumartími á Kyrrahafssvæðinu", "PDTM": "Sumartími í Mexíkó á Kyrrahafssvæðinu", "PETDT": "Sumartími í Petropavlovsk-Kamchatski", "PETST": "Staðaltími í Petropavlovsk-Kamchatski", "PGT": "Tími á Papúa Nýju-Gíneu", "PHOT": "Fönixeyjatími", "PKT": "Staðaltími í Pakistan", "PKT DST": "Sumartími í Pakistan", "PMDT": "Sumartími á Sankti Pierre og Miquelon", "PMST": "Staðaltími á Sankti Pierre og Miquelon", "PONT": "Ponape-tími", "PST": "Staðaltími á Kyrrahafssvæðinu", "PST Philippine": "Staðaltími á Filippseyjum", "PST Philippine DST": "Sumartími á Filippseyjum", "PST Pitcairn": "Pitcairn-tími", "PSTM": "Staðaltími í Mexíkó á Kyrrahafssvæðinu", "PWT": "Palátími", "PYST": "Sumartími í Paragvæ", "PYT": "Staðaltími í Paragvæ", "PYT Korea": "Tími í Pjongjang", "RET": "Réunion-tími", "ROTT": "Rothera-tími", "SAKST": "Sumartími á Sakhalin", "SAKT": "Staðaltími á Sakhalin", "SAMST": "Sumartími í Samara", "SAMT": "Staðaltími í Samara", "SAST": "Suður-Afríkutími", "SBT": "Salómonseyjatími", "SCT": "Seychelles-eyjatími", "SGT": "Singapúrtími", "SLST": "SLST", "SRT": "Súrinamtími", "SST Samoa": "Staðaltími á Samóa", "SST Samoa Apia": "Staðaltími í Apía", "SST Samoa Apia DST": "Sumartími í Apía", "SST Samoa DST": "Sumartími á Samóa", "SYOT": "Syowa-tími", "TAAF": "Tími á frönsku suðurhafssvæðum og Suðurskautslandssvæði", "TAHT": "Tahítí-tími", "TJT": "Tadsjíkistan-tími", "TKT": "Tókelá-tími", "TLT": "Tími á Tímor-Leste", "TMST": "Sumartími í Túrkmenistan", "TMT": "Staðaltími í Túrkmenistan", "TOST": "Sumartími á Tonga", "TOT": "Staðaltími á Tonga", "TVT": "Túvalútími", "TWT": "Staðaltími í Taipei", "TWT DST": "Sumartími í Taipei", "ULAST": "Sumartími í Úlan Bator", "ULAT": "Staðaltími í Úlan Bator", "UYST": "Sumartími í Úrúgvæ", "UYT": "Staðaltími í Úrúgvæ", "UZT": "Staðaltími í Úsbekistan", "UZT DST": "Sumartími í Úsbekistan", "VET": "Venesúelatími", "VLAST": "Sumartími í Vladivostok", "VLAT": "Staðaltími í Vladivostok", "VOLST": "Sumartími í Volgograd", "VOLT": "Staðaltími í Volgograd", "VOST": "Vostok-tími", "VUT": "Staðaltími á Vanúatú", "VUT DST": "Sumartími á Vanúatú", "WAKT": "Tími á Wake-eyju", "WARST": "Sumartími í Vestur-Argentínu", "WART": "Staðaltími í Vestur-Argentínu", "WAST": "Vestur-Afríkutími", "WAT": "Vestur-Afríkutími", "WESZ": "Sumartími í Vestur-Evrópu", "WEZ": "Staðaltími í Vestur-Evrópu", "WFT": "Tími á Wallis- og Fútúnaeyjum", "WGST": "Sumartími á Vestur-Grænlandi", "WGT": "Staðaltími á Vestur-Grænlandi", "WIB": "Vestur-Indónesíutími", "WIT": "Austur-Indónesíutími", "WITA": "Mið-Indónesíutími", "YAKST": "Sumartími í Yakutsk", "YAKT": "Staðaltími í Yakutsk", "YEKST": "Sumartími í Yekaterinburg", "YEKT": "Staðaltími í Yekaterinborg", "YST": "Tími í Júkon", "МСК": "Staðaltími í Moskvu", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Tími í Vestur-Kasakstan", "شىعىش قازاق ەلى": "Tími í Austur-Kasakstan", "قازاق ەلى": "Tími í Kasakstan", "قىرعىزستان": "Kirgistan-tími", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Sumartími í Perú"},
	}
}

// Locale returns the current translators string locale
func (is *is_IS) Locale() string {
	return is.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'is_IS'
func (is *is_IS) PluralsCardinal() []locales.PluralRule {
	return is.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'is_IS'
func (is *is_IS) PluralsOrdinal() []locales.PluralRule {
	return is.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'is_IS'
func (is *is_IS) PluralsRange() []locales.PluralRule {
	return is.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'is_IS'
func (is *is_IS) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	t := locales.T(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	tMod100 := t % 100
	tMod10 := t % 10

	if (t == 0 && iMod10 == 1 && iMod100 != 11) || (tMod10 == 1 && tMod100 != 11) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'is_IS'
func (is *is_IS) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'is_IS'
func (is *is_IS) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := is.CardinalPluralRule(num1, v1)
	end := is.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (is *is_IS) MonthAbbreviated(month time.Month) string {
	return is.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (is *is_IS) MonthsAbbreviated() []string {
	return is.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (is *is_IS) MonthNarrow(month time.Month) string {
	return is.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (is *is_IS) MonthsNarrow() []string {
	return is.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (is *is_IS) MonthWide(month time.Month) string {
	return is.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (is *is_IS) MonthsWide() []string {
	return is.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (is *is_IS) WeekdayAbbreviated(weekday time.Weekday) string {
	return is.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (is *is_IS) WeekdaysAbbreviated() []string {
	return is.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (is *is_IS) WeekdayNarrow(weekday time.Weekday) string {
	return is.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (is *is_IS) WeekdaysNarrow() []string {
	return is.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (is *is_IS) WeekdayShort(weekday time.Weekday) string {
	return is.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (is *is_IS) WeekdaysShort() []string {
	return is.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (is *is_IS) WeekdayWide(weekday time.Weekday) string {
	return is.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (is *is_IS) WeekdaysWide() []string {
	return is.daysWide
}

// Decimal returns the decimal point of number
func (is *is_IS) Decimal() string {
	return is.decimal
}

// Group returns the group of number
func (is *is_IS) Group() string {
	return is.group
}

// Group returns the minus sign of number
func (is *is_IS) Minus() string {
	return is.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'is_IS' and handles both Whole and Real numbers based on 'v'
func (is *is_IS) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, is.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, is.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, is.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'is_IS' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (is *is_IS) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, is.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, is.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, is.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'is_IS'
func (is *is_IS) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := is.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, is.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, is.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, is.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, is.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, is.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'is_IS'
// in accounting notation.
func (is *is_IS) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := is.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, is.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, is.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, is.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, is.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, is.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, is.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'is_IS'
func (is *is_IS) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'is_IS'
func (is *is_IS) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, is.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'is_IS'
func (is *is_IS) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, is.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'is_IS'
func (is *is_IS) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, is.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, is.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'is_IS'
func (is *is_IS) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, is.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'is_IS'
func (is *is_IS) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, is.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, is.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'is_IS'
func (is *is_IS) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, is.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, is.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'is_IS'
func (is *is_IS) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, is.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, is.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := is.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
