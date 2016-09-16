package main

import (
	"io/ioutil"
	"log"

	pb "github.com/dougreese/sdgopher-201609/vectorTile"
	"github.com/golang/protobuf/proto"
)

// printValue prints key/value of a tile layer.
func printValue(key string, val *pb.Tile_Value) {
	switch {
	case nil != val.StringValue:
		log.Printf("- %s: %s", key, *val.StringValue)
	case nil != val.FloatValue:
		log.Printf("- %s: %f", key, *val.FloatValue)
	case nil != val.DoubleValue:
		log.Printf("- %s: %f", key, *val.DoubleValue)
	case nil != val.IntValue:
		log.Printf("- %s: %d", key, *val.IntValue)
	case nil != val.UintValue:
		log.Printf("- %s: %d", key, *val.UintValue)
	case nil != val.SintValue:
		log.Printf("- %s: %d", key, *val.SintValue)
	default:
		log.Println("I'm confused!")
	}
}

func main() {
	fname := "2751.vector.pbf"

	log.Printf("Reading protobuf file: %s\n", fname)
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	tile := &pb.Tile{}
	if err := proto.Unmarshal(in, tile); err != nil {
		log.Fatalln("Failed to parse tile:", err)
	}

	layers := tile.Layers
	log.Println("Num layers", len(layers))

	for i := 0; i < len(layers); i++ {
		layer := layers[i]
		log.Println("Name:", layer.GetName())

		keys := layer.GetKeys()
		values := layer.GetValues()

		for j := 0; j < len(keys); j++ {
			val := values[j]
			printValue(keys[j], val)
		}
	}

	log.Println("Success!")
}
