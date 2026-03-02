package ml_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ml_IN struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ml_IN' locale
func New() locales.Translator {
	return &ml_IN{
		locale:                 "ml_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "ജനു", "ഫെബ്രു", "മാർ", "ഏപ്രി", "മേയ്", "ജൂൺ", "ജൂലൈ", "ഓഗ", "സെപ്റ്റം", "ഒക്ടോ", "നവം", "ഡിസം"},
		monthsNarrow:           []string{"", "ജ", "ഫെ", "മാ", "ഏ", "മെ", "ജൂൺ", "ജൂ", "ഓ", "സെ", "ഒ", "ന", "ഡി"},
		monthsWide:             []string{"", "ജനുവരി", "ഫെബ്രുവരി", "മാർച്ച്", "ഏപ്രിൽ", "മേയ്", "ജൂൺ", "ജൂലൈ", "ഓഗസ്റ്റ്", "സെപ്റ്റംബർ", "ഒക്\u200cടോബർ", "നവംബർ", "ഡിസംബർ"},
		daysAbbreviated:        []string{"ഞായർ", "തിങ്കൾ", "ചൊവ്വ", "ബുധൻ", "വ്യാഴം", "വെള്ളി", "ശനി"},
		daysNarrow:             []string{"ഞ", "തി", "ചൊ", "ബു", "വ്യാ", "വെ", "ശ"},
		daysShort:              []string{"ഞാ", "തി", "ചൊ", "ബു", "വ്യാ", "വെ", "ശ"},
		daysWide:               []string{"ഞായറാഴ്\u200cച", "തിങ്കളാഴ്\u200cച", "ചൊവ്വാഴ്ച", "ബുധനാഴ്\u200cച", "വ്യാഴാഴ്\u200cച", "വെള്ളിയാഴ്\u200cച", "ശനിയാഴ്\u200cച"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "ഓസ്ട്രേലിയൻ സെൻട്രൽ ഡേലൈറ്റ് സമയം", "ACST": "എയ്ക്കർ വേനൽക്കാല സമയം", "ACT": "എയ്ക്കർ സ്റ്റാൻഡേർഡ് സമയം", "ACWDT": "ഓസ്ട്രേലിയൻ സെൻട്രൽ പടിഞ്ഞാറൻ ഡേലൈറ്റ് സമയം", "ACWST": "ഓസ്ട്രേലിയൻ സെൻട്രൽ പടിഞ്ഞാറൻ സ്റ്റാൻഡേർഡ് സമയം", "ADT": "അറ്റ്\u200cലാന്റിക് ഡേലൈറ്റ് സമയം", "ADT Arabia": "അറേബ്യൻ ഡേലൈറ്റ് സമയം", "AEDT": "ഓസ്\u200cട്രേലിയൻ കിഴക്കൻ ഡേലൈറ്റ് സമയം", "AEST": "ഓസ്\u200cട്രേലിയൻ കിഴക്കൻ സ്റ്റാൻഡേർഡ് സമയം", "AFT": "അഫ്\u200cഗാനിസ്ഥാൻ സമയം", "AKDT": "അലാസ്\u200cക ഡേലൈറ്റ് സമയം", "AKST": "അലാസ്ക സ്റ്റാൻഡേർഡ് സമയം", "AMST": "ആമസോൺ ഗ്രീഷ്\u200cമകാല സമയം", "AMST Armenia": "അർമേനിയ ഗ്രീഷ്\u200cമകാല സമയം", "AMT": "ആമസോൺ സ്റ്റാൻഡേർഡ് സമയം", "AMT Armenia": "അർമേനിയ സ്റ്റാൻഡേർഡ് സമയം", "ANAST": "അനാഡിർ വേനൽക്കാല സമയം", "ANAT": "അനാഡിർ സ്റ്റാൻഡേർഡ് സമയം", "ARST": "അർജന്റീന ഗ്രീഷ്\u200cമകാല സമയം", "ART": "അർജന്റീന സ്റ്റാൻഡേർഡ് സമയം", "AST": "അറ്റ്\u200cലാന്റിക് സ്റ്റാൻഡേർഡ് സമയം", "AST Arabia": "അറേബ്യൻ സ്റ്റാൻഡേർഡ് സമയം", "AWDT": "ഓസ്\u200cട്രേലിയൻ പടിഞ്ഞാറൻ ഡേലൈറ്റ് സമയം", "AWST": "ഓസ്\u200cട്രേലിയൻ പടിഞ്ഞാറൻ സ്റ്റാൻഡേർഡ് സമയം", "AZST": "അസർബൈജാൻ ഗ്രീഷ്\u200cമകാല സമയം", "AZT": "അസർബൈജാൻ സ്റ്റാൻഡേർഡ് സമയം", "BDT Bangladesh": "ബംഗ്ലാദേശ് ഗ്രീഷ്\u200cമകാല സമയം", "BNT": "ബ്രൂണൈ ദാറുസ്സലാം സമയം", "BOT": "ബൊളീവിയ സമയം", "BRST": "ബ്രസീലിയ ഗ്രീഷ്\u200cമകാല സമയം", "BRT": "ബ്രസീലിയ സ്റ്റാൻഡേർഡ് സമയം", "BST Bangladesh": "ബംഗ്ലാദേശ് സ്റ്റാൻഡേർഡ് സമയം", "BT": "ഭൂട്ടാൻ സമയം", "CAST": "CAST", "CAT": "മധ്യ ആഫ്രിക്ക സമയം", "CCT": "കൊക്കോസ് ദ്വീപുകൾ സമയം", "CDT": "സെൻട്രൽ ഡേലൈറ്റ് സമയം", "CHADT": "ചാത്തം ഗ്രീഷ്\u200cമകാല സമയം", "CHAST": "ചാത്തം സ്റ്റാൻഡേർഡ് സമയം", "CHUT": "ചൂക്ക് സമയം", "CKT": "കുക്ക് ദ്വീപുകൾ സ്റ്റാൻഡേർഡ് സമയം", "CKT DST": "കുക്ക് ദ്വീപുകൾ അർദ്ധ ഗ്രീഷ്\u200cമകാല സമയം", "CLST": "ചിലി ഗ്രീഷ്\u200cമകാല സമയം", "CLT": "ചിലി സ്റ്റാൻഡേർഡ് സമയം", "COST": "കൊളംബിയ ഗ്രീഷ്\u200cമകാല സമയം", "COT": "കൊളംബിയ സ്റ്റാൻഡേർഡ് സമയം", "CST": "സെൻട്രൽ സ്റ്റാൻഡേർഡ് സമയം", "CST China": "ചൈന സ്റ്റാൻഡേർഡ് സമയം", "CST China DST": "ചൈന ഡേലൈറ്റ് സമയം", "CVST": "കേപ് വെർദെ ഗ്രീഷ്\u200cമകാല സമയം", "CVT": "കേപ് വെർദെ സ്റ്റാൻഡേർഡ് സമയം", "CXT": "ക്രിസ്\u200cമസ് ദ്വീപ് സമയം", "ChST": "ചമോറോ സ്റ്റാൻഡേർഡ് സമയം", "ChST NMI": "നോർത്ത് മറിയാനാ ദ്വീപുകൾ സമയം", "CuDT": "ക്യൂബ ഡേലൈറ്റ് സമയം", "CuST": "ക്യൂബ സ്റ്റാൻഡേർഡ് സമയം", "DAVT": "ഡേവിസ് സമയം", "DDUT": "ഡുമോണ്ട് ഡി ഉർവില്ലെ സമയം", "EASST": "ഈസ്റ്റർ ദ്വീപ് ഗ്രീഷ്\u200cമകാല സമയം", "EAST": "ഈസ്റ്റർ ദ്വീപ് സ്റ്റാൻഡേർഡ് സമയം", "EAT": "കിഴക്കൻ ആഫ്രിക്ക സമയം", "ECT": "ഇക്വഡോർ സമയം", "EDT": "ഈസ്റ്റേൺ ഡേലൈറ്റ് സമയം", "EGDT": "കിഴക്കൻ ഗ്രീൻലാൻഡ് ഗ്രീഷ്\u200cമകാല സമയം", "EGST": "കിഴക്കൻ ഗ്രീൻലാൻഡ് സ്റ്റാൻഡേർഡ് സമയം", "EST": "ഈസ്റ്റേൺ സ്റ്റാൻഡേർഡ് സമയം", "FEET": "കിഴക്കേയറ്റത്തെ യൂറോപ്യൻ സമയം", "FJT": "ഫിജി സ്റ്റാൻഡേർഡ് സമയം", "FJT Summer": "ഫിജി ഗ്രീഷ്\u200cമകാല സമയം", "FKST": "ഫാക്ക്\u200cലാൻഡ് ദ്വീപുകൾ ഗ്രീഷ്\u200cമകാല സമയം", "FKT": "ഫാക്ക്\u200cലാൻഡ് ദ്വീപുകൾ സ്റ്റാൻഡേർഡ് സമയം", "FNST": "ഫെർണാഡോ ഡി നൊറോൻഹ ഗ്രീഷ്\u200cമകാല സമയം", "FNT": "ഫെർണാഡോ ഡി നൊറോൻഹ സ്റ്റാൻഡേർഡ് സമയം", "GALT": "ഗാലപ്പഗോസ് സമയം", "GAMT": "ഗാമ്പിയർ സമയം", "GEST": "ജോർജ്ജിയ ഗ്രീഷ്\u200cമകാല സമയം", "GET": "ജോർജ്ജിയ സ്റ്റാൻഡേർഡ് സമയം", "GFT": "ഫ്രഞ്ച് ഗയാന സമയം", "GIT": "ഗിൽബേർട്ട് ദ്വീപുകൾ സമയം", "GMT": "ഗ്രീൻവിച്ച് മീൻ സമയം", "GNSST": "GNSST", "GNST": "GNST", "GST": "ഗൾഫ് സ്റ്റാൻഡേർഡ് സമയം", "GST Guam": "ഗ്വാം സ്റ്റാൻഡേർഡ് സമയം", "GYT": "ഗയാന സമയം", "HADT": "ഹവായ്-അലൂഷ്യൻ സ്റ്റാൻഡേർഡ് സമയം", "HAST": "ഹവായ്-അലൂഷ്യൻ സ്റ്റാൻഡേർഡ് സമയം", "HKST": "ഹോങ്കോങ്ങ് ഗ്രീഷ്\u200cമകാല സമയം", "HKT": "ഹോങ്കോങ്ങ് സ്റ്റാൻഡേർഡ് സമയം", "HOVST": "ഹോഡ് ഗ്രീഷ്\u200cമകാല സമയം", "HOVT": "ഹോഡ് സ്റ്റാൻഡേർഡ് സമയം", "ICT": "ഇൻഡോചൈന സമയം", "IDT": "ഇസ്രായേൽ ഡേലൈറ്റ് സമയം", "IOT": "ഇന്ത്യൻ മഹാസമുദ്ര സമയം", "IRKST": "ഇർകസ്\u200cക് ഗ്രീഷ്\u200cമകാല സമയം", "IRKT": "ഇർകസ്ക് സ്റ്റാൻഡേർഡ് സമയം", "IRST": "ഇറാൻ സ്റ്റാൻഡേർഡ് സമയം", "IRST DST": "ഇറാൻ ഡേലൈറ്റ് സമയം", "IST": "ഇന്ത്യൻ സ്റ്റാൻഡേർഡ് സമയം", "IST Israel": "ഇസ്രായേൽ സ്റ്റാൻഡേർഡ് സമയം", "JDT": "ജപ്പാൻ ഡേലൈറ്റ് സമയം", "JST": "ജപ്പാൻ സ്റ്റാൻഡേർഡ് സമയം", "KOST": "കൊസ്ര സമയം", "KRAST": "ക്രാസ്\u200cനോയാർസ്\u200cക് ഗ്രീഷ്\u200cമകാല സമയം", "KRAT": "ക്രാസ്\u200cനോയാർസ്\u200cക് സ്റ്റാൻഡേർഡ് സമയം", "KST": "കൊറിയൻ സ്റ്റാൻഡേർഡ് സമയം", "KST DST": "കൊറിയൻ ഡേലൈറ്റ് സമയം", "LHDT": "ലോർഡ് ഹോവ് ഡേലൈറ്റ് സമയം", "LHST": "ലോർഡ് ഹോവ് സ്റ്റാൻഡേർഡ് സമയം", "LINT": "ലൈൻ ദ്വീപുകൾ സമയം", "MAGST": "മഗാദൻ ഗ്രീഷ്\u200cമകാല സമയം", "MAGT": "മഗാദൻ സ്റ്റാൻഡേർഡ് സമയം", "MART": "മർക്കസസ് സമയം", "MAWT": "മാസൺ സമയം", "MDT": "മൗണ്ടൻ ഡേലൈറ്റ് സമയം", "MESZ": "സെൻട്രൽ യൂറോപ്യൻ ഗ്രീഷ്മകാല സമയം", "MEZ": "സെൻട്രൽ യൂറോപ്യൻ സ്റ്റാൻഡേർഡ് സമയം", "MHT": "മാർഷൽ ദ്വീപുകൾ സമയം", "MMT": "മ്യാൻമാർ സമയം", "MSD": "മോസ്\u200cകോ ഗ്രീഷ്\u200cമകാല സമയം", "MST": "മൗണ്ടൻ സ്റ്റാൻഡേർഡ് സമയം", "MUST": "മൗറീഷ്യസ് ഗ്രീഷ്\u200cമകാല സമയം", "MUT": "മൗറീഷ്യസ് സ്റ്റാൻഡേർഡ് സമയം", "MVT": "മാലിദ്വീപ് സമയം", "MYT": "മലേഷ്യ സമയം", "NCT": "ന്യൂ കാലിഡോണിയ സ്റ്റാൻഡേർഡ് സമയം", "NDT": "ന്യൂഫൗണ്ട്\u200cലാന്റ് ഡേലൈറ്റ് സമയം", "NDT New Caledonia": "ന്യൂ കാലിഡോണിയ ഗ്രീഷ്\u200cമകാല സമയം", "NFDT": "നോർ\u200cഫോക്ക് ദ്വീപ് ഡേലൈറ്റ് സമയം", "NFT": "നോർ\u200cഫോക്ക് ദ്വീപ് സ്റ്റാൻഡേർഡ് സമയം", "NOVST": "നോവോസിബിർസ്\u200cക് ഗ്രീഷ്\u200cമകാല സമയം", "NOVT": "നോവോസിബിർസ്ക് സ്റ്റാൻഡേർഡ് സമയം", "NPT": "നേപ്പാൾ സമയം", "NRT": "നൗറു സമയം", "NST": "ന്യൂഫൗണ്ട്\u200cലാന്റ് സ്റ്റാൻഡേർഡ് സമയം", "NUT": "ന്യൂയി സമയം", "NZDT": "ന്യൂസിലാൻഡ് ഡേലൈറ്റ് സമയം", "NZST": "ന്യൂസിലാൻഡ് സ്റ്റാൻഡേർഡ് സമയം", "OESZ": "കിഴക്കൻ യൂറോപ്യൻ ഗ്രീഷ്മകാല സമയം", "OEZ": "കിഴക്കൻ യൂറോപ്യൻ സ്റ്റാൻഡേർഡ് സമയം", "OMSST": "ഓംസ്\u200cക്ക് ഗ്രീഷ്\u200cമകാല സമയം", "OMST": "ഓംസ്\u200cക്ക് സ്റ്റാൻഡേർഡ് സമയം", "PDT": "പസഫിക് ഡേലൈറ്റ് സമയം", "PDTM": "മെക്സിക്കൻ പസഫിക് ഡേലൈറ്റ് സമയം", "PETDT": "പെട്രോപാവ്\u200cലോസ്ക് കംചാസ്കി വേനൽക്കാല സമയം", "PETST": "പെട്രോപാവ്\u200cലോസ്ക് കംചാസ്കി സ്റ്റാൻഡേർഡ് സമയം", "PGT": "പാപ്പുവ ന്യൂ ഗിനിയ സമയം", "PHOT": "ഫിനിക്\u200cസ് ദ്വീപ് സമയം", "PKT": "പാക്കിസ്ഥാൻ സ്റ്റാൻഡേർഡ് സമയം", "PKT DST": "പാക്കിസ്ഥാൻ ഗ്രീഷ്\u200cമകാല സമയം", "PMDT": "സെന്റ് പിയറി ആൻഡ് മിക്വലൻ ഡേലൈറ്റ് സമയം", "PMST": "സെന്റ് പിയറി ആൻഡ് മിക്വലൻ സ്റ്റാൻഡേർഡ് സമയം", "PONT": "പൊനാപ്പ് സമയം", "PST": "പസഫിക് സ്റ്റാൻഡേർഡ് സമയം", "PST Philippine": "ഫിലിപ്പൈൻ സ്റ്റാൻഡേർഡ് സമയം", "PST Philippine DST": "ഫിലിപ്പൈൻ ഗ്രീഷ്\u200cമകാല സമയം", "PST Pitcairn": "പിറ്റ്കേൻ സമയം", "PSTM": "മെക്\u200cസിക്കൻ പസഫിക് സ്റ്റാൻഡേർഡ് സമയം", "PWT": "പലാവു സമയം", "PYST": "പരാഗ്വേ ഗ്രീഷ്\u200cമകാല സമയം", "PYT": "പരാഗ്വേ സ്റ്റാൻഡേർഡ് സമയം", "PYT Korea": "പ്യോംഗ്\u200cയാംഗ് സമയം", "RET": "റീയൂണിയൻ സമയം", "ROTT": "റോഥെറ സമയം", "SAKST": "സഖാലിൻ ഗ്രീഷ്\u200cമകാല സമയം", "SAKT": "സഖാലിൻ സ്റ്റാൻഡേർഡ് സമയം", "SAMST": "സമാറ വേനൽക്കാല സമയം", "SAMT": "സമാറ സ്റ്റാൻഡേർഡ് സമയം", "SAST": "ദക്ഷിണാഫ്രിക്ക സ്റ്റാൻഡേർഡ് സമയം", "SBT": "സോളമൻ ദ്വീപുകൾ സമയം", "SCT": "സീഷെൽസ് സമയം", "SGT": "സിംഗപ്പൂർ സ്റ്റാൻഡേർഡ് സമയം", "SLST": "ലങ്ക സമയം", "SRT": "സുരിനെയിം സമയം", "SST Samoa": "സമോവ സ്റ്റാൻഡേർഡ് സമയം", "SST Samoa Apia": "അപിയ സ്റ്റാൻഡേർഡ് സമയം", "SST Samoa Apia DST": "അപിയ ഡേലൈറ്റ് സമയം", "SST Samoa DST": "സമോവാ ഗ്രീഷ്\u200cമകാല സമയം", "SYOT": "സയോവ സമയം", "TAAF": "ഫ്രഞ്ച് സതേൺ, അന്റാർട്ടിക് സമയം", "TAHT": "താഹിതി സമയം", "TJT": "താജിക്കിസ്ഥാൻ സമയം", "TKT": "ടോക്കെലൂ സമയം", "TLT": "കിഴക്കൻ തിമോർ സമയം", "TMST": "തുർക്ക്\u200cമെനിസ്ഥാൻ ഗ്രീഷ്\u200cമകാല സമയം", "TMT": "തുർക്ക്\u200cമെനിസ്ഥാൻ സ്റ്റാൻഡേർഡ് സമയം", "TOST": "ടോംഗ ഗ്രീഷ്\u200cമകാല സമയം", "TOT": "ടോംഗ സ്റ്റാൻഡേർഡ് സമയം", "TVT": "ടുവാലു സമയം", "TWT": "തായ്\u200cപെയ് സ്റ്റാൻഡേർഡ് സമയം", "TWT DST": "തായ്\u200cപെയ് ഡേലൈറ്റ് സമയം", "ULAST": "ഉലാൻബാത്തർ ഗ്രീഷ്\u200cമകാല സമയം", "ULAT": "ഉലാൻബാത്തർ സ്റ്റാൻഡേർഡ് സമയം", "UYST": "ഉറുഗ്വേ ഗ്രീഷ്\u200cമകാല സമയം", "UYT": "ഉറുഗ്വേ സ്റ്റാൻഡേർഡ് സമയം", "UZT": "ഉസ്\u200cബെക്കിസ്ഥാൻ സ്റ്റാൻഡേർഡ് സമയം", "UZT DST": "ഉസ്\u200cബെക്കിസ്ഥാൻ ഗ്രീഷ്\u200cമകാല സമയം", "VET": "വെനിസ്വേല സമയം", "VLAST": "വ്ലാഡിവോസ്റ്റോക് ഗ്രീഷ്\u200cമകാല സമയം", "VLAT": "വ്ലാഡിവോസ്റ്റോക് സ്റ്റാൻഡേർഡ് സമയം", "VOLST": "വോൾഗോഗ്രാഡ് ഗ്രീഷ്\u200cമകാല സമയം", "VOLT": "വോൾഗോഗ്രാഡ് സ്റ്റാൻഡേർഡ് സമയം", "VOST": "വോസ്റ്റോക് സമയം", "VUT": "വന്വാതു സ്റ്റാൻഡേർഡ് സമയം", "VUT DST": "വന്വാതു ഗ്രീഷ്\u200cമകാല സമയം", "WAKT": "വേക്ക് ദ്വീപ് സമയം", "WARST": "പടിഞ്ഞാറൻ അർജന്റീന ഗ്രീഷ്\u200cമകാല സമയം", "WART": "പടിഞ്ഞാറൻ അർജന്റീന സ്റ്റാൻഡേർഡ് സമയം", "WAST": "പടിഞ്ഞാറൻ ആഫ്രിക്ക സമയം", "WAT": "പടിഞ്ഞാറൻ ആഫ്രിക്ക സമയം", "WESZ": "പടിഞ്ഞാറൻ യൂറോപ്യൻ ഗ്രീഷ്\u200cമകാല സമയം", "WEZ": "പടിഞ്ഞാറൻ യൂറോപ്യൻ സ്റ്റാൻഡേർഡ് സമയം", "WFT": "വാലിസ് ആന്റ് ഫ്യൂച്യുന സമയം", "WGST": "പടിഞ്ഞാറൻ ഗ്രീൻലാൻഡ് ഗ്രീഷ്\u200cമകാല സമയം", "WGT": "പടിഞ്ഞാറൻ ഗ്രീൻലാൻഡ് സ്റ്റാൻഡേർഡ് സമയം", "WIB": "പടിഞ്ഞാറൻ ഇന്തോനേഷ്യ സമയം", "WIT": "കിഴക്കൻ ഇന്തോനേഷ്യ സമയം", "WITA": "മധ്യ ഇന്തോനേഷ്യ സമയം", "YAKST": "യാകസ്\u200cക്ക് ഗ്രീഷ്\u200cമകാല സമയം", "YAKT": "യാകസ്\u200cക്ക് സ്റ്റാൻഡേർഡ് സമയം", "YEKST": "യെക്കാറ്റരിൻബർഗ് ഗ്രീഷ്\u200cമകാല സമയം", "YEKT": "യെക്കാറ്റരിൻബർഗ് സ്റ്റാൻഡേർഡ് സമയം", "YST": "യൂക്കോൺ സമയം", "МСК": "മോസ്കോ സ്റ്റാൻഡേർഡ് സമയം", "اقتاۋ": "അഖ്തൌ സ്റ്റാൻഡേർഡ് സമയം", "اقتاۋ قالاسى": "അഖ്തൌ വേനൽക്കാല സമയം", "اقتوبە": "അഖ്തോബ് സ്റ്റാൻഡേർഡ് സമയം", "اقتوبە قالاسى": "അഖ്തോബ് വേനൽക്കാല സമയം", "الماتى": "അൽമതി സ്റ്റാൻഡേർഡ് സമയം", "الماتى قالاسى": "അൽമതി വേനൽക്കാല സമയം", "باتىس قازاق ەلى": "പടിഞ്ഞാറൻ കസാഖിസ്ഥാൻ സമയം", "شىعىش قازاق ەلى": "കിഴക്കൻ കസാഖിസ്ഥാൻ സമയം", "قازاق ەلى": "കസാഖിസ്ഥാൻ സമയം", "قىرعىزستان": "കിർഗിസ്ഥാൻ സമയം", "قىزىلوردا": "ഖിസിലോർഡ സ്റ്റാൻഡേർഡ് സമയം", "قىزىلوردا قالاسى": "ഖിസിലോർഡ വേനൽക്കാല സമയം", "∅∅∅": "അസോർസ് ഗ്രീഷ്\u200cമകാല സമയം"},
	}
}

