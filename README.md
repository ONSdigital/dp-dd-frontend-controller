# Data Discovery Front-End Controller

Handles requests from users for data discovery, calling the DD API to fetch metadata
(dimensions, variables, value domains, etc) about datasets a user is interested in, and
ultimately to submit and await data file generation jobs (i.e., CSV/XSL/JSON data files).
Coordinates with the front-end renderer to make it all pretty.

In other words, this is the "C" in a traditional [MVC](https://en.wikipedia.org/wiki/Model–view–controller)
front-end, where the model is provided by the DD API, and the view is implemented by the front-end
renderer. The twist is that each component of the triad is a separately deployed microservice.

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under the MIT license, see [LICENSE](LICENSE.md) for details.
