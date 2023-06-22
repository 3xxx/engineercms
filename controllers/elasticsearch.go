package controllers

import (
	//"bytes"
	//"context"
	//"encoding/json"
	//"fmt"
	//"log"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/elastic/go-elasticsearch/esapi"
	//"github.com/beego/beego/v2/adapter/logs"
	"regexp"

	// "bytes"
	// "context"
	// "encoding/json"
	// "log"
	"strconv"
	//"strings"
	"sync"

	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"bytes"
	"encoding/json"
	"flag"
	// "strconv"
	// "sync/atomic"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-tika/tika"

	//"github.com/cenkalti/backoff/v4"
	// "github.com/dustin/go-humanize"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	// "github.com/elastic/go-elasticsearch/v8/esutil"
	"crypto/tls"
	"io"
	"math/rand"
	"net"
	"net/http"
	"path"
)

//type Article struct {
//	ID        int       `json:"id"`
//	Title     string    `json:"title"`
//	Body      string    `json:"body"`
//	Published time.Time `json:"published"`
//	Author    Author    `json:"author"`
//}
//
//type Author struct {
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//}

// const baseURL = "https://xkcd.com"
// http://127.0.0.1:8081/v1/wx/standardpdf?file=/v1/wx/downloadstandard/16
// 注意点：
// 发起的请求，如果成功了，一定要记得关闭返回Response的Body,否则会占用一个连接。
// 全局变量和函数
var es *elasticsearch.Client

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const baseURL = "/v1/wx/standardpdf?file=/v1/wx/downloadstandard/"

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
)

func init() {
	var err error
	config := elasticsearch.Config{}
	config.Addresses = []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	es, err = elasticsearch.NewClient(config)
	checkError(err)

	flag.StringVar(&indexName, "index", "standard-index", "Index name")
	//flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	//flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	//flag.IntVar(&numItems, "count", 10, "Number of documents to generate")
	flag.Parse()
	//setupIndex(indexName)
	openElasticsearch, err := web.AppConfig.String("openElasticsearch")
	if err != nil {
		logs.Error(err)
	}
	if openElasticsearch == "true" {
		// beego.Info("建立index")
		if err = setupIndex(indexName); err != nil { //这里建立索引（相当于表）
			log.Fatalf("Error creating the client: %s", err)
		}
	}
	//rand.Seed(time.Now().UnixNano())
}

// Document wraps an xkcd.com comic.
type Document struct {
	ID        string `json:"id"`
	ImageURL  string `json:"image_url"`
	Published string `json:"published"`

	Title      string `json:"title"`
	Body       string `json:"body"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	Link       string `json:"link,omitempty"`
	News       string `json:"news,omitempty"`
}

// SearchResults wraps the Elasticsearch search response.
type SearchResults struct {
	Total int    `json:"total"`
	Hits  []*Hit `json:"hits"`
}

// Hit wraps the document returned in search response.
type Hit struct {
	//Document
	URL        string        `json:"url"`
	Sort       []interface{} `json:"sort"`
	Highlights *struct {
		Title      []string `json:"title"`
		Body       []string `json:"body"`
		Alt        []string `json:"alt"`
		Transcript []string `json:"transcript"`
	} `json:"highlights,omitempty"`
	ID        string `json:"id"`
	ImageURL  string `json:"image_url"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Published string `json:"published"`
	Author    Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//"hits" : [
//{
//"_index" : "test-bulk-example",
//"_type" : "_doc",
//"_id" : "2",
//"_score" : 0.8938179,
//"_source" : {
//"id" : 2,
//"title" : "Title 2",
//"body" : "Lorem ipsum dolor sit amet...",
//"published" : "2021-10-29T11:34:32Z",
//"author" : {
//"first_name" : "John",
//"last_name" : "Smith"
//}
//}
//},

// CMSELASTIC API
type ElasticController struct {
	web.Controller
}

// @Title get elasticsearch web
// @Description get elasticsearch web
// @Success 200 {object} models.GetElastic
// @Failure 400 Invalid page supplied
// @Failure 404 Elastic not found
// @router /get [get]
// 进入搜索主页面
func (c *ElasticController) Get() {
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "elastic/index.html"
	} else {
		c.TplName = "elastic/index.html"
	}
}

// @Title upload file to tika&elastic html
// @Description get upload file to tika&elastic html
// @Param id path string true "The id of project"
// @Success 200 {object} models.Elastic
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /uploadelastic/:id [get]
// 进入上传主页面
func (c *ElasticController) UploadElastic() {
	_, _, _, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "elastic/upload_elastic.tpl"
	} else {
		c.TplName = "elastic/upload_elastic.tpl"
	}
}

