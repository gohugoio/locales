package chr

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type chr struct {
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

// New returns a new instance of translator for the 'chr' locale
func New() locales.Translator {
	return &chr{
		locale:                 "chr",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "ᎤᏃ", "ᎧᎦ", "ᎠᏅ", "ᎧᏬ", "ᎠᏂ", "ᏕᎭ", "ᎫᏰ", "ᎦᎶ", "ᏚᎵ", "ᏚᏂ", "ᏅᏓ", "ᎥᏍ"},
		monthsNarrow:           []string{"", "Ꭴ", "Ꭷ", "Ꭰ", "Ꭷ", "Ꭰ", "Ꮥ", "Ꭻ", "Ꭶ", "Ꮪ", "Ꮪ", "Ꮕ", "Ꭵ"},
		monthsWide:             []string{"", "ᎤᏃᎸᏔᏅ", "ᎧᎦᎵ", "ᎠᏅᏱ", "ᎧᏬᏂ", "ᎠᏂᏍᎬᏘ", "ᏕᎭᎷᏱ", "ᎫᏰᏉᏂ", "ᎦᎶᏂ", "ᏚᎵᏍᏗ", "ᏚᏂᏅᏗ", "ᏅᏓᏕᏆ", "ᎥᏍᎩᏱ"},
		daysAbbreviated:        []string{"ᏆᏍᎬ", "ᏉᏅᎯ", "ᏔᎵᏁ", "ᏦᎢᏁ", "ᏅᎩᏁ", "ᏧᎾᎩ", "ᏈᏕᎾ"},
		daysNarrow:             []string{"Ꮖ", "Ꮙ", "Ꮤ", "Ꮶ", "Ꮕ", "Ꮷ", "Ꭴ"},
		daysShort:              []string{"ᏍᎬ", "ᏅᎯ", "ᏔᎵ", "ᏦᎢ", "ᏅᎩ", "ᏧᎾ", "ᏕᎾ"},
		daysWide:               []string{"ᎤᎾᏙᏓᏆᏍᎬ", "ᎤᎾᏙᏓᏉᏅᎯ", "ᏔᎵᏁᎢᎦ", "ᏦᎢᏁᎢᎦ", "ᏅᎩᏁᎢᎦ", "ᏧᎾᎩᎶᏍᏗ", "ᎤᎾᏙᏓᏈᏕᎾ"},
		periodsAbbreviated:     []string{"ᏌᎾᎴ", "ᏒᎯᏱᎢ"},
		periodsNarrow:          []string{"Ꮜ", "Ꮢ"},
		periodsWide:            []string{"ᏌᎾᎴ", "ᏒᎯᏱᎢᏗᏢ"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ᏧᏓᎷᎸ ᎤᎷᎯᏍᏗ ᎦᎶᏁᏛ", "ᎠᏃ ᏙᎻᏂ"},
		timezones:              map[string]string{"ACDT": "ᎠᏰᏟ ᎡᎳᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ᎠᏰᏟ ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "ACWST": "ᎠᏰᏟ ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "ADT": "ᏗᎧᎸᎬ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "ADT Arabia": "ᎠᎴᏈᏯ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "AEDT": "ᎡᎳᏗᏜ ᏗᎧᎸᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "AEST": "ᎡᎳᏗᏜ ᏗᎧᎸᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AFT": "ᎠᏫᎨᏂᏍᏖᏂ ᎠᏟᎢᎵᏒ", "AKDT": "ᎠᎳᏍᎦ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "AKST": "ᎠᎳᏍᎦ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AMST": "ᎠᎺᏌᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "AMST Armenia": "ᎠᎵᎻᏂᎠ ᎪᎩ ᎠᏟᎢᎵᏒ", "AMT": "ᎠᎺᏌᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AMT Armenia": "ᎠᎵᎻᏂᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ᎠᏥᏂᏘᏂᎠ ᎪᎩ ᎠᏟᎢᎵᏒ", "ART": "ᎠᏥᏂᏘᏂᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AST": "ᏗᎧᎸᎬ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AST Arabia": "ᎠᎴᏈᏯ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AWDT": "ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "AWST": "ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AZST": "ᎠᏏᎵᏆᏌᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "AZT": "ᎠᏏᎵᏆᏌᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "BDT Bangladesh": "ᏆᏂᎦᎵᏕᏍ ᎪᎩ ᎠᏟᎢᎵᏒ", "BNT": "ᏊᎾᎢ ᏓᎷᏌᎳᎻ ᎠᏟᎢᎵᏒ", "BOT": "ᏉᎵᏫᎠ ᎠᏟᎢᎵᏒ", "BRST": "ᏆᏏᎵᏯ ᎪᎩ ᎠᏟᎢᎵᏒ", "BRT": "ᏆᏏᎵᏯ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "BST Bangladesh": "ᏆᏂᎦᎵᏕᏍ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "BT": "ᏊᏔᏂ ᎠᏟᎢᎵᏒ", "CAST": "CAST", "CAT": "ᎠᏰᏟ ᎬᎿᎨᏍᏛ ᎠᏟᎢᎵᏒ", "CCT": "ᎪᎪᏍ ᏚᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "CDT": "ᎠᏰᏟ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "CHADT": "ᏣᏝᎻ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "CHAST": "ᏣᏝᎻ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CHUT": "ᏧᎩ ᎠᏟᎢᎵᏒ", "CKT": "ᎠᏓᏍᏓᏴᎲᏍᎩ ᏚᎦᏚᏛᎢ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CKT DST": "ᎠᏓᏍᏓᏴᎲᏍᎩ ᏚᎦᏚᏛᎢ ᎠᏰᏟ ᎪᎩ ᎠᏟᎢᎵᏒ", "CLST": "ᏥᎵ ᎪᎩ ᎠᏟᎢᎵᏒ", "CLT": "ᏥᎵ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "COST": "ᎪᎸᎻᏈᎢᎠ ᎪᎩ ᎠᏟᎢᎵᏒ", "COT": "ᎪᎸᎻᏈᎢᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CST": "ᎠᏰᏟ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CST China": "ᏓᎶᏂᎨᏍᏛ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CST China DST": "ᏓᎶᏂᎨᏍᏛ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "CVST": "ᎢᎬᎾᏕᎾ ᎢᏤᏳᏍᏗ ᎪᎩ ᎠᏟᎢᎵᏒ", "CVT": "ᎢᎬᎾᏕᎾ ᎢᏤᏳᏍᏗ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CXT": "ᏓᏂᏍᏓᏲᎯᎲ ᎤᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "ChST": "ᏣᎼᎶ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "ChST NMI": "ChST NMI", "CuDT": "ᎫᏆ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "CuST": "ᎫᏆ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "DAVT": "ᏕᏫᏏ ᎠᏟᎢᎵᏒ", "DDUT": "ᏚᎼᎾᏘ-Ꮧ’ᎤᎵᏫᎵ ᎠᏟᎢᎵᏒ", "EASST": "ᏥᏌᏕᎴᎯᏌᏅ ᎤᎦᏚᏛᎢ ᎪᎩ ᎠᏟᎢᎵᏒ", "EAST": "ᏥᏌᏕᎴᎯᏌᏅ ᎤᎦᏚᏛᎢ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "EAT": "ᏗᎧᎸᎬ ᎬᎿᎨᏍᏛ ᎠᏟᎢᎵᏒ", "ECT": "ᎡᏆᏙᎵ ᎠᏟᎢᎵᏒ", "EDT": "ᏗᎧᎸᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "EGDT": "ᏗᎧᎸᎬ ᎢᏤᏍᏛᏱ ᎪᎩ ᎠᏟᎢᎵᏒ", "EGST": "ᏗᎧᎸᎬ ᎢᏤᏍᏛᏱ ᎠᏟᎶᏍᏗ ᎠᎵᎢᎵᏒ", "EST": "ᏗᎧᎸᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "FEET": "ᏗᎧᎸᎬ ᏳᎳᏈ ᎠᏟᎢᎵᏒ", "FJT": "ᏫᏥ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "FJT Summer": "ᏫᏥ ᎪᎩ ᎠᏟᎢᎵᏒ", "FKST": "ᏩᎩ ᏚᎦᏚᏛᎢ ᎪᎩ ᎠᏟᎢᎵᏒ", "FKT": "ᏩᎩ ᏚᎦᏚᏛᎢ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "FNST": "ᏪᎾᏅᏙ Ꮥ ᏃᎶᎾᎭ ᎪᎩ ᎠᏟᎢᎵᏒ", "FNT": "ᏪᎾᏅᏙ Ꮥ ᏃᎶᎾᎭ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "GALT": "ᎡᏆ ᏓᎦᏏ ᎤᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "GAMT": "ᎦᎻᏇᎵ ᎠᏟᎢᎵᏒ", "GEST": "ᏣᎠᏥᎢ ᎪᎩ ᎠᏟᎢᎵᏒ", "GET": "ᏣᎠᏥᎢ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "GFT": "ᎠᏂᎦᎸ ᏈᏯᎾ ᎠᏟᎢᎵᏒ", "GIT": "ᎩᎵᏇᏘ ᏚᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "GMT": "ᎢᏤ ᎢᏳᏍᏗ ᎠᏟᎢᎵᏒ", "GNSST": "GNSST", "GNST": "GNST", "GST": "ᎡᏉᏄᎸᏗ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "GST Guam": "GST Guam", "GYT": "ᎦᏯᎾ ᎠᏟᎢᎵᏒ", "HADT": "ᎭᏩᏱ-ᎠᎵᏳᏏᎠᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HAST": "ᎭᏩᏱ-ᎠᎵᏳᏏᎠᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HKST": "ᎰᏂᎩ ᎪᏂᎩ ᎪᎩ ᎠᏟᎢᎵᏒ", "HKT": "ᎰᏂᎩ ᎪᏂᎩ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HOVST": "ᎰᏩᏗ ᎪᎩ ᎠᏟᎢᎵᏒ", "HOVT": "ᎰᏩᏗ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "ICT": "ᎢᏂᏙᏓᎶᏂᎨᏍᏛ ᎠᏟᎢᎵᏒ", "IDT": "ᎢᏏᎵᏱ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "IOT": "ᎠᏂᏴᏫᏯ ᎠᎺᏉᎯ ᎠᏟᎢᎵᏒ", "IRKST": "ᎢᎫᏥᎧ ᎪᎩ ᎠᏟᎢᎵᏒ", "IRKT": "ᎢᎫᏥᎧ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "IRST": "ᎢᎳᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "IRST DST": "ᎢᎳᏂ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "IST": "ᎢᏂᏗᎢᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "IST Israel": "ᎢᏏᎵᏱ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "JDT": "ᏣᏩᏂᏏ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "JST": "ᏣᏩᏂᏏ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "KOST": "ᎪᏍᎴ ᎠᏟᎢᎵᏒ", "KRAST": "ᏝᏍᏃᏯᏍᎧ ᎪᎩ ᎠᏟᎢᎵᏒ", "KRAT": "ᏝᏍᏃᏯᏍᎧ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "KST": "ᎪᎵᎠᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "KST DST": "ᎪᎵᎠᏂ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "LHDT": "ᎤᎬᏫᏳᎯ ᎭᏫ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "LHST": "ᎤᎬᏫᏳᎯ ᎭᏫ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "LINT": "ᎠᏍᏓᏅᏅ ᏚᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "MAGST": "ᎹᎦᏓᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "MAGT": "ᎹᎦᏓᏂ ᎠᏟᎢᎵᏒ ᎠᏟᎢᎵᏒ", "MART": "ᎹᎵᎨᏌᏏ ᎠᏟᎢᎵᏒ", "MAWT": "ᎹᏌᏂ ᎠᏟᎢᎵᏒ", "MDT": "MDT", "MESZ": "ᎠᏰᏟ ᏳᎳᏈ ᎪᎩ ᎠᏟᎢᎵᏒ", "MEZ": "ᎠᏰᏟ ᏳᎳᏈ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "MHT": "ᎹᏌᎵ ᏚᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "MMT": "ᎹᏯᎹᎵ ᎠᏟᎢᎵᏒ", "MSD": "ᎹᏍᎦᏫ ᎪᎩ ᎠᏟᎢᎵᏒ", "MST": "MST", "MUST": "ᎼᎵᏏᎥᏍ ᎪᎩ ᎠᏟᎢᎵᏒ", "MUT": "ᎼᎵᏏᎥᏍ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "MVT": "ᎹᎵᏗᏫᏍ ᎠᏟᎢᎵᏒ", "MYT": "ᎹᎴᏏᎢᎠ ᎠᏟᎢᎵᏒ", "NCT": "ᎢᏤ ᎧᎵᏙᏂᎠᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "NDT": "ᎢᏤᎤᏂᏩᏛᏓᎦᏙᎯ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "NDT New Caledonia": "ᎢᏤ ᎧᎵᏙᏂᎠᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "NFDT": "ᏃᎵᏬᎵᎩ ᎤᎦᏚᏛᎢ ᎪᎩ ᎠᏟᎢᎵᏒ", "NFT": "ᏃᎵᏬᎵᎩ ᎤᎦᏚᏛᎢ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "NOVST": "ᏃᏬᏏᏈᏍᎧ ᎪᎩ ᎠᏟᎢᎵᏒ", "NOVT": "ᏃᏬᏏᏈᏍᎧ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "NPT": "ᏁᏆᎵ ᎠᏟᎢᎵᏒ", "NRT": "ᎾᎷ ᎠᏟᎢᎵᏒ", "NST": "ᎢᏤᎤᏂᏩᏛᏓᎦᏙᎯ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "NUT": "ᏂᏳ ᎠᏟᎢᎵᏒ", "NZDT": "ᎢᏤ ᏏᎢᎴᏂᏗ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "NZST": "ᎢᏤ ᏏᎢᎴᏂᏗ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "OESZ": "ᏗᎧᎸᎬ ᏗᏜ ᏳᎳᏈ ᎪᎩ ᎠᏟᎢᎵᏒ", "OEZ": "ᏗᎧᎸᎬ ᏗᏜ ᏳᎳᏈ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "OMSST": "ᎣᎻᏍᎧ ᎪᎩ ᎠᏟᎢᎵᏒ", "OMST": "ᎣᎻᏍᎧ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PDT": "ᏭᏕᎵᎬ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "PDTM": "ᎠᏂᏍᏆᏂ ᏭᏕᎵᎬ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ᏆᏇ ᎢᏤ ᎩᎢᏂ ᎠᏟᎢᎵᏒ", "PHOT": "ᏧᎴᎯᏌᏅᎯ ᏚᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "PKT": "ᏆᎩᏍᏖᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PKT DST": "ᏆᎩᏍᏖᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "PMDT": "ᎤᏓᏅᏘ ᏈᏰ & ᎻᏇᎶᏂ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "PMST": "ᎤᏓᏅᏘ ᏈᏰ & ᎻᏇᎶᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PONT": "ᏉᎾᏇ ᎠᏟᎢᎵᏒ", "PST": "ᏭᏕᎵᎬ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PST Philippine": "ᎠᏂᏈᎵᎩᏃ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PST Philippine DST": "ᎠᏂᏈᎵᎩᏃ ᎪᎩ ᎠᏟᎢᎵᏒ", "PST Pitcairn": "ᏈᎧᎵᏂ ᎠᏟᎢᎵᏒ", "PSTM": "ᎠᏂᏍᏆᏂ ᏭᏕᎵᎬ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PWT": "ᏆᎷ ᎠᏟᎢᎵᏒ", "PYST": "ᏆᎵᏇ ᎪᎩ ᎠᏟᎢᎵᏒ", "PYT": "ᏆᎵᏇ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PYT Korea": "ᏈᏯᏂᎩᏰᏂᎩ ᎠᏟᎢᎵᏒ", "RET": "ᎴᏳᏂᎠᏂ ᎠᏟᎢᎵᏒ", "ROTT": "ᎳᏞᎳ ᎠᏟᎢᎵᏒ", "SAKST": "ᏌᎧᎵᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "SAKT": "ᏌᎧᎵᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ᏧᎦᎾᏮ ᎬᎿᎨᏍᏛ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "SBT": "ᏐᎶᎹᏂ ᏚᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "SCT": "ᏎᏤᎴᏏ ᎠᏟᎢᎵᏒ", "SGT": "ᏏᏂᎦᏉᎵ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "SLST": "SLST", "SRT": "ᏒᎵᎾᎻ ᎠᏟᎢᎵᏒ", "SST Samoa": "ᏌᎼᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "SST Samoa Apia": "ᎠᏈᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "SST Samoa Apia DST": "ᎠᏈᎠ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "SST Samoa DST": "ᏌᎼᎠ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "SYOT": "ᏏᏲᏩ ᎠᏟᎢᎵᏒ", "TAAF": "ᎠᏂᎦᎸᏥ ᎤᎦᏃᏮ & ᎤᏁᏍᏓᎶ ᎠᏟᎢᎵᏒ", "TAHT": "ᏔᎯᏘ ᎠᏟᎢᎵᏒ", "TJT": "ᏔᏥᎩᏍᏕᏂ ᎠᏟᎢᎵᏒ", "TKT": "ᏙᎨᎳᎤ ᎠᏟᎢᎵᏒ", "TLT": "ᏗᎧᎸᎬ ᏘᎼᎵ ᎠᏟᎢᎵᏒ", "TMST": "ᏛᎵᎩᎺᏂᏍᏔᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "TMT": "ᏛᎵᎩᎺᏂᏍᏔᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "TOST": "ᏙᎾᎦ ᎪᎩ ᎠᏟᎢᎵᏒ", "TOT": "ᏙᎾᎦ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "TVT": "ᏚᏩᎷ ᎠᏟᎢᎵᏒ", "TWT": "ᏔᏱᏇ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "TWT DST": "ᏔᏱᏇ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒ", "ULAST": "ᎤᎳᏂ ᏆᏙᎸ ᎪᎩ ᎠᏟᎢᎵᏒ", "ULAT": "ᎤᎳᏂ ᏆᏙᎸ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "UYST": "ᏳᎷᏇ ᎪᎩ ᎠᏟᎢᎵᏒ", "UYT": "ᏳᎷᏇ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "UZT": "ᎤᏍᏇᎩᏍᏖᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "UZT DST": "ᎤᏍᏇᎩᏍᏖᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "VET": "ᏪᏁᏑᏪᎳ ᎠᏟᎢᎵᏒ", "VLAST": "ᏭᎳᏗᏬᏍᏙᎩ ᎪᎩ ᎠᏟᎢᎵᏒ", "VLAT": "ᏭᎳᏗᏬᏍᏙᎩ ᎠᏟᎢᎵᏒ ᎠᏟᎢᎵᏒ", "VOLST": "ᏬᎶᎪᏝᏗ ᎪᎩ ᎠᏟᎢᎵᏒ", "VOLT": "ᏬᎶᎪᏝᏗ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "VOST": "ᏬᏍᏙᎧ ᎠᏟᎢᎵᏒ", "VUT": "ᏩᏄᏩᏚ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "VUT DST": "ᏩᏄᏩᏚ ᎪᎩ ᎠᏟᎢᎵᏒ", "WAKT": "ᎤᏰᏨ ᎤᎦᏚᏛᎢ ᎠᏟᎢᎵᏒ", "WARST": "ᏭᏕᎵᎬ ᏗᏜ ᎠᏥᏂᏘᏂᎠ ᎪᎩ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "WART": "ᏭᏕᎵᎬ ᏗᏜ ᎠᏥᏂᏘᏂᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "WAST": "ᏭᏕᎵᎬ ᎬᎿᎨᏍᏛ ᎠᏟᎢᎵᏒ", "WAT": "ᏭᏕᎵᎬ ᎬᎿᎨᏍᏛ ᎠᏟᎢᎵᏒ", "WESZ": "ᏭᏕᎵᎬ ᏗᏜ ᏳᎳᏈ ᎪᎩ ᎠᏟᎢᎵᏒ", "WEZ": "ᏭᏕᎵᎬ ᏗᏜ ᏳᎳᏈ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "WFT": "ᏩᎵᏍ ᎠᎴ ᏊᏚᎾ ᎠᏟᎢᎵᏒ", "WGST": "ᏭᏕᎵᎬ ᎢᏤᏍᏛᏱ ᎪᎩ ᎠᏟᎢᎵᏒ", "WGT": "ᏭᏕᎵᎬ ᎢᏤᏍᏛᏱ ᎠᏟᎶᏍᏗ ᎠᎵᎢᎵᏒ", "WIB": "ᏭᏕᎵᎬ ᏗᏜ ᎢᏂᏙᏂᏍᏯ ᎠᏟᎢᎵᏒ", "WIT": "ᏗᎧᎸᎬ ᏗᏜ ᎢᏂᏙᏂᏍᏯ ᎠᏟᎢᎵᏒ", "WITA": "ᎠᏰᏟ ᎢᏂᏙᏂᏍᏯ ᎠᏟᎢᎵᏒ", "YAKST": "ᏯᎫᏥᎧ ᎪᎩ ᎠᏟᎢᎵᏒ", "YAKT": "ᏯᎫᏥᎧ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "YEKST": "ᏰᎧᏖᎵᏂᏊᎦ ᎪᎩ ᎠᏟᎢᎵᏒ", "YEKT": "ᏰᎧᏖᎵᏂᏊᎦ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "YST": "ᏳᎧᏂ ᎠᏟᎢᎵᏒ", "МСК": "ᎹᏍᎦᏫ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ᏭᏕᎵᎬ ᎧᏎᎧᏍᏕᏂ ᎠᏟᎢᎵᏒ", "شىعىش قازاق ەلى": "ᏗᎧᎸᎬ ᎧᏎᎧᏍᏕᏂ ᎠᏟᎢᎵᏒ", "قازاق ەلى": "ᎧᏎᎧᏍᏕᏂ ᎠᏟᎢᎵᏒ", "قىرعىزستان": "ᎩᎵᏣᎢᏍ ᎠᏟᎢᎵᏒ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ᏇᎷ ᎪᎩ ᎠᏟᎢᎵᏒ"},
	}
}

// Locale returns the current translators string locale
func (chr *chr) Locale() string {
	return chr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'chr'
func (chr *chr) PluralsCardinal() []locales.PluralRule {
	return chr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'chr'
func (chr *chr) PluralsOrdinal() []locales.PluralRule {
	return chr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'chr'
func (chr *chr) PluralsRange() []locales.PluralRule {
	return chr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'chr'
func (chr *chr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'chr'
func (chr *chr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'chr'
func (chr *chr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (chr *chr) MonthAbbreviated(month time.Month) string {
	return chr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (chr *chr) MonthsAbbreviated() []string {
	return chr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (chr *chr) MonthNarrow(month time.Month) string {
	return chr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (chr *chr) MonthsNarrow() []string {
	return chr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (chr *chr) MonthWide(month time.Month) string {
	return chr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (chr *chr) MonthsWide() []string {
	return chr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (chr *chr) WeekdayAbbreviated(weekday time.Weekday) string {
	return chr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (chr *chr) WeekdaysAbbreviated() []string {
	return chr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (chr *chr) WeekdayNarrow(weekday time.Weekday) string {
	return chr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (chr *chr) WeekdaysNarrow() []string {
	return chr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (chr *chr) WeekdayShort(weekday time.Weekday) string {
	return chr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (chr *chr) WeekdaysShort() []string {
	return chr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (chr *chr) WeekdayWide(weekday time.Weekday) string {
	return chr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (chr *chr) WeekdaysWide() []string {
	return chr.daysWide
}

// Decimal returns the decimal point of number
func (chr *chr) Decimal() string {
	return chr.decimal
}

// Group returns the group of number
func (chr *chr) Group() string {
	return chr.group
}

// Group returns the minus sign of number
func (chr *chr) Minus() string {
	return chr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'chr' and handles both Whole and Real numbers based on 'v'
func (chr *chr) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'chr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (chr *chr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, chr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'chr'
func (chr *chr) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := chr.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
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
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, chr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'chr'
// in accounting notation.
func (chr *chr) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := chr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
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

		b = append(b, chr.currencyNegativePrefix[0])

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
			b = append(b, chr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, chr.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'chr'
func (chr *chr) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'chr'
func (chr *chr) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, chr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'chr'
func (chr *chr) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, chr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'chr'
func (chr *chr) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, chr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, chr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'chr'
func (chr *chr) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'chr'
func (chr *chr) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'chr'
func (chr *chr) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'chr'
func (chr *chr) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := chr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
