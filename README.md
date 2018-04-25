# Go implementation of libcardutils (v0.1)

## Purpose
This software library allows extraction of first name, last name, 
and student ID number from ID cards used at Florida Polytechnic University. 
This can be freely used, as long as appropriate credit is given.
At this time, the University Registrar has requested that the email 
not be able to be generated by this library. If you modify the library to 
provide this functionality, that's on you.

## Importing it into your project
In your imports, add `github.com/libpoly/libcardutils-go` then run `go get`.

## Using it in your project
Using a USB magstripe reader, keep the card's data in a string.

For example, if your string's name is `cardData`, and you want to represent 
the card as an object named `card`, you would do this:

`card := cardutils.CreateCard(cardData)`

You will then be able to pull the data off of the card like this:

`card.GetFirstName()`

`card.GetLastName()`

`card.GetID()`

## License
Oh, yeah, and this is licensed under APACHE2 or something.

```
Copyright 2018 Greg Willard (r3pwn)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```