# cortis

> A project to manipulate an image and send it to a laser for printing.

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

For detailed explanation on how things work, consult the [docs for vue-loader](http://vuejs.github.io/vue-loader).

In order to quickly get started use the following command:
```
npm run build && go run *.go -statics dist -settings ./settings.yml
```

To build the go distribution for raspberry pi:
```
env GOOS=linux GOARCH=arm GOARM=7 go build .
```