package az_Latn

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type az_Latn struct {
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

// New returns a new instance of translator for the 'az_Latn' locale
func New() locales.Translator {
	return &az_Latn{
		locale:                 "az_Latn",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 4, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "₼", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "ley", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "S£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "yan", "fev", "mar", "apr", "may", "iyn", "iyl", "avq", "sen", "okt", "noy", "dek"},
		monthsWide:             []string{"", "yanvar", "fevral", "mart", "aprel", "may", "iyun", "iyul", "avqust", "sentyabr", "oktyabr", "noyabr", "dekabr"},
		daysAbbreviated:        []string{"B.", "B.e.", "Ç.a.", "Ç.", "C.a.", "C.", "Ş."},
		daysNarrow:             []string{"7", "1", "2", "3", "4", "5", "6"},
		daysShort:              []string{"B.", "B.E.", "Ç.A.", "Ç.", "C.A.", "C.", "Ş."},
		daysWide:               []string{"bazar", "bazar ertəsi", "çərşənbə axşamı", "çərşənbə", "cümə axşamı", "cümə", "şənbə"},
		timezones:              map[string]string{"ACDT": "Mərkəzi Avstraliya Yay Vaxtı", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Mərkəzi Qərbi Avstraliya Yay Vaxtı", "ACWST": "Mərkəzi Qərbi Avstraliya Standart Vaxtı", "ADT": "Atlantik Yay Vaxtı", "ADT Arabia": "Ərəbistan Yay Vaxtı", "AEDT": "Şərqi Avstraliya Yay Vaxtı", "AEST": "Şərqi Avstraliya Standart Vaxtı", "AFT": "Əfqanıstan Vaxtı", "AKDT": "Alyaska Yay Vaxtı", "AKST": "Alyaska Standart Vaxtı", "AMST": "Amazon Yay Vaxtı", "AMST Armenia": "Ermənistan Yay Vaxtı", "AMT": "Amazon Standart Vaxtı", "AMT Armenia": "Ermənistan Standart Vaxtı", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Argentina Yay Vaxtı", "ART": "Argentina Standart Vaxtı", "AST": "Atlantik Standart Vaxt", "AST Arabia": "Ərəbistan Standart Vaxtı", "AWDT": "Qərbi Avstraliya Yay Vaxtı", "AWST": "Qərbi Avstraliya Standart Vaxtı", "AZST": "Azərbaycan Yay Vaxtı", "AZT": "Azərbaycan Standart Vaxtı", "BDT Bangladesh": "Banqladeş Yay Vaxtı", "BNT": "Brunei Darussalam vaxtı", "BOT": "Boliviya Vaxtı", "BRST": "Braziliya Yay Vaxtı", "BRT": "Braziliya Standart Vaxtı", "BST Bangladesh": "Banqladeş Standart Vaxtı", "BT": "Butan Vaxtı", "CAST": "CAST", "CAT": "Mərkəzi Afrika Vaxtı", "CCT": "Kokos Adaları Vaxtı", "CDT": "Mərkəzi Amerika yay vaxtı", "CHADT": "Çatham Yay Vaxtı", "CHAST": "Çatham Standart Vaxtı", "CHUT": "Çuuk Vaxtı", "CKT": "Kuk Adaları Standart Vaxtı", "CKT DST": "Kuk Adaları Yarım Yay Vaxtı", "CLST": "Çili Yay Vaxtı", "CLT": "Çili Standart Vaxtı", "COST": "Kolumbiya Yay Vaxtı", "COT": "Kolumbiya Standart Vaxtı", "CST": "Mərkəzi Amerika Standart Vaxtı", "CST China": "Çin Standart Vaxtı", "CST China DST": "Çin Yay Vaxtı", "CVST": "Kape Verde Yay Vaxtı", "CVT": "Kape Verde Standart Vaxtı", "CXT": "Milad Adası Vaxtı", "ChST": "Çamorro Vaxtı", "ChST NMI": "ChST NMI", "CuDT": "Kuba Yay Vaxtı", "CuST": "Kuba Standart Vaxtı", "DAVT": "Devis Vaxtı", "DDUT": "Dümon-d’Ürvil Vaxtı", "EASST": "Pasxa Adası Yay Vaxtı", "EAST": "Pasxa Adası Standart Vaxtı", "EAT": "Şərqi Afrika Vaxtı", "ECT": "Ekvador Vaxtı", "EDT": "Şimali Şərqi Amerika Yay Vaxtı", "EGDT": "Şərqi Qrenlandiya Yay Vaxtı", "EGST": "Şərqi Qrenlandiya Standart Vaxtı", "EST": "Şimali Şərqi Amerika Standart Vaxtı", "FEET": "Kənar Şərqi Avropa Vaxtı", "FJT": "Fici Standart Vaxtı", "FJT Summer": "Fici Yay Vaxtı", "FKST": "Folklend Adaları Yay Vaxtı", "FKT": "Folklend Adaları Standart Vaxtı", "FNST": "Fernando de Noronya Yay Vaxtı", "FNT": "Fernando de Noronya Standart Vaxtı", "GALT": "Qalapaqos Vaxtı", "GAMT": "Qambier Vaxtı", "GEST": "Gurcüstan Yay Vaxtı", "GET": "Gurcüstan Standart Vaxtı", "GFT": "Fransız Qvianası Vaxtı", "GIT": "Gilbert Adaları Vaxtı", "GMT": "Qrinviç Orta Vaxtı", "GNSST": "GNSST", "GNST": "GNST", "GST": "Cənubi Corciya Vaxtı", "GST Guam": "GST Guam", "GYT": "Qayana Vaxtı", "HADT": "Havay-Aleut Standart Vaxtı", "HAST": "Havay-Aleut Standart Vaxtı", "HKST": "Honq Konq Yay Vaxtı", "HKT": "Honq Konq Standart Vaxtı", "HOVST": "Hovd Yay Vaxtı", "HOVT": "Hovd Standart Vaxtı", "ICT": "Hindçin Vaxtı", "IDT": "İsrail Yay Vaxtı", "IOT": "Hind Okeanı Vaxtı", "IRKST": "İrkutsk Yay Vaxtı", "IRKT": "İrkutsk Standart Vaxtı", "IRST": "İran Standart Vaxtı", "IRST DST": "İran Yay Vaxtı", "IST": "Hindistan Vaxtı", "IST Israel": "İsrail Standart Vaxtı", "JDT": "Yaponiya Yay Vaxtı", "JST": "Yaponiya Standart Vaxtı", "KOST": "Korse Vaxtı", "KRAST": "Krasnoyarsk Yay Vaxtı", "KRAT": "Krasnoyarsk Standart Vaxtı", "KST": "Koreya Standart Vaxtı", "KST DST": "Koreya Yay Vaxtı", "LHDT": "Lord Hau Yay vaxtı", "LHST": "Lord Hau Standart Vaxtı", "LINT": "Layn Adaları Vaxtı", "MAGST": "Maqadan Yay Vaxtı", "MAGT": "Maqadan Standart Vaxtı", "MART": "Markesas Vaxtı", "MAWT": "Mouson Vaxtı", "MDT": "MDT", "MESZ": "Mərkəzi Avropa Yay Vaxtı", "MEZ": "Mərkəzi Avropa Standart Vaxtı", "MHT": "Marşal Adaları Vaxtı", "MMT": "Myanma Vaxtı", "MSD": "Moskva Yay vaxtı", "MST": "MST", "MUST": "Mavriki Yay Vaxtı", "MUT": "Mavriki Standart Vaxtı", "MVT": "Maldiv Vaxtı", "MYT": "Malayziya Vaxtı", "NCT": "Yeni Kaledoniya Standart Vaxtı", "NDT": "Nyufaundlend Yay Vaxtı", "NDT New Caledonia": "Yeni Kaledoniya Yay Vaxtı", "NFDT": "Norfolk Adası Yay Vaxtı", "NFT": "Norfolk Adası Standart Vaxtı", "NOVST": "Novosibirsk Yay Vaxtı", "NOVT": "Novosibirsk Standart Vaxtı", "NPT": "Nepal vaxtı", "NRT": "Nauru Vaxtı", "NST": "Nyufaundlend Standart Vaxtı", "NUT": "Niue Vaxtı", "NZDT": "Yeni Zelandiya Yay Vaxtı", "NZST": "Yeni Zelandiya Standart Vaxtı", "OESZ": "Şərqi Avropa Yay Vaxtı", "OEZ": "Şərqi Avropa Standart Vaxtı", "OMSST": "Omsk Yay Vaxtı", "OMST": "Omsk Standart Vaxtı", "PDT": "Şimali Amerika Sakit Okean Yay Vaxtı", "PDTM": "Meksika Sakit Okean Yay Vaxtı", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papua Yeni Qvineya Vaxtı", "PHOT": "Feniks Adaları Vaxtı", "PKT": "Pakistan Standart vaxtı", "PKT DST": "Pakistan Yay Vaxtı", "PMDT": "Müqəddəs Pyer və Mikelon Yay Vaxtı", "PMST": "Müqəddəs Pyer və Mikelon Standart Vaxtı", "PONT": "Ponape Vaxtı", "PST": "Şimali Amerika Sakit Okean Standart Vaxtı", "PST Philippine": "Filippin Standart Vaxtı", "PST Philippine DST": "Filippin Yay Vaxtı", "PST Pitcairn": "Pitkern Vaxtı", "PSTM": "Meksika Sakit Okean Standart Vaxtı", "PWT": "Palau Vaxtı", "PYST": "Paraqvay Yay Vaxtı", "PYT": "Paraqvay Standart Vaxtı", "PYT Korea": "Pxenyan Vaxtı", "RET": "Reyunyon", "ROTT": "Rotera Vaxtı", "SAKST": "Saxalin Yay Vaxtı", "SAKT": "Saxalin Standart Vaxtı", "SAMST": "Samara yay vaxtı", "SAMT": "Samara standart vaxtı", "SAST": "Cənubi Afrika Vaxtı", "SBT": "Solomon Adaları Vaxtı", "SCT": "Seyşel Adaları Vaxtı", "SGT": "Sinqapur Vaxtı", "SLST": "SLST", "SRT": "Surinam Vaxtı", "SST Samoa": "Samoa Standart Vaxtı", "SST Samoa Apia": "Apia Standart Vaxtı", "SST Samoa Apia DST": "Apia Yay Vaxtı", "SST Samoa DST": "Samoa Yay Vaxtı", "SYOT": "Syova Vaxtı", "TAAF": "Fransız Cənubi və Antarktik Vaxtı", "TAHT": "Tahiti Vaxtı", "TJT": "Tacikistan Vaxtı", "TKT": "Tokelau Vaxtı", "TLT": "Şərqi Timor Vaxtı", "TMST": "Türkmənistan Yay Vaxtı", "TMT": "Türkmənistan Standart Vaxtı", "TOST": "Tonqa Yay Vaxtı", "TOT": "Tonqa Standart Vaxtı", "TVT": "Tuvalu Vaxtı", "TWT": "Taybey Standart Vaxtı", "TWT DST": "Taybey Yay Vaxtı", "ULAST": "Ulanbator Yay Vaxtı", "ULAT": "Ulanbator Standart Vaxtı", "UYST": "Uruqvay Yay Vaxtı", "UYT": "Uruqvay Standart Vaxtı", "UZT": "Özbəkistan Standart Vaxtı", "UZT DST": "Özbəkistan Yay Vaxtı", "VET": "Venesuela Vaxtı", "VLAST": "Vladivostok Yay Vaxtı", "VLAT": "Vladivostok Standart Vaxtı", "VOLST": "Volqoqrad Yay Vaxtı", "VOLT": "Volqoqrad Standart Vaxtı", "VOST": "Vostok Vaxtı", "VUT": "Vanuatu Standart Vaxtı", "VUT DST": "Vaunatu Yay Vaxtı", "WAKT": "Ueyk Vaxtı", "WARST": "Qərbi Argentina Yay Vaxtı", "WART": "Qərbi Argentina Standart Vaxtı", "WAST": "Qərbi Afrika Vaxtı", "WAT": "Qərbi Afrika Vaxtı", "WESZ": "Qərbi Avropa Yay Vaxtı", "WEZ": "Qərbi Avropa Standart Vaxtı", "WFT": "Uollis və Futuna Vaxtı", "WGST": "Qərbi Qrenlandiya Yay Vaxtı", "WGT": "Qərbi Qrenlandiya Standart Vaxtı", "WIB": "Qərbi İndoneziya Vaxtı", "WIT": "Şərqi İndoneziya Vaxtı", "WITA": "Mərkəzi İndoneziya Vaxtı", "YAKST": "Yakutsk Yay Vaxtı", "YAKT": "Yakutsk Standart Vaxtı", "YEKST": "Yekaterinburq Yay Vaxtı", "YEKT": "Yekaterinburq Standart Vaxtı", "YST": "Yukon Vaxtı", "МСК": "Moskva Standart Vaxtı", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Qərbi Qazaxıstan Vaxtı", "شىعىش قازاق ەلى": "Şərqi Qazaxıstan Vaxtı", "قازاق ەلى": "Qazaxıstan vaxtı", "قىرعىزستان": "Qırğızıstan Vaxtı", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Azor Yay Vaxtı"},
	}
}

// Locale returns the current translators string locale
func (az *az_Latn) Locale() string {
	return az.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'az_Latn'
func (az *az_Latn) PluralsCardinal() []locales.PluralRule {
	return az.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'az_Latn'
func (az *az_Latn) PluralsOrdinal() []locales.PluralRule {
	return az.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'az_Latn'
func (az *az_Latn) PluralsRange() []locales.PluralRule {
	return az.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'az_Latn'
func (az *az_Latn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'az_Latn'
func (az *az_Latn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod1000 := i % 1000
	iMod10 := i % 10
	iMod100 := i % 100

	if (iMod10 == 1 || iMod10 == 2 || iMod10 == 5 || iMod10 == 7 || iMod10 == 8) || (iMod100 == 20 || iMod100 == 50 || iMod100 == 70 || iMod100 == 80) {
		return locales.PluralRuleOne
	} else if (iMod10 == 3 || iMod10 == 4) || (iMod1000 == 100 || iMod1000 == 200 || iMod1000 == 300 || iMod1000 == 400 || iMod1000 == 500 || iMod1000 == 600 || iMod1000 == 700 || iMod1000 == 800 || iMod1000 == 900) {
		return locales.PluralRuleFew
	} else if (i == 0) || (iMod10 == 6) || (iMod100 == 40 || iMod100 == 60 || iMod100 == 90) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'az_Latn'
func (az *az_Latn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := az.CardinalPluralRule(num1, v1)
	end := az.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (az *az_Latn) MonthAbbreviated(month time.Month) string {
	if len(az.monthsAbbreviated) == 0 {
		return ""
	}
	return az.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (az *az_Latn) MonthsAbbreviated() []string {
	return az.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (az *az_Latn) MonthNarrow(month time.Month) string {
	if len(az.monthsNarrow) == 0 {
		return ""
	}
	return az.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (az *az_Latn) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (az *az_Latn) MonthWide(month time.Month) string {
	if len(az.monthsWide) == 0 {
		return ""
	}
	return az.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (az *az_Latn) MonthsWide() []string {
	return az.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (az *az_Latn) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(az.daysAbbreviated) == 0 {
		return ""
	}
	return az.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (az *az_Latn) WeekdaysAbbreviated() []string {
	return az.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (az *az_Latn) WeekdayNarrow(weekday time.Weekday) string {
	if len(az.daysNarrow) == 0 {
		return ""
	}
	return az.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (az *az_Latn) WeekdaysNarrow() []string {
	return az.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (az *az_Latn) WeekdayShort(weekday time.Weekday) string {
	if len(az.daysShort) == 0 {
		return ""
	}
	return az.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (az *az_Latn) WeekdaysShort() []string {
	return az.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (az *az_Latn) WeekdayWide(weekday time.Weekday) string {
	if len(az.daysWide) == 0 {
		return ""
	}
	return az.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (az *az_Latn) WeekdaysWide() []string {
	return az.daysWide
}

// Decimal returns the decimal point of number
func (az *az_Latn) Decimal() string {
	return az.decimal
}

// Group returns the group of number
func (az *az_Latn) Group() string {
	return az.group
}

// Group returns the minus sign of number
func (az *az_Latn) Minus() string {
	return az.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'az_Latn' and handles both Whole and Real numbers based on 'v'
func (az *az_Latn) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'az_Latn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (az *az_Latn) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, az.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'az_Latn'
func (az *az_Latn) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := az.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, az.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, az.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'az_Latn'
// in accounting notation.
func (az *az_Latn) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := az.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, az.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, az.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, az.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, az.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'az_Latn'
func (az *az_Latn) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'az_Latn'
func (az *az_Latn) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'az_Latn'
func (az *az_Latn) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'az_Latn'
func (az *az_Latn) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, az.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'az_Latn'
func (az *az_Latn) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'az_Latn'
func (az *az_Latn) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'az_Latn'
func (az *az_Latn) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'az_Latn'
func (az *az_Latn) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := az.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
