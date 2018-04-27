package cardutils

import (
    "regexp"
    "strings"
)

type polyCard struct{
    fname string
    lname string
    id string
}

func New() *polyCard {
   return &polyCard{}
}

func (pc *polyCard) Swipe(cardData string) error {
    // regex pattern to find last/first name
    // \\ = Go's string escape squence.
    nameReg, _ := regexp.Compile("\\^([a-zA-Z/ \\-]+)")
    name := nameReg.FindStringSubmatch(cardData)[1]

    // regex pattern to find student id num
    numReg, _ := regexp.Compile("(?:0{6,8})([0-9]+)")
    stuNum := numReg.FindStringSubmatch(cardData)[1]

    // split the name on the "/"
    nameSlice := strings.Split(name, "/")
    first := strings.Split(nameSlice[1], " ")[0]
    last := nameSlice[0]

    pc.fname = first
    pc.lname = last
    pc.id = stuNum
    return nil
}

func (pc polyCard) GetFirstName() string {
    return pc.fname
}

func (pc polyCard) GetLastName() string {
    return pc.lname
}

func (pc polyCard) GetFullName() string {
    return pc.fname + " " + pc.lname
}

func (pc polyCard) GetID() string {
    return pc.id
}

func (pc polyCard) GetEmail() string {
    // The registrar has requested that this not be implemented for now
    return "This function was not implemented"
}
