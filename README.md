# go-zeroweb
[![GoDoc](https://godoc.org/github.com/godofdream/go-zeroweb?status.png)](http://godoc.org/github.com/godofdream/go-zeroweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/godofdream/go-zeroweb)](https://goreportcard.com/report/github.com/godofdream/go-zeroweb)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8e13858e0e064c77902b082966520a60)](https://www.codacy.com/app/godofdream/go-zeroweb?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=godofdream/go-zeroweb&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/0df317e347fb9cc9747c/maintainability)](https://codeclimate.com/github/godofdream/go-zeroweb/maintainability)
[![Coverage](http://gocover.io/_badge/github.com/godofdream/go-zeroweb)](http://gocover.io/github.com/godofdream/go-zeroweb)
[![Build Status](https://travis-ci.org/godofdream/go-zeroweb.svg)](https://travis-ci.org/godofdream/go-zeroweb)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgodofdream%2Fgo-zeroweb.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgodofdream%2Fgo-zeroweb?ref=badge_shield)


WebFramework Stack written in golang optimized for lowest latency and high throughput

# Components
| Component     | Description   | License  |
| ------------- |:-------------:| -----:|
| Config        | [Viper](https://godoc.org/github.com/spf13/viper) can take configs from various formats and sources and supports reload triggers | [MIT License](https://github.com/spf13/viper/blob/master/LICENSE) |
| DB (Postgres) | [PGX](https://godoc.org/github.com/jackc/pgx) is a database driver for Postgres/Cockroachdb including a Connectionpool and prepared statements  | [MIT License](https://github.com/jackc/pgx/blob/master/LICENSE)   |
| Log           | [zerolog](https://godoc.org/github.com/rs/zerolog) is a static typed json logger with nearly no allocations | [MIT License](https://github.com/rs/zerolog/blob/master/LICENSE)    |
| Router        | [fasthttprouter](https://godoc.org/github.com/godofdream/fasthttprouter)      | [BSD 3-Clause "Revised" License](https://github.com/godofdream/fasthttprouter/blob/master/LICENSE)    |
| Server        | [fasthttp](https://godoc.org/github.com/godofdream/fasthttp) is a http library      | [MIT License](https://github.com/godofdream/fasthttp/blob/master/LICENSE)  |
| View          | [jet](https://godoc.org/github.com/godofdream/jet) is a zero allocation tempalte engine with jinja2 familiar syntax | [Apache License 2.0](https://github.com/godofdream/jet/blob/master/LICENSE)   |

# Why these component?
| Component     | Description   |
| ------------- |:-------------:|
| Config        | To provide a simple way for handling configurations and reloads on configchange Viper is directly included. Viper is not optimized for speed but for comfortability  |
| DB            | Postgres is mature, Cockroachdb is extremly scaleable but both speak the postgres protocol, this provides an common sql syntax with the possibility to scale  |
| Log           | Fastest Logger by [benchmark](https://github.com/rs/zerolog#benchmarks) |
| Router        | The Router is a Trie Based router with a Ο(\|s\|) Complexity with s being the length of the string |
| Server        | Fasthttp is the fastest http-server for golang [benchmark](https://www.techempower.com/benchmarks/) and uses an optimized flate and gzip implementation [compress](https://github.com/klauspost/compress)|
| View          | jinja2 familiar Syntax And fastest templating without precompilation to golang [benchmark](https://github.com/SlinSo/goTemplateBenchmark#full-featured-template-engines-1) |


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgodofdream%2Fgo-zeroweb.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgodofdream%2Fgo-zeroweb?ref=badge_large)