// @Title post bootstrapfileinput
// @Description post file by BootstrapFileInput
// @Param id path string true "The id of project"
// @Param prodlabel query string false "The prodlabel"
// @Param prodprincipal query string false "The prodprincipal"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 page not found
// @router /upload/:id [post]
// 上传
func (c *ElasticController) Upload() {
	// 取得用户名
	_, _, uid, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}

	openElasticsearch, err := web.AppConfig.String("openElasticsearch")
	if err != nil {
		logs.Error(err)
	}
	if openElasticsearch == "false" {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "用户未开启elastic全文检索服务！"}
		c.ServeJSON()
		return
	}

	//获取上传的文件
	_, h, err := c.GetFile("input-ke-2[]")
	// beego.Info(h.Filename)自动将英文括号改成了_下划线
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	if fileSuffix != ".DOC" && fileSuffix != ".doc" && fileSuffix != ".DOCX" && fileSuffix != ".docx" && fileSuffix != ".pdf" && fileSuffix != ".PDF" {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件类型错误，请上传doc或pdf"}
		c.ServeJSON()
		return
	}
	prodlabel := c.GetString("prodlabel")
	prodprincipal := c.GetString("prodprincipal")

	id := c.Ctx.Input.Param(":id")
	id_int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	var filepath, DiskDirectory, Url, attachmentname, article_body, filename1, filename2 string
	var filesize, attachmentid int64
	//由proj id取得url
	Url, DiskDirectory, err = GetUrlPath(id_int64)
	if err != nil {
		logs.Error(err)
	}

	//取项目本身
	category, err := models.GetProj(id_int64)
	if err != nil {
		logs.Error(err)
	}
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	if h != nil {
		//保存附件——要防止重名覆盖！！！！先判断是否存在！！！
		// filepath := "./attachment/mathcad/" + username + "/" + h.Filename
		// filepath := "./attachment/mathcad/" + username + "/" + h.Filename
		// Url := "/attachment/mathcad/" + username + "/"

		//将附件的编号和名称写入数据库
		_, filename1, filename2, _, _, _, _ = Record(h.Filename)
		// filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		//20190728改名，替换文件名中的#和斜杠
		filename2 = strings.Replace(filename2, "#", "号", -1)
		filename2 = strings.Replace(filename2, "/", "-", -1)
		FileSuffix := path.Ext(h.Filename)
		attachmentname = filename1 + filename2 + FileSuffix
		// code := filename1
		// title := filename2
		//存入成果数据库
		//如果编号重复，则不写入，只返回Id值。
		//根据id添加成果code, title, label, principal, content string, projectid int64
		prodId, err := models.AddProduct(filename1, filename2, prodlabel, prodprincipal, uid, id_int64, topprojectid)
		if err != nil {
			logs.Error(err)
		}

		filepath = DiskDirectory + "/" + attachmentname
		//如果附件名称相同，则覆盖上传，但数据库不追加
		attachmentid, err = models.AddAttachment(attachmentname, filesize, 0, prodId)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "写入附件数据库错误"}
			c.ServeJSON()
		} else {
			// 如果文件存在，则返回
			if PathisExist(filepath) {
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件已存在！"}
				c.ServeJSON()
				return
			}
			err = c.SaveToFile("input-ke-2[]", filepath) //.Join("attachment", attachment)) //存文件
			if err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件保存错误！"}
				c.ServeJSON()
				return
			}
			filesize, _ = FileSize(filepath)
			filesize = filesize / 1000.0

			//存入文件夹
			// err = c.SaveToFile("file", filepath) //.Join("attachment", attachment)) //存文件    WaterMark(filepath)    //给文件加水印
			// if err != nil {
			// 	logs.Error(err)
			// 	c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "存入文件夹错误"}
			// 	c.ServeJSON()
			// }
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": attachmentname, "original": attachmentname, "url": Url + "/" + attachmentname}
			c.ServeJSON()
		}
	}

	cwd, _ := os.Getwd()
	// writeExample(excel, workbooks, cwd+"\\write.xls")
	// readExample(cwd+"\\excel97-2003.xls", excel, workbooks)
	//坑1，必须要输入磁盘，如果是当前目录里的模板，则无法打开！！
	// DiskDirectory = "./attachment/" + proj1.Code + proj1.Title
	DiskDirectory = strings.Replace(DiskDirectory, "./attachment/", "/attachment/", -1)
	f, err := os.Open(cwd + DiskDirectory + "/" + attachmentname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(f.Name())
	client := tika.NewClient(nil, "http://localhost:9998")
	body, err := client.Parse(context.Background(), f)
	// body, err := client.Detect(context.Background(), f) //application/pdf
	// fmt.Println(err)
	// fmt.Println(body)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		// if selection.Text() != " " || selection.Text() != "\n" {
		// fmt.Println(selection.Text())
		text := strings.Replace(selection.Text(), "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)
		article_body = article_body + text //selection.Text()
		// }
	})

	now := time.Now()
	year, month, day := now.Date()
	today_str := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	rand.Seed(time.Now().Unix())
	doc := &Document{
		//是productid还是attachmentid
		ID:        strconv.FormatInt(attachmentid, 10),
		ImageURL:  "/static/images/" + strconv.Itoa(rand.Intn(8)) + "s.jpg", //1s.jpg
		Published: today_str,                                                //fmt.Sprintf("%04d-%02d-%02d", jYear, jMonth, jDay),
		Title:     filename1 + filename2,
		Body:      article_body,
	}

	err = Createitem(indexName, doc)
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	r  map[string]interface{}
	wg sync.WaitGroup
)