// Locale returns the current translators string locale
func (ml *ml_IN) Locale() string {
	return ml.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ml_IN'
func (ml *ml_IN) PluralsCardinal() []locales.PluralRule {
	return ml.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ml_IN'
func (ml *ml_IN) PluralsOrdinal() []locales.PluralRule {
	return ml.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ml_IN'
func (ml *ml_IN) PluralsRange() []locales.PluralRule {
	return ml.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ml_IN'
func (ml *ml_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ml_IN'
func (ml *ml_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ml_IN'
func (ml *ml_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := ml.CardinalPluralRule(num1, v1)
	end := ml.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ml *ml_IN) MonthAbbreviated(month time.Month) string {
	return ml.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ml *ml_IN) MonthsAbbreviated() []string {
	return ml.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ml *ml_IN) MonthNarrow(month time.Month) string {
	return ml.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ml *ml_IN) MonthsNarrow() []string {
	return ml.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ml *ml_IN) MonthWide(month time.Month) string {
	return ml.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ml *ml_IN) MonthsWide() []string {
	return ml.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ml *ml_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return ml.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ml *ml_IN) WeekdaysAbbreviated() []string {
	return ml.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ml *ml_IN) WeekdayNarrow(weekday time.Weekday) string {
	return ml.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ml *ml_IN) WeekdaysNarrow() []string {
	return ml.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ml *ml_IN) WeekdayShort(weekday time.Weekday) string {
	return ml.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ml *ml_IN) WeekdaysShort() []string {
	return ml.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ml *ml_IN) WeekdayWide(weekday time.Weekday) string {
	return ml.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ml *ml_IN) WeekdaysWide() []string {
	return ml.daysWide
}

// Decimal returns the decimal point of number
func (ml *ml_IN) Decimal() string {
	return ml.decimal
}

// Group returns the group of number
func (ml *ml_IN) Group() string {
	return ml.group
}

// Group returns the minus sign of number
func (ml *ml_IN) Minus() string {
	return ml.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ml_IN' and handles both Whole and Real numbers based on 'v'
func (ml *ml_IN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, ml.group[0])
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
		b = append(b, ml.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ml_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ml *ml_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ml.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ml.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ml_IN'
func (ml *ml_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ml.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ml.group[0])
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
		b = append(b, ml.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ml.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ml_IN'
// in accounting notation.
func (ml *ml_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ml.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ml.group[0])
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

		b = append(b, ml.currencyNegativePrefix[0])

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
			b = append(b, ml.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ml.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ml.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ml.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ml.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ml.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ml.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ml.periodsAbbreviated[0]...)
	} else {
		b = append(b, ml.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ml.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ml.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ml.periodsAbbreviated[0]...)
	} else {
		b = append(b, ml.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ml.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ml.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ml.periodsAbbreviated[0]...)
	} else {
		b = append(b, ml.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ml_IN'
func (ml *ml_IN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	tz, _ := t.Zone()

	if btz, ok := ml.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ml.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ml.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ml.periodsAbbreviated[0]...)
	} else {
		b = append(b, ml.periodsAbbreviated[1]...)
	}

	return string(b)
}
