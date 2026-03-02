package blo_BJ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type blo_BJ struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentPrefix          string
	timeSeparator          string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'blo_BJ' locale
func New() locales.Translator {
	return &blo_BJ{
		locale:                 "blo_BJ",
		pluralsCardinal:        []locales.PluralRule{1, 2, 6},
		pluralsOrdinal:         []locales.PluralRule{1, 2, 4, 6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentPrefix:          " ",
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " -",
		monthsAbbreviated:      []string{"", "kaw", "kpa", "ci", "ɖʊ", "ɖu5", "ɖu6", "la", "kǝu", "fʊm", "cim", "pom", "bʊn"},
		monthsWide:             []string{"", "ɩjikawǝrka kaŋɔrɔ", "ɩjikpaka kaŋɔrɔ", "arɛ́cika kaŋɔrɔ", "njɩbɔ nɖʊka kaŋɔrɔ", "acafʊnɖuka kaŋɔrɔ", "anɔɔɖuka kaŋɔrɔ", "alàlaka kaŋɔrɔ", "ɩjikǝuka kaŋɔrɔ", "abofʊmka kaŋɔrɔ", "ɩjicimka kaŋɔrɔ", "acapomka kaŋɔrɔ", "anɔɔbʊnka kaŋɔrɔ"},
		daysAbbreviated:        []string{"alah", "aɖɩt", "atal", "alar", "alam", "arɩs", "asib"},
		daysNarrow:             []string{"lh", "ɖt", "tl", "lr", "lm", "rs", "sb"},
		daysShort:              []string{"alh", "aɖt", "atl", "alr", "alm", "ars", "asb"},
		daysWide:               []string{"alahaɖɩ", "aɖɩtɛnɛɛ", "atalaata", "alaarba", "alaamɩshɩ", "arɩsǝma", "asiibi"},
		timezones:              map[string]string{"ACDT": "Ɔstraliya kagɩcɩɩca kaakɔŋkɔŋɔ̀ gafʊbaka", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Ɔstraliya kagɩcɩɩca gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "ACWST": "Ɔstraliya kagɩcɩɩca gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "ADT": "Amalɩka katǝlantika kaakɔŋkɔŋɔ̀ gafʊbaka", "ADT Arabia": "Galaaributǝna kaakɔŋkɔŋɔ̀ gafʊbaka", "AEDT": "Ɔstraliya kaajakalaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "AEST": "Ɔstraliya kaajakalaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "AFT": "Afganistan kaakɔŋkɔŋɔ̀", "AKDT": "Alaskaa kaakɔŋkɔŋɔ̀ gafʊbaka", "AKST": "Alaskaa kaakɔŋkɔŋɔ̀ ɖeiɖei", "AMST": "Amasɔn kaakɔŋkɔŋɔ̀ gafʊbaka", "AMST Armenia": "Armenii kaakɔŋkɔŋɔ̀ gafʊbaka", "AMT": "Amasɔn kaakɔŋkɔŋɔ̀ ɖeiɖei", "AMT Armenia": "Armenii kaakɔŋkɔŋɔ̀ ɖeiɖei", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Arjantin kaakɔŋkɔŋɔ̀ gafʊbaka", "ART": "Arjantin kaakɔŋkɔŋɔ̀ ɖeiɖei", "AST": "Amalɩka katǝlantika kaakɔŋkɔŋɔ̀ ɖeiɖei", "AST Arabia": "Galaaributǝna kaakɔŋkɔŋɔ̀ ɖeiɖei", "AWDT": "Ɔstraliya kagɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "AWST": "Ɔstraliya kagɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "AZST": "Asɛrbaɩjaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "AZT": "Asɛrbaɩjaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "BDT Bangladesh": "Baŋglaɖɛɛshɩ kaakɔŋkɔŋɔ̀ gafʊbaka", "BNT": "Brunɛɩ kaakɔŋkɔŋɔ̀", "BOT": "Bolifiya kaakɔŋkɔŋɔ̀", "BRST": "Brasiliya kaakɔŋkɔŋɔ̀ gafʊbaka", "BRT": "Brasiliya kaakɔŋkɔŋɔ̀ ɖeiɖei", "BST Bangladesh": "Baŋglaɖɛɛshɩ kaakɔŋkɔŋɔ̀ ɖeiɖei", "BT": "Butan kaakɔŋkɔŋɔ̀", "CAST": "CAST", "CAT": "Garɩɖontǝna gɩcɩɩca kaakɔŋkɔŋɔ̀", "CCT": "Kokoos kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "CDT": "Amalɩka gʊnyɩpɛnɛlaŋ kagɩcɩɩca kaakɔŋkɔŋɔ̀ gafʊbaka", "CHADT": "Shatam kaakɔŋkɔŋɔ̀ gafʊbaka", "CHAST": "Shatam kaakɔŋkɔŋɔ̀ ɖeiɖei", "CHUT": "Cuuk kaakɔŋkɔŋɔ̀", "CKT": "Kʊkʊ kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀ ɖeiɖei", "CKT DST": "Kʊkʊ kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀ gafʊbaka", "CLST": "Shilii kaakɔŋkɔŋɔ̀ gafʊbaka", "CLT": "Shilii kaakɔŋkɔŋɔ̀ ɖeiɖei", "COST": "Kolɔmbii kaakɔŋkɔŋɔ̀ gafʊbaka", "COT": "Kolɔmbii kaakɔŋkɔŋɔ̀ ɖeiɖei", "CST": "Amalɩka gʊnyɩpɛnɛlaŋ kagɩcɩɩca kaakɔŋkɔŋɔ̀ ɖeiɖei", "CST China": "Caɩna kaakɔŋkɔŋɔ̀ ɖeiɖei", "CST China DST": "Caɩna kaakɔŋkɔŋɔ̀ gafʊbaka", "CVST": "Kapfɛɛr kaakɔŋkɔŋɔ̀ gafʊbaka", "CVT": "Kapfɛɛr kaakɔŋkɔŋɔ̀ ɖeiɖei", "CXT": "Nowɛl kaAtukǝltǝna kaakɔŋkɔŋɔ̀", "ChST": "Shamoroo kaakɔŋkɔŋɔ̀", "ChST NMI": "Mariyan kǝbʊtukǝltǝna gʊnyɩpɛnɛlaŋ kaakɔŋkɔŋɔ̀", "CuDT": "Kubaa kaakɔŋkɔŋɔ̀ gafʊbaka", "CuST": "Kubaa kaakɔŋkɔŋɔ̀ ɖeiɖei", "DAVT": "Ɖefis kaakɔŋkɔŋɔ̀", "DDUT": "Ɖimɔn Ɖirfil kaakɔŋkɔŋɔ̀", "EASST": "Paakɩ kaAtukǝltǝna kaakɔŋkɔŋɔ̀ gafʊbaka", "EAST": "Paakɩ kaAtukǝltǝna kaakɔŋkɔŋɔ̀ ɖeiɖei", "EAT": "Garɩɖontǝna gajakalaŋ kaakɔŋkɔŋɔ̀", "ECT": "Ekuwaɖɔɔr kaakɔŋkɔŋɔ̀", "EDT": "Amalɩka gʊnyɩpɛnɛlaŋ kaajakalaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "EGDT": "Grinlanɖ gajakalaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "EGST": "Grinlanɖ gajakalaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "EST": "Amalɩka gʊnyɩpɛnɛlaŋ kaajakalaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "FEET": "Garɩfɔntǝna gajakalaŋ kaajakalaŋ kaakɔŋkɔŋɔ̀", "FJT": "Fiji kaakɔŋkɔŋɔ̀ ɖeiɖei", "FJT Summer": "Fiji kaakɔŋkɔŋɔ̀ gafʊbaka", "FKST": "Fɔklanɖ kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀ gafʊbaka", "FKT": "Fɔklanɖ kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀ ɖeiɖei", "FNST": "Fɛrnanɖo ɖe Norɔnya kaakɔŋkɔŋɔ̀ gafʊbaka", "FNT": "Fɛrnanɖo ɖe Norɔnya kaakɔŋkɔŋɔ̀ ɖeiɖei", "GALT": "Galapagɔs kaakɔŋkɔŋɔ̀", "GAMT": "Gambiyee kaakɔŋkɔŋɔ̀", "GEST": "Jɔrjiya kaakɔŋkɔŋɔ̀ gafʊbaka", "GET": "Jɔrjiya kaakɔŋkɔŋɔ̀ ɖeiɖei", "GFT": "Guyanaa Gafɔntǝna kaja kaakɔŋkɔŋɔ̀", "GIT": "Jilbɛɛr kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "GMT": "Griinwish kaakɔŋkɔŋɔ̀", "GNSST": "GNSST", "GNST": "GNST", "GST": "Gɔlf kaakɔŋkɔŋɔ̀", "GST Guam": "Guwam kaakɔŋkɔŋɔ̀", "GYT": "Guyanaa kaakɔŋkɔŋɔ̀", "HADT": "Awayɩɩ n’Alewutii kaakɔŋkɔŋɔ̀ ɖeiɖei", "HAST": "Awayɩɩ n’Alewutii kaakɔŋkɔŋɔ̀ ɖeiɖei", "HKST": "Hɔŋ Kɔŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "HKT": "Hɔŋ Kɔŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "HOVST": "Kɔfɖǝ kaakɔŋkɔŋɔ̀ gafʊbaka", "HOVT": "Kɔfɖǝ kaakɔŋkɔŋɔ̀ ɖeiɖei", "ICT": "Inɖicaɩna kaakɔŋkɔŋɔ̀", "IDT": "Yishraɛl kaakɔŋkɔŋɔ̀ gafʊbaka", "IOT": "Inɖiya kateŋku kaakɔŋkɔŋɔ̀", "IRKST": "Irkut kaakɔŋkɔŋɔ̀ gafʊbaka", "IRKT": "Irkut kaakɔŋkɔŋɔ̀ ɖeiɖei", "IRST": "Iraŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "IRST DST": "Iraŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "IST": "Inɖiya kaakɔŋkɔŋɔ̀", "IST Israel": "Yishraɛl kaakɔŋkɔŋɔ̀ ɖeiɖei", "JDT": "Japaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "JST": "Japaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "KOST": "Kɔsrɛɛ kaakɔŋkɔŋɔ̀", "KRAST": "Krasnoyark kaakɔŋkɔŋɔ̀ gafʊbaka", "KRAT": "Krasnoyark kaakɔŋkɔŋɔ̀ ɖeiɖei", "KST": "Koree kaakɔŋkɔŋɔ̀ ɖeiɖei", "KST DST": "Koree kaakɔŋkɔŋɔ̀ gafʊbaka", "LHDT": "Lɔrɖ Hoo kaakɔŋkɔŋɔ̀ gafʊbaka", "LHST": "Lɔrɖ Hoo kaakɔŋkɔŋɔ̀ ɖeiɖei", "LINT": "Laɩn kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "MAGST": "Magaɖan kaakɔŋkɔŋɔ̀ gafʊbaka", "MAGT": "Magaɖan kaakɔŋkɔŋɔ̀ ɖeiɖei", "MART": "Markesas kaakɔŋkɔŋɔ̀", "MAWT": "Mɔsɔn kaakɔŋkɔŋɔ̀", "MDT": "Amalɩka gʊnyɩpɛnɛlaŋ kabʊnʊ kaakɔŋkɔŋɔ̀ gafʊbaka", "MESZ": "Garɩfɔntǝna gɩcɩɩca kaakɔŋkɔŋɔ̀ gafʊbaka", "MEZ": "Garɩfɔntǝna gɩcɩɩca kaakɔŋkɔŋɔ̀ ɖeiɖei", "MHT": "Marshal kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "MMT": "Miyanmaa kaakɔŋkɔŋɔ̀", "MSD": "Moskuu kaakɔŋkɔŋɔ̀ gafʊbaka", "MST": "Amalɩka gʊnyɩpɛnɛlaŋ kabʊnʊ kaakɔŋkɔŋɔ̀ ɖeiɖei", "MUST": "Imoris kaakɔŋkɔŋɔ̀ gafʊbaka", "MUT": "Imoris kaakɔŋkɔŋɔ̀ ɖeiɖei", "MVT": "Malɖiifu kaakɔŋkɔŋɔ̀", "MYT": "Malɛsii kaakɔŋkɔŋɔ̀", "NCT": "Kaleɖonii afɔlɩ kaakɔŋkɔŋɔ̀ ɖeiɖei", "NDT": "Faʊnɖlanɖ afɔlɩ kaakɔŋkɔŋɔ̀ gafʊbaka", "NDT New Caledonia": "Kaleɖonii afɔlɩ kaakɔŋkɔŋɔ̀ gafʊbaka", "NFDT": "Nɔrfook kaAtukǝltǝna kaakɔŋkɔŋɔ̀ gafʊbaka", "NFT": "Nɔrfook kaAtukǝltǝna kaakɔŋkɔŋɔ̀ ɖeiɖei", "NOVST": "Nofosibirk kaakɔŋkɔŋɔ̀ gafʊbaka", "NOVT": "Nofosibirk kaakɔŋkɔŋɔ̀ ɖeiɖei", "NPT": "Neepal kaakɔŋkɔŋɔ̀", "NRT": "Nawuru kaakɔŋkɔŋɔ̀", "NST": "Faʊnɖlanɖ afɔlɩ kaakɔŋkɔŋɔ̀ ɖeiɖei", "NUT": "Niwuye kaakɔŋkɔŋɔ̀", "NZDT": "Selanɖ afɔlɩ kaakɔŋkɔŋɔ̀ gafʊbaka", "NZST": "Selanɖ afɔlɩ kaakɔŋkɔŋɔ̀ ɖeiɖei", "OESZ": "Garɩfɔntǝna gajakalaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "OEZ": "Garɩfɔntǝna gajakalaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "OMSST": "Ɔmsǝkǝ kaakɔŋkɔŋɔ̀ gafʊbaka", "OMST": "Ɔmsǝkǝ kaakɔŋkɔŋɔ̀ ɖeiɖei", "PDT": "Amalɩka kapasifika kaakɔŋkɔŋɔ̀ gafʊbaka", "PDTM": "Mɛsik kapasifika kaakɔŋkɔŋɔ̀ gafʊbaka", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papuasii Ginee afɔlɩ kaakɔŋkɔŋɔ̀", "PHOT": "Foeniis kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "PKT": "Pakistan kaakɔŋkɔŋɔ̀ ɖeiɖei", "PKT DST": "Pakistan kaakɔŋkɔŋɔ̀ gafʊbaka", "PMDT": "Sɛŋ-Petrɔs na Mikelɔŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "PMST": "Sɛŋ-Petrɔs na Mikelɔŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "PONT": "Poonpei kaakɔŋkɔŋɔ̀", "PST": "Amalɩka kapasifika kaakɔŋkɔŋɔ̀ ɖeiɖei", "PST Philippine": "Filipiin kaakɔŋkɔŋɔ̀ ɖeiɖei", "PST Philippine DST": "Filipiin kaakɔŋkɔŋɔ̀ gafʊbaka", "PST Pitcairn": "Pɩtkɛɛn kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "PSTM": "Mɛsik kapasifika kaakɔŋkɔŋɔ̀ ɖeiɖei", "PWT": "Palawoo kaakɔŋkɔŋɔ̀", "PYST": "Paraguwee kaakɔŋkɔŋɔ̀ gafʊbaka", "PYT": "Paraguwee kaakɔŋkɔŋɔ̀ ɖeiɖei", "PYT Korea": "Koree gʊnyɩpɛnɛlaŋ kaakɔŋkɔŋɔ̀", "RET": "Reeniyɔŋ kaakɔŋkɔŋɔ̀", "ROTT": "Roteraa kaakɔŋkɔŋɔ̀", "SAKST": "Sakalin kaakɔŋkɔŋɔ̀ gafʊbaka", "SAKT": "Sakalin kaakɔŋkɔŋɔ̀ ɖeiɖei", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Garɩɖontǝna gʊnyɩsonolaŋ kaakɔŋkɔŋɔ̀", "SBT": "Salomɔɔn kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "SCT": "Seshɛl kaakɔŋkɔŋɔ̀", "SGT": "Siŋgapuur kaakɔŋkɔŋɔ̀", "SLST": "Siri Laŋkaa kaakɔŋkɔŋɔ̀", "SRT": "Surinam kaakɔŋkɔŋɔ̀", "SST Samoa": "Samowa Amalɩka kaja kaakɔŋkɔŋɔ̀ ɖeiɖei", "SST Samoa Apia": "Samowa kaakɔŋkɔŋɔ̀ ɖeiɖei", "SST Samoa Apia DST": "Samowa kaakɔŋkɔŋɔ̀ gafʊbaka", "SST Samoa DST": "Samowa Amalɩka kaja kaakɔŋkɔŋɔ̀ gafʊbaka", "SYOT": "Siyowaa kaakɔŋkɔŋɔ̀", "TAAF": "Gafɔntǝna gʊnyɩsonolaŋ na Gatutaltǝna kaakɔŋkɔŋɔ̀", "TAHT": "Tahitii kaakɔŋkɔŋɔ̀", "TJT": "Tajikistan kaakɔŋkɔŋɔ̀", "TKT": "Tokelaʊ kaakɔŋkɔŋɔ̀", "TLT": "Timɔɔ gajakalaŋ kaakɔŋkɔŋɔ̀", "TMST": "Turkmenistan kaakɔŋkɔŋɔ̀ gafʊbaka", "TMT": "Turkmenistan kaakɔŋkɔŋɔ̀ ɖeiɖei", "TOST": "Tɔŋga kaakɔŋkɔŋɔ̀ gafʊbaka", "TOT": "Tɔŋga kaakɔŋkɔŋɔ̀ ɖeiɖei", "TVT": "Tufalu kaakɔŋkɔŋɔ̀", "TWT": "Taɩwan kaakɔŋkɔŋɔ̀ ɖeiɖei", "TWT DST": "Taɩwan kaakɔŋkɔŋɔ̀ gafʊbaka", "ULAST": "Ulanbatɔɔr kaakɔŋkɔŋɔ̀ gafʊbaka", "ULAT": "Ulanbatɔɔr kaakɔŋkɔŋɔ̀ ɖeiɖei", "UYST": "Uruguwee kaakɔŋkɔŋɔ̀ gafʊbaka", "UYT": "Uruguwee kaakɔŋkɔŋɔ̀ ɖeiɖei", "UZT": "Usbeekistan kaakɔŋkɔŋɔ̀ ɖeiɖei", "UZT DST": "Usbeekistan kaakɔŋkɔŋɔ̀ gafʊbaka", "VET": "Fenesuwelaa kaakɔŋkɔŋɔ̀", "VLAST": "Flaɖifɔstɔk kaakɔŋkɔŋɔ̀ gafʊbaka", "VLAT": "Flaɖifɔstɔk kaakɔŋkɔŋɔ̀ ɖeiɖei", "VOLST": "Fɔlgograaɖ kaakɔŋkɔŋɔ̀ gafʊbaka", "VOLT": "Fɔlgograaɖ kaakɔŋkɔŋɔ̀ ɖeiɖei", "VOST": "Fɔstɔk kaakɔŋkɔŋɔ̀", "VUT": "Fanuwatu kaakɔŋkɔŋɔ̀ ɖeiɖei", "VUT DST": "Fanuwatu kaakɔŋkɔŋɔ̀ gafʊbaka", "WAKT": "Week kaBʊtukǝltǝna kaakɔŋkɔŋɔ̀", "WARST": "Arjantin gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "WART": "Arjantin gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "WAST": "Garɩɖontǝna gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀", "WAT": "Garɩɖontǝna gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀", "WESZ": "Garɩfɔntǝna gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "WEZ": "Garɩfɔntǝna gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "WFT": "Walis na Futuna kaakɔŋkɔŋɔ̀", "WGST": "Grinlanɖ gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ gafʊbaka", "WGT": "Grinlanɖ gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀ ɖeiɖei", "WIB": "Ɛnɖonosii kagɩteŋshilelaŋ kaakɔŋkɔŋɔ̀", "WIT": "Ɛnɖonosii kaajakalaŋ kaakɔŋkɔŋɔ̀", "WITA": "Ɛnɖonosii kagɩcɩɩca kaakɔŋkɔŋɔ̀", "YAKST": "Yakut kaakɔŋkɔŋɔ̀ gafʊbaka", "YAKT": "Yakut kaakɔŋkɔŋɔ̀ ɖeiɖei", "YEKST": "Yekaterinbuu kaakɔŋkɔŋɔ̀ gafʊbaka", "YEKT": "Yekaterinbuu kaakɔŋkɔŋɔ̀ ɖeiɖei", "YST": "Yukɔn kaakɔŋkɔŋɔ̀", "МСК": "Moskuu kaakɔŋkɔŋɔ̀ ɖeiɖei", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Kasastan gɩteŋshilelaŋ kaakɔŋkɔŋɔ̀", "شىعىش قازاق ەلى": "Kasastan gajakalaŋ kaakɔŋkɔŋɔ̀", "قازاق ەلى": "Kasastan kaakɔŋkɔŋɔ̀", "قىرعىزستان": "Kirgistan kaakɔŋkɔŋɔ̀", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Asɔɔr kaakɔŋkɔŋɔ̀ gafʊbaka"},
	}
}

// Locale returns the current translators string locale
func (blo *blo_BJ) Locale() string {
	return blo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'blo_BJ'
func (blo *blo_BJ) PluralsCardinal() []locales.PluralRule {
	return blo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'blo_BJ'
func (blo *blo_BJ) PluralsOrdinal() []locales.PluralRule {
	return blo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'blo_BJ'
func (blo *blo_BJ) PluralsRange() []locales.PluralRule {
	return blo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'blo_BJ'
func (blo *blo_BJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'blo_BJ'
func (blo *blo_BJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 0 {
		return locales.PluralRuleZero
	} else if i == 1 {
		return locales.PluralRuleOne
	} else if i == 2 || i == 3 || i == 4 || i == 5 || i == 6 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'blo_BJ'
func (blo *blo_BJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (blo *blo_BJ) MonthAbbreviated(month time.Month) string {
	return blo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (blo *blo_BJ) MonthsAbbreviated() []string {
	return blo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (blo *blo_BJ) MonthNarrow(month time.Month) string {
	return blo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (blo *blo_BJ) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (blo *blo_BJ) MonthWide(month time.Month) string {
	return blo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (blo *blo_BJ) MonthsWide() []string {
	return blo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (blo *blo_BJ) WeekdayAbbreviated(weekday time.Weekday) string {
	return blo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (blo *blo_BJ) WeekdaysAbbreviated() []string {
	return blo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (blo *blo_BJ) WeekdayNarrow(weekday time.Weekday) string {
	return blo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (blo *blo_BJ) WeekdaysNarrow() []string {
	return blo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (blo *blo_BJ) WeekdayShort(weekday time.Weekday) string {
	return blo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (blo *blo_BJ) WeekdaysShort() []string {
	return blo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (blo *blo_BJ) WeekdayWide(weekday time.Weekday) string {
	return blo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (blo *blo_BJ) WeekdaysWide() []string {
	return blo.daysWide
}

// Decimal returns the decimal point of number
func (blo *blo_BJ) Decimal() string {
	return blo.decimal
}

// Group returns the group of number
func (blo *blo_BJ) Group() string {
	return blo.group
}

// Group returns the minus sign of number
func (blo *blo_BJ) Minus() string {
	return blo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'blo_BJ' and handles both Whole and Real numbers based on 'v'
func (blo *blo_BJ) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, blo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(blo.group) - 1; j >= 0; j-- {
					b = append(b, blo.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, blo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'blo_BJ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (blo *blo_BJ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 2*len(s[:len(s)-int(v)-1])/2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, blo.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 2 {
				for j := len(blo.group) - 1; j >= 0; j-- {
					b = append(b, blo.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, blo.minus[0])
	}

	for j := len(blo.percentPrefix) - 1; j >= 0; j-- {
		b = append(b, blo.percentPrefix[j])
	}

	b = append(b, blo.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'blo_BJ'
func (blo *blo_BJ) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := blo.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, blo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(blo.group) - 1; j >= 0; j-- {
					b = append(b, blo.group[j])
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

	for j := len(blo.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, blo.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, blo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, blo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'blo_BJ'
// in accounting notation.
func (blo *blo_BJ) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := blo.currencies[currency]
	l := len(s) + len(symbol) + 5 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, blo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(blo.group) - 1; j >= 0; j-- {
					b = append(b, blo.group[j])
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

		for j := len(blo.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, blo.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(blo.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, blo.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, blo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, blo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, blo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, blo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, blo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'blo_BJ'
func (blo *blo_BJ) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, blo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := blo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