//const mapping = `{
//    "mappings": {
//      "_doc": {
//        "properties": {
//          "id":         { "type": "keyword" },
//          "image_url":  { "type": "keyword" },
//          "title":      { "type": "text", "analyzer": "ik_max_word" },
//          "body":       { "type": "text", "analyzer": "ik_max_word" },
//          "alt":        { "type": "text", "analyzer": "ik_max_word" },
//          "transcript": { "type": "text", "analyzer": "ik_max_word" },
//          "published":  { "type": "date" },
//          "link":       { "type": "keyword" },
//          "news":       { "type": "text", "analyzer": "ik_max_word" }
//        }
//      }
//    }
//}`

func setupIndex(string) error {
	mapping := `{
    "mappings": {
       "properties": {
          "id":         { "type": "keyword" },
          "image_url":  { "type": "keyword" },
          "title":      { "type": "text", "analyzer": "ik_max_word" },
          "body":       { "type": "text", "analyzer": "ik_max_word" },
          "alt":        { "type": "text", "analyzer": "ik_max_word" },
          "transcript": { "type": "text", "analyzer": "ik_max_word" },
          "published":  { "type": "date" },
          "link":       { "type": "keyword" },
          "news":       { "type": "text", "analyzer": "ik_max_word" }
       }
    }
	}`
	return CreateIndex(mapping)
}

// CreateIndex creates a new index with mapping.
// func (c *ElasticController) CreateIndex(mapping string) error {
func CreateIndex(mapping string) error {
	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}

	// // Re-create the index
	// if res, err = es.Indices.Delete([]string{indexName}, es.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
	// 	log.Fatalf("Cannot delete index: %s", err)
	// }
	// res.Body.Close()

	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		logs.Error(err)
		return err
	}
	//s, _ := ioutil.ReadAll(res.Body)
	// beego.Info(res.Body)
	// beego.Info(res.Status()) //1.存在：200 OK 2.不存在：404 Not Found 3.未启动：
	res.Body.Close()
	if res.Status() == "404 Not Found" {
		//beego.Info(err)
		res2, err := es.Indices.Create(indexName, es.Indices.Create.WithBody(strings.NewReader(mapping)))
		if err != nil {
			return err
		}
		if res2.IsError() {
			return fmt.Errorf("error: %s", res2)
		}
		return nil
	}
	return nil
}

// Create indexes a new document into store.
// 这里插入数据
func Createitem(indexName string, item *Document) error {
	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	// logs.Error(err)
	//	return err
	//	// log.Fatalf("Error creating the client: %s", err)
	//}
	payload, err := json.Marshal(item)
	if err != nil {
		return err
	}

	ctx := context.Background()
	res, err := esapi.CreateRequest{ //这里插入数据
		Index:      indexName,
		DocumentID: item.ID,
		Body:       bytes.NewReader(payload),
	}.Do(ctx, es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}
	return nil
}

// Exists returns true when a document with id already exists in the store.
func (c *ElasticController) Exists(id string) (bool, error) {
	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}
	res, err := es.Exists(indexName, id)
	if err != nil {
		return false, err
	}
	switch res.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	default:
		return false, fmt.Errorf("[%s]", res.Status())
	}
}

