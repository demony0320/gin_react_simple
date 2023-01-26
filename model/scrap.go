package model
import (
    "encoding/json"
    "fmt"
    "net/http"
    "log" 
	"github.com/gin-gonic/gin"
    "github.com/PuerkitoBio/goquery"

    )
type article struct {
    Category string
    Subject string
}

func (a article) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "category": a.Category,
        "subject": a.Subject,
    })
}

//Refactoring with goquery
func (mh *ModelHandler) ScrapHumor(c *gin.Context){
    res, err := http.Get("https://chimhaha.net")
        if err != nil {
            log.Fatal(err)
        }
    defer res.Body.Close()
        if res.StatusCode != 200 {
            log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
        }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    var notices []article
    doc.Find("#boardList.notice > a").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the title
        category := s.Find("div > div.titleContainer > span:nth-child(2)").Text()
        title := s.Find("div > div.titleContainer > span.title > span.text").Text()
        article := article{ Category : category, Subject : title };
        fmt.Println(article)
        notices = append(notices,article)
    })
    doc.Find("#boardList > a").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the title
        category := s.Find("div.info > div.titleContainer > span.category").Text()
        fmt.Println(category)
        title := s.Find("div.info > div.titleContainer > span.title > span.text").Text()
        article := article{ Category : category, Subject : title };
        notices = append(notices,article)
    })
    c.JSON(http.StatusOK,notices);
}

func (mh *ModelHandler) ScrapNews(c *gin.Context){
    res, err := http.Get("https://clien.net")
        if err != nil {
            log.Fatal(err)
        }
    defer res.Body.Close()
        if res.StatusCode != 200 {
            log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
        }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    var notices []article
    doc.Find("#div_content > div.contents_main > div.section_contents.top > div.section_list.recommended > div > div").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the title
        //div:nth-child(1) > div > a.list_subject
        category := s.Find("div > a.list_subject > span.shortname ").Text()
        fmt.Println(category)
        title := s.Find("div > a.list_subject > span.subject").Text()
        article := article{ Category : category, Subject : title };
        notices = append(notices,article)
    })
    c.JSON(http.StatusOK,notices);
}
