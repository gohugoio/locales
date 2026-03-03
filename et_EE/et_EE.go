package et_EE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type et_EE struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'et_EE' locale
func New() locales.Translator {
	return &et_EE{
		locale:                 "et_EE",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "kr", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jaan", "veebr", "märts", "apr", "mai", "juuni", "juuli", "aug", "sept", "okt", "nov", "dets"},
		monthsNarrow:           []string{"", "J", "V", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "jaanuar", "veebruar", "märts", "aprill", "mai", "juuni", "juuli", "august", "september", "oktoober", "november", "detsember"},
		daysAbbreviated:        []string{"P", "E", "T", "K", "N", "R", "L"},
		daysNarrow:             []string{"P", "E", "T", "K", "N", "R", "L"},
		daysWide:               []string{"pühapäev", "esmaspäev", "teisipäev", "kolmapäev", "neljapäev", "reede", "laupäev"},
		timezones:              map[string]string{"ACDT": "Kesk-Austraalia suveaeg", "ACST": "Kesk-Austraalia standardaeg", "ACT": "Acre standardaeg", "ACWDT": "Austraalia Kesk-Lääne suveaeg", "ACWST": "Austraalia Kesk-Lääne standardaeg", "ADT": "Atlandi suveaeg", "ADT Arabia": "Araabia suveaeg", "AEDT": "Ida-Austraalia suveaeg", "AEST": "Ida-Austraalia standardaeg", "AFT": "Afganistani aeg", "AKDT": "Alaska suveaeg", "AKST": "Alaska standardaeg", "AMST": "Amazonase suveaeg", "AMST Armenia": "Armeenia suveaeg", "AMT": "Amazonase standardaeg", "AMT Armenia": "Armeenia standardaeg", "ANAST": "Anadõri suveaeg", "ANAT": "Anadõri standardaeg", "ARST": "Argentina suveaeg", "ART": "Argentina standardaeg", "AST": "Atlandi standardaeg", "AST Arabia": "Araabia standardaeg", "AWDT": "Lääne-Austraalia suveaeg", "AWST": "Lääne-Austraalia standardaeg", "AZST": "Aserbaidžaani suveaeg", "AZT": "Aserbaidžaani standardaeg", "BDT Bangladesh": "Bangladeshi suveaeg", "BNT": "Brunei aeg", "BOT": "Boliivia aeg", "BRST": "Brasiilia suveaeg", "BRT": "Brasiilia standardaeg", "BST Bangladesh": "Bangladeshi standardaeg", "BT": "Bhutani aeg", "CAST": "Casey aeg", "CAT": "Kesk-Aafrika aeg", "CCT": "Kookossaarte aeg", "CDT": "Kesk-Ameerika suveaeg", "CHADT": "Chathami suveaeg", "CHAST": "Chathami standardaeg", "CHUT": "Chuuki aeg", "CKT": "Cooki saarte standardaeg", "CKT DST": "Cooki saarte osaline suveaeg", "CLST": "Tšiili suveaeg", "CLT": "Tšiili standardaeg", "COST": "Colombia suveaeg", "COT": "Colombia standardaeg", "CST": "Kesk-Ameerika standardaeg", "CST China": "Hiina standardaeg", "CST China DST": "Hiina suveaeg", "CVST": "Roheneemesaarte suveaeg", "CVT": "Roheneemesaarte standardaeg", "CXT": "Jõulusaare aeg", "ChST": "Tšamorro standardaeg", "ChST NMI": "Põhja-Mariaani aeg", "CuDT": "Kuuba suveaeg", "CuST": "Kuuba standardaeg", "DAVT": "Davise aeg", "DDUT": "Dumont d’Urville’i aeg", "EASST": "Lihavõttesaare suveaeg", "EAST": "Lihavõttesaare standardaeg", "EAT": "Ida-Aafrika aeg", "ECT": "Ecuadori aeg", "EDT": "Idaranniku suveaeg", "EGDT": "Ida-Gröönimaa suveaeg", "EGST": "Ida-Gröönimaa standardaeg", "EST": "Idaranniku standardaeg", "FEET": "Kaliningradi ja Valgevene aeg", "FJT": "Fidži standardaeg", "FJT Summer": "Fidži suveaeg", "FKST": "Falklandi saarte suveaeg", "FKT": "Falklandi saarte standardaeg", "FNST": "Fernando de Noronha suveaeg", "FNT": "Fernando de Noronha standardaeg", "GALT": "Galapagose aeg", "GAMT": "Gambier’ aeg", "GEST": "Gruusia suveaeg", "GET": "Gruusia standardaeg", "GFT": "Prantsuse Guajaana aeg", "GIT": "Gilberti saarte aeg", "GMT": "Greenwichi aeg", "GNSST": "Gröönimaa suveaeg", "GNST": "Gröönimaa standardaeg", "GST": "Lõuna-Georgia aeg", "GST Guam": "Guami standardaeg", "GYT": "Guyana aeg", "HADT": "Hawaii-Aleuudi suveaeg", "HAST": "Hawaii-Aleuudi standardaeg", "HKST": "Hongkongi suveaeg", "HKT": "Hongkongi standardaeg", "HOVST": "Hovdi suveaeg", "HOVT": "Hovdi standardaeg", "ICT": "Indohiina aeg", "IDT": "Iisraeli suveaeg", "IOT": "India ookeani aeg", "IRKST": "Irkutski suveaeg", "IRKT": "Irkutski standardaeg", "IRST": "Iraani standardaeg", "IRST DST": "Iraani suveaeg", "IST": "India aeg", "IST Israel": "Iisraeli standardaeg", "JDT": "Jaapani suveaeg", "JST": "Jaapani standardaeg", "KOST": "Kosrae aeg", "KRAST": "Krasnojarski suveaeg", "KRAT": "Krasnojarski standardaeg", "KST": "Korea standardaeg", "KST DST": "Korea suveaeg", "LHDT": "Lord Howe’i suveaeg", "LHST": "Lord Howe’i standardaeg", "LINT": "Line’i saarte aeg", "MAGST": "Magadani suveaeg", "MAGT": "Magadani standardaeg", "MART": "Markiisaarte aeg", "MAWT": "Mawsoni aeg", "MDT": "Mäestikuvööndi suveaeg", "MESZ": "Kesk-Euroopa suveaeg", "MEZ": "Kesk-Euroopa standardaeg", "MHT": "Marshalli Saarte aeg", "MMT": "Birma aeg", "MSD": "Moskva suveaeg", "MST": "Mäestikuvööndi standardaeg", "MUST": "Mauritiuse suveaeg", "MUT": "Mauritiuse standardaeg", "MVT": "Maldiivi aeg", "MYT": "Malaisia aeg", "NCT": "Uus-Kaledoonia standardaeg", "NDT": "Newfoundlandi suveaeg", "NDT New Caledonia": "Uus-Kaledoonia suveaeg", "NFDT": "Norfolki saare suveaeg", "NFT": "Norfolki saare standardaeg", "NOVST": "Novosibirski suveaeg", "NOVT": "Novosibirski standardaeg", "NPT": "Nepali aeg", "NRT": "Nauru aeg", "NST": "Newfoundlandi standardaeg", "NUT": "Niue aeg", "NZDT": "Uus-Meremaa suveaeg", "NZST": "Uus-Meremaa standardaeg", "OESZ": "Ida-Euroopa suveaeg", "OEZ": "Ida-Euroopa standardaeg", "OMSST": "Omski suveaeg", "OMST": "Omski standardaeg", "PDT": "Vaikse ookeani suveaeg", "PDTM": "Mehhiko Vaikse ookeani suveaeg", "PETDT": "Kamtšatka suveaeg", "PETST": "Kamtšatka standardaeg", "PGT": "Paapua Uus-Guinea aeg", "PHOT": "Fööniksisaarte aeg", "PKT": "Pakistani standardaeg", "PKT DST": "Pakistani suveaeg", "PMDT": "Saint-Pierre’i ja Miqueloni suveaeg", "PMST": "Saint-Pierre’i ja Miqueloni standardaeg", "PONT": "Pohnpei aeg", "PST": "Vaikse ookeani standardaeg", "PST Philippine": "Filipiini standardaeg", "PST Philippine DST": "Filipiini suveaeg", "PST Pitcairn": "Pitcairni aeg", "PSTM": "Mehhiko Vaikse ookeani standardaeg", "PWT": "Belau aeg", "PYST": "Paraguay suveaeg", "PYT": "Paraguay standardaeg", "PYT Korea": "Pyongyangi aeg", "RET": "Réunioni aeg", "ROTT": "Rothera aeg", "SAKST": "Sahhalini suveaeg", "SAKT": "Sahhalini standardaeg", "SAMST": "Samara suveaeg", "SAMT": "Samara standardaeg", "SAST": "Lõuna-Aafrika standardaeg", "SBT": "Saalomoni Saarte aeg", "SCT": "Seišelli aeg", "SGT": "Singapuri standardaeg", "SLST": "Sri Lanka aeg", "SRT": "Suriname aeg", "SST Samoa": "Samoa standardaeg", "SST Samoa Apia": "Apia standardaeg", "SST Samoa Apia DST": "Apia suveaeg", "SST Samoa DST": "Samoa suveaeg", "SYOT": "Syowa aeg", "TAAF": "Prantsuse Antarktiliste ja Lõunaalade aeg", "TAHT": "Tahiti aeg", "TJT": "Tadžikistani aeg", "TKT": "Tokelau aeg", "TLT": "Ida-Timori aeg", "TMST": "Türkmenistani suveaeg", "TMT": "Türkmenistani standardaeg", "TOST": "Tonga suveaeg", "TOT": "Tonga standardaeg", "TVT": "Tuvalu aeg", "TWT": "Taipei standardaeg", "TWT DST": "Taipei suveaeg", "ULAST": "Ulaanbaatari suveaeg", "ULAT": "Ulaanbaatari standardaeg", "UYST": "Uruguay suveaeg", "UYT": "Uruguay standardaeg", "UZT": "Usbekistani standardaeg", "UZT DST": "Usbekistani suveaeg", "VET": "Venezuela aeg", "VLAST": "Vladivostoki suveaeg", "VLAT": "Vladivostoki standardaeg", "VOLST": "Volgogradi suveaeg", "VOLT": "Volgogradi standardaeg", "VOST": "Vostoki aeg", "VUT": "Vanuatu standardaeg", "VUT DST": "Vanuatu suveaeg", "WAKT": "Wake’i aeg", "WARST": "Lääne-Argentina suveaeg", "WART": "Lääne-Argentina standardaeg", "WAST": "Lääne-Aafrika aeg", "WAT": "Lääne-Aafrika aeg", "WESZ": "Lääne-Euroopa suveaeg", "WEZ": "Lääne-Euroopa standardaeg", "WFT": "Wallise ja Futuna aeg", "WGST": "Lääne-Gröönimaa suveaeg", "WGT": "Lääne-Gröönimaa standardaeg", "WIB": "Lääne-Indoneesia aeg", "WIT": "Ida-Indoneesia aeg", "WITA": "Kesk-Indoneesia aeg", "YAKST": "Jakutski suveaeg", "YAKT": "Jakutski standardaeg", "YEKST": "Jekaterinburgi suveaeg", "YEKT": "Jekaterinburgi standardaeg", "YST": "Yukoni aeg", "МСК": "Moskva standardaeg", "اقتاۋ": "Aktau standardaeg", "اقتاۋ قالاسى": "Aktau suveaeg", "اقتوبە": "Aktöbe standardaeg", "اقتوبە قالاسى": "Aktöbe suveaeg", "الماتى": "Almatõ standardaeg", "الماتى قالاسى": "Almatõ suveaeg", "باتىس قازاق ەلى": "Lääne-Kasahstani aeg", "شىعىش قازاق ەلى": "Ida-Kasahstani aeg", "قازاق ەلى": "Kasahstani aeg", "قىرعىزستان": "Kõrgõzstani aeg", "قىزىلوردا": "Kõzõlorda standardaeg", "قىزىلوردا قالاسى": "Kõzõlorda suveaeg", "∅∅∅": "Peruu suveaeg"},
	}
}

