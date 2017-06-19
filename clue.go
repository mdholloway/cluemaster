package main

import (
    "os"
    "fmt"
    "log"
    "flag"
    "time"
    "strings"
    "strconv"
    "io/ioutil"
    "net/url"
    "net/http"
    "github.com/mmcdole/gofeed"
    "github.com/olekukonko/tablewriter"
    "github.com/briandowns/spinner"
)

// TODO: Parse from https://www.craigslist.org/about/sites

var testPrefixes = [...]string { "annarbor", "cleveland", "grandrapids", "chicago" }

var prefixes = [...]string {
            "auburn", "bham", "dothan", "shoals", "gadsden", "huntsville", "mobile", "montgomery", "tuscaloosa", // AL
            "anchorage", "fairbanks", "kenai", "juneau", // AK
            "flagstaff", "mohave", "phoenix", "prescott", "showlow", "sierravista", "tucson", "yuma", // AZ
            "fayar", "fortsmith", "jonesboro", "littlerock", "texarkana", // AR
            "bakersfield", "chico", "fresno", "goldcountry", "hanford", "humboldt", "imperial", "inlandempire", "losangeles", "mendocino", "merced", "modesto", "monterey", "orangecounty", "palmsprings", "redding", "sacramento", "sandiego", "sfbay", "slo", "santabarbara", "santamaria", "siskiyou", "stockton", "susanville", "ventura", "visalia", "yubasutter", // CA
            "boulder", "cosprings", "denver", "eastco", "fortcollins", "rockies", "pueblo", "westslope", // CO
            "newlondon", "hartford", "newhaven", "nwct", // CT
            "delaware", // DE
            "washingtondc", // DC
            "daytona", "keys", "fortlauderdale", "fortmyers", "gainesville", "cfl", "jacksonville", "lakeland", "lakecity", "miami", "ocala", "okaloosa", "orlando", "panamacity", "pensacola", "sarasota", "spacecoast", "staugustine", "tallahassee", "tampa", "treasure", // FL
            "albanyga", "athensga", "atlanta", "augusta", "brunswick", "columbusga", "macon", "nwga", "savannah", "statesboro", "valdosta", // GA
            "honolulu", // HI
            "boise", "eastidaho", "lewiston", "twinfalls", // ID
            "bn", "chambana", "chicago", "decatur", "lasalle", "mattoon", "peoria", "rockford", "carbondale", "springfieldil", "quincy", // IL
            "bloomington", "evansville", "fortwayne", "indianapolis", "kokomo", "tippecanoe", "muncie", "richmondin", "southbend", "terrehaute", // IN
            "ames", "cedarrapids", "desmoines", "dubuque", "fortdodge", "iowacity", "masoncity", "quadcities", "siouxcity", "ottumwa", "waterloo", // IA
            "lawrence", "ksu", "nwks", "salina", "seks", "swks", "topeka", "wichita", // KS
            "bgky", "eastky", "lexington", "louisville", "owensboro", "westky", // KY
            "batonrouge", "cenla", "houma", "lafayette", "lakecharles", "monroe", "neworleans", "shreveport", // LA
            "maine", // ME
            "annapolis", "baltimore", "easternshore", "frederick", "smd", "westmd", // MD
            "boston", "capecod", "southcoast", "westernmass", "worcester", // MA
            "annarbor", "battlecreek", "centralmich", "detroit", "flint", "grandrapids", "holland", "jxn", "kalamazoo", "lansing", "monroemi", "muskegon", "nmi", "porthuron", "saginaw", "swmi", "thumb", "up", // MI
            "bemidji", "brainerd", "duluth", "mankato", "minneapolis", "rmn", "marshall", "stcloud", // MN
            "gulfport", "hattiesburg", "jackson", "meridian", "northmiss", "natchez", // MS
            "columbiamo", "joplin", "kansascity", "kirksville", "loz", "semo", "springfield", "stjoseph", "stlouis", // MO
            "billings", "bozeman", "butte", "greatfalls", "helena", "kalispell", "missoula", "montana", // MT
            "grandisland", "lincoln", "northplatte", "omaha", "scottsbluff", // NE
            "elko", "lasvegas", "reno", // NV
            "nh", // NH
            "cnj", "jerseyshore", "newjersey", "southjersey", // NJ
            "albuquerque", "clovis", "farmington", "lascruces", "roswell", "santafe", // NM
            "albany", "binghamton", "buffalo", "catskills", "chautauqua", "elmira", "fingerlakes", "glensfalls", "hudsonvalley", "ithaca", "longisland", "newyork", "oneonta", "plattsburgh", "potsdam", "rochester", "syracuse", "twintiers", "utica", "watertown", // NY
            "asheville", "boone", "charlotte", "eastnc", "fayetteville", "greensboro", "hickory", "onslow", "outerbanks", "raleigh", "wilmington", "winstonsalem", // NC
            "bismarck", "fargo", "grandforks", "nd", // ND
            "akroncanton", "ashtabula", "athensohio", "chillicothe", "cincinnati", "cleveland", "columbus", "dayton", "limaohio", "mansfield", "sandusky", "toledo", "tuscarawas", "youngstown", "zanesville", // OH
            "lawton", "enid", "oklahomacity", "stillwater", "tulsa", // OK
            "bend", "corvallis", "eastoregon", "eugene", "klamath", "medford", "oregoncoast", "portland", "roseburg", "salem", // OR
            "altoona", "chambersburg", "erie", "harrisburg", "lancaster", "allentown", "meadville", "philadelphia", "pittsburgh", "poconos", "reading", "scranton", "pennstate", "williamsport", "york", // PA
            "providence", // RI
            "charleston", "columbia", "florencesc", "greenville", "hiltonhead", "myrtlebeach", // SC
            "nesd", "csd", "rapidcity", "siouxfalls", "sd", // SD
            "chattanooga", "clarksville", "cookeville", "jacksontn", "knoxville", "memphis", "nashville", "tricities", // TN
            "abilene", "amarillo", "austin", "beaumont", "brownsville", "collegestation", "corpuschristi", "dallas", "nacogdoches", "delrio", "elpaso", "galveston", "houston", "killeen", "laredo", "lubbock", "mcallen", "odessa", "sanangelo", "sanantonio", "sanmarcos", "bigbend", "texoma", "easttexas", "victoriatx", "waco", "wichitafalls", // TX
            "logan", "ogden", "provo", "saltlakecity", "stgeorge", // UT
            "vermont", // VT
            "charlottesville", "danville", "fredericksburg", "norfolk", "harrisonburg", "lynchburg", "blacksburg", "richmond", "roanoke", "swva", "winchester", // VA
            "bellingham", "kpr", "moseslake", "olympic", "pullman", "seattle", "skagit", "spokane", "wenatchee", "yakima", // WA
            "charlestonwv", "martinsburg", "huntington", "morgantown", "wheeling", "parkersburg", "swv", "wv", // WV
            "appleton", "eauclaire", "greenbay", "janesville", "racine", "lacrosse", "madison", "milwaukee", "northernwi", "sheboygan", "wausau", // WI
            "wyoming" } // WY

