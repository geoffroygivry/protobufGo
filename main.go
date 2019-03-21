package main

import (
    fmt "fmt"
    "./protobuf"
    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
    "log"
    "io/ioutil"
)

func main() {
    sm := doSimple()
    readAndWrite(sm)
}

func toJSON(pb proto.Message) string {
    marshaler := jsonpb.Marshaler{}
    out, err := marshaler.MarshalToString(pb)
    if (err != nil) {
        log.Fatalln("We can't put this to JSON!")
        return ""
    }
    output := "This is the toJSON function\n"+ out
    return output
}

func readAndWrite(sm proto.Message) {
//     writeToFile("data.bin", sm)
//     sm2 := &simplepb.SimpleMessage{}
//     readFromFile("data.bin", sm2)
//     fmt.Println(sm2)
    sm3 := toJSON(sm)
    fmt.Println(sm3)
}

func readFromFile(fname string, pb proto.Message) error {
    in, err := ioutil.ReadFile(fname)
    if (err != nil) {
        log.Fatalln("Can't read the file...", err)
        return err
    }
    err2 := proto.Unmarshal(in, pb)
    if (err2 != nil) {
        log.Fatalln("couldn't read it, sorry.", err2)
        return err2
    }
    return nil
}

func writeToFile(fname string, pb proto.Message) error{
    out, err := proto.Marshal(pb)
    if (err != nil) {
        log.Fatalln("sorry can't Write to file...", err)
        return err
    }
    if err := ioutil.WriteFile(fname, out, 0644); err != nil {
        log.Fatalln("Sorry can't write again to file!", err)
        return err
    }
    fmt.Println("Data have been written! yay!")
    return nil
}

func doSimple() *simplepb.SimpleMessage {
    sm := simplepb.SimpleMessage{
        Id : 1975,
        IsSimple : true,
        Name : "My First ProtoBuf in Go!",
        SampleList : []int32 {1,2,3,4,5,6},
    }
//     fmt.Println(sm)
    
//     sm.Name = "I renamed You!"
//     fmt.Println(sm)
    
//     fmt.Println("The Id of this message is: ", sm.GetId())
    
    return &sm
}





