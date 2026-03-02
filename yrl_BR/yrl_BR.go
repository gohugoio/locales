package yrl_BR

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type yrl_BR struct {
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
	currencyPositivePrefix string
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'yrl_BR' locale
func New() locales.Translator {
	return &yrl_BR{
		locale:                 "yrl_BR",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: " miu",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: " miu",
		monthsAbbreviated:      []string{"", "ye", "mk", "ms", "id", "pu", "py", "pm", "ps", "pi", "yp", "yy", "ym"},
		monthsNarrow:           []string{"", "Y", "M", "M", "I", "P", "P", "P", "P", "P", "Y", "Y", "Y"},
		monthsWide:             []string{"", "yepé", "mukũi", "musapíri", "irũdí", "pú", "pú-yepé", "pú-mukũi", "pú-musapíri", "pú-irũdí", "yepé-putimaã", "yepé-yepé", "yepé-mukũi"},
		daysAbbreviated:        []string{"mit", "mur", "mmk", "mms", "sup", "yuk", "sau"},
		daysNarrow:             []string{"M", "M", "M", "M", "S", "Y", "S"},
		daysWide:               []string{"mituú", "murakipí", "murakí-mukũi", "murakí-musapíri", "supapá", "yukuakú", "saurú"},
		periodsAbbreviated:     []string{"", ""},
		erasAbbreviated:        []string{"K.s.", "K.a."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Kiristu senũdé", "Kiristu ariré"},
		timezones:              map[string]string{"ACDT": "Ausitaraliya Piterawara Kurasí Ara Hurariyu", "ACST": "Hurariyu Kurasí Ara Acre yara", "ACT": "Hurariyu Retewa Acre yara", "ACWDT": "Ausitaraliya Piterawara-Usidẽtawara Kurasí Ara Hurariyu", "ACWST": "Ausitaraliya Piterawara-Usidẽtawara Hurariyu Retewa", "ADT": "Atalãtku Kurasí Ara Hurariyu", "ADT Arabia": "Arabiya Kurasí Ara Hurariyu", "AEDT": "Ausitaraliya Uriẽtawara Kurasí Ara Hurariyu", "AEST": "Ausitaraliya Uriẽtawara Hurariyu Retewa", "AFT": "Afegãniretãma Hurariyu", "AKDT": "Alasika Kurasí Ara Hurariyu", "AKST": "Alasika Hurariyu Eté", "AMST": "Amasuna Kurasí Ara Hurariyu", "AMST Armenia": "Arimẽniya Kurasí Ara Hurariyu", "AMT": "Amasuna Hurariyu Eté", "AMT Armenia": "Arimẽniya Hurariyu Retewa", "ANAST": "Anadí Kurasí Ara Hurariyu", "ANAT": "Anadí Hurariyu Retewa", "ARST": "Argẽtina Kurasí Ara Hurariyu", "ART": "Argẽtina Hurariyu Retewa", "AST": "Atalãtiku Hurariyu Retewa", "AST Arabia": "Arábiya Hurariyu Retewa", "AWDT": "Ausitaraliya Usidẽtawara Kurasí Ara Hurariyu", "AWST": "Ausitaraliya Usidẽtawara Hurariyu Retewa", "AZST": "Aseriretãma Kurasí Ara Hurariyu", "AZT": "Aseriretãma Hurariyu Retewa", "BDT Bangladesh": "Bãkaradexi Kurasí Ara Horariyu", "BNT": "Burunei Darusaram Hurariyu", "BOT": "Buríwia Hurariyu", "BRST": "Rỹ Kã Óra Brasília tá", "BRT": "Óra Pã Brasília tá", "BST Bangladesh": "Bãkaradexi Horariyu Retewa", "BT": "Butãu Hurariyu", "CAST": "CAST", "CAT": "Afirika Piterawara Hurariyu", "CCT": "Kapuã-ita Kuku-ita Hurariyu", "CDT": "Hurariyu Kurasí Ara Piterawara", "CHADT": "Xatham Kurasí Ara Hurariyu", "CHAST": "Xatham Hurariyu Retewa", "CHUT": "Chuuk Hurariyu", "CKT": "Kapuã-ita Kooki Hurariyu Retewa", "CKT DST": "Kapuã-ita Kooki Kurasí Ara Pitera Hurariyu", "CLST": "Xiri Kurasí Ara Hurariyu", "CLT": "Xiri Hurariyu Retewa", "COST": "Kurũbia Kurasí Ara Hurariyu", "COT": "Kurũbia Hurariyu Retewa", "CST": "Hurariyu Retewa Piterawara", "CST China": "Xina Hurariyu Retewa", "CST China DST": "Xina Kurasí Ara Hurariyu", "CVST": "Kabu Suikiri Kurasí Ara Hurariyu", "CVT": "Kabu Suikiri Hurariyu Retewa", "CXT": "KapuãmaKiritima Hurariyu", "ChST": "Xamoro Hurariyu", "ChST NMI": "Kapuã-ita Mariyãna Nuti suí Hurariyu", "CuDT": "Kuba Kurasí Ara Hurariyu", "CuST": "Kuba Hurariyu Retewa", "DAVT": "Dawi Hurariyu", "DDUT": "Dumont-d’Urville Hurariyu", "EASST": "Pasikuwa Kapuãma Kurasí Ara Hurariyu", "EAST": "Pasikuwa Kapuãma Hurariyu Retewa", "EAT": "Afirika Uriẽtawara Hurariyu", "ECT": "Ekuadú Hurariyu", "EDT": "Hurariyu Kurasí Ara Lesiti", "EGDT": "Guruẽrãdiya Uriẽtawara Kurasí Ara Hurariyu", "EGST": "Guruẽrãdiya Uriẽtawara Hurariyu Retewa", "EST": "Hurariyu Retewa Lesiti yara", "FEET": "Eurupa Lesiti-eté Hurariyu", "FJT": "Fiyi Hurariyu Retewa", "FJT Summer": "Fiyi Kurasí Ara Hurariyu", "FKST": "Kapuã-ita Mawina Kurasí Ara Hurariyu", "FKT": "Kapuã-ita Mawina Hurariyu Retewa", "FNST": "Fenãdu Nuruyã Kurasí Ara Hurariyu", "FNT": "Fenãdu Nuruyã Hurariyu Retewa", "GALT": "Garapagu-ita Hurariyu", "GAMT": "Gãbiere Hurariyu", "GEST": "Geugiya Kurasí Ara Hurariyu", "GET": "Geugiya Hurariyu Retewa", "GFT": "Giyana Frãsa yara Hurariyu", "GIT": "Kapuã-ita Yubetu Hurariyu", "GMT": "Greenwich Miridiyanu yara Hurariyu", "GNSST": "GNSST", "GNST": "GNST", "GST": "Golfo Hurariyu", "GST Guam": "Guwã Hurariyu Retewa", "GYT": "Giyana Hurariyu", "HADT": "Hawaí asuí Kapuã-ita Areuta-ita Kurasí Ara Hurariyu", "HAST": "Hawaí asuí Kapuã-ita Areuta-ita Hurariyu Retewa", "HKST": "Hũg Kũg Kurasí Ara Hurariyu", "HKT": "Hũg Kũg Hurariyu Retewa", "HOVST": "Howidi Kurasí Ara Hurariyu", "HOVT": "Howidi Hurariyu Retewa", "ICT": "Ĩdoxina Hurariyu", "IDT": "Isirayeu Kurasí Ara Hurariyu", "IOT": "Useyanu Ĩdiku Hurariyu", "IRKST": "Irkutisiki Kurasí Ara Hurariyu", "IRKT": "Irkutisiki Hurariyu Retewa", "IRST": "Irã Hurariyu Retewa", "IRST DST": "Irã Kurasí Ara Hurariyu", "IST": "Ĩdia Hurariyu Retewa", "IST Israel": "Isirayeu Hurariyu Retewa", "JDT": "Nipõ Kurasí Ara Hurariyu", "JST": "Nipõ Hurariyu Retewa", "KOST": "Kusirai Hurariyu", "KRAST": "Karasinoyarisiki Kurasí Ara Hurariyu", "KRAT": "Karasinoyarisiki Hurariyu Retewa", "KST": "Kureya Hurariyu Retewa", "KST DST": "Kureya Kurasí Ara Hurariyu", "LHDT": "Lurudi Howe Kurasí Ara Hurariyu", "LHST": "Lurudi Howe Hurariyu Retewa", "LINT": "Kapuã-ita Inĩbu Hurariyu", "MAGST": "Magadã Kurasí Ara Hurariyu", "MAGT": "Magadã Hurariyu Retewa", "MART": "Makesa-ita Hurariyu", "MAWT": "Mausũ Hurariyu", "MDT": "Makau Kurasí Ara Hurariyu", "MESZ": "Eurupa Piterawara Kurasí Ara Hurariyu", "MEZ": "Eurupa Piterawara Hurariyu Retewa", "MHT": "Kapuã-ita Marshall Hurariyu", "MMT": "Miyamá Hurariyu", "MSD": "Moskou Kurasí Ara Hurariyu", "MST": "Makau Hurariyu Retewa", "MUST": "Maurisiyu Kurasí Ara Hurariyu", "MUT": "Maurisiyu Hurariyu Retewa", "MVT": "Kapuã-ita Maudiwa-ita Hurariyu", "MYT": "Malasiya Hurariyu", "NCT": "Karedũniya Pisasú Hurariyu Retewa", "NDT": "Iwí Pisasú Kurasí Ara Hurariyu", "NDT New Caledonia": "Karedũniya Pisasú Kurasí Ara Hurariyu", "NFDT": "Kapuãma Norfolk Kurasí Ara Hurariyu", "NFT": "Kapuãma Norfolk Hurariyu Retewa", "NOVST": "Sibirisiki Pisasú Kurasí Ara Hurariyu", "NOVT": "Sibirisiki Pisasú Hurariyu Retewa", "NPT": "Nepau Hurariyu", "NRT": "Nauru Hurariyu", "NST": "Iwí Pisasú Hurariyu Retewa", "NUT": "Niwe Hurariyu", "NZDT": "Serãdiya Pisasú Kurasí Ara Hurariyu", "NZST": "Serãdiya Pisasú Hurariyu Retewa", "OESZ": "Eurupa Uriẽtawara Kurasí Ara Hurariyu", "OEZ": "Eurupa Uriẽtawara Hurariyu Retewa", "OMSST": "Omisiki Kurasí Ara Hurariyu", "OMST": "Omisiki Hurariyu Retewa", "PDT": "Rỹ Kã Óra Pasifiku tá", "PDTM": "Pasifiku Mexikanu Kurasí Ara Hurariyu", "PETDT": "Petropavlovsk-Kamchatski Kurasí Ara Hurariyu", "PETST": "Petropavlovsk-Kamchatski Hurariyu Retewa", "PGT": "Papuwa-Giné Pisasú Hurariyu", "PHOT": "Kapuã-ita Fẽnix", "PKT": "Pakiretãma Hurariyu Retewa", "PKT DST": "Pakiretãma Kurasí Ara Hurariyu", "PMDT": "Sã Peduru asuí Mikirãu Kurasí Ara Hurariyu", "PMST": "Sã Peduru asuí Mikirãu Hurariyu Retewa", "PONT": "Ponape Hurariyu", "PST": "Óra Pã Pasifiku tá", "PST Philippine": "Firipina Hurariyu Retewa", "PST Philippine DST": "Firipina Kurasí Ara Hurariyu", "PST Pitcairn": "Pitcairn Hurariyu", "PSTM": "Pasifiku Mexikanu Hurariyu Retewa", "PWT": "Parau Hurariyu", "PYST": "ParaguwaiKurasí Ara Hurariyu", "PYT": "Paraguwai Hurariyu Retewa", "PYT Korea": "Piúgiãgi Hurariyu", "RET": "Yumuatirisawa Hurariyu", "ROTT": "Rotera Hurariyu", "SAKST": "Sakarina Kurasí Ara Hurariyu", "SAKT": "Sakarina Hurariyu Retewa", "SAMST": "Samara Kurasí Ara Hurariyu", "SAMT": "Samara Hurariyu Retewa", "SAST": "Afirika Su suí Hurariyu", "SBT": "Kapuãma-ita Sarumũ Hurariyu", "SCT": "Seixeri Hurariyu", "SGT": "Sĩgapura Hurariyu Retewa", "SLST": "Rãka Hurariyu", "SRT": "Suriname Hurariyu", "SST Samoa": "Samowa Hurariyu Retewa", "SST Samoa Apia": "Apiya Uraruiyu Retewa", "SST Samoa Apia DST": "Apiya Kurasí Ara Hurariyu", "SST Samoa DST": "Samowa Kurasí Ara Hurariyu", "SYOT": "Siyowa Hurariyu", "TAAF": "Tetãma Su suí Frãsa yara asuí Ãtartida Hurariyu", "TAHT": "Taiti Hurariyu", "TJT": "Tayikiretãma Hurariyu", "TKT": "Tokerau Hurariyu", "TLT": "Timu-Semusawa Hurariyu", "TMST": "Turkuranaretãma Kurasí Ara Hurariyu", "TMT": "Turkuranaretãma Hurariyu Retewa", "TOST": "Tõga Kurasí Ara Hurariyu", "TOT": "Tõga Hurariyu Retewa", "TVT": "Tuvaru Hurariyu", "TWT": "Taipei Hurariyu Retewa", "TWT DST": "Taipei Kurasí Ara Hurariyu", "ULAST": "Urã Baturu Kurasí Ara Hurariyu", "ULAT": "Urã Baturu Hurariyu Retewa", "UYST": "Uruguwai Kurasí Ara Hurariyu", "UYT": "Uruguwai Hurariyu Retewa", "UZT": "Yũbuesara-retãma Hurariyu Retewa", "UZT DST": "Yũbuesara-retãma Kurasí Ara Hurariyu", "VET": "Wenẽsuera Hurariyu", "VLAST": "Waradiwosituki Kurasí Ara Hurariyu", "VLAT": "Waradiwosituki Hurariyu Retewa", "VOLST": "Worwogaradu Kurasí Ara Hurariyu", "VOLT": "Worwogaradu Hurariyu Retewa", "VOST": "Wosituki Hurariyu", "VUT": "Wanuatu Hurariyu Retewa", "VUT DST": "Wanuatu Kurasí Ara Hurariyu", "WAKT": "Kapuã-ita Wake", "WARST": "Argẽtina Usidẽtawara Kurasí Ara Hurariyu", "WART": "Argẽtina Usidẽtawara Hurariyu Retewa", "WAST": "Afirikia Usidẽtawara Hurariyu", "WAT": "Afirikia Usidẽtawara Hurariyu", "WESZ": "Eurupa Usidẽtawara Kurasí Ara Hurariyu", "WEZ": "Eurupa Usidẽtawara Hurariyu Retewa", "WFT": "Wari asuí Futuna Hurariyu", "WGST": "Guruẽrãdiya Usidẽtawara Kurasí Ara Hurariyu", "WGT": "Guruẽrãdiya Usidẽtawara Hurariyu Retewa", "WIB": "Ĩdonesiya Usidẽtawara Hurariyu", "WIT": "Ĩdonesiya Uriẽtawara Hurariyu", "WITA": "Ĩdonesiya Piterawara Hurariyu", "YAKST": "Yakutisiki Kurasí Ara Hurariyu", "YAKT": "Yakutisiki Hurariyu Retewa", "YEKST": "Ekaterĩbugu Kurasí Ara Hurariyu", "YEKT": "Ekaterĩbugu Hurariyu Retewa", "YST": "YST", "МСК": "Moskou Hurariyu Retewa", "اقتاۋ": "Akitau Hurariyu Retewa", "اقتاۋ قالاسى": "Akitau Kurasí Ara Hurariyu", "اقتوبە": "Akitubi Hurariyu Retewa", "اقتوبە قالاسى": "Akitubi Kurasí Ara Hurariyu", "الماتى": "Aumati Hurariyu Eté", "الماتى قالاسى": "Aumati Kurasí Ara Hurariyu", "باتىس قازاق ەلى": "Kasakiretãma Usidẽtawara Hurariyu", "شىعىش قازاق ەلى": "Kasakiretãma Uriẽtawara Hurariyu", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Kirigiretãma Hurariyu", "قىزىلوردا": "Kisiroda Hurariyu Retewa", "قىزىلوردا قالاسى": "Kisiroda Kurasí Ara Hurariyu", "∅∅∅": "Asori-ita Kurasí Ara Hurariyu"},
	}
}

// Locale returns the current translators string locale
func (yrl *yrl_BR) Locale() string {
	return yrl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yrl_BR'
func (yrl *yrl_BR) PluralsCardinal() []locales.PluralRule {
	return yrl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yrl_BR'
func (yrl *yrl_BR) PluralsOrdinal() []locales.PluralRule {
	return yrl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yrl_BR'
func (yrl *yrl_BR) PluralsRange() []locales.PluralRule {
	return yrl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yrl_BR'
func (yrl *yrl_BR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yrl_BR'
func (yrl *yrl_BR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yrl_BR'
func (yrl *yrl_BR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yrl *yrl_BR) MonthAbbreviated(month time.Month) string {
	return yrl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yrl *yrl_BR) MonthsAbbreviated() []string {
	return yrl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yrl *yrl_BR) MonthNarrow(month time.Month) string {
	return yrl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yrl *yrl_BR) MonthsNarrow() []string {
	return yrl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yrl *yrl_BR) MonthWide(month time.Month) string {
	return yrl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yrl *yrl_BR) MonthsWide() []string {
	return yrl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yrl *yrl_BR) WeekdayAbbreviated(weekday time.Weekday) string {
	return yrl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yrl *yrl_BR) WeekdaysAbbreviated() []string {
	return yrl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yrl *yrl_BR) WeekdayNarrow(weekday time.Weekday) string {
	return yrl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yrl *yrl_BR) WeekdaysNarrow() []string {
	return yrl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yrl *yrl_BR) WeekdayShort(weekday time.Weekday) string {
	return yrl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yrl *yrl_BR) WeekdaysShort() []string {
	return yrl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yrl *yrl_BR) WeekdayWide(weekday time.Weekday) string {
	return yrl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yrl *yrl_BR) WeekdaysWide() []string {
	return yrl.daysWide
}

// Decimal returns the decimal point of number
func (yrl *yrl_BR) Decimal() string {
	return yrl.decimal
}

// Group returns the group of number
func (yrl *yrl_BR) Group() string {
	return yrl.group
}

// Group returns the minus sign of number
func (yrl *yrl_BR) Minus() string {
	return yrl.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yrl_BR' and handles both Whole and Real numbers based on 'v'
func (yrl *yrl_BR) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yrl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yrl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yrl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yrl_BR' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yrl *yrl_BR) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yrl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yrl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yrl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yrl_BR'
func (yrl *yrl_BR) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yrl.currencies[currency]
	l := len(s) + len(symbol) + 9

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yrl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(yrl.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, yrl.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, yrl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yrl.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yrl_BR'
// in accounting notation.
func (yrl *yrl_BR) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yrl.currencies[currency]
	l := len(s) + len(symbol) + 9

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yrl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(yrl.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, yrl.currencyNegativePrefix[j])
		}

		b = append(b, yrl.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(yrl.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, yrl.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, yrl.currencyNegativeSuffix...)
	} else {
		b = append(b, yrl.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yrl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yrl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, yrl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yrl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'yrl_BR'
func (yrl *yrl_BR) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yrl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := yrl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
