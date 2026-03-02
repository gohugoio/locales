package lb

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type lb struct {
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

// New returns a new instance of translator for the 'lb' locale
func New() locales.Translator {
	return &lb{
		locale:                 "lb",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "öS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "Jan.", "Feb.", "Mäe.", "Abr.", "Mee", "Juni", "Juli", "Aug.", "Sep.", "Okt.", "Nov.", "Dez."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januar", "Februar", "Mäerz", "Abrëll", "Mee", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
		daysAbbreviated:        []string{"Son.", "Méi.", "Dën.", "Mët.", "Don.", "Fre.", "Sam."},
		daysNarrow:             []string{"S", "M", "D", "M", "D", "F", "S"},
		daysShort:              []string{"So.", "Mé.", "Dë.", "Më.", "Do.", "Fr.", "Sa."},
		daysWide:               []string{"Sonndeg", "Méindeg", "Dënschdeg", "Mëttwoch", "Donneschdeg", "Freideg", "Samschdeg"},
		periodsAbbreviated:     []string{"moies", "nomëttes"},
		periodsNarrow:          []string{"mo.", "nomë."},
		erasAbbreviated:        []string{"v. Chr.", "n. Chr."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "Zentralaustralesch Summerzäit", "ACST": "Acre-Summerzäit", "ACT": "Acre-Normalzäit", "ACWDT": "Zentral-/Westaustralesch Summerzäit", "ACWST": "Zentral-/Westaustralesch Normalzäit", "ADT": "Atlantik-Summerzäit", "ADT Arabia": "Arabesch Summerzäit", "AEDT": "Ostaustralesch Summerzäit", "AEST": "Ostaustralesch Normalzäit", "AFT": "Afghanistan-Zäit", "AKDT": "Alaska-Summerzäit", "AKST": "Alaska-Normalzäit", "AMST": "Amazonas-Summerzäit", "AMST Armenia": "Armenesch Summerzäit", "AMT": "Amazonas-Normalzäit", "AMT Armenia": "Armenesch Normalzäit", "ANAST": "Anadyr-Summerzäit", "ANAT": "Anadyr-Normalzäit", "ARST": "Argentinesch Summerzäit", "ART": "Argentinesch Normalzäit", "AST": "Atlantik-Normalzäit", "AST Arabia": "Arabesch Normalzäit", "AWDT": "Westaustralesch Summerzäit", "AWST": "Westaustralesch Normalzäit", "AZST": "Aserbaidschanesch Summerzäit", "AZT": "Aserbeidschanesch Normalzäit", "BDT Bangladesh": "Bangladesch-Summerzäit", "BNT": "Brunei-Zäit", "BOT": "Bolivianesch Zäit", "BRST": "Brasília-Summerzäit", "BRT": "Brasília-Normalzäit", "BST Bangladesh": "Bangladesch-Normalzäit", "BT": "Bhutan-Zäit", "CAST": "CAST", "CAT": "Zentralafrikanesch Zäit", "CCT": "Kokosinselen-Zäit", "CDT": "Nordamerikanesch Inland-Summerzäit", "CHADT": "Chatham-Summerzäit", "CHAST": "Chatham-Normalzäit", "CHUT": "Chuuk-Zäit", "CKT": "Cookinselen-Normalzäit", "CKT DST": "Cookinselen-Summerzäit", "CLST": "Chilenesch Summerzäit", "CLT": "Chilenesch Normalzäit", "COST": "Kolumbianesch Summerzäit", "COT": "Kolumbianesch Normalzäit", "CST": "Nordamerikanesch Inland-Normalzäit", "CST China": "Chinesesch Normalzäit", "CST China DST": "Chinesesch Summerzäit", "CVST": "Kap-Verde-Summerzäit", "CVT": "Kap-Verde-Normalzäit", "CXT": "Chrëschtdagsinsel-Zäit", "ChST": "Chamorro-Zäit", "ChST NMI": "ChST NMI", "CuDT": "Kubanesch Summerzäit", "CuST": "Kubanesch Normalzäit", "DAVT": "Davis-Zäit", "DDUT": "Dumont-d’Urville-Zäit", "EASST": "Ouschterinsel-Summerzäit", "EAST": "Ouschterinsel-Normalzäit", "EAT": "Ostafrikanesch Zäit", "ECT": "Ecuadorianesch Zäit", "EDT": "Nordamerikanesch Ostküsten-Summerzäit", "EGDT": "Ostgrönland-Summerzäit", "EGST": "Ostgrönland-Normalzäit", "EST": "Nordamerikanesch Ostküsten-Normalzäit", "FEET": "FEET", "FJT": "Fidschi-Normalzäit", "FJT Summer": "Fidschi-Summerzäit", "FKST": "Falklandinselen-Summerzäit", "FKT": "Falklandinselen-Normalzäit", "FNST": "Fernando-de-Noronha-Summerzäit", "FNT": "Fernando-de-Noronha-Normalzäit", "GALT": "Galapagos-Zäit", "GAMT": "Gambier-Zäit", "GEST": "Georgesch Summerzäit", "GET": "Georgesch Normalzäit", "GFT": "Franséisch-Guayane-Zäit", "GIT": "Gilbert-Inselen-Zäit", "GMT": "Mëttler Greenwich-Zäit", "GNSST": "GNSST", "GNST": "GNST", "GST": "Südgeorgesch Zäit", "GST Guam": "Guam-Zäit", "GYT": "Guyana-Zäit", "HADT": "Hawaii-Aleuten-Summerzäit", "HAST": "Hawaii-Aleuten-Normalzäit", "HKST": "Hong-Kong-Summerzäit", "HKT": "Hong-Kong-Normalzäit", "HOVST": "Hovd-Summerzäit", "HOVT": "Hovd-Normalzäit", "ICT": "Indochina-Zäit", "IDT": "Israelesch Summerzäit", "IOT": "Indeschen Ozean-Zäit", "IRKST": "Irkutsk-Summerzäit", "IRKT": "Irkutsk-Normalzäit", "IRST": "Iranesch Normalzäit", "IRST DST": "Iranesch Summerzäit", "IST": "Indesch Zäit", "IST Israel": "Israelesch Normalzäit", "JDT": "Japanesch Summerzäit", "JST": "Japanesch Normalzäit", "KOST": "Kosrae-Zäit", "KRAST": "Krasnojarsk-Summerzäit", "KRAT": "Krasnojarsk-Normalzäit", "KST": "Koreanesch Normalzäit", "KST DST": "Koreanesch Summerzäit", "LHDT": "Lord-Howe-Summerzäit", "LHST": "Lord-Howe-Normalzäit", "LINT": "Linneninselen-Zäit", "MAGST": "Magadan-Summerzäit", "MAGT": "Magadan-Normalzäit", "MART": "Marquesas-Zäit", "MAWT": "Mawson-Zäit", "MDT": "Rocky-Mountain-Summerzäit", "MESZ": "Mëtteleuropäesch Summerzäit", "MEZ": "Mëtteleuropäesch Normalzäit", "MHT": "Marshallinselen-Zäit", "MMT": "Myanmar-Zäit", "MSD": "Moskauer Summerzäit", "MST": "Rocky-Mountain-Normalzäit", "MUST": "Mauritius-Summerzäit", "MUT": "Mauritius-Normalzäit", "MVT": "Maldiven-Zäit", "MYT": "Malaysesch Zäit", "NCT": "Neikaledonesch Normalzäit", "NDT": "Neifundland-Summerzäit", "NDT New Caledonia": "Neikaledonesch Summerzäit", "NFDT": "Norfolkinselen-Summerzäit", "NFT": "Norfolkinselen-Normalzäit", "NOVST": "Nowosibirsk-Summerzäit", "NOVT": "Nowosibirsk-Normalzäit", "NPT": "Nepalesesch Zäit", "NRT": "Nauru-Zäit", "NST": "Neifundland-Normalzäit", "NUT": "Niue-Zäit", "NZDT": "Neiséiland-Summerzäit", "NZST": "Neiséiland-Normalzäit", "OESZ": "Osteuropäesch Summerzäit", "OEZ": "Osteuropäesch Normalzäit", "OMSST": "Omsk-Summerzäit", "OMST": "Omsk-Normalzäit", "PDT": "Nordamerikanesch Westküsten-Summerzäit", "PDTM": "Mexikanesch Pazifik-Summerzäit", "PETDT": "Kamtschatka-Summerzäit", "PETST": "Kamtschatka-Normalzäit", "PGT": "Papua-Neiguinea-Zäit", "PHOT": "Phoenixinselen-Zäit", "PKT": "Pakistanesch Normalzäit", "PKT DST": "Pakistanesch Summerzäit", "PMDT": "Saint-Pierre-a-Miquelon-Summerzäit", "PMST": "Saint-Pierre-a-Miquelon-Normalzäit", "PONT": "Ponape-Zäit", "PST": "Nordamerikanesch Westküsten-Normalzäit", "PST Philippine": "Philippinnesch Normalzäit", "PST Philippine DST": "Philippinnesch Summerzäit", "PST Pitcairn": "Pitcairninselen-Zäit", "PSTM": "Mexikanesch Pazifik-Normalzäit", "PWT": "Palau-Zäit", "PYST": "Paraguayanesch Summerzäit", "PYT": "Paraguayanesch Normalzäit", "PYT Korea": "PYT Korea", "RET": "Réunion-Zäit", "ROTT": "Rothera-Zäit", "SAKST": "Sakhalin-Summerzäit", "SAKT": "Sakhalin-Normalzäit", "SAMST": "Samara-Summerzäit", "SAMT": "Samara-Normalzäit", "SAST": "Südafrikanesch Zäit", "SBT": "Salomoninselen-Zäit", "SCT": "Seychellen-Zäit", "SGT": "Singapur-Standardzäit", "SLST": "SLST", "SRT": "Suriname-Zäit", "SST Samoa": "Samoa-Normalzäit", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "Samoa-Summerzäit", "SYOT": "Syowa-Zäit", "TAAF": "Franséisch Süd- an Antarktisgebidder-Zäit", "TAHT": "Tahiti-Zäit", "TJT": "Tadschikistan-Zäit", "TKT": "Tokelau-Zäit", "TLT": "Osttimor-Zäit", "TMST": "Turkmenistan-Summerzäit", "TMT": "Turkmenistan-Normalzäit", "TOST": "Tonganesch Summerzäit", "TOT": "Tonganesch Normalzäit", "TVT": "Tuvalu-Zäit", "TWT": "Taipei-Normalzäit", "TWT DST": "Taipei-Summerzäit", "ULAST": "Ulaanbaatar-Summerzäit", "ULAT": "Ulaanbaatar-Normalzäit", "UYST": "Uruguayanesch Summerzäit", "UYT": "Uruguyanesch Normalzäit", "UZT": "Usbekistan-Normalzäit", "UZT DST": "Usbekistan-Summerzäit", "VET": "Venezuela-Zäit", "VLAST": "Wladiwostok-Summerzäit", "VLAT": "Wladiwostok-Normalzäit", "VOLST": "Wolgograd-Summerzäit", "VOLT": "Wolgograd-Normalzäit", "VOST": "Wostok-Zäit", "VUT": "Vanuatu-Normalzäit", "VUT DST": "Vanuatu-Summerzäit", "WAKT": "Wake-Insel-Zäit", "WARST": "Westargentinesch Summerzäit", "WART": "Westargentinesch Normalzäit", "WAST": "Westafrikanesch Zäit", "WAT": "Westafrikanesch Zäit", "WESZ": "Westeuropäesch Summerzäit", "WEZ": "Westeuropäesch Normalzäit", "WFT": "Wallis-a-Futuna-Zäit", "WGST": "Westgrönland-Summerzäit", "WGT": "Westgrönland-Normalzäit", "WIB": "Westindonesesch Zäit", "WIT": "Ostindonesesch Zäit", "WITA": "Zentralindonesesch Zäit", "YAKST": "Jakutsk-Summerzäit", "YAKT": "Jakutsk-Normalzäit", "YEKST": "Jekaterinbuerg-Summerzäit", "YEKT": "Jekaterinbuerg-Normalzäit", "YST": "YST", "МСК": "Moskauer Normalzäit", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "Almaty-Normalzäit", "الماتى قالاسى": "Almaty-Summerzäit", "باتىس قازاق ەلى": "Westkasachesch Zäit", "شىعىش قازاق ەلى": "Ostkasachesch Zäit", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Kirgisistan-Zäit", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Peruanesch Summerzäit"},
	}
}

// Locale returns the current translators string locale
func (lb *lb) Locale() string {
	return lb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lb'
func (lb *lb) PluralsCardinal() []locales.PluralRule {
	return lb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lb'
func (lb *lb) PluralsOrdinal() []locales.PluralRule {
	return lb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lb'
func (lb *lb) PluralsRange() []locales.PluralRule {
	return lb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lb'
func (lb *lb) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lb'
func (lb *lb) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lb'
func (lb *lb) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lb *lb) MonthAbbreviated(month time.Month) string {
	return lb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lb *lb) MonthsAbbreviated() []string {
	return lb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lb *lb) MonthNarrow(month time.Month) string {
	return lb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lb *lb) MonthsNarrow() []string {
	return lb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lb *lb) MonthWide(month time.Month) string {
	return lb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lb *lb) MonthsWide() []string {
	return lb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lb *lb) WeekdayAbbreviated(weekday time.Weekday) string {
	return lb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lb *lb) WeekdaysAbbreviated() []string {
	return lb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lb *lb) WeekdayNarrow(weekday time.Weekday) string {
	return lb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lb *lb) WeekdaysNarrow() []string {
	return lb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lb *lb) WeekdayShort(weekday time.Weekday) string {
	return lb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lb *lb) WeekdaysShort() []string {
	return lb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lb *lb) WeekdayWide(weekday time.Weekday) string {
	return lb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lb *lb) WeekdaysWide() []string {
	return lb.daysWide
}

// Decimal returns the decimal point of number
func (lb *lb) Decimal() string {
	return lb.decimal
}

// Group returns the group of number
func (lb *lb) Group() string {
	return lb.group
}

// Group returns the minus sign of number
func (lb *lb) Minus() string {
	return lb.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lb' and handles both Whole and Real numbers based on 'v'
func (lb *lb) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lb' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lb *lb) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lb.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lb.percentSuffix...)

	b = append(b, lb.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lb'
func (lb *lb) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lb.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, lb.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lb'
// in accounting notation.
func (lb *lb) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lb.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, lb.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, lb.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, lb.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'lb'
func (lb *lb) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'lb'
func (lb *lb) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lb.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'lb'
func (lb *lb) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'lb'
func (lb *lb) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, lb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'lb'
func (lb *lb) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'lb'
func (lb *lb) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'lb'
func (lb *lb) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'lb'
func (lb *lb) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
