package se_FI

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type se_FI struct {
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

// New returns a new instance of translator for the 'se_FI' locale
func New() locales.Translator {
	return &se_FI{
		locale:                 "se_FI",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "cuoŋ"},
		monthsNarrow:           []string{"", "O", "G", "N", "C", "M", "G", "S", "B", "Č", "G", "S", "J"},
		monthsWide:             []string{"", "ođđajagemánnu", "guovvamánnu", "njukčamánnu", "cuoŋománnu", "miessemánnu", "geassemánnu", "suoidnemánnu", "borgemánnu", "čakčamánnu", "golggotmánnu", "skábmamánnu", "juovlamánnu"},
		daysAbbreviated:        []string{"so", "má", "di", "ga", "du", "be", "lá"},
		daysNarrow:             []string{"M", "D"},
		daysShort:              []string{"so", "má", "di", "ga", "du", "be", "lá"},
		daysWide:               []string{"mánnodat", "disdat", "gaskavahkku", "duorastat", "bearjadat", "lávvordat"},
		periodsAbbreviated:     []string{"ib", "eb"},
		periodsNarrow:          []string{"i", "e"},
		periodsWide:            []string{"ib", "eb"},
		erasAbbreviated:        []string{"oKr.", "mKr."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ovdal Kristusa", "maŋŋel Kristusa"},
		timezones:              map[string]string{"ACDT": "Gaska-Austrália geasseáigi", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Gaska-Austrália oarjjabeali geasseáigi", "ACWST": "Gaska-Austrália oarjjabeali dálveáigi", "ADT": "atlántalaš geasseáigi", "ADT Arabia": "Arábia geasseáigi", "AEDT": "Nuorta-Austrália geasseáigi", "AEST": "Nuorta-Austrália dálveáigi", "AFT": "Afganisthana áigi", "AKDT": "Alaska geasseáigi", "AKST": "Alaska dálveáigi", "AMST": "Amazona geasseáigi", "AMST Armenia": "Armenia geasseáigi", "AMT": "Amazona dálveáigi", "AMT Armenia": "Armenia dálveáigi", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Argentina geasseáigi", "ART": "Argentina dálveáigi", "AST": "atlántalaš dálveáigi", "AST Arabia": "Arábia dálveáigi", "AWDT": "Oarje-Austrália geasseáigi", "AWST": "Oarje-Austrália dálveáigi", "AZST": "Aserbaižana geasseáigi", "AZT": "Aserbaižana dálveáigi", "BDT Bangladesh": "Bangladesha geasseáigi", "BNT": "Brunei Darussalama áigi", "BOT": "Bolivia áigi", "BRST": "Brasilia geasseáigi", "BRT": "Brasilia dálveáigi", "BST Bangladesh": "Bangladesha dálveáigi", "BT": "Bhutana áigi", "CAST": "CAST", "CAT": "Gaska-Afrihká áigi", "CCT": "Kokossulloid áigi", "CDT": "dábálaš geasseáigi", "CHADT": "Chathama geasseáigi", "CHAST": "Chathama dálveáigi", "CHUT": "Chuuka áigi", "CKT": "Cooksulloid dálveáigi", "CKT DST": "Cooksulloid geasi beallemuttu áigi", "CLST": "Chile geasseáigi", "CLT": "Chile dálveáigi", "COST": "Colombia geasseáigi", "COT": "Colombia dálveáigi", "CST": "dábálaš dálveáigi", "CST China": "Kiinná dálveáigi", "CST China DST": "Kiinná geasseáigi", "CVST": "Kap Verde geasseáigi", "CVT": "Kap Verde dálveáigi", "CXT": "Juovlasullo áigi", "ChST": "Čamorro dálveáigi", "ChST NMI": "ChST NMI", "CuDT": "Cuba geasseáigi", "CuST": "Cuba dálveáigi", "DAVT": "Davisa áigi", "DDUT": "Dumont-d’Urville áigi", "EASST": "Beassášsullo geasseáigi", "EAST": "Beassášsullo dálveáigi", "EAT": "Nuorta-Afrihká áigi", "ECT": "Ecuadora áigi", "EDT": "geasseáigi nuortan", "EGDT": "Nuorta-Ruonáeatnama geasseáigi", "EGST": "Nuorta-Ruonáeatnama dálveáigi", "EST": "dálveáigi nuortan", "FEET": "Gáiddus-Nuortti eurohpalaš áigi", "FJT": "Fiji dálveáigi", "FJT Summer": "Fiji geasseáigi", "FKST": "Falklandsulluid geasseáigi", "FKT": "Falklandsulluid dálveáigi", "FNST": "Fernando de Noronha geasseáigi", "FNT": "Fernando de Noronha dálveáigi", "GALT": "Galapagosa áigi", "GAMT": "Gambiera áigi", "GEST": "Georgia geasseáigi", "GET": "Georgia dálveáigi", "GFT": "Frankriikka Guyana áigi", "GIT": "Gilbertsulloid áigi", "GMT": "Greenwicha áigi", "GNSST": "GNSST", "GNST": "GNST", "GST": "Lulli-Georgia áigi", "GST Guam": "GST Guam", "GYT": "Guyana áigi", "HADT": "Hawaii-aleuhtalaš geasseáigi", "HAST": "Hawaii-aleuhtalaš dálveáigi", "HKST": "Hong Konga geasseáigi", "HKT": "Hong Konga dálveáigi", "HOVST": "Hovda geasseáigi", "HOVT": "Hovda dálveáigi", "ICT": "Indokiinná áigi", "IDT": "Israela geasseáigi", "IOT": "Indiaábi áigi", "IRKST": "Irkucka geasseáigi", "IRKT": "Irkucka dálveáigi", "IRST": "Irana dálveáigi", "IRST DST": "Irana geasseáigi", "IST": "India dálveáigi", "IST Israel": "Israela dálveáigi", "JDT": "Japána geasseáigi", "JST": "Japána dálveáigi", "KOST": "Kosraea áigi", "KRAST": "Krasnojarska geasseáigi", "KRAT": "Krasnojarska dálveáigi", "KST": "Korea dálveáigi", "KST DST": "Korea geasseáigi", "LHDT": "Lord Howe geasseáigi", "LHST": "Lord Howe dálveáigi", "LINT": "Linesulloid áigi", "MAGST": "Magadana geasseáigi", "MAGT": "Magadana dálveáigi", "MART": "Marquesasiid áigi", "MAWT": "Mawsona áigi", "MDT": "geasseduottaráigi", "MESZ": "Gaska-Eurohpá geasseáigi", "MEZ": "Gaska-Eurohpá dálveáigi", "MHT": "Marshallsulloid áigi", "MMT": "Myanmara áigi", "MSD": "Moskva geasseáigi", "MST": "dálveduottaráigi", "MUST": "Mauritiusa geasseáigi", "MUT": "Mauritiusa dálveáigi", "MVT": "Malediivvaid áigi", "MYT": "Malesia áigi", "NCT": "Ođđa-Kaledonia dálveáigi", "NDT": "Newfoundlanda geasseáigi", "NDT New Caledonia": "Ođđa-Kaledonia geasseáigi", "NFDT": "Norfolksullo áigi", "NFT": "Norfolksullo áigi", "NOVST": "Novosibirska geasseáigi", "NOVT": "Novosibirska dálveáigi", "NPT": "Nepala áigi", "NRT": "Nauru áigi", "NST": "Newfoundlanda dálveáigi", "NUT": "Niuea áigi", "NZDT": "Ođđa-Selánda geasseáigi", "NZST": "Ođđa-Selánda dálveáigi", "OESZ": "Nuorta-Eurohpa geasseáigi", "OEZ": "Nuorta-Eurohpa dálveáigi", "OMSST": "Omska geasseáigi", "OMST": "Omska dálveáigi", "PDT": "Jaskesábi geasseáigi", "PDTM": "Meksiko Jáskesábi geasseáigi", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papua Ođđa-Guinea áigi", "PHOT": "Phoenixsulloid áigi", "PKT": "Pakistana dálveáigi", "PKT DST": "Pakistana geasseáigi", "PMDT": "St. Pierre & Miquelo geasseáigi", "PMST": "St. Pierre & Miquelo dálveáigi", "PONT": "Ponape áigi", "PST": "Jaskesábi dálveáigi", "PST Philippine": "Filippiinnaid dálveáigi", "PST Philippine DST": "Filippiinnaid geasseáigi", "PST Pitcairn": "Pitcairnsulloid áigi", "PSTM": "Meksiko Jáskesábi dálveáigi", "PWT": "Palaua áigi", "PYST": "Paraguaya geasseáigi", "PYT": "Paraguaya dálveáigi", "PYT Korea": "Pyongyanga áigi", "RET": "Reuniona áigi", "ROTT": "Rothera áigi", "SAKST": "Sahalina geasseáigi", "SAKT": "Sahalina dálveáigi", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Lulli-Afrihká dálveáigi", "SBT": "Salomonsulloid áigi", "SCT": "Seychellaid áigi", "SGT": "Singapore dálveáigi", "SLST": "SLST", "SRT": "Suriname áigi", "SST Samoa": "Samoa dálveáigi", "SST Samoa Apia": "Apia dálveáigi", "SST Samoa Apia DST": "Apia geasseáigi", "SST Samoa DST": "Samoa geasseáigi", "SYOT": "Syowa áigi", "TAAF": "Frankriikka lulli & antárktisa áigi", "TAHT": "Tahiti áigi", "TJT": "Tažikistana áigi", "TKT": "Tokelaua áigi", "TLT": "Nuorta-Timora áigi", "TMST": "Turkmenistana geasseáigi", "TMT": "Turkmenistana dálveáigi", "TOST": "Tonga geasseáigi", "TOT": "Tonga dálveáigi", "TVT": "Tuvalu áigi", "TWT": "Taipeia dálveáigi", "TWT DST": "Taipeia geasseáigi", "ULAST": "Ulan-Batora geasseáigi", "ULAT": "Ulan-Batora dálveáigi", "UYST": "Uruguaya geasseáigi", "UYT": "Uruguaya dálveáigi", "UZT": "Usbekistana dálveáigi", "UZT DST": "Usbekistana geasseáigi", "VET": "Venezuela áigi", "VLAST": "Vladivostoka geasseáigi", "VLAT": "Vladivostoka dálveáigi", "VOLST": "Volgograda geasseáigi", "VOLT": "Volgograda dálveáigi", "VOST": "Vostoka áigi", "VUT": "Vanuatu dálveáigi", "VUT DST": "Vanuatu geasseáigi", "WAKT": "Wakesullo áigi", "WARST": "Oarje-Argentina geasseáigi", "WART": "Oarje-Argentina dálveáigi", "WAST": "Oarje-Afrihká áigi", "WAT": "Oarje-Afrihká áigi", "WESZ": "Oarje-Eurohpá geasseáigi", "WEZ": "Oarje-Eurohpá dálveáigi", "WFT": "Wallis- ja Futuna áigi", "WGST": "Oarje-Ruonáeatnama geasseáigi", "WGT": "Oarje-Ruonáeatnama dálveáigi", "WIB": "Oarje-Indonesia áigi", "WIT": "Nuorta-Indonesia áigi", "WITA": "Gaska-Indonesia áigi", "YAKST": "Jakucka geasseáigi", "YAKT": "Jakucka dálveáigi", "YEKST": "Jekaterinburga geasseáigi", "YEKT": "Jekaterinburga dálveáigi", "YST": "YST", "МСК": "Moskva dálveáigi", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Oarje-Kasakstana áigi", "شىعىش قازاق ەلى": "Nuorta-Kasakstana áigi", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Kirgisia áigi", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Peru geasseáigi"},
	}
}

// Locale returns the current translators string locale
func (se *se_FI) Locale() string {
	return se.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'se_FI'
func (se *se_FI) PluralsCardinal() []locales.PluralRule {
	return se.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'se_FI'
func (se *se_FI) PluralsOrdinal() []locales.PluralRule {
	return se.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'se_FI'
func (se *se_FI) PluralsRange() []locales.PluralRule {
	return se.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'se_FI'
func (se *se_FI) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'se_FI'
func (se *se_FI) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'se_FI'
func (se *se_FI) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (se *se_FI) MonthAbbreviated(month time.Month) string {
	return se.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (se *se_FI) MonthsAbbreviated() []string {
	return se.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (se *se_FI) MonthNarrow(month time.Month) string {
	return se.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (se *se_FI) MonthsNarrow() []string {
	return se.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (se *se_FI) MonthWide(month time.Month) string {
	return se.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (se *se_FI) MonthsWide() []string {
	return se.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (se *se_FI) WeekdayAbbreviated(weekday time.Weekday) string {
	return se.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (se *se_FI) WeekdaysAbbreviated() []string {
	return se.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (se *se_FI) WeekdayNarrow(weekday time.Weekday) string {
	return se.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (se *se_FI) WeekdaysNarrow() []string {
	return se.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (se *se_FI) WeekdayShort(weekday time.Weekday) string {
	return se.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (se *se_FI) WeekdaysShort() []string {
	return se.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (se *se_FI) WeekdayWide(weekday time.Weekday) string {
	return se.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (se *se_FI) WeekdaysWide() []string {
	return se.daysWide
}

// Decimal returns the decimal point of number
func (se *se_FI) Decimal() string {
	return se.decimal
}

// Group returns the group of number
func (se *se_FI) Group() string {
	return se.group
}

// Group returns the minus sign of number
func (se *se_FI) Minus() string {
	return se.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'se_FI' and handles both Whole and Real numbers based on 'v'
func (se *se_FI) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'se_FI' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (se *se_FI) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, se.percentSuffix...)

	b = append(b, se.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'se_FI'
func (se *se_FI) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := se.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, se.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, se.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'se_FI'
// in accounting notation.
func (se *se_FI) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := se.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, se.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, se.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, se.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'se_FI'
func (se *se_FI) FmtDateShort(t time.Time) string {

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

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'se_FI'
func (se *se_FI) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, se.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'se_FI'
func (se *se_FI) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, se.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'se_FI'
func (se *se_FI) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, se.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, se.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'se_FI'
func (se *se_FI) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'se_FI'
func (se *se_FI) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'se_FI'
func (se *se_FI) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'se_FI'
func (se *se_FI) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}
