package mapreduce

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"sort"
)

// doReduce manages one reduce task: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	//
	// You will need to write this function.
	//
	// You'll need to read one intermediate file from each map task;
	// reduceName(jobName, m, reduceTaskNumber) yields the file
	// name from map task m.
	//
	// Your doMap() encoded the key/value pairs in the intermediate
	// files, so you will need to decode them. If you used JSON, you can
	// read and decode by creating a decoder and repeatedly calling
	// .Decode(&kv) on it until it returns an error.
	//
	// You may find the first example in the golang sort package
	// documentation useful.
	//
	// reduceF() is the application's reduce function. You should
	// call it once per distinct key, with a slice of all the values
	// for that key. reduceF() returns the reduced value for that key.
	//
	// You should write the reduce output as JSON encoded KeyValue
	// objects to the file named outFile. We require you to use JSON
	// because that is what the merger than combines the output
	// from all the reduce tasks expects. There is nothing special about
	// JSON -- it is just the marshalling format we chose to use. Your
	// output code will look something like this:
	//
	// enc := json.NewEncoder(file)
	// for key := ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//

	//zhenhuli code
	fmt.Printf("Reduce: jobName=%s, outFile=%s, reduceTaskNumber=%d, nMap=%d\n",jobName, outFile, reduceTaskNumber, nMap)

	keyvalues_map := make(map[string]([]string))

	for i := 0; i< nMap; i++ {
		fileName := reduceName(jobName, i, reduceTaskNumber)
		filePtr, error := os.Open(fileName)

		if (error != nil) {
			log.Fatal("read intermediate file error:",fileName)
		}

		defer filePtr.Close()

		decoder := json.NewDecoder(filePtr)
		var keyvalueTmp KeyValue

		for ; decoder.More() ; {
			err := decoder.Decode(&keyvalueTmp)
			if (err != nil) {
				log.Fatal("decode intermediate file failed")
			}
			keyvalues_map[keyvalueTmp.Key] = append(keyvalues_map[keyvalueTmp.Key], keyvalueTmp.Value)
		}


	}

	//sort the key in keyvalues_map
	keys := make([]string, 0, len(keyvalues_map))
	for k,_ := range keyvalues_map {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//write to file in json formate
	outFilePtr, err := os.Create(outFile)
	if (err != nil) {
		log.Fatal("write to file in json formate failed")
	}
	defer outFilePtr.Close()
	encoder := json.NewEncoder(outFilePtr)

	for _,k := range keys {
		encoder.Encode(KeyValue{k, reduceF(k, keyvalues_map[k])})

	}

	//zhenhuli code
}
