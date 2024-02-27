humanround is a simple Go package that tries to round human-scale numbers to more "human" values. For example, rather than rounding 50.8 to 51, it rounds it to 50. The precision varies by the size of the number, and unit idiosyncracies are accounted for (ie inches get rounded to halves/quarters/eighths, depending on their size).

See https://pkg.go.dev/github.com/reillywatson/humanround for documentation.

