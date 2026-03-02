package af

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type af struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'af' locale
func New() locales.Translator {
	return &af{
		locale:                 "af",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "leu", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Jan.", "Feb.", "Mrt.", "Apr.", "Mei", "Jun.", "Jul.", "Aug.", "Sep.", "Okt.", "Nov.", "Des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januarie", "Februarie", "Maart", "April", "Mei", "Junie", "Julie", "Augustus", "September", "Oktober", "November", "Desember"},
		daysAbbreviated:        []string{"So.", "Ma.", "Di.", "Wo.", "Do.", "Vr.", "Sa."},
		daysNarrow:             []string{"S", "M", "D", "W", "D", "V", "S"},
		daysWide:               []string{"Sondag", "Maandag", "Dinsdag", "Woensdag", "Donderdag", "Vrydag", "Saterdag"},
		periodsAbbreviated:     []string{"vm.", "nm."},
		periodsNarrow:          []string{"v", "n"},
		erasAbbreviated:        []string{"v.C.", "n.C."},
		erasNarrow:             []string{"vgj", "gj"},
		erasWide:               []string{"voor Christus", "ná Christus"},
		timezones:              map[string]string{"ACDT": "Sentraal-Australiese dagligtyd", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Sentraal-westelike Australiese dagligtyd", "ACWST": "Sentraal-westelike Australiese standaard-tyd", "ADT": "Atlantiese dagligtyd", "ADT Arabia": "Arabiese dagligtyd", "AEDT": "Oostelike Australiese dagligtyd", "AEST": "Oostelike Australiese standaardtyd", "AFT": "Afganistan-tyd", "AKDT": "Alaska-dagligtyd", "AKST": "Alaska-standaardtyd", "AMST": "Amasone-somertyd", "AMST Armenia": "Armenië-somertyd", "AMT": "Amasone-standaardtyd", "AMT Armenia": "Armenië-standaardtyd", "ANAST": "Anadyr-somertyd", "ANAT": "Anadyr-standaardtyd", "ARST": "Argentinië-somertyd", "ART": "Argentinië-standaardtyd", "AST": "Atlantiese standaardtyd", "AST Arabia": "Arabiese standaardtyd", "AWDT": "Westelike Australiese dagligtyd", "AWST": "Westelike Australiese standaardtyd", "AZST": "Azerbeidjan-somertyd", "AZT": "Azerbeidjan-standaardtyd", "BDT Bangladesh": "Bangladesj-somertyd", "BNT": "Broenei Darussalam-tyd", "BOT": "Bolivië-tyd", "BRST": "Brasilia-somertyd", "BRT": "Brasilia-standaardtyd", "BST Bangladesh": "Bangladesj-standaardtyd", "BT": "Bhoetan-tyd", "CAST": "CAST", "CAT": "Sentraal-Afrika-tyd", "CCT": "Kokoseilande-tyd", "CDT": "Noord-Amerikaanse sentrale dagligtyd", "CHADT": "Chatham-dagligtyd", "CHAST": "Chatham-standaardtyd", "CHUT": "Chuuk-tyd", "CKT": "Cookeilande-standaardtyd", "CKT DST": "Cookeilande-halfsomertyd", "CLST": "Chili-somertyd", "CLT": "Chili-standaardtyd", "COST": "Colombië-somertyd", "COT": "Colombië-standaardtyd", "CST": "Noord-Amerikaanse sentrale standaardtyd", "CST China": "China-standaardtyd", "CST China DST": "China-dagligtyd", "CVST": "Kaap Verde-somertyd", "CVT": "Kaap Verde-standaardtyd", "CXT": "Christmaseiland-tyd", "ChST": "Chamorro-standaardtyd", "ChST NMI": "ChST NMI", "CuDT": "Kuba-dagligtyd", "CuST": "Kuba-standaardtyd", "DAVT": "Davis-tyd", "DDUT": "Dumont d’Urville-tyd", "EASST": "Paaseiland-somertyd", "EAST": "Paaseiland-standaardtyd", "EAT": "Oos-Afrika-tyd", "ECT": "Ecuador-tyd", "EDT": "Noord-Amerikaanse oostelike dagligtyd", "EGDT": "Oos-Groenland-somertyd", "EGST": "Oos-Groenland-standaardtyd", "EST": "Noord-Amerikaanse oostelike standaardtyd", "FEET": "Verder-oos-Europese tyd", "FJT": "Fidji-standaardtyd", "FJT Summer": "Fidji-somertyd", "FKST": "Falklandeilande-somertyd", "FKT": "Falklandeilande-standaardtyd", "FNST": "Fernando de Noronha-somertyd", "FNT": "Fernando de Noronha-standaardtyd", "GALT": "Galapagos-tyd", "GAMT": "Gambier-tyd", "GEST": "Georgië-somertyd", "GET": "Georgië-standaardtyd", "GFT": "Frans-Guiana-tyd", "GIT": "Gilberteilande-tyd", "GMT": "Greenwich-tyd", "GNSST": "GNSST", "GNST": "GNST", "GST": "Suid-Georgië-tyd", "GST Guam": "GST Guam", "GYT": "Guiana-tyd", "HADT": "Hawaii-Aleoete-dagligtyd", "HAST": "Hawaii-Aleoete-standaardtyd", "HKST": "Hongkong-somertyd", "HKT": "Hongkong-standaardtyd", "HOVST": "Hovd-somertyd", "HOVT": "Hovd-standaardtyd", "ICT": "Indosjina-tyd", "IDT": "Israel-dagligtyd", "IOT": "Indiese Oseaan-tyd", "IRKST": "Irkoetsk-somertyd", "IRKT": "Irkoetsk-standaardtyd", "IRST": "Iran-standaardtyd", "IRST DST": "Iran-dagligtyd", "IST": "Indië-standaardtyd", "IST Israel": "Israel-standaardtyd", "JDT": "Japan-dagligtyd", "JST": "Japan-standaardtyd", "KOST": "Kosrae-tyd", "KRAST": "Krasnojarsk-somertyd", "KRAT": "Krasnojarsk-standaardtyd", "KST": "Koreaanse standaardtyd", "KST DST": "Koreaanse dagligtyd", "LHDT": "Lord Howe-dagligtyd", "LHST": "Lord Howe-standaardtyd", "LINT": "Line-eilande-tyd", "MAGST": "Magadan-somertyd", "MAGT": "Magadan-standaardtyd", "MART": "Marquesas-tyd", "MAWT": "Mawson-tyd", "MDT": "MDT", "MESZ": "Sentraal-Europese somertyd", "MEZ": "Sentraal-Europese standaardtyd", "MHT": "Marshalleilande-tyd", "MMT": "Mianmar-tyd", "MSD": "Moskou-somertyd", "MST": "MST", "MUST": "Mauritius-somertyd", "MUT": "Mauritius-standaardtyd", "MVT": "Maledive-tyd", "MYT": "Maleisië-tyd", "NCT": "Nieu-Kaledonië-standaardtyd", "NDT": "Newfoundland-dagligtyd", "NDT New Caledonia": "Nieu-Kaledonië-somertyd", "NFDT": "Norfolkeiland-dagligtyd", "NFT": "Norfolkeiland-standaardtyd", "NOVST": "Novosibirsk-somertyd", "NOVT": "Novosibirsk-standaardtyd", "NPT": "Nepal-tyd", "NRT": "Nauru-tyd", "NST": "Newfoundland-standaardtyd", "NUT": "Niue-tyd", "NZDT": "Nieu-Seeland-dagligtyd", "NZST": "Nieu-Seeland-standaardtyd", "OESZ": "Oos-Europese somertyd", "OEZ": "Oos-Europese standaardtyd", "OMSST": "Omsk-somertyd", "OMST": "Omsk-standaardtyd", "PDT": "Pasifiese dagligtyd", "PDTM": "Meksikaanse Pasifiese dagligtyd", "PETDT": "Petropavlovsk-Kamchatski-somertyd", "PETST": "Petropavlovsk-Kamchatski-standaardtyd", "PGT": "Papoea-Nieu-Guinee-tyd", "PHOT": "Fenikseilande-tyd", "PKT": "Pakistan-standaardtyd", "PKT DST": "Pakistan-somertyd", "PMDT": "Sint-Pierre en Miquelon-dagligtyd", "PMST": "Sint-Pierre en Miquelon-standaardtyd", "PONT": "Ponape-tyd", "PST": "Pasifiese standaardtyd", "PST Philippine": "Filippynse standaardtyd", "PST Philippine DST": "Filippynse somertyd", "PST Pitcairn": "Pitcairn-tyd", "PSTM": "Meksikaanse Pasifiese standaardtyd", "PWT": "Palau-tyd", "PYST": "Paraguay-somertyd", "PYT": "Paraguay-standaardtyd", "PYT Korea": "Pyongyang-tyd", "RET": "Réunion-tyd", "ROTT": "Rothera-tyd", "SAKST": "Sakhalin-somertyd", "SAKT": "Sakhalin-standaardtyd", "SAMST": "Samara-dagligtyd", "SAMT": "Samara-standaardtyd", "SAST": "Suid-Afrika-standaardtyd", "SBT": "Salomonseilande-tyd", "SCT": "Seychelle-tyd", "SGT": "Singapoer-standaardtyd", "SLST": "SLST", "SRT": "Suriname-tyd", "SST Samoa": "Samoa-standaardtyd", "SST Samoa Apia": "Apia-standaardtyd", "SST Samoa Apia DST": "Apia-dagligtyd", "SST Samoa DST": "Samoa-dagligtyd", "SYOT": "Syowa-tyd", "TAAF": "Franse Suider- en Antarktiese tyd", "TAHT": "Tahiti-tyd", "TJT": "Tadjikistan-tyd", "TKT": "Tokelau-tyd", "TLT": "Oos-Timor-tyd", "TMST": "Turkmenistan-somertyd", "TMT": "Turkmenistan-standaardtyd", "TOST": "Tonga-somertyd", "TOT": "Tonga-standaardtyd", "TVT": "Tuvalu-tyd", "TWT": "Taipei-standaardtyd", "TWT DST": "Taipei-dagligtyd", "ULAST": "Ulaanbaatar-somertyd", "ULAT": "Ulaanbaatar-standaardtyd", "UYST": "Uruguay-somertyd", "UYT": "Uruguay-standaardtyd", "UZT": "Oesbekistan-standaardtyd", "UZT DST": "Oesbekistan-somertyd", "VET": "Venezuela-tyd", "VLAST": "Wladiwostok-somertyd", "VLAT": "Wladiwostok-standaardtyd", "VOLST": "Wolgograd-somertyd", "VOLT": "Wolgograd-standaardtyd", "VOST": "Wostok-tyd", "VUT": "Vanuatu-standaardtyd", "VUT DST": "Vanuatu-somertyd", "WAKT": "Wake-eiland-tyd", "WARST": "Wes-Argentinië-somertyd", "WART": "Wes-Argentinië-standaardtyd", "WAST": "Wes-Afrika-tyd", "WAT": "Wes-Afrika-tyd", "WESZ": "Wes-Europese somertyd", "WEZ": "Wes-Europese standaardtyd", "WFT": "Wallis en Futuna-tyd", "WGST": "Wes-Groenland-somertyd", "WGT": "Wes-Groenland-standaardtyd", "WIB": "Wes-Indonesië-tyd", "WIT": "Oos-Indonesië-tyd", "WITA": "Sentraal-Indonesiese tyd", "YAKST": "Jakoetsk-somertyd", "YAKT": "Jakoetsk-standaardtyd", "YEKST": "Jekaterinburg-somertyd", "YEKT": "Jekaterinburg-standaardtyd", "YST": "Yukontyd", "МСК": "Moskou-standaardtyd", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Wes-Kazakstan-tyd", "شىعىش قازاق ەلى": "Oos-Kazakstan-tyd", "قازاق ەلى": "Kazakstan-tyd", "قىرعىزستان": "Kirgistan-tyd", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Asore-somertyd"},
	}
}

// Locale returns the current translators string locale
func (af *af) Locale() string {
	return af.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'af'
func (af *af) PluralsCardinal() []locales.PluralRule {
	return af.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'af'
func (af *af) PluralsOrdinal() []locales.PluralRule {
	return af.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'af'
func (af *af) PluralsRange() []locales.PluralRule {
	return af.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'af'
func (af *af) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'af'
func (af *af) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'af'
func (af *af) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (af *af) MonthAbbreviated(month time.Month) string {
	return af.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (af *af) MonthsAbbreviated() []string {
	return af.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (af *af) MonthNarrow(month time.Month) string {
	return af.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (af *af) MonthsNarrow() []string {
	return af.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (af *af) MonthWide(month time.Month) string {
	return af.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (af *af) MonthsWide() []string {
	return af.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (af *af) WeekdayAbbreviated(weekday time.Weekday) string {
	return af.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (af *af) WeekdaysAbbreviated() []string {
	return af.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (af *af) WeekdayNarrow(weekday time.Weekday) string {
	return af.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (af *af) WeekdaysNarrow() []string {
	return af.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (af *af) WeekdayShort(weekday time.Weekday) string {
	return af.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (af *af) WeekdaysShort() []string {
	return af.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (af *af) WeekdayWide(weekday time.Weekday) string {
	return af.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (af *af) WeekdaysWide() []string {
	return af.daysWide
}

// Decimal returns the decimal point of number
func (af *af) Decimal() string {
	return af.decimal
}

// Group returns the group of number
func (af *af) Group() string {
	return af.group
}

// Group returns the minus sign of number
func (af *af) Minus() string {
	return af.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'af' and handles both Whole and Real numbers based on 'v'
func (af *af) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'af' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (af *af) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'af'
func (af *af) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := af.currencies[currency]
	l := len(s) + len(symbol) + 1 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, af.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(af.group) - 1; j >= 0; j-- {
					b = append(b, af.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, af.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, af.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'af'
// in accounting notation.
func (af *af) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := af.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, af.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(af.group) - 1; j >= 0; j-- {
					b = append(b, af.group[j])
				}
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

		b = append(b, af.currencyNegativePrefix[0])

	} else {

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
			b = append(b, af.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, af.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'af'
func (af *af) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'af'
func (af *af) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, af.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'af'
func (af *af) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, af.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'af'
func (af *af) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, af.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, af.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'af'
func (af *af) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, af.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'af'
func (af *af) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, af.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, af.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'af'
func (af *af) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, af.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, af.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'af'
func (af *af) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, af.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, af.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := af.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
