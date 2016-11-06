export CSV_DATA_INPUT=./tmp/src.csv
export JSON_DATA_OUTPUT=./tmp/output.json

output: gorun
	test ! -r ${JSON_DATA_OUTPUT} || cat ${JSON_DATA_OUTPUT} | python -m json.tool

gorun:
	test ! -r ${JSON_DATA_OUTPUT} || rm -rf ${JSON_DATA_OUTPUT}
	go run main.go

