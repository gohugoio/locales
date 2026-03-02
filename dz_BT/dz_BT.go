package dz_BT

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type dz_BT struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	percentSuffix      string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'dz_BT' locale
func New() locales.Translator {
	return &dz_BT{
		locale:             "dz_BT",
		pluralsCardinal:    []locales.PluralRule{6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		inifinity:          "གྲངས་མེད",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:      " ",
		monthsAbbreviated:  []string{"", "༡", "༢", "༣", "༤", "༥", "༦", "༧", "༨", "༩", "༡༠", "༡༡", "12"},
		monthsNarrow:       []string{"", "༡", "༢", "༣", "4", "༥", "༦", "༧", "༨", "9", "༡༠", "༡༡", "༡༢"},
		monthsWide:         []string{"", "ཟླ་དངཔ་", "ཟླ་གཉིས་པ་", "ཟླ་གསུམ་པ་", "ཟླ་བཞི་པ་", "ཟླ་ལྔ་པ་", "ཟླ་དྲུག་པ", "ཟླ་བདུན་པ་", "ཟླ་བརྒྱད་པ་", "ཟླ་དགུ་པ་", "ཟླ་བཅུ་པ་", "ཟླ་བཅུ་གཅིག་པ་", "ཟླ་བཅུ་གཉིས་པ་"},
		daysAbbreviated:    []string{"ཟླ་", "མིར་", "ལྷག་", "ཕུར་", "སངས་", "སྤེན་", "ཉི་"},
		daysNarrow:         []string{"ཟླ", "མིར", "ལྷག", "ཕུར", "སངྶ", "སྤེན", "ཉི"},
		daysWide:           []string{"གཟའ་ཟླ་བ་", "གཟའ་མིག་དམར་", "གཟའ་ལྷག་པ་", "གཟའ་ཕུར་བུ་", "གཟའ་པ་སངས་", "གཟའ་སྤེན་པ་", "གཟའ་ཉི་མ་"},
		periodsAbbreviated: []string{"སྔ་ཆ་", "ཕྱི་ཆ་"},
		timezones:          map[string]string{"ACDT": "དབུས་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཉིན་སྲུང་ཆུ་ཚོད", "ACST": "ACST", "ACT": "ACT", "ACWDT": "དབུས་ནུབ་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཉིན་སྲུང་ཆུ་ཚོད", "ACWST": "དབུས་ནུབ་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "ADT": "ཨེཊ་ལེན་ཊིཀ་ཉིན་སྲུང་ཆུ་ཚོད", "ADT Arabia": "ཨ་རེ་བྷི་ཡན་སྲུང་ཆུ་ཚོད", "AEDT": "ཤར་ཕྱོགས་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཉིན་སྲུང་ཆུ་ཚོད", "AEST": "ཤར་ཕྱོགས་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "AFT": "ཨཕ་ག་ནི་ས྄ཏཱནཆུ་ཚོད", "AKDT": "ཨ་ལསི་ཀ་ཉིན་སྲུང་ཆུ་ཚོད", "AKST": "ཨ་ལསི་ཀ་ཚད་ལྡན་ཆུ་ཚོད", "AMST": "ཨེ་མ་ཛཱོན་བྱཱར་དུས་ཆུ་ཚོད", "AMST Armenia": "ཨར་མི་ནི་ཡ་བྱཱར་དུས་ཆུ་ཚོད", "AMT": "ཨེ་མ་ཛཱོན་ཚད་ལྡན་ཆུ་ཚོད", "AMT Armenia": "ཨར་མི་ནི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ཨར་ཇེན་ཊི་ན་བྱཱར་དུས་ཆུ་ཚོད", "ART": "ཨར་ཇེན་ཊི་ན་ཚད་ལྡན་ཆུ་ཚོད", "AST": "ཨེཊ་ལེན་ཊིཀ་ཚད་ལྡན་ཆུ་ཚོད", "AST Arabia": "ཨ་རེ་བྷི་ཡན་ཚད་ལྡན་ཆུ་ཚོད", "AWDT": "ནུབ་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཉིན་སྲུང་ཆུ་ཚོད", "AWST": "ནུབ་ཕྱོགས་ཨཱོས་ཊྲེལ་ལི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "AZST": "ཨ་ཛར་བྷའི་ཇཱན་བྱཱར་དུས་ཆུ་ཚོད", "AZT": "ཨ་ཛར་བྷའི་ཇཱན་ཚད་ལྡན་ཆུ་ཚོད", "BDT Bangladesh": "བངྒ་ལ་དེཤ་བྱཱར་དུས་ཆུ་ཚོད", "BNT": "BNT", "BOT": "བྷོ་ལི་བི་ཡ་ཆུ་ཚོད", "BRST": "བྲ་ཛི་ལི་ཡ་བྱཱར་དུས་ཆུ་ཚོད", "BRT": "བྲ་ཛི་ལི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "BST Bangladesh": "བངྒ་ལ་དེཤ་ཚད་ལྡན་ཆུ་ཚོད", "BT": "འབྲུག་ཡུལ་ཆུ་ཚོད", "CAST": "CAST", "CAT": "དབུས་ཕྱོགས་ཨཕ་རི་ཀཱ་ཆུ་ཚོད", "CCT": "CCT", "CDT": "བྱང་ཨ་མི་རི་ཀ་དབུས་ཕྱོགས་ཉིན་སྲུང་ཆུ་ཚོད", "CHADT": "CHADT", "CHAST": "CHAST", "CHUT": "CHUT", "CKT": "CKT", "CKT DST": "CKT DST", "CLST": "ཅི་ལི་བྱཱར་དུས་ཆུ་ཚོད", "CLT": "ཅི་ལི་ཚད་ལྡན་ཆུ་ཚོད", "COST": "ཀོ་ལོམ་བྷི་ཡ་བྱཱར་དུས་ཆུ་ཚོད", "COT": "ཀོ་ལོམ་བྷི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "CST": "བྱང་ཨ་མི་རི་ཀ་དབུས་ཕྱོགས་ཚད་ལྡན་ཆུ་ཚོད", "CST China": "རྒྱ་ནག་ཚད་ལྡན་ཆུ་ཚོད", "CST China DST": "རྒྱ་ནག་ཉིན་སྲུང་ཆུ་ཚོད", "CVST": "ཀེཔ་བཱཌ་བྱཱར་དུས་ཆུ་ཚོད", "CVT": "ཀེཔ་བཱཌ་ཚད་ལྡན་ཆུ་ཚོད", "CXT": "ཁི་རིསྟ་མེས་མཚོ་གླིང་ཆུ་ཚོད", "ChST": "ChST", "ChST NMI": "ChST NMI", "CuDT": "ཀིའུ་བྷ་ཉིན་སྲུང་ཆུ་ཚོད", "CuST": "ཀིའུ་བྷ་ཚད་ལྡན་ཆུ་ཚོད", "DAVT": "DAVT", "DDUT": "DDUT", "EASST": "ཨིསི་ཊར་ཨཱའི་ལེནཌ་བྱཱར་དུས་ཆུ་ཚོད", "EAST": "ཨིསི་ཊར་ཨཱའི་ལེནཌ་ཚད་ལྡན་ཆུ་ཚོད", "EAT": "ཤར་ཕྱོགས་ཨཕ་རི་ཀཱ་ཆུ་ཚོད", "ECT": "ཨེ་ཀུ་ཌཽ་ཆུ་ཚོད", "EDT": "བྱང་ཨ་མི་རི་ཀ་ཤར་ཕྱོགས་ཉིན་སྲུང་ཆུ་ཚོད", "EGDT": "ཤར་ཕྱོགས་གིརིན་ལེནཌ་བྱཱར་དུས་ཆུ་ཚོད", "EGST": "ཤར་ཕྱོགས་གིརིན་ལེནཌ་ཚད་ལྡན་ཆུ་ཚོད", "EST": "བྱང་ཨ་མི་རི་ཀ་ཤར་ཕྱོགས་ཚད་ལྡན་ཆུ་ཚོད", "FEET": "FEET", "FJT": "FJT", "FJT Summer": "FJT Summer", "FKST": "ཕལཀ་ལེནཌ་ཨཱའི་ལེནཌས་བྱཱར་དུས་ཆུ་ཚོད", "FKT": "ཕལཀ་ལེནཌ་ཨཱའི་ལེནཌས་ཚད་ལྡན་ཆུ་ཚོད", "FNST": "ཕར་ནེན་ཌོ་ ཌི་ ནོ་རཱོན་ཧ་བྱཱར་དུས་ཆུ་ཚོད", "FNT": "ཕར་ནེན་ཌོ་ ཌི་ ནོ་རཱོན་ཧ་ཚད་ལྡན་ཆུ་ཚོད", "GALT": "ག་ལ་པ་གོསི་ཆུ་ཚོད", "GAMT": "GAMT", "GEST": "ཇཽ་ཇཱ་བྱཱར་དུས་ཆུ་ཚོད", "GET": "ཇཽ་ཇཱ་ཚད་ལྡན་ཆུ་ཚོད", "GFT": "ཕིརེནཅ་གི་ཡ་ན་ཆུ་ཚོད", "GIT": "GIT", "GMT": "གིརིན་ཝིཆ་ལུ་ཡོད་པའི་ཆུ་ཚོད", "GNSST": "GNSST", "GNST": "GNST", "GST": "GST", "GST Guam": "GST Guam", "GYT": "གུ་ཡ་ན་ཆུ་ཚོད", "HADT": "ཧ་ཝའི་-ཨེ་ལིའུ་ཤེན་ཉིན་སྲུང་ཆུ་ཚོད", "HAST": "ཧ་ཝའི་-ཨེ་ལིའུ་ཤེན་ཚད་ལྡན་ཆུ་ཚོད", "HKST": "HKST", "HKT": "HKT", "HOVST": "HOVST", "HOVT": "HOVT", "ICT": "ཨིན་ཌོ་ཅཱའི་ན་ཆུ་ཚོད", "IDT": "ཨིས་རེལ་ཉིན་སྲུང་ཆུ་ཚོད", "IOT": "རྒྱ་གར་གྱི་རྒྱ་མཚོ་ཆུ་ཚོད", "IRKST": "ཨར་ཀུཙི་བྱཱར་དུས་ཆུ་ཚོད", "IRKT": "ཨར་ཀུཙི་ཚད་ལྡན་ཆུ་ཚོད", "IRST": "ཨི་རཱན་ཚད་ལྡན་ཆུ་ཚོད", "IRST DST": "ཨི་རཱན་ཉིན་སྲུང་ཆུ་ཚོད", "IST": "རྒྱ་གར་ཆུ་ཚོད", "IST Israel": "ཨིས་རེལ་ཚད་ལྡན་ཆུ་ཚོད", "JDT": "ཇ་པཱན་ཉིན་སྲུང་ཆུ་ཚོད", "JST": "ཇ་པཱན་ཚད་ལྡན་ཆུ་ཚོད", "KOST": "KOST", "KRAST": "ཀརསི་ནོ་ཡརསཀི་བྱཱར་དུས་ཆུ་ཚོད", "KRAT": "ཀརསི་ནོ་ཡརསཀི་ཚད་ལྡན་ཆུ་ཚོད", "KST": "ཀོ་རི་ཡ་ཚད་ལྡན་ཆུ་ཚོད", "KST DST": "ཀོ་རི་ཡ་ཉིན་སྲུང་ཆུ་ཚོད", "LHDT": "LHDT", "LHST": "LHST", "LINT": "LINT", "MAGST": "མ་གྷ་དཱན་བྱཱར་དུས་ཆུ་ཚོད", "MAGT": "མ་གྷ་དཱན་ཚད་ལྡན་ཆུ་ཚོད", "MART": "MART", "MAWT": "MAWT", "MDT": "བྱང་ཨ་མི་རི་ཀ་མའུ་ཊེན་ཉིན་སྲུང་ཆུ་ཚོད", "MESZ": "དབུས་ཕྱོགས་ཡུ་རོ་པེན་བྱཱར་དུས་ཆུ་ཚོད", "MEZ": "དབུས་ཕྱོགས་ཡུ་རོ་པེན་ཚད་ལྡན་ཆུ་ཚོད", "MHT": "MHT", "MMT": "MMT", "MSD": "མཽས་ཀོ་བྱཱར་དུས་ཆུ་ཚོད", "MST": "བྱང་ཨ་མི་རི་ཀ་མའུ་ཊེན་ཚད་ལྡན་ཆུ་ཚོད", "MUST": "མོ་རི་ཤཱས་བྱཱར་དུས་ཆུ་ཚོད", "MUT": "མོ་རི་ཤཱས་ཚད་ལྡན་ཆུ་ཚོད", "MVT": "མཱལ་དིབས་ཆུ་ཚོད", "MYT": "MYT", "NCT": "NCT", "NDT": "ནིའུ་ཕའུནཌ་ལེནཌ་ཉིན་སྲུང་ཆུ་ཚོད", "NDT New Caledonia": "NDT New Caledonia", "NFDT": "NFDT", "NFT": "NFT", "NOVST": "ནོ་བོ་སི་བིརསཀི་བྱཱར་དུས་ཆུ་ཚོད", "NOVT": "ནོ་བོ་སི་བིརསཀི་ཚད་ལྡན་ཆུ་ཚོད", "NPT": "ནེ་པཱལ་ཆུ་ཚོད", "NRT": "NRT", "NST": "ནིའུ་ཕའུནཌ་ལེནཌ་ཚད་ལྡན་ཆུ་ཚོད", "NUT": "NUT", "NZDT": "ནིའུ་ཛི་ལེནཌ་ཉིན་སྲུང་ཆུ་ཚོད", "NZST": "ནིའུ་ཛི་ལེནཌ་ཚད་ལྡན་ཆུ་ཚོད", "OESZ": "ཤར་ཕྱོགས་ཡུ་རོ་པེན་བྱཱར་དུས་ཆུ་ཚོད", "OEZ": "ཤར་ཕྱོགས་ཡུ་རོ་པེན་ཚད་ལྡན་ཆུ་ཚོད", "OMSST": "ཨོམསཀི་བྱཱར་དུས་ཆུ་ཚོད", "OMST": "ཨོམསཀི་ཚད་ལྡན་ཆུ་ཚོད", "PDT": "བྱང་ཨ་མི་རི་ཀ་པེ་སི་ཕིག་ཉིན་སྲུང་ཆུ་ཚོད", "PDTM": "PDTM", "PETDT": "PETDT", "PETST": "PETST", "PGT": "PGT", "PHOT": "PHOT", "PKT": "པ་ཀི་ས྄ཏཱན་ཚད་ལྡན་ཆུ་ཚོད", "PKT DST": "པ་ཀི་ས྄ཏཱན་བྱཱར་དུས་ཆུ་ཚོད", "PMDT": "པའི་རི་དང་མི་ཀི་ལཱོན་ཉིན་སྲུང་ཆུ་ཚོད", "PMST": "པའི་རི་དང་མི་ཀི་ལཱོན་ཚད་ལྡན་ཆུ་ཚོད", "PONT": "PONT", "PST": "བྱང་ཨ་མི་རི་ཀ་པེ་སི་ཕིག་ཚད་ལྡན་ཆུ་ཚོད", "PST Philippine": "PST Philippine", "PST Philippine DST": "PST Philippine DST", "PST Pitcairn": "PST Pitcairn", "PSTM": "PSTM", "PWT": "PWT", "PYST": "པ་ར་གུ་ཝའི་བྱཱར་དུས་ཆུ་ཚོད", "PYT": "པ་ར་གུ་ཝའི་ཚད་ལྡན་ཆུ་ཚོད", "PYT Korea": "PYT Korea", "RET": "རི་ཡུ་ནི་ཡཱན་ཆུ་ཚོད", "ROTT": "ROTT", "SAKST": "ས་ཁ་ལིན་བྱཱར་དུས་ཆུ་ཚོད", "SAKT": "ས་ཁ་ལིན་ཚད་ལྡན་ཆུ་ཚོད", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ལྷོ་ཕྱོགས་ཨཕ་རི་ཀཱ་ཆུ་ཚོད", "SBT": "SBT", "SCT": "སེ་ཤཱལས་ཆུ་ཚོད", "SGT": "SGT", "SLST": "SLST", "SRT": "སུ་རི་ནཱམ་ཆུ་ཚོད", "SST Samoa": "SST Samoa", "SST Samoa Apia": "SST Samoa Apia", "SST Samoa Apia DST": "SST Samoa Apia DST", "SST Samoa DST": "SST Samoa DST", "SYOT": "SYOT", "TAAF": "TAAF", "TAHT": "TAHT", "TJT": "TJT", "TKT": "TKT", "TLT": "TLT", "TMST": "TMST", "TMT": "TMT", "TOST": "TOST", "TOT": "TOT", "TVT": "TVT", "TWT": "TWT", "TWT DST": "TWT DST", "ULAST": "ULAST", "ULAT": "ULAT", "UYST": "ཡུ་རུ་གུ་ཝཱའི་བྱཱར་དུས་ཆུ་ཚོད", "UYT": "ཡུ་རུ་གུ་ཝཱའི་ཚད་ལྡན་ཆུ་ཚོད", "UZT": "UZT", "UZT DST": "UZT DST", "VET": "བེ་ནི་ཛུ་ཝེ་ལ་ཆུ་ཚོད", "VLAST": "བ་ལ་ཌི་བོསི་ཏོཀ་བྱཱར་དུས་ཆུ་ཚོད", "VLAT": "བ་ལ་ཌི་བོསི་ཏོཀ་ཚད་ལྡན་ཆུ་ཚོད", "VOLST": "བཱོལ་གོ་གིརེཌ་བྱཱར་དུས་ཆུ་ཚོད", "VOLT": "བཱོལ་གོ་གིརེཌ་ཚད་ལྡན་ཆུ་ཚོད", "VOST": "VOST", "VUT": "VUT", "VUT DST": "VUT DST", "WAKT": "WAKT", "WARST": "ནུབ་ཕྱོགས་ཨར་ཇེན་ཊི་ན་བྱཱར་དུས་ཆུ་ཚོད", "WART": "ནུབ་ཕྱོགས་ཨར་ཇེན་ཊི་ན་ཚད་ལྡན་ཆུ་ཚོད", "WAST": "ནུབ་ཕྱོགས་ཨཕ་རི་ཀཱ་ཆུ་ཚོད", "WAT": "ནུབ་ཕྱོགས་ཨཕ་རི་ཀཱ་ཆུ་ཚོད", "WESZ": "ནུབ་ཕྱོགས་ཡུ་རོ་པེན་བྱཱར་དུས་ཆུ་ཚོད", "WEZ": "ནུབ་ཕྱོགས་ཡུ་རོ་པེན་ཚད་ལྡན་ཆུ་ཚོད", "WFT": "WFT", "WGST": "ནུབ་ཕྱོགས་གིརིན་ལེནཌ་བྱཱར་དུས་ཆུ་ཚོད", "WGT": "ནུབ་ཕྱོགས་གིརིན་ལེནཌ་ཚད་ལྡན་ཆུ་ཚོད", "WIB": "ནུབ་ཕྱོགས་ཨིན་ཌོ་ནེ་ཤི་ཡ་ཆུ་ཚོད", "WIT": "ཤར་ཕྱོགས་ཨིན་ཌོ་ནེ་ཤི་ཡ་ཆུ་ཚོད", "WITA": "དབུས་ཕྱོགས་ཨིན་ཌོ་ནེ་ཤི་ཡ་ཆུ་ཚོད", "YAKST": "ཡ་ཀུཙིཀི་བྱཱར་དུས་ཆུ་ཚོད", "YAKT": "ཡ་ཀུཙིཀི་ཚད་ལྡན་ཆུ་ཚོད", "YEKST": "ཡེ་ཀ་ཏེ་རིན་བརག་བྱཱར་དུས་ཆུ་ཚོད", "YEKT": "ཡེ་ཀ་ཏེ་རིན་བརག་ཚད་ལྡན་ཆུ་ཚོད", "YST": "YST", "МСК": "མཽས་ཀོ་ཚད་ལྡན་ཆུ་ཚོད", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "باتىس قازاق ەلى", "شىعىش قازاق ەلى": "شىعىش قازاق ەلى", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "قىرعىزستان", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "པ་རུ་བྱཱར་དུས་ཆུ་ཚོད"},
	}
}

// Locale returns the current translators string locale
func (dz *dz_BT) Locale() string {
	return dz.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'dz_BT'
func (dz *dz_BT) PluralsCardinal() []locales.PluralRule {
	return dz.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'dz_BT'
func (dz *dz_BT) PluralsOrdinal() []locales.PluralRule {
	return dz.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'dz_BT'
func (dz *dz_BT) PluralsRange() []locales.PluralRule {
	return dz.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'dz_BT'
func (dz *dz_BT) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'dz_BT'
func (dz *dz_BT) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'dz_BT'
func (dz *dz_BT) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (dz *dz_BT) MonthAbbreviated(month time.Month) string {
	return dz.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (dz *dz_BT) MonthsAbbreviated() []string {
	return dz.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (dz *dz_BT) MonthNarrow(month time.Month) string {
	return dz.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (dz *dz_BT) MonthsNarrow() []string {
	return dz.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (dz *dz_BT) MonthWide(month time.Month) string {
	return dz.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (dz *dz_BT) MonthsWide() []string {
	return dz.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (dz *dz_BT) WeekdayAbbreviated(weekday time.Weekday) string {
	return dz.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (dz *dz_BT) WeekdaysAbbreviated() []string {
	return dz.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (dz *dz_BT) WeekdayNarrow(weekday time.Weekday) string {
	return dz.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (dz *dz_BT) WeekdaysNarrow() []string {
	return dz.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (dz *dz_BT) WeekdayShort(weekday time.Weekday) string {
	return dz.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (dz *dz_BT) WeekdaysShort() []string {
	return dz.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (dz *dz_BT) WeekdayWide(weekday time.Weekday) string {
	return dz.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (dz *dz_BT) WeekdaysWide() []string {
	return dz.daysWide
}

// Decimal returns the decimal point of number
func (dz *dz_BT) Decimal() string {
	return dz.decimal
}

// Group returns the group of number
func (dz *dz_BT) Group() string {
	return dz.group
}

// Group returns the minus sign of number
func (dz *dz_BT) Minus() string {
	return dz.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'dz_BT' and handles both Whole and Real numbers based on 'v'
func (dz *dz_BT) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, dz.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, dz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'dz_BT' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (dz *dz_BT) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dz.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, dz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, dz.percentSuffix...)

	b = append(b, dz.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'dz_BT'
func (dz *dz_BT) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dz.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, dz.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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
		b = append(b, dz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'dz_BT'
// in accounting notation.
func (dz *dz_BT) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dz.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, dz.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		b = append(b, dz.minus[0])

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
			b = append(b, dz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b}...)
	b = append(b, dz.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, dz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, dz.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20, 0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, dz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0x86, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xbc, 0xe0, 0xbd, 0x91, 0xe0, 0xbc, 0x8b, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0x90, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x98, 0xe0, 0xbc, 0x8b, 0x20}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, dz.periodsAbbreviated[0]...)
	} else {
		b = append(b, dz.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0x86, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xbc, 0xe0, 0xbd, 0x91, 0xe0, 0xbc, 0x8b}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, dz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, dz.periodsAbbreviated[0]...)
	} else {
		b = append(b, dz.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0x86, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xbc, 0xe0, 0xbd, 0x91, 0xe0, 0xbc, 0x8b, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0x90, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x98, 0xe0, 0xbc, 0x8b, 0x20}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, dz.periodsAbbreviated[0]...)
	} else {
		b = append(b, dz.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'dz_BT'
func (dz *dz_BT) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0x86, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xbc, 0xe0, 0xbd, 0x91, 0xe0, 0xbc, 0x8b, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0x90, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x98, 0xe0, 0xbc, 0x8b, 0x20}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, dz.periodsAbbreviated[0]...)
	} else {
		b = append(b, dz.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := dz.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
