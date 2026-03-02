package kgp_BR

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kgp_BR struct {
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

// New returns a new instance of translator for the 'kgp_BR' locale
func New() locales.Translator {
	return &kgp_BR{
		locale:                 "kgp_BR",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: " mil",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: " mil",
		monthsAbbreviated:      []string{"", "1Ky.", "2Ky.", "3Ky.", "4Ky.", "5Ky.", "6Ky.", "7Ky.", "8Ky.", "9Ky.", "10Ky.", "11Ky.", "12Ky."},
		monthsNarrow:           []string{"", "1K", "2K", "3K", "4K", "5K", "6K", "7K", "8K", "9K", "10K", "11K", "12K"},
		monthsWide:             []string{"", "1-Kysã", "2-Kysã", "3-Kysã", "4-Kysã", "5-Kysã", "6-Kysã", "7-Kysã", "8-Kysã", "9-Kysã", "10-Kysã", "11-Kysã", "12-Kysã"},
		daysAbbreviated:        []string{"num.", "pir.", "rég.", "tẽg.", "vẽn.", "pén.", "sav."},
		daysNarrow:             []string{"N.", "P.", "R.", "T.", "V.", "P.", "S."},
		daysShort:              []string{"N.", "1kh.", "2kh.", "3kh.", "4kh.", "5kh.", "S."},
		daysWide:               []string{"numĩggu", "pir-kurã-há", "régre-kurã-há", "tẽgtũ-kurã-há", "vẽnhkãgra-kurã-há", "pénkar-kurã-há", "savnu"},
		periodsAbbreviated:     []string{"", ""},
		erasAbbreviated:        []string{"C.j.", "C.kk."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Cristo jo", "Cristo kar kỹ"},
		timezones:              map[string]string{"ACDT": "Rỹ Kã óra Avotyraria Kuju tá", "ACST": "Óra Pã Avotyraria Kuju tá", "ACT": "Óra Pã Akre tá", "ACWDT": "Rỹ Kã óra Avotyraria Kuju-Rãpur tá", "ACWST": "Óra Pã Avotyraria Kuju-Rãpur tá", "ADT": "Rỹ Kã óra Atrỹtiku tá", "ADT Arabia": "Rỹ Kã óra Aramija tá", "AEDT": "Rỹ Kã óra Avotyraria Rãjur tá", "AEST": "Óra Pã Avotyraria Rãjur tá", "AFT": "Afeganĩtã tá óra", "AKDT": "Rỹ Kã óra Aranhka tá", "AKST": "Óra Pã Aranhka tá", "AMST": "Rỹ Kã óra Amỹjonỹ tá", "AMST Armenia": "Rỹ Kã óra Armẽnĩja tá", "AMT": "Óra Pã Amỹjonỹ tá", "AMT Armenia": "Óra Pã Armẽnĩja tá", "ANAST": "Rỹ Kã óra Anỹnhyr tá", "ANAT": "Óra Pã Anỹnhyr tá", "ARST": "Rỹ Kã óra Arjẽtĩnỹ tá", "ART": "Óra Opã Arjẽtĩnỹ tá", "AST": "Óra Pã Atrỹtiku tá", "AST Arabia": "Óra Pã Aramija tá", "AWDT": "Rỹ Kã óra Avotyraria Rãpur tá", "AWST": "Óra Pã Avotyraria Rãpur tá", "AZST": "Rỹ Kã óra Ajermajjáv tá", "AZT": "Óra Pã Ajermajjáv tá", "BDT Bangladesh": "Rỹ Kã óra Mỹngranési tá", "BNT": "Óra Mrunẽj Narusarỹ tá", "BOT": "Óra Morivia tá", "BRST": "Rỹ Kã óra Mrasirja tá", "BRT": "Óra Pã Mrasirja tá", "BST Bangladesh": "Óra Pã Mỹngranési tá", "BT": "Óra Mutỹv tá", "CAST": "CAST", "CAT": "Afrika-Kuju tá óra", "CCT": "Óra Kokonh Goj-vẽso tá", "CDT": "Rỹ Kã óra Kuju tá", "CHADT": "Rỹ Kã óra San-hỹm tá", "CHAST": "Óra Pã San-hỹm tá", "CHUT": "Óra Suuki tá", "CKT": "Óra Pã Kuki Goj-vẽso tá", "CKT DST": "Rỹ Kã óra Kuki Goj-vẽso tá", "CLST": "Rỹ Kã óra Sire tá", "CLT": "Óra Pã Sire tá", "COST": "Rỹ Kã óra Korãmija tá", "COT": "Óra Pã Korãmija tá", "CST": "Óra Pã Kuju tá", "CST China": "Óra Pã Sĩnỹ tá", "CST China DST": "Rỹ Kã óra Sĩnỹ tá", "CVST": "Rỹ Kã óra Pu tánh tá", "CVT": "Óra Pã Pu Tánh tá", "CXT": "Óra Krĩtimỹnh Goj-vẽso tá", "ChST": "Óra Samãho tá", "ChST NMI": "Óra Nãrti-Mỹrijỹnỹ Goj-vẽso tá", "CuDT": "Rỹ Kã óra Kuma tá", "CuST": "Óra Pã Kuma tá", "DAVT": "Óra Navinh tá", "DDUT": "Óra Numã-Nurviri tá", "EASST": "Rỹ Kã óra Panhkuva Goj-vẽso tá", "EAST": "Óra Pã Panhkuva Goj-vẽso tá", "EAT": "Afrika Rãjur tá óra", "ECT": "Óra Ekuvanor tá", "EDT": "Rỹ Kã óra Rãjur tá", "EGDT": "Rỹ Kã óra Groẽrỹnija Rãjur tá", "EGST": "Óra Pã Groẽrỹnija Rãjur tá", "EST": "Óra Pã Rãjur tá", "FEET": "Rãjur tỹ Orópa jã há tá óra", "FJT": "Óra Pã Fiji tá", "FJT Summer": "Rỹ Kã óra Fiji tá", "FKST": "Rỹ Kã óra Mỹrvĩnỹ Goj-vẽso tá", "FKT": "Ór Pã Mỹrvĩnỹ Goj-vẽso tá", "FNST": "Rỹ Kã óra Fernỹnu Nãrãja-tá tá", "FNT": "Óra Pã Fernỹnu Nãrãja-tá tá", "GALT": "Óra Gara Pago tá", "GAMT": "Óra Gỹmmijer", "GEST": "Rỹ Kã óra Jeórja tá", "GET": "Óra Pã Jeórja tá", "GFT": "Óra Frỹsa Gijanỹ tá", "GIT": "Óra Jirmértu Goj-vẽso tá", "GMT": "Óra Mirinjỹnũ Grinũvisi tá", "GNSST": "GNSST", "GNST": "GNST", "GST": "Óra Jiórja tỹ Sur tá", "GST Guam": "Óra Pã Guvỹm tá", "GYT": "Óra Gijỹnỹ tá", "HADT": "Rỹ kã óra Hava’i kar Arevta Goj-vẽso tá", "HAST": "Óra Pã Hava’i kar Arevta Goj-vẽso tá", "HKST": "Rỹ Kã óra Hãg Kãg tá", "HKT": "Óra Pã Hãg Kãg tá", "HOVST": "Rỹ kã óra Hóvin tá", "HOVT": "Óra Pã Hóvin tá", "ICT": "Óra Ĩnosĩnỹ tá", "IDT": "Rỹ Kã óra Isihaé tá", "IOT": "Óra Osiỹno Ĩniko tá", "IRKST": "Rỹ Kã óra Irkutinhki tá", "IRKT": "Óra Pã Irkutinhki tá", "IRST": "Óra Pã Irỹ tá", "IRST DST": "Rỹ Kã óra Irỹ tá", "IST": "Óra Pã Ĩnija tá", "IST Israel": "Óra Pã Isihaé tá", "JDT": "Rỹ Kã óra Japã tá", "JST": "Óra Pã Japã tá", "KOST": "Óra de Kosiraje tá", "KRAST": "Rỹ Kã óra Kranhnãjarki tá", "KRAT": "Óra Pã Kranhnãjarki tá", "KST": "Óra Pã Koréja tá", "KST DST": "Rỹ Kã óra Koréja tá", "LHDT": "Rỹ Kã óra Rórni Hove tá", "LHST": "Óra Pã Rórni Hove tá", "LINT": "Óra Vãfe Goj-vẽso tỹ tá", "MAGST": "Rỹ Kã óra Mỹganan tá", "MAGT": "Óra Pã Mỹganan tá", "MART": "Óra Mỹrkeja Fag tá", "MAWT": "Óra Mỹusãn tá", "MDT": "Rỹ Kã óra Krĩ tá", "MESZ": "Rỹ Kã óra Orópa Kuju tá", "MEZ": "Óra Pã Orópa Kuju tá", "MHT": "Óra MỹrSar Goj-vẽso tá", "MMT": "Óra Mĩjỹmỹr tá", "MSD": "Rỹ Kã óra Mãnhkov tá", "MST": "Óra Pã Krĩ tá", "MUST": "Rỹ Kã óra Mãriso tá", "MUT": "Óra Pã Mãriso tá", "MVT": "Óra Goj Vẽso Mỹrniva tá", "MYT": "Óra Mỹraja tá", "NCT": "Óra Pã Karenonĩja Tãg tá", "NDT": "Rỹ Kã óra Ga Tãg tá", "NDT New Caledonia": "Rỹ Kã óra Karenonĩja Tãg tá", "NFDT": "Rỹ Kã óra Nãrforki Goj-vẽso tá", "NFT": "Óra Pã Nãrforki Goj-vẽso tá", "NOVST": "Rỹ Kã óra Pã Simirsiki Tãg tá", "NOVT": "Óra Pã Simirsiki Tãg tá", "NPT": "Óra Nẽpar tá", "NRT": "Óra Nỹvuru tá", "NST": "Óra Pã Ga Tãg tá", "NUT": "Óra Nĩve tá", "NZDT": "Rỹ Kã óra Jerỹnija Tãg tá", "NZST": "Óra Pã Jerỹnija Tãg tá", "OESZ": "Rỹ Kã óra Orópa Rãjur tá", "OEZ": "Óra Pã Orópa Rãjur tá", "OMSST": "Rỹ Kã óra Omĩnhki tá", "OMST": "Óra Pã Omĩnhki tá", "PDT": "Rỹ Kã óra Rãpur tá", "PDTM": "Rỹ Kã óra Mẽsiku Pasifiku tá", "PETDT": "Rỹ Kã óra Petrupaviróvinhki-Kỹmsatinhki", "PETST": "Óra Pã Petrupaviróvinhki-Kỹmsatinhki", "PGT": "Óra Papuva-Ginẽ Tãg tá", "PHOT": "Óra Fẽnĩg Goj-vẽso tá", "PKT": "Óra Pã Pakinhtỹv tá", "PKT DST": "Rỹ Kã óra Pakinhtỹv tá", "PMDT": "Rỹ Kã óra Sỹ Pedro kar Mĩkerỹv tá", "PMST": "Óra Pã Sỹ Pedro kar Mĩkerỹv tá", "PONT": "Óra Ponỹpe tá", "PST": "Óra Pã Rãpur tá", "PST Philippine": "Óra Pã Firipinỹ tá", "PST Philippine DST": "Rỹ Kã Firipinỹ tá", "PST Pitcairn": "Óra Pinkajir Goj-vẽso tá", "PSTM": "Óra Pã Mẽsiku Pasifiku tá", "PWT": "Óra Paravu tá", "PYST": "Rỹ Kã óra Paraguvaj tá", "PYT": "Óra Pã Paraguvaj tá", "PYT Korea": "Óra Piãgiỹng tá", "RET": "Óra Hujáv tá", "ROTT": "Óra Rotera tá", "SAKST": "Rỹ Kã óra Sakarinỹ tá", "SAKT": "Óra Pã Sakarinỹ tá", "SAMST": "Rỹ Kã óra Samỹra tá", "SAMT": "Óra Pã Samỹra tá", "SAST": "Sur-Afrika tá óra", "SBT": "Óra Saromỹv Goj-vẽso tá", "SCT": "Óra Sejserenh tá", "SGT": "Óra Pã Sĩgapura tá", "SLST": "Óra Rỹnka tá", "SRT": "Óra Surinỹmĩ tá", "SST Samoa": "Óra Pã Samãva tá", "SST Samoa Apia": "Óra Pã Apija tá", "SST Samoa Apia DST": "Rỹ Kã óra Apija tá", "SST Samoa DST": "Rỹ Kã óra Samãva tá", "SYOT": "Óra Siova tá", "TAAF": "Óra Frỹsa Ga Sur kar Ỹtartina tá", "TAHT": "Óra Tajti tá", "TJT": "Óra Tajikinhtỹv tá", "TKT": "Óra Tokeravu tá", "TLT": "Óra Tĩmãr-Rãjur tá", "TMST": "Rỹ Kã óra Turkomẽnĩnhtỹv tá", "TMT": "Óra Pã Turkomẽnĩnhtỹv tá", "TOST": "Rỹ Kã óra Tãga tá", "TOT": "Óra Pã Tãga tá", "TVT": "Óra Tuvaru tá", "TWT": "Óra Pã Tajpej tá", "TWT DST": "Rỹ Kã óra Tajpej tá", "ULAST": "Rỹ Kã óra Uran Mator tá", "ULAT": "Óra Pã Uran Mator tá", "UYST": "Rỹ Kã óra Uruguvaj tá", "UYT": "Óra Pã Uruguvaj tá", "UZT": "Óra Pã Unnmekinhtỹv tá", "UZT DST": "Rỹ Kã óra Unhmekinhtỹv tá", "VET": "Óra Venẽjuvéra tá", "VLAST": "Rỹ Kã óra Uranivónhtóki tá", "VLAT": "Óra Pã Uranivónhtóki tá", "VOLST": "Rỹ Kã óra Vorgugrano tá", "VOLT": "Óra Pã Vorgugrano tá", "VOST": "Óra Vonhtóki tá", "VUT": "Óra Pã Vanũvatu tá", "VUT DST": "Rỹ Kã óra Vanũvatu tá", "WAKT": "Óra Vejki Goj-vẽso tá", "WARST": "Rỹ Kã óra Arjẽtĩnỹ Rãpurtá", "WART": "Óra Pã Arjẽtĩnỹ Rãpur tá", "WAST": "Afrika Rãpur tá óra", "WAT": "Afrika Rãpur tá óra", "WESZ": "Rỹ Kã óra Orópa Rãpur tá", "WEZ": "Óra Pã Orópa Rãpur tá", "WFT": "Óra Varinh kar Futunỹ tá", "WGST": "Rỹ Kã óra Groẽrỹnija Rãpur tá", "WGT": "Óra Pã Groẽrỹnija Rãpur tá", "WIB": "Óra Ĩnonẽja Rãpur tá", "WIT": "Óra Ĩnonẽja Rãjur tá", "WITA": "Óra Ĩnonẽja Kuju tá", "YAKST": "Rỹ Kã óra Yjakutinhki tá", "YAKT": "Óra Pã Yjakutinhkii tá", "YEKST": "Rỹ Kã óra Ekaterĩnmurgu tá", "YEKT": "Óra Pã Ekaterĩnmurgu tá", "YST": "YST", "МСК": "Óra Pã Mãnhkov tá", "اقتاۋ": "Óra Pã Agtav tá", "اقتاۋ قالاسى": "Rỹ Kã óra Agtav tá", "اقتوبە": "Óra Pã Agtóme tá", "اقتوبە قالاسى": "Rỹ Kã óra Agtóme tá", "الماتى": "Óra Pã Armaty tá", "الماتى قالاسى": "Rỹ Kã óra Armaty tá", "باتىس قازاق ەلى": "Óra Kajakinhtỹv Rãpur tá", "شىعىش قازاق ەلى": "Óra Kajakinhtỹv Rãjur tá", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "Óra Kirginhtỹv tá", "قىزىلوردا": "Óra Pã Kysyrorna tá", "قىزىلوردا قالاسى": "Rỹ Kã óra Kysyrorna tá", "∅∅∅": "Rỹ Kã óra Piru tá"},
	}
}

// Locale returns the current translators string locale
func (kgp *kgp_BR) Locale() string {
	return kgp.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kgp_BR'
func (kgp *kgp_BR) PluralsCardinal() []locales.PluralRule {
	return kgp.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kgp_BR'
func (kgp *kgp_BR) PluralsOrdinal() []locales.PluralRule {
	return kgp.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kgp_BR'
func (kgp *kgp_BR) PluralsRange() []locales.PluralRule {
	return kgp.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kgp_BR'
func (kgp *kgp_BR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kgp_BR'
func (kgp *kgp_BR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kgp_BR'
func (kgp *kgp_BR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kgp *kgp_BR) MonthAbbreviated(month time.Month) string {
	return kgp.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kgp *kgp_BR) MonthsAbbreviated() []string {
	return kgp.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kgp *kgp_BR) MonthNarrow(month time.Month) string {
	return kgp.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kgp *kgp_BR) MonthsNarrow() []string {
	return kgp.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kgp *kgp_BR) MonthWide(month time.Month) string {
	return kgp.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kgp *kgp_BR) MonthsWide() []string {
	return kgp.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kgp *kgp_BR) WeekdayAbbreviated(weekday time.Weekday) string {
	return kgp.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kgp *kgp_BR) WeekdaysAbbreviated() []string {
	return kgp.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kgp *kgp_BR) WeekdayNarrow(weekday time.Weekday) string {
	return kgp.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kgp *kgp_BR) WeekdaysNarrow() []string {
	return kgp.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kgp *kgp_BR) WeekdayShort(weekday time.Weekday) string {
	return kgp.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kgp *kgp_BR) WeekdaysShort() []string {
	return kgp.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kgp *kgp_BR) WeekdayWide(weekday time.Weekday) string {
	return kgp.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kgp *kgp_BR) WeekdaysWide() []string {
	return kgp.daysWide
}

// Decimal returns the decimal point of number
func (kgp *kgp_BR) Decimal() string {
	return kgp.decimal
}

// Group returns the group of number
func (kgp *kgp_BR) Group() string {
	return kgp.group
}

// Group returns the minus sign of number
func (kgp *kgp_BR) Minus() string {
	return kgp.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kgp_BR' and handles both Whole and Real numbers based on 'v'
func (kgp *kgp_BR) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kgp_BR' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kgp *kgp_BR) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kgp_BR'
func (kgp *kgp_BR) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kgp.currencies[currency]
	l := len(s) + len(symbol) + 8

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kgp.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(kgp.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kgp.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, kgp.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kgp.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kgp_BR'
// in accounting notation.
func (kgp *kgp_BR) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kgp.currencies[currency]
	l := len(s) + len(symbol) + 8

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kgp.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kgp.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kgp.currencyNegativePrefix[j])
		}

		b = append(b, kgp.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kgp.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kgp.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, kgp.currencyNegativeSuffix...)
	} else {

		b = append(b, kgp.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6e, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, kgp.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6e, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, kgp.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, kgp.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6e, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, kgp.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kgp_BR'
func (kgp *kgp_BR) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kgp.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kgp.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