// @Title get search
// @Description get earch
// @Param q query string false "The query="
// @Param a formData string false "The after..."
// @Success 200 {object} models.GetSearch
// @Failure 400 Invalid page supplied
// @Failure 404 Search not found
// @router /search [get]
// 上传文档——tika解析——构造document——存入es
// Search returns results matching a query, paginated by after.
// //(*SearchResults, error)
func (c *ElasticController) Search() {
	// 3. Search for the indexed documents
	// Build the request body.
	query := c.GetString("q")
	after := c.GetStrings("a")
	//beego.Info(after)
	//if query == "" {
	//	query = "目标"
	//}
	// *************
	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}
	var results SearchResults
	//Search returns results matching a query, paginated by after.
	res, err := es.Search(
		es.Search.WithIndex(indexName),
		es.Search.WithBody(buildQuery(query, after...)),
	)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"status": 0, "info": "ERR"}
		c.ServeJSON()
		//	return &results, err
	}

	//***************
	//var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			//"title": "Title",
	//			"body": q,
	//			//"author.first_name": "John",
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"fields": map[string]interface{}{
	//			"title":      map[string]interface{}{"number_of_fragments": 0},
	//			"alt":        map[string]interface{}{"number_of_fragments": 0},
	//			"body":       map[string]interface{}{"number_of_fragments": 0},
	//			"transcript": map[string]interface{}{"number_of_fragments": 5, "fragment_size": 25},
	//		},
	//		//"size" : 25,
	//		//"pre_tags": "<font color='red'>",
	//		//"post_tags": "</font>",
	//		"fragment_size": 10,
	//	},
	//	"size": 25,
	//}
	//if err := json.NewEncoder(&buf).Encode(query); err != nil {
	//	log.Fatalf("Error encoding query: %s", err)
	//}
	//// Perform the search request.
	//res, err := es.Search(
	//	es.Search.WithContext(context.Background()),
	//	es.Search.WithIndex(indexName), // default indexname
	//	es.Search.WithBody(&buf),
	//	es.Search.WithTrackTotalHits(true),
	//	es.Search.WithPretty(),
	//)
	//if err != nil {
	//	//return //&results, err
	//	c.Data["json"] = map[string]interface{}{"status": 0, "info": "ERR"}
	//	c.ServeJSON()
	//}
	//**********

	defer res.Body.Close()
	//log.Printf(res.String())

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			//return //&results, err
			c.Data["json"] = map[string]interface{}{"status": 1, "info": "ERR"}
			c.ServeJSON()
		}
		//return //&results, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		c.Data["json"] = map[string]interface{}{"status": 2, "info": "ERR"}
		c.ServeJSON()
	}

	type envelopeResponse struct {
		Took int
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				ID         string          `json:"_id"`
				Source     json.RawMessage `json:"_source"`
				Highlights json.RawMessage `json:"highlight"`
				Sort       []interface{}   `json:"sort"`
			}
		}
	}

	var r envelopeResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		//return //&results, err
		c.Data["json"] = map[string]interface{}{"status": 3, "info": "ERR"}
		c.ServeJSON()
	}

	results.Total = r.Hits.Total.Value

	if len(r.Hits.Hits) < 1 {
		results.Hits = []*Hit{}
		//return //&results, nil
		c.Data["json"] = &results //map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}

	for _, hit := range r.Hits.Hits {
		var h Hit
		h.ID = hit.ID
		h.Sort = hit.Sort
		//如果是office文件，则/officeview/id，如果是pdf文件，则/pdf?id=
		//根据id（attatchment id）判断附件类型
		//h.URL = strings.Join([]string{baseURL, h.ID, ""}, "/")
		h.URL = strings.Join([]string{baseURL, h.ID}, "?id=")

		if err := json.Unmarshal(hit.Source, &h); err != nil {
			//return //&results, err
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"status": 5, "info": "ERR"}
			c.ServeJSON()
		}

		if len(hit.Highlights) > 0 {
			if err := json.Unmarshal(hit.Highlights, &h.Highlights); err != nil {
				//return //&results, err
				c.Data["json"] = map[string]interface{}{"status": 6, "info": "ERR"}
				c.ServeJSON()
			}
		}

		results.Hits = append(results.Hits, &h)
	}
	//beego.Info(&results)
	//return &results, nil
	c.Data["json"] = &results //map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
	c.ServeJSON()
}

func buildQuery(query string, after ...string) io.Reader {
	var b strings.Builder

	b.WriteString("{\n")

	if query == "" {
		b.WriteString(searchAll)
	} else {
		b.WriteString(fmt.Sprintf(searchMatch, query))
	}

	if len(after) > 0 && after[0] != "" && after[0] != "null" {
		b.WriteString(",\n")
		b.WriteString(fmt.Sprintf(`	"search_after": %s`, after))
	}

	b.WriteString("\n}")

	// fmt.Printf("%s\n", b.String())
	return strings.NewReader(b.String())
}

const searchAll = `
	"query" : {
		"match_all" : {

		}
	},
	"size" : 10,
	"sort" : { "published" : "desc", "_doc" : "asc" }`

const searchMatch = `
	"query" : {
		"multi_match" : {
			"query" : %q,
			"fields" : ["title^100", "body^100","alt^10", "transcript"],
			"operator" : "and"
		}
	},
	"highlight" : {
		"fields" : {
			"title" : { "number_of_fragments" : 5 },
			"body" : { "number_of_fragments" : 5, "fragment_size" : 60 },
			"alt" : { "number_of_fragments" : 5 },
			"transcript" : { "number_of_fragments" : 5, "fragment_size" : 25 }
		},
		"max_analyzed_offset":900000
	},
	"size" : 10,
	"sort" : [ { "_score" : "desc" }, { "_doc" : "asc" } ]`

