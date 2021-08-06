[ElasticSearch]
- docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.14.0
- curl localhost:9200/


[Use Case]
- Create Employee Index with attributes
  - Id,Name,Age,Designation
- Operations performed in this code.
  - create
  - select
  
[ElasticSearch Curls]
- curl localhost:9200/_cat/indices
- curl  localhost:9200/INDEX_NAME/_search?pretty=true&q="*"


[Execute Tests]
- go test