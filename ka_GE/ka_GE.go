package ka_GE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ka_GE struct {
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

// New returns a new instance of translator for the 'ka_GE' locale
func New() locales.Translator {
	return &ka_GE{
		locale:                 "ka_GE",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "იან", "თებ", "მარ", "აპრ", "მაი", "ივნ", "ივლ", "აგვ", "სექ", "ოქტ", "ნოე", "დეკ"},
		monthsNarrow:           []string{"", "ი", "თ", "მ", "ა", "მ", "ი", "ი", "ა", "ს", "ო", "ნ", "დ"},
		monthsWide:             []string{"", "იანვარი", "თებერვალი", "მარტი", "აპრილი", "მაისი", "ივნისი", "ივლისი", "აგვისტო", "სექტემბერი", "ოქტომბერი", "ნოემბერი", "დეკემბერი"},
		daysAbbreviated:        []string{"კვი", "ორშ", "სამ", "ოთხ", "ხუთ", "პარ", "შაბ"},
		daysNarrow:             []string{"კ", "ო", "ს", "ო", "ხ", "პ", "შ"},
		daysShort:              []string{"კვ", "ორ", "სმ", "ოთ", "ხთ", "პრ", "შბ"},
		daysWide:               []string{"კვირა", "ორშაბათი", "სამშაბათი", "ოთხშაბათი", "ხუთშაბათი", "პარასკევი", "შაბათი"},
		timezones:              map[string]string{"ACDT": "ავსტრალიის ცენტრალური ზაფხულის დრო", "ACST": "ავსტრალიის ცენტრალური სტანდარტული დრო", "ACT": "ACT", "ACWDT": "ცენტრალური და დასავლეთ ავსტრალიის ზაფხულის დრო", "ACWST": "ცენტრალური და დასავლეთ ავსტრალიის სტანდარტული დრო", "ADT": "ატლანტიკის ოკეანის ზაფხულის დრო", "ADT Arabia": "არაბეთის ზაფხულის დრო", "AEDT": "აღმოსავლეთ ავსტრალიის ზაფხულის დრო", "AEST": "აღმოსავლეთ ავსტრალიის სტანდარტული დრო", "AFT": "ავღანეთის დრო", "AKDT": "ალასკის ზაფხულის დრო", "AKST": "ალასკის სტანდარტული დრო", "AMST": "ამაზონიის ზაფხულის დრო", "AMST Armenia": "სომხეთის ზაფხულის დრო", "AMT": "ამაზონიის სტანდარტული დრო", "AMT Armenia": "სომხეთის სტანდარტული დრო", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "არგენტინის ზაფხულის დრო", "ART": "არგენტინის სტანდარტული დრო", "AST": "ატლანტიკის ოკეანის სტანდარტული დრო", "AST Arabia": "არაბეთის სტანდარტული დრო", "AWDT": "დასავლეთ ავსტრალიის ზაფხულის დრო", "AWST": "დასავლეთ ავსტრალიის სტანდარტული დრო", "AZST": "აზერბაიჯანის ზაფხულის დრო", "AZT": "აზერბაიჯანის სტანდარტული დრო", "BDT Bangladesh": "ბანგლადეშის ზაფხულის დრო", "BNT": "ბრუნეი-დარუსალამის დრო", "BOT": "ბოლივიის დრო", "BRST": "ბრაზილიის ზაფხულის დრო", "BRT": "ბრაზილიის სტანდარტული დრო", "BST Bangladesh": "ბანგლადეშის სტანდარტული დრო", "BT": "ბუტანის დრო", "CAST": "CAST", "CAT": "ცენტრალური აფრიკის დრო", "CCT": "ქოქოსის კუნძულების დრო", "CDT": "ჩრდილოეთ ამერიკის ცენტრალური ზაფხულის დრო", "CHADT": "ჩატემის ზაფხულის დრო", "CHAST": "ჩატემის სტანდარტული დრო", "CHUT": "ჩუუკის დრო", "CKT": "კუკის კუნძულების სტანდარტული დრო", "CKT DST": "კუკის კუნძულების ნახევრად ზაფხულის დრო", "CLST": "ჩილეს ზაფხულის დრო", "CLT": "ჩილეს სტანდარტული დრო", "COST": "კოლუმბიის ზაფხულის დრო", "COT": "კოლუმბიის სტანდარტული დრო", "CST": "ჩრდილოეთ ამერიკის ცენტრალური სტანდარტული დრო", "CST China": "ჩინეთის სტანდარტული დრო", "CST China DST": "ჩინეთის დროის სარტყელი", "CVST": "კაბო-ვერდეს ზაფხულის დრო", "CVT": "კაბო-ვერდეს სტანდარტული დრო", "CXT": "შობის კუნძულის დრო", "ChST": "ჩამოროს დრო", "ChST NMI": "ChST NMI", "CuDT": "კუბის ზაფხულის დრო", "CuST": "კუბის სტანდარტული დრო", "DAVT": "დევისის დრო", "DDUT": "დუმონ-დურვილის დრო", "EASST": "აღდგომის კუნძულის ზაფხულის დრო", "EAST": "აღდგომის კუნძულის სტანდარტული დრო", "EAT": "აღმოსავლეთ აფრიკის დრო", "ECT": "ეკვადორის დრო", "EDT": "ჩრდილოეთ ამერიკის აღმოსავლეთის ზაფხულის დრო", "EGDT": "აღმოსავლეთ გრენლანდიის ზაფხულის დრო", "EGST": "აღმოსავლეთ გრენლანდიის სტანდარტული დრო", "EST": "ჩრდილოეთ ამერიკის აღმოსავლეთის სტანდარტული დრო", "FEET": "შორეული აღმოსავლეთ ევროპის დრო", "FJT": "ფიჯის სტანდარტული დრო", "FJT Summer": "ფიჯის ზაფხულის დრო", "FKST": "ფოლკლენდის კუნძულების ზაფხულის დრო", "FKT": "ფოლკლენდის კუნძულების სტანდარტული დრო", "FNST": "ფერნანდო-დე-ნორონიას ზაფხულის დრო", "FNT": "ფერნანდო-დე-ნორონიას სტანდარტული დრო", "GALT": "გალაპაგოსის დრო", "GAMT": "გამბიერის დრო", "GEST": "საქართველოს ზაფხულის დრო", "GET": "საქართველოს სტანდარტული დრო", "GFT": "საფრანგეთის გვიანის დრო", "GIT": "გილბერტის კუნძულების დრო", "GMT": "გრინვიჩის საშუალო დრო", "GNSST": "GNSST", "GNST": "GNST", "GST": "სპარსეთის ყურის სტანდარტული დრო", "GST Guam": "GST Guam", "GYT": "გაიანის დრო", "HADT": "ჰავაისა და ალეუტის ზაფხულის დრო", "HAST": "ჰავაისა და ალეუტის სტანდარტული დრო", "HKST": "ჰონკონგის ზაფხულის დრო", "HKT": "ჰონკონგის სტანდარტული დრო", "HOVST": "ჰოვდის ზაფხულის დრო", "HOVT": "ჰოვდის სტანდარტული დრო", "ICT": "ინდოჩინეთის დრო", "IDT": "ისრაელის ზაფხულის დრო", "IOT": "ინდოეთის ოკეანის კუნძულების დრო", "IRKST": "ირკუტსკის ზაფხულის დრო", "IRKT": "ირკუტსკის სტანდარტული დრო", "IRST": "ირანის სტანდარტული დრო", "IRST DST": "ირანის დროის სარტყელი", "IST": "ინდოეთის დრო", "IST Israel": "ისრაელის სტანდარტული დრო", "JDT": "იაპონიის ზაფხულის დრო", "JST": "იაპონიის სტანდარტული დრო", "KOST": "კოსრეს დრო", "KRAST": "კრასნოიარსკის ზაფხულის დრო", "KRAT": "კრასნოიარსკის სტანდარტული დრო", "KST": "კორეის სტანდარტული დრო", "KST DST": "კორეის ზაფხულის დრო", "LHDT": "ლორდ-ჰაუს ზაფხულის დრო", "LHST": "ლორდ-ჰაუს სტანდარტული დრო", "LINT": "ლაინის კუნძულების დრო", "MAGST": "მაგადანის ზაფხულის დრო", "MAGT": "მაგადანის სტანდარტული დრო", "MART": "მარკიზის კუნძულების დრო", "MAWT": "მოუსონის დრო", "MDT": "MDT", "MESZ": "ცენტრალური ევროპის ზაფხულის დრო", "MEZ": "ცენტრალური ევროპის სტანდარტული დრო", "MHT": "მარშალის კუნძულების დრო", "MMT": "მიანმარის დრო", "MSD": "მოსკოვის ზაფხულის დრო", "MST": "MST", "MUST": "მავრიკის ზაფხულის დრო", "MUT": "მავრიკის სტანდარტული დრო", "MVT": "მალდივების დრო", "MYT": "მალაიზიის დრო", "NCT": "ახალი კალედონიის სტანდარტული დრო", "NDT": "ნიუფაუნდლენდის ზაფხულის დრო", "NDT New Caledonia": "ახალი კალედონიის ზაფხულის დრო", "NFDT": "ნორფოლკის კუნძულის ზაფხულის დრო", "NFT": "ნორფოლკის კუნძულის სტანდარტული დრო", "NOVST": "ნოვოსიბირსკის ზაფხულის დრო", "NOVT": "ნოვოსიბირსკის სტანდარტული დრო", "NPT": "ნეპალის დრო", "NRT": "ნაურუს დრო", "NST": "ნიუფაუნდლენდის სტანდარტული დრო", "NUT": "ნიუეს დრო", "NZDT": "ახალი ზელანდიის ზაფხულის დრო", "NZST": "ახალი ზელანდიის სტანდარტული დრო", "OESZ": "აღმოსავლეთ ევროპის ზაფხულის დრო", "OEZ": "აღმოსავლეთ ევროპის სტანდარტული დრო", "OMSST": "ომსკის ზაფხულის დრო", "OMST": "ომსკის სტანდარტული დრო", "PDT": "ჩრდილოეთ ამერიკის წყნარი ოკეანის ზაფხულის დრო", "PDTM": "მექსიკის წყნარი ოკეანის ზაფხულის დრო", "PETDT": "PETDT", "PETST": "PETST", "PGT": "პაპუა-ახალი გვინეის დრო", "PHOT": "ფენიქსის კუნძულების დრო", "PKT": "პაკისტანის სტანდარტული დრო", "PKT DST": "პაკისტანის ზაფხულის დრო", "PMDT": "სენ-პიერის და მიკელონის ზაფხულის დრო", "PMST": "სენ-პიერის და მიკელონის სტანდარტული დრო", "PONT": "პონაპეს დრო", "PST": "ჩრდილოეთ ამერიკის წყნარი ოკეანის სტანდარტული დრო", "PST Philippine": "ფილიპინების სტანდარტული დრო", "PST Philippine DST": "ფილიპინების ზაფხულის დრო", "PST Pitcairn": "პიტკერნის დრო", "PSTM": "მექსიკის წყნარი ოკეანის სტანდარტული დრო", "PWT": "პალაუს დრო", "PYST": "პარაგვაის ზაფხულის დრო", "PYT": "პარაგვაის სტანდარტული დრო", "PYT Korea": "ფხენიანის დრო", "RET": "რეიუნიონის დრო", "ROTT": "როთერის დრო", "SAKST": "სახალინის ზაფხულის დრო", "SAKT": "სახალინის სტანდარტული დრო", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "სამხრეთ აფრიკის დრო", "SBT": "სოლომონის კუნძულების დრო", "SCT": "სეიშელის კუნძულების დრო", "SGT": "სინგაპურის დრო", "SLST": "SLST", "SRT": "სურინამის დრო", "SST Samoa": "სამოას სტანდარტული დრო", "SST Samoa Apia": "აპიას სტანდარტული დრო", "SST Samoa Apia DST": "აპიას ზაფხულის დრო", "SST Samoa DST": "სამოას ზაფხულის დრო", "SYOT": "სიოვას დრო", "TAAF": "ფრანგული სამხრეთის და ანტარქტიკის დრო", "TAHT": "ტაიტის დრო", "TJT": "ტაჯიკეთის დრო", "TKT": "ტოკელაუს დრო", "TLT": "აღმოსავლეთ ტიმორის დრო", "TMST": "თურქმენეთის ზაფხულის დრო", "TMT": "თურქმენეთის სტანდარტული დრო", "TOST": "ტონგის ზაფხულის დრო", "TOT": "ტონგის სტანდარტული დრო", "TVT": "ტუვალუს დრო", "TWT": "ტაიბეის სტანდარტული დრო", "TWT DST": "ტაიბეის ზაფხულის დრო", "ULAST": "ულან-ბატორის ზაფხულის დრო", "ULAT": "ულან-ბატორის სტანდარტული დრო", "UYST": "ურუგვაის ზაფხულის დრო", "UYT": "ურუგვაის სტანდარტული დრო", "UZT": "უზბეკეთის სტანდარტული დრო", "UZT DST": "უზბეკეთის ზაფხულის დრო", "VET": "ვენესუელის დრო", "VLAST": "ვლადივოსტოკის ზაფხულის დრო", "VLAT": "ვლადივოსტოკის სტანდარტული დრო", "VOLST": "ვოლგოგრადის ზაფხულის დრო", "VOLT": "ვოლგოგრადის სტანდარტული დრო", "VOST": "ვოსტოკის დრო", "VUT": "ვანუატუს სტანდარტული დრო", "VUT DST": "ვანუატუს ზაფხულის დრო", "WAKT": "ვეიკის კუნძულის დრო", "WARST": "დასავლეთ არგენტინის ზაფხულის დრო", "WART": "დასავლეთ არგენტინის სტანდარტული დრო", "WAST": "დასავლეთ აფრიკის დრო", "WAT": "დასავლეთ აფრიკის დრო", "WESZ": "დასავლეთ ევროპის ზაფხულის დრო", "WEZ": "დასავლეთ ევროპის სტანდარტული დრო", "WFT": "ვოლისი და ფუტუნას დრო", "WGST": "დასავლეთ გრენლანდიის ზაფხულის დრო", "WGT": "დასავლეთ გრენლანდიის სტანდარტული დრო", "WIB": "დასავლეთ ინდონეზიის დრო", "WIT": "აღმოსავლეთ ინდონეზიის დრო", "WITA": "ცენტრალური ინდონეზიის დრო", "YAKST": "იაკუტსკის ზაფხულის დრო", "YAKT": "იაკუტსკის სტანდარტული დრო", "YEKST": "ეკატერინბურგის ზაფხულის დრო", "YEKT": "ეკატერინბურგის სტანდარტული დრო", "YST": "იუკონის დრო", "МСК": "მოსკოვის სტანდარტული დრო", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "დასავლეთ ყაზახეთის დრო", "شىعىش قازاق ەلى": "აღმოსავლეთ ყაზახეთის დრო", "قازاق ەلى": "ყაზახეთის დრო", "قىرعىزستان": "ყირგიზეთის დრო", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "აზორის კუნძულების ზაფხულის დრო"},
	}
}