// @Title post tika
// @Description post tika
// @Failure 400 Invalid page supplied
// @Failure 404 Tika not found
// @router /tika [post]
// 上传文档——tika解析——构造document——存入es
func (c *ElasticController) Tika() {
	// Optionally pass a port as the second argument.
	f, err := os.Open("./test.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(f.Name())
	client := tika.NewClient(nil, "http://localhost:9998")
	body, err := client.Parse(context.Background(), f)
	// body, err := client.Detect(context.Background(), f) //application/pdf
	// fmt.Println(err)
	// fmt.Println(body)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != " " || selection.Text() != "\n" {
			fmt.Println(selection.Text())
		}
	})

	// 解析出来的文字生成文档，准备存入es
	//defer res.Body.Close()
	//var doc xkcdsearch.Document
	//type jsonResponse struct {
	//	Num        int `json:"num"`
	//	Year       string
	//	Month      string
	//	Day        string
	//	Title      string `json:"title"`
	//	Transcript string
	//	Alt        string
	//	Link       string
	//	News       string
	//	Img        string
	//}
	////if res.StatusCode != 200 {
	////	return doc, fmt.Errorf("[%s]", res.Status)
	////}
	//
	//var j jsonResponse
	//if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
	//	return doc, err
	//}
	//
	//jYear, err := strconv.ParseInt(j.Year, 10, 16)
	//if err != nil {
	//	return doc, fmt.Errorf("strconv: %s", err)
	//}
	//jMonth, err := strconv.ParseInt(j.Month, 10, 8)
	//if err != nil {
	//	return doc, fmt.Errorf("strconv: %s", err)
	//}
	//jDay, err := strconv.ParseInt(j.Day, 10, 8)
	//if err != nil {
	//	return doc, fmt.Errorf("strconv: %s", err)
	//}
	//
	//doc = xkcdsearch.Document{
	//	ID:         strconv.Itoa(j.Num),
	//	ImageURL:   j.Img,
	//	Published:  fmt.Sprintf("%04d-%02d-%02d", jYear, jMonth, jDay),
	//	Title:      j.Title,
	//	Alt:        j.Alt,
	//	Transcript: j.Transcript,
	//	Link:       j.Link,
	//	News:       j.News,
	//}
}

// 官方用法
// https://blog.csdn.net/weixin_52025712/article/details/126253327
func esdefaultconfig() {
	es, _ := elasticsearch.NewDefaultClient()
	res, _ := es.Info()
	defer res.Body.Close()
	log.Println(res)
}

func esapicustomconfig() {
	// 自定义配置
	cfg := elasticsearch.Config{
		// 有多个节点时需要配置
		Addresses: []string{
			"http://localhost:9200",
		},
		// 配置HTTP传输对象
		Transport: &http.Transport{
			//MaxIdleConnsPerHost 如果非零，控制每个主机保持的最大空闲(keep-alive)连接。如果为零，则使用默认配置2。
			MaxIdleConnsPerHost: 10,
			//ResponseHeaderTimeout 如果非零，则指定在写完请求(包括请求体，如果有)后等待服务器响应头的时间。
			ResponseHeaderTimeout: time.Second,
			//DialContext 指定拨号功能，用于创建不加密的TCP连接。如果DialContext为nil(下面已弃用的Dial也为nil)，那么传输拨号使用包网络。
			DialContext: (&net.Dialer{Timeout: time.Second}).DialContext,
			// TLSClientConfig指定TLS.client使用的TLS配置。
			//如果为空，则使用默认配置。
			//如果非nil，默认情况下可能不启用HTTP/2支持。
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS11,
				//InsecureSkipVerify 控制客户端是否验证服务器的证书链和主机名。
				InsecureSkipVerify: true,
			},
		},
	}
	es, _ := elasticsearch.NewClient(cfg)
	res, _ := es.Info()
	defer res.Body.Close()
	log.Println(res)
}