// Locale returns the current translators string locale
func (et *et_EE) Locale() string {
	return et.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'et_EE'
func (et *et_EE) PluralsCardinal() []locales.PluralRule {
	return et.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'et_EE'
func (et *et_EE) PluralsOrdinal() []locales.PluralRule {
	return et.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'et_EE'
func (et *et_EE) PluralsRange() []locales.PluralRule {
	return et.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'et_EE'
func (et *et_EE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'et_EE'
func (et *et_EE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'et_EE'
func (et *et_EE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (et *et_EE) MonthAbbreviated(month time.Month) string {
	if len(et.monthsAbbreviated) == 0 {
		return ""
	}
	return et.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (et *et_EE) MonthsAbbreviated() []string {
	return et.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (et *et_EE) MonthNarrow(month time.Month) string {
	if len(et.monthsNarrow) == 0 {
		return ""
	}
	return et.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (et *et_EE) MonthsNarrow() []string {
	return et.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (et *et_EE) MonthWide(month time.Month) string {
	if len(et.monthsWide) == 0 {
		return ""
	}
	return et.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (et *et_EE) MonthsWide() []string {
	return et.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (et *et_EE) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(et.daysAbbreviated) == 0 {
		return ""
	}
	return et.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (et *et_EE) WeekdaysAbbreviated() []string {
	return et.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (et *et_EE) WeekdayNarrow(weekday time.Weekday) string {
	if len(et.daysNarrow) == 0 {
		return ""
	}
	return et.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (et *et_EE) WeekdaysNarrow() []string {
	return et.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (et *et_EE) WeekdayShort(weekday time.Weekday) string {
	if len(et.daysShort) == 0 {
		return ""
	}
	return et.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (et *et_EE) WeekdaysShort() []string {
	return et.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (et *et_EE) WeekdayWide(weekday time.Weekday) string {
	if len(et.daysWide) == 0 {
		return ""
	}
	return et.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (et *et_EE) WeekdaysWide() []string {
	return et.daysWide
}

// Decimal returns the decimal point of number
func (et *et_EE) Decimal() string {
	return et.decimal
}

// Group returns the group of number
func (et *et_EE) Group() string {
	return et.group
}

// Group returns the minus sign of number
func (et *et_EE) Minus() string {
	return et.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'et_EE' and handles both Whole and Real numbers based on 'v'
func (et *et_EE) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(et.group) - 1; j >= 0; j-- {
					b = append(b, et.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(et.minus) - 1; j >= 0; j-- {
			b = append(b, et.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'et_EE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (et *et_EE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(et.minus) - 1; j >= 0; j-- {
			b = append(b, et.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, et.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'et_EE'
func (et *et_EE) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := et.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(et.group) - 1; j >= 0; j-- {
					b = append(b, et.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(et.minus) - 1; j >= 0; j-- {
			b = append(b, et.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, et.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, et.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'et_EE'
// in accounting notation.
func (et *et_EE) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := et.currencies[currency]
	l := len(s) + len(symbol) + 8 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, et.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(et.group) - 1; j >= 0; j-- {
					b = append(b, et.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, et.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, et.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, et.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, et.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'et_EE'
func (et *et_EE) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'et_EE'
func (et *et_EE) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, et.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'et_EE'
func (et *et_EE) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, et.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'et_EE'
func (et *et_EE) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, et.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, et.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'et_EE'
func (et *et_EE) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'et_EE'
func (et *et_EE) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, et.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'et_EE'
func (et *et_EE) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, et.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'et_EE'
func (et *et_EE) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, et.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, et.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := et.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