// Locale returns the current translators string locale
func (ka *ka_GE) Locale() string {
	return ka.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ka_GE'
func (ka *ka_GE) PluralsCardinal() []locales.PluralRule {
	return ka.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ka_GE'
func (ka *ka_GE) PluralsOrdinal() []locales.PluralRule {
	return ka.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ka_GE'
func (ka *ka_GE) PluralsRange() []locales.PluralRule {
	return ka.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ka_GE'
func (ka *ka_GE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ka_GE'
func (ka *ka_GE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	iMod100 := i % 100

	if i == 1 {
		return locales.PluralRuleOne
	} else if (i == 0) || (iMod100 >= 2 && iMod100 <= 20 && (iMod100 == 40 || iMod100 == 60 || iMod100 == 80)) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ka_GE'
func (ka *ka_GE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := ka.CardinalPluralRule(num1, v1)
	end := ka.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ka *ka_GE) MonthAbbreviated(month time.Month) string {
	if len(ka.monthsAbbreviated) == 0 {
		return ""
	}
	return ka.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ka *ka_GE) MonthsAbbreviated() []string {
	return ka.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ka *ka_GE) MonthNarrow(month time.Month) string {
	if len(ka.monthsNarrow) == 0 {
		return ""
	}
	return ka.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ka *ka_GE) MonthsNarrow() []string {
	return ka.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ka *ka_GE) MonthWide(month time.Month) string {
	if len(ka.monthsWide) == 0 {
		return ""
	}
	return ka.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ka *ka_GE) MonthsWide() []string {
	return ka.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ka *ka_GE) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ka.daysAbbreviated) == 0 {
		return ""
	}
	return ka.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ka *ka_GE) WeekdaysAbbreviated() []string {
	return ka.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ka *ka_GE) WeekdayNarrow(weekday time.Weekday) string {
	if len(ka.daysNarrow) == 0 {
		return ""
	}
	return ka.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ka *ka_GE) WeekdaysNarrow() []string {
	return ka.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ka *ka_GE) WeekdayShort(weekday time.Weekday) string {
	if len(ka.daysShort) == 0 {
		return ""
	}
	return ka.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ka *ka_GE) WeekdaysShort() []string {
	return ka.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ka *ka_GE) WeekdayWide(weekday time.Weekday) string {
	if len(ka.daysWide) == 0 {
		return ""
	}
	return ka.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ka *ka_GE) WeekdaysWide() []string {
	return ka.daysWide
}

// Decimal returns the decimal point of number
func (ka *ka_GE) Decimal() string {
	return ka.decimal
}

// Group returns the group of number
func (ka *ka_GE) Group() string {
	return ka.group
}

// Group returns the minus sign of number
func (ka *ka_GE) Minus() string {
	return ka.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ka_GE' and handles both Whole and Real numbers based on 'v'
func (ka *ka_GE) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ka.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ka.group) - 1; j >= 0; j-- {
					b = append(b, ka.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ka.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ka_GE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ka *ka_GE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ka.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ka.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ka.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ka_GE'
func (ka *ka_GE) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ka.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ka.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ka.group) - 1; j >= 0; j-- {
					b = append(b, ka.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ka.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ka.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ka.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ka_GE'
// in accounting notation.
func (ka *ka_GE) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ka.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ka.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ka.group) - 1; j >= 0; j-- {
					b = append(b, ka.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ka.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ka.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ka.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ka.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ka.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ka.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ka.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ka.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ka_GE'
func (ka *ka_GE) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ka.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ka.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
