### Gophercise \#4

HTML Link Parser: https://github.com/gophercises/link


#### Notes
- better to use io.Read interface to abstract behaviour since you are going to 
feed data from the web/os.File/etc.
- think before of how the API is used. Currently is a tee bit convoluted.
- maybe separate the majority of methods in a file separately from main.