// CRUD
// 新增文档
// 使用 index api对文档进行增添或是修改操作。如果id不存在为创建文档，如果文档存在则进行修改。
func esindex() {
	es, _ := elasticsearch.NewDefaultClient()
	// 在索引中创建或更新文档。
	res, err := es.Index(
		"test",                                  // Index name
		strings.NewReader(`{"title" : "Test"}`), // Document body
		es.Index.WithDocumentID("1"),            // Document ID
		//es.Index.WithRefresh("true"),               // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

// 也可以使用esapi进行请求的包装，然后使用Do()方法执行请求。我们做同上面一样的操作，如下
func esapiindexrequest() {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.IndexRequest{
		Index:      "test",                                  // Index name
		Body:       strings.NewReader(`{"title" : "Test"}`), // Document body
		DocumentID: "1",                                     // Document ID
		Refresh:    "true",                                  // Refresh
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

// 下面都将之使用esapi方法实现。

// 不覆盖的创建文档
// 如果不想因为在创建文档填写错了id而对不想进行操作的文档进行了修改，那么可以使用CreateRequest包装请求。
func esapicreaterequest() {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.CreateRequest{
		Index: "learn",
		// DocumentType: "user",
		DocumentID: "1",
		Body: strings.NewReader(`
{
   "name": "张三",
   "age": 25,
   "about": "一个热爱刑法的男人，但是不精通唱跳Rap"
}`),
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Println("出错了，这个你麻麻滴错误是", err)
	}
	log.Println(res)
}

// 查询文档
// 查询单个文档
// 使用GetRequest包装请求。
func esapigetrequest() {
	es, _ := elasticsearch.NewDefaultClient()

	req := esapi.GetRequest{
		Index: "learn",
		// DocumentType: "user",
		DocumentID: "1",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

// 查询多个文档
// 使用MgetRequest包装请求。
func esapiMgetrequest() {
	es, _ := elasticsearch.NewDefaultClient()
	request := esapi.MgetRequest{
		Index: "learn",
		// DocumentType: "user",
		Body: strings.NewReader(`{
  "docs": [
    {
      "_id": "1"
    },
    {
      "_id": "2"
    }
  ]
}`),
	}
	res, err := request.Do(context.Background(), es)
	if err != nil {
		log.Println("出错了，错误是", err)
	}
	log.Println(res)
}

// 修改文档
// 在上面我们已经进行了创建或者修改的操作，但是使用 index api进行的修改操作需要提供所有的字段，不然会返回 400。
// 但我们大多数时候只是进行单个字段或多个字段的修改，并不会修改整个文档，这时候我们可以使用UpdateRequest包装请求。
func esapiupdaterequest() {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.UpdateRequest{
		Index: "learn",
		// DocumentType: "user",
		DocumentID: "1",
		Body: strings.NewReader(`
{
   "doc": {
   "name": "张三"
   }
}`),
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Println("出错了，这个你麻麻滴错误是", err)
	}
	log.Println(res)
}

// 删除文档
// 使用DeleteRequest包装请求。
func esapideleterequest() {
	// 创建一个默认配置的客户端
	es, _ := elasticsearch.NewDefaultClient()

	// 使用index请求
	req := esapi.DeleteRequest{
		Index: "test",
		// DocumentType: "_doc",
		DocumentID: "1",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

// 批量操作
// 使用BulkRequest包装请求。注意：格式一定要按照bulk api的格式来写，不然会400，最后别忘了回车
func esapibulkrequest() {
	es, _ := elasticsearch.NewDefaultClient()

	// 使用index请求
	req := esapi.BulkRequest{
		// 在body中写入bulk请求
		Body: strings.NewReader(`{ "index" : { "_index" : "test", "_id" : "1" } }
{ "title" : "Test2" }
{ "delete" : { "_index" : "test", "_id" : "2" } }
{ "create" : { "_index" : "test", "_id" : "3" } }
{ "field1" : "value3" }
{ "update" : {"_id" : "1", "_index" : "test"} }
{ "doc" : {"field2" : "value2"} }
`),
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

// 搜索
// 使用SearchRequest对请求进行包装。
func esapisearch() {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.SearchRequest{
		Index: []string{"learn"},
		// DocumentType: []string{"user"},
		Body: strings.NewReader(`{
  "query": {
    "match": {
      "about": "唱跳"
    }
  }
}`),
	}
	response, err := req.Do(context.Background(), es)
	if err != nil {
		log.Println("出错了，这个你麻麻滴错误是", err)
	}
	log.Println(response)
}

// 一下作废，非官方用法
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Index() {
	// es, err := elasticsearch.NewDefaultClient()
	// if err != nil {
	// 	log.Fatalf("Error creating the client: %s", err)
	// }

	// res, err := es.Info()
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }

	// defer res.Body.Close()
	// log.Println(res)

	// cfg := elasticsearch.Config{
	// 	Addresses: []string{
	// 		"http://localhost:9200",
	// 		"http://localhost:9201",
	// 	},
	// 	// ...
	// }
	// es, err = elasticsearch.NewClient(cfg)

	// cfg = elasticsearch.Config{
	// 	// ...
	// 	Username: "foo",
	// 	Password: "bar",
	// }

	// Initialize a client with the default settings.
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}

	// 1. Get cluster info
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}

//func Index2() {
//	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
//	config := elasticsearch.Config{
//		Addresses: addresses,
//	}
//	// new client
//	es, err := elasticsearch.NewClient(config)
//	failOnError(err, "Error creating the client")
//	// Index creates or updates a document in an index
//	var buf bytes.Buffer
//	doc := map[string]interface{}{
//		"title":   "你看到外面的世界是什么样的？",
//		"content": "外面的世界真的很精彩",
//	}
//	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
//		failOnError(err, "Error encoding doc")
//	}
//	res, err := es.Index("demo", &buf, es.Index.WithDocumentType("doc"))
//	if err != nil {
//		failOnError(err, "Error Index response")
//	}
//	defer res.Body.Close()
//	fmt.Println(res.String())
//}

func Search() {
	//addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	//config := elasticsearch.Config{
	//	Addresses: addresses,
	//	// Username:  "",
	//	// Password:  "",
	//	// CloudID:   "",
	//	// APIKey:    "",
	//}
	//// new client
	//es, err := elasticsearch.NewClient(config)
	//failOnError(err, "Error creating the client")
	// info
	res, err := es.Info()
	failOnError(err, "Error getting response")
	fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "中国",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='red'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("demo"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func DeleteByQuery() {
	//addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	//config := elasticsearch.Config{
	//	Addresses: addresses,
	//	// Username:  "",
	//	// Password:  "",
	//	// CloudID:   "",
	//	// APIKey:    "",
	//}
	//// new client
	//es, err := elasticsearch.NewClient(config)
	//failOnError(err, "Error creating the client")
	// DeleteByQuery deletes documents matching the provided query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}
	index := []string{"demo"}
	res, err := es.DeleteByQuery(index, &buf)
	if err != nil {
		failOnError(err, "Error delete by query response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func Delete() {
	//addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	//config := elasticsearch.Config{
	//	Addresses: addresses,
	//	// Username:  "",
	//	// Password:  "",
	//	// CloudID:   "",
	//	// APIKey:    "",
	//}
	//// new client
	//es, err := elasticsearch.NewClient(config)
	//failOnError(err, "Error creating the client")
	// Delete removes a document from the index
	res, err := es.Delete("demo", "POcKSHIBX-ZyL96-ywQO")
	if err != nil {
		failOnError(err, "Error delete by id response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

//func CreateEs() {
//	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
//	config := elasticsearch.Config{
//		Addresses: addresses,
//		// Username:  "",
//		// Password:  "",
//		// CloudID:   "",
//		// APIKey:    "",
//	}
//	// new client
//	es, err := elasticsearch.NewClient(config)
//	failOnError(err, "Error creating the client")
//	// Create creates a new document in the index.
//	// Returns a 409 response when a document with a same ID already exists in the index.
//	var buf bytes.Buffer
//	doc := map[string]interface{}{
//		"title":   "你看到外面的世界是什么样的？",
//		"content": "外面的世界真的很精彩",
//	}
//	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
//		failOnError(err, "Error encoding doc")
//	}
//	res, err := es.Create("demo", "esd", &buf, es.Create.WithDocumentType("doc"))
//	if err != nil {
//		failOnError(err, "Error create response")
//	}
//	defer res.Body.Close()
//	fmt.Println(res.String())
//}

func Get() {
	//addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	//config := elasticsearch.Config{
	//	Addresses: addresses,
	//	// Username:  "",
	//	// Password:  "",
	//	// CloudID:   "",
	//	// APIKey:    "",
	//}
	//// new client
	//es, err := elasticsearch.NewClient(config)
	//failOnError(err, "Error creating the client")
	res, err := es.Get("demo", "esd")
	if err != nil {
		failOnError(err, "Error get response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

//func Update() {
//	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
//	config := elasticsearch.Config{
//		Addresses: addresses,
//		// Username:  "",
//		// Password:  "",
//		// CloudID:   "",
//		// APIKey:    "",
//	}
//	// new client
//	es, err := elasticsearch.NewClient(config)
//	failOnError(err, "Error creating the client")
//	// Update updates a document with a script or partial document.
//	var buf bytes.Buffer
//	doc := map[string]interface{}{
//		"doc": map[string]interface{}{
//			"title":   "更新你看到外面的世界是什么样的？",
//			"content": "更新外面的世界真的很精彩",
//		},
//	}
//	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
//		failOnError(err, "Error encoding doc")
//	}
//	res, err := es.Update("demo", "esd", &buf, es.Update.WithDocumentType("doc"))
//	if err != nil {
//		failOnError(err, "Error Update response")
//	}
//	defer res.Body.Close()
//	fmt.Println(res.String())
//}

func UpdateByQuery() {
	//addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	//config := elasticsearch.Config{
	//	Addresses: addresses,
	//	// Username:  "",
	//	// Password:  "",
	//	// CloudID:   "",
	//	// APIKey:    "",
	//}
	//// new client
	//es, err := elasticsearch.NewClient(config)
	//failOnError(err, "Error creating the client")
	// UpdateByQuery performs an update on every document in the index without changing the source,
	// for example to pick up a mapping change.
	index := []string{"demo"}
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
		// 根据搜索条件更新title
		/*
			"script": map[string]interface{}{
				"source": "ctx._source['title']='更新你看到外面的世界是什么样的？'",
			},
		*/
		// 根据搜索条件更新title、content
		/*
			"script": map[string]interface{}{
				"source": "ctx._source=params",
				"params": map[string]interface{}{
					"title": "外面的世界真的很精彩",
					"content": "你看到外面的世界是什么样的？",
				},
				"lang": "painless",
			},
		*/
		// 根据搜索条件更新title、content
		"script": map[string]interface{}{
			"source": "ctx._source.title=params.title;ctx._source.content=params.content;",
			"params": map[string]interface{}{
				"title":   "看看外面的世界真的很精彩",
				"content": "他们和你看到外面的世界是什么样的？",
			},
			"lang": "painless",
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.UpdateByQuery(
		index,
		es.UpdateByQuery.WithAnalyzer("doc"),
		es.UpdateByQuery.WithBody(&buf),
		es.UpdateByQuery.WithContext(context.Background()),
		es.UpdateByQuery.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error Update response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// https://www.cnblogs.com/Me1onRind/p/11534544.html
// 删除索引
func deleteIndex() {
	req := esapi.IndicesDeleteRequest{
		Index: []string{"test_index"},
	}
	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

// 往索引插入数据
// 插入单条数据
func insertSingle() {
	body := map[string]interface{}{
		"num": 0,
		"v":   0,
		"str": "test",
	}
	jsonBody, _ := json.Marshal(body)

	req := esapi.CreateRequest{ // 如果是esapi.IndexRequest则是插入/替换
		Index: "test_index",
		//DocumentType: "test_type",
		DocumentID: "test_1",
		Body:       bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

//[201 Created] {"_index":"test_index","_type":"test_type","_id":"test_1","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}

// 批量插入(很明显，也可以批量做其他操作)
func insertBatch() {
	var bodyBuf bytes.Buffer
	for i := 2; i < 10; i++ {
		createLine := map[string]interface{}{
			"create": map[string]interface{}{
				"_index": "test_index",
				"_id":    "test_" + strconv.Itoa(i),
				"_type":  "test_type",
			},
		}
		jsonStr, _ := json.Marshal(createLine)
		bodyBuf.Write(jsonStr)
		bodyBuf.WriteByte('\n')

		body := map[string]interface{}{
			"num": i % 3,
			"v":   i,
			"str": "test" + strconv.Itoa(i),
		}
		jsonStr, _ = json.Marshal(body)
		bodyBuf.Write(jsonStr)
		bodyBuf.WriteByte('\n')
	}

	req := esapi.BulkRequest{
		Body: &bodyBuf,
	}
	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

// 查询
// 通过sql查询
//
//	func selectBySql() {
//		query := map[string]interface{}{
//			"query": "select count(*) as cnt, max(v) as value, num from test_index where num > 0 group by num",
//		}
//		jsonBody, _ := json.Marshal(query)
//		req := esapi.XPackSQLQueryRequest{
//			Body: bytes.NewReader(jsonBody),
//		}
//		res, err := req.Do(context.Background(), es)
//		checkError(err)
//		defer res.Body.Close()
//		fmt.Println(res.String())
//	}
//
// 通过Search Api查询
func selectBySearch() {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"num": map[string]interface{}{
							"gt": 0,
						},
					},
				},
			},
		},
		"size": 0,
		"aggs": map[string]interface{}{
			"num": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "num",
					//"size":  1,
				},
				"aggs": map[string]interface{}{
					"max_v": map[string]interface{}{
						"max": map[string]interface{}{
							"field": "v",
						},
					},
				},
			},
		},
	}
	jsonBody, _ := json.Marshal(query)
	req := esapi.SearchRequest{
		Index: []string{"test_index"},
		//DocumentType: []string{"test_type"},
		Body: bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

//但是elasticsearch对聚合查询分页并不是很友好，基本上都是得自己手动分页。

// 局部更新(批量更新略)
// 根据id更新
func updateSingle() {
	body := map[string]interface{}{
		"doc": map[string]interface{}{
			"v": 100,
		},
	}
	jsonBody, _ := json.Marshal(body)
	req := esapi.UpdateRequest{
		Index: "test_index",
		//DocumentType: "test_type",
		DocumentID: "test_1",
		Body:       bytes.NewReader(jsonBody),
	}

	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

//[200 OK] {"_index":"test_index","_type":"test_type","_id":"test_1","_version":2,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":3,"_primary_term":1}

//除了doc方式之外，还有script方式

// 根据条件更新
func updateByQuery() {
	body := map[string]interface{}{
		"script": map[string]interface{}{
			"lang": "painless",
			"source": `
                ctx._source.v = params.value;
            `,
			"params": map[string]interface{}{
				"value": 101,
			},
		},
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	jsonBody, _ := json.Marshal(body)
	req := esapi.UpdateByQueryRequest{
		Index: []string{"test_index"},
		Body:  bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

//[200 OK] {"took":109,"timed_out":false,"total":9,"updated":9,"deleted":0,"batches":1,"version_conflicts":0,"noops":0,"retries":{"bulk":0,"search":0},"throttled_millis":0,"requests_per_second":-1.0,"throttled_until_millis":0,"failures":[]}

// 删除
// 根据id删除
func deleteSingle() {
	req := esapi.DeleteRequest{
		Index: "test_index",
		//DocumentType: "test_type",
		DocumentID: "test_1",
	}

	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

//[200 OK] {"_index":"test_index","_type":"test_type","_id":"test_1","_version":6,"result":"deleted","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":7,"_primary_term":1}

// 根据条件删除
func deleteByQuery() {
	body := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	jsonBody, _ := json.Marshal(body)
	req := esapi.DeleteByQueryRequest{
		Index: []string{"test_index"},
		Body:  bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), es)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

// [200 OK] {"took":17,"timed_out":false,"total":9,"deleted":9,"batches":1,"version_conflicts":0,"noops":0,"retries":{"bulk":0,"search":0},"throttled_millis":0,"requests_per_second":-1.0,"throttled_until_millis":0,"failures":[]}
func PathisExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
