[![GoDoc](https://godoc.org/github.com/cjtoolkit/validate?status.svg)](https://godoc.org/github.com/cjtoolkit/validate)
[![Build Status](https://travis-ci.org/cjtoolkit/validate.svg?branch=master)](https://travis-ci.org/cjtoolkit/validate)

# CJToolkit Validate

A simple functional data validation system

## Installation

``` sh
$ go get github.com/cjtoolkit/validate
```

## Example

* https://github.com/CJ-Jackson/scratchpad/blob/master/example/validation/main.go

## Benchmark

* https://github.com/CJ-Jackson/scratchpad/tree/master/benchmark/validation

## Documentation

* https://godoc.org/github.com/cjtoolkit/validate/vBool
* https://godoc.org/github.com/cjtoolkit/validate/vError
* https://godoc.org/github.com/cjtoolkit/validate/vFile
* https://godoc.org/github.com/cjtoolkit/validate/vFloat
* https://godoc.org/github.com/cjtoolkit/validate/vInt
* https://godoc.org/github.com/cjtoolkit/validate/vString
* https://godoc.org/github.com/cjtoolkit/validate/vTime
* https://godoc.org/github.com/cjtoolkit/validate/vUint

## FAQ

### Why did I build a validation system?

I needed a validation system that I can easily work on with my IDE of choice, which happens to be
[GoLand](https://www.jetbrains.com/go/), while working with the system, all I have to remember is
those following vocabularies.

* vBool
* vError
* vFile
* vFloat
* vInt
* vString
* vTime
* vUint

Than GoLand auto-complete just click in quite nicely.  Anytime I want to validate, all I have to
think of is V for Vendatta, I couldn't resist.  All the validation rules are closures/lambda,
therefore the system is very easy to extend on and more importantly easier to test; also all the
rules are nicely documented in alphabetical order courtesy of godoc.org. (See documentation above)

I also did look at other alternative, such as https://github.com/thedevsaddam/govalidator and
https://github.com/gorilla/schema , with those alternative you to have specify the validation rules
as string rather than a simple function call, I would actually find that rather awkward, I would rather
have the CPU being able to assist me with writing validation rules, with simple string it's not really
possible without the use of IDE plugin, but I prefer to keep the IDE vanilla as possible!

I also benchmark those two and compared them to mine and here are the
[benchmark](https://github.com/CJ-Jackson/scratchpad/tree/master/benchmark/validation), those numbers
 do not lie! :D 
 
Also it's nice to have pinpoint accuracy with error message and knowing where to put it onto
the html template, have a close look at the template in the
[example](https://github.com/CJ-Jackson/scratchpad/blob/master/example/validation/main.go), pretty
cool isn't it?  Multi-value return is actually quite nice when used properly.

### Is it going to have a translation system?

Sorry no, I don't want to make the same mistake as I did with 
[form](https://github.com/cjtoolkit/form), which I have archived and it's was a nightmare to maintain.
However I did do something to make it easier for you to integrate this validation system into a
translation system of your choice (a map or something), I kept the constant variable and
[error data type](https://godoc.org/github.com/cjtoolkit/validate/vError#Errors)
visible; for example you can use the constant as key values for your map and have the translations as
values.

### What about automated html form rendering?

Sorry no, I'm not a fan of automated html form, I think it's will be better for you to build your own
html form and pick your own css framework of your choice, you will get more flexibility.  Because of
the number of css framework it's going to be a nightmare for me to build such a system;  it's just not
worth the effort.

Also I like to keep the system generic and simple.

### Why did you only cover int64, float64 and uint64 and not anything smaller in terms of bits?

I like to keep the system very easy to maintain, also it's not rocket science to type case int8 to
int64 and back, it's very easy.  If I did the entire thing, it's would be a nightmare to maintain.
Just deal with it.
