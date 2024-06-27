### Gophercise \#4

HTML Link Parser: https://github.com/gophercises/link


#### Notes
- better to use io.Read interface to abstract behaviour since you are going to 
feed data from the web/os.File/etc
- think before of how the API is used - currently is a tee bit convoluted
- maybe separate the majority of methods in a file separately from main
- for more efficient string building we could have used a byte buffer or something
similar
