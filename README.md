# CJToolkit Validate

A simple functional data validation system

## Example

* https://github.com/CJ-Jackson/scratchpad/blob/master/example/validation/main.go

## Benchmark

* https://github.com/CJ-Jackson/scratchpad/tree/master/benchmark/validation

## Documentation

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

* vError
* vFile
* vFloat
* vInt
* vString
* vTime
* vUnit

Than GoLand auto-complete just click in quite nicely.  Anytime I want to validate, all I have to
think of is V for Vendatta, I couldn't resist.  All the validation rules are closures/lambda,
therefore the system is very easy to extend on and more importantly easier to test; also all the
rules are nicely documented in alphabetical order courtesy of godoc.org. (See documentation above)

I also did look at other alternative, such at https://github.com/thedevsaddam/govalidator and
https://github.com/gorilla/schema , with those alternative you to have specify the validation rules
as string rather than call, I would actually find that rather awkward, I would rather have the CPU
being able to assist me with writing validation rules, with simple string it's not really possible
without the use of IDE plugin, but I prefer to keep the IDE vanilla as possible!

I also benchmark those two and compared them to mime and if you're up for a laugh have a look at the
[benchmark](https://github.com/CJ-Jackson/scratchpad/tree/master/benchmark/validation), those numbers
 do not lie! :D 