const QPS = 5

type RawResponse struct {
    Locale string
    Content string
}

type Result struct {
    Locale string
    Price string
    Title string
    Link string
}

func avg(values []float64) float64 {
    var total float64;
    for _, val := range values {
        total += val
    }
    return total / float64(len(values))
}

func getListings(prefix string, url string, client *http.Client, ch chan<-*RawResponse) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Println(err)
    }
    req.Header.Set("user-agent", "cluemaster/0.1 contact mdh@hollowlog.co =)")
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
    }
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
    }
    content := string(bytes)
    ch <- &RawResponse{Locale: prefix, Content: content}
    resp.Body.Close()
}

func handleQuery(query string) {
    ch := make(chan *RawResponse)
    throttle := time.Tick(time.Second / QPS)
    client := &http.Client{}
    parser := gofeed.NewParser()
    var results []*Result
    var prices []float64

    s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
    s.Start()

    for _, prefix := range prefixes {
        <-throttle
        url := fmt.Sprintf("https://%s.craigslist.org/search/bia?format=rss&srchType=T&query=%s",
            prefix, url.QueryEscape(query))
        go getListings(prefix, url, client, ch)

        response := <- ch
        locale := response.Locale
        feed, err := parser.ParseString(response.Content)
        if err != nil {
            log.Println(err)
        }

        for _, item := range feed.Items {
            // TODO: decode correctly, don't just string-replace
            title := strings.Replace(item.Title, "&#x0024;", "$", -1)
            var displayTitle string
            var displayPrice string
            if strings.Contains(title, "$") {
                idx := strings.LastIndex(title, "$")
                displayTitle = strings.TrimSpace(title[:idx])
                displayPrice = title[idx+1:]
            } else {
                displayTitle = title
            }
            price, _ := strconv.ParseFloat(displayPrice, 64)
            result := &Result{Locale: locale, Price: displayPrice, Title: displayTitle, Link: item.Link}
            results = append(results, result)
            prices = append(prices, price)
        }
    }
    s.Stop()
    if len(prices) > 0 {
        table := tablewriter.NewWriter(os.Stdout)
        table.SetHeader([]string{"Locale", "Price", "Title", "Link"})
        for _, result := range results {
            table.Append([]string{result.Locale, result.Price, result.Title, result.Link})
        }
        table.Render()
        fmt.Println("Total results: " + strconv.Itoa(len(results)))
        fmt.Println("Average price: $" + strconv.FormatFloat(avg(prices), 'f', 2, 64))
    } else {
        fmt.Println("No results found. :[")
    }
}

func main() {
    query := flag.String("q", "", "Search query")
    flag.Parse()
    if *query == "" {
        fmt.Println("Please specify a query!\nUsage:\n  ./clue -q=query")
        return
    }
    fmt.Printf("Searching: %s (please be patient)\n\n", *query)
    handleQuery(*query)
}
