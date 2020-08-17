package es

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic"
)

var client *esClient

const (
	host = "http://127.0.0.1:9200"
)

type esClient struct {
	*elastic.Client
	ctx context.Context
}

func NewElasticSearchClient() error {
	errLog := log.New(os.Stdout, "Elastic", log.LstdFlags)

	esCli, err := elastic.NewClient(elastic.SetErrorLog(errLog), elastic.SetURL(host))
	if err != nil {
		return err
	}
	client = &esClient{Client: esCli, ctx: context.Background()}
	result, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		return err
	}
	log.Printf("Elasticsearch returned with code: %d and version: %s", code, result.Version.Number)

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		return err
	}
	log.Printf("Elasticsearch version :%s", version)

	return nil
}

func init() {
	err := NewElasticSearchClient()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// insert a document to the index
func (client *esClient) Insert(index string, value interface{}) (*elastic.IndexResponse, error) {
	// access by the http://localhost:9700/pibigstar/employee/id
	response, err := client.Index().
		Index(index).
		BodyJson(value).
		Do(client.ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// get the document by id
func (client *esClient) GetById(index string, id string) ([]byte, error) {
	// id 必须存在，不然会报错，如果想查找请用search
	result, err := client.Get().Index(index).Id(id).Do(client.ctx)
	if err != nil && !result.Found {
		return nil, err
	}
	if result.Found {
		bytes, _ := result.Source.MarshalJSON()
		return bytes, nil
	}
	return nil, nil
}

// search the result by query strings
func (client *esClient) Query(index, keyword string) (*elastic.SearchResult, error) {
	// 根据名字查询
	query := elastic.NewQueryStringQuery(keyword)
	result, err := client.Search().Index(index).Query(query).Do(client.ctx)
	return result, err
}

// Aggregate query
func (client *esClient) AggQuery(index, keyword string) (*elastic.SearchResult, error) {
	agg := elastic.NewDateHistogramAggregation().
		Field("@timestamp").
		TimeZone("Asia/Shanghai").
		MinDocCount(1).
		Interval("1m")

	// 查询一分钟前是否出现关键字keyword
	boolQuery := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery("@timestamp").
			Format("strict_date_optional_time").
			Gte(time.Now().Add(time.Minute * -1).Format(time.RFC3339)).
			Lte(time.Now().Format(time.RFC3339))).
		Filter(elastic.NewMultiMatchQuery(keyword).
			Type("best_fields").
			Lenient(true))

	result, err := client.Search().
		Index(index).
		Query(boolQuery).
		Timeout("30000ms").
		IgnoreUnavailable(true).
		Size(500).
		Aggregation("aggs", agg).
		Version(true).
		StoredFields("*").
		Do(client.ctx)

	return result, err
}

// delete the document by id
func (client *esClient) DeleteById(index, id string) (*elastic.DeleteResponse, error) {
	response, err := client.Delete().Index(index).Id(id).Do(client.ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// update the document by id
// values : map[string]interface{}{"age": 12}
func (client *esClient) UpdateById(index, id string, values map[string]interface{}) (*elastic.UpdateResponse, error) {
	response, err := client.Update().
		Index(index).
		Id(id).
		Doc(values).Do(client.ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